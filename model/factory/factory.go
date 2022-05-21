package factory

import (
	v8 "github.com/moontrade/dtc-go/model/v8"
)

type Factory struct {
	AccountBalanceAdjustmentBuilder AccountBalanceAdjustment
}

//func (f *Factory) ProcessPointer(b []byte) message.Pointer {
//
//}

type AccountBalanceAdjustment struct {
}

func (AccountBalanceAdjustment) New() v8.AccountBalanceAdjustmentBuilder {
	return v8.NewAccountBalanceAdjustment()
}

func (AccountBalanceAdjustment) Alloc() v8.AccountBalanceAdjustmentPointerBuilder {
	return v8.AllocAccountBalanceAdjustment()
}

type AccountBalanceAdjustmentFixed struct {
}

func (AccountBalanceAdjustmentFixed) New() v8.AccountBalanceAdjustmentBuilder {
	return v8.NewAccountBalanceAdjustment()
}

func (AccountBalanceAdjustmentFixed) Alloc() v8.AccountBalanceAdjustmentPointerBuilder {
	return v8.AllocAccountBalanceAdjustment()
}

//func (AccountBalanceAdjustment) AllocFromFixed(f fixed.AccountBalanceAdjustment) model.AccountBalanceAdjustmentPointer {
//	to := vls.AccountBalanceAdjustmentPointerBuilder{VLSPointerBuilder: message.AllocVLSBuilder(40)}
//	to.SetRequestID(f.RequestID())
//	to.SetCurrency(f.Currency())
//	return vls.AccountBalanceAdjustmentPointer{VLSPointer: to.VLSBuilder.Finish()}
//}
