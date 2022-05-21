//go:build !tinygo

// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"github.com/moontrade/dtc-go/message/pb"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 204 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderPointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 204 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 204 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m *CancelReplaceOrderPointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 204 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m *CancelReplaceOrderBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderPointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":204,\"F\":["...)
	w.String(m.ServerOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price1())
	w.Data = append(w.Data, ',')
	w.Float64(m.Price2())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
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
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.Price1AsString())
	w.Data = append(w.Data, ',')
	w.String(m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

func (m CancelReplaceOrderFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 204)
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.BoolField("Price1IsSet", m.Price1IsSet())
	w.BoolField("Price2IsSet", m.Price2IsSet())
	w.Int32Field("Unused", m.Unused())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.Uint8Field("UpdatePrice1OffsetToParent", m.UpdatePrice1OffsetToParent())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixedBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 204 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixedPointerBuilder) UnmarshalJSONCompact(r *json.Reader) error {
	if r.Type != 204 {
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
	m.SetPrice1(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
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
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice1AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPrice2AsString(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixedBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 204 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m *CancelReplaceOrderFixedPointerBuilder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 204 {
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
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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

func (m *CancelReplaceOrderFixedBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixedPointerBuilder) UnmarshalJSONReader(r *json.Reader) error {
	if r.IsCompact {
		return m.UnmarshalJSONCompact(r)
	}
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 4:
			m.SetPrice2(r.ReadFloat64())
		case 5:
			m.SetQuantity(r.ReadFloat64())
		case 6:
			m.SetPrice1IsSet(r.ReadBool())
		case 7:
			m.SetPrice2IsSet(r.ReadBool())
		case 9:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 10:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 11:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetPrice1AsString(r.ReadString())
		case 14:
			m.SetPrice2AsString(r.ReadString())
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

func (m *CancelReplaceOrderPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 4:
			m.SetPrice2(r.ReadFloat64())
		case 5:
			m.SetQuantity(r.ReadFloat64())
		case 6:
			m.SetPrice1IsSet(r.ReadBool())
		case 7:
			m.SetPrice2IsSet(r.ReadBool())
		case 9:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 10:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 11:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetPrice1AsString(r.ReadString())
		case 14:
			m.SetPrice2AsString(r.ReadString())
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

func (m CancelReplaceOrderBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 204)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteFixed64Float64(3, m.Price1())
	w.WriteFixed64Float64(4, m.Price2())
	w.WriteFixed64Float64(5, m.Quantity())
	w.WriteBool(6, m.Price1IsSet())
	w.WriteBool(7, m.Price2IsSet())
	w.WriteVarint32(9, int32(m.TimeInForce()))
	w.WriteVarint64(10, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(11, m.UpdatePrice1OffsetToParent())
	w.WriteString(12, m.TradeAccount())
	w.WriteString(13, m.Price1AsString())
	w.WriteString(14, m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 204)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteFixed64Float64(3, m.Price1())
	w.WriteFixed64Float64(4, m.Price2())
	w.WriteFixed64Float64(5, m.Quantity())
	w.WriteBool(6, m.Price1IsSet())
	w.WriteBool(7, m.Price2IsSet())
	w.WriteVarint32(9, int32(m.TimeInForce()))
	w.WriteVarint64(10, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(11, m.UpdatePrice1OffsetToParent())
	w.WriteString(12, m.TradeAccount())
	w.WriteString(13, m.Price1AsString())
	w.WriteString(14, m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *CancelReplaceOrderFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 4:
			m.SetPrice2(r.ReadFloat64())
		case 5:
			m.SetQuantity(r.ReadFloat64())
		case 6:
			m.SetPrice1IsSet(r.ReadBool())
		case 7:
			m.SetPrice2IsSet(r.ReadBool())
		case 9:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 10:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 11:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetPrice1AsString(r.ReadString())
		case 14:
			m.SetPrice2AsString(r.ReadString())
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

func (m *CancelReplaceOrderFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetPrice1(r.ReadFloat64())
		case 4:
			m.SetPrice2(r.ReadFloat64())
		case 5:
			m.SetQuantity(r.ReadFloat64())
		case 6:
			m.SetPrice1IsSet(r.ReadBool())
		case 7:
			m.SetPrice2IsSet(r.ReadBool())
		case 9:
			m.SetTimeInForce(TimeInForceEnum(r.ReadInt32()))
		case 10:
			m.SetGoodTillDateTime(DateTime(r.ReadInt64()))
		case 11:
			m.SetUpdatePrice1OffsetToParent(r.ReadUint8())
		case 12:
			m.SetTradeAccount(r.ReadString())
		case 13:
			m.SetPrice1AsString(r.ReadString())
		case 14:
			m.SetPrice2AsString(r.ReadString())
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

func (m CancelReplaceOrderFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 204)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteFixed64Float64(3, m.Price1())
	w.WriteFixed64Float64(4, m.Price2())
	w.WriteFixed64Float64(5, m.Quantity())
	w.WriteBool(6, m.Price1IsSet())
	w.WriteBool(7, m.Price2IsSet())
	w.WriteVarint32(9, int32(m.TimeInForce()))
	w.WriteVarint64(10, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(11, m.UpdatePrice1OffsetToParent())
	w.WriteString(12, m.TradeAccount())
	w.WriteString(13, m.Price1AsString())
	w.WriteString(14, m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m CancelReplaceOrderFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 204)
	w.WriteString(1, m.ServerOrderID())
	w.WriteString(2, m.ClientOrderID())
	w.WriteFixed64Float64(3, m.Price1())
	w.WriteFixed64Float64(4, m.Price2())
	w.WriteFixed64Float64(5, m.Quantity())
	w.WriteBool(6, m.Price1IsSet())
	w.WriteBool(7, m.Price2IsSet())
	w.WriteVarint32(9, int32(m.TimeInForce()))
	w.WriteVarint64(10, int64(m.GoodTillDateTime()))
	w.WriteUvarint8(11, m.UpdatePrice1OffsetToParent())
	w.WriteString(12, m.TradeAccount())
	w.WriteString(13, m.Price1AsString())
	w.WriteString(14, m.Price2AsString())
	return w.Finish(), nil
}
