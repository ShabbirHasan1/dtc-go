package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen"
	"reflect"
	"strings"
)

type Package struct {
	Name      string
	Path      string
	Namespace *codegen.Namespace
}

type Message struct {
	Fixed *Struct
	VLS   *Struct
}

type Enum struct {
	*codegen.Enum
}

type EnumOption struct {
	*codegen.EnumOption
}

type Struct struct {
	*codegen.Struct
}

type Field struct {
	*codegen.Field
}

type Type struct {
	Import *Import
	Name   string
	DTC    codegen.Type
	Kind   reflect.Kind
}

func (t *Type) String() string {
	if t.Import != nil {
		return fmt.Sprintf("%s.%s", t.Import.Prefix(), t.Name)
	}
	return t.Name
}

func (t *Type) IsVLS() bool {
	return t.DTC.Kind == codegen.KindStringVLS
}

type Import struct {
	Alias string
	Name  string
	Path  string
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

func TypeOf(typ codegen.Type) Type {
	switch typ.Kind {
	case codegen.KindUnknown:
	case codegen.KindInt8:
		return Type{
			Import: nil,
			Name:   "int8",
			DTC:    typ,
			Kind:   reflect.Int8,
		}
	case codegen.KindUint8:
		return Type{
			Import: nil,
			Name:   "uint8",
			DTC:    typ,
			Kind:   reflect.Uint8,
		}
	case codegen.KindInt16:
		return Type{
			Import: nil,
			Name:   "int16",
			DTC:    typ,
			Kind:   reflect.Int16,
		}
	case codegen.KindUint16:
		return Type{
			Import: nil,
			Name:   "uint16",
			DTC:    typ,
			Kind:   reflect.Uint16,
		}
	case codegen.KindInt32:
		return Type{
			Import: nil,
			Name:   "int32",
			DTC:    typ,
			Kind:   reflect.Int32,
		}
	case codegen.KindUint32:
		return Type{
			Import: nil,
			Name:   "uint32",
			DTC:    typ,
			Kind:   reflect.Uint32,
		}
	case codegen.KindInt64:
		return Type{
			Import: nil,
			Name:   "int64",
			DTC:    typ,
			Kind:   reflect.Int64,
		}
	case codegen.KindUint64:
		return Type{
			Import: nil,
			Name:   "uint64",
			DTC:    typ,
			Kind:   reflect.Uint64,
		}
	case codegen.KindFloat32:
		return Type{
			Import: nil,
			Name:   "float32",
			DTC:    typ,
			Kind:   reflect.Float32,
		}
	case codegen.KindFloat64:
		return Type{
			Import: nil,
			Name:   "float64",
			DTC:    typ,
			Kind:   reflect.Float64,
		}
	case codegen.KindStringFixed:
		return Type{
			Import: nil,
			Name:   "string",
			DTC:    typ,
			Kind:   reflect.String,
		}
	case codegen.KindStringVLS:
		return Type{
			Import: nil,
			Name:   "string",
			DTC:    typ,
			Kind:   reflect.String,
		}
	case codegen.KindAlias:
		underlying := TypeOf(typ.Alias.Type)
		return Type{
			Import: nil,
			Name:   cleanName(typ.Alias.Name),
			DTC:    typ,
			Kind:   underlying.Kind,
		}
	case codegen.KindEnum:
		underlying := primitiveKind(typ.Kind)
		return Type{
			Import: nil,
			Name:   cleanName(typ.Enum.Name),
			DTC:    typ,
			Kind:   underlying,
		}

	case codegen.KindUnion:
	case codegen.KindStruct:
	}
	return Type{}
}

func primitiveKind(k codegen.Kind) reflect.Kind {
	switch k {
	case codegen.KindInt8:
		return reflect.Int8
	case codegen.KindUint8:
		return reflect.Uint8
	case codegen.KindInt16:
		return reflect.Int16
	case codegen.KindUint16:
		return reflect.Uint16
	case codegen.KindInt32:
		return reflect.Int32
	case codegen.KindUint32:
		return reflect.Uint32
	case codegen.KindInt64:
		return reflect.Int64
	case codegen.KindUint64:
		return reflect.Uint64
	case codegen.KindFloat32:
		return reflect.Float32
	case codegen.KindFloat64:
		return reflect.Float64
	case codegen.KindStringFixed:
		return reflect.String
	case codegen.KindStringVLS:
		return reflect.String
	case codegen.KindStruct:
		return reflect.Struct
	default:
		return reflect.Invalid
	}
}

func cleanName(n string) string {
	index := strings.IndexByte(n, '_')
	if index > -1 {
		n = n[index+1:]
	}
	n = strings.TrimSpace(n)
	if strings.HasSuffix(n, "Enum") {
		n = n[0 : len(n)-4]
	}
	return n
}
