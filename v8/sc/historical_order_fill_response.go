// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const HistoricalOrderFillResponseSize = 112

const HistoricalOrderFillResponseFixedSize = 384

//     Size                     uint16              = HistoricalOrderFillResponseSize  (112)
//     Type                     uint16              = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     BaseSize                 uint16              = HistoricalOrderFillResponseSize  (112)
//     RequestID                int32               = 0
//     TotalNumberMessages      int32               = 0
//     MessageNumber            int32               = 0
//     Symbol                   string              = ""
//     Exchange                 string              = ""
//     ServerOrderID            string              = ""
//     BuySell                  BuySellEnum         = BUY_SELL_UNSET  (0)
//     Price                    float64             = 0.000000
//     DateTime                 DateTime            = 0
//     Quantity                 float64             = 0.000000
//     UniqueExecutionID        string              = ""
//     TradeAccount             string              = ""
//     OpenClose                OpenCloseTradeEnum  = TRADE_UNSET  (0)
//     NoOrderFills             uint8               = 0
//     InfoText                 string              = ""
//     HighPriceDuringPosition  float64             = 0.000000
//     LowPriceDuringPosition   float64             = 0.000000
//     PositionQuantity         float64             = math.MaxFloat64
var _HistoricalOrderFillResponseDefault = []byte{112, 0, 48, 1, 112, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127}

//     Size                     uint16              = HistoricalOrderFillResponseFixedSize  (384)
//     Type                     uint16              = HISTORICAL_ORDER_FILL_RESPONSE  (304)
//     RequestID                int32               = 0
//     TotalNumberMessages      int32               = 0
//     MessageNumber            int32               = 0
//     Symbol                   string[64]          = ""
//     Exchange                 string[16]          = ""
//     ServerOrderID            string[32]          = ""
//     BuySell                  BuySellEnum         = BUY_SELL_UNSET  (0)
//     Price                    float64             = 0.000000
//     DateTime                 DateTime            = 0
//     Quantity                 float64             = 0.000000
//     UniqueExecutionID        string[64]          = ""
//     TradeAccount             string[32]          = ""
//     OpenClose                OpenCloseTradeEnum  = TRADE_UNSET  (0)
//     NoOrderFills             uint8               = 0
//     InfoText                 string[96]          = ""
//     HighPriceDuringPosition  float64             = 0.000000
//     LowPriceDuringPosition   float64             = 0.000000
//     PositionQuantity         float64             = math.MaxFloat64
var _HistoricalOrderFillResponseFixedDefault = []byte{128, 1, 48, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 255, 255, 239, 127}

// HistoricalOrderFillResponse This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponse struct {
	p message.VLS
}

// HistoricalOrderFillResponseFixed This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
type HistoricalOrderFillResponseFixed struct {
	p message.Fixed
}

func NewHistoricalOrderFillResponseFrom(b []byte) HistoricalOrderFillResponse {
	return HistoricalOrderFillResponse{p: message.NewVLS(b)}
}

func WrapHistoricalOrderFillResponse(b []byte) HistoricalOrderFillResponse {
	return HistoricalOrderFillResponse{p: message.WrapVLS(b)}
}

func NewHistoricalOrderFillResponse() *HistoricalOrderFillResponse {
	return &HistoricalOrderFillResponse{p: message.NewVLS(_HistoricalOrderFillResponseDefault)}
}

