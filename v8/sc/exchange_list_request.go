// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const ExchangeListRequestSize = 8

//     Size       uint16  = ExchangeListRequestSize  (8)
//     Type       uint16  = EXCHANGE_LIST_REQUEST  (500)
//     RequestID  int32   = 0
var _ExchangeListRequestDefault = []byte{8, 0, 244, 1, 0, 0, 0, 0}

// ExchangeListRequest This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
type ExchangeListRequest struct {
	p message.Fixed
}

func NewExchangeListRequestFrom(b []byte) ExchangeListRequest {
	return ExchangeListRequest{p: message.NewFixed(b)}
}

func WrapExchangeListRequest(b []byte) ExchangeListRequest {
	return ExchangeListRequest{p: message.WrapFixed(b)}
}

func NewExchangeListRequest() *ExchangeListRequest {
	return &ExchangeListRequest{p: message.NewFixed(_ExchangeListRequestDefault)}
}

func ParseExchangeListRequest(b []byte) (ExchangeListRequest, error) {
	if len(b) < 4 {
		return ExchangeListRequest{}, message.ErrShortBuffer
	}
	m := WrapExchangeListRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return ExchangeListRequest{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return ExchangeListRequest{}, message.ErrBaseSizeOverflow
	}
	if size < 8 {
		clone := *NewExchangeListRequest()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _ExchangeListRequestDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m ExchangeListRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m ExchangeListRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m ExchangeListRequest) RequestID() int32 {
	return m.p.Int32LE(4)
}

// SetRequestID The unique identifier for this request. This same identifier will be returned
// in the ExchangeListResponse message.
func (m *ExchangeListRequest) SetRequestID(value int32) *ExchangeListRequest {
	m.p.SetInt32LE(4, value)
	return m
}

func (m ExchangeListRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m ExchangeListRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m ExchangeListRequest) Copy(to ExchangeListRequest) {
	to.SetRequestID(m.RequestID())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m ExchangeListRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 500)
	w.Int32Field("RequestID", m.RequestID())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *ExchangeListRequest) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 500 {
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

func (m *ExchangeListRequest) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
