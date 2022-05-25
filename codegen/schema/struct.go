package schema

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"unsafe"

	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/pb"
	"github.com/moontrade/nogc"
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Message struct {
	Fixed       *Struct
	VLS         *Struct
	NonStandard bool
}

type Struct struct {
	Namespace    *Namespace
	File         *CHeaderFile
	Type         uint16
	Name         string
	NamePretty   string
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
	p := nogc.Pointer(unsafe.Pointer(&s.Init[0]))
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

func (f *Field) Init(p nogc.Pointer) {
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
		pointerSetUint(f.Type.Kind, p+nogc.Pointer(offset), uint64(f.Struct.Size))
	case ValueTypeFloat32Max:
		p.SetFloat32LE(offset, math.MaxFloat32)
	case ValueTypeFloat64Max:
		p.SetFloat64LE(offset, math.MaxFloat64)
	case ValueTypeInt:
		pointerSetInt(f.Type.Kind, p+nogc.Pointer(offset), v.Int)
	case ValueTypeUint:
		pointerSetUint(f.Type.Kind, p+nogc.Pointer(offset), v.Uint)
	case ValueTypeFloat:
		switch kind {
		case KindFloat32:
			p.SetFloat32LE(offset, float32(v.Float32))
		case KindFloat64:
			p.SetFloat64LE(offset, v.Float64)
		default:
			panic("expected float32 or float64")
		}
	case ValueTypeBool:
		p.SetUint8(offset, uint8(v.Int))
	case ValueTypeEnumOption:
		pointerSetInt(kind, p+nogc.Pointer(offset), v.EnumOption.Value)
	case ValueTypeString:
		if f.Struct.VLS {
			panic("implement Value.Set for VLS strings")
		}
		if len(v.Str) > 0 {
			message.SetStringFixed(p, uint16(offset+f.Type.Size), offset, v.Str)
		}
	case ValueTypeConst:
		panic("ValueTypeConst needs to be unwrapped first")
	default:
		panic("Value.Type is not supported")
	}
}

func pointerSetInt(kind Kind, p nogc.Pointer, value int64) {
	switch kind {
	case KindInt8:
		p.SetInt8(0, int8(value))
	case KindInt16:
		p.SetInt16LE(0, int16(value))
	case KindInt32:
		p.SetInt32LE(0, int32(value))
	case KindInt64:
		p.SetInt64LE(0, value)
	case KindUint8:
		p.SetUint8(0, uint8(value))
	case KindUint16:
		p.SetUint16LE(0, uint16(value))
	case KindUint32:
		p.SetUint32LE(0, uint32(value))
	case KindUint64:
		p.SetUint64LE(0, uint64(value))
	}
}

func pointerSetUint(kind Kind, p nogc.Pointer, value uint64) {
	switch kind {
	case KindInt8:
		p.SetInt8(0, int8(value))
	case KindInt16:
		p.SetInt16LE(0, int16(value))
	case KindInt32:
		p.SetInt32LE(0, int32(value))
	case KindInt64:
		p.SetInt64LE(0, int64(value))
	case KindUint8:
		p.SetUint8(0, uint8(value))
	case KindUint16:
		p.SetUint16LE(0, uint16(value))
	case KindUint32:
		p.SetUint32LE(0, uint32(value))
	case KindUint64:
		p.SetUint64LE(0, value)
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
