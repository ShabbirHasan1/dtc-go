package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = SubmitNewSingleOrderSize  (104)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     BaseSize         = SubmitNewSingleOrderSize  (104)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = 0
//     BuySell          = 0
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = 0
//     GoodTillDateTime = 0
//     IsAutomatedOrder = 0
//     IsParentOrder    = 0
//     FreeFormText     = ""
//     OpenOrClose      = 0
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
var _SubmitNewSingleOrderDefault = []byte{104, 0, 208, 0, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SubmitNewSingleOrderSize = 104

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

// Clear
// {
//     Size             = SubmitNewSingleOrderSize  (104)
//     Type             = SUBMIT_NEW_SINGLE_ORDER  (208)
//     BaseSize         = SubmitNewSingleOrderSize  (104)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = 0
//     BuySell          = 0
//     Price1           = 0
//     Price2           = 0
//     Quantity         = 0
//     TimeInForce      = 0
//     GoodTillDateTime = 0
//     IsAutomatedOrder = 0
//     IsParentOrder    = 0
//     FreeFormText     = ""
//     OpenOrClose      = 0
//     MaxShowQuantity  = 0
//     Price1AsString   = ""
//     Price2AsString   = ""
// }
func (m SubmitNewSingleOrderBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewSingleOrderDefault)
}

// ToBuilder
func (m SubmitNewSingleOrder) ToBuilder() SubmitNewSingleOrderBuilder {
	return SubmitNewSingleOrderBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SubmitNewSingleOrderBuilder) Finish() SubmitNewSingleOrder {
	return SubmitNewSingleOrder{m.VLSBuilder.Finish()}
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewSingleOrderBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// TradeAccount This is the trade account as a text string that the order belongs to.
func (m SubmitNewSingleOrderBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
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

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 32, 28))
}

// BuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderBuilder) BuySell() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 32, 28))
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

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 73, 72)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 73, 72)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrder) IsParentOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 74, 73)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderBuilder) IsParentOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 74, 73)
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

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 84, 80))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 84, 80))
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

// SetSymbol The symbol for the order.
func (m *SubmitNewSingleOrderBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewSingleOrderBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetTradeAccount This is the trade account as a text string that the order belongs to.
func (m *SubmitNewSingleOrderBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
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

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 32, 28, int32(value))
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderBuilder) SetPrice1(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 40, 32, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderBuilder) SetPrice2(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 48, 40, value)
}

// SetQuantity The quantity of the order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewSingleOrderBuilder) SetQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 56, 48, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewSingleOrderBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 60, 56, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 72, 64, int64(value))
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 73, 72, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderBuilder) SetIsParentOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 74, 73, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m *SubmitNewSingleOrderBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 78, 74, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 84, 80, int32(value))
}

// SetMaxShowQuantity This field is provided to the exchange and represents the maximum quantity
// to show in the limit order book for the order.
func (m SubmitNewSingleOrderBuilder) SetMaxShowQuantity(value float64) {
	message.SetFloat64VLS(m.Unsafe(), 96, 88, value)
}

// SetPrice1AsString This is an optional field which may be used by the Server.
//
// This field is the order price 1 as a string.
func (m *SubmitNewSingleOrderBuilder) SetPrice1AsString(value string) {
	message.SetStringVLS(&m.VLSBuilder, 100, 96, value)
}

// SetPrice2AsString This is an optional field which may be used by the Server.
//
// This field is the order price 2 as a string.
func (m *SubmitNewSingleOrderBuilder) SetPrice2AsString(value string) {
	message.SetStringVLS(&m.VLSBuilder, 104, 100, value)
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
