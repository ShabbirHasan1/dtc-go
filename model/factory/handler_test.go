package factory

import (
	"github.com/moontrade/dtc-go/model"
	"testing"
)

func TestHandler(t *testing.T) {
	h := &Handler{AccountBalanceAdjustment: func(msg model.AccountBalanceAdjustment) {
		if msg.Type() == model.ACCOUNT_BALANCE_REJECT {

		}
	}}

	_ = h

	f := &Factory{}
	f.AccountBalanceAdjustmentBuilder.Alloc()
}
