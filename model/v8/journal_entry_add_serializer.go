//go:build !tinygo

// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAdd) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAdd) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 703 {
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddPointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 703 {
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 703 {
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

func (m *JournalEntryAddPointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 703 {
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

func (m *JournalEntryAddBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddPointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":703,\"F\":["...)
	w.String(m.JournalEntry())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

func (m JournalEntryAddFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 703)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddFixedBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 703 {
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddFixedPointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 703 {
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
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddFixedBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 703 {
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

func (m *JournalEntryAddFixedPointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 703 {
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

func (m *JournalEntryAddFixedBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddFixedPointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddBuilder) UnmarshalProtobuf(b []byte) error {
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

func (m *JournalEntryAddPointerBuilder) UnmarshalProtobuf(b []byte) error {
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

func (m JournalEntryAddBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 703)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 703)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryAddFixedBuilder) UnmarshalProtobuf(b []byte) error {
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

func (m *JournalEntryAddFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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

func (m JournalEntryAddFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 703)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryAddFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 703)
	w.WriteString(1, m.JournalEntry())
	w.WriteVarint64(2, int64(m.DateTime()))
	return w.Finish(), nil
}
