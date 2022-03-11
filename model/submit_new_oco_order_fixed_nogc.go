package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SubmitNewOCOOrderFixedPointer This is a message from the Client to the Server for submitting an order
// cancels order (OCO) pair into the market. What this means is when one
// of the orders is filled or canceled, the other order will be canceled.
// If one order partially fills, the other order will be reduced in quantity
// by the fill amount of the order that partially filled.
//
// A service provider must implement OCO orders on the server so that they
// can independently be modified (Cancel/Replace) and canceled independently
// using each order's distinct ServerOrderID. Although, if one of the orders
// is canceled by the Client, the other order will be canceled as well unless
// they have a parent order, as specified through the ParentTriggerClientOrderID
// field, in which case the other order should remain open.
//
// If the OCO order pair is rejected, this must be communicated through two
// separate OrderUpdate messages, 1 for each order, with the OrderUpdateReason
// set to NEW_ORDER_REJECTED.
type SubmitNewOCOOrderFixedPointer struct {
	message.FixedPointer
}

// SubmitNewOCOOrderFixedPointerBuilder This is a message from the Client to the Server for submitting an order
// cancels order (OCO) pair into the market. What this means is when one
// of the orders is filled or canceled, the other order will be canceled.
// If one order partially fills, the other order will be reduced in quantity
// by the fill amount of the order that partially filled.
//
// A service provider must implement OCO orders on the server so that they
// can independently be modified (Cancel/Replace) and canceled independently
// using each order's distinct ServerOrderID. Although, if one of the orders
// is canceled by the Client, the other order will be canceled as well unless
// they have a parent order, as specified through the ParentTriggerClientOrderID
// field, in which case the other order should remain open.
//
// If the OCO order pair is rejected, this must be communicated through two
// separate OrderUpdate messages, 1 for each order, with the OrderUpdateReason
// set to NEW_ORDER_REJECTED.
type SubmitNewOCOOrderFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSubmitNewOCOOrderFixed() SubmitNewOCOOrderFixedPointerBuilder {
	a := SubmitNewOCOOrderFixedPointerBuilder{message.AllocFixed(448)}
	a.Ptr.SetBytes(0, _SubmitNewOCOOrderFixedDefault)
	return a
}

func AllocSubmitNewOCOOrderFixedFrom(b []byte) SubmitNewOCOOrderFixedPointer {
	return SubmitNewOCOOrderFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                           = SubmitNewOCOOrderFixedSize  (448)
//     Type                           = SUBMIT_NEW_OCO_ORDER  (201)
//     Symbol                         = ""
//     Exchange                       = ""
//     ClientOrderID_1                = ""
//     OrderType_1                    = 0
//     BuySell_1                      = 0
//     Price1_1                       = 0
//     Price2_1                       = 0
//     Quantity_1                     = 0
//     ClientOrderID_2                = ""
//     OrderType_2                    = 0
//     BuySell_2                      = 0
//     Price1_2                       = 0
//     Price2_2                       = 0
//     Quantity_2                     = 0
//     TimeInForce                    = 0
//     GoodTillDateTime               = 0
//     TradeAccount                   = ""
//     IsAutomatedOrder               = 0
//     ParentTriggerClientOrderID     = ""
//     FreeFormText                   = ""
//     OpenOrClose                    = 0
//     PartialFillHandling            = 0
//     UseOffsets                     = 0
//     OffsetFromParent1              = 0
//     OffsetFromParent2              = 0
//     MaintainSamePricesOnParentFill = 0
//     Price1_1AsString               = ""
//     Price2_1AsString               = ""
//     Price1_2AsString               = ""
//     Price2_2AsString               = ""
// }
func (m SubmitNewOCOOrderFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewOCOOrderFixedDefault)
}

