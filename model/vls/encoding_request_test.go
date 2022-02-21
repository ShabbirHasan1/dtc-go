package vls

import "testing"

func TestAllocEncodingRequest(t *testing.T) {
	b := AllocEncodingRequest()
	b.SetEncoding(10)

	b.Finish()
}
