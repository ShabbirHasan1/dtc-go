package model

import (
	"github.com/moontrade/dtc-go/message"
)

type CorrectingOrderFillResponsePointer struct {
	message.VLSPointer
}

type CorrectingOrderFillResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocCorrectingOrderFillResponse() CorrectingOrderFillResponsePointerBuilder {
	a := CorrectingOrderFillResponsePointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _CorrectingOrderFillResponseDefault)
	return a
}

func AllocCorrectingOrderFillResponseFrom(b []byte) CorrectingOrderFillResponsePointer {
	return CorrectingOrderFillResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size          = CorrectingOrderFillResponseSize  (16)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     BaseSize      = CorrectingOrderFillResponseSize  (16)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = 0
// }
func (m CorrectingOrderFillResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CorrectingOrderFillResponseDefault)
}

// ToBuilder
func (m CorrectingOrderFillResponsePointer) ToBuilder() CorrectingOrderFillResponsePointerBuilder {
	return CorrectingOrderFillResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *CorrectingOrderFillResponsePointerBuilder) Finish() CorrectingOrderFillResponsePointer {
	return CorrectingOrderFillResponsePointer{m.VLSPointerBuilder.Finish()}
}

// ClientOrderID
func (m CorrectingOrderFillResponsePointer) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// ClientOrderID
func (m CorrectingOrderFillResponsePointerBuilder) ClientOrderID() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// ResultText
func (m CorrectingOrderFillResponsePointer) ResultText() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// ResultText
func (m CorrectingOrderFillResponsePointerBuilder) ResultText() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// IsError
func (m CorrectingOrderFillResponsePointer) IsError() uint8 {
	return message.Uint8VLS(m.Ptr, 15, 14)
}

// IsError
func (m CorrectingOrderFillResponsePointerBuilder) IsError() uint8 {
	return message.Uint8VLS(m.Ptr, 15, 14)
}

// SetClientOrderID
func (m *CorrectingOrderFillResponsePointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetResultText
func (m *CorrectingOrderFillResponsePointerBuilder) SetResultText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetIsError
func (m CorrectingOrderFillResponsePointerBuilder) SetIsError(value uint8) {
	message.SetUint8VLS(m.Ptr, 15, 14, value)
}

// Copy
func (m CorrectingOrderFillResponsePointer) Copy(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// Copy
func (m CorrectingOrderFillResponsePointerBuilder) Copy(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponsePointer) CopyPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponsePointerBuilder) CopyPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponsePointer) CopyTo(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponsePointerBuilder) CopyTo(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponsePointer) CopyToPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponsePointerBuilder) CopyToPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}
