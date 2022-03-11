//go:build !tinygo

package sc

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m FlattenPositionsForTradeAccount) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":210,\"F\":["...)
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	return w.Finish(), nil
}

func (m FlattenPositionsForTradeAccountBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":210,\"F\":["...)
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m FlattenPositionsForTradeAccountPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":210,\"F\":["...)
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	return w.Finish(), nil
}

func (m FlattenPositionsForTradeAccountPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":210,\"F\":["...)
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	w.Data = append(w.Data, ',')
	w.Bool(m.IsAutomatedOrder())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m FlattenPositionsForTradeAccount) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 210)
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	return w.Finish(), nil
}

func (m FlattenPositionsForTradeAccountBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 210)
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m FlattenPositionsForTradeAccountPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 210)
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	return w.Finish(), nil
}

func (m FlattenPositionsForTradeAccountPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 210)
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *FlattenPositionsForTradeAccountBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 210 {
		return message.ErrWrongType
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *FlattenPositionsForTradeAccountPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 210 {
		return message.ErrWrongType
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetIsAutomatedOrder(r.Bool())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *FlattenPositionsForTradeAccountBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 210 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
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

func (m *FlattenPositionsForTradeAccountPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 210 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
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
