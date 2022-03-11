package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const UserInformationSize = 14

// {
//     Size       = UserInformationSize  (14)
//     Type       = USER_INFORMATION  (10136)
//     BaseSize   = UserInformationSize  (14)
//     OperatorID = ""
//     LocationID = ""
// }
var _UserInformationDefault = []byte{14, 0, 152, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0}

type UserInformation struct {
	message.VLS
}

type UserInformationBuilder struct {
	message.VLSBuilder
}

type UserInformationPointer struct {
	message.VLSPointer
}

type UserInformationPointerBuilder struct {
	message.VLSPointerBuilder
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
func (m UserInformationBuilder) Clear() {
	m.Unsafe().SetBytes(0, _UserInformationDefault)
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
func (m UserInformation) ToBuilder() UserInformationBuilder {
	return UserInformationBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m UserInformationPointer) ToBuilder() UserInformationPointerBuilder {
	return UserInformationPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m UserInformationBuilder) Finish() UserInformation {
	return UserInformation{m.VLSBuilder.Finish()}
}

// Finish
func (m *UserInformationPointerBuilder) Finish() UserInformationPointer {
	return UserInformationPointer{m.VLSPointerBuilder.Finish()}
}

// OperatorID
func (m UserInformation) OperatorID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// OperatorID
func (m UserInformationBuilder) OperatorID() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
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
func (m UserInformation) LocationID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
}

// LocationID
func (m UserInformationBuilder) LocationID() string {
	return message.StringVLS(m.Unsafe(), 14, 10)
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
func (m *UserInformationBuilder) SetOperatorID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetOperatorID
func (m *UserInformationPointerBuilder) SetOperatorID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// SetLocationID
func (m *UserInformationBuilder) SetLocationID(value string) {
	message.SetStringVLS(&m.VLSBuilder, 14, 10, value)
}

// SetLocationID
func (m *UserInformationPointerBuilder) SetLocationID(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 14, 10, value)
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