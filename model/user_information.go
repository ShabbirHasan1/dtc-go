package model

import (
	"github.com/moontrade/dtc-go/message"
)

// {
//     Size       = UserInformationSize  (14)
//     Type       = USER_INFORMATION  (10136)
//     BaseSize   = UserInformationSize  (14)
//     OperatorID = ""
//     LocationID = ""
// }
var _UserInformationDefault = []byte{14, 0, 152, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

const UserInformationSize = 14

type UserInformation struct {
	message.VLS
}

type UserInformationBuilder struct {
	message.VLSBuilder
}

func NewUserInformationFrom(b []byte) UserInformation {
	return UserInformation{VLS: message.NewVLSFrom(b)}
}

func WrapUserInformation(b []byte) UserInformation {
	return UserInformation{VLS: message.WrapVLS(b)}
}

func NewUserInformation() UserInformationBuilder {
	a := UserInformationBuilder{message.NewVLSBuilder(14)}
	a.Unsafe().SetBytes(0, _UserInformationDefault)
	return a
}

// Clear
// {
//     Size       = UserInformationSize  (14)
//     Type       = USER_INFORMATION  (10136)
//     BaseSize   = UserInformationSize  (14)
//     OperatorID = ""
//     LocationID = ""
// }
func (m UserInformationBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UserInformationDefault)
}

// ToBuilder
func (m UserInformation) ToBuilder() UserInformationBuilder {
	return UserInformationBuilder{m.VLS.ToBuilder()}
}

// Finish
func (m UserInformationBuilder) Finish() UserInformation {
	return UserInformation{m.VLSBuilder.Finish()}
}

// OperatorID
func (m UserInformation) OperatorID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// OperatorID
func (m UserInformationBuilder) OperatorID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// LocationID
func (m UserInformation) LocationID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// LocationID
func (m UserInformationBuilder) LocationID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// SetOperatorID
func (m *UserInformationBuilder) SetOperatorID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetLocationID
func (m *UserInformationBuilder) SetLocationID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// Copy
func (m UserInformation) Copy(to UserInformationBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}

// Copy
func (m UserInformationBuilder) Copy(to UserInformationBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}

// CopyPointer
func (m UserInformation) CopyPointer(to UserInformationPointerBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}

// CopyPointer
func (m UserInformationBuilder) CopyPointer(to UserInformationPointerBuilder) {
	to.SetOperatorID(m.OperatorID())
	to.SetLocationID(m.LocationID())
}
