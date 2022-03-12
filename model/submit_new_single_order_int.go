package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SubmitNewSingleOrderIntSize = 96

const SubmitNewSingleOrderIntFixedSize = 264

// {
//     Size             = SubmitNewSingleOrderIntSize  (96)
//     Type             = SUBMIT_NEW_SINGLE_ORDER_INT  (206)
//     BaseSize         = SubmitNewSingleOrderIntSize  (96)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Divisor          = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
// }
var _SubmitNewSingleOrderIntDefault = []byte{96, 0, 206, 0, 96, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size             = SubmitNewSingleOrderIntFixedSize  (264)
//     Type             = SUBMIT_NEW_SINGLE_ORDER_INT  (206)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Divisor          = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
// }
var _SubmitNewSingleOrderIntFixedDefault = []byte{8, 1, 206, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SubmitNewSingleOrderInt The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderInt struct {
	message.VLS
}

// SubmitNewSingleOrderIntBuilder The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderIntBuilder struct {
	message.VLSBuilder
}

// SubmitNewSingleOrderIntFixed The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderIntFixed struct {
	message.Fixed
}

// SubmitNewSingleOrderIntFixedBuilder The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderIntFixedBuilder struct {
	message.Fixed
}

// SubmitNewSingleOrderIntPointer The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderIntPointer struct {
	message.VLSPointer
}

// SubmitNewSingleOrderIntPointerBuilder The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderIntPointerBuilder struct {
	message.VLSPointerBuilder
}

// SubmitNewSingleOrderIntFixedPointer The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderIntFixedPointer struct {
	message.FixedPointer
}

// SubmitNewSingleOrderIntFixedPointerBuilder The SubmitNewSingleOrder_INT message is identical to SubmitNewSingleOrder
// except that the order prices are as integers.
//
// The Client will only send this message to the Server if UseIntegerPriceOrderMessages
// is set in the LogonResponse message.
//
// When setting the Price1 and Price2 fields, multiply the order price by
// the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded.
//
// This message also contains the Divisor field. This is the FloatToIntPriceMultiplier
// value provided in the SecurityDefinitionResponse message for the symbol
// being traded. The Server needs to divide the Price1 and Price2 fields
// by this Divisor to get the prices as float values.
type SubmitNewSingleOrderIntFixedPointerBuilder struct {
	message.FixedPointer
}

func NewSubmitNewSingleOrderIntFrom(b []byte) SubmitNewSingleOrderInt {
	return SubmitNewSingleOrderInt{VLS: message.NewVLSFrom(b)}
}

func WrapSubmitNewSingleOrderInt(b []byte) SubmitNewSingleOrderInt {
	return SubmitNewSingleOrderInt{VLS: message.WrapVLS(b)}
}

func NewSubmitNewSingleOrderInt() SubmitNewSingleOrderIntBuilder {
	a := SubmitNewSingleOrderIntBuilder{message.NewVLSBuilder(96)}
	a.Unsafe().SetBytes(0, _SubmitNewSingleOrderIntDefault)
	return a
}

func NewSubmitNewSingleOrderIntFixedFrom(b []byte) SubmitNewSingleOrderIntFixed {
	return SubmitNewSingleOrderIntFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSubmitNewSingleOrderIntFixed(b []byte) SubmitNewSingleOrderIntFixed {
	return SubmitNewSingleOrderIntFixed{Fixed: message.WrapFixed(b)}
}

func NewSubmitNewSingleOrderIntFixed() SubmitNewSingleOrderIntFixedBuilder {
	a := SubmitNewSingleOrderIntFixedBuilder{message.NewFixed(264)}
	a.Unsafe().SetBytes(0, _SubmitNewSingleOrderIntFixedDefault)
	return a
}

func AllocSubmitNewSingleOrderInt() SubmitNewSingleOrderIntPointerBuilder {
	a := SubmitNewSingleOrderIntPointerBuilder{message.AllocVLSBuilder(96)}
	a.Ptr.SetBytes(0, _SubmitNewSingleOrderIntDefault)
	return a
}

func AllocSubmitNewSingleOrderIntFrom(b []byte) SubmitNewSingleOrderIntPointer {
	return SubmitNewSingleOrderIntPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocSubmitNewSingleOrderIntFixed() SubmitNewSingleOrderIntFixedPointerBuilder {
	a := SubmitNewSingleOrderIntFixedPointerBuilder{message.AllocFixed(264)}
	a.Ptr.SetBytes(0, _SubmitNewSingleOrderIntFixedDefault)
	return a
}

func AllocSubmitNewSingleOrderIntFixedFrom(b []byte) SubmitNewSingleOrderIntFixedPointer {
	return SubmitNewSingleOrderIntFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size             = SubmitNewSingleOrderIntSize  (96)
//     Type             = SUBMIT_NEW_SINGLE_ORDER_INT  (206)
//     BaseSize         = SubmitNewSingleOrderIntSize  (96)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Divisor          = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
// }
func (m SubmitNewSingleOrderIntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewSingleOrderIntDefault)
}

// Clear
// {
//     Size             = SubmitNewSingleOrderIntFixedSize  (264)
//     Type             = SUBMIT_NEW_SINGLE_ORDER_INT  (206)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Divisor          = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
// }
func (m SubmitNewSingleOrderIntFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewSingleOrderIntFixedDefault)
}

// Clear
// {
//     Size             = SubmitNewSingleOrderIntSize  (96)
//     Type             = SUBMIT_NEW_SINGLE_ORDER_INT  (206)
//     BaseSize         = SubmitNewSingleOrderIntSize  (96)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Divisor          = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
// }
func (m SubmitNewSingleOrderIntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewSingleOrderIntDefault)
}

