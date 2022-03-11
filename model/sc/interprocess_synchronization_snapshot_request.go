package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const InterprocessSynchronizationSnapshotRequestSize = 8

// {
//     Size      = InterprocessSynchronizationSnapshotRequestSize  (8)
//     Type      = INTERPROCESS_SYNCHRONIZATION_SNAPSHOT_REQUEST  (10133)
//     RequestID = 0
// }
var _InterprocessSynchronizationSnapshotRequestDefault = []byte{8, 0, 149, 39, 0, 0, 0, 0}

type InterprocessSynchronizationSnapshotRequest struct {
	message.Fixed
}

type InterprocessSynchronizationSnapshotRequestBuilder struct {
	message.Fixed
}

type InterprocessSynchronizationSnapshotRequestPointer struct {
	message.FixedPointer
}

type InterprocessSynchronizationSnapshotRequestPointerBuilder struct {
	message.FixedPointer
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
func (m InterprocessSynchronizationSnapshotRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _InterprocessSynchronizationSnapshotRequestDefault)
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
func (m InterprocessSynchronizationSnapshotRequest) ToBuilder() InterprocessSynchronizationSnapshotRequestBuilder {
	return InterprocessSynchronizationSnapshotRequestBuilder{m.Fixed}
}

// ToBuilder
func (m InterprocessSynchronizationSnapshotRequestPointer) ToBuilder() InterprocessSynchronizationSnapshotRequestPointerBuilder {
	return InterprocessSynchronizationSnapshotRequestPointerBuilder{m.FixedPointer}
}

// Finish
func (m InterprocessSynchronizationSnapshotRequestBuilder) Finish() InterprocessSynchronizationSnapshotRequest {
	return InterprocessSynchronizationSnapshotRequest{m.Fixed}
}

// Finish
func (m *InterprocessSynchronizationSnapshotRequestPointerBuilder) Finish() InterprocessSynchronizationSnapshotRequestPointer {
	return InterprocessSynchronizationSnapshotRequestPointer{m.FixedPointer}
}

// RequestID
func (m InterprocessSynchronizationSnapshotRequest) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m InterprocessSynchronizationSnapshotRequestBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
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
func (m InterprocessSynchronizationSnapshotRequestBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m InterprocessSynchronizationSnapshotRequestPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
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
