//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ProcessedFillIdentifier) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10132,\"F\":["...)
	w.String(m.FillIdentifier())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m ProcessedFillIdentifierBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10132,\"F\":["...)
	w.String(m.FillIdentifier())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ProcessedFillIdentifierPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10132,\"F\":["...)
	w.String(m.FillIdentifier())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m ProcessedFillIdentifierPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10132,\"F\":["...)
	w.String(m.FillIdentifier())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ProcessedFillIdentifier) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10132)
	w.StringField("FillIdentifier", m.FillIdentifier())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m ProcessedFillIdentifierBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10132)
	w.StringField("FillIdentifier", m.FillIdentifier())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ProcessedFillIdentifierPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10132)
	w.StringField("FillIdentifier", m.FillIdentifier())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m ProcessedFillIdentifierPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10132)
	w.StringField("FillIdentifier", m.FillIdentifier())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.BoolField("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ProcessedFillIdentifierBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10132 {
		return message.ErrWrongType
	}
	m.SetFillIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ProcessedFillIdentifierPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10132 {
		return message.ErrWrongType
	}
	m.SetFillIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ProcessedFillIdentifierBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10132 {
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
		case "FillIdentifier":
			m.SetFillIdentifier(r.String())
		case "IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Bool())
		case "IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
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

func (m *ProcessedFillIdentifierPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10132 {
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
		case "FillIdentifier":
			m.SetFillIdentifier(r.String())
		case "IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Bool())
		case "IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
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
