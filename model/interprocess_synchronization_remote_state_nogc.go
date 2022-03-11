package model

import (
	"github.com/moontrade/dtc-go/message"
)

type InterprocessSynchronizationRemoteStatePointer struct {
	message.FixedPointer
}

type InterprocessSynchronizationRemoteStatePointerBuilder struct {
	message.FixedPointer
}

func AllocInterprocessSynchronizationRemoteState() InterprocessSynchronizationRemoteStatePointerBuilder {
	a := InterprocessSynchronizationRemoteStatePointerBuilder{message.AllocFixed(7)}
	a.Ptr.SetBytes(0, _InterprocessSynchronizationRemoteStateDefault)
	return a
}

func AllocInterprocessSynchronizationRemoteStateFrom(b []byte) InterprocessSynchronizationRemoteStatePointer {
	return InterprocessSynchronizationRemoteStatePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size                      = InterprocessSynchronizationRemoteStateSize  (7)
//     Type                      = INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE  (10134)
//     IsPrimary                 = 0
//     IsSecondary               = 0
//     IsSecondaryFailoverActive = 0
// }
func (m InterprocessSynchronizationRemoteStatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _InterprocessSynchronizationRemoteStateDefault)
}

// ToBuilder
func (m InterprocessSynchronizationRemoteStatePointer) ToBuilder() InterprocessSynchronizationRemoteStatePointerBuilder {
	return InterprocessSynchronizationRemoteStatePointerBuilder{m.FixedPointer}
}

// Finish
func (m *InterprocessSynchronizationRemoteStatePointerBuilder) Finish() InterprocessSynchronizationRemoteStatePointer {
	return InterprocessSynchronizationRemoteStatePointer{m.FixedPointer}
}

// IsPrimary
func (m InterprocessSynchronizationRemoteStatePointer) IsPrimary() uint8 {
	return message.Uint8Fixed(m.Ptr, 5, 4)
}

// IsPrimary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) IsPrimary() uint8 {
	return message.Uint8Fixed(m.Ptr, 5, 4)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteStatePointer) IsSecondary() uint8 {
	return message.Uint8Fixed(m.Ptr, 6, 5)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) IsSecondary() uint8 {
	return message.Uint8Fixed(m.Ptr, 6, 5)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStatePointer) IsSecondaryFailoverActive() uint8 {
	return message.Uint8Fixed(m.Ptr, 7, 6)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStatePointerBuilder) IsSecondaryFailoverActive() uint8 {
	return message.Uint8Fixed(m.Ptr, 7, 6)
}

// SetIsPrimary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) SetIsPrimary(value uint8) {
	message.SetUint8Fixed(m.Ptr, 5, 4, value)
}

// SetIsSecondary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) SetIsSecondary(value uint8) {
	message.SetUint8Fixed(m.Ptr, 6, 5, value)
}

// SetIsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStatePointerBuilder) SetIsSecondaryFailoverActive(value uint8) {
	message.SetUint8Fixed(m.Ptr, 7, 6, value)
}

// Copy
func (m InterprocessSynchronizationRemoteStatePointer) Copy(to InterprocessSynchronizationRemoteStateBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}

// Copy
func (m InterprocessSynchronizationRemoteStatePointerBuilder) Copy(to InterprocessSynchronizationRemoteStateBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}

// CopyPointer
func (m InterprocessSynchronizationRemoteStatePointer) CopyPointer(to InterprocessSynchronizationRemoteStatePointerBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}

// CopyPointer
func (m InterprocessSynchronizationRemoteStatePointerBuilder) CopyPointer(to InterprocessSynchronizationRemoteStatePointerBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}
