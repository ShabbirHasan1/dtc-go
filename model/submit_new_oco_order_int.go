package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SubmitNewOCOOrderIntSize = 144

const SubmitNewOCOOrderIntFixedSize = 360

// {
//     Size                       = SubmitNewOCOOrderIntSize  (144)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     BaseSize                   = SubmitNewOCOOrderIntSize  (144)
//     Symbol                     = ""
//     Exchange                   = ""
//     ClientOrderID_1            = ""
//     OrderType_1                = ORDER_TYPE_UNSET  (0)
//     BuySell_1                  = BUY_SELL_UNSET  (0)
//     Price1_1                   = 0
//     Price2_1                   = 0
//     Quantity_1                 = 0
//     ClientOrderID_2            = ""
//     OrderType_2                = ORDER_TYPE_UNSET  (0)
//     BuySell_2                  = BUY_SELL_UNSET  (0)
//     Price1_2                   = 0
//     Price2_2                   = 0
//     Quantity_2                 = 0
//     TimeInForce                = TIF_UNSET  (0)
//     GoodTillDateTime           = 0
//     TradeAccount               = ""
//     IsAutomatedOrder           = false
//     ParentTriggerClientOrderID = ""
//     FreeFormText               = ""
//     Divisor                    = 0
//     OpenOrClose                = TRADE_UNSET  (0)
//     PartialFillHandling        = PARTIAL_FILL_UNSET  (0)
// }
var _SubmitNewOCOOrderIntDefault = []byte{144, 0, 207, 0, 144, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                       = SubmitNewOCOOrderIntFixedSize  (360)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     Symbol                     = ""
//     Exchange                   = ""
//     ClientOrderID_1            = ""
//     OrderType_1                = ORDER_TYPE_UNSET  (0)
//     BuySell_1                  = BUY_SELL_UNSET  (0)
//     Price1_1                   = 0
//     Price2_1                   = 0
//     Quantity_1                 = 0
//     ClientOrderID_2            = ""
//     OrderType_2                = ORDER_TYPE_UNSET  (0)
//     BuySell_2                  = BUY_SELL_UNSET  (0)
//     Price1_2                   = 0
//     Price2_2                   = 0
//     Quantity_2                 = 0
//     TimeInForce                = TIF_UNSET  (0)
//     GoodTillDateTime           = 0
//     TradeAccount               = ""
//     IsAutomatedOrder           = false
//     ParentTriggerClientOrderID = ""
//     FreeFormText               = ""
//     Divisor                    = 0
//     OpenOrClose                = TRADE_UNSET  (0)
//     PartialFillHandling        = PARTIAL_FILL_UNSET  (0)
// }
var _SubmitNewOCOOrderIntFixedDefault = []byte{104, 1, 207, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func AllocSubmitNewOCOOrderInt() SubmitNewOCOOrderIntPointerBuilder {
	a := SubmitNewOCOOrderIntPointerBuilder{message.AllocVLSBuilder(144)}
	a.Ptr.SetBytes(0, _SubmitNewOCOOrderIntDefault)
	return a
}

func AllocSubmitNewOCOOrderIntFrom(b []byte) SubmitNewOCOOrderIntPointer {
	return SubmitNewOCOOrderIntPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size                       = SubmitNewOCOOrderIntSize  (144)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     BaseSize                   = SubmitNewOCOOrderIntSize  (144)
//     Symbol                     = ""
//     Exchange                   = ""
//     ClientOrderID_1            = ""
//     OrderType_1                = ORDER_TYPE_UNSET  (0)
//     BuySell_1                  = BUY_SELL_UNSET  (0)
//     Price1_1                   = 0
//     Price2_1                   = 0
//     Quantity_1                 = 0
//     ClientOrderID_2            = ""
//     OrderType_2                = ORDER_TYPE_UNSET  (0)
//     BuySell_2                  = BUY_SELL_UNSET  (0)
//     Price1_2                   = 0
//     Price2_2                   = 0
//     Quantity_2                 = 0
//     TimeInForce                = TIF_UNSET  (0)
//     GoodTillDateTime           = 0
//     TradeAccount               = ""
//     IsAutomatedOrder           = false
//     ParentTriggerClientOrderID = ""
//     FreeFormText               = ""
//     Divisor                    = 0
//     OpenOrClose                = TRADE_UNSET  (0)
//     PartialFillHandling        = PARTIAL_FILL_UNSET  (0)
// }
func (m SubmitNewOCOOrderIntBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewOCOOrderIntDefault)
}

// Clear
// {
//     Size                       = SubmitNewOCOOrderIntFixedSize  (360)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     Symbol                     = ""
//     Exchange                   = ""
//     ClientOrderID_1            = ""
//     OrderType_1                = ORDER_TYPE_UNSET  (0)
//     BuySell_1                  = BUY_SELL_UNSET  (0)
//     Price1_1                   = 0
//     Price2_1                   = 0
//     Quantity_1                 = 0
//     ClientOrderID_2            = ""
//     OrderType_2                = ORDER_TYPE_UNSET  (0)
//     BuySell_2                  = BUY_SELL_UNSET  (0)
//     Price1_2                   = 0
//     Price2_2                   = 0
//     Quantity_2                 = 0
//     TimeInForce                = TIF_UNSET  (0)
//     GoodTillDateTime           = 0
//     TradeAccount               = ""
//     IsAutomatedOrder           = false
//     ParentTriggerClientOrderID = ""
//     FreeFormText               = ""
//     Divisor                    = 0
//     OpenOrClose                = TRADE_UNSET  (0)
//     PartialFillHandling        = PARTIAL_FILL_UNSET  (0)
// }
func (m SubmitNewOCOOrderIntFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SubmitNewOCOOrderIntFixedDefault)
}

// Clear
// {
//     Size                       = SubmitNewOCOOrderIntSize  (144)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     BaseSize                   = SubmitNewOCOOrderIntSize  (144)
//     Symbol                     = ""
//     Exchange                   = ""
//     ClientOrderID_1            = ""
//     OrderType_1                = ORDER_TYPE_UNSET  (0)
//     BuySell_1                  = BUY_SELL_UNSET  (0)
//     Price1_1                   = 0
//     Price2_1                   = 0
//     Quantity_1                 = 0
//     ClientOrderID_2            = ""
//     OrderType_2                = ORDER_TYPE_UNSET  (0)
//     BuySell_2                  = BUY_SELL_UNSET  (0)
//     Price1_2                   = 0
//     Price2_2                   = 0
//     Quantity_2                 = 0
//     TimeInForce                = TIF_UNSET  (0)
//     GoodTillDateTime           = 0
//     TradeAccount               = ""
//     IsAutomatedOrder           = false
//     ParentTriggerClientOrderID = ""
//     FreeFormText               = ""
//     Divisor                    = 0
//     OpenOrClose                = TRADE_UNSET  (0)
//     PartialFillHandling        = PARTIAL_FILL_UNSET  (0)
// }
func (m SubmitNewOCOOrderIntPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewOCOOrderIntDefault)
}

