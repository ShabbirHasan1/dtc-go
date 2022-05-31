// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const MarketDepthRequestSize = 28

const MarketDepthRequestFixedSize = 96

//     Size           uint16             = MarketDepthRequestSize  (28)
//     Type           uint16             = MARKET_DEPTH_REQUEST  (102)
//     BaseSize       uint16             = MarketDepthRequestSize  (28)
//     RequestAction  RequestActionEnum  = SUBSCRIBE  (1)
//     SymbolID       uint32             = 0
//     Symbol         string             = ""
//     Exchange       string             = ""
//     NumLevels      int32              = 10
var _MarketDepthRequestDefault = []byte{28, 0, 102, 0, 28, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 0, 0, 0}

//     Size           uint16             = MarketDepthRequestFixedSize  (96)
//     Type           uint16             = MARKET_DEPTH_REQUEST  (102)
//     RequestAction  RequestActionEnum  = SUBSCRIBE  (1)
//     SymbolID       uint32             = 0
//     Symbol         string[64]         = ""
//     Exchange       string[16]         = ""
//     NumLevels      int32              = 0
var _MarketDepthRequestFixedDefault = []byte{96, 0, 102, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type MarketDepthRequest struct {
	p message.VLS
}

type MarketDepthRequestFixed struct {
	p message.Fixed
}

func NewMarketDepthRequestFrom(b []byte) MarketDepthRequest {
	return MarketDepthRequest{p: message.NewVLS(b)}
}

func WrapMarketDepthRequest(b []byte) MarketDepthRequest {
	return MarketDepthRequest{p: message.WrapVLS(b)}
}

func NewMarketDepthRequest() *MarketDepthRequest {
	return &MarketDepthRequest{p: message.NewVLS(_MarketDepthRequestDefault)}
}

func ParseMarketDepthRequest(b []byte) (MarketDepthRequest, error) {
	if len(b) < 6 {
		return MarketDepthRequest{}, message.ErrShortBuffer
	}
	m := WrapMarketDepthRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDepthRequest{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return MarketDepthRequest{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 28 {
		newSize := len(b) + (28 - baseSize)
		if newSize > message.MaxSize {
			return MarketDepthRequest{}, message.ErrOverflow
		}
		clone := MarketDepthRequest{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _MarketDepthRequestDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(28 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(16)
			if offset > 0 {
				clone.p.SetUint16LE(16, offset+shift)
			}
			offset = clone.p.Uint16LE(20)
			if offset > 0 {
				clone.p.SetUint16LE(20, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewMarketDepthRequestFixedFrom(b []byte) MarketDepthRequestFixed {
	return MarketDepthRequestFixed{p: message.NewFixed(b)}
}

func WrapMarketDepthRequestFixed(b []byte) MarketDepthRequestFixed {
	return MarketDepthRequestFixed{p: message.WrapFixed(b)}
}

func NewMarketDepthRequestFixed() *MarketDepthRequestFixed {
	return &MarketDepthRequestFixed{p: message.NewFixed(_MarketDepthRequestFixedDefault)}
}

func ParseMarketDepthRequestFixed(b []byte) (MarketDepthRequestFixed, error) {
	if len(b) < 4 {
		return MarketDepthRequestFixed{}, message.ErrShortBuffer
	}
	m := WrapMarketDepthRequestFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return MarketDepthRequestFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return MarketDepthRequestFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 96 {
		clone := *NewMarketDepthRequestFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _MarketDepthRequestFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDepthRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDepthRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m MarketDepthRequest) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequest) RequestAction() RequestActionEnum {
	return RequestActionEnum(m.p.Int32LE(8))
}

// SymbolID This is the identifier which will be used in all of the market depth data
// response messages.
//
// This SymbolID can be the same as the one used in the MARKET_DATA_REQUEST
// message for the same Symbol and Exchange.
//
// This identifier is used so that the Symbol does not have to be passed
// back in response messages from the Server. If the Server receives a MARKET_DEPTH_REQUEST
// for a Symbol and Exchange to subscribe to market depth data for, that
// is currently subscribed to and this SymbolID is different, then the Server
// should reject it.
func (m MarketDepthRequest) SymbolID() uint32 {
	return m.p.Uint32LE(12)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequest) Symbol() string {
	return m.p.StringVLS(16)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequest) Exchange() string {
	return m.p.StringVLS(20)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequest) NumLevels() int32 {
	return m.p.Int32LE(24)
}

// Size The standard message size field. Automatically set by constructor.
func (m MarketDepthRequestFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m MarketDepthRequestFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m MarketDepthRequestFixed) RequestAction() RequestActionEnum {
	return RequestActionEnum(m.p.Int32LE(4))
}

// SymbolID This is the identifier which will be used in all of the market depth data
// response messages.
//
// This SymbolID can be the same as the one used in the MARKET_DATA_REQUEST
// message for the same Symbol and Exchange.
//
// This identifier is used so that the Symbol does not have to be passed
// back in response messages from the Server. If the Server receives a MARKET_DEPTH_REQUEST
// for a Symbol and Exchange to subscribe to market depth data for, that
// is currently subscribed to and this SymbolID is different, then the Server
// should reject it.
func (m MarketDepthRequestFixed) SymbolID() uint32 {
	return m.p.Uint32LE(8)
}

// Symbol The symbol for the market depth request. Not set when unsubscribing.
func (m MarketDepthRequestFixed) Symbol() string {
	return m.p.StringFixed(12, 64)
}

// Exchange The optional exchange for the symbol. Not set when unsubscribing.
func (m MarketDepthRequestFixed) Exchange() string {
	return m.p.StringFixed(76, 16)
}

// NumLevels Number of depth levels requested. Not set when unsubscribing.
func (m MarketDepthRequestFixed) NumLevels() int32 {
	return m.p.Int32LE(92)
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m *MarketDepthRequest) SetRequestAction(value RequestActionEnum) *MarketDepthRequest {
	m.p.SetInt32LE(8, int32(value))
	return m
}

// SetSymbolID This is the identifier which will be used in all of the market depth data
// response messages.
//
// This SymbolID can be the same as the one used in the MARKET_DATA_REQUEST
// message for the same Symbol and Exchange.
//
// This identifier is used so that the Symbol does not have to be passed
// back in response messages from the Server. If the Server receives a MARKET_DEPTH_REQUEST
// for a Symbol and Exchange to subscribe to market depth data for, that
// is currently subscribed to and this SymbolID is different, then the Server
// should reject it.
func (m *MarketDepthRequest) SetSymbolID(value uint32) *MarketDepthRequest {
	m.p.SetUint32LE(12, value)
	return m
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m *MarketDepthRequest) SetSymbol(value string) *MarketDepthRequest {
	m.p.SetStringVLS(16, value)
	return m
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m *MarketDepthRequest) SetExchange(value string) *MarketDepthRequest {
	m.p.SetStringVLS(20, value)
	return m
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m *MarketDepthRequest) SetNumLevels(value int32) *MarketDepthRequest {
	m.p.SetInt32LE(24, value)
	return m
}

// SetRequestAction This needs to be set to SUBSCRIBE to subscribe to market data for the
// Symbol from the Server. The server will respond with an initial MARKET_DEPTH_SNAPSHOT_LEVEL
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
// message and then provide MARKET_DEPTH_UPDATE_LEVEL updates as necessary.
//
// To unsubscribe, use UNSUBSCRIBE.
func (m *MarketDepthRequestFixed) SetRequestAction(value RequestActionEnum) *MarketDepthRequestFixed {
	m.p.SetInt32LE(4, int32(value))
	return m
}

// SetSymbolID This is the identifier which will be used in all of the market depth data
// response messages.
//
// This SymbolID can be the same as the one used in the MARKET_DATA_REQUEST
// message for the same Symbol and Exchange.
//
// This identifier is used so that the Symbol does not have to be passed
// back in response messages from the Server. If the Server receives a MARKET_DEPTH_REQUEST
// for a Symbol and Exchange to subscribe to market depth data for, that
// is currently subscribed to and this SymbolID is different, then the Server
// should reject it.
func (m *MarketDepthRequestFixed) SetSymbolID(value uint32) *MarketDepthRequestFixed {
	m.p.SetUint32LE(8, value)
	return m
}

// SetSymbol The symbol for the market depth request. Not set when unsubscribing.
func (m *MarketDepthRequestFixed) SetSymbol(value string) *MarketDepthRequestFixed {
	m.p.SetStringFixed(12, 64, value)
	return m
}

// SetExchange The optional exchange for the symbol. Not set when unsubscribing.
func (m *MarketDepthRequestFixed) SetExchange(value string) *MarketDepthRequestFixed {
	m.p.SetStringFixed(76, 16, value)
	return m
}

// SetNumLevels Number of depth levels requested. Not set when unsubscribing.
func (m *MarketDepthRequestFixed) SetNumLevels(value int32) *MarketDepthRequestFixed {
	m.p.SetInt32LE(92, value)
	return m
}

func (m MarketDepthRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketDepthRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m MarketDepthRequestFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m MarketDepthRequestFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m MarketDepthRequest) Copy(to MarketDepthRequest) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequest) CopyTo(to MarketDepthRequestFixed) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// Copy
func (m MarketDepthRequestFixed) Copy(to MarketDepthRequestFixed) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

// CopyTo
func (m MarketDepthRequestFixed) CopyTo(to MarketDepthRequest) {
	to.SetRequestAction(m.RequestAction())
	to.SetSymbolID(m.SymbolID())
	to.SetSymbol(m.Symbol())
	to.SetExchange(m.Exchange())
	to.SetNumLevels(m.NumLevels())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequest) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 102 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "NumLevels":
			m.SetNumLevels(r.Int32())
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

func (m *MarketDepthRequest) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m MarketDepthRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 102)
	w.Int32Field("RequestAction", int32(m.RequestAction()))
	w.Uint32Field("SymbolID", m.SymbolID())
	w.StringField("Symbol", m.Symbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("NumLevels", m.NumLevels())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *MarketDepthRequestFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 102 {
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
		case "RequestAction":
			m.SetRequestAction(RequestActionEnum(r.Int32()))
		case "SymbolID":
			m.SetSymbolID(r.Uint32())
		case "Symbol":
			m.SetSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "NumLevels":
			m.SetNumLevels(r.Int32())
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

func (m *MarketDepthRequestFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
