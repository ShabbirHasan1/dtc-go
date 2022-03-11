package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const SCConfigurationSynchronizationSize = 31

// {
//     Size                            = SCConfigurationSynchronizationSize  (31)
//     Type                            = SC_CONFIGURATION_SYNCHRONIZATION  (10109)
//     BaseSize                        = SCConfigurationSynchronizationSize  (31)
//     MessageID                       = 0
//     CurrentInboundSequenceNumber    = 0
//     CurrentOutboundSequenceNumber   = 0
//     CurrentInternalOrderID          = 0
//     IsSnapshot                      = false
//     LastReceivedExecutionIdentifier = ""
// }
var _SCConfigurationSynchronizationDefault = []byte{31, 0, 125, 39, 31, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type SCConfigurationSynchronization struct {
	message.VLS
}

type SCConfigurationSynchronizationBuilder struct {
	message.VLSBuilder
}

type SCConfigurationSynchronizationPointer struct {
	message.VLSPointer
}

type SCConfigurationSynchronizationPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewSCConfigurationSynchronizationFrom(b []byte) SCConfigurationSynchronization {
	return SCConfigurationSynchronization{VLS: message.NewVLSFrom(b)}
}

func WrapSCConfigurationSynchronization(b []byte) SCConfigurationSynchronization {
	return SCConfigurationSynchronization{VLS: message.WrapVLS(b)}
}

func NewSCConfigurationSynchronization() SCConfigurationSynchronizationBuilder {
	a := SCConfigurationSynchronizationBuilder{message.NewVLSBuilder(31)}
	a.Unsafe().SetBytes(0, _SCConfigurationSynchronizationDefault)
	return a
}

func AllocSCConfigurationSynchronization() SCConfigurationSynchronizationPointerBuilder {
	a := SCConfigurationSynchronizationPointerBuilder{message.AllocVLSBuilder(31)}
	a.Ptr.SetBytes(0, _SCConfigurationSynchronizationDefault)
	return a
}

func AllocSCConfigurationSynchronizationFrom(b []byte) SCConfigurationSynchronizationPointer {
	return SCConfigurationSynchronizationPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size                            = SCConfigurationSynchronizationSize  (31)
//     Type                            = SC_CONFIGURATION_SYNCHRONIZATION  (10109)
//     BaseSize                        = SCConfigurationSynchronizationSize  (31)
//     MessageID                       = 0
//     CurrentInboundSequenceNumber    = 0
//     CurrentOutboundSequenceNumber   = 0
//     CurrentInternalOrderID          = 0
//     IsSnapshot                      = false
//     LastReceivedExecutionIdentifier = ""
// }
func (m SCConfigurationSynchronizationBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SCConfigurationSynchronizationDefault)
}

// Clear
// {
//     Size                            = SCConfigurationSynchronizationSize  (31)
//     Type                            = SC_CONFIGURATION_SYNCHRONIZATION  (10109)
//     BaseSize                        = SCConfigurationSynchronizationSize  (31)
//     MessageID                       = 0
//     CurrentInboundSequenceNumber    = 0
//     CurrentOutboundSequenceNumber   = 0
//     CurrentInternalOrderID          = 0
//     IsSnapshot                      = false
//     LastReceivedExecutionIdentifier = ""
// }
func (m SCConfigurationSynchronizationPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SCConfigurationSynchronizationDefault)
}

// ToBuilder
func (m SCConfigurationSynchronization) ToBuilder() SCConfigurationSynchronizationBuilder {
	return SCConfigurationSynchronizationBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SCConfigurationSynchronizationPointer) ToBuilder() SCConfigurationSynchronizationPointerBuilder {
	return SCConfigurationSynchronizationPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m SCConfigurationSynchronizationBuilder) Finish() SCConfigurationSynchronization {
	return SCConfigurationSynchronization{m.VLSBuilder.Finish()}
}

// Finish
func (m *SCConfigurationSynchronizationPointerBuilder) Finish() SCConfigurationSynchronizationPointer {
	return SCConfigurationSynchronizationPointer{m.VLSPointerBuilder.Finish()}
}

// MessageID
func (m SCConfigurationSynchronization) MessageID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// MessageID
func (m SCConfigurationSynchronizationBuilder) MessageID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// MessageID
func (m SCConfigurationSynchronizationPointer) MessageID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// MessageID
func (m SCConfigurationSynchronizationPointerBuilder) MessageID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// CurrentInboundSequenceNumber
func (m SCConfigurationSynchronization) CurrentInboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Unsafe(), 14, 10)
}

// CurrentInboundSequenceNumber
func (m SCConfigurationSynchronizationBuilder) CurrentInboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Unsafe(), 14, 10)
}

// CurrentInboundSequenceNumber
func (m SCConfigurationSynchronizationPointer) CurrentInboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Ptr, 14, 10)
}

// CurrentInboundSequenceNumber
func (m SCConfigurationSynchronizationPointerBuilder) CurrentInboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Ptr, 14, 10)
}

// CurrentOutboundSequenceNumber
func (m SCConfigurationSynchronization) CurrentOutboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Unsafe(), 18, 14)
}

// CurrentOutboundSequenceNumber
func (m SCConfigurationSynchronizationBuilder) CurrentOutboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Unsafe(), 18, 14)
}

// CurrentOutboundSequenceNumber
func (m SCConfigurationSynchronizationPointer) CurrentOutboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Ptr, 18, 14)
}

