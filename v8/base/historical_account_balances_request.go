// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const HistoricalAccountBalancesRequestSize = 24

const HistoricalAccountBalancesRequestFixedSize = 48

//     Size           uint16    = HistoricalAccountBalancesRequestSize  (24)
//     Type           uint16    = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     BaseSize       uint16    = HistoricalAccountBalancesRequestSize  (24)
//     RequestID      int32     = 0
//     TradeAccount   string    = ""
//     StartDateTime  DateTime  = 0
var _HistoricalAccountBalancesRequestDefault = []byte{24, 0, 91, 2, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size           uint16      = HistoricalAccountBalancesRequestFixedSize  (48)
//     Type           uint16      = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
//     RequestID      int32       = 0
//     TradeAccount   string[32]  = ""
//     StartDateTime  DateTime    = 0
var _HistoricalAccountBalancesRequestFixedDefault = []byte{48, 0, 91, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalAccountBalancesRequest This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequest struct {
	p message.VLS
}

// HistoricalAccountBalancesRequestFixed This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
type HistoricalAccountBalancesRequestFixed struct {
	p message.Fixed
}

func NewHistoricalAccountBalancesRequestFrom(b []byte) HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{p: message.NewVLS(b)}
}

func WrapHistoricalAccountBalancesRequest(b []byte) HistoricalAccountBalancesRequest {
	return HistoricalAccountBalancesRequest{p: message.WrapVLS(b)}
}

func NewHistoricalAccountBalancesRequest() *HistoricalAccountBalancesRequest {
	return &HistoricalAccountBalancesRequest{p: message.NewVLS(_HistoricalAccountBalancesRequestDefault)}
}

func ParseHistoricalAccountBalancesRequest(b []byte) (HistoricalAccountBalancesRequest, error) {
	if len(b) < 6 {
		return HistoricalAccountBalancesRequest{}, message.ErrShortBuffer
	}
	m := WrapHistoricalAccountBalancesRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalAccountBalancesRequest{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return HistoricalAccountBalancesRequest{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 24 {
		newSize := len(b) + (24 - baseSize)
		if newSize > message.MaxSize {
			return HistoricalAccountBalancesRequest{}, message.ErrOverflow
		}
		clone := HistoricalAccountBalancesRequest{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _HistoricalAccountBalancesRequestDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(24 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(12)
			if offset > 0 {
				clone.p.SetUint16LE(12, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewHistoricalAccountBalancesRequestFixedFrom(b []byte) HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{p: message.NewFixed(b)}
}

func WrapHistoricalAccountBalancesRequestFixed(b []byte) HistoricalAccountBalancesRequestFixed {
	return HistoricalAccountBalancesRequestFixed{p: message.WrapFixed(b)}
}

func NewHistoricalAccountBalancesRequestFixed() *HistoricalAccountBalancesRequestFixed {
	return &HistoricalAccountBalancesRequestFixed{p: message.NewFixed(_HistoricalAccountBalancesRequestFixedDefault)}
}

func ParseHistoricalAccountBalancesRequestFixed(b []byte) (HistoricalAccountBalancesRequestFixed, error) {
	if len(b) < 4 {
		return HistoricalAccountBalancesRequestFixed{}, message.ErrShortBuffer
	}
	m := WrapHistoricalAccountBalancesRequestFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalAccountBalancesRequestFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return HistoricalAccountBalancesRequestFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 48 {
		clone := *NewHistoricalAccountBalancesRequestFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _HistoricalAccountBalancesRequestFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m HistoricalAccountBalancesRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m HistoricalAccountBalancesRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m HistoricalAccountBalancesRequest) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequest) RequestID() int32 {
	return m.p.Int32LE(8)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequest) TradeAccount() string {
	return m.p.StringVLS(12)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequest) StartDateTime() DateTime {
	return DateTime(m.p.Int64LE(16))
}

// Size The standard message size field. Automatically set by constructor.
func (m HistoricalAccountBalancesRequestFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m HistoricalAccountBalancesRequestFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID A unique request identifier for this request.
func (m HistoricalAccountBalancesRequestFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// TradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m HistoricalAccountBalancesRequestFixed) TradeAccount() string {
	return m.p.StringFixed(8, 32)
}

// StartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m HistoricalAccountBalancesRequestFixed) StartDateTime() DateTime {
	return DateTime(m.p.Int64LE(40))
}

// SetRequestID A unique request identifier for this request.
func (m *HistoricalAccountBalancesRequest) SetRequestID(value int32) *HistoricalAccountBalancesRequest {
	m.p.SetInt32LE(8, value)
	return m
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m *HistoricalAccountBalancesRequest) SetTradeAccount(value string) *HistoricalAccountBalancesRequest {
	m.p.SetStringVLS(12, value)
	return m
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m *HistoricalAccountBalancesRequest) SetStartDateTime(value DateTime) *HistoricalAccountBalancesRequest {
	m.p.SetInt64LE(16, int64(value))
	return m
}

// SetRequestID A unique request identifier for this request.
func (m *HistoricalAccountBalancesRequestFixed) SetRequestID(value int32) *HistoricalAccountBalancesRequestFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetTradeAccount This is a required field. Set to the particular Trade Account for which
// to request historical Account Balance data.
func (m *HistoricalAccountBalancesRequestFixed) SetTradeAccount(value string) *HistoricalAccountBalancesRequestFixed {
	m.p.SetStringFixed(8, 32, value)
	return m
}

// SetStartDateTime Set this to the Date-Time that the server is to send historical cash balance
// updates starting with.
func (m *HistoricalAccountBalancesRequestFixed) SetStartDateTime(value DateTime) *HistoricalAccountBalancesRequestFixed {
	m.p.SetInt64LE(40, int64(value))
	return m
}

func (m *HistoricalAccountBalancesRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *HistoricalAccountBalancesRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *HistoricalAccountBalancesRequest) Clone() *HistoricalAccountBalancesRequest {
	return &HistoricalAccountBalancesRequest{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *HistoricalAccountBalancesRequestFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *HistoricalAccountBalancesRequestFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *HistoricalAccountBalancesRequestFixed) Clone() *HistoricalAccountBalancesRequestFixed {
	return &HistoricalAccountBalancesRequestFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m HistoricalAccountBalancesRequest) Copy(to HistoricalAccountBalancesRequest) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequest) CopyTo(to HistoricalAccountBalancesRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// Copy
func (m HistoricalAccountBalancesRequestFixed) Copy(to HistoricalAccountBalancesRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

// CopyTo
func (m HistoricalAccountBalancesRequestFixed) CopyTo(to HistoricalAccountBalancesRequest) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetStartDateTime(m.StartDateTime())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalancesRequest) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *HistoricalAccountBalancesRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 603)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalancesRequest) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *HistoricalAccountBalancesRequest) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 603 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
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

func (m *HistoricalAccountBalancesRequestFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *HistoricalAccountBalancesRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 603)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int64Field("StartDateTime", int64(m.StartDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalancesRequestFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *HistoricalAccountBalancesRequestFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 603 {
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
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "StartDateTime":
			m.SetStartDateTime(DateTime(r.Int64()))
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
