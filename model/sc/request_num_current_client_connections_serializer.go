//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m RequestNumCurrentClientConnections) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m RequestNumCurrentClientConnectionsPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m RequestNumCurrentClientConnections) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m RequestNumCurrentClientConnectionsPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *RequestNumCurrentClientConnectionsBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
	m.SetDeviceIdentifier(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *RequestNumCurrentClientConnectionsPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
	m.SetDeviceIdentifier(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *RequestNumCurrentClientConnectionsBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
		case "DeviceIdentifier":
			m.SetDeviceIdentifier(r.Int64())
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

func (m *RequestNumCurrentClientConnectionsPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
		case "DeviceIdentifier":
			m.SetDeviceIdentifier(r.Int64())
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

func (m RequestNumCurrentClientConnectionsFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m RequestNumCurrentClientConnectionsFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10107,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.Int64(m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m RequestNumCurrentClientConnectionsFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m RequestNumCurrentClientConnectionsFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10107)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("Username", m.Username())
	w.Int64Field("DeviceIdentifier", m.DeviceIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *RequestNumCurrentClientConnectionsFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
	m.SetDeviceIdentifier(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *RequestNumCurrentClientConnectionsFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
	m.SetDeviceIdentifier(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *RequestNumCurrentClientConnectionsFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
		case "DeviceIdentifier":
			m.SetDeviceIdentifier(r.Int64())
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

func (m *RequestNumCurrentClientConnectionsFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10107 {
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
		case "DeviceIdentifier":
			m.SetDeviceIdentifier(r.Int64())
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
