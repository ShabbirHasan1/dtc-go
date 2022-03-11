package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
//     Size             = SubmitNewSingleOrderIntFixedSize  (264)
//     Type             = SUBMIT_NEW_SINGLE_ORDER_INT  (206)
//     Symbol           = ""
//     Exchange         = ""
//     TradeAccount     = ""
//     ClientOrderID    = ""
//     OrderType        = 0
//     BuySell          = 0
//     Price1           = 0
//     Price2           = 0
//     Divisor          = 0
//     Quantity         = 0
//     TimeInForce      = 0
//     GoodTillDateTime = 0
//     IsAutomatedOrder = 0
//     IsParentOrder    = 0
//     FreeFormText     = ""
//     OpenOrClose      = 0
// }
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewSingleOrderIntFixedDefault)
}

// ToBuilder
func (m SubmitNewSingleOrderIntFixedPointer) ToBuilder() SubmitNewSingleOrderIntFixedPointerBuilder {
	return SubmitNewSingleOrderIntFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SubmitNewSingleOrderIntFixedPointerBuilder) Finish() SubmitNewSingleOrderIntFixedPointer {
	return SubmitNewSingleOrderIntFixedPointer{m.FixedPointer}
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
func (m SubmitNewSingleOrderIntFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
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
func (m SubmitNewSingleOrderIntFixedPointer) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 152, 148))
}

// OrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) OrderType() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 152, 148))
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
func (m SubmitNewSingleOrderIntFixedPointer) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 209, 208)
}

// IsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 209, 208)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedPointer) IsParentOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 210, 209)
}

// IsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) IsParentOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 210, 209)
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
func (m SubmitNewSingleOrderIntFixedPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 264, 260))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 264, 260))
}

// SetSymbol The symbol for the order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
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
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 148, 116, value)
}

// SetOrderType The order type. For list of order types, refer to OrderTypeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetOrderType(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 152, 148, int32(value))
}

// SetBuySell The side of the order. Either Buy or Sell.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetBuySell(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 156, 152, int32(value))
}

// SetPrice1 This is the price of the order. This is the limit price for a Limit order,
// the stop price for a Stop order, or the trigger price for a Market if
// Touched order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetPrice1(value int64) {
	message.SetInt64Fixed(m.Ptr, 168, 160, value)
}

// SetPrice2 For a Stop-Limit order, this is the limit price. This only applies to
// Stop-Limit orders.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetPrice2(value int64) {
	message.SetInt64Fixed(m.Ptr, 176, 168, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the symbol being traded. The Server needs to divide the Price1
// and Price2 fields by this Divisor to get the prices as float values.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 180, 176, value)
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
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 196, 192, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 208, 200, int64(value))
}

// SetIsAutomatedOrder This is set 1 to signify the order has been submitted by an automated
// trading process.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8Fixed(m.Ptr, 209, 208, value)
}

// SetIsParentOrder The Client will set this to 1 when the order is part of a bracket order.
// This indicates that this is the parent order. A bracket order will consist
// of a SubmitNewSingleOrder message followed by a SubmitNewOCOOrder message.
// The Server will use IsParentOrder as a flag to know that this message
// is a parent order. The Server will hold onto this order until it receives
// the subsequent SubmitNewOCOOrder message and then process all of the orders
// as one complete set.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetIsParentOrder(value uint8) {
	message.SetUint8Fixed(m.Ptr, 210, 209, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order. The Server is not under any obligation
// to use this text and it may place a limitation on the length of this text.
// to use this text and it may place a limitation on the length of this text.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 258, 210, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewSingleOrderIntFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 264, 260, int32(value))
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
