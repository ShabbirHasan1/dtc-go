package factory

import (
	"fmt"
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/model"
	"github.com/moontrade/dtc-go/model/sc"
	"testing"
)

func TestHandler(t *testing.T) {
	h := &Handler{AccountBalanceAdjustment: func(msg model.AccountBalanceAdjustment) {
		if msg.Type() == model.ACCOUNT_BALANCE_REJECT {

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
