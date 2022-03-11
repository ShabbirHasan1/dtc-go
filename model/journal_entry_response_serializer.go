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

func (m JournalEntryResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 706 {
		return message.ErrWrongType
	}
	m.SetJournalEntry(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 706 {
		return message.ErrWrongType
	}
	m.SetJournalEntry(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 706 {
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
		case "JournalEntry":
			m.SetJournalEntry(r.String())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
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

func (m *JournalEntryResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 706 {
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
		case "JournalEntry":
			m.SetJournalEntry(r.String())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
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

func (m JournalEntryResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":706,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

func (m JournalEntryResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 706 {
		return message.ErrWrongType
	}
	m.SetJournalEntry(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 706 {
		return message.ErrWrongType
	}
	m.SetJournalEntry(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 706 {
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
		case "JournalEntry":
			m.SetJournalEntry(r.String())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
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

func (m *JournalEntryResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 706 {
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
		case "JournalEntry":
			m.SetJournalEntry(r.String())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
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

func (m *JournalEntryResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetJournalEntry(r.ReadString())
		case 2:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 3:
			m.SetIsFinalResponse(r.ReadBool())
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

func (m *JournalEntryResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetJournalEntry(r.ReadString())
		case 2:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 3:
			m.SetIsFinalResponse(r.ReadBool())
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

func (m JournalEntryResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 706)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	w.WriteBool(3, m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 706)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	w.WriteBool(3, m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetJournalEntry(r.ReadString())
		case 2:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 3:
			m.SetIsFinalResponse(r.ReadBool())
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

func (m *JournalEntryResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetJournalEntry(r.ReadString())
		case 2:
			m.SetDateTime(DateTime(r.ReadInt64()))
		case 3:
			m.SetIsFinalResponse(r.ReadBool())
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

func (m JournalEntryResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 706)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	w.WriteBool(3, m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 706)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	w.WriteBool(3, m.IsFinalResponse())
	return w.Finish(), nil
}