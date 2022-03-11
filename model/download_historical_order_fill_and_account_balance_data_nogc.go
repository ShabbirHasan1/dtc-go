package model

import (
	"github.com/moontrade/dtc-go/message"
)

type DownloadHistoricalOrderFillAndAccountBalanceDataPointer struct {
	message.VLSPointer
}

type DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder struct {
	message.VLSPointerBuilder
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
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) Clear() {
	m.Ptr.SetBytes(0, _DownloadHistoricalOrderFillAndAccountBalanceDataDefault)
}

// ToBuilder
func (m DownloadHistoricalOrderFillAndAccountBalanceDataPointer) ToBuilder() DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder {
	return DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder{m.VLSPointer.ToBuilder()}
}

// Finish
func (m *DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) Finish() DownloadHistoricalOrderFillAndAccountBalanceDataPointer {
	return DownloadHistoricalOrderFillAndAccountBalanceDataPointer{m.VLSPointerBuilder.Finish()}
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
func (m *DownloadHistoricalOrderFillAndAccountBalanceDataPointerBuilder) SetTradeAccount(value string) {
	message.SetStringVLSPointer(&m.VLSPointerBuilder, 10, 6, value)
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
