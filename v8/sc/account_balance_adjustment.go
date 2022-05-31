// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const AccountBalanceAdjustmentSize = 40

const AccountBalanceAdjustmentFixedSize = 160

//     Size          uint16   = AccountBalanceAdjustmentSize  (40)
//     Type          uint16   = ACCOUNT_BALANCE_ADJUSTMENT  (607)
//     BaseSize      uint16   = AccountBalanceAdjustmentSize  (40)
//     RequestID     int32    = 0
//     TradeAccount  string   = ""
//     CreditAmount  float64  = 0.000000
//     DebitAmount   float64  = 0.000000
//     Currency      string   = ""
//     Reason        string   = ""
var _AccountBalanceAdjustmentDefault = []byte{40, 0, 95, 2, 40, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size          uint16      = AccountBalanceAdjustmentFixedSize  (160)
//     Type          uint16      = ACCOUNT_BALANCE_ADJUSTMENT  (607)
//     RequestID     int32       = 0
//     TradeAccount  string[32]  = ""
//     CreditAmount  float64     = 0.000000
//     DebitAmount   float64     = 0.000000
//     Currency      string[8]   = ""
//     Reason        string[96]  = ""
var _AccountBalanceAdjustmentFixedDefault = []byte{160, 0, 95, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type AccountBalanceAdjustment struct {
	p message.VLS
}

type AccountBalanceAdjustmentFixed struct {
	p message.Fixed
}

func NewAccountBalanceAdjustmentFrom(b []byte) AccountBalanceAdjustment {
	return AccountBalanceAdjustment{p: message.NewVLS(b)}
}

func WrapAccountBalanceAdjustment(b []byte) AccountBalanceAdjustment {
	return AccountBalanceAdjustment{p: message.WrapVLS(b)}
}

func NewAccountBalanceAdjustment() *AccountBalanceAdjustment {
	return &AccountBalanceAdjustment{p: message.NewVLS(_AccountBalanceAdjustmentDefault)}
}

func ParseAccountBalanceAdjustment(b []byte) (AccountBalanceAdjustment, error) {
	if len(b) < 6 {
		return AccountBalanceAdjustment{}, message.ErrShortBuffer
	}
	m := WrapAccountBalanceAdjustment(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return AccountBalanceAdjustment{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return AccountBalanceAdjustment{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 40 {
		newSize := len(b) + (40 - baseSize)
		if newSize > message.MaxSize {
			return AccountBalanceAdjustment{}, message.ErrOverflow
		}
		clone := AccountBalanceAdjustment{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _AccountBalanceAdjustmentDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(40 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(12)
			if offset > 0 {
				clone.p.SetUint16LE(12, offset+shift)
			}
			offset = clone.p.Uint16LE(32)
			if offset > 0 {
				clone.p.SetUint16LE(32, offset+shift)
			}
			offset = clone.p.Uint16LE(36)
			if offset > 0 {
				clone.p.SetUint16LE(36, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewAccountBalanceAdjustmentFixedFrom(b []byte) AccountBalanceAdjustmentFixed {
	return AccountBalanceAdjustmentFixed{p: message.NewFixed(b)}
}

func WrapAccountBalanceAdjustmentFixed(b []byte) AccountBalanceAdjustmentFixed {
	return AccountBalanceAdjustmentFixed{p: message.WrapFixed(b)}
}

func NewAccountBalanceAdjustmentFixed() *AccountBalanceAdjustmentFixed {
	return &AccountBalanceAdjustmentFixed{p: message.NewFixed(_AccountBalanceAdjustmentFixedDefault)}
}

func ParseAccountBalanceAdjustmentFixed(b []byte) (AccountBalanceAdjustmentFixed, error) {
	if len(b) < 4 {
		return AccountBalanceAdjustmentFixed{}, message.ErrShortBuffer
	}
	m := WrapAccountBalanceAdjustmentFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return AccountBalanceAdjustmentFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return AccountBalanceAdjustmentFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 160 {
		clone := *NewAccountBalanceAdjustmentFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _AccountBalanceAdjustmentFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m AccountBalanceAdjustment) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m AccountBalanceAdjustment) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m AccountBalanceAdjustment) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID
func (m AccountBalanceAdjustment) RequestID() int32 {
	return m.p.Int32LE(8)
}

// TradeAccount
func (m AccountBalanceAdjustment) TradeAccount() string {
	return m.p.StringVLS(12)
}

// CreditAmount
func (m AccountBalanceAdjustment) CreditAmount() float64 {
	return m.p.Float64LE(16)
}

// DebitAmount
func (m AccountBalanceAdjustment) DebitAmount() float64 {
	return m.p.Float64LE(24)
}

// Currency
func (m AccountBalanceAdjustment) Currency() string {
	return m.p.StringVLS(32)
}

// Reason
func (m AccountBalanceAdjustment) Reason() string {
	return m.p.StringVLS(36)
}

// Size
func (m AccountBalanceAdjustmentFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m AccountBalanceAdjustmentFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID
func (m AccountBalanceAdjustmentFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// TradeAccount
func (m AccountBalanceAdjustmentFixed) TradeAccount() string {
	return m.p.StringFixed(8, 32)
}

// CreditAmount
func (m AccountBalanceAdjustmentFixed) CreditAmount() float64 {
	return m.p.Float64LE(40)
}

// DebitAmount
func (m AccountBalanceAdjustmentFixed) DebitAmount() float64 {
	return m.p.Float64LE(48)
}

// Currency
func (m AccountBalanceAdjustmentFixed) Currency() string {
	return m.p.StringFixed(56, 8)
}

// Reason
func (m AccountBalanceAdjustmentFixed) Reason() string {
	return m.p.StringFixed(64, 96)
}

// SetRequestID
func (m *AccountBalanceAdjustment) SetRequestID(value int32) *AccountBalanceAdjustment {
	m.p.SetInt32LE(8, value)
	return m
}

// SetTradeAccount
func (m *AccountBalanceAdjustment) SetTradeAccount(value string) *AccountBalanceAdjustment {
	m.p.SetStringVLS(12, value)
	return m
}

// SetCreditAmount
func (m *AccountBalanceAdjustment) SetCreditAmount(value float64) *AccountBalanceAdjustment {
	m.p.SetFloat64LE(16, value)
	return m
}

// SetDebitAmount
func (m *AccountBalanceAdjustment) SetDebitAmount(value float64) *AccountBalanceAdjustment {
	m.p.SetFloat64LE(24, value)
	return m
}

// SetCurrency
func (m *AccountBalanceAdjustment) SetCurrency(value string) *AccountBalanceAdjustment {
	m.p.SetStringVLS(32, value)
	return m
}

// SetReason
func (m *AccountBalanceAdjustment) SetReason(value string) *AccountBalanceAdjustment {
	m.p.SetStringVLS(36, value)
	return m
}

// SetRequestID
func (m *AccountBalanceAdjustmentFixed) SetRequestID(value int32) *AccountBalanceAdjustmentFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetTradeAccount
func (m *AccountBalanceAdjustmentFixed) SetTradeAccount(value string) *AccountBalanceAdjustmentFixed {
	m.p.SetStringFixed(8, 32, value)
	return m
}

// SetCreditAmount
func (m *AccountBalanceAdjustmentFixed) SetCreditAmount(value float64) *AccountBalanceAdjustmentFixed {
	m.p.SetFloat64LE(40, value)
	return m
}

// SetDebitAmount
func (m *AccountBalanceAdjustmentFixed) SetDebitAmount(value float64) *AccountBalanceAdjustmentFixed {
	m.p.SetFloat64LE(48, value)
	return m
}

// SetCurrency
func (m *AccountBalanceAdjustmentFixed) SetCurrency(value string) *AccountBalanceAdjustmentFixed {
	m.p.SetStringFixed(56, 8, value)
	return m
}

// SetReason
func (m *AccountBalanceAdjustmentFixed) SetReason(value string) *AccountBalanceAdjustmentFixed {
	m.p.SetStringFixed(64, 96, value)
	return m
}

func (m AccountBalanceAdjustment) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m AccountBalanceAdjustment) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m AccountBalanceAdjustmentFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m AccountBalanceAdjustmentFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m AccountBalanceAdjustment) Copy(to AccountBalanceAdjustment) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustment) CopyTo(to AccountBalanceAdjustmentFixed) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// Copy
func (m AccountBalanceAdjustmentFixed) Copy(to AccountBalanceAdjustmentFixed) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

// CopyTo
func (m AccountBalanceAdjustmentFixed) CopyTo(to AccountBalanceAdjustment) {
	to.SetRequestID(m.RequestID())
	to.SetTradeAccount(m.TradeAccount())
	to.SetCreditAmount(m.CreditAmount())
	to.SetDebitAmount(m.DebitAmount())
	to.SetCurrency(m.Currency())
	to.SetReason(m.Reason())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustment) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustment) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 607 {
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
		case "CreditAmount":
			m.SetCreditAmount(r.Float64())
		case "DebitAmount":
			m.SetDebitAmount(r.Float64())
		case "Currency":
			m.SetCurrency(r.String())
		case "Reason":
			m.SetReason(r.String())
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

func (m *AccountBalanceAdjustment) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m AccountBalanceAdjustmentFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 607)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("TradeAccount", m.TradeAccount())
	w.Float64Field("CreditAmount", m.CreditAmount())
	w.Float64Field("DebitAmount", m.DebitAmount())
	w.StringField("Currency", m.Currency())
	w.StringField("Reason", m.Reason())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *AccountBalanceAdjustmentFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 607 {
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
		case "CreditAmount":
			m.SetCreditAmount(r.Float64())
		case "DebitAmount":
			m.SetDebitAmount(r.Float64())
		case "Currency":
			m.SetCurrency(r.String())
		case "Reason":
			m.SetReason(r.String())
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

func (m *AccountBalanceAdjustmentFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}