// CurrentOutboundSequenceNumber
func (m SCConfigurationSynchronizationPointerBuilder) CurrentOutboundSequenceNumber() uint32 {
	return message.Uint32VLS(m.Ptr, 18, 14)
}

// CurrentInternalOrderID
func (m SCConfigurationSynchronization) CurrentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 26, 18)
}

// CurrentInternalOrderID
func (m SCConfigurationSynchronizationBuilder) CurrentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Unsafe(), 26, 18)
}

// CurrentInternalOrderID
func (m SCConfigurationSynchronizationPointer) CurrentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 26, 18)
}

// CurrentInternalOrderID
func (m SCConfigurationSynchronizationPointerBuilder) CurrentInternalOrderID() uint64 {
	return message.Uint64VLS(m.Ptr, 26, 18)
}

// IsSnapshot
func (m SCConfigurationSynchronization) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 27, 26)
}

// IsSnapshot
func (m SCConfigurationSynchronizationBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Unsafe(), 27, 26)
}

// IsSnapshot
func (m SCConfigurationSynchronizationPointer) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 27, 26)
}

// IsSnapshot
func (m SCConfigurationSynchronizationPointerBuilder) IsSnapshot() bool {
	return message.BoolVLS(m.Ptr, 27, 26)
}

// LastReceivedExecutionIdentifier
func (m SCConfigurationSynchronization) LastReceivedExecutionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 31, 27)
}

// LastReceivedExecutionIdentifier
func (m SCConfigurationSynchronizationBuilder) LastReceivedExecutionIdentifier() string {
	return message.StringVLS(m.Unsafe(), 31, 27)
}

// LastReceivedExecutionIdentifier
func (m SCConfigurationSynchronizationPointer) LastReceivedExecutionIdentifier() string {
	return message.StringVLS(m.Ptr, 31, 27)
}

// LastReceivedExecutionIdentifier
func (m SCConfigurationSynchronizationPointerBuilder) LastReceivedExecutionIdentifier() string {
	return message.StringVLS(m.Ptr, 31, 27)
}

// SetMessageID
func (m SCConfigurationSynchronizationBuilder) SetMessageID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetMessageID
func (m SCConfigurationSynchronizationPointerBuilder) SetMessageID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetCurrentInboundSequenceNumber
func (m SCConfigurationSynchronizationBuilder) SetCurrentInboundSequenceNumber(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 14, 10, value)
}

// SetCurrentInboundSequenceNumber
func (m SCConfigurationSynchronizationPointerBuilder) SetCurrentInboundSequenceNumber(value uint32) {
	message.SetUint32VLS(m.Ptr, 14, 10, value)
}

// SetCurrentOutboundSequenceNumber
func (m SCConfigurationSynchronizationBuilder) SetCurrentOutboundSequenceNumber(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 18, 14, value)
}

// SetCurrentOutboundSequenceNumber
func (m SCConfigurationSynchronizationPointerBuilder) SetCurrentOutboundSequenceNumber(value uint32) {
	message.SetUint32VLS(m.Ptr, 18, 14, value)
}

// SetCurrentInternalOrderID
func (m SCConfigurationSynchronizationBuilder) SetCurrentInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Unsafe(), 26, 18, value)
}

// SetCurrentInternalOrderID
func (m SCConfigurationSynchronizationPointerBuilder) SetCurrentInternalOrderID(value uint64) {
	message.SetUint64VLS(m.Ptr, 26, 18, value)
}

// SetIsSnapshot
func (m SCConfigurationSynchronizationBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Unsafe(), 27, 26, value)
}

// SetIsSnapshot
func (m SCConfigurationSynchronizationPointerBuilder) SetIsSnapshot(value bool) {
	message.SetBoolVLS(m.Ptr, 27, 26, value)
}

// SetLastReceivedExecutionIdentifier
func (m *SCConfigurationSynchronizationBuilder) SetLastReceivedExecutionIdentifier(value string) {
	message.SetStringVLS(&m.VLSBuilder, 31, 27, value)
}

// SetLastReceivedExecutionIdentifier
func (m *SCConfigurationSynchronizationPointerBuilder) SetLastReceivedExecutionIdentifier(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 31, 27, value)
}

// Copy
func (m SCConfigurationSynchronization) Copy(to SCConfigurationSynchronizationBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}

// Copy
func (m SCConfigurationSynchronizationBuilder) Copy(to SCConfigurationSynchronizationBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}

// CopyPointer
func (m SCConfigurationSynchronization) CopyPointer(to SCConfigurationSynchronizationPointerBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}

// CopyPointer
func (m SCConfigurationSynchronizationBuilder) CopyPointer(to SCConfigurationSynchronizationPointerBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}

// Copy
func (m SCConfigurationSynchronizationPointer) Copy(to SCConfigurationSynchronizationBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}

// Copy
func (m SCConfigurationSynchronizationPointerBuilder) Copy(to SCConfigurationSynchronizationBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}

// CopyPointer
func (m SCConfigurationSynchronizationPointer) CopyPointer(to SCConfigurationSynchronizationPointerBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}

// CopyPointer
func (m SCConfigurationSynchronizationPointerBuilder) CopyPointer(to SCConfigurationSynchronizationPointerBuilder) {
	to.SetMessageID(m.MessageID())
	to.SetCurrentInboundSequenceNumber(m.CurrentInboundSequenceNumber())
	to.SetCurrentOutboundSequenceNumber(m.CurrentOutboundSequenceNumber())
	to.SetCurrentInternalOrderID(m.CurrentInternalOrderID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetLastReceivedExecutionIdentifier(m.LastReceivedExecutionIdentifier())
}
