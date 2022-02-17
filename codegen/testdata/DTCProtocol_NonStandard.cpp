#ifndef _CRT_SECURE_NO_WARNINGS
#define  _CRT_SECURE_NO_WARNINGS
#endif

#include <float.h>
#include <limits.h>
#include <string.h>
#include <memory.h> 
#include <stddef.h>

#include "DTCProtocol_NonStandard.h"


#ifndef max
#define max(a,b)            (((a) > (b)) ? (a) : (b))
#endif

#ifndef min
#define min(a,b)            (((a) < (b)) ? (a) : (b))
#endif


/****************************************************************************/
// s_EncodingRequestExtended

/*==========================================================================*/
uint16_t DTC::s_EncodingRequestExtended::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
void DTC::s_EncodingRequestExtended::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_EncodingRequestExtended), *static_cast<uint16_t*>(p_SourceData)));
}
/*==========================================================================*/
int32_t DTC::s_EncodingRequestExtended::GetProtocolVersion() const
{
	if (Size < offsetof(s_EncodingRequestExtended, ProtocolVersion) + sizeof(ProtocolVersion))
		return 0;

	return ProtocolVersion;
}
/*==========================================================================*/
DTC::EncodingEnum DTC::s_EncodingRequestExtended::GetEncoding() const
{
	if (Size < offsetof(s_EncodingRequestExtended, Encoding) + sizeof(Encoding))
		return DTC::BINARY_ENCODING;

	return Encoding;
}
/*==========================================================================*/
const char* DTC::s_EncodingRequestExtended::GetProtocolType()
{
	if (Size < offsetof(s_EncodingRequestExtended, ProtocolType) + sizeof(ProtocolType))
		return "";

	ProtocolType[sizeof(ProtocolType) - 1] = '\0';  // Ensure that the null terminator exists

	return ProtocolType;
}

/*==========================================================================*/
void DTC::s_EncodingRequestExtended::SetProtocolType(const char* NewValue)
{
	//Do not use the secure version of this function. This version of the function will set the remaining bytes in the destination after the null terminator to nulls. The secure version does not do this.
	strncpy(ProtocolType, NewValue, sizeof(ProtocolType) - 1);
}
/*==========================================================================*/
n_DTCNonStandard::UseCompressionEnum DTC::s_EncodingRequestExtended::GetUseCompression() const
{
	if (Size < offsetof(s_EncodingRequestExtended, UseCompression) + sizeof(UseCompression))
		return n_DTCNonStandard::USE_COMPRESSION_DISABLED;

	return UseCompression;
}


/****************************************************************************/
// s_HeartbeatExtended

/*==========================================================================*/
uint16_t DTC::s_HeartbeatExtended::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_HeartbeatExtended::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_HeartbeatExtended), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_HeartbeatExtended::GetNumDroppedMessages() const
{
	if (Size < offsetof(s_HeartbeatExtended, NumDroppedMessages) + sizeof(NumDroppedMessages))
		return 0;

	return NumDroppedMessages;
}

/*==========================================================================*/
DTC::t_DateTime DTC::s_HeartbeatExtended::GetCurrentDateTime() const
{
	if (Size < offsetof(s_HeartbeatExtended, CurrentDateTime) + sizeof(CurrentDateTime))
		return 0;

	return CurrentDateTime;
}

/*==========================================================================*/
uint16_t DTC::s_HeartbeatExtended::GetSecondsSinceLastReceivedHeartbeat() const
{
	if (Size < offsetof(s_HeartbeatExtended, SecondsSinceLastReceivedHeartbeat) + sizeof(SecondsSinceLastReceivedHeartbeat))
		return 0;

	return SecondsSinceLastReceivedHeartbeat;
}

/*==========================================================================*/
uint16_t DTC::s_HeartbeatExtended::GetNumberOfOutstandingSentBuffers() const
{
	if (Size < offsetof(s_HeartbeatExtended, NumberOfOutstandingSentBuffers) + sizeof(NumberOfOutstandingSentBuffers))
		return 0;

	return NumberOfOutstandingSentBuffers;
}

/*==========================================================================*/
uint16_t DTC::s_HeartbeatExtended::GetPendingTransmissionDelayInMilliseconds() const
{
	if (Size < offsetof(s_HeartbeatExtended, PendingTransmissionDelayInMilliseconds) + sizeof(PendingTransmissionDelayInMilliseconds))
		return 0;

	return PendingTransmissionDelayInMilliseconds;
}

/*==========================================================================*/
uint32_t DTC::s_HeartbeatExtended::GetCurrentSendBufferSizeInBytes() const
{
	if (Size < offsetof(s_HeartbeatExtended, CurrentSendBufferSizeInBytes) + sizeof(CurrentSendBufferSizeInBytes))
		return 0;

	return CurrentSendBufferSizeInBytes;
}

/*==========================================================================*/
DTC::t_DateTimeWithMicrosecondsInt DTC::s_HeartbeatExtended::GetSendingDateTimeMicroseconds() const
{
	if (Size < offsetof(s_HeartbeatExtended, SendingDateTimeMicroseconds) + sizeof(SendingDateTimeMicroseconds))
		return 0;

	return SendingDateTimeMicroseconds;

}

/*==========================================================================*/
float DTC::s_HeartbeatExtended::GetDataCompressionRatio() const
{
	if (Size < offsetof(s_HeartbeatExtended, DataCompressionRatio) + sizeof(DataCompressionRatio))
		return 0;

	return DataCompressionRatio;
}

/*==========================================================================*/
uint64_t DTC::s_HeartbeatExtended::GetTotalUncompressedBytes() const
{
	if (Size < offsetof(s_HeartbeatExtended, TotalUncompressedBytes) + sizeof(TotalUncompressedBytes))
		return 0;

	return TotalUncompressedBytes;
}

/*==========================================================================*/
double DTC::s_HeartbeatExtended::GetTotalCompressionTime() const
{
	if (Size < offsetof(s_HeartbeatExtended, TotalCompressionTime) + sizeof(TotalCompressionTime))
		return 0;

	return TotalCompressionTime;
}

/*==========================================================================*/
uint32_t DTC::s_HeartbeatExtended::GetNumberOfCompressions() const
{
	if (Size < offsetof(s_HeartbeatExtended, NumberOfCompressions) + sizeof(NumberOfCompressions))
		return 0;

	return NumberOfCompressions;
}

/*==========================================================================*/
uint32_t DTC::s_HeartbeatExtended::GetSourceIPAddress() const
{
	if (Size < offsetof(s_HeartbeatExtended, SourceIPAddress) + sizeof(SourceIPAddress))
		return 0;

	return SourceIPAddress;
}

/****************************************************************************/
// s_HistoricalTradesRequest

/*==========================================================================*/
uint16_t DTC::s_HistoricalTradesRequest::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesRequest::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_HistoricalTradesRequest), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
int32_t DTC::s_HistoricalTradesRequest::GetRequestID() const
{
	if (Size < offsetof(s_HistoricalTradesRequest, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
const char* DTC::s_HistoricalTradesRequest::GetSymbol()
{
	if (Size < offsetof(s_HistoricalTradesRequest, Symbol) + sizeof(Symbol))
		return "";

	Symbol[sizeof(Symbol) - 1] = '\0';  // Ensure that the null terminator exists

	return Symbol;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesRequest::SetSymbol(const char* NewValue)
{
	strncpy(Symbol, NewValue, sizeof(Symbol) - 1);
}

/*==========================================================================*/
const char* DTC::s_HistoricalTradesRequest::GetTradeAccount()
{
	if (Size < offsetof(s_HistoricalTradesRequest, TradeAccount) + sizeof(TradeAccount))
		return "";

	TradeAccount[sizeof(TradeAccount) - 1] = '\0';  // Ensure that the null terminator exists

	return TradeAccount;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesRequest::SetTradeAccount(const char* NewValue)
{
	strncpy(TradeAccount, NewValue, sizeof(TradeAccount) - 1);
}

/*==========================================================================*/
DTC::t_DateTime DTC::s_HistoricalTradesRequest::GetStartDateTime() const
{
	if (Size < offsetof(s_HistoricalTradesRequest, StartDateTime) + sizeof(StartDateTime))
		return 0;

	return StartDateTime;
}

/*==========================================================================*/
uint32_t DTC::s_HistoricalTradesRequest::GetSubAccountIdentifier() const
{
	if (Size < offsetof(s_HistoricalTradesRequest, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/*==========================================================================*/
uint8_t DTC::s_HistoricalTradesRequest::GetCreateFlatToFlatTrades() const
{
	if (Size < offsetof(s_HistoricalTradesRequest, CreateFlatToFlatTrades) + sizeof(CreateFlatToFlatTrades))
		return 0;

	return CreateFlatToFlatTrades;
}

/****************************************************************************/
// s_HistoricalTradesReject

/*==========================================================================*/
uint16_t DTC::s_HistoricalTradesReject::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesReject::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_HistoricalTradesReject), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_HistoricalTradesReject::GetRequestID() const
{
	if (Size < offsetof(s_HistoricalTradesReject, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
const char* DTC::s_HistoricalTradesReject::GetRejectText()
{
	if (Size < offsetof(s_HistoricalTradesReject, RejectText) + sizeof(RejectText))
		return "";

	RejectText[sizeof(RejectText) - 1] = '\0';  // Ensure that the null terminator exists

	return RejectText;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesReject::SetRejectText(const char* NewValue)
{
	strncpy(RejectText, NewValue, sizeof(RejectText) - 1);
}

/****************************************************************************/
// s_HistoricalTradesResponse

/*==========================================================================*/
uint16_t DTC::s_HistoricalTradesResponse::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesResponse::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_HistoricalTradesResponse), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
int32_t DTC::s_HistoricalTradesResponse::GetRequestID() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC::s_HistoricalTradesResponse::GetIsFinalMessage() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, IsFinalMessage) + sizeof(IsFinalMessage))
		return 0;

	return IsFinalMessage;
}

/*==========================================================================*/
const char* DTC::s_HistoricalTradesResponse::GetSymbol()
{
	if (Size < offsetof(s_HistoricalTradesResponse, Symbol) + sizeof(Symbol))
		return "";

	Symbol[sizeof(Symbol) - 1] = '\0';  // Ensure that the null terminator exists

	return Symbol;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesResponse::SetSymbol(const char* NewValue)
{
	strncpy(Symbol, NewValue, sizeof(Symbol) - 1);
}

/*==========================================================================*/
const char* DTC::s_HistoricalTradesResponse::GetTradeAccount()
{
	if (Size < offsetof(s_HistoricalTradesResponse, TradeAccount) + sizeof(TradeAccount))
		return "";

	TradeAccount[sizeof(TradeAccount) - 1] = '\0';  // Ensure that the null terminator exists

	return TradeAccount;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesResponse::SetTradeAccount(const char* NewValue)
{
	strncpy(TradeAccount, NewValue, sizeof(TradeAccount) - 1);
}

/*==========================================================================*/
DTC::t_DateTimeWithMilliseconds DTC::s_HistoricalTradesResponse::GetEntryDateTime() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, EntryDateTime) + sizeof(EntryDateTime))
		return 0.0;

	return EntryDateTime;
}

/*==========================================================================*/
DTC::t_DateTimeWithMilliseconds DTC::s_HistoricalTradesResponse::GetExitDateTime() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, ExitDateTime) + sizeof(ExitDateTime))
		return 0.0;

	return ExitDateTime;
}

/*==========================================================================*/
double DTC::s_HistoricalTradesResponse::GetEntryPrice() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, EntryPrice) + sizeof(EntryPrice))
		return 0.0;

	return EntryPrice;
}

/*==========================================================================*/
double DTC::s_HistoricalTradesResponse::GetExitPrice() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, ExitPrice) + sizeof(ExitPrice))
		return 0.0;

	return ExitPrice;
}

/*==========================================================================*/
DTC::BuySellEnum DTC::s_HistoricalTradesResponse::GetTradeType() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, TradeType) + sizeof(TradeType))
		return DTC::BUY_SELL_UNSET;

	return TradeType;
}

/*==========================================================================*/
uint32_t DTC::s_HistoricalTradesResponse::GetEntryQuantity() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, EntryQuantity) + sizeof(EntryQuantity))
		return 0;

	return EntryQuantity;
}

/*==========================================================================*/
uint32_t DTC::s_HistoricalTradesResponse::GetExitQuantity() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, ExitQuantity) + sizeof(ExitQuantity))
		return 0;

	return ExitQuantity;
}

/*==========================================================================*/
uint32_t DTC::s_HistoricalTradesResponse::GetMaxOpenQuantity() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, MaxOpenQuantity) + sizeof(MaxOpenQuantity))
		return 0;

	return MaxOpenQuantity;
}

/*==========================================================================*/
double DTC::s_HistoricalTradesResponse::GetClosedProfitLoss() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, ClosedProfitLoss) + sizeof(ClosedProfitLoss))
		return 0.0;

	return ClosedProfitLoss;
}

/*==========================================================================*/
double DTC::s_HistoricalTradesResponse::GetMaximumOpenPositionLoss() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, MaximumOpenPositionLoss) + sizeof(MaximumOpenPositionLoss))
		return 0.0;

	return MaximumOpenPositionLoss;
}

/*==========================================================================*/
double DTC::s_HistoricalTradesResponse::GetMaximumOpenPositionProfit() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, MaximumOpenPositionProfit) + sizeof(MaximumOpenPositionProfit))
		return 0.0;

	return MaximumOpenPositionProfit;
}

/*==========================================================================*/
double DTC::s_HistoricalTradesResponse::GetCommission() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, Commission) + sizeof(Commission))
		return 0.0;

	return Commission;
}

