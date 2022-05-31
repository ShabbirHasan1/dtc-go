// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const TradeAccountDataUpdateOperationCompleteSize = 22

//     Size                               uint16  = TradeAccountDataUpdateOperationCompleteSize  (22)
//     Type                               uint16  = TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE  (10131)
//     BaseSize                           uint16  = TradeAccountDataUpdateOperationCompleteSize  (22)
//     RequestID                          uint32  = 0
//     IsError                            bool    = false
//     ErrorText                          string  = ""
//     IsDeleteOperation                  bool    = false
//     IsSymbolLimitsUpdateOperation      bool    = false
//     IsSymbolCommissionUpdateOperation  bool    = false
//     TradeAccount                       string  = ""
var _TradeAccountDataUpdateOperationCompleteDefault = []byte{22, 0, 147, 39, 22, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type TradeAccountDataUpdateOperationComplete struct {
	p message.VLS
}

func NewTradeAccountDataUpdateOperationCompleteFrom(b []byte) TradeAccountDataUpdateOperationComplete {
	return TradeAccountDataUpdateOperationComplete{p: message.NewVLS(b)}
}

func WrapTradeAccountDataUpdateOperationComplete(b []byte) TradeAccountDataUpdateOperationComplete {
	return TradeAccountDataUpdateOperationComplete{p: message.WrapVLS(b)}
}

func NewTradeAccountDataUpdateOperationComplete() *TradeAccountDataUpdateOperationComplete {
	return &TradeAccountDataUpdateOperationComplete{p: message.NewVLS(_TradeAccountDataUpdateOperationCompleteDefault)}
}

func ParseTradeAccountDataUpdateOperationComplete(b []byte) (TradeAccountDataUpdateOperationComplete, error) {
	if len(b) < 6 {
		return TradeAccountDataUpdateOperationComplete{}, message.ErrShortBuffer
	}
	m := WrapTradeAccountDataUpdateOperationComplete(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return TradeAccountDataUpdateOperationComplete{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return TradeAccountDataUpdateOperationComplete{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 22 {
		newSize := len(b) + (22 - baseSize)
		if newSize > message.MaxSize {
			return TradeAccountDataUpdateOperationComplete{}, message.ErrOverflow
		}
		clone := TradeAccountDataUpdateOperationComplete{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _TradeAccountDataUpdateOperationCompleteDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(22 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(11)
			if offset > 0 {
				clone.p.SetUint16LE(11, offset+shift)
			}
			offset = clone.p.Uint16LE(18)
			if offset > 0 {
				clone.p.SetUint16LE(18, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

// Size
func (m TradeAccountDataUpdateOperationComplete) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m TradeAccountDataUpdateOperationComplete) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m TradeAccountDataUpdateOperationComplete) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m TradeAccountDataUpdateOperationComplete) RequestID() uint32 {
	return m.p.Uint32LE(6)
}

// IsError
func (m TradeAccountDataUpdateOperationComplete) IsError() bool {
	return m.p.Bool(10)
}

// ErrorText
func (m TradeAccountDataUpdateOperationComplete) ErrorText() string {
	return m.p.StringVLS(11)
}

// IsDeleteOperation
func (m TradeAccountDataUpdateOperationComplete) IsDeleteOperation() bool {
	return m.p.Bool(15)
}

// IsSymbolLimitsUpdateOperation
func (m TradeAccountDataUpdateOperationComplete) IsSymbolLimitsUpdateOperation() bool {
	return m.p.Bool(16)
}

// IsSymbolCommissionUpdateOperation
func (m TradeAccountDataUpdateOperationComplete) IsSymbolCommissionUpdateOperation() bool {
	return m.p.Bool(17)
}

// TradeAccount
func (m TradeAccountDataUpdateOperationComplete) TradeAccount() string {
	return m.p.StringVLS(18)
}

// SetRequestID
func (m *TradeAccountDataUpdateOperationComplete) SetRequestID(value uint32) *TradeAccountDataUpdateOperationComplete {
	m.p.SetUint32LE(6, value)
	return m
}

// SetIsError
func (m *TradeAccountDataUpdateOperationComplete) SetIsError(value bool) *TradeAccountDataUpdateOperationComplete {
	m.p.SetBool(10, value)
	return m
}

// SetErrorText
func (m *TradeAccountDataUpdateOperationComplete) SetErrorText(value string) *TradeAccountDataUpdateOperationComplete {
	m.p.SetStringVLS(11, value)
	return m
}

// SetIsDeleteOperation
func (m *TradeAccountDataUpdateOperationComplete) SetIsDeleteOperation(value bool) *TradeAccountDataUpdateOperationComplete {
	m.p.SetBool(15, value)
	return m
}

// SetIsSymbolLimitsUpdateOperation
func (m *TradeAccountDataUpdateOperationComplete) SetIsSymbolLimitsUpdateOperation(value bool) *TradeAccountDataUpdateOperationComplete {
	m.p.SetBool(16, value)
	return m
}

// SetIsSymbolCommissionUpdateOperation
func (m *TradeAccountDataUpdateOperationComplete) SetIsSymbolCommissionUpdateOperation(value bool) *TradeAccountDataUpdateOperationComplete {
	m.p.SetBool(17, value)
	return m
}

// SetTradeAccount
func (m *TradeAccountDataUpdateOperationComplete) SetTradeAccount(value string) *TradeAccountDataUpdateOperationComplete {
	m.p.SetStringVLS(18, value)
	return m
}

func (m TradeAccountDataUpdateOperationComplete) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m TradeAccountDataUpdateOperationComplete) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m TradeAccountDataUpdateOperationComplete) Copy(to TradeAccountDataUpdateOperationComplete) {
	to.SetRequestID(m.RequestID())
	to.SetIsError(m.IsError())
	to.SetErrorText(m.ErrorText())
	to.SetIsDeleteOperation(m.IsDeleteOperation())
	to.SetIsSymbolLimitsUpdateOperation(m.IsSymbolLimitsUpdateOperation())
	to.SetIsSymbolCommissionUpdateOperation(m.IsSymbolCommissionUpdateOperation())
	to.SetTradeAccount(m.TradeAccount())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m TradeAccountDataUpdateOperationComplete) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 10131)
	w.Uint32Field("RequestID", m.RequestID())
	w.BoolField("IsError", m.IsError())
	w.StringField("ErrorText", m.ErrorText())
	w.BoolField("IsDeleteOperation", m.IsDeleteOperation())
	w.BoolField("IsSymbolLimitsUpdateOperation", m.IsSymbolLimitsUpdateOperation())
	w.BoolField("IsSymbolCommissionUpdateOperation", m.IsSymbolCommissionUpdateOperation())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *TradeAccountDataUpdateOperationComplete) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 10131 {
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
		case "IsError":
			m.SetIsError(r.Bool())
		case "ErrorText":
			m.SetErrorText(r.String())
		case "IsDeleteOperation":
			m.SetIsDeleteOperation(r.Bool())
		case "IsSymbolLimitsUpdateOperation":
			m.SetIsSymbolLimitsUpdateOperation(r.Bool())
		case "IsSymbolCommissionUpdateOperation":
			m.SetIsSymbolCommissionUpdateOperation(r.Bool())
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

func (m *TradeAccountDataUpdateOperationComplete) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
