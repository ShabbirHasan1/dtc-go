package schema

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type cppLayout struct {
	Structs []cppStructLayout
}

type cppFieldLayout struct {
	Name   string
	Offset int
	Size   int
}
type cppStructLayout struct {
	Namespace string
	Name      string
	Size      int
	Fields    []cppFieldLayout
}

func cppNamespaceOf(namespace NamespaceKind) string {
	switch namespace {
	case NamespaceKindFixed:
		return "DTC"
	case NamespaceKindVLS:
		return "DTC_VLS"
	case NamespaceKindNonStandard:
		return "n_DTCNonStandard"
	default:
		return ""
	}
}

func printCPPLayoutErrors(schema *Schema, data []byte) error {
	layout := parseCPPLayout(string(data))

	for _, cppStruct := range layout.Structs {
		namespace := schema.GetNamespace(cppStruct.Namespace)
		if namespace == nil {
			panic("namespace: " + cppStruct.Namespace + "  was not found")
		}
		s := namespace.StructsByName[cppStruct.Name]
		if s == nil {
			panic(cppStruct.Name + " not found!!!")
		}
		if s.Size != cppStruct.Size {
			fmt.Println(cppNamespaceOf(s.Namespace)+"::"+s.Name, " ", s.Size, " != ", cppStruct.Size)
		}
		for _, cppField := range cppStruct.Fields {
			f := s.FieldsByName[cppField.Name]
			if f == nil {
				return errors.New(cppStruct.Namespace + "::" + cppStruct.Name + "." + cppField.Name + " not found!!!")
			}

			if f.Type.Offset != cppField.Offset {
				fmt.Println("\t"+cppStruct.Namespace+"::"+cppStruct.Name+"."+cppField.Name, " ", f.Type.Offset, " != ", cppField.Offset)
			}
			if f.Type.Size != cppField.Size {
				fmt.Println("\t"+cppStruct.Namespace+"::"+cppStruct.Name+"."+cppField.Name, " ", f.Type.Size, " != ", cppField.Size)
			}
		}
	}
	return nil
}

func parseCPPLayout(contents string) cppLayout {
	contents = strings.TrimSpace(contents)
	lines := strings.Split(contents, "\n")
	file := cppLayout{}
	structLayout := cppStructLayout{}

	for _, line := range lines {
		if !strings.HasPrefix(line, "\t") {
			if len(structLayout.Name) > 0 {
				file.Structs = append(file.Structs, structLayout)
			}

			parts := strings.Split(line, "=")
			sz, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
			if err != nil {
				panic(err)
			}

			name := strings.TrimSpace(parts[0])
			index := strings.LastIndex(name, "::")
			namespaceName := ""
			if index > -1 {
				namespaceName = strings.TrimSpace(name[0:index])
				name = strings.TrimSpace(name[index+2:])
			}

			structLayout = cppStructLayout{
				Namespace: namespaceName,
				Name:      name,
				Size:      int(sz),
			}
		} else {
			parts := strings.Split(line, "=")

			nums := strings.Split(strings.TrimSpace(parts[1]), ",")
			offset, err := strconv.ParseInt(strings.TrimSpace(nums[0]), 10, 64)
			if err != nil {
				panic(err)
			}
			sz, err := strconv.ParseInt(strings.TrimSpace(nums[1]), 10, 64)
			if err != nil {
				panic(err)
			}

			structLayout.Fields = append(structLayout.Fields, cppFieldLayout{
				Name:   strings.TrimSpace(parts[0]),
				Offset: int(offset),
				Size:   int(sz),
			})
		}
	}

	if len(structLayout.Name) > 0 {
		file.Structs = append(file.Structs, structLayout)
	}
	return file
}

func printCPPLayoutCode(schema *Schema) {
	namespaces := schema.GetNamespaces()
	for _, namespace := range namespaces {
		for _, s := range namespace.Structs {
			fmt.Println("std::cout << \"" + cppNamespaceOf(namespace.Kind) + "::" + s.Name + "\" << \" = \" << sizeof(" + cppNamespaceOf(s.Namespace) + "::" + s.Name + ") << std::endl;")
			for _, field := range s.Fields {
				if field.Type.Union != nil {
					for _, f := range field.Type.Union.Fields {
						fmt.Println("std::cout << \"\t" + f.Name + "\" << \" = \" << offsetof(" + cppNamespaceOf(s.Namespace) + "::" + s.Name + ", " + f.Name + ") << \",\" << sizeof(" + cppNamespaceOf(s.Namespace) + "::" + s.Name + "::" + f.Name + ") << std::endl;")
					}
				} else {
					fmt.Println("std::cout << \"\t" + field.Name + "\" << \" = \" << offsetof(" + cppNamespaceOf(s.Namespace) + "::" + s.Name + ", " + field.Name + ") << \",\" << sizeof(" + cppNamespaceOf(s.Namespace) + "::" + s.Name + "::" + field.Name + ") << std::endl;")
				}
			}
		}
	}
}
