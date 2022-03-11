//go:build !tinygo

package sc

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
	w.Bool(m.IsPrimary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationRemoteStateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10134,\"F\":["...)
	w.Bool(m.IsPrimary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationRemoteStatePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10134,\"F\":["...)
	w.Bool(m.IsPrimary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationRemoteStatePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10134,\"F\":["...)
	w.Bool(m.IsPrimary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondary())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationRemoteState) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10134)
	w.BoolField("IsPrimary", m.IsPrimary())
	w.BoolField("IsSecondary", m.IsSecondary())
	w.BoolField("IsSecondaryFailoverActive", m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationRemoteStateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10134)
	w.BoolField("IsPrimary", m.IsPrimary())
	w.BoolField("IsSecondary", m.IsSecondary())
	w.BoolField("IsSecondaryFailoverActive", m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m InterprocessSynchronizationRemoteStatePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10134)
	w.BoolField("IsPrimary", m.IsPrimary())
	w.BoolField("IsSecondary", m.IsSecondary())
	w.BoolField("IsSecondaryFailoverActive", m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

func (m InterprocessSynchronizationRemoteStatePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10134)
	w.BoolField("IsPrimary", m.IsPrimary())
	w.BoolField("IsSecondary", m.IsSecondary())
	w.BoolField("IsSecondaryFailoverActive", m.IsSecondaryFailoverActive())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationRemoteStateBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10134 {
		return message.ErrWrongType
	}
	m.SetIsPrimary(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSecondary(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSecondaryFailoverActive(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *InterprocessSynchronizationRemoteStatePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10134 {
		return message.ErrWrongType
	}
	m.SetIsPrimary(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSecondary(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSecondaryFailoverActive(r.Bool())
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
			m.SetIsPrimary(r.Bool())
		case "IsSecondary":
			m.SetIsSecondary(r.Bool())
		case "IsSecondaryFailoverActive":
			m.SetIsSecondaryFailoverActive(r.Bool())
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

func (m *InterprocessSynchronizationRemoteStatePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetIsPrimary(r.Bool())
		case "IsSecondary":
			m.SetIsSecondary(r.Bool())
		case "IsSecondaryFailoverActive":
			m.SetIsSecondaryFailoverActive(r.Bool())
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
