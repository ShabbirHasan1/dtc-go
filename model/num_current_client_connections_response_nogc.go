package model

import (
	"github.com/moontrade/dtc-go/message"
)

type NumCurrentClientConnectionsResponsePointer struct {
	message.VLSPointer
}

type NumCurrentClientConnectionsResponsePointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocNumCurrentClientConnectionsResponse() NumCurrentClientConnectionsResponsePointerBuilder {
	a := NumCurrentClientConnectionsResponsePointerBuilder{message.AllocVLSBuilder(22)}
	a.Ptr.SetBytes(0, _NumCurrentClientConnectionsResponseDefault)
	return a
}

func AllocNumCurrentClientConnectionsResponseFrom(b []byte) NumCurrentClientConnectionsResponsePointer {
	return NumCurrentClientConnectionsResponsePointer{VLSPointer: message.AllocVLSFrom(b)}
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

// ToBuilder
func (m NumCurrentClientConnectionsResponsePointer) ToBuilder() NumCurrentClientConnectionsResponsePointerBuilder {
	return NumCurrentClientConnectionsResponsePointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *NumCurrentClientConnectionsResponsePointerBuilder) Finish() NumCurrentClientConnectionsResponsePointer {
	return NumCurrentClientConnectionsResponsePointer{m.VLSPointerBuilder.Finish()}
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
func (m NumCurrentClientConnectionsResponsePointer) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m NumCurrentClientConnectionsResponsePointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
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
func (m NumCurrentClientConnectionsResponsePointer) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Ptr, 22, 18)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponsePointerBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32VLS(m.Ptr, 22, 18)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponsePointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetUsername
func (m *NumCurrentClientConnectionsResponsePointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponsePointerBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32VLS(m.Ptr, 18, 14, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponsePointerBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32VLS(m.Ptr, 22, 18, value)
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
