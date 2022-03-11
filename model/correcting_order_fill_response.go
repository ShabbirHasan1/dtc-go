package model

import (
	"github.com/moontrade/dtc-go/message"
)

const CorrectingOrderFillResponseSize = 16

const CorrectingOrderFillResponseFixedSize = 294

// {
//     Size          = CorrectingOrderFillResponseSize  (16)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     BaseSize      = CorrectingOrderFillResponseSize  (16)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = false
// }
var _CorrectingOrderFillResponseDefault = []byte{16, 0, 54, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size          = CorrectingOrderFillResponseFixedSize  (294)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = false
// }
var _CorrectingOrderFillResponseFixedDefault = []byte{38, 1, 54, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type CorrectingOrderFillResponse struct {
	message.VLS
}

type CorrectingOrderFillResponseBuilder struct {
	message.VLSBuilder
}

type CorrectingOrderFillResponseFixed struct {
	message.Fixed
}

type CorrectingOrderFillResponseFixedBuilder struct {
	message.Fixed
}

type CorrectingOrderFillResponsePointer struct {
	message.VLSPointer
}

type CorrectingOrderFillResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

type CorrectingOrderFillResponseFixedPointer struct {
	message.FixedPointer
}

type CorrectingOrderFillResponseFixedPointerBuilder struct {
	message.FixedPointer
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

func AllocCorrectingOrderFillResponse() CorrectingOrderFillResponsePointerBuilder {
	a := CorrectingOrderFillResponsePointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _CorrectingOrderFillResponseDefault)
	return a
}

func AllocCorrectingOrderFillResponseFrom(b []byte) CorrectingOrderFillResponsePointer {
	return CorrectingOrderFillResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size          = CorrectingOrderFillResponseSize  (16)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     BaseSize      = CorrectingOrderFillResponseSize  (16)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = false
// }
func (m CorrectingOrderFillResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CorrectingOrderFillResponseDefault)
}

// Clear
// {
//     Size          = CorrectingOrderFillResponseFixedSize  (294)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = false
// }
func (m CorrectingOrderFillResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CorrectingOrderFillResponseFixedDefault)
}

// Clear
// {
//     Size          = CorrectingOrderFillResponseSize  (16)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     BaseSize      = CorrectingOrderFillResponseSize  (16)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = false
// }
func (m CorrectingOrderFillResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CorrectingOrderFillResponseDefault)
}

// Clear
// {
//     Size          = CorrectingOrderFillResponseFixedSize  (294)
//     Type          = CORRECTING_ORDER_FILL_RESPONSE  (310)
//     ClientOrderID = ""
//     ResultText    = ""
//     IsError       = false
// }
func (m CorrectingOrderFillResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CorrectingOrderFillResponseFixedDefault)
}

// ToBuilder
func (m CorrectingOrderFillResponse) ToBuilder() CorrectingOrderFillResponseBuilder {
	return CorrectingOrderFillResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m CorrectingOrderFillResponseFixed) ToBuilder() CorrectingOrderFillResponseFixedBuilder {
	return CorrectingOrderFillResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m CorrectingOrderFillResponsePointer) ToBuilder() CorrectingOrderFillResponsePointerBuilder {
	return CorrectingOrderFillResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m CorrectingOrderFillResponseFixedPointer) ToBuilder() CorrectingOrderFillResponseFixedPointerBuilder {
	return CorrectingOrderFillResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m CorrectingOrderFillResponseBuilder) Finish() CorrectingOrderFillResponse {
	return CorrectingOrderFillResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m CorrectingOrderFillResponseFixedBuilder) Finish() CorrectingOrderFillResponseFixed {
	return CorrectingOrderFillResponseFixed{m.Fixed}
}

// Finish
func (m *CorrectingOrderFillResponsePointerBuilder) Finish() CorrectingOrderFillResponsePointer {
	return CorrectingOrderFillResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *CorrectingOrderFillResponseFixedPointerBuilder) Finish() CorrectingOrderFillResponseFixedPointer {
	return CorrectingOrderFillResponseFixedPointer{m.FixedPointer}
}

