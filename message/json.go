//go:build !tinygo

package message

type JsonMarshaller interface {
	MarshalJSONTo(b []byte) ([]byte, error)
}

func ToJsonString(m JsonMarshaller) string {
	if m == nil {
		return ""
	}
	b, _ := m.MarshalJSONTo(nil)
	return string(b)
}

func JsonString(m JsonMarshaller) string {
	b, _ := m.MarshalJSONTo(nil)
	return string(b)
}
