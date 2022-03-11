package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SubmitNewSingleOrderSize = 104

const SubmitNewSingleOrderFixedSize = 296

// {
//     Size             = SubmitNewSingleOrderSize  (104)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     BaseSize         = SubmitNewSingleOrderSize  (104)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
var _SubmitNewSingleOrderDefault = []byte{104, 0, 208, 0, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size             = SubmitNewSingleOrderFixedSize  (296)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
var _SubmitNewSingleOrderFixedDefault = []byte{40, 1, 208, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SubmitNewSingleOrder This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrder struct {
	message.VLS
}

// SubmitNewSingleOrderBuilder This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderBuilder struct {
	message.VLSBuilder
}

// SubmitNewSingleOrderFixed This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderFixed struct {
	message.Fixed
}

// SubmitNewSingleOrderFixedBuilder This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderFixedBuilder struct {
	message.Fixed
}

// SubmitNewSingleOrderPointer This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderPointer struct {
	message.VLSPointer
}

// SubmitNewSingleOrderPointerBuilder This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderPointerBuilder struct {
	message.VLSPointerBuilder
}

// SubmitNewSingleOrderFixedPointer This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderFixedPointer struct {
	message.FixedPointer
}

// SubmitNewSingleOrderFixedPointerBuilder This message is used to submit a new single order into the market from
// the Client to the Server.
type SubmitNewSingleOrderFixedPointerBuilder struct {
	message.FixedPointer
}

func NewSubmitNewSingleOrderFrom(b []byte) SubmitNewSingleOrder {
	return SubmitNewSingleOrder{VLS: message.NewVLSFrom(b)}
}

func WrapSubmitNewSingleOrder(b []byte) SubmitNewSingleOrder {
	return SubmitNewSingleOrder{VLS: message.WrapVLS(b)}
}

func NewSubmitNewSingleOrder() SubmitNewSingleOrderBuilder {
	a := SubmitNewSingleOrderBuilder{message.NewVLSBuilder(104)}
	a.Unsafe().SetBytes(0, _SubmitNewSingleOrderDefault)
	return a
}

func NewSubmitNewSingleOrderFixedFrom(b []byte) SubmitNewSingleOrderFixed {
	return SubmitNewSingleOrderFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSubmitNewSingleOrderFixed(b []byte) SubmitNewSingleOrderFixed {
	return SubmitNewSingleOrderFixed{Fixed: message.WrapFixed(b)}
}

func NewSubmitNewSingleOrderFixed() SubmitNewSingleOrderFixedBuilder {
	a := SubmitNewSingleOrderFixedBuilder{message.NewFixed(296)}
	a.Unsafe().SetBytes(0, _SubmitNewSingleOrderFixedDefault)
	return a
}

func AllocSubmitNewSingleOrder() SubmitNewSingleOrderPointerBuilder {
	a := SubmitNewSingleOrderPointerBuilder{message.AllocVLSBuilder(104)}
	a.Ptr.SetBytes(0, _SubmitNewSingleOrderDefault)
	return a
}

func AllocSubmitNewSingleOrderFrom(b []byte) SubmitNewSingleOrderPointer {
	return SubmitNewSingleOrderPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocSubmitNewSingleOrderFixed() SubmitNewSingleOrderFixedPointerBuilder {
	a := SubmitNewSingleOrderFixedPointerBuilder{message.AllocFixed(296)}
	a.Ptr.SetBytes(0, _SubmitNewSingleOrderFixedDefault)
	return a
}

func AllocSubmitNewSingleOrderFixedFrom(b []byte) SubmitNewSingleOrderFixedPointer {
	return SubmitNewSingleOrderFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size             = SubmitNewSingleOrderSize  (104)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     BaseSize         = SubmitNewSingleOrderSize  (104)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
func (m SubmitNewSingleOrderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewSingleOrderDefault)
}

// Clear
// {
//     Size             = SubmitNewSingleOrderFixedSize  (296)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
func (m SubmitNewSingleOrderFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewSingleOrderFixedDefault)
}

// Clear
// {
//     Size             = SubmitNewSingleOrderSize  (104)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     BaseSize         = SubmitNewSingleOrderSize  (104)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
func (m SubmitNewSingleOrderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewSingleOrderDefault)
}

