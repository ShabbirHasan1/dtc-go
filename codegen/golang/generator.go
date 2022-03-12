package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen"
	"github.com/moontrade/dtc-go/message/pb"
	"io/ioutil"
	"os"
	"os/exec"
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
	Dir            string
	RootPackage    string
	FactoryPackage string
	GoFmt          bool
	NonStandard    bool
	Json           bool
	Protobuf       bool
	GC             bool
	NoGC           bool
	TinyGo         bool
	Debug          bool
}

func DefaultConfig(dir string) *Config {
	return &Config{
		Dir:            dir,
		RootPackage:    "github.com/moontrade/dtc-go/model",
		FactoryPackage: "factory",
		GoFmt:          true,
		NonStandard:    false,
		Json:           true,
		Protobuf:       true,
		GC:             true,
		NoGC:           true,
		TinyGo:         false,
	}
}

func TinyGoConfig(dir string) *Config {
	return &Config{
		Dir:            dir,
		RootPackage:    "github.com/moontrade/dtc-go/model",
		FactoryPackage: "factory",
		GoFmt:          true,
		NonStandard:    false,
		Json:           false,
		Protobuf:       false,
		GC:             true,
		NoGC:           true,
		TinyGo:         true,
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
		if !st.VLS && schema.MessagesByName[st.Name].VLS != nil {
			s.Name = fmt.Sprintf("%sFixed", s.Name)
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
		if !config.NonStandard && message.NonStandard {
			continue
		}
		if config.NonStandard && !message.NonStandard {
			continue
		}
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
	return nil
}

func (g *Generator) writeFile(name string, data []byte) error {
	path := filepath.Join(g.config.Dir, name)
	if err := os.WriteFile(path, data, 0755); err != nil {
		return err
	}
	if g.config.GoFmt {
		if err := exec.Command("gofmt", "-w", path).Run(); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateAliases() error {
	w := &Writer{}
	w.Line("package %s", g.root.Name)
	w.Line("")

	w.Line("type (")
	for _, alias := range g.aliases {
		if alias.Doc != nil {
			_ = g.writeComments(w, 1, alias.Name, alias.Doc.Description)
		}
		w.IndentLine(1, "%s %s", alias.Name, g.GoTypeName(&alias.Type))

		if alias != g.aliases[len(g.aliases)-1] {
			w.Line("")
		}
	}
	w.Line(")")
	w.Line("")

	return g.writeFile("alias.go", w.b)
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
		w.IndentLine(1, "%s %s = %s", constant.Name, g.GoTypeName(&constant.Type), valueToGo(constant.Const.Value))
	}
	w.Line(")")
	w.Line("")

	return g.writeFile("constants.go", w.b)
}

func (g *Generator) generateEnums() error {
	w := &Writer{}
	w.Line("package %s", g.root.Name)
	w.Line("")

	for _, enum := range g.enums {
		if enum.Doc != nil {
			_ = g.writeComments(w, 0, enum.Name, enum.Doc.Description)
		}
		w.Line("type %s %s", enum.Name, primitiveKindName(enum.Type))
		w.Line("")
		w.Line("const (")
		for _, option := range enum.Options {
			if option.Doc != nil {
				_ = g.writeComments(w, 1, option.Name, option.Doc.Description)
			}
			w.IndentLine(1, "%s %s = %d", option.Name, enum.Name, option.Value)
		}
		w.Line(")")
		w.Line("")
	}

	return g.writeFile("enums.go", w.b)
}

func (g *Generator) generateStructDebug(msg *Struct) error {
	return nil
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
					w.IndentLine(1, "%s() %s", field.Name, g.GoTypeName(&field.Type))
				}
			} else {
				w.IndentLine(1, "%s() %s", field.Name, g.GoTypeName(&field.Type))
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
		w.IndentLine(1, "Finish() %s", msg.Name)
		w.Line("")

		for i, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}

			if field.Fields != nil {
				for _, field := range field.Fields {
					if field.Type.Kind.IsString() {
						w.IndentLine(1, "Set%s(value %s) %sBuilder", field.Name, g.GoTypeName(&field.Type), msg.Name)
					} else {
						w.IndentLine(1, "Set%s(value %s)", field.Name, g.GoTypeName(&field.Type))
					}
				}
			} else {
				if field.Type.Kind.IsString() {
					w.IndentLine(1, "Set%s(value %s) %sBuilder", field.Name, g.GoTypeName(&field.Type), msg.Name)
				} else {
					w.IndentLine(1, "Set%s(value %s)", field.Name, g.GoTypeName(&field.Type))
				}
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

		if err := g.writeFile(fmt.Sprintf("%s_interfaces.go", toSnakeCase(msg.Name)), w.b); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) writeComments(w *Writer, indent int, name string, comments []string) error {
	if len(comments) == 0 {
		return nil
	}
	comments = g.normalizeComments(comments)
	if !strings.HasPrefix(comments[0], name) {
		comments[0] = fmt.Sprintf("%s %s", name, strings.TrimSpace(comments[0]))
	}
	for _, comment := range comments {
		comment = strings.TrimSpace(comment)
		if len(comment) == 0 {
			w.IndentLine(indent, "//")
		} else {
			w.IndentLine(indent, "// %s", comment)
		}
	}
	return nil
}

func (g *Generator) normalizeComments(comments []string) []string {
	to := make([]string, 0, len(comments))
	for _, comment := range comments {
		for _, alias := range g.aliases {
			comment = strings.ReplaceAll(comment, alias.Alias.Name, alias.Name)
		}
		for _, enum := range g.enums {
			comment = strings.ReplaceAll(comment, enum.Enum.Name, enum.Name)
		}
		for _, msg := range g.messages {
			m := msg.VLS
			if m == nil {
				m = msg.Fixed
			}
			comment = strings.ReplaceAll(comment, m.Struct.Name, m.Name)
		}

		for len(comment) > 70 {
			index := 69
		LOOP:
			for ; index < len(comment); index++ {
				c := comment[index]
				switch c {
				case ' ', '\r', '\t':
					break LOOP

				case '.':
					index += 1

					break LOOP
				}
			}

			if index >= len(comment) {
				comment = strings.TrimSpace(comment)
				if len(comment) > 0 {
					to = append(to, comment)
				}
				break
			}

			cutoff := strings.TrimSpace(comment[0:index])
			to = append(to, cutoff)
			comment = strings.TrimSpace(comment[index:])
		}

		to = append(to, comment)
	}
	return to
}

func (g *Generator) generateStructs() error {
	for _, m := range g.messages {
		var (
			gcWriter   *Writer
			nogcWriter *Writer
		)
		if g.config.GC {
			gcWriter = &Writer{}
		}
		if g.config.NoGC {
			if gcWriter == nil {
				nogcWriter = &Writer{}
			} else {
				nogcWriter = gcWriter
			}
		}

		writeImports := func(w *Writer) {
			w.Line("package %s", g.root.Name)
			w.Line("")
			w.Line("import (")
			w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
			w.Line(")")
			w.Line("")
		}

		if gcWriter != nil {
			writeImports(gcWriter)
		}
		if nogcWriter != nil && nogcWriter != gcWriter {
			writeImports(nogcWriter)
		}

		writeMessageDefault := func(w *Writer, msg *Struct) {
			if msg == nil {
				return
			}
			g.writeClearComment(w, msg, "")
			w.Write("var _%sDefault = ", msg.Name)
			w.WriteByteSlice(msg.Init)
			w.Write("\n\n")
		}
		writeMessageSize := func(w *Writer, msg *Struct) {
			if msg == nil {
				return
			}
			w.Line("const %sSize = %d", msg.Name, msg.Size)
			w.Line("")
		}

		if gcWriter != nil {
			writeMessageSize(gcWriter, m.VLS)
			writeMessageSize(gcWriter, m.Fixed)
			writeMessageDefault(gcWriter, m.VLS)
			writeMessageDefault(gcWriter, m.Fixed)
		}
		if nogcWriter != nil && gcWriter != nogcWriter {
			writeMessageSize(nogcWriter, m.VLS)
			writeMessageSize(nogcWriter, m.Fixed)
			writeMessageDefault(nogcWriter, m.VLS)
			writeMessageDefault(nogcWriter, m.Fixed)
		}

		// Type declarations
		if gcWriter != nil {
			write := func(msg *Struct) {
				if msg.Doc != nil {
					_ = g.writeComments(gcWriter, 0, msg.Name, msg.Doc.Description)
				}
				gcWriter.Line("type %s struct {", msg.Name)
				gcWriter.IndentLine(1, "message.%s", msg.Suffix())
				gcWriter.Line("}")
				gcWriter.Line("")

				if msg.Doc != nil {
					_ = g.writeComments(gcWriter, 0, msg.Name+"Builder", msg.Doc.Description)
				}
				gcWriter.Line("type %sBuilder struct {", msg.Name)
				if msg.VLS {
					gcWriter.IndentLine(1, "message.%sBuilder", msg.Suffix())
				} else {
					gcWriter.IndentLine(1, "message.%s", msg.Suffix())
				}
				gcWriter.Line("}")
				gcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}
		if nogcWriter != nil {
			write := func(msg *Struct) {
				if msg.Doc != nil {
					_ = g.writeComments(nogcWriter, 0, msg.Name+"Pointer", msg.Doc.Description)
				}
				nogcWriter.Line("type %sPointer struct {", msg.Name)
				nogcWriter.IndentLine(1, "message.%sPointer", msg.Suffix())
				nogcWriter.Line("}")
				nogcWriter.Line("")

				if msg.Doc != nil {
					_ = g.writeComments(nogcWriter, 0, msg.Name+"PointerBuilder", msg.Doc.Description)
				}
				nogcWriter.Line("type %sPointerBuilder struct {", msg.Name)
				if msg.VLS {
					nogcWriter.IndentLine(1, "message.%sPointerBuilder", msg.Suffix())
				} else {
					nogcWriter.IndentLine(1, "message.%sPointer", msg.Suffix())
				}
				nogcWriter.Line("}")
				nogcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Constructors
		if gcWriter != nil {
			write := func(msg *Struct) {
				/*
					func NewAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustment {
						return AccountBalanceAdjustment{VLS: message.NewVLSFrom(b)}
					}

					func AccountBalanceAdjustmentWrap(b []byte) AccountBalanceAdjustment {
						return AccountBalanceAdjustment{VLS: message.WrapVLS(b)}
					}
				*/
				gcWriter.Line("func New%sFrom(b []byte) %s {", msg.Name, msg.Name)
				gcWriter.IndentLine(1, "return %s{%s: message.New%sFrom(b)}", msg.Name, msg.Suffix(), msg.Suffix())
				gcWriter.Line("}")
				gcWriter.Line("")

				gcWriter.Line("func Wrap%s(b []byte) %s {", msg.Name, msg.Name)
				gcWriter.IndentLine(1, "return %s{%s: message.Wrap%s(b)}", msg.Name, msg.Suffix(), msg.Suffix())
				gcWriter.Line("}")
				gcWriter.Line("")

				gcWriter.Line("func New%s() %sBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					gcWriter.IndentLine(1, "a := %sBuilder{message.New%sBuilder(%d)}", msg.Name, msg.Suffix(), msg.Size)
				} else {
					gcWriter.IndentLine(1, "a := %sBuilder{message.New%s(%d)}", msg.Name, msg.Suffix(), msg.Size)
				}
				gcWriter.IndentLine(1, "a.Unsafe().SetBytes(0, _%sDefault)", msg.Name)
				gcWriter.IndentLine(1, "return a")
				gcWriter.Line("}")
				gcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}
		if nogcWriter != nil {
			write := func(msg *Struct) {
				nogcWriter.Line("func Alloc%s() %sPointerBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					nogcWriter.IndentLine(1, "a := %sPointerBuilder{message.Alloc%sBuilder(%d)}", msg.Name, msg.Suffix(), msg.Size)
				} else {
					nogcWriter.IndentLine(1, "a := %sPointerBuilder{message.Alloc%s(%d)}", msg.Name, msg.Suffix(), msg.Size)
				}
				nogcWriter.IndentLine(1, "a.Ptr.SetBytes(0, _%sDefault)", msg.Name)
				nogcWriter.IndentLine(1, "return a")
				nogcWriter.Line("}")
				nogcWriter.Line("")

				/*
					func AllocAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustmentPointer {
						return AccountBalanceAdjustmentPointer{VLSPointer: message.AllocVLSFrom(b)}
					}
				*/
				nogcWriter.Line("func Alloc%sFrom(b []byte) %sPointer {", msg.Name, msg.Name)
				nogcWriter.IndentLine(1, "return %sPointer{%sPointer: message.Alloc%sFrom(b)}", msg.Name, msg.Suffix(), msg.Suffix())
				nogcWriter.Line("}")
				nogcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Clear
		if gcWriter != nil {
			write := func(msg *Struct) {
				g.writeClearComment(gcWriter, msg, "Clear")
				gcWriter.Line("func (m %sBuilder) Clear() {", msg.Name)
				gcWriter.IndentLine(1, "m.Unsafe().SetBytes(0, _%sDefault)", msg.Name)
				gcWriter.Line("}")
				gcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}
		if nogcWriter != nil {
			write := func(msg *Struct) {
				g.writeClearComment(nogcWriter, msg, "Clear")
				nogcWriter.Line("func (m %sPointerBuilder) Clear() {", msg.Name)
				nogcWriter.IndentLine(1, "m.Ptr.SetBytes(0, _%sDefault)", msg.Name)
				nogcWriter.Line("}")
				nogcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// ToBuilder
		if gcWriter != nil {
			write := func(msg *Struct) {
				gcWriter.Line("// ToBuilder")
				gcWriter.Line("func (m %s) ToBuilder() %sBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					gcWriter.IndentLine(1, "return %sBuilder{m.%s.ToBuilder()}", msg.Name, msg.Suffix())
				} else {
					gcWriter.IndentLine(1, "return %sBuilder{m.%s}", msg.Name, msg.Suffix())
				}
				gcWriter.Line("}")
				gcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}
		if nogcWriter != nil {
			write := func(msg *Struct) {
				nogcWriter.Line("// ToBuilder")
				nogcWriter.Line("func (m %sPointer) ToBuilder() %sPointerBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					nogcWriter.IndentLine(1, "return %sPointerBuilder{m.%sPointer.ToBuilder()}", msg.Name, msg.Suffix())
				} else {
					nogcWriter.IndentLine(1, "return %sPointerBuilder{m.%sPointer}", msg.Name, msg.Suffix())
				}
				nogcWriter.Line("}")
				nogcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Finish
		if gcWriter != nil {
			write := func(msg *Struct) {
				gcWriter.Line("// Finish")
				gcWriter.Line("func (m %sBuilder) Finish() %s {", msg.Name, msg.Name)
				if msg.VLS {
					gcWriter.IndentLine(1, "return %s{m.%sBuilder.Finish()}", msg.Name, msg.Suffix())
				} else {
					gcWriter.IndentLine(1, "return %s{m.%s}", msg.Name, msg.Suffix())
				}
				gcWriter.Line("}")
				gcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}
		if nogcWriter != nil {
			write := func(msg *Struct) {
				nogcWriter.Line("// Finish")
				nogcWriter.Line("func (m *%sPointerBuilder) Finish() %sPointer {", msg.Name, msg.Name)
				if msg.VLS {
					nogcWriter.IndentLine(1, "return %sPointer{m.%sPointerBuilder.Finish()}", msg.Name, msg.Suffix())
				} else {
					nogcWriter.IndentLine(1, "return %sPointer{m.%sPointer}", msg.Name, msg.Suffix())
				}
				nogcWriter.Line("}")
				nogcWriter.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Getters
		writeGetters := func(msg *Struct) {
			for _, f := range msg.Fields {
				if isHeaderField(f) {
					continue
				}

				if f.Fields != nil {
					for _, field := range f.Fields {
						if gcWriter != nil {
							if field.Doc != nil {
								_ = g.writeComments(gcWriter, 0, field.Name, field.Doc.Description)
							} else {
								gcWriter.Line("// %s", field.Name)
							}
							g.writeFieldGetter(gcWriter, field, "", "Unsafe()")
							if field.Doc != nil {
								_ = g.writeComments(gcWriter, 0, field.Name, field.Doc.Description)
							} else {
								gcWriter.Line("// %s", field.Name)
							}
							g.writeFieldGetter(gcWriter, field, "Builder", "Unsafe()")
						}
						if nogcWriter != nil {
							if field.Doc != nil {
								_ = g.writeComments(nogcWriter, 0, field.Name, field.Doc.Description)
							} else {
								nogcWriter.Line("// %s", field.Name)
							}
							g.writeFieldGetter(nogcWriter, field, "Pointer", "Ptr")
							if field.Doc != nil {
								_ = g.writeComments(nogcWriter, 0, field.Name, field.Doc.Description)
							} else {
								nogcWriter.Line("// %s", field.Name)
							}
							g.writeFieldGetter(nogcWriter, field, "PointerBuilder", "Ptr")
						}
					}
				} else {
					if gcWriter != nil {
						if f.Doc != nil {
							_ = g.writeComments(gcWriter, 0, f.Name, f.Doc.Description)
						} else {
							gcWriter.Line("// %s", f.Name)
						}
						g.writeFieldGetter(gcWriter, f, "", "Unsafe()")
						if f.Doc != nil {
							_ = g.writeComments(gcWriter, 0, f.Name, f.Doc.Description)
						} else {
							gcWriter.Line("// %s", f.Name)
						}
						g.writeFieldGetter(gcWriter, f, "Builder", "Unsafe()")
					}
					if nogcWriter != nil {
						if f.Doc != nil {
							_ = g.writeComments(nogcWriter, 0, f.Name, f.Doc.Description)
						} else {
							nogcWriter.Line("// %s", f.Name)
						}
						g.writeFieldGetter(nogcWriter, f, "Pointer", "Ptr")
						if f.Doc != nil {
							_ = g.writeComments(nogcWriter, 0, f.Name, f.Doc.Description)
						} else {
							nogcWriter.Line("// %s", f.Name)
						}
						g.writeFieldGetter(nogcWriter, f, "PointerBuilder", "Ptr")
					}
				}
			}
		}
		if m.VLS != nil {
			writeGetters(m.VLS)
		}
		if m.Fixed != nil {
			writeGetters(m.Fixed)
		}

		// Setters
		writeSetters := func(msg *Struct) {
			for _, f := range msg.Fields {
				if isHeaderField(f) {
					continue
				}
				if f.Fields != nil {
					for _, field := range f.Fields {
						if field.Field.Type.Kind == codegen.KindStringVLS {
							if gcWriter != nil {
								if field.Doc != nil {
									_ = g.writeComments(gcWriter, 0, "Set"+field.Name, field.Doc.Description)
								} else {
									gcWriter.Line("// Set%s", field.Name)
								}
								g.writeFieldSetter(gcWriter, field, "Builder", "VLSBuilder")
							}
							if nogcWriter != nil {
								if field.Doc != nil {
									_ = g.writeComments(nogcWriter, 0, "Set"+field.Name, field.Doc.Description)
								} else {
									nogcWriter.Line("// Set%s", field.Name)
								}
								g.writeFieldSetter(nogcWriter, field, "PointerBuilder", "VLSPointerBuilder")
							}
						} else {
							if gcWriter != nil {
								if field.Doc != nil {
									_ = g.writeComments(gcWriter, 0, "Set"+field.Name, field.Doc.Description)
								} else {
									gcWriter.Line("// Set%s", field.Name)
								}
								g.writeFieldSetter(gcWriter, field, "Builder", "Unsafe()")
							}

							if nogcWriter != nil {
								if field.Doc != nil {
									_ = g.writeComments(nogcWriter, 0, "Set"+field.Name, field.Doc.Description)
								} else {
									nogcWriter.Line("// Set%s", field.Name)
								}
								g.writeFieldSetter(nogcWriter, field, "PointerBuilder", "Ptr")
							}
						}
					}
				} else {

					if f.Field.Type.Kind == codegen.KindStringVLS {
						if gcWriter != nil {
							if f.Doc != nil {
								_ = g.writeComments(gcWriter, 0, "Set"+f.Name, f.Doc.Description)
							} else {
								gcWriter.Line("// Set%s", f.Name)
							}
							g.writeFieldSetter(gcWriter, f, "Builder", "VLSBuilder")
						}

						if nogcWriter != nil {
							if f.Doc != nil {
								_ = g.writeComments(nogcWriter, 0, "Set"+f.Name, f.Doc.Description)
							} else {
								nogcWriter.Line("// Set%s", f.Name)
							}
							g.writeFieldSetter(nogcWriter, f, "PointerBuilder", "VLSPointerBuilder")
						}
					} else {
						if gcWriter != nil {
							if f.Doc != nil {
								_ = g.writeComments(gcWriter, 0, "Set"+f.Name, f.Doc.Description)
							} else {
								gcWriter.Line("// Set%s", f.Name)
							}
							g.writeFieldSetter(gcWriter, f, "Builder", "Unsafe()")
						}

						if nogcWriter != nil {
							if f.Doc != nil {
								_ = g.writeComments(nogcWriter, 0, "Set"+f.Name, f.Doc.Description)
							} else {
								nogcWriter.Line("// Set%s", f.Name)
							}
							g.writeFieldSetter(nogcWriter, f, "PointerBuilder", "Ptr")
						}
					}
				}
			}
		}
		if m.VLS != nil {
			writeSetters(m.VLS)
		}
		if m.Fixed != nil {
			writeSetters(m.Fixed)
		}

		copyTo := func(w *Writer, msg *Struct, fn, name, toName string) {
			w.Line("// %s", fn)
			w.Line("func (m %s) %s(to %s) {", name, fn, toName)
			for _, field := range msg.Fields {
				if isHeaderField(field) {
					continue
				}
				if len(field.Fields) > 0 {
					for _, field := range field.Fields {
						w.IndentLine(1, "to.Set%s(m.%s())", field.Name, field.Name)
					}
				} else {
					w.IndentLine(1, "to.Set%s(m.%s())", field.Name, field.Name)
				}
			}
			w.Line("}")
			w.Line("")
		}

		// Copy
		if gcWriter != nil {
			write := func(msg *Struct) {
				copyTo(gcWriter, msg, "Copy", msg.Name, msg.Name+"Builder")
				copyTo(gcWriter, msg, "Copy", msg.Name+"Builder", msg.Name+"Builder")

				if nogcWriter != nil {
					copyTo(gcWriter, msg, "CopyPointer", msg.Name, msg.Name+"PointerBuilder")
					copyTo(gcWriter, msg, "CopyPointer", msg.Name+"Builder", msg.Name+"PointerBuilder")
				}

				if msg.Message != nil {
					to := msg.Message.VLS
					if msg.VLS {
						to = msg.Message.Fixed
					}
					if to != nil {
						if nogcWriter != nil {
							copyTo(gcWriter, msg, "CopyToPointer", msg.Name, to.Name+"PointerBuilder")
							copyTo(gcWriter, msg, "CopyToPointer", msg.Name+"Builder", to.Name+"PointerBuilder")
						}
						copyTo(gcWriter, msg, "CopyTo", msg.Name, to.Name+"Builder")
						copyTo(gcWriter, msg, "CopyTo", msg.Name+"Builder", to.Name+"Builder")
					}
				}
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}
		if nogcWriter != nil {
			write := func(msg *Struct) {
				if gcWriter != nil {
					copyTo(nogcWriter, msg, "Copy", msg.Name+"Pointer", msg.Name+"Builder")
					copyTo(nogcWriter, msg, "Copy", msg.Name+"PointerBuilder", msg.Name+"Builder")
				}

				copyTo(nogcWriter, msg, "CopyPointer", msg.Name+"Pointer", msg.Name+"PointerBuilder")
				copyTo(nogcWriter, msg, "CopyPointer", msg.Name+"PointerBuilder", msg.Name+"PointerBuilder")

				if msg.Message != nil {
					to := msg.Message.VLS
					if msg.VLS {
						to = msg.Message.Fixed
					}
					if to != nil {
						if gcWriter != nil {
							copyTo(nogcWriter, msg, "CopyTo", msg.Name+"Pointer", to.Name+"Builder")
							copyTo(nogcWriter, msg, "CopyTo", msg.Name+"PointerBuilder", to.Name+"Builder")
						}
						copyTo(nogcWriter, msg, "CopyToPointer", msg.Name+"Pointer", to.Name+"PointerBuilder")
						copyTo(nogcWriter, msg, "CopyToPointer", msg.Name+"PointerBuilder", to.Name+"PointerBuilder")
					}
				}
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		if gcWriter == nogcWriter {
			if gcWriter != nil {
				if err := g.writeFile(fmt.Sprintf("%s.go", toSnakeCase(m.Name())), gcWriter.b); err != nil {
					return err
				}
			}
		} else {
			if gcWriter != nil {
				if err := g.writeFile(fmt.Sprintf("%s.go", toSnakeCase(m.Name())), gcWriter.b); err != nil {
					return err
				}
			}
			if nogcWriter != nil {
				if err := g.writeFile(fmt.Sprintf("%s_nogc.go", toSnakeCase(m.Name())), nogcWriter.b); err != nil {
					return err
				}
			}
		}

		if err := g.generateSerializers(m); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) generateSerializers(msg *Message) error {
	if !g.config.Json && !g.config.Protobuf {
		return nil
	}
	if !msg.HasSerializers() {
		return nil
	}
	var (
		gcWriter   *Writer
		nogcWriter *Writer
	)
	if g.config.GC {
		gcWriter = &Writer{}
	}
	if g.config.NoGC {
		if gcWriter == nil {
			nogcWriter = &Writer{}
		} else {
			nogcWriter = gcWriter
		}
	}

	writeImports := func(w *Writer) {
		w.Line("//go:build !tinygo")
		w.Line("")
		w.Line("package %s", g.root.Name)
		w.Line("")
		w.Line("import (")
		w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
		if g.config.Json {
			w.IndentLine(1, "\"github.com/moontrade/dtc-go/message/json\"")
		}
		if g.config.Protobuf && msg.HasProtobuf() {
			w.IndentLine(1, "\"github.com/moontrade/dtc-go/message/pb\"")
		}
		w.Line(")")
		w.Line("")
	}
	if gcWriter != nil {
		writeImports(gcWriter)
	}
	if nogcWriter != nil && nogcWriter != gcWriter {
		writeImports(nogcWriter)
	}

	if err := g.generateJson(msg.VLS, gcWriter, nogcWriter); err != nil {
		return err
	}
	if err := g.generateJson(msg.Fixed, gcWriter, nogcWriter); err != nil {
		return err
	}
	if err := g.generateProtobuf(msg.VLS, gcWriter, nogcWriter); err != nil {
		return err
	}
	if err := g.generateProtobuf(msg.Fixed, gcWriter, nogcWriter); err != nil {
		return err
	}

	_ = os.MkdirAll(g.config.Dir, 0755)
	if gcWriter == nogcWriter {
		if gcWriter != nil {
			if err := g.writeFile(fmt.Sprintf("%s_serializer.go", toSnakeCase(msg.Name())), gcWriter.b); err != nil {
				return err
			}
		}
	} else {
		if gcWriter != nil {
			if err := g.writeFile(fmt.Sprintf("%s_serializer.go", toSnakeCase(msg.Name())), gcWriter.b); err != nil {
				return err
			}
		}
		if nogcWriter != nil {
			if err := g.writeFile(fmt.Sprintf("%s_nogc_serializer.go", toSnakeCase(msg.Name())), nogcWriter.b); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *Generator) generateJson(msg *Struct, gcWriter, nogcWriter *Writer) error {
	if msg == nil || !g.config.Json {
		return nil
	}

	marshalJSONCompact := func(w *Writer, name string) {
		w.Line("func (m %s) MarshalJSONCompactTo(b []byte) ([]byte, error) {", name)
		w.IndentLine(1, "w := json.Writer{Data: b, Compact: true}")
		w.IndentLine(1, "w.Data = append(w.Data, \"{\\\"Type\\\":%d,\\\"F\\\":[\"...)", msg.Type)
		first := true
		for _, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}
			if first {
				first = false
			} else {
				w.IndentLine(1, "w.Data = append(w.Data, ',')")
			}

			if len(field.Fields) > 0 {
				maxField := field.Fields[0]
				for _, f := range field.Fields {
					if f.Type.Size > maxField.Type.Size {
						maxField = f
					}
				}
				field = maxField
			}

			switch field.Type.Kind {
			case codegen.KindEnum:
				w.IndentLine(1, "w.%s(%s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Enum.Type), primitiveTypeName(field.Type.Enum.Type), field.Name)
			case codegen.KindAlias:
				w.IndentLine(1, "w.%s(%s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind), primitiveTypeName(field.Type.Alias.Type.Kind), field.Name)
			default:
				w.IndentLine(1, "w.%s(m.%s())", jsonGetterFuncNamePrimitive(field.Type.Kind), field.Name)
			}
		}
		w.IndentLine(1, "return w.Finish(), nil")
		w.Line("}")
		w.Line("")
	}

	if gcWriter != nil {
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("// JSON Compact Marshal")
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("")

		marshalJSONCompact(gcWriter, msg.Name)
		marshalJSONCompact(gcWriter, msg.Name+"Builder")
	}

	if nogcWriter != nil {
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("// JSON Compact Marshal")
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("")

		marshalJSONCompact(nogcWriter, msg.Name+"Pointer")
		marshalJSONCompact(nogcWriter, msg.Name+"PointerBuilder")
	}

	marshalJSON := func(w *Writer, name string) {
		w.Line("func (m %s) MarshalJSONTo(b []byte) ([]byte, error) {", name)
		w.IndentLine(1, "w := json.NewWriter(b, %d)", msg.Struct.Type)
		for _, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}
			if len(field.Fields) > 0 {
				for _, field := range field.Fields {
					switch field.Type.Kind {
					case codegen.KindEnum:
						w.IndentLine(1, "w.%sField(\"%s\", %s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Enum.Type), field.Field.Name, primitiveTypeName(field.Type.Enum.Type), field.Name)
					case codegen.KindAlias:
						w.IndentLine(1, "w.%sField(\"%s\", %s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind), field.Field.Name, primitiveTypeName(field.Type.Alias.Type.Kind), field.Name)
					default:
						w.IndentLine(1, "w.%sField(\"%s\", m.%s())", jsonGetterFuncNamePrimitive(field.Type.Kind), field.Field.Name, field.Name)
					}
				}
			} else {
				switch field.Type.Kind {
				case codegen.KindEnum:
					w.IndentLine(1, "w.%sField(\"%s\", %s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Enum.Type), field.Field.Name, primitiveTypeName(field.Type.Enum.Type), field.Name)
				case codegen.KindAlias:
					w.IndentLine(1, "w.%sField(\"%s\", %s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind), field.Field.Name, primitiveTypeName(field.Type.Alias.Type.Kind), field.Name)
				default:
					w.IndentLine(1, "w.%sField(\"%s\", m.%s())", jsonGetterFuncNamePrimitive(field.Type.Kind), field.Field.Name, field.Name)
				}
			}
		}
		w.IndentLine(1, "return w.Finish(), nil")
		w.Line("}")
		w.Line("")
	}

	if gcWriter != nil {
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("// JSON Marshal")
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("")
		marshalJSON(gcWriter, msg.Name)
		marshalJSON(gcWriter, msg.Name+"Builder")
	}
	if nogcWriter != nil {
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("// JSON Marshal")
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("")
		marshalJSON(nogcWriter, msg.Name+"Pointer")
		marshalJSON(nogcWriter, msg.Name+"PointerBuilder")
	}

	unmarshalJSONCompact := func(w *Writer, name string) {
		w.Line("func (m *%s) UnmarshalJSONCompactFrom(r *json.Reader) error {", name)
		//w.IndentLine(1, "r, err := json.OpenReader(b)")
		//w.IndentLine(1, "if err != nil {")
		//w.IndentLine(2, "return err")
		//w.IndentLine(1, "}")
		w.IndentLine(1, "if r.Type != %d {", msg.Type)
		w.IndentLine(2, "return message.ErrWrongType")
		w.IndentLine(1, "}")
		for _, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}
			if len(field.Fields) > 0 {
				maxField := field.Fields[0]
				for _, f := range field.Fields {
					if f.Type.Size > maxField.Type.Size {
						maxField = f
					}
				}
				field = maxField
			}
			switch field.Type.Kind {
			case codegen.KindEnum:
				w.IndentLine(1, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type))
			case codegen.KindAlias:
				w.IndentLine(1, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind))
			default:
				w.IndentLine(1, "m.Set%s(r.%s())", field.Name, jsonGetterFuncNamePrimitive(field.Type.Kind))
			}
			w.IndentLine(1, "if r.IsError() {")
			w.IndentLine(2, "return r.Error()")
			w.IndentLine(1, "}")
		}
		w.IndentLine(1, "return nil")
		w.Line("}")
		w.Line("")
	}

	if gcWriter != nil {
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("// JSON Compact Unmarshal")
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("")
		unmarshalJSONCompact(gcWriter, msg.Name+"Builder")
	}
	if nogcWriter != nil {
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("// JSON Compact Unmarshal")
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("")
		unmarshalJSONCompact(nogcWriter, msg.Name+"PointerBuilder")
	}

	unmarshalJSON := func(w *Writer, name string) {
		w.Line("func (m *%s) UnmarshalJSONFrom(r *json.Reader) error {", name)
		w.IndentLine(1, "if r.Type != %d {", msg.Type)
		w.IndentLine(2, "return message.ErrWrongType")
		w.IndentLine(1, "}")
		w.IndentLine(1, "in := &r.Lexer")
		w.IndentLine(0, "LOOP:")
		w.IndentLine(1, "for !in.IsDelim('}') {")
		w.IndentLine(2, "key, err := r.FieldName()")
		w.IndentLine(3, "if err != nil {")
		w.IndentLine(2, "return err")
		w.IndentLine(2, "}")
		w.IndentLine(2, "switch key {")
		for _, field := range msg.Fields {
			if isHeaderField(field) {
				continue
			}
			if len(field.Fields) > 0 {
				for _, field := range field.Fields {
					w.IndentLine(2, "case \"%s\":", field.Field.Name)
					switch field.Type.Kind {
					case codegen.KindEnum:
						w.IndentLine(3, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type))
					case codegen.KindAlias:
						w.IndentLine(3, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind))
					default:
						w.IndentLine(3, "m.Set%s(r.%s())", field.Name, jsonGetterFuncNamePrimitive(field.Type.Kind))
					}
				}
			} else {
				w.IndentLine(2, "case \"%s\":", field.Field.Name)
				switch field.Type.Kind {
				case codegen.KindEnum:
					w.IndentLine(3, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type))
				case codegen.KindAlias:
					w.IndentLine(3, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind))
				default:
					w.IndentLine(3, "m.Set%s(r.%s())", field.Name, jsonGetterFuncNamePrimitive(field.Type.Kind))
				}
			}
		}
		w.IndentLine(2, "case \"f\", \"F\":")
		w.IndentLine(3, "return message.ErrJSONCompactDetected")
		w.IndentLine(2, "case \"\":")
		w.IndentLine(3, "break LOOP")
		w.IndentLine(2, "default:")
		w.IndentLine(3, "in.SkipRecursive()")
		w.IndentLine(2, "}")
		w.IndentLine(2, "if r.IsError() {")
		w.IndentLine(3, "return r.Error()")
		w.IndentLine(2, "}")
		w.IndentLine(2, "in.WantComma()")
		w.IndentLine(1, "}")
		w.IndentLine(1, "return nil")
		w.Line("}")
		w.Line("")
	}

	if gcWriter != nil {
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("// JSON Unmarshal")
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("")

		unmarshalJSON(gcWriter, msg.Name+"Builder")
	}
	if nogcWriter != nil {
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("// JSON Unmarshal")
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("")

		unmarshalJSON(nogcWriter, msg.Name+"PointerBuilder")
	}
	return nil
}

func (g *Generator) generateProtobuf(msg *Struct, gcWriter, nogcWriter *Writer) error {
	if msg == nil || !g.config.Protobuf || msg.Proto == nil {
		return nil
	}

	var fields FieldsSlice
	fields = append(fields, msg.Fields...)
	fields.Sort()

	unmarshalProtobuf := func(w *Writer, name string) {
		w.Line("func (m *%s) UnmarshalProtobuf(b []byte) error {", name)
		w.IndentLine(1, "var (")
		w.IndentLine(2, "r = pb.NewBuffer(b)")
		w.IndentLine(2, "f pb.Number")
		w.IndentLine(2, "err error")
		w.IndentLine(1, ")")
		w.IndentLine(1, "for !r.EOF() {")
		w.IndentLine(2, "f, _, err = r.DecodeTag()")
		w.IndentLine(2, "if err != nil {")
		w.IndentLine(3, "break")
		w.IndentLine(2, "}")
		w.IndentLine(2, "switch f {")

		for _, field := range fields {
			if isHeaderField(field) {
				continue
			}
			if field.Field.Proto == nil {
				continue
			}
			zigzag := ""
			if field.Field.ProtoZigzag {
				zigzag = "Zigzag"
			}
			w.IndentLine(2, "case %s:", field.Field.Proto.FieldNumber)
			if len(field.Fields) > 0 {
				for _, field := range field.Fields {
					switch field.Type.Kind {
					case codegen.KindEnum:
						w.IndentLine(3, "m.Set%s(%s(r.Read%s%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type), zigzag)
					case codegen.KindAlias:
						w.IndentLine(3, "m.Set%s(%s(r.Read%s%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind), zigzag)
					default:
						w.IndentLine(3, "m.Set%s(r.Read%s%s())", field.Name, jsonGetterFuncNamePrimitive(field.Type.Kind), zigzag)
					}
				}
			} else {
				switch field.Type.Kind {
				case codegen.KindEnum:
					w.IndentLine(3, "m.Set%s(%s(r.Read%s%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type), zigzag)
				case codegen.KindAlias:
					w.IndentLine(3, "m.Set%s(%s(r.Read%s%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind), zigzag)
				default:
					w.IndentLine(3, "m.Set%s(r.Read%s%s())", field.Name, jsonGetterFuncNamePrimitive(field.Type.Kind), zigzag)
				}
			}
		}
		w.IndentLine(2, "default:")
		w.IndentLine(3, "r.Consume()")
		w.IndentLine(2, "}")
		w.IndentLine(2, "if err = r.Error(); err != nil {")
		w.IndentLine(3, "return err")
		w.IndentLine(2, "}")
		w.IndentLine(1, "}")
		w.IndentLine(1, "return err")
		w.Line("}")
		w.Line("")
	}

	if gcWriter != nil {
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("// Protocol Buffers Unmarshal")
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("")
		unmarshalProtobuf(gcWriter, msg.Name+"Builder")
	}

	if nogcWriter != nil {
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("// Protocol Buffers Unmarshal")
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("")

		unmarshalProtobuf(nogcWriter, msg.Name+"PointerBuilder")
	}

	marshalProtobuf := func(w *Writer, name string) {
		w.Line("func (m %s) MarshalProtobuf(b []byte) ([]byte, error) {", name)
		w.IndentLine(1, "w := pb.NewWriter(b, %d)", msg.Struct.Type)
		for _, field := range fields {
			if isHeaderField(field) {
				continue
			}

			if len(field.Fields) > 0 {
				for _, field := range field.Fields {
					if field.Proto == nil {
						continue
					}
					w.IndentLine(1, "w.%s", protobufMarshalField(field, fmt.Sprintf("m.%s()", field.Name)))
				}
			} else {
				if field.Proto == nil {
					continue
				}
				w.IndentLine(1, "w.%s", protobufMarshalField(field, fmt.Sprintf("m.%s()", field.Name)))
			}
		}
		w.IndentLine(1, "return w.Finish(), nil")
		w.Line("}")
		w.Line("")
	}

	if gcWriter != nil {
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("// Protocol Buffers Marshal")
		gcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		gcWriter.Line("")

		marshalProtobuf(gcWriter, msg.Name+"Builder")
	}

	if nogcWriter != nil {
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("// Protocol Buffers Marshal")
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("")

		marshalProtobuf(nogcWriter, msg.Name+"PointerBuilder")
	}

	return nil
}

func (g *Generator) writeClearComment(w *Writer, msg *Struct, name string) {
	if len(name) > 0 {
		w.Line("// %s", name)
	}
	w.Line("// {")
	maxWidth := 0
	for _, f := range msg.Fields {
		if len(f.Fields) > 0 {
			continue
		}
		if maxWidth < len(f.Name) {
			maxWidth = len(f.Name)
		}
	}

	maxWidth++
	pad := func(s string) string {
		for len(s) < maxWidth {
			s = s + " "
		}
		return s
	}
	for _, f := range msg.Fields {
		if len(f.Fields) > 0 {
			continue
		}

		name := pad(f.Name)

		if f.Initial == nil {
			switch f.Type.Kind {
			case codegen.KindStringFixed, codegen.KindStringVLS:
				w.Line("//     %s= \"\"", name)
			case codegen.KindBool:
				w.Line("//     %s= false", name)
			case codegen.KindEnum:
				option := f.Type.Enum.OptionByValue(0)
				if option != nil {
					w.Line("//     %s= %s  (%d)", name, option.Name, option.Value)
				} else {
					w.Line("//     %s= 0", name)
				}
			default:
				w.Line("//     %s= 0", name)
			}
			continue
		}

		switch f.Initial.Type {
		case codegen.ValueTypeSizeof:
			w.Line("//     %s= %sSize  (%d)", name, msg.Name, msg.Size)
		case codegen.ValueTypeString:
			w.Line("//     %s= \"%s\"", name, f.Initial.Str)
		case codegen.ValueTypeBool:
			if f.Initial.Int == 0 {
				w.Line("//     %s= false", name)
			} else {
				w.Line("//     %s= true", name)
			}
		case codegen.ValueTypeInt:
			if f.Type.Kind == codegen.KindBool {
				if f.Initial.Int == 0 {
					w.Line("//     %s= false", name)
				} else {
					w.Line("//     %s= true", name)
				}
			} else {
				w.Line("//     %s= %d", name, f.Initial.Int)
			}
		case codegen.ValueTypeUint:
			if f.Type.Kind == codegen.KindBool {
				if f.Initial.Uint == 0 {
					w.Line("//     %s= false", name)
				} else {
					w.Line("//     %s= true", name)
				}
			} else {
				w.Line("//     %s= %d", name, f.Initial.Int)
			}
		case codegen.ValueTypeFloat:
			w.Line("//     %s= %f", name, f.Initial.Float64)
		case codegen.ValueTypeFloat32Max:
			w.Line("//     %s= math.MaxFloat32", name)
		case codegen.ValueTypeFloat64Max:
			w.Line("//     %s= math.MaxFloat64", name)
		case codegen.ValueTypeConst:
			w.Line("//     %s= %s  (%d)", name, f.Initial.Const.Name, f.Initial.Int)
		case codegen.ValueTypeEnumOption:
			w.Line("//     %s= %s  (%d)", name, f.Initial.EnumOption.Name, f.Initial.Int)
		}
	}
	w.Line("// }")
}

func protobufMarshalField(f *Field, getterExpression string) string {
	kind := f.Type.Kind
	switch kind {
	case codegen.KindEnum:
		kind = f.Type.Enum.Type
		getterExpression = fmt.Sprintf("%s(%s)", primitiveTypeName(kind), getterExpression)
	case codegen.KindAlias:
		kind = f.Type.Alias.Type.Kind
		getterExpression = fmt.Sprintf("%s(%s)", primitiveTypeName(kind), getterExpression)
	}
	switch kind {
	case codegen.KindBool:
		return fmt.Sprintf("WriteBool(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindInt8:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Int8(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Int8(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Int8(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteVarint8(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindInt16:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Int16(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Int16(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Int16(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteVarint16(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindInt32:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Int32(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Int32(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Int32(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteVarint32(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindInt64:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag64Int64(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Int64(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Int64(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteVarint64(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindUint8:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Uint8(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Uint8(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Uint8(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteUvarint8(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindUint16:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Uint16(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Uint16(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Uint16(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteUvarint16(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindUint32:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Uint32(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Uint32(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Uint32(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteUvarint32(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindUint64:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Uint64(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Uint64(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Uint64(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteUvarint64(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindFloat32:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag32Float32(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Float32(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Float32(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteVarintFloat32(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindFloat64:
		switch f.ProtoType {
		case pb.VarintType:
			if f.ProtoZigzag {
				return fmt.Sprintf("WriteZigzag64Float64(%s, %s)", f.Proto.FieldNumber, getterExpression)
			}
		case pb.Fixed32Type:
			return fmt.Sprintf("WriteFixed32Float64(%s, %s)", f.Proto.FieldNumber, getterExpression)
		case pb.Fixed64Type:
			return fmt.Sprintf("WriteFixed64Float64(%s, %s)", f.Proto.FieldNumber, getterExpression)
		}
		return fmt.Sprintf("WriteVarintFloat64(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case codegen.KindStringFixed, codegen.KindStringVLS:
		return fmt.Sprintf("WriteString(%s, %s)", f.Proto.FieldNumber, getterExpression)
	default:
		return "Invalid"
	}
}

func jsonGetterFuncNamePrimitive(kind codegen.Kind) string {
	switch kind {
	case codegen.KindBool:
		return "Bool"
	case codegen.KindInt8:
		return "Int8"
	case codegen.KindInt16:
		return "Int16"
	case codegen.KindInt32:
		return "Int32"
	case codegen.KindInt64:
		return "Int64"
	case codegen.KindUint8:
		return "Uint8"
	case codegen.KindUint16:
		return "Uint16"
	case codegen.KindUint32:
		return "Uint32"
	case codegen.KindUint64:
		return "Uint64"
	case codegen.KindFloat32:
		return "Float32"
	case codegen.KindFloat64:
		return "Float64"
	case codegen.KindStringFixed, codegen.KindStringVLS:
		return "String"
	default:
		return "Invalid"
	}
}

func (g *Generator) generateFactory() error {
	return nil
}

func (g *Generator) writeFieldGetter(w *Writer, f *Field, suffix, ptrAccessor string) {
	w.Line("func (m %s%s) %s() %s {", f.Struct.Name, suffix, f.Name, g.GoTypeName(&f.Type))
	w.IndentLine(1, "return %s", g.fieldGetValue(f, ptrAccessor))
	w.Line("}")
	w.Line("")
}

func (g *Generator) fieldGetValue(f *Field, ptrAccessor string) string {
	if f.Field.Type.Alias != nil {
		return fmt.Sprintf("%s(%s)",
			g.GoTypeName(&f.Field.Type),
			//fmt.Sprintf("%s.%s", g.root.Name, cleanName(f.Field.Type.Alias.Name)),
			g.primitiveGetter(
				f.Field.Type.Alias.Type.Kind,
				f.Field.Type.Offset,
				f.Field.Type.Size,
				"m",
				ptrAccessor,
				f.Struct.VLS,
			),
		)
	}
	if f.Field.Type.Enum != nil {
		return fmt.Sprintf("%s(%s)",
			g.GoTypeName(&f.Field.Type),
			g.primitiveGetter(
				f.Field.Type.Enum.Type,
				f.Field.Type.Offset,
				f.Field.Type.Size,
				"m",
				ptrAccessor,
				f.Struct.VLS,
			),
		)
	}

	return g.primitiveGetter(
		f.Field.Type.Kind,
		f.Field.Type.Offset,
		f.Field.Type.Size,
		"m",
		ptrAccessor,
		f.Struct.VLS,
	)
}

func (g *Generator) primitiveGetter(kind codegen.Kind, offset, size int, messageVar, ptrAccessor string, vls bool) string {
	var (
		bounds = offset + size
		suffix = "Fixed"
		prefix = ""
	)
	if vls {
		suffix = "VLS"
	}
	switch kind {
	case codegen.KindBool:
		prefix = "Bool"
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

	//messagePackage := ""
	//if pkg != g.root {
	//	messagePackage = pkg.Name + "."
	//}
	return fmt.Sprintf("message.%s%s(%s.%s, %d, %d)", prefix, suffix, messageVar, ptrAccessor, bounds, offset)
}

func (g *Generator) writeFieldSetter(w *Writer, f *Field, suffix, ptrAccessor string) {
	if f.Field.Type.Kind.IsString() && f.Struct.VLS {
		w.Line("func (m *%s%s) Set%s(value %s) {", f.Struct.Name, suffix, f.Name, g.GoTypeName(&f.Type))
	} else {
		w.Line("func (m %s%s) Set%s(value %s) {", f.Struct.Name, suffix, f.Name, g.GoTypeName(&f.Type))
	}
	w.IndentLine(1, "%s", g.fieldSetValue(f, ptrAccessor))
	w.Line("}")
	w.Line("")
}

func (g *Generator) fieldSetValue(f *Field, ptrAccessor string) string {
	if f.Field.Type.Alias != nil {
		return g.primitiveSetter(
			f.Field.Type.Alias.Type.Kind,
			f.Field.Type.Offset,
			f.Field.Type.Size,
			"m",
			ptrAccessor,
			fmt.Sprintf("%s(value)", primitiveKindName(f.Field.Type.Alias.Type.Kind)),
			f.Struct.VLS,
		)
	}
	if f.Field.Type.Enum != nil {
		return g.primitiveSetter(
			f.Field.Type.Enum.Type,
			f.Field.Type.Offset,
			f.Field.Type.Size,
			"m",
			ptrAccessor,
			fmt.Sprintf("%s(value)", primitiveKindName(f.Field.Type.Enum.Type)),
			f.Struct.VLS,
		)
	}
	return g.primitiveSetter(
		f.Field.Type.Kind,
		f.Field.Type.Offset,
		f.Field.Type.Size,
		"m",
		ptrAccessor,
		"value",
		f.Struct.VLS,
	)
}

func (g *Generator) primitiveSetter(kind codegen.Kind, offset, size int, messageVar, ptrAccessor, valueExpr string, vls bool) string {
	var (
		bounds = offset + size
		suffix = "Fixed"
		prefix = ""
	)
	if vls {
		suffix = "VLS"
	}
	switch kind {
	case codegen.KindBool:
		prefix = "SetBool"
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

	//messagePackage := ""
	//if pkg != g.root {
	//	messagePackage = pkg.Name + "."
	//}

	if kind == codegen.KindStringVLS {
		if ptrAccessor == "VLSPointerBuilder" {
			return fmt.Sprintf("message.%s%sPointer(&%s.%s, %d, %d, %s)", prefix, suffix, messageVar, ptrAccessor, bounds, offset, valueExpr)
		} else {
			return fmt.Sprintf("message.%s%s(&%s.%s, %d, %d, %s)", prefix, suffix, messageVar, ptrAccessor, bounds, offset, valueExpr)
		}

	} else {
		return fmt.Sprintf("message.%s%s(%s.%s, %d, %d, %s)", prefix, suffix, messageVar, ptrAccessor, bounds, offset, valueExpr)
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

func (g *Generator) GoTypeName(t *codegen.Type) string {
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
		return cleanName(t.Enum.Name)
	case codegen.KindAlias:
		return cleanName(t.Alias.Name)
	default:
		return "invalid"
	}
}
