// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const OpenOrdersRequestSize = 24

const OpenOrdersRequestFixedSize = 76

//     Size              uint16  = OpenOrdersRequestSize  (24)
//     Type              uint16  = OPEN_ORDERS_REQUEST  (300)
//     BaseSize          uint16  = OpenOrdersRequestSize  (24)
//     RequestID         int32   = 0
//     RequestAllOrders  int32   = 1
//     ServerOrderID     string  = ""
//     TradeAccount      string  = ""
var _OpenOrdersRequestDefault = []byte{24, 0, 44, 1, 24, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size              uint16      = OpenOrdersRequestFixedSize  (76)
//     Type              uint16      = OPEN_ORDERS_REQUEST  (300)
//     RequestID         int32       = 0
//     RequestAllOrders  int32       = 1
//     ServerOrderID     string[32]  = ""
//     TradeAccount      string[32]  = ""
var _OpenOrdersRequestFixedDefault = []byte{76, 0, 44, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// OpenOrdersRequest This is a message from the Client to the Server requesting the currently
// open orders.
//
// The Server will send open/working orders in response to this request through
// OrderUpdate messages.
//
// The Server will not return canceled or filled orders.
//
// When the Server responds to this request, it needs to respond with a separate
// OrderUpdate for each order.
//
// When the Server responds to this request, OrderUpdateReason in the OrderUpdate
// message must be set to OpenOrdersRequest_RESPONSE indicating the orders
// are being restated.
//
// If there are no Open orders, the Server will send back 1 OrderUpdate message
// with only the TotalNumberMessages, MessageNumber, RequestID, OrderUpdateReason,
// NoOrders = 1 fields set in the OrderUpdate message.
type OpenOrdersRequest struct {
	p message.VLS
}

// OpenOrdersRequestFixed This is a message from the Client to the Server requesting the currently
// open orders.
//
// The Server will send open/working orders in response to this request through
// OrderUpdate messages.
//
// The Server will not return canceled or filled orders.
//
// When the Server responds to this request, it needs to respond with a separate
// OrderUpdate for each order.
//
// When the Server responds to this request, OrderUpdateReason in the OrderUpdate
// message must be set to OpenOrdersRequest_RESPONSE indicating the orders
// are being restated.
//
// If there are no Open orders, the Server will send back 1 OrderUpdate message
// with only the TotalNumberMessages, MessageNumber, RequestID, OrderUpdateReason,
// NoOrders = 1 fields set in the OrderUpdate message.
type OpenOrdersRequestFixed struct {
	p message.Fixed
}

func NewOpenOrdersRequestFrom(b []byte) OpenOrdersRequest {
	return OpenOrdersRequest{p: message.NewVLS(b)}
}

func WrapOpenOrdersRequest(b []byte) OpenOrdersRequest {
	return OpenOrdersRequest{p: message.WrapVLS(b)}
}

func NewOpenOrdersRequest() *OpenOrdersRequest {
	return &OpenOrdersRequest{p: message.NewVLS(_OpenOrdersRequestDefault)}
}

func ParseOpenOrdersRequest(b []byte) (OpenOrdersRequest, error) {
	if len(b) < 6 {
		return OpenOrdersRequest{}, message.ErrShortBuffer
	}
	m := WrapOpenOrdersRequest(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return OpenOrdersRequest{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return OpenOrdersRequest{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 24 {
		newSize := len(b) + (24 - baseSize)
		if newSize > message.MaxSize {
			return OpenOrdersRequest{}, message.ErrOverflow
		}
		clone := OpenOrdersRequest{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _OpenOrdersRequestDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(24 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(16)
			if offset > 0 {
				clone.p.SetUint16LE(16, offset+shift)
			}
			offset = clone.p.Uint16LE(20)
			if offset > 0 {
				clone.p.SetUint16LE(20, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewOpenOrdersRequestFixedFrom(b []byte) OpenOrdersRequestFixed {
	return OpenOrdersRequestFixed{p: message.NewFixed(b)}
}

func WrapOpenOrdersRequestFixed(b []byte) OpenOrdersRequestFixed {
	return OpenOrdersRequestFixed{p: message.WrapFixed(b)}
}

func NewOpenOrdersRequestFixed() *OpenOrdersRequestFixed {
	return &OpenOrdersRequestFixed{p: message.NewFixed(_OpenOrdersRequestFixedDefault)}
}

func ParseOpenOrdersRequestFixed(b []byte) (OpenOrdersRequestFixed, error) {
	if len(b) < 4 {
		return OpenOrdersRequestFixed{}, message.ErrShortBuffer
	}
	m := WrapOpenOrdersRequestFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return OpenOrdersRequestFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return OpenOrdersRequestFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 76 {
		clone := *NewOpenOrdersRequestFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _OpenOrdersRequestFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m OpenOrdersRequest) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m OpenOrdersRequest) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m OpenOrdersRequest) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequest) RequestID() int32 {
	return m.p.Int32LE(8)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequest) RequestAllOrders() int32 {
	return m.p.Int32LE(12)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequest) ServerOrderID() string {
	return m.p.StringVLS(16)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequest) TradeAccount() string {
	return m.p.StringVLS(20)
}

// Size The standard message size field. Automatically set by constructor.
func (m OpenOrdersRequestFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m OpenOrdersRequestFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// RequestID A unique request identifier for this request.
func (m OpenOrdersRequestFixed) RequestID() int32 {
	return m.p.Int32LE(4)
}

// RequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m OpenOrdersRequestFixed) RequestAllOrders() int32 {
	return m.p.Int32LE(8)
}

// ServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m OpenOrdersRequestFixed) ServerOrderID() string {
	return m.p.StringFixed(12, 32)
}

// TradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m OpenOrdersRequestFixed) TradeAccount() string {
	return m.p.StringFixed(44, 32)
}

// SetRequestID A unique request identifier for this request.
func (m *OpenOrdersRequest) SetRequestID(value int32) *OpenOrdersRequest {
	m.p.SetInt32LE(8, value)
	return m
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m *OpenOrdersRequest) SetRequestAllOrders(value int32) *OpenOrdersRequest {
	m.p.SetInt32LE(12, value)
	return m
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m *OpenOrdersRequest) SetServerOrderID(value string) *OpenOrdersRequest {
	m.p.SetStringVLS(16, value)
	return m
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m *OpenOrdersRequest) SetTradeAccount(value string) *OpenOrdersRequest {
	m.p.SetStringVLS(20, value)
	return m
}

// SetRequestID A unique request identifier for this request.
func (m *OpenOrdersRequestFixed) SetRequestID(value int32) *OpenOrdersRequestFixed {
	m.p.SetInt32LE(4, value)
	return m
}

// SetRequestAllOrders 0 = request a specific order. 1 = for all orders (default).
func (m *OpenOrdersRequestFixed) SetRequestAllOrders(value int32) *OpenOrdersRequestFixed {
	m.p.SetInt32LE(8, value)
	return m
}

// SetServerOrderID Leave this field empty if Client wants all orders. Otherwise, specify
// the specific ServerServerID that want an OrderUpdate message for.
func (m *OpenOrdersRequestFixed) SetServerOrderID(value string) *OpenOrdersRequestFixed {
	m.p.SetStringFixed(12, 32, value)
	return m
}

// SetTradeAccount Leave this field empty if Client wants open orders for all Trade Accounts.
// Otherwise, specify the specific TradeAccount that want OrderUpdate messages
// for.
func (m *OpenOrdersRequestFixed) SetTradeAccount(value string) *OpenOrdersRequestFixed {
	m.p.SetStringFixed(44, 32, value)
	return m
}

func (m OpenOrdersRequest) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m OpenOrdersRequest) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m OpenOrdersRequestFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m OpenOrdersRequestFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m OpenOrdersRequest) Copy(to OpenOrdersRequest) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequest) CopyTo(to OpenOrdersRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m OpenOrdersRequestFixed) Copy(to OpenOrdersRequestFixed) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

// CopyTo
func (m OpenOrdersRequestFixed) CopyTo(to OpenOrdersRequest) {
	to.SetRequestID(m.RequestID())
	to.SetRequestAllOrders(m.RequestAllOrders())
	to.SetServerOrderID(m.ServerOrderID())
	to.SetTradeAccount(m.TradeAccount())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequest) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersRequest) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 300 {
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
		case "RequestAllOrders":
			m.SetRequestAllOrders(r.Int32())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *OpenOrdersRequest) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m OpenOrdersRequestFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 300)
	w.Int32Field("RequestID", m.RequestID())
	w.Int32Field("RequestAllOrders", m.RequestAllOrders())
	w.StringField("ServerOrderID", m.ServerOrderID())
	w.StringField("TradeAccount", m.TradeAccount())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *OpenOrdersRequestFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 300 {
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
		case "RequestAllOrders":
			m.SetRequestAllOrders(r.Int32())
		case "ServerOrderID":
			m.SetServerOrderID(r.String())
		case "TradeAccount":
			m.SetTradeAccount(r.String())
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

func (m *OpenOrdersRequestFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
