//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SCConfigurationSynchronization) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10109,\"F\":["...)
	w.Uint32(m.MessageID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentInboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentOutboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint64(m.CurrentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.String(m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

func (m SCConfigurationSynchronizationBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10109,\"F\":["...)
	w.Uint32(m.MessageID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentInboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentOutboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint64(m.CurrentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.String(m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SCConfigurationSynchronizationPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10109,\"F\":["...)
	w.Uint32(m.MessageID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentInboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentOutboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint64(m.CurrentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.String(m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

func (m SCConfigurationSynchronizationPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10109,\"F\":["...)
	w.Uint32(m.MessageID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentInboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentOutboundSequenceNumber())
	w.Data = append(w.Data, ',')
	w.Uint64(m.CurrentInternalOrderID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.String(m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SCConfigurationSynchronization) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10109)
	w.Uint32Field("MessageID", m.MessageID())
	w.Uint32Field("CurrentInboundSequenceNumber", m.CurrentInboundSequenceNumber())
	w.Uint32Field("CurrentOutboundSequenceNumber", m.CurrentOutboundSequenceNumber())
	w.Uint64Field("CurrentInternalOrderID", m.CurrentInternalOrderID())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.StringField("LastReceivedExecutionIdentifier", m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

func (m SCConfigurationSynchronizationBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10109)
	w.Uint32Field("MessageID", m.MessageID())
	w.Uint32Field("CurrentInboundSequenceNumber", m.CurrentInboundSequenceNumber())
	w.Uint32Field("CurrentOutboundSequenceNumber", m.CurrentOutboundSequenceNumber())
	w.Uint64Field("CurrentInternalOrderID", m.CurrentInternalOrderID())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.StringField("LastReceivedExecutionIdentifier", m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SCConfigurationSynchronizationPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10109)
	w.Uint32Field("MessageID", m.MessageID())
	w.Uint32Field("CurrentInboundSequenceNumber", m.CurrentInboundSequenceNumber())
	w.Uint32Field("CurrentOutboundSequenceNumber", m.CurrentOutboundSequenceNumber())
	w.Uint64Field("CurrentInternalOrderID", m.CurrentInternalOrderID())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.StringField("LastReceivedExecutionIdentifier", m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

func (m SCConfigurationSynchronizationPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10109)
	w.Uint32Field("MessageID", m.MessageID())
	w.Uint32Field("CurrentInboundSequenceNumber", m.CurrentInboundSequenceNumber())
	w.Uint32Field("CurrentOutboundSequenceNumber", m.CurrentOutboundSequenceNumber())
	w.Uint64Field("CurrentInternalOrderID", m.CurrentInternalOrderID())
	w.BoolField("IsSnapshot", m.IsSnapshot())
	w.StringField("LastReceivedExecutionIdentifier", m.LastReceivedExecutionIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SCConfigurationSynchronizationBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10109 {
		return message.ErrWrongType
	}
	m.SetMessageID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentInboundSequenceNumber(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentOutboundSequenceNumber(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastReceivedExecutionIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SCConfigurationSynchronizationPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10109 {
		return message.ErrWrongType
	}
	m.SetMessageID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentInboundSequenceNumber(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentOutboundSequenceNumber(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentInternalOrderID(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetLastReceivedExecutionIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SCConfigurationSynchronizationBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10109 {
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
		case "MessageID":
			m.SetMessageID(r.Uint32())
		case "CurrentInboundSequenceNumber":
			m.SetCurrentInboundSequenceNumber(r.Uint32())
		case "CurrentOutboundSequenceNumber":
			m.SetCurrentOutboundSequenceNumber(r.Uint32())
		case "CurrentInternalOrderID":
			m.SetCurrentInternalOrderID(r.Uint64())
		case "IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "LastReceivedExecutionIdentifier":
			m.SetLastReceivedExecutionIdentifier(r.String())
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

func (m *SCConfigurationSynchronizationPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10109 {
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
		case "MessageID":
			m.SetMessageID(r.Uint32())
		case "CurrentInboundSequenceNumber":
			m.SetCurrentInboundSequenceNumber(r.Uint32())
		case "CurrentOutboundSequenceNumber":
			m.SetCurrentOutboundSequenceNumber(r.Uint32())
		case "CurrentInternalOrderID":
			m.SetCurrentInternalOrderID(r.Uint64())
		case "IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "LastReceivedExecutionIdentifier":
			m.SetLastReceivedExecutionIdentifier(r.String())
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
