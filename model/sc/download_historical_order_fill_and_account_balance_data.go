package sc

import (
	"github.com/moontrade/dtc-go/message"
)

const DownloadHistoricalOrderFillAndAccountBalanceDataSize = 10

// {
//     Size         = DownloadHistoricalOrderFillAndAccountBalanceDataSize  (10)
//     Type         = DOWNLOAD_HISTORICAL_ORDER_FILL_AND_ACCOUNT_BALANCE_DATA  (10138)
//     BaseSize     = DownloadHistoricalOrderFillAndAccountBalanceDataSize  (10)
//     TradeAccount = ""
// }
var _DownloadHistoricalOrderFillAndAccountBalanceDataDefault = []byte{10, 0, 154, 39, 10, 0, 0, 0, 0, 0}

type DownloadHistoricalOrderFillAndAccountBalanceData struct {
	message.VLS
}

type DownloadHistoricalOrderFillAndAccountBalanceDataBuilder struct {
	message.VLSBuilder
}

type DownloadHistoricalOrderFillAndAccountBalanceDataPointer struct {
	message.VLSPointer
}

type DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder struct {
	message.VLSPointerBuilder
}

func NewDownloadHistoricalOrderFillAndAccountBalanceDataFrom(b []byte) DownloadHistoricalOrderFillAndAccountBalanceData {
	return DownloadHistoricalOrderFillAndAccountBalanceData{VLS: message.NewVLSFrom(b)}
}

func WrapDownloadHistoricalOrderFillAndAccountBalanceData(b []byte) DownloadHistoricalOrderFillAndAccountBalanceData {
	return DownloadHistoricalOrderFillAndAccountBalanceData{VLS: message.WrapVLS(b)}
}

func NewDownloadHistoricalOrderFillAndAccountBalanceData() DownloadHistoricalOrderFillAndAccountBalanceDataBuilder {
	a := DownloadHistoricalOrderFillAndAccountBalanceDataBuilder{message.NewVLSBuilder(10)}
	a.Unsafe().SetBytes(0, _DownloadHistoricalOrderFillAndAccountBalanceDataDefault)
	return a
}

func AllocDownloadHistoricalOrderFillAndAccountBalanceData() DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder {
	a := DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder{message.AllocVLSBuilder(10)}
	a.Ptr.SetBytes(0, _DownloadHistoricalOrderFillAndAccountBalanceDataDefault)
	return a
}

func AllocDownloadHistoricalOrderFillAndAccountBalanceDataFrom(b []byte) DownloadHistoricalOrderFillAndAccountBalanceDataPointer {
	return DownloadHistoricalOrderFillAndAccountBalanceDataPointer{VLSPointer: message.AllocVLSFrom(b)}
}

// Clear
// {
//     Size         = DownloadHistoricalOrderFillAndAccountBalanceDataSize  (10)
//     Type         = DOWNLOAD_HISTORICAL_ORDER_FILL_AND_ACCOUNT_BALANCE_DATA  (10138)
//     BaseSize     = DownloadHistoricalOrderFillAndAccountBalanceDataSize  (10)
//     TradeAccount = ""
// }
func (m DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) Clear() {
	m.Unsafe().SetBytes(0, _DownloadHistoricalOrderFillAndAccountBalanceDataDefault)
}

// Clear
// {
//     Size         = DownloadHistoricalOrderFillAndAccountBalanceDataSize  (10)
//     Type         = DOWNLOAD_HISTORICAL_ORDER_FILL_AND_ACCOUNT_BALANCE_DATA  (10138)
//     BaseSize     = DownloadHistoricalOrderFillAndAccountBalanceDataSize  (10)
//     TradeAccount = ""
// }
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _DownloadHistoricalOrderFillAndAccountBalanceDataDefault)
}

// ToBuilder
func (m DownloadHistoricalOrderFillAndAccountBalanceData) ToBuilder() DownloadHistoricalOrderFillAndAccountBalanceDataBuilder {
	return DownloadHistoricalOrderFillAndAccountBalanceDataBuilder{m.VLS.ToBuilder()}
}

// ToBuilder
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointer) ToBuilder() DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder {
	return DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) Finish() DownloadHistoricalOrderFillAndAccountBalanceData {
	return DownloadHistoricalOrderFillAndAccountBalanceData{m.VLSBuilder.Finish()}
}

// Finish
func (m *DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) Finish() DownloadHistoricalOrderFillAndAccountBalanceDataPointer {
	return DownloadHistoricalOrderFillAndAccountBalanceDataPointer{m.VLSPointerBuilder.Finish()}
}

// TradeAccount
func (m DownloadHistoricalOrderFillAndAccountBalanceData) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) TradeAccount() string {
	return message.StringVLS(m.Unsafe(), 10, 6)
}

// TradeAccount
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointer) TradeAccount() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// TradeAccount
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) TradeAccount() string {
	return message.StringVLS(m.Ptr, 10, 6)
}

// SetTradeAccount
func (m *DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) SetTradeAccount(value string) {
	message.SetStringVLS(&m.VLSBuilder, 10, 6, value)
}

// SetTradeAccount
func (m *DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
}

// Copy
func (m DownloadHistoricalOrderFillAndAccountBalanceData) Copy(to DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) Copy(to DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m DownloadHistoricalOrderFillAndAccountBalanceData) CopyPointer(to DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) CopyPointer(to DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointer) Copy(to DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}

// Copy
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) Copy(to DownloadHistoricalOrderFillAndAccountBalanceDataBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointer) CopyPointer(to DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}

// CopyPointer
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) CopyPointer(to DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) {
	to.SetTradeAccount(m.TradeAccount())
}
