package model

import (
	"github.com/moontrade/dtc-go/message"
)

type WriteIntradayDataFileSessionValuePointer struct {
	message.FixedPointer
}

type WriteIntradayDataFileSessionValuePointerBuilder struct {
	message.FixedPointer
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
func (m WriteIntradayDataFileSessionValuePointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _WriteIntradayDataFileSessionValueDefault)
}

// ToBuilder
func (m WriteIntradayDataFileSessionValuePointer) ToBuilder() WriteIntradayDataFileSessionValuePointerBuilder {
	return WriteIntradayDataFileSessionValuePointerBuilder{m.FixedPointer}
}

// Finish
func (m *WriteIntradayDataFileSessionValuePointerBuilder) Finish() WriteIntradayDataFileSessionValuePointer {
	return WriteIntradayDataFileSessionValuePointer{m.FixedPointer}
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
func (m WriteIntradayDataFileSessionValuePointer) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Ptr, 12, 8))
}

// ValueType
func (m WriteIntradayDataFileSessionValuePointerBuilder) ValueType() WriteIntradayDataFileSessionValueTypeEnum {
	return WriteIntradayDataFileSessionValueTypeEnum(message.Int32Fixed(m.Ptr, 12, 8))
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
func (m WriteIntradayDataFileSessionValuePointer) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 28, 20))
}

// Date
func (m WriteIntradayDataFileSessionValuePointerBuilder) Date() DateTime {
	return DateTime(message.Int64Fixed(m.Ptr, 28, 20))
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
func (m WriteIntradayDataFileSessionValuePointer) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 44, 36)
}

// Volume
func (m WriteIntradayDataFileSessionValuePointerBuilder) Volume() float64 {
	return message.Float64Fixed(m.Ptr, 44, 36)
}

// SetSymbolID
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetSymbolID(value uint32) {
	message.SetUint32Fixed(m.Ptr, 8, 4, value)
}

// SetValueType
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetValueType(value WriteIntradayDataFileSessionValueTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 12, 8, int32(value))
}

// SetDateTime
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetDateTime(value DateTimeWithMicrosecondsInt) {
	message.SetInt64Fixed(m.Ptr, 20, 12, int64(value))
}

// SetDate
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetDate(value DateTime) {
	message.SetInt64Fixed(m.Ptr, 28, 20, int64(value))
}

// SetPrice
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetPrice(value float64) {
	message.SetFloat64Fixed(m.Ptr, 36, 28, value)
}

// SetVolume
func (m WriteIntradayDataFileSessionValuePointerBuilder) SetVolume(value float64) {
	message.SetFloat64Fixed(m.Ptr, 44, 36, value)
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
