package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                       = SubmitNewOCOOrderIntSize  (144)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     BaseSize                   = SubmitNewOCOOrderIntSize  (144)
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
var _SubmitNewOCOOrderIntDefault = []byte{144, 0, 207, 0, 144, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SubmitNewOCOOrderIntSize = 144

// SubmitNewOCOOrderInt The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
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
type SubmitNewOCOOrderInt struct {
	message.VLS
}

// SubmitNewOCOOrderIntBuilder The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
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
type SubmitNewOCOOrderIntBuilder struct {
	message.VLSBuilder
}

func NewSubmitNewOCOOrderIntFrom(b []byte) SubmitNewOCOOrderInt {
	return SubmitNewOCOOrderInt{VLS: message.NewVLSFrom(b)}
}

func WrapSubmitNewOCOOrderInt(b []byte) SubmitNewOCOOrderInt {
	return SubmitNewOCOOrderInt{VLS: message.WrapVLS(b)}
}

func NewSubmitNewOCOOrderInt() SubmitNewOCOOrderIntBuilder {
	a := SubmitNewOCOOrderIntBuilder{message.NewVLSBuilder(144)}
	a.Unsafe().SetBytes(0, _SubmitNewOCOOrderIntDefault)
	return a
}

// Clear
// {
//     Size                       = SubmitNewOCOOrderIntSize  (144)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     BaseSize                   = SubmitNewOCOOrderIntSize  (144)
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
func (m SubmitNewOCOOrderIntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewOCOOrderIntDefault)
}

// ToBuilder
func (m SubmitNewOCOOrderInt) ToBuilder() SubmitNewOCOOrderIntBuilder {
	return SubmitNewOCOOrderIntBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SubmitNewOCOOrderIntBuilder) Finish() SubmitNewOCOOrderInt {
	return SubmitNewOCOOrderInt{m.VLSBuilder.Finish()}
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderInt) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderInt) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderInt) ClientOrderID_1() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntBuilder) ClientOrderID_1() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderInt) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderInt) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderInt) Price1_1() int64 {
	return message.Int64VLS(m.Unsafe(), 40, 32)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) Price1_1() int64 {
	return message.Int64VLS(m.Unsafe(), 40, 32)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderInt) Price2_1() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) Price2_1() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderInt) Quantity_1() int64 {
	return message.Int64VLS(m.Unsafe(), 56, 48)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) Quantity_1() int64 {
	return message.Int64VLS(m.Unsafe(), 56, 48)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderInt) ClientOrderID_2() string {
	return message.StringVLS(m.Unsafe(), 60, 56)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntBuilder) ClientOrderID_2() string {
	return message.StringVLS(m.Unsafe(), 60, 56)
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderInt) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 64, 60))
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 64, 60))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderInt) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 68, 64))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 68, 64))
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderInt) Price1_2() int64 {
	return message.Int64VLS(m.Unsafe(), 80, 72)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) Price1_2() int64 {
	return message.Int64VLS(m.Unsafe(), 80, 72)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderInt) Price2_2() int64 {
	return message.Int64VLS(m.Unsafe(), 88, 80)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) Price2_2() int64 {
	return message.Int64VLS(m.Unsafe(), 88, 80)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderInt) Quantity_2() int64 {
	return message.Int64VLS(m.Unsafe(), 96, 88)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) Quantity_2() int64 {
	return message.Int64VLS(m.Unsafe(), 96, 88)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderInt) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 100, 96))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Unsafe(), 100, 96))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderInt) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 112, 104))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 112, 104))
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderInt) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 116, 112)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 116, 112)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderInt) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 117, 116)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Unsafe(), 117, 116)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderInt) ParentTriggerClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 122, 118)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntBuilder) ParentTriggerClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 122, 118)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderInt) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 126, 122)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntBuilder) FreeFormText() string {
	return message.StringVLS(m.Unsafe(), 126, 122)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderInt) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 132, 128)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntBuilder) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 132, 128)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderInt) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 136, 132))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 136, 132))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderInt) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8VLS(m.Unsafe(), 137, 136))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntBuilder) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8VLS(m.Unsafe(), 137, 136))
}

// SetSymbol The symbol for the order.
func (m *SubmitNewOCOOrderIntBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewOCOOrderIntBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m *SubmitNewOCOOrderIntBuilder) SetClientOrderID_1(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) SetPrice1_1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 40, 32, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) SetPrice2_1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 48, 40, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) SetQuantity_1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 56, 48, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m *SubmitNewOCOOrderIntBuilder) SetClientOrderID_2(value string) {
	message.SetStringVLS(&m.VLSBuilder, 60, 56, value)
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 64, 60, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 68, 64, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) SetPrice1_2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 80, 72, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) SetPrice2_2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 88, 80, value)
}

// SetQuantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) SetQuantity_2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 96, 88, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 100, 96, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 112, 104, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m *SubmitNewOCOOrderIntBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 116, 112, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 117, 116, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m *SubmitNewOCOOrderIntBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 122, 118, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m *SubmitNewOCOOrderIntBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 126, 122, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 132, 128, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 136, 132, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8VLS(m.Unsafe(), 137, 136, int8(value))
}

// Copy
func (m SubmitNewOCOOrderInt) Copy(to SubmitNewOCOOrderIntBuilder) {
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
func (m SubmitNewOCOOrderIntBuilder) Copy(to SubmitNewOCOOrderIntBuilder) {
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
func (m SubmitNewOCOOrderInt) CopyPointer(to SubmitNewOCOOrderIntPointerBuilder) {
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
func (m SubmitNewOCOOrderIntBuilder) CopyPointer(to SubmitNewOCOOrderIntPointerBuilder) {
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
func (m SubmitNewOCOOrderInt) CopyToPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
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
func (m SubmitNewOCOOrderIntBuilder) CopyToPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
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
func (m SubmitNewOCOOrderInt) CopyTo(to SubmitNewOCOOrderIntFixedBuilder) {
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
func (m SubmitNewOCOOrderIntBuilder) CopyTo(to SubmitNewOCOOrderIntFixedBuilder) {
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
