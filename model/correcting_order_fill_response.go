package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size          = CorrectingOrderFillResponseSize  (16)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     BaseSize      = CorrectingOrderFillResponseSize  (16)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = 0
// }
var _CorrectingOrderFillResponseDefault = []byte{16, 0, 54, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const CorrectingOrderFillResponseSize = 16

type CorrectingOrderFillResponse struct {
	message.VLS
}

type CorrectingOrderFillResponseBuilder struct {
	message.VLSBuilder
}

func NewCorrectingOrderFillResponseFrom(b []byte) CorrectingOrderFillResponse {
	return CorrectingOrderFillResponse{VLS: message.NewVLSFrom(b)}
}

func WrapCorrectingOrderFillResponse(b []byte) CorrectingOrderFillResponse {
	return CorrectingOrderFillResponse{VLS: message.WrapVLS(b)}
}

func NewCorrectingOrderFillResponse() CorrectingOrderFillResponseBuilder {
	a := CorrectingOrderFillResponseBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _CorrectingOrderFillResponseDefault)
	return a
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
func (m CorrectingOrderFillResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CorrectingOrderFillResponseDefault)
}

// ToBuilder
func (m CorrectingOrderFillResponse) ToBuilder() CorrectingOrderFillResponseBuilder {
	return CorrectingOrderFillResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m CorrectingOrderFillResponseBuilder) Finish() CorrectingOrderFillResponse {
	return CorrectingOrderFillResponse{m.VLSBuilder.Finish()}
}

// ClientOrderID
func (m CorrectingOrderFillResponse) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ClientOrderID
func (m CorrectingOrderFillResponseBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ResultText
func (m CorrectingOrderFillResponse) ResultText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ResultText
func (m CorrectingOrderFillResponseBuilder) ResultText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// IsError
func (m CorrectingOrderFillResponse) IsError() uint8 {
	return message.Uint8VLS(m.Unsafe(), 15, 14)
}

// IsError
func (m CorrectingOrderFillResponseBuilder) IsError() uint8 {
	return message.Uint8VLS(m.Unsafe(), 15, 14)
}

// SetClientOrderID
func (m *CorrectingOrderFillResponseBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetResultText
func (m *CorrectingOrderFillResponseBuilder) SetResultText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetIsError
func (m CorrectingOrderFillResponseBuilder) SetIsError(value uint8) {
	message.SetUint8VLS(m.Unsafe(), 15, 14, value)
}

// Copy
func (m CorrectingOrderFillResponse) Copy(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// Copy
func (m CorrectingOrderFillResponseBuilder) Copy(to CorrectingOrderFillResponseBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponse) CopyPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyPointer
func (m CorrectingOrderFillResponseBuilder) CopyPointer(to CorrectingOrderFillResponsePointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponse) CopyToPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyToPointer
func (m CorrectingOrderFillResponseBuilder) CopyToPointer(to CorrectingOrderFillResponseFixedPointerBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponse) CopyTo(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}

// CopyTo
func (m CorrectingOrderFillResponseBuilder) CopyTo(to CorrectingOrderFillResponseFixedBuilder) {
	to.SetClientOrderID(m.ClientOrderID())
	to.SetResultText(m.ResultText())
	to.SetIsError(m.IsError())
}