// ToBuilder
func (m SubmitNewOCOOrderFixedPointer) ToBuilder() SubmitNewOCOOrderFixedPointerBuilder {
	return SubmitNewOCOOrderFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SubmitNewOCOOrderFixedPointerBuilder) Finish() SubmitNewOCOOrderFixedPointer {
	return SubmitNewOCOOrderFixedPointer{m.FixedPointer}
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderFixedPointer) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderFixedPointerBuilder) Symbol() string {
	return message.StringFixed(m.Ptr, 68, 4)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 84, 68)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderFixedPointer) ClientOrderID_1() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderFixedPointerBuilder) ClientOrderID_1() string {
	return message.StringFixed(m.Ptr, 116, 84)
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderFixedPointer) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 120, 116))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 120, 116))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderFixedPointer) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 124, 120))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderFixedPointerBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 124, 120))
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderFixedPointer) Price1_1() float64 {
	return message.Float64Fixed(m.Ptr, 136, 128)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderFixedPointerBuilder) Price1_1() float64 {
	return message.Float64Fixed(m.Ptr, 136, 128)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. Price2_1 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderFixedPointer) Price2_1() float64 {
	return message.Float64Fixed(m.Ptr, 144, 136)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. Price2_1 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderFixedPointerBuilder) Price2_1() float64 {
	return message.Float64Fixed(m.Ptr, 144, 136)
}

// Quantity_1 The quantity for the first order.
func (m SubmitNewOCOOrderFixedPointer) Quantity_1() float64 {
	return message.Float64Fixed(m.Ptr, 152, 144)
}

// Quantity_1 The quantity for the first order.
func (m SubmitNewOCOOrderFixedPointerBuilder) Quantity_1() float64 {
	return message.Float64Fixed(m.Ptr, 152, 144)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderFixedPointer) ClientOrderID_2() string {
	return message.StringFixed(m.Ptr, 184, 152)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderFixedPointerBuilder) ClientOrderID_2() string {
	return message.StringFixed(m.Ptr, 184, 152)
}

// OrderType_2 The order type for the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderFixedPointer) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 188, 184))
}

// OrderType_2 The order type for the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Ptr, 188, 184))
}

// BuySell_2 The side for the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderFixedPointer) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 192, 188))
}

// BuySell_2 The side for the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderFixedPointerBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Ptr, 192, 188))
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderFixedPointer) Price1_2() float64 {
	return message.Float64Fixed(m.Ptr, 200, 192)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderFixedPointerBuilder) Price1_2() float64 {
	return message.Float64Fixed(m.Ptr, 200, 192)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. Price2_2 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderFixedPointer) Price2_2() float64 {
	return message.Float64Fixed(m.Ptr, 208, 200)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. Price2_2 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderFixedPointerBuilder) Price2_2() float64 {
	return message.Float64Fixed(m.Ptr, 208, 200)
}

// Quantity_2 The quantity for the second order.
func (m SubmitNewOCOOrderFixedPointer) Quantity_2() float64 {
	return message.Float64Fixed(m.Ptr, 216, 208)
}

// Quantity_2 The quantity for the second order.
func (m SubmitNewOCOOrderFixedPointerBuilder) Quantity_2() float64 {
	return message.Float64Fixed(m.Ptr, 216, 208)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderFixedPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 220, 216))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32Fixed(m.Ptr, 220, 216))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderFixedPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 232, 224))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderFixedPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 232, 224))
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderFixedPointer) TradeAccount() string {
	return message.StringFixed(m.Ptr, 264, 232)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderFixedPointerBuilder) TradeAccount() string {
	return message.StringFixed(m.Ptr, 264, 232)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderFixedPointer) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 265, 264)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderFixedPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8Fixed(m.Ptr, 265, 264)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderFixedPointer) ParentTriggerClientOrderID() string {
	return message.StringFixed(m.Ptr, 297, 265)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderFixedPointerBuilder) ParentTriggerClientOrderID() string {
	return message.StringFixed(m.Ptr, 297, 265)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderFixedPointer) FreeFormText() string {
	return message.StringFixed(m.Ptr, 345, 297)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderFixedPointerBuilder) FreeFormText() string {
	return message.StringFixed(m.Ptr, 345, 297)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderFixedPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 352, 348))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Ptr, 352, 348))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderFixedPointer) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8Fixed(m.Ptr, 353, 352))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8Fixed(m.Ptr, 353, 352))
}

