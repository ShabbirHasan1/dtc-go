// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const SubmitFlattenPositionOrderSize = 28

const SubmitFlattenPositionOrderFixedSize = 198

//     Size              uint16  = SubmitFlattenPositionOrderSize  (28)
//     Type              uint16  = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     BaseSize          uint16  = SubmitFlattenPositionOrderSize  (28)
//     Symbol            string  = ""
//     Exchange          string  = ""
//     TradeAccount      string  = ""
//     ClientOrderID     string  = ""
//     FreeFormText      string  = ""
//     IsAutomatedOrder  bool    = false
var _SubmitFlattenPositionOrderDefault = []byte{28, 0, 209, 0, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size              uint16      = SubmitFlattenPositionOrderFixedSize  (198)
//     Type              uint16      = SUBMIT_FLATTEN_POSITION_ORDER  (209)
//     Symbol            string[64]  = ""
//     Exchange          string[16]  = ""
//     TradeAccount      string[32]  = ""
//     ClientOrderID     string[32]  = ""
//     FreeFormText      string[48]  = ""
//     IsAutomatedOrder  bool        = false
var _SubmitFlattenPositionOrderFixedDefault = []byte{198, 0, 209, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type SubmitFlattenPositionOrder struct {
	p message.VLS
}

type SubmitFlattenPositionOrderFixed struct {
	p message.Fixed
}

func NewSubmitFlattenPositionOrderFrom(b []byte) SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{p: message.NewVLS(b)}
}

func WrapSubmitFlattenPositionOrder(b []byte) SubmitFlattenPositionOrder {
	return SubmitFlattenPositionOrder{p: message.WrapVLS(b)}
}

func NewSubmitFlattenPositionOrder() *SubmitFlattenPositionOrder {
	return &SubmitFlattenPositionOrder{p: message.NewVLS(_SubmitFlattenPositionOrderDefault)}
}

func ParseSubmitFlattenPositionOrder(b []byte) (SubmitFlattenPositionOrder, error) {
	if len(b) < 6 {
		return SubmitFlattenPositionOrder{}, message.ErrShortBuffer
	}
	m := WrapSubmitFlattenPositionOrder(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SubmitFlattenPositionOrder{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return SubmitFlattenPositionOrder{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 28 {
		newSize := len(b) + (28 - baseSize)
		if newSize > message.MaxSize {
			return SubmitFlattenPositionOrder{}, message.ErrOverflow
		}
		clone := SubmitFlattenPositionOrder{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _SubmitFlattenPositionOrderDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(28 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(6)
			if offset > 0 {
				clone.p.SetUint16LE(6, offset+shift)
			}
			offset = clone.p.Uint16LE(10)
			if offset > 0 {
				clone.p.SetUint16LE(10, offset+shift)
			}
			offset = clone.p.Uint16LE(14)
			if offset > 0 {
				clone.p.SetUint16LE(14, offset+shift)
			}
			offset = clone.p.Uint16LE(18)
			if offset > 0 {
				clone.p.SetUint16LE(18, offset+shift)
			}
			offset = clone.p.Uint16LE(22)
			if offset > 0 {
				clone.p.SetUint16LE(22, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewSubmitFlattenPositionOrderFixedFrom(b []byte) SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{p: message.NewFixed(b)}
}

func WrapSubmitFlattenPositionOrderFixed(b []byte) SubmitFlattenPositionOrderFixed {
	return SubmitFlattenPositionOrderFixed{p: message.WrapFixed(b)}
}

func NewSubmitFlattenPositionOrderFixed() *SubmitFlattenPositionOrderFixed {
	return &SubmitFlattenPositionOrderFixed{p: message.NewFixed(_SubmitFlattenPositionOrderFixedDefault)}
}

func ParseSubmitFlattenPositionOrderFixed(b []byte) (SubmitFlattenPositionOrderFixed, error) {
	if len(b) < 4 {
		return SubmitFlattenPositionOrderFixed{}, message.ErrShortBuffer
	}
	m := WrapSubmitFlattenPositionOrderFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SubmitFlattenPositionOrderFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return SubmitFlattenPositionOrderFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 198 {
		clone := *NewSubmitFlattenPositionOrderFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _SubmitFlattenPositionOrderFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m SubmitFlattenPositionOrder) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SubmitFlattenPositionOrder) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m SubmitFlattenPositionOrder) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrder) Symbol() string {
	return m.p.StringVLS(6)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrder) Exchange() string {
	return m.p.StringVLS(10)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrder) TradeAccount() string {
	return m.p.StringVLS(14)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrder) ClientOrderID() string {
	return m.p.StringVLS(18)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrder) FreeFormText() string {
	return m.p.StringVLS(22)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrder) IsAutomatedOrder() bool {
	return m.p.Bool(26)
}

// Size The standard message size field. Automatically set by constructor.
func (m SubmitFlattenPositionOrderFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SubmitFlattenPositionOrderFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// Symbol The symbol of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixed) Symbol() string {
	return m.p.StringFixed(4, 64)
}

// Exchange The optional exchange for the Symbol.
func (m SubmitFlattenPositionOrderFixed) Exchange() string {
	return m.p.StringFixed(68, 16)
}

// TradeAccount The trade account as a text string of the Trade Position to flatten.
func (m SubmitFlattenPositionOrderFixed) TradeAccount() string {
	return m.p.StringFixed(84, 32)
}

// ClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m SubmitFlattenPositionOrderFixed) ClientOrderID() string {
	return m.p.StringFixed(116, 32)
}

// FreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m SubmitFlattenPositionOrderFixed) FreeFormText() string {
	return m.p.StringFixed(148, 48)
}

// IsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m SubmitFlattenPositionOrderFixed) IsAutomatedOrder() bool {
	return m.p.Bool(196)
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrder) SetSymbol(value string) *SubmitFlattenPositionOrder {
	m.p.SetStringVLS(6, value)
	return m
}

// SetExchange The optional exchange for the Symbol.
func (m *SubmitFlattenPositionOrder) SetExchange(value string) *SubmitFlattenPositionOrder {
	m.p.SetStringVLS(10, value)
	return m
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrder) SetTradeAccount(value string) *SubmitFlattenPositionOrder {
	m.p.SetStringVLS(14, value)
	return m
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m *SubmitFlattenPositionOrder) SetClientOrderID(value string) *SubmitFlattenPositionOrder {
	m.p.SetStringVLS(18, value)
	return m
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m *SubmitFlattenPositionOrder) SetFreeFormText(value string) *SubmitFlattenPositionOrder {
	m.p.SetStringVLS(22, value)
	return m
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m *SubmitFlattenPositionOrder) SetIsAutomatedOrder(value bool) *SubmitFlattenPositionOrder {
	m.p.SetBool(26, value)
	return m
}

// SetSymbol The symbol of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderFixed) SetSymbol(value string) *SubmitFlattenPositionOrderFixed {
	m.p.SetStringFixed(4, 64, value)
	return m
}

// SetExchange The optional exchange for the Symbol.
func (m *SubmitFlattenPositionOrderFixed) SetExchange(value string) *SubmitFlattenPositionOrderFixed {
	m.p.SetStringFixed(68, 16, value)
	return m
}

// SetTradeAccount The trade account as a text string of the Trade Position to flatten.
func (m *SubmitFlattenPositionOrderFixed) SetTradeAccount(value string) *SubmitFlattenPositionOrderFixed {
	m.p.SetStringFixed(84, 32, value)
	return m
}

// SetClientOrderID The Client supplied order identifier for the order which will be created
// to flatten the Trade Position.
//
// The Server will remember this for the life of the order.
func (m *SubmitFlattenPositionOrderFixed) SetClientOrderID(value string) *SubmitFlattenPositionOrderFixed {
	m.p.SetStringFixed(116, 32, value)
	return m
}

// SetFreeFormText Optional: This is an optional text string which can be set by the Client
// to associate text with the order which will be created to flatten the
// Trade Position.
//
// The Server is not under any obligation to use this text and it may place
// a limitation on the length of this text.
func (m *SubmitFlattenPositionOrderFixed) SetFreeFormText(value string) *SubmitFlattenPositionOrderFixed {
	m.p.SetStringFixed(148, 48, value)
	return m
}

// SetIsAutomatedOrder Set to 1 for an order submitted by an automated trading system.
func (m *SubmitFlattenPositionOrderFixed) SetIsAutomatedOrder(value bool) *SubmitFlattenPositionOrderFixed {
	m.p.SetBool(196, value)
	return m
}

func (m SubmitFlattenPositionOrder) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m SubmitFlattenPositionOrder) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m SubmitFlattenPositionOrderFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m SubmitFlattenPositionOrderFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m SubmitFlattenPositionOrder) Copy(to SubmitFlattenPositionOrder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrder) CopyTo(to SubmitFlattenPositionOrderFixed) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// Copy
func (m SubmitFlattenPositionOrderFixed) Copy(to SubmitFlattenPositionOrderFixed) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

// CopyTo
func (m SubmitFlattenPositionOrderFixed) CopyTo(to SubmitFlattenPositionOrder) {
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetTradeAccount(m.TradeAccount())
	to.SetClientOrderID(m.ClientOrderID())
	to.SetFreeFormText(m.FreeFormText())
	to.SetIsAutomatedOrder(m.IsAutomatedOrder())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitFlattenPositionOrder) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 209)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitFlattenPositionOrder) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 209 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
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

func (m *SubmitFlattenPositionOrder) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SubmitFlattenPositionOrderFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 209)
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.StringField("TradeAccount", m.TradeAccount())
	w.StringField("ClientOrderID", m.ClientOrderID())
	w.StringField("FreeFormText", m.FreeFormText())
	w.BoolField("IsAutomatedOrder", m.IsAutomatedOrder())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SubmitFlattenPositionOrderFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 209 {
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
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
		case "ClientOrderID":
			m.SetClientOrderID(r.String())
		case "FreeFormText":
			m.SetFreeFormText(r.String())
		case "IsAutomatedOrder":
			m.SetIsAutomatedOrder(r.Bool())
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

func (m *SubmitFlattenPositionOrderFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}