func ParseHistoricalOrderFillResponse(b []byte) (HistoricalOrderFillResponse, error) {
	if len(b) < 6 {
		return HistoricalOrderFillResponse{}, message.ErrShortBuffer
	}
	m := WrapHistoricalOrderFillResponse(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalOrderFillResponse{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return HistoricalOrderFillResponse{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 112 {
		newSize := len(b) + (112 - baseSize)
		if newSize > message.MaxSize {
			return HistoricalOrderFillResponse{}, message.ErrOverflow
		}
		clone := HistoricalOrderFillResponse{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _HistoricalOrderFillResponseDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(112 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(20)
			if offset > 0 {
				clone.p.SetUint16LE(20, offset+shift)
			}
			offset = clone.p.Uint16LE(24)
			if offset > 0 {
				clone.p.SetUint16LE(24, offset+shift)
			}
			offset = clone.p.Uint16LE(28)
			if offset > 0 {
				clone.p.SetUint16LE(28, offset+shift)
			}
			offset = clone.p.Uint16LE(64)
			if offset > 0 {
				clone.p.SetUint16LE(64, offset+shift)
			}
			offset = clone.p.Uint16LE(68)
			if offset > 0 {
				clone.p.SetUint16LE(68, offset+shift)
			}
			offset = clone.p.Uint16LE(78)
			if offset > 0 {
				clone.p.SetUint16LE(78, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewHistoricalOrderFillResponseFixedFrom(b []byte) HistoricalOrderFillResponseFixed {
	return HistoricalOrderFillResponseFixed{p: message.NewFixed(b)}
}

func WrapHistoricalOrderFillResponseFixed(b []byte) HistoricalOrderFillResponseFixed {
	return HistoricalOrderFillResponseFixed{p: message.WrapFixed(b)}
}

func NewHistoricalOrderFillResponseFixed() *HistoricalOrderFillResponseFixed {
	return &HistoricalOrderFillResponseFixed{p: message.NewFixed(_HistoricalOrderFillResponseFixedDefault)}
}

func ParseHistoricalOrderFillResponseFixed(b []byte) (HistoricalOrderFillResponseFixed, error) {
	if len(b) < 4 {
		return HistoricalOrderFillResponseFixed{}, message.ErrShortBuffer
	}
	m := WrapHistoricalOrderFillResponseFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalOrderFillResponseFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return HistoricalOrderFillResponseFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 384 {
		clone := *NewHistoricalOrderFillResponseFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _HistoricalOrderFillResponseFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m HistoricalOrderFillResponse) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m HistoricalOrderFillResponse) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m HistoricalOrderFillResponse) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponse) RequestID() int32 {
	return m.p.Int32LE(8)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponse) TotalNumberMessages() int32 {
	return m.p.Int32LE(12)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponse) MessageNumber() int32 {
	return m.p.Int32LE(16)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponse) Symbol() string {
	return m.p.StringVLS(20)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponse) Exchange() string {
	return m.p.StringVLS(24)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponse) ServerOrderID() string {
	return m.p.StringVLS(28)
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponse) BuySell() BuySellEnum {
	return BuySellEnum(m.p.Int32LE(32))
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponse) Price() float64 {
	return m.p.Float64LE(40)
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponse) DateTime() DateTime {
	return DateTime(m.p.Int64LE(48))
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponse) Quantity() float64 {
	return m.p.Float64LE(56)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponse) UniqueExecutionID() string {
	return m.p.StringVLS(64)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponse) TradeAccount() string {
	return m.p.StringVLS(68)
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponse) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(m.p.Int32LE(72))
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponse) NoOrderFills() uint8 {
	return m.p.Uint8(76)
}

// InfoText
func (m HistoricalOrderFillResponse) InfoText() string {
	return m.p.StringVLS(78)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponse) HighPriceDuringPosition() float64 {
	return m.p.Float64LE(88)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponse) LowPriceDuringPosition() float64 {
	return m.p.Float64LE(96)
}

// PositionQuantity
func (m HistoricalOrderFillResponse) PositionQuantity() float64 {
	return m.p.Float64LE(104)
}

// Size The standard message size field. Automatically set by constructor.
func (m HistoricalOrderFillResponseFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m HistoricalOrderFillResponseFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m HistoricalOrderFillResponseFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// TotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m HistoricalOrderFillResponseFixed) TotalNumberMessages() int32 {
	return m.p.Int32LE(8)
}

// MessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m HistoricalOrderFillResponseFixed) MessageNumber() int32 {
	return m.p.Int32LE(12)
}

// Symbol The symbol the order fill is for.
func (m HistoricalOrderFillResponseFixed) Symbol() string {
	return m.p.StringFixed(16, 64)
}

// Exchange The optional exchange for the symbol.
func (m HistoricalOrderFillResponseFixed) Exchange() string {
	return m.p.StringFixed(80, 16)
}

// ServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m HistoricalOrderFillResponseFixed) ServerOrderID() string {
	return m.p.StringFixed(96, 32)
}

// BuySell This is the fill side and can be either BUY or SELL.
func (m HistoricalOrderFillResponseFixed) BuySell() BuySellEnum {
	return BuySellEnum(m.p.Int32LE(128))
}

// Price This is the price of the order fill.
func (m HistoricalOrderFillResponseFixed) Price() float64 {
	return m.p.Float64LE(136)
}

// DateTime This is the Date and Time of the order fill.
func (m HistoricalOrderFillResponseFixed) DateTime() DateTime {
	return DateTime(m.p.Int64LE(144))
}

// Quantity This is the quantity of the order fill.
func (m HistoricalOrderFillResponseFixed) Quantity() float64 {
	return m.p.Float64LE(152)
}

// UniqueExecutionID This is the unique execution identifier for the order fill.
func (m HistoricalOrderFillResponseFixed) UniqueExecutionID() string {
	return m.p.StringFixed(160, 64)
}

// TradeAccount This is the trade account that the order fill is associated with.
func (m HistoricalOrderFillResponseFixed) TradeAccount() string {
	return m.p.StringFixed(224, 32)
}

// OpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m HistoricalOrderFillResponseFixed) OpenClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(m.p.Int32LE(256))
}

// NoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m HistoricalOrderFillResponseFixed) NoOrderFills() uint8 {
	return m.p.Uint8(260)
}

// InfoText
func (m HistoricalOrderFillResponseFixed) InfoText() string {
	return m.p.StringFixed(261, 96)
}

// HighPriceDuringPosition
func (m HistoricalOrderFillResponseFixed) HighPriceDuringPosition() float64 {
	return m.p.Float64LE(360)
}

// LowPriceDuringPosition
func (m HistoricalOrderFillResponseFixed) LowPriceDuringPosition() float64 {
	return m.p.Float64LE(368)
}

// PositionQuantity
func (m HistoricalOrderFillResponseFixed) PositionQuantity() float64 {
	return m.p.Float64LE(376)
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m *HistoricalOrderFillResponse) SetRequestID(value int32) *HistoricalOrderFillResponse {
	m.p.SetInt32LE(8, value)
	return m
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m *HistoricalOrderFillResponse) SetTotalNumberMessages(value int32) *HistoricalOrderFillResponse {
	m.p.SetInt32LE(12, value)
	return m
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m *HistoricalOrderFillResponse) SetMessageNumber(value int32) *HistoricalOrderFillResponse {
	m.p.SetInt32LE(16, value)
	return m
}

// SetSymbol The symbol the order fill is for.
func (m *HistoricalOrderFillResponse) SetSymbol(value string) *HistoricalOrderFillResponse {
	m.p.SetStringVLS(20, value)
	return m
}

// SetExchange The optional exchange for the symbol.
func (m *HistoricalOrderFillResponse) SetExchange(value string) *HistoricalOrderFillResponse {
	m.p.SetStringVLS(24, value)
	return m
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m *HistoricalOrderFillResponse) SetServerOrderID(value string) *HistoricalOrderFillResponse {
	m.p.SetStringVLS(28, value)
	return m
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m *HistoricalOrderFillResponse) SetBuySell(value BuySellEnum) *HistoricalOrderFillResponse {
	m.p.SetInt32LE(32, int32(value))
	return m
}

// SetPrice This is the price of the order fill.
func (m *HistoricalOrderFillResponse) SetPrice(value float64) *HistoricalOrderFillResponse {
	m.p.SetFloat64LE(40, value)
	return m
}

// SetDateTime This is the Date and Time of the order fill.
func (m *HistoricalOrderFillResponse) SetDateTime(value DateTime) *HistoricalOrderFillResponse {
	m.p.SetInt64LE(48, int64(value))
	return m
}

// SetQuantity This is the quantity of the order fill.
func (m *HistoricalOrderFillResponse) SetQuantity(value float64) *HistoricalOrderFillResponse {
	m.p.SetFloat64LE(56, value)
	return m
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m *HistoricalOrderFillResponse) SetUniqueExecutionID(value string) *HistoricalOrderFillResponse {
	m.p.SetStringVLS(64, value)
	return m
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m *HistoricalOrderFillResponse) SetTradeAccount(value string) *HistoricalOrderFillResponse {
	m.p.SetStringVLS(68, value)
	return m
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m *HistoricalOrderFillResponse) SetOpenClose(value OpenCloseTradeEnum) *HistoricalOrderFillResponse {
	m.p.SetInt32LE(72, int32(value))
	return m
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m *HistoricalOrderFillResponse) SetNoOrderFills(value uint8) *HistoricalOrderFillResponse {
	m.p.SetUint8(76, value)
	return m
}

// SetInfoText
func (m *HistoricalOrderFillResponse) SetInfoText(value string) *HistoricalOrderFillResponse {
	m.p.SetStringVLS(78, value)
	return m
}

// SetHighPriceDuringPosition
func (m *HistoricalOrderFillResponse) SetHighPriceDuringPosition(value float64) *HistoricalOrderFillResponse {
	m.p.SetFloat64LE(88, value)
	return m
}

// SetLowPriceDuringPosition
func (m *HistoricalOrderFillResponse) SetLowPriceDuringPosition(value float64) *HistoricalOrderFillResponse {
	m.p.SetFloat64LE(96, value)
	return m
}

// SetPositionQuantity
func (m *HistoricalOrderFillResponse) SetPositionQuantity(value float64) *HistoricalOrderFillResponse {
	m.p.SetFloat64LE(104, value)
	return m
}

// SetRequestID The RequestID specified in the HistoricalOrderFillsRequest message from
// the Client.
func (m *HistoricalOrderFillResponseFixed) SetRequestID(value int32) *HistoricalOrderFillResponseFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetTotalNumberMessages This indicates the total number of order fill reports when a batch of
// reports is being sent. If there is only one order fill report being sent,
// this will be 1.
func (m *HistoricalOrderFillResponseFixed) SetTotalNumberMessages(value int32) *HistoricalOrderFillResponseFixed {
	m.p.SetInt32LE(8, value)
	return m
}

// SetMessageNumber This indicates the 1-based index of the order fill report when a batch
// of reports is being sent. If there is only one order fill report being
// sent, this will be 1.
func (m *HistoricalOrderFillResponseFixed) SetMessageNumber(value int32) *HistoricalOrderFillResponseFixed {
	m.p.SetInt32LE(12, value)
	return m
}

// SetSymbol The symbol the order fill is for.
func (m *HistoricalOrderFillResponseFixed) SetSymbol(value string) *HistoricalOrderFillResponseFixed {
	m.p.SetStringFixed(16, 64, value)
	return m
}

// SetExchange The optional exchange for the symbol.
func (m *HistoricalOrderFillResponseFixed) SetExchange(value string) *HistoricalOrderFillResponseFixed {
	m.p.SetStringFixed(80, 16, value)
	return m
}

// SetServerOrderID This is the Server provided order identifier that the order fill was associated
// with.
func (m *HistoricalOrderFillResponseFixed) SetServerOrderID(value string) *HistoricalOrderFillResponseFixed {
	m.p.SetStringFixed(96, 32, value)
	return m
}

// SetBuySell This is the fill side and can be either BUY or SELL.
func (m *HistoricalOrderFillResponseFixed) SetBuySell(value BuySellEnum) *HistoricalOrderFillResponseFixed {
	m.p.SetInt32LE(128, int32(value))
	return m
}

// SetPrice This is the price of the order fill.
func (m *HistoricalOrderFillResponseFixed) SetPrice(value float64) *HistoricalOrderFillResponseFixed {
	m.p.SetFloat64LE(136, value)
	return m
}

// SetDateTime This is the Date and Time of the order fill.
func (m *HistoricalOrderFillResponseFixed) SetDateTime(value DateTime) *HistoricalOrderFillResponseFixed {
	m.p.SetInt64LE(144, int64(value))
	return m
}

// SetQuantity This is the quantity of the order fill.
func (m *HistoricalOrderFillResponseFixed) SetQuantity(value float64) *HistoricalOrderFillResponseFixed {
	m.p.SetFloat64LE(152, value)
	return m
}

// SetUniqueExecutionID This is the unique execution identifier for the order fill.
func (m *HistoricalOrderFillResponseFixed) SetUniqueExecutionID(value string) *HistoricalOrderFillResponseFixed {
	m.p.SetStringFixed(160, 64, value)
	return m
}

// SetTradeAccount This is the trade account that the order fill is associated with.
func (m *HistoricalOrderFillResponseFixed) SetTradeAccount(value string) *HistoricalOrderFillResponseFixed {
	m.p.SetStringFixed(224, 32, value)
	return m
}

// SetOpenClose Indicates whether this is an opening or closing order fill.
//
// This field is optional.
func (m *HistoricalOrderFillResponseFixed) SetOpenClose(value OpenCloseTradeEnum) *HistoricalOrderFillResponseFixed {
	m.p.SetInt32LE(256, int32(value))
	return m
}

// SetNoOrderFills Set to a numeric 1 to indicate there are no historical order fills.
//
// If there are no order fills to return, the Server needs to set this to
// 1 and send through 1 HistoricalOrderFillResponse message to indicate there
// are no order fills. Otherwise, leave this field at the default of 0.
func (m *HistoricalOrderFillResponseFixed) SetNoOrderFills(value uint8) *HistoricalOrderFillResponseFixed {
	m.p.SetUint8(260, value)
	return m
}

// SetInfoText
func (m *HistoricalOrderFillResponseFixed) SetInfoText(value string) *HistoricalOrderFillResponseFixed {
	m.p.SetStringFixed(261, 96, value)
	return m
}

// SetHighPriceDuringPosition
func (m *HistoricalOrderFillResponseFixed) SetHighPriceDuringPosition(value float64) *HistoricalOrderFillResponseFixed {
	m.p.SetFloat64LE(360, value)
	return m
}

// SetLowPriceDuringPosition
func (m *HistoricalOrderFillResponseFixed) SetLowPriceDuringPosition(value float64) *HistoricalOrderFillResponseFixed {
	m.p.SetFloat64LE(368, value)
	return m
}

// SetPositionQuantity
func (m *HistoricalOrderFillResponseFixed) SetPositionQuantity(value float64) *HistoricalOrderFillResponseFixed {
	m.p.SetFloat64LE(376, value)
	return m
}

func (m HistoricalOrderFillResponse) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m HistoricalOrderFillResponse) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m HistoricalOrderFillResponseFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m HistoricalOrderFillResponseFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m HistoricalOrderFillResponse) Copy(to HistoricalOrderFillResponse) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponse) CopyTo(to HistoricalOrderFillResponseFixed) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// Copy
func (m HistoricalOrderFillResponseFixed) Copy(to HistoricalOrderFillResponseFixed) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

