package schema

import (
	"errors"
	"fmt"
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/pb"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"math"
	"sort"
	"strconv"
)

type Message struct {
	Fixed       *Struct
	VLS         *Struct
	Extension   *Message
	Extends     *Message
	NonStandard bool
}

func (m *Message) Name() string {
	if m.Fixed != nil {
		return m.Fixed.Name
	}
	if m.VLS != nil {
		return m.VLS.Name
	}
	return ""
}

func (m *Message) Type() uint16 {
	if m.Fixed != nil {
		return m.Fixed.Type
	}
	if m.VLS != nil {
		return m.VLS.Type
	}
	return 0
}

type Struct struct {
	Namespace    NamespaceKind
	Type         uint16
	Name         string
	Doc          *MessageDoc
	MaxAlign     int
	Size         int
	VLS          bool
	Fields       []*Field
	FieldsByName map[string]*Field
	Proto        *parser.Message
	Init         []byte // Exact initial binary value
}

// FieldsSlice attaches the methods of Interface to []*Field, sorting in increasing order.
type FieldsSlice []*Field

func (x FieldsSlice) Len() int { return len(x) }
func (x FieldsSlice) Less(i, j int) bool {
	if x[i].Proto == nil {
		return true
	}
	if x[j].Proto == nil {
		return false
	}
	in, err := strconv.ParseInt(x[i].Proto.FieldNumber, 10, 64)
	if err != nil {
		return true
	}
	jn, err := strconv.ParseInt(x[j].Proto.FieldNumber, 10, 64)
	if err != nil {
		return false
	}
	return in < jn
}
func (x FieldsSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x FieldsSlice) Sort() { sort.Sort(x) }

type Field struct {
	Struct          *Struct
	Name            string
	Doc             *FieldDoc
	Type            Type
	Initial         *Value
	InitExpression  string
	ClearExpression string
	Proto           *parser.Field
	ProtoType       pb.Type
	ProtoZigzag     bool
}

const MaxMessageSize = math.MaxUint16 - 4