// Clear
// {
//     Size                       = SubmitNewOCOOrderIntFixedSize  (360)
//     Type                       = SUBMIT_NEW_OCO_ORDER_INT  (207)
//     Symbol                     = ""
//     Exchange                   = ""
//     ClientOrderID_1            = ""
//     OrderType_1                = ORDER_TYPE_UNSET  (0)
//     BuySell_1                  = BUY_SELL_UNSET  (0)
//     Price1_1                   = 0
//     Price2_1                   = 0
//     Quantity_1                 = 0
//     ClientOrderID_2            = ""
//     OrderType_2                = ORDER_TYPE_UNSET  (0)
//     BuySell_2                  = BUY_SELL_UNSET  (0)
//     Price1_2                   = 0
//     Price2_2                   = 0
//     Quantity_2                 = 0
//     TimeInForce                = TIF_UNSET  (0)
//     GoodTillDateTime           = 0
//     TradeAccount               = ""
//     IsAutomatedOrder           = false
//     ParentTriggerClientOrderID = ""
//     FreeFormText               = ""
//     Divisor                    = 0
//     OpenOrClose                = TRADE_UNSET  (0)
//     PartialFillHandling        = PARTIAL_FILL_UNSET  (0)
// }
func (m SubmitNewOCOOrderIntFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SubmitNewOCOOrderIntFixedDefault)
}

