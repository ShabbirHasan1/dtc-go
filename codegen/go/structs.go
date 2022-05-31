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
		w := &Writer{}
		if len(g.config.GeneratedComment) > 0 {
			w.Line("// %s", g.config.GeneratedComment)
			w.Line("")
		}

		w.Line("package %s", g.packageName)
		w.Line("")
		w.Line("import (")
		w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
		if g.config.Json {
			w.IndentLine(1, "\"github.com/moontrade/dtc-go/message/json\"")
		}
		w.IndentLine(1, "\"io\"")
		w.Line(")")
		w.Line("")

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

		if w != nil {
			writeMessageSize(w, m.VLS)
			writeMessageSize(w, m.Fixed)
			writeMessageDefault(w, m.VLS)
			writeMessageDefault(w, m.Fixed)
		}

		// Type declarations
		{
			write := func(msg *Struct) {
				if msg.Doc != nil {
					_ = g.writeComments(w, 0, msg.Name, msg.Doc.Description)
				}
				w.Line("type %s struct {", msg.Name)
				w.IndentLine(1, "p message.%s", msg.Suffix())
				w.Line("}")
				w.Line("")
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
				w.Line("func New%sFrom(b []byte) %s {", msg.Name, msg.Name)
				w.IndentLine(1, "return %s{p: message.New%s(b)}", msg.Name, msg.Suffix())
				w.Line("}")
				w.Line("")

				w.Line("func Wrap%s(b []byte) %s {", msg.Name, msg.Name)
				w.IndentLine(1, "return %s{p: message.Wrap%s(b)}", msg.Name, msg.Suffix())
				w.Line("}")
				w.Line("")

				w.Line("func New%s() *%s {", msg.Name, msg.Name)
				if msg.VLS {
					w.IndentLine(1, "return &%s{p: message.New%s(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
				} else {
					w.IndentLine(1, "return &%s{p: message.New%s(_%sDefault)}", msg.Name, msg.Suffix(), msg.Name)
				}
				w.Line("}")
				w.Line("")

				if msg.VLS {
					w.Line("func Parse%s(b []byte) (%s, error) {", msg.Name, msg.Name)
					w.IndentLine(1, "if len(b) < 6 {")
					w.IndentLine(2, "return %s{}, message.ErrShortBuffer", msg.Name)
					w.IndentLine(1, "}")

					w.IndentLine(1, "m := Wrap%s(b)", msg.Name)
					w.IndentLine(1, "if int(m.p.AsUint16LE()) != len(b) {")
					w.IndentLine(2, "return %s{}, message.ErrOverflow", msg.Name)
					w.IndentLine(1, "}")
					w.IndentLine(1, "baseSize := int(m.p.Uint16LE(4))")
					w.IndentLine(1, "if baseSize > len(b) {")
					w.IndentLine(2, "return %s{}, message.ErrBaseSizeOverflow", msg.Name)
					w.IndentLine(1, "}")
					w.IndentLine(1, "if baseSize < %d {", msg.Size)
					w.IndentLine(2, "newSize := len(b)+(%d-baseSize)", msg.Size)
					w.IndentLine(2, "if newSize > message.MaxSize {")
					w.IndentLine(3, "return %s{}, message.ErrOverflow", msg.Name)
					w.IndentLine(2, "}")
					w.IndentLine(2, "clone := %s{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}", msg.Name)
					w.IndentLine(2, "clone.p.SetBytes(0, b[0:baseSize])")
					w.IndentLine(2, "clone.p.SetBytes(baseSize, _%sDefault[baseSize:])", msg.Name)
					vlsFields := msg.Filter(func(field *Field) bool {
						return field.Type.Kind == schema.KindStringVLS
					})
					if len(vlsFields) > 0 {
						w.IndentLine(2, "if len(b) > baseSize {")
						w.IndentLine(3, "shift := uint16(%d - baseSize)", msg.Size)
						//w.IndentLine(3, "maxOffset := uint16(len(b)) + shift")
						w.IndentLine(3, "var offset uint16")
						for _, field := range vlsFields {
							w.IndentLine(3, "offset = clone.p.Uint16LE(%d)", field.Type.Offset)
							w.IndentLine(3, "if offset > 0 {")
							w.IndentLine(4, "clone.p.SetUint16LE(%d, offset+shift)", field.Type.Offset)
							w.IndentLine(3, "}")
						}
						w.IndentLine(2, "}")
					}
					w.IndentLine(2, "return clone, nil")
					w.IndentLine(1, "}")
					w.IndentLine(1, "return m, nil")
					w.Line("}")
					w.Line("")
				} else {
					w.Line("func Parse%s(b []byte) (%s, error) {", msg.Name, msg.Name)
					w.IndentLine(1, "if len(b) < 4 {")
					w.IndentLine(2, "return %s{}, message.ErrShortBuffer", msg.Name)
					w.IndentLine(1, "}")
					w.IndentLine(1, "m := Wrap%s(b)", msg.Name)
					w.IndentLine(1, "if int(m.p.AsUint16LE()) != len(b) {")
					w.IndentLine(2, "return %s{}, message.ErrOverflow", msg.Name)
					w.IndentLine(1, "}")
					w.IndentLine(1, "size := int(m.p.AsUint16LE())")
					w.IndentLine(1, "if size > len(b) {")
					w.IndentLine(2, "return %s{}, message.ErrBaseSizeOverflow", msg.Name)
					w.IndentLine(1, "}")
					w.IndentLine(1, "if size < %d {", msg.Size)
					w.IndentLine(2, "clone := *New%s()", msg.Name)
					w.IndentLine(2, "clone.p.SetBytes(0, b[0:size])")
					w.IndentLine(2, "clone.p.SetBytes(size, _%sDefault[size:])", msg.Name)
					w.IndentLine(2, "return clone, nil")
					w.IndentLine(2, "}")
					w.IndentLine(2, "return m, nil")
					w.Line("}")
					w.Line("")
				}
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		//// Clear
		//if w != nil {
		//	write := func(msg *Struct) {
		//		g.writeClearComment(w, msg, "Clear")
		//		w.Line("func (m %sBuilder) Clear() {", msg.Name)
		//		w.IndentLine(1, "m.Unsafe().SetBytes(0, _%sDefault)", msg.Name)
		//		w.Line("}")
		//		w.Line("")
		//	}
		//	if m.VLS != nil {
		//		write(m.VLS)
		//	}
		//	if m.Fixed != nil {
		//		write(m.Fixed)
		//	}
		//}
		//if nogcWriter != nil {
		//	write := func(msg *Struct) {
		//		g.writeClearComment(nogcWriter, msg, "Clear")
		//		nogcWriter.Line("func (m %sPointerBuilder) Clear() {", msg.Name)
		//		nogcWriter.IndentLine(1, "m.p.SetBytes(0, _%sDefault)", msg.Name)
		//		nogcWriter.Line("}")
		//		nogcWriter.Line("")
		//	}
		//	if m.VLS != nil {
		//		write(m.VLS)
		//	}
		//	if m.Fixed != nil {
		//		write(m.Fixed)
		//	}
		//}

		// Getters
		writeGetters := func(msg *Struct) {
			for _, f := range msg.Fields {
				if f.Fields != nil {
					for _, field := range f.Fields {
						if w != nil {
							if field.Doc != nil {
								_ = g.writeComments(w, 0, field.Name, field.Doc.Description)
							} else {
								w.Line("// %s", field.Name)
							}
							g.writeFieldGetter(w, field, "", "p")
						}
					}
				} else {
					if w != nil {
						if f.Doc != nil {
							_ = g.writeComments(w, 0, f.Name, f.Doc.Description)
						} else {
							w.Line("// %s", f.Name)
						}
						g.writeFieldGetter(w, f, "", "p")
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
							if w != nil {
								if field.Doc != nil {
									_ = g.writeComments(w, 0, "Set"+field.Name, field.Doc.Description)
								} else {
									w.Line("// Set%s", field.Name)
								}
								g.writeFieldSetter(w, field, "", "p")
							}
						} else {
							if w != nil {
								if field.Doc != nil {
									_ = g.writeComments(w, 0, "Set"+field.Name, field.Doc.Description)
								} else {
									w.Line("// Set%s", field.Name)
								}
								g.writeFieldSetter(w, field, "", "p")
							}
						}
					}
				} else {
					if f.Field.Type.Kind == schema.KindStringVLS {
						if w != nil {
							if f.Doc != nil {
								_ = g.writeComments(w, 0, "Set"+f.Name, f.Doc.Description)
							} else {
								w.Line("// Set%s", f.Name)
							}
							g.writeFieldSetter(w, f, "", "p")
						}
					} else {
						if w != nil {
							if f.Doc != nil {
								_ = g.writeComments(w, 0, "Set"+f.Name, f.Doc.Description)
							} else {
								w.Line("// Set%s", f.Name)
							}
							g.writeFieldSetter(w, f, "", "p")
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

		{
			// Write
			write := func(msg *Struct) {
				w.Line("func (m *%s) WriteTo(w io.Writer) (int64, error) {", msg.Name)
				w.IndentLine(1, "s := int(m.Size())")
				w.IndentLine(1, "n, err := w.Write(m.p.AsBytes(s))")
				w.IndentLine(1, "return int64(n), err")
				w.Line("}")
				w.Line("")

				w.Line("func (m *%s) MarshalBinary() ([]byte, error) {", msg.Name)
				w.IndentLine(1, "return m.p.AsBytes(int(m.Size())), nil")
				w.Line("}")
				w.Line("")

				w.Line("func (m *%s) Clone() *%s {", msg.Name, msg.Name)
				if msg.VLS {
					w.IndentLine(1, "return &%s{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}", msg.Name)
				} else {
					w.IndentLine(1, "return &%s{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}", msg.Name)
				}
				w.Line("}")
				w.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
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

			// Copy
			write := func(msg *Struct) {
				copyTo(w, msg, "Copy", msg.Name, msg.Name)

				if msg.Message != nil {
					to := msg.Message.VLS
					if msg.VLS {
						to = msg.Message.Fixed
					}
					if to != nil {
						copyTo(w, msg, "CopyTo", msg.Name, to.Name)
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

		if err := g.generateJson(m.VLS, w); err != nil {
			return err
		}
		if err := g.generateJson(m.Fixed, w); err != nil {
			return err
		}

		if err := g.writeFile(fmt.Sprintf("%s.go", toSnakeCase(m.Name())), w.b); err != nil {
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
	if f.Field.Type.Alias != nil {
		w.IndentLine(1, "return %s(%s)",
			g.GoTypeName(&f.Field.Type),
			g.primitiveGetter(
				f.Field.Type.Alias.Type.Kind,
				f.Field.Type.Offset,
				f.Field.Type.Size,
				f.Struct.VLS,
			))
	} else if f.Field.Type.Enum != nil {
		w.IndentLine(1, "return %s(%s)",
			g.GoTypeName(&f.Field.Type),
			g.primitiveGetter(
				f.Field.Type.Enum.Type,
				f.Field.Type.Offset,
				f.Field.Type.Size,
				f.Struct.VLS,
			))
	} else {
		w.IndentLine(1, "return %s",
			g.primitiveGetter(f.Field.Type.Kind,
				f.Field.Type.Offset,
				f.Field.Type.Size,
				f.Struct.VLS,
			))
	}
	w.Line("}")
	w.Line("")
}

func (g *Generator) primitiveGetter(kind schema.Kind, offset, size int, vls bool) string {
	switch kind {
	case schema.KindBool:
		return fmt.Sprintf("m.p.Bool(%d)", offset)
	case schema.KindInt8:
		return fmt.Sprintf("m.p.Int8(%d)", offset)
	case schema.KindInt16:
		return fmt.Sprintf("m.p.Int16LE(%d)", offset)
	case schema.KindInt32:
		return fmt.Sprintf("m.p.Int32LE(%d)", offset)
	case schema.KindInt64:
		return fmt.Sprintf("m.p.Int64LE(%d)", offset)
	case schema.KindUint8:
		return fmt.Sprintf("m.p.Uint8(%d)", offset)
	case schema.KindUint16:
		return fmt.Sprintf("m.p.Uint16LE(%d)", offset)
	case schema.KindUint32:
		return fmt.Sprintf("m.p.Uint32LE(%d)", offset)
	case schema.KindUint64:
		return fmt.Sprintf("m.p.Uint64LE(%d)", offset)
	case schema.KindFloat32:
		return fmt.Sprintf("m.p.Float32LE(%d)", offset)
	case schema.KindFloat64:
		return fmt.Sprintf("m.p.Float64LE(%d)", offset)
	case schema.KindStringFixed:
		return fmt.Sprintf("m.p.StringFixed(%d, %d)", offset, size)
	case schema.KindStringVLS:
		return fmt.Sprintf("m.p.StringVLS(%d)", offset)
	}

	panic("unknown type")
}

func (g *Generator) writeFieldSetter(w *Writer, f *Field, suffix, ptrAccessor string) {
	if f.Field.Type.Kind.IsString() && f.Struct.VLS {
		w.Line("func (m *%s%s) Set%s(value %s) *%s {", f.Struct.Name, suffix, f.Name, g.GoTypeName(&f.Type), f.Struct.Name)
	} else {
		w.Line("func (m *%s%s) Set%s(value %s) *%s {", f.Struct.Name, suffix, f.Name, g.GoTypeName(&f.Type), f.Struct.Name)
	}
	w.IndentLine(1, "%s", g.fieldSetValue(f, ptrAccessor))
	w.IndentLine(1, "return m")
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
	switch kind {
	case schema.KindBool:
		return fmt.Sprintf("m.p.SetBool(%d, %s)", offset, valueExpr)
	case schema.KindInt8:
		return fmt.Sprintf("m.p.SetInt8(%d, %s)", offset, valueExpr)
	case schema.KindInt16:
		return fmt.Sprintf("m.p.SetInt16LE(%d, %s)", offset, valueExpr)
	case schema.KindInt32:
		return fmt.Sprintf("m.p.SetInt32LE(%d, %s)", offset, valueExpr)
	case schema.KindInt64:
		return fmt.Sprintf("m.p.SetInt64LE(%d, %s)", offset, valueExpr)
	case schema.KindUint8:
		return fmt.Sprintf("m.p.SetUint8(%d, %s)", offset, valueExpr)
	case schema.KindUint16:
		return fmt.Sprintf("m.p.SetUint16LE(%d, %s)", offset, valueExpr)
	case schema.KindUint32:
		return fmt.Sprintf("m.p.SetUint32LE(%d, %s)", offset, valueExpr)
	case schema.KindUint64:
		return fmt.Sprintf("m.p.SetUint64LE(%d, %s)", offset, valueExpr)
	case schema.KindFloat32:
		return fmt.Sprintf("m.p.SetFloat32LE(%d, %s)", offset, valueExpr)
	case schema.KindFloat64:
		return fmt.Sprintf("m.p.SetFloat64LE(%d, %s)", offset, valueExpr)
	case schema.KindStringFixed:
		return fmt.Sprintf("m.p.SetStringFixed(%d, %d, %s)", offset, size, valueExpr)
	case schema.KindStringVLS:
		return fmt.Sprintf("m.p.SetStringVLS(%d, %s)", offset, valueExpr)
	}
	return "<UNKNOWN TYPE>"
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
