// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const SymbolSearchRequestSize = 28

const SymbolSearchRequestFixedSize = 96

//     Size          uint16            = SymbolSearchRequestSize  (28)
//     Type          uint16            = SYMBOL_SEARCH_REQUEST  (508)
//     BaseSize      uint16            = SymbolSearchRequestSize  (28)
//     RequestID     int32             = 0
//     SearchText    string            = ""
//     Exchange      string            = ""
//     SecurityType  SecurityTypeEnum  = SECURITY_TYPE_UNSET  (0)
//     SearchType    SearchTypeEnum    = SEARCH_TYPE_UNSET  (0)
var _SymbolSearchRequestDefault = []byte{28, 0, 252, 1, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size          uint16            = SymbolSearchRequestFixedSize  (96)
//     Type          uint16            = SYMBOL_SEARCH_REQUEST  (508)
//     RequestID     int32             = 0
//     SearchText    string[64]        = ""
//     Exchange      string[16]        = ""
//     SecurityType  SecurityTypeEnum  = SECURITY_TYPE_UNSET  (0)
//     SearchType    SearchTypeEnum    = SEARCH_TYPE_UNSET  (0)
var _SymbolSearchRequestFixedDefault = []byte{96, 0, 252, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// SymbolSearchRequest The SymbolSearchRequest message is sent by the Client to the Server to
// return Security Definitions matching the specified SecurityType and Exchange
// and where the Symbol or Description contains the specified SearchText.
//
// The SearchText can search either the Symbol or the Description field in
// the SecurityDefinitionResponse message.
//
// In either case there does not need to be an exact match. The SearchText
// only needs to be contained within the Symbol or the Description depending
// upon which field is being searched.
//
// The Server returns SecurityDefinitionResponse messages for all Symbols
// which match.
//
// If there are no matches, the Server needs to send a SecurityDefinitionResponse
// message to the Client with with all fields at their default values except
// for the RequestID and IsFinalMessage fields set. This will be a clear
// indication to the Client that the request returned no matches.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolSearchRequest struct {
	p message.VLS
}

// SymbolSearchRequestFixed The SymbolSearchRequest message is sent by the Client to the Server to
// return Security Definitions matching the specified SecurityType and Exchange
// and where the Symbol or Description contains the specified SearchText.
//
// The SearchText can search either the Symbol or the Description field in
// the SecurityDefinitionResponse message.
//
// In either case there does not need to be an exact match. The SearchText
// only needs to be contained within the Symbol or the Description depending
// upon which field is being searched.
//
// The Server returns SecurityDefinitionResponse messages for all Symbols
// which match.
//
// If there are no matches, the Server needs to send a SecurityDefinitionResponse
// message to the Client with with all fields at their default values except
// for the RequestID and IsFinalMessage fields set. This will be a clear
// indication to the Client that the request returned no matches.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
type SymbolSearchRequestFixed struct {
	p message.Fixed
}

func NewSymbolSearchRequestFrom(b []byte) SymbolSearchRequest {
	return SymbolSearchRequest{p: message.NewVLS(b)}
}

func WrapSymbolSearchRequest(b []byte) SymbolSearchRequest {
	return SymbolSearchRequest{p: message.WrapVLS(b)}
}

func NewSymbolSearchRequest() *SymbolSearchRequest {
	return &SymbolSearchRequest{p: message.NewVLS(_SymbolSearchRequestDefault)}
}

func ParseSymbolSearchRequest(b []byte) (SymbolSearchRequest, error) {
	if len(b) < 6 {
		return SymbolSearchRequest{}, message.ErrShortBuffer
	}
	m := WrapSymbolSearchRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SymbolSearchRequest{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return SymbolSearchRequest{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 28 {
		newSize := len(b) + (28 - baseSize)
		if newSize > message.MaxSize {
			return SymbolSearchRequest{}, message.ErrOverflow
		}
		clone := SymbolSearchRequest{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _SymbolSearchRequestDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(28 - baseSize)
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

func NewSymbolSearchRequestFixedFrom(b []byte) SymbolSearchRequestFixed {
	return SymbolSearchRequestFixed{p: message.NewFixed(b)}
}

func WrapSymbolSearchRequestFixed(b []byte) SymbolSearchRequestFixed {
	return SymbolSearchRequestFixed{p: message.WrapFixed(b)}
}

func NewSymbolSearchRequestFixed() *SymbolSearchRequestFixed {
	return &SymbolSearchRequestFixed{p: message.NewFixed(_SymbolSearchRequestFixedDefault)}
}

func ParseSymbolSearchRequestFixed(b []byte) (SymbolSearchRequestFixed, error) {
	if len(b) < 4 {
		return SymbolSearchRequestFixed{}, message.ErrShortBuffer
	}
	m := WrapSymbolSearchRequestFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return SymbolSearchRequestFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return SymbolSearchRequestFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 96 {
		clone := *NewSymbolSearchRequestFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _SymbolSearchRequestFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m SymbolSearchRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SymbolSearchRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m SymbolSearchRequest) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequest) RequestID() int32 {
	return m.p.Int32LE(8)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequest) SearchText() string {
	return m.p.StringVLS(12)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequest) Exchange() string {
	return m.p.StringVLS(16)
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(m.p.Int32LE(20))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequest) SearchType() SearchTypeEnum {
	return SearchTypeEnum(m.p.Int32LE(24))
}

// Size The standard message size field. Automatically set by constructor.
func (m SymbolSearchRequestFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m SymbolSearchRequestFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixed) SearchText() string {
	return m.p.StringFixed(8, 64)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixed) Exchange() string {
	return m.p.StringFixed(72, 16)
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(m.p.Int32LE(88))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixed) SearchType() SearchTypeEnum {
	return SearchTypeEnum(m.p.Int32LE(92))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m *SymbolSearchRequest) SetRequestID(value int32) *SymbolSearchRequest {
	m.p.SetInt32LE(8, value)
	return m
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m *SymbolSearchRequest) SetSearchText(value string) *SymbolSearchRequest {
	m.p.SetStringVLS(12, value)
	return m
}

// SetExchange The Exchange of the Symbol to search for.
func (m *SymbolSearchRequest) SetExchange(value string) *SymbolSearchRequest {
	m.p.SetStringVLS(16, value)
	return m
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m *SymbolSearchRequest) SetSecurityType(value SecurityTypeEnum) *SymbolSearchRequest {
	m.p.SetInt32LE(20, int32(value))
	return m
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m *SymbolSearchRequest) SetSearchType(value SearchTypeEnum) *SymbolSearchRequest {
	m.p.SetInt32LE(24, int32(value))
	return m
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m *SymbolSearchRequestFixed) SetRequestID(value int32) *SymbolSearchRequestFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m *SymbolSearchRequestFixed) SetSearchText(value string) *SymbolSearchRequestFixed {
	m.p.SetStringFixed(8, 64, value)
	return m
}

// SetExchange The Exchange of the Symbol to search for.
func (m *SymbolSearchRequestFixed) SetExchange(value string) *SymbolSearchRequestFixed {
	m.p.SetStringFixed(72, 16, value)
	return m
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m *SymbolSearchRequestFixed) SetSecurityType(value SecurityTypeEnum) *SymbolSearchRequestFixed {
	m.p.SetInt32LE(88, int32(value))
	return m
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m *SymbolSearchRequestFixed) SetSearchType(value SearchTypeEnum) *SymbolSearchRequestFixed {
	m.p.SetInt32LE(92, int32(value))
	return m
}

func (m SymbolSearchRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m SymbolSearchRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m SymbolSearchRequestFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m SymbolSearchRequestFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m SymbolSearchRequest) Copy(to SymbolSearchRequest) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequest) CopyTo(to SymbolSearchRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// Copy
func (m SymbolSearchRequestFixed) Copy(to SymbolSearchRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestFixed) CopyTo(to SymbolSearchRequest) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolSearchRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 508)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("SearchText", m.SearchText())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("SearchType", int32(m.SearchType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolSearchRequest) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 508 {
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
		case "SearchText":
			m.SetSearchText(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "SearchType":
			m.SetSearchType(SearchTypeEnum(r.Int32()))
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

func (m *SymbolSearchRequest) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m SymbolSearchRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 508)
	w.Int32Field("RequestID", m.RequestID())
	w.StringField("SearchText", m.SearchText())
	w.StringField("Exchange", m.Exchange())
	w.Int32Field("SecurityType", int32(m.SecurityType()))
	w.Int32Field("SearchType", int32(m.SearchType()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *SymbolSearchRequestFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 508 {
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
		case "SearchText":
			m.SetSearchText(r.String())
		case "Exchange":
			m.SetExchange(r.String())
		case "SecurityType":
			m.SetSecurityType(SecurityTypeEnum(r.Int32()))
		case "SearchType":
			m.SetSearchType(SearchTypeEnum(r.Int32()))
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

func (m *SymbolSearchRequestFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
