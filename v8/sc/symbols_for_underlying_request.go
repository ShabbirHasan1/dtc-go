// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const SymbolsForUnderlyingRequestSize = 24

const SymbolsForUnderlyingRequestFixedSize = 60

//     Size              uint16            = SymbolsForUnderlyingRequestSize  (24)
//     Type              uint16            = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     BaseSize          uint16            = SymbolsForUnderlyingRequestSize  (24)
//     RequestID         int32             = 0
//     UnderlyingSymbol  string            = ""
//     Exchange          string            = ""
//     SecurityType      SecurityTypeEnum  = SECURITY_TYPE_UNSET  (0)
var _SymbolsForUnderlyingRequestDefault = []byte{24, 0, 248, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size              uint16            = SymbolsForUnderlyingRequestFixedSize  (60)
//     Type              uint16            = SYMBOLS_FOR_UNDERLYING_REQUEST  (504)
//     RequestID         int32             = 0
//     UnderlyingSymbol  string[32]        = ""
//     Exchange          string[16]        = ""
//     SecurityType      SecurityTypeEnum  = SECURITY_TYPE_UNSET  (0)
var _SymbolsForUnderlyingRequestFixedDefault = []byte{60, 0, 248, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SymbolsForUnderlyingRequest This is a message from the Client to the Server for requesting all of
// the symbols for a particular underlying symbol.
//
// For example, all of the futures contracts for a particular underlying
// futures symbol or all of the option symbols for a specific futures or
// stock symbol.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForUnderlyingRequest struct {
	p message.VLS
}

// SymbolsForUnderlyingRequestFixed This is a message from the Client to the Server for requesting all of
// the symbols for a particular underlying symbol.
//
// For example, all of the futures contracts for a particular underlying
// futures symbol or all of the option symbols for a specific futures or
// stock symbol.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolsForUnderlyingRequestFixed struct {
	p message.Fixed
}

func NewSymbolsForUnderlyingRequestFrom(b []byte) SymbolsForUnderlyingRequest {
	return SymbolsForUnderlyingRequest{p: message.NewVLS(b)}
}

func WrapSymbolsForUnderlyingRequest(b []byte) SymbolsForUnderlyingRequest {
	return SymbolsForUnderlyingRequest{p: message.WrapVLS(b)}
}

func NewSymbolsForUnderlyingRequest() *SymbolsForUnderlyingRequest {
	return &SymbolsForUnderlyingRequest{p: message.NewVLS(_SymbolsForUnderlyingRequestDefault)}
}

