package model

import (
	"github.com/moontrade/dtc-go/message"
)

type NumCurrentClientConnectionsResponseFixedPointer struct {
	message.FixedPointer
}

type NumCurrentClientConnectionsResponseFixedPointerBuilder struct {
	message.FixedPointer
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
func (m NumCurrentClientConnectionsResponseFixedPointer) ToBuilder() NumCurrentClientConnectionsResponseFixedPointerBuilder {
	return NumCurrentClientConnectionsResponseFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *NumCurrentClientConnectionsResponseFixedPointerBuilder) Finish() NumCurrentClientConnectionsResponseFixedPointer {
	return NumCurrentClientConnectionsResponseFixedPointer{m.FixedPointer}
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
func (m NumCurrentClientConnectionsResponseFixedPointer) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Username
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
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
func (m NumCurrentClientConnectionsResponseFixedPointer) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Ptr, 48, 44)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetUsername
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32Fixed(m.Ptr, 44, 40, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedPointerBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32Fixed(m.Ptr, 48, 44, value)
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
