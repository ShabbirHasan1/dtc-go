package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen/schema"
)

func (g *Generator) generateJson(msg *Struct, writer *Writer) error {
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

	if g.config.JsonCompact && writer != nil {
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("// JSON Compact Marshal")
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("")

		marshalJSONCompact(writer, msg.Name)
	}

	marshalJSON := func(w *Writer, name string) {
		w.Line("func (m *%s) MarshalJSON() ([]byte, error) {", name)
		w.IndentLine(1, "return m.MarshalJSONTo(nil)")
		w.Line("}")
		w.Line("")
		w.Line("func (m *%s) MarshalJSONTo(b []byte) ([]byte, error) {", name)
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

	if g.config.Json && writer != nil {
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("// JSON Marshal")
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("")
		marshalJSON(writer, msg.Name)
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

	if g.config.JsonCompact && writer != nil {
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("// JSON Compact Unmarshal")
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("")
		unmarshalJSONCompact(writer, msg.Name)
	}

	unmarshalJSONFrom := func(w *Writer, name string) {
		w.Line("func (m *%s) UnmarshalJSON(b []byte) error {", name)
		w.IndentLine(1, "r, err := json.OpenReader(b)")
		w.IndentLine(1, "if err != nil {")
		w.IndentLine(2, "return err")
		w.IndentLine(1, "}")
		w.IndentLine(1, "return m.UnmarshalJSONFromReader(&r)")
		w.Line("}")
		w.Line("")
		w.Line("func (m *%s) UnmarshalJSONFromReader(r *json.Reader) error {", name)
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

	if g.config.Json && writer != nil {
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("// JSON Unmarshal")
		writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
		writer.Line("")

		unmarshalJSONFrom(writer, msg.Name)
	}

	//unmarshalJSON := func(w *Writer, name string) {
	//	w.Line("func (m *%s) UnmarshalJSONReader(r *json.Reader) error {", name)
	//	if g.config.JsonCompact {
	//		w.IndentLine(1, "if r.IsCompact {")
	//		w.IndentLine(2, "return m.UnmarshalJSONCompact(r)")
	//		w.IndentLine(1, "}")
	//	}
	//	w.IndentLine(2, "return m.UnmarshalJSONDoc(r)")
	//	w.Line("}")
	//	w.Line("")
	//}
	//if g.config.Json && writer != nil {
	//	writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
	//	writer.Line("// JSON Unmarshal")
	//	writer.Line("//////////////////////////////////////////////////////////////////////////////////////////")
	//	writer.Line("")
	//
	//	unmarshalJSON(writer, msg.Name)
	//}
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
