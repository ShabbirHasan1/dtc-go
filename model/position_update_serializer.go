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

func (m PositionUpdate) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdateBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdatePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdatePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdate) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdateBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdatePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdatePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdateBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 306 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAveragePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaxPotentialPostionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdatePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 306 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAveragePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaxPotentialPostionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdateBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 306 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "AveragePrice":
			m.SetAveragePrice(r.Float64())
		case "PositionIdentifier":
			m.SetPositionIdentifier(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "NoPositions":
			m.SetNoPositions(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "EntryDateTime":
			m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
		case "OpenProfitLoss":
			m.SetOpenProfitLoss(r.Float64())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "QuantityLimit":
			m.SetQuantityLimit(r.Float64())
		case "MaxPotentialPostionQuantity":
			m.SetMaxPotentialPostionQuantity(r.Float64())
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

func (m *PositionUpdatePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 306 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "AveragePrice":
			m.SetAveragePrice(r.Float64())
		case "PositionIdentifier":
			m.SetPositionIdentifier(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "NoPositions":
			m.SetNoPositions(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "EntryDateTime":
			m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
		case "OpenProfitLoss":
			m.SetOpenProfitLoss(r.Float64())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "QuantityLimit":
			m.SetQuantityLimit(r.Float64())
		case "MaxPotentialPostionQuantity":
			m.SetMaxPotentialPostionQuantity(r.Float64())
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

func (m PositionUpdateFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdateFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdateFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdateFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":306,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Int32(m.TotalNumberMessages())
	w.Data = append(w.Data, ',')
	w.Int32(m.MessageNumber())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.AveragePrice())
	w.Data = append(w.Data, ',')
	w.String(m.PositionIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoPositions())
	w.Data = append(w.Data, ',')
	w.Uint8(m.Unsolicited())
	w.Data = append(w.Data, ',')
	w.Float64(m.MarginRequirement())
	w.Data = append(w.Data, ',')
	w.Uint32(uint32(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.HighPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.LowPriceDuringPosition())
	w.Data = append(w.Data, ',')
	w.Float64(m.QuantityLimit())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdateFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdateFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdateFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

func (m PositionUpdateFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 306)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Float64Field("Quantity", m.Quantity())
	w.Float64Field("AveragePrice", m.AveragePrice())
	w.StringField("PositionIdentifier", m.PositionIdentifier())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Uint8Field("NoPositions", m.NoPositions())
	w.Uint8Field("Unsolicited", m.Unsolicited())
	w.Float64Field("MarginRequirement", m.MarginRequirement())
	w.Uint32Field("EntryDateTime", uint32(m.EntryDateTime()))
	w.Float64Field("OpenProfitLoss", m.OpenProfitLoss())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("QuantityLimit", m.QuantityLimit())
	w.Float64Field("MaxPotentialPostionQuantity", m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdateFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 306 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAveragePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaxPotentialPostionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdateFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 306 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTotalNumberMessages(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMessageNumber(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAveragePrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoPositions(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetUnsolicited(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarginRequirement(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetHighPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetLowPriceDuringPosition(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantityLimit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaxPotentialPostionQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdateFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 306 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "AveragePrice":
			m.SetAveragePrice(r.Float64())
		case "PositionIdentifier":
			m.SetPositionIdentifier(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "NoPositions":
			m.SetNoPositions(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "EntryDateTime":
			m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
		case "OpenProfitLoss":
			m.SetOpenProfitLoss(r.Float64())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "QuantityLimit":
			m.SetQuantityLimit(r.Float64())
		case "MaxPotentialPostionQuantity":
			m.SetMaxPotentialPostionQuantity(r.Float64())
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

func (m *PositionUpdateFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 306 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "AveragePrice":
			m.SetAveragePrice(r.Float64())
		case "PositionIdentifier":
			m.SetPositionIdentifier(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "NoPositions":
			m.SetNoPositions(r.Uint8())
		case "Unsolicited":
			m.SetUnsolicited(r.Uint8())
		case "MarginRequirement":
			m.SetMarginRequirement(r.Float64())
		case "EntryDateTime":
			m.SetEntryDateTime(DateTime4Byte(r.Uint32()))
		case "OpenProfitLoss":
			m.SetOpenProfitLoss(r.Float64())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "QuantityLimit":
			m.SetQuantityLimit(r.Float64())
		case "MaxPotentialPostionQuantity":
			m.SetMaxPotentialPostionQuantity(r.Float64())
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

func (m *PositionUpdateBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetQuantity(r.ReadFloat64())
		case 7:
			m.SetAveragePrice(r.ReadFloat64())
		case 8:
			m.SetPositionIdentifier(r.ReadString())
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetNoPositions(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetMarginRequirement(r.ReadFloat64())
		case 13:
			m.SetEntryDateTime(DateTime4Byte(r.ReadUint32()))
		case 14:
			m.SetOpenProfitLoss(r.ReadFloat64())
		case 15:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 16:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetQuantityLimit(r.ReadFloat64())
		case 18:
			m.SetMaxPotentialPostionQuantity(r.ReadFloat64())
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

func (m *PositionUpdatePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetQuantity(r.ReadFloat64())
		case 7:
			m.SetAveragePrice(r.ReadFloat64())
		case 8:
			m.SetPositionIdentifier(r.ReadString())
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetNoPositions(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetMarginRequirement(r.ReadFloat64())
		case 13:
			m.SetEntryDateTime(DateTime4Byte(r.ReadUint32()))
		case 14:
			m.SetOpenProfitLoss(r.ReadFloat64())
		case 15:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 16:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetQuantityLimit(r.ReadFloat64())
		case 18:
			m.SetMaxPotentialPostionQuantity(r.ReadFloat64())
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

func (m PositionUpdateBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 306)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteFixed64Float64(6, m.Quantity())
	w.WriteFixed64Float64(7, m.AveragePrice())
	w.WriteString(8, m.PositionIdentifier())
	w.WriteString(9, m.TradeAccount())
	w.WriteUvarint8(10, m.NoPositions())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.MarginRequirement())
	w.WriteUvarint32(13, uint32(m.EntryDateTime()))
	w.WriteFixed64Float64(14, m.OpenProfitLoss())
	w.WriteFixed64Float64(15, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(16, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(17, m.QuantityLimit())
	w.WriteFixed64Float64(18, m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdatePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 306)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteFixed64Float64(6, m.Quantity())
	w.WriteFixed64Float64(7, m.AveragePrice())
	w.WriteString(8, m.PositionIdentifier())
	w.WriteString(9, m.TradeAccount())
	w.WriteUvarint8(10, m.NoPositions())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.MarginRequirement())
	w.WriteUvarint32(13, uint32(m.EntryDateTime()))
	w.WriteFixed64Float64(14, m.OpenProfitLoss())
	w.WriteFixed64Float64(15, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(16, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(17, m.QuantityLimit())
	w.WriteFixed64Float64(18, m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *PositionUpdateFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetQuantity(r.ReadFloat64())
		case 7:
			m.SetAveragePrice(r.ReadFloat64())
		case 8:
			m.SetPositionIdentifier(r.ReadString())
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetNoPositions(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetMarginRequirement(r.ReadFloat64())
		case 13:
			m.SetEntryDateTime(DateTime4Byte(r.ReadUint32()))
		case 14:
			m.SetOpenProfitLoss(r.ReadFloat64())
		case 15:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 16:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetQuantityLimit(r.ReadFloat64())
		case 18:
			m.SetMaxPotentialPostionQuantity(r.ReadFloat64())
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

func (m *PositionUpdateFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetTotalNumberMessages(r.ReadInt32())
		case 3:
			m.SetMessageNumber(r.ReadInt32())
		case 4:
			m.SetSymbol(r.ReadString())
		case 5:
			m.SetExchange(r.ReadString())
		case 6:
			m.SetQuantity(r.ReadFloat64())
		case 7:
			m.SetAveragePrice(r.ReadFloat64())
		case 8:
			m.SetPositionIdentifier(r.ReadString())
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetNoPositions(r.ReadUint8())
		case 11:
			m.SetUnsolicited(r.ReadUint8())
		case 12:
			m.SetMarginRequirement(r.ReadFloat64())
		case 13:
			m.SetEntryDateTime(DateTime4Byte(r.ReadUint32()))
		case 14:
			m.SetOpenProfitLoss(r.ReadFloat64())
		case 15:
			m.SetHighPriceDuringPosition(r.ReadFloat64())
		case 16:
			m.SetLowPriceDuringPosition(r.ReadFloat64())
		case 17:
			m.SetQuantityLimit(r.ReadFloat64())
		case 18:
			m.SetMaxPotentialPostionQuantity(r.ReadFloat64())
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

func (m PositionUpdateFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 306)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteFixed64Float64(6, m.Quantity())
	w.WriteFixed64Float64(7, m.AveragePrice())
	w.WriteString(8, m.PositionIdentifier())
	w.WriteString(9, m.TradeAccount())
	w.WriteUvarint8(10, m.NoPositions())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.MarginRequirement())
	w.WriteUvarint32(13, uint32(m.EntryDateTime()))
	w.WriteFixed64Float64(14, m.OpenProfitLoss())
	w.WriteFixed64Float64(15, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(16, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(17, m.QuantityLimit())
	w.WriteFixed64Float64(18, m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m PositionUpdateFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 306)
	w.WriteVarint32(1, m.RequestID())
	w.WriteVarint32(2, m.TotalNumberMessages())
	w.WriteVarint32(3, m.MessageNumber())
	w.WriteString(4, m.Symbol())
	w.WriteString(5, m.Exchange())
	w.WriteFixed64Float64(6, m.Quantity())
	w.WriteFixed64Float64(7, m.AveragePrice())
	w.WriteString(8, m.PositionIdentifier())
	w.WriteString(9, m.TradeAccount())
	w.WriteUvarint8(10, m.NoPositions())
	w.WriteUvarint8(11, m.Unsolicited())
	w.WriteFixed64Float64(12, m.MarginRequirement())
	w.WriteUvarint32(13, uint32(m.EntryDateTime()))
	w.WriteFixed64Float64(14, m.OpenProfitLoss())
	w.WriteFixed64Float64(15, m.HighPriceDuringPosition())
	w.WriteFixed64Float64(16, m.LowPriceDuringPosition())
	w.WriteFixed64Float64(17, m.QuantityLimit())
	w.WriteFixed64Float64(18, m.MaxPotentialPostionQuantity())
	return w.Finish(), nil
}
