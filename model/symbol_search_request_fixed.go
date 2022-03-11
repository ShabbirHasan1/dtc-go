package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size         = SymbolSearchRequestFixedSize  (96)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = 0
//     SearchType   = 0
// }
var _SymbolSearchRequestFixedDefault = []byte{96, 0, 252, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SymbolSearchRequestFixedSize = 96

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
	message.Fixed
}

// SymbolSearchRequestFixedBuilder The SymbolSearchRequest message is sent by the Client to the Server to
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
type SymbolSearchRequestFixedBuilder struct {
	message.Fixed
}

func NewSymbolSearchRequestFixedFrom(b []byte) SymbolSearchRequestFixed {
	return SymbolSearchRequestFixed{Fixed: message.NewFixedFrom(b)}
}

func WrapSymbolSearchRequestFixed(b []byte) SymbolSearchRequestFixed {
	return SymbolSearchRequestFixed{Fixed: message.WrapFixed(b)}
}

func NewSymbolSearchRequestFixed() SymbolSearchRequestFixedBuilder {
	a := SymbolSearchRequestFixedBuilder{message.NewFixed(96)}
	a.Unsafe().SetBytes(0, _SymbolSearchRequestFixedDefault)
	return a
}

// Clear
// {
//     Size         = SymbolSearchRequestFixedSize  (96)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = 0
//     SearchType   = 0
// }
func (m SymbolSearchRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolSearchRequestFixedDefault)
}

// ToBuilder
func (m SymbolSearchRequestFixed) ToBuilder() SymbolSearchRequestFixedBuilder {
	return SymbolSearchRequestFixedBuilder{m.Fixed}
}

// Finish
func (m SymbolSearchRequestFixedBuilder) Finish() SymbolSearchRequestFixed {
	return SymbolSearchRequestFixed{m.Fixed}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixed) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixedBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Unsafe(), 8, 4)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixed) SearchText() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixedBuilder) SearchText() string {
	return message.StringFixed(m.Unsafe(), 72, 8)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixed) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32Fixed(m.Unsafe(), 96, 92))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedBuilder) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32Fixed(m.Unsafe(), 96, 92))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixedBuilder) SetSearchText(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 92, 88, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 96, 92, int32(value))
}

// Copy
func (m SymbolSearchRequestFixed) Copy(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// Copy
func (m SymbolSearchRequestFixedBuilder) Copy(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequestFixed) CopyPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequestFixedBuilder) CopyPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequestFixed) CopyToPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequestFixedBuilder) CopyToPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestFixed) CopyTo(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestFixedBuilder) CopyTo(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}
