//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdateOperationComplete) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10131,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleteOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolLimitsUpdateOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolCommissionUpdateOperation())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdateOperationCompleteBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10131,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleteOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolLimitsUpdateOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolCommissionUpdateOperation())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdateOperationCompletePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10131,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleteOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolLimitsUpdateOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolCommissionUpdateOperation())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdateOperationCompletePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":10131,\"F\":["...)
	w.Uint32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsError())
	w.Data = append(w.Data, ',')
	w.String(m.ErrorText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsDeleteOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolLimitsUpdateOperation())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsSymbolCommissionUpdateOperation())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdateOperationComplete) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10131)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsError", m.IsError())
	w.StringField("ErrorText", m.ErrorText())
	w.BoolField("IsDeleteOperation", m.IsDeleteOperation())
	w.BoolField("IsSymbolLimitsUpdateOperation", m.IsSymbolLimitsUpdateOperation())
	w.BoolField("IsSymbolCommissionUpdateOperation", m.IsSymbolCommissionUpdateOperation())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdateOperationCompleteBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10131)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsError", m.IsError())
	w.StringField("ErrorText", m.ErrorText())
	w.BoolField("IsDeleteOperation", m.IsDeleteOperation())
	w.BoolField("IsSymbolLimitsUpdateOperation", m.IsSymbolLimitsUpdateOperation())
	w.BoolField("IsSymbolCommissionUpdateOperation", m.IsSymbolCommissionUpdateOperation())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdateOperationCompletePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10131)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsError", m.IsError())
	w.StringField("ErrorText", m.ErrorText())
	w.BoolField("IsDeleteOperation", m.IsDeleteOperation())
	w.BoolField("IsSymbolLimitsUpdateOperation", m.IsSymbolLimitsUpdateOperation())
	w.BoolField("IsSymbolCommissionUpdateOperation", m.IsSymbolCommissionUpdateOperation())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

func (m TradeAccountDataUpdateOperationCompletePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10131)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsError", m.IsError())
	w.StringField("ErrorText", m.ErrorText())
	w.BoolField("IsDeleteOperation", m.IsDeleteOperation())
	w.BoolField("IsSymbolLimitsUpdateOperation", m.IsSymbolLimitsUpdateOperation())
	w.BoolField("IsSymbolCommissionUpdateOperation", m.IsSymbolCommissionUpdateOperation())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUpdateOperationCompleteBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10131 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsError(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetErrorText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDeleteOperation(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSymbolLimitsUpdateOperation(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSymbolCommissionUpdateOperation(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 10131 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Uint32())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsError(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetErrorText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsDeleteOperation(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSymbolLimitsUpdateOperation(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsSymbolCommissionUpdateOperation(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUpdateOperationCompleteBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10131 {
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
		case "IsError":
			m.SetIsError(r.Bool())
		case "ErrorText":
			m.SetErrorText(r.String())
		case "IsDeleteOperation":
			m.SetIsDeleteOperation(r.Bool())
		case "IsSymbolLimitsUpdateOperation":
			m.SetIsSymbolLimitsUpdateOperation(r.Bool())
		case "IsSymbolCommissionUpdateOperation":
			m.SetIsSymbolCommissionUpdateOperation(r.Bool())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *TradeAccountDataUpdateOperationCompletePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 10131 {
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
		case "IsError":
			m.SetIsError(r.Bool())
		case "ErrorText":
			m.SetErrorText(r.String())
		case "IsDeleteOperation":
			m.SetIsDeleteOperation(r.Bool())
		case "IsSymbolLimitsUpdateOperation":
			m.SetIsSymbolLimitsUpdateOperation(r.Bool())
		case "IsSymbolCommissionUpdateOperation":
			m.SetIsSymbolCommissionUpdateOperation(r.Bool())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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
