package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const WriteIntradayDataFileSessionValueSize = 44

// {
//     Size      = WriteIntradayDataFileSessionValueSize  (44)
//     Type      = WRITE_INTRADAY_DATA_FILE_SESSION_VALUE  (10140)
//     SymbolID  = 0
//     ValueType = INTRADAY_DATA_FILE_SESSION_VALUE_UNSET  (0)
//     DateTime  = 0
//     Date      = 0
//     Price     = 0
//     Volume    = 0
// }
var _WriteIntradayDataFileSessionValueDefault = []byte{44, 0, 156, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type WriteIntradayDataFileSessionValue struct {
	message.Fixed
}

type WriteIntradayDataFileSessionValueBuilder struct {
	message.Fixed
}

type WriteIntradayDataFileSessionValuePointer struct {
	message.FixedPointer
}

type WriteIntradayDataFileSessionValuePointerBuilder struct {
	message.FixedPointer
}

func NewWriteIntradayDataFileSessionValueFrom(b []byte) WriteIntradayDataFileSessionValue {
	return WriteIntradayDataFileSessionValue{Fixed: message.NewFixedFrom(b)}
}

func WrapWriteIntradayDataFileSessionValue(b []byte) WriteIntradayDataFileSessionValue {
	return WriteIntradayDataFileSessionValue{Fixed: message.WrapFixed(b)}
}

func NewWriteIntradayDataFileSessionValue() WriteIntradayDataFileSessionValueBuilder {
	a := WriteIntradayDataFileSessionValueBuilder{message.NewFixed(44)}
	a.Unsafe().SetBytes(0, _WriteIntradayDataFileSessionValueDefault)
	return a
}

func AllocWriteIntradayDataFileSessionValue() WriteIntradayDataFileSessionValuePointerBuilder {
	a := WriteIntradayDataFileSessionValuePointerBuilder{message.AllocFixed(44)}
	a.Ptr.SetBytes(0, _WriteIntradayDataFileSessionValueDefault)
	return a
}

func AllocWriteIntradayDataFileSessionValueFrom(b []byte) WriteIntradayDataFileSessionValuePointer {
	return WriteIntradayDataFileSessionValuePointer{FixedPointer: message.AllocFixedFrom(b)}
}

// Clear
// {
//     Size      = WriteIntradayDataFileSessionValueSize  (44)
//     Type      = WRITE_INTRADAY_DATA_FILE_SESSION_VALUE  (10140)
//     SymbolID  = 0
//     ValueType = INTRADAY_DATA_FILE_SESSION_VALUE_UNSET  (0)
//     DateTime  = 0
//     Date      = 0
//     Price     = 0
//     Volume    = 0
// }
func (m WriteIntradayDataFileSessionValueBuilder) Clear() {
	m.Unsafe().SetBytes(0, _WriteIntradayDataFileSessionValueDefault)
}

// Clear
// {
//     Size      = WriteIntradayDataFileSessionValueSize  (44)
//     Type      = WRITE_INTRADAY_DATA_FILE_SESSION_VALUE  (10140)
//     SymbolID  = 0
//     ValueType = INTRADAY_DATA_FILE_SESSION_VALUE_UNSET  (0)
//     DateTime  = 0
//     Date      = 0
//     Price     = 0
//     Volume    = 0
// }
func (m WriteIntradayDataFileSessionValuePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _WriteIntradayDataFileSessionValueDefault)
}

// ToBuilder
func (m WriteIntradayDataFileSessionValue) ToBuilder() WriteIntradayDataFileSessionValueBuilder {
	return WriteIntradayDataFileSessionValueBuilder{m.Fixed}
}

// ToBuilder
func (m WriteIntradayDataFileSessionValuePointer) ToBuilder() WriteIntradayDataFileSessionValuePointerBuilder {
	return WriteIntradayDataFileSessionValuePointerBuilder{m.FixedPointer}
}

// Finish
func (m WriteIntradayDataFileSessionValueBuilder) Finish() WriteIntradayDataFileSessionValue {
	return WriteIntradayDataFileSessionValue{m.Fixed}
}

// Finish
func (m *WriteIntradayDataFileSessionValuePointerBuilder) Finish() WriteIntradayDataFileSessionValuePointer {
	return WriteIntradayDataFileSessionValuePointer{m.FixedPointer}
}

// SymbolID
func (m WriteIntradayDataFileSessionValue) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m WriteIntradayDataFileSessionValueBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m WriteIntradayDataFileSessionValuePointer) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// SymbolID
func (m WriteIntradayDataFileSessionValuePointerBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Ptr, 8, 4)
}

