// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const SubmitNewSingleOrderSize = 104

const SubmitNewSingleOrderFixedSize = 296

//     Size              uint16              = SubmitNewSingleOrderSize  (104)
//     Type              uint16              = SUBMIT_NEW_SINGLE_ORDER  (208)
//     BaseSize          uint16              = SubmitNewSingleOrderSize  (104)
//     Symbol            string              = ""
//     Exchange          string              = ""
//     TradeAccount      string              = ""
//     ClientOrderID     string              = ""
//     OrderType         OrderTypeEnum       = ORDER_TYPE_UNSET  (0)
//     BuySell           BuySellEnum         = BUY_SELL_UNSET  (0)
//     Price1            float64             = 0.000000
//     Price2            float64             = 0.000000
//     Quantity          float64             = 0.000000
//     TimeInForce       TimeInForceEnum     = TIF_UNSET  (0)
//     GoodTillDateTime  DateTime            = 0
//     IsAutomatedOrder  bool                = false
//     IsParentOrder     bool                = false
//     FreeFormText      string              = ""
//     OpenOrClose       OpenCloseTradeEnum  = TRADE_UNSET  (0)
//     MaxShowQuantity   float64             = 0.000000
//     Price1AsString    string              = ""
//     Price2AsString    string              = ""
var _SubmitNewSingleOrderDefault = []byte{104, 0, 208, 0, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size              uint16              = SubmitNewSingleOrderFixedSize  (296)
//     Type              uint16              = SUBMIT_NEW_SINGLE_ORDER  (208)
//     Symbol            string[64]          = ""
//     Exchange          string[16]          = ""
//     TradeAccount      string[32]          = ""
//     ClientOrderID     string[32]          = ""
//     OrderType         OrderTypeEnum       = ORDER_TYPE_UNSET  (0)
//     BuySell           BuySellEnum         = BUY_SELL_UNSET  (0)
//     Price1            float64             = 0.000000
//     Price2            float64             = 0.000000
//     Quantity          float64             = 0.000000
//     TimeInForce       TimeInForceEnum     = TIF_UNSET  (0)
//     GoodTillDateTime  DateTime            = 0
//     IsAutomatedOrder  bool                = false
//     IsParentOrder     bool                = false
//     FreeFormText      string[48]          = ""
//     OpenOrClose       OpenCloseTradeEnum  = TRADE_UNSET  (0)
//     MaxShowQuantity   float64             = 0.000000
//     Price1AsString    string[16]          = ""
//     Price2AsString    string[16]          = ""
var _SubmitNewSingleOrderFixedDefault = []byte{40, 1, 208, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SubmitNewSingleOrder This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrder struct {
	p message.VLS
}

// SubmitNewSingleOrderFixed This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderFixed struct {
	p message.Fixed
}

func NewSubmitNewSingleOrderFrom(b []byte) SubmitNewSingleOrder {
	return SubmitNewSingleOrder{p: message.NewVLS(b)}
}

func WrapSubmitNewSingleOrder(b []byte) SubmitNewSingleOrder {
	return SubmitNewSingleOrder{p: message.WrapVLS(b)}
}

func NewSubmitNewSingleOrder() *SubmitNewSingleOrder {
	return &SubmitNewSingleOrder{p: message.NewVLS(_SubmitNewSingleOrderDefault)}
}

func ParseSubmitNewSingleOrder(b []byte) (SubmitNewSingleOrder, error) {
	if len(b) < 6 {
		return SubmitNewSingleOrder{}, message.ErrShortBuffer
	}
	m := WrapSubmitNewSingleOrder(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SubmitNewSingleOrder{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return SubmitNewSingleOrder{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 104 {
		newSize := len(b) + (104 - baseSize)
		if newSize > message.MaxSize {
			return SubmitNewSingleOrder{}, message.ErrOverflow
		}
		clone := SubmitNewSingleOrder{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _SubmitNewSingleOrderDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(104 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(6)
			if offset > 0 {
				clone.p.SetUint16LE(6, offset+shift)
			}
			offset = clone.p.Uint16LE(10)
			if offset > 0 {
				clone.p.SetUint16LE(10, offset+shift)
			}
			offset = clone.p.Uint16LE(14)
			if offset > 0 {
				clone.p.SetUint16LE(14, offset+shift)
			}
			offset = clone.p.Uint16LE(18)
			if offset > 0 {
				clone.p.SetUint16LE(18, offset+shift)
			}
			offset = clone.p.Uint16LE(74)
			if offset > 0 {
				clone.p.SetUint16LE(74, offset+shift)
			}
			offset = clone.p.Uint16LE(96)
			if offset > 0 {
				clone.p.SetUint16LE(96, offset+shift)
			}
			offset = clone.p.Uint16LE(100)
			if offset > 0 {
				clone.p.SetUint16LE(100, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewSubmitNewSingleOrderFixedFrom(b []byte) SubmitNewSingleOrderFixed {
	return SubmitNewSingleOrderFixed{p: message.NewFixed(b)}
}

func WrapSubmitNewSingleOrderFixed(b []byte) SubmitNewSingleOrderFixed {
	return SubmitNewSingleOrderFixed{p: message.WrapFixed(b)}
}

func NewSubmitNewSingleOrderFixed() *SubmitNewSingleOrderFixed {
	return &SubmitNewSingleOrderFixed{p: message.NewFixed(_SubmitNewSingleOrderFixedDefault)}
}

func ParseSubmitNewSingleOrderFixed(b []byte) (SubmitNewSingleOrderFixed, error) {
	if len(b) < 4 {
		return SubmitNewSingleOrderFixed{}, message.ErrShortBuffer
	}
	m := WrapSubmitNewSingleOrderFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SubmitNewSingleOrderFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return SubmitNewSingleOrderFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 296 {
		clone := *NewSubmitNewSingleOrderFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _SubmitNewSingleOrderFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m SubmitNewSingleOrder) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SubmitNewSingleOrder) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m SubmitNewSingleOrder) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrder) Symbol() string {
	return m.p.StringVLS(6)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrder) Exchange() string {
	return m.p.StringVLS(10)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrder) TradeAccount() string {
	return m.p.StringVLS(14)
}

// ClientOrderID This is the Client supplied order identifier. The Server will maintain
// this order identifier throughout the life of the order and always provide
// it back through the ClientOrderID field in the OrderUpdate messages for
// the order.
//
// This identifier cannot be an identifier used for a currently open order
// and it cannot be an identifier previously used in the current trading
// session. The trading session typically will be a 24-hour period defined
// by the Server. The Server shall reject an order with a client order identifier
// that is for a currently open order or which has already been used during
// the current trading session.
func (m SubmitNewSingleOrder) ClientOrderID() string {
	return m.p.StringVLS(18)
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(m.p.Int32LE(24))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrder) BuySell() BuySellEnum {
	return BuySellEnum(m.p.Int32LE(28))
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrder) Price1() float64 {
	return m.p.Float64LE(32)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrder) Price2() float64 {
	return m.p.Float64LE(40)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrder) Quantity() float64 {
	return m.p.Float64LE(48)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(m.p.Int32LE(56))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrder) GoodTillDateTime() DateTime {
	return DateTime(m.p.Int64LE(64))
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrder) IsAutomatedOrder() bool {
	return m.p.Bool(72)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrder) IsParentOrder() bool {
	return m.p.Bool(73)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrder) FreeFormText() string {
	return m.p.StringVLS(74)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(m.p.Int32LE(80))
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrder) MaxShowQuantity() float64 {
	return m.p.Float64LE(88)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrder) Price1AsString() string {
	return m.p.StringVLS(96)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrder) Price2AsString() string {
	return m.p.StringVLS(100)
}

// Size The standard message size field. Automatically set by constructor.
func (m SubmitNewSingleOrderFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SubmitNewSingleOrderFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderFixed) Symbol() string {
	return m.p.StringFixed(4, 64)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderFixed) Exchange() string {
	return m.p.StringFixed(68, 16)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderFixed) TradeAccount() string {
	return m.p.StringFixed(84, 32)
}

// ClientOrderID This is the Client supplied order identifier. The Server will maintain
// this order identifier throughout the life of the order and always provide
// it back through the ClientOrderID field in the OrderUpdate messages for
// the order.
//
// This identifier cannot be an identifier used for a currently open order
// and it cannot be an identifier previously used in the current trading
// session. The trading session typically will be a 24-hour period defined
// by the Server. The Server shall reject an order with a client order identifier
// that is for a currently open order or which has already been used during
// the current trading session.
func (m SubmitNewSingleOrderFixed) ClientOrderID() string {
	return m.p.StringFixed(116, 32)
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderFixed) OrderType() OrderTypeEnum {
	return OrderTypeEnum(m.p.Int32LE(148))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderFixed) BuySell() BuySellEnum {
	return BuySellEnum(m.p.Int32LE(152))
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderFixed) Price1() float64 {
	return m.p.Float64LE(160)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderFixed) Price2() float64 {
	return m.p.Float64LE(168)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderFixed) Quantity() float64 {
	return m.p.Float64LE(176)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderFixed) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(m.p.Int32LE(184))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderFixed) GoodTillDateTime() DateTime {
	return DateTime(m.p.Int64LE(192))
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderFixed) IsAutomatedOrder() bool {
	return m.p.Bool(200)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderFixed) IsParentOrder() bool {
	return m.p.Bool(201)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderFixed) FreeFormText() string {
	return m.p.StringFixed(202, 48)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderFixed) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(m.p.Int32LE(252))
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderFixed) MaxShowQuantity() float64 {
	return m.p.Float64LE(256)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderFixed) Price1AsString() string {
	return m.p.StringFixed(264, 16)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderFixed) Price2AsString() string {
	return m.p.StringFixed(280, 16)
}

// SetSymbol The symbol for the order.
func (m *SubmitNewSingleOrder) SetSymbol(value string) *SubmitNewSingleOrder {
	m.p.SetStringVLS(6, value)
	return m
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewSingleOrder) SetExchange(value string) *SubmitNewSingleOrder {
	m.p.SetStringVLS(10, value)
	return m
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m *SubmitNewSingleOrder) SetTradeAccount(value string) *SubmitNewSingleOrder {
	m.p.SetStringVLS(14, value)
	return m
}

// SetClientOrderID This is the Client supplied order identifier. The Server will maintain
// this order identifier throughout the life of the order and always provide
// it back through the ClientOrderID field in the OrderUpdate messages for
// the order.
//
// This identifier cannot be an identifier used for a currently open order
// and it cannot be an identifier previously used in the current trading
// session. The trading session typically will be a 24-hour period defined
// by the Server. The Server shall reject an order with a client order identifier
// that is for a currently open order or which has already been used during
// the current trading session.
func (m *SubmitNewSingleOrder) SetClientOrderID(value string) *SubmitNewSingleOrder {
	m.p.SetStringVLS(18, value)
	return m
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m *SubmitNewSingleOrder) SetOrderType(value OrderTypeEnum) *SubmitNewSingleOrder {
	m.p.SetInt32LE(24, int32(value))
	return m
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m *SubmitNewSingleOrder) SetBuySell(value BuySellEnum) *SubmitNewSingleOrder {
	m.p.SetInt32LE(28, int32(value))
	return m
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m *SubmitNewSingleOrder) SetPrice1(value float64) *SubmitNewSingleOrder {
	m.p.SetFloat64LE(32, value)
	return m
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m *SubmitNewSingleOrder) SetPrice2(value float64) *SubmitNewSingleOrder {
	m.p.SetFloat64LE(40, value)
	return m
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m *SubmitNewSingleOrder) SetQuantity(value float64) *SubmitNewSingleOrder {
	m.p.SetFloat64LE(48, value)
	return m
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m *SubmitNewSingleOrder) SetTimeInForce(value TimeInForceEnum) *SubmitNewSingleOrder {
	m.p.SetInt32LE(56, int32(value))
	return m
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m *SubmitNewSingleOrder) SetGoodTillDateTime(value DateTime) *SubmitNewSingleOrder {
	m.p.SetInt64LE(64, int64(value))
	return m
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m *SubmitNewSingleOrder) SetIsAutomatedOrder(value bool) *SubmitNewSingleOrder {
	m.p.SetBool(72, value)
	return m
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m *SubmitNewSingleOrder) SetIsParentOrder(value bool) *SubmitNewSingleOrder {
	m.p.SetBool(73, value)
	return m
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m *SubmitNewSingleOrder) SetFreeFormText(value string) *SubmitNewSingleOrder {
	m.p.SetStringVLS(74, value)
	return m
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m *SubmitNewSingleOrder) SetOpenOrClose(value OpenCloseTradeEnum) *SubmitNewSingleOrder {
	m.p.SetInt32LE(80, int32(value))
	return m
}

// SetMaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m *SubmitNewSingleOrder) SetMaxShowQuantity(value float64) *SubmitNewSingleOrder {
	m.p.SetFloat64LE(88, value)
	return m
}

// SetPrice1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m *SubmitNewSingleOrder) SetPrice1AsString(value string) *SubmitNewSingleOrder {
	m.p.SetStringVLS(96, value)
	return m
}

// SetPrice2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m *SubmitNewSingleOrder) SetPrice2AsString(value string) *SubmitNewSingleOrder {
	m.p.SetStringVLS(100, value)
	return m
}

// SetSymbol The symbol for the order.
func (m *SubmitNewSingleOrderFixed) SetSymbol(value string) *SubmitNewSingleOrderFixed {
	m.p.SetStringFixed(4, 64, value)
	return m
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewSingleOrderFixed) SetExchange(value string) *SubmitNewSingleOrderFixed {
	m.p.SetStringFixed(68, 16, value)
	return m
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m *SubmitNewSingleOrderFixed) SetTradeAccount(value string) *SubmitNewSingleOrderFixed {
	m.p.SetStringFixed(84, 32, value)
	return m
}

// SetClientOrderID This is the Client supplied order identifier. The Server will maintain
// this order identifier throughout the life of the order and always provide
// it back through the ClientOrderID field in the OrderUpdate messages for
// the order.
//
// This identifier cannot be an identifier used for a currently open order
// and it cannot be an identifier previously used in the current trading
// session. The trading session typically will be a 24-hour period defined
// by the Server. The Server shall reject an order with a client order identifier
// that is for a currently open order or which has already been used during
// the current trading session.
func (m *SubmitNewSingleOrderFixed) SetClientOrderID(value string) *SubmitNewSingleOrderFixed {
	m.p.SetStringFixed(116, 32, value)
	return m
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m *SubmitNewSingleOrderFixed) SetOrderType(value OrderTypeEnum) *SubmitNewSingleOrderFixed {
	m.p.SetInt32LE(148, int32(value))
	return m
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m *SubmitNewSingleOrderFixed) SetBuySell(value BuySellEnum) *SubmitNewSingleOrderFixed {
	m.p.SetInt32LE(152, int32(value))
	return m
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m *SubmitNewSingleOrderFixed) SetPrice1(value float64) *SubmitNewSingleOrderFixed {
	m.p.SetFloat64LE(160, value)
	return m
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m *SubmitNewSingleOrderFixed) SetPrice2(value float64) *SubmitNewSingleOrderFixed {
	m.p.SetFloat64LE(168, value)
	return m
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m *SubmitNewSingleOrderFixed) SetQuantity(value float64) *SubmitNewSingleOrderFixed {
	m.p.SetFloat64LE(176, value)
	return m
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m *SubmitNewSingleOrderFixed) SetTimeInForce(value TimeInForceEnum) *SubmitNewSingleOrderFixed {
	m.p.SetInt32LE(184, int32(value))
	return m
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m *SubmitNewSingleOrderFixed) SetGoodTillDateTime(value DateTime) *SubmitNewSingleOrderFixed {
	m.p.SetInt64LE(192, int64(value))
	return m
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m *SubmitNewSingleOrderFixed) SetIsAutomatedOrder(value bool) *SubmitNewSingleOrderFixed {
	m.p.SetBool(200, value)
	return m
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m *SubmitNewSingleOrderFixed) SetIsParentOrder(value bool) *SubmitNewSingleOrderFixed {
	m.p.SetBool(201, value)
	return m
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m *SubmitNewSingleOrderFixed) SetFreeFormText(value string) *SubmitNewSingleOrderFixed {
	m.p.SetStringFixed(202, 48, value)
	return m
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m *SubmitNewSingleOrderFixed) SetOpenOrClose(value OpenCloseTradeEnum) *SubmitNewSingleOrderFixed {
	m.p.SetInt32LE(252, int32(value))
	return m
}

// SetMaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m *SubmitNewSingleOrderFixed) SetMaxShowQuantity(value float64) *SubmitNewSingleOrderFixed {
	m.p.SetFloat64LE(256, value)
	return m
}

// SetPrice1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m *SubmitNewSingleOrderFixed) SetPrice1AsString(value string) *SubmitNewSingleOrderFixed {
	m.p.SetStringFixed(264, 16, value)
	return m
}

// SetPrice2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m *SubmitNewSingleOrderFixed) SetPrice2AsString(value string) *SubmitNewSingleOrderFixed {
	m.p.SetStringFixed(280, 16, value)
	return m
}

