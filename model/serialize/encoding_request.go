package serialize

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
	"github.com/moontrade/dtc-go/model"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func EncodingRequestUnmarshalJSONCompact(r *json.Reader, to model.EncodingRequestBuilder) error {
	if r.Type != 6 {
		return message.ErrWrongType
	}
	to.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	r.Lexer.WantComma()

	to.SetEncoding(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	r.Lexer.WantComma()

	to.SetProtocolType(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

func EncodingRequestFromJSON(r *json.Reader, to model.EncodingRequestBuilder) error {
	if r.Type != 6 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "ProtocolVersion":
			to.SetProtocolVersion(r.Int32())
		case "Encoding":
			to.SetEncoding(r.Int32())
		case "ProtocolType":
			to.SetProtocolType(r.String())
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func EncodingRequestMarshalJSONCompact(m model.EncodingRequest, b []byte) ([]byte, error) {
	w := json.NewCompactWriterI32(b, 6, m.ProtocolVersion())
	w.Int32Compact(m.Encoding()).StringCompact(m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func EncodingRequestMarshalJSON(m model.EncodingRequest, b []byte) ([]byte, error) {
	w := json.NewWriter(b, m.Type())
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.Int32Field("Encoding", m.Encoding())
	w.StringField("ProtocolType", m.ProtocolType())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func EncodingRequestUnmarshalProtobuf(b []byte, builder model.EncodingRequestBuilder) error {
	var (
		r   = pb.NewBuffer(b)
		f   pb.Number
		err error
	)
	for !r.EOF() {
		f, _, err = r.DecodeTag()
		if err != nil {
			break
		}
		switch f {
		case 1:
			builder.SetProtocolVersion(r.ReadInt32())
		case 2:
			builder.SetEncoding(r.ReadInt32())
		case 3:
			builder.SetProtocolType(r.ReadString())
		default:
			r.Consume()
		}
		if err = r.Error(); err != nil {
			return err
		}
	}
	return err
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func EncodingRequestMarshalProtobuf(m model.EncodingRequest, b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 6)
	/*
		message EncodingRequest
		{
		  int32 ProtocolVersion = 1;
		  EncodingEnum Encoding = 2;
		  string ProtocolType = 3;
		}
	*/
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteVarint32(2, m.Encoding())
	w.WriteString(3, m.ProtocolType())
	return w.Finish(), nil
}
