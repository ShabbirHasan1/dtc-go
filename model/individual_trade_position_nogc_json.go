//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m IndividualTradePositionPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10112,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Int64(m.TradeDateTime())
	w.Data = append(w.Data, ',')
	w.String(m.FillExecutionIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m IndividualTradePositionPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10112,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSimulated())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Float64(m.Quantity())
	w.Data = append(w.Data, ',')
	w.Float64(m.PositionPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.OpenProfitLoss())
	w.Data = append(w.Data, ',')
	w.Int64(m.TradeDateTime())
	w.Data = append(w.Data, ',')
	w.String(m.FillExecutionIdentifier())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsSnapshot())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsFirstMessageInBatch())
	w.Data = append(w.Data, ',')
	w.Uint8(m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m IndividualTradePositionPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10112)
	w.StringField("m_Symbol", m.Symbol())
	w.Uint8Field("m_IsSimulated", m.IsSimulated())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Float64Field("m_Quantity", m.Quantity())
	w.Float64Field("m_PositionPrice", m.PositionPrice())
	w.Float64Field("m_OpenProfitLoss", m.OpenProfitLoss())
	w.Int64Field("m_TradeDateTime", m.TradeDateTime())
	w.StringField("m_FillExecutionIdentifier", m.FillExecutionIdentifier())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

func (m IndividualTradePositionPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10112)
	w.StringField("m_Symbol", m.Symbol())
	w.Uint8Field("m_IsSimulated", m.IsSimulated())
	w.StringField("m_TradeAccount", m.TradeAccount())
	w.Float64Field("m_Quantity", m.Quantity())
	w.Float64Field("m_PositionPrice", m.PositionPrice())
	w.Float64Field("m_OpenProfitLoss", m.OpenProfitLoss())
	w.Int64Field("m_TradeDateTime", m.TradeDateTime())
	w.StringField("m_FillExecutionIdentifier", m.FillExecutionIdentifier())
	w.Uint8Field("m_IsSnapshot", m.IsSnapshot())
	w.Uint8Field("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.Uint8Field("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *IndividualTradePositionPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10112 {
		return message.ErrWrongType
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSimulated(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetPositionPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetOpenProfitLoss(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeDateTime(r.Int64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillExecutionIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSnapshot(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFirstMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsLastMessageInBatch(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *IndividualTradePositionPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10112 {
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
		case "m_Symbol":
			m.SetSymbol(r.String())
		case "m_IsSimulated":
			m.SetIsSimulated(r.Uint8())
		case "m_TradeAccount":
			m.SetTradeAccount(r.String())
		case "m_Quantity":
			m.SetQuantity(r.Float64())
		case "m_PositionPrice":
			m.SetPositionPrice(r.Float64())
		case "m_OpenProfitLoss":
			m.SetOpenProfitLoss(r.Float64())
		case "m_TradeDateTime":
			m.SetTradeDateTime(r.Int64())
		case "m_FillExecutionIdentifier":
			m.SetFillExecutionIdentifier(r.String())
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Uint8())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Uint8())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Uint8())
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
