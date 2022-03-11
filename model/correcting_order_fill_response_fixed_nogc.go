package model

import (
	"github.com/moontrade/dtc-go/message"
)

type CorrectingOrderFillResponseFixedPointer struct {
	message.FixedPointer
}

type CorrectingOrderFillResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocCorrectingOrderFillResponseFixed() CorrectingOrderFillResponseFixedPointerBuilder {
	a := CorrectingOrderFillResponseFixedPointerBuilder{message.AllocFixed(294)}
	a.Ptr.SetBytes(0, _CorrectingOrderFillResponseFixedDefault)
	return a
}

func AllocCorrectingOrderFillResponseFixedFrom(b []byte) CorrectingOrderFillResponseFixedPointer {
	return CorrectingOrderFillResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size          = CorrectingOrderFillResponseFixedSize  (294)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = 0
// }
func (m CorrectingOrderFillResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CorrectingOrderFillResponseFixedDefault)
}

// ToBuilder
func (m CorrectingOrderFillResponseFixedPointer) ToBuilder() CorrectingOrderFillResponseFixedPointerBuilder {
	return CorrectingOrderFillResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *CorrectingOrderFillResponseFixedPointerBuilder) Finish() CorrectingOrderFillResponseFixedPointer {
	return CorrectingOrderFillResponseFixedPointer{m.FixedPointer}
}

// ClientOrderID
func (m CorrectingOrderFillResponseFixedPointer) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 36, 4)
}

// ClientOrderID
func (m CorrectingOrderFillResponseFixedPointerBuilder) ClientOrderID() string {
	return message.StringFixed(m.Ptr, 36, 4)
}

// ResultText
func (m CorrectingOrderFillResponseFixedPointer) ResultText() string {
	return message.StringFixed(m.Ptr, 292, 36)
}

// ResultText
func (m CorrectingOrderFillResponseFixedPointerBuilder) ResultText() string {
	return message.StringFixed(m.Ptr, 292, 36)
}

// IsError
func (m CorrectingOrderFillResponseFixedPointer) IsError() uint8 {
	return message.Uint8Fixed(m.Ptr, 293, 292)
}

// IsError
func (m CorrectingOrderFillResponseFixedPointerBuilder) IsError() uint8 {
	return message.Uint8Fixed(m.Ptr, 293, 292)
}

// SetClientOrderID
func (m CorrectingOrderFillResponseFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 36, 4, value)
}

// SetResultText
func (m CorrectingOrderFillResponseFixedPointerBuilder) SetResultText(value string) {
	message.SetStringFixed(m.Ptr, 292, 36, value)
}

// SetIsError
func (m CorrectingOrderFillResponseFixedPointerBuilder) SetIsError(value uint8) {
	message.SetUint8Fixed(m.Ptr, 293, 292, value)
}

// Copy
func (m CorrectingOrderFillResponseFixedPointer) Copy(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// Copy
func (m CorrectingOrderFillResponseFixedPointerBuilder) Copy(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponseFixedPointer) CopyPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponseFixedPointerBuilder) CopyPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponseFixedPointer) CopyTo(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponseFixedPointerBuilder) CopyTo(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponseFixedPointer) CopyToPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponseFixedPointerBuilder) CopyToPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}
