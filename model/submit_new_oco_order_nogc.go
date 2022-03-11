package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SubmitNewOCOOrderPointer This is a message from the Client to the Server for submitting an order
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
type SubmitNewOCOOrderPointer struct {
	message.VLSPointer
}

// SubmitNewOCOOrderPointerBuilder This is a message from the Client to the Server for submitting an order
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
type SubmitNewOCOOrderPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSubmitNewOCOOrder() SubmitNewOCOOrderPointerBuilder {
	a := SubmitNewOCOOrderPointerBuilder{message.AllocVLSBuilder(176)}
	a.Ptr.SetBytes(0, _SubmitNewOCOOrderDefault)
	return a
}

func AllocSubmitNewOCOOrderFrom(b []byte) SubmitNewOCOOrderPointer {
	return SubmitNewOCOOrderPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                           = SubmitNewOCOOrderSize  (176)
//     Type                           = SUBMIT_NEW_OCO_ORDER  (201)
//     BaseSize                       = SubmitNewOCOOrderSize  (176)
//     Symbol                         = ""
//     Exchange                       = ""
//     ClientOrderID_1                = ""
//     OrderType_1                    = ORDER_TYPE_UNSET  (0)
//     BuySell_1                      = BUY_SELL_UNSET  (0)
//     Price1_1                       = 0
//     Price2_1                       = 0
//     Quantity_1                     = 0
//     ClientOrderID_2                = ""
//     OrderType_2                    = ORDER_TYPE_UNSET  (0)
//     BuySell_2                      = BUY_SELL_UNSET  (0)
//     Price1_2                       = 0
//     Price2_2                       = 0
//     Quantity_2                     = 0
//     TimeInForce                    = TIF_UNSET  (0)
//     GoodTillDateTime               = 0
//     TradeAccount                   = ""
//     IsAutomatedOrder               = 0
//     ParentTriggerClientOrderID     = ""
//     FreeFormText                   = ""
//     OpenOrClose                    = TRADE_UNSET  (0)
//     PartialFillHandling            = PARTIAL_FILL_UNSET  (0)
//     UseOffsets                     = 0
//     OffsetFromParent1              = 0
//     OffsetFromParent2              = 0
//     MaintainSamePricesOnParentFill = 0
//     Price1_1AsString               = ""
//     Price2_1AsString               = ""
//     Price1_2AsString               = ""
//     Price2_2AsString               = ""
// }
func (m SubmitNewOCOOrderPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewOCOOrderDefault)
}

// ToBuilder
func (m SubmitNewOCOOrderPointer) ToBuilder() SubmitNewOCOOrderPointerBuilder {
	return SubmitNewOCOOrderPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SubmitNewOCOOrderPointerBuilder) Finish() SubmitNewOCOOrderPointer {
	return SubmitNewOCOOrderPointer{m.VLSPointerBuilder.Finish()}
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderPointer) ClientOrderID_1() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderPointerBuilder) ClientOrderID_1() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderPointer) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderPointerBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderPointer) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderPointerBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderPointer) Price1_1() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderPointerBuilder) Price1_1() float64 {
	return message.Float64VLS(m.Ptr, 40, 32)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. Price2_1 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderPointer) Price2_1() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. Price2_1 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderPointerBuilder) Price2_1() float64 {
	return message.Float64VLS(m.Ptr, 48, 40)
}

// Quantity_1 The quantity for the first order.
func (m SubmitNewOCOOrderPointer) Quantity_1() float64 {
	return message.Float64VLS(m.Ptr, 56, 48)
}

// Quantity_1 The quantity for the first order.
func (m SubmitNewOCOOrderPointerBuilder) Quantity_1() float64 {
	return message.Float64VLS(m.Ptr, 56, 48)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderPointer) ClientOrderID_2() string {
	return message.StringVLS(m.Ptr, 60, 56)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderPointerBuilder) ClientOrderID_2() string {
	return message.StringVLS(m.Ptr, 60, 56)
}

// OrderType_2 The order type for the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderPointer) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 64, 60))
}

// OrderType_2 The order type for the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderPointerBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 64, 60))
}

// BuySell_2 The side for the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderPointer) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 68, 64))
}

