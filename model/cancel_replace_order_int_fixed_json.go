//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":205,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float32(m.Divisor())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Price2IsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.Unused())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

func (m CancelReplaceOrderIntFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":205,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Int64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float32(m.Divisor())
	w.Data = append(w.Data, ',')
	w.Int64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Price2IsSet())
	w.Data = append(w.Data, ',')
	w.Int32(m.Unused())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TimeInForce()))
	w.Data = append(w.Data, ',')
	w.Int64(int64(m.GoodTillDateTime()))
	w.Data = append(w.Data, ',')
	w.Uint8(m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.Uint8Field("Price1IsSet", m.Price1IsSet())
	w.Uint8Field("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

func (m CancelReplaceOrderIntFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.Uint8Field("Price1IsSet", m.Price1IsSet())
	w.Uint8Field("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 205 {
		return message.ErrWrongType
	}
	m.SetServerOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetDivisor(r.Float32())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1IsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2IsSet(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnused(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTimeInForce(TimeInForceEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetGoodTillDateTime(DateTime(r.Int64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetUpdatePrice1OffsetToParent(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 205 {
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
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "Price1":
			m.SetPrice1(r.Int64())
		case "Price2":
			m.SetPrice2(r.Int64())
		case "Divisor":
			m.SetDivisor(r.Float32())
		case "Quantity":
			m.SetQuantity(r.Int64())
		case "Price1IsSet":
			m.SetPrice1IsSet(r.Uint8())
		case "Price2IsSet":
			m.SetPrice2IsSet(r.Uint8())
		case "Unused":
			m.SetUnused(r.Int32())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "UpdatePrice1OffsetToParent":
			m.SetUpdatePrice1OffsetToParent(r.Uint8())
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
