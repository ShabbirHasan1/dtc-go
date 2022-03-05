package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen"
	"os"
	"path/filepath"
	"strconv"
)

type Generator struct {
	config          *Config
	schema          *codegen.Schema
	root            *Package
	fixed           *Package
	vls             *Package
	nonStandard     *Package
	serialize       *Package
	factory         *Package
	all             map[string]interface{}
	packages        map[string]*Package
	aliases         []*Alias
	aliasesByName   map[string]*Alias
	constants       []*Constant
	constantsByName map[string]*Constant
	enums           []*Enum
	enumsByName     map[string]*Enum
	messages        []*Message
	messagesByName  map[string]*Message
}

type Config struct {
	Dir                string
	RootPackage        string
	FixedPackage       string
	VLSPackage         string
	NonStandardPackage string
	SerializePackage   string
	FactoryPackage     string
}

func DefaultConfig(dir string) *Config {
	return &Config{
		Dir:                dir,
		RootPackage:        "github.com/moontrade/dtc-go/model",
		FixedPackage:       "fixed",
		VLSPackage:         "vls",
		NonStandardPackage: "ns",
		SerializePackage:   "serialize",
		FactoryPackage:     "factory",
	}
}

func NewPackage(name, path string) *Package {
	return &Package{
		Name:            name,
		Path:            path,
		EnumsByName:     make(map[string]*Enum),
		ConstantsByName: make(map[string]*Constant),
		AliasesByMame:   make(map[string]*Alias),
		StructsByName:   make(map[string]*Struct),
	}
}

func (g *Generator) typeOf(typ codegen.Type) (*Type, error) {
	switch typ.Kind {
	case codegen.KindUnknown:
	case codegen.KindInt8:
		return &Type{
			Import:    nil,
			Name:      "int8",
			DTC:       typ,
			Kind:      Int8,
			Primitive: true,
		}, nil
	case codegen.KindUint8:
		return &Type{
			Import:    nil,
			Name:      "uint8",
			DTC:       typ,
			Kind:      Uint8,
			Primitive: true,
		}, nil
	case codegen.KindInt16:
		return &Type{
			Import:    nil,
			Name:      "int16",
			DTC:       typ,
			Kind:      Int16,
			Primitive: true,
		}, nil
	case codegen.KindUint16:
		return &Type{
			Import:    nil,
			Name:      "uint16",
			DTC:       typ,
			Kind:      Uint16,
			Primitive: true,
		}, nil
	case codegen.KindInt32:
		return &Type{
			Import:    nil,
			Name:      "int32",
			DTC:       typ,
			Kind:      Int32,
			Primitive: true,
		}, nil
	case codegen.KindUint32:
		return &Type{
			Import:    nil,
			Name:      "uint32",
			DTC:       typ,
			Kind:      Uint32,
			Primitive: true,
		}, nil
	case codegen.KindInt64:
		return &Type{
			Import:    nil,
			Name:      "int64",
			DTC:       typ,
			Kind:      Int64,
			Primitive: true,
		}, nil
	case codegen.KindUint64:
		return &Type{
			Import:    nil,
			Name:      "uint64",
			DTC:       typ,
			Kind:      Uint64,
			Primitive: true,
		}, nil
	case codegen.KindFloat32:
		return &Type{
			Import:    nil,
			Name:      "float32",
			DTC:       typ,
			Kind:      Float32,
			Primitive: true,
		}, nil
	case codegen.KindFloat64:
		return &Type{
			Import:    nil,
			Name:      "float64",
			DTC:       typ,
			Kind:      Float64,
			Primitive: true,
		}, nil
	case codegen.KindStringFixed:
		return &Type{
			Import:    nil,
			Name:      "string",
			DTC:       typ,
			Kind:      String,
			Primitive: true,
		}, nil
	case codegen.KindStringVLS:
		return &Type{
			Import:    nil,
			Name:      "string",
			DTC:       typ,
			Kind:      String,
			Primitive: true,
		}, nil
	case codegen.KindAlias:
		underlying := primitiveKind(typ.Alias.Type.Kind)
		if underlying == Invalid {
			return nil, fmt.Errorf("alias kind is not primitive: %d", typ.Alias.Type.Kind)
		}
		return &Type{
			Import: &Import{
				Alias: g.root.Name,
				Name:  cleanName(typ.Alias.Name),
				Path:  g.root.Path,
			},
			Name: cleanName(typ.Alias.Name),
			DTC:  typ,
			Kind: underlying,
		}, nil
	case codegen.KindEnum:
		underlying := primitiveKind(typ.Kind)
		return &Type{
			Import: &Import{
				Alias: g.root.Name,
				Name:  cleanName(typ.Enum.Name),
				Path:  g.root.Path,
			},
			Name: cleanName(typ.Enum.Name),
			DTC:  typ,
			Kind: underlying,
		}, nil

	case codegen.KindUnion:
	case codegen.KindStruct:
	}
	return nil, fmt.Errorf("invalid kind %d", typ.Kind)
}

