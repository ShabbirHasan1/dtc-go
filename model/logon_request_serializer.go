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

func (m LogonRequest) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m LogonRequestBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonRequestPointer) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m LogonRequestPointerBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m LogonRequest) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m LogonRequestBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonRequestPointer) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m LogonRequestPointerBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m *LogonRequestBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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
// JSON Compact Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonRequestPointerBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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

func (m *LogonRequestBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonRequestPointerBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Compact Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonRequestFixed) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m LogonRequestFixedBuilder) MarshalJSONCompactTo(b []byte) ([]byte, error) {
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

func (m LogonRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m LogonRequestFixedBuilder) MarshalJSONTo(b []byte) ([]byte, error) {
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

func (m *LogonRequestFixedBuilder) UnmarshalJSONCompactFrom(r *json.Reader) error {
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

func (m *LogonRequestFixedBuilder) UnmarshalJSONFrom(r *json.Reader) error {
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

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonRequestBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetUsername(r.ReadString())
		case 3:
			m.SetPassword(r.ReadString())
		case 4:
			m.SetGeneralTextData(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetInteger_2(r.ReadInt32())
		case 7:
			m.SetHeartbeatIntervalInSeconds(r.ReadInt32())
		case 8:
			m.SetTradeMode(TradeModeEnum(r.ReadInt32()))
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetHardwareIdentifier(r.ReadString())
		case 11:
			m.SetClientName(r.ReadString())
		case 12:
			m.SetMarketDataTransmissionInterval(r.ReadInt32())
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

func (m *LogonRequestPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetUsername(r.ReadString())
		case 3:
			m.SetPassword(r.ReadString())
		case 4:
			m.SetGeneralTextData(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetInteger_2(r.ReadInt32())
		case 7:
			m.SetHeartbeatIntervalInSeconds(r.ReadInt32())
		case 8:
			m.SetTradeMode(TradeModeEnum(r.ReadInt32()))
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetHardwareIdentifier(r.ReadString())
		case 11:
			m.SetClientName(r.ReadString())
		case 12:
			m.SetMarketDataTransmissionInterval(r.ReadInt32())
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

func (m LogonRequestBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 1)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteString(2, m.Username())
	w.WriteString(3, m.Password())
	w.WriteString(4, m.GeneralTextData())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteVarint32(6, m.Integer_2())
	w.WriteVarint32(7, m.HeartbeatIntervalInSeconds())
	w.WriteVarint32(8, int32(m.TradeMode()))
	w.WriteString(9, m.TradeAccount())
	w.WriteString(10, m.HardwareIdentifier())
	w.WriteString(11, m.ClientName())
	w.WriteVarint32(12, m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonRequestPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 1)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteString(2, m.Username())
	w.WriteString(3, m.Password())
	w.WriteString(4, m.GeneralTextData())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteVarint32(6, m.Integer_2())
	w.WriteVarint32(7, m.HeartbeatIntervalInSeconds())
	w.WriteVarint32(8, int32(m.TradeMode()))
	w.WriteString(9, m.TradeAccount())
	w.WriteString(10, m.HardwareIdentifier())
	w.WriteString(11, m.ClientName())
	w.WriteVarint32(12, m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *LogonRequestFixedBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetUsername(r.ReadString())
		case 3:
			m.SetPassword(r.ReadString())
		case 4:
			m.SetGeneralTextData(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetInteger_2(r.ReadInt32())
		case 7:
			m.SetHeartbeatIntervalInSeconds(r.ReadInt32())
		case 8:
			m.SetTradeMode(TradeModeEnum(r.ReadInt32()))
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetHardwareIdentifier(r.ReadString())
		case 11:
			m.SetClientName(r.ReadString())
		case 12:
			m.SetMarketDataTransmissionInterval(r.ReadInt32())
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

func (m *LogonRequestFixedPointerBuilder) UnmarshalProtobuf(b []byte) error {
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
			m.SetProtocolVersion(r.ReadInt32())
		case 2:
			m.SetUsername(r.ReadString())
		case 3:
			m.SetPassword(r.ReadString())
		case 4:
			m.SetGeneralTextData(r.ReadString())
		case 5:
			m.SetInteger_1(r.ReadInt32())
		case 6:
			m.SetInteger_2(r.ReadInt32())
		case 7:
			m.SetHeartbeatIntervalInSeconds(r.ReadInt32())
		case 8:
			m.SetTradeMode(TradeModeEnum(r.ReadInt32()))
		case 9:
			m.SetTradeAccount(r.ReadString())
		case 10:
			m.SetHardwareIdentifier(r.ReadString())
		case 11:
			m.SetClientName(r.ReadString())
		case 12:
			m.SetMarketDataTransmissionInterval(r.ReadInt32())
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

func (m LogonRequestFixedBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 1)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteString(2, m.Username())
	w.WriteString(3, m.Password())
	w.WriteString(4, m.GeneralTextData())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteVarint32(6, m.Integer_2())
	w.WriteVarint32(7, m.HeartbeatIntervalInSeconds())
	w.WriteVarint32(8, int32(m.TradeMode()))
	w.WriteString(9, m.TradeAccount())
	w.WriteString(10, m.HardwareIdentifier())
	w.WriteString(11, m.ClientName())
	w.WriteVarint32(12, m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Protocol Buffers Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m LogonRequestFixedPointerBuilder) MarshalProtobuf(b []byte) ([]byte, error) {
	w := pb.NewWriter(b, 1)
	w.WriteVarint32(1, m.ProtocolVersion())
	w.WriteString(2, m.Username())
	w.WriteString(3, m.Password())
	w.WriteString(4, m.GeneralTextData())
	w.WriteVarint32(5, m.Integer_1())
	w.WriteVarint32(6, m.Integer_2())
	w.WriteVarint32(7, m.HeartbeatIntervalInSeconds())
	w.WriteVarint32(8, int32(m.TradeMode()))
	w.WriteString(9, m.TradeAccount())
	w.WriteString(10, m.HardwareIdentifier())
	w.WriteString(11, m.ClientName())
	w.WriteVarint32(12, m.MarketDataTransmissionInterval())
	return w.Finish(), nil
}