// ToBuilder
func (m SubmitNewOCOOrderInt) ToBuilder() SubmitNewOCOOrderIntBuilder {
	return SubmitNewOCOOrderIntBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SubmitNewOCOOrderIntFixed) ToBuilder() SubmitNewOCOOrderIntFixedBuilder {
	return SubmitNewOCOOrderIntFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SubmitNewOCOOrderIntPointer) ToBuilder() SubmitNewOCOOrderIntPointerBuilder {
	return SubmitNewOCOOrderIntPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SubmitNewOCOOrderIntFixedPointer) ToBuilder() SubmitNewOCOOrderIntFixedPointerBuilder {
	return SubmitNewOCOOrderIntFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SubmitNewOCOOrderIntBuilder) Finish() SubmitNewOCOOrderInt {
	return SubmitNewOCOOrderInt{m.VLSBuilder.Finish()}
}

// Finish
func (m SubmitNewOCOOrderIntFixedBuilder) Finish() SubmitNewOCOOrderIntFixed {
	return SubmitNewOCOOrderIntFixed{m.Fixed}
}

// Finish
func (m *SubmitNewOCOOrderIntPointerBuilder) Finish() SubmitNewOCOOrderIntPointer {
	return SubmitNewOCOOrderIntPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SubmitNewOCOOrderIntFixedPointerBuilder) Finish() SubmitNewOCOOrderIntFixedPointer {
	return SubmitNewOCOOrderIntFixedPointer{m.FixedPointer}
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderInt) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntBuilder) Symbol() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m SubmitNewOCOOrderInt) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m SubmitNewOCOOrderInt) ClientOrderID_1() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntBuilder) ClientOrderID_1() string {
	return message.StringVLS(m.Unsafe(), 18, 14)
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
func (m SubmitNewOCOOrderInt) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
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
func (m SubmitNewOCOOrderInt) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 28, 24))
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
func (m SubmitNewOCOOrderInt) Price1_1() int64 {
	return message.Int64VLS(m.Unsafe(), 40, 32)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) Price1_1() int64 {
	return message.Int64VLS(m.Unsafe(), 40, 32)
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
func (m SubmitNewOCOOrderInt) Price2_1() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) Price2_1() int64 {
	return message.Int64VLS(m.Unsafe(), 48, 40)
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
func (m SubmitNewOCOOrderInt) Quantity_1() int64 {
	return message.Int64VLS(m.Unsafe(), 56, 48)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) Quantity_1() int64 {
	return message.Int64VLS(m.Unsafe(), 56, 48)
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
func (m SubmitNewOCOOrderInt) ClientOrderID_2() string {
	return message.StringVLS(m.Unsafe(), 60, 56)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntBuilder) ClientOrderID_2() string {
	return message.StringVLS(m.Unsafe(), 60, 56)
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
func (m SubmitNewOCOOrderInt) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 64, 60))
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32VLS(m.Unsafe(), 64, 60))
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
func (m SubmitNewOCOOrderInt) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 68, 64))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32VLS(m.Unsafe(), 68, 64))
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
func (m SubmitNewOCOOrderInt) Price1_2() int64 {
	return message.Int64VLS(m.Unsafe(), 80, 72)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) Price1_2() int64 {
	return message.Int64VLS(m.Unsafe(), 80, 72)
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
func (m SubmitNewOCOOrderInt) Price2_2() int64 {
	return message.Int64VLS(m.Unsafe(), 88, 80)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) Price2_2() int64 {
	return message.Int64VLS(m.Unsafe(), 88, 80)
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
func (m SubmitNewOCOOrderInt) Quantity_2() int64 {
	return message.Int64VLS(m.Unsafe(), 96, 88)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) Quantity_2() int64 {
	return message.Int64VLS(m.Unsafe(), 96, 88)
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
func (m SubmitNewOCOOrderInt) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 112, 104))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64VLS(m.Unsafe(), 112, 104))
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
func (m SubmitNewOCOOrderInt) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 116, 112)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 116, 112)
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
func (m SubmitNewOCOOrderInt) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 117, 116)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Unsafe(), 117, 116)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntPointer) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 117, 116)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolVLS(m.Ptr, 117, 116)
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
func (m SubmitNewOCOOrderInt) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 132, 128)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntBuilder) Divisor() float32 {
	return message.Float32VLS(m.Unsafe(), 132, 128)
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
func (m SubmitNewOCOOrderInt) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 136, 132))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32VLS(m.Unsafe(), 136, 132))
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

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixed) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
}

// Symbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedBuilder) Symbol() string {
	return message.StringFixed(m.Unsafe(), 68, 4)
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
func (m SubmitNewOCOOrderIntFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
}

// Exchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 84, 68)
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
func (m SubmitNewOCOOrderIntFixed) ClientOrderID_1() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
}

// ClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) ClientOrderID_1() string {
	return message.StringFixed(m.Unsafe(), 116, 84)
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
func (m SubmitNewOCOOrderIntFixed) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 120, 116))
}

// OrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) OrderType_1() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 120, 116))
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
func (m SubmitNewOCOOrderIntFixed) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 124, 120))
}

// BuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) BuySell_1() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 124, 120))
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
func (m SubmitNewOCOOrderIntFixed) Price1_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 136, 128)
}

// Price1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) Price1_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 136, 128)
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
func (m SubmitNewOCOOrderIntFixed) Price2_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 144, 136)
}

// Price2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) Price2_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 144, 136)
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
func (m SubmitNewOCOOrderIntFixed) Quantity_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 152, 144)
}

// Quantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) Quantity_1() int64 {
	return message.Int64Fixed(m.Unsafe(), 152, 144)
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
func (m SubmitNewOCOOrderIntFixed) ClientOrderID_2() string {
	return message.StringFixed(m.Unsafe(), 184, 152)
}

// ClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) ClientOrderID_2() string {
	return message.StringFixed(m.Unsafe(), 184, 152)
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
func (m SubmitNewOCOOrderIntFixed) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 188, 184))
}

// OrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) OrderType_2() OrderTypeEnum {
	return OrderTypeEnum(message.Int32Fixed(m.Unsafe(), 188, 184))
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
func (m SubmitNewOCOOrderIntFixed) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 192, 188))
}

// BuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) BuySell_2() BuySellEnum {
	return BuySellEnum(message.Int32Fixed(m.Unsafe(), 192, 188))
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
func (m SubmitNewOCOOrderIntFixed) Price1_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 200, 192)
}

// Price1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) Price1_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 200, 192)
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
func (m SubmitNewOCOOrderIntFixed) Price2_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 208, 200)
}

// Price2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) Price2_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 208, 200)
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
func (m SubmitNewOCOOrderIntFixed) Quantity_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 216, 208)
}

// Quantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) Quantity_2() int64 {
	return message.Int64Fixed(m.Unsafe(), 216, 208)
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
func (m SubmitNewOCOOrderIntFixed) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 232, 224))
}

// GoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixedBuilder) GoodTillDateTime() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 232, 224))
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
func (m SubmitNewOCOOrderIntFixed) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 264, 232)
}

// TradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedBuilder) TradeAccount() string {
	return message.StringFixed(m.Unsafe(), 264, 232)
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
func (m SubmitNewOCOOrderIntFixed) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 265, 264)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Unsafe(), 265, 264)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedPointer) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 265, 264)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) IsAutomatedOrder() bool {
	return message.BoolFixed(m.Ptr, 265, 264)
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
func (m SubmitNewOCOOrderIntFixed) Divisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 352, 348)
}

// Divisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixedBuilder) Divisor() float32 {
	return message.Float32Fixed(m.Unsafe(), 352, 348)
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
func (m SubmitNewOCOOrderIntFixed) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 356, 352))
}

// OpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) OpenOrClose() OpenCloseTradeEnum {
	return OpenCloseTradeEnum(message.Int32Fixed(m.Unsafe(), 356, 352))
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
func (m *SubmitNewOCOOrderIntBuilder) SetSymbol(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetSymbol The symbol for the order.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetSymbol(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewOCOOrderIntBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetExchange The optional exchange for the symbol.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m *SubmitNewOCOOrderIntBuilder) SetClientOrderID_1(value string) {
	message.SetStringVLS(&m.VLSBuilder, 18, 14, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetClientOrderID_1(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 18, 14, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointerBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) SetPrice1_1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 40, 32, value)
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice1_1(value int64) {
	message.SetInt64VLS(m.Ptr, 40, 32, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) SetPrice2_1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 48, 40, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice2_1(value int64) {
	message.SetInt64VLS(m.Ptr, 48, 40, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) SetQuantity_1(value int64) {
	message.SetInt64VLS(m.Unsafe(), 56, 48, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntPointerBuilder) SetQuantity_1(value int64) {
	message.SetInt64VLS(m.Ptr, 56, 48, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m *SubmitNewOCOOrderIntBuilder) SetClientOrderID_2(value string) {
	message.SetStringVLS(&m.VLSBuilder, 60, 56, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetClientOrderID_2(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 60, 56, value)
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 64, 60, int32(value))
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32VLS(m.Ptr, 64, 60, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32VLS(m.Unsafe(), 68, 64, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntPointerBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32VLS(m.Ptr, 68, 64, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntBuilder) SetPrice1_2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 80, 72, value)
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice1_2(value int64) {
	message.SetInt64VLS(m.Ptr, 80, 72, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntBuilder) SetPrice2_2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 88, 80, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPrice2_2(value int64) {
	message.SetInt64VLS(m.Ptr, 88, 80, value)
}

// SetQuantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntBuilder) SetQuantity_2(value int64) {
	message.SetInt64VLS(m.Unsafe(), 96, 88, value)
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
func (m SubmitNewOCOOrderIntBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32VLS(m.Unsafe(), 100, 96, int32(value))
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
func (m SubmitNewOCOOrderIntBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Unsafe(), 112, 104, int64(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64VLS(m.Ptr, 112, 104, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m *SubmitNewOCOOrderIntBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 116, 112, value)
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m *SubmitNewOCOOrderIntPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 116, 112, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Unsafe(), 117, 116, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolVLS(m.Ptr, 117, 116, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m *SubmitNewOCOOrderIntBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 122, 118, value)
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
func (m *SubmitNewOCOOrderIntBuilder) SetFreeFormText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 126, 122, value)
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
func (m SubmitNewOCOOrderIntBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Unsafe(), 132, 128, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32VLS(m.Ptr, 132, 128, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Unsafe(), 136, 132, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32VLS(m.Ptr, 136, 132, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8VLS(m.Unsafe(), 137, 136, int8(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntPointerBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8VLS(m.Ptr, 137, 136, int8(value))
}

// SetSymbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Unsafe(), 68, 4, value)
}

// SetSymbol The symbol for the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetSymbol(value string) {
	message.SetStringFixed(m.Ptr, 68, 4, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 84, 68, value)
}

// SetExchange The optional exchange for the symbol.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 84, 68, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetClientOrderID_1(value string) {
	message.SetStringFixed(m.Unsafe(), 116, 84, value)
}

// SetClientOrderID_1 The Client supplied order identifier for the first order. The Server will
// remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetClientOrderID_1(value string) {
	message.SetStringFixed(m.Ptr, 116, 84, value)
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 120, 116, int32(value))
}

// SetOrderType_1 The order type for the first order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetOrderType_1(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 120, 116, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 124, 120, int32(value))
}

// SetBuySell_1 The side for the first order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetBuySell_1(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 124, 120, int32(value))
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice1_1(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 136, 128, value)
}

// SetPrice1_1 This is the price of the first order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice1_1(value int64) {
	message.SetInt64Fixed(m.Ptr, 136, 128, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice2_1(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 144, 136, value)
}

// SetPrice2_1 This is the second price for the first order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice2_1(value int64) {
	message.SetInt64Fixed(m.Ptr, 144, 136, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) SetQuantity_1(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 152, 144, value)
}

// SetQuantity_1 The quantity of the first order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetQuantity_1(value int64) {
	message.SetInt64Fixed(m.Ptr, 152, 144, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetClientOrderID_2(value string) {
	message.SetStringFixed(m.Unsafe(), 184, 152, value)
}

// SetClientOrderID_2 The Client supplied order identifier for the second order. The Server
// will remember this for the life of the order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetClientOrderID_2(value string) {
	message.SetStringFixed(m.Ptr, 184, 152, value)
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 188, 184, int32(value))
}

// SetOrderType_2 The order type of the second order. For list of order types, refer to
// OrderTypeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetOrderType_2(value OrderTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 188, 184, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32Fixed(m.Unsafe(), 192, 188, int32(value))
}

// SetBuySell_2 The side of the second order. Either Buy or Sell.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetBuySell_2(value BuySellEnum) {
	message.SetInt32Fixed(m.Ptr, 192, 188, int32(value))
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice1_2(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 200, 192, value)
}

// SetPrice1_2 This is the price of the second order. This is the limit price for a Limit
// order, the stop price for a Stop order, or the trigger price for a Market
// if Touched order.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice1_2(value int64) {
	message.SetInt64Fixed(m.Ptr, 200, 192, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPrice2_2(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 208, 200, value)
}

// SetPrice2_2 This is the second price for the second order. For a Stop-Limit order,
// this is the limit price. This only applies to Stop-Limit orders.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPrice2_2(value int64) {
	message.SetInt64Fixed(m.Ptr, 208, 200, value)
}

// SetQuantity_2 The quantity of the second order. The exact meaning of this will be specified
// by the Server implementation.
func (m SubmitNewOCOOrderIntFixedBuilder) SetQuantity_2(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 216, 208, value)
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
func (m SubmitNewOCOOrderIntFixedBuilder) SetTimeInForce(value TimeInForceEnum) {
	message.SetInt32Fixed(m.Unsafe(), 220, 216, int32(value))
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
func (m SubmitNewOCOOrderIntFixedBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 232, 224, int64(value))
}

// SetGoodTillDateTime In the case of when the TimeInForce is TIF_GOOD_TILL_DATE_TIME, this specifies
// the expiration Date-Time of the order. This applies to both of the orders
// in the OCO pair.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetGoodTillDateTime(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 232, 224, int64(value))
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Unsafe(), 264, 232, value)
}

// SetTradeAccount This is the trade account as a text string that the orders belong to.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetTradeAccount(value string) {
	message.SetStringFixed(m.Ptr, 264, 232, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Unsafe(), 265, 264, value)
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetIsAutomatedOrder(value bool) {
	message.SetBoolFixed(m.Ptr, 265, 264, value)
}

// SetParentTriggerClientOrderID Optional: This field supports the submission of an OCO order pair which
// has a parent. This is known as a Bracket order.
//
// For complete documentation for bracket orders, refer to Bracket Order
// Procedures.
func (m SubmitNewOCOOrderIntFixedBuilder) SetParentTriggerClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 297, 265, value)
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
func (m SubmitNewOCOOrderIntFixedBuilder) SetFreeFormText(value string) {
	message.SetStringFixed(m.Unsafe(), 345, 297, value)
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
func (m SubmitNewOCOOrderIntFixedBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Unsafe(), 352, 348, value)
}

// SetDivisor This is the FloatToIntPriceMultiplier value provided in the SecurityDefinitionResponse
// message for the Symbol being traded. The Server needs to divide the Price1_*
// and Price2_* fields by this Divisor to get the prices as float values.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetDivisor(value float32) {
	message.SetFloat32Fixed(m.Ptr, 352, 348, value)
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 356, 352, int32(value))
}

// SetOpenOrClose For the description for this field, refer to OpenCloseTradeEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetOpenOrClose(value OpenCloseTradeEnum) {
	message.SetInt32Fixed(m.Ptr, 356, 352, int32(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixedBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8Fixed(m.Unsafe(), 357, 356, int8(value))
}

// SetPartialFillHandling Specifies how partial fills should be handled when when one of the orders
// in the OCO order set partially fills.
//
// For the possible values, refer to PartialFillHandlingEnum.
func (m SubmitNewOCOOrderIntFixedPointerBuilder) SetPartialFillHandling(value PartialFillHandlingEnum) {
	message.SetInt8Fixed(m.Ptr, 357, 356, int8(value))
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