// UseOffsets This field is only relevant to a Bracket order which is the case when
// the ParentTriggerClientOrderID field is set.
//
// UseOffsets can be set to 1 and indicates that the OffsetFromParent1 and
// OffsetFromParent2 fields specify the two OCO order prices as a price offset
// from the parent order Price1 field, rather than an absolute price. In
// this case Price1_1 and Price1_2 are not used.
//
// When UseOffsets is set to 0, the default, then the OCO order prices are
// specified with Price1_1 and Price1_2.
//
// When UseOffsets is set to 1 and the OffsetFromParent1 and OffsetFromParent2
// fields are set, it is necessary that the parent order Price1 field be
// set even in the case of a Market order type. In the case of a Market order
// type use the current order price. This is so that the server has a reference
// price for the offsets in case it needs to translate them to actual prices.
// price for the offsets in case it needs to translate them to actual prices.
//
// When the parent order fills, it is expected the Server will maintain the
// specified offsets to the parent order fill price for the Target and Stop
// orders.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointer) UseOffsets() uint8 {
	return message.Uint8Fixed(m.Ptr, 354, 353)
}

// UseOffsets This field is only relevant to a Bracket order which is the case when
// the ParentTriggerClientOrderID field is set.
//
// UseOffsets can be set to 1 and indicates that the OffsetFromParent1 and
// OffsetFromParent2 fields specify the two OCO order prices as a price offset
// from the parent order Price1 field, rather than an absolute price. In
// this case Price1_1 and Price1_2 are not used.
//
// When UseOffsets is set to 0, the default, then the OCO order prices are
// specified with Price1_1 and Price1_2.
//
// When UseOffsets is set to 1 and the OffsetFromParent1 and OffsetFromParent2
// fields are set, it is necessary that the parent order Price1 field be
// set even in the case of a Market order type. In the case of a Market order
// type use the current order price. This is so that the server has a reference
// price for the offsets in case it needs to translate them to actual prices.
// price for the offsets in case it needs to translate them to actual prices.
//
// When the parent order fills, it is expected the Server will maintain the
// specified offsets to the parent order fill price for the Target and Stop
// orders.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointerBuilder) UseOffsets() uint8 {
	return message.Uint8Fixed(m.Ptr, 354, 353)
}

// OffsetFromParent1 When UseOffsets is set to 1, then this field specifies the Price1_1 price
// as an offset from the parent order. In this case Price1_1 will not be
// set in the message. Instead the Server calculates that price from this
// offset and parent order price.
//
// This needs to always be set to a positive price value which is an offset
// from the parent order price. The Server will make the correct calculation
// based upon the Side and Order Type.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointer) OffsetFromParent1() float64 {
	return message.Float64Fixed(m.Ptr, 368, 360)
}

// OffsetFromParent1 When UseOffsets is set to 1, then this field specifies the Price1_1 price
// as an offset from the parent order. In this case Price1_1 will not be
// set in the message. Instead the Server calculates that price from this
// offset and parent order price.
//
// This needs to always be set to a positive price value which is an offset
// from the parent order price. The Server will make the correct calculation
// based upon the Side and Order Type.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointerBuilder) OffsetFromParent1() float64 {
	return message.Float64Fixed(m.Ptr, 368, 360)
}

// OffsetFromParent2 When UseOffsets is set to 1, then this field specifies the Price1_2 price
// as an offset from the parent order. In this case Price1_2 will not be
// set in the message. Instead the Server calculates that price from this
// offset and parent order price.
//
// This needs to always be set to a positive price value which is an offset
// from the parent order price. The Server will make the correct calculation
// based upon the Side and Order Type.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointer) OffsetFromParent2() float64 {
	return message.Float64Fixed(m.Ptr, 376, 368)
}