// Clear
// {
//     Size             = SubmitNewSingleOrderIntFixedSize  (264)
//     Type             = SUBMIT_NEW_SINGLE_ORDER_INT  (206)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Divisor          = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
// }
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewSingleOrderIntFixedDefault)
}

// ToBuilder
func (m SubmitNewSingleOrderInt) ToBuilder() SubmitNewSingleOrderIntBuilder {
	return SubmitNewSingleOrderIntBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SubmitNewSingleOrderIntFixed) ToBuilder() SubmitNewSingleOrderIntFixedBuilder {
	return SubmitNewSingleOrderIntFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SubmitNewSingleOrderIntPointer) ToBuilder() SubmitNewSingleOrderIntPointerBuilder {
	return SubmitNewSingleOrderIntPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SubmitNewSingleOrderIntFixedPointer) ToBuilder() SubmitNewSingleOrderIntFixedPointerBuilder {
	return SubmitNewSingleOrderIntFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SubmitNewSingleOrderIntBuilder) Finish() SubmitNewSingleOrderInt {
	return SubmitNewSingleOrderInt{m.VLSBuilder.Finish()}
}

// Finish
func (m SubmitNewSingleOrderIntFixedBuilder) Finish() SubmitNewSingleOrderIntFixed {
	return SubmitNewSingleOrderIntFixed{m.Fixed}
}