// CopyTo
func (m HistoricalOrderFillResponseFixed) CopyTo(to HistoricalOrderFillResponse) {
	to.SetRequestID(m.RequestID())
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetBuySell(m.BuySell())
	to.SetPrice(m.Price())
	to.SetDateTime(m.DateTime())
	to.SetQuantity(m.Quantity())
	to.SetUniqueExecutionID(m.UniqueExecutionID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetOpenClose(m.OpenClose())
	to.SetNoOrderFills(m.NoOrderFills())
	to.SetInfoText(m.InfoText())
	to.SetHighPriceDuringPosition(m.HighPriceDuringPosition())
	to.SetLowPriceDuringPosition(m.LowPriceDuringPosition())
	to.SetPositionQuantity(m.PositionQuantity())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponse) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 304 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price":
			m.SetPrice(r.Float64())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UniqueExecutionID":
			m.SetUniqueExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "OpenClose":
			m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
		case "NoOrderFills":
			m.SetNoOrderFills(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
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

func (m *HistoricalOrderFillResponse) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalOrderFillResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 304)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price", m.Price())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.Float64Field("Quantity", m.Quantity())
	w.StringField("UniqueExecutionID", m.UniqueExecutionID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("OpenClose", int32(m.OpenClose()))
	w.Uint8Field("NoOrderFills", m.NoOrderFills())
	w.StringField("InfoText", m.InfoText())
	w.Float64Field("HighPriceDuringPosition", m.HighPriceDuringPosition())
	w.Float64Field("LowPriceDuringPosition", m.LowPriceDuringPosition())
	w.Float64Field("PositionQuantity", m.PositionQuantity())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalOrderFillResponseFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 304 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price":
			m.SetPrice(r.Float64())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "UniqueExecutionID":
			m.SetUniqueExecutionID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "OpenClose":
			m.SetOpenClose(OpenCloseTradeEnum(r.Int32()))
		case "NoOrderFills":
			m.SetNoOrderFills(r.Uint8())
		case "InfoText":
			m.SetInfoText(r.String())
		case "HighPriceDuringPosition":
			m.SetHighPriceDuringPosition(r.Float64())
		case "LowPriceDuringPosition":
			m.SetLowPriceDuringPosition(r.Float64())
		case "PositionQuantity":
			m.SetPositionQuantity(r.Float64())
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

func (m *HistoricalOrderFillResponseFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