/*==========================================================================*/
const char* DTC::s_HistoricalTradesResponse::GetOpenFillExecutionID()
{
	if (Size < offsetof(s_HistoricalTradesResponse, OpenFillExecutionID) + sizeof(OpenFillExecutionID))
		return "";

	OpenFillExecutionID[sizeof(OpenFillExecutionID) - 1] = '\0';  // Ensure that the null terminator exists

	return OpenFillExecutionID;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesResponse::SetOpenFillExecutionID(const char* NewValue)
{
	strncpy(OpenFillExecutionID, NewValue, sizeof(OpenFillExecutionID) - 1);
}

/*==========================================================================*/
const char* DTC::s_HistoricalTradesResponse::GetCloseFillExecutionID()
{
	if (Size < offsetof(s_HistoricalTradesResponse, CloseFillExecutionID) + sizeof(CloseFillExecutionID))
		return "";

	CloseFillExecutionID[sizeof(CloseFillExecutionID) - 1] = '\0';  // Ensure that the null terminator exists

	return CloseFillExecutionID;
}

/*==========================================================================*/
void DTC::s_HistoricalTradesResponse::SetCloseFillExecutionID(const char* NewValue)
{
	strncpy(CloseFillExecutionID, NewValue, sizeof(CloseFillExecutionID) - 1);
}

/*==========================================================================*/
uint32_t DTC::s_HistoricalTradesResponse::GetSubAccountIdentifier() const
{
	if (Size < offsetof(s_HistoricalTradesResponse, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/****************************************************************************/
// s_ReplayChartData

/*==========================================================================*/
uint16_t DTC::s_ReplayChartData::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_ReplayChartData::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_ReplayChartData), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_ReplayChartData::GetRequestID() const
{
	if (Size < offsetof(s_ReplayChartData, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
const char* DTC::s_ReplayChartData::GetSymbol()
{
	if (Size < offsetof(s_ReplayChartData, Symbol) + sizeof(Symbol))
		return "";

	Symbol[sizeof(Symbol) - 1] = '\0';  // Ensure that the null terminator exists

	return Symbol;
}

/*==========================================================================*/
void DTC::s_ReplayChartData::SetSymbol(const char* NewValue)
{
	strncpy(Symbol, NewValue, sizeof(Symbol) - 1);
}

/*==========================================================================*/
const char* DTC::s_ReplayChartData::GetTradeAccount()
{
	if (Size < offsetof(s_ReplayChartData, TradeAccount) + sizeof(TradeAccount))
		return "";

	TradeAccount[sizeof(TradeAccount) - 1] = '\0';  // Ensure that the null terminator exists

	return TradeAccount;
}

/*==========================================================================*/
void DTC::s_ReplayChartData::SetTradeAccount(const char* NewValue)
{
	strncpy(TradeAccount, NewValue, sizeof(TradeAccount) - 1);
}

/*==========================================================================*/
const char* DTC::s_ReplayChartData::GetTimeZone()
{
	if (Size < offsetof(s_ReplayChartData, TimeZone) + sizeof(TimeZone))
		return "";

	TimeZone[sizeof(TimeZone) - 1] = '\0';  // Ensure that the null terminator exists

	return TimeZone;
}

/*==========================================================================*/
void DTC::s_ReplayChartData::SetTimeZone(const char* NewValue)
{
	strncpy(TimeZone, NewValue, sizeof(TimeZone) - 1);
}

/*==========================================================================*/
DTC::t_DateTimeWithMillisecondsInt DTC::s_ReplayChartData::GetStartDateTimeForInitialData() const
{
	if (Size < offsetof(s_ReplayChartData, StartDateTimeForInitialData) + sizeof(StartDateTimeForInitialData))
		return 0;

	return StartDateTimeForInitialData;
}

/*==========================================================================*/
DTC::t_DateTimeWithMillisecondsInt DTC::s_ReplayChartData::GetStartDateTime() const
{
	if (Size < offsetof(s_ReplayChartData, StartDateTime) + sizeof(StartDateTime))
		return 0;

	return StartDateTime;
}

/*==========================================================================*/
DTC::t_DateTimeWithMillisecondsInt DTC::s_ReplayChartData::GetStopDateTime() const
{
	if (Size < offsetof(s_ReplayChartData, StopDateTime) + sizeof(StopDateTime))
		return 0;

	return StopDateTime;
}

/*==========================================================================*/
uint16_t DTC::s_ReplayChartData::GetSessionBeginTimeInSeconds() const
{
	if (Size < offsetof(s_ReplayChartData, SessionBeginTimeInSeconds) + sizeof(SessionBeginTimeInSeconds))
		return 0;

	return SessionBeginTimeInSeconds;
}

/*==========================================================================*/
uint16_t DTC::s_ReplayChartData::GetSessionEndTimeInSeconds() const
{
	if (Size < offsetof(s_ReplayChartData, SessionEndTimeInSeconds) + sizeof(SessionEndTimeInSeconds))
		return 0;

	return SessionEndTimeInSeconds;
}

/*==========================================================================*/
float DTC::s_ReplayChartData::GetReplaySpeed() const
{
	if (Size < offsetof(s_ReplayChartData, ReplaySpeed) + sizeof(ReplaySpeed))
		return 0.0;

	return ReplaySpeed;
}

/*==========================================================================*/
int32_t DTC::s_ReplayChartData::GetBarTimeInSeconds() const
{
	if (Size < offsetof(s_ReplayChartData, BarTimeInSeconds) + sizeof(BarTimeInSeconds))
		return 0;

	return BarTimeInSeconds;
}

/*==========================================================================*/
uint8_t DTC::s_ReplayChartData::GetPauseReplayAfterInitialDataSent() const
{
	if (Size < offsetof(s_ReplayChartData, PauseReplayAfterInitialDataSent) + sizeof(PauseReplayAfterInitialDataSent))
		return 0;

	return PauseReplayAfterInitialDataSent;
}

/*==========================================================================*/
uint8_t DTC::s_ReplayChartData::GetUseSavedPriorState() const
{
	if (Size < offsetof(s_ReplayChartData, UseSavedPriorState) + sizeof(UseSavedPriorState))
		return 0;

	return UseSavedPriorState;
}

/*==========================================================================*/
float DTC::s_ReplayChartData::GetSymbolVolatility() const
{
	if (Size < offsetof(s_ReplayChartData, SymbolVolatility) + sizeof(SymbolVolatility))
		return 0.0;

	return SymbolVolatility;
}

/*==========================================================================*/
float DTC::s_ReplayChartData::GetInterestRate() const
{
	if (Size < offsetof(s_ReplayChartData, InterestRate) + sizeof(InterestRate))
		return 0.0;

	return InterestRate;
}

/*==========================================================================*/
int32_t DTC::s_ReplayChartData::GetNumberOfOrdersToTriggerFinishToStopDateTime() const
{
	if (Size < offsetof(s_ReplayChartData, NumberOfOrdersToTriggerFinishToStopDateTime) + sizeof(NumberOfOrdersToTriggerFinishToStopDateTime))
		return 0;

	return NumberOfOrdersToTriggerFinishToStopDateTime;
}

/*==========================================================================*/
int32_t DTC::s_ReplayChartData::GetMaximumNumberOfOrdersPerReplaySession() const
{
	if (Size < offsetof(s_ReplayChartData, MaximumNumberOfOrdersPerReplaySession) + sizeof(MaximumNumberOfOrdersPerReplaySession))
		return 0;

	return MaximumNumberOfOrdersPerReplaySession;
}

/*==========================================================================*/
int32_t DTC::s_ReplayChartData::GetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime() const
{
	if (Size < offsetof(s_ReplayChartData, NumberOfDaysForInitialDataFromBeforeLastSavedDateTime) + sizeof(NumberOfDaysForInitialDataFromBeforeLastSavedDateTime))
		return 0;

	return NumberOfDaysForInitialDataFromBeforeLastSavedDateTime;
}

/*==========================================================================*/
uint32_t DTC::s_ReplayChartData::GetSubAccountIdentifier() const
{
	if (Size < offsetof(s_ReplayChartData, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/*==========================================================================*/
int32_t DTC::s_ReplayChartData::GetOptionsPriceSendIntervalInSeconds() const
{
	if (Size < offsetof(s_ReplayChartData, OptionsPriceSendIntervalInSeconds) + sizeof(OptionsPriceSendIntervalInSeconds))
		return 0;

	return OptionsPriceSendIntervalInSeconds;
}

/****************************************************************************/
// s_ReplayChartDataPerformAction

/*==========================================================================*/
uint16_t DTC::s_ReplayChartDataPerformAction::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_ReplayChartDataPerformAction::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_ReplayChartDataPerformAction), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_ReplayChartDataPerformAction::GetRequestID() const
{
	if (Size < offsetof(s_ReplayChartDataPerformAction, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
n_DTCNonStandard::ReplayChartDataActionEnum DTC::s_ReplayChartDataPerformAction::GetAction() const
{
	if (Size < offsetof(s_ReplayChartDataPerformAction, Action) + sizeof(Action))
		return n_DTCNonStandard::REPLAY_CHART_DATA_ACTION_NONE;

	return Action;
}

/*==========================================================================*/
float DTC::s_ReplayChartDataPerformAction::GetReplaySpeed() const
{
	if (Size < offsetof(s_ReplayChartDataPerformAction, ReplaySpeed) + sizeof(ReplaySpeed))
		return 0.0;

	return ReplaySpeed;
}

/****************************************************************************/
// s_ReplayChartDataStatus

/*==========================================================================*/
uint16_t DTC::s_ReplayChartDataStatus::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_ReplayChartDataStatus::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_ReplayChartDataStatus), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_ReplayChartDataStatus::GetRequestID() const
{
	if (Size < offsetof(s_ReplayChartDataStatus, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
const char* DTC::s_ReplayChartDataStatus::GetErrorMessage()
{
	if (Size < offsetof(s_ReplayChartDataStatus, ErrorMessage) + sizeof(ErrorMessage))
		return "";

	ErrorMessage[sizeof(ErrorMessage) - 1] = '\0';  // Ensure that the null terminator exists

	return ErrorMessage;
}

/*==========================================================================*/
void DTC::s_ReplayChartDataStatus::SetErrorMessage(const char* NewValue)
{
	strncpy(ErrorMessage, NewValue, sizeof(ErrorMessage) - 1);
}

/*==========================================================================*/
n_DTCNonStandard::ReplayChartDataStatusEnum DTC::s_ReplayChartDataStatus::GetStatus() const
{
	if (Size < offsetof(s_ReplayChartDataStatus, Status) + sizeof(Status))
		return n_DTCNonStandard::REPLAY_CHART_DATA_STATUS_UNSET;

	return Status;
}

/*==========================================================================*/
uint32_t DTC::s_ReplayChartDataStatus::GetSubAccountIdentifier() const
{
	if (Size < offsetof(s_ReplayChartDataStatus, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/****************************************************************************/
// s_RequestNumCurrentClientConnections

/*==========================================================================*/
uint16_t DTC::s_RequestNumCurrentClientConnections::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_RequestNumCurrentClientConnections::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_RequestNumCurrentClientConnections), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_RequestNumCurrentClientConnections::GetRequestID() const
{
	if (Size < offsetof(s_RequestNumCurrentClientConnections, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
const char* DTC::s_RequestNumCurrentClientConnections::GetUsername()
{
	if (Size < offsetof(s_RequestNumCurrentClientConnections, Username) + sizeof(Username))
		return "";

	Username[sizeof(Username) - 1] = '\0';  // Ensure that the null terminator exists

	return Username;
}

/*==========================================================================*/
void DTC::s_RequestNumCurrentClientConnections::SetUsername(const char* NewValue)
{
	strncpy(Username, NewValue, sizeof(Username) - 1);
}

/*==========================================================================*/
int64_t DTC::s_RequestNumCurrentClientConnections::GetDeviceIdentifier() const
{
	if (Size < offsetof(s_RequestNumCurrentClientConnections, DeviceIdentifier) + sizeof(DeviceIdentifier))
		return 0;

	return DeviceIdentifier;
}

/****************************************************************************/
// s_NumCurrentClientConnectionsResponse

/*==========================================================================*/
uint16_t DTC::s_NumCurrentClientConnectionsResponse::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_NumCurrentClientConnectionsResponse::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_NumCurrentClientConnectionsResponse), *static_cast<uint16_t*>( p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_NumCurrentClientConnectionsResponse::GetRequestID() const
{
	if (Size < offsetof(s_NumCurrentClientConnectionsResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
const char* DTC::s_NumCurrentClientConnectionsResponse::GetUsername()
{
	if (Size < offsetof(s_NumCurrentClientConnectionsResponse, Username) + sizeof(Username))
		return "";

	Username[sizeof(Username) - 1] = '\0';  // Ensure that the null terminator exists

	return Username;
}

/*==========================================================================*/
void DTC::s_NumCurrentClientConnectionsResponse::SetUsername(const char* NewValue)
{
	strncpy(Username, NewValue, sizeof(Username) - 1);
}

/*==========================================================================*/
int32_t DTC::s_NumCurrentClientConnectionsResponse::GetNumConnectionsForDifferentDevices() const
{
	if (Size < offsetof(s_NumCurrentClientConnectionsResponse, NumConnectionsForDifferentDevices) + sizeof(NumConnectionsForDifferentDevices))
		return 0;

	return NumConnectionsForDifferentDevices;
}

/*==========================================================================*/
int32_t DTC::s_NumCurrentClientConnectionsResponse::GetNumConnectionsForSameDevice() const
{
	if (Size < offsetof(s_NumCurrentClientConnectionsResponse, NumConnectionsForSameDevice) + sizeof(NumConnectionsForSameDevice))
		return 0;

	return NumConnectionsForSameDevice;
}


/****************************************************************************/
// s_ClientDeviceUpdate

/*==========================================================================*/
uint16_t DTC::s_ClientDeviceUpdate::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_ClientDeviceUpdate::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_ClientDeviceUpdate), *static_cast<uint16_t*>(p_SourceData)));
}

/*==========================================================================*/
int64_t DTC::s_ClientDeviceUpdate::GetClientDeviceIdentifier() const
{
	if (Size < offsetof(s_ClientDeviceUpdate, ClientDeviceIdentifier) + sizeof(ClientDeviceIdentifier))
		return 0;

	return ClientDeviceIdentifier;
}

/****************************************************************************/
// s_InterprocessSynchronizationRemoteState

/*==========================================================================*/
uint16_t DTC::s_InterprocessSynchronizationRemoteState::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_InterprocessSynchronizationRemoteState::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_InterprocessSynchronizationRemoteState), *static_cast<uint16_t*>(p_SourceData)));
}

/*==========================================================================*/
uint8_t DTC::s_InterprocessSynchronizationRemoteState::GetIsPrimary() const
{
	if (Size < offsetof(s_InterprocessSynchronizationRemoteState, IsPrimary) + sizeof(IsPrimary))
		return 0;

	return IsPrimary;
}

/*==========================================================================*/
uint8_t DTC::s_InterprocessSynchronizationRemoteState::GetIsSecondary() const
{
	if (Size < offsetof(s_InterprocessSynchronizationRemoteState, IsSecondary) + sizeof(IsSecondary))
		return 0;

	return IsSecondary;
}

/*==========================================================================*/
uint8_t DTC::s_InterprocessSynchronizationRemoteState::GetIsSecondaryFailoverActive() const
{
	if (Size < offsetof(s_InterprocessSynchronizationRemoteState, IsSecondaryFailoverActive) + sizeof(IsSecondaryFailoverActive))
		return 0;

	return IsSecondaryFailoverActive;
}

/****************************************************************************/
// s_InterprocessSynchronizationSnapshotRequest

/*==========================================================================*/
uint16_t DTC::s_InterprocessSynchronizationSnapshotRequest::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_InterprocessSynchronizationSnapshotRequest::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_InterprocessSynchronizationSnapshotRequest), *static_cast<uint16_t*>(p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_InterprocessSynchronizationSnapshotRequest::GetRequestID() const
{
	if (Size < offsetof(s_InterprocessSynchronizationSnapshotRequest, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/****************************************************************************/
// s_InterprocessSynchronizationTradeActivityRequest

/*==========================================================================*/
uint16_t DTC::s_InterprocessSynchronizationTradeActivityRequest::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_InterprocessSynchronizationTradeActivityRequest::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_InterprocessSynchronizationTradeActivityRequest), *static_cast<uint16_t*>(p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_InterprocessSynchronizationTradeActivityRequest::GetRequestID() const
{
	if (Size < offsetof(s_InterprocessSynchronizationTradeActivityRequest, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
int64_t DTC::s_InterprocessSynchronizationTradeActivityRequest::GetStartDateTimeUTC() const
{
	if (Size < offsetof(s_InterprocessSynchronizationTradeActivityRequest, StartDateTimeUTC) + sizeof(StartDateTimeUTC))
		return 0;

	return StartDateTimeUTC;
}


/****************************************************************************/
// s_SCConfigurationSynchronization

/*==========================================================================*/
uint16_t DTC_VLS::s_SCConfigurationSynchronization::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_SCConfigurationSynchronization::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_SCConfigurationSynchronization::GetMessageID() const
{
	if (BaseSize < offsetof(s_SCConfigurationSynchronization, MessageID) + sizeof(MessageID))
		return 0;

	return MessageID;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_SCConfigurationSynchronization::GetCurrentInboundSequenceNumber() const
{
	if (BaseSize < offsetof(s_SCConfigurationSynchronization, CurrentInboundSequenceNumber) + sizeof(CurrentInboundSequenceNumber))
		return 0;

	return CurrentInboundSequenceNumber;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_SCConfigurationSynchronization::GetCurrentOutboundSequenceNumber() const
{
	if (BaseSize < offsetof(s_SCConfigurationSynchronization, CurrentOutboundSequenceNumber) + sizeof(CurrentOutboundSequenceNumber))
		return 0;

	return CurrentOutboundSequenceNumber;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_SCConfigurationSynchronization::GetCurrentInternalOrderID() const
{
	if (BaseSize < offsetof(s_SCConfigurationSynchronization, CurrentInternalOrderID) + sizeof(CurrentInternalOrderID))
		return 0;

	return CurrentInternalOrderID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_SCConfigurationSynchronization::GetIsSnapshot() const
{
	if (BaseSize < offsetof(s_SCConfigurationSynchronization, IsSnapshot) + sizeof(IsSnapshot))
		return 0;

	return IsSnapshot;
}

/****************************************************************************/
// s_DownloadHistoricalOrderFillAndAccountBalanceData

/*==========================================================================*/
uint16_t DTC_VLS::s_DownloadHistoricalOrderFillAndAccountBalanceData::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_DownloadHistoricalOrderFillAndAccountBalanceData::GetBaseSize() const
{
	return BaseSize;
}

/****************************************************************************/
// s_HistoricalTradesRequest

/*==========================================================================*/
uint16_t DTC_VLS::s_HistoricalTradesRequest::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_HistoricalTradesRequest::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
int32_t DTC_VLS::s_HistoricalTradesRequest::GetRequestID() const
{
	if (BaseSize < offsetof(s_HistoricalTradesRequest, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
DTC::t_DateTime DTC_VLS::s_HistoricalTradesRequest::GetStartDateTime() const
{
	if (BaseSize < offsetof(s_HistoricalTradesRequest, StartDateTime) + sizeof(StartDateTime))
		return 0;

	return StartDateTime;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_HistoricalTradesRequest::GetSubAccountIdentifier() const
{
	if (BaseSize < offsetof(s_HistoricalTradesRequest, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_HistoricalTradesRequest::GetCreateFlatToFlatTrades() const
{
	if (BaseSize < offsetof(s_HistoricalTradesRequest, CreateFlatToFlatTrades) + sizeof(CreateFlatToFlatTrades))
		return 0;

	return CreateFlatToFlatTrades;
}

/****************************************************************************/
// s_HistoricalTradesReject

/*==========================================================================*/
uint16_t DTC_VLS::s_HistoricalTradesReject::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_HistoricalTradesReject::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_HistoricalTradesReject::GetRequestID() const
{
	if (BaseSize < offsetof(s_HistoricalTradesReject, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/****************************************************************************/
// s_HistoricalTradesResponse

/*==========================================================================*/
uint16_t DTC_VLS::s_HistoricalTradesResponse::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_HistoricalTradesResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
int32_t DTC_VLS::s_HistoricalTradesResponse::GetRequestID() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_HistoricalTradesResponse::GetIsFinalMessage() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, IsFinalMessage) + sizeof(IsFinalMessage))
		return 0;

	return IsFinalMessage;
}

/*==========================================================================*/
DTC::t_DateTimeWithMilliseconds DTC_VLS::s_HistoricalTradesResponse::GetEntryDateTime() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, EntryDateTime) + sizeof(EntryDateTime))
		return 0.0;

	return EntryDateTime;
}

/*==========================================================================*/
DTC::t_DateTimeWithMilliseconds DTC_VLS::s_HistoricalTradesResponse::GetExitDateTime() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, ExitDateTime) + sizeof(ExitDateTime))
		return DTC::BUY_SELL_UNSET;

	return ExitDateTime;
}

/*==========================================================================*/
double DTC_VLS::s_HistoricalTradesResponse::GetEntryPrice() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, EntryPrice) + sizeof(EntryPrice))
		return 0.0;

	return EntryPrice;
}

/*==========================================================================*/
double DTC_VLS::s_HistoricalTradesResponse::GetExitPrice() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, ExitPrice) + sizeof(ExitPrice))
		return 0.0;

	return ExitPrice;
}

/*==========================================================================*/
DTC::BuySellEnum DTC_VLS::s_HistoricalTradesResponse::GetTradeType() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, TradeType) + sizeof(TradeType))
		return DTC::BUY_SELL_UNSET;

	return TradeType;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_HistoricalTradesResponse::GetEntryQuantity() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, EntryQuantity) + sizeof(EntryQuantity))
		return 0;

	return EntryQuantity;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_HistoricalTradesResponse::GetExitQuantity() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, ExitQuantity) + sizeof(ExitQuantity))
		return 0;

	return ExitQuantity;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_HistoricalTradesResponse::GetMaxOpenQuantity() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, MaxOpenQuantity) + sizeof(MaxOpenQuantity))
		return 0;

	return MaxOpenQuantity;
}

/*==========================================================================*/
double DTC_VLS::s_HistoricalTradesResponse::GetClosedProfitLoss() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, ClosedProfitLoss) + sizeof(ClosedProfitLoss))
		return 0.0;

	return ClosedProfitLoss;
}

/*==========================================================================*/
double DTC_VLS::s_HistoricalTradesResponse::GetMaximumOpenPositionLoss() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, MaximumOpenPositionLoss) + sizeof(MaximumOpenPositionLoss))
		return 0.0;

	return MaximumOpenPositionLoss;
}

/*==========================================================================*/
double DTC_VLS::s_HistoricalTradesResponse::GetMaximumOpenPositionProfit() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, MaximumOpenPositionProfit) + sizeof(MaximumOpenPositionProfit))
		return 0.0;

	return MaximumOpenPositionProfit;
}

