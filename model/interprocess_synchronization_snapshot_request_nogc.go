package model

import (
	"github.com/moontrade/dtc-go/message"
)

type InterprocessSynchronizationSnapshotRequestPointer struct {
	message.FixedPointer
}

type InterprocessSynchronizationSnapshotRequestPointerBuilder struct {
	message.FixedPointer
}

func AllocInterprocessSynchronizationSnapshotRequest() InterprocessSynchronizationSnapshotRequestPointerBuilder {
	a := InterprocessSynchronizationSnapshotRequestPointerBuilder{message.AllocFixed(8)}
	a.Ptr.SetBytes(0, _InterprocessSynchronizationSnapshotRequestDefault)
	return a
}

func AllocInterprocessSynchronizationSnapshotRequestFrom(b []byte) InterprocessSynchronizationSnapshotRequestPointer {
	return InterprocessSynchronizationSnapshotRequestPointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size      = InterprocessSynchronizationSnapshotRequestSize  (8)
//     Type      = INTERPROCESS_SYNCHRONIZATION_SNAPSHOT_REQUEST  (10133)
//     RequestID = 0
// }
func (m InterprocessSynchronizationSnapshotRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _InterprocessSynchronizationSnapshotRequestDefault)
}

// ToBuilder
func (m InterprocessSynchronizationSnapshotRequestPointer) ToBuilder() InterprocessSynchronizationSnapshotRequestPointerBuilder {
	return InterprocessSynchronizationSnapshotRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m *InterprocessSynchronizationSnapshotRequestPointerBuilder) Finish() InterprocessSynchronizationSnapshotRequestPointer {
	return InterprocessSynchronizationSnapshotRequestPointer{m.FixedPointer}
}

// RequestID
func (m InterprocessSynchronizationSnapshotRequestPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m InterprocessSynchronizationSnapshotRequestPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SetRequestID
func (m InterprocessSynchronizationSnapshotRequestPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// Copy
func (m InterprocessSynchronizationSnapshotRequestPointer) Copy(to InterprocessSynchronizationSnapshotRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m InterprocessSynchronizationSnapshotRequestPointerBuilder) Copy(to InterprocessSynchronizationSnapshotRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m InterprocessSynchronizationSnapshotRequestPointer) CopyPointer(to InterprocessSynchronizationSnapshotRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m InterprocessSynchronizationSnapshotRequestPointerBuilder) CopyPointer(to InterprocessSynchronizationSnapshotRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}
