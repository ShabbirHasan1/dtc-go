package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size                      = InterprocessSynchronizationRemoteStateSize  (7)
//     Type                      = INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE  (10134)
//     IsPrimary                 = 0
//     IsSecondary               = 0
//     IsSecondaryFailoverActive = 0
// }
var _InterprocessSynchronizationRemoteStateDefault = []byte{7, 0, 150, 39, 0, 0, 0}

const InterprocessSynchronizationRemoteStateSize = 7

type InterprocessSynchronizationRemoteState struct {
	message.Fixed
}

type InterprocessSynchronizationRemoteStateBuilder struct {
	message.Fixed
}

func NewInterprocessSynchronizationRemoteStateFrom(b []byte) InterprocessSynchronizationRemoteState {
	return InterprocessSynchronizationRemoteState{Fixed: message.NewFixedFrom(b)}
}

func WrapInterprocessSynchronizationRemoteState(b []byte) InterprocessSynchronizationRemoteState {
	return InterprocessSynchronizationRemoteState{Fixed: message.WrapFixed(b)}
}

func NewInterprocessSynchronizationRemoteState() InterprocessSynchronizationRemoteStateBuilder {
	a := InterprocessSynchronizationRemoteStateBuilder{message.NewFixed(7)}
	a.Unsafe().SetBytes(0, _InterprocessSynchronizationRemoteStateDefault)
	return a
}

// Clear
// {
//     Size                      = InterprocessSynchronizationRemoteStateSize  (7)
//     Type                      = INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE  (10134)
//     IsPrimary                 = 0
//     IsSecondary               = 0
//     IsSecondaryFailoverActive = 0
// }
func (m InterprocessSynchronizationRemoteStateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _InterprocessSynchronizationRemoteStateDefault)
}

// ToBuilder
func (m InterprocessSynchronizationRemoteState) ToBuilder() InterprocessSynchronizationRemoteStateBuilder {
	return InterprocessSynchronizationRemoteStateBuilder{m.Fixed}
}

// Finish
func (m InterprocessSynchronizationRemoteStateBuilder) Finish() InterprocessSynchronizationRemoteState {
	return InterprocessSynchronizationRemoteState{m.Fixed}
}

// IsPrimary
func (m InterprocessSynchronizationRemoteState) IsPrimary() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 5, 4)
}

// IsPrimary
func (m InterprocessSynchronizationRemoteStateBuilder) IsPrimary() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 5, 4)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteState) IsSecondary() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 6, 5)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteStateBuilder) IsSecondary() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 6, 5)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteState) IsSecondaryFailoverActive() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 7, 6)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStateBuilder) IsSecondaryFailoverActive() uint8 {
	return message.Uint8Fixed(m.Unsafe(), 7, 6)
}

// SetIsPrimary
func (m InterprocessSynchronizationRemoteStateBuilder) SetIsPrimary(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 5, 4, value)
}

// SetIsSecondary
func (m InterprocessSynchronizationRemoteStateBuilder) SetIsSecondary(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 6, 5, value)
}

// SetIsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStateBuilder) SetIsSecondaryFailoverActive(value uint8) {
	message.SetUint8Fixed(m.Unsafe(), 7, 6, value)
}

// Copy
func (m InterprocessSynchronizationRemoteState) Copy(to InterprocessSynchronizationRemoteStateBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}

// Copy
func (m InterprocessSynchronizationRemoteStateBuilder) Copy(to InterprocessSynchronizationRemoteStateBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}

// CopyPointer
func (m InterprocessSynchronizationRemoteState) CopyPointer(to InterprocessSynchronizationRemoteStatePointerBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}

// CopyPointer
func (m InterprocessSynchronizationRemoteStateBuilder) CopyPointer(to InterprocessSynchronizationRemoteStatePointerBuilder) {
	to.SetIsPrimary(m.IsPrimary())
	to.SetIsSecondary(m.IsSecondary())
	to.SetIsSecondaryFailoverActive(m.IsSecondaryFailoverActive())
}
