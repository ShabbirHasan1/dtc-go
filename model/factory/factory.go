package factory

import (
	"github.com/moontrade/dtc-go/message"
	"github.com/moontrade/dtc-go/model"
)

type Factory struct {
	AccountBalanceAdjustmentBuilder AccountBalanceAdjustment
}

//func (f *Factory) ProcessPointer(b []byte) message.Pointer {
//
//}

type AccountBalanceAdjustment struct {
}

func (AccountBalanceAdjustment) New() model.AccountBalanceAdjustmentBuilder {
	return model.NewAccountBalanceAdjustment()
}

func (AccountBalanceAdjustment) Alloc() model.AccountBalanceAdjustmentPointerBuilder {
	return model.AllocAccountBalanceAdjustment()
}

type AccountBalanceAdjustmentFixed struct {
}

func (AccountBalanceAdjustmentFixed) New() model.AccountBalanceAdjustmentBuilder {
	return model.NewAccountBalanceAdjustment()
}

func (AccountBalanceAdjustmentFixed) Alloc() model.AccountBalanceAdjustmentPointerBuilder {
	return model.AllocAccountBalanceAdjustment()
}

func (AccountBalanceAdjustmentFixed) AllocFromBytes(b []byte) model.AccountBalanceAdjustmentPointerBuilder {
	to := model.AllocAccountBalanceAdjustment()
	model.AccountBalanceAdjustmentFixed{Fixed: message.WrapFixed(b)}.CopyToPointer(to)
	return to
}

//func (AccountBalanceAdjustment) AllocFromFixed(f fixed.AccountBalanceAdjustment) model.AccountBalanceAdjustmentPointer {
//	to := vls.AccountBalanceAdjustmentPointerBuilder{VLSPointerBuilder: message.AllocVLSBuilder(40)}
//	to.SetRequestID(f.RequestID())
//	to.SetCurrency(f.Currency())
//	return vls.AccountBalanceAdjustmentPointer{VLSPointer: to.VLSBuilder.Finish()}
//}
