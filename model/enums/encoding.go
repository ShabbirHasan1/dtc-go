package enums

type Encoding int32

const (
	EncodingBinaryEncoding                  Encoding = 0
	EncodingBinaryWithVariableLengthStrings Encoding = 1
	EncodingJson                            Encoding = 2
	EncodingJsonCompact                     Encoding = 3
	EncodingProtocolBuffers                 Encoding = 4
)
