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

func (m CancelReplaceOrderInt) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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

func (m CancelReplaceOrderIntBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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

func (m CancelReplaceOrderIntPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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

func (m CancelReplaceOrderInt) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

func (m CancelReplaceOrderIntBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

func (m CancelReplaceOrderIntPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetPrice1IsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2IsSet(r.Bool())
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetPrice1IsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2IsSet(r.Bool())
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

func (m *CancelReplaceOrderIntBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetPrice1IsSet(r.Bool())
		case "Price2IsSet":
			m.SetPrice2IsSet(r.Bool())
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetPrice1IsSet(r.Bool())
		case "Price2IsSet":
			m.SetPrice2IsSet(r.Bool())
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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

func (m CancelReplaceOrderIntFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
	w.Bool(m.Price1IsSet())
	w.Data = append(w.Data, ',')
	w.Bool(m.Price2IsSet())
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
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
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
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

func (m CancelReplaceOrderIntFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 205)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int64Field("Price1", m.Price1())
	w.Int64Field("Price2", m.Price2())
	w.Float32Field("Divisor", m.Divisor())
	w.Int64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
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
	m.SetPrice1IsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2IsSet(r.Bool())
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
	m.SetPrice1IsSet(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2IsSet(r.Bool())
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
			m.SetPrice1IsSet(r.Bool())
		case "Price2IsSet":
			m.SetPrice2IsSet(r.Bool())
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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
			m.SetPrice1IsSet(r.Bool())
		case "Price2IsSet":
			m.SetPrice2IsSet(r.Bool())
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetPrice1(r.ReadInt64())
		case 4:
			m.SetPrice2(r.ReadInt64())
		case 5:
			m.SetDivisor(r.ReadFloat32())
		case 6:
			m.SetQuantity(r.ReadInt64())
		case 7:
			m.SetPrice1IsSet(r.ReadBool())
		case 8:
			m.SetPrice2IsSet(r.ReadBool())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
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

func (m *CancelReplaceOrderIntPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetPrice1(r.ReadInt64())
		case 4:
			m.SetPrice2(r.ReadInt64())
		case 5:
			m.SetDivisor(r.ReadFloat32())
		case 6:
			m.SetQuantity(r.ReadInt64())
		case 7:
			m.SetPrice1IsSet(r.ReadBool())
		case 8:
			m.SetPrice2IsSet(r.ReadBool())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
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

func (m CancelReplaceOrderIntBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 205)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteVarint64(3, m.Price1())
	w.WriteVarint64(4, m.Price2())
	w.WriteFixed32Float32(5, m.Divisor())
	w.WriteVarint64(6, m.Quantity())
	w.WriteBool(7, m.Price1IsSet())
	w.WriteBool(8, m.Price2IsSet())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteVarint64(11, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(12, m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 205)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteVarint64(3, m.Price1())
	w.WriteVarint64(4, m.Price2())
	w.WriteFixed32Float32(5, m.Divisor())
	w.WriteVarint64(6, m.Quantity())
	w.WriteBool(7, m.Price1IsSet())
	w.WriteBool(8, m.Price2IsSet())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteVarint64(11, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(12, m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderIntFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetPrice1(r.ReadInt64())
		case 4:
			m.SetPrice2(r.ReadInt64())
		case 5:
			m.SetDivisor(r.ReadFloat32())
		case 6:
			m.SetQuantity(r.ReadInt64())
		case 7:
			m.SetPrice1IsSet(r.ReadBool())
		case 8:
			m.SetPrice2IsSet(r.ReadBool())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
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

func (m *CancelReplaceOrderIntFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetServerOrderID(r.ReadString())
		case 2:
			m.SetClientOrderID(r.ReadString())
		case 3:
			m.SetPrice1(r.ReadInt64())
		case 4:
			m.SetPrice2(r.ReadInt64())
		case 5:
			m.SetDivisor(r.ReadFloat32())
		case 6:
			m.SetQuantity(r.ReadInt64())
		case 7:
			m.SetPrice1IsSet(r.ReadBool())
		case 8:
			m.SetPrice2IsSet(r.ReadBool())
		case 10:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 11:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 12:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
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

func (m CancelReplaceOrderIntFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 205)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteVarint64(3, m.Price1())
	w.WriteVarint64(4, m.Price2())
	w.WriteFixed32Float32(5, m.Divisor())
	w.WriteVarint64(6, m.Quantity())
	w.WriteBool(7, m.Price1IsSet())
	w.WriteBool(8, m.Price2IsSet())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteVarint64(11, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(12, m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderIntFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 205)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteVarint64(3, m.Price1())
	w.WriteVarint64(4, m.Price2())
	w.WriteFixed32Float32(5, m.Divisor())
	w.WriteVarint64(6, m.Quantity())
	w.WriteBool(7, m.Price1IsSet())
	w.WriteBool(8, m.Price2IsSet())
	w.WriteVarint32(10, int32(m.TimeInForce()))
	w.WriteVarint64(11, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(12, m.UpdatePrice1OffsetToParent())
	return w.Finish(), nil
}
