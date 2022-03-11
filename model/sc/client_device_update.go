package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const ClientDeviceUpdateSize = 12

// {
//     Size                   = ClientDeviceUpdateSize  (12)
//     Type                   = CLIENT_DEVICE_UPDATE  (10139)
//     ClientDeviceIdentifier = 0
// }
var _ClientDeviceUpdateDefault = []byte{12, 0, 155, 39, 0, 0, 0, 0, 0, 0, 0, 0}

type ClientDeviceUpdate struct {
	message.Fixed
}

type ClientDeviceUpdateBuilder struct {
	message.Fixed
}

type ClientDeviceUpdatePointer struct {
	message.FixedPointer
}

type ClientDeviceUpdatePointerBuilder struct {
	message.FixedPointer
}

func NewClientDeviceUpdateFrom(b []byte) ClientDeviceUpdate {
	return ClientDeviceUpdate{Fixed: message.NewFixedFrom(b)}
}

func WrapClientDeviceUpdate(b []byte) ClientDeviceUpdate {
	return ClientDeviceUpdate{Fixed: message.WrapFixed(b)}
}

func NewClientDeviceUpdate() ClientDeviceUpdateBuilder {
	a := ClientDeviceUpdateBuilder{message.NewFixed(12)}
	a.Unsafe().SetBytes(0, _ClientDeviceUpdateDefault)
	return a
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
func (m ClientDeviceUpdateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _ClientDeviceUpdateDefault)
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
func (m ClientDeviceUpdate) ToBuilder() ClientDeviceUpdateBuilder {
	return ClientDeviceUpdateBuilder{m.Fixed}
}

// ToBuilder
func (m ClientDeviceUpdatePointer) ToBuilder() ClientDeviceUpdatePointerBuilder {
	return ClientDeviceUpdatePointerBuilder{m.FixedPointer}
}

// Finish
func (m ClientDeviceUpdateBuilder) Finish() ClientDeviceUpdate {
	return ClientDeviceUpdate{m.Fixed}
}

// Finish
func (m *ClientDeviceUpdatePointerBuilder) Finish() ClientDeviceUpdatePointer {
	return ClientDeviceUpdatePointer{m.FixedPointer}
}

// ClientDeviceIdentifier
func (m ClientDeviceUpdate) ClientDeviceIdentifier() int64 {
	return message.Int64Fixed(m.Unsafe(), 12, 4)
}

// ClientDeviceIdentifier
func (m ClientDeviceUpdateBuilder) ClientDeviceIdentifier() int64 {
	return message.Int64Fixed(m.Unsafe(), 12, 4)
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
func (m ClientDeviceUpdateBuilder) SetClientDeviceIdentifier(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 12, 4, value)
}

// SetClientDeviceIdentifier
func (m ClientDeviceUpdatePointerBuilder) SetClientDeviceIdentifier(value int64) {
	message.SetInt64Fixed(m.Ptr, 12, 4, value)
}

// Copy
func (m ClientDeviceUpdate) Copy(to ClientDeviceUpdateBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
}

// Copy
func (m ClientDeviceUpdateBuilder) Copy(to ClientDeviceUpdateBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
}

// CopyPointer
func (m ClientDeviceUpdate) CopyPointer(to ClientDeviceUpdatePointerBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
}

// CopyPointer
func (m ClientDeviceUpdateBuilder) CopyPointer(to ClientDeviceUpdatePointerBuilder) {
	to.SetClientDeviceIdentifier(m.ClientDeviceIdentifier())
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
