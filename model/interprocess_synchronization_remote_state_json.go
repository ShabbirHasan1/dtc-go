//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationRemoteState) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10134,\"F\":["...)
	w.Uint8(m.IsPrimary())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSecondary())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationRemoteStateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10134,\"F\":["...)
	w.Uint8(m.IsPrimary())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSecondary())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationRemoteState) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10134)
	w.Uint8Field("IsPrimary", m.IsPrimary())
	w.Uint8Field("IsSecondary", m.IsSecondary())
	w.Uint8Field("IsSecondaryFailoverActive", m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationRemoteStateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10134)
	w.Uint8Field("IsPrimary", m.IsPrimary())
	w.Uint8Field("IsSecondary", m.IsSecondary())
	w.Uint8Field("IsSecondaryFailoverActive", m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationRemoteStateBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10134 {
		return message.ErrWrongType
	}
	m.SetIsPrimary(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSecondary(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSecondaryFailoverActive(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationRemoteStateBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10134 {
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
		case "IsPrimary":
			m.SetIsPrimary(r.Uint8())
		case "IsSecondary":
			m.SetIsSecondary(r.Uint8())
		case "IsSecondaryFailoverActive":
			m.SetIsSecondaryFailoverActive(r.Uint8())
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