// OffsetFromParent2 When UseOffsets is set to 1, then this field specifies the Price1_2 price
// as an offset from the parent order. In this case Price1_2 will not be
// set in the message. Instead the Server calculates that price from this
// offset and parent order price.
//
// This needs to always be set to a positive price value which is an offset
// from the parent order price. The Server will make the correct calculation
// based upon the Side and Order Type.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointerBuilder) OffsetFromParent2() float64 {
	return message.Float64Fixed(m.Ptr, 376, 368)
}

// MaintainSamePricesOnParentFill
func (m SubmitNewOCOOrderFixedPointer) MaintainSamePricesOnParentFill() uint8 {
	return message.Uint8Fixed(m.Ptr, 377, 376)
}

// MaintainSamePricesOnParentFill
func (m SubmitNewOCOOrderFixedPointerBuilder) MaintainSamePricesOnParentFill() uint8 {
	return message.Uint8Fixed(m.Ptr, 377, 376)
}

// Price1_1AsString
func (m SubmitNewOCOOrderFixedPointer) Price1_1AsString() string {
	return message.StringFixed(m.Ptr, 393, 377)
}

// Price1_1AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) Price1_1AsString() string {
	return message.StringFixed(m.Ptr, 393, 377)
}

// Price2_1AsString
func (m SubmitNewOCOOrderFixedPointer) Price2_1AsString() string {
	return message.StringFixed(m.Ptr, 409, 393)
}

// Price2_1AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) Price2_1AsString() string {
	return message.StringFixed(m.Ptr, 409, 393)
}

// Price1_2AsString
func (m SubmitNewOCOOrderFixedPointer) Price1_2AsString() string {
	return message.StringFixed(m.Ptr, 425, 409)
}

// Price1_2AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) Price1_2AsString() string {
	return message.StringFixed(m.Ptr, 425, 409)
}

// Price2_2AsString
func (m SubmitNewOCOOrderFixedPointer) Price2_2AsString() string {
	return message.StringFixed(m.Ptr, 441, 425)
}

// Price2_2AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) Price2_2AsString() string {
	return message.StringFixed(m.Ptr, 441, 425)
}

// SetSymbol The symbol for the order.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetClientOrderID_1(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 120, 116, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 124, 120, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice1_1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 136, 128, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. Price2_1 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice2_1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 144, 136, value)
}

// SetQuantity_1 The quantity for the first order.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetQuantity_1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 152, 144, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetClientOrderID_2(value string) {
	message.SetStringFixed(m.Ptr, 184, 152, value)
}

// SetOrderType_2 The order type for the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 188, 184, int32(value))
}

// SetBuySell_2 The side for the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 192, 188, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice1_2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 200, 192, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. Price2_2 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice2_2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 208, 200, value)
}

// SetQuantity_2 The quantity for the second order.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetQuantity_2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 216, 208, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Ptr, 220, 216, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 232, 224, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 264, 232, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8Fixed(m.Ptr, 265, 264, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 297, 265, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Ptr, 345, 297, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 352, 348, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8Fixed(m.Ptr, 353, 352, int8(value))
}

// SetUseOffsets This field is only relevant to a Bracket order which is the case when
// the ParentTriggerClientOrderID field is set.
//
// UseOffsets can be set to 1 and indicates that the OffsetFromParent1 and
// OffsetFromParent2 fields specify the two OCO order prices as a price offset
// from the parent order Price1 field, rather than an absolute price. In
// this case Price1_1 and Price1_2 are not used.
//
// When UseOffsets is set to 0, the default, then the OCO order prices are
// specified with Price1_1 and Price1_2.
//
// When UseOffsets is set to 1 and the OffsetFromParent1 and OffsetFromParent2
// fields are set, it is necessary that the parent order Price1 field be
// set even in the case of a Market order type. In the case of a Market order
// type use the current order price. This is so that the server has a reference
// price for the offsets in case it needs to translate them to actual prices.
// price for the offsets in case it needs to translate them to actual prices.
//
// When the parent order fills, it is expected the Server will maintain the
// specified offsets to the parent order fill price for the Target and Stop
// orders.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetUseOffsets(value uint8) {
	message.SetUint8Fixed(m.Ptr, 354, 353, value)
}

