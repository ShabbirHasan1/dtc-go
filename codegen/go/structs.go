package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen/schema"
	"github.com/moontrade/dtc-go/message/pb"
	"strings"
)

func (g *Generator) generateStructDebug(msg *Struct) error {
	return nil
}

func (g *Generator) generateStructs() error {
	for _, m := range g.messages {
		var (
			gcWriter   *Writer
			nogcWriter *Writer
		)
		if g.config.GC {
			gcWriter = &Writer{}
			if len(g.config.GeneratedComment) > 0 {
				gcWriter.Line("// %s", g.config.GeneratedComment)
				gcWriter.Line("")
			}
		}
		if g.config.NoGC {
			if gcWriter == nil {
				nogcWriter = &Writer{}
				if len(g.config.GeneratedComment) > 0 {
					nogcWriter.Line("// %s", g.config.GeneratedComment)
					nogcWriter.Line("")
				}
			} else {
				nogcWriter = gcWriter
			}
		}

		writeImports := func(w *Writer) {
			w.Line("package %s", g.packageName)
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
				gcWriter.IndentLine(1, "return %s{%s: message.New%s(b)}", msg.Name, msg.Suffix(), msg.Suffix())
				gcWriter.Line("}")
				gcWriter.Line("")

				gcWriter.Line("func Wrap%s(b []byte) %s {", msg.Name, msg.Name)
				gcWriter.IndentLine(1, "return %s{%s: message.Wrap%s(b)}", msg.Name, msg.Suffix(), msg.Suffix())
				gcWriter.Line("}")
				gcWriter.Line("")

				gcWriter.Line("func New%s() %sBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					gcWriter.IndentLine(1, "return %sBuilder{message.New%sBuilder(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
				} else {
					gcWriter.IndentLine(1, "return %sBuilder{message.New%s(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
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
				nogcWriter.Line("func Alloc%s() %sPointerBuilder {", msg.Name, msg.Name)
				if msg.VLS {
					nogcWriter.IndentLine(1, "return %sPointerBuilder{message.Alloc%sBuilder(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
				} else {
					nogcWriter.IndentLine(1, "return %sPointerBuilder{message.Alloc%s(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
				}
				nogcWriter.Line("}")
				nogcWriter.Line("")

				/*
					func AllocAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustmentPointer {
						return AccountBalanceAdjustmentPointer{VLSPointer: message.AllocVLSFrom(b)}
					}
				*/
				nogcWriter.Line("func Alloc%sFrom(b []byte) %sPointer {", msg.Name, msg.Name)
				nogcWriter.IndentLine(1, "return %sPointer{%sPointer: message.Alloc%s(b)}", msg.Name, msg.Suffix(), msg.Suffix())
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
						if field.Field.Type.Kind == schema.KindStringVLS {
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

					if f.Field.Type.Kind == schema.KindStringVLS {
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

func protobufMarshalField(f *Field, getterExpression string) string {
	kind := f.Type.Kind
	switch kind {
	case schema.KindEnum:
		kind = f.Type.Enum.Type
		getterExpression = fmt.Sprintf("%s(%s)", primitiveTypeName(kind), getterExpression)
	case schema.KindAlias:
		kind = f.Type.Alias.Type.Kind
		getterExpression = fmt.Sprintf("%s(%s)", primitiveTypeName(kind), getterExpression)
	}
	switch kind {
	case schema.KindBool:
		return fmt.Sprintf("WriteBool(%s, %s)", f.Proto.FieldNumber, getterExpression)
	case schema.KindInt8:
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
	case schema.KindInt16:
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
	case schema.KindInt32:
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
	case schema.KindInt64:
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
	case schema.KindUint8:
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
	case schema.KindUint16:
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
	case schema.KindUint32:
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
	case schema.KindUint64:
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
	case schema.KindFloat32:
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
	case schema.KindFloat64:
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
	case schema.KindStringFixed, schema.KindStringVLS:
		return fmt.Sprintf("WriteString(%s, %s)", f.Proto.FieldNumber, getterExpression)
	default:
		return "Invalid"
	}
}

func jsonGetterFuncNamePrimitive(kind schema.Kind) string {
	switch kind {
	case schema.KindBool:
		return "Bool"
	case schema.KindInt8:
		return "Int8"
	case schema.KindInt16:
		return "Int16"
	case schema.KindInt32:
		return "Int32"
	case schema.KindInt64:
		return "Int64"
	case schema.KindUint8:
		return "Uint8"
	case schema.KindUint16:
		return "Uint16"
	case schema.KindUint32:
		return "Uint32"
	case schema.KindUint64:
		return "Uint64"
	case schema.KindFloat32:
		return "Float32"
	case schema.KindFloat64:
		return "Float64"
	case schema.KindStringFixed, schema.KindStringVLS:
		return "String"
	default:
		return "Invalid"
	}
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
