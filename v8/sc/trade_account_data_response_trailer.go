// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const TradeAccountDataResponseTrailerSize = 17

//     Size                   uint16  = TradeAccountDataResponseTrailerSize  (17)
//     Type                   uint16  = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER  (10130)
//     BaseSize               uint16  = TradeAccountDataResponseTrailerSize  (17)
//     RequestID              uint32  = 0
//     IsSnapshot             bool    = false
//     IsFirstMessageInBatch  bool    = false
//     IsLastMessageInBatch   bool    = false
//     TradeAccount           string  = ""
var _TradeAccountDataResponseTrailerDefault = []byte{17, 0, 146, 39, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataResponseTrailer struct {
	p message.VLS
}

func NewTradeAccountDataResponseTrailerFrom(b []byte) TradeAccountDataResponseTrailer {
	return TradeAccountDataResponseTrailer{p: message.NewVLS(b)}
}

func WrapTradeAccountDataResponseTrailer(b []byte) TradeAccountDataResponseTrailer {
	return TradeAccountDataResponseTrailer{p: message.WrapVLS(b)}
}

func NewTradeAccountDataResponseTrailer() *TradeAccountDataResponseTrailer {
	return &TradeAccountDataResponseTrailer{p: message.NewVLS(_TradeAccountDataResponseTrailerDefault)}
}

func ParseTradeAccountDataResponseTrailer(b []byte) (TradeAccountDataResponseTrailer, error) {
	if len(b) < 6 {
		return TradeAccountDataResponseTrailer{}, message.ErrShortBuffer
	}
	m := WrapTradeAccountDataResponseTrailer(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return TradeAccountDataResponseTrailer{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return TradeAccountDataResponseTrailer{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 17 {
		newSize := len(b) + (17 - baseSize)
		if newSize > message.MaxSize {
			return TradeAccountDataResponseTrailer{}, message.ErrOverflow
		}
		clone := TradeAccountDataResponseTrailer{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _TradeAccountDataResponseTrailerDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(17 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(13)
			if offset > 0 {
				clone.p.SetUint16LE(13, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

// Size
func (m TradeAccountDataResponseTrailer) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m TradeAccountDataResponseTrailer) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m TradeAccountDataResponseTrailer) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m TradeAccountDataResponseTrailer) RequestID() uint32 {
	return m.p.Uint32LE(6)
}

// IsSnapshot
func (m TradeAccountDataResponseTrailer) IsSnapshot() bool {
	return m.p.Bool(10)
}

// IsFirstMessageInBatch
func (m TradeAccountDataResponseTrailer) IsFirstMessageInBatch() bool {
	return m.p.Bool(11)
}

// IsLastMessageInBatch
func (m TradeAccountDataResponseTrailer) IsLastMessageInBatch() bool {
	return m.p.Bool(12)
}

// TradeAccount
func (m TradeAccountDataResponseTrailer) TradeAccount() string {
	return m.p.StringVLS(13)
}

// SetRequestID
func (m *TradeAccountDataResponseTrailer) SetRequestID(value uint32) *TradeAccountDataResponseTrailer {
	m.p.SetUint32LE(6, value)
	return m
}

// SetIsSnapshot
func (m *TradeAccountDataResponseTrailer) SetIsSnapshot(value bool) *TradeAccountDataResponseTrailer {
	m.p.SetBool(10, value)
	return m
}

// SetIsFirstMessageInBatch
func (m *TradeAccountDataResponseTrailer) SetIsFirstMessageInBatch(value bool) *TradeAccountDataResponseTrailer {
	m.p.SetBool(11, value)
	return m
}

// SetIsLastMessageInBatch
func (m *TradeAccountDataResponseTrailer) SetIsLastMessageInBatch(value bool) *TradeAccountDataResponseTrailer {
	m.p.SetBool(12, value)
	return m
}

// SetTradeAccount
func (m *TradeAccountDataResponseTrailer) SetTradeAccount(value string) *TradeAccountDataResponseTrailer {
	m.p.SetStringVLS(13, value)
	return m
}

func (m TradeAccountDataResponseTrailer) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m TradeAccountDataResponseTrailer) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m TradeAccountDataResponseTrailer) Copy(to TradeAccountDataResponseTrailer) {
	to.SetRequestID(m.RequestID())
	to.SetIsSnapshot(m.IsSnapshot())
	to.SetIsFirstMessageInBatch(m.IsFirstMessageInBatch())
	to.SetIsLastMessageInBatch(m.IsLastMessageInBatch())
	to.SetTradeAccount(m.TradeAccount())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataResponseTrailer) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10130)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("m_IsSnapshot", m.IsSnapshot())
	w.BoolField("m_IsFirstMessageInBatch", m.IsFirstMessageInBatch())
	w.BoolField("m_IsLastMessageInBatch", m.IsLastMessageInBatch())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataResponseTrailer) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10130 {
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
		case "m_IsSnapshot":
			m.SetIsSnapshot(r.Bool())
		case "m_IsFirstMessageInBatch":
			m.SetIsFirstMessageInBatch(r.Bool())
		case "m_IsLastMessageInBatch":
			m.SetIsLastMessageInBatch(r.Bool())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *TradeAccountDataResponseTrailer) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
