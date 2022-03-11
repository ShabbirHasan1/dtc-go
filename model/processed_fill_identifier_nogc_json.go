//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ProcessedFillIdentifierPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10132,\"F\":["...)
	w.String(m.FillIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m ProcessedFillIdentifierPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10132,\"F\":["...)
	w.String(m.FillIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ProcessedFillIdentifierPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10132)
	w.StringField("FillIdentifier", m.FillIdentifier())
	w.Uint8Field("IsSnapshot", m.IsSnapshot())
	w.Uint8Field("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m ProcessedFillIdentifierPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10132)
	w.StringField("FillIdentifier", m.FillIdentifier())
	w.Uint8Field("IsSnapshot", m.IsSnapshot())
	w.Uint8Field("IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
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
	m.SetIsSnapshot(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
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
			m.SetIsSnapshot(r.Uint8())
		case "IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Uint8())
		case "IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Uint8())
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
