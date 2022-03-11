//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
)

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonRequestFixedPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":1,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.String(m.Password())
	w.Data = append(w.Data, ',')
	w.String(m.GeneralTextData())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_2())
	w.Data = append(w.Data, ',')
	w.Int32(m.HeartbeatIntervalInSeconds())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TradeMode()))
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.HardwareIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.ClientName())
	w.Data = append(w.Data, ',')
	w.Int32(m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}

func (m LogonRequestFixedPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
	w := json.Writer{Data: b, Compact: true}
	w.Data = append(w.Data, "{\"Type\":1,\"F\":["...)
	w.Int32(m.ProtocolVersion())
	w.Data = append(w.Data, ',')
	w.String(m.Username())
	w.Data = append(w.Data, ',')
	w.String(m.Password())
	w.Data = append(w.Data, ',')
	w.String(m.GeneralTextData())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_1())
	w.Data = append(w.Data, ',')
	w.Int32(m.Integer_2())
	w.Data = append(w.Data, ',')
	w.Int32(m.HeartbeatIntervalInSeconds())
	w.Data = append(w.Data, ',')
	w.Int32(int32(m.TradeMode()))
	w.Data = append(w.Data, ',')
	w.String(m.TradeAccount())
	w.Data = append(w.Data, ',')
	w.String(m.HardwareIdentifier())
	w.Data = append(w.Data, ',')
	w.String(m.ClientName())
	w.Data = append(w.Data, ',')
	w.Int32(m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonRequestFixedPointer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 1)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.StringField("Username", m.Username())
	w.StringField("Password", m.Password())
	w.StringField("GeneralTextData", m.GeneralTextData())
	w.Int32Field("Integer_1", m.Integer_1())
	w.Int32Field("Integer_2", m.Integer_2())
	w.Int32Field("HeartbeatIntervalInSeconds", m.HeartbeatIntervalInSeconds())
	w.Int32Field("TradeMode", int32(m.TradeMode()))
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("HardwareIdentifier", m.HardwareIdentifier())
	w.StringField("ClientName", m.ClientName())
	w.Int32Field("MarketDataTransmissionInterval", m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}

func (m LogonRequestFixedPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 1)
	w.Int32Field("ProtocolVersion", m.ProtocolVersion())
	w.StringField("Username", m.Username())
	w.StringField("Password", m.Password())
	w.StringField("GeneralTextData", m.GeneralTextData())
	w.Int32Field("Integer_1", m.Integer_1())
	w.Int32Field("Integer_2", m.Integer_2())
	w.Int32Field("HeartbeatIntervalInSeconds", m.HeartbeatIntervalInSeconds())
	w.Int32Field("TradeMode", int32(m.TradeMode()))
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("HardwareIdentifier", m.HardwareIdentifier())
	w.StringField("ClientName", m.ClientName())
	w.Int32Field("MarketDataTransmissionInterval", m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonRequestFixedPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
	if r.Type != 1 {
		return message.ErrWrongType
	}
	m.SetProtocolVersion(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetUsername(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetPassword(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetGeneralTextData(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_1(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetInteger_2(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetHeartbeatIntervalInSeconds(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeMode(TradeModeEnum(r.Int32()))
	if r.IsError() {
		return r.Error()
	}
	m.SetTradeAccount(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetHardwareIdentifier(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetClientName(r.String())
	if r.IsError() {
		return r.Error()
	}
	m.SetMarketDataTransmissionInterval(r.Int32())
	if r.IsError() {
		return r.Error()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonRequestFixedPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
	if r.Type != 1 {
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
		case "ProtocolVersion":
			m.SetProtocolVersion(r.Int32())
		case "Username":
			m.SetUsername(r.String())
		case "Password":
			m.SetPassword(r.String())
		case "GeneralTextData":
			m.SetGeneralTextData(r.String())
		case "Integer_1":
			m.SetInteger_1(r.Int32())
		case "Integer_2":
			m.SetInteger_2(r.Int32())
		case "HeartbeatIntervalInSeconds":
			m.SetHeartbeatIntervalInSeconds(r.Int32())
		case "TradeMode":
			m.SetTradeMode(TradeModeEnum(r.Int32()))
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "HardwareIdentifier":
			m.SetHardwareIdentifier(r.String())
		case "ClientName":
			m.SetClientName(r.String())
		case "MarketDataTransmissionInterval":
			m.SetMarketDataTransmissionInterval(r.Int32())
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
