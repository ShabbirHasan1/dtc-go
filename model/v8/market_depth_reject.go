// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

import (
	"github.com/moontrade/dtc-go/message"
)

const MarketDepthRejectSize = 16

const MarketDepthRejectFixedSize = 104

//     Size        uint16  = MarketDepthRejectSize  (16)
//     Type        uint16  = MARKET_DEPTH_REJECT  (121)
//     BaseSize    uint16  = MarketDepthRejectSize  (16)
//     SymbolID    uint32  = 0
//     RejectText  string  = ""
var _MarketDepthRejectDefault = []byte{16, 0, 121, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = MarketDepthRejectFixedSize  (104)
//     Type        uint16      = MARKET_DEPTH_REJECT  (121)
//     SymbolID    uint32      = 0
//     RejectText  string[96]  = ""
var _MarketDepthRejectFixedDefault = []byte{104, 0, 121, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// MarketDepthReject The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthReject struct {
	message.VLS
}

// MarketDepthRejectBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectBuilder struct {
	message.VLSBuilder
}

// MarketDepthRejectFixed The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixed struct {
	message.Fixed
}

// MarketDepthRejectFixedBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixedBuilder struct {
	message.Fixed
}

// MarketDepthRejectPointer The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectPointer struct {
	message.VLSPointer
}

// MarketDepthRejectPointerBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

// MarketDepthRejectFixedPointer The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixedPointer struct {
	message.FixedPointer
}

// MarketDepthRejectFixedPointerBuilder The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
type MarketDepthRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewMarketDepthRejectFrom(b []byte) MarketDepthReject {
	return MarketDepthReject{VLS: message.NewVLS(b)}
}

func WrapMarketDepthReject(b []byte) MarketDepthReject {
	return MarketDepthReject{VLS: message.WrapVLS(b)}
}

func NewMarketDepthReject() MarketDepthRejectBuilder {
	return MarketDepthRejectBuilder{message.NewVLSBuilder(_MarketDepthRejectDefault)}
}

func NewMarketDepthRejectFixedFrom(b []byte) MarketDepthRejectFixed {
	return MarketDepthRejectFixed{Fixed: message.NewFixed(b)}
}

func WrapMarketDepthRejectFixed(b []byte) MarketDepthRejectFixed {
	return MarketDepthRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewMarketDepthRejectFixed() MarketDepthRejectFixedBuilder {
	return MarketDepthRejectFixedBuilder{message.NewFixed(_MarketDepthRejectFixedDefault)}
}

func AllocMarketDepthReject() MarketDepthRejectPointerBuilder {
	return MarketDepthRejectPointerBuilder{message.AllocVLSBuilder(_MarketDepthRejectDefault)}
}

func AllocMarketDepthRejectFrom(b []byte) MarketDepthRejectPointer {
	return MarketDepthRejectPointer{VLSPointer: message.AllocVLS(b)}
}

func AllocMarketDepthRejectFixed() MarketDepthRejectFixedPointerBuilder {
	return MarketDepthRejectFixedPointerBuilder{message.AllocFixed(_MarketDepthRejectFixedDefault)}
}

func AllocMarketDepthRejectFixedFrom(b []byte) MarketDepthRejectFixedPointer {
	return MarketDepthRejectFixedPointer{FixedPointer: message.AllocFixed(b)}
}

// Clear
//     Size        uint16  = MarketDepthRejectSize  (16)
//     Type        uint16  = MARKET_DEPTH_REJECT  (121)
//     BaseSize    uint16  = MarketDepthRejectSize  (16)
//     SymbolID    uint32  = 0
//     RejectText  string  = ""
func (m MarketDepthRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthRejectDefault)
}

// Clear
//     Size        uint16      = MarketDepthRejectFixedSize  (104)
//     Type        uint16      = MARKET_DEPTH_REJECT  (121)
//     SymbolID    uint32      = 0
//     RejectText  string[96]  = ""
func (m MarketDepthRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _MarketDepthRejectFixedDefault)
}

// Clear
//     Size        uint16  = MarketDepthRejectSize  (16)
//     Type        uint16  = MARKET_DEPTH_REJECT  (121)
//     BaseSize    uint16  = MarketDepthRejectSize  (16)
//     SymbolID    uint32  = 0
//     RejectText  string  = ""
func (m MarketDepthRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthRejectDefault)
}

// Clear
//     Size        uint16      = MarketDepthRejectFixedSize  (104)
//     Type        uint16      = MARKET_DEPTH_REJECT  (121)
//     SymbolID    uint32      = 0
//     RejectText  string[96]  = ""
func (m MarketDepthRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _MarketDepthRejectFixedDefault)
}

// ToBuilder
func (m MarketDepthReject) ToBuilder() MarketDepthRejectBuilder {
	return MarketDepthRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m MarketDepthRejectFixed) ToBuilder() MarketDepthRejectFixedBuilder {
	return MarketDepthRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m MarketDepthRejectPointer) ToBuilder() MarketDepthRejectPointerBuilder {
	return MarketDepthRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m MarketDepthRejectFixedPointer) ToBuilder() MarketDepthRejectFixedPointerBuilder {
	return MarketDepthRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m MarketDepthRejectBuilder) Finish() MarketDepthReject {
	return MarketDepthReject{m.VLSBuilder.Finish()}
}

// Finish
func (m MarketDepthRejectFixedBuilder) Finish() MarketDepthRejectFixed {
	return MarketDepthRejectFixed{m.Fixed}
}

// Finish
func (m *MarketDepthRejectPointerBuilder) Finish() MarketDepthRejectPointer {
	return MarketDepthRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *MarketDepthRejectFixedPointerBuilder) Finish() MarketDepthRejectFixedPointer {
	return MarketDepthRejectFixedPointer{m.FixedPointer}
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthReject) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 12, 8)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectPointer) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectPointerBuilder) SymbolID() uint32 {
	return message.Uint32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixed) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedPointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedPointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 12, 8, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDepthRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m *MarketDepthRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID This is the same SymbolID sent by the Client in the MarketDepthRequest
// message which corresponds to the Symbol that the data in this message
// is for.
func (m MarketDepthRejectFixedPointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Free-form text explaining the reason for the reject.
func (m MarketDepthRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}