// Clear
// {
//     Size             = SubmitNewSingleOrderFixedSize  (296)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = ORDER_TYPE_UNSET  (0)
//     BuySell          = BUY_SELL_UNSET  (0)
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = TIF_UNSET  (0)
//     GoodTillDateTime = 0
//     IsAutomatedOrder = false
//     IsParentOrder    = false
//     FreeFormText     = ""
//     OpenOrClose      = TRADE_UNSET  (0)
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
func (m SubmitNewSingleOrderFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewSingleOrderFixedDefault)
}

// ToBuilder
func (m SubmitNewSingleOrder) ToBuilder() SubmitNewSingleOrderBuilder {
	return SubmitNewSingleOrderBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SubmitNewSingleOrderFixed) ToBuilder() SubmitNewSingleOrderFixedBuilder {
	return SubmitNewSingleOrderFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SubmitNewSingleOrderPointer) ToBuilder() SubmitNewSingleOrderPointerBuilder {
	return SubmitNewSingleOrderPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SubmitNewSingleOrderFixedPointer) ToBuilder() SubmitNewSingleOrderFixedPointerBuilder {
	return SubmitNewSingleOrderFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SubmitNewSingleOrderBuilder) Finish() SubmitNewSingleOrder {
	return SubmitNewSingleOrder{m.VLSBuilder.Finish()}
}

// Finish
func (m SubmitNewSingleOrderFixedBuilder) Finish() SubmitNewSingleOrderFixed {
	return SubmitNewSingleOrderFixed{m.Fixed}
}

// Finish
func (m *SubmitNewSingleOrderPointerBuilder) Finish() SubmitNewSingleOrderPointer {
	return SubmitNewSingleOrderPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SubmitNewSingleOrderFixedPointerBuilder) Finish() SubmitNewSingleOrderFixedPointer {
	return SubmitNewSingleOrderFixedPointer{m.FixedPointer}
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderPointerBuilder) TradeAccount() string {
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
func (m SubmitNewSingleOrder) ClientOrderID() string {
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
func (m SubmitNewSingleOrderBuilder) ClientOrderID() string {
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
func (m SubmitNewSingleOrderPointer) ClientOrderID() string {
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
func (m SubmitNewSingleOrderPointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 22, 18)
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderPointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderPointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 32, 28))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 32, 28))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 32, 28))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 32, 28))
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderBuilder) Price1() float64 {
	return message.Float64VLS(m.Unsafe(), 40, 32)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderPointer) Price1() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderPointerBuilder) Price1() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderBuilder) Price2() float64 {
	return message.Float64VLS(m.Unsafe(), 48, 40)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderPointer) Price2() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderPointerBuilder) Price2() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 56, 48)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderBuilder) Quantity() float64 {
	return message.Float64VLS(m.Unsafe(), 56, 48)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderPointer) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 56, 48)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderPointerBuilder) Quantity() float64 {
	return message.Float64VLS(m.Ptr, 56, 48)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 60, 56))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 60, 56))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 60, 56))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 60, 56))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 72, 64))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 72, 64))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 72, 64))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 72, 64))
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 73, 72)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 73, 72)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderPointer) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 73, 72)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 73, 72)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrder) IsParentOrder() bool {
	return message.BoolVLS(m.Unsafe(), 74, 73)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderBuilder) IsParentOrder() bool {
	return message.BoolVLS(m.Unsafe(), 74, 73)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderPointer) IsParentOrder() bool {
	return message.BoolVLS(m.Ptr, 74, 73)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderPointerBuilder) IsParentOrder() bool {
	return message.BoolVLS(m.Ptr, 74, 73)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 78, 74)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 78, 74)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 78, 74)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 78, 74)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 84, 80))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 84, 80))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 84, 80))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 84, 80))
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrder) MaxShowQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderBuilder) MaxShowQuantity() float64 {
	return message.Float64VLS(m.Unsafe(), 96, 88)
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderPointer) MaxShowQuantity() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderPointerBuilder) MaxShowQuantity() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrder) Price1AsString() string {
	return message.StringVLS(m.Unsafe(), 100, 96)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderBuilder) Price1AsString() string {
	return message.StringVLS(m.Unsafe(), 100, 96)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderPointer) Price1AsString() string {
	return message.StringVLS(m.Ptr, 100, 96)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderPointerBuilder) Price1AsString() string {
	return message.StringVLS(m.Ptr, 100, 96)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrder) Price2AsString() string {
	return message.StringVLS(m.Unsafe(), 104, 100)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderBuilder) Price2AsString() string {
	return message.StringVLS(m.Unsafe(), 104, 100)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderPointer) Price2AsString() string {
	return message.StringVLS(m.Ptr, 104, 100)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderPointerBuilder) Price2AsString() string {
	return message.StringVLS(m.Ptr, 104, 100)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderFixedPointerBuilder) TradeAccount() string {
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
func (m SubmitNewSingleOrderFixed) ClientOrderID() string {
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
func (m SubmitNewSingleOrderFixedBuilder) ClientOrderID() string {
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
func (m SubmitNewSingleOrderFixedPointer) ClientOrderID() string {
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
func (m SubmitNewSingleOrderFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 148, 116)
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderFixed) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderFixedBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 152, 148))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderFixedPointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderFixedPointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderFixed) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 156, 152))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderFixedBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 156, 152))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderFixedPointer) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 156, 152))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderFixedPointerBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 156, 152))
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderFixed) Price1() float64 {
	return message.Float64Fixed(m.Unsafe(), 168, 160)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderFixedBuilder) Price1() float64 {
	return message.Float64Fixed(m.Unsafe(), 168, 160)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderFixedPointer) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 168, 160)
}

