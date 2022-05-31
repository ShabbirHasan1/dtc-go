package golang

import (
	"github.com/moontrade/dtc-go/codegen/schema"
	"sort"
	"strconv"
	"strings"
)

type Message struct {
	Fixed       *Struct
	VLS         *Struct
	Extension   *Message
	Extends     *Message
	NonStandard bool
}

func (m *Message) HasSerializers() bool {
	s := m.VLS
	if s == nil {
		s = m.Fixed
	}
	if s == nil {
		return false
	}
	return true
}

func (m *Message) HasProtobuf() bool {
	s := m.VLS
	if s == nil {
		s = m.Fixed
	}
	if s == nil {
		return false
	}
	return s.Proto != nil
}

func (m *Message) Name() string {
	if m.VLS != nil {
		return m.VLS.Name
	}
	return m.Fixed.Name
}

func (m *Message) Doc() *schema.MessageDoc {
	if m.VLS != nil {
		return m.VLS.Doc
	}
	if m.Fixed != nil {
		return m.Fixed.Doc
	}
	return nil
}

type Alias struct {
	*schema.Alias
	Name string
}

type Constant struct {
	*schema.Const
	Name string
}

type Enum struct {
	*schema.Enum
	Name          string
	Options       []*EnumOption
	OptionsByName map[string]*EnumOption
}

type EnumOption struct {
	*schema.EnumOption
	Enum *Enum
	Name string
}

type Struct struct {
	*schema.Struct
	Message      *Message
	Name         string
	Fields       []*Field
	FieldsByName map[string]*Field
}

func (s *Struct) Suffix() string {
	if s.VLS {
		return "VLS"
	} else {
		return "Fixed"
	}
}

func (s *Struct) Filter(fn func(field *Field) bool) []*Field {
	if fn == nil {
		fn = func(field *Field) bool {
			return true
		}
	}
	var r []*Field
	for _, field := range s.Fields {
		if fn(field) {
			r = append(r, field)
		}
	}
	return r
}

type Field struct {
	*schema.Field
	Struct *Struct
	Name   string
	Fields []*Field
}

// FieldsByProtoNumber attaches the methods of Interface to []*Field, sorting in increasing order.
type FieldsByProtoNumber []*Field

func (x FieldsByProtoNumber) Len() int { return len(x) }
func (x FieldsByProtoNumber) Less(i, j int) bool {
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
func (x FieldsByProtoNumber) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x FieldsByProtoNumber) Sort() { sort.Sort(x) }

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
	if strings.HasPrefix(n, "n_") {
		return strings.TrimSpace(n[2:])
	}
	n = strings.TrimSpace(n)
	return n
}

func primitiveKindName(kind schema.Kind) string {
	switch kind {
	case schema.KindInt8:
		return "int8"
	case schema.KindInt16:
		return "int16"
	case schema.KindInt32:
		return "int32"
	case schema.KindInt64:
		return "int64"
	case schema.KindUint8:
		return "uint8"
	case schema.KindUint16:
		return "uint16"
	case schema.KindUint32:
		return "uint32"
	case schema.KindUint64:
		return "uint64"
	case schema.KindFloat32:
		return "float32"
	case schema.KindFloat64:
		return "float64"
	case schema.KindStringFixed:
		return "string"
	case schema.KindStringVLS:
		return "string"
	}
	return "invalid"
}
