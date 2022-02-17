package dtc

import (
	"unsafe"
)

const (
	// DTC protocol version
	CurrentVersion int32 = 8

	// Text string lengths when using fixed length string binary encoding.
	UsernamePasswordLength       int32 = 32
	SymbolExchangeDelimterLength int32 = 4
	ExchangeLength               int32 = 16
	UnderlyingSymbolLength       int32 = 32
	SymbolDescriptionLength      int32 = 64
	ExchangeDescriptionLength    int32 = 48
	OrderIDLength                int32 = 32
	TradeAccountLength           int32 = 32
	TradeDescriptionLength       int32 = 96
	TextMessageLength            int32 = 256
	OrderFreeFormTextLength      int32 = 48
	ClientServerNameLength       int32 = 48
	GeneralIdentifierLength      int32 = 64
	CurrencyCodeLength           int32 = 8
	OrderFillExecutionLength     int32 = 64
	PriceStringLength            int32 = 16
)

type _string struct {
	Data uintptr
	Len  int
}

func strlen(b []byte) int {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return i
		}
	}
	return len(b)
}

/*
struct s_EncodingRequest
	{
		uint16_t Size;
		uint16_t Type;
		int32_t ProtocolVersion;
		EncodingEnum Encoding;
		char ProtocolType[4];

		s_EncodingRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(*this));
			Type = ENCODING_REQUEST;
			Size = sizeof(*this);

			ProtocolVersion = CURRENT_VERSION;
			Encoding = BINARY_ENCODING;
			SetProtocolType("DTC");
		}

		int32_t GetProtocolVersion() const;
		EncodingEnum GetEncoding() const;
		const char* GetProtocolType();
		void SetProtocolType(const char* NewValue);
	};
*/
type EncodingRequest struct {
	Size            uint16
	Type            uint16
	ProtocolVersion int32
	Encoding        EncodingEnum
	ProtocolType    [4]byte
}

func NewEncodingRequest() EncodingRequest {
	var er EncodingRequest
	er.Clear()
	return er
}

func (er *EncodingRequest) Clear() {
	*er = EncodingRequest{}
	er.ProtocolVersion = CurrentVersion
	er.Encoding = EncodingEnum_BINARY
}

func (er *EncodingRequest) GetProtocolType() string {
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&er.ProtocolType[0])),
		Len:  strlen(er.ProtocolType[0:cap(er.ProtocolType)]),
	}))
}

type EncodingResponse struct {
	Size            uint16
	Type            uint16
	ProtocolVersion int32
	Encoding        EncodingEnum
	ProtocolType    [4]byte
}

func NewEncodingResponse() EncodingResponse {
	var er EncodingResponse
	er.Clear()
	return er
}

func (er *EncodingResponse) Clear() {
	*er = EncodingResponse{}
	er.ProtocolVersion = CurrentVersion
	er.Encoding = EncodingEnum_BINARY
}

func (er *EncodingResponse) GetProtocolType() string {
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&er.ProtocolType[0])),
		Len:  strlen(er.ProtocolType[0:cap(er.ProtocolType)]),
	}))
}