/*==========================================================================*/
double DTC_VLS::s_HistoricalTradesResponse::GetCommission() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, Commission) + sizeof(Commission))
		return 0.0;

	return Commission;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_HistoricalTradesResponse::GetSubAccountIdentifier() const
{
	if (BaseSize < offsetof(s_HistoricalTradesResponse, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/****************************************************************************/
// s_ReplayChartData

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartData::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartData::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_ReplayChartData::GetRequestID() const
{
	if (BaseSize < offsetof(s_ReplayChartData, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
DTC::t_DateTimeWithMillisecondsInt DTC_VLS::s_ReplayChartData::GetStartDateTimeForInitialData() const
{
	if (BaseSize < offsetof(s_ReplayChartData, StartDateTimeForInitialData) + sizeof(StartDateTimeForInitialData))
		return 0;

	return StartDateTimeForInitialData;
}

/*==========================================================================*/
DTC::t_DateTimeWithMillisecondsInt DTC_VLS::s_ReplayChartData::GetStartDateTime() const
{
	if (BaseSize < offsetof(s_ReplayChartData, StartDateTime) + sizeof(StartDateTime))
		return 0;

	return StartDateTime;
}

/*==========================================================================*/
DTC::t_DateTimeWithMillisecondsInt DTC_VLS::s_ReplayChartData::GetStopDateTime() const
{
	if (BaseSize < offsetof(s_ReplayChartData, StopDateTime) + sizeof(StopDateTime))
		return 0;

	return StopDateTime;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartData::GetSessionBeginTimeInSeconds() const
{
	if (BaseSize < offsetof(s_ReplayChartData, SessionBeginTimeInSeconds) + sizeof(SessionBeginTimeInSeconds))
		return 0;

	return SessionBeginTimeInSeconds;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartData::GetSessionEndTimeInSeconds() const
{
	if (BaseSize < offsetof(s_ReplayChartData, SessionEndTimeInSeconds) + sizeof(SessionEndTimeInSeconds))
		return 0;

	return SessionEndTimeInSeconds;
}

/*==========================================================================*/
float DTC_VLS::s_ReplayChartData::GetReplaySpeed() const
{
	if (BaseSize < offsetof(s_ReplayChartData, ReplaySpeed) + sizeof(ReplaySpeed))
		return 0.0;

	return ReplaySpeed;
}

/*==========================================================================*/
int32_t DTC_VLS::s_ReplayChartData::GetBarTimeInSeconds() const
{
	if (BaseSize < offsetof(s_ReplayChartData, BarTimeInSeconds) + sizeof(BarTimeInSeconds))
		return 0;

	return BarTimeInSeconds;
}
/*==========================================================================*/
uint8_t DTC_VLS::s_ReplayChartData::GetPauseReplayAfterInitialDataSent() const
{
	if (BaseSize < offsetof(s_ReplayChartData, PauseReplayAfterInitialDataSent) + sizeof(PauseReplayAfterInitialDataSent))
		return 0;

	return PauseReplayAfterInitialDataSent;
}
/*==========================================================================*/
uint8_t DTC_VLS::s_ReplayChartData::GetUseSavedPriorState() const
{
	if (BaseSize < offsetof(s_ReplayChartData, UseSavedPriorState) + sizeof(UseSavedPriorState))
		return 0;

	return UseSavedPriorState;
}

/*==========================================================================*/
float DTC_VLS::s_ReplayChartData::GetSymbolVolatility() const
{
	if (BaseSize < offsetof(s_ReplayChartData, SymbolVolatility) + sizeof(SymbolVolatility))
		return 0.0;

	return SymbolVolatility;
}

/*==========================================================================*/
float DTC_VLS::s_ReplayChartData::GetInterestRate() const
{
	if (BaseSize < offsetof(s_ReplayChartData, InterestRate) + sizeof(InterestRate))
		return 0.0;

	return InterestRate;
}

/*==========================================================================*/
int32_t DTC_VLS::s_ReplayChartData::GetNumberOfOrdersToTriggerFinishToStopDateTime() const
{
	if (BaseSize < offsetof(s_ReplayChartData, NumberOfOrdersToTriggerFinishToStopDateTime) + sizeof(NumberOfOrdersToTriggerFinishToStopDateTime))
		return 0;

	return NumberOfOrdersToTriggerFinishToStopDateTime;
}

/*==========================================================================*/
int32_t DTC_VLS::s_ReplayChartData::GetMaximumNumberOfOrdersPerReplaySession() const
{
	if (BaseSize < offsetof(s_ReplayChartData, MaximumNumberOfOrdersPerReplaySession) + sizeof(MaximumNumberOfOrdersPerReplaySession))
		return 0;

	return MaximumNumberOfOrdersPerReplaySession;
}

/*==========================================================================*/
int32_t DTC_VLS::s_ReplayChartData::GetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime() const
{
	if (BaseSize < offsetof(s_ReplayChartData, NumberOfDaysForInitialDataFromBeforeLastSavedDateTime) + sizeof(NumberOfDaysForInitialDataFromBeforeLastSavedDateTime))
		return 0;

	return NumberOfDaysForInitialDataFromBeforeLastSavedDateTime;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_ReplayChartData::GetSubAccountIdentifier() const
{
	if (BaseSize < offsetof(s_ReplayChartData, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/*==========================================================================*/
int32_t DTC_VLS::s_ReplayChartData::GetOptionsPriceSendIntervalInSeconds() const
{
	if (BaseSize < offsetof(s_ReplayChartData, OptionsPriceSendIntervalInSeconds) + sizeof(OptionsPriceSendIntervalInSeconds))
		return 0;

	return OptionsPriceSendIntervalInSeconds;
}

/****************************************************************************/
// s_ReplayChartDataPerformAction

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartDataPerformAction::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartDataPerformAction::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_ReplayChartDataPerformAction::GetRequestID() const
{
	if (BaseSize < offsetof(s_ReplayChartDataPerformAction, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
n_DTCNonStandard::ReplayChartDataActionEnum DTC_VLS::s_ReplayChartDataPerformAction::GetAction() const
{
	if (BaseSize < offsetof(s_ReplayChartDataPerformAction, Action) + sizeof(Action))
		return n_DTCNonStandard::REPLAY_CHART_DATA_ACTION_NONE;

	return Action;
}

/*==========================================================================*/
float DTC_VLS::s_ReplayChartDataPerformAction::GetReplaySpeed() const
{
	if (BaseSize < offsetof(s_ReplayChartDataPerformAction, ReplaySpeed) + sizeof(ReplaySpeed))
		return 0.0;

	return ReplaySpeed;
}

/****************************************************************************/
// s_ReplayChartDataStatus

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartDataStatus::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_ReplayChartDataStatus::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_ReplayChartDataStatus::GetRequestID() const
{
	if (BaseSize < offsetof(s_ReplayChartDataStatus, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
n_DTCNonStandard::ReplayChartDataStatusEnum DTC_VLS::s_ReplayChartDataStatus::GetStatus() const
{
	if (BaseSize < offsetof(s_ReplayChartDataStatus, Status) + sizeof(Status))
		return n_DTCNonStandard::REPLAY_CHART_DATA_STATUS_UNSET;

	return Status;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_ReplayChartDataStatus::GetSubAccountIdentifier() const
{
	if (BaseSize < offsetof(s_ReplayChartDataStatus, SubAccountIdentifier) + sizeof(SubAccountIdentifier))
		return 0;

	return SubAccountIdentifier;
}

/****************************************************************************/
// s_RequestNumCurrentClientConnections

/*==========================================================================*/
uint16_t DTC_VLS::s_RequestNumCurrentClientConnections::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_RequestNumCurrentClientConnections::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_RequestNumCurrentClientConnections::GetRequestID() const
{
	if (BaseSize < offsetof(s_RequestNumCurrentClientConnections, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
int64_t DTC_VLS::s_RequestNumCurrentClientConnections::GetDeviceIdentifier() const
{
	if (BaseSize < offsetof(s_RequestNumCurrentClientConnections, DeviceIdentifier) + sizeof(DeviceIdentifier))
		return 0;

	return DeviceIdentifier;
}

/****************************************************************************/
// s_NumCurrentClientConnectionsResponse

/*==========================================================================*/
uint16_t DTC_VLS::s_NumCurrentClientConnectionsResponse::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_NumCurrentClientConnectionsResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_NumCurrentClientConnectionsResponse::GetRequestID() const
{
	if (BaseSize < offsetof(s_NumCurrentClientConnectionsResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
int32_t DTC_VLS::s_NumCurrentClientConnectionsResponse::GetNumConnectionsForDifferentDevices() const
{
	if (BaseSize < offsetof(s_NumCurrentClientConnectionsResponse, NumConnectionsForDifferentDevices) + sizeof(NumConnectionsForDifferentDevices))
		return 0;

	return NumConnectionsForDifferentDevices;
}

/*==========================================================================*/
int32_t DTC_VLS::s_NumCurrentClientConnectionsResponse::GetNumConnectionsForSameDevice() const
{
	if (BaseSize < offsetof(s_NumCurrentClientConnectionsResponse, NumConnectionsForSameDevice) + sizeof(NumConnectionsForSameDevice))
		return 0;

	return NumConnectionsForSameDevice;
}

/****************************************************************************/
// s_SCTradeOrder

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeOrder::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeOrder::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsOrderDeleted() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsOrderDeleted) + sizeof(m_IsOrderDeleted))
		return 0;

	return m_IsOrderDeleted;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_InternalOrderID) + sizeof(m_InternalOrderID))
		return 0;

	return m_InternalOrderID;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeOrder::GetOrderStatusCode() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderStatusCode) + sizeof(m_OrderStatusCode))
		return 0;

	return m_OrderStatusCode;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeOrder::GetOrderStatusBeforePendingModify() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderStatusBeforePendingModify) + sizeof(m_OrderStatusBeforePendingModify))
		return 0;

	return m_OrderStatusBeforePendingModify;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeOrder::GetOrderStatusBeforePendingCancel() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderStatusBeforePendingCancel) + sizeof(m_OrderStatusBeforePendingCancel))
		return 0;

	return m_OrderStatusBeforePendingCancel;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetOrderType() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderType) + sizeof(m_OrderType))
		return 0;

	return m_OrderType;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeOrder::GetBuySell() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_BuySell) + sizeof(m_BuySell))
		return 0;

	return m_BuySell;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetPrice1() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_Price1) + sizeof(m_Price1))
		return 0;

	return m_Price1;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetPrice2() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_Price2) + sizeof(m_Price2))
		return 0;

	return m_Price2;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetOrderQuantity() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderQuantity) + sizeof(m_OrderQuantity))
		return 0;

	return m_OrderQuantity;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetFilledQuantity() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_FilledQuantity) + sizeof(m_FilledQuantity))
		return 0;

	return m_FilledQuantity;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetAverageFillPrice() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_AverageFillPrice) + sizeof(m_AverageFillPrice))
		return 0;

	return m_AverageFillPrice;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetRealtimeFillStatus() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_RealtimeFillStatus) + sizeof(m_RealtimeFillStatus))
		return 0;

	return m_RealtimeFillStatus;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsRestingOrderDuringFill() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsRestingOrderDuringFill) + sizeof(m_IsRestingOrderDuringFill))
		return 0;

	return m_IsRestingOrderDuringFill;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetOrderRejectType() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderRejectType) + sizeof(m_OrderRejectType))
		return 0;

	return m_OrderRejectType;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeOrder::GetSubAccountIdentifier() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SubAccountIdentifier) + sizeof(m_SubAccountIdentifier))
		return 0;

	return m_SubAccountIdentifier;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetInternalOrderIDModifierForService() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_InternalOrderIDModifierForService) + sizeof(m_InternalOrderIDModifierForService))
		return 0;

	return m_InternalOrderIDModifierForService;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeOrder::GetSequenceNumberBasedClientOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SequenceNumberBasedClientOrderID) + sizeof(m_SequenceNumberBasedClientOrderID))
		return 0;

	return m_SequenceNumberBasedClientOrderID;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetEntryDateTime() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_EntryDateTime) + sizeof(m_EntryDateTime))
		return 0;

	return m_EntryDateTime;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetLastActionDateTime() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastActionDateTime) + sizeof(m_LastActionDateTime))
		return 0;

	return m_LastActionDateTime;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetServiceUpdateDateTimeUTC() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_ServiceUpdateDateTimeUTC) + sizeof(m_ServiceUpdateDateTimeUTC))
		return 0;

	return m_ServiceUpdateDateTimeUTC;
}

