package model

import (
	"github.com/moontrade/dtc-go/message"
)

type RequestNumCurrentClientConnectionsFixedPointer struct {
	message.FixedPointer
}

type RequestNumCurrentClientConnectionsFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocRequestNumCurrentClientConnectionsFixed() RequestNumCurrentClientConnectionsFixedPointerBuilder {
	a := RequestNumCurrentClientConnectionsFixedPointerBuilder{message.AllocFixed(48)}
	a.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
	return a
}

func AllocRequestNumCurrentClientConnectionsFixedFrom(b []byte) RequestNumCurrentClientConnectionsFixedPointer {
	return RequestNumCurrentClientConnectionsFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size             = RequestNumCurrentClientConnectionsFixedSize  (48)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
}

// ToBuilder
func (m RequestNumCurrentClientConnectionsFixedPointer) ToBuilder() RequestNumCurrentClientConnectionsFixedPointerBuilder {
	return RequestNumCurrentClientConnectionsFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *RequestNumCurrentClientConnectionsFixedPointerBuilder) Finish() RequestNumCurrentClientConnectionsFixedPointer {
	return RequestNumCurrentClientConnectionsFixedPointer{m.FixedPointer}
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixedPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Username
func (m RequestNumCurrentClientConnectionsFixedPointer) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Username
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedPointer) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Ptr, 48, 40)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Ptr, 48, 40)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetUsername
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64Fixed(m.Ptr, 48, 40, value)
}

// Copy
func (m RequestNumCurrentClientConnectionsFixedPointer) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixedPointer) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixedPointer) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixedPointer) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}
