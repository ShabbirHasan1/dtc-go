package generator

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen/schema"
	"sort"
	"strconv"
	"strings"
)

type NameTransformer func(name string) string

type Namer struct {
	Struct     NameTransformer
	Enum       NameTransformer
	EnumOption NameTransformer
	Alias      NameTransformer
	Field      NameTransformer
}

type Generator struct {
	namer           *Namer
	schema          *schema.Schema
	packageName     string
	all             map[string]interface{}
	aliases         []*Alias
	aliasesByName   map[string]*Alias
	constants       []*Constant
	constantsByName map[string]*Constant
	enums           []*Enum
	enumsByName     map[string]*Enum
	messages        []*Message
	messagesByName  map[string]*Message
}

func cleanName(n string) string {
	n = strings.TrimPrefix(n, "s_")
	n = strings.TrimPrefix(n, "m_")
	n = strings.TrimPrefix(n, "t_")
	n = strings.TrimSpace(n)
	return n
}

func NewGenerator(s *schema.Schema, namer *Namer) (*Generator, error) {
	if namer == nil {
		namer = &Namer{}
	}
	if namer.Struct == nil {
		namer.Struct = cleanName
	}
	if namer.Enum == nil {
		namer.Enum = cleanName
	}
	if namer.EnumOption == nil {
		namer.EnumOption = func(name string) string {
			return name
		}
	}
	if namer.Alias == nil {
		namer.Alias = cleanName
	}
	if namer.Field == nil {
		namer.Field = cleanName
	}
	generator := &Generator{
		namer:           namer,
		schema:          s,
		all:             make(map[string]interface{}),
		aliasesByName:   make(map[string]*Alias),
		messagesByName:  make(map[string]*Message),
		enumsByName:     make(map[string]*Enum),
		constantsByName: make(map[string]*Constant),
	}

	for _, alias := range s.Aliases {
		a := &Alias{
			Alias: alias,
			Name:  cleanName(alias.Name),
		}
		if !alias.Type.Kind.IsPrimitive() {
			return nil, fmt.Errorf("alias kind is not primitive: %d", alias.Type.Kind)
		}
		if generator.all[a.Name] != nil {
			return nil, fmt.Errorf("alias name conflict: %s", a.Name)
		}
		generator.all[a.Name] = a
		if generator.aliasesByName[a.Name] != nil {
			return nil, fmt.Errorf("duplicate alias name: %s", a.Name)
		}
		generator.aliases = append(generator.aliases, a)
		generator.aliasesByName[a.Name] = a
	}

	for _, constant := range s.Constants {
		c := &Constant{
			Const: constant,
			Name:  constant.Name,
		}
		if !constant.Type.Kind.IsPrimitive() {
			return nil, fmt.Errorf("constant kind is not primitive: %d", constant.Type.Kind)
		}

		if generator.all[c.Name] != nil {
			return nil, fmt.Errorf("constant name conflict: %s", c.Name)
		}
		generator.all[c.Name] = c
		if generator.constantsByName[c.Name] != nil {
			return nil, fmt.Errorf("duplicate constant name: %s", c.Name)
		}
		generator.constants = append(generator.constants, c)
		generator.constantsByName[c.Name] = c
	}
	for _, enum := range s.Enums {
		e := &Enum{
			Enum:          enum,
			Name:          cleanName(enum.Name),
			OptionsByName: make(map[string]*EnumOption),
		}
		if !e.Type.IsPrimitive() {
			return nil, fmt.Errorf("enum kind is not primitive: %d", enum.Type)
		}

		for _, option := range enum.Options {
			o := &EnumOption{
				EnumOption: option,
				Enum:       e,
				Name:       option.Name,
			}
			if generator.all[o.Name] != nil {
				return nil, fmt.Errorf("EnumOption name conflict: %s", o.Name)
			}
			generator.all[o.Name] = o
			e.Options = append(e.Options, o)
			e.OptionsByName[o.Name] = o
		}

		if generator.all[e.Name] != nil {
			return nil, fmt.Errorf("enum name conflict: %s", e.Name)
		}
		generator.all[e.Name] = e
		if generator.enumsByName[e.Name] != nil {
			return nil, fmt.Errorf("duplicate enum name: %s", e.Name)
		}

		generator.enums = append(generator.enums, e)
		generator.enumsByName[e.Name] = e
	}

	createStruct := func(st *schema.Struct) (*Struct, error) {
		if st == nil {
			return nil, nil
		}

		msg := &Struct{
			Struct:       st,
			Name:         cleanName(st.Name),
			FieldsByName: make(map[string]*Field),
		}
		if !st.VLS && s.MessagesByName[st.Name].VLS != nil {
			msg.Name = fmt.Sprintf("%sFixed", msg.Name)
		}
		for _, field := range st.Fields {
			var (
				f = &Field{
					Struct: msg,
					Field:  field,
					Name:   cleanName(field.Name),
				}
			)
			if field.Type.Union != nil {
				for _, field := range field.Type.Union.Fields {
					ff := &Field{
						Struct: msg,
						Field:  field,
						Name:   cleanName(field.Name),
					}
					if msg.FieldsByName[ff.Name] != nil {
						return nil, fmt.Errorf("%s.%s duplicate field name after cleanName", st.Name, field.Name)
					}
					f.Fields = append(f.Fields, ff)
					msg.FieldsByName[ff.Name] = ff
				}
			} else {
				if msg.FieldsByName[f.Name] != nil {
					return nil, fmt.Errorf("%s.%s duplicate field name after cleanName", st.Name, field.Name)
				}
				msg.FieldsByName[f.Name] = f
			}
			msg.Fields = append(msg.Fields, f)
		}
		return msg, nil
	}
	for _, message := range s.Messages {
		var (
			m = &Message{
				NonStandard: message.NonStandard,
			}
			err error
		)
		if m.Fixed, err = createStruct(message.Fixed); err != nil {
			return nil, err
		}
		if m.VLS, err = createStruct(message.VLS); err != nil {
			return nil, err
		}
		name := ""
		if m.Fixed != nil {
			name = m.Fixed.Name
			m.Fixed.Message = m
		}
		if m.VLS != nil {
			name = m.VLS.Name
			m.VLS.Message = m
		}

		if len(name) == 0 {
			return nil, fmt.Errorf("nil message")
		}
		if generator.messagesByName[name] != nil {
			return nil, fmt.Errorf("duplicate message name after cleanName: %s", name)
		}
		generator.messages = append(generator.messages, m)
		generator.messagesByName[name] = m
	}

	return generator, nil
}

type Message struct {
	Fixed       *Struct
	VLS         *Struct
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