func NewGenerator(config *Config, schema *codegen.Schema) (*Generator, error) {
	generator := &Generator{
		config: config,
		schema: schema,
		root:   NewPackage(filepath.Base(config.RootPackage), config.RootPackage),
		fixed: NewPackage(config.FixedPackage,
			fmt.Sprintf("%s/%s", config.RootPackage, config.FixedPackage)),
		vls: NewPackage(config.VLSPackage,
			fmt.Sprintf("%s/%s", config.RootPackage, config.VLSPackage)),
		nonStandard: NewPackage(config.NonStandardPackage,
			fmt.Sprintf("%s/%s", config.RootPackage, config.NonStandardPackage)),
		serialize: NewPackage(config.SerializePackage,
			fmt.Sprintf("%s/%s", config.RootPackage, config.SerializePackage)),
		factory: NewPackage(config.FactoryPackage,
			fmt.Sprintf("%s/%s", config.RootPackage, config.FactoryPackage)),
		all:             make(map[string]interface{}),
		packages:        make(map[string]*Package),
		aliasesByName:   make(map[string]*Alias),
		messagesByName:  make(map[string]*Message),
		enumsByName:     make(map[string]*Enum),
		constantsByName: make(map[string]*Constant),
	}

	for _, alias := range schema.Aliases {
		a := &Alias{
			Package: generator.root,
			Alias:   alias,
			Name:    cleanName(alias.Name),
			Kind:    primitiveKind(alias.Type.Kind),
		}
		if a.Kind == Invalid {
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

	for _, constant := range schema.Constants {
		c := &Constant{
			Package: generator.root,
			Const:   constant,
			Name:    constant.Name,
			Kind:    primitiveKind(constant.Type.Kind),
		}
		if c.Kind == Invalid {
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
	for _, enum := range schema.Enums {
		e := &Enum{
			Package:       generator.root,
			Enum:          enum,
			Name:          cleanName(enum.Name),
			Type:          primitiveKind(enum.Type),
			OptionsByName: make(map[string]*EnumOption),
		}
		if e.Type == Invalid {
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

	createStruct := func(st *codegen.Struct) (*Struct, error) {
		if st == nil {
			return nil, nil
		}
		s := &Struct{
			Struct:       st,
			Name:         cleanName(st.Name),
			FieldsByName: make(map[string]*Field),
		}
		for _, field := range st.Fields {
			var (
				f = &Field{
					Struct: s,
					Field:  field,
					Name:   cleanName(field.Name),
				}
				err error
			)
			if field.Type.Union != nil {
				for _, ufield := range field.Type.Union.Fields {
					ff := &Field{
						Struct: s,
						Field:  ufield,
						Name:   cleanName(ufield.Name),
					}
					if ff.Type, err = generator.typeOf(ufield.Type); err != nil {
						return nil, err
					}
					if s.FieldsByName[ff.Name] != nil {
						return nil, fmt.Errorf("%s.%s duplicate field name after cleanName", st.Name, field.Name)
					}
					f.Fields = append(f.Fields, ff)
					s.FieldsByName[ff.Name] = ff
				}
			} else {
				if f.Type, err = generator.typeOf(field.Type); err != nil {
					return nil, err
				}
				if s.FieldsByName[f.Name] != nil {
					return nil, fmt.Errorf("%s.%s duplicate field name after cleanName", st.Name, field.Name)
				}
				s.FieldsByName[f.Name] = f
			}
			s.Fields = append(s.Fields, f)
		}
		return s, nil
	}
	for _, message := range schema.Messages {
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
		} else if m.VLS != nil {
			name = m.VLS.Name
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

func (g *Generator) Run() error {
	if err := g.generateConstants(); err != nil {
		return err
	}
	if err := g.generateEnums(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) generateConstants() error {
	w := &Writer{}
	w.W("package %s", g.root.Name)
	w.W("")

	importMath := false
	for _, constant := range g.constants {
		switch constant.Const.Value.Type {
		case codegen.ValueTypeFloat32Max, codegen.ValueTypeFloat64Max:
			importMath = true
		}
	}
	if importMath {
		w.W("import \"math\"")
		w.W("")
	}

	w.W("const (")
	for _, constant := range g.constants {
		w.Indent(1, "%s %s = %s", constant.Name, constant.Kind.String(), valueToGo(constant.Const.Value))
	}
	w.W(")")
	w.W("")

	return os.WriteFile(filepath.Join(g.config.Dir, "constants.go"), w.b, 0755)
}

func (g *Generator) generateEnums() error {
	w := &Writer{}
	w.W("package %s", g.root.Name)
	w.W("")

	for _, enum := range g.enums {
		w.W("type %s %s", enum.Name, enum.Type.String())
		w.W("")
		w.W("const (")
		for _, option := range enum.Options {
			w.Indent(1, "%s %s = %d", option.Name, enum.Name, option.Value)
		}
		w.W(")")
		w.W("")
	}

	return os.WriteFile(filepath.Join(g.config.Dir, "enums.go"), w.b, 0755)
}

func (g *Generator) generateFixed() error {
	return nil
}
func (g *Generator) generateVLS() error {
	return nil
}
func (g *Generator) generateSerialize() error {
	return nil
}
func (g *Generator) generateFactory() error {
	return nil
}

func valueToGo(v *codegen.Value) string {
	switch v.Type {
	case codegen.ValueTypeBool:
		if v.Int == 0 {
			return "false"
		} else {
			return "true"
		}
	case codegen.ValueTypeInt:
		return strconv.FormatInt(v.Int, 10)
	case codegen.ValueTypeUint:
		return strconv.FormatUint(v.Uint, 10)
	case codegen.ValueTypeFloat:
		return strconv.FormatFloat(v.Float64, 'g', 10, 64)
	case codegen.ValueTypeFloat32Max:
		return "math.MaxFloat32"
	case codegen.ValueTypeFloat64Max:
		return "math.MaxFloat64"
	case codegen.ValueTypeString:
		return fmt.Sprintf("\"%s\"", v.Str)
	case codegen.ValueTypeConst:
		return cleanName(v.Const.Name)
	case codegen.ValueTypeEnumOption:
		return cleanName(v.EnumOption.Name)
	case codegen.ValueTypeSizeof:
		return fmt.Sprintf("%d", v.Struct.Size)
	}
	return ""
}

type Writer struct {
	b []byte
}

func (w *Writer) Indent(times int, format string, args ...interface{}) {
	for i := 0; i < times; i++ {
		w.b = append(w.b, "    "...)
		w.b = append(w.b, fmt.Sprintf(format, args...)...)
		w.b = append(w.b, "\n"...)
	}
}

func (w *Writer) W(format string, args ...interface{}) {
	w.b = append(w.b, fmt.Sprintf(format, args...)...)
	w.b = append(w.b, "\n"...)
}
