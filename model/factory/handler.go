package factory

import "github.com/moontrade/dtc-go/model"

type Handler struct {
	AccountBalanceAdjustment func(msg model.AccountBalanceAdjustment)
}

type HandlerNoGC struct {
	AccountBalanceAdjustment func(msg model.AccountBalanceAdjustmentPointer)
}
