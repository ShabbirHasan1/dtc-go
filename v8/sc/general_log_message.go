// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const GeneralLogMessageSize = 10

const GeneralLogMessageFixedSize = 132

//     Size         uint16  = GeneralLogMessageSize  (10)
//     Type         uint16  = GENERAL_LOG_MESSAGE  (701)
//     BaseSize     uint16  = GeneralLogMessageSize  (10)
//     MessageText  string  = ""
var _GeneralLogMessageDefault = []byte{10, 0, 189, 2, 10, 0, 0, 0, 0, 0}

//     Size         uint16       = GeneralLogMessageFixedSize  (132)
//     Type         uint16       = GENERAL_LOG_MESSAGE  (701)
//     MessageText  string[128]  = ""
var _GeneralLogMessageFixedDefault = []byte{132, 0, 189, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// GeneralLogMessage This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessage struct {
	p message.VLS
}

// GeneralLogMessageFixed This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
type GeneralLogMessageFixed struct {
	p message.Fixed
}

func NewGeneralLogMessageFrom(b []byte) GeneralLogMessage {
	return GeneralLogMessage{p: message.NewVLS(b)}
}

func WrapGeneralLogMessage(b []byte) GeneralLogMessage {
	return GeneralLogMessage{p: message.WrapVLS(b)}
}

func NewGeneralLogMessage() *GeneralLogMessage {
	return &GeneralLogMessage{p: message.NewVLS(_GeneralLogMessageDefault)}
}

func ParseGeneralLogMessage(b []byte) (GeneralLogMessage, error) {
	if len(b) < 6 {
		return GeneralLogMessage{}, message.ErrShortBuffer
	}
	m := WrapGeneralLogMessage(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return GeneralLogMessage{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return GeneralLogMessage{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 10 {
		newSize := len(b) + (10 - baseSize)
		if newSize > message.MaxSize {
			return GeneralLogMessage{}, message.ErrOverflow
		}
		clone := GeneralLogMessage{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _GeneralLogMessageDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(10 - baseSize)
			var offset uint16
			offset = clone.p.Uint16LE(6)
			if offset > 0 {
				clone.p.SetUint16LE(6, offset+shift)
			}
		}
		return clone, nil
	}
	return m, nil
}

func NewGeneralLogMessageFixedFrom(b []byte) GeneralLogMessageFixed {
	return GeneralLogMessageFixed{p: message.NewFixed(b)}
}

func WrapGeneralLogMessageFixed(b []byte) GeneralLogMessageFixed {
	return GeneralLogMessageFixed{p: message.WrapFixed(b)}
}

func NewGeneralLogMessageFixed() *GeneralLogMessageFixed {
	return &GeneralLogMessageFixed{p: message.NewFixed(_GeneralLogMessageFixedDefault)}
}

func ParseGeneralLogMessageFixed(b []byte) (GeneralLogMessageFixed, error) {
	if len(b) < 4 {
		return GeneralLogMessageFixed{}, message.ErrShortBuffer
	}
	m := WrapGeneralLogMessageFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return GeneralLogMessageFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return GeneralLogMessageFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 132 {
		clone := *NewGeneralLogMessageFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _GeneralLogMessageFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m GeneralLogMessage) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m GeneralLogMessage) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m GeneralLogMessage) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessage) MessageText() string {
	return m.p.StringVLS(6)
}

// Size The standard message size field. Automatically set by constructor.
func (m GeneralLogMessageFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m GeneralLogMessageFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// MessageText The message text to the Client which it should add to its log.
//
func (m GeneralLogMessageFixed) MessageText() string {
	return m.p.StringFixed(4, 128)
}

// SetMessageText The message text to the Client which it should add to its log.
//
func (m *GeneralLogMessage) SetMessageText(value string) *GeneralLogMessage {
	m.p.SetStringVLS(6, value)
	return m
}

// SetMessageText The message text to the Client which it should add to its log.
//
func (m *GeneralLogMessageFixed) SetMessageText(value string) *GeneralLogMessageFixed {
	m.p.SetStringFixed(4, 128, value)
	return m
}

func (m GeneralLogMessage) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m GeneralLogMessage) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m GeneralLogMessageFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m GeneralLogMessageFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m GeneralLogMessage) Copy(to GeneralLogMessage) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessage) CopyTo(to GeneralLogMessageFixed) {
	to.SetMessageText(m.MessageText())
}

// Copy
func (m GeneralLogMessageFixed) Copy(to GeneralLogMessageFixed) {
	to.SetMessageText(m.MessageText())
}

// CopyTo
func (m GeneralLogMessageFixed) CopyTo(to GeneralLogMessage) {
	to.SetMessageText(m.MessageText())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m GeneralLogMessage) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 701)
	w.StringField("MessageText", m.MessageText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *GeneralLogMessage) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 701 {
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
		case "MessageText":
			m.SetMessageText(r.String())
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

func (m *GeneralLogMessage) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m GeneralLogMessageFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 701)
	w.StringField("MessageText", m.MessageText())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *GeneralLogMessageFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 701 {
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
		case "MessageText":
			m.SetMessageText(r.String())
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

func (m *GeneralLogMessageFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