func ParseSymbolsForUnderlyingRequest(b []byte) (SymbolsForUnderlyingRequest, error) {
	if len(b) < 6 {
		return SymbolsForUnderlyingRequest{}, message.ErrShortBuffer
	}
	m := WrapSymbolsForUnderlyingRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SymbolsForUnderlyingRequest{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return SymbolsForUnderlyingRequest{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 24 {
		newSize := len(b) + (24 - baseSize)
		if newSize > message.MaxSize {
			return SymbolsForUnderlyingRequest{}, message.ErrOverflow
		}
		clone := SymbolsForUnderlyingRequest{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _SymbolsForUnderlyingRequestDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(24 - baseSize)
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

func NewSymbolsForUnderlyingRequestFixedFrom(b []byte) SymbolsForUnderlyingRequestFixed {
	return SymbolsForUnderlyingRequestFixed{p: message.NewFixed(b)}
}

func WrapSymbolsForUnderlyingRequestFixed(b []byte) SymbolsForUnderlyingRequestFixed {
	return SymbolsForUnderlyingRequestFixed{p: message.WrapFixed(b)}
}

func NewSymbolsForUnderlyingRequestFixed() *SymbolsForUnderlyingRequestFixed {
	return &SymbolsForUnderlyingRequestFixed{p: message.NewFixed(_SymbolsForUnderlyingRequestFixedDefault)}
}

func ParseSymbolsForUnderlyingRequestFixed(b []byte) (SymbolsForUnderlyingRequestFixed, error) {
	if len(b) < 4 {
		return SymbolsForUnderlyingRequestFixed{}, message.ErrShortBuffer
	}
	m := WrapSymbolsForUnderlyingRequestFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SymbolsForUnderlyingRequestFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return SymbolsForUnderlyingRequestFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 60 {
		clone := *NewSymbolsForUnderlyingRequestFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _SymbolsForUnderlyingRequestFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m SymbolsForUnderlyingRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SymbolsForUnderlyingRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m SymbolsForUnderlyingRequest) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequest) RequestID() int32 {
	return m.p.Int32LE(8)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequest) UnderlyingSymbol() string {
	return m.p.StringVLS(12)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequest) Exchange() string {
	return m.p.StringVLS(16)
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(m.p.Int32LE(20))
}

// Size The standard message size field. Automatically set by constructor.
func (m SymbolsForUnderlyingRequestFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SymbolsForUnderlyingRequestFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolsForUnderlyingRequestFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// UnderlyingSymbol The underlying symbol.
func (m SymbolsForUnderlyingRequestFixed) UnderlyingSymbol() string {
	return m.p.StringFixed(8, 32)
}

// Exchange The exchange of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixed) Exchange() string {
	return m.p.StringFixed(40, 16)
}

// SecurityType The security type of the symbols to search for.
func (m SymbolsForUnderlyingRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(m.p.Int32LE(56))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m *SymbolsForUnderlyingRequest) SetRequestID(value int32) *SymbolsForUnderlyingRequest {
	m.p.SetInt32LE(8, value)
	return m
}

// SetUnderlyingSymbol The underlying symbol.
func (m *SymbolsForUnderlyingRequest) SetUnderlyingSymbol(value string) *SymbolsForUnderlyingRequest {
	m.p.SetStringVLS(12, value)
	return m
}

// SetExchange The exchange of the symbols to search for.
func (m *SymbolsForUnderlyingRequest) SetExchange(value string) *SymbolsForUnderlyingRequest {
	m.p.SetStringVLS(16, value)
	return m
}

// SetSecurityType The security type of the symbols to search for.
func (m *SymbolsForUnderlyingRequest) SetSecurityType(value SecurityTypeEnum) *SymbolsForUnderlyingRequest {
	m.p.SetInt32LE(20, int32(value))
	return m
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m *SymbolsForUnderlyingRequestFixed) SetRequestID(value int32) *SymbolsForUnderlyingRequestFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetUnderlyingSymbol The underlying symbol.
func (m *SymbolsForUnderlyingRequestFixed) SetUnderlyingSymbol(value string) *SymbolsForUnderlyingRequestFixed {
	m.p.SetStringFixed(8, 32, value)
	return m
}

// SetExchange The exchange of the symbols to search for.
func (m *SymbolsForUnderlyingRequestFixed) SetExchange(value string) *SymbolsForUnderlyingRequestFixed {
	m.p.SetStringFixed(40, 16, value)
	return m
}

// SetSecurityType The security type of the symbols to search for.
func (m *SymbolsForUnderlyingRequestFixed) SetSecurityType(value SecurityTypeEnum) *SymbolsForUnderlyingRequestFixed {
	m.p.SetInt32LE(56, int32(value))
	return m
}

func (m SymbolsForUnderlyingRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m SymbolsForUnderlyingRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m SymbolsForUnderlyingRequestFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m SymbolsForUnderlyingRequestFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m SymbolsForUnderlyingRequest) Copy(to SymbolsForUnderlyingRequest) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequest) CopyTo(to SymbolsForUnderlyingRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// Copy
func (m SymbolsForUnderlyingRequestFixed) Copy(to SymbolsForUnderlyingRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

// CopyTo
func (m SymbolsForUnderlyingRequestFixed) CopyTo(to SymbolsForUnderlyingRequest) {
	to.SetRequestID(m.RequestID())
	to.SetUnderlyingSymbol(m.UnderlyingSymbol())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForUnderlyingRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 504)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForUnderlyingRequest) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 504 {
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
		case "UnderlyingSymbol":
			m.SetUnderlyingSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
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

func (m *SymbolsForUnderlyingRequest) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolsForUnderlyingRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 504)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("UnderlyingSymbol", m.UnderlyingSymbol())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolsForUnderlyingRequestFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 504 {
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
		case "UnderlyingSymbol":
			m.SetUnderlyingSymbol(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
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

func (m *SymbolsForUnderlyingRequestFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