// Price1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderFixedPointerBuilder) Price1() float64 {
	return message.Float64Fixed(m.Ptr, 168, 160)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderFixed) Price2() float64 {
	return message.Float64Fixed(m.Unsafe(), 176, 168)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderFixedBuilder) Price2() float64 {
	return message.Float64Fixed(m.Unsafe(), 176, 168)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderFixedPointer) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 176, 168)
}

// Price2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderFixedPointerBuilder) Price2() float64 {
	return message.Float64Fixed(m.Ptr, 176, 168)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderFixed) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 184, 176)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderFixedBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 184, 176)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderFixedPointer) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 184, 176)
}

// Quantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderFixedPointerBuilder) Quantity() float64 {
	return message.Float64Fixed(m.Ptr, 184, 176)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderFixed) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 188, 184))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderFixedBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 188, 184))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 188, 184))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 188, 184))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderFixed) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 200, 192))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderFixedBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 200, 192))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 200, 192))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 200, 192))
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderFixed) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 201, 200)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderFixedBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 201, 200)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderFixedPointer) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 201, 200)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderFixedPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 201, 200)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderFixed) IsParentOrder() bool {
	return message.BoolFixed(m.Unsafe(), 202, 201)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderFixedBuilder) IsParentOrder() bool {
	return message.BoolFixed(m.Unsafe(), 202, 201)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderFixedPointer) IsParentOrder() bool {
	return message.BoolFixed(m.Ptr, 202, 201)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderFixedPointerBuilder) IsParentOrder() bool {
	return message.BoolFixed(m.Ptr, 202, 201)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 250, 202)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 250, 202)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 250, 202)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 250, 202)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderFixed) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 256, 252))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderFixedBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 256, 252))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderFixedPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 256, 252))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderFixedPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 256, 252))
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderFixed) MaxShowQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 264, 256)
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderFixedBuilder) MaxShowQuantity() float64 {
	return message.Float64Fixed(m.Unsafe(), 264, 256)
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderFixedPointer) MaxShowQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 264, 256)
}

// MaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderFixedPointerBuilder) MaxShowQuantity() float64 {
	return message.Float64Fixed(m.Ptr, 264, 256)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderFixed) Price1AsString() string {
	return message.StringFixed(m.Unsafe(), 280, 264)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderFixedBuilder) Price1AsString() string {
	return message.StringFixed(m.Unsafe(), 280, 264)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderFixedPointer) Price1AsString() string {
	return message.StringFixed(m.Ptr, 280, 264)
}

// Price1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderFixedPointerBuilder) Price1AsString() string {
	return message.StringFixed(m.Ptr, 280, 264)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderFixed) Price2AsString() string {
	return message.StringFixed(m.Unsafe(), 296, 280)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderFixedBuilder) Price2AsString() string {
	return message.StringFixed(m.Unsafe(), 296, 280)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderFixedPointer) Price2AsString() string {
	return message.StringFixed(m.Ptr, 296, 280)
}

// Price2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderFixedPointerBuilder) Price2AsString() string {
	return message.StringFixed(m.Ptr, 296, 280)
}

// SetSymbol The symbol for the order.
func (m *SubmitNewSingleOrderBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetSymbol The symbol for the order.
func (m *SubmitNewSingleOrderPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewSingleOrderBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewSingleOrderPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m *SubmitNewSingleOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m *SubmitNewSingleOrderPointerBuilder) SetTradeAccount(value string) {
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
func (m *SubmitNewSingleOrderBuilder) SetClientOrderID(value string) {
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
func (m *SubmitNewSingleOrderPointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 22, 18, value)
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderPointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 32, 28, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 32, 28, int32(value))
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 40, 32, value)
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderPointerBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Ptr, 40, 32, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderPointerBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 56, 48, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 56, 48, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 60, 56, int32(value))
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Ptr, 60, 56, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 72, 64, int64(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 72, 64, int64(value))
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 73, 72, value)
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 73, 72, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderBuilder) SetIsParentOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 74, 73, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderPointerBuilder) SetIsParentOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 74, 73, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m *SubmitNewSingleOrderBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 78, 74, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m *SubmitNewSingleOrderPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 78, 74, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 84, 80, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Ptr, 84, 80, int32(value))
}

// SetMaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderBuilder) SetMaxShowQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 96, 88, value)
}

// SetMaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderPointerBuilder) SetMaxShowQuantity(value float64) {
	message.SetFloat64VLS(m.Ptr, 96, 88, value)
}

// SetPrice1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m *SubmitNewSingleOrderBuilder) SetPrice1AsString(value string) {
	message.SetStringVLS(&m.VLSBuilder, 100, 96, value)
}

// SetPrice1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m *SubmitNewSingleOrderPointerBuilder) SetPrice1AsString(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 100, 96, value)
}

// SetPrice2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m *SubmitNewSingleOrderBuilder) SetPrice2AsString(value string) {
	message.SetStringVLS(&m.VLSBuilder, 104, 100, value)
}

// SetPrice2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m *SubmitNewSingleOrderPointerBuilder) SetPrice2AsString(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 104, 100, value)
}

// SetSymbol The symbol for the order.
func (m SubmitNewSingleOrderFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetSymbol The symbol for the order.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetTradeAccount(value string) {
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
func (m SubmitNewSingleOrderFixedBuilder) SetClientOrderID(value string) {
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
func (m SubmitNewSingleOrderFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 148, 116, value)
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderFixedBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 152, 148, int32(value))
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 152, 148, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderFixedBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 156, 152, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 156, 152, int32(value))
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderFixedBuilder) SetPrice1(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 168, 160, value)
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetPrice1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 168, 160, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderFixedBuilder) SetPrice2(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 176, 168, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetPrice2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 176, 168, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderFixedBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 184, 176, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 184, 176, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderFixedBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Unsafe(), 188, 184, int32(value))
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 188, 184, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderFixedBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 200, 192, int64(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 200, 192, int64(value))
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderFixedBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Unsafe(), 201, 200, value)
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Ptr, 201, 200, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderFixedBuilder) SetIsParentOrder(value bool) {
	message.SetBoolFixed(m.Unsafe(), 202, 201, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetIsParentOrder(value bool) {
	message.SetBoolFixed(m.Ptr, 202, 201, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 250, 202, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 250, 202, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderFixedBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 256, 252, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 256, 252, int32(value))
}

// SetMaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderFixedBuilder) SetMaxShowQuantity(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 264, 256, value)
}

// SetMaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetMaxShowQuantity(value float64) {
	message.SetFloat64Fixed(m.Ptr, 264, 256, value)
}

