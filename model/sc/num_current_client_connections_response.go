package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const NumCurrentClientConnectionsResponseSize = 22

const NumCurrentClientConnectionsResponseFixedSize = 48

// {
//     Size                              = NumCurrentClientConnectionsResponseSize  (22)
//     Type                              = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
//     BaseSize                          = NumCurrentClientConnectionsResponseSize  (22)
//     RequestID                         = 0
//     Username                          = ""
//     NumConnectionsForDifferentDevices = 0
//     NumConnectionsForSameDevice       = 0
// }
var _NumCurrentClientConnectionsResponseDefault = []byte{22, 0, 124, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size                              = NumCurrentClientConnectionsResponseFixedSize  (48)
//     Type                              = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
//     RequestID                         = 0
//     Username                          = ""
//     NumConnectionsForDifferentDevices = 0
//     NumConnectionsForSameDevice       = 0
// }
var _NumCurrentClientConnectionsResponseFixedDefault = []byte{48, 0, 124, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type NumCurrentClientConnectionsResponse struct {
	message.VLS
}

type NumCurrentClientConnectionsResponseBuilder struct {
	message.VLSBuilder
}

type NumCurrentClientConnectionsResponseFixed struct {
	message.Fixed
}

type NumCurrentClientConnectionsResponseFixedBuilder struct {
	message.Fixed
}

type NumCurrentClientConnectionsResponsePointer struct {
	message.VLSPointer
}

type NumCurrentClientConnectionsResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

type NumCurrentClientConnectionsResponseFixedPointer struct {
	message.FixedPointer
}

type NumCurrentClientConnectionsResponseFixedPointerBuilder struct {
	message.FixedPointer
}

func NewNumCurrentClientConnectionsResponseFrom(b []byte) NumCurrentClientConnectionsResponse {
	return NumCurrentClientConnectionsResponse{VLS: message.NewVLSFrom(b)}
}

func WrapNumCurrentClientConnectionsResponse(b []byte) NumCurrentClientConnectionsResponse {
	return NumCurrentClientConnectionsResponse{VLS: message.WrapVLS(b)}
}

func NewNumCurrentClientConnectionsResponse() NumCurrentClientConnectionsResponseBuilder {
	a := NumCurrentClientConnectionsResponseBuilder{message.NewVLSBuilder(22)}
	a.Unsafe().SetBytes(0, _NumCurrentClientConnectionsResponseDefault)
	return a
}

func NewNumCurrentClientConnectionsResponseFixedFrom(b []byte) NumCurrentClientConnectionsResponseFixed {
	return NumCurrentClientConnectionsResponseFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapNumCurrentClientConnectionsResponseFixed(b []byte) NumCurrentClientConnectionsResponseFixed {
	return NumCurrentClientConnectionsResponseFixed{Fixed: message.WrapFixed(b)}
}

func NewNumCurrentClientConnectionsResponseFixed() NumCurrentClientConnectionsResponseFixedBuilder {
	a := NumCurrentClientConnectionsResponseFixedBuilder{message.NewFixed(48)}
	a.Unsafe().SetBytes(0, _NumCurrentClientConnectionsResponseFixedDefault)
	return a
}

func AllocNumCurrentClientConnectionsResponse() NumCurrentClientConnectionsResponsePointerBuilder {
	a := NumCurrentClientConnectionsResponsePointerBuilder{message.AllocVLSBuilder(22)}
	a.Ptr.SetBytes(0, _NumCurrentClientConnectionsResponseDefault)
	return a
}

func AllocNumCurrentClientConnectionsResponseFrom(b []byte) NumCurrentClientConnectionsResponsePointer {
	return NumCurrentClientConnectionsResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocNumCurrentClientConnectionsResponseFixed() NumCurrentClientConnectionsResponseFixedPointerBuilder {
	a := NumCurrentClientConnectionsResponseFixedPointerBuilder{message.AllocFixed(48)}
	a.Ptr.SetBytes(0, _NumCurrentClientConnectionsResponseFixedDefault)
	return a
}

func AllocNumCurrentClientConnectionsResponseFixedFrom(b []byte) NumCurrentClientConnectionsResponseFixedPointer {
	return NumCurrentClientConnectionsResponseFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                              = NumCurrentClientConnectionsResponseSize  (22)
//     Type                              = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
//     BaseSize                          = NumCurrentClientConnectionsResponseSize  (22)
//     RequestID                         = 0
//     Username                          = ""
//     NumConnectionsForDifferentDevices = 0
//     NumConnectionsForSameDevice       = 0
// }
func (m NumCurrentClientConnectionsResponseBuilder) Clear() {
	m.Unsafe().SetBytes(0, _NumCurrentClientConnectionsResponseDefault)
}

// Clear
// {
//     Size                              = NumCurrentClientConnectionsResponseFixedSize  (48)
//     Type                              = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
//     RequestID                         = 0
//     Username                          = ""
//     NumConnectionsForDifferentDevices = 0
//     NumConnectionsForSameDevice       = 0
// }
func (m NumCurrentClientConnectionsResponseFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _NumCurrentClientConnectionsResponseFixedDefault)
}

// Clear
// {
//     Size                              = NumCurrentClientConnectionsResponseSize  (22)
//     Type                              = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
//     BaseSize                          = NumCurrentClientConnectionsResponseSize  (22)
//     RequestID                         = 0
//     Username                          = ""
//     NumConnectionsForDifferentDevices = 0
//     NumConnectionsForSameDevice       = 0
// }
func (m NumCurrentClientConnectionsResponsePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _NumCurrentClientConnectionsResponseDefault)
}

// Clear
// {
//     Size                              = NumCurrentClientConnectionsResponseFixedSize  (48)
//     Type                              = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
//     RequestID                         = 0
//     Username                          = ""
//     NumConnectionsForDifferentDevices = 0
//     NumConnectionsForSameDevice       = 0
// }
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _NumCurrentClientConnectionsResponseFixedDefault)
}

// ToBuilder
func (m NumCurrentClientConnectionsResponse) ToBuilder() NumCurrentClientConnectionsResponseBuilder {
	return NumCurrentClientConnectionsResponseBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m NumCurrentClientConnectionsResponseFixed) ToBuilder() NumCurrentClientConnectionsResponseFixedBuilder {
	return NumCurrentClientConnectionsResponseFixedBuilder{m.Fixed}
}

// ToBuilder
func (m NumCurrentClientConnectionsResponsePointer) ToBuilder() NumCurrentClientConnectionsResponsePointerBuilder {
	return NumCurrentClientConnectionsResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m NumCurrentClientConnectionsResponseFixedPointer) ToBuilder() NumCurrentClientConnectionsResponseFixedPointerBuilder {
	return NumCurrentClientConnectionsResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m NumCurrentClientConnectionsResponseBuilder) Finish() NumCurrentClientConnectionsResponse {
	return NumCurrentClientConnectionsResponse{m.VLSBuilder.Finish()}
}

// Finish
func (m NumCurrentClientConnectionsResponseFixedBuilder) Finish() NumCurrentClientConnectionsResponseFixed {
	return NumCurrentClientConnectionsResponseFixed{m.Fixed}
}

// Finish
func (m *NumCurrentClientConnectionsResponsePointerBuilder) Finish() NumCurrentClientConnectionsResponsePointer {
	return NumCurrentClientConnectionsResponsePointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *NumCurrentClientConnectionsResponseFixedPointerBuilder) Finish() NumCurrentClientConnectionsResponseFixedPointer {
	return NumCurrentClientConnectionsResponseFixedPointer{m.FixedPointer}
}

// RequestID
func (m NumCurrentClientConnectionsResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m NumCurrentClientConnectionsResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m NumCurrentClientConnectionsResponsePointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m NumCurrentClientConnectionsResponsePointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// Username
func (m NumCurrentClientConnectionsResponse) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m NumCurrentClientConnectionsResponseBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m NumCurrentClientConnectionsResponsePointer) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m NumCurrentClientConnectionsResponsePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponse) NumConnectionsForDifferentDevices() int32 {
	return message.Int32VLS(m.Unsafe(), 18, 14)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseBuilder) NumConnectionsForDifferentDevices() int32 {
	return message.Int32VLS(m.Unsafe(), 18, 14)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponsePointer) NumConnectionsForDifferentDevices() int32 {
	return message.Int32VLS(m.Ptr, 18, 14)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponsePointerBuilder) NumConnectionsForDifferentDevices() int32 {
	return message.Int32VLS(m.Ptr, 18, 14)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponse) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Unsafe(), 22, 18)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Unsafe(), 22, 18)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponsePointer) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Ptr, 22, 18)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponsePointerBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Ptr, 22, 18)
}

// RequestID
func (m NumCurrentClientConnectionsResponseFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m NumCurrentClientConnectionsResponseFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m NumCurrentClientConnectionsResponseFixedPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Username
func (m NumCurrentClientConnectionsResponseFixed) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username
func (m NumCurrentClientConnectionsResponseFixedBuilder) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username
func (m NumCurrentClientConnectionsResponseFixedPointer) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Username
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixed) NumConnectionsForDifferentDevices() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedBuilder) NumConnectionsForDifferentDevices() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedPointer) NumConnectionsForDifferentDevices() int32 {
	return message.Int32Fixed(m.Ptr, 44, 40)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) NumConnectionsForDifferentDevices() int32 {
	return message.Int32Fixed(m.Ptr, 44, 40)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixed) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedPointer) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetUsername
func (m *NumCurrentClientConnectionsResponseBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *NumCurrentClientConnectionsResponsePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32VLS(m.Unsafe(), 18, 14, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponsePointerBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32VLS(m.Ptr, 18, 14, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32VLS(m.Unsafe(), 22, 18, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponsePointerBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32VLS(m.Ptr, 22, 18, value)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetUsername
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetUsername
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 44, 40, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32Fixed(m.Ptr, 44, 40, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 48, 44, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32Fixed(m.Ptr, 48, 44, value)
}

// Copy
func (m NumCurrentClientConnectionsResponse) Copy(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// Copy
func (m NumCurrentClientConnectionsResponseBuilder) Copy(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponse) CopyPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponseBuilder) CopyPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponse) CopyToPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponseBuilder) CopyToPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponse) CopyTo(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponseBuilder) CopyTo(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// Copy
func (m NumCurrentClientConnectionsResponseFixed) Copy(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// Copy
func (m NumCurrentClientConnectionsResponseFixedBuilder) Copy(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponseFixed) CopyPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponseFixedBuilder) CopyPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponseFixed) CopyToPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponseFixedBuilder) CopyToPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponseFixed) CopyTo(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponseFixedBuilder) CopyTo(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// Copy
func (m NumCurrentClientConnectionsResponsePointer) Copy(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// Copy
func (m NumCurrentClientConnectionsResponsePointerBuilder) Copy(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponsePointer) CopyPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponsePointerBuilder) CopyPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponsePointer) CopyTo(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponsePointerBuilder) CopyTo(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponsePointer) CopyToPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponsePointerBuilder) CopyToPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// Copy
func (m NumCurrentClientConnectionsResponseFixedPointer) Copy(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// Copy
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) Copy(to NumCurrentClientConnectionsResponseFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponseFixedPointer) CopyPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyPointer
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) CopyPointer(to NumCurrentClientConnectionsResponseFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponseFixedPointer) CopyTo(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyTo
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) CopyTo(to NumCurrentClientConnectionsResponseBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponseFixedPointer) CopyToPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}

// CopyToPointer
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) CopyToPointer(to NumCurrentClientConnectionsResponsePointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetNumConnectionsForDifferentDevices(m.NumConnectionsForDifferentDevices())
	to.SetNumConnectionsForSameDevice(m.NumConnectionsForSameDevice())
}
