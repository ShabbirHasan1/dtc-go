// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDepthUpdateLevelSize = 56

//     Size        uint16                     = MarketDepthUpdateLevelSize  (56)
//     Type        uint16                     = MARKET_DEPTH_UPDATE_LEVEL  (106)
//     SymbolID    uint32                     = 0
//     Side        AtBidOrAskEnum             = BID_ASK_UNSET  (0)
//     Price       float64                    = 0.000000
//     Quantity    float64                    = 0.000000
//     UpdateType  MarketDepthUpdateTypeEnum  = MARKET_DEPTH_UNSET  (0)
//     DateTime    DateTimeWithMilliseconds   = 0.000000
//     NumOrders   uint32                     = 0
var _MarketDepthUpdateLevelDefault = []byte{56, 0, 106, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDepthUpdateLevel Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// Each MarketDepthUpdateLevel message updates one level of market depth
// on one side. An insert/update/delete model is used for market depth.
//
// The Client will need to determine the based upon the price, what particular
// market depth level is being updated, inserted or deleted.
//
// It is for this reason, that an insert/update is considered as one update
// type since it is possible to determine whether it is an insert or update
// based upon the existence of the price level in the existing market depth
// book on the Client side.
//
// What this means is that when the UpdateType field is MARKET_DEPTH_INSERT_UPDATE_LEVEL,
// it is considered an insert if the price level is not found on the particular
// side of the market depth being updated. It is considered an update, if
// the price level is found on the particular side of market depth being
// updated.
//
// This message uses a double datatype for the Price field. There is no level
// index. It is the responsibility of the Client to determine where in its
// market depth array it is maintaining where the insert/update/delete operation
// needs to occur.
//
// Since floating-point comparisons are not always precise, there should
// be a comparison made only to the number of decimal places the symbol specifies
// in its security definition. This can be determined through the SecurityDefinitionResponse::PriceDisplayFormat
// field.
type MarketDepthUpdateLevel struct {
	p message.Fixed
}

func NewMarketDepthUpdateLevelFrom(b []byte) MarketDepthUpdateLevel {
	return MarketDepthUpdateLevel{p: message.NewFixed(b)}
}

func WrapMarketDepthUpdateLevel(b []byte) MarketDepthUpdateLevel {
	return MarketDepthUpdateLevel{p: message.WrapFixed(b)}
}

func NewMarketDepthUpdateLevel() *MarketDepthUpdateLevel {
	return &MarketDepthUpdateLevel{p: message.NewFixed(_MarketDepthUpdateLevelDefault)}
}

func ParseMarketDepthUpdateLevel(b []byte) (MarketDepthUpdateLevel, error) {
	if len(b) < 4 {
		return MarketDepthUpdateLevel{}, message.ErrShortBuffer
	}
	m := WrapMarketDepthUpdateLevel(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDepthUpdateLevel{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDepthUpdateLevel{}, message.ErrBaseSizeOverflow
	}
	if size < 56 {
		clone := *NewMarketDepthUpdateLevel()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDepthUpdateLevelDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDepthUpdateLevel) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDepthUpdateLevel) Type() uint16 {
	return m.p.Uint16LE(2)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthUpdateLevel) SymbolID() uint32 {
	return m.p.Uint32LE(4)
}

// Side Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m MarketDepthUpdateLevel) Side() AtBidOrAskEnum {
	return AtBidOrAskEnum(m.p.Uint16LE(8))
}

// Price The price level to insert, update or delete.
func (m MarketDepthUpdateLevel) Price() float64 {
	return m.p.Float64LE(16)
}

// Quantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m MarketDepthUpdateLevel) Quantity() float64 {
	return m.p.Float64LE(24)
}

// UpdateType Specifies whether this is a MARKET_DEPTH_INSERT_UPDATE_LEVEL operation
// or a MARKET_DEPTH_DELETE_LEVEL operation.
//
// MARKET_DEPTH_INSERT_UPDATE_LEVEL: Insert or update in the market depth
// book on the specified side, the particular Price and Volume specified.
// It is an insert operation of the price level does not exist. It is an
// update operation if the price level already exists. In the case of insert,
// the other levels in the market depth book need to be shifted to make room
// for the new level.
//
// MARKET_DEPTH_DELETE_LEVEL: Remove from the market depth book on the specified
// side, the specified Price level. The other levels need to be shifted to
// fill in the missing level. In this particular case the Quantity is ignored
// and will be 0.
func (m MarketDepthUpdateLevel) UpdateType() MarketDepthUpdateTypeEnum {
	return MarketDepthUpdateTypeEnum(m.p.Uint8(32))
}

