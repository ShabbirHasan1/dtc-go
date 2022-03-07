package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen"
	"strings"
)

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
//type Kind uint
//
//const (
//	Invalid Kind = iota
//	Bool
//	Int
//	Int8
//	Int16
//	Int32
//	Int64
//	Uint8
//	Uint16
//	Uint32
//	Uint64
//	Float32
//	Float64
//	String
//)
//
//func (k Kind) String() string {
//	switch k {
//	case Int8:
//		return "int8"
//	case Int16:
//		return "int16"
//	case Int32:
//		return "int32"
//	case Int64:
//		return "int64"
//	case Uint8:
//		return "uint8"
//	case Uint16:
//		return "uint16"
//	case Uint32:
//		return "uint32"
//	case Uint64:
//		return "uint64"
//	case Float32:
//		return "float32"
//	case Float64:
//		return "float64"
//	case String:
//		return "string"
//	}
//	return "invalid"
//}

func primitiveKindName(kind codegen.Kind) string {
	switch kind {
	case codegen.KindInt8:
		return "int8"
	case codegen.KindInt16:
		return "int16"
	case codegen.KindInt32:
		return "int32"
	case codegen.KindInt64:
		return "int64"
	case codegen.KindUint8:
		return "uint8"
	case codegen.KindUint16:
		return "uint16"
	case codegen.KindUint32:
		return "uint32"
	case codegen.KindUint64:
		return "uint64"
	case codegen.KindFloat32:
		return "float32"
	case codegen.KindFloat64:
		return "float64"
	case codegen.KindStringFixed:
		return "string"
	case codegen.KindStringVLS:
		return "string"
	}
	return "invalid"
}

type Package struct {
	Name            string
	Path            string
	Enums           []*Enum
	EnumsByName     map[string]*Enum
	Constants       []*Constant
	ConstantsByName map[string]*Constant
	Aliases         []*Alias
	AliasesByMame   map[string]*Alias
	Structs         []*Struct
	StructsByName   map[string]*Struct
}

type Message struct {
	Fixed       *Struct
	VLS         *Struct
	NonStandard bool
}

type Alias struct {
	Package *Package
	*codegen.Alias
	Name string
}

type Constant struct {
	Package *Package
	*codegen.Const
	Name string
}

type Enum struct {
	Package *Package
	*codegen.Enum
	Name          string
	Options       []*EnumOption
	OptionsByName map[string]*EnumOption
}

type EnumOption struct {
	*codegen.EnumOption
	Enum *Enum
	Name string
}

type Struct struct {
	*codegen.Struct
	Name         string
	Fields       []*Field
	FieldsByName map[string]*Field
}

type Field struct {
	*codegen.Field
	Struct *Struct
	Name   string
	Fields []*Field
}

//type Type struct {
//	Package   *Package
//	Import    *Import
//	Name      string
//	DTC       codegen.Type
//	Primitive bool
//}
//
//func (t *Type) String() string {
//	if t == nil {
//		return ""
//	}
//	if t.Import != nil {
//		return fmt.Sprintf("%s%s", t.Import.Prefix(), t.Name)
//	}
//	return t.Name
//}
//
//func (t *Type) IsVLS() bool {
//	return t.DTC.Kind == codegen.KindStringVLS
//}

type Import struct {
	Package *Package
	Alias   string
	Name    string
	Path    string
}

func (imp *Import) Clause() string {
	if len(imp.Path) == 0 {
		return ""
	}
	if len(imp.Alias) == 0 {
		return fmt.Sprintf("\"%s\"", imp.Path)
	}
	return fmt.Sprintf("%s \"%s\"", imp.Alias, imp.Path)
}

func (imp *Import) Prefix() string {
	if len(imp.Alias) > 0 {
		return fmt.Sprintf("%s.", imp.Alias)
	}
	if len(imp.Name) > 0 {
		return fmt.Sprintf("%s.", imp.Name)
	}
	return ""
}

func cleanName(n string) string {
	if strings.HasPrefix(n, "s_") {
		return strings.TrimSpace(n[2:])
	}
	if strings.HasPrefix(n, "m_") {
		return strings.TrimSpace(n[2:])
	}
	if strings.HasPrefix(n, "t_") {
		return strings.TrimSpace(n[2:])
	}
	n = strings.TrimSpace(n)
	//if strings.HasSuffix(n, "Enum") {
	//	n = n[0 : len(n)-4]
	//}
	return n
}
