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

func (m HistoricalAccountBalanceResponse) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponsePointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponsePointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponsePointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponsePointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 606 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionId(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponsePointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 606 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionId(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 606 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionId":
			m.SetTransactionId(r.String())
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

func (m *HistoricalAccountBalanceResponsePointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 606 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionId":
			m.SetTransactionId(r.String())
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

func (m HistoricalAccountBalanceResponseFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponseFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":606,\"F\":["...)
	w.Int32(m.RequestID())
	w.Data = append(w.Data, ',')
	w.Float64(float64(m.DateTime()))
	w.Data = append(w.Data, ',')
	w.Float64(m.CashBalance())
	w.Data = append(w.Data, ',')
	w.String(m.AccountCurrency())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsFinalResponse())
	w.Data = append(w.Data, ',')
	w.Uint8(m.NoAccountBalances())
	w.Data = append(w.Data, ',')
	w.String(m.InfoText())
	w.Data = append(w.Data, ',')
	w.String(m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponseFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

func (m HistoricalAccountBalanceResponseFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 606)
	w.Int32Field("RequestID", m.RequestID())
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Float64Field("CashBalance", m.CashBalance())
	w.StringField("AccountCurrency", m.AccountCurrency())
	w.StringField("TradeAccount", m.TradeAccount())
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	w.Uint8Field("NoAccountBalances", m.NoAccountBalances())
	w.StringField("InfoText", m.InfoText())
	w.StringField("TransactionId", m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 606 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionId(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 606 {
		return message.ErrWrongType
	}
	m.SetRequestID(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
	if r.IsError() {
		return r.Error()
	}
	m.SetCashBalance(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetAccountCurrency(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsFinalResponse(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	m.SetNoAccountBalances(r.Uint8())
	if r.IsError() {
		return r.Error()
	}
	m.SetInfoText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTransactionId(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 606 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionId":
			m.SetTransactionId(r.String())
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

func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 606 {
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
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "CashBalance":
			m.SetCashBalance(r.Float64())
		case "AccountCurrency":
			m.SetAccountCurrency(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
		case "NoAccountBalances":
			m.SetNoAccountBalances(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "TransactionId":
			m.SetTransactionId(r.String())
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

func (m *HistoricalAccountBalanceResponseBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetCashBalance(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetIsFinalResponse(r.ReadBool())
		case 7:
			m.SetNoAccountBalances(r.ReadUint8())
		case 8:
			m.SetInfoText(r.ReadString())
		case 9:
			m.SetTransactionId(r.ReadString())
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

func (m *HistoricalAccountBalanceResponsePointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetCashBalance(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetIsFinalResponse(r.ReadBool())
		case 7:
			m.SetNoAccountBalances(r.ReadUint8())
		case 8:
			m.SetInfoText(r.ReadString())
		case 9:
			m.SetTransactionId(r.ReadString())
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

func (m HistoricalAccountBalanceResponseBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 606)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteFixed64Float64(3, m.CashBalance())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteBool(6, m.IsFinalResponse())
	w.WriteUvarint8(7, m.NoAccountBalances())
	w.WriteString(8, m.InfoText())
	w.WriteString(9, m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponsePointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 606)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteFixed64Float64(3, m.CashBalance())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteBool(6, m.IsFinalResponse())
	w.WriteUvarint8(7, m.NoAccountBalances())
	w.WriteString(8, m.InfoText())
	w.WriteString(9, m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalanceResponseFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetCashBalance(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetIsFinalResponse(r.ReadBool())
		case 7:
			m.SetNoAccountBalances(r.ReadUint8())
		case 8:
			m.SetInfoText(r.ReadString())
		case 9:
			m.SetTransactionId(r.ReadString())
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

func (m *HistoricalAccountBalanceResponseFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetDateTime(DateTimeWithMilliseconds(r.ReadFloat64()))
		case 3:
			m.SetCashBalance(r.ReadFloat64())
		case 4:
			m.SetAccountCurrency(r.ReadString())
		case 5:
			m.SetTradeAccount(r.ReadString())
		case 6:
			m.SetIsFinalResponse(r.ReadBool())
		case 7:
			m.SetNoAccountBalances(r.ReadUint8())
		case 8:
			m.SetInfoText(r.ReadString())
		case 9:
			m.SetTransactionId(r.ReadString())
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

func (m HistoricalAccountBalanceResponseFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 606)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteFixed64Float64(3, m.CashBalance())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteBool(6, m.IsFinalResponse())
	w.WriteUvarint8(7, m.NoAccountBalances())
	w.WriteString(8, m.InfoText())
	w.WriteString(9, m.TransactionId())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalanceResponseFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 606)
	w.WriteVarint32(1, m.RequestID())
	w.WriteFixed64Float64(2, float64(m.DateTime()))
	w.WriteFixed64Float64(3, m.CashBalance())
	w.WriteString(4, m.AccountCurrency())
	w.WriteString(5, m.TradeAccount())
	w.WriteBool(6, m.IsFinalResponse())
	w.WriteUvarint8(7, m.NoAccountBalances())
	w.WriteString(8, m.InfoText())
	w.WriteString(9, m.TransactionId())
	return w.Finish(), nil
}
