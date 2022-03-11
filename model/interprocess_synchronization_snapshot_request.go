package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size      = InterprocessSynchronizationSnapshotRequestSize  (8)
//     Type      = INTERPROCESS_SYNCHRONIZATION_SNAPSHOT_REQUEST  (10133)
//     RequestID = 0
// }
var _InterprocessSynchronizationSnapshotRequestDefault = []byte{8, 0, 149, 39, 0, 0, 0, 0}

const InterprocessSynchronizationSnapshotRequestSize = 8

type InterprocessSynchronizationSnapshotRequest struct {
	message.Fixed
}

type InterprocessSynchronizationSnapshotRequestBuilder struct {
	message.Fixed
}

func NewInterprocessSynchronizationSnapshotRequestFrom(b []byte) InterprocessSynchronizationSnapshotRequest {
	return InterprocessSynchronizationSnapshotRequest{Fixed: message.NewFixedFrom(b)}
}

func WrapInterprocessSynchronizationSnapshotRequest(b []byte) InterprocessSynchronizationSnapshotRequest {
	return InterprocessSynchronizationSnapshotRequest{Fixed: message.WrapFixed(b)}
}

func NewInterprocessSynchronizationSnapshotRequest() InterprocessSynchronizationSnapshotRequestBuilder {
	a := InterprocessSynchronizationSnapshotRequestBuilder{message.NewFixed(8)}
	a.Unsafe().SetBytes(0, _InterprocessSynchronizationSnapshotRequestDefault)
	return a
}

// Clear
// {
//     Size      = InterprocessSynchronizationSnapshotRequestSize  (8)
//     Type      = INTERPROCESS_SYNCHRONIZATION_SNAPSHOT_REQUEST  (10133)
//     RequestID = 0
// }
func (m InterprocessSynchronizationSnapshotRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _InterprocessSynchronizationSnapshotRequestDefault)
}

// ToBuilder
func (m InterprocessSynchronizationSnapshotRequest) ToBuilder() InterprocessSynchronizationSnapshotRequestBuilder {
	return InterprocessSynchronizationSnapshotRequestBuilder{m.Fixed}
}

// Finish
func (m InterprocessSynchronizationSnapshotRequestBuilder) Finish() InterprocessSynchronizationSnapshotRequest {
	return InterprocessSynchronizationSnapshotRequest{m.Fixed}
}

// RequestID
func (m InterprocessSynchronizationSnapshotRequest) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m InterprocessSynchronizationSnapshotRequestBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SetRequestID
func (m InterprocessSynchronizationSnapshotRequestBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// Copy
func (m InterprocessSynchronizationSnapshotRequest) Copy(to InterprocessSynchronizationSnapshotRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// Copy
func (m InterprocessSynchronizationSnapshotRequestBuilder) Copy(to InterprocessSynchronizationSnapshotRequestBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m InterprocessSynchronizationSnapshotRequest) CopyPointer(to InterprocessSynchronizationSnapshotRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}

// CopyPointer
func (m InterprocessSynchronizationSnapshotRequestBuilder) CopyPointer(to InterprocessSynchronizationSnapshotRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
}