// DateTime The Date-Time of the market depth update.
func (m MarketDepthUpdateLevel) DateTime() DateTimeWithMilliseconds {
	return DateTimeWithMilliseconds(m.p.Float64LE(40))
}

// NumOrders The number of orders at the Price.
func (m MarketDepthUpdateLevel) NumOrders() uint32 {
	return m.p.Uint32LE(48)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m *MarketDepthUpdateLevel) SetSymbolID(value uint32) *MarketDepthUpdateLevel {
	m.p.SetUint32LE(4, value)
	return m
}

// SetSide Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
// Specifies whether the side being updated is the Bid (AT_BID) or Ask (AT_ASK).
func (m *MarketDepthUpdateLevel) SetSide(value AtBidOrAskEnum) *MarketDepthUpdateLevel {
	m.p.SetUint16LE(8, uint16(value))
	return m
}

// SetPrice The price level to insert, update or delete.
func (m *MarketDepthUpdateLevel) SetPrice(value float64) *MarketDepthUpdateLevel {
	m.p.SetFloat64LE(16, value)
	return m
}

// SetQuantity The number of shares/contracts at the Price level. This will be 0 in the
// case when UpdateType is set to MARKET_DEPTH_DELETE_LEVEL.
func (m *MarketDepthUpdateLevel) SetQuantity(value float64) *MarketDepthUpdateLevel {
	m.p.SetFloat64LE(24, value)
	return m
}

// SetUpdateType Specifies whether this is a MARKET_DEPTH_INSERT_UPDATE_LEVEL operation
// or a MARKET_DEPTH_DELETE_LEVEL operation.
//
// MARKET_DEPTH_INSERT_UPDATE_LEVEL: Insert or update in the market depth
// book on the specified side, the particular Price and Volume specified.
// It is an insert operation of the price level does not exist. It is an
// update operation if the price level already exists. In the case of insert,
// the other levels in the market depth book need to be shifted to make room
// for the new level.
//
// MARKET_DEPTH_DELETE_LEVEL: Remove from the market depth book on the specified
// side, the specified Price level. The other levels need to be shifted to
// fill in the missing level. In this particular case the Quantity is ignored
// and will be 0.
func (m *MarketDepthUpdateLevel) SetUpdateType(value MarketDepthUpdateTypeEnum) *MarketDepthUpdateLevel {
	m.p.SetUint8(32, uint8(value))
	return m
}

// SetDateTime The Date-Time of the market depth update.
func (m *MarketDepthUpdateLevel) SetDateTime(value DateTimeWithMilliseconds) *MarketDepthUpdateLevel {
	m.p.SetFloat64LE(40, float64(value))
	return m
}

// SetNumOrders The number of orders at the Price.
func (m *MarketDepthUpdateLevel) SetNumOrders(value uint32) *MarketDepthUpdateLevel {
	m.p.SetUint32LE(48, value)
	return m
}

func (m *MarketDepthUpdateLevel) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *MarketDepthUpdateLevel) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *MarketDepthUpdateLevel) Clone() *MarketDepthUpdateLevel {
	return &MarketDepthUpdateLevel{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m MarketDepthUpdateLevel) Copy(to MarketDepthUpdateLevel) {
	to.SetSymbolID(m.SymbolID())
	to.SetSide(m.Side())
	to.SetPrice(m.Price())
	to.SetQuantity(m.Quantity())
	to.SetUpdateType(m.UpdateType())
	to.SetDateTime(m.DateTime())
	to.SetNumOrders(m.NumOrders())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevel) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *MarketDepthUpdateLevel) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 106)
	w.Uint32Field("SymbolID", m.SymbolID())
	w.Uint16Field("Side", uint16(m.Side()))
	w.Float64Field("Price", m.Price())
	w.Float64Field("Quantity", m.Quantity())
	w.Uint8Field("UpdateType", uint8(m.UpdateType()))
	w.Float64Field("DateTime", float64(m.DateTime()))
	w.Uint32Field("NumOrders", m.NumOrders())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthUpdateLevel) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *MarketDepthUpdateLevel) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 106 {
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
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Side":
			m.SetSide(AtBidOrAskEnum(r.Uint16()))
		case "Price":
			m.SetPrice(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UpdateType":
			m.SetUpdateType(MarketDepthUpdateTypeEnum(r.Uint8()))
		case "DateTime":
			m.SetDateTime(DateTimeWithMilliseconds(r.Float64()))
		case "NumOrders":
			m.SetNumOrders(r.Uint32())
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
