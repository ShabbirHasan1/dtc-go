package factory

import (
	"github.com/moontrade/dtc-go/model/fixed"
	"github.com/moontrade/dtc-go/model/types"
)

type Factory interface {
	EncodingRequest() types.EncodingRequestFactory
}

type fixedImpl struct {
}

func (fixedImpl) EncodingRequest() types.EncodingRequestFactory {
	return fixed.EncodingRequestFactory{}
}
