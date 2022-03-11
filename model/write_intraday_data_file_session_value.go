package model

import (
	"github.com/moontrade/dtc-go/message"
)

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

const WriteIntradayDataFileSessionValueSize = 44

type WriteIntradayDataFileSessionValue struct {
	message.Fixed
}

type WriteIntradayDataFileSessionValueBuilder struct {
	message.Fixed
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

// ToBuilder
func (m WriteIntradayDataFileSessionValue) ToBuilder() WriteIntradayDataFileSessionValueBuilder {
	return WriteIntradayDataFileSessionValueBuilder{m.Fixed}
}

// Finish
func (m WriteIntradayDataFileSessionValueBuilder) Finish() WriteIntradayDataFileSessionValue {
	return WriteIntradayDataFileSessionValue{m.Fixed}
}

// SymbolID
func (m WriteIntradayDataFileSessionValue) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// SymbolID
func (m WriteIntradayDataFileSessionValueBuilder) SymbolID() uint32 {
	return message.Uint32Fixed(m.Unsafe(), 8, 4)
}

// ValueType
func (m WriteIntradayDataFileSessionValue) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// ValueType
func (m WriteIntradayDataFileSessionValueBuilder) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Unsafe(), 12, 8))
}

// DateTime
func (m WriteIntradayDataFileSessionValue) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 20, 12))
}

// DateTime
func (m WriteIntradayDataFileSessionValueBuilder) DateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(message.Int64Fixed(m.Unsafe(), 20, 12))
}

// Date
func (m WriteIntradayDataFileSessionValue) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 28, 20))
}

// Date
func (m WriteIntradayDataFileSessionValueBuilder) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Unsafe(), 28, 20))
}

// Price
func (m WriteIntradayDataFileSessionValue) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 36, 28)
}

// Price
func (m WriteIntradayDataFileSessionValueBuilder) Price() float64 {
	return message.Float64Fixed(m.Unsafe(), 36, 28)
}

// Volume
func (m WriteIntradayDataFileSessionValue) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 44, 36)
}

// Volume
func (m WriteIntradayDataFileSessionValueBuilder) Volume() float64 {
	return message.Float64Fixed(m.Unsafe(), 44, 36)
}

// SetSymbolID
func (m WriteIntradayDataFileSessionValueBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Unsafe(), 8, 4, value)
}

// SetValueType
func (m WriteIntradayDataFileSessionValueBuilder) SetValueType(value WriteIntradayDataFileSessionValueTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 12, 8, int32(value))
}

// SetDateTime
func (m WriteIntradayDataFileSessionValueBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Unsafe(), 20, 12, int64(value))
}

// SetDate
func (m WriteIntradayDataFileSessionValueBuilder) SetDate(value DateTime) {
	message.SetInt64Fixed(m.Unsafe(), 28, 20, int64(value))
}

// SetPrice
func (m WriteIntradayDataFileSessionValueBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 36, 28, value)
}

// SetVolume
func (m WriteIntradayDataFileSessionValueBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Unsafe(), 44, 36, value)
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
