//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10102,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.ExitDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.EntryPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.ExitPrice())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TradeType()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.EntryQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.ExitQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MaxOpenQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.ClosedProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumOpenPositionLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumOpenPositionProfit())
	w.Data = append(w.Data, ',')
	w.Float64(m.Commission())
	w.Data = append(w.Data, ',')
	w.String(m.OpenFillExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.CloseFillExecutionID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m HistoricalTradesResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10102,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFinalMessage())
	w.Data = append(w.Data, ',')
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.EntryDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.ExitDateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.EntryPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.ExitPrice())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TradeType()))
	w.Data = append(w.Data, ',')
	w.Uint32(m.EntryQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.ExitQuantity())
	w.Data = append(w.Data, ',')
	w.Uint32(m.MaxOpenQuantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.ClosedProfitLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumOpenPositionLoss())
	w.Data = append(w.Data, ',')
	w.Float64(m.MaximumOpenPositionProfit())
	w.Data = append(w.Data, ',')
	w.Float64(m.Commission())
	w.Data = append(w.Data, ',')
	w.String(m.OpenFillExecutionID())
	w.Data = append(w.Data, ',')
	w.String(m.CloseFillExecutionID())
	w.Data = append(w.Data, ',')
	w.Uint32(m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalTradesResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10102)
	w.Int32Field("RequestID", m.RequestID())
	w.Uint8Field("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("EntryDateTime", float64(m.EntryDateTime()))
	w.Float64Field("ExitDateTime", float64(m.ExitDateTime()))
	w.Float64Field("EntryPrice", m.EntryPrice())
	w.Float64Field("ExitPrice", m.ExitPrice())
	w.Int32Field("TradeType", int32(m.TradeType()))
	w.Uint32Field("EntryQuantity", m.EntryQuantity())
	w.Uint32Field("ExitQuantity", m.ExitQuantity())
	w.Uint32Field("MaxOpenQuantity", m.MaxOpenQuantity())
	w.Float64Field("ClosedProfitLoss", m.ClosedProfitLoss())
	w.Float64Field("MaximumOpenPositionLoss", m.MaximumOpenPositionLoss())
	w.Float64Field("MaximumOpenPositionProfit", m.MaximumOpenPositionProfit())
	w.Float64Field("Commission", m.Commission())
	w.StringField("OpenFillExecutionID", m.OpenFillExecutionID())
	w.StringField("CloseFillExecutionID", m.CloseFillExecutionID())
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

func (m HistoricalTradesResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10102)
	w.Int32Field("RequestID", m.RequestID())
	w.Uint8Field("IsFinalMessage", m.IsFinalMessage())
	w.StringField("Symbol", m.Symbol())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("EntryDateTime", float64(m.EntryDateTime()))
	w.Float64Field("ExitDateTime", float64(m.ExitDateTime()))
	w.Float64Field("EntryPrice", m.EntryPrice())
	w.Float64Field("ExitPrice", m.ExitPrice())
	w.Int32Field("TradeType", int32(m.TradeType()))
	w.Uint32Field("EntryQuantity", m.EntryQuantity())
	w.Uint32Field("ExitQuantity", m.ExitQuantity())
	w.Uint32Field("MaxOpenQuantity", m.MaxOpenQuantity())
	w.Float64Field("ClosedProfitLoss", m.ClosedProfitLoss())
	w.Float64Field("MaximumOpenPositionLoss", m.MaximumOpenPositionLoss())
	w.Float64Field("MaximumOpenPositionProfit", m.MaximumOpenPositionProfit())
	w.Float64Field("Commission", m.Commission())
	w.StringField("OpenFillExecutionID", m.OpenFillExecutionID())
	w.StringField("CloseFillExecutionID", m.CloseFillExecutionID())
	w.Uint32Field("SubAccountIdentifier", m.SubAccountIdentifier())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10102 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalMessage(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetExitDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetExitPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeType(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetEntryQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetExitQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaxOpenQuantity(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetClosedProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumOpenPositionLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetMaximumOpenPositionProfit(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetCommission(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenFillExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetCloseFillExecutionID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetSubAccountIdentifier(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalTradesResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10102 {
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
		case "IsFinalMessage":
			m.SetIsFinalMessage(r.Uint8())
		case "Symbol":
			m.SetSymbol(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "EntryDateTime":
			m.SetEntryDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "ExitDateTime":
			m.SetExitDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "EntryPrice":
			m.SetEntryPrice(r.Float64())
		case "ExitPrice":
			m.SetExitPrice(r.Float64())
		case "TradeType":
			m.SetTradeType(BuySellEnum(r.Int32()))
		case "EntryQuantity":
			m.SetEntryQuantity(r.Uint32())
		case "ExitQuantity":
			m.SetExitQuantity(r.Uint32())
		case "MaxOpenQuantity":
			m.SetMaxOpenQuantity(r.Uint32())
		case "ClosedProfitLoss":
			m.SetClosedProfitLoss(r.Float64())
		case "MaximumOpenPositionLoss":
			m.SetMaximumOpenPositionLoss(r.Float64())
		case "MaximumOpenPositionProfit":
			m.SetMaximumOpenPositionProfit(r.Float64())
		case "Commission":
			m.SetCommission(r.Float64())
		case "OpenFillExecutionID":
			m.SetOpenFillExecutionID(r.String())
		case "CloseFillExecutionID":
			m.SetCloseFillExecutionID(r.String())
		case "SubAccountIdentifier":
			m.SetSubAccountIdentifier(r.Uint32())
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
