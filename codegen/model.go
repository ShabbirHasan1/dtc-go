package codegen

import "github.com/yoheimuta/go-protoparser/v4/parser"

type Kind byte

const (
	KindUnknown     Kind = 0
	KindInt8        Kind = 1
	KindUint8       Kind = 2
	KindInt16       Kind = 3
	KindUint16      Kind = 4
	KindInt32       Kind = 5
	KindUint32      Kind = 6
	KindInt64       Kind = 7
	KindUint64      Kind = 8
	KindFloat32     Kind = 9
	KindFloat64     Kind = 10
	KindStringFixed Kind = 11
	KindStringVLS   Kind = 12
	KindAlias       Kind = 19
	KindEnum        Kind = 20
	KindUnion       Kind = 21
	KindStruct      Kind = 30
)

type Type struct {
	Namespace *Namespace
	File      *File
	Kind      Kind
	Offset    int
	Align     int
	Size      int
	SizeConst *Const
	Enum      *Enum
	Message   *Struct
	Alias     *Alias
	Union     *Union
}

type Namespaces struct {
	Files      map[string]*File
	Namespaces map[string]*Namespace
}

type Namespace struct {
	Namespaces      *Namespaces
	Name            string
	Alias           []*Alias
	AliasByName     map[string]*Alias
	Constants       []*Const
	ConstantsByName map[string]*Const
	Enums           []*Enum
	EnumsByName     map[string]*Enum
	EnumOptions     map[string]*EnumOption
	Structs         []*Struct
	StructsByName   map[string]*Struct
}

type File struct {
	Path             string
	Namespaces       *Namespaces
	currentNamespace *Namespace
	Alias            []*Alias
	AliasByName      map[string]*Alias
	Constants        []*Const
	ConstantsByName  map[string]*Const
	Enums            []*Enum
	EnumsByName      map[string]*Enum
	EnumOptions      map[string]*EnumOption
	Structs          []*Struct
	StructsByName    map[string]*Struct
}

type Alias struct {
	Namespace *Namespace
	File      *File
	Name      string
	Type      Type
}

type Const struct {
	Namespace *Namespace
	File      *File
	Name      string
	Type      Type
	Length    int
	Value     *Value
	Comment   string
}

type Struct struct {
	Namespace    *Namespace
	File         *File
	Name         string
	MaxAlign     int
	Size         int
	VLS          bool
	Fields       []*Field
	FieldsByName map[string]*Field
	Proto        *parser.Message
}

type Field struct {
	Name  string
	Type  Type
	Init  *Value
	Proto *parser.Field
}

type ValueType int

const (
	ValueTypeInt        ValueType = 0
	ValueTypeFloat      ValueType = 1
	ValueTypeString     ValueType = 2
	ValueTypeConst      ValueType = 3
	ValueTypeEnumOption ValueType = 4
	ValueTypeSizeof     ValueType = 5
)

type Value struct {
	File       *File
	Namespace  *Namespace
	Type       ValueType
	Int        int64
	Float      float64
	String     string
	Const      *Const
	EnumOption *EnumOption
	Sizeof     string
}

type Enum struct {
	File          *File
	Namespace     *Namespace
	Name          string
	Type          Kind
	Options       []*EnumOption
	OptionsByName map[string]*EnumOption
	Proto         *parser.Enum
}

type EnumOption struct {
	File      *File
	Namespace *Namespace
	Enum      *Enum
	Name      string
	Value     int64
	Comment   string
	Proto     *parser.EnumField
}

type Union struct {
	Fields []*Field
}

func (s *Struct) Layout() {
	maxAlign := 0
	// Ensure length and alignment is set
	for _, field := range s.Fields {
		setLengthAndAlign(&field.Type)
		if maxAlign < field.Type.Align {
			maxAlign = field.Type.Align
		}
		if field.Type.Kind == KindStringVLS {
			s.VLS = true
		}
		if s.MaxAlign < field.Type.Align {
			field.Type.Align = s.MaxAlign
			if field.Type.Union != nil {
				for _, f := range field.Type.Union.Fields {
					f.Type.Align = s.MaxAlign
				}
				if field.Type.Kind == KindStringVLS {
					s.VLS = true
				}
			}
		}
	}

	offset := 0

	for i := 0; i < len(s.Fields); i++ {
		field := s.Fields[i]
		offset = align(offset, field.Type.Align)
		field.Type.Offset = offset

		if field.Type.Union != nil {
			for _, f := range field.Type.Union.Fields {
				f.Type.Offset = offset
			}
		}

		offset += field.Type.Size
	}
	if maxAlign < s.MaxAlign {
		s.Size = align(offset, maxAlign)
	} else {
		s.Size = align(offset, s.MaxAlign)
	}
}

func align(offset, align int) int {
	if offset == 0 || align == 0 {
		return 0
	}
	extras := offset % align
	if extras == 0 {
		return offset
	}
	return ((offset / align) + 1) * align
}