// SetOffsetFromParent1 When UseOffsets is set to 1, then this field specifies the Price1_1 price
// as an offset from the parent order. In this case Price1_1 will not be
// set in the message. Instead the Server calculates that price from this
// offset and parent order price.
//
// This needs to always be set to a positive price value which is an offset
// from the parent order price. The Server will make the correct calculation
// based upon the Side and Order Type.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetOffsetFromParent1(value float64) {
	message.SetFloat64Fixed(m.Ptr, 368, 360, value)
}

// SetOffsetFromParent2 When UseOffsets is set to 1, then this field specifies the Price1_2 price
// as an offset from the parent order. In this case Price1_2 will not be
// set in the message. Instead the Server calculates that price from this
// offset and parent order price.
//
// This needs to always be set to a positive price value which is an offset
// from the parent order price. The Server will make the correct calculation
// based upon the Side and Order Type.
//
// A Server is not required to support this field.
func (m SubmitNewOCOOrderFixedPointerBuilder) SetOffsetFromParent2(value float64) {
	message.SetFloat64Fixed(m.Ptr, 376, 368, value)
}

// SetMaintainSamePricesOnParentFill
func (m SubmitNewOCOOrderFixedPointerBuilder) SetMaintainSamePricesOnParentFill(value uint8) {
	message.SetUint8Fixed(m.Ptr, 377, 376, value)
}

// SetPrice1_1AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice1_1AsString(value string) {
	message.SetStringFixed(m.Ptr, 393, 377, value)
}

// SetPrice2_1AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice2_1AsString(value string) {
	message.SetStringFixed(m.Ptr, 409, 393, value)
}

// SetPrice1_2AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice1_2AsString(value string) {
	message.SetStringFixed(m.Ptr, 425, 409, value)
}

// SetPrice2_2AsString
func (m SubmitNewOCOOrderFixedPointerBuilder) SetPrice2_2AsString(value string) {
	message.SetStringFixed(m.Ptr, 441, 425, value)
}

// Copy
func (m SubmitNewOCOOrderFixedPointer) Copy(to SubmitNewOCOOrderFixedBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}

// Copy
func (m SubmitNewOCOOrderFixedPointerBuilder) Copy(to SubmitNewOCOOrderFixedBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}

// CopyPointer
func (m SubmitNewOCOOrderFixedPointer) CopyPointer(to SubmitNewOCOOrderFixedPointerBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}

// CopyPointer
func (m SubmitNewOCOOrderFixedPointerBuilder) CopyPointer(to SubmitNewOCOOrderFixedPointerBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}

// CopyTo
func (m SubmitNewOCOOrderFixedPointer) CopyTo(to SubmitNewOCOOrderBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}

// CopyTo
func (m SubmitNewOCOOrderFixedPointerBuilder) CopyTo(to SubmitNewOCOOrderBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}

// CopyToPointer
func (m SubmitNewOCOOrderFixedPointer) CopyToPointer(to SubmitNewOCOOrderPointerBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}

// CopyToPointer
func (m SubmitNewOCOOrderFixedPointerBuilder) CopyToPointer(to SubmitNewOCOOrderPointerBuilder) {
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
	to.SetOpenOrClose(m.OpenOrClose())
	to.SetPartialFillHandling(m.PartialFillHandling())
	to.SetUseOffsets(m.UseOffsets())
	to.SetOffsetFromParent1(m.OffsetFromParent1())
	to.SetOffsetFromParent2(m.OffsetFromParent2())
	to.SetMaintainSamePricesOnParentFill(m.MaintainSamePricesOnParentFill())
	to.SetPrice1_1AsString(m.Price1_1AsString())
	to.SetPrice2_1AsString(m.Price2_1AsString())
	to.SetPrice1_2AsString(m.Price1_2AsString())
	to.SetPrice2_2AsString(m.Price2_2AsString())
}
