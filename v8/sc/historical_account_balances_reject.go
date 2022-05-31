// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const HistoricalAccountBalancesRejectSize = 16

const HistoricalAccountBalancesRejectFixedSize = 104

//     Size        uint16  = HistoricalAccountBalancesRejectSize  (16)
//     Type        uint16  = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     BaseSize    uint16  = HistoricalAccountBalancesRejectSize  (16)
//     RequestID   int32   = 0
//     RejectText  string  = ""
var _HistoricalAccountBalancesRejectDefault = []byte{16, 0, 92, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size        uint16      = HistoricalAccountBalancesRejectFixedSize  (104)
//     Type        uint16      = HISTORICAL_ACCOUNT_BALANCES_REJECT  (604)
//     RequestID   int32       = 0
//     RejectText  string[96]  = ""
var _HistoricalAccountBalancesRejectFixedDefault = []byte{104, 0, 92, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// HistoricalAccountBalancesReject This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesReject struct {
	p message.VLS
}

// HistoricalAccountBalancesRejectFixed This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
type HistoricalAccountBalancesRejectFixed struct {
	p message.Fixed
}

func NewHistoricalAccountBalancesRejectFrom(b []byte) HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{p: message.NewVLS(b)}
}

func WrapHistoricalAccountBalancesReject(b []byte) HistoricalAccountBalancesReject {
	return HistoricalAccountBalancesReject{p: message.WrapVLS(b)}
}

func NewHistoricalAccountBalancesReject() *HistoricalAccountBalancesReject {
	return &HistoricalAccountBalancesReject{p: message.NewVLS(_HistoricalAccountBalancesRejectDefault)}
}

func ParseHistoricalAccountBalancesReject(b []byte) (HistoricalAccountBalancesReject, error) {
	if len(b) < 6 {
		return HistoricalAccountBalancesReject{}, message.ErrShortBuffer
	}
	m := WrapHistoricalAccountBalancesReject(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalAccountBalancesReject{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return HistoricalAccountBalancesReject{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 16 {
		newSize := len(b) + (16 - baseSize)
		if newSize > message.MaxSize {
			return HistoricalAccountBalancesReject{}, message.ErrOverflow
		}
		clone := HistoricalAccountBalancesReject{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _HistoricalAccountBalancesRejectDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(16 - baseSize)
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

func NewHistoricalAccountBalancesRejectFixedFrom(b []byte) HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{p: message.NewFixed(b)}
}

func WrapHistoricalAccountBalancesRejectFixed(b []byte) HistoricalAccountBalancesRejectFixed {
	return HistoricalAccountBalancesRejectFixed{p: message.WrapFixed(b)}
}

func NewHistoricalAccountBalancesRejectFixed() *HistoricalAccountBalancesRejectFixed {
	return &HistoricalAccountBalancesRejectFixed{p: message.NewFixed(_HistoricalAccountBalancesRejectFixedDefault)}
}

func ParseHistoricalAccountBalancesRejectFixed(b []byte) (HistoricalAccountBalancesRejectFixed, error) {
	if len(b) < 4 {
		return HistoricalAccountBalancesRejectFixed{}, message.ErrShortBuffer
	}
	m := WrapHistoricalAccountBalancesRejectFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return HistoricalAccountBalancesRejectFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return HistoricalAccountBalancesRejectFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 104 {
		clone := *NewHistoricalAccountBalancesRejectFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _HistoricalAccountBalancesRejectFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m HistoricalAccountBalancesReject) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m HistoricalAccountBalancesReject) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m HistoricalAccountBalancesReject) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesReject) RequestID() int32 {
	return m.p.Int32LE(8)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesReject) RejectText() string {
	return m.p.StringVLS(12)
}

// Size The standard message size field. Automatically set by constructor.
func (m HistoricalAccountBalancesRejectFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m HistoricalAccountBalancesRejectFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m HistoricalAccountBalancesRejectFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// RejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m HistoricalAccountBalancesRejectFixed) RejectText() string {
	return m.p.StringFixed(8, 96)
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m *HistoricalAccountBalancesReject) SetRequestID(value int32) *HistoricalAccountBalancesReject {
	m.p.SetInt32LE(8, value)
	return m
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m *HistoricalAccountBalancesReject) SetRejectText(value string) *HistoricalAccountBalancesReject {
	m.p.SetStringVLS(12, value)
	return m
}

// SetRequestID The unique request identifier sent in the corresponding request message.
// The unique request identifier sent in the corresponding request message.
func (m *HistoricalAccountBalancesRejectFixed) SetRequestID(value int32) *HistoricalAccountBalancesRejectFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetRejectText The text reason the HistoricalAccountBalancesRequest message was rejected.
// The text reason the HistoricalAccountBalancesRequest message was rejected.
func (m *HistoricalAccountBalancesRejectFixed) SetRejectText(value string) *HistoricalAccountBalancesRejectFixed {
	m.p.SetStringFixed(8, 96, value)
	return m
}

func (m HistoricalAccountBalancesReject) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m HistoricalAccountBalancesReject) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m HistoricalAccountBalancesRejectFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m HistoricalAccountBalancesRejectFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m HistoricalAccountBalancesReject) Copy(to HistoricalAccountBalancesReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesReject) CopyTo(to HistoricalAccountBalancesRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// Copy
func (m HistoricalAccountBalancesRejectFixed) Copy(to HistoricalAccountBalancesRejectFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

// CopyTo
func (m HistoricalAccountBalancesRejectFixed) CopyTo(to HistoricalAccountBalancesReject) {
	to.SetRequestID(m.RequestID())
	to.SetRejectText(m.RejectText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalancesReject) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 604)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalancesReject) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 604 {
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
		case "RejectText":
			m.SetRejectText(r.String())
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
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalancesReject) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m HistoricalAccountBalancesRejectFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 604)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("RejectText", m.RejectText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalancesRejectFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 604 {
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
		case "RejectText":
			m.SetRejectText(r.String())
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
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *HistoricalAccountBalancesRejectFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
