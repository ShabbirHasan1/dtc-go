// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const TradeAccountResponseSize = 24

const TradeAccountResponseFixedSize = 48

//     Size                 uint16  = TradeAccountResponseSize  (24)
//     Type                 uint16  = TRADE_ACCOUNT_RESPONSE  (401)
//     BaseSize             uint16  = TradeAccountResponseSize  (24)
//     TotalNumberMessages  int32   = 0
//     MessageNumber        int32   = 0
//     TradeAccount         string  = ""
//     RequestID            int32   = 0
var _TradeAccountResponseDefault = []byte{24, 0, 145, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size                 uint16      = TradeAccountResponseFixedSize  (48)
//     Type                 uint16      = TRADE_ACCOUNT_RESPONSE  (401)
//     TotalNumberMessages  int32       = 0
//     MessageNumber        int32       = 0
//     TradeAccount         string[32]  = ""
//     RequestID            int32       = 0
var _TradeAccountResponseFixedDefault = []byte{48, 0, 145, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// TradeAccountResponse This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponse struct {
	p message.VLS
}

// TradeAccountResponseFixed This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
type TradeAccountResponseFixed struct {
	p message.Fixed
}

func NewTradeAccountResponseFrom(b []byte) TradeAccountResponse {
	return TradeAccountResponse{p: message.NewVLS(b)}
}

func WrapTradeAccountResponse(b []byte) TradeAccountResponse {
	return TradeAccountResponse{p: message.WrapVLS(b)}
}

func NewTradeAccountResponse() *TradeAccountResponse {
	return &TradeAccountResponse{p: message.NewVLS(_TradeAccountResponseDefault)}
}

func ParseTradeAccountResponse(b []byte) (TradeAccountResponse, error) {
	if len(b) < 6 {
		return TradeAccountResponse{}, message.ErrShortBuffer
	}
	m := WrapTradeAccountResponse(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return TradeAccountResponse{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return TradeAccountResponse{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 24 {
		newSize := len(b) + (24 - baseSize)
		if newSize > message.MaxSize {
			return TradeAccountResponse{}, message.ErrOverflow
		}
		clone := TradeAccountResponse{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _TradeAccountResponseDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(24 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(16)
			if offset > 0 {
				clone.p.SetUint16LE(16, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewTradeAccountResponseFixedFrom(b []byte) TradeAccountResponseFixed {
	return TradeAccountResponseFixed{p: message.NewFixed(b)}
}

func WrapTradeAccountResponseFixed(b []byte) TradeAccountResponseFixed {
	return TradeAccountResponseFixed{p: message.WrapFixed(b)}
}

func NewTradeAccountResponseFixed() *TradeAccountResponseFixed {
	return &TradeAccountResponseFixed{p: message.NewFixed(_TradeAccountResponseFixedDefault)}
}

func ParseTradeAccountResponseFixed(b []byte) (TradeAccountResponseFixed, error) {
	if len(b) < 4 {
		return TradeAccountResponseFixed{}, message.ErrShortBuffer
	}
	m := WrapTradeAccountResponseFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return TradeAccountResponseFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return TradeAccountResponseFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 48 {
		clone := *NewTradeAccountResponseFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _TradeAccountResponseFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m TradeAccountResponse) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m TradeAccountResponse) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m TradeAccountResponse) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponse) TotalNumberMessages() int32 {
	return m.p.Int32LE(8)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponse) MessageNumber() int32 {
	return m.p.Int32LE(12)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponse) TradeAccount() string {
	return m.p.StringVLS(16)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponse) RequestID() int32 {
	return m.p.Int32LE(20)
}

// Size The standard message size field. Automatically set by constructor.
func (m TradeAccountResponseFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m TradeAccountResponseFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// TotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixed) TotalNumberMessages() int32 {
	return m.p.Int32LE(4)
}

// MessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m TradeAccountResponseFixed) MessageNumber() int32 {
	return m.p.Int32LE(8)
}

// TradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m TradeAccountResponseFixed) TradeAccount() string {
	return m.p.StringFixed(12, 32)
}

// RequestID The RequestID sent in the corresponding request by the Client.
//
func (m TradeAccountResponseFixed) RequestID() int32 {
	return m.p.Int32LE(44)
}

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m *TradeAccountResponse) SetTotalNumberMessages(value int32) *TradeAccountResponse {
	m.p.SetInt32LE(8, value)
	return m
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m *TradeAccountResponse) SetMessageNumber(value int32) *TradeAccountResponse {
	m.p.SetInt32LE(12, value)
	return m
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m *TradeAccountResponse) SetTradeAccount(value string) *TradeAccountResponse {
	m.p.SetStringVLS(16, value)
	return m
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m *TradeAccountResponse) SetRequestID(value int32) *TradeAccountResponse {
	m.p.SetInt32LE(20, value)
	return m
}

// SetTotalNumberMessages This indicates the total number of Account list messages when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m *TradeAccountResponseFixed) SetTotalNumberMessages(value int32) *TradeAccountResponseFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetMessageNumber This indicates the 1-based index of the Account list message when a batch
// of messages is being sent. If there is only one message being sent, this
// will be 1.
func (m *TradeAccountResponseFixed) SetMessageNumber(value int32) *TradeAccountResponseFixed {
	m.p.SetInt32LE(8, value)
	return m
}

// SetTradeAccount The trade account identifier.
//
// In the case when there are no Trade Accounts available for the logged
// in Username, the Server will send a single TradeAccountResponse message
// to the Client and leave this field empty.
func (m *TradeAccountResponseFixed) SetTradeAccount(value string) *TradeAccountResponseFixed {
	m.p.SetStringFixed(12, 32, value)
	return m
}

// SetRequestID The RequestID sent in the corresponding request by the Client.
//
func (m *TradeAccountResponseFixed) SetRequestID(value int32) *TradeAccountResponseFixed {
	m.p.SetInt32LE(44, value)
	return m
}

func (m *TradeAccountResponse) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *TradeAccountResponse) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *TradeAccountResponse) Clone() *TradeAccountResponse {
	return &TradeAccountResponse{message.WrapVLSPointer(m.p.Clone(uintptr(m.Size())), int(m.Size()))}
}

func (m *TradeAccountResponseFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *TradeAccountResponseFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *TradeAccountResponseFixed) Clone() *TradeAccountResponseFixed {
	return &TradeAccountResponseFixed{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m TradeAccountResponse) Copy(to TradeAccountResponse) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponse) CopyTo(to TradeAccountResponseFixed) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// Copy
func (m TradeAccountResponseFixed) Copy(to TradeAccountResponseFixed) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

// CopyTo
func (m TradeAccountResponseFixed) CopyTo(to TradeAccountResponse) {
	to.SetTotalNumberMessages(m.TotalNumberMessages())
	to.SetMessageNumber(m.MessageNumber())
	to.SetTradeAccount(m.TradeAccount())
	to.SetRequestID(m.RequestID())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponse) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *TradeAccountResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponse) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *TradeAccountResponse) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 401 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "RequestID":
			m.SetRequestID(r.Int32())
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

func (m *TradeAccountResponseFixed) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *TradeAccountResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 401)
	w.Int32Field("TotalNumberMessages", m.TotalNumberMessages())
	w.Int32Field("MessageNumber", m.MessageNumber())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountResponseFixed) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *TradeAccountResponseFixed) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 401 {
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
		case "TotalNumberMessages":
			m.SetTotalNumberMessages(r.Int32())
		case "MessageNumber":
			m.SetMessageNumber(r.Int32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "RequestID":
			m.SetRequestID(r.Int32())
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