// SetPrice1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderFixedBuilder) SetPrice1AsString(value string) {
	message.SetStringFixed(m.Unsafe(), 280, 264, value)
}

// SetPrice1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetPrice1AsString(value string) {
	message.SetStringFixed(m.Ptr, 280, 264, value)
}

// SetPrice2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderFixedBuilder) SetPrice2AsString(value string) {
	message.SetStringFixed(m.Unsafe(), 296, 280, value)
}

// SetPrice2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m SubmitNewSingleOrderFixedPointerBuilder) SetPrice2AsString(value string) {
	message.SetStringFixed(m.Ptr, 296, 280, value)
}

// Copy
func (m SubmitNewSingleOrder) Copy(to SubmitNewSingleOrderBuilder) {
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
func (m SubmitNewSingleOrderBuilder) Copy(to SubmitNewSingleOrderBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrder) CopyPointer(to SubmitNewSingleOrderPointerBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrderBuilder) CopyPointer(to SubmitNewSingleOrderPointerBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrder) CopyToPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrderBuilder) CopyToPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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
func (m SubmitNewSingleOrder) CopyTo(to SubmitNewSingleOrderFixedBuilder) {
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
func (m SubmitNewSingleOrderBuilder) CopyTo(to SubmitNewSingleOrderFixedBuilder) {
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
func (m SubmitNewSingleOrderFixed) Copy(to SubmitNewSingleOrderFixedBuilder) {
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
func (m SubmitNewSingleOrderFixedBuilder) Copy(to SubmitNewSingleOrderFixedBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrderFixed) CopyPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrderFixedBuilder) CopyPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrderFixed) CopyToPointer(to SubmitNewSingleOrderPointerBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrderFixedBuilder) CopyToPointer(to SubmitNewSingleOrderPointerBuilder) {
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
func (m SubmitNewSingleOrderFixed) CopyTo(to SubmitNewSingleOrderBuilder) {
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
func (m SubmitNewSingleOrderFixedBuilder) CopyTo(to SubmitNewSingleOrderBuilder) {
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
func (m SubmitNewSingleOrderPointer) Copy(to SubmitNewSingleOrderBuilder) {
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
func (m SubmitNewSingleOrderPointerBuilder) Copy(to SubmitNewSingleOrderBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrderPointer) CopyPointer(to SubmitNewSingleOrderPointerBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrderPointerBuilder) CopyPointer(to SubmitNewSingleOrderPointerBuilder) {
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
func (m SubmitNewSingleOrderPointer) CopyTo(to SubmitNewSingleOrderFixedBuilder) {
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
func (m SubmitNewSingleOrderPointerBuilder) CopyTo(to SubmitNewSingleOrderFixedBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrderPointer) CopyToPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrderPointerBuilder) CopyToPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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
func (m SubmitNewSingleOrderFixedPointer) Copy(to SubmitNewSingleOrderFixedBuilder) {
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
func (m SubmitNewSingleOrderFixedPointerBuilder) Copy(to SubmitNewSingleOrderFixedBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrderFixedPointer) CopyPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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

// CopyPointer
func (m SubmitNewSingleOrderFixedPointerBuilder) CopyPointer(to SubmitNewSingleOrderFixedPointerBuilder) {
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
func (m SubmitNewSingleOrderFixedPointer) CopyTo(to SubmitNewSingleOrderBuilder) {
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
func (m SubmitNewSingleOrderFixedPointerBuilder) CopyTo(to SubmitNewSingleOrderBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrderFixedPointer) CopyToPointer(to SubmitNewSingleOrderPointerBuilder) {
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

// CopyToPointer
func (m SubmitNewSingleOrderFixedPointerBuilder) CopyToPointer(to SubmitNewSingleOrderPointerBuilder) {
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
