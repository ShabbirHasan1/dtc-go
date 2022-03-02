package codegen

import (
	"fmt"
	"testing"
)

func TestParseAll(t *testing.T) {
	namespaces := NewNamespaces()
	err := namespaces.AddHeaders("testdata/DTCProtocol.h", "testdata/DTCProtocolVLS.h", "testdata/DTCProtocol_NonStandard.h")
	if err != nil {
		t.Fatal(err)
	}

	allConstants := make(map[string]*Const)
	allEnums := make(map[string]*Enum)

	for _, namespace := range namespaces.Namespaces {
		//fmt.Println("Namespace:", namespace.Name)
		//fmt.Println("\tConstants")
		for _, constant := range namespace.Constants {
			//fmt.Println("\t\t", constant.Name)
			if allConstants[constant.Name] != nil {
				dup := allConstants[constant.Name]
				if dup.File.Path != constant.File.Path {
					fmt.Println("duplicate constant named: "+constant.Name, " in file:", constant.File.Path, " and file:", allConstants[constant.Name].File.Path)
				}
			}
			allConstants[constant.Name] = constant
		}
		//fmt.Println("\tEnums")
		for _, enum := range namespace.Enums {
			//fmt.Println("\t\t", enum.Name)
			if allEnums[enum.Name] != nil {
				dup := allEnums[enum.Name]
				if enum.File.Path != dup.File.Path {
					fmt.Println("duplicate constant named: "+enum.Name, " in file:", enum.File.Path, " and file:", allEnums[enum.Name].File.Path)
				}
			}
			allEnums[enum.Name] = enum
		}
		//fmt.Println("\tStructs")
		for _, st := range namespace.Structs {
			//fmt.Println("\t\t", st.Name)
			_ = st
		}
	}

	fmt.Println(namespaces)
}

func TestParse(t *testing.T) {
	namespaces := NewNamespaces()
	file, err := namespaces.AddHeader("testdata/DTCProtocol.h")
	if err != nil {
		t.Fatal(err)
	}
	msg := file.StructsByName["s_SecurityDefinitionResponse"]
	fmt.Println(msg)

	msg = file.StructsByName["s_MarketOrdersAdd"]
	fmt.Println(msg)

	msg = file.StructsByName["s_HistoricalPriceDataRecordResponse"]
	fmt.Println(msg)
	fmt.Println(file)

	//for _, s := range file.Structs {
	//	fmt.Println(s.Name, " = ", s.Size)
	//}
	//for _, s := range file.Structs {
	//	fmt.Println("std::cout << \"" + s.Name + "\" << \" = \" << sizeof(DTC::" + s.Name + ") << std::endl;")
	//	for _, field := range s.Fields {
	//		if field.Type.Union != nil {
	//			for _, f := range field.Type.Union.Fields {
	//				fmt.Println("std::cout << \"\t" + f.Name + "\" << \" = \" << offsetof(DTC::" + s.Name + ", " + f.Name + ") << std::endl;")
	//			}
	//		} else {
	//			fmt.Println("std::cout << \"\t" + field.Name + "\" << \" = \" << offsetof(DTC::" + s.Name + ", " + field.Name + ") << std::endl;")
	//		}
	//	}
	//}

	//data, err := os.ReadFile("testdata/DTCLayout.txt")
	//if err != nil {
	//	t.Fatal(err)
	//}

	//fmt.Println("Mismatch Count: ", len(badSizes), " of ", len(file.Structs))
	//for _, s := range badSizes {
	//	fmt.Println(s.Struct.Name, " ", s.Struct.Size, " != ", s.CPPSize)
	//}
}
