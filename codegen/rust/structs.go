package rust

import (
	"fmt"
	"strings"

	"github.com/moontrade/dtc-go/codegen/schema"
)

func (g *Generator) generateStructDebug(msg *Struct) error {
	return nil
}

func (g *Generator) generateStructs() error {
	for _, m := range g.messages {
		writer := &Writer{}
		if len(g.config.GeneratedComment) > 0 {
			writer.Line("// %s", g.config.GeneratedComment)
			writer.Line("")
		}

		// writeImports := func(w *Writer) {
		// 	w.Line("package %s", g.packageName)
		// 	w.Line("")
		// 	w.Line("import (")
		// 	w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
		// 	w.Line(")")
		// 	w.Line("")
		// }
		// writeImports(writer)

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

		writeMessageSize(writer, m.VLS)
		writeMessageSize(writer, m.Fixed)
		writeMessageDefault(writer, m.VLS)
		writeMessageDefault(writer, m.Fixed)

		// Type declarations
		{
			write := func(msg *Struct) {
				if msg.Doc != nil {
					_ = g.writeComments(writer, 0, msg.Name, msg.Doc.Description)
				}
				writer.Line("type %s struct {", msg.Name)
				writer.IndentLine(1, "message.%s", msg.Suffix())
				writer.Line("}")
				writer.Line("")

				if msg.Doc != nil {
					_ = g.writeComments(writer, 0, msg.Name+"Builder", msg.Doc.Description)
				}
				writer.Line("type %sBuilder struct {", msg.Name)
				if msg.VLS {
					writer.IndentLine(1, "message.%sBuilder", msg.Suffix())
				} else {
					writer.IndentLine(1, "message.%s", msg.Suffix())
				}
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Constructors
		{
			write := func(msg *Struct) {
				/*
					func NewAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustment {
						return AccountBalanceAdjustment{VLS: message.NewVLSFrom(b)}
					}

					func AccountBalanceAdjustmentWrap(b []byte) AccountBalanceAdjustment {
						return AccountBalanceAdjustment{VLS: message.WrapVLS(b)}
					}
				*/
				writer.Line("func New%sFrom(b []byte) %s {", msg.Name, msg.Name)
				writer.IndentLine(1, "return %s{%s: message.New%s(b)}", msg.Name, msg.Suffix(), msg.Suffix())
				writer.Line("}")
				writer.Line("")

				writer.Line("func Wrap%s(b []byte) %s {", msg.Name, msg.Name)
				writer.IndentLine(1, "return %s{%s: message.Wrap%s(b)}", msg.Name, msg.Suffix(), msg.Suffix())
				writer.Line("}")
				writer.Line("")

				writer.Line("func New%s() %sBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					writer.IndentLine(1, "return %sBuilder{message.New%sBuilder(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
				} else {
					writer.IndentLine(1, "return %sBuilder{message.New%s(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
				}
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Clear
		{
			write := func(msg *Struct) {
				g.writeClearComment(writer, msg, "Clear")
				writer.Line("func (m %sBuilder) Clear() {", msg.Name)
				writer.IndentLine(1, "m.Unsafe().SetBytes(0, _%sDefault)", msg.Name)
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// ToBuilder
		{
			write := func(msg *Struct) {
				writer.Line("// ToBuilder")
				writer.Line("func (m %s) ToBuilder() %sBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					writer.IndentLine(1, "return %sBuilder{m.%s.ToBuilder()}", msg.Name, msg.Suffix())
				} else {
					writer.IndentLine(1, "return %sBuilder{m.%s}", msg.Name, msg.Suffix())
				}
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Finish
		{
			write := func(msg *Struct) {
				writer.Line("// Finish")
				writer.Line("func (m %sBuilder) Finish() %s {", msg.Name, msg.Name)
				if msg.VLS {
					writer.IndentLine(1, "return %s{m.%sBuilder.Finish()}", msg.Name, msg.Suffix())
				} else {
					writer.IndentLine(1, "return %s{m.%s}", msg.Name, msg.Suffix())
				}
				writer.Line("}")
				writer.Line("")
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

						if field.Doc != nil {
							_ = g.writeComments(writer, 0, field.Name, field.Doc.Description)
						} else {
							writer.Line("// %s", field.Name)
						}
						g.writeFieldGetter(writer, field, "", "Unsafe()")
						if field.Doc != nil {
							_ = g.writeComments(writer, 0, field.Name, field.Doc.Description)
						} else {
							writer.Line("// %s", field.Name)
						}
						g.writeFieldGetter(writer, field, "Builder", "Unsafe()")

					}
				} else {
					if f.Doc != nil {
						_ = g.writeComments(writer, 0, f.Name, f.Doc.Description)
					} else {
						writer.Line("// %s", f.Name)
					}
					g.writeFieldGetter(writer, f, "", "Unsafe()")
					if f.Doc != nil {
						_ = g.writeComments(writer, 0, f.Name, f.Doc.Description)
					} else {
						writer.Line("// %s", f.Name)
					}
					g.writeFieldGetter(writer, f, "Builder", "Unsafe()")

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
						if field.Field.Type.Kind == schema.KindStringVLS {
							if field.Doc != nil {
								_ = g.writeComments(writer, 0, "Set"+field.Name, field.Doc.Description)
							} else {
								writer.Line("// Set%s", field.Name)
							}
							g.writeFieldSetter(writer, field, "Builder", "VLSBuilder")
						} else {
							if field.Doc != nil {
								_ = g.writeComments(writer, 0, "Set"+field.Name, field.Doc.Description)
							} else {
								writer.Line("// Set%s", field.Name)
							}
							g.writeFieldSetter(writer, field, "Builder", "Unsafe()")
						}
					}
				} else {

					if f.Field.Type.Kind == schema.KindStringVLS {
						if f.Doc != nil {
							_ = g.writeComments(writer, 0, "Set"+f.Name, f.Doc.Description)
						} else {
							writer.Line("// Set%s", f.Name)
						}
						g.writeFieldSetter(writer, f, "Builder", "VLSBuilder")
					} else {
						if f.Doc != nil {
							_ = g.writeComments(writer, 0, "Set"+f.Name, f.Doc.Description)
						} else {
							writer.Line("// Set%s", f.Name)
						}
						g.writeFieldSetter(writer, f, "Builder", "Unsafe()")
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

		if g.config.GenerateCopyMethods {
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

			//copy := func(w *Writer, msg *Struct, fn, name, toName string) {
			//	w.Line("// %s", fn)
			//	w.Line("func (m %s) %s(to %s) {", name, fn, toName)
			//	w.IndentLine(1, "")
			//	for _, field := range msg.Fields {
			//		if isHeaderField(field) {
			//			continue
			//		}
			//		if len(field.Fields) > 0 {
			//			for _, field := range field.Fields {
			//				w.IndentLine(1, "to.Set%s(m.%s())", field.Name, field.Name)
			//			}
			//		} else {
			//			w.IndentLine(1, "to.Set%s(m.%s())", field.Name, field.Name)
			//		}
			//	}
			//	w.Line("}")
			//	w.Line("")
			//}

			// Copy
			{
				write := func(msg *Struct) {
					copyTo(writer, msg, "Copy", msg.Name, msg.Name+"Builder")
					copyTo(writer, msg, "Copy", msg.Name+"Builder", msg.Name+"Builder")

					if msg.Message != nil {
						to := msg.Message.VLS
						if msg.VLS {
							to = msg.Message.Fixed
						}
						if to != nil {
							copyTo(writer, msg, "CopyTo", msg.Name, to.Name+"Builder")
							copyTo(writer, msg, "CopyTo", msg.Name+"Builder", to.Name+"Builder")
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
		}

		if err := g.writeFile(fmt.Sprintf("%s.rs", toSnakeCase(m.Name())), writer.b); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) writeFieldGetter(w *Writer, f *Field, suffix, ptrAccessor string) {
	w.Line("func (m %s%s) %s() %s {", f.Struct.Name, suffix, f.Name, g.RustTypeName(&f.Type))
	w.IndentLine(1, "return %s", g.fieldGetValue(f, ptrAccessor))
	w.Line("}")
	w.Line("")
}

func (g *Generator) fieldGetValue(f *Field, ptrAccessor string) string {
	if f.Field.Type.Alias != nil {
		return fmt.Sprintf("%s(%s)",
			g.RustTypeName(&f.Field.Type),
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
			g.RustTypeName(&f.Field.Type),
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

func (g *Generator) primitiveGetter(kind schema.Kind, offset, size int, messageVar, ptrAccessor string, vls bool) string {
	var (
		bounds = offset + size
		suffix = "Fixed"
		prefix = ""
	)
	if vls {
		suffix = "VLS"
	}
	switch kind {
	case schema.KindBool:
		prefix = "Bool"
	case schema.KindInt8:
		prefix = "Int8"
	case schema.KindInt16:
		prefix = "Int16"
	case schema.KindInt32:
		prefix = "Int32"
	case schema.KindInt64:
		prefix = "Int64"
	case schema.KindUint8:
		prefix = "Uint8"
	case schema.KindUint16:
		prefix = "Uint16"
	case schema.KindUint32:
		prefix = "Uint32"
	case schema.KindUint64:
		prefix = "Uint64"
	case schema.KindFloat32:
		prefix = "Float32"
	case schema.KindFloat64:
		prefix = "Float64"
	case schema.KindStringFixed:
		prefix = "String"
	case schema.KindStringVLS:
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
		w.Line("func (m *%s%s) Set%s(value %s) {", f.Struct.Name, suffix, f.Name, g.RustTypeName(&f.Type))
	} else {
		w.Line("func (m %s%s) Set%s(value %s) {", f.Struct.Name, suffix, f.Name, g.RustTypeName(&f.Type))
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

func (g *Generator) primitiveSetter(kind schema.Kind, offset, size int, messageVar, ptrAccessor, valueExpr string, vls bool) string {
	var (
		bounds = offset + size
		suffix = "Fixed"
		prefix = ""
	)
	if vls {
		suffix = "VLS"
	}
	switch kind {
	case schema.KindBool:
		prefix = "SetBool"
	case schema.KindInt8:
		prefix = "SetInt8"
	case schema.KindInt16:
		prefix = "SetInt16"
	case schema.KindInt32:
		prefix = "SetInt32"
	case schema.KindInt64:
		prefix = "SetInt64"
	case schema.KindUint8:
		prefix = "SetUint8"
	case schema.KindUint16:
		prefix = "SetUint16"
	case schema.KindUint32:
		prefix = "SetUint32"
	case schema.KindUint64:
		prefix = "SetUint64"
	case schema.KindFloat32:
		prefix = "SetFloat32"
	case schema.KindFloat64:
		prefix = "SetFloat64"
	case schema.KindStringFixed:
		prefix = "SetString"
	case schema.KindStringVLS:
		prefix = "SetString"
	}

	//messagePackage := ""
	//if pkg != g.root {
	//	messagePackage = pkg.Name + "."
	//}

	if kind == schema.KindStringVLS {
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
