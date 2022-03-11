package model

import (
	"github.com/moontrade/dtc-go/message"
)

const CurrentPositionsRejectSize = 16

const CurrentPositionsRejectFixedSize = 104

// {
//     Size       = CurrentPositionsRejectSize  (16)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     BaseSize   = CurrentPositionsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
var _CurrentPositionsRejectDefault = []byte{16, 0, 51, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size       = CurrentPositionsRejectFixedSize  (104)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     RequestID  = 0
//     RejectText = ""
// }
var _CurrentPositionsRejectFixedDefault = []byte{104, 0, 51, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// CurrentPositionsReject If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsReject struct {
	message.VLS
}

// CurrentPositionsRejectBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectBuilder struct {
	message.VLSBuilder
}

// CurrentPositionsRejectFixed If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixed struct {
	message.Fixed
}

// CurrentPositionsRejectFixedBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixedBuilder struct {
	message.Fixed
}

// CurrentPositionsRejectPointer If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectPointer struct {
	message.VLSPointer
}

// CurrentPositionsRejectPointerBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectPointerBuilder struct {
	message.VLSPointerBuilder
}

// CurrentPositionsRejectFixedPointer If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixedPointer struct {
	message.FixedPointer
}

// CurrentPositionsRejectFixedPointerBuilder If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
type CurrentPositionsRejectFixedPointerBuilder struct {
	message.FixedPointer
}

func NewCurrentPositionsRejectFrom(b []byte) CurrentPositionsReject {
	return CurrentPositionsReject{VLS: message.NewVLSFrom(b)}
}

func WrapCurrentPositionsReject(b []byte) CurrentPositionsReject {
	return CurrentPositionsReject{VLS: message.WrapVLS(b)}
}

func NewCurrentPositionsReject() CurrentPositionsRejectBuilder {
	a := CurrentPositionsRejectBuilder{message.NewVLSBuilder(16)}
	a.Unsafe().SetBytes(0, _CurrentPositionsRejectDefault)
	return a
}

func NewCurrentPositionsRejectFixedFrom(b []byte) CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapCurrentPositionsRejectFixed(b []byte) CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{Fixed: message.WrapFixed(b)}
}

func NewCurrentPositionsRejectFixed() CurrentPositionsRejectFixedBuilder {
	a := CurrentPositionsRejectFixedBuilder{message.NewFixed(104)}
	a.Unsafe().SetBytes(0, _CurrentPositionsRejectFixedDefault)
	return a
}

func AllocCurrentPositionsReject() CurrentPositionsRejectPointerBuilder {
	a := CurrentPositionsRejectPointerBuilder{message.AllocVLSBuilder(16)}
	a.Ptr.SetBytes(0, _CurrentPositionsRejectDefault)
	return a
}

func AllocCurrentPositionsRejectFrom(b []byte) CurrentPositionsRejectPointer {
	return CurrentPositionsRejectPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocCurrentPositionsRejectFixed() CurrentPositionsRejectFixedPointerBuilder {
	a := CurrentPositionsRejectFixedPointerBuilder{message.AllocFixed(104)}
	a.Ptr.SetBytes(0, _CurrentPositionsRejectFixedDefault)
	return a
}

func AllocCurrentPositionsRejectFixedFrom(b []byte) CurrentPositionsRejectFixedPointer {
	return CurrentPositionsRejectFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size       = CurrentPositionsRejectSize  (16)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     BaseSize   = CurrentPositionsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CurrentPositionsRejectDefault)
}

// Clear
// {
//     Size       = CurrentPositionsRejectFixedSize  (104)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _CurrentPositionsRejectFixedDefault)
}

// Clear
// {
//     Size       = CurrentPositionsRejectSize  (16)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     BaseSize   = CurrentPositionsRejectSize  (16)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CurrentPositionsRejectDefault)
}

// Clear
// {
//     Size       = CurrentPositionsRejectFixedSize  (104)
//     Type       = CURRENT_POSITIONS_REJECT  (307)
//     RequestID  = 0
//     RejectText = ""
// }
func (m CurrentPositionsRejectFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _CurrentPositionsRejectFixedDefault)
}

// ToBuilder
func (m CurrentPositionsReject) ToBuilder() CurrentPositionsRejectBuilder {
	return CurrentPositionsRejectBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m CurrentPositionsRejectFixed) ToBuilder() CurrentPositionsRejectFixedBuilder {
	return CurrentPositionsRejectFixedBuilder{m.Fixed}
}

// ToBuilder
func (m CurrentPositionsRejectPointer) ToBuilder() CurrentPositionsRejectPointerBuilder {
	return CurrentPositionsRejectPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m CurrentPositionsRejectFixedPointer) ToBuilder() CurrentPositionsRejectFixedPointerBuilder {
	return CurrentPositionsRejectFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m CurrentPositionsRejectBuilder) Finish() CurrentPositionsReject {
	return CurrentPositionsReject{m.VLSBuilder.Finish()}
}

// Finish
func (m CurrentPositionsRejectFixedBuilder) Finish() CurrentPositionsRejectFixed {
	return CurrentPositionsRejectFixed{m.Fixed}
}

// Finish
func (m *CurrentPositionsRejectPointerBuilder) Finish() CurrentPositionsRejectPointer {
	return CurrentPositionsRejectPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *CurrentPositionsRejectFixedPointerBuilder) Finish() CurrentPositionsRejectFixedPointer {
	return CurrentPositionsRejectFixedPointer{m.FixedPointer}
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsReject) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsReject) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectBuilder) RejectText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectPointer) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectPointerBuilder) RejectText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixed) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedBuilder) RejectText() string {
	return message.StringFixed(m.Unsafe(), 104, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedPointer) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// RejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedPointerBuilder) RejectText() string {
	return message.StringFixed(m.Ptr, 104, 8)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m *CurrentPositionsRejectBuilder) SetRejectText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m *CurrentPositionsRejectPointerBuilder) SetRejectText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID This is set to the RequestID field sent in the CurrentPositionsRequest
// message.
func (m CurrentPositionsRejectFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Unsafe(), 104, 8, value)
}

// SetRejectText Free-form text indicating the reason for the rejection.
//
func (m CurrentPositionsRejectFixedPointerBuilder) SetRejectText(value string) {
	message.SetStringFixed(m.Ptr, 104, 8, value)
}

// Copy
func (m CurrentPositionsReject) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectBuilder) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsReject) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectBuilder) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsReject) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectBuilder) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsReject) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectBuilder) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectFixed) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectFixedBuilder) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixed) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixedBuilder) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixed) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixedBuilder) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixed) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixedBuilder) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectPointer) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectPointerBuilder) Copy(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectPointer) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectPointerBuilder) CopyPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectPointer) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectPointerBuilder) CopyTo(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectPointer) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectPointerBuilder) CopyToPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectFixedPointer) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m CurrentPositionsRejectFixedPointerBuilder) Copy(to CurrentPositionsRejectFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixedPointer) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyPointer
func (m CurrentPositionsRejectFixedPointerBuilder) CopyPointer(to CurrentPositionsRejectFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixedPointer) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m CurrentPositionsRejectFixedPointerBuilder) CopyTo(to CurrentPositionsRejectBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixedPointer) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyToPointer
func (m CurrentPositionsRejectFixedPointerBuilder) CopyToPointer(to CurrentPositionsRejectPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}
