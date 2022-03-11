//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HeartbeatExtended) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":3,\"F\":["...)
	w.Uint32(m.NumDroppedMessages())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.CurrentDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.SecondsSinceLastReceivedHeartbeat())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumberOfOutstandingSentBuffers())
	w.Data = append(w.Data, ',')
	w.Uint16(m.PendingTransmissionDelayInMilliseconds())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentSendBufferSizeInBytes())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.SendingDateTimeMicroseconds()))
	w.Data = append(w.Data, ',')
	w.Float32(m.DataCompressionRatio())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TotalUncompressedBytes())
	w.Data = append(w.Data, ',')
	w.Float64(m.TotalCompressionTime())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfCompressions())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SourceIPAddress())
	return w.Finish(), nil
}

func (m HeartbeatExtendedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":3,\"F\":["...)
	w.Uint32(m.NumDroppedMessages())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.CurrentDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.SecondsSinceLastReceivedHeartbeat())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumberOfOutstandingSentBuffers())
	w.Data = append(w.Data, ',')
	w.Uint16(m.PendingTransmissionDelayInMilliseconds())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentSendBufferSizeInBytes())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.SendingDateTimeMicroseconds()))
	w.Data = append(w.Data, ',')
	w.Float32(m.DataCompressionRatio())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TotalUncompressedBytes())
	w.Data = append(w.Data, ',')
	w.Float64(m.TotalCompressionTime())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfCompressions())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SourceIPAddress())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HeartbeatExtendedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":3,\"F\":["...)
	w.Uint32(m.NumDroppedMessages())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.CurrentDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.SecondsSinceLastReceivedHeartbeat())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumberOfOutstandingSentBuffers())
	w.Data = append(w.Data, ',')
	w.Uint16(m.PendingTransmissionDelayInMilliseconds())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentSendBufferSizeInBytes())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.SendingDateTimeMicroseconds()))
	w.Data = append(w.Data, ',')
	w.Float32(m.DataCompressionRatio())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TotalUncompressedBytes())
	w.Data = append(w.Data, ',')
	w.Float64(m.TotalCompressionTime())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfCompressions())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SourceIPAddress())
	return w.Finish(), nil
}

func (m HeartbeatExtendedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":3,\"F\":["...)
	w.Uint32(m.NumDroppedMessages())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.CurrentDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.SecondsSinceLastReceivedHeartbeat())
	w.Data = append(w.Data, ',')
	w.Uint16(m.NumberOfOutstandingSentBuffers())
	w.Data = append(w.Data, ',')
	w.Uint16(m.PendingTransmissionDelayInMilliseconds())
	w.Data = append(w.Data, ',')
	w.Uint32(m.CurrentSendBufferSizeInBytes())
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.SendingDateTimeMicroseconds()))
	w.Data = append(w.Data, ',')
	w.Float32(m.DataCompressionRatio())
	w.Data = append(w.Data, ',')
	w.Uint64(m.TotalUncompressedBytes())
	w.Data = append(w.Data, ',')
	w.Float64(m.TotalCompressionTime())
	w.Data = append(w.Data, ',')
	w.Uint32(m.NumberOfCompressions())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SourceIPAddress())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HeartbeatExtended) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 3)
	w.Uint32Field("NumDroppedMessages", m.NumDroppedMessages())
	w.Int64Field("CurrentDateTime", int64(m.CurrentDateTime()))
	w.Uint16Field("SecondsSinceLastReceivedHeartbeat", m.SecondsSinceLastReceivedHeartbeat())
	w.Uint16Field("NumberOfOutstandingSentBuffers", m.NumberOfOutstandingSentBuffers())
	w.Uint16Field("PendingTransmissionDelayInMilliseconds", m.PendingTransmissionDelayInMilliseconds())
	w.Uint32Field("CurrentSendBufferSizeInBytes", m.CurrentSendBufferSizeInBytes())
	w.Int64Field("SendingDateTimeMicroseconds", int64(m.SendingDateTimeMicroseconds()))
	w.Float32Field("DataCompressionRatio", m.DataCompressionRatio())
	w.Uint64Field("TotalUncompressedBytes", m.TotalUncompressedBytes())
	w.Float64Field("TotalCompressionTime", m.TotalCompressionTime())
	w.Uint32Field("NumberOfCompressions", m.NumberOfCompressions())
	w.Uint32Field("SourceIPAddress", m.SourceIPAddress())
	return w.Finish(), nil
}

