package model

import (
	"github.com/moontrade/dtc-go/message"
)

type ClientDeviceUpdatePointer struct {
	message.FixedPointer
}

type ClientDeviceUpdatePointerBuilder struct {
	message.FixedPointer
}

func AllocClientDeviceUpdate() ClientDeviceUpdatePointerBuilder {
	a := ClientDeviceUpdatePointerBuilder{message.AllocFixed(12)}
	a.Ptr.SetBytes(0, _ClientDeviceUpdateDefault)
	return a
}

func AllocClientDeviceUpdateFrom(b []byte) ClientDeviceUpdatePointer {
	return ClientDeviceUpdatePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                   = ClientDeviceUpdateSize  (12)
//     Type                   = CLIENT_DEVICE_UPDATE  (10139)
//     ClientDeviceIdentifier = 0
// }
func (m ClientDeviceUpdatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _ClientDeviceUpdateDefault)
}

// ToBuilder
func (m ClientDeviceUpdatePointer) ToBuilder() ClientDeviceUpdatePointerBuilder {
	return ClientDeviceUpdatePointerBuilder{m.FixedPointer}
}

// Finish
func (m *ClientDeviceUpdatePointerBuilder) Finish() ClientDeviceUpdatePointer {
	return ClientDeviceUpdatePointer{m.FixedPointer}
}

// ClientDeviceIdentifier
func (m ClientDeviceUpdatePointer) ClientDeviceIdentifier() int64 {
	return message.Int64Fixed(m.Ptr, 12, 4)
}

// ClientDeviceIdentifier
func (m ClientDeviceUpdatePointerBuilder) ClientDeviceIdentifier() int64 {
	return message.Int64Fixed(m.Ptr, 12, 4)
}

// SetClientDeviceIdentifier
func (m ClientDeviceUpdatePointerBuilder) SetClientDeviceIdentifier(value int64) {
	message.SetInt64Fixed(m.Ptr, 12, 4, value)
}

// Copy
func (m ClientDeviceUpdatePointer) Copy(to ClientDeviceUpdateBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
}

// Copy
func (m ClientDeviceUpdatePointerBuilder) Copy(to ClientDeviceUpdateBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
}

// CopyPointer
func (m ClientDeviceUpdatePointer) CopyPointer(to ClientDeviceUpdatePointerBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
}

// CopyPointer
func (m ClientDeviceUpdatePointerBuilder) CopyPointer(to ClientDeviceUpdatePointerBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
}
