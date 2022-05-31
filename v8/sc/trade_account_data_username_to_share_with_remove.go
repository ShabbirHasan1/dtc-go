// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const TradeAccountDataUsernameToShareWithRemoveSize = 19

//     Size          uint16  = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     Type          uint16  = TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE  (10129)
//     BaseSize      uint16  = TradeAccountDataUsernameToShareWithRemoveSize  (19)
//     RequestID     uint32  = 0
//     TradeAccount  string  = ""
//     Username      string  = ""
//     IsReadOnly    bool    = false
var _TradeAccountDataUsernameToShareWithRemoveDefault = []byte{19, 0, 145, 39, 19, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataUsernameToShareWithRemove struct {
	p message.VLS
}

func NewTradeAccountDataUsernameToShareWithRemoveFrom(b []byte) TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{p: message.NewVLS(b)}
}

func WrapTradeAccountDataUsernameToShareWithRemove(b []byte) TradeAccountDataUsernameToShareWithRemove {
	return TradeAccountDataUsernameToShareWithRemove{p: message.WrapVLS(b)}
}

func NewTradeAccountDataUsernameToShareWithRemove() *TradeAccountDataUsernameToShareWithRemove {
	return &TradeAccountDataUsernameToShareWithRemove{p: message.NewVLS(_TradeAccountDataUsernameToShareWithRemoveDefault)}
}

func ParseTradeAccountDataUsernameToShareWithRemove(b []byte) (TradeAccountDataUsernameToShareWithRemove, error) {
	if len(b) < 6 {
		return TradeAccountDataUsernameToShareWithRemove{}, message.ErrShortBuffer
	}
	m := WrapTradeAccountDataUsernameToShareWithRemove(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return TradeAccountDataUsernameToShareWithRemove{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return TradeAccountDataUsernameToShareWithRemove{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 19 {
		newSize := len(b) + (19 - baseSize)
		if newSize > message.MaxSize {
			return TradeAccountDataUsernameToShareWithRemove{}, message.ErrOverflow
		}
		clone := TradeAccountDataUsernameToShareWithRemove{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _TradeAccountDataUsernameToShareWithRemoveDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(19 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(10)
			if offset > 0 {
				clone.p.SetUint16LE(10, offset+shift)
			}
			offset = clone.p.Uint16LE(14)
			if offset > 0 {
				clone.p.SetUint16LE(14, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

// Size
func (m TradeAccountDataUsernameToShareWithRemove) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m TradeAccountDataUsernameToShareWithRemove) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m TradeAccountDataUsernameToShareWithRemove) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m TradeAccountDataUsernameToShareWithRemove) RequestID() uint32 {
	return m.p.Uint32LE(6)
}

// TradeAccount
func (m TradeAccountDataUsernameToShareWithRemove) TradeAccount() string {
	return m.p.StringVLS(10)
}

// Username
func (m TradeAccountDataUsernameToShareWithRemove) Username() string {
	return m.p.StringVLS(14)
}

// IsReadOnly
func (m TradeAccountDataUsernameToShareWithRemove) IsReadOnly() bool {
	return m.p.Bool(18)
}

// SetRequestID
func (m *TradeAccountDataUsernameToShareWithRemove) SetRequestID(value uint32) *TradeAccountDataUsernameToShareWithRemove {
	m.p.SetUint32LE(6, value)
	return m
}

// SetTradeAccount
func (m *TradeAccountDataUsernameToShareWithRemove) SetTradeAccount(value string) *TradeAccountDataUsernameToShareWithRemove {
	m.p.SetStringVLS(10, value)
	return m
}

// SetUsername
func (m *TradeAccountDataUsernameToShareWithRemove) SetUsername(value string) *TradeAccountDataUsernameToShareWithRemove {
	m.p.SetStringVLS(14, value)
	return m
}

// SetIsReadOnly
func (m *TradeAccountDataUsernameToShareWithRemove) SetIsReadOnly(value bool) *TradeAccountDataUsernameToShareWithRemove {
	m.p.SetBool(18, value)
	return m
}

func (m TradeAccountDataUsernameToShareWithRemove) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m TradeAccountDataUsernameToShareWithRemove) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m TradeAccountDataUsernameToShareWithRemove) Copy(to TradeAccountDataUsernameToShareWithRemove) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetUsername(m.Username())
	to.SetIsReadOnly(m.IsReadOnly())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUsernameToShareWithRemove) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10129)
	w.Uint32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("Username", m.Username())
	w.BoolField("IsReadOnly", m.IsReadOnly())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUsernameToShareWithRemove) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10129 {
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
			m.SetRequestID(r.Uint32())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "Username":
			m.SetUsername(r.String())
		case "IsReadOnly":
			m.SetIsReadOnly(r.Bool())
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

func (m *TradeAccountDataUsernameToShareWithRemove) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
