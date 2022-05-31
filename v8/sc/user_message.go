// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const UserMessageSize = 12

const UserMessageFixedSize = 262

//     Size            uint16  = UserMessageSize  (12)
//     Type            uint16  = USER_MESSAGE  (700)
//     BaseSize        uint16  = UserMessageSize  (12)
//     UserMessage     string  = ""
//     IsPopupMessage  bool    = true
var _UserMessageDefault = []byte{12, 0, 188, 2, 12, 0, 0, 0, 0, 0, 1, 0}

//     Size            uint16       = UserMessageFixedSize  (262)
//     Type            uint16       = USER_MESSAGE  (700)
//     UserMessage     string[256]  = ""
//     IsPopupMessage  bool         = true
var _UserMessageFixedDefault = []byte{6, 1, 188, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}

// UserMessage This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessage struct {
	p message.VLS
}

// UserMessageFixed This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
type UserMessageFixed struct {
	p message.Fixed
}

func NewUserMessageFrom(b []byte) UserMessage {
	return UserMessage{p: message.NewVLS(b)}
}

func WrapUserMessage(b []byte) UserMessage {
	return UserMessage{p: message.WrapVLS(b)}
}

func NewUserMessage() *UserMessage {
	return &UserMessage{p: message.NewVLS(_UserMessageDefault)}
}

func ParseUserMessage(b []byte) (UserMessage, error) {
	if len(b) < 6 {
		return UserMessage{}, message.ErrShortBuffer
	}
	m := WrapUserMessage(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return UserMessage{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return UserMessage{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 12 {
		newSize := len(b) + (12 - baseSize)
		if newSize > message.MaxSize {
			return UserMessage{}, message.ErrOverflow
		}
		clone := UserMessage{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _UserMessageDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(12 - baseSize)
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

func NewUserMessageFixedFrom(b []byte) UserMessageFixed {
	return UserMessageFixed{p: message.NewFixed(b)}
}

func WrapUserMessageFixed(b []byte) UserMessageFixed {
	return UserMessageFixed{p: message.WrapFixed(b)}
}

func NewUserMessageFixed() *UserMessageFixed {
	return &UserMessageFixed{p: message.NewFixed(_UserMessageFixedDefault)}
}

func ParseUserMessageFixed(b []byte) (UserMessageFixed, error) {
	if len(b) < 4 {
		return UserMessageFixed{}, message.ErrShortBuffer
	}
	m := WrapUserMessageFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return UserMessageFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return UserMessageFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 262 {
		clone := *NewUserMessageFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _UserMessageFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size The standard message size field. Automatically set by constructor.
func (m UserMessage) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m UserMessage) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m UserMessage) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// UserMessage General message to present to user in the Client.
func (m UserMessage) UserMessage() string {
	return m.p.StringVLS(6)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessage) IsPopupMessage() bool {
	return m.p.Bool(10)
}

// Size The standard message size field. Automatically set by constructor.
func (m UserMessageFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type The standard message type field. Automatically set by constructor.
func (m UserMessageFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// UserMessage General message to present to user in the Client.
func (m UserMessageFixed) UserMessage() string {
	return m.p.StringFixed(4, 256)
}

// IsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m UserMessageFixed) IsPopupMessage() bool {
	return m.p.Bool(260)
}

// SetUserMessage General message to present to user in the Client.
func (m *UserMessage) SetUserMessage(value string) *UserMessage {
	m.p.SetStringVLS(6, value)
	return m
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m *UserMessage) SetIsPopupMessage(value bool) *UserMessage {
	m.p.SetBool(10, value)
	return m
}

// SetUserMessage General message to present to user in the Client.
func (m *UserMessageFixed) SetUserMessage(value string) *UserMessageFixed {
	m.p.SetStringFixed(4, 256, value)
	return m
}

// SetIsPopupMessage The default for this is 1 which signifies that the Server would like the
// Client to present the message to the user in a way which will get their
// attention. Otherwise, set this to 0 to give the message lower priority
// (just add to a log).
func (m *UserMessageFixed) SetIsPopupMessage(value bool) *UserMessageFixed {
	m.p.SetBool(260, value)
	return m
}

func (m UserMessage) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m UserMessage) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m UserMessageFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m UserMessageFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m UserMessage) Copy(to UserMessage) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessage) CopyTo(to UserMessageFixed) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// Copy
func (m UserMessageFixed) Copy(to UserMessageFixed) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

// CopyTo
func (m UserMessageFixed) CopyTo(to UserMessage) {
	to.SetUserMessage(m.UserMessage())
	to.SetIsPopupMessage(m.IsPopupMessage())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessage) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessage) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 700 {
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
		case "UserMessage":
			m.SetUserMessage(r.String())
		case "IsPopupMessage":
			m.SetIsPopupMessage(r.Bool())
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

func (m *UserMessage) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m UserMessageFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 700)
	w.StringField("UserMessage", m.UserMessage())
	w.BoolField("IsPopupMessage", m.IsPopupMessage())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *UserMessageFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 700 {
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
		case "UserMessage":
			m.SetUserMessage(r.String())
		case "IsPopupMessage":
			m.SetIsPopupMessage(r.Bool())
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

func (m *UserMessageFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}