// Finish
func (m *SubmitNewSingleOrderIntPointerBuilder) Finish() SubmitNewSingleOrderIntPointer {
	return SubmitNewSingleOrderIntPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SubmitNewSingleOrderIntFixedPointerBuilder) Finish() SubmitNewSingleOrderIntFixedPointer {
	return SubmitNewSingleOrderIntFixedPointer{m.FixedPointer}
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderInt) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderIntBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderIntPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderIntPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderInt) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderInt) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
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
func (m SubmitNewSingleOrderInt) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
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
func (m SubmitNewSingleOrderIntBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 22, 18)
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
func (m SubmitNewSingleOrderIntPointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
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
func (m SubmitNewSingleOrderIntPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderInt) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntPointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntPointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderInt) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 32, 28))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 32, 28))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 32, 28))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 32, 28))
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderInt) Price1() int64 {
	return message.Int64VLS(m.Unsafe(), 40, 32)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntBuilder) Price1() int64 {
	return message.Int64VLS(m.Unsafe(), 40, 32)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntPointer) Price1() int64 {
	return message.Int64VLS(m.Ptr, 40, 32)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntPointerBuilder) Price1() int64 {
	return message.Int64VLS(m.Ptr, 40, 32)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderInt) Price2() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntBuilder) Price2() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntPointer) Price2() int64 {
	return message.Int64VLS(m.Ptr, 48, 40)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntPointerBuilder) Price2() int64 {
	return message.Int64VLS(m.Ptr, 48, 40)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderInt) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 52, 48)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntBuilder) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 52, 48)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntPointer) Divisor() float32 {
	return message.Float32VLS(m.Ptr, 52, 48)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntPointerBuilder) Divisor() float32 {
	return message.Float32VLS(m.Ptr, 52, 48)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderInt) Quantity() int64 {
	return message.Int64VLS(m.Unsafe(), 64, 56)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntBuilder) Quantity() int64 {
	return message.Int64VLS(m.Unsafe(), 64, 56)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntPointer) Quantity() int64 {
	return message.Int64VLS(m.Ptr, 64, 56)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntPointerBuilder) Quantity() int64 {
	return message.Int64VLS(m.Ptr, 64, 56)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderInt) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 68, 64))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 68, 64))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 68, 64))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 68, 64))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderInt) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 80, 72))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 80, 72))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 80, 72))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 80, 72))
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderInt) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 81, 80)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 81, 80)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntPointer) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 81, 80)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 81, 80)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderInt) IsParentOrder() bool {
	return message.BoolVLS(m.Unsafe(), 82, 81)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntBuilder) IsParentOrder() bool {
	return message.BoolVLS(m.Unsafe(), 82, 81)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntPointer) IsParentOrder() bool {
	return message.BoolVLS(m.Ptr, 82, 81)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntPointerBuilder) IsParentOrder() bool {
	return message.BoolVLS(m.Ptr, 82, 81)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderInt) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 86, 82)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 86, 82)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 86, 82)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 86, 82)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderInt) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 92, 88))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 92, 88))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 92, 88))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 92, 88))
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderIntFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderIntFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderIntFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
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
func (m SubmitNewSingleOrderIntFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
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
func (m SubmitNewSingleOrderIntFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 148, 116)
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
func (m SubmitNewSingleOrderIntFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
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
func (m SubmitNewSingleOrderIntFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixed) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixedBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixedPointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntFixed) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 156, 152))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntFixedBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 156, 152))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntFixedPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 156, 152))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 156, 152))
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntFixed) Price1() int64 {
	return message.Int64Fixed(m.Unsafe(), 168, 160)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntFixedBuilder) Price1() int64 {
	return message.Int64Fixed(m.Unsafe(), 168, 160)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntFixedPointer) Price1() int64 {
	return message.Int64Fixed(m.Ptr, 168, 160)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Price1() int64 {
	return message.Int64Fixed(m.Ptr, 168, 160)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntFixed) Price2() int64 {
	return message.Int64Fixed(m.Unsafe(), 176, 168)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntFixedBuilder) Price2() int64 {
	return message.Int64Fixed(m.Unsafe(), 176, 168)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntFixedPointer) Price2() int64 {
	return message.Int64Fixed(m.Ptr, 176, 168)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Price2() int64 {
	return message.Int64Fixed(m.Ptr, 176, 168)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntFixed) Divisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 180, 176)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntFixedBuilder) Divisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 180, 176)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntFixedPointer) Divisor() float32 {
	return message.Float32Fixed(m.Ptr, 180, 176)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Divisor() float32 {
	return message.Float32Fixed(m.Ptr, 180, 176)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntFixed) Quantity() int64 {
	return message.Int64Fixed(m.Unsafe(), 192, 184)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntFixedBuilder) Quantity() int64 {
	return message.Int64Fixed(m.Unsafe(), 192, 184)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntFixedPointer) Quantity() int64 {
	return message.Int64Fixed(m.Ptr, 192, 184)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Quantity() int64 {
	return message.Int64Fixed(m.Ptr, 192, 184)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntFixed) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 196, 192))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntFixedBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 196, 192))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 196, 192))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 196, 192))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntFixed) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 208, 200))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntFixedBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 208, 200))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 208, 200))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 208, 200))
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixed) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 209, 208)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixedBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 209, 208)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixedPointer) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 209, 208)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 209, 208)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixed) IsParentOrder() bool {
	return message.BoolFixed(m.Unsafe(), 210, 209)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedBuilder) IsParentOrder() bool {
	return message.BoolFixed(m.Unsafe(), 210, 209)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedPointer) IsParentOrder() bool {
	return message.BoolFixed(m.Ptr, 210, 209)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) IsParentOrder() bool {
	return message.BoolFixed(m.Ptr, 210, 209)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 258, 210)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 258, 210)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 258, 210)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 258, 210)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixed) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 264, 260))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixedBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 264, 260))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixedPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 264, 260))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 264, 260))
}