/*==========================================================================*/
unsigned int DTC_VLS::s_TradeOrder::GetOrderEntryTimeForService() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderEntryTimeForService) + sizeof(m_OrderEntryTimeForService))
		return 0;

	return m_OrderEntryTimeForService;
}

/*==========================================================================*/
unsigned int DTC_VLS::s_TradeOrder::GetLastModifyTimeForService() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastModifyTimeForService) + sizeof(m_LastModifyTimeForService))
		return 0;

	return m_LastModifyTimeForService;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetGoodTillDateTime() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_GoodTillDateTime) + sizeof(m_GoodTillDateTime))
		return 0;

	return m_GoodTillDateTime;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetTimeInForce() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TimeInForce) + sizeof(m_TimeInForce))
		return 0;

	return m_TimeInForce;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeOrder::GetOpenClose() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OpenClose) + sizeof(m_OpenClose))
		return 0;

	return m_OpenClose;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetTrailStopOffset1() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TrailStopOffset1) + sizeof(m_TrailStopOffset1))
		return 0;

	return m_TrailStopOffset1;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetTrailStopStep() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TrailStopStep) + sizeof(m_TrailStopStep))
		return 0;

	return m_TrailStopStep;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetTrailTriggerPrice() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TrailTriggerPrice) + sizeof(m_TrailTriggerPrice))
		return 0;

	return m_TrailTriggerPrice;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetTrailingStopTriggerOffset() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TrailingStopTriggerOffset) + sizeof(m_TrailingStopTriggerOffset))
		return 0;

	return m_TrailingStopTriggerOffset;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetTrailTriggerHit() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TrailTriggerHit) + sizeof(m_TrailTriggerHit))
		return 0;

	return m_TrailTriggerHit;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetTrailToBreakEvenStopOffset() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TrailToBreakEvenStopOffset) + sizeof(m_TrailToBreakEvenStopOffset))
		return 0;

	return m_TrailToBreakEvenStopOffset;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetMaximumChaseAmountAsPrice() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_MaximumChaseAmountAsPrice) + sizeof(m_MaximumChaseAmountAsPrice))
		return 0;

	return m_MaximumChaseAmountAsPrice;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetInitialChaseOrderPrice1() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_InitialChaseOrderPrice1) + sizeof(m_InitialChaseOrderPrice1))
		return 0;

	return m_InitialChaseOrderPrice1;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetInitialLastTradePriceForChaseOrders() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_InitialLastTradePriceForChaseOrders) + sizeof(m_InitialLastTradePriceForChaseOrders))
		return 0;

	return m_InitialLastTradePriceForChaseOrders;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetTrailingStopTriggerOCOGroupNumber() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TrailingStopTriggerOCOGroupNumber) + sizeof(m_TrailingStopTriggerOCOGroupNumber))
		return 0;

	return m_TrailingStopTriggerOCOGroupNumber;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetLastModifyPrice1() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastModifyPrice1) + sizeof(m_LastModifyPrice1))
		return 0;

	return m_LastModifyPrice1;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetLastModifyQuantity() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastModifyQuantity) + sizeof(m_LastModifyQuantity))
		return 0;

	return m_LastModifyQuantity;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetCumulativeOrderQuantityFromParentFills() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_CumulativeOrderQuantityFromParentFills) + sizeof(m_CumulativeOrderQuantityFromParentFills))
		return 0;

	return m_CumulativeOrderQuantityFromParentFills;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetPriorFilledQuantity() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_PriorFilledQuantity) + sizeof(m_PriorFilledQuantity))
		return 0;

	return m_PriorFilledQuantity;
}

/*==========================================================================*/
float DTC_VLS::s_TradeOrder::GetTickSize() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TickSize) + sizeof(m_TickSize))
		return 0;

	return m_TickSize;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetValueFormat() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_ValueFormat) + sizeof(m_ValueFormat))
		return 0;

	return m_ValueFormat;
}

/*==========================================================================*/
float DTC_VLS::s_TradeOrder::GetPriceMultiplier() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_PriceMultiplier) + sizeof(m_PriceMultiplier))
		return 0;

	return m_PriceMultiplier;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetParentInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_ParentInternalOrderID) + sizeof(m_ParentInternalOrderID))
		return 0;

	return m_ParentInternalOrderID;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetTargetChildInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TargetChildInternalOrderID) + sizeof(m_TargetChildInternalOrderID))
		return 0;

	return m_TargetChildInternalOrderID;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetStopChildInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_StopChildInternalOrderID) + sizeof(m_StopChildInternalOrderID))
		return 0;

	return m_StopChildInternalOrderID;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetAttachedOrderPriceOffset1() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_AttachedOrderPriceOffset1) + sizeof(m_AttachedOrderPriceOffset1))
		return 0;

	return m_AttachedOrderPriceOffset1;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetLinkInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LinkInternalOrderID) + sizeof(m_LinkInternalOrderID))
		return 0;

	return m_LinkInternalOrderID;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetOCOGroupInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OCOGroupInternalOrderID) + sizeof(m_OCOGroupInternalOrderID))
		return 0;

	return m_OCOGroupInternalOrderID;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetOCOSiblingInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OCOSiblingInternalOrderID) + sizeof(m_OCOSiblingInternalOrderID))
		return 0;

	return m_OCOSiblingInternalOrderID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetDisableChildAndSiblingRelatedActions() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_DisableChildAndSiblingRelatedActions) + sizeof(m_DisableChildAndSiblingRelatedActions))
		return 0;

	return m_DisableChildAndSiblingRelatedActions;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetOCOManagedByService() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OCOManagedByService) + sizeof(m_OCOManagedByService))
		return 0;

	return m_OCOManagedByService;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetBracketOrderServerManaged() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_BracketOrderServerManaged) + sizeof(m_BracketOrderServerManaged))
		return 0;

	return m_BracketOrderServerManaged;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetStopLimitOrderStopPriceTriggered() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_StopLimitOrderStopPriceTriggered) + sizeof(m_StopLimitOrderStopPriceTriggered))
		return 0;

	return m_StopLimitOrderStopPriceTriggered;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetOCOFullSiblingCancelOnPartialFill() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OCOFullSiblingCancelOnPartialFill) + sizeof(m_OCOFullSiblingCancelOnPartialFill))
		return 0;

	return m_OCOFullSiblingCancelOnPartialFill;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetReverseOnCompleteFill() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_ReverseOnCompleteFill) + sizeof(m_ReverseOnCompleteFill))
		return 0;

	return m_ReverseOnCompleteFill;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetSupportScaleIn() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SupportScaleIn) + sizeof(m_SupportScaleIn))
		return 0;

	return m_SupportScaleIn;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetSupportScaleOut() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SupportScaleOut) + sizeof(m_SupportScaleOut))
		return 0;

	return m_SupportScaleOut;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetSourceChartNumber() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SourceChartNumber) + sizeof(m_SourceChartNumber))
		return 0;

	return m_SourceChartNumber;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsAutomatedOrder() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsAutomatedOrder) + sizeof(m_IsAutomatedOrder))
		return 0;

	return m_IsAutomatedOrder;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetSimulatedOrder() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SimulatedOrder) + sizeof(m_SimulatedOrder))
		return 0;

	return m_SimulatedOrder;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsChartReplaying() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsChartReplaying) + sizeof(m_IsChartReplaying))
		return 0;

	return m_IsChartReplaying;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetAttachedOrderOCOGroupNumber() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_AttachedOrderOCOGroupNumber) + sizeof(m_AttachedOrderOCOGroupNumber))
		return 0;

	return m_AttachedOrderOCOGroupNumber;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetFillCount() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_FillCount) + sizeof(m_FillCount))
		return 0;

	return m_FillCount;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetLastFillQuantity() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastFillQuantity) + sizeof(m_LastFillQuantity))
		return 0;

	return m_LastFillQuantity;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetLastFillPrice() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastFillPrice) + sizeof(m_LastFillPrice))
		return 0;

	return m_LastFillPrice;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetLastFillDateTimeUTC() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastFillDateTimeUTC) + sizeof(m_LastFillDateTimeUTC))
		return 0;

	return m_LastFillDateTimeUTC;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeOrder::GetRejectedStopOCOSiblingInternalOrderID() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_RejectedStopOCOSiblingInternalOrderID) + sizeof(m_RejectedStopOCOSiblingInternalOrderID))
		return 0;

	return m_RejectedStopOCOSiblingInternalOrderID;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetRejectedStopReplacementMarketOrderQuantity() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_RejectedStopReplacementMarketOrderQuantity) + sizeof(m_RejectedStopReplacementMarketOrderQuantity))
		return 0;

	return m_RejectedStopReplacementMarketOrderQuantity;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetEvaluatingForFill() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_EvaluatingForFill) + sizeof(m_EvaluatingForFill))
		return 0;

	return m_EvaluatingForFill;
}

