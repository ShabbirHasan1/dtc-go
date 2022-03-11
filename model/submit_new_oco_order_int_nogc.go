package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SubmitNewOCOOrderIntPointer The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
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
type SubmitNewOCOOrderIntPointer struct {
	message.VLSPointer
}

// SubmitNewOCOOrderIntPointerBuilder The SubmitNewOCOOrder_INT message is identical to SubmitNewOCOOrder except
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
type SubmitNewOCOOrderIntPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSubmitNewOCOOrderInt() SubmitNewOCOOrderIntPointerBuilder {
	a := SubmitNewOCOOrderIntPointerBuilder{message.AllocVLSBuilder(144)}
	a.Ptr.SetBytes(0, _SubmitNewOCOOrderIntDefault)
	return a
}

func AllocSubmitNewOCOOrderIntFrom(b []byte) SubmitNewOCOOrderIntPointer {
	return SubmitNewOCOOrderIntPointer{VLSPointer: message.AllocVLSFrom(b)}
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
func (m SubmitNewOCOOrderIntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewOCOOrderIntDefault)
}

// ToBuilder
func (m SubmitNewOCOOrderIntPointer) ToBuilder() SubmitNewOCOOrderIntPointerBuilder {
	return SubmitNewOCOOrderIntPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SubmitNewOCOOrderIntPointerBuilder) Finish() SubmitNewOCOOrderIntPointer {
	return SubmitNewOCOOrderIntPointer{m.VLSPointerBuilder.Finish()}
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntPointer) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntPointerBuilder) Symbol() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntPointer) ClientOrderID_1() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntPointerBuilder) ClientOrderID_1() string {
	return message.StringVLS(m.Ptr, 18, 14)
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointer) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointer) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointerBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointer) Price1_1() int64 {
	return message.Int64VLS(m.Ptr, 40, 32)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointerBuilder) Price1_1() int64 {
	return message.Int64VLS(m.Ptr, 40, 32)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointer) Price2_1() int64 {
	return message.Int64VLS(m.Ptr, 48, 40)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointerBuilder) Price2_1() int64 {
	return message.Int64VLS(m.Ptr, 48, 40)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntPointer) Quantity_1() int64 {
	return message.Int64VLS(m.Ptr, 56, 48)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntPointerBuilder) Quantity_1() int64 {
	return message.Int64VLS(m.Ptr, 56, 48)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntPointer) ClientOrderID_2() string {
	return message.StringVLS(m.Ptr, 60, 56)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntPointerBuilder) ClientOrderID_2() string {
	return message.StringVLS(m.Ptr, 60, 56)
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointer) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 64, 60))
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Ptr, 64, 60))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointer) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 68, 64))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointerBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Ptr, 68, 64))
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointer) Price1_2() int64 {
	return message.Int64VLS(m.Ptr, 80, 72)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointerBuilder) Price1_2() int64 {
	return message.Int64VLS(m.Ptr, 80, 72)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointer) Price2_2() int64 {
	return message.Int64VLS(m.Ptr, 88, 80)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointerBuilder) Price2_2() int64 {
	return message.Int64VLS(m.Ptr, 88, 80)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntPointer) Quantity_2() int64 {
	return message.Int64VLS(m.Ptr, 96, 88)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntPointerBuilder) Quantity_2() int64 {
	return message.Int64VLS(m.Ptr, 96, 88)
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntPointer) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 100, 96))
}

// TimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) TimeInForce() TimeInForceEnum {
	return TimeInForceEnum(message.Int32VLS(m.Ptr, 100, 96))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntPointer) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 112, 104))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntPointerBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Ptr, 112, 104))
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 116, 112)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 116, 112)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntPointer) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 117, 116)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntPointerBuilder) IsAutomatedOrder() uint8 {
	return message.Uint8VLS(m.Ptr, 117, 116)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntPointer) ParentTriggerClientOrderID() string {
	return message.StringVLS(m.Ptr, 122, 118)
}

// ParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntPointerBuilder) ParentTriggerClientOrderID() string {
	return message.StringVLS(m.Ptr, 122, 118)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntPointer) FreeFormText() string {
	return message.StringVLS(m.Ptr, 126, 122)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m SubmitNewOCOOrderIntPointerBuilder) FreeFormText() string {
	return message.StringVLS(m.Ptr, 126, 122)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntPointer) Divisor() float32 {
	return message.Float32VLS(m.Ptr, 132, 128)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntPointerBuilder) Divisor() float32 {
	return message.Float32VLS(m.Ptr, 132, 128)
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntPointer) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 136, 132))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Ptr, 136, 132))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntPointer) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8VLS(m.Ptr, 137, 136))
}

// PartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) PartialFillHandling() PartialFillHandlingEnum {
	return PartialFillHandlingEnum(message.Int8VLS(m.Ptr, 137, 136))
}

// SetSymbol The symbol for the order.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetClientOrderID_1(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointerBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice1_1(value int64) {
	message.SetInt64VLS(m.Ptr, 40, 32, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice2_1(value int64) {
	message.SetInt64VLS(m.Ptr, 48, 40, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntPointerBuilder) SetQuantity_1(value int64) {
	message.SetInt64VLS(m.Ptr, 56, 48, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetClientOrderID_2(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 60, 56, value)
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 64, 60, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointerBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 68, 64, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice1_2(value int64) {
	message.SetInt64VLS(m.Ptr, 80, 72, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice2_2(value int64) {
	message.SetInt64VLS(m.Ptr, 88, 80, value)
}

// SetQuantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntPointerBuilder) SetQuantity_2(value int64) {
	message.SetInt64VLS(m.Ptr, 96, 88, value)
}

// SetTimeInForce The Time in Force for the order or orders (in the case of an OCO order).
// The Time in Force for the order or orders (in the case of an OCO order).
//
// For more information, refer to TimeInForceEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Ptr, 100, 96, int32(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 112, 104, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 116, 112, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntPointerBuilder) SetIsAutomatedOrder(value uint8) {
	message.SetUint8VLS(m.Ptr, 117, 116, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 122, 118, value)
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with each of the OCO orders. The Server is not under
// any obligation to use this text and it may place a limitation on the length
// of this text.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetFreeFormText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 126, 122, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Ptr, 132, 128, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Ptr, 136, 132, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8VLS(m.Ptr, 137, 136, int8(value))
}

// Copy
func (m SubmitNewOCOOrderIntPointer) Copy(to SubmitNewOCOOrderIntBuilder) {
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
func (m SubmitNewOCOOrderIntPointerBuilder) Copy(to SubmitNewOCOOrderIntBuilder) {
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
func (m SubmitNewOCOOrderIntPointer) CopyPointer(to SubmitNewOCOOrderIntPointerBuilder) {
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
func (m SubmitNewOCOOrderIntPointerBuilder) CopyPointer(to SubmitNewOCOOrderIntPointerBuilder) {
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
func (m SubmitNewOCOOrderIntPointer) CopyTo(to SubmitNewOCOOrderIntFixedBuilder) {
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
func (m SubmitNewOCOOrderIntPointerBuilder) CopyTo(to SubmitNewOCOOrderIntFixedBuilder) {
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
func (m SubmitNewOCOOrderIntPointer) CopyToPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
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
func (m SubmitNewOCOOrderIntPointerBuilder) CopyToPointer(to SubmitNewOCOOrderIntFixedPointerBuilder) {
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
