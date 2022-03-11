//go:build !tinygo

package model

import (
	"github.com/moontrade/dtc-go/message/pb"
)

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
