package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = CorrectingOrderFillResponseFixedSize  (294)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = 0
// }
var _CorrectingOrderFillResponseFixedDefault = []byte{38, 1, 54, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CorrectingOrderFillResponseFixedSize = 294

type CorrectingOrderFillResponseFixed struct {
	message.Fixed
}

type CorrectingOrderFillResponseFixedBuilder struct {
	message.Fixed
}

func NewCorrectingOrderFillResponseFixedFrom(b []byte) CorrectingOrderFillResponseFixed {
	return CorrectingOrderFillResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapCorrectingOrderFillResponseFixed(b []byte) CorrectingOrderFillResponseFixed {
	return CorrectingOrderFillResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewCorrectingOrderFillResponseFixed() CorrectingOrderFillResponseFixedBuilder {
	a := CorrectingOrderFillResponseFixedBuilder{message.NewFixed(294)}
	a.Unsafe().SetBytes(0, _CorrectingOrderFillResponseFixedDefault)
	return a
}

// Clear
// {
//     Size          = CorrectingOrderFillResponseFixedSize  (294)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = 0
// }
func (m CorrectingOrderFillResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CorrectingOrderFillResponseFixedDefault)
}

// ToBuilder
func (m CorrectingOrderFillResponseFixed) ToBuilder() CorrectingOrderFillResponseFixedBuilder {
	return CorrectingOrderFillResponseFixedBuilder{m.Fixed}
}

// Finish
func (m CorrectingOrderFillResponseFixedBuilder) Finish() CorrectingOrderFillResponseFixed {
	return CorrectingOrderFillResponseFixed{m.Fixed}
}

// ClientOrderID
func (m CorrectingOrderFillResponseFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 36, 4)
}

// ClientOrderID
func (m CorrectingOrderFillResponseFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 36, 4)
}

// ResultText
func (m CorrectingOrderFillResponseFixed) ResultText() string {
	return message.StringFixed(m.Unsafe(), 292, 36)
}

// ResultText
func (m CorrectingOrderFillResponseFixedBuilder) ResultText() string {
	return message.StringFixed(m.Unsafe(), 292, 36)
}

// IsError
func (m CorrectingOrderFillResponseFixed) IsError() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 293, 292)
}

// IsError
func (m CorrectingOrderFillResponseFixedBuilder) IsError() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 293, 292)
}

// SetClientOrderID
func (m CorrectingOrderFillResponseFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 36, 4, value)
}

// SetResultText
func (m CorrectingOrderFillResponseFixedBuilder) SetResultText(value string) {
	message.SetStringFixed(m.Unsafe(), 292, 36, value)
}

// SetIsError
func (m CorrectingOrderFillResponseFixedBuilder) SetIsError(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 293, 292, value)
}

// Copy
func (m CorrectingOrderFillResponseFixed) Copy(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// Copy
func (m CorrectingOrderFillResponseFixedBuilder) Copy(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponseFixed) CopyPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponseFixedBuilder) CopyPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponseFixed) CopyToPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponseFixedBuilder) CopyToPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponseFixed) CopyTo(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponseFixedBuilder) CopyTo(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}
