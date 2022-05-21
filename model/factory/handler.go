package factory

import (
	"encoding/binary"
	"github.com/moontrade/dtc-go/message/json"
	v8 "github.com/moontrade/dtc-go/model/v8"
	"io"
)

type Handler struct {
	AccountBalanceAdjustment func(msg v8.AccountBalanceAdjustment)
}

type HandlerNoGC struct {
	Unhandled                func(msg interface{}) error
	AccountBalanceAdjustment func(msg v8.AccountBalanceAdjustmentPointer) error
}

func (h *HandlerNoGC) Process(b []byte) {
	if len(b) == 0 {
		return
	}
	switch binary.LittleEndian.Uint16(b[2:]) {
	case v8.ACCOUNT_BALANCE_ADJUSTMENT:
		h.AccountBalanceAdjustment(v8.AllocAccountBalanceAdjustmentFrom(b))
	}
}

func (h *HandlerNoGC) ProcessJson(b []byte) error {
	r, err := json.OpenReader(b)
	if err != nil {
		return err
	}
	switch r.Type {
	case v8.ACCOUNT_BALANCE_ADJUSTMENT:
		m := v8.AllocAccountBalanceAdjustment()
		if err := m.UnmarshalJSONReader(&r); err != nil {
			return err
		}
		if h.AccountBalanceAdjustment == nil {
			return h.AccountBalanceAdjustment(m.Finish())
		} else if h.Unhandled != nil {
			return h.Unhandled(m.Finish())
		}
	}
	return nil
}

func (h *HandlerNoGC) ProcessProtobuf(b []byte) error {
	if len(b) < 4 {
		return io.ErrShortBuffer
	}
	switch binary.LittleEndian.Uint16(b[2:]) {
	case v8.ACCOUNT_BALANCE_ADJUSTMENT:
		m := v8.AllocAccountBalanceAdjustment()
		if err := m.UnmarshalProtobuf(b[4:]); err != nil {
			return err
		}
		if h.AccountBalanceAdjustment == nil {
			return h.AccountBalanceAdjustment(m.Finish())
		} else if h.Unhandled != nil {
			return h.Unhandled(m.Finish())
		}
	}
	return nil
}
