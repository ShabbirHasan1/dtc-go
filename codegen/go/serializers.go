package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen/schema"
	"os"
)

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
			case schema.KindEnum:
				w.IndentLine(1, "w.%s(%s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Enum.Type), primitiveTypeName(field.Type.Enum.Type), field.Name)
			case schema.KindAlias:
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
					case schema.KindEnum:
						w.IndentLine(1, "w.%sField(\"%s\", %s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Enum.Type), field.Field.Name, primitiveTypeName(field.Type.Enum.Type), field.Name)
					case schema.KindAlias:
						w.IndentLine(1, "w.%sField(\"%s\", %s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind), field.Field.Name, primitiveTypeName(field.Type.Alias.Type.Kind), field.Name)
					default:
						w.IndentLine(1, "w.%sField(\"%s\", m.%s())", jsonGetterFuncNamePrimitive(field.Type.Kind), field.Field.Name, field.Name)
					}
				}
			} else {
				switch field.Type.Kind {
				case schema.KindEnum:
					w.IndentLine(1, "w.%sField(\"%s\", %s(m.%s()))", jsonGetterFuncNamePrimitive(field.Type.Enum.Type), field.Field.Name, primitiveTypeName(field.Type.Enum.Type), field.Name)
				case schema.KindAlias:
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
		w.Line("func (m *%s) UnmarshalJSONCompact(r *json.Reader) error {", name)
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
			case schema.KindEnum:
				w.IndentLine(1, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type))
			case schema.KindAlias:
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

	unmarshalJSONFrom := func(w *Writer, name string) {
		w.Line("func (m *%s) UnmarshalJSONDoc(r *json.Reader) error {", name)
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
					case schema.KindEnum:
						w.IndentLine(3, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type))
					case schema.KindAlias:
						w.IndentLine(3, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind))
					default:
						w.IndentLine(3, "m.Set%s(r.%s())", field.Name, jsonGetterFuncNamePrimitive(field.Type.Kind))
					}
				}
			} else {
				w.IndentLine(2, "case \"%s\":", field.Field.Name)
				switch field.Type.Kind {
				case schema.KindEnum:
					w.IndentLine(3, "m.Set%s(%s(r.%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type))
				case schema.KindAlias:
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

		unmarshalJSONFrom(gcWriter, msg.Name+"Builder")
	}
	if nogcWriter != nil {
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("// JSON Unmarshal")
		nogcWriter.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		nogcWriter.Line("")

		unmarshalJSONFrom(nogcWriter, msg.Name+"PointerBuilder")
	}

	unmarshalJSON := func(w *Writer, name string) {
		w.Line("func (m *%s) UnmarshalJSONReader(r *json.Reader) error {", name)
		w.IndentLine(1, "if r.IsCompact {")
		w.IndentLine(2, "return m.UnmarshalJSONCompact(r)")
		w.IndentLine(1, "}")
		w.IndentLine(2, "return m.UnmarshalJSONDoc(r)")
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

	var fields FieldsByProtoNumber
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
					case schema.KindEnum:
						w.IndentLine(3, "m.Set%s(%s(r.Read%s%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type), zigzag)
					case schema.KindAlias:
						w.IndentLine(3, "m.Set%s(%s(r.Read%s%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Alias.Type.Kind), zigzag)
					default:
						w.IndentLine(3, "m.Set%s(r.Read%s%s())", field.Name, jsonGetterFuncNamePrimitive(field.Type.Kind), zigzag)
					}
				}
			} else {
				switch field.Type.Kind {
				case schema.KindEnum:
					w.IndentLine(3, "m.Set%s(%s(r.Read%s%s()))", field.Name, g.GoTypeName(&field.Type), jsonGetterFuncNamePrimitive(field.Type.Enum.Type), zigzag)
				case schema.KindAlias:
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
