package model

import (
	"github.com/moontrade/dtc-go/message"
)

const SymbolSearchRequestSize = 28

const SymbolSearchRequestFixedSize = 96

// {
//     Size         = SymbolSearchRequestSize  (28)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     BaseSize     = SymbolSearchRequestSize  (28)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
//     SearchType   = SEARCH_TYPE_UNSET  (0)
// }
var _SymbolSearchRequestDefault = []byte{28, 0, 252, 1, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// {
//     Size         = SymbolSearchRequestFixedSize  (96)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
//     SearchType   = SEARCH_TYPE_UNSET  (0)
// }
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

func AllocSymbolSearchRequest() SymbolSearchRequestPointerBuilder {
	a := SymbolSearchRequestPointerBuilder{message.AllocVLSBuilder(28)}
	a.Ptr.SetBytes(0, _SymbolSearchRequestDefault)
	return a
}

func AllocSymbolSearchRequestFrom(b []byte) SymbolSearchRequestPointer {
	return SymbolSearchRequestPointer{VLSPointer: message.AllocVLSFrom(b)}
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
//     Size         = SymbolSearchRequestSize  (28)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     BaseSize     = SymbolSearchRequestSize  (28)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
//     SearchType   = SEARCH_TYPE_UNSET  (0)
// }
func (m SymbolSearchRequestBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolSearchRequestDefault)
}

// Clear
// {
//     Size         = SymbolSearchRequestFixedSize  (96)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
//     SearchType   = SEARCH_TYPE_UNSET  (0)
// }
func (m SymbolSearchRequestFixedBuilder) Clear() {
	m.Unsafe().SetBytes(0, _SymbolSearchRequestFixedDefault)
}

// Clear
// {
//     Size         = SymbolSearchRequestSize  (28)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     BaseSize     = SymbolSearchRequestSize  (28)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
//     SearchType   = SEARCH_TYPE_UNSET  (0)
// }
func (m SymbolSearchRequestPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolSearchRequestDefault)
}

// Clear
// {
//     Size         = SymbolSearchRequestFixedSize  (96)
//     Type         = SYMBOL_SEARCH_REQUEST  (508)
//     RequestID    = 0
//     SearchText   = ""
//     Exchange     = ""
//     SecurityType = SECURITY_TYPE_UNSET  (0)
//     SearchType   = SEARCH_TYPE_UNSET  (0)
// }
func (m SymbolSearchRequestFixedPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _SymbolSearchRequestFixedDefault)
}

// ToBuilder
func (m SymbolSearchRequest) ToBuilder() SymbolSearchRequestBuilder {
	return SymbolSearchRequestBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m SymbolSearchRequestFixed) ToBuilder() SymbolSearchRequestFixedBuilder {
	return SymbolSearchRequestFixedBuilder{m.Fixed}
}

// ToBuilder
func (m SymbolSearchRequestPointer) ToBuilder() SymbolSearchRequestPointerBuilder {
	return SymbolSearchRequestPointerBuilder{m.VLSPointer.ToBuilder()}
}

// ToBuilder
func (m SymbolSearchRequestFixedPointer) ToBuilder() SymbolSearchRequestFixedPointerBuilder {
	return SymbolSearchRequestFixedPointerBuilder{m.FixedPointer}
}

// Finish
func (m SymbolSearchRequestBuilder) Finish() SymbolSearchRequest {
	return SymbolSearchRequest{m.VLSBuilder.Finish()}
}

// Finish
func (m SymbolSearchRequestFixedBuilder) Finish() SymbolSearchRequestFixed {
	return SymbolSearchRequestFixed{m.Fixed}
}

// Finish
func (m *SymbolSearchRequestPointerBuilder) Finish() SymbolSearchRequestPointer {
	return SymbolSearchRequestPointer{m.VLSPointerBuilder.Finish()}
}

// Finish
func (m *SymbolSearchRequestFixedPointerBuilder) Finish() SymbolSearchRequestFixedPointer {
	return SymbolSearchRequestFixedPointer{m.FixedPointer}
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
func (m SymbolSearchRequest) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestBuilder) Exchange() string {
	return message.StringVLS(m.Unsafe(), 20, 16)
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
func (m SymbolSearchRequest) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32VLS(m.Unsafe(), 24, 20))
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
func (m SymbolSearchRequest) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestBuilder) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32VLS(m.Unsafe(), 28, 24))
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
func (m SymbolSearchRequestFixed) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
}

// Exchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) Exchange() string {
	return message.StringFixed(m.Unsafe(), 88, 72)
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
func (m SymbolSearchRequestFixed) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
}

// SecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) SecurityType() SecurityTypeEnum {
	return SecurityTypeEnum(message.Int32Fixed(m.Unsafe(), 92, 88))
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
func (m SymbolSearchRequestFixed) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32Fixed(m.Unsafe(), 96, 92))
}

// SearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedBuilder) SearchType() SearchTypeEnum {
	return SearchTypeEnum(message.Int32Fixed(m.Unsafe(), 96, 92))
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
func (m SymbolSearchRequestBuilder) SetRequestID(value int32) {
	message.SetInt32VLS(m.Unsafe(), 12, 8, value)
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
func (m *SymbolSearchRequestBuilder) SetSearchText(value string) {
	message.SetStringVLS(&m.VLSBuilder, 16, 12, value)
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m *SymbolSearchRequestPointerBuilder) SetSearchText(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 16, 12, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m *SymbolSearchRequestBuilder) SetExchange(value string) {
	message.SetStringVLS(&m.VLSBuilder, 20, 16, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m *SymbolSearchRequestPointerBuilder) SetExchange(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 20, 16, value)
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 24, 20, int32(value))
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32VLS(m.Ptr, 24, 20, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32VLS(m.Unsafe(), 28, 24, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestPointerBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32VLS(m.Ptr, 28, 24, int32(value))
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the SecurityDefinitionResponse message.
func (m SymbolSearchRequestFixedBuilder) SetRequestID(value int32) {
	message.SetInt32Fixed(m.Unsafe(), 8, 4, value)
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
func (m SymbolSearchRequestFixedBuilder) SetSearchText(value string) {
	message.SetStringFixed(m.Unsafe(), 72, 8, value)
}

// SetSearchText The search text to search the Symbol or the Description for.
//
// If the SearchText field is an empty text string, then the server should
// reject the request with a SecurityDefinitionReject message.
func (m SymbolSearchRequestFixedPointerBuilder) SetSearchText(value string) {
	message.SetStringFixed(m.Ptr, 72, 8, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Unsafe(), 88, 72, value)
}

// SetExchange The Exchange of the Symbol to search for.
func (m SymbolSearchRequestFixedPointerBuilder) SetExchange(value string) {
	message.SetStringFixed(m.Ptr, 88, 72, value)
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 92, 88, int32(value))
}

// SetSecurityType The Security Type of the Symbol to search for.
func (m SymbolSearchRequestFixedPointerBuilder) SetSecurityType(value SecurityTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 92, 88, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32Fixed(m.Unsafe(), 96, 92, int32(value))
}

// SetSearchType This field is the search type. Can be one of two possible values. Can
// specify to search the Symbol or the Description.
func (m SymbolSearchRequestFixedPointerBuilder) SetSearchType(value SearchTypeEnum) {
	message.SetInt32Fixed(m.Ptr, 96, 92, int32(value))
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
