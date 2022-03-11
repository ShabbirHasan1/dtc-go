//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m NumCurrentClientConnectionsResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m NumCurrentClientConnectionsResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m NumCurrentClientConnectionsResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m NumCurrentClientConnectionsResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *NumCurrentClientConnectionsResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10108 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForDifferentDevices(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForSameDevice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *NumCurrentClientConnectionsResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10108 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForDifferentDevices(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForSameDevice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *NumCurrentClientConnectionsResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10108 {
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
		case "RequestID":
			m.SetRequestID(r.Uint32())
		case "Username":
			m.SetUsername(r.String())
		case "NumConnectionsForDifferentDevices":
			m.SetNumConnectionsForDifferentDevices(r.Int32())
		case "NumConnectionsForSameDevice":
			m.SetNumConnectionsForSameDevice(r.Int32())
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

func (m *NumCurrentClientConnectionsResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10108 {
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
		case "RequestID":
			m.SetRequestID(r.Uint32())
		case "Username":
			m.SetUsername(r.String())
		case "NumConnectionsForDifferentDevices":
			m.SetNumConnectionsForDifferentDevices(r.Int32())
		case "NumConnectionsForSameDevice":
			m.SetNumConnectionsForSameDevice(r.Int32())
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

func (m NumCurrentClientConnectionsResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m NumCurrentClientConnectionsResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10108,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForDifferentDevices())
	w.Data = append(w.Data, ',')
	w.Int32(m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m NumCurrentClientConnectionsResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m NumCurrentClientConnectionsResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10108)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int32Field("NumConnectionsForDifferentDevices", m.NumConnectionsForDifferentDevices())
	w.Int32Field("NumConnectionsForSameDevice", m.NumConnectionsForSameDevice())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *NumCurrentClientConnectionsResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10108 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForDifferentDevices(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForSameDevice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *NumCurrentClientConnectionsResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10108 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForDifferentDevices(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumConnectionsForSameDevice(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *NumCurrentClientConnectionsResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10108 {
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
		case "RequestID":
			m.SetRequestID(r.Uint32())
		case "Username":
			m.SetUsername(r.String())
		case "NumConnectionsForDifferentDevices":
			m.SetNumConnectionsForDifferentDevices(r.Int32())
		case "NumConnectionsForSameDevice":
			m.SetNumConnectionsForSameDevice(r.Int32())
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

func (m *NumCurrentClientConnectionsResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10108 {
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
		case "RequestID":
			m.SetRequestID(r.Uint32())
		case "Username":
			m.SetUsername(r.String())
		case "NumConnectionsForDifferentDevices":
			m.SetNumConnectionsForDifferentDevices(r.Int32())
		case "NumConnectionsForSameDevice":
			m.SetNumConnectionsForSameDevice(r.Int32())
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
