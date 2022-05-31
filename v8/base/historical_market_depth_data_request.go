// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const HistoricalMarketDepthDataRequestSize = 48

const HistoricalMarketDepthDataRequestFixedSize = 112

//     Size                uint16                       = HistoricalMarketDepthDataRequestSize  (48)
//     Type                uint16                       = HISTORICAL_MARKET_DEPTH_DATA_REQUEST  (900)
//     BaseSize            uint16                       = HistoricalMarketDepthDataRequestSize  (48)
//     RequestID           int32                        = 0
//     Symbol              string                       = ""
//     Exchange            string                       = ""
//     StartDateTime       DateTimeWithMicrosecondsInt  = 0
//     EndDateTime         DateTimeWithMicrosecondsInt  = 0
//     UseZLibCompression  bool                         = false
//     Integer_1           uint8                        = 0
var _HistoricalMarketDepthDataRequestDefault = []byte{48, 0, 132, 3, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size                uint16                       = HistoricalMarketDepthDataRequestFixedSize  (112)
//     Type                uint16                       = HISTORICAL_MARKET_DEPTH_DATA_REQUEST  (900)
//     RequestID           int32                        = 0
//     Symbol              string[64]                   = ""
//     Exchange            string[16]                   = ""
//     StartDateTime       DateTimeWithMicrosecondsInt  = 0
//     EndDateTime         DateTimeWithMicrosecondsInt  = 0
//     UseZLibCompression  bool                         = false
//     Integer_1           uint8                        = 0
var _HistoricalMarketDepthDataRequestFixedDefault = []byte{112, 0, 132, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type HistoricalMarketDepthDataRequest struct {
	p message.VLS
}

type HistoricalMarketDepthDataRequestFixed struct {
	p message.Fixed
}

func NewHistoricalMarketDepthDataRequestFrom(b []byte) HistoricalMarketDepthDataRequest {
	return HistoricalMarketDepthDataRequest{p: message.NewVLS(b)}
}

func WrapHistoricalMarketDepthDataRequest(b []byte) HistoricalMarketDepthDataRequest {
	return HistoricalMarketDepthDataRequest{p: message.WrapVLS(b)}
}

func NewHistoricalMarketDepthDataRequest() *HistoricalMarketDepthDataRequest {
	return &HistoricalMarketDepthDataRequest{p: message.NewVLS(_HistoricalMarketDepthDataRequestDefault)}
}

func ParseHistoricalMarketDepthDataRequest(b []byte) (HistoricalMarketDepthDataRequest, error) {
	if len(b) < 6 {
		return HistoricalMarketDepthDataRequest{}, message.ErrShortBuffer
	}
	m := WrapHistoricalMarketDepthDataRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalMarketDepthDataRequest{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return HistoricalMarketDepthDataRequest{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 48 {
		newSize := len(b) + (48 - baseSize)
		if newSize > message.MaxSize {
			return HistoricalMarketDepthDataRequest{}, message.ErrOverflow
		}
		clone := HistoricalMarketDepthDataRequest{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _HistoricalMarketDepthDataRequestDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(48 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(12)
			if offset > 0 {
				clone.p.SetUint16LE(12, offset+shift)
			}
			offset = clone.p.Uint16LE(16)
			if offset > 0 {
				clone.p.SetUint16LE(16, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewHistoricalMarketDepthDataRequestFixedFrom(b []byte) HistoricalMarketDepthDataRequestFixed {
	return HistoricalMarketDepthDataRequestFixed{p: message.NewFixed(b)}
}

func WrapHistoricalMarketDepthDataRequestFixed(b []byte) HistoricalMarketDepthDataRequestFixed {
	return HistoricalMarketDepthDataRequestFixed{p: message.WrapFixed(b)}
}

func NewHistoricalMarketDepthDataRequestFixed() *HistoricalMarketDepthDataRequestFixed {
	return &HistoricalMarketDepthDataRequestFixed{p: message.NewFixed(_HistoricalMarketDepthDataRequestFixedDefault)}
}

func ParseHistoricalMarketDepthDataRequestFixed(b []byte) (HistoricalMarketDepthDataRequestFixed, error) {
	if len(b) < 4 {
		return HistoricalMarketDepthDataRequestFixed{}, message.ErrShortBuffer
	}
	m := WrapHistoricalMarketDepthDataRequestFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalMarketDepthDataRequestFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return HistoricalMarketDepthDataRequestFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 112 {
		clone := *NewHistoricalMarketDepthDataRequestFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _HistoricalMarketDepthDataRequestFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m HistoricalMarketDepthDataRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m HistoricalMarketDepthDataRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m HistoricalMarketDepthDataRequest) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m HistoricalMarketDepthDataRequest) RequestID() int32 {
	return m.p.Int32LE(8)
}

// Symbol
func (m HistoricalMarketDepthDataRequest) Symbol() string {
	return m.p.StringVLS(12)
}

// Exchange
func (m HistoricalMarketDepthDataRequest) Exchange() string {
	return m.p.StringVLS(16)
}

// StartDateTime
func (m HistoricalMarketDepthDataRequest) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(m.p.Int64LE(24))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequest) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(m.p.Int64LE(32))
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequest) UseZLibCompression() bool {
	return m.p.Bool(40)
}

// Integer_1
func (m HistoricalMarketDepthDataRequest) Integer_1() uint8 {
	return m.p.Uint8(41)
}

// Size
func (m HistoricalMarketDepthDataRequestFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m HistoricalMarketDepthDataRequestFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID
func (m HistoricalMarketDepthDataRequestFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// Symbol
func (m HistoricalMarketDepthDataRequestFixed) Symbol() string {
	return m.p.StringFixed(8, 64)
}

// Exchange
func (m HistoricalMarketDepthDataRequestFixed) Exchange() string {
	return m.p.StringFixed(72, 16)
}

// StartDateTime
func (m HistoricalMarketDepthDataRequestFixed) StartDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(m.p.Int64LE(88))
}

// EndDateTime
func (m HistoricalMarketDepthDataRequestFixed) EndDateTime() DateTimeWithMicrosecondsInt {
	return DateTimeWithMicrosecondsInt(m.p.Int64LE(96))
}

// UseZLibCompression
func (m HistoricalMarketDepthDataRequestFixed) UseZLibCompression() bool {
	return m.p.Bool(104)
}

// Integer_1
func (m HistoricalMarketDepthDataRequestFixed) Integer_1() uint8 {
	return m.p.Uint8(105)
}

// SetRequestID
func (m *HistoricalMarketDepthDataRequest) SetRequestID(value int32) *HistoricalMarketDepthDataRequest {
	m.p.SetInt32LE(8, value)
	return m
}

// SetSymbol
func (m *HistoricalMarketDepthDataRequest) SetSymbol(value string) *HistoricalMarketDepthDataRequest {
	m.p.SetStringVLS(12, value)
	return m
}

// SetExchange
func (m *HistoricalMarketDepthDataRequest) SetExchange(value string) *HistoricalMarketDepthDataRequest {
	m.p.SetStringVLS(16, value)
	return m
}

// SetStartDateTime
func (m *HistoricalMarketDepthDataRequest) SetStartDateTime(value DateTimeWithMicrosecondsInt) *HistoricalMarketDepthDataRequest {
	m.p.SetInt64LE(24, int64(value))
	return m
}

// SetEndDateTime
func (m *HistoricalMarketDepthDataRequest) SetEndDateTime(value DateTimeWithMicrosecondsInt) *HistoricalMarketDepthDataRequest {
	m.p.SetInt64LE(32, int64(value))
	return m
}

// SetUseZLibCompression
func (m *HistoricalMarketDepthDataRequest) SetUseZLibCompression(value bool) *HistoricalMarketDepthDataRequest {
	m.p.SetBool(40, value)
	return m
}

// SetInteger_1
func (m *HistoricalMarketDepthDataRequest) SetInteger_1(value uint8) *HistoricalMarketDepthDataRequest {
	m.p.SetUint8(41, value)
	return m
}

// SetRequestID
func (m *HistoricalMarketDepthDataRequestFixed) SetRequestID(value int32) *HistoricalMarketDepthDataRequestFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetSymbol
func (m *HistoricalMarketDepthDataRequestFixed) SetSymbol(value string) *HistoricalMarketDepthDataRequestFixed {
	m.p.SetStringFixed(8, 64, value)
	return m
}

// SetExchange
func (m *HistoricalMarketDepthDataRequestFixed) SetExchange(value string) *HistoricalMarketDepthDataRequestFixed {
	m.p.SetStringFixed(72, 16, value)
	return m
}

// SetStartDateTime
func (m *HistoricalMarketDepthDataRequestFixed) SetStartDateTime(value DateTimeWithMicrosecondsInt) *HistoricalMarketDepthDataRequestFixed {
	m.p.SetInt64LE(88, int64(value))
	return m
}

// SetEndDateTime
func (m *HistoricalMarketDepthDataRequestFixed) SetEndDateTime(value DateTimeWithMicrosecondsInt) *HistoricalMarketDepthDataRequestFixed {
	m.p.SetInt64LE(96, int64(value))
	return m
}

// SetUseZLibCompression
func (m *HistoricalMarketDepthDataRequestFixed) SetUseZLibCompression(value bool) *HistoricalMarketDepthDataRequestFixed {
	m.p.SetBool(104, value)
	return m
}

// SetInteger_1
func (m *HistoricalMarketDepthDataRequestFixed) SetInteger_1(value uint8) *HistoricalMarketDepthDataRequestFixed {
	m.p.SetUint8(105, value)
	return m
}

func (m *HistoricalMarketDepthDataRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *HistoricalMarketDepthDataRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *HistoricalMarketDepthDataRequest) Clone() *HistoricalMarketDepthDataRequest {
	return &HistoricalMarketDepthDataRequest{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *HistoricalMarketDepthDataRequestFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *HistoricalMarketDepthDataRequestFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *HistoricalMarketDepthDataRequestFixed) Clone() *HistoricalMarketDepthDataRequestFixed {
	return &HistoricalMarketDepthDataRequestFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m HistoricalMarketDepthDataRequest) Copy(to HistoricalMarketDepthDataRequest) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequest) CopyTo(to HistoricalMarketDepthDataRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// Copy
func (m HistoricalMarketDepthDataRequestFixed) Copy(to HistoricalMarketDepthDataRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

// CopyTo
func (m HistoricalMarketDepthDataRequestFixed) CopyTo(to HistoricalMarketDepthDataRequest) {
	to.SetRequestID(m.RequestID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetStartDateTime(m.StartDateTime())
	to.SetEndDateTime(m.EndDateTime())
	to.SetUseZLibCompression(m.UseZLibCompression())
	to.SetInteger_1(m.Integer_1())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequest) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *HistoricalMarketDepthDataRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequest) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *HistoricalMarketDepthDataRequest) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 900 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "EndDateTime":
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Bool())
		case "Integer_1":
			m.SetInteger_1(r.Uint8())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *HistoricalMarketDepthDataRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 900)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	w.Int64Field("EndDateTime", int64(m.EndDateTime()))
	w.BoolField("UseZLibCompression", m.UseZLibCompression())
	w.Uint8Field("Integer_1", m.Integer_1())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalMarketDepthDataRequestFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *HistoricalMarketDepthDataRequestFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 900 {
		return message.ErrWrongType
	}
	in := &r.Lexer
LOOP:
	for !in.IsDelim('}') {
		key, err := r.FieldName()
		if err != nil {
			return err
		}
		switch key {
		case "RequestID":
			m.SetRequestID(r.Int32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "EndDateTime":
			m.SetEndDateTime(DateTimeWithMicrosecondsInt(r.Int64()))
		case "UseZLibCompression":
			m.SetUseZLibCompression(r.Bool())
		case "Integer_1":
			m.SetInteger_1(r.Uint8())
		case "f", "F":
			return message.ErrJSONCompactDetected
		case "":
			break LOOP
		default:
			in.SkipRecursive()
		}
		if r.IsError() {
			return r.Error()
		}
		in.WantComma()
	}
	return nil
}
