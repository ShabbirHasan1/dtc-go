package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const RequestNumCurrentClientConnectionsSize = 22

const RequestNumCurrentClientConnectionsFixedSize = 48

// {
//     Size             = RequestNumCurrentClientConnectionsSize  (22)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     BaseSize         = RequestNumCurrentClientConnectionsSize  (22)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
var _RequestNumCurrentClientConnectionsDefault = []byte{22, 0, 123, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size             = RequestNumCurrentClientConnectionsFixedSize  (48)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
var _RequestNumCurrentClientConnectionsFixedDefault = []byte{48, 0, 123, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type RequestNumCurrentClientConnections struct {
	message.VLS
}

type RequestNumCurrentClientConnectionsBuilder struct {
	message.VLSBuilder
}

type RequestNumCurrentClientConnectionsFixed struct {
	message.Fixed
}

type RequestNumCurrentClientConnectionsFixedBuilder struct {
	message.Fixed
}

type RequestNumCurrentClientConnectionsPointer struct {
	message.VLSPointer
}

type RequestNumCurrentClientConnectionsPointerBuilder struct {
	message.VLSPointerBuilder
}

type RequestNumCurrentClientConnectionsFixedPointer struct {
	message.FixedPointer
}

type RequestNumCurrentClientConnectionsFixedPointerBuilder struct {
	message.FixedPointer
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

func NewRequestNumCurrentClientConnectionsFixedFrom(b []byte) RequestNumCurrentClientConnectionsFixed {
	return RequestNumCurrentClientConnectionsFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapRequestNumCurrentClientConnectionsFixed(b []byte) RequestNumCurrentClientConnectionsFixed {
	return RequestNumCurrentClientConnectionsFixed{Fixed: message.WrapFixed(b)}
}

func NewRequestNumCurrentClientConnectionsFixed() RequestNumCurrentClientConnectionsFixedBuilder {
	a := RequestNumCurrentClientConnectionsFixedBuilder{message.NewFixed(48)}
	a.Unsafe().SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
	return a
}

func AllocRequestNumCurrentClientConnections() RequestNumCurrentClientConnectionsPointerBuilder {
	a := RequestNumCurrentClientConnectionsPointerBuilder{message.AllocVLSBuilder(22)}
	a.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsDefault)
	return a
}

func AllocRequestNumCurrentClientConnectionsFrom(b []byte) RequestNumCurrentClientConnectionsPointer {
	return RequestNumCurrentClientConnectionsPointer{VLSPointer: message.AllocVLSFrom(b)}
}

func AllocRequestNumCurrentClientConnectionsFixed() RequestNumCurrentClientConnectionsFixedPointerBuilder {
	a := RequestNumCurrentClientConnectionsFixedPointerBuilder{message.AllocFixed(48)}
	a.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
	return a
}

func AllocRequestNumCurrentClientConnectionsFixedFrom(b []byte) RequestNumCurrentClientConnectionsFixedPointer {
	return RequestNumCurrentClientConnectionsFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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

// Clear
// {
//     Size             = RequestNumCurrentClientConnectionsFixedSize  (48)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
func (m RequestNumCurrentClientConnectionsFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
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
func (m RequestNumCurrentClientConnectionsPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsDefault)
}

// Clear
// {
//     Size             = RequestNumCurrentClientConnectionsFixedSize  (48)
//     Type             = REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS  (10107)
//     RequestID        = 0
//     Username         = ""
//     DeviceIdentifier = 0
// }
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _RequestNumCurrentClientConnectionsFixedDefault)
}

// ToBuilder
func (m RequestNumCurrentClientConnections) ToBuilder() RequestNumCurrentClientConnectionsBuilder {
	return RequestNumCurrentClientConnectionsBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m RequestNumCurrentClientConnectionsFixed) ToBuilder() RequestNumCurrentClientConnectionsFixedBuilder {
	return RequestNumCurrentClientConnectionsFixedBuilder{m.Fixed}
}

// ToBuilder
func (m RequestNumCurrentClientConnectionsPointer) ToBuilder() RequestNumCurrentClientConnectionsPointerBuilder {
	return RequestNumCurrentClientConnectionsPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m RequestNumCurrentClientConnectionsFixedPointer) ToBuilder() RequestNumCurrentClientConnectionsFixedPointerBuilder {
	return RequestNumCurrentClientConnectionsFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m RequestNumCurrentClientConnectionsBuilder) Finish() RequestNumCurrentClientConnections {
	return RequestNumCurrentClientConnections{m.VLSBuilder.Finish()}
}

// Finish
func (m RequestNumCurrentClientConnectionsFixedBuilder) Finish() RequestNumCurrentClientConnectionsFixed {
	return RequestNumCurrentClientConnectionsFixed{m.Fixed}
}

// Finish
func (m *RequestNumCurrentClientConnectionsPointerBuilder) Finish() RequestNumCurrentClientConnectionsPointer {
	return RequestNumCurrentClientConnectionsPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *RequestNumCurrentClientConnectionsFixedPointerBuilder) Finish() RequestNumCurrentClientConnectionsFixedPointer {
	return RequestNumCurrentClientConnectionsFixedPointer{m.FixedPointer}
}

// RequestID
func (m RequestNumCurrentClientConnections) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m RequestNumCurrentClientConnectionsBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Unsafe(), 10, 6)
}

