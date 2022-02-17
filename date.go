package dtc

type (
	// DateTime Standard UNIX date-time value. In seconds.
	DateTime int64

	// DateTime4Byte This is a 32 bit UNIX date-time value used in messages where compactness is an important consideration.
	// Or, where only the Date is needed. In seconds.
	DateTime4Byte uint32

	// DateTimeWithMilliseconds UNIX date-time value with fractional portion for milliseconds.
	DateTimeWithMilliseconds float64

	// DateTimeWithMillisecondsInt Standard UNIX date-time value in milliseconds.
	DateTimeWithMillisecondsInt int64

	// DateTimeWithMicrosecondsInt Standard UNIX date-time value in microseconds.
	DateTimeWithMicrosecondsInt int64
)

func (d DateTimeWithMilliseconds) ToMicros() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(d*1000000.0 + 0.5)
}

//func DateTimeWithMillisecondsToDateTimeWithMicrosecondsInt(DateTime DateTimeWithMilliseconds) int64 {
//	return int64(DateTime*1000000.0 + 0.5)
//}