// ValueType
func (m WriteIntradayDataFileSessionValue) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// ValueType
func (m WriteIntradayDataFileSessionValueBuilder) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// ValueType
func (m WriteIntradayDataFileSessionValuePointer) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// ValueType
func (m WriteIntradayDataFileSessionValuePointerBuilder) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// DateTime
func (m WriteIntradayDataFileSessionValue) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 20, 12))
}

// DateTime
func (m WriteIntradayDataFileSessionValueBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 20, 12))
}

// DateTime
func (m WriteIntradayDataFileSessionValuePointer) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 20, 12))
}

// DateTime
func (m WriteIntradayDataFileSessionValuePointerBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Ptr, 20, 12))
}

// Date
func (m WriteIntradayDataFileSessionValue) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 28, 20))
}

// Date
func (m WriteIntradayDataFileSessionValueBuilder) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 28, 20))
}

// Date
func (m WriteIntradayDataFileSessionValuePointer) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 28, 20))
}

// Date
func (m WriteIntradayDataFileSessionValuePointerBuilder) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 28, 20))
}

// Price
func (m WriteIntradayDataFileSessionValue) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 36, 28)
}

// Price
func (m WriteIntradayDataFileSessionValueBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 36, 28)
}

// Price
func (m WriteIntradayDataFileSessionValuePointer) Price() float64 {
	return message.Float64Fixed(m.Ptr, 36, 28)
}

// Price
func (m WriteIntradayDataFileSessionValuePointerBuilder) Price() float64 {
	return message.Float64Fixed(m.Ptr, 36, 28)
}

// Volume
func (m WriteIntradayDataFileSessionValue) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 44, 36)
}

// Volume
func (m WriteIntradayDataFileSessionValueBuilder) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 44, 36)
}

// Volume
func (m WriteIntradayDataFileSessionValuePointer) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 44, 36)
}

// Volume
func (m WriteIntradayDataFileSessionValuePointerBuilder) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 44, 36)
}

// SetSymbolID
func (m WriteIntradayDataFileSessionValueBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSymbolID
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetValueType
func (m WriteIntradayDataFileSessionValueBuilder) SetValueType(value WriteIntradayDataFileSessionValueTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetValueType
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetValueType(value WriteIntradayDataFileSessionValueTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetDateTime
func (m WriteIntradayDataFileSessionValueBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 20, 12, int64(value))
}

// SetDateTime
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 20, 12, int64(value))
}

// SetDate
func (m WriteIntradayDataFileSessionValueBuilder) SetDate(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 28, 20, int64(value))
}

// SetDate
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetDate(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 28, 20, int64(value))
}

// SetPrice
func (m WriteIntradayDataFileSessionValueBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 36, 28, value)
}

// SetPrice
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 36, 28, value)
}

// SetVolume
func (m WriteIntradayDataFileSessionValueBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 44, 36, value)
}

// SetVolume
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 44, 36, value)
}

// Copy
func (m WriteIntradayDataFileSessionValue) Copy(to WriteIntradayDataFileSessionValueBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}

// Copy
func (m WriteIntradayDataFileSessionValueBuilder) Copy(to WriteIntradayDataFileSessionValueBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}

// CopyPointer
func (m WriteIntradayDataFileSessionValue) CopyPointer(to WriteIntradayDataFileSessionValuePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}

// CopyPointer
func (m WriteIntradayDataFileSessionValueBuilder) CopyPointer(to WriteIntradayDataFileSessionValuePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}

// Copy
func (m WriteIntradayDataFileSessionValuePointer) Copy(to WriteIntradayDataFileSessionValueBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}

// Copy
func (m WriteIntradayDataFileSessionValuePointerBuilder) Copy(to WriteIntradayDataFileSessionValueBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}

// CopyPointer
func (m WriteIntradayDataFileSessionValuePointer) CopyPointer(to WriteIntradayDataFileSessionValuePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}

// CopyPointer
func (m WriteIntradayDataFileSessionValuePointerBuilder) CopyPointer(to WriteIntradayDataFileSessionValuePointerBuilder) {
	to.SetSymbolID(m.SymbolID())
	to.SetValueType(m.ValueType())
	to.SetDateTime(m.DateTime())
	to.SetDate(m.Date())
	to.SetPrice(m.Price())
	to.SetVolume(m.Volume())
}
