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

type Value struct {
	Namespace  *Namespace
	Type       ValueType
	Int        int64
	Uint       uint64
	Float32    float64
	Float64    float64
	Str        string
	Const      *Const
	EnumOption *EnumOption
	Sizeof     string
	Struct     *Struct
}
