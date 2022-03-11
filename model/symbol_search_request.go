package model

import (
	"github.com/moontrade/dtc-go/message"
)

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
var _SymbolSearchRequestDefault = []byte{28, 0, 252, 1, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const SymbolSearchRequestSize = 28

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
	message.VLS
}

// SymbolSearchRequestBuilder The SymbolSearchRequest message is sent by the Client to the Server to
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
type SymbolSearchRequestBuilder struct {
	message.VLSBuilder
}

func NewSymbolSearchRequestFrom(b []byte) SymbolSearchRequest {
	return SymbolSearchRequest{VLS: message.NewVLSFrom(b)}
}

func WrapSymbolSearchRequest(b []byte) SymbolSearchRequest {
	return SymbolSearchRequest{VLS: message.WrapVLS(b)}
}

func NewSymbolSearchRequest() SymbolSearchRequestBuilder {
	a := SymbolSearchRequestBuilder{message.NewVLSBuilder(28)}
	a.Unsafe().SetBytes(0, _SymbolSearchRequestDefault)
	return a
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
func (m SymbolSearchRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolSearchRequestDefault)
}

// ToBuilder
func (m SymbolSearchRequest) ToBuilder() SymbolSearchRequestBuilder {
	return SymbolSearchRequestBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m SymbolSearchRequestBuilder) Finish() SymbolSearchRequest {
	return SymbolSearchRequest{m.VLSBuilder.Finish()}
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequest) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestBuilder) RequestID() int32 {
	return message.Int32VLS(m.Unsafe(), 12, 8)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequest) SearchText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// SearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestBuilder) SearchText() string {
	return message.StringVLS(m.Unsafe(), 16, 12)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequest) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestBuilder) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m *SymbolSearchRequestBuilder) SetSearchText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m *SymbolSearchRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// Copy
func (m SymbolSearchRequest) Copy(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// Copy
func (m SymbolSearchRequestBuilder) Copy(to SymbolSearchRequestBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequest) CopyPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyPointer
func (m SymbolSearchRequestBuilder) CopyPointer(to SymbolSearchRequestPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequest) CopyToPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyToPointer
func (m SymbolSearchRequestBuilder) CopyToPointer(to SymbolSearchRequestFixedPointerBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequest) CopyTo(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}

// CopyTo
func (m SymbolSearchRequestBuilder) CopyTo(to SymbolSearchRequestFixedBuilder) {
	to.SetRequestID(m.RequestID())
	to.SetSearchText(m.SearchText())
	to.SetExchange(m.Exchange())
	to.SetSecurityType(m.SecurityType())
	to.SetSearchType(m.SearchType())
}