func (m HeartbeatExtendedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 3)
	w.Uint32Field("NumDroppedMessages", m.NumDroppedMessages())
	w.Int64Field("CurrentDateTime", int64(m.CurrentDateTime()))
	w.Uint16Field("SecondsSinceLastReceivedHeartbeat", m.SecondsSinceLastReceivedHeartbeat())
	w.Uint16Field("NumberOfOutstandingSentBuffers", m.NumberOfOutstandingSentBuffers())
	w.Uint16Field("PendingTransmissionDelayInMilliseconds", m.PendingTransmissionDelayInMilliseconds())
	w.Uint32Field("CurrentSendBufferSizeInBytes", m.CurrentSendBufferSizeInBytes())
	w.Int64Field("SendingDateTimeMicroseconds", int64(m.SendingDateTimeMicroseconds()))
	w.Float32Field("DataCompressionRatio", m.DataCompressionRatio())
	w.Uint64Field("TotalUncompressedBytes", m.TotalUncompressedBytes())
	w.Float64Field("TotalCompressionTime", m.TotalCompressionTime())
	w.Uint32Field("NumberOfCompressions", m.NumberOfCompressions())
	w.Uint32Field("SourceIPAddress", m.SourceIPAddress())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HeartbeatExtendedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 3)
	w.Uint32Field("NumDroppedMessages", m.NumDroppedMessages())
	w.Int64Field("CurrentDateTime", int64(m.CurrentDateTime()))
	w.Uint16Field("SecondsSinceLastReceivedHeartbeat", m.SecondsSinceLastReceivedHeartbeat())
	w.Uint16Field("NumberOfOutstandingSentBuffers", m.NumberOfOutstandingSentBuffers())
	w.Uint16Field("PendingTransmissionDelayInMilliseconds", m.PendingTransmissionDelayInMilliseconds())
	w.Uint32Field("CurrentSendBufferSizeInBytes", m.CurrentSendBufferSizeInBytes())
	w.Int64Field("SendingDateTimeMicroseconds", int64(m.SendingDateTimeMicroseconds()))
	w.Float32Field("DataCompressionRatio", m.DataCompressionRatio())
	w.Uint64Field("TotalUncompressedBytes", m.TotalUncompressedBytes())
	w.Float64Field("TotalCompressionTime", m.TotalCompressionTime())
	w.Uint32Field("NumberOfCompressions", m.NumberOfCompressions())
	w.Uint32Field("SourceIPAddress", m.SourceIPAddress())
	return w.Finish(), nil
}

func (m HeartbeatExtendedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 3)
	w.Uint32Field("NumDroppedMessages", m.NumDroppedMessages())
	w.Int64Field("CurrentDateTime", int64(m.CurrentDateTime()))
	w.Uint16Field("SecondsSinceLastReceivedHeartbeat", m.SecondsSinceLastReceivedHeartbeat())
	w.Uint16Field("NumberOfOutstandingSentBuffers", m.NumberOfOutstandingSentBuffers())
	w.Uint16Field("PendingTransmissionDelayInMilliseconds", m.PendingTransmissionDelayInMilliseconds())
	w.Uint32Field("CurrentSendBufferSizeInBytes", m.CurrentSendBufferSizeInBytes())
	w.Int64Field("SendingDateTimeMicroseconds", int64(m.SendingDateTimeMicroseconds()))
	w.Float32Field("DataCompressionRatio", m.DataCompressionRatio())
	w.Uint64Field("TotalUncompressedBytes", m.TotalUncompressedBytes())
	w.Float64Field("TotalCompressionTime", m.TotalCompressionTime())
	w.Uint32Field("NumberOfCompressions", m.NumberOfCompressions())
	w.Uint32Field("SourceIPAddress", m.SourceIPAddress())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HeartbeatExtendedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 3 {
		return message.ErrWrongType
	}
	m.SetNumDroppedMessages(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSecondsSinceLastReceivedHeartbeat(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfOutstandingSentBuffers(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetPendingTransmissionDelayInMilliseconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentSendBufferSizeInBytes(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSendingDateTimeMicroseconds(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDataCompressionRatio(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalUncompressedBytes(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalCompressionTime(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfCompressions(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSourceIPAddress(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HeartbeatExtendedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 3 {
		return message.ErrWrongType
	}
	m.SetNumDroppedMessages(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetSecondsSinceLastReceivedHeartbeat(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfOutstandingSentBuffers(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetPendingTransmissionDelayInMilliseconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	m.SetCurrentSendBufferSizeInBytes(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSendingDateTimeMicroseconds(DateTimeWithMicrosecondsInt(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetDataCompressionRatio(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalUncompressedBytes(r.Uint64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalCompressionTime(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetNumberOfCompressions(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSourceIPAddress(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HeartbeatExtendedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 3 {
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
		case "NumDroppedMessages":
			m.SetNumDroppedMessages(r.Uint32())
		case "CurrentDateTime":
			m.SetCurrentDateTime(DateTime(r.Int64()))
		case "SecondsSinceLastReceivedHeartbeat":
			m.SetSecondsSinceLastReceivedHeartbeat(r.Uint16())
		case "NumberOfOutstandingSentBuffers":
			m.SetNumberOfOutstandingSentBuffers(r.Uint16())
		case "PendingTransmissionDelayInMilliseconds":
			m.SetPendingTransmissionDelayInMilliseconds(r.Uint16())
		case "CurrentSendBufferSizeInBytes":
			m.SetCurrentSendBufferSizeInBytes(r.Uint32())
		case "SendingDateTimeMicroseconds":
			m.SetSendingDateTimeMicroseconds(DateTimeWithMicrosecondsInt(r.Int64()))
		case "DataCompressionRatio":
			m.SetDataCompressionRatio(r.Float32())
		case "TotalUncompressedBytes":
			m.SetTotalUncompressedBytes(r.Uint64())
		case "TotalCompressionTime":
			m.SetTotalCompressionTime(r.Float64())
		case "NumberOfCompressions":
			m.SetNumberOfCompressions(r.Uint32())
		case "SourceIPAddress":
			m.SetSourceIPAddress(r.Uint32())
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

func (m *HeartbeatExtendedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 3 {
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
		case "NumDroppedMessages":
			m.SetNumDroppedMessages(r.Uint32())
		case "CurrentDateTime":
			m.SetCurrentDateTime(DateTime(r.Int64()))
		case "SecondsSinceLastReceivedHeartbeat":
			m.SetSecondsSinceLastReceivedHeartbeat(r.Uint16())
		case "NumberOfOutstandingSentBuffers":
			m.SetNumberOfOutstandingSentBuffers(r.Uint16())
		case "PendingTransmissionDelayInMilliseconds":
			m.SetPendingTransmissionDelayInMilliseconds(r.Uint16())
		case "CurrentSendBufferSizeInBytes":
			m.SetCurrentSendBufferSizeInBytes(r.Uint32())
		case "SendingDateTimeMicroseconds":
			m.SetSendingDateTimeMicroseconds(DateTimeWithMicrosecondsInt(r.Int64()))
		case "DataCompressionRatio":
			m.SetDataCompressionRatio(r.Float32())
		case "TotalUncompressedBytes":
			m.SetTotalUncompressedBytes(r.Uint64())
		case "TotalCompressionTime":
			m.SetTotalCompressionTime(r.Float64())
		case "NumberOfCompressions":
			m.SetNumberOfCompressions(r.Uint32())
		case "SourceIPAddress":
			m.SetSourceIPAddress(r.Uint32())
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