/*==========================================================================*/
unsigned int DTC_VLS::s_TradeOrder::GetLastProcessedTimeSalesRecordSequenceForPrices() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_LastProcessedTimeSalesRecordSequenceForPrices) + sizeof(m_LastProcessedTimeSalesRecordSequenceForPrices))
		return 0;

	return m_LastProcessedTimeSalesRecordSequenceForPrices;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsMarketDataManagementOfOrderEnabled() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsMarketDataManagementOfOrderEnabled) + sizeof(m_IsMarketDataManagementOfOrderEnabled))
		return 0;

	return m_IsMarketDataManagementOfOrderEnabled;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetTimedOutOrderRequestedStatusDateTime() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_TimedOutOrderRequestedStatusDateTime) + sizeof(m_TimedOutOrderRequestedStatusDateTime))
		return 0;

	return m_TimedOutOrderRequestedStatusDateTime;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetRequestedStatusForTimedOutOrder() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_RequestedStatusForTimedOutOrder) + sizeof(m_RequestedStatusForTimedOutOrder))
		return 0;

	return m_RequestedStatusForTimedOutOrder;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled) + sizeof(m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled))
		return 0;

	return m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetQuantityToIncreaseFromParentFill() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_QuantityToIncreaseFromParentFill) + sizeof(m_QuantityToIncreaseFromParentFill))
		return 0;

	return m_QuantityToIncreaseFromParentFill;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetMoveToBreakevenStopReferencePrice() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_MoveToBreakevenStopReferencePrice) + sizeof(m_MoveToBreakevenStopReferencePrice))
		return 0;

	return m_MoveToBreakevenStopReferencePrice;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetQuantityTriggeredStop_QuantityForTrigger() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_QuantityTriggeredStop_QuantityForTrigger) + sizeof(m_QuantityTriggeredStop_QuantityForTrigger))
		return 0;

	return m_QuantityTriggeredStop_QuantityForTrigger;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetAccumulatedTradeVolumeForTriggeredStop() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_AccumulatedTradeVolumeForTriggeredStop) + sizeof(m_AccumulatedTradeVolumeForTriggeredStop))
		return 0;

	return m_AccumulatedTradeVolumeForTriggeredStop;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetBidAskQuantityStopInitialTriggerMet() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_BidAskQuantityStopInitialTriggerMet) + sizeof(m_BidAskQuantityStopInitialTriggerMet))
		return 0;

	return m_BidAskQuantityStopInitialTriggerMet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetNeedToOverrideLock() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_NeedToOverrideLock) + sizeof(m_NeedToOverrideLock))
		return 0;

	return m_NeedToOverrideLock;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetCurrentMarketPrice() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_CurrentMarketPrice) + sizeof(m_CurrentMarketPrice))
		return 0;

	return m_CurrentMarketPrice;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetCurrentMarketDateTime() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_CurrentMarketDateTime) + sizeof(m_CurrentMarketDateTime))
		return 0;

	return m_CurrentMarketDateTime;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetSupportOrderFillBilling() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_SupportOrderFillBilling) + sizeof(m_SupportOrderFillBilling))
		return 0;

	return m_SupportOrderFillBilling;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsBillable() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsBillable) + sizeof(m_IsBillable))
		return 0;

	return m_IsBillable;
}

/*==========================================================================*/
int DTC_VLS::s_TradeOrder::GetQuantityForBilling() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_QuantityForBilling) + sizeof(m_QuantityForBilling))
		return 0;

	return m_QuantityForBilling;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeOrder::GetNumberOfFailedOrderModifications() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_NumberOfFailedOrderModifications) + sizeof(m_NumberOfFailedOrderModifications))
		return 0;

	return m_NumberOfFailedOrderModifications;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeOrder::GetDTCServerIndex() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_DTCServerIndex) + sizeof(m_DTCServerIndex))
		return 0;

	return m_DTCServerIndex;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeOrder::GetCTICode() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_CTICode) + sizeof(m_CTICode))
		return 0;

	return m_CTICode;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetObtainOrderActionDateTimeFromLastTradeTimeInChart() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_ObtainOrderActionDateTimeFromLastTradeTimeInChart) + sizeof(m_ObtainOrderActionDateTimeFromLastTradeTimeInChart))
		return 0;

	return m_ObtainOrderActionDateTimeFromLastTradeTimeInChart;
}

/*==========================================================================*/
double DTC_VLS::s_TradeOrder::GetMaximumShowQuantity() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_MaximumShowQuantity) + sizeof(m_MaximumShowQuantity))
		return 0;

	return m_MaximumShowQuantity;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetOrderSubmitted() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_OrderSubmitted) + sizeof(m_OrderSubmitted))
		return 0;

	return m_OrderSubmitted;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsSnapshot() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsSnapshot) + sizeof(m_IsSnapshot))
		return 0;

	return m_IsSnapshot;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsFirstMessageInBatch() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsFirstMessageInBatch) + sizeof(m_IsFirstMessageInBatch))
		return 0;

	return m_IsFirstMessageInBatch;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeOrder::GetIsLastMessageInBatch() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_IsLastMessageInBatch) + sizeof(m_IsLastMessageInBatch))
		return 0;

	return m_IsLastMessageInBatch;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeOrder::GetExternalLastActionDateTimeUTC() const
{
	if (BaseSize < offsetof(s_TradeOrder, m_ExternalLastActionDateTimeUTC) + sizeof(m_ExternalLastActionDateTimeUTC))
		return 0;

	return m_ExternalLastActionDateTimeUTC;
}

/****************************************************************************/
//s_IndividualTradePosition

/*==========================================================================*/
uint16_t DTC_VLS::s_IndividualTradePosition::GetMessageSize() const
{	
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_IndividualTradePosition::GetBaseSize() const
{
	return BaseSize;
}
/*==========================================================================*/
/****************************************************************************/
//s_TradePositionConsolidated

/*==========================================================================*/
uint16_t DTC_VLS::s_TradePositionConsolidated::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradePositionConsolidated::GetBaseSize() const
{
	return BaseSize;
}
/*==========================================================================*/

/****************************************************************************/
//s_TradeActivityData

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeActivityData::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeActivityData::GetBaseSize() const
{
	return BaseSize;
}
/*==========================================================================*/

/****************************************************************************/
//s_RequestTradeAccountData

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataRequest::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataRequest::GetBaseSize() const
{
	return BaseSize;
}
/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataRequest::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataRequest, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataResponse

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataResponse::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataResponse::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_RequestID) + sizeof(m_RequestID))
		return 0;

	return m_RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetTradeAccountNotExist() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_TradeAccountNotExist) + sizeof(m_TradeAccountNotExist))
		return 0;

	return m_TradeAccountNotExist;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsSimulated() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsSimulated) + sizeof(m_IsSimulated))
		return 0;

	return m_IsSimulated;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetCurrentCashBalance() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_CurrentCashBalance) + sizeof(m_CurrentCashBalance))
		return 0;

	return m_CurrentCashBalance;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetAvailableFundsForNewPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_AvailableFundsForNewPositions) + sizeof(m_AvailableFundsForNewPositions))
		return 0;

	return m_AvailableFundsForNewPositions;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetMarginRequirement() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MarginRequirement) + sizeof(m_MarginRequirement))
		return 0;

	return m_MarginRequirement;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetAccountValue() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_AccountValue) + sizeof(m_AccountValue))
		return 0;

	return m_AccountValue;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetOpenPositionsProfitLoss() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_OpenPositionsProfitLoss) + sizeof(m_OpenPositionsProfitLoss))
		return 0;

	return m_OpenPositionsProfitLoss;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetDailyProfitLoss() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_DailyProfitLoss) + sizeof(m_DailyProfitLoss))
		return 0;

	return m_DailyProfitLoss;
}

/*==========================================================================*/
uint64_t DTC_VLS::s_TradeAccountDataResponse::GetTransactionIdentifierForCashBalanceAdjustment() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_TransactionIdentifierForCashBalanceAdjustment) + sizeof(m_TransactionIdentifierForCashBalanceAdjustment))
		return 0;

	return m_TransactionIdentifierForCashBalanceAdjustment;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataResponse::GetLastTransactionDateTime() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_LastTransactionDateTime) + sizeof(m_LastTransactionDateTime))
		return 0;

	return m_LastTransactionDateTime;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetTrailingAccountValueAtWhichToNotAllowNewPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_TrailingAccountValueAtWhichToNotAllowNewPositions) + sizeof(m_TrailingAccountValueAtWhichToNotAllowNewPositions))
		return 0;

	return m_TrailingAccountValueAtWhichToNotAllowNewPositions;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetCalculatedDailyNetLossLimitInAccountCurrency() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_CalculatedDailyNetLossLimitInAccountCurrency) + sizeof(m_CalculatedDailyNetLossLimitInAccountCurrency))
		return 0;

	return m_CalculatedDailyNetLossLimitInAccountCurrency;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetDailyNetLossLimitHasBeenReached() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_DailyNetLossLimitHasBeenReached) + sizeof(m_DailyNetLossLimitHasBeenReached))
		return 0;

	return m_DailyNetLossLimitHasBeenReached;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataResponse::GetLastResetDailyNetLossManagementVariablesDateTimeUTC() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_LastResetDailyNetLossManagementVariablesDateTimeUTC) + sizeof(m_LastResetDailyNetLossManagementVariablesDateTimeUTC))
		return 0;

	return m_LastResetDailyNetLossManagementVariablesDateTimeUTC;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsUnderRequiredMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsUnderRequiredMargin) + sizeof(m_IsUnderRequiredMargin))
		return 0;

	return m_IsUnderRequiredMargin;
}

