package factory

import (
	"github.com/moontrade/dtc-go/model"
	"github.com/moontrade/dtc-go/model/fixed"
)

type Factory interface {
	EncodingRequest() model.EncodingRequestFactory
}

type fixedImpl struct {
}

func (fixedImpl) EncodingRequest() model.EncodingRequestFactory {
	return fixed.EncodingRequestFactoryImpl{}
}