// BuySell_2 The side for the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderPointerBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 68, 64))
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderPointer) Price1_2() float64 {
	return message.Float64VLS(m.Ptr, 80, 72)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderPointerBuilder) Price1_2() float64 {
	return message.Float64VLS(m.Ptr, 80, 72)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. Price2_2 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderPointer) Price2_2() float64 {
	return message.Float64VLS(m.Ptr, 88, 80)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. Price2_2 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderPointerBuilder) Price2_2() float64 {
	return message.Float64VLS(m.Ptr, 88, 80)
}

// Quantity_2 The quantity for the second order.
func (m SubmitNewOCOOrderPointer) Quantity_2() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// Quantity_2 The quantity for the second order.
func (m SubmitNewOCOOrderPointerBuilder) Quantity_2() float64 {
	return message.Float64VLS(m.Ptr, 96, 88)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 100, 96))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 100, 96))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 112, 104))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 112, 104))
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 116, 112)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 116, 112)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderPointer) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 117, 116)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 117, 116)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderPointer) ParentTriggerClientOrderID() string {
	return message.StringVLS(m.Ptr, 122, 118)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderPointerBuilder) ParentTriggerClientOrderID() string {
	return message.StringVLS(m.Ptr, 122, 118)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 126, 122)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 126, 122)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 132, 128))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 132, 128))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderPointer) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8VLS(m.Ptr, 133, 132))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderPointerBuilder) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8VLS(m.Ptr, 133, 132))
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
func (m SubmitNewOCOOrderPointer) UseOffsets() uint8 {
	return message.Uint8VLS(m.Ptr, 134, 133)
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
func (m SubmitNewOCOOrderPointerBuilder) UseOffsets() uint8 {
	return message.Uint8VLS(m.Ptr, 134, 133)
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
func (m SubmitNewOCOOrderPointer) OffsetFromParent1() float64 {
	return message.Float64VLS(m.Ptr, 144, 136)
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
func (m SubmitNewOCOOrderPointerBuilder) OffsetFromParent1() float64 {
	return message.Float64VLS(m.Ptr, 144, 136)
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
func (m SubmitNewOCOOrderPointer) OffsetFromParent2() float64 {
	return message.Float64VLS(m.Ptr, 152, 144)
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
func (m SubmitNewOCOOrderPointerBuilder) OffsetFromParent2() float64 {
	return message.Float64VLS(m.Ptr, 152, 144)
}

// MaintainSamePricesOnParentFill
func (m SubmitNewOCOOrderPointer) MaintainSamePricesOnParentFill() uint8 {
	return message.Uint8VLS(m.Ptr, 153, 152)
}

// MaintainSamePricesOnParentFill
func (m SubmitNewOCOOrderPointerBuilder) MaintainSamePricesOnParentFill() uint8 {
	return message.Uint8VLS(m.Ptr, 153, 152)
}

// Price1_1AsString
func (m SubmitNewOCOOrderPointer) Price1_1AsString() string {
	return message.StringVLS(m.Ptr, 158, 154)
}

// Price1_1AsString
func (m SubmitNewOCOOrderPointerBuilder) Price1_1AsString() string {
	return message.StringVLS(m.Ptr, 158, 154)
}

// Price2_1AsString
func (m SubmitNewOCOOrderPointer) Price2_1AsString() string {
	return message.StringVLS(m.Ptr, 162, 158)
}

// Price2_1AsString
func (m SubmitNewOCOOrderPointerBuilder) Price2_1AsString() string {
	return message.StringVLS(m.Ptr, 162, 158)
}

// Price1_2AsString
func (m SubmitNewOCOOrderPointer) Price1_2AsString() string {
	return message.StringVLS(m.Ptr, 166, 162)
}

// Price1_2AsString
func (m SubmitNewOCOOrderPointerBuilder) Price1_2AsString() string {
	return message.StringVLS(m.Ptr, 166, 162)
}

// Price2_2AsString
func (m SubmitNewOCOOrderPointer) Price2_2AsString() string {
	return message.StringVLS(m.Ptr, 170, 166)
}

// Price2_2AsString
func (m SubmitNewOCOOrderPointerBuilder) Price2_2AsString() string {
	return message.StringVLS(m.Ptr, 170, 166)
}

// SetSymbol The symbol for the order.
func (m *SubmitNewOCOOrderPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewOCOOrderPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m *SubmitNewOCOOrderPointerBuilder) SetClientOrderID_1(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderPointerBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderPointerBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderPointerBuilder) SetPrice1_1(value float64) {
	message.SetFloat64VLS(m.Ptr, 40, 32, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. Price2_1 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderPointerBuilder) SetPrice2_1(value float64) {
	message.SetFloat64VLS(m.Ptr, 48, 40, value)
}

// SetQuantity_1 The quantity for the first order.
func (m SubmitNewOCOOrderPointerBuilder) SetQuantity_1(value float64) {
	message.SetFloat64VLS(m.Ptr, 56, 48, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m *SubmitNewOCOOrderPointerBuilder) SetClientOrderID_2(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 60, 56, value)
}

// SetOrderType_2 The order type for the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderPointerBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 64, 60, int32(value))
}

// SetBuySell_2 The side for the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderPointerBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 68, 64, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderPointerBuilder) SetPrice1_2(value float64) {
	message.SetFloat64VLS(m.Ptr, 80, 72, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. Price2_2 only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderPointerBuilder) SetPrice2_2(value float64) {
	message.SetFloat64VLS(m.Ptr, 88, 80, value)
}

// SetQuantity_2 The quantity for the second order.
func (m SubmitNewOCOOrderPointerBuilder) SetQuantity_2(value float64) {
	message.SetFloat64VLS(m.Ptr, 96, 88, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Ptr, 100, 96, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 112, 104, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m *SubmitNewOCOOrderPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 116, 112, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Ptr, 117, 116, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m *SubmitNewOCOOrderPointerBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 122, 118, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m *SubmitNewOCOOrderPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 126, 122, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Ptr, 132, 128, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderPointerBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8VLS(m.Ptr, 133, 132, int8(value))
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
func (m SubmitNewOCOOrderPointerBuilder) SetUseOffsets(value uint8) {
	message.SetUint8VLS(m.Ptr, 134, 133, value)
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
func (m SubmitNewOCOOrderPointerBuilder) SetOffsetFromParent1(value float64) {
	message.SetFloat64VLS(m.Ptr, 144, 136, value)
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
func (m SubmitNewOCOOrderPointerBuilder) SetOffsetFromParent2(value float64) {
	message.SetFloat64VLS(m.Ptr, 152, 144, value)
}

// SetMaintainSamePricesOnParentFill
func (m SubmitNewOCOOrderPointerBuilder) SetMaintainSamePricesOnParentFill(value uint8) {
	message.SetUint8VLS(m.Ptr, 153, 152, value)
}

// SetPrice1_1AsString
func (m *SubmitNewOCOOrderPointerBuilder) SetPrice1_1AsString(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 158, 154, value)
}

// SetPrice2_1AsString
func (m *SubmitNewOCOOrderPointerBuilder) SetPrice2_1AsString(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 162, 158, value)
}

// SetPrice1_2AsString
func (m *SubmitNewOCOOrderPointerBuilder) SetPrice1_2AsString(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 166, 162, value)
}

// SetPrice2_2AsString
func (m *SubmitNewOCOOrderPointerBuilder) SetPrice2_2AsString(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 170, 166, value)
}

// Copy
func (m SubmitNewOCOOrderPointer) Copy(to SubmitNewOCOOrderBuilder) {
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
func (m SubmitNewOCOOrderPointerBuilder) Copy(to SubmitNewOCOOrderBuilder) {
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
func (m SubmitNewOCOOrderPointer) CopyPointer(to SubmitNewOCOOrderPointerBuilder) {
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
func (m SubmitNewOCOOrderPointerBuilder) CopyPointer(to SubmitNewOCOOrderPointerBuilder) {
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
func (m SubmitNewOCOOrderPointer) CopyTo(to SubmitNewOCOOrderFixedBuilder) {
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
func (m SubmitNewOCOOrderPointerBuilder) CopyTo(to SubmitNewOCOOrderFixedBuilder) {
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
func (m SubmitNewOCOOrderPointer) CopyToPointer(to SubmitNewOCOOrderFixedPointerBuilder) {
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
func (m SubmitNewOCOOrderPointerBuilder) CopyToPointer(to SubmitNewOCOOrderFixedPointerBuilder) {
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
