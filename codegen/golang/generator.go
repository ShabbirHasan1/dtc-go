package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

//func (g *Generator) typeOf(typ codegen.Type) (*Type, error) {
//	switch typ.Kind {
//	case codegen.KindUnknown:
//	case codegen.KindInt8:
//		return &Type{
//			Import:    nil,
//			Name:      "int8",
//			DTC:       typ,
//			Primitive: true,
//		}, nil
//	case codegen.KindUint8:
//		return &Type{
//			Import:    nil,
//			Name:      "uint8",
//			DTC:       typ,
//			Primitive: true,
//		}, nil
//	case codegen.KindInt16:
//		return &Type{
//			Import:    nil,
//			Name:      "int16",
//			DTC:       typ,
//			Kind:      Int16,
//			Primitive: true,
//		}, nil
//	case codegen.KindUint16:
//		return &Type{
//			Import:    nil,
//			Name:      "uint16",
//			DTC:       typ,
//			Kind:      Uint16,
//			Primitive: true,
//		}, nil
//	case codegen.KindInt32:
//		return &Type{
//			Import:    nil,
//			Name:      "int32",
//			DTC:       typ,
//			Kind:      Int32,
//			Primitive: true,
//		}, nil
//	case codegen.KindUint32:
//		return &Type{
//			Import:    nil,
//			Name:      "uint32",
//			DTC:       typ,
//			Kind:      Uint32,
//			Primitive: true,
//		}, nil
//	case codegen.KindInt64:
//		return &Type{
//			Import:    nil,
//			Name:      "int64",
//			DTC:       typ,
//			Kind:      Int64,
//			Primitive: true,
//		}, nil
//	case codegen.KindUint64:
//		return &Type{
//			Import:    nil,
//			Name:      "uint64",
//			DTC:       typ,
//			Kind:      Uint64,
//			Primitive: true,
//		}, nil
//	case codegen.KindFloat32:
//		return &Type{
//			Import:    nil,
//			Name:      "float32",
//			DTC:       typ,
//			Kind:      Float32,
//			Primitive: true,
//		}, nil
//	case codegen.KindFloat64:
//		return &Type{
//			Import:    nil,
//			Name:      "float64",
//			DTC:       typ,
//			Kind:      Float64,
//			Primitive: true,
//		}, nil
//	case codegen.KindStringFixed:
//		return &Type{
//			Import:    nil,
//			Name:      "string",
//			DTC:       typ,
//			Kind:      String,
//			Primitive: true,
//		}, nil
//	case codegen.KindStringVLS:
//		return &Type{
//			Import:    nil,
//			Name:      "string",
//			DTC:       typ,
//			Kind:      String,
//			Primitive: true,
//		}, nil
//	case codegen.KindAlias:
//		underlying := primitiveKind(typ.Alias.Type.Kind)
//		if underlying == Invalid {
//			return nil, fmt.Errorf("alias kind is not primitive: %d", typ.Alias.Type.Kind)
//		}
//		return &Type{
//			Import: &Import{
//				Alias: g.root.Name,
//				Name:  cleanName(typ.Alias.Name),
//				Path:  g.root.Path,
//			},
//			Name: cleanName(typ.Alias.Name),
//			DTC:  typ,
//			Kind: underlying,
//		}, nil
//	case codegen.KindEnum:
//		underlying := primitiveKind(typ.Kind)
//		return &Type{
//			Import: &Import{
//				Alias: g.root.Name,
//				Name:  cleanName(typ.Enum.Name),
//				Path:  g.root.Path,
//			},
//			Name: cleanName(typ.Enum.Name),
//			DTC:  typ,
//			Kind: underlying,
//		}, nil
//
//	case codegen.KindUnion:
//	case codegen.KindStruct:
//	}
//	return nil, fmt.Errorf("invalid kind %d", typ.Kind)
//}

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

	for _, constant := range schema.Constants {
		c := &Constant{
			Package: generator.root,
			Const:   constant,
			Name:    constant.Name,
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
	for _, enum := range schema.Enums {
		e := &Enum{
			Package:       generator.root,
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
			)
			if field.Type.Union != nil {
				for _, field := range field.Type.Union.Fields {
					ff := &Field{
						Struct: s,
						Field:  field,
						Name:   cleanName(field.Name),
					}
					if s.FieldsByName[ff.Name] != nil {
						return nil, fmt.Errorf("%s.%s duplicate field name after cleanName", st.Name, field.Name)
					}
					f.Fields = append(f.Fields, ff)
					s.FieldsByName[ff.Name] = ff
				}
			} else {
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
	if err := g.generateAliases(); err != nil {
		return err
	}
	if err := g.generateConstants(); err != nil {
		return err
	}
	if err := g.generateEnums(); err != nil {
		return err
	}
	if err := g.generateMessages(); err != nil {
		return err
	}
	if err := g.generateSerialize(); err != nil {
		return err
	}
	if err := g.generateFixed(); err != nil {
		return err
	}
	if err := g.generateVLS(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) generateAliases() error {
	w := &Writer{}
	w.Line("package %s", g.root.Name)
	w.Line("")

	w.Line("type (")
	for _, alias := range g.aliases {
		w.IndentLine(1, "%s %s", alias.Name, g.GoTypeName(&alias.Type, nil))
	}
	w.Line(")")
	w.Line("")

	return os.WriteFile(filepath.Join(g.config.Dir, "alias.go"), w.b, 0755)
}

func (g *Generator) generateConstants() error {
	w := &Writer{}
	w.Line("package %s", g.root.Name)
	w.Line("")

	importMath := false
	for _, constant := range g.constants {
		switch constant.Const.Value.Type {
		case codegen.ValueTypeFloat32Max, codegen.ValueTypeFloat64Max:
			importMath = true
		}
	}
	if importMath {
		w.Line("import \"math\"")
		w.Line("")
	}

	w.Line("const (")
	for _, constant := range g.constants {
		w.IndentLine(1, "%s %s = %s", constant.Name, g.GoTypeName(&constant.Type, nil), valueToGo(constant.Const.Value))
	}
	w.Line(")")
	w.Line("")

	return os.WriteFile(filepath.Join(g.config.Dir, "constants.go"), w.b, 0755)
}

func (g *Generator) generateEnums() error {
	w := &Writer{}
	w.Line("package %s", g.root.Name)
	w.Line("")

	for _, enum := range g.enums {
		w.Line("type %s %s", enum.Name, primitiveKindName(enum.Type))
		w.Line("")
		w.Line("const (")
		for _, option := range enum.Options {
			w.IndentLine(1, "%s %s = %d", option.Name, enum.Name, option.Value)
		}
		w.Line(")")
		w.Line("")
	}

	return os.WriteFile(filepath.Join(g.config.Dir, "enums.go"), w.b, 0755)
}

func (g *Generator) generateMessages() error {
	for _, model := range g.messages {
		msg := model.Fixed
		if msg == nil {
			msg = model.VLS
		}
		if msg == nil {
			continue
		}
		w := &Writer{}
		w.Line("package %s", g.root.Name)
		w.Line("")
		w.Line("import (")
		w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
		w.Line(")")
		w.Line("")

		w.Line("type %s interface {", msg.Name)
		w.IndentLine(1, "message.Message")
		w.Line("")
		w.IndentLine(1, "ToBuilder() %sBuilder", msg.Name)
		w.Line("")

		for i, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}

			if field.Fields != nil {
				for _, field := range field.Fields {
					w.IndentLine(1, "%s() %s", field.Name, g.GoTypeName(&field.Type, nil))
				}
			} else {
				w.IndentLine(1, "%s() %s", field.Name, g.GoTypeName(&field.Type, nil))
			}

			if i < len(msg.Fields)-1 {
				w.Line("")
			}
		}
		w.Line("}")
		w.Line("")

		w.Line("type %sBuilder interface {", msg.Name)
		w.IndentLine(1, "message.Builder")
		w.Line("")
		w.IndentLine(1, "%s", msg.Name)
		w.Line("")

		for i, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}

			if field.Fields != nil {
				for _, field := range field.Fields {
					w.IndentLine(1, "Set%s(value %s)", field.Name, g.GoTypeName(&field.Type, nil))
				}
			} else {
				w.IndentLine(1, "Set%s(value %s)", field.Name, g.GoTypeName(&field.Type, nil))
			}

			if i < len(msg.Fields)-1 {
				w.Line("")
			}
		}
		w.Line("}")
		w.Line("")

		w.Line("type %sFactory interface {", msg.Name)
		w.IndentLine(1, "New() %sBuilder", msg.Name)
		w.Line("")
		w.IndentLine(1, "Alloc() %sBuilder", msg.Name)
		w.Line("}")
		w.Line("")

		if err := os.WriteFile(filepath.Join(g.config.Dir, fmt.Sprintf("%s.go", toSnakeCase(msg.Name))), w.b, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateFixed() error {
	for _, model := range g.messages {
		msg := model.Fixed
		if msg == nil {
			continue
		}
		//importMath := false
		//for _, field := range msg.Fields {
		//	if len(field.Fields) > 0 {
		//		for _, field := range field.Fields {
		//			if field.Initial != nil {
		//				switch field.Initial.Type {
		//				case codegen.ValueTypeFloat32Max, codegen.ValueTypeFloat64Max:
		//					importMath = true
		//				}
		//			}
		//		}
		//	} else if field.Initial != nil {
		//		switch field.Initial.Type {
		//		case codegen.ValueTypeFloat32Max, codegen.ValueTypeFloat64Max:
		//			importMath = true
		//		}
		//	}
		//}

		w := &Writer{}
		w.Line("package %s", g.fixed.Name)
		w.Line("")
		w.Line("import (")
		w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
		w.IndentLine(1, "\"%s\"", g.root.Path)
		//if importMath {
		//	w.IndentLine(1, "\"math\"")
		//}
		w.Line(")")
		w.Line("")

		w.Line("var (")
		w.IndentLine(1, "_ model.%s = %s{}", msg.Name, msg.Name)
		w.IndentLine(1, "_ model.%sBuilder = %sBuilder{}", msg.Name, msg.Name)
		w.IndentLine(1, "_ model.%s = %sPointer{}", msg.Name, msg.Name)
		w.IndentLine(1, "_ model.%sBuilder = %sPointerBuilder{}", msg.Name, msg.Name)
		w.Line(")")
		w.Line("")
		w.Write("var _%sDefault = ", msg.Name)
		w.WriteByteSlice(msg.Init)
		w.Write("\n\n")

		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("// Types")
		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("")
		w.Line("type %s struct {", msg.Name)
		w.IndentLine(1, "message.Fixed")
		w.Line("}")
		w.Line("")

		w.Line("type %sBuilder struct {", msg.Name)
		w.IndentLine(1, "message.Fixed")
		w.Line("}")
		w.Line("")

		w.Line("type %sPointer struct {", msg.Name)
		w.IndentLine(1, "message.FixedPointer")
		w.Line("}")
		w.Line("")

		w.Line("type %sPointerBuilder struct {", msg.Name)
		w.IndentLine(1, "message.FixedPointer")
		w.Line("}")
		w.Line("")

		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("// Constructors")
		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("")

		w.Line("func New%s() %sBuilder {", msg.Name, msg.Name)
		w.IndentLine(1, "a := %sBuilder{message.NewFixed(%d)}", msg.Name, msg.Size)
		w.IndentLine(1, "a.Unsafe().SetBytes(0, _%sDefault)", msg.Name)
		w.IndentLine(1, "return a")
		w.Line("}")
		w.Line("")

		w.Line("func Alloc%s() %sPointerBuilder {", msg.Name, msg.Name)
		w.IndentLine(1, "a := %sPointerBuilder{message.AllocFixed(%d)}", msg.Name, msg.Size)
		w.IndentLine(1, "a.Ptr.SetBytes(0, _%sDefault)", msg.Name)
		w.IndentLine(1, "return a")
		w.Line("}")
		w.Line("")

		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("// Clear")
		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("")

		w.Line("func (e %sBuilder) Clear() {", msg.Name)
		w.IndentLine(1, "e.Unsafe().SetBytes(0, _%sDefault)", msg.Name)
		w.Line("}")
		w.Line("")

		w.Line("func (e %sPointerBuilder) Clear() {", msg.Name)
		w.IndentLine(1, "e.Ptr.SetBytes(0, _%sDefault)", msg.Name)
		w.Line("}")
		w.Line("")

		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("// ToBuilder")
		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("")

		w.Line("func (e %s) ToBuilder() model.%sBuilder {", msg.Name, msg.Name)
		w.IndentLine(1, "return %sBuilder{Fixed: e.Fixed}", msg.Name)
		w.Line("}")
		w.Line("")

		w.Line("func (e %sBuilder) ToBuilder() model.%sBuilder {", msg.Name, msg.Name)
		w.IndentLine(1, "return e")
		w.Line("}")
		w.Line("")

		w.Line("func (e %sPointer) ToBuilder() model.%sBuilder {", msg.Name, msg.Name)
		w.IndentLine(1, "return %sPointerBuilder{FixedPointer: e.FixedPointer}", msg.Name)
		w.Line("}")
		w.Line("")

		w.Line("func (e %sPointerBuilder) ToBuilder() model.%sBuilder {", msg.Name, msg.Name)
		w.IndentLine(1, "return e")
		w.Line("}")
		w.Line("")

		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("// Getters")
		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("")

		for i, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}

			g.writeFieldGetters(w, field)

			if i < len(msg.Fields)-1 {
				w.Line("")
			}
		}
		w.Line("")

		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("// Setters")
		w.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		w.Line("")

		for i, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}

			g.writeFieldSetters(w, field)

			if i < len(msg.Fields)-1 {
				w.Line("")
			}
		}
		w.Line("")

		if err := os.WriteFile(filepath.Join(g.config.Dir, g.fixed.Name, fmt.Sprintf("%s.go", toSnakeCase(msg.Name))), w.b, 0755); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateVLS() error {
	return nil
}

func (g *Generator) generateSerialize() error {
	for _, model := range g.messages {
		msg := model.Fixed
		if msg == nil {
			msg = model.VLS
		}
		if msg == nil {
			continue
		}
	}
	return nil
}

func (g *Generator) generateFactory() error {
	return nil
}

func (g *Generator) writeFieldGetters(w *Writer, f *Field) {
	if f.Fields != nil {
		for _, field := range f.Fields {
			w.Line("///////////////////////////////////////////")
			w.Line("// %s", field.Name)
			w.Line("///////////////////////////////////////////")
			w.Line("")
			g.writeFieldGetter(w, field, "", "Unsafe()")
			g.writeFieldGetter(w, field, "Builder", "Unsafe()")
			g.writeFieldGetter(w, field, "Pointer", "Ptr")
			g.writeFieldGetter(w, field, "PointerBuilder", "Ptr")
		}
	} else {
		w.Line("///////////////////////////////////////////")
		w.Line("// %s", f.Name)
		w.Line("///////////////////////////////////////////")
		w.Line("")
		g.writeFieldGetter(w, f, "", "Unsafe()")
		g.writeFieldGetter(w, f, "Builder", "Unsafe()")
		g.writeFieldGetter(w, f, "Pointer", "Ptr")
		g.writeFieldGetter(w, f, "PointerBuilder", "Ptr")
	}
}

func (g *Generator) writeFieldGetter(w *Writer, f *Field, suffix, ptrAccessor string) {
	w.Line("func (m %s%s) %s() %s {", f.Struct.Name, suffix, f.Name, g.GoTypeName(&f.Type, g.fixed))
	w.IndentLine(1, "return %s", g.fieldGetValue(f, ptrAccessor))
	w.Line("}")
	w.Line("")
}

func (g *Generator) fieldGetValue(f *Field, ptrAccessor string) string {
	if f.Field.Type.Alias != nil {
		return fmt.Sprintf("%s(%s)",
			fmt.Sprintf("%s.%s", g.root.Name, cleanName(f.Field.Type.Alias.Name)),
			primitiveGetter(
				f.Field.Type.Alias.Type.Kind,
				f.Field.Type.Offset,
				f.Field.Type.Size,
				"message",
				"m",
				ptrAccessor,
				f.Struct.VLS,
			),
		)
	}
	if f.Field.Type.Enum != nil {
		return fmt.Sprintf("%s(%s)",
			fmt.Sprintf("%s.%s", g.root.Name, cleanName(f.Field.Type.Enum.Name)),
			primitiveGetter(
				f.Field.Type.Enum.Type,
				f.Field.Type.Offset,
				f.Field.Type.Size,
				"message",
				"m",
				ptrAccessor,
				f.Struct.VLS,
			),
		)
	}

	return primitiveGetter(
		f.Field.Type.Kind,
		f.Field.Type.Offset,
		f.Field.Type.Size,
		"message",
		"m",
		ptrAccessor,
		f.Struct.VLS,
	)
}

func primitiveGetter(kind codegen.Kind, offset, size int, messagePackage, messageVar, ptrAccessor string, vls bool) string {
	var (
		bounds = offset + size
		suffix = "Fixed"
		prefix = ""
	)
	if vls {
		suffix = "VLS"
	}
	switch kind {
	case codegen.KindInt8:
		prefix = "Int8"
	case codegen.KindInt16:
		prefix = "Int16"
	case codegen.KindInt32:
		prefix = "Int32"
	case codegen.KindInt64:
		prefix = "Int64"
	case codegen.KindUint8:
		prefix = "Uint8"
	case codegen.KindUint16:
		prefix = "Uint16"
	case codegen.KindUint32:
		prefix = "Uint32"
	case codegen.KindUint64:
		prefix = "Uint64"
	case codegen.KindFloat32:
		prefix = "Float32"
	case codegen.KindFloat64:
		prefix = "Float64"
	case codegen.KindStringFixed:
		prefix = "String"
	case codegen.KindStringVLS:
		prefix = "String"
	}

	if strings.HasSuffix(messagePackage, ".") {
		messagePackage = messagePackage[0 : len(messagePackage)-1]
	}
	return fmt.Sprintf("%s.%s%s(%s.%s, %d, %d)", messagePackage, prefix, suffix, messageVar, ptrAccessor, bounds, offset)
}

func (g *Generator) writeFieldSetters(w *Writer, f *Field) {
	if f.Fields != nil {
		for _, field := range f.Fields {
			w.Line("///////////////////////////////////////////")
			w.Line("// Set%s", field.Name)
			w.Line("///////////////////////////////////////////")
			w.Line("")
			if field.Field.Type.Kind == codegen.KindStringVLS {
				g.writeFieldSetter(w, field, "Builder", "VLSBuilder")
				g.writeFieldSetter(w, field, "PointerBuilder", "VLSPointerBuilder")
			} else {
				g.writeFieldSetter(w, field, "Builder", "Unsafe()")
				g.writeFieldSetter(w, field, "PointerBuilder", "Ptr")
			}
		}
	} else {
		w.Line("///////////////////////////////////////////")
		w.Line("// Set%s", f.Name)
		w.Line("///////////////////////////////////////////")
		w.Line("")
		if f.Field.Type.Kind == codegen.KindStringVLS {
			g.writeFieldSetter(w, f, "Builder", "VLSBuilder")
			g.writeFieldSetter(w, f, "PointerBuilder", "VLSPointerBuilder")
		} else {
			g.writeFieldSetter(w, f, "Builder", "Unsafe()")
			g.writeFieldSetter(w, f, "PointerBuilder", "Ptr")
		}
	}
}

func (g *Generator) writeFieldSetter(w *Writer, f *Field, suffix, ptrAccessor string) {
	w.Line("func (m %s%s) Set%s(value %s) {", f.Struct.Name, suffix, f.Name, g.GoTypeName(&f.Type, g.fixed))
	w.IndentLine(1, "%s", g.fieldSetValue(f, ptrAccessor))
	w.Line("}")
	w.Line("")
}

func (g *Generator) fieldSetValue(f *Field, ptrAccessor string) string {
	if f.Field.Type.Alias != nil {
		return primitiveSetter(
			f.Field.Type.Alias.Type.Kind,
			f.Field.Type.Offset,
			f.Field.Type.Size,
			"message",
			"m",
			ptrAccessor,
			fmt.Sprintf("%s(value)", primitiveKindName(f.Field.Type.Alias.Type.Kind)),
			f.Struct.VLS,
		)
	}
	if f.Field.Type.Enum != nil {
		return primitiveSetter(
			f.Field.Type.Enum.Type,
			f.Field.Type.Offset,
			f.Field.Type.Size,
			"message",
			"m",
			ptrAccessor,
			fmt.Sprintf("%s(value)", primitiveKindName(f.Field.Type.Enum.Type)),
			f.Struct.VLS,
		)
	}
	return primitiveSetter(
		f.Field.Type.Kind,
		f.Field.Type.Offset,
		f.Field.Type.Size,
		"message",
		"m",
		ptrAccessor,
		"value",
		f.Struct.VLS,
	)
}

func primitiveSetter(kind codegen.Kind, offset, size int, messagePackage, messageVar, ptrAccessor, valueExpr string, vls bool) string {
	var (
		bounds = offset + size
		suffix = "Fixed"
		prefix = ""
	)
	if vls {
		suffix = "VLS"
	}
	switch kind {
	case codegen.KindInt8:
		prefix = "SetInt8"
	case codegen.KindInt16:
		prefix = "SetInt16"
	case codegen.KindInt32:
		prefix = "SetInt32"
	case codegen.KindInt64:
		prefix = "SetInt64"
	case codegen.KindUint8:
		prefix = "SetUint8"
	case codegen.KindUint16:
		prefix = "SetUint16"
	case codegen.KindUint32:
		prefix = "SetUint32"
	case codegen.KindUint64:
		prefix = "SetUint64"
	case codegen.KindFloat32:
		prefix = "SetFloat32"
	case codegen.KindFloat64:
		prefix = "SetFloat64"
	case codegen.KindStringFixed:
		prefix = "SetString"
	case codegen.KindStringVLS:
		prefix = "SetString"
	}

	if strings.HasSuffix(messagePackage, ".") {
		messagePackage = messagePackage[0 : len(messagePackage)-1]
	}

	if kind == codegen.KindStringVLS {
		return fmt.Sprintf("%s.%s%s(%s.%s, %d, %d, %s)", messagePackage, prefix, suffix, messageVar, ptrAccessor, bounds, offset, valueExpr)
	} else {
		return fmt.Sprintf("%s.%s%s(%s.%s, %d, %d, %s)", messagePackage, prefix, suffix, messageVar, ptrAccessor, bounds, offset, valueExpr)
	}
}

func isHeaderField(f *Field) bool {
	if f == nil {
		return false
	}
	switch strings.ToLower(f.Name) {
	case "size", "type", "basesize":
		return true
	}
	return false
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

func primitiveTypeName(kind codegen.Kind) string {
	switch kind {
	case codegen.KindBool:
		return "bool"
	case codegen.KindInt8:
		return "int8"
	case codegen.KindInt16:
		return "int16"
	case codegen.KindInt32:
		return "int32"
	case codegen.KindInt64:
		return "int64"
	case codegen.KindUint8:
		return "uint8"
	case codegen.KindUint16:
		return "uint16"
	case codegen.KindUint32:
		return "uint32"
	case codegen.KindUint64:
		return "uint64"
	case codegen.KindFloat32:
		return "float32"
	case codegen.KindFloat64:
		return "float64"
	case codegen.KindStringFixed, codegen.KindStringVLS:
		return "string"
	default:
		return "invalid"
	}
}

func (g *Generator) GoTypeName(t *codegen.Type, fromPkg *Package) string {
	switch t.Kind {
	case codegen.KindBool:
		return "bool"
	case codegen.KindInt8:
		return "int8"
	case codegen.KindInt16:
		return "int16"
	case codegen.KindInt32:
		return "int32"
	case codegen.KindInt64:
		return "int64"
	case codegen.KindUint8:
		return "uint8"
	case codegen.KindUint16:
		return "uint16"
	case codegen.KindUint32:
		return "uint32"
	case codegen.KindUint64:
		return "uint64"
	case codegen.KindFloat32:
		return "float32"
	case codegen.KindFloat64:
		return "float64"
	case codegen.KindStringFixed, codegen.KindStringVLS:
		return "string"
	case codegen.KindEnum:
		if fromPkg == nil || g.root != fromPkg {
			return fmt.Sprintf("%s.%s", g.root.Name, cleanName(t.Enum.Name))
		}
		return t.Enum.Name
	case codegen.KindAlias:
		if fromPkg == nil || g.root != fromPkg {
			return fmt.Sprintf("%s.%s", g.root.Name, cleanName(t.Alias.Name))
		}
		return t.Alias.Name
	default:
		return "invalid"
	}
}
