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

func (m AddCorrectingOrderFill) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFill) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 309 {
		return message.ErrWrongType
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 309 {
		return message.ErrWrongType
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 309 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "FillPrice":
			m.SetFillPrice(r.Float64())
		case "FillQuantity":
			m.SetFillQuantity(r.Float64())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
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

func (m *AddCorrectingOrderFillPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 309 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "FillPrice":
			m.SetFillPrice(r.Float64())
		case "FillQuantity":
			m.SetFillQuantity(r.Float64())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
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

func (m AddCorrectingOrderFillFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":309,\"F\":["...)
	w.String(m.Symbol())
	w.Data = append(w.Data, ',')
	w.String(m.Exchange())
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.ClientOrderID())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.BuySell()))
	w.Data = append(w.Data, ',')
	w.Float64(m.FillPrice())
	w.Data = append(w.Data, ',')
	w.Float64(m.FillQuantity())
	w.Data = append(w.Data, ',')
	w.String(m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

func (m AddCorrectingOrderFillFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 309)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("FillPrice", m.FillPrice())
	w.Float64Field("FillQuantity", m.FillQuantity())
	w.StringField("FreeFormText", m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 309 {
		return message.ErrWrongType
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 309 {
		return message.ErrWrongType
	}
	m.SetSymbol(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetExchange(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientOrderID(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetBuySell(BuySellEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetFillPrice(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFillQuantity(r.Float64())
	if r.IsError() {
		return r.Error()
	}
	m.SetFreeFormText(r.String())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 309 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "FillPrice":
			m.SetFillPrice(r.Float64())
		case "FillQuantity":
			m.SetFillQuantity(r.Float64())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
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

func (m *AddCorrectingOrderFillFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 309 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "FillPrice":
			m.SetFillPrice(r.Float64())
		case "FillQuantity":
			m.SetFillQuantity(r.Float64())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
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

func (m *AddCorrectingOrderFillBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 6:
			m.SetFillPrice(r.ReadFloat64())
		case 7:
			m.SetFillQuantity(r.ReadFloat64())
		case 8:
			m.SetFreeFormText(r.ReadString())
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

func (m *AddCorrectingOrderFillPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 6:
			m.SetFillPrice(r.ReadFloat64())
		case 7:
			m.SetFillQuantity(r.ReadFloat64())
		case 8:
			m.SetFreeFormText(r.ReadString())
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

func (m AddCorrectingOrderFillBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 309)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.BuySell()))
	w.WriteFixed64Float64(6, m.FillPrice())
	w.WriteFixed64Float64(7, m.FillQuantity())
	w.WriteString(8, m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 309)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.BuySell()))
	w.WriteFixed64Float64(6, m.FillPrice())
	w.WriteFixed64Float64(7, m.FillQuantity())
	w.WriteString(8, m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AddCorrectingOrderFillFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 6:
			m.SetFillPrice(r.ReadFloat64())
		case 7:
			m.SetFillQuantity(r.ReadFloat64())
		case 8:
			m.SetFreeFormText(r.ReadString())
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

func (m *AddCorrectingOrderFillFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetSymbol(r.ReadString())
		case 2:
			m.SetExchange(r.ReadString())
		case 3:
			m.SetTradeAccount(r.ReadString())
		case 4:
			m.SetClientOrderID(r.ReadString())
		case 5:
			m.SetBuySell(BuySellEnum(r.ReadInt32()))
		case 6:
			m.SetFillPrice(r.ReadFloat64())
		case 7:
			m.SetFillQuantity(r.ReadFloat64())
		case 8:
			m.SetFreeFormText(r.ReadString())
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

func (m AddCorrectingOrderFillFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 309)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.BuySell()))
	w.WriteFixed64Float64(6, m.FillPrice())
	w.WriteFixed64Float64(7, m.FillQuantity())
	w.WriteString(8, m.FreeFormText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AddCorrectingOrderFillFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 309)
	w.WriteString(1, m.Symbol())
	w.WriteString(2, m.Exchange())
	w.WriteString(3, m.TradeAccount())
	w.WriteString(4, m.ClientOrderID())
	w.WriteVarint32(5, int32(m.BuySell()))
	w.WriteFixed64Float64(6, m.FillPrice())
	w.WriteFixed64Float64(7, m.FillQuantity())
	w.WriteString(8, m.FreeFormText())
	return w.Finish(), nil
}
