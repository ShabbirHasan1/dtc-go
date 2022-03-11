package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _SubmitNewOCOOrderIntFixedDefault = []byte{104, 1, 207, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SubmitNewOCOOrderIntFixedSize = 360

// SubmitNewOCOOrderIntFixed The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
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
type SubmitNewOCOOrderIntFixed struct {
	message.Fixed
}

// SubmitNewOCOOrderIntFixedBuilder The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
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
type SubmitNewOCOOrderIntFixedBuilder struct {
	message.Fixed
}

func NewSubmitNewOCOOrderIntFixedFrom(b []byte) SubmitNewOCOOrderIntFixed {
	return SubmitNewOCOOrderIntFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSubmitNewOCOOrderIntFixed(b []byte) SubmitNewOCOOrderIntFixed {
	return SubmitNewOCOOrderIntFixed{Fixed: message.WrapFixed(b)}
}

func NewSubmitNewOCOOrderIntFixed() SubmitNewOCOOrderIntFixedBuilder {
	a := SubmitNewOCOOrderIntFixedBuilder{message.NewFixed(360)}
	a.Unsafe().SetBytes(0, _SubmitNewOCOOrderIntFixedDefault)
	return a
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
func (m SubmitNewOCOOrderIntFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewOCOOrderIntFixedDefault)
}

// ToBuilder
func (m SubmitNewOCOOrderIntFixed) ToBuilder() SubmitNewOCOOrderIntFixedBuilder {
	return SubmitNewOCOOrderIntFixedBuilder{m.Fixed}
}

// Finish
func (m SubmitNewOCOOrderIntFixedBuilder) Finish() SubmitNewOCOOrderIntFixed {
	return SubmitNewOCOOrderIntFixed{m.Fixed}
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixed) ClientOrderID_1() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) ClientOrderID_1() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixed) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 120, 116))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 120, 116))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixed) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 124, 120))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 124, 120))
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixed) Price1_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 136, 128)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) Price1_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 136, 128)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixed) Price2_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 144, 136)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) Price2_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 144, 136)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixed) Quantity_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 152, 144)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) Quantity_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 152, 144)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixed) ClientOrderID_2() string {
	return message.StringFixed(m.Unsafe(), 184, 152)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) ClientOrderID_2() string {
	return message.StringFixed(m.Unsafe(), 184, 152)
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixed) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 188, 184))
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 188, 184))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixed) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 192, 188))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 192, 188))
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixed) Price1_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 200, 192)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) Price1_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 200, 192)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixed) Price2_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 208, 200)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) Price2_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 208, 200)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixed) Quantity_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 216, 208)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) Quantity_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 216, 208)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntFixed) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 220, 216))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Unsafe(), 220, 216))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixed) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 232, 224))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixedBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 232, 224))
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 264, 232)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 264, 232)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixed) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 265, 264)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 265, 264)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntFixed) ParentTriggerClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 297, 265)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntFixedBuilder) ParentTriggerClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 297, 265)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntFixed) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 345, 297)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntFixedBuilder) FreeFormText() string {
	return message.StringFixed(m.Unsafe(), 345, 297)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixed) Divisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 352, 348)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixedBuilder) Divisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 352, 348)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixed) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 356, 352))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 356, 352))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixed) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8Fixed(m.Unsafe(), 357, 356))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8Fixed(m.Unsafe(), 357, 356))
}

// SetSymbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetClientOrderID_1(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 120, 116, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 124, 120, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice1_1(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 136, 128, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice2_1(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 144, 136, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) SetQuantity_1(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 152, 144, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetClientOrderID_2(value string) {
	message.SetStringFixed(m.Unsafe(), 184, 152, value)
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 188, 184, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 192, 188, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice1_2(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 200, 192, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice2_2(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 208, 200, value)
}

// SetQuantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) SetQuantity_2(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 216, 208, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Unsafe(), 220, 216, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixedBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 232, 224, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 264, 232, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 265, 264, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntFixedBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 297, 265, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 345, 297, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixedBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 352, 348, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 356, 352, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8Fixed(m.Unsafe(), 357, 356, int8(value))
}

// Copy
func (m SubmitNewOCOOrderIntFixed) Copy(to SubmitNewOCOOrderIntFixedBuilder) {
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
func (m SubmitNewOCOOrderIntFixedBuilder) Copy(to SubmitNewOCOOrderIntFixedBuilder) {
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
func (m SubmitNewOCOOrderIntFixed) CopyPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
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
func (m SubmitNewOCOOrderIntFixedBuilder) CopyPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
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
func (m SubmitNewOCOOrderIntFixed) CopyToPointer(to SubmitNewOCOOrderIntPointerBuilder) {
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
func (m SubmitNewOCOOrderIntFixedBuilder) CopyToPointer(to SubmitNewOCOOrderIntPointerBuilder) {
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
func (m SubmitNewOCOOrderIntFixed) CopyTo(to SubmitNewOCOOrderIntBuilder) {
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
func (m SubmitNewOCOOrderIntFixedBuilder) CopyTo(to SubmitNewOCOOrderIntBuilder) {
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
