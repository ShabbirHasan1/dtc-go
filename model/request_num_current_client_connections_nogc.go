package model

import (
	"github.com/moontrade/dtc-go/message"
)

type RequestNumCurrentClientConnectionsPointer struct {
	message.VLSPointer
}

type RequestNumCurrentClientConnectionsPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocRequestNumCurrentClientConnections() RequestNumCurrentClientConnectionsPointerBuilder {
	a := RequestNumCurrentClientConnectionsPointerBuilder{message.AllocVLSBuilder(22)}
	a.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsDefault)
	return a
}

func AllocRequestNumCurrentClientConnectionsFrom(b []byte) RequestNumCurrentClientConnectionsPointer {
	return RequestNumCurrentClientConnectionsPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size             = RequestNumCurrentClientConnectionsSize  (22)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     BaseSize         = RequestNumCurrentClientConnectionsSize  (22)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
func (m RequestNumCurrentClientConnectionsPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsDefault)
}

// ToBuilder
func (m RequestNumCurrentClientConnectionsPointer) ToBuilder() RequestNumCurrentClientConnectionsPointerBuilder {
	return RequestNumCurrentClientConnectionsPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *RequestNumCurrentClientConnectionsPointerBuilder) Finish() RequestNumCurrentClientConnectionsPointer {
	return RequestNumCurrentClientConnectionsPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID
func (m RequestNumCurrentClientConnectionsPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m RequestNumCurrentClientConnectionsPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// Username
func (m RequestNumCurrentClientConnectionsPointer) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m RequestNumCurrentClientConnectionsPointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsPointer) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Ptr, 22, 14)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsPointerBuilder) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Ptr, 22, 14)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetUsername
func (m *RequestNumCurrentClientConnectionsPointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsPointerBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64VLS(m.Ptr, 22, 14, value)
}

// Copy
func (m RequestNumCurrentClientConnectionsPointer) Copy(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsPointerBuilder) Copy(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsPointer) CopyPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsPointerBuilder) CopyPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsPointer) CopyTo(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsPointerBuilder) CopyTo(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsPointer) CopyToPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsPointerBuilder) CopyToPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}
