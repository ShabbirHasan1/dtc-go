package codegen

import (
	"fmt"
	"strconv"
	"strings"
)

type CPPLayout struct {
	Structs []CPPStructLayout
}

type CPPFieldLayout struct {
	Name   string
	Offset int
	Size   int
}
type CPPStructLayout struct {
	Namespace string
	Name      string
	Size      int
	Fields    []CPPFieldLayout
}

func (schema *Schema) PrintCPPLayoutErrors(data []byte) {
	layout := parseCPPLayout(string(data))

	type structCPPSize struct {
		Namespace *Namespace
		Struct    *Struct
		CPPSize   int
	}

	//goodSizes := make([]CPPStructLayout, 0, len(file.Structs))
	//badSizes := make([]CPPStructLayout, 0, len(file.Structs))
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
			fmt.Println(s.Namespace.Name+"::"+s.Name, " ", s.Size, " != ", cppStruct.Size)
		}
		for _, cppField := range cppStruct.Fields {
			f := s.FieldsByName[cppField.Name]
			if f == nil {
				panic(cppStruct.Namespace + "::" + cppStruct.Name + "." + cppField.Name + " not found!!!")
			}

			if f.Type.Offset != cppField.Offset {
				fmt.Println("\t"+cppStruct.Namespace+"::"+cppStruct.Name+"."+cppField.Name, " ", f.Type.Offset, " != ", cppField.Offset)
			}
			if f.Type.Size != cppField.Size {
				fmt.Println("\t"+cppStruct.Namespace+"::"+cppStruct.Name+"."+cppField.Name, " ", f.Type.Size, " != ", cppField.Size)
			}
		}
	}
}

func parseCPPLayout(contents string) CPPLayout {
	contents = strings.TrimSpace(contents)
	lines := strings.Split(contents, "\n")
	file := CPPLayout{}
	structLayout := CPPStructLayout{}

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

			structLayout = CPPStructLayout{
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

			structLayout.Fields = append(structLayout.Fields, CPPFieldLayout{
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

func (schema *Schema) PrintCPPLayoutCode() {
	namespaces := schema.GetNamespaces()
	for _, namespace := range namespaces {
		for _, s := range namespace.Structs {
			fmt.Println("std::cout << \"" + s.Namespace.Name + "::" + s.Name + "\" << \" = \" << sizeof(" + s.Namespace.Name + "::" + s.Name + ") << std::endl;")
			for _, field := range s.Fields {
				if field.Type.Union != nil {
					for _, f := range field.Type.Union.Fields {
						fmt.Println("std::cout << \"\t" + f.Name + "\" << \" = \" << offsetof(" + s.Namespace.Name + "::" + s.Name + ", " + f.Name + ") << \",\" << sizeof(" + s.Namespace.Name + "::" + s.Name + "::" + f.Name + ") << std::endl;")
					}
				} else {
					fmt.Println("std::cout << \"\t" + field.Name + "\" << \" = \" << offsetof(" + s.Namespace.Name + "::" + s.Name + ", " + field.Name + ") << \",\" << sizeof(" + s.Namespace.Name + "::" + s.Name + "::" + field.Name + ") << std::endl;")
				}
			}
		}
	}
}