// SetSymbol The symbol for the order.
func (m *SubmitNewSingleOrderIntBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetSymbol The symbol for the order.
func (m *SubmitNewSingleOrderIntPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewSingleOrderIntBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewSingleOrderIntPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m *SubmitNewSingleOrderIntBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m *SubmitNewSingleOrderIntPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
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
func (m *SubmitNewSingleOrderIntBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 22, 18, value)
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
func (m *SubmitNewSingleOrderIntPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntPointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 32, 28, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 32, 28, int32(value))
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntBuilder) SetPrice1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 40, 32, value)
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntPointerBuilder) SetPrice1(value int64) {
	message.SetInt64VLS(m.Ptr, 40, 32, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntBuilder) SetPrice2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 48, 40, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntPointerBuilder) SetPrice2(value int64) {
	message.SetInt64VLS(m.Ptr, 48, 40, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 52, 48, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Ptr, 52, 48, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntBuilder) SetQuantity(value int64) {
	message.SetInt64VLS(m.Unsafe(), 64, 56, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntPointerBuilder) SetQuantity(value int64) {
	message.SetInt64VLS(m.Ptr, 64, 56, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 68, 64, int32(value))
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Ptr, 68, 64, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 80, 72, int64(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 80, 72, int64(value))
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 81, 80, value)
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 81, 80, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntBuilder) SetIsParentOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 82, 81, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntPointerBuilder) SetIsParentOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 82, 81, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m *SubmitNewSingleOrderIntBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 86, 82, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m *SubmitNewSingleOrderIntPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 86, 82, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 92, 88, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Ptr, 92, 88, int32(value))
}

// SetSymbol The symbol for the order.
func (m SubmitNewSingleOrderIntFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetSymbol The symbol for the order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
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
func (m SubmitNewSingleOrderIntFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 148, 116, value)
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
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 148, 116, value)
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixedBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 152, 148, int32(value))
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 152, 148, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntFixedBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 156, 152, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 156, 152, int32(value))
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntFixedBuilder) SetPrice1(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 168, 160, value)
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetPrice1(value int64) {
	message.SetInt64Fixed(m.Ptr, 168, 160, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntFixedBuilder) SetPrice2(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 176, 168, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetPrice2(value int64) {
	message.SetInt64Fixed(m.Ptr, 176, 168, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntFixedBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 180, 176, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 180, 176, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntFixedBuilder) SetQuantity(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 192, 184, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetQuantity(value int64) {
	message.SetInt64Fixed(m.Ptr, 192, 184, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntFixedBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Unsafe(), 196, 192, int32(value))
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 196, 192, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntFixedBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 208, 200, int64(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 208, 200, int64(value))
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixedBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Unsafe(), 209, 208, value)
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Ptr, 209, 208, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedBuilder) SetIsParentOrder(value bool) {
	message.SetBoolFixed(m.Unsafe(), 210, 209, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetIsParentOrder(value bool) {
	message.SetBoolFixed(m.Ptr, 210, 209, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 258, 210, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 258, 210, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixedBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 264, 260, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 264, 260, int32(value))
}

// Copy
func (m SubmitNewSingleOrderInt) Copy(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// Copy
func (m SubmitNewSingleOrderIntBuilder) Copy(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderInt) CopyPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderIntBuilder) CopyPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderInt) CopyToPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderIntBuilder) CopyToPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderInt) CopyTo(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderIntBuilder) CopyTo(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// Copy
func (m SubmitNewSingleOrderIntFixed) Copy(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// Copy
func (m SubmitNewSingleOrderIntFixedBuilder) Copy(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderIntFixed) CopyPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderIntFixedBuilder) CopyPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderIntFixed) CopyToPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderIntFixedBuilder) CopyToPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderIntFixed) CopyTo(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderIntFixedBuilder) CopyTo(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// Copy
func (m SubmitNewSingleOrderIntPointer) Copy(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// Copy
func (m SubmitNewSingleOrderIntPointerBuilder) Copy(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderIntPointer) CopyPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderIntPointerBuilder) CopyPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderIntPointer) CopyTo(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderIntPointerBuilder) CopyTo(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderIntPointer) CopyToPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderIntPointerBuilder) CopyToPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// Copy
func (m SubmitNewSingleOrderIntFixedPointer) Copy(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// Copy
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Copy(to SubmitNewSingleOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderIntFixedPointer) CopyPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyPointer
func (m SubmitNewSingleOrderIntFixedPointerBuilder) CopyPointer(to SubmitNewSingleOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderIntFixedPointer) CopyTo(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyTo
func (m SubmitNewSingleOrderIntFixedPointerBuilder) CopyTo(to SubmitNewSingleOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderIntFixedPointer) CopyToPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}

// CopyToPointer
func (m SubmitNewSingleOrderIntFixedPointerBuilder) CopyToPointer(to SubmitNewSingleOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetOrderType(m.OrderType())
	to.SetBuySell(m.BuySell())
	to.SetPrice1(m.Price1())
	to.SetPrice2(m.Price2())
	to.SetDivisor(m.Divisor())
	to.SetQuantity(m.Quantity())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetIsParentOrder(m.IsParentOrder())
	to.SetFreeFormText(m.FreeFormText())
	to.SetOpenOrClose(m.OpenOrClose())
}
