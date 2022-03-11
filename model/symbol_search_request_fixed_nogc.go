package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SymbolSearchRequestFixedPointer The SymbolSearchRequest message is sent by the Client to the Server to
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
type SymbolSearchRequestFixedPointer struct {
	message.FixedPointer
}

// SymbolSearchRequestFixedPointerBuilder The SymbolSearchRequest message is sent by the Client to the Server to
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
type SymbolSearchRequestFixedPointerBuilder struct {
	message.FixedPointer
}

func AllocSymbolSearchRequestFixed() SymbolSearchRequestFixedPointerBuilder {
	a := SymbolSearchRequestFixedPointerBuilder{message.AllocFixed(96)}
	a.Ptr.SetBytes(0, _SymbolSearchRequestFixedDefault)
	return a
}

func AllocSymbolSearchRequestFixedFrom(b []byte) SymbolSearchRequestFixedPointer {
	return SymbolSearchRequestFixedPointer{FixedPointer: message.AllocFixedFrom(b)}
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
func (m SymbolSearchRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolSearchRequestFixedDefault)
}

// ToBuilder
func (m SymbolSearchRequestFixedPointer) ToBuilder() SymbolSearchRequestFixedPointerBuilder {
	return SymbolSearchRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m *SymbolSearchRequestFixedPointerBuilder) Finish() SymbolSearchRequestFixedPointer {
	return SymbolSearchRequestFixedPointer{m.FixedPointer}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixedPointer) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixedPointerBuilder) RequestID() int32 {
	return message.Int32Fixed(m.Ptr, 8, 4)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixedPointer) SearchText() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixedPointerBuilder) SearchText() string {
	return message.StringFixed(m.Ptr, 72, 8)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedPointer) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedPointerBuilder) Exchange() string {
	return message.StringFixed(m.Ptr, 88, 72)
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 92, 88))
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Ptr, 92, 88))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedPointer) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32Fixed(m.Ptr, 96, 92))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedPointerBuilder) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32Fixed(m.Ptr, 96, 92))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixedPointerBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Ptr, 8, 4, value)
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixedPointerBuilder) SetSearchText(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 92, 88, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedPointerBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 96, 92, int32(value))
}

// Copy
func (m SymbolSearchRequestFixedPointer) Copy(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// Copy
func (m SymbolSearchRequestFixedPointerBuilder) Copy(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequestFixedPointer) CopyPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequestFixedPointerBuilder) CopyPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestFixedPointer) CopyTo(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestFixedPointerBuilder) CopyTo(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequestFixedPointer) CopyToPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequestFixedPointerBuilder) CopyToPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}