/*==========================================================================*/
float DTC_VLS::s_TradeAccountDataResponse::GetDailyNetLossLimitInAccountCurrency() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_DailyNetLossLimitInAccountCurrency) + sizeof(m_DailyNetLossLimitInAccountCurrency))
		return 0;

	return m_DailyNetLossLimitInAccountCurrency;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataResponse::GetPercentOfCashBalanceForDailyNetLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_PercentOfCashBalanceForDailyNetLossLimit) + sizeof(m_PercentOfCashBalanceForDailyNetLossLimit))
		return 0;

	return m_PercentOfCashBalanceForDailyNetLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseTrailingAccountValueToNotAllowIncreaseInPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseTrailingAccountValueToNotAllowIncreaseInPositions) + sizeof(m_UseTrailingAccountValueToNotAllowIncreaseInPositions))
		return 0;

	return m_UseTrailingAccountValueToNotAllowIncreaseInPositions;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetDoNotAllowIncreaseInPositionsAtDailyLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_DoNotAllowIncreaseInPositionsAtDailyLossLimit) + sizeof(m_DoNotAllowIncreaseInPositionsAtDailyLossLimit))
		return 0;

	return m_DoNotAllowIncreaseInPositionsAtDailyLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetFlattenPositionsAtDailyLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_FlattenPositionsAtDailyLossLimit) + sizeof(m_FlattenPositionsAtDailyLossLimit))
		return 0;

	return m_FlattenPositionsAtDailyLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetClosePositionsAtEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_ClosePositionsAtEndOfDay) + sizeof(m_ClosePositionsAtEndOfDay))
		return 0;

	return m_ClosePositionsAtEndOfDay;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetFlattenPositionsWhenUnderMarginIntraday() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_FlattenPositionsWhenUnderMarginIntraday) + sizeof(m_FlattenPositionsWhenUnderMarginIntraday))
		return 0;

	return m_FlattenPositionsWhenUnderMarginIntraday;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetFlattenPositionsWhenUnderMarginAtEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_FlattenPositionsWhenUnderMarginAtEndOfDay) + sizeof(m_FlattenPositionsWhenUnderMarginAtEndOfDay))
		return 0;

	return m_FlattenPositionsWhenUnderMarginAtEndOfDay;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataResponse::GetCTICode() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_CTICode) + sizeof(m_CTICode))
		return 0;

	return m_CTICode;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetTradeAccountIsReadOnly() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_TradeAccountIsReadOnly) + sizeof(m_TradeAccountIsReadOnly))
		return 0;

	return m_TradeAccountIsReadOnly;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataResponse::GetMaximumGlobalPositionQuantity() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MaximumGlobalPositionQuantity) + sizeof(m_MaximumGlobalPositionQuantity))
		return 0;

	return m_MaximumGlobalPositionQuantity;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetTradeFeePerContract() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_TradeFeePerContract) + sizeof(m_TradeFeePerContract))
		return 0;

	return m_TradeFeePerContract;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetTradeFeePerShare() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_TradeFeePerShare) + sizeof(m_TradeFeePerShare))
		return 0;

	return m_TradeFeePerShare;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetRequireSufficientMarginForNewPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_RequireSufficientMarginForNewPositions) + sizeof(m_RequireSufficientMarginForNewPositions))
		return 0;

	return m_RequireSufficientMarginForNewPositions;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataResponse::GetUsePercentOfMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UsePercentOfMargin) + sizeof(m_UsePercentOfMargin))
		return 0;

	return m_UsePercentOfMargin;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataResponse::GetUsePercentOfMarginForDayTrading() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UsePercentOfMarginForDayTrading) + sizeof(m_UsePercentOfMarginForDayTrading))
		return 0;

	return m_UsePercentOfMarginForDayTrading;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataResponse::GetMaximumAllowedAccountBalanceForPositionsAsPercentage() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MaximumAllowedAccountBalanceForPositionsAsPercentage) + sizeof(m_MaximumAllowedAccountBalanceForPositionsAsPercentage))
		return 0;

	return m_MaximumAllowedAccountBalanceForPositionsAsPercentage;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetTradingIsDisabled() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_TradingIsDisabled) + sizeof(m_TradingIsDisabled))
		return 0;

	return m_TradingIsDisabled;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsMasterFirmControlAccount() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsMasterFirmControlAccount) + sizeof(m_IsMasterFirmControlAccount))
		return 0;

	return m_IsMasterFirmControlAccount;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetMinimumRequiredAccountValue() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MinimumRequiredAccountValue) + sizeof(m_MinimumRequiredAccountValue))
		return 0;

	return m_MinimumRequiredAccountValue;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataResponse::GetBeginTimeForDayMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_BeginTimeForDayMargin) + sizeof(m_BeginTimeForDayMargin))
		return 0;

	return m_BeginTimeForDayMargin;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataResponse::GetEndTimeForDayMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_EndTimeForDayMargin) + sizeof(m_EndTimeForDayMargin))
		return 0;

	return m_EndTimeForDayMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsSnapshot() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsSnapshot) + sizeof(m_IsSnapshot))
		return 0;

	return m_IsSnapshot;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsFirstMessageInBatch() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsFirstMessageInBatch) + sizeof(m_IsFirstMessageInBatch))
		return 0;

	return m_IsFirstMessageInBatch;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsLastMessageInBatch() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsLastMessageInBatch) + sizeof(m_IsLastMessageInBatch))
		return 0;

	return m_IsLastMessageInBatch;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsDeleted() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsDeleted) + sizeof(m_IsDeleted))
		return 0;

	return m_IsDeleted;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday) + sizeof(m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday))
		return 0;

	return m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay) + sizeof(m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay))
		return 0;

	return m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_SymbolLimitsArray() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_SymbolLimitsArray) + sizeof(m_UseMasterFirm_SymbolLimitsArray))
		return 0;

	return m_UseMasterFirm_SymbolLimitsArray;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_TradeFees() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_TradeFees) + sizeof(m_UseMasterFirm_TradeFees))
		return 0;

	return m_UseMasterFirm_TradeFees;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_TradeFeePerShare() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_TradeFeePerShare) + sizeof(m_UseMasterFirm_TradeFeePerShare))
		return 0;

	return m_UseMasterFirm_TradeFeePerShare;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_RequireSufficientMarginForNewPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_RequireSufficientMarginForNewPositions) + sizeof(m_UseMasterFirm_RequireSufficientMarginForNewPositions))
		return 0;

	return m_UseMasterFirm_RequireSufficientMarginForNewPositions;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_UsePercentOfMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_UsePercentOfMargin) + sizeof(m_UseMasterFirm_UsePercentOfMargin))
		return 0;

	return m_UseMasterFirm_UsePercentOfMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage) + sizeof(m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage))
		return 0;

	return m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_MinimumRequiredAccountValue() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_MinimumRequiredAccountValue) + sizeof(m_UseMasterFirm_MinimumRequiredAccountValue))
		return 0;

	return m_UseMasterFirm_MinimumRequiredAccountValue;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_MarginTimeSettings() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_MarginTimeSettings) + sizeof(m_UseMasterFirm_MarginTimeSettings))
		return 0;

	return m_UseMasterFirm_MarginTimeSettings;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_TradingIsDisabled() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_TradingIsDisabled) + sizeof(m_UseMasterFirm_TradingIsDisabled))
		return 0;

	return m_UseMasterFirm_TradingIsDisabled;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsTradeStatisticsPublicallyShared() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsTradeStatisticsPublicallyShared) + sizeof(m_IsTradeStatisticsPublicallyShared))
		return 0;

	return m_IsTradeStatisticsPublicallyShared;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsReadOnlyFollowingRequestsAllowed() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsReadOnlyFollowingRequestsAllowed) + sizeof(m_IsReadOnlyFollowingRequestsAllowed))
		return 0;

	return m_IsReadOnlyFollowingRequestsAllowed;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetIsTradeAccountSharingAllowed() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_IsTradeAccountSharingAllowed) + sizeof(m_IsTradeAccountSharingAllowed))
		return 0;

	return m_IsTradeAccountSharingAllowed;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetReadOnlyShareToAllSCUsernames() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_ReadOnlyShareToAllSCUsernames) + sizeof(m_ReadOnlyShareToAllSCUsernames))
		return 0;

	return m_ReadOnlyShareToAllSCUsernames;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_SymbolCommissionsArray() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_SymbolCommissionsArray) + sizeof(m_UseMasterFirm_SymbolCommissionsArray))
		return 0;

	return m_UseMasterFirm_SymbolCommissionsArray;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit) + sizeof(m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit))
		return 0;

	return m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_UsePercentOfMarginForDayTrading() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_UsePercentOfMarginForDayTrading) + sizeof(m_UseMasterFirm_UsePercentOfMarginForDayTrading))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginForDayTrading;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetOpenPositionsProfitLossBasedOnSettlement() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_OpenPositionsProfitLossBasedOnSettlement) + sizeof(m_OpenPositionsProfitLossBasedOnSettlement))
		return 0;

	return m_OpenPositionsProfitLossBasedOnSettlement;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetMarginRequirementFull() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MarginRequirementFull) + sizeof(m_MarginRequirementFull))
		return 0;

	return m_MarginRequirementFull;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetMarginRequirementFullPositionsOnly() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MarginRequirementFullPositionsOnly) + sizeof(m_MarginRequirementFullPositionsOnly))
		return 0;

	return m_MarginRequirementFullPositionsOnly;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_TradeFeesFullOverride() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_TradeFeesFullOverride) + sizeof(m_UseMasterFirm_TradeFeesFullOverride))
		return 0;

	return m_UseMasterFirm_TradeFeesFullOverride;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders) + sizeof(m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders))
		return 0;

	return m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_UsePercentOfMarginFullOverride() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_UsePercentOfMarginFullOverride) + sizeof(m_UseMasterFirm_UsePercentOfMarginFullOverride))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginFullOverride;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride) + sizeof(m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetPeakMarginRequirement() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_PeakMarginRequirement) + sizeof(m_PeakMarginRequirement))
		return 0;

	return m_PeakMarginRequirement;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetLiquidationOnlyMode() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_LiquidationOnlyMode) + sizeof(m_LiquidationOnlyMode))
		return 0;

	return m_LiquidationOnlyMode;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetFlattenPositionsWhenUnderMarginIntradayTriggered() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_FlattenPositionsWhenUnderMarginIntradayTriggered) + sizeof(m_FlattenPositionsWhenUnderMarginIntradayTriggered))
		return 0;

	return m_FlattenPositionsWhenUnderMarginIntradayTriggered;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetFlattenPositionsWhenUnderMinimumAccountValueTriggered() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_FlattenPositionsWhenUnderMinimumAccountValueTriggered) + sizeof(m_FlattenPositionsWhenUnderMinimumAccountValueTriggered))
		return 0;

	return m_FlattenPositionsWhenUnderMinimumAccountValueTriggered;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataResponse::GetAccountValueAtEndOfDayCaptureTime() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_AccountValueAtEndOfDayCaptureTime) + sizeof(m_AccountValueAtEndOfDayCaptureTime))
		return 0;

	return m_AccountValueAtEndOfDayCaptureTime;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataResponse::GetEndOfDayCaptureTime() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_EndOfDayCaptureTime) + sizeof(m_EndOfDayCaptureTime))
		return 0;

	return m_EndOfDayCaptureTime;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetCustomerOrFirm() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_CustomerOrFirm) + sizeof(m_CustomerOrFirm))
		return 0;

	return m_CustomerOrFirm;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet) + sizeof(m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet))
		return 0;

	return m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetMasterFirm_FlattenCancelWhenUnderMarginIntraday() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MasterFirm_FlattenCancelWhenUnderMarginIntraday) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMarginIntraday))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMarginIntraday;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataResponse::GetMasterFirm_MaximumOrderQuantity() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_MasterFirm_MaximumOrderQuantity) + sizeof(m_MasterFirm_MaximumOrderQuantity))
		return 0;

	return m_MasterFirm_MaximumOrderQuantity;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataResponse::GetLastTriggerDateTimeUTCForDailyLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_LastTriggerDateTimeUTCForDailyLossLimit) + sizeof(m_LastTriggerDateTimeUTCForDailyLossLimit))
		return 0;

	return m_LastTriggerDateTimeUTCForDailyLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponse::GetOpenPositionsProfitLossIsDelayed() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponse, m_OpenPositionsProfitLossIsDelayed) + sizeof(m_OpenPositionsProfitLossIsDelayed))
		return 0;

	return m_OpenPositionsProfitLossIsDelayed;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataUpdate

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUpdate::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUpdate::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataUpdate::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsNewAccount() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsNewAccount) + sizeof(IsNewAccount))
		return 0;

	return IsNewAccount;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetCurrencyCodeIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, CurrencyCodeIsSet) + sizeof(CurrencyCodeIsSet))
		return 0;

	return CurrencyCodeIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetDailyNetLossLimitInAccountCurrencyIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, DailyNetLossLimitInAccountCurrencyIsSet) + sizeof(DailyNetLossLimitInAccountCurrencyIsSet))
		return 0;

	return DailyNetLossLimitInAccountCurrencyIsSet;
}

