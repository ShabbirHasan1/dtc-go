package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SubmitNewOCOOrderIntFixedPointer The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
// that the order prices are as integers.
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
type SubmitNewOCOOrderIntFixedPointer struct {
	message.FixedPointer
}

// SubmitNewOCOOrderIntFixedPointerBuilder The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
// that the order prices are as integers.
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
type SubmitNewOCOOrderIntFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSubmitNewOCOOrderIntFixed() SubmitNewOCOOrderIntFixedPointerBuilder {
	a := SubmitNewOCOOrderIntFixedPointerBuilder{message.AllocFixed(360)}
	a.Ptr.SetBytes(0, _SubmitNewOCOOrderIntFixedDefault)
	return a
}

func AllocSubmitNewOCOOrderIntFixedFrom(b []byte) SubmitNewOCOOrderIntFixedPointer {
	return SubmitNewOCOOrderIntFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                       = SubmitNewOCOOrderIntFixedSize  (360)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     Symbol                     = ""
//     Exchange                   = ""
//     ClientOrderID_1            = ""
//     OrderType_1                = 0
//     BuySell_1                  = 0
//     Price1_1                   = 0
//     Price2_1                   = 0
//     Quantity_1                 = 0
//     ClientOrderID_2            = ""
//     OrderType_2                = 0
//     BuySell_2                  = 0
//     Price1_2                   = 0
//     Price2_2                   = 0
//     Quantity_2                 = 0
//     TimeInForce                = 0
//     GoodTillDateTime           = 0
//     TradeAccount               = ""
//     IsAutomatedOrder           = 0
//     ParentTriggerClientOrderID = ""
//     FreeFormText               = ""
//     Divisor                    = 0
//     OpenOrClose                = 0
//     PartialFillHandling        = 0
// }
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewOCOOrderIntFixedDefault)
}

// ToBuilder
func (m SubmitNewOCOOrderIntFixedPointer) ToBuilder() SubmitNewOCOOrderIntFixedPointerBuilder {
	return SubmitNewOCOOrderIntFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SubmitNewOCOOrderIntFixedPointerBuilder) Finish() SubmitNewOCOOrderIntFixedPointer {
	return SubmitNewOCOOrderIntFixedPointer{m.FixedPointer}
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointer) ClientOrderID_1() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) ClientOrderID_1() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointer) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 120, 116))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 120, 116))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointer) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 124, 120))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 124, 120))
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointer) Price1_1() int64 {
	return message.Int64Fixed(m.Ptr, 136, 128)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Price1_1() int64 {
	return message.Int64Fixed(m.Ptr, 136, 128)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointer) Price2_1() int64 {
	return message.Int64Fixed(m.Ptr, 144, 136)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Price2_1() int64 {
	return message.Int64Fixed(m.Ptr, 144, 136)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedPointer) Quantity_1() int64 {
	return message.Int64Fixed(m.Ptr, 152, 144)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Quantity_1() int64 {
	return message.Int64Fixed(m.Ptr, 152, 144)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointer) ClientOrderID_2() string {
	return message.StringFixed(m.Ptr, 184, 152)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) ClientOrderID_2() string {
	return message.StringFixed(m.Ptr, 184, 152)
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointer) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 188, 184))
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 188, 184))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointer) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 192, 188))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 192, 188))
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointer) Price1_2() int64 {
	return message.Int64Fixed(m.Ptr, 200, 192)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Price1_2() int64 {
	return message.Int64Fixed(m.Ptr, 200, 192)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointer) Price2_2() int64 {
	return message.Int64Fixed(m.Ptr, 208, 200)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Price2_2() int64 {
	return message.Int64Fixed(m.Ptr, 208, 200)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedPointer) Quantity_2() int64 {
	return message.Int64Fixed(m.Ptr, 216, 208)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Quantity_2() int64 {
	return message.Int64Fixed(m.Ptr, 216, 208)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 220, 216))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 220, 216))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 232, 224))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 232, 224))
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 264, 232)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 264, 232)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedPointer) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 265, 264)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 265, 264)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntFixedPointer) ParentTriggerClientOrderID() string {
	return message.StringFixed(m.Ptr, 297, 265)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) ParentTriggerClientOrderID() string {
	return message.StringFixed(m.Ptr, 297, 265)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 345, 297)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 345, 297)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixedPointer) Divisor() float32 {
	return message.Float32Fixed(m.Ptr, 352, 348)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Divisor() float32 {
	return message.Float32Fixed(m.Ptr, 352, 348)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 356, 352))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 356, 352))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixedPointer) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8Fixed(m.Ptr, 357, 356))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8Fixed(m.Ptr, 357, 356))
}

// SetSymbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetClientOrderID_1(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 120, 116, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 124, 120, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice1_1(value int64) {
	message.SetInt64Fixed(m.Ptr, 136, 128, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice2_1(value int64) {
	message.SetInt64Fixed(m.Ptr, 144, 136, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetQuantity_1(value int64) {
	message.SetInt64Fixed(m.Ptr, 152, 144, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetClientOrderID_2(value string) {
	message.SetStringFixed(m.Ptr, 184, 152, value)
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 188, 184, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 192, 188, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice1_2(value int64) {
	message.SetInt64Fixed(m.Ptr, 200, 192, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice2_2(value int64) {
	message.SetInt64Fixed(m.Ptr, 208, 200, value)
}

// SetQuantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetQuantity_2(value int64) {
	message.SetInt64Fixed(m.Ptr, 216, 208, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 220, 216, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 232, 224, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 264, 232, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8Fixed(m.Ptr, 265, 264, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 297, 265, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 345, 297, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 352, 348, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 356, 352, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8Fixed(m.Ptr, 357, 356, int8(value))
}

// Copy
func (m SubmitNewOCOOrderIntFixedPointer) Copy(to SubmitNewOCOOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}

// Copy
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Copy(to SubmitNewOCOOrderIntFixedBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}

// CopyPointer
func (m SubmitNewOCOOrderIntFixedPointer) CopyPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}

// CopyPointer
func (m SubmitNewOCOOrderIntFixedPointerBuilder) CopyPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}

// CopyTo
func (m SubmitNewOCOOrderIntFixedPointer) CopyTo(to SubmitNewOCOOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}

// CopyTo
func (m SubmitNewOCOOrderIntFixedPointerBuilder) CopyTo(to SubmitNewOCOOrderIntBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}

// CopyToPointer
func (m SubmitNewOCOOrderIntFixedPointer) CopyToPointer(to SubmitNewOCOOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}

// CopyToPointer
func (m SubmitNewOCOOrderIntFixedPointerBuilder) CopyToPointer(to SubmitNewOCOOrderIntPointerBuilder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetClientOrderID_1(m.ClientOrderID_1())
	to.SetOrderType_1(m.OrderType_1())
	to.SetBuySell_1(m.BuySell_1())
	to.SetPrice1_1(m.Price1_1())
	to.SetPrice2_1(m.Price2_1())
	to.SetQuantity_1(m.Quantity_1())
	to.SetClientOrderID_2(m.ClientOrderID_2())
	to.SetOrderType_2(m.OrderType_2())
	to.SetBuySell_2(m.BuySell_2())
	to.SetPrice1_2(m.Price1_2())
	to.SetPrice2_2(m.Price2_2())
	to.SetQuantity_2(m.Quantity_2())
	to.SetTimeInForce(m.TimeInForce())
	to.SetGoodTillDateTime(m.GoodTillDateTime())
	to.SetTradeAccount(m.TradeAccount())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
	to.SetParentTriggerClientOrderID(m.ParentTriggerClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetDivisor(m.Divisor())
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
}
