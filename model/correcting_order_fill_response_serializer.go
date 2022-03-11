//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 310 {
		return message.ErrWrongType
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsError(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 310 {
		return message.ErrWrongType
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsError(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 310 {
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
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "ResultText":
			m.SetResultText(r.String())
		case "IsError":
			m.SetIsError(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 310 {
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
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "ResultText":
			m.SetResultText(r.String())
		case "IsError":
			m.SetIsError(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":310,\"F\":["...)
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ResultText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

func (m CorrectingOrderFillResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 310)
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("ResultText", m.ResultText())
	w.BoolField("IsError", m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 310 {
		return message.ErrWrongType
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsError(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 310 {
		return message.ErrWrongType
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetResultText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsError(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 310 {
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
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "ResultText":
			m.SetResultText(r.String())
		case "IsError":
			m.SetIsError(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 310 {
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
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "ResultText":
			m.SetResultText(r.String())
		case "IsError":
			m.SetIsError(r.Bool())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetClientOrderID(r.ReadString())
		case 2:
			m.SetResultText(r.ReadString())
		case 3:
			m.SetIsError(r.ReadBool())
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
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetClientOrderID(r.ReadString())
		case 2:
			m.SetResultText(r.ReadString())
		case 3:
			m.SetIsError(r.ReadBool())
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

func (m CorrectingOrderFillResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 310)
	w.WriteString(1, m.ClientOrderID())
	w.WriteString(2, m.ResultText())
	w.WriteBool(3, m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 310)
	w.WriteString(1, m.ClientOrderID())
	w.WriteString(2, m.ResultText())
	w.WriteBool(3, m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetClientOrderID(r.ReadString())
		case 2:
			m.SetResultText(r.ReadString())
		case 3:
			m.SetIsError(r.ReadBool())
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
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CorrectingOrderFillResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetClientOrderID(r.ReadString())
		case 2:
			m.SetResultText(r.ReadString())
		case 3:
			m.SetIsError(r.ReadBool())
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

func (m CorrectingOrderFillResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 310)
	w.WriteString(1, m.ClientOrderID())
	w.WriteString(2, m.ResultText())
	w.WriteBool(3, m.IsError())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CorrectingOrderFillResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 310)
	w.WriteString(1, m.ClientOrderID())
	w.WriteString(2, m.ResultText())
	w.WriteBool(3, m.IsError())
	return w.Finish(), nil
}
