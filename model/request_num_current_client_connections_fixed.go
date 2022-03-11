package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = RequestNumCurrentClientConnectionsFixedSize  (48)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
var _RequestNumCurrentClientConnectionsFixedDefault = []byte{48, 0, 123, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const RequestNumCurrentClientConnectionsFixedSize = 48

type RequestNumCurrentClientConnectionsFixed struct {
	message.Fixed
}

type RequestNumCurrentClientConnectionsFixedBuilder struct {
	message.Fixed
}

func NewRequestNumCurrentClientConnectionsFixedFrom(b []byte) RequestNumCurrentClientConnectionsFixed {
	return RequestNumCurrentClientConnectionsFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapRequestNumCurrentClientConnectionsFixed(b []byte) RequestNumCurrentClientConnectionsFixed {
	return RequestNumCurrentClientConnectionsFixed{Fixed: message.WrapFixed(b)}
}

func NewRequestNumCurrentClientConnectionsFixed() RequestNumCurrentClientConnectionsFixedBuilder {
	a := RequestNumCurrentClientConnectionsFixedBuilder{message.NewFixed(48)}
	a.Unsafe().SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
	return a
}

// Clear
// {
//     Size             = RequestNumCurrentClientConnectionsFixedSize  (48)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
func (m RequestNumCurrentClientConnectionsFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
}

// ToBuilder
func (m RequestNumCurrentClientConnectionsFixed) ToBuilder() RequestNumCurrentClientConnectionsFixedBuilder {
	return RequestNumCurrentClientConnectionsFixedBuilder{m.Fixed}
}

// Finish
func (m RequestNumCurrentClientConnectionsFixedBuilder) Finish() RequestNumCurrentClientConnectionsFixed {
	return RequestNumCurrentClientConnectionsFixed{m.Fixed}
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// Username
func (m RequestNumCurrentClientConnectionsFixed) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username
func (m RequestNumCurrentClientConnectionsFixedBuilder) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixed) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Unsafe(), 48, 40)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedBuilder) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Unsafe(), 48, 40)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetUsername
func (m RequestNumCurrentClientConnectionsFixedBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 48, 40, value)
}

// Copy
func (m RequestNumCurrentClientConnectionsFixed) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsFixedBuilder) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixed) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixedBuilder) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixed) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixedBuilder) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixed) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixedBuilder) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}
