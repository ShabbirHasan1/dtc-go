package schema

type ValueType int

const (
	ValueTypeUnknown    ValueType = 0
	ValueTypeInt        ValueType = 1
	ValueTypeUint       ValueType = 2
	ValueTypeBool       ValueType = 9
	ValueTypeFloat      ValueType = 10
	ValueTypeFloat32Max ValueType = 11
	ValueTypeFloat64Max ValueType = 12
	ValueTypeString     ValueType = 20
	ValueTypeConst      ValueType = 30
	ValueTypeEnumOption ValueType = 40
	ValueTypeSizeof     ValueType = 50
)

func (v ValueType) String() string {
	switch v {
	case ValueTypeUnknown:
		return "ValueTypeUnknown"
	case ValueTypeInt:
		return "ValueTypeInt"
	case ValueTypeUint:
		return "ValueTypeUint"
	case ValueTypeBool:
		return "ValueTypeBool"
	case ValueTypeFloat:
		return "ValueTypeFloat"
	case ValueTypeFloat32Max:
		return "ValueTypeFloat32Max"
	case ValueTypeFloat64Max:
		return "ValueTypeFloat64Max"
	case ValueTypeString:
		return "ValueTypeString"
	case ValueTypeConst:
		return "ValueTypeConst"
	case ValueTypeEnumOption:
		return "ValueTypeEnumOption"
	case ValueTypeSizeof:
		return "ValueTypeSizeof"
	default:
		return "ValueTypeUnknown"
	}
}

type Value struct {
	Namespace NamespaceKind
	Type      ValueType
	Int       int64
	Uint      uint64
	//Float32    float64
	Float      float64
	Str        string
	Const      *Const
	EnumOption *EnumOption
	Sizeof     string
	Struct     *Struct
}
