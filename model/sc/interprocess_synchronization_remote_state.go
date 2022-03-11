package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const InterprocessSynchronizationRemoteStateSize = 7

// {
//     Size                      = InterprocessSynchronizationRemoteStateSize  (7)
//     Type                      = INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE  (10134)
//     IsPrimary                 = false
//     IsSecondary               = false
//     IsSecondaryFailoverActive = false
// }
var _InterprocessSynchronizationRemoteStateDefault = []byte{7, 0, 150, 39, 0, 0, 0}

type InterprocessSynchronizationRemoteState struct {
	message.Fixed
}

type InterprocessSynchronizationRemoteStateBuilder struct {
	message.Fixed
}

type InterprocessSynchronizationRemoteStatePointer struct {
	message.FixedPointer
}

type InterprocessSynchronizationRemoteStatePointerBuilder struct {
	message.FixedPointer
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
//     IsPrimary                 = false
//     IsSecondary               = false
//     IsSecondaryFailoverActive = false
// }
func (m InterprocessSynchronizationRemoteStateBuilder) Clear() {
	m.Unsafe().SetBytes(0, _InterprocessSynchronizationRemoteStateDefault)
}

// Clear
// {
//     Size                      = InterprocessSynchronizationRemoteStateSize  (7)
//     Type                      = INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE  (10134)
//     IsPrimary                 = false
//     IsSecondary               = false
//     IsSecondaryFailoverActive = false
// }
func (m InterprocessSynchronizationRemoteStatePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _InterprocessSynchronizationRemoteStateDefault)
}

// ToBuilder
func (m InterprocessSynchronizationRemoteState) ToBuilder() InterprocessSynchronizationRemoteStateBuilder {
	return InterprocessSynchronizationRemoteStateBuilder{m.Fixed}
}

// ToBuilder
func (m InterprocessSynchronizationRemoteStatePointer) ToBuilder() InterprocessSynchronizationRemoteStatePointerBuilder {
	return InterprocessSynchronizationRemoteStatePointerBuilder{m.FixedPointer}
}

// Finish
func (m InterprocessSynchronizationRemoteStateBuilder) Finish() InterprocessSynchronizationRemoteState {
	return InterprocessSynchronizationRemoteState{m.Fixed}
}

// Finish
func (m *InterprocessSynchronizationRemoteStatePointerBuilder) Finish() InterprocessSynchronizationRemoteStatePointer {
	return InterprocessSynchronizationRemoteStatePointer{m.FixedPointer}
}

// IsPrimary
func (m InterprocessSynchronizationRemoteState) IsPrimary() bool {
	return message.BoolFixed(m.Unsafe(), 5, 4)
}

// IsPrimary
func (m InterprocessSynchronizationRemoteStateBuilder) IsPrimary() bool {
	return message.BoolFixed(m.Unsafe(), 5, 4)
}

// IsPrimary
func (m InterprocessSynchronizationRemoteStatePointer) IsPrimary() bool {
	return message.BoolFixed(m.Ptr, 5, 4)
}

// IsPrimary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) IsPrimary() bool {
	return message.BoolFixed(m.Ptr, 5, 4)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteState) IsSecondary() bool {
	return message.BoolFixed(m.Unsafe(), 6, 5)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteStateBuilder) IsSecondary() bool {
	return message.BoolFixed(m.Unsafe(), 6, 5)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteStatePointer) IsSecondary() bool {
	return message.BoolFixed(m.Ptr, 6, 5)
}

// IsSecondary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) IsSecondary() bool {
	return message.BoolFixed(m.Ptr, 6, 5)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteState) IsSecondaryFailoverActive() bool {
	return message.BoolFixed(m.Unsafe(), 7, 6)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStateBuilder) IsSecondaryFailoverActive() bool {
	return message.BoolFixed(m.Unsafe(), 7, 6)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStatePointer) IsSecondaryFailoverActive() bool {
	return message.BoolFixed(m.Ptr, 7, 6)
}

// IsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStatePointerBuilder) IsSecondaryFailoverActive() bool {
	return message.BoolFixed(m.Ptr, 7, 6)
}

// SetIsPrimary
func (m InterprocessSynchronizationRemoteStateBuilder) SetIsPrimary(value bool) {
	message.SetBoolFixed(m.Unsafe(), 5, 4, value)
}

// SetIsPrimary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) SetIsPrimary(value bool) {
	message.SetBoolFixed(m.Ptr, 5, 4, value)
}

// SetIsSecondary
func (m InterprocessSynchronizationRemoteStateBuilder) SetIsSecondary(value bool) {
	message.SetBoolFixed(m.Unsafe(), 6, 5, value)
}

// SetIsSecondary
func (m InterprocessSynchronizationRemoteStatePointerBuilder) SetIsSecondary(value bool) {
	message.SetBoolFixed(m.Ptr, 6, 5, value)
}

// SetIsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStateBuilder) SetIsSecondaryFailoverActive(value bool) {
	message.SetBoolFixed(m.Unsafe(), 7, 6, value)
}

// SetIsSecondaryFailoverActive
func (m InterprocessSynchronizationRemoteStatePointerBuilder) SetIsSecondaryFailoverActive(value bool) {
	message.SetBoolFixed(m.Ptr, 7, 6, value)
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
