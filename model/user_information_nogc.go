package model

import (
	"github.com/moontrade/dtc-go/message"
)

type UserInformationPointer struct {
	message.VLSPointer
}

type UserInformationPointerBuilder struct {
	message.VLSPointerBuilder
}

func AllocUserInformation() UserInformationPointerBuilder {
	a := UserInformationPointerBuilder{message.AllocVLSBuilder(14)}
	a.Ptr.SetBytes(0, _UserInformationDefault)
	return a
}

func AllocUserInformationFrom(b []byte) UserInformationPointer {
	return UserInformationPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size       = UserInformationSize  (14)
//     Type       = USER_INFORMATION  (10136)
//     BaseSize   = UserInformationSize  (14)
//     OperatorID = ""
//     LocationID = ""
// }
func (m UserInformationPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _UserInformationDefault)
}

// ToBuilder
func (m UserInformationPointer) ToBuilder() UserInformationPointerBuilder {
	return UserInformationPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *UserInformationPointerBuilder) Finish() UserInformationPointer {
	return UserInformationPointer{m.VLSPointerBuilder.Finish()}
}

// OperatorID
func (m UserInformationPointer) OperatorID() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// OperatorID
func (m UserInformationPointerBuilder) OperatorID() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// LocationID
func (m UserInformationPointer) LocationID() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// LocationID
func (m UserInformationPointerBuilder) LocationID() string {
	return message.StringVLS(m.Ptr, 14, 10)
}

// SetOperatorID
func (m *UserInformationPointerBuilder) SetOperatorID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetLocationID
func (m *UserInformationPointerBuilder) SetLocationID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
}

// Copy
func (m UserInformationPointer) Copy(to UserInformationBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}

// Copy
func (m UserInformationPointerBuilder) Copy(to UserInformationBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}

// CopyPointer
func (m UserInformationPointer) CopyPointer(to UserInformationPointerBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}

// CopyPointer
func (m UserInformationPointerBuilder) CopyPointer(to UserInformationPointerBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}
