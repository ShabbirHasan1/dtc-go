package rust

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/moontrade/dtc-go/codegen/schema"
)

type Generator struct {
	config          *Config
	schema          *schema.Schema
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

func NewGenerator(config *Config, s *schema.Schema) (*Generator, error) {
	generator := &Generator{
		config:          config,
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
			Name:  cleanStructName(alias.Name),
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
			Name:          cleanStructName(enum.Name),
			OptionsByName: make(map[string]*EnumOption),
		}
		if !e.Type.IsPrimitive() {
			return nil, fmt.Errorf("enum kind is not primitive: %d", enum.Type)
		}

		for _, option := range enum.Options {
			o := &EnumOption{
				EnumOption: option,
				Enum:       e,
				Name:       toCamelCase(option.Name),
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

		suffix := "VLS"
		if !st.VLS {
			suffix = "Fixed"
		}
		msg := &Struct{
			Struct:       st,
			Name:         fmt.Sprintf("%s%s", cleanStructName(st.Name), suffix),
			UnsafeName:   fmt.Sprintf("%s%s%s", cleanStructName(st.Name), suffix, "Unsafe"),
			DataName:     fmt.Sprintf("%s%s%s", cleanStructName(st.Name), suffix, "Data"),
			FieldsByName: make(map[string]*Field),
		}
		for _, field := range st.Fields {
			var (
				f = &Field{
					Struct: msg,
					Field:  field,
					Name:   cleanFieldName(field.Name),
					Init:   initValue(field),
				}
			)
			if strings.ToLower(field.Name) == "type" {
				if field.Initial == nil {
					panic("type must have a value")
				}
				msg.TypeConst = generator.constantsByName[field.Initial.Const.Name]
				if msg.TypeConst == nil {
					panic("type must have a type const")
				}
			}
			if field.Type.Union != nil {
				for _, field := range field.Type.Union.Fields {
					ff := &Field{
						Struct: msg,
						Field:  field,
						Name:   cleanFieldName(field.Name),
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

	byType := make(map[uint16]*Message)

	for _, message := range s.Messages {
		var (
			m = &Message{
				NonStandard: message.NonStandard,
			}
			err error
		)
		// if !config.NonStandard && message.NonStandard {
		// 	continue
		// }
		// if config.NonStandard && !message.NonStandard {
		// 	continue
		// }
		if m.Fixed, err = createStruct(message.Fixed); err != nil {
			return nil, err
		}
		if m.VLS, err = createStruct(message.VLS); err != nil {
			return nil, err
		}
		name := ""
		if m.Fixed != nil {
			if m.Fixed.TypeConst == nil {
				fmt.Printf("[WARN] type %s does not have a TypeConst\n", m.Fixed.Name)
				continue
			}
			name = m.Fixed.Name
			m.Fixed.Message = m
		}
		if m.VLS != nil {
			if m.VLS.TypeConst == nil {
				fmt.Printf("[WARN] type %s does not have a TypeConst\n", m.VLS.Name)
				continue
			}
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

		existing := byType[m.TypeCode()]
		if existing != nil {
			if existing.Fixed != nil && m.Fixed != nil {
				if existing.Fixed.Size > m.Fixed.Size {
					m.Extension = existing
					m.IsExtension = false
					existing.IsExtension = true
					existing.Extension = nil
					byType[m.TypeCode()] = m
				} else {
					m.Extension = nil
					m.IsExtension = true
					existing.IsExtension = false
					existing.Extension = m
				}
			}
			if existing.VLS != nil && m.VLS != nil {
				if existing.VLS.Size > m.VLS.Size {
					m.Extension = existing
					m.IsExtension = false
					existing.IsExtension = true
					existing.Extension = nil
					byType[m.TypeCode()] = m
				} else {
					m.Extension = nil
					m.IsExtension = true
					existing.IsExtension = false
					existing.Extension = m
				}
			}
		} else {
			byType[m.TypeCode()] = m
		}
	}

	return generator, nil
}

func (g *Generator) Run() error {
	_ = os.MkdirAll(g.config.Dir, 0755)
	files, err := ioutil.ReadDir(g.config.Dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		n := file.Name()
		if err = os.Remove(filepath.Join(g.config.Dir, n)); err != nil {
			return err
		}
	}
	if err := g.generateAliases(); err != nil {
		return err
	}
	if err := g.generateConstants(); err != nil {
		return err
	}
	if err := g.generateEnums(); err != nil {
		return err
	}
	if err := g.generateStructs(); err != nil {
		return err
	}
	if err := g.generateMod(); err != nil {
		return err
	}
	if err := g.generateFactory(); err != nil {
		return err
	}
	if err := g.generateHandler(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) Overwrite() error {
	_ = os.MkdirAll(g.config.Dir, 0755)
	files, err := ioutil.ReadDir(g.config.Dir)
	if err != nil {
		return err
	}
	_ = files
	if err := g.generateAliases(); err != nil {
		return err
	}
	if err := g.generateConstants(); err != nil {
		return err
	}
	if err := g.generateEnums(); err != nil {
		return err
	}
	if err := g.generateStructs(); err != nil {
		return err
	}
	if err := g.generateMod(); err != nil {
		return err
	}
	if err := g.generateFactory(); err != nil {
		return err
	}
	if err := g.generateHandler(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) writeFile(name string, data []byte) error {
	path := filepath.Join(g.config.Dir, name)
	if err := os.WriteFile(path, data, 0755); err != nil {
		return err
	}
	return nil
}

func toCamelCase(name string) string {
	out := make([]byte, 0, len(name))
	last := byte(0)
	for i := 0; i < len(name); i++ {
		if i == 0 {
			out = append(out, toUpper(name[i]))
			last = name[0]
			continue
		}
		if name[i] != '_' {
			if last == '_' {
				out = append(out, toUpper(name[i]))
			} else {
				out = append(out, toLower(name[i]))
			}
		}

		last = name[i]
	}
	return string(out)
}

func toUpper(b byte) byte {
	switch b {
	case 'a':
		return 'A'
	case 'b':
		return 'B'
	case 'c':
		return 'C'
	case 'd':
		return 'D'
	case 'e':
		return 'E'
	case 'f':
		return 'F'
	case 'g':
		return 'G'
	case 'h':
		return 'H'
	case 'i':
		return 'I'
	case 'j':
		return 'J'
	case 'k':
		return 'K'
	case 'l':
		return 'L'
	case 'm':
		return 'M'
	case 'n':
		return 'N'
	case 'o':
		return 'O'
	case 'p':
		return 'P'
	case 'q':
		return 'Q'
	case 'r':
		return 'R'
	case 's':
		return 'S'
	case 't':
		return 'T'
	case 'u':
		return 'U'
	case 'v':
		return 'V'
	case 'w':
		return 'W'
	case 'x':
		return 'X'
	case 'y':
		return 'Y'
	case 'z':
		return 'Z'
	}
	return b
}

func toLower(b byte) byte {
	switch b {
	case 'A':
		return 'a'
	case 'B':
		return 'b'
	case 'C':
		return 'c'
	case 'D':
		return 'd'
	case 'E':
		return 'e'
	case 'F':
		return 'f'
	case 'G':
		return 'g'
	case 'H':
		return 'h'
	case 'I':
		return 'i'
	case 'J':
		return 'j'
	case 'K':
		return 'k'
	case 'L':
		return 'l'
	case 'M':
		return 'm'
	case 'N':
		return 'n'
	case 'O':
		return 'o'
	case 'P':
		return 'p'
	case 'Q':
		return 'q'
	case 'R':
		return 'r'
	case 'S':
		return 's'
	case 'T':
		return 't'
	case 'U':
		return 'u'
	case 'V':
		return 'v'
	case 'W':
		return 'w'
	case 'X':
		return 'x'
	case 'Y':
		return 'y'
	case 'Z':
		return 'z'
	}
	return b
}

func toSnakeCase(name string) string {
	out := make([]byte, 0, len(name)*2)

	// NewOCOOrder
	// TradeOrder
	for i := 0; i < len(name); i++ {
		var (
			isUpper     = string(name[i]) != strings.ToLower(string(name[i]))
			nextIsLower = false
			prevIsLower = false
		)
		if i > 0 {
			prevIsLower = string(name[i-1]) == strings.ToLower(string(name[i-1]))
		}
		if i < len(name)-1 {
			nextIsLower = strings.ToLower(string(name[i+1])) == string(name[i+1])
		}
		if name[i] == '_' {
			continue
		}
		if isUpper {
			if prevIsLower || nextIsLower {
				if len(out) > 0 {
					out = append(out, '_')
				}
			}
		}

		out = append(out, strings.ToLower(string(name[i]))...)
	}

	return string(out)
}

func valueToRust(v *schema.Value) string {
	switch v.Type {
	case schema.ValueTypeBool:
		if v.Int == 0 {
			return "false"
		} else {
			return "true"
		}
	case schema.ValueTypeInt:
		return strconv.FormatInt(v.Int, 10)
	case schema.ValueTypeUint:
		return strconv.FormatUint(v.Uint, 10)
	case schema.ValueTypeFloat:
		return strconv.FormatFloat(v.Float64, 'g', 10, 64)
	case schema.ValueTypeFloat32Max:
		return "f32::MAX"
	case schema.ValueTypeFloat64Max:
		return "f64::MAX"
	case schema.ValueTypeString:
		return fmt.Sprintf("\"%s\"", v.Str)
	case schema.ValueTypeConst:
		return cleanFieldName(v.Const.Name)
	case schema.ValueTypeEnumOption:
		return cleanFieldName(v.EnumOption.Name)
	case schema.ValueTypeSizeof:
		return fmt.Sprintf("%d", v.Struct.Size)
	}
	return ""
}

func primitiveTypeName(kind schema.Kind) string {
	switch kind {
	case schema.KindBool:
		return "bool"
	case schema.KindInt8:
		return "i8"
	case schema.KindInt16:
		return "i16"
	case schema.KindInt32:
		return "i32"
	case schema.KindInt64:
		return "i64"
	case schema.KindUint8:
		return "u8"
	case schema.KindUint16:
		return "u16"
	case schema.KindUint32:
		return "u32"
	case schema.KindUint64:
		return "u64"
	case schema.KindFloat32:
		return "f32"
	case schema.KindFloat64:
		return "f64"
	case schema.KindStringFixed, schema.KindStringVLS:
		return "&str"
	default:
		return "invalid"
	}
}

func (g *Generator) PublicType(t *schema.Type) string {
	switch t.Kind {
	case schema.KindBool:
		return "bool"
	case schema.KindInt8:
		return "i8"
	case schema.KindInt16:
		return "i16"
	case schema.KindInt32:
		return "i32"
	case schema.KindInt64:
		return "i64"
	case schema.KindUint8:
		return "u8"
	case schema.KindUint16:
		return "u16"
	case schema.KindUint32:
		return "u32"
	case schema.KindUint64:
		return "u64"
	case schema.KindFloat32:
		return "f32"
	case schema.KindFloat64:
		return "f64"
	case schema.KindStringFixed, schema.KindStringVLS:
		return "&str"
	case schema.KindEnum:
		return cleanStructName(t.Enum.Name)
	case schema.KindAlias:
		return cleanStructName(t.Alias.Name)
	default:
		return "invalid"
	}
}

func (g *Generator) RustTypeName(t *schema.Type) string {
	switch t.Kind {
	case schema.KindBool:
		return "bool"
	case schema.KindInt8:
		return "i8"
	case schema.KindInt16:
		return "i16"
	case schema.KindInt32:
		return "i32"
	case schema.KindInt64:
		return "i64"
	case schema.KindUint8:
		return "u8"
	case schema.KindUint16:
		return "u16"
	case schema.KindUint32:
		return "u32"
	case schema.KindUint64:
		return "u64"
	case schema.KindFloat32:
		return "f32"
	case schema.KindFloat64:
		return "f64"
	case schema.KindStringFixed:
		return fmt.Sprintf("[u8; %d]", t.Size)
	case schema.KindStringVLS:
		return "VLS"
	case schema.KindEnum:
		return cleanStructName(t.Enum.Name)
	case schema.KindAlias:
		return cleanStructName(t.Alias.Name)
	default:
		return "invalid"
	}
}

func initValue(field *schema.Field) string {
	if field.Type.Union != nil {
		largest := field.Type.Union.Fields[0]
		for _, f := range field.Type.Union.Fields {
			if f.Type.Size > largest.Type.Size {
				largest = f
			}
		}
		return initValue(largest)
	}
	value := field.Initial

	t := &field.Type
	if t.Kind == schema.KindAlias {
		t = &t.Alias.Type
	}

	if value != nil {
		switch value.Type {
		case schema.ValueTypeUnknown:
			panic("unknown type")
		case schema.ValueTypeInt:
			if t.Kind == schema.KindFloat32 || t.Kind == schema.KindFloat64 {
				return fmt.Sprintf("%d.0", value.Uint)
			}
			if t.Kind == schema.KindBool {
				if value.Uint != 0 || value.Int != 0 {
					return "true"
				} else {
					return "false"
				}
			}
			switch value.Int {
			case math.MaxInt8:
				if t.Kind != schema.KindInt8 {
					return fmt.Sprintf("(i8::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "i8::MAX"
			case math.MaxInt16:
				if t.Kind != schema.KindInt16 {
					return fmt.Sprintf("(i16::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "i16::MAX.to_le()"
			case math.MaxInt32:
				if t.Kind != schema.KindInt32 {
					return fmt.Sprintf("(i32::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "i32::MAX.to_le()"
			case math.MaxInt64:
				if t.Kind != schema.KindInt64 {
					return fmt.Sprintf("(i64::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "i64::MAX.to_le()"
			default:
				switch t.Kind {
				case schema.KindInt8:
					return fmt.Sprintf("%d", value.Int)
				case schema.KindInt16:
					return fmt.Sprintf("%di16.to_le()", value.Int)
				case schema.KindInt32:
					return fmt.Sprintf("%di32.to_le()", value.Int)
				case schema.KindInt64:
					return fmt.Sprintf("%di64.to_le()", value.Int)
				case schema.KindUint8:
					return fmt.Sprintf("%d", value.Int)
				case schema.KindUint16:
					return fmt.Sprintf("%du16.to_le()", value.Int)
				case schema.KindUint32:
					return fmt.Sprintf("%du32.to_le()", value.Int)
				case schema.KindUint64:
					return fmt.Sprintf("%du64.to_le()", value.Int)
				}
			}
			return fmt.Sprintf("%d", value.Int)
		case schema.ValueTypeUint:
			if t.Kind == schema.KindFloat32 || t.Kind == schema.KindFloat64 {
				return fmt.Sprintf("%d.0", value.Uint)
			}
			if t.Kind == schema.KindBool {
				if value.Uint != 0 || value.Int != 0 {
					return "true"
				} else {
					return "false"
				}
			}
			switch value.Uint {
			case math.MaxUint8:
				if t.Kind != schema.KindUint8 {
					return fmt.Sprintf("(u8::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "u8::MAX"
			case math.MaxUint16:
				if t.Kind != schema.KindUint16 {
					return fmt.Sprintf("(u16::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "u16::MAX.to_le()"
			case math.MaxUint32:
				if t.Kind != schema.KindUint32 {
					return fmt.Sprintf("(u32::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "u32::MAX.to_le()"
			case math.MaxUint64:
				if t.Kind != schema.KindUint64 {
					return fmt.Sprintf("(u64)::MAX as %s).to_le()", primitiveTypeName(t.Kind))
				}
				return "u64::MAX.to_le()"
			default:
				switch t.Kind {
				case schema.KindInt8:
					return fmt.Sprintf("%d", value.Uint)
				case schema.KindInt16:
					return fmt.Sprintf("%di16.to_le()", value.Uint)
				case schema.KindInt32:
					return fmt.Sprintf("%di32.to_le()", value.Uint)
				case schema.KindInt64:
					return fmt.Sprintf("%di64.to_le()", value.Uint)
				case schema.KindUint8:
					return fmt.Sprintf("%d", value.Int)
				case schema.KindUint16:
					return fmt.Sprintf("%du16.to_le()", value.Uint)
				case schema.KindUint32:
					return fmt.Sprintf("%du32.to_le()", value.Uint)
				case schema.KindUint64:
					return fmt.Sprintf("%du64.to_le()", value.Uint)
				}
			}
			return fmt.Sprintf("%d", value.Uint)
		case schema.ValueTypeBool:
			if value.Int > 0 || value.Uint > 0 {
				if t.Kind != schema.KindBool {
					return "1"
				} else {
					return "true"
				}
			} else {
				if t.Kind != schema.KindBool {
					return "0"
				} else {
					return "false"
				}
			}
		case schema.ValueTypeFloat:
			if value.Float64 == 0.0 {
				return "0.0"
			} else {
				v := strconv.FormatFloat(value.Float64, 'g', 4, 64)
				if strings.IndexByte(v, '.') == -1 {
					return fmt.Sprintf("%s.0", v)
				}
				return v
			}
		case schema.ValueTypeFloat32Max:
			return "f32_le(f32::MAX)"
		case schema.ValueTypeFloat64Max:
			return "f64_le(f64::MAX)"
		case schema.ValueTypeString:
			if field.Struct.VLS {
				panic("default VLS string")
			}

			b := make([]byte, 0, 128)
			b = append(b, '[')
			for i := 0; i < field.Type.Size; i++ {
				if i < len(value.Str) {
					b = strconv.AppendInt(b, int64(value.Str[i]), 10)
				} else {
					b = strconv.AppendInt(b, int64(0), 10)
				}
				if i < field.Type.Size-1 {
					b = append(b, ',')
				}
			}
			b = append(b, ']')

			return string(b)

		case schema.ValueTypeConst:
			// return value.Const.Name
			// switch field.Type.Kind {
			// case schema.KindInt8:
			// 	return fmt.Sprintf("%d", value.Const.Value.Int)
			// case schema.KindInt16:
			// 	return fmt.Sprintf("%di16.to_le()", value.Const.Value.Int)
			// case schema.KindInt32:
			// 	return fmt.Sprintf("%di32.to_le()", value.Const.Value.Int)
			// case schema.KindInt64:
			// 	return fmt.Sprintf("%di64.to_le()", value.Const.Value.Int)
			// case schema.KindUint8:
			// 	return fmt.Sprintf("%d", value.Const.Value.Int)
			// case schema.KindUint16:
			// 	return fmt.Sprintf("%du16.to_le()", value.Const.Value.Int)
			// case schema.KindUint32:
			// 	return fmt.Sprintf("%du32.to_le()", value.Const.Value.Int)
			// case schema.KindUint64:
			// 	return fmt.Sprintf("%du64.to_le()", value.Const.Value.Int)
			// }

			return fmt.Sprintf("%s.to_le()", value.Const.Name)
		case schema.ValueTypeEnumOption:
			return fmt.Sprintf("%s::%s.to_le()", cleanStructName(value.EnumOption.Enum.Name), toCamelCase(value.EnumOption.Name))
		case schema.ValueTypeSizeof:
			return fmt.Sprintf("%du16.to_le()", field.Struct.Size)
		}

		return defaultValue(t)
	} else {
		return defaultValue(t)
	}
}

func defaultValue(t *schema.Type) string {
	if t.Kind == schema.KindAlias {
		return defaultValue(&t.Alias.Type)
	} else if t.Kind == schema.KindUnion {
		largest := t.Union.Fields[0]
		for _, f := range t.Union.Fields {
			if f.Type.Size > largest.Type.Size {
				largest = f
			}
		}
		return defaultValue(&largest.Type)
	}
	switch t.Kind {
	case schema.KindBool:
		return "false"
	case schema.KindInt8:
		return "0i8"
	case schema.KindInt16:
		return "0i16"
	case schema.KindInt32:
		return "0i32"
	case schema.KindInt64:
		return "0i64"
	case schema.KindUint8:
		return "0u8"
	case schema.KindUint16:
		return "0u16"
	case schema.KindUint32:
		return "0u32"
	case schema.KindUint64:
		return "0u64"
	case schema.KindFloat32:
		return "0.0f32"
	case schema.KindFloat64:
		return "0.0f64"
	case schema.KindStringFixed:
		return fmt.Sprintf("[0; %d]", t.Size)
	case schema.KindStringVLS:
		return "crate::message::VLS::new()"
	case schema.KindEnum:
		return fmt.Sprintf("%s::%s.to_le()", cleanStructName(t.Enum.Name), toCamelCase(t.Enum.Options[0].Name))
	default:
		return "invalid"
	}
}
