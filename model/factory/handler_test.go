package factory

import (
	"encoding/binary"
	"fmt"
	"github.com/moontrade/dtc-go/message"
	v8 "github.com/moontrade/dtc-go/model/v8"
	"github.com/moontrade/dtc-go/model/v8/sc"
	"github.com/moontrade/nogc"
	"testing"
	"unsafe"
)

func TestHandler(t *testing.T) {
	h := &Handler{AccountBalanceAdjustment: func(msg v8.AccountBalanceAdjustment) {
		if msg.Type() == v8.ACCOUNT_BALANCE_REJECT {

		}
	}}

	sc.AllocClientDeviceUpdate()

	_ = h

	f := &Factory{}
	b := f.AccountBalanceAdjustmentBuilder.Alloc()
	b.SetRequestID(101)
	a := b.Finish()

	fmt.Println(message.ToJsonString(b))
	a.Free()
}

func TestLayout(t *testing.T) {
	fmt.Println(unsafe.Sizeof(Struct{}))
}

type Struct struct {
	b [163]byte
}

func (s *Struct) Unsafe() nogc.Pointer {
	return nogc.Pointer(unsafe.Pointer(s))
}
func (s *Struct) Size() uint16 {
	message.Uint16Fixed(nogc.Pointer(unsafe.Pointer(&s)), 4, 0)
	return *(*uint16)(unsafe.Add(unsafe.Pointer(s), 0))
}
func (s *Struct) Type() uint16 {
	return *(*uint16)(unsafe.Add(unsafe.Pointer(s), 2))
}
func (s *Struct) BaseSize() uint16 {
	return *(*uint16)(unsafe.Add(unsafe.Pointer(s), 0))
}
func (s *Struct) A() uint16 {
	return *(*uint16)(unsafe.Add(unsafe.Pointer(s), 0))
}
func (s *Struct) IsGC() bool { return true }

func handleGCCopy(b []byte, h *Handler) {
	switch binary.LittleEndian.Uint16(b[2:]) {
	case v8.ACCOUNT_BALANCE_ADJUSTMENT:
		h.AccountBalanceAdjustment(v8.NewAccountBalanceAdjustmentFrom(b))
	}
}

func handleNoGC(b []byte, h *HandlerNoGC) {
	switch binary.LittleEndian.Uint16(b[2:]) {
	case v8.ACCOUNT_BALANCE_ADJUSTMENT:
		h.AccountBalanceAdjustment(v8.AllocAccountBalanceAdjustmentFrom(b))
	}
}
