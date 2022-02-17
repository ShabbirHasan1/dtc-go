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
	Name             string
	Offset           int
	CalculatedOffset int
}
type CPPStructLayout struct {
	Namespace      string
	Name           string
	Size           int
	CalculatedSize int
	Fields         []CPPFieldLayout
}

func (namespaces *Namespaces) PrintCPPLayoutErrors(data []byte) {
	layout := parseCPPLayout(string(data))

	type structCPPSize struct {
		Namespace *Namespace
		Struct    *Struct
		CPPSize   int
	}

	//goodSizes := make([]CPPStructLayout, 0, len(file.Structs))
	//badSizes := make([]CPPStructLayout, 0, len(file.Structs))
	for _, cppStruct := range layout.Structs {
		namespace := namespaces.Namespaces[cppStruct.Namespace]
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
			cppField.CalculatedOffset = f.Type.Offset

			if f.Type.Offset != cppField.Offset {
				fmt.Println("\t"+cppStruct.Namespace+"::"+cppStruct.Name+"."+cppField.Name, " ", f.Type.Offset, " != ", cppField.Offset)
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
			sz, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
			if err != nil {
				panic(err)
			}

			structLayout.Fields = append(structLayout.Fields, CPPFieldLayout{
				Name:   strings.TrimSpace(parts[0]),
				Offset: int(sz),
			})
		}
	}

	if len(structLayout.Name) > 0 {
		file.Structs = append(file.Structs, structLayout)
	}
	return file
}

func (namespaces *Namespaces) PrintCPPLayoutCode() {
	for _, namespace := range namespaces.Namespaces {
		for _, s := range namespace.Structs {
			fmt.Println("std::cout << \"" + s.Namespace.Name + "::" + s.Name + "\" << \" = \" << sizeof(" + s.Namespace.Name + "::" + s.Name + ") << std::endl;")
			for _, field := range s.Fields {
				if field.Type.Union != nil {
					for _, f := range field.Type.Union.Fields {
						fmt.Println("std::cout << \"\t" + f.Name + "\" << \" = \" << offsetof(" + s.Namespace.Name + "::" + s.Name + ", " + f.Name + ") << std::endl;")
					}
				} else {
					fmt.Println("std::cout << \"\t" + field.Name + "\" << \" = \" << offsetof(" + s.Namespace.Name + "::" + s.Name + ", " + field.Name + ") << std::endl;")
				}
			}
		}
	}
}