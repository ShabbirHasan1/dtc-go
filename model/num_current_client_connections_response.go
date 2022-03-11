package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

const NumCurrentClientConnectionsResponseSize = 22

type NumCurrentClientConnectionsResponse struct {
	message.VLS
}

type NumCurrentClientConnectionsResponseBuilder struct {
	message.VLSBuilder
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

// ToBuilder
func (m NumCurrentClientConnectionsResponse) ToBuilder() NumCurrentClientConnectionsResponseBuilder {
	return NumCurrentClientConnectionsResponseBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m NumCurrentClientConnectionsResponseBuilder) Finish() NumCurrentClientConnectionsResponse {
	return NumCurrentClientConnectionsResponse{m.VLSBuilder.Finish()}
}

// RequestID
func (m NumCurrentClientConnectionsResponse) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m NumCurrentClientConnectionsResponseBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// Username
func (m NumCurrentClientConnectionsResponse) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m NumCurrentClientConnectionsResponseBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponse) NumConnectionsForDifferentDevices() int32 {
	return message.Int32VLS(m.Unsafe(), 18, 14)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseBuilder) NumConnectionsForDifferentDevices() int32 {
	return message.Int32VLS(m.Unsafe(), 18, 14)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponse) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Unsafe(), 22, 18)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Unsafe(), 22, 18)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponseBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetUsername
func (m *NumCurrentClientConnectionsResponseBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32VLS(m.Unsafe(), 18, 14, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32VLS(m.Unsafe(), 22, 18, value)
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