func (f *Field) Validate() error {
	kind := f.Type.Kind
	if kind == KindAlias {
		if f.Type.Alias == nil {
			return errors.New(fmt.Sprintf("%s.%s is an alias type but Alias details are nil",
				f.Struct.Name, f.Name))
		}
		kind = f.Type.Alias.Type.Kind
	}

	if f.Initial != nil && f.Initial.Type == ValueTypeBool {
		if kind != KindBool {
			f.Type.Kind = KindBool
			kind = KindBool
		}
	}
	switch kind {
	default:
		return errors.New(fmt.Sprintf("%s.%s is KindUnknown",
			f.Struct.Name, f.Name))
	case KindBool:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeBool,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeBool:
				if f.Initial.Int > 0 {
					f.Initial.Int = 1
					f.Initial.Uint = 1
				}
				if f.Initial.Uint > 0 {
					f.Initial.Int = 1
					f.Initial.Uint = 1
				}
			case ValueTypeInt:
				if f.Initial.Int > 0 {
					f.Initial.Int = 1
				}
				f.Initial.Type = ValueTypeBool

			case ValueTypeUint:
				if f.Initial.Uint > 0 {
					f.Initial.Uint = 1
				}
				f.Initial.Type = ValueTypeBool

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeBool but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}
	case KindInt8:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeInt,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Uint = uint64(f.Initial.Int)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}
			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)
					if f.Initial.Int < math.MinInt8 || f.Initial.Int > math.MaxInt8 {
						return errors.New(fmt.Sprintf("%s.%s default value for int8 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Uint))
					}

				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)

					if f.Initial.Uint > math.MaxInt8 {
						return errors.New(fmt.Sprintf("%s.%s default value for int8 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Uint))
					}
				}
			case ValueTypeInt:
				if f.Initial.Int < math.MinInt8 || f.Initial.Int > math.MaxInt8 {
					return errors.New(fmt.Sprintf("%s.%s default value for int8 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
			case ValueTypeUint:
				if f.Initial.Uint > math.MaxInt8 {
					return errors.New(fmt.Sprintf("%s.%s default value for int8 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Uint))
				}
				f.Initial.Int = int64(f.Initial.Uint)
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeInt

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeInt but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindUint8:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeUint,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Int = int64(f.Initial.Uint)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}
			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)

					if f.Initial.Int < 0 || f.Initial.Uint > math.MaxUint8 {
						return errors.New(fmt.Sprintf("%s.%s default value for uint8 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Int))
					}
				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)

					if f.Initial.Uint > math.MaxUint8 {
						return errors.New(fmt.Sprintf("%s.%s default value for uint8 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Int))
					}
				}
			case ValueTypeUint:
				if f.Initial.Uint > math.MaxUint8 {
					return errors.New(fmt.Sprintf("%s.%s default value for uint8 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Uint))
				}
			case ValueTypeInt:
				if f.Initial.Int < 0 || f.Initial.Int > math.MaxUint8 {
					return errors.New(fmt.Sprintf("%s.%s default value for uint8 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
				f.Initial.Uint = uint64(f.Initial.Int)
				f.Initial.Int = 0
				f.Initial.Type = ValueTypeUint

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindInt16:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeInt,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Uint = uint64(f.Initial.Int)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}
			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)
					if f.Initial.Int < math.MinInt16 || f.Initial.Int > math.MaxInt16 {
						return errors.New(fmt.Sprintf("%s.%s default value for int32 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Uint))
					}

				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)

					if f.Initial.Uint > math.MaxInt16 {
						return errors.New(fmt.Sprintf("%s.%s default value for int16 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Uint))
					}
				}
			case ValueTypeInt:
				if f.Initial.Int < math.MinInt16 || f.Initial.Int > math.MaxInt16 {
					return errors.New(fmt.Sprintf("%s.%s default value for int16 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
			case ValueTypeUint:
				if f.Initial.Uint > math.MaxInt16 {
					return errors.New(fmt.Sprintf("%s.%s default value for int16 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Uint))
				}
				f.Initial.Int = int64(f.Initial.Uint)
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeInt

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeInt but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindUint16:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeUint,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Int = int64(f.Initial.Uint)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}
			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)

					if f.Initial.Int < 0 || f.Initial.Uint > math.MaxUint16 {
						return errors.New(fmt.Sprintf("%s.%s default value for uint16 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Int))
					}
				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)

					if f.Initial.Uint > math.MaxUint16 {
						return errors.New(fmt.Sprintf("%s.%s default value for uint16 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Int))
					}
				}
			case ValueTypeUint:
				if f.Initial.Uint > math.MaxUint16 {
					return errors.New(fmt.Sprintf("%s.%s default value for uint16 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Uint))
				}
			case ValueTypeInt:
				if f.Initial.Int < 0 || f.Initial.Int > math.MaxUint16 {
					return errors.New(fmt.Sprintf("%s.%s default value for uint16 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
				f.Initial.Uint = uint64(f.Initial.Int)
				f.Initial.Int = 0
				f.Initial.Type = ValueTypeUint

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindInt32:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeInt,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Uint = uint64(f.Initial.Int)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}

			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)
					if f.Initial.Int < math.MinInt32 || f.Initial.Int > math.MaxInt32 {
						return errors.New(fmt.Sprintf("%s.%s default value for int32 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Uint))
					}

				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)

					if f.Initial.Uint > math.MaxInt32 {
						return errors.New(fmt.Sprintf("%s.%s default value for int32 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Uint))
					}
				}

			case ValueTypeInt:
				if f.Initial.Int < math.MinInt32 || f.Initial.Int > math.MaxInt32 {
					return errors.New(fmt.Sprintf("%s.%s default value for int32 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
			case ValueTypeUint:
				if f.Initial.Uint > math.MaxInt32 {
					return errors.New(fmt.Sprintf("%s.%s default value for int32 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Uint))
				}
				f.Initial.Int = int64(f.Initial.Uint)
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeInt

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeInt but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindUint32:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeUint,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Int = int64(f.Initial.Uint)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}
			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)

					if f.Initial.Int < 0 || f.Initial.Uint > math.MaxUint32 {
						return errors.New(fmt.Sprintf("%s.%s default value for uint32 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Int))
					}
				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)

					if f.Initial.Uint > math.MaxUint32 {
						return errors.New(fmt.Sprintf("%s.%s default value for uint32 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Int))
					}
				}
			case ValueTypeUint:
				if f.Initial.Uint > math.MaxUint32 {
					return errors.New(fmt.Sprintf("%s.%s default value for uint32 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Uint))
				}
			case ValueTypeInt:
				if f.Initial.Int < 0 || f.Initial.Int > math.MaxUint32 {
					return errors.New(fmt.Sprintf("%s.%s default value for uint32 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
				f.Initial.Uint = uint64(f.Initial.Int)
				f.Initial.Int = 0
				f.Initial.Type = ValueTypeUint

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindInt64:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeInt,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Uint = uint64(f.Initial.Int)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}

			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)

				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)

					if f.Initial.Uint > math.MaxInt64 {
						return errors.New(fmt.Sprintf("%s.%s default value for int64 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Uint))
					}
				}

			case ValueTypeInt:
			case ValueTypeUint:
				if f.Initial.Uint > math.MaxInt64 {
					return errors.New(fmt.Sprintf("%s.%s default value for int64 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Uint))
				}
				f.Initial.Int = int64(f.Initial.Uint)
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeInt

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeInt but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindUint64:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeUint,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeSizeof:
				f.Initial.Int = int64(f.Initial.Uint)
				if f.Initial.Uint > MaxMessageSize {
					return errors.New(fmt.Sprintf("%s.%s default is a sizeof which cannot be larger than %d",
						f.Struct.Name, f.Name, MaxMessageSize))
				}
			case ValueTypeConst:
				if f.Initial.Const == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const details are nil",
						f.Struct.Name, f.Name))
				}
				if f.Initial.Const.Value == nil {
					return errors.New(fmt.Sprintf("%s.%s default is a const however Const.Value details are nil",
						f.Struct.Name, f.Name))
				}
				switch f.Initial.Const.Value.Type {
				case ValueTypeInt:
					f.Initial.Int = f.Initial.Const.Value.Int
					f.Initial.Uint = uint64(f.Initial.Const.Value.Int)

					if f.Initial.Int < 0 {
						return errors.New(fmt.Sprintf("%s.%s default value for uint64 is out of range %d",
							f.Struct.Name, f.Name, f.Initial.Int))
					}
				case ValueTypeUint:
					f.Initial.Uint = f.Initial.Const.Value.Uint
					f.Initial.Int = int64(f.Initial.Const.Value.Uint)
				}
			case ValueTypeUint:
			case ValueTypeInt:
				if f.Initial.Int < 0 {
					return errors.New(fmt.Sprintf("%s.%s default value for uint64 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
				f.Initial.Uint = uint64(f.Initial.Int)
				f.Initial.Int = 0
				f.Initial.Type = ValueTypeUint

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindFloat32:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeFloat,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeFloat:
				if f.Initial.Float > math.MaxFloat32 {
					return errors.New(fmt.Sprintf("%s.%s default value for float32 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
			case ValueTypeFloat32Max:
				f.Initial.Float = math.MaxFloat32
			case ValueTypeFloat64Max:
				return errors.New(fmt.Sprintf("%s.%s default value for float32 cannot be ValueTypeFloat64Max",
					f.Struct.Name, f.Name))
			case ValueTypeInt:
				v := float64(f.Initial.Int)
				if v > math.MaxFloat32 {
					return errors.New(fmt.Sprintf("%s.%s default value for float32 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
				f.Initial.Float = v
				f.Initial.Int = 0
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeFloat

			case ValueTypeUint:
				v := float64(f.Initial.Uint)
				if v > math.MaxFloat32 {
					return errors.New(fmt.Sprintf("%s.%s default value for float32 is out of range %d",
						f.Struct.Name, f.Name, f.Initial.Int))
				}
				f.Initial.Float = v
				f.Initial.Int = 0
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeFloat

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindFloat64:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeFloat,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeFloat:
			case ValueTypeFloat32Max:
				f.Initial.Float = math.MaxFloat32
			case ValueTypeFloat64Max:
				f.Initial.Float = math.MaxFloat64
			case ValueTypeInt:
				f.Initial.Float = float64(f.Initial.Int)
				f.Initial.Int = 0
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeFloat

			case ValueTypeUint:
				f.Initial.Float = float64(f.Initial.Uint)
				f.Initial.Int = 0
				f.Initial.Uint = 0
				f.Initial.Type = ValueTypeFloat

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindStringFixed:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeString,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeString:
			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}
	case KindStringVLS:
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace: f.Struct.Namespace,
				Type:      ValueTypeString,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeString:
			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindEnum:
		if f.Type.Enum == nil {
			return errors.New(fmt.Sprintf("%s.%s is an enum type but Enum details are nil",
				f.Struct.Name, f.Name))
		}
		if len(f.Type.Enum.Options) == 0 {
			return errors.New(fmt.Sprintf("%s.%s is an enum type but Enum has no options",
				f.Struct.Name, f.Name))
		}
		var zeroOption *EnumOption
		for _, option := range f.Type.Enum.Options {
			if option.Value == 0 {
				zeroOption = option
				break
			}
		}
		if zeroOption == nil {
			zeroOption = f.Type.Enum.Options[0]
		}
		if f.Initial == nil || f.Initial.Type == ValueTypeUnknown {
			f.Initial = &Value{
				Namespace:  f.Struct.Namespace,
				Type:       ValueTypeEnumOption,
				EnumOption: zeroOption,
			}
		} else {
			switch f.Initial.Type {
			case ValueTypeEnumOption:
			case ValueTypeInt:
				for _, option := range f.Type.Enum.Options {
					if option.Value == f.Initial.Int {
						zeroOption = option
						break
					}
				}
				f.Initial.Type = ValueTypeEnumOption
				f.Initial.EnumOption = zeroOption
			case ValueTypeUint:
				for _, option := range f.Type.Enum.Options {
					if option.Value == int64(f.Initial.Uint) {
						zeroOption = option
						break
					}
				}
				f.Initial.Type = ValueTypeEnumOption
				f.Initial.EnumOption = zeroOption
			case ValueTypeString:
				if len(f.Initial.Str) > 0 {
					var found *EnumOption
					for _, option := range f.Type.Enum.Options {
						if option.Name == f.Initial.Str {
							found = option
							break
						}
					}
					if found == nil {
						return errors.New(fmt.Sprintf("%s.%s is an enum type defaulted to option %s which does not exist",
							f.Struct.Name, f.Name, f.Initial.Str))
					}
					f.Initial.EnumOption = found
					f.Initial.Type = ValueTypeEnumOption
				} else {
					f.Initial.Type = ValueTypeEnumOption
					f.Initial.EnumOption = zeroOption
				}

			default:
				return errors.New(fmt.Sprintf("%s.%s default value type expected ValueTypeUint but found %s",
					f.Struct.Name, f.Name, f.Initial.Type))
			}
		}

	case KindUnion:
		if f.Type.Union == nil {
			return errors.New(fmt.Sprintf("%s.%s is a union type but the Union details are nil",
				f.Struct.Name, f.Name))
		}
		for _, field := range f.Type.Union.Fields {
			if err := field.Validate(); err != nil {
				return err
			}
		}

	case KindStruct:
		return errors.New("currently structs cannot be embedded in other structs")
	case KindList:
		return errors.New("list types not supported")
	}
	if f.Initial == nil {
		f.Initial = &Value{}
	}
	return nil
}

func (s *Struct) Validate() error {
	for _, field := range s.Fields {
		if err := field.Validate(); err != nil {
			return err
		}
	}
	return nil
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

	s.Init = make([]byte, s.Size)

	if s.Size <= 0 {
		return
	}
	p := message.PointerOf(s.Init)
	p.Zero(uintptr(s.Size))
	for _, field := range s.Fields {
		if field.Type.Union != nil {
			for _, field := range field.Type.Union.Fields {
				if field.Initial == nil {
					continue
				}
				field.Init(p)
			}
		} else {
			if field.Initial == nil {
				continue
			}
			field.Init(p)
		}
	}
}

func (f *Field) Init(p message.Pointer) {
	v := f.Initial
	if v == nil {
		// Zero
		return
	}
	if f.Type.Union != nil {
		return
	}
	kind := f.Type.Kind
	if f.Type.Alias != nil {
		kind = f.Type.Alias.Type.Kind
	} else if f.Type.Enum != nil {
		kind = f.Type.Enum.Type
	}
	offset := f.Type.Offset
	for v.Type == ValueTypeConst {
		v = v.Const.Value
	}
	switch v.Type {
	case ValueTypeSizeof:
		pointerSetUint(f.Type.Kind, p, offset, uint64(f.Struct.Size))
	case ValueTypeFloat32Max:
		p.SetFloat32LE(offset, math.MaxFloat32)
	case ValueTypeFloat64Max:
		p.SetFloat64LE(offset, math.MaxFloat64)
	case ValueTypeInt:
		pointerSetInt(f.Type.Kind, p, offset, v.Int)
	case ValueTypeUint:
		pointerSetUint(f.Type.Kind, p, offset, v.Uint)
	case ValueTypeFloat:
		switch kind {
		case KindFloat32:
			p.SetFloat32LE(offset, float32(float32(v.Float)))
		case KindFloat64:
			p.SetFloat64LE(offset, v.Float)
		default:
			panic("expected float32 or float64")
		}
	case ValueTypeBool:
		p.SetUint8(offset, uint8(v.Int))
	case ValueTypeEnumOption:
		pointerSetInt(kind, p, offset, v.EnumOption.Value)
	case ValueTypeString:
		if f.Struct.VLS {
			if len(v.Str) > 0 {
				panic("implement Value.Set for VLS strings")
			}
		} else if len(v.Str) > 0 {
			p.SetStringFixed(offset, offset+f.Type.Size, v.Str)
		}
	case ValueTypeConst:
		panic("ValueTypeConst needs to be unwrapped first")
	default:
		panic("Value.Type is not supported")
	}
}

func pointerSetInt(kind Kind, p message.Pointer, offset int, value int64) {
	switch kind {
	case KindInt8:
		p.SetInt8(offset, int8(value))
	case KindInt16:
		p.SetInt16LE(offset, int16(value))
	case KindInt32:
		p.SetInt32LE(offset, int32(value))
	case KindInt64:
		p.SetInt64LE(offset, value)
	case KindUint8:
		p.SetUint8(offset, uint8(value))
	case KindUint16:
		p.SetUint16LE(offset, uint16(value))
	case KindUint32:
		p.SetUint32LE(offset, uint32(value))
	case KindUint64:
		p.SetUint64LE(offset, uint64(value))
	}
}

func pointerSetUint(kind Kind, p message.Pointer, offset int, value uint64) {
	switch kind {
	case KindInt8:
		p.SetInt8(offset, int8(value))
	case KindInt16:
		p.SetInt16LE(offset, int16(value))
	case KindInt32:
		p.SetInt32LE(offset, int32(value))
	case KindInt64:
		p.SetInt64LE(offset, int64(value))
	case KindUint8:
		p.SetUint8(offset, uint8(value))
	case KindUint16:
		p.SetUint16LE(offset, uint16(value))
	case KindUint32:
		p.SetUint32LE(offset, uint32(value))
	case KindUint64:
		p.SetUint64LE(offset, value)
	}
}

func (s *Struct) Bind(proto *parser.Message) error {
	for _, pf := range proto.MessageBody {
		switch pf := pf.(type) {
		case *parser.Field:
			f := s.FieldsByName[pf.FieldName]
			if f == nil {
				f = s.FieldsByName["m_"+pf.FieldName]
			}
			if f == nil {
				return fmt.Errorf("proto: %s.%s does not match any fields for %s", proto.MessageName, pf.FieldName, s.Name)
			}

			kind := f.Type.Kind
			if f.Type.Alias != nil {
				kind = f.Type.Alias.Type.Kind
			}
			if f.Type.Enum != nil {
				kind = f.Type.Enum.Type
			}

			switch pf.Type {
			case "sint32", "sint64":
				f.ProtoType = pb.VarintType
				f.ProtoZigzag = true
				switch kind {
				case KindBool, KindInt8, KindUint8, KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindEnum:
				default:
					return fmt.Errorf("proto: %s.%s type %s does not match field type: KindBool | KindInt8 | KindUint8 | KindInt16 | KindUint16 | KindInt32 | KindUint32 | KindInt64 | KindUint64", proto.MessageName, pf.FieldName, pf.Type)
				}

			case "fixed32", "sfixed32":
				f.ProtoType = pb.Fixed32Type
				switch kind {
				case KindBool, KindInt8, KindUint8, KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindEnum:
				default:
					return fmt.Errorf("proto: %s.%s type %s does not match field type: KindBool | KindInt8 | KindUint8 | KindInt16 | KindUint16 | KindInt32 | KindUint32 | KindInt64 | KindUint64", proto.MessageName, pf.FieldName, pf.Type)
				}

			case "fixed64", "sfixed64":
				f.ProtoType = pb.Fixed64Type
				switch kind {
				case KindBool, KindInt8, KindUint8, KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindEnum:
				default:
					return fmt.Errorf("proto: %s.%s type %s does not match field type: KindBool | KindInt8 | KindUint8 | KindInt16 | KindUint16 | KindInt32 | KindUint32 | KindInt64 | KindUint64", proto.MessageName, pf.FieldName, pf.Type)
				}

			case "int32", "uint32", "short", "long", "int", "int64", "uint64", "bool":
				f.ProtoType = pb.VarintType
				switch kind {
				case KindBool, KindInt8, KindUint8, KindInt16, KindUint16, KindInt32, KindUint32, KindInt64, KindUint64, KindEnum:
				default:
					return fmt.Errorf("proto: %s.%s type %s does not match field type: KindBool | KindInt8 | KindUint8 | KindInt16 | KindUint16 | KindInt32 | KindUint32 | KindInt64 | KindUint64", proto.MessageName, pf.FieldName, pf.Type)
				}

			case "float":
				f.ProtoType = pb.Fixed32Type
				switch kind {
				case KindFloat32, KindFloat64:
				default:
					return fmt.Errorf("proto: %s.%s type %s does not match field type: KindFloat32 | KindFloat64", proto.MessageName, pf.FieldName, pf.Type)
				}

			case "double":
				f.ProtoType = pb.Fixed64Type
				switch kind {
				case KindFloat32, KindFloat64:
				default:
					return fmt.Errorf("proto: %s.%s type %s does not match field type: KindFloat32 | KindFloat64", proto.MessageName, pf.FieldName, pf.Type)
				}

			case "string", "bytes":
				switch kind {
				case KindStringFixed, KindStringVLS:
				default:
					return fmt.Errorf("proto: %s.%s type %s does not match field type KindStringFixed | KindStringVLS", proto.MessageName, pf.FieldName, pf.Type)
				}

			default:
				enum := f.Type.Enum
				if enum != nil {
					if enum.Name != pf.Type {
						return fmt.Errorf("proto: %s.%s type %s does not match field type enum: %s", proto.MessageName, pf.FieldName, pf.Type, enum.Name)
					}
					f.ProtoType = pb.VarintType
				} else {
					return fmt.Errorf("proto: %s.%s type %s does not match field type enum: %s", proto.MessageName, pf.FieldName, pf.Type, enum.Name)
				}
			}
			f.Proto = pf
		}
	}
	s.Proto = proto
	return nil
}