// RequestID
func (m RequestNumCurrentClientConnectionsPointer) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// RequestID
func (m RequestNumCurrentClientConnectionsPointerBuilder) RequestID() uint32 {
	return message.Uint32VLS(m.Ptr, 10, 6)
}

// Username
func (m RequestNumCurrentClientConnections) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m RequestNumCurrentClientConnectionsBuilder) Username() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// Username
func (m RequestNumCurrentClientConnectionsPointer) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// Username
func (m RequestNumCurrentClientConnectionsPointerBuilder) Username() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnections) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Unsafe(), 22, 14)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsBuilder) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Unsafe(), 22, 14)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsPointer) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Ptr, 22, 14)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsPointerBuilder) DeviceIdentifier() int64 {
	return message.Int64VLS(m.Ptr, 22, 14)
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixed) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixedBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixedPointer) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// RequestID
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) RequestID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// Username
func (m RequestNumCurrentClientConnectionsFixed) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username
func (m RequestNumCurrentClientConnectionsFixedBuilder) Username() string {
	return message.StringFixed(m.Unsafe(), 40, 8)
}

// Username
func (m RequestNumCurrentClientConnectionsFixedPointer) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// Username
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) Username() string {
	return message.StringFixed(m.Ptr, 40, 8)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixed) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Unsafe(), 48, 40)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedBuilder) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Unsafe(), 48, 40)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedPointer) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Ptr, 48, 40)
}

// DeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) DeviceIdentifier() int64 {
	return message.Int64Fixed(m.Ptr, 48, 40)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Unsafe(), 10, 6, value)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32VLS(m.Ptr, 10, 6, value)
}

// SetUsername
func (m *RequestNumCurrentClientConnectionsBuilder) SetUsername(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetUsername
func (m *RequestNumCurrentClientConnectionsPointerBuilder) SetUsername(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64VLS(m.Unsafe(), 22, 14, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsPointerBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64VLS(m.Ptr, 22, 14, value)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsFixedBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetRequestID
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) SetRequestID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetUsername
func (m RequestNumCurrentClientConnectionsFixedBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Unsafe(), 40, 8, value)
}

// SetUsername
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) SetUsername(value string) {
	message.SetStringFixed(m.Ptr, 40, 8, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64Fixed(m.Unsafe(), 48, 40, value)
}

// SetDeviceIdentifier
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) SetDeviceIdentifier(value int64) {
	message.SetInt64Fixed(m.Ptr, 48, 40, value)
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

// Copy
func (m RequestNumCurrentClientConnectionsFixed) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsFixedBuilder) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixed) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixedBuilder) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixed) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixedBuilder) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixed) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixedBuilder) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsPointer) Copy(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsPointerBuilder) Copy(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsPointer) CopyPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsPointerBuilder) CopyPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsPointer) CopyTo(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsPointerBuilder) CopyTo(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsPointer) CopyToPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsPointerBuilder) CopyToPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsFixedPointer) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// Copy
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) Copy(to RequestNumCurrentClientConnectionsFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixedPointer) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyPointer
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) CopyPointer(to RequestNumCurrentClientConnectionsFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixedPointer) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyTo
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) CopyTo(to RequestNumCurrentClientConnectionsBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixedPointer) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}

// CopyToPointer
func (m RequestNumCurrentClientConnectionsFixedPointerBuilder) CopyToPointer(to RequestNumCurrentClientConnectionsPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetUsername(m.Username())
	to.SetDeviceIdentifier(m.DeviceIdentifier())
}