// ClientOrderID
func (m CorrectingOrderFillResponse) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// ClientOrderID
func (m CorrectingOrderFillResponseBuilder) ClientOrderID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m CorrectingOrderFillResponse) ResultText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// ResultText
func (m CorrectingOrderFillResponseBuilder) ResultText() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m CorrectingOrderFillResponse) IsError() bool {
	return message.BoolVLS(m.Unsafe(), 15, 14)
}

// IsError
func (m CorrectingOrderFillResponseBuilder) IsError() bool {
	return message.BoolVLS(m.Unsafe(), 15, 14)
}

// IsError
func (m CorrectingOrderFillResponsePointer) IsError() bool {
	return message.BoolVLS(m.Ptr, 15, 14)
}

// IsError
func (m CorrectingOrderFillResponsePointerBuilder) IsError() bool {
	return message.BoolVLS(m.Ptr, 15, 14)
}

// ClientOrderID
func (m CorrectingOrderFillResponseFixed) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 36, 4)
}

// ClientOrderID
func (m CorrectingOrderFillResponseFixedBuilder) ClientOrderID() string {
	return message.StringFixed(m.Unsafe(), 36, 4)
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
func (m CorrectingOrderFillResponseFixed) ResultText() string {
	return message.StringFixed(m.Unsafe(), 292, 36)
}

// ResultText
func (m CorrectingOrderFillResponseFixedBuilder) ResultText() string {
	return message.StringFixed(m.Unsafe(), 292, 36)
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
func (m CorrectingOrderFillResponseFixed) IsError() bool {
	return message.BoolFixed(m.Unsafe(), 293, 292)
}

// IsError
func (m CorrectingOrderFillResponseFixedBuilder) IsError() bool {
	return message.BoolFixed(m.Unsafe(), 293, 292)
}

// IsError
func (m CorrectingOrderFillResponseFixedPointer) IsError() bool {
	return message.BoolFixed(m.Ptr, 293, 292)
}

// IsError
func (m CorrectingOrderFillResponseFixedPointerBuilder) IsError() bool {
	return message.BoolFixed(m.Ptr, 293, 292)
}

// SetClientOrderID
func (m *CorrectingOrderFillResponseBuilder) SetClientOrderID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetClientOrderID
func (m *CorrectingOrderFillResponsePointerBuilder) SetClientOrderID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetResultText
func (m *CorrectingOrderFillResponseBuilder) SetResultText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetResultText
func (m *CorrectingOrderFillResponsePointerBuilder) SetResultText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetIsError
func (m CorrectingOrderFillResponseBuilder) SetIsError(value bool) {
	message.SetBoolVLS(m.Unsafe(), 15, 14, value)
}

// SetIsError
func (m CorrectingOrderFillResponsePointerBuilder) SetIsError(value bool) {
	message.SetBoolVLS(m.Ptr, 15, 14, value)
}

// SetClientOrderID
func (m CorrectingOrderFillResponseFixedBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Unsafe(), 36, 4, value)
}

// SetClientOrderID
func (m CorrectingOrderFillResponseFixedPointerBuilder) SetClientOrderID(value string) {
	message.SetStringFixed(m.Ptr, 36, 4, value)
}

// SetResultText
func (m CorrectingOrderFillResponseFixedBuilder) SetResultText(value string) {
	message.SetStringFixed(m.Unsafe(), 292, 36, value)
}

// SetResultText
func (m CorrectingOrderFillResponseFixedPointerBuilder) SetResultText(value string) {
	message.SetStringFixed(m.Ptr, 292, 36, value)
}

// SetIsError
func (m CorrectingOrderFillResponseFixedBuilder) SetIsError(value bool) {
	message.SetBoolFixed(m.Unsafe(), 293, 292, value)
}

// SetIsError
func (m CorrectingOrderFillResponseFixedPointerBuilder) SetIsError(value bool) {
	message.SetBoolFixed(m.Ptr, 293, 292, value)
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