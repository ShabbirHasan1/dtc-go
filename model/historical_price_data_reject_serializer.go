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

func (m HistoricalPriceDataReject) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRejectPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRejectPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 802 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRetryTimeInSeconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 802 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRetryTimeInSeconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 802 {
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
			m.SetRequestID(r.Int32())
		case "RejectText":
			m.SetRejectText(r.String())
		case "RejectReasonCode":
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
		case "RetryTimeInSeconds":
			m.SetRetryTimeInSeconds(r.Uint16())
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

func (m *HistoricalPriceDataRejectPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 802 {
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
			m.SetRequestID(r.Int32())
		case "RejectText":
			m.SetRejectText(r.String())
		case "RejectReasonCode":
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
		case "RetryTimeInSeconds":
			m.SetRetryTimeInSeconds(r.Uint16())
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

func (m HistoricalPriceDataRejectFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRejectFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":802,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.String(m.RejectText())
	w.Data = append(w.Data, ',')
	w.Int16(int16(m.RejectReasonCode()))
	w.Data = append(w.Data, ',')
	w.Uint16(m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRejectFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

func (m HistoricalPriceDataRejectFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 802)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	w.Int16Field("RejectReasonCode", int16(m.RejectReasonCode()))
	w.Uint16Field("RetryTimeInSeconds", m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 802 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRetryTimeInSeconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 802 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
	if r.IsError() {
		return r.Error()
	}
	m.SetRetryTimeInSeconds(r.Uint16())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 802 {
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
			m.SetRequestID(r.Int32())
		case "RejectText":
			m.SetRejectText(r.String())
		case "RejectReasonCode":
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
		case "RetryTimeInSeconds":
			m.SetRetryTimeInSeconds(r.Uint16())
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

func (m *HistoricalPriceDataRejectFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 802 {
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
			m.SetRequestID(r.Int32())
		case "RejectText":
			m.SetRejectText(r.String())
		case "RejectReasonCode":
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.Int16()))
		case "RetryTimeInSeconds":
			m.SetRetryTimeInSeconds(r.Uint16())
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

func (m *HistoricalPriceDataRejectBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetRejectText(r.ReadString())
		case 3:
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.ReadInt16()))
		case 4:
			m.SetRetryTimeInSeconds(r.ReadUint16())
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

func (m *HistoricalPriceDataRejectPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetRejectText(r.ReadString())
		case 3:
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.ReadInt16()))
		case 4:
			m.SetRetryTimeInSeconds(r.ReadUint16())
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

func (m HistoricalPriceDataRejectBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 802)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.RejectText())
	w.WriteVarint16(3, int16(m.RejectReasonCode()))
	w.WriteUvarint16(4, m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRejectPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 802)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.RejectText())
	w.WriteVarint16(3, int16(m.RejectReasonCode()))
	w.WriteUvarint16(4, m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalPriceDataRejectFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetRejectText(r.ReadString())
		case 3:
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.ReadInt16()))
		case 4:
			m.SetRetryTimeInSeconds(r.ReadUint16())
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

func (m *HistoricalPriceDataRejectFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetRequestID(r.ReadInt32())
		case 2:
			m.SetRejectText(r.ReadString())
		case 3:
			m.SetRejectReasonCode(HistoricalPriceDataRejectReasonCodeEnum(r.ReadInt16()))
		case 4:
			m.SetRetryTimeInSeconds(r.ReadUint16())
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

func (m HistoricalPriceDataRejectFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 802)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.RejectText())
	w.WriteVarint16(3, int16(m.RejectReasonCode()))
	w.WriteUvarint16(4, m.RetryTimeInSeconds())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalPriceDataRejectFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 802)
	w.WriteVarint32(1, m.RequestID())
	w.WriteString(2, m.RejectText())
	w.WriteVarint16(3, int16(m.RejectReasonCode()))
	w.WriteUvarint16(4, m.RetryTimeInSeconds())
	return w.Finish(), nil
}
