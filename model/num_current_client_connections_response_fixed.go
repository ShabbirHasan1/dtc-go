package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                              = NumCurrentClientConnectionsResponseFixedSize  (48)
//     Type                              = NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE  (10108)
//     RequestID                         = 0
//     Username                          = ""
//     NumConnectionsForDifferentDevices = 0
//     NumConnectionsForSameDevice       = 0
// }
var _NumCurrentClientConnectionsResponseFixedDefault = []byte{48, 0, 124, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const NumCurrentClientConnectionsResponseFixedSize = 48

type NumCurrentClientConnectionsResponseFixed struct {
	message.Fixed
}

type NumCurrentClientConnectionsResponseFixedBuilder struct {
	message.Fixed
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

// ToBuilder
func (m NumCurrentClientConnectionsResponseFixed) ToBuilder() NumCurrentClientConnectionsResponseFixedBuilder {
	return NumCurrentClientConnectionsResponseFixedBuilder{m.Fixed}
}

// Finish
func (m NumCurrentClientConnectionsResponseFixedBuilder) Finish() NumCurrentClientConnectionsResponseFixed {
	return NumCurrentClientConnectionsResponseFixed{m.Fixed}
}

// RequestID
func (m NumCurrentClientConnectionsResponseFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m NumCurrentClientConnectionsResponseFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Username
func (m NumCurrentClientConnectionsResponseFixed) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username
func (m NumCurrentClientConnectionsResponseFixedBuilder) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixed) NumConnectionsForDifferentDevices() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// NumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedBuilder) NumConnectionsForDifferentDevices() int32 {
	return message.Int32Fixed(m.Unsafe(), 44, 40)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixed) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// NumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedBuilder) NumConnectionsForSameDevice() int32 {
	return message.Int32Fixed(m.Unsafe(), 48, 44)
}

// SetRequestID
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetUsername
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetNumConnectionsForDifferentDevices
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetNumConnectionsForDifferentDevices(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 44, 40, value)
}

// SetNumConnectionsForSameDevice
func (m NumCurrentClientConnectionsResponseFixedBuilder) SetNumConnectionsForSameDevice(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 48, 44, value)
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
