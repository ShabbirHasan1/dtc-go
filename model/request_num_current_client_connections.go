package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size             = RequestNumCurrentClientConnectionsSize  (22)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     BaseSize         = RequestNumCurrentClientConnectionsSize  (22)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
var _RequestNumCurrentClientConnectionsDefault = []byte{22, 0, 123, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const RequestNumCurrentClientConnectionsSize = 22

type RequestNumCurrentClientConnections struct {
	message.VLS
}

type RequestNumCurrentClientConnectionsBuilder struct {
	message.VLSBuilder
}

func NewRequestNumCurrentClientConnectionsFrom(b []byte) RequestNumCurrentClientConnections {
	return RequestNumCurrentClientConnections{VLS: message.NewVLSFrom(b)}
}

func WrapRequestNumCurrentClientConnections(b []byte) RequestNumCurrentClientConnections {
	return RequestNumCurrentClientConnections{VLS: message.WrapVLS(b)}
}

func NewRequestNumCurrentClientConnections() RequestNumCurrentClientConnectionsBuilder {
	a := RequestNumCurrentClientConnectionsBuilder{message.NewVLSBuilder(22)}
	a.Unsafe().SetBytes(0, _RequestNumCurrentClientConnectionsDefault)
	return a
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
func (m RequestNumCurrentClientConnectionsBuilder) Clear() {
	m.Unsafe().SetBytes(0, _RequestNumCurrentClientConnectionsDefault)
}

// ToBuilder
func (m RequestNumCurrentClientConnections) ToBuilder() RequestNumCurrentClientConnectionsBuilder {
	return RequestNumCurrentClientConnectionsBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m RequestNumCurrentClientConnectionsBuilder) Finish() RequestNumCurrentClientConnections {
	return RequestNumCurrentClientConnections{m.VLSBuilder.Finish()}
}

// RequestID
func (m RequestNumCurrentClientConnections) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m RequestNumCurrentClientConnectionsBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// Username
func (m RequestNumCurrentClientConnections) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m RequestNumCurrentClientConnectionsBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnections) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Unsafe(), 22, 14)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsBuilder) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Unsafe(), 22, 14)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetUsername
func (m *RequestNumCurrentClientConnectionsBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64VLS(m.Unsafe(), 22, 14, value)
}

// Copy
func (m RequestNumCurrentClientConnections) Copy(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsBuilder) Copy(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnections) CopyPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsBuilder) CopyPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnections) CopyToPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsBuilder) CopyToPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnections) CopyTo(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsBuilder) CopyTo(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}
