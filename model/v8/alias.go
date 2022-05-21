// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-16 08:00:54.519198 +0800 WITA m=+0.055446543

package v8

type (
	// DateTime This is a 64-bit integer UNIX time value.
	//
	// This is the number of seconds since the UNIX epoch (January 1, 1970, 00:00:00
	// UTC).
	//
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	//
	// In the case of Google Protocol Buffer encoding the equivalent data type
	// used is sfixed64.
	DateTime int64

	// DateTime4Byte This is a 32 bit integer UNIX time value.
	//
	// This is the number of seconds since the UNIX epoch (January 1, 1970, 00:00:00
	// UTC).
	//
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	//
	// In the case of Google Protocol Buffer encoding, the equivalent data type
	// used is sfixed32.
	DateTime4Byte uint32

	// DateTimeWithMilliseconds This is a 64-bit floating-point UNIX time value.
	//
	// The integer portion is the number of seconds since the UNIX epoch (January
	// 1, 1970, 00:00:00 UTC).
	//
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	//
	// The portion of this value to the right of the decimal point is the optional
	// number of milliseconds. Where one millisecond equals .001.
	//
	// In the case of Google Protocol Buffer encoding, the equivalent data type
	// used is double.
	DateTimeWithMilliseconds float64

	DateTimeWithMillisecondsInt int64

	// DateTimeWithMicrosecondsInt This is a 64-bit integer UNIX time value.
	//
	// The integer portion is the number of microseconds since the UNIX epoch
	// (January 1, 1970, 00:00:00 UTC).
	//
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	// With the DTC Protocol, the time zone is always UTC for Date-Time values.
	DateTimeWithMicrosecondsInt int64
)
