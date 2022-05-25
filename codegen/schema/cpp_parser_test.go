package schema

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	schema, err := LoadSchemaFromCHeaders(
		"../testdata/docs.json",
		"../testdata/DTCProtocol.proto",
		"../testdata/DTCProtocol.h",
		"../testdata/DTCProtocolVLS.h",
		"../testdata/DTCProtocol_NonStandard.h",
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(schema)
}

func cppPrint(file *CHeaderFile, namespace string) {
	// for _, s := range file.Structs {
	// fmt.Println(s.Name, " = ", s.Size)
	// }
	for _, s := range file.Structs {
		fmt.Println("std::cout << \"" + namespace + "::" + s.Name + "\" << \" = \" << sizeof(" + namespace + "::" + s.Name + ") << std::endl;")
		for _, field := range s.Fields {
			if field.Type.Union != nil {
				for _, f := range field.Type.Union.Fields {
					fmt.Println("std::cout << \"\t" + f.Name + "\" << \" = \" << offsetof(" + namespace + "::" + s.Name + ", " + f.Name + ") << std::endl;")
				}
			} else {
				fmt.Println("std::cout << \"\t" + field.Name + "\" << \" = \" << offsetof(" + namespace + "::" + s.Name + ", " + field.Name + ") << std::endl;")
			}
		}
	}
}

func TestParse(t *testing.T) {
	schema := NewSchema()
	file, err := schema.AddCHeader("../testdata/DTCProtocol.h")
	if err != nil {
		t.Fatal(err)
	}
	cppPrint(file, "DTC")
	// schema = NewSchema()
	file, err = schema.AddCHeader("../testdata/DTCProtocolVLS.h")
	if err != nil {
		t.Fatal(err)
	}
	cppPrint(file, "DTC_VLS")
	// msg := file.StructsByName["s_SecurityDefinitionResponse"]
	// fmt.Println(msg)

	// msg = file.StructsByName["s_MarketOrdersAdd"]
	// fmt.Println(msg)

	// msg = file.StructsByName["s_HistoricalPriceDataRecordResponse"]
	// fmt.Println(msg)
	// fmt.Println(file)

	// data, err := os.ReadFile("testdata/DTCLayout.txt")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.Println("Mismatch Count: ", len(badSizes), " of ", len(file.Structs))
	// for _, s := range badSizes {
	// 	fmt.Println(s.Struct.Name, " ", s.Struct.Size, " != ", s.CPPSize)
	// }
}