/*==========================================================================*/
float DTC_VLS::s_TradeAccountDataUpdate::GetDailyNetLossLimitInAccountCurrency() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, DailyNetLossLimitInAccountCurrency) + sizeof(DailyNetLossLimitInAccountCurrency))
		return 0;

	return DailyNetLossLimitInAccountCurrency;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetPercentOfCashBalanceForDailyNetLossLimitIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, PercentOfCashBalanceForDailyNetLossLimitIsSet) + sizeof(PercentOfCashBalanceForDailyNetLossLimitIsSet))
		return 0;

	return PercentOfCashBalanceForDailyNetLossLimitIsSet;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataUpdate::GetPercentOfCashBalanceForDailyNetLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, PercentOfCashBalanceForDailyNetLossLimit) + sizeof(PercentOfCashBalanceForDailyNetLossLimit))
		return 0;

	return PercentOfCashBalanceForDailyNetLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet) + sizeof(UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet))
		return 0;

	return UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseTrailingAccountValueToNotAllowIncreaseInPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, UseTrailingAccountValueToNotAllowIncreaseInPositions) + sizeof(UseTrailingAccountValueToNotAllowIncreaseInPositions))
		return 0;

	return UseTrailingAccountValueToNotAllowIncreaseInPositions;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet) + sizeof(DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet))
		return 0;

	return DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetDoNotAllowIncreaseInPositionsAtDailyLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, DoNotAllowIncreaseInPositionsAtDailyLossLimit) + sizeof(DoNotAllowIncreaseInPositionsAtDailyLossLimit))
		return 0;

	return DoNotAllowIncreaseInPositionsAtDailyLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetFlattenPositionsAtDailyLossLimitIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, FlattenPositionsAtDailyLossLimitIsSet) + sizeof(FlattenPositionsAtDailyLossLimitIsSet))
		return 0;

	return FlattenPositionsAtDailyLossLimitIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetFlattenPositionsAtDailyLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, FlattenPositionsAtDailyLossLimit) + sizeof(FlattenPositionsAtDailyLossLimit))
		return 0;

	return FlattenPositionsAtDailyLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetClosePositionsAtEndOfDayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, ClosePositionsAtEndOfDayIsSet) + sizeof(ClosePositionsAtEndOfDayIsSet))
		return 0;

	return ClosePositionsAtEndOfDayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetClosePositionsAtEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, ClosePositionsAtEndOfDay) + sizeof(ClosePositionsAtEndOfDay))
		return 0;

	return ClosePositionsAtEndOfDay;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetFlattenPositionsWhenUnderMarginIntradayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, FlattenPositionsWhenUnderMarginIntradayIsSet) + sizeof(FlattenPositionsWhenUnderMarginIntradayIsSet))
		return 0;

	return FlattenPositionsWhenUnderMarginIntradayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetFlattenPositionsWhenUnderMarginIntraday() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, FlattenPositionsWhenUnderMarginIntraday) + sizeof(FlattenPositionsWhenUnderMarginIntraday))
		return 0;

	return FlattenPositionsWhenUnderMarginIntraday;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, FlattenPositionsWhenUnderMarginAtEndOfDayIsSet) + sizeof(FlattenPositionsWhenUnderMarginAtEndOfDayIsSet))
		return 0;

	return FlattenPositionsWhenUnderMarginAtEndOfDayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetFlattenPositionsWhenUnderMarginAtEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, FlattenPositionsWhenUnderMarginAtEndOfDay) + sizeof(FlattenPositionsWhenUnderMarginAtEndOfDay))
		return 0;

	return FlattenPositionsWhenUnderMarginAtEndOfDay;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetSenderSubIDIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, SenderSubIDIsSet) + sizeof(SenderSubIDIsSet))
		return 0;

	return SenderSubIDIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetSenderLocationIdIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, SenderLocationIdIsSet) + sizeof(SenderLocationIdIsSet))
		return 0;

	return SenderLocationIdIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetSelfMatchPreventionIDIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, SelfMatchPreventionIDIsSet) + sizeof(SelfMatchPreventionIDIsSet))
		return 0;

	return SelfMatchPreventionIDIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetCTICodeIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, CTICodeIsSet) + sizeof(CTICodeIsSet))
		return 0;

	return CTICodeIsSet;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataUpdate::GetCTICode() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, CTICode) + sizeof(CTICode))
		return 0;

	return CTICode;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetTradeAccountIsReadOnlyIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradeAccountIsReadOnlyIsSet) + sizeof(TradeAccountIsReadOnlyIsSet))
		return 0;

	return TradeAccountIsReadOnlyIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetTradeAccountIsReadOnly() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradeAccountIsReadOnly) + sizeof(TradeAccountIsReadOnly))
		return 0;

	return TradeAccountIsReadOnly;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMaximumGlobalPositionQuantityIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, MaximumGlobalPositionQuantityIsSet) + sizeof(MaximumGlobalPositionQuantityIsSet))
		return 0;

	return MaximumGlobalPositionQuantityIsSet;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataUpdate::GetMaximumGlobalPositionQuantity() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, MaximumGlobalPositionQuantity) + sizeof(MaximumGlobalPositionQuantity))
		return 0;

	return MaximumGlobalPositionQuantity;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetTradeFeePerContractIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradeFeePerContractIsSet) + sizeof(TradeFeePerContractIsSet))
		return 0;

	return TradeFeePerContractIsSet;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataUpdate::GetTradeFeePerContract() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradeFeePerContract) + sizeof(TradeFeePerContract))
		return 0;

	return TradeFeePerContract;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetTradeFeePerShareIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradeFeePerShareIsSet) + sizeof(TradeFeePerShareIsSet))
		return 0;

	return TradeFeePerShareIsSet;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataUpdate::GetTradeFeePerShare() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradeFeePerShare) + sizeof(TradeFeePerShare))
		return 0;

	return TradeFeePerShare;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetRequireSufficientMarginForNewPositionsIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, RequireSufficientMarginForNewPositionsIsSet) + sizeof(RequireSufficientMarginForNewPositionsIsSet))
		return 0;

	return RequireSufficientMarginForNewPositionsIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetRequireSufficientMarginForNewPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, RequireSufficientMarginForNewPositions) + sizeof(RequireSufficientMarginForNewPositions))
		return 0;

	return RequireSufficientMarginForNewPositions;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUsePercentOfMarginIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, UsePercentOfMarginIsSet) + sizeof(UsePercentOfMarginIsSet))
		return 0;

	return UsePercentOfMarginIsSet;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataUpdate::GetUsePercentOfMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, UsePercentOfMargin) + sizeof(UsePercentOfMargin))
		return 0;

	return UsePercentOfMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUsePercentOfMarginForDayTradingIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, UsePercentOfMarginForDayTradingIsSet) + sizeof(UsePercentOfMarginForDayTradingIsSet))
		return 0;

	return UsePercentOfMarginForDayTradingIsSet;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataUpdate::GetUsePercentOfMarginForDayTrading() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, UsePercentOfMarginForDayTrading) + sizeof(UsePercentOfMarginForDayTrading))
		return 0;

	return UsePercentOfMarginForDayTrading;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet) + sizeof(MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet))
		return 0;

	return MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataUpdate::GetMaximumAllowedAccountBalanceForPositionsAsPercentage() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, MaximumAllowedAccountBalanceForPositionsAsPercentage) + sizeof(MaximumAllowedAccountBalanceForPositionsAsPercentage))
		return 0;

	return MaximumAllowedAccountBalanceForPositionsAsPercentage;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetFirmIDIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, FirmIDIsSet) + sizeof(FirmIDIsSet))
		return 0;

	return FirmIDIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetTradingIsDisabledIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradingIsDisabledIsSet) + sizeof(TradingIsDisabledIsSet))
		return 0;

	return TradingIsDisabledIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetTradingIsDisabled() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, TradingIsDisabled) + sizeof(TradingIsDisabled))
		return 0;

	return TradingIsDisabled;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetDescriptiveNameIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, DescriptiveNameIsSet) + sizeof(DescriptiveNameIsSet))
		return 0;

	return DescriptiveNameIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsMasterFirmControlAccountIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsMasterFirmControlAccountIsSet) + sizeof(IsMasterFirmControlAccountIsSet))
		return 0;

	return IsMasterFirmControlAccountIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsMasterFirmControlAccount() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsMasterFirmControlAccount) + sizeof(IsMasterFirmControlAccount))
		return 0;

	return IsMasterFirmControlAccount;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMinimumRequiredAccountValueIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, MinimumRequiredAccountValueIsSet) + sizeof(MinimumRequiredAccountValueIsSet))
		return 0;

	return MinimumRequiredAccountValueIsSet;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataUpdate::GetMinimumRequiredAccountValue() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, MinimumRequiredAccountValue) + sizeof(MinimumRequiredAccountValue))
		return 0;

	return MinimumRequiredAccountValue;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetBeginTimeForDayMarginIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, BeginTimeForDayMarginIsSet) + sizeof(BeginTimeForDayMarginIsSet))
		return 0;

	return BeginTimeForDayMarginIsSet;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataUpdate::GetBeginTimeForDayMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, BeginTimeForDayMargin) + sizeof(BeginTimeForDayMargin))
		return 0;

	return BeginTimeForDayMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetEndTimeForDayMarginIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, EndTimeForDayMarginIsSet) + sizeof(EndTimeForDayMarginIsSet))
		return 0;

	return EndTimeForDayMarginIsSet;
}

/*==========================================================================*/
int64_t DTC_VLS::s_TradeAccountDataUpdate::GetEndTimeForDayMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, EndTimeForDayMargin) + sizeof(EndTimeForDayMargin))
		return 0;

	return EndTimeForDayMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetDayMarginTimeZoneIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, DayMarginTimeZoneIsSet) + sizeof(DayMarginTimeZoneIsSet))
		return 0;

	return DayMarginTimeZoneIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet) + sizeof(m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet))
		return 0;

	return m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday) + sizeof(m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday))
		return 0;

	return m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet) + sizeof(m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet))
		return 0;

	return m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay) + sizeof(m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay))
		return 0;

	return m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_SymbolLimitsArrayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_SymbolLimitsArrayIsSet) + sizeof(m_UseMasterFirm_SymbolLimitsArrayIsSet))
		return 0;

	return m_UseMasterFirm_SymbolLimitsArrayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_SymbolLimitsArray() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_SymbolLimitsArray) + sizeof(m_UseMasterFirm_SymbolLimitsArray))
		return 0;

	return m_UseMasterFirm_SymbolLimitsArray;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradeFeesIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradeFeesIsSet) + sizeof(m_UseMasterFirm_TradeFeesIsSet))
		return 0;

	return m_UseMasterFirm_TradeFeesIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradeFees() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradeFees) + sizeof(m_UseMasterFirm_TradeFees))
		return 0;

	return m_UseMasterFirm_TradeFees;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradeFeePerShareIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradeFeePerShareIsSet) + sizeof(m_UseMasterFirm_TradeFeePerShareIsSet))
		return 0;

	return m_UseMasterFirm_TradeFeePerShareIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradeFeePerShare() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradeFeePerShare) + sizeof(m_UseMasterFirm_TradeFeePerShare))
		return 0;

	return m_UseMasterFirm_TradeFeePerShare;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet) + sizeof(m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet))
		return 0;

	return m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_RequireSufficientMarginForNewPositions() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_RequireSufficientMarginForNewPositions) + sizeof(m_UseMasterFirm_RequireSufficientMarginForNewPositions))
		return 0;

	return m_UseMasterFirm_RequireSufficientMarginForNewPositions;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMarginIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMarginIsSet) + sizeof(m_UseMasterFirm_UsePercentOfMarginIsSet))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMargin) + sizeof(m_UseMasterFirm_UsePercentOfMargin))
		return 0;

	return m_UseMasterFirm_UsePercentOfMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet) + sizeof(m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet))
		return 0;

	return m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage) + sizeof(m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage))
		return 0;

	return m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_MinimumRequiredAccountValueIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_MinimumRequiredAccountValueIsSet) + sizeof(m_UseMasterFirm_MinimumRequiredAccountValueIsSet))
		return 0;

	return m_UseMasterFirm_MinimumRequiredAccountValueIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_MinimumRequiredAccountValue() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_MinimumRequiredAccountValue) + sizeof(m_UseMasterFirm_MinimumRequiredAccountValue))
		return 0;

	return m_UseMasterFirm_MinimumRequiredAccountValue;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_MarginTimeSettingsIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_MarginTimeSettingsIsSet) + sizeof(m_UseMasterFirm_MarginTimeSettingsIsSet))
		return 0;

	return m_UseMasterFirm_MarginTimeSettingsIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_MarginTimeSettings() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_MarginTimeSettings) + sizeof(m_UseMasterFirm_MarginTimeSettings))
		return 0;

	return m_UseMasterFirm_MarginTimeSettings;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradingIsDisabledIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradingIsDisabledIsSet) + sizeof(m_UseMasterFirm_TradingIsDisabledIsSet))
		return 0;

	return m_UseMasterFirm_TradingIsDisabledIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradingIsDisabled() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradingIsDisabled) + sizeof(m_UseMasterFirm_TradingIsDisabled))
		return 0;

	return m_UseMasterFirm_TradingIsDisabled;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsTradeStatisticsPublicallySharedIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsTradeStatisticsPublicallySharedIsSet) + sizeof(IsTradeStatisticsPublicallySharedIsSet))
		return 0;

	return IsTradeStatisticsPublicallySharedIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsTradeStatisticsPublicallyShared() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsTradeStatisticsPublicallyShared) + sizeof(IsTradeStatisticsPublicallyShared))
		return 0;

	return IsTradeStatisticsPublicallyShared;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsReadOnlyFollowingRequestsAllowedIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsReadOnlyFollowingRequestsAllowedIsSet) + sizeof(IsReadOnlyFollowingRequestsAllowedIsSet))
		return 0;

	return IsReadOnlyFollowingRequestsAllowedIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsReadOnlyFollowingRequestsAllowed() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsReadOnlyFollowingRequestsAllowed) + sizeof(IsReadOnlyFollowingRequestsAllowed))
		return 0;

	return IsReadOnlyFollowingRequestsAllowed;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsTradeAccountSharingAllowedIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsTradeAccountSharingAllowedIsSet) + sizeof(IsTradeAccountSharingAllowedIsSet))
		return 0;

	return IsTradeAccountSharingAllowedIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetIsTradeAccountSharingAllowed() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, IsTradeAccountSharingAllowed) + sizeof(IsTradeAccountSharingAllowed))
		return 0;

	return IsTradeAccountSharingAllowed;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetReadOnlyShareToAllSCUsernamesIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, ReadOnlyShareToAllSCUsernamesIsSet) + sizeof(ReadOnlyShareToAllSCUsernamesIsSet))
		return 0;

	return ReadOnlyShareToAllSCUsernamesIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetReadOnlyShareToAllSCUsernames() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, ReadOnlyShareToAllSCUsernames) + sizeof(ReadOnlyShareToAllSCUsernames))
		return 0;

	return ReadOnlyShareToAllSCUsernames;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_SymbolCommissionsArrayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_SymbolCommissionsArrayIsSet) + sizeof(m_UseMasterFirm_SymbolCommissionsArrayIsSet))
		return 0;

	return m_UseMasterFirm_SymbolCommissionsArrayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_SymbolCommissionsArray() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_SymbolCommissionsArray) + sizeof(m_UseMasterFirm_SymbolCommissionsArray))
		return 0;

	return m_UseMasterFirm_SymbolCommissionsArray;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet) + sizeof(m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet))
		return 0;

	return m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit) + sizeof(m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit))
		return 0;

	return m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet) + sizeof(m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMarginForDayTrading() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMarginForDayTrading) + sizeof(m_UseMasterFirm_UsePercentOfMarginForDayTrading))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginForDayTrading;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet) + sizeof(m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet))
		return 0;

	return m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_SymbolCommissionsArrayFullOverride() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_SymbolCommissionsArrayFullOverride) + sizeof(m_UseMasterFirm_SymbolCommissionsArrayFullOverride))
		return 0;

	return m_UseMasterFirm_SymbolCommissionsArrayFullOverride;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet) + sizeof(m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet))
		return 0;

	return m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders) + sizeof(m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders))
		return 0;

	return m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet) + sizeof(m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMarginFullOverride() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMarginFullOverride) + sizeof(m_UseMasterFirm_UsePercentOfMarginFullOverride))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginFullOverride;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradeFeesFullOverrideIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradeFeesFullOverrideIsSet) + sizeof(m_UseMasterFirm_TradeFeesFullOverrideIsSet))
		return 0;

	return m_UseMasterFirm_TradeFeesFullOverrideIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_TradeFeesFullOverride() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_TradeFeesFullOverride) + sizeof(m_UseMasterFirm_TradeFeesFullOverride))
		return 0;

	return m_UseMasterFirm_TradeFeesFullOverride;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet) + sizeof(m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride) + sizeof(m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride))
		return 0;

	return m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetLiquidationOnlyModeIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_LiquidationOnlyModeIsSet) + sizeof(m_LiquidationOnlyModeIsSet))
		return 0;

	return m_LiquidationOnlyModeIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetLiquidationOnlyMode() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_LiquidationOnlyMode) + sizeof(m_LiquidationOnlyMode))
		return 0;

	return m_LiquidationOnlyMode;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetCustomerOrFirmIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_CustomerOrFirmIsSet) + sizeof(m_CustomerOrFirmIsSet))
		return 0;

	return m_CustomerOrFirmIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetCustomerOrFirm() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_CustomerOrFirm) + sizeof(m_CustomerOrFirm))
		return 0;

	return m_CustomerOrFirm;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet) + sizeof(m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet))
		return 0;

	return m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet) + sizeof(m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet))
		return 0;

	return m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelWhenUnderMarginIntraday() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelWhenUnderMarginIntraday) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMarginIntraday))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMarginIntraday;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay) + sizeof(m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay))
		return 0;

	return m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_MaximumOrderQuantityIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_MaximumOrderQuantityIsSet) + sizeof(m_MasterFirm_MaximumOrderQuantityIsSet))
		return 0;

	return m_MasterFirm_MaximumOrderQuantityIsSet;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataUpdate::GetMasterFirm_MaximumOrderQuantity() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_MasterFirm_MaximumOrderQuantity) + sizeof(m_MasterFirm_MaximumOrderQuantity))
		return 0;

	return m_MasterFirm_MaximumOrderQuantity;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdate::GetExchangeTraderIdIsSet() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdate, m_ExchangeTraderIdIsSet) + sizeof(m_ExchangeTraderIdIsSet))
		return 0;

	return m_ExchangeTraderIdIsSet;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataDelete

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataDelete::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataDelete::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataDelete::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataDelete, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataSymbolLimits

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetTradePositionLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsResponse, TradePositionLimit) + sizeof(TradePositionLimit))
		return 0;

	return TradePositionLimit;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetOrderQuantityLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsResponse, OrderQuantityLimit) + sizeof(OrderQuantityLimit))
		return 0;

	return OrderQuantityLimit;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetUsePercentOfMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsResponse, UsePercentOfMargin) + sizeof(UsePercentOfMargin))
		return 0;

	return UsePercentOfMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetOverrideMarginOtherAccounts() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsResponse, OverrideMarginOtherAccounts) + sizeof(OverrideMarginOtherAccounts))
		return 0;

	return OverrideMarginOtherAccounts;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetUsePercentOfMarginForDayTrading() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsResponse, UsePercentOfMarginForDayTrading) + sizeof(UsePercentOfMarginForDayTrading))
		return 0;

	return UsePercentOfMarginForDayTrading;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataSymbolLimitsResponse::GetNumberOfDaysBeforeLastTradingDateToDisallowOrders() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsResponse, NumberOfDaysBeforeLastTradingDateToDisallowOrders) + sizeof(NumberOfDaysBeforeLastTradingDateToDisallowOrders))
		return 0;

	return NumberOfDaysBeforeLastTradingDateToDisallowOrders;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataSymbolLimitsUpdate

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetIsDeleted() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, IsDeleted) + sizeof(IsDeleted))
		return 0;

	return IsDeleted;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetTradePositionLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, TradePositionLimit) + sizeof(TradePositionLimit))
		return 0;

	return TradePositionLimit;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetOrderQuantityLimit() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, OrderQuantityLimit) + sizeof(OrderQuantityLimit))
		return 0;

	return OrderQuantityLimit;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetUsePercentOfMargin() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, UsePercentOfMargin) + sizeof(UsePercentOfMargin))
		return 0;

	return UsePercentOfMargin;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetOverrideMarginOtherAccounts() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, OverrideMarginOtherAccounts) + sizeof(OverrideMarginOtherAccounts))
		return 0;

	return OverrideMarginOtherAccounts;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetUsePercentOfMarginForDayTrading() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, UsePercentOfMarginForDayTrading) + sizeof(UsePercentOfMarginForDayTrading))
		return 0;

	return UsePercentOfMarginForDayTrading;
}

