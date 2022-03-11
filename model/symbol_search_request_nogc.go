package model

import (
	"github.com/moontrade/dtc-go/message"
)

// SymbolSearchRequestPointer The SymbolSearchRequest message is sent by the Client to the Server to
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
type SymbolSearchRequestPointer struct {
	message.VLSPointer
}

// SymbolSearchRequestPointerBuilder The SymbolSearchRequest message is sent by the Client to the Server to
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
type SymbolSearchRequestPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocSymbolSearchRequest() SymbolSearchRequestPointerBuilder {
	a := SymbolSearchRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _SymbolSearchRequestDefault)
	return a
}

func AllocSymbolSearchRequestFrom(b []byte) SymbolSearchRequestPointer {
	return SymbolSearchRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = SymbolSearchRequestSize  (28)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     BaseSize     = SymbolSearchRequestSize  (28)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = 0
//     SearchType   = 0
// }
func (m SymbolSearchRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolSearchRequestDefault)
}

// ToBuilder
func (m SymbolSearchRequestPointer) ToBuilder() SymbolSearchRequestPointerBuilder {
	return SymbolSearchRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *SymbolSearchRequestPointerBuilder) Finish() SymbolSearchRequestPointer {
	return SymbolSearchRequestPointer{m.VLSPointerBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestPointer) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestPointerBuilder) RequestID() int32 {
	return message.Int32VLS(m.Ptr, 12, 8)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestPointer) SearchText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestPointerBuilder) SearchText() string {
	return message.StringVLS(m.Ptr, 16, 12)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestPointer) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestPointerBuilder) Exchange() string {
	return message.StringVLS(m.Ptr, 20, 16)
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestPointer) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestPointerBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Ptr, 24, 20))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestPointer) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestPointerBuilder) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32VLS(m.Ptr, 28, 24))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestPointerBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Ptr, 12, 8, value)
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m *SymbolSearchRequestPointerBuilder) SetSearchText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m *SymbolSearchRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestPointerBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// Copy
func (m SymbolSearchRequestPointer) Copy(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// Copy
func (m SymbolSearchRequestPointerBuilder) Copy(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequestPointer) CopyPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequestPointerBuilder) CopyPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestPointer) CopyTo(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestPointerBuilder) CopyTo(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequestPointer) CopyToPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequestPointerBuilder) CopyToPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}
