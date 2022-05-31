package schema

type Kind byte

const (
	KindUnknown     Kind = 0
	KindBool        Kind = 1
	KindInt8        Kind = 2
	KindUint8       Kind = 3
	KindInt16       Kind = 4
	KindUint16      Kind = 5
	KindInt32       Kind = 6
	KindUint32      Kind = 7
	KindInt64       Kind = 8
	KindUint64      Kind = 9
	KindFloat32     Kind = 10
	KindFloat64     Kind = 11
	KindStringFixed Kind = 12
	KindStringVLS   Kind = 13
	KindAlias       Kind = 20
	KindEnum        Kind = 21
	KindUnion       Kind = 22
	KindStruct      Kind = 30
	KindList        Kind = 40
)

func (k Kind) IsPrimitive() bool {
	switch k {
	case KindBool, KindInt8, KindInt16, KindInt32, KindInt64, KindUint8,
		KindUint16, KindUint32, KindUint64, KindFloat32, KindFloat64,
		KindStringFixed, KindStringVLS:
		return true
	default:
		return false
	}
}

func (k Kind) IsString() bool {
	switch k {
	case KindStringFixed, KindStringVLS:
		return true
	default:
		return false
	}
}

type Type struct {
	Namespace NamespaceKind
	Kind      Kind
	Offset    int
	Align     int
	Size      int
	SizeConst *Const
	Enum      *Enum
	Struct    *Struct
	Alias     *Alias
	Union     *Union
}

func (t *Type) IsNonStandard() bool {
	if t.Enum != nil {
		return t.Enum.Namespace == NamespaceKindNonStandard
	}
	if t.Struct != nil {
		return t.Struct.Namespace == NamespaceKindNonStandard
	}
	if t.Alias != nil {
		return t.Alias.Namespace == NamespaceKindNonStandard
	}
	return false
}