func (m *SubmitNewSingleOrder) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *SubmitNewSingleOrder) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *SubmitNewSingleOrder) Clone() *SubmitNewSingleOrder {
	return &SubmitNewSingleOrder{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *SubmitNewSingleOrderFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *SubmitNewSingleOrderFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *SubmitNewSingleOrderFixed) Clone() *SubmitNewSingleOrderFixed {
	return &SubmitNewSingleOrderFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m SubmitNewSingleOrder) Copy(to SubmitNewSingleOrder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetMaxShowQuantity(m.MaxShowQuantity())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyTo
func (m SubmitNewSingleOrder) CopyTo(to SubmitNewSingleOrderFixed) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetMaxShowQuantity(m.MaxShowQuantity())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// Copy
func (m SubmitNewSingleOrderFixed) Copy(to SubmitNewSingleOrderFixed) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetMaxShowQuantity(m.MaxShowQuantity())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

// CopyTo
func (m SubmitNewSingleOrderFixed) CopyTo(to SubmitNewSingleOrder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetMaxShowQuantity(m.MaxShowQuantity())
	to.SetPrice1AsString(m.Price1AsString())
	to.SetPrice2AsString(m.Price2AsString())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrder) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *SubmitNewSingleOrder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrder) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *SubmitNewSingleOrder) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 208 {
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
		case "OrderType":
			m.SetOrderType(OrderTypeEnum(r.Int32()))
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price1":
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
		case "IsParentOrder":
			m.SetIsParentOrder(r.Bool())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "OpenOrClose":
			m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
		case "MaxShowQuantity":
			m.SetMaxShowQuantity(r.Float64())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *SubmitNewSingleOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 208)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.Int32Field("OrderType", int32(m.OrderType()))
	w.Int32Field("BuySell", int32(m.BuySell()))
	w.Float64Field("Price1", m.Price1())
	w.Float64Field("Price2", m.Price2())
	w.Float64Field("Quantity", m.Quantity())
	w.Int32Field("TimeInForce", int32(m.TimeInForce()))
	w.Int64Field("GoodTillDateTime", int64(m.GoodTillDateTime()))
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	w.BoolField("IsParentOrder", m.IsParentOrder())
	w.StringField("FreeFormText", m.FreeFormText())
	w.Int32Field("OpenOrClose", int32(m.OpenOrClose()))
	w.Float64Field("MaxShowQuantity", m.MaxShowQuantity())
	w.StringField("Price1AsString", m.Price1AsString())
	w.StringField("Price2AsString", m.Price2AsString())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitNewSingleOrderFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *SubmitNewSingleOrderFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 208 {
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
		case "OrderType":
			m.SetOrderType(OrderTypeEnum(r.Int32()))
		case "BuySell":
			m.SetBuySell(BuySellEnum(r.Int32()))
		case "Price1":
			m.SetPrice1(r.Float64())
		case "Price2":
			m.SetPrice2(r.Float64())
		case "Quantity":
			m.SetQuantity(r.Float64())
		case "TimeInForce":
			m.SetTimeInForce(TimeInForceEnum(r.Int32()))
		case "GoodTillDateTime":
			m.SetGoodTillDateTime(DateTime(r.Int64()))
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
		case "IsParentOrder":
			m.SetIsParentOrder(r.Bool())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "OpenOrClose":
			m.SetOpenOrClose(OpenCloseTradeEnum(r.Int32()))
		case "MaxShowQuantity":
			m.SetMaxShowQuantity(r.Float64())
		case "Price1AsString":
			m.SetPrice1AsString(r.String())
		case "Price2AsString":
			m.SetPrice2AsString(r.String())
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
