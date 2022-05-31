// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-31 21:46:24.141427 +0800 WITA m=+0.018610501

package v8

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/message/json"
	"io"
)

const JournalEntryResponseSize = 32

const JournalEntryResponseFixedSize = 280

//     Size             uint16    = JournalEntryResponseSize  (32)
//     Type             uint16    = JOURNAL_ENTRY_RESPONSE  (706)
//     BaseSize         uint16    = JournalEntryResponseSize  (32)
//     JournalEntry     string    = ""
//     DateTime         DateTime  = 0
//     IsFinalResponse  bool      = false
var _JournalEntryResponseDefault = []byte{32, 0, 194, 2, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//     Size             uint16       = JournalEntryResponseFixedSize  (280)
//     Type             uint16       = JOURNAL_ENTRY_RESPONSE  (706)
//     JournalEntry     string[256]  = ""
//     DateTime         DateTime     = 0
//     IsFinalResponse  bool         = false
var _JournalEntryResponseFixedDefault = []byte{24, 1, 194, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type JournalEntryResponse struct {
	p message.VLS
}

type JournalEntryResponseFixed struct {
	p message.Fixed
}

func NewJournalEntryResponseFrom(b []byte) JournalEntryResponse {
	return JournalEntryResponse{p: message.NewVLS(b)}
}

func WrapJournalEntryResponse(b []byte) JournalEntryResponse {
	return JournalEntryResponse{p: message.WrapVLS(b)}
}

func NewJournalEntryResponse() *JournalEntryResponse {
	return &JournalEntryResponse{p: message.NewVLS(_JournalEntryResponseDefault)}
}

func ParseJournalEntryResponse(b []byte) (JournalEntryResponse, error) {
	if len(b) < 6 {
		return JournalEntryResponse{}, message.ErrShortBuffer
	}
	m := WrapJournalEntryResponse(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return JournalEntryResponse{}, message.ErrOverflow
	}
	baseSize := int(m.p.Uint16LE(4))
	if baseSize > len(b) {
		return JournalEntryResponse{}, message.ErrBaseSizeOverflow
	}
	if baseSize < 32 {
		newSize := len(b) + (32 - baseSize)
		if newSize > message.MaxSize {
			return JournalEntryResponse{}, message.ErrOverflow
		}
		clone := JournalEntryResponse{message.WrapVLSUnsafe(message.Alloc(uintptr(newSize)), len(b))}
		clone.p.SetBytes(0, b[0:baseSize])
		clone.p.SetBytes(baseSize, _JournalEntryResponseDefault[baseSize:])
		if len(b) > baseSize {
			shift := uint16(32 - baseSize)
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

func NewJournalEntryResponseFixedFrom(b []byte) JournalEntryResponseFixed {
	return JournalEntryResponseFixed{p: message.NewFixed(b)}
}

func WrapJournalEntryResponseFixed(b []byte) JournalEntryResponseFixed {
	return JournalEntryResponseFixed{p: message.WrapFixed(b)}
}

func NewJournalEntryResponseFixed() *JournalEntryResponseFixed {
	return &JournalEntryResponseFixed{p: message.NewFixed(_JournalEntryResponseFixedDefault)}
}

func ParseJournalEntryResponseFixed(b []byte) (JournalEntryResponseFixed, error) {
	if len(b) < 4 {
		return JournalEntryResponseFixed{}, message.ErrShortBuffer
	}
	m := WrapJournalEntryResponseFixed(b)
	if int(m.p.AsUint16LE()) != len(b) {
		return JournalEntryResponseFixed{}, message.ErrOverflow
	}
	size := int(m.p.AsUint16LE())
	if size > len(b) {
		return JournalEntryResponseFixed{}, message.ErrBaseSizeOverflow
	}
	if size < 280 {
		clone := *NewJournalEntryResponseFixed()
		clone.p.SetBytes(0, b[0:size])
		clone.p.SetBytes(size, _JournalEntryResponseFixedDefault[size:])
		return clone, nil
	}
	return m, nil
}

// Size
func (m JournalEntryResponse) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m JournalEntryResponse) Type() uint16 {
	return m.p.Uint16LE(2)
}

// BaseSize
func (m JournalEntryResponse) BaseSize() uint16 {
	return m.p.Uint16LE(4)
}

// JournalEntry
func (m JournalEntryResponse) JournalEntry() string {
	return m.p.StringVLS(6)
}

// DateTime
func (m JournalEntryResponse) DateTime() DateTime {
	return DateTime(m.p.Int64LE(16))
}

// IsFinalResponse
func (m JournalEntryResponse) IsFinalResponse() bool {
	return m.p.Bool(24)
}

// Size
func (m JournalEntryResponseFixed) Size() uint16 {
	return m.p.Uint16LE(0)
}

// Type
func (m JournalEntryResponseFixed) Type() uint16 {
	return m.p.Uint16LE(2)
}

// JournalEntry
func (m JournalEntryResponseFixed) JournalEntry() string {
	return m.p.StringFixed(4, 256)
}

// DateTime
func (m JournalEntryResponseFixed) DateTime() DateTime {
	return DateTime(m.p.Int64LE(264))
}

// IsFinalResponse
func (m JournalEntryResponseFixed) IsFinalResponse() bool {
	return m.p.Bool(272)
}

// SetJournalEntry
func (m *JournalEntryResponse) SetJournalEntry(value string) *JournalEntryResponse {
	m.p.SetStringVLS(6, value)
	return m
}

// SetDateTime
func (m *JournalEntryResponse) SetDateTime(value DateTime) *JournalEntryResponse {
	m.p.SetInt64LE(16, int64(value))
	return m
}

// SetIsFinalResponse
func (m *JournalEntryResponse) SetIsFinalResponse(value bool) *JournalEntryResponse {
	m.p.SetBool(24, value)
	return m
}

// SetJournalEntry
func (m *JournalEntryResponseFixed) SetJournalEntry(value string) *JournalEntryResponseFixed {
	m.p.SetStringFixed(4, 256, value)
	return m
}

// SetDateTime
func (m *JournalEntryResponseFixed) SetDateTime(value DateTime) *JournalEntryResponseFixed {
	m.p.SetInt64LE(264, int64(value))
	return m
}

// SetIsFinalResponse
func (m *JournalEntryResponseFixed) SetIsFinalResponse(value bool) *JournalEntryResponseFixed {
	m.p.SetBool(272, value)
	return m
}

func (m JournalEntryResponse) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m JournalEntryResponse) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

func (m JournalEntryResponseFixed) WriteTo(w io.Writer) (int64, error) {
	s := int(m.Size())
	n, err := w.Write(m.p.AsBytes(s))
	return int64(n), err
}

func (m JournalEntryResponseFixed) MarshalBinary() ([]byte, error) {
	return m.p.AsBytes(int(m.Size())), nil
}

// Copy
func (m JournalEntryResponse) Copy(to JournalEntryResponse) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponse) CopyTo(to JournalEntryResponseFixed) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// Copy
func (m JournalEntryResponseFixed) Copy(to JournalEntryResponseFixed) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

// CopyTo
func (m JournalEntryResponseFixed) CopyTo(to JournalEntryResponse) {
	to.SetJournalEntry(m.JournalEntry())
	to.SetDateTime(m.DateTime())
	to.SetIsFinalResponse(m.IsFinalResponse())
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponse) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponse) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 706 {
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
		case "JournalEntry":
			m.SetJournalEntry(r.String())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
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

func (m *JournalEntryResponse) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Marshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m JournalEntryResponseFixed) MarshalJSONTo(b []byte) ([]byte, error) {
	w := json.NewWriter(b, 706)
	w.StringField("JournalEntry", m.JournalEntry())
	w.Int64Field("DateTime", int64(m.DateTime()))
	w.BoolField("IsFinalResponse", m.IsFinalResponse())
	return w.Finish(), nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// JSON Unmarshal
//////////////////////////////////////////////////////////////////////////////////////////

func (m *JournalEntryResponseFixed) UnmarshalJSONDoc(r *json.Reader) error {
	if r.Type != 706 {
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
		case "JournalEntry":
			m.SetJournalEntry(r.String())
		case "DateTime":
			m.SetDateTime(DateTime(r.Int64()))
		case "IsFinalResponse":
			m.SetIsFinalResponse(r.Bool())
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

func (m *JournalEntryResponseFixed) UnmarshalJSONReader(r *json.Reader) error {
	return m.UnmarshalJSONDoc(r)
}
