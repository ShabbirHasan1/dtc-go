// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 22:08:18.145964 +0800 WITA m=+0.011497918

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const HeartbeatSize = 16

//     Size                uint16    = HeartbeatSize  (16)
//     Type                uint16    = HEARTBEAT  (3)
//     NumDroppedMessages  uint32    = 0
//     CurrentDateTime     DateTime  = 0
var _HeartbeatDefault = []byte{16, 0, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// Heartbeat Both the Client and the Server need to send to the other side a heartbeat
// at the interval specified by the HeartbeatIntervalInSeconds member in
// the LogonRequest.
//
// There are no required member fields to set in this message. The purpose
// of the Heartbeat message is so that the Client or the Server can determine
// whether the other side is still connected.
//
// It is recommended that if there is a loss of Heartbeat messages from the
// other side, for twice the amount of the HeartbeatIntervalInSeconds time
// that it is safe to assume that the other side is no longer present and
// the network socket should be then gracefully closed.
//
// The Server may choose to send a heartbeat message every second to the
// Client. In this particular case, it is recommended the Client use a minimum
// time of about 5 to 10 seconds without a heartbeat to determine the loss
// of the connection rather than the standard of twice the amount of the
// heartbeat time interval.
type Heartbeat struct {
	p message.Fixed
}

func NewHeartbeatFrom(b []byte) Heartbeat {
	return Heartbeat{p: message.NewFixed(b)}
}

func WrapHeartbeat(b []byte) Heartbeat {
	return Heartbeat{p: message.WrapFixed(b)}
}

func NewHeartbeat() *Heartbeat {
	return &Heartbeat{p: message.NewFixed(_HeartbeatDefault)}
}

func ParseHeartbeat(b []byte) (Heartbeat, error) {
	if len(b) < 4 {
		return Heartbeat{}, message.ErrShortBuffer
	}
	m := WrapHeartbeat(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return Heartbeat{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return Heartbeat{}, message.ErrBaseSizeOverflow
	}
	if size < 16 {
		clone := *NewHeartbeat()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _HeartbeatDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m Heartbeat) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
//
// To determine the field number for JSON, refer to this message type constant
// in the DTCProtocol.h file.
func (m Heartbeat) Type() uint16 {
	return m.p.Uint16LE(2)
}

// NumDroppedMessages
func (m Heartbeat) NumDroppedMessages() uint32 {
	return m.p.Uint32LE(4)
}

// CurrentDateTime
func (m Heartbeat) CurrentDateTime() DateTime {
	return DateTime(m.p.Int64LE(8))
}

// SetNumDroppedMessages
func (m *Heartbeat) SetNumDroppedMessages(value uint32) *Heartbeat {
	m.p.SetUint32LE(4, value)
	return m
}

// SetCurrentDateTime
func (m *Heartbeat) SetCurrentDateTime(value DateTime) *Heartbeat {
	m.p.SetInt64LE(8, int64(value))
	return m
}

func (m *Heartbeat) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m *Heartbeat) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m *Heartbeat) Clone() *Heartbeat {
	return &Heartbeat{message.WrapFixedPointer(m.p.Clone(uintptr(m.Size())))}
}

// Copy
func (m Heartbeat) Copy(to Heartbeat) {
	to.SetNumDroppedMessages(m.NumDroppedMessages())
	to.SetCurrentDateTime(m.CurrentDateTime())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *Heartbeat) MarshalJSON() ([]byte, error) {
	return m.MarshalJSONTo(nil)
}

func (m *Heartbeat) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 3)
	w.Uint32Field("NumDroppedMessages", m.NumDroppedMessages())
	w.Int64Field("CurrentDateTime", int64(m.CurrentDateTime()))
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *Heartbeat) UnmarshalJSON(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONFromReader(&r)
}

func (m *Heartbeat) UnmarshalJSONFromReader(r *json.Reader) error {
	if r.Type != 3 {
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
		case "NumDroppedMessages":
			m.SetNumDroppedMessages(r.Uint32())
		case "CurrentDateTime":
			m.SetCurrentDateTime(DateTime(r.Int64()))
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