/*==========================================================================*/
int32_t DTC_VLS::s_TradeAccountDataSymbolLimitsUpdate::GetNumberOfDaysBeforeLastTradingDateToDisallowOrders() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolLimitsUpdate, NumberOfDaysBeforeLastTradingDateToDisallowOrders) + sizeof(NumberOfDaysBeforeLastTradingDateToDisallowOrders))
		return 0;

	return NumberOfDaysBeforeLastTradingDateToDisallowOrders;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataSymbolCommissionResponse

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolCommissionResponse::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolCommissionResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataSymbolCommissionResponse::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolCommissionResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataSymbolCommissionResponse::GetTradeFeePerContract() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolCommissionResponse, TradeFeePerContract) + sizeof(TradeFeePerContract))
		return 0;

	return TradeFeePerContract;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataSymbolCommissionUpdate

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolCommissionUpdate::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataSymbolCommissionUpdate::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataSymbolCommissionUpdate::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolCommissionUpdate, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataSymbolCommissionUpdate::GetIsDeleted() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolCommissionUpdate, IsDeleted) + sizeof(IsDeleted))
		return 0;

	return IsDeleted;
}

/*==========================================================================*/
double DTC_VLS::s_TradeAccountDataSymbolCommissionUpdate::GetTradeFeePerContract() const
{
	if (BaseSize < offsetof(s_TradeAccountDataSymbolCommissionUpdate, TradeFeePerContract) + sizeof(TradeFeePerContract))
		return 0;

	return TradeFeePerContract;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataAuthorizedUsernameResponse

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameResponse::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameResponse::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataAuthorizedUsernameResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameResponse::GetUpdateOperationComplete() const
{
	if (BaseSize < offsetof(s_TradeAccountDataAuthorizedUsernameResponse, UpdateOperationComplete) + sizeof(UpdateOperationComplete))
		return 0;

	return UpdateOperationComplete;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameResponse::GetUpdateOperationMessageType() const
{
	if (BaseSize < offsetof(s_TradeAccountDataAuthorizedUsernameResponse, UpdateOperationMessageType) + sizeof(UpdateOperationMessageType))
		return 0;

	return UpdateOperationMessageType;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataAuthorizedUsernameAdd

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameAdd::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameAdd::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameAdd::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataAuthorizedUsernameAdd, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataAuthorizedUsernameRemove

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameRemove::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameRemove::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataAuthorizedUsernameRemove::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataAuthorizedUsernameRemove, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataUsernameToShareWithResponse

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUsernameToShareWithResponse::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUsernameToShareWithResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataUsernameToShareWithResponse::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUsernameToShareWithResponse::GetIsReadOnly() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithResponse, IsReadOnly) + sizeof(IsReadOnly))
		return 0;

	return IsReadOnly;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUsernameToShareWithResponse::GetUpdateOperationComplete() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithResponse, UpdateOperationComplete) + sizeof(UpdateOperationComplete))
		return 0;

	return UpdateOperationComplete;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUsernameToShareWithResponse::GetUpdateOperationMessageType() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithResponse, UpdateOperationMessageType) + sizeof(UpdateOperationMessageType))
		return 0;

	return UpdateOperationMessageType;
}

/*==========================================================================*/

/****************************************************************************/
// s_TradeAccountDataUsernameToShareWithAdd

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUsernameToShareWithAdd::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUsernameToShareWithAdd::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataUsernameToShareWithAdd::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithAdd, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUsernameToShareWithAdd::GetIsReadOnly() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithAdd, IsReadOnly) + sizeof(IsReadOnly))
		return 0;

	return IsReadOnly;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataUsernameToShareWithRemove

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUsernameToShareWithRemove::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUsernameToShareWithRemove::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataUsernameToShareWithRemove::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithRemove, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUsernameToShareWithRemove::GetIsReadOnly() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUsernameToShareWithRemove, IsReadOnly) + sizeof(IsReadOnly))
		return 0;

	return IsReadOnly;
}

/*==========================================================================*/


/****************************************************************************/
//s_TradeAccountDataResponseTrailer

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataResponseTrailer::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataResponseTrailer::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataResponseTrailer::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponseTrailer, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponseTrailer::GetIsSnapshot() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponseTrailer, m_IsSnapshot) + sizeof(m_IsSnapshot))
		return 0;

	return m_IsSnapshot;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponseTrailer::GetIsFirstMessageInBatch() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponseTrailer, m_IsFirstMessageInBatch) + sizeof(m_IsFirstMessageInBatch))
		return 0;

	return m_IsFirstMessageInBatch;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataResponseTrailer::GetIsLastMessageInBatch() const
{
	if (BaseSize < offsetof(s_TradeAccountDataResponseTrailer, m_IsLastMessageInBatch) + sizeof(m_IsLastMessageInBatch))
		return 0;

	return m_IsLastMessageInBatch;
}

/*==========================================================================*/

/****************************************************************************/
//s_TradeAccountDataUpdateOperationComplete

/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUpdateOperationComplete::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_TradeAccountDataUpdateOperationComplete::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_TradeAccountDataUpdateOperationComplete::GetRequestID() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdateOperationComplete, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdateOperationComplete::GetIsError() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdateOperationComplete, IsError) + sizeof(IsError))
		return 0;

	return IsError;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdateOperationComplete::GetIsDeleteOperation() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdateOperationComplete, IsDeleteOperation) + sizeof(IsDeleteOperation))
		return 0;

	return IsDeleteOperation;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdateOperationComplete::GetIsSymbolLimitsUpdateOperation() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdateOperationComplete, IsSymbolLimitsUpdateOperation) + sizeof(IsSymbolLimitsUpdateOperation))
		return 0;

	return IsSymbolLimitsUpdateOperation;
}

/*==========================================================================*/
uint8_t DTC_VLS::s_TradeAccountDataUpdateOperationComplete::GetIsSymbolCommissionUpdateOperation() const
{
	if (BaseSize < offsetof(s_TradeAccountDataUpdateOperationComplete, IsSymbolCommissionUpdateOperation) + sizeof(IsSymbolCommissionUpdateOperation))
		return 0;

	return IsSymbolCommissionUpdateOperation;
}

/*==========================================================================*/


/****************************************************************************/
//s_ProcessedFillIdentifier

/*==========================================================================*/
uint16_t DTC_VLS::s_ProcessedFillIdentifier::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_ProcessedFillIdentifier::GetBaseSize() const
{
	return BaseSize;
}
/*==========================================================================*/

/****************************************************************************/
//s_FlattenPositionsForTradeAccount

/*==========================================================================*/
uint16_t DTC_VLS::s_FlattenPositionsForTradeAccount::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_FlattenPositionsForTradeAccount::GetBaseSize() const
{
	return BaseSize;
}
/*==========================================================================*/

/****************************************************************************/
//s_UserInformation

/*==========================================================================*/
uint16_t DTC_VLS::s_UserInformation::GetMessageSize() const
{
	return Size;
}
/*==========================================================================*/
uint16_t DTC_VLS::s_UserInformation::GetBaseSize() const
{
	return BaseSize;
}
/*==========================================================================*/

/****************************************************************************/
//s_WriteIntradayDataFileSessionValue

/*==========================================================================*/
uint16_t DTC::s_WriteIntradayDataFileSessionValue::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
void DTC::s_WriteIntradayDataFileSessionValue::CopyFrom(void* p_SourceData)
{
	memcpy(this, p_SourceData, min(sizeof(s_WriteIntradayDataFileSessionValue), *static_cast<uint16_t*>(p_SourceData)));
}

/*==========================================================================*/
uint32_t DTC::s_WriteIntradayDataFileSessionValue::GetSymbolID() const
{
	if (Size < offsetof(s_WriteIntradayDataFileSessionValue, SymbolID) + sizeof(SymbolID))
		return 0;

	return SymbolID;
}

/*==========================================================================*/
n_DTCNonStandard::WriteIntradayDataFileSessionValueTypeEnum DTC::s_WriteIntradayDataFileSessionValue::GetValueType() const
{
	if (Size < offsetof(s_WriteIntradayDataFileSessionValue, ValueType) + sizeof(ValueType))
		return n_DTCNonStandard::INTRADAY_DATA_FILE_SESSION_VALUE_UNSET;

	return ValueType;
}

/*==========================================================================*/
DTC::t_DateTimeWithMicrosecondsInt DTC::s_WriteIntradayDataFileSessionValue::GetDateTime() const
{
	if (Size < offsetof(s_WriteIntradayDataFileSessionValue, DateTime) + sizeof(DateTime))
		return 0;

	return DateTime;
}

/*==========================================================================*/
DTC::t_DateTime DTC::s_WriteIntradayDataFileSessionValue::GetDate() const
{
	if (Size < offsetof(s_WriteIntradayDataFileSessionValue, Date) + sizeof(Date))
		return 0;

	return Date;
}

/*==========================================================================*/
double DTC::s_WriteIntradayDataFileSessionValue::GetPrice() const
{
	if (Size < offsetof(s_WriteIntradayDataFileSessionValue, Price) + sizeof(Price))
		return 0;

	return Price;
}

/*==========================================================================*/
double DTC::s_WriteIntradayDataFileSessionValue::GetVolume() const
{
	if (Size < offsetof(s_WriteIntradayDataFileSessionValue, Volume) + sizeof(Volume))
		return 0;

	return Volume;
}

/****************************************************************************/
//s_MarginDataRequest

/*==========================================================================*/
uint16_t DTC_VLS::s_MarginDataRequest::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_MarginDataRequest::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_MarginDataRequest::GetRequestID() const
{
	if (BaseSize < offsetof(s_MarginDataRequest, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/****************************************************************************/
//s_MarginDataResponse

/*==========================================================================*/
uint16_t DTC_VLS::s_MarginDataResponse::GetMessageSize() const
{
	return Size;
}

/*==========================================================================*/
uint16_t DTC_VLS::s_MarginDataResponse::GetBaseSize() const
{
	return BaseSize;
}

/*==========================================================================*/
uint32_t DTC_VLS::s_MarginDataResponse::GetRequestID() const
{
	if (Size < offsetof(s_MarginDataResponse, RequestID) + sizeof(RequestID))
		return 0;

	return RequestID;
}

/*==========================================================================*/
double DTC_VLS::s_MarginDataResponse::GetInitialExchangeMargin() const
{
	if (Size < offsetof(s_MarginDataResponse, InitialExchangeMargin) + sizeof(InitialExchangeMargin))
		return 0;

	return InitialExchangeMargin;
}

/*==========================================================================*/
double DTC_VLS::s_MarginDataResponse::GetMaintenanceExchangeMargin() const
{
	if (Size < offsetof(s_MarginDataResponse, MaintenanceExchangeMargin) + sizeof(MaintenanceExchangeMargin))
		return 0;

	return MaintenanceExchangeMargin;
}

/*==========================================================================*/
double DTC_VLS::s_MarginDataResponse::GetInitialAccountMargin() const
{
	if (Size < offsetof(s_MarginDataResponse, InitialAccountMargin) + sizeof(InitialAccountMargin))
		return 0;

	return InitialAccountMargin;
}

/*==========================================================================*/
double DTC_VLS::s_MarginDataResponse::GetMaintenanceAccountMargin() const
{
	if (Size < offsetof(s_MarginDataResponse, MaintenanceAccountMargin) + sizeof(MaintenanceAccountMargin))
		return 0;

	return MaintenanceAccountMargin;
}
