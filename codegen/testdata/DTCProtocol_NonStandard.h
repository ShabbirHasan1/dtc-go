#pragma once

#include "DTCProtocolVLS.h"

// Contains both binary and binary with variable length string messages
namespace n_DTCNonStandard
{

	/*==========================================================================*/
	//Nonstandard messages.
	const uint16_t HISTORICAL_TRADES_REQUEST = 10100;
	const uint16_t HISTORICAL_TRADES_REJECT = 10101;
	const uint16_t HISTORICAL_TRADES_RESPONSE = 10102;
	const uint16_t REPLAY_CHART_DATA = 10104;
	const uint16_t REPLAY_CHART_DATA_PERFORM_ACTION = 10105;
	const uint16_t REPLAY_CHART_DATA_STATUS = 10106;
	const uint16_t REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS = 10107;
	const uint16_t NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE = 10108;

	const uint16_t SC_CONFIGURATION_SYNCHRONIZATION = 10109;
	const uint16_t SC_TRADE_ORDER = 10110;
	const uint16_t INDIVIDUAL_TRADE_POSITION = 10112;
	const uint16_t TRADE_POSITION_CONSOLIDATED = 10113;
	const uint16_t TRADE_ACTIVITY_DATA = 10114;

	const uint16_t TRADE_ACCOUNT_DATA_REQUEST = 10115;// Requests all trade account data including sharing data.

	const uint16_t TRADE_ACCOUNT_DATA_RESPONSE = 10116;
	const uint16_t TRADE_ACCOUNT_DATA_UPDATE = 10117;
	const uint16_t TRADE_ACCOUNT_DATA_DELETE = 10118;

	const uint16_t TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_RESPONSE = 10119;
	const uint16_t TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_UPDATE = 10120;

	const uint16_t TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_RESPONSE = 10121;
	const uint16_t TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE = 10122;

	const uint16_t TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_RESPONSE = 10124;
	const uint16_t TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_ADD = 10125;
	const uint16_t TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE = 10126;

	const uint16_t TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_RESPONSE = 10127;
	const uint16_t TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD = 10128;
	const uint16_t TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE = 10129;

	const uint16_t TRADE_ACCOUNT_DATA_RESPONSE_TRAILER = 10130;// m_ inconsistencies
	const uint16_t TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE = 10131;

	const uint16_t PROCESSED_FILL_IDENTIFIER = 10132;
	const uint16_t INTERPROCESS_SYNCHRONIZATION_SNAPSHOT_REQUEST = 10133;

	const uint16_t INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE = 10134;
	const uint16_t FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT = 10135;

	const uint16_t USER_INFORMATION = 10136;

	const uint16_t INTERPROCESS_SYNCHRONIZATION_TRADE_ACTIVITY_REQUEST = 10137;
	const uint16_t DOWNLOAD_HISTORICAL_ORDER_FILL_AND_ACCOUNT_BALANCE_DATA = 10138;
	const uint16_t CLIENT_DEVICE_UPDATE = 10139;
	const uint16_t WRITE_INTRADAY_DATA_FILE_SESSION_VALUE = 10140;

	const uint16_t MARGIN_DATA_REQUEST = 10141;
	const uint16_t MARGIN_DATA_RESPONSE = 10142;

	enum ReplayChartDataActionEnum : int32_t
	{
		REPLAY_CHART_DATA_ACTION_NONE = 0
		, REPLAY_CHART_DATA_ACTION_STOP = 1
		, REPLAY_CHART_DATA_ACTION_PAUSE = 2
		, REPLAY_CHART_DATA_ACTION_RESUME = 3
		, REPLAY_CHART_DATA_ACTION_FINISH = 4
		, REPLAY_CHART_DATA_ACTION_CHANGE_SPEED = 5
	};

	enum ReplayChartDataStatusEnum : int32_t
	{
		REPLAY_CHART_DATA_STATUS_UNSET = 0
		, REPLAY_CHART_DATA_STATUS_STARTED = 1
		, REPLAY_CHART_DATA_STATUS_ERROR = 2
		, REPLAY_CHART_DATA_STATUS_COMPLETE = 3
	};

	enum UseCompressionEnum : int32_t
	{
		USE_COMPRESSION_DISABLED = 0
		, USE_COMPRESSION_BLOCK_COMPRESSION = 1
		, USE_COMPRESSION_STREAMING_COMPRESSION = 2
	};

	enum WriteIntradayDataFileSessionValueTypeEnum : int32_t
	{
		INTRADAY_DATA_FILE_SESSION_VALUE_UNSET = 0
		, INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_OPEN = 1
		, INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_HIGH
		, INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_LOW
		, INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_VOLUME
		, INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_SETTLEMENT_PRICE
		, INTRADAY_DATA_FILE_SESSION_VALUE_OPEN_INTEREST
		, INTRADAY_DATA_FILE_SESSION_VALUE_ODD_LOT_TRADE
		, INTRADAY_DATA_FILE_SESSION_VALUE_NON_LAST_UPDATE_EQUITY_TRADE
	};
}


/*==========================================================================*/
namespace DTC
{
#pragma pack(push, 1)

	/*==========================================================================*/
	struct s_EncodingRequestExtended
	{
		uint16_t Size = 0;
		uint16_t Type = 0;
		int32_t ProtocolVersion = DTC::CURRENT_VERSION;
		DTC::EncodingEnum Encoding = DTC::BINARY_ENCODING;
		char ProtocolType[4];
		n_DTCNonStandard::UseCompressionEnum UseCompression = n_DTCNonStandard::USE_COMPRESSION_DISABLED;

		s_EncodingRequestExtended()
		{
			Size = sizeof(*this);
			Type = DTC::ENCODING_REQUEST;
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);
		int32_t GetProtocolVersion() const;
		DTC::EncodingEnum GetEncoding() const;
		const char* GetProtocolType();
		void SetProtocolType(const char* NewValue);
		n_DTCNonStandard::UseCompressionEnum GetUseCompression() const;
	};

	/*==========================================================================*/
	struct s_HeartbeatExtended
	{
		uint16_t Size = 0;
		uint16_t Type = 0;

		uint32_t NumDroppedMessages = 0;
		DTC::t_DateTime CurrentDateTime = 0;

		uint16_t SecondsSinceLastReceivedHeartbeat = 0;
		uint16_t NumberOfOutstandingSentBuffers = 0;
		uint16_t PendingTransmissionDelayInMilliseconds = 0;
		uint32_t CurrentSendBufferSizeInBytes = 0;

		DTC::t_DateTimeWithMicrosecondsInt SendingDateTimeMicroseconds = 0;
		float DataCompressionRatio = 0;
		uint64_t TotalUncompressedBytes = 0;
		double TotalCompressionTime = 0;//In seconds
		uint32_t NumberOfCompressions = 0;
		uint32_t SourceIPAddress = 0;

		s_HeartbeatExtended()
		{
			Size = sizeof(*this);
			Type = DTC::HEARTBEAT;
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetNumDroppedMessages() const;
		DTC::t_DateTime GetCurrentDateTime() const;

		uint16_t GetSecondsSinceLastReceivedHeartbeat() const;
		uint16_t GetNumberOfOutstandingSentBuffers() const;
		uint16_t GetPendingTransmissionDelayInMilliseconds() const;
		uint32_t GetCurrentSendBufferSizeInBytes() const;

		DTC::t_DateTimeWithMicrosecondsInt GetSendingDateTimeMicroseconds() const;

		float GetDataCompressionRatio() const;
		uint64_t GetTotalUncompressedBytes() const;
		double GetTotalCompressionTime() const;
		uint32_t GetNumberOfCompressions() const;
		uint32_t GetSourceIPAddress() const;
	};

	/*==========================================================================*/
	struct s_HistoricalTradesRequest
	{
		uint16_t Size;
		uint16_t Type;

		int32_t RequestID;
		char Symbol[SYMBOL_LENGTH];
		char TradeAccount[TRADE_ACCOUNT_LENGTH];
		DTC::t_DateTime StartDateTime;
		uint32_t SubAccountIdentifier;
		uint8_t CreateFlatToFlatTrades;

		s_HistoricalTradesRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalTradesRequest));
			Size = sizeof(s_HistoricalTradesRequest);
			Type = n_DTCNonStandard::HISTORICAL_TRADES_REQUEST;
		}

		int32_t GetRequestID() const;

		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);

		DTC::t_DateTime GetStartDateTime() const;
		uint32_t GetSubAccountIdentifier() const;
		uint8_t GetCreateFlatToFlatTrades() const;
	};

	/*==========================================================================*/
	struct s_HistoricalTradesReject
	{
		uint16_t Size;
		uint16_t Type;

		int32_t RequestID;
		char RejectText[TEXT_DESCRIPTION_LENGTH];

		s_HistoricalTradesReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalTradesReject));
			Size = sizeof(s_HistoricalTradesReject);
			Type = n_DTCNonStandard::HISTORICAL_TRADES_REJECT;
		}

		uint32_t GetRequestID() const;
		const char* GetRejectText();
		void SetRejectText(const char* NewValue);
	};

	/*==========================================================================*/
	struct s_HistoricalTradesResponse
	{
		uint16_t Size;
		uint16_t Type;

		int32_t RequestID;
		uint8_t IsFinalMessage;
		char Symbol[SYMBOL_LENGTH];
		char TradeAccount[TRADE_ACCOUNT_LENGTH];
		DTC::t_DateTimeWithMilliseconds EntryDateTime;
		DTC::t_DateTimeWithMilliseconds ExitDateTime;
		double EntryPrice;
		double ExitPrice;
		DTC::BuySellEnum TradeType;
		uint32_t EntryQuantity;
		uint32_t ExitQuantity;
		uint32_t MaxOpenQuantity;
		double ClosedProfitLoss;
		double MaximumOpenPositionLoss;
		double MaximumOpenPositionProfit;
		double Commission;
		char OpenFillExecutionID[ORDER_FILL_EXECUTION_LENGTH];
		char CloseFillExecutionID[ORDER_FILL_EXECUTION_LENGTH];
		uint32_t SubAccountIdentifier;

		s_HistoricalTradesResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalTradesResponse));
			Size = sizeof(s_HistoricalTradesResponse);
			Type = n_DTCNonStandard::HISTORICAL_TRADES_RESPONSE;

			TradeType = DTC::BUY_SELL_UNSET;
		}

		int32_t GetRequestID() const;
		uint8_t GetIsFinalMessage() const;
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
		DTC::t_DateTimeWithMilliseconds GetEntryDateTime() const;
		DTC::t_DateTimeWithMilliseconds GetExitDateTime() const;
		double GetEntryPrice() const;
		double GetExitPrice() const;
		DTC::BuySellEnum GetTradeType() const;
		uint32_t GetEntryQuantity() const;
		uint32_t GetExitQuantity() const;
		uint32_t GetMaxOpenQuantity() const;
		double GetClosedProfitLoss() const;
		double GetMaximumOpenPositionLoss() const;
		double GetMaximumOpenPositionProfit() const;
		double GetCommission() const;
		const char* GetOpenFillExecutionID();
		void SetOpenFillExecutionID(const char* NewValue);
		const char* GetCloseFillExecutionID();
		void SetCloseFillExecutionID(const char* NewValue);
		uint32_t GetSubAccountIdentifier() const;
	};

	/*==========================================================================*/
	struct s_ReplayChartData
	{
		uint16_t Size;
		uint16_t Type;

		uint32_t RequestID;
		char Symbol[SYMBOL_LENGTH];
		char TradeAccount[TRADE_ACCOUNT_LENGTH];
		char TimeZone[TEXT_MESSAGE_LENGTH];
		DTC::t_DateTimeWithMillisecondsInt StartDateTimeForInitialData;
		DTC::t_DateTimeWithMillisecondsInt StartDateTime;
		DTC::t_DateTimeWithMillisecondsInt StopDateTime;
		uint16_t SessionBeginTimeInSeconds;
		uint16_t SessionEndTimeInSeconds;
		float ReplaySpeed;
		int32_t BarTimeInSeconds;
		uint8_t PauseReplayAfterInitialDataSent;
		uint8_t UseSavedPriorState;
		float SymbolVolatility;
		float InterestRate;
		int32_t NumberOfOrdersToTriggerFinishToStopDateTime;
		int32_t MaximumNumberOfOrdersPerReplaySession;
		int32_t NumberOfDaysForInitialDataFromBeforeLastSavedDateTime;
		uint32_t SubAccountIdentifier = 0;
		int32_t OptionsPriceSendIntervalInSeconds = 0;

		s_ReplayChartData()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(s_ReplayChartData));
			Size = sizeof(s_ReplayChartData);
			Type = n_DTCNonStandard::REPLAY_CHART_DATA;

			ReplaySpeed = 1.0;
		}

		uint32_t GetRequestID() const;
		const char* GetSymbol();
		void SetSymbol(const char* NewValue);
		const char* GetTradeAccount();
		void SetTradeAccount(const char* NewValue);
		const char* GetTimeZone();
		void SetTimeZone(const char* NewValue);
		DTC::t_DateTimeWithMillisecondsInt GetStartDateTimeForInitialData() const;
		DTC::t_DateTimeWithMillisecondsInt GetStartDateTime() const;
		DTC::t_DateTimeWithMillisecondsInt GetStopDateTime() const;
		uint16_t GetSessionBeginTimeInSeconds() const;
		uint16_t GetSessionEndTimeInSeconds() const;
		float GetReplaySpeed() const;
		int32_t GetBarTimeInSeconds() const;
		uint8_t GetPauseReplayAfterInitialDataSent() const;
		uint8_t GetUseSavedPriorState() const;
		float GetSymbolVolatility() const;
		float GetInterestRate() const;
		int32_t GetNumberOfOrdersToTriggerFinishToStopDateTime() const;
		int32_t GetMaximumNumberOfOrdersPerReplaySession() const;
		int32_t GetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime() const;
		uint32_t GetSubAccountIdentifier() const;
		int32_t GetOptionsPriceSendIntervalInSeconds() const;
	};

	/*==========================================================================*/
	struct s_ReplayChartDataPerformAction
	{
		uint16_t Size;
		uint16_t Type;

		uint32_t RequestID = 0;
		n_DTCNonStandard::ReplayChartDataActionEnum Action = n_DTCNonStandard::REPLAY_CHART_DATA_ACTION_NONE;
		float ReplaySpeed = 0;

		s_ReplayChartDataPerformAction()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(s_ReplayChartDataPerformAction));
			Size = sizeof(s_ReplayChartDataPerformAction);
			Type = n_DTCNonStandard::REPLAY_CHART_DATA_PERFORM_ACTION;
		}

		uint32_t GetRequestID() const;
		n_DTCNonStandard::ReplayChartDataActionEnum GetAction() const;
		float GetReplaySpeed() const;
	};

	/*==========================================================================*/
	struct s_ReplayChartDataStatus
	{
		uint16_t Size;
		uint16_t Type;

		uint32_t RequestID = 0;
		char ErrorMessage[TEXT_MESSAGE_LENGTH];
		n_DTCNonStandard::ReplayChartDataStatusEnum Status = n_DTCNonStandard::REPLAY_CHART_DATA_STATUS_UNSET;
		uint32_t  SubAccountIdentifier = 0;

		s_ReplayChartDataStatus()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(s_ReplayChartDataStatus));
			Size = sizeof(s_ReplayChartDataStatus);
			Type = n_DTCNonStandard::REPLAY_CHART_DATA_STATUS;
		}

		uint32_t GetRequestID() const;
		const char* GetErrorMessage();
		void SetErrorMessage(const char* NewValue);
		n_DTCNonStandard::ReplayChartDataStatusEnum GetStatus() const;
		uint32_t GetSubAccountIdentifier() const;
	};

	/*==========================================================================*/
	struct s_RequestNumCurrentClientConnections
	{
		uint16_t Size;
		uint16_t Type;

		uint32_t RequestID = 0;
		char Username[USERNAME_PASSWORD_LENGTH];
		int64_t DeviceIdentifier = 0;

		s_RequestNumCurrentClientConnections()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);
		void Clear()
		{
			memset(this, 0, sizeof(s_RequestNumCurrentClientConnections));
			Size = sizeof(s_RequestNumCurrentClientConnections);
			Type = n_DTCNonStandard::REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS;
		}

		uint32_t GetRequestID() const;
		const char* GetUsername();
		void SetUsername(const char* NewValue);
		int64_t GetDeviceIdentifier() const;
	};

	/*==========================================================================*/
	struct s_NumCurrentClientConnectionsResponse
	{
		uint16_t Size = sizeof(s_NumCurrentClientConnectionsResponse);
		uint16_t Type = n_DTCNonStandard::NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE;

		uint32_t RequestID = 0;
		char Username[USERNAME_PASSWORD_LENGTH] = {0};
		int32_t NumConnectionsForDifferentDevices = 0;
		int32_t NumConnectionsForSameDevice = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void * p_SourceData);

		uint32_t GetRequestID() const;
		const char* GetUsername();
		void SetUsername(const char* NewValue);
		int32_t GetNumConnectionsForDifferentDevices() const;
		int32_t GetNumConnectionsForSameDevice() const;
	};

	/*==========================================================================*/
	struct s_ClientDeviceUpdate
	{
		uint16_t Size = sizeof(s_ClientDeviceUpdate);
		uint16_t Type = n_DTCNonStandard::CLIENT_DEVICE_UPDATE;

		int64_t ClientDeviceIdentifier = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		int64_t GetClientDeviceIdentifier() const;
	
	};

	struct s_InterprocessSynchronizationRemoteState
	{
		uint16_t Size = sizeof(s_InterprocessSynchronizationRemoteState);
		uint16_t Type = n_DTCNonStandard::INTERPROCESS_SYNCHRONIZATION_REMOTE_STATE;

		uint8_t IsPrimary = 0;
		uint8_t IsSecondary = 0;
		uint8_t IsSecondaryFailoverActive = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint8_t GetIsPrimary() const;
		uint8_t GetIsSecondary() const;
		uint8_t GetIsSecondaryFailoverActive() const;
	};

	struct s_InterprocessSynchronizationSnapshotRequest
	{
		uint16_t Size = sizeof(s_InterprocessSynchronizationSnapshotRequest);
		uint16_t Type = n_DTCNonStandard::INTERPROCESS_SYNCHRONIZATION_SNAPSHOT_REQUEST;

		uint32_t RequestID = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetRequestID() const;
	};

	struct s_InterprocessSynchronizationTradeActivityRequest
	{
		uint16_t Size = sizeof(s_InterprocessSynchronizationTradeActivityRequest);
		uint16_t Type = n_DTCNonStandard::INTERPROCESS_SYNCHRONIZATION_TRADE_ACTIVITY_REQUEST;

		uint32_t RequestID = 0;
		int64_t StartDateTimeUTC = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetRequestID() const;
		int64_t GetStartDateTimeUTC() const;
	};

	/*==========================================================================*/
	struct s_WriteIntradayDataFileSessionValue
	{
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::WRITE_INTRADAY_DATA_FILE_SESSION_VALUE;

		uint32_t SymbolID = 0;
		n_DTCNonStandard::WriteIntradayDataFileSessionValueTypeEnum ValueType = n_DTCNonStandard::INTRADAY_DATA_FILE_SESSION_VALUE_UNSET;
		DTC::t_DateTimeWithMicrosecondsInt DateTime = 0;
		DTC::t_DateTime Date = 0;
		double Price = 0;
		double Volume = 0;

		uint16_t GetMessageSize() const;
		void CopyFrom(void* p_SourceData);

		uint32_t GetSymbolID() const;
		n_DTCNonStandard::WriteIntradayDataFileSessionValueTypeEnum GetValueType() const;
		DTC::t_DateTimeWithMicrosecondsInt GetDateTime() const;
		DTC::t_DateTime GetDate() const;
		double GetPrice() const;
		double GetVolume() const;
	};

#pragma pack(pop)
}

/*==========================================================================*/
namespace DTC_VLS
{
#pragma pack(push, 1)

	/*==========================================================================*/
	struct s_SCConfigurationSynchronization
	{
		uint16_t Size = sizeof(s_SCConfigurationSynchronization);
		uint16_t Type = n_DTCNonStandard::SC_CONFIGURATION_SYNCHRONIZATION;
		uint16_t BaseSize = Size;

		uint32_t MessageID = 0;
		uint32_t CurrentInboundSequenceNumber = 0;
		uint32_t CurrentOutboundSequenceNumber = 0;
		uint64_t CurrentInternalOrderID = 0;
		uint8_t IsSnapshot = 0;
		DTC_VLS::vls_t LastReceivedExecutionIdentifier;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetMessageID() const;
		uint32_t GetCurrentInboundSequenceNumber() const;
		uint32_t GetCurrentOutboundSequenceNumber() const;
		uint64_t GetCurrentInternalOrderID() const;
		uint8_t GetIsSnapshot() const;

		const char* GetLastReceivedExecutionIdentifier() const
		{
			return GetVariableLengthStringField(Size, BaseSize, LastReceivedExecutionIdentifier, offsetof(s_SCConfigurationSynchronization, LastReceivedExecutionIdentifier));
		}

		void AddLastReceivedExecutionIdentifier(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, LastReceivedExecutionIdentifier, StringLength);
		}
	};

	struct s_DownloadHistoricalOrderFillAndAccountBalanceData
	{
		uint16_t Size = sizeof(s_DownloadHistoricalOrderFillAndAccountBalanceData);
		uint16_t Type = n_DTCNonStandard::DOWNLOAD_HISTORICAL_ORDER_FILL_AND_ACCOUNT_BALANCE_DATA;
		uint16_t BaseSize = Size;

		DTC_VLS::vls_t TradeAccount;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_DownloadHistoricalOrderFillAndAccountBalanceData, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_HistoricalTradesRequest
	{
		uint16_t Size = sizeof(s_HistoricalTradesRequest);
		uint16_t Type = n_DTCNonStandard::HISTORICAL_TRADES_REQUEST;
		uint16_t BaseSize = Size;

		int32_t RequestID = 0;
		DTC_VLS::vls_t Symbol;
		DTC_VLS::vls_t TradeAccount;
		DTC::t_DateTime StartDateTime = 0;
		uint32_t SubAccountIdentifier = 0;
		uint8_t CreateFlatToFlatTrades = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		int32_t GetRequestID() const;

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_HistoricalTradesRequest, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_HistoricalTradesRequest, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		DTC::t_DateTime GetStartDateTime() const;
		uint32_t GetSubAccountIdentifier() const;
		uint8_t GetCreateFlatToFlatTrades() const;
	};

	/*==========================================================================*/
	struct s_HistoricalTradesReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		DTC_VLS::vls_t RejectText;

		s_HistoricalTradesReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalTradesReject));
			Size = sizeof(s_HistoricalTradesReject);
			Type = n_DTCNonStandard::HISTORICAL_TRADES_REJECT;
			BaseSize = Size;
		}

		uint32_t GetRequestID() const;

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_HistoricalTradesReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_HistoricalTradesResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		uint8_t IsFinalMessage;
		DTC_VLS::vls_t Symbol;
		DTC_VLS::vls_t TradeAccount;
		DTC::t_DateTimeWithMilliseconds EntryDateTime;
		DTC::t_DateTimeWithMilliseconds ExitDateTime;
		double EntryPrice;
		double ExitPrice;
		DTC::BuySellEnum TradeType;
		uint32_t EntryQuantity;
		uint32_t ExitQuantity;
		uint32_t MaxOpenQuantity;
		double ClosedProfitLoss;
		double MaximumOpenPositionLoss;
		double MaximumOpenPositionProfit;
		double Commission;
		DTC_VLS::vls_t OpenFillExecutionID;
		DTC_VLS::vls_t CloseFillExecutionID;
		uint32_t SubAccountIdentifier;

		s_HistoricalTradesResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalTradesResponse));
			Size = sizeof(s_HistoricalTradesResponse);
			Type = n_DTCNonStandard::HISTORICAL_TRADES_RESPONSE;
			BaseSize = Size;

			TradeType = DTC::BUY_SELL_UNSET;
		}

		int32_t GetRequestID() const;
		uint8_t GetIsFinalMessage() const;

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_HistoricalTradesResponse, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_HistoricalTradesResponse, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		DTC::t_DateTimeWithMilliseconds GetEntryDateTime() const;
		DTC::t_DateTimeWithMilliseconds GetExitDateTime() const;
		double GetEntryPrice() const;
		double GetExitPrice() const;
		DTC::BuySellEnum GetTradeType() const;
		uint32_t GetEntryQuantity() const;
		uint32_t GetExitQuantity() const;
		uint32_t GetMaxOpenQuantity() const;
		double GetClosedProfitLoss() const;
		double GetMaximumOpenPositionLoss() const;
		double GetMaximumOpenPositionProfit() const;
		double GetCommission() const;

		const char* GetOpenFillExecutionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, OpenFillExecutionID, offsetof(s_HistoricalTradesResponse, OpenFillExecutionID));
		}

		void AddOpenFillExecutionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, OpenFillExecutionID, StringLength);
		}

		const char* GetCloseFillExecutionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, CloseFillExecutionID, offsetof(s_HistoricalTradesResponse, CloseFillExecutionID));
		}

		void AddCloseFillExecutionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, CloseFillExecutionID, StringLength);
		}

		uint32_t GetSubAccountIdentifier() const;
	};

	/*==========================================================================*/
	struct s_ReplayChartData
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		uint32_t RequestID;
		DTC_VLS::vls_t Symbol;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t TimeZone;
		DTC::t_DateTimeWithMillisecondsInt StartDateTimeForInitialData;
		DTC::t_DateTimeWithMillisecondsInt StartDateTime;
		DTC::t_DateTimeWithMillisecondsInt StopDateTime;
		uint16_t SessionBeginTimeInSeconds;
		uint16_t SessionEndTimeInSeconds;
		float ReplaySpeed;
		int32_t BarTimeInSeconds;
		uint8_t PauseReplayAfterInitialDataSent;
		uint8_t UseSavedPriorState;
		float SymbolVolatility;
		float InterestRate;
		int32_t NumberOfOrdersToTriggerFinishToStopDateTime;
		int32_t MaximumNumberOfOrdersPerReplaySession;
		int32_t NumberOfDaysForInitialDataFromBeforeLastSavedDateTime;
		uint32_t SubAccountIdentifier = 0;
		int32_t OptionsPriceSendIntervalInSeconds = 0;

		s_ReplayChartData()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_ReplayChartData));
			Size = sizeof(s_ReplayChartData);
			Type = n_DTCNonStandard::REPLAY_CHART_DATA;
			BaseSize = Size;

			ReplaySpeed = 1.0;
		}

		uint32_t GetRequestID() const;

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_ReplayChartData, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_ReplayChartData, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTimeZone() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TimeZone, offsetof(s_ReplayChartData, TimeZone));
		}

		void AddTimeZone(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TimeZone, StringLength);
		}

		DTC::t_DateTimeWithMillisecondsInt GetStartDateTimeForInitialData() const;
		DTC::t_DateTimeWithMillisecondsInt GetStartDateTime() const;
		DTC::t_DateTimeWithMillisecondsInt GetStopDateTime() const;
		uint16_t GetSessionBeginTimeInSeconds() const;
		uint16_t GetSessionEndTimeInSeconds() const;
		float GetReplaySpeed() const;
		int32_t GetBarTimeInSeconds() const;
		uint8_t GetPauseReplayAfterInitialDataSent() const;
		uint8_t GetUseSavedPriorState() const;
		float GetSymbolVolatility() const;
		float GetInterestRate() const;
		int32_t GetNumberOfOrdersToTriggerFinishToStopDateTime() const;
		int32_t GetMaximumNumberOfOrdersPerReplaySession() const;
		int32_t GetNumberOfDaysForInitialDataFromBeforeLastSavedDateTime() const;
		uint32_t GetSubAccountIdentifier() const;
		int32_t GetOptionsPriceSendIntervalInSeconds() const;
	};

	/*==========================================================================*/
	struct s_ReplayChartDataPerformAction
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		uint32_t RequestID = 0;
		n_DTCNonStandard::ReplayChartDataActionEnum Action = n_DTCNonStandard::REPLAY_CHART_DATA_ACTION_NONE;
		float ReplaySpeed = 0;

		s_ReplayChartDataPerformAction()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_ReplayChartDataPerformAction));
			Size = sizeof(s_ReplayChartDataPerformAction);
			Type = n_DTCNonStandard::REPLAY_CHART_DATA_PERFORM_ACTION;
			BaseSize = Size;
		}

		uint32_t GetRequestID() const;
		n_DTCNonStandard::ReplayChartDataActionEnum GetAction() const;
		float GetReplaySpeed() const;
	};

	/*==========================================================================*/
	struct s_ReplayChartDataStatus
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		uint32_t RequestID = 0;
		DTC_VLS::vls_t ErrorMessage;
		n_DTCNonStandard::ReplayChartDataStatusEnum Status = n_DTCNonStandard::REPLAY_CHART_DATA_STATUS_UNSET;
		uint32_t  SubAccountIdentifier = 0;

		s_ReplayChartDataStatus()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_ReplayChartDataStatus));
			Size = sizeof(s_ReplayChartDataStatus);
			Type = n_DTCNonStandard::REPLAY_CHART_DATA_STATUS;
			BaseSize = Size;
		}

		uint32_t GetRequestID() const;

		const char* GetErrorMessage() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ErrorMessage, offsetof(s_ReplayChartDataStatus, ErrorMessage));
		}

		void AddErrorMessage(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ErrorMessage, StringLength);
		}

		n_DTCNonStandard::ReplayChartDataStatusEnum GetStatus() const;
		uint32_t GetSubAccountIdentifier() const;
	};


	/*==========================================================================*/
	struct s_RequestNumCurrentClientConnections
	{
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::REQUEST_NUM_CURRENT_CLIENT_CONNECTIONS;
		uint16_t BaseSize = sizeof(*this);

		uint32_t RequestID = 0;
		DTC_VLS::vls_t Username;
		int64_t DeviceIdentifier = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;

		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_RequestNumCurrentClientConnections, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}

		int64_t GetDeviceIdentifier() const;
	};

	/*==========================================================================*/
	struct s_NumCurrentClientConnectionsResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		uint32_t RequestID = 0;
		DTC_VLS::vls_t Username;
		int32_t NumConnectionsForDifferentDevices = 0;
		int32_t NumConnectionsForSameDevice = 0;

		s_NumCurrentClientConnectionsResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_NumCurrentClientConnectionsResponse));
			Size = sizeof(s_NumCurrentClientConnectionsResponse);
			Type = n_DTCNonStandard::NUM_CURRENT_CLIENT_CONNECTIONS_RESPONSE;
			BaseSize = Size;
		}

		uint32_t GetRequestID() const;

		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_NumCurrentClientConnectionsResponse, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}

		int32_t GetNumConnectionsForDifferentDevices() const;
		int32_t GetNumConnectionsForSameDevice() const;
	};

	/*==========================================================================*/
	struct s_TradeOrder
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::SC_TRADE_ORDER;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint8_t m_IsOrderDeleted = 0;
		uint64_t m_InternalOrderID = 0;

		uint16_t m_OrderStatusCode = 0;
		uint16_t m_OrderStatusBeforePendingModify = 0;
		uint16_t m_OrderStatusBeforePendingCancel = 0;

		DTC_VLS::vls_t m_ServiceOrderID;

		DTC_VLS::vls_t m_ActualSymbol;

		int m_OrderType = 0;
		uint16_t m_BuySell = 0;

		double m_Price1 = 0;
		double m_Price2 = 0;

		double	m_OrderQuantity = 0;
		double	m_FilledQuantity = 0;

		double m_AverageFillPrice = 0;

		int m_RealtimeFillStatus = 0;

		uint8_t m_IsRestingOrderDuringFill = 0;

		int m_OrderRejectType = 0;

		DTC_VLS::vls_t m_TradeAccount;

		uint32_t m_SubAccountIdentifier = 0;

		int m_InternalOrderIDModifierForService = 0;

		DTC_VLS::vls_t m_FIXClientOrderID;

		uint32_t m_SequenceNumberBasedClientOrderID = 0;

		DTC_VLS::vls_t m_ClientOrderIDForDTCServer;
		DTC_VLS::vls_t m_PreviousClientOrderIDForDTCServer;

		DTC_VLS::vls_t m_ExchangeOrderID;

		DTC_VLS::vls_t m_OriginatingClientUsername;

		int64_t	m_EntryDateTime = 0;
		int64_t	m_LastActionDateTime = 0;
		int64_t m_ServiceUpdateDateTimeUTC = 0;

		unsigned int m_OrderEntryTimeForService = 0;
		unsigned int m_LastModifyTimeForService = 0;

		int64_t m_GoodTillDateTime = 0;
		int m_TimeInForce = 0;
		uint16_t m_OpenClose = 0;

		double m_TrailStopOffset1 = 0;
		double m_TrailStopStep = 0;
		double m_TrailTriggerPrice = 0;
		double m_TrailingStopTriggerOffset = 0;

		uint8_t m_TrailTriggerHit = 0;
		double m_TrailToBreakEvenStopOffset = 0;

		double m_MaximumChaseAmountAsPrice = 0;
		double m_InitialChaseOrderPrice1 = 0;
		double m_InitialLastTradePriceForChaseOrders = 0;

		int m_TrailingStopTriggerOCOGroupNumber = 0;

		double m_LastModifyPrice1 = 0;

		double m_LastModifyQuantity = 0;

		double m_CumulativeOrderQuantityFromParentFills = 0;

		double m_PriorFilledQuantity = 0;

		float m_TickSize = 0;

		int m_ValueFormat = 0;

		float m_PriceMultiplier = 0;

		uint64_t m_ParentInternalOrderID = 0;
		uint64_t m_TargetChildInternalOrderID = 0;
		uint64_t m_StopChildInternalOrderID = 0;

		double m_AttachedOrderPriceOffset1 = 0;

		uint64_t m_LinkInternalOrderID = 0;

		uint64_t m_OCOGroupInternalOrderID = 0;

		uint64_t m_OCOSiblingInternalOrderID = 0;

		uint8_t m_DisableChildAndSiblingRelatedActions = 0;

		uint8_t m_OCOManagedByService = 0;
		uint8_t m_BracketOrderServerManaged = 0;

		DTC_VLS::vls_t m_LastOrderActionSource;

		uint8_t m_StopLimitOrderStopPriceTriggered = 0;

		uint8_t m_OCOFullSiblingCancelOnPartialFill = 0;

		uint8_t m_ReverseOnCompleteFill = 0;

		uint8_t m_SupportScaleIn = 0;

		uint8_t m_SupportScaleOut = 0;

		int m_SourceChartNumber = 0;
		DTC_VLS::vls_t m_SourceChartbookFileName;

		uint8_t m_IsAutomatedOrder = 0;
		uint8_t m_SimulatedOrder = 0;
		uint8_t m_IsChartReplaying = 0;
		int m_AttachedOrderOCOGroupNumber = 0;

		DTC_VLS::vls_t m_LastFillExecutionServiceID;

		int m_FillCount = 0;

		double m_LastFillQuantity = 0;

		double m_LastFillPrice = 0;

		int64_t m_LastFillDateTimeUTC = 0;

		uint64_t m_RejectedStopOCOSiblingInternalOrderID = 0;

		double m_RejectedStopReplacementMarketOrderQuantity = 0;

		uint8_t m_EvaluatingForFill = 0;

		unsigned int m_LastProcessedTimeSalesRecordSequenceForPrices = 0;

		uint8_t m_IsMarketDataManagementOfOrderEnabled = 0;

		DTC_VLS::vls_t m_TextTag;

		int64_t m_TimedOutOrderRequestedStatusDateTime = 0;

		uint8_t m_RequestedStatusForTimedOutOrder = 0;

		uint8_t m_SendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled = 0;

		double m_QuantityToIncreaseFromParentFill = 0;

		double m_MoveToBreakevenStopReferencePrice = 0;

		double m_QuantityTriggeredStop_QuantityForTrigger = 0;
		double m_AccumulatedTradeVolumeForTriggeredStop = 0;
		uint8_t m_BidAskQuantityStopInitialTriggerMet = 0;

		uint8_t m_NeedToOverrideLock = 0;

		double m_CurrentMarketPrice = 0;
		int64_t m_CurrentMarketDateTime = 0;

		uint8_t m_SupportOrderFillBilling = 0;
		uint8_t m_IsBillable = 0;
		int m_QuantityForBilling = 0;

		uint32_t m_NumberOfFailedOrderModifications = 0;
		int32_t m_DTCServerIndex = 0;

		DTC_VLS::vls_t m_ClearingFirmID;
		DTC_VLS::vls_t m_SenderSubID;
		DTC_VLS::vls_t m_SenderLocationId;
		DTC_VLS::vls_t m_SelfMatchPreventionID;
		int32_t m_CTICode = 0;

		uint8_t m_ObtainOrderActionDateTimeFromLastTradeTimeInChart = 0;
		double m_MaximumShowQuantity = 0;

		uint8_t m_OrderSubmitted = 0;

		uint8_t m_IsSnapshot = 0;
		uint8_t m_IsFirstMessageInBatch = 0;
		uint8_t m_IsLastMessageInBatch = 0;

		int64_t m_ExternalLastActionDateTimeUTC = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint8_t GetIsOrderDeleted() const;
		uint64_t GetInternalOrderID() const;
		uint16_t GetOrderStatusCode() const;
		uint16_t GetOrderStatusBeforePendingModify() const;
		uint16_t GetOrderStatusBeforePendingCancel() const;
		int GetOrderType() const;
		uint16_t GetBuySell() const;
		double GetPrice1() const;
		double GetPrice2() const;
		double	GetOrderQuantity() const;
		double	GetFilledQuantity() const;
		double GetAverageFillPrice() const;
		int GetRealtimeFillStatus() const;
		uint8_t GetIsRestingOrderDuringFill() const;
		int GetOrderRejectType() const;
		uint32_t GetSubAccountIdentifier() const;
		int GetInternalOrderIDModifierForService() const;
		uint32_t GetSequenceNumberBasedClientOrderID() const;
		int64_t	GetEntryDateTime() const;
		int64_t	GetLastActionDateTime() const;
		int64_t GetServiceUpdateDateTimeUTC() const;
		unsigned int GetOrderEntryTimeForService() const;
		unsigned int GetLastModifyTimeForService() const;
		int64_t GetGoodTillDateTime() const;
		int GetTimeInForce() const;
		uint16_t GetOpenClose() const;
		double GetTrailStopOffset1() const;
		double GetTrailStopStep() const;
		double GetTrailTriggerPrice() const;
		double GetTrailingStopTriggerOffset() const;
		uint8_t GetTrailTriggerHit() const;
		double GetTrailToBreakEvenStopOffset() const;
		double GetMaximumChaseAmountAsPrice() const;
		double GetInitialChaseOrderPrice1() const;
		double GetInitialLastTradePriceForChaseOrders() const;
		int GetTrailingStopTriggerOCOGroupNumber() const;
		double GetLastModifyPrice1() const;
		double GetLastModifyQuantity() const;
		double GetCumulativeOrderQuantityFromParentFills() const;
		double GetPriorFilledQuantity() const;
		float GetTickSize() const;
		int GetValueFormat() const;
		float GetPriceMultiplier() const;
		uint64_t GetParentInternalOrderID() const;
		uint64_t GetTargetChildInternalOrderID() const;
		uint64_t GetStopChildInternalOrderID() const;
		double GetAttachedOrderPriceOffset1() const;
		uint64_t GetLinkInternalOrderID() const;
		uint64_t GetOCOGroupInternalOrderID() const;
		uint64_t GetOCOSiblingInternalOrderID() const;
		uint8_t GetDisableChildAndSiblingRelatedActions() const;
		uint8_t GetOCOManagedByService() const;
		uint8_t GetBracketOrderServerManaged() const;
		uint8_t GetStopLimitOrderStopPriceTriggered() const;
		uint8_t GetOCOFullSiblingCancelOnPartialFill() const;
		uint8_t GetReverseOnCompleteFill() const;
		uint8_t GetSupportScaleIn() const;
		uint8_t GetSupportScaleOut() const;
		int GetSourceChartNumber() const;
		uint8_t GetIsAutomatedOrder() const;
		uint8_t GetSimulatedOrder() const;
		uint8_t GetIsChartReplaying() const;
		int GetAttachedOrderOCOGroupNumber() const;
		int GetFillCount() const;
		double GetLastFillQuantity() const;
		double GetLastFillPrice() const;
		int64_t GetLastFillDateTimeUTC() const;
		uint64_t GetRejectedStopOCOSiblingInternalOrderID() const;
		double GetRejectedStopReplacementMarketOrderQuantity() const;
		uint8_t GetEvaluatingForFill() const;
		unsigned int GetLastProcessedTimeSalesRecordSequenceForPrices() const;
		uint8_t GetIsMarketDataManagementOfOrderEnabled() const;
		int64_t GetTimedOutOrderRequestedStatusDateTime() const;
		uint8_t GetRequestedStatusForTimedOutOrder() const;
		uint8_t GetSendFlattenMarketOrderWhenRelatedOrdersConfirmedCanceled() const;
		double GetQuantityToIncreaseFromParentFill() const;
		double GetMoveToBreakevenStopReferencePrice() const;
		double GetQuantityTriggeredStop_QuantityForTrigger() const;
		double GetAccumulatedTradeVolumeForTriggeredStop() const;
		uint8_t GetBidAskQuantityStopInitialTriggerMet() const;
		uint8_t GetNeedToOverrideLock() const;
		double GetCurrentMarketPrice() const;
		int64_t GetCurrentMarketDateTime() const;
		uint8_t GetSupportOrderFillBilling() const;
		uint8_t GetIsBillable() const;
		int GetQuantityForBilling() const;
		uint32_t GetNumberOfFailedOrderModifications() const;
		int32_t GetDTCServerIndex() const;
		int32_t GetCTICode() const;
		uint8_t GetObtainOrderActionDateTimeFromLastTradeTimeInChart() const;
		double GetMaximumShowQuantity() const;
		uint8_t GetOrderSubmitted() const;

		uint8_t GetIsSnapshot() const;
		uint8_t GetIsFirstMessageInBatch() const;
		uint8_t GetIsLastMessageInBatch() const;

		int64_t GetExternalLastActionDateTimeUTC() const;

		//Implementing Get Add inline functions------------------------------
		const char* GetServiceOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_ServiceOrderID, offsetof(s_TradeOrder, m_ServiceOrderID));
		}

		void AddServiceOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_ServiceOrderID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetActualSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_ActualSymbol, offsetof(s_TradeOrder, m_ActualSymbol));
		}

		void AddActualSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_ActualSymbol, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_TradeAccount, offsetof(s_TradeOrder, m_TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetFIXClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_FIXClientOrderID, offsetof(s_TradeOrder, m_FIXClientOrderID));
		}

		void AddFIXClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_FIXClientOrderID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetClientOrderIDForDTCServer() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_ClientOrderIDForDTCServer, offsetof(s_TradeOrder, m_ClientOrderIDForDTCServer));
		}

		void AddClientOrderIDForDTCServer(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_ClientOrderIDForDTCServer, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetPreviousClientOrderIDForDTCServer() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_PreviousClientOrderIDForDTCServer, offsetof(s_TradeOrder, m_PreviousClientOrderIDForDTCServer));
		}

		void AddPreviousClientOrderIDForDTCServer(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_PreviousClientOrderIDForDTCServer, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetExchangeOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_ExchangeOrderID, offsetof(s_TradeOrder, m_ExchangeOrderID));
		}

		void AddExchangeOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_ExchangeOrderID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetOriginatingClientUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_OriginatingClientUsername, offsetof(s_TradeOrder, m_OriginatingClientUsername));
		}

		void AddOriginatingClientUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_OriginatingClientUsername, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetLastOrderActionSource() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_LastOrderActionSource, offsetof(s_TradeOrder, m_LastOrderActionSource));
		}

		void AddLastOrderActionSource(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_LastOrderActionSource, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSourceChartbookFileName() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_SourceChartbookFileName, offsetof(s_TradeOrder, m_SourceChartbookFileName));
		}

		void AddSourceChartbookFileName(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_SourceChartbookFileName, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetLastFillExecutionServiceID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_LastFillExecutionServiceID, offsetof(s_TradeOrder, m_LastFillExecutionServiceID));
		}

		void AddLastFillExecutionServiceID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_LastFillExecutionServiceID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetTextTag() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_TextTag, offsetof(s_TradeOrder, m_TextTag));
		}

		void AddTextTag(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_TextTag, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetClearingFirmID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_ClearingFirmID, offsetof(s_TradeOrder, m_ClearingFirmID));
		}

		void AddClearingFirmID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_ClearingFirmID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSenderSubID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_SenderSubID, offsetof(s_TradeOrder, m_SenderSubID));
		}

		void AddSenderSubID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_SenderSubID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSenderLocationId() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_SenderLocationId, offsetof(s_TradeOrder, m_SenderLocationId));
		}

		void AddSenderLocationId(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_SenderLocationId, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSelfMatchPreventionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_SelfMatchPreventionID, offsetof(s_TradeOrder, m_SelfMatchPreventionID));
		}

		void AddSelfMatchPreventionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_SelfMatchPreventionID, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_IndividualTradePosition
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::INDIVIDUAL_TRADE_POSITION;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		DTC_VLS::vls_t m_Symbol;
		uint8_t m_IsSimulated = 0;
		DTC_VLS::vls_t m_TradeAccount;
		double m_Quantity = 0;
		double m_PositionPrice = 0;
		double m_OpenProfitLoss = 0;
		int64_t m_TradeDateTime = 0;
		DTC_VLS::vls_t m_FillExecutionIdentifier;

		uint8_t m_IsSnapshot = 0;
		uint8_t m_IsFirstMessageInBatch = 0;
		uint8_t m_IsLastMessageInBatch = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		//Implementing Get Add inline functions------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_Symbol, offsetof(s_IndividualTradePosition, m_Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_Symbol, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_TradeAccount, offsetof(s_IndividualTradePosition, m_TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetFillExecutionIdentifier() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_FillExecutionIdentifier, offsetof(s_IndividualTradePosition, m_FillExecutionIdentifier));
		}

		void AddFillExecutionIdentifier(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_FillExecutionIdentifier, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradePositionConsolidated
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_POSITION_CONSOLIDATED;
		uint16_t BaseSize = sizeof(*this);
	
		//message field variables
		uint8_t m_IsDeleted = 0;
		DTC_VLS::vls_t m_Symbol;
		uint8_t m_IsSimulated = 0;
		DTC_VLS::vls_t m_TradeAccount;
		DTC_VLS::vls_t m_CurrencyCode;
		double m_Quantity;
		double m_AveragePrice = 0;
		double m_OpenProfitLoss = 0;

		double m_DailyProfitLoss = 0;
		int64_t m_LastDailyProfitLossResetDateTimeUTC = 0;

		double m_ServicePositionQuantity = 0;
		uint8_t m_PositionHasBeenUpdatedByService = 0;

		double m_PriceHighDuringPosition = 0;
		double m_PriceLowDuringPosition = 0;
		double m_PriceLastDuringPosition = 0;

		int64_t m_LastProcessedTimeAndSalesSequence = 0;

		double m_TotalMarginRequirement = 0.0;

		int64_t m_InitialEntryDateTimeUTC = 0;

		uint8_t m_IsFromDTCServerReplay = 0;
		int64_t m_MostRecentPositionIncreaseDateTimeUTC = 0;

		uint8_t m_IsSnapshot = 0;
		uint8_t m_IsFirstMessageInBatch = 0;
		uint8_t m_IsLastMessageInBatch = 0;

		double m_MarginRequirementFull = 0;
		double m_MarginRequirementFullPositionsOnly = 0;

		double m_MaxPotentialPositionQuantity = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		//Implementing Get Add inline functions------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_Symbol, offsetof(s_TradePositionConsolidated, m_Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_Symbol, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_TradeAccount, offsetof(s_TradePositionConsolidated, m_TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetCurrencyCode() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_CurrencyCode, offsetof(s_TradePositionConsolidated, m_CurrencyCode));
		}

		void AddCurrencyCode(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_CurrencyCode, StringLength);
		}
		//-------------------------------------------------------------------

	};

	/*==========================================================================*/
	struct s_TradeActivityData
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACTIVITY_DATA;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint8_t ActivityType = 0;
		int64_t DataDateTimeUTC = 0;
		DTC_VLS::vls_t Symbol;

		DTC_VLS::vls_t OrderActionSource;
		uint64_t InternalOrderID = 0;
		DTC_VLS::vls_t ServiceOrderID;
		DTC_VLS::vls_t ExchangeOrderID;
		DTC_VLS::vls_t FIXClientOrderID;
		DTC_VLS::vls_t OrderTypeName;

		double Quantity = 0;

		uint8_t BuySell = 0;
		double Price1 = 0;
		double Price2 = 0;
		uint8_t NewOrderStatus = 0;
		double FillPrice = 0.0;
		double OrderFilledQuantity = 0;
		double HighPriceDuringPosition = 0;
		double LowPriceDuringPosition = 0;
		double LastPriceDuringPosition = 0;
		DTC_VLS::vls_t TradeAccount;
		uint64_t ParentInternalOrderID = 0;

		uint8_t OpenClose = 0;

		uint8_t IsSimulated = 0;
		uint8_t IsAutomatedOrder = 0;
		uint8_t IsChartReplaying = 0;

		DTC_VLS::vls_t FillExecutionServiceID;
		double PositionQuantity = 0;

		int SourceChartNumber = 0;
		DTC_VLS::vls_t SourceChartbookFileName;

		int TimeInForce = 0;

		DTC_VLS::vls_t SymbolServiceCode;

		DTC_VLS::vls_t Note;
		DTC_VLS::vls_t OriginatingClientUsername;

		double TradeAccountBalance = 0;

		uint8_t SupportsPositionQuantityField = 0;

		uint8_t IsBillable = 0;
		int32_t QuantityForBilling = 0;
		DTC_VLS::vls_t OrderRoutingServiceCode;
		uint32_t SubAccountIdentifier = 0;

		int64_t AuditTrail_TransactDateTimeUTC = 0;
		int AuditTrail_MessageDirection = 0;
		DTC_VLS::vls_t AuditTrail_OperatorID;
		DTC_VLS::vls_t AuditTrail_SelfMatchPreventionID;
		DTC_VLS::vls_t AuditTrail_SessionID;
		DTC_VLS::vls_t AuditTrail_ExecutingFirmID;
		DTC_VLS::vls_t AuditTrail_FixMessageType;
		int16_t AuditTrail_CustomerTypeIndicator = 0;
		int16_t AuditTrail_CustomerOrFirm = 0;
		DTC_VLS::vls_t AuditTrail_ExecutionReportID;
		DTC_VLS::vls_t AuditTrail_SpreadLegLinkID;
		DTC_VLS::vls_t AuditTrail_SecurityDesc;
		DTC_VLS::vls_t AuditTrail_MarketSegmentID;
		uint8_t AuditTrail_IFMFlag = 0;
		double AuditTrail_DisplayQuantity = 0;
		DTC_VLS::vls_t AuditTrail_CountryOfOrigin;
		double AuditTrail_FillQuantity = 0;
		double AuditTrail_RemainingQuantity = 0;
		uint8_t AuditTrail_AggressorFlag = 0;
		int32_t AuditTrail_SourceOfCancellation = 0;
		DTC_VLS::vls_t AuditTrail_OrdRejReason;

		uint8_t IsSnapshot = 0;
		uint8_t IsFirstMessageInBatch = 0;
		uint8_t IsLastMessageInBatch = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		//Implementing Get Add inline functions------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_TradeActivityData, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetOrderActionSource() const
		{
			return GetVariableLengthStringField(Size, BaseSize, OrderActionSource, offsetof(s_TradeActivityData, OrderActionSource));
		}

		void AddOrderActionSource(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, OrderActionSource, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetServiceOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServiceOrderID, offsetof(s_TradeActivityData, ServiceOrderID));
		}

		void AddServiceOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServiceOrderID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetExchangeOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ExchangeOrderID, offsetof(s_TradeActivityData, ExchangeOrderID));
		}

		void AddExchangeOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ExchangeOrderID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetFIXClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FIXClientOrderID, offsetof(s_TradeActivityData, FIXClientOrderID));
		}

		void AddFIXClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FIXClientOrderID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetOrderTypeName() const
		{
			return GetVariableLengthStringField(Size, BaseSize, OrderTypeName, offsetof(s_TradeActivityData, OrderTypeName));
		}

		void AddOrderTypeName(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, OrderTypeName, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeActivityData, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetFillExecutionServiceID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FillExecutionServiceID, offsetof(s_TradeActivityData, FillExecutionServiceID));
		}

		void AddFillExecutionServiceID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FillExecutionServiceID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSourceChartbookFileName() const
		{
			return GetVariableLengthStringField(Size, BaseSize, SourceChartbookFileName, offsetof(s_TradeActivityData, SourceChartbookFileName));
		}

		void AddSourceChartbookFileName(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, SourceChartbookFileName, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSymbolServiceCode() const
		{
			return GetVariableLengthStringField(Size, BaseSize, SymbolServiceCode, offsetof(s_TradeActivityData, SymbolServiceCode));
		}

		void AddSymbolServiceCode(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, SymbolServiceCode, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetNote() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Note, offsetof(s_TradeActivityData, Note));
		}

		void AddNote(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Note, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetOriginatingClientUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, OriginatingClientUsername, offsetof(s_TradeActivityData, OriginatingClientUsername));
		}

		void AddOriginatingClientUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, OriginatingClientUsername, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetOrderRoutingServiceCode() const
		{
			return GetVariableLengthStringField(Size, BaseSize, OrderRoutingServiceCode, offsetof(s_TradeActivityData, OrderRoutingServiceCode));
		}

		void AddOrderRoutingServiceCode(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, OrderRoutingServiceCode, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_OperatorID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_OperatorID, offsetof(s_TradeActivityData, AuditTrail_OperatorID));
		}

		void AddAuditTrail_OperatorID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_OperatorID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_SelfMatchPreventionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_SelfMatchPreventionID, offsetof(s_TradeActivityData, AuditTrail_SelfMatchPreventionID));
		}

		void AddAuditTrail_SelfMatchPreventionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_SelfMatchPreventionID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_SessionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_SessionID, offsetof(s_TradeActivityData, AuditTrail_SessionID));
		}

		void AddAuditTrail_SessionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_SessionID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_ExecutingFirmID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_ExecutingFirmID, offsetof(s_TradeActivityData, AuditTrail_ExecutingFirmID));
		}

		void AddAuditTrail_ExecutingFirmID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_ExecutingFirmID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_FixMessageType() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_FixMessageType, offsetof(s_TradeActivityData, AuditTrail_FixMessageType));
		}

		void AddAuditTrail_FixMessageType(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_FixMessageType, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_ExecutionReportID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_ExecutionReportID, offsetof(s_TradeActivityData, AuditTrail_ExecutionReportID));
		}

		void AddAuditTrail_ExecutionReportID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_ExecutionReportID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_SpreadLegLinkID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_SpreadLegLinkID, offsetof(s_TradeActivityData, AuditTrail_SpreadLegLinkID));
		}

		void AddAuditTrail_SpreadLegLinkID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_SpreadLegLinkID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_SecurityDesc() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_SecurityDesc, offsetof(s_TradeActivityData, AuditTrail_SecurityDesc));
		}

		void AddAuditTrail_SecurityDesc(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_SecurityDesc, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_MarketSegmentID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_MarketSegmentID, offsetof(s_TradeActivityData, AuditTrail_MarketSegmentID));
		}

		void AddAuditTrail_MarketSegmentID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_MarketSegmentID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetAuditTrail_CountryOfOrigin() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_CountryOfOrigin, offsetof(s_TradeActivityData, AuditTrail_CountryOfOrigin));
		}

		void AddAuditTrail_CountryOfOrigin(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_CountryOfOrigin, StringLength);
		}		

		//-------------------------------------------------------------------
		const char* GetAuditTrail_OrdRejReason() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AuditTrail_OrdRejReason, offsetof(s_TradeActivityData, AuditTrail_OrdRejReason));
		}

		void AddAuditTrail_OrdRejReason(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AuditTrail_OrdRejReason, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataRequest
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_REQUEST;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		DTC_VLS::vls_t TradeAccount;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataRequest, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//--End of Get Add functions----------------------------------------

	};

	/*==========================================================================*/
	struct s_TradeAccountDataResponse
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_RESPONSE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t m_RequestID = 0;
		uint8_t m_TradeAccountNotExist = 0;

		DTC_VLS::vls_t m_TradeAccount;
		uint8_t m_IsSimulated = false;
		DTC_VLS::vls_t m_CurrencyCode;

		double m_CurrentCashBalance = 0;
		double m_AvailableFundsForNewPositions = 0;
		double m_MarginRequirement = 0;
		double m_AccountValue = 0;
		double m_OpenPositionsProfitLoss = 0;
		double m_DailyProfitLoss = 0;
		uint64_t m_TransactionIdentifierForCashBalanceAdjustment = 0;
		int64_t m_LastTransactionDateTime = 0;
		double m_TrailingAccountValueAtWhichToNotAllowNewPositions = 0;
		double m_CalculatedDailyNetLossLimitInAccountCurrency = 0;
		uint8_t m_DailyNetLossLimitHasBeenReached = 0;
		int64_t m_LastResetDailyNetLossManagementVariablesDateTimeUTC = 0;
		uint8_t m_IsUnderRequiredMargin = 0;

		float m_DailyNetLossLimitInAccountCurrency = 0;

		int32_t m_PercentOfCashBalanceForDailyNetLossLimit = 0;

		uint8_t m_UseTrailingAccountValueToNotAllowIncreaseInPositions = false;

		uint8_t m_DoNotAllowIncreaseInPositionsAtDailyLossLimit = false;

		uint8_t m_FlattenPositionsAtDailyLossLimit = false;
		uint8_t m_ClosePositionsAtEndOfDay = false;

		uint8_t m_FlattenPositionsWhenUnderMarginIntraday = true;
		uint8_t m_FlattenPositionsWhenUnderMarginAtEndOfDay = false;

		DTC_VLS::vls_t m_SenderSubID;
		DTC_VLS::vls_t m_SenderLocationId;
		DTC_VLS::vls_t m_SelfMatchPreventionID;
		int32_t m_CTICode = 0;
		uint8_t m_TradeAccountIsReadOnly = false;
		int32_t m_MaximumGlobalPositionQuantity = 0;
		double m_TradeFeePerContract = 0;
		double m_TradeFeePerShare = 0;
		uint8_t m_RequireSufficientMarginForNewPositions = true;
		int32_t m_UsePercentOfMargin = 100;
		int32_t m_UsePercentOfMarginForDayTrading = 100;
		int32_t m_MaximumAllowedAccountBalanceForPositionsAsPercentage = 100;
		DTC_VLS::vls_t m_FirmID;

		uint8_t m_TradingIsDisabled = 0;
		DTC_VLS::vls_t m_DescriptiveName;

		uint8_t m_IsMasterFirmControlAccount = 0;
		double m_MinimumRequiredAccountValue = 0;
		int64_t m_BeginTimeForDayMargin = 0;
		int64_t m_EndTimeForDayMargin = 0;
		DTC_VLS::vls_t m_DayMarginTimeZone;

		uint8_t m_IsSnapshot = 0;
		uint8_t m_IsFirstMessageInBatch = 0;
		uint8_t m_IsLastMessageInBatch = 0;
		uint8_t m_IsDeleted = 0;

		uint8_t m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday = 0;
		uint8_t m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay = 0;
		uint8_t m_UseMasterFirm_SymbolLimitsArray = 0;// Unused
		uint8_t m_UseMasterFirm_TradeFees = 0;
		uint8_t m_UseMasterFirm_TradeFeePerShare = 0;
		uint8_t m_UseMasterFirm_RequireSufficientMarginForNewPositions = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMargin = 0;
		uint8_t m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = 0;
		uint8_t m_UseMasterFirm_MinimumRequiredAccountValue = 0;
		uint8_t m_UseMasterFirm_MarginTimeSettings = 0;
		uint8_t m_UseMasterFirm_TradingIsDisabled = 0;
		
		uint8_t m_IsTradeStatisticsPublicallyShared = 0;
		uint8_t m_IsReadOnlyFollowingRequestsAllowed = 0;
		uint8_t m_IsTradeAccountSharingAllowed = 0;
		uint8_t m_ReadOnlyShareToAllSCUsernames = 0;

		uint8_t m_UseMasterFirm_SymbolCommissionsArray = 0;
		uint8_t m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginForDayTrading = 0;
		double m_OpenPositionsProfitLossBasedOnSettlement = 0;
		double m_MarginRequirementFull = 0;
		double m_MarginRequirementFullPositionsOnly = 0;
		// End read only
				
		uint8_t m_UseMasterFirm_TradeFeesFullOverride = 0;
		uint8_t m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginFullOverride = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride = 0;

		double m_PeakMarginRequirement = 0;// Read only

		uint8_t m_LiquidationOnlyMode = 0;

		uint8_t m_FlattenPositionsWhenUnderMarginIntradayTriggered = 0;// Read only
		uint8_t m_FlattenPositionsWhenUnderMinimumAccountValueTriggered = 0;// Read only

		double m_AccountValueAtEndOfDayCaptureTime = 0;// Read only
		int64_t m_EndOfDayCaptureTime = 0;// Read only

		uint8_t m_CustomerOrFirm = 0;
		uint8_t m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMarginIntraday = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay = 0;
		uint32_t m_MasterFirm_MaximumOrderQuantity = 0;

		int64_t m_LastTriggerDateTimeUTCForDailyLossLimit = 0;// Read only
		uint8_t m_OpenPositionsProfitLossIsDelayed = 0;// Read only

		DTC_VLS::vls_t m_ExchangeTraderId;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetTradeAccountNotExist() const;
		uint8_t GetIsSimulated() const;
		double GetCurrentCashBalance() const;
		double GetAvailableFundsForNewPositions() const;
		double GetMarginRequirement() const;
		double GetAccountValue() const;
		double GetOpenPositionsProfitLoss() const;
		double GetDailyProfitLoss() const;
		uint64_t GetTransactionIdentifierForCashBalanceAdjustment() const;
		int64_t GetLastTransactionDateTime() const;
		double GetTrailingAccountValueAtWhichToNotAllowNewPositions() const;
		double GetCalculatedDailyNetLossLimitInAccountCurrency() const;
		uint8_t GetDailyNetLossLimitHasBeenReached() const;
		int64_t GetLastResetDailyNetLossManagementVariablesDateTimeUTC() const;
		uint8_t GetIsUnderRequiredMargin() const;
		float GetDailyNetLossLimitInAccountCurrency() const;
		int32_t GetPercentOfCashBalanceForDailyNetLossLimit() const;
		uint8_t GetUseTrailingAccountValueToNotAllowIncreaseInPositions() const;
		uint8_t GetDoNotAllowIncreaseInPositionsAtDailyLossLimit() const;
		uint8_t GetFlattenPositionsAtDailyLossLimit() const;
		uint8_t GetClosePositionsAtEndOfDay() const;
		uint8_t GetFlattenPositionsWhenUnderMarginIntraday() const;
		uint8_t GetFlattenPositionsWhenUnderMarginAtEndOfDay() const;
		int32_t GetCTICode() const;
		uint8_t GetTradeAccountIsReadOnly() const;
		int32_t GetMaximumGlobalPositionQuantity() const;
		double GetTradeFeePerContract() const;
		double GetTradeFeePerShare() const;
		uint8_t GetRequireSufficientMarginForNewPositions() const;
		int32_t GetUsePercentOfMargin() const;
		int32_t GetUsePercentOfMarginForDayTrading() const;
		int32_t GetMaximumAllowedAccountBalanceForPositionsAsPercentage() const;
		uint8_t GetTradingIsDisabled() const;
		uint8_t GetIsMasterFirmControlAccount() const;
		double GetMinimumRequiredAccountValue() const;
		int64_t GetBeginTimeForDayMargin() const;
		int64_t GetEndTimeForDayMargin() const;
		uint8_t GetIsSnapshot() const;
		uint8_t GetIsFirstMessageInBatch() const;
		uint8_t GetIsLastMessageInBatch() const;
		uint8_t GetIsDeleted() const;
		uint8_t GetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() const;
		uint8_t GetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() const;
		uint8_t GetUseMasterFirm_SymbolLimitsArray() const;
		uint8_t GetUseMasterFirm_TradeFees() const;
		uint8_t GetUseMasterFirm_TradeFeePerShare() const;
		uint8_t GetUseMasterFirm_RequireSufficientMarginForNewPositions() const;
		uint8_t GetUseMasterFirm_UsePercentOfMargin() const;
		uint8_t GetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() const;
		uint8_t GetUseMasterFirm_MinimumRequiredAccountValue() const;
		uint8_t GetUseMasterFirm_MarginTimeSettings() const;
		uint8_t GetUseMasterFirm_TradingIsDisabled() const;
		uint8_t GetIsTradeStatisticsPublicallyShared() const;
		uint8_t GetIsReadOnlyFollowingRequestsAllowed() const;
		uint8_t GetIsTradeAccountSharingAllowed() const;
		uint8_t GetReadOnlyShareToAllSCUsernames() const;
		uint8_t GetUseMasterFirm_SymbolCommissionsArray() const;
		uint8_t GetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginForDayTrading() const;
		double GetOpenPositionsProfitLossBasedOnSettlement() const;
		double GetMarginRequirementFull() const;
		double GetMarginRequirementFullPositionsOnly() const;

		uint8_t GetUseMasterFirm_TradeFeesFullOverride() const;
		uint8_t GetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginFullOverride() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() const;
		double GetPeakMarginRequirement() const;
		uint8_t GetLiquidationOnlyMode() const;

		uint8_t GetFlattenPositionsWhenUnderMarginIntradayTriggered() const;
		uint8_t GetFlattenPositionsWhenUnderMinimumAccountValueTriggered() const;

		double GetAccountValueAtEndOfDayCaptureTime() const;
		int64_t GetEndOfDayCaptureTime() const;

		uint8_t GetCustomerOrFirm() const;
		uint8_t GetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMarginIntraday() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay() const;
		uint32_t GetMasterFirm_MaximumOrderQuantity() const;

		int64_t GetLastTriggerDateTimeUTCForDailyLossLimit() const;
		uint8_t GetOpenPositionsProfitLossIsDelayed() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_TradeAccount, offsetof(s_TradeAccountDataResponse, m_TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetCurrencyCode() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_CurrencyCode, offsetof(s_TradeAccountDataResponse, m_CurrencyCode));
		}

		void AddCurrencyCode(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_CurrencyCode, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSenderSubID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_SenderSubID, offsetof(s_TradeAccountDataResponse, m_SenderSubID));
		}

		void AddSenderSubID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_SenderSubID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSenderLocationId() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_SenderLocationId, offsetof(s_TradeAccountDataResponse, m_SenderLocationId));
		}

		void AddSenderLocationId(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_SenderLocationId, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSelfMatchPreventionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_SelfMatchPreventionID, offsetof(s_TradeAccountDataResponse, m_SelfMatchPreventionID));
		}

		void AddSelfMatchPreventionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_SelfMatchPreventionID, StringLength);
		}
		//-------------------------------------------------------------------
			const char* GetFirmID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_FirmID, offsetof(s_TradeAccountDataResponse, m_FirmID));
		}

		void AddFirmID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_FirmID, StringLength);
		}
		
		//-------------------------------------------------------------------
			const char* GetDescriptiveName() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_DescriptiveName, offsetof(s_TradeAccountDataResponse, m_DescriptiveName));
		}

		void AddDescriptiveName(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_DescriptiveName, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetDayMarginTimeZone() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_DayMarginTimeZone, offsetof(s_TradeAccountDataResponse, m_DayMarginTimeZone));
		}

		void AddDayMarginTimeZone(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_DayMarginTimeZone, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetExchangeTraderId() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_ExchangeTraderId, offsetof(s_TradeAccountDataResponse, m_ExchangeTraderId));
		}

		void AddExchangeTraderId(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_ExchangeTraderId, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataUpdate
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_UPDATE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		uint8_t IsNewAccount = 0;

		DTC_VLS::vls_t NewAccountAuthorizedUsername;
		DTC_VLS::vls_t TradeAccount;

		uint8_t CurrencyCodeIsSet = 0;
		DTC_VLS::vls_t CurrencyCode;

		uint8_t DailyNetLossLimitInAccountCurrencyIsSet = 0;
		float DailyNetLossLimitInAccountCurrency = 0;

		uint8_t PercentOfCashBalanceForDailyNetLossLimitIsSet = 0;
		int32_t PercentOfCashBalanceForDailyNetLossLimit = 0;

		uint8_t UseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet = 0;
		uint8_t UseTrailingAccountValueToNotAllowIncreaseInPositions = 0;

		uint8_t DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet = 0;
		uint8_t DoNotAllowIncreaseInPositionsAtDailyLossLimit = 0;

		uint8_t FlattenPositionsAtDailyLossLimitIsSet = 0;
		uint8_t FlattenPositionsAtDailyLossLimit = 0;
		uint8_t ClosePositionsAtEndOfDayIsSet = 0;
		uint8_t ClosePositionsAtEndOfDay = 0;

		uint8_t FlattenPositionsWhenUnderMarginIntradayIsSet = 0;
		uint8_t FlattenPositionsWhenUnderMarginIntraday = 0;
		uint8_t FlattenPositionsWhenUnderMarginAtEndOfDayIsSet = 0;
		uint8_t FlattenPositionsWhenUnderMarginAtEndOfDay = 0;

		uint8_t SenderSubIDIsSet = 0;
		DTC_VLS::vls_t SenderSubID;
		uint8_t SenderLocationIdIsSet = 0;
		DTC_VLS::vls_t SenderLocationId;
		uint8_t SelfMatchPreventionIDIsSet = 0;
		DTC_VLS::vls_t SelfMatchPreventionID;

		uint8_t CTICodeIsSet = 0;
		int32_t CTICode = 0;
		uint8_t TradeAccountIsReadOnlyIsSet = 0;
		uint8_t TradeAccountIsReadOnly = 0;
		uint8_t MaximumGlobalPositionQuantityIsSet = 0;
		int32_t MaximumGlobalPositionQuantity = 0;
		uint8_t TradeFeePerContractIsSet = 0;
		double TradeFeePerContract = 0;
		uint8_t TradeFeePerShareIsSet = 0;
		double TradeFeePerShare = 0;
		uint8_t RequireSufficientMarginForNewPositionsIsSet = 0;
		uint8_t RequireSufficientMarginForNewPositions = 1;
		uint8_t UsePercentOfMarginIsSet = 0;
		int32_t UsePercentOfMargin = 100;
		uint8_t UsePercentOfMarginForDayTradingIsSet = 0;
		int32_t UsePercentOfMarginForDayTrading = 100;
		uint8_t MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet = 0;
		int32_t MaximumAllowedAccountBalanceForPositionsAsPercentage = 100;
		uint8_t FirmIDIsSet = 0;
		DTC_VLS::vls_t FirmID;

		uint8_t TradingIsDisabledIsSet = 0;
		uint8_t TradingIsDisabled = 0;
		uint8_t DescriptiveNameIsSet = 0;
		DTC_VLS::vls_t DescriptiveName;
		uint8_t IsMasterFirmControlAccountIsSet = 0;
		uint8_t IsMasterFirmControlAccount = 0;
		uint8_t MinimumRequiredAccountValueIsSet = 0;
		double MinimumRequiredAccountValue = 0;

		uint8_t BeginTimeForDayMarginIsSet = 0;
		int64_t BeginTimeForDayMargin = 0;
		uint8_t EndTimeForDayMarginIsSet = 0;
		int64_t EndTimeForDayMargin = 0;
		uint8_t DayMarginTimeZoneIsSet = 0;
		DTC_VLS::vls_t DayMarginTimeZone;

		uint8_t m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet = 0;
		uint8_t m_UseMasterFirm_FlattenPositionsWhenUnderMarginIntraday = 0;
		uint8_t m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet = 0;
		uint8_t m_UseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay = 0;
		uint8_t m_UseMasterFirm_SymbolLimitsArrayIsSet = 0;// Unused
		uint8_t m_UseMasterFirm_SymbolLimitsArray = 0;// Unused
		uint8_t m_UseMasterFirm_TradeFeesIsSet = 0;
		uint8_t m_UseMasterFirm_TradeFees = 0;
		uint8_t m_UseMasterFirm_TradeFeePerShareIsSet = 0;
		uint8_t m_UseMasterFirm_TradeFeePerShare = 0;
		uint8_t m_UseMasterFirm_RequireSufficientMarginForNewPositionsIsSet = 0;
		uint8_t m_UseMasterFirm_RequireSufficientMarginForNewPositions = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginIsSet = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMargin = 0;
		uint8_t m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet = 0;
		uint8_t m_UseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage = 0;
		uint8_t m_UseMasterFirm_MinimumRequiredAccountValueIsSet = 0;
		uint8_t m_UseMasterFirm_MinimumRequiredAccountValue = 0;
		uint8_t m_UseMasterFirm_MarginTimeSettingsIsSet = 0;
		uint8_t m_UseMasterFirm_MarginTimeSettings = 0;
		uint8_t m_UseMasterFirm_TradingIsDisabledIsSet = 0;
		uint8_t m_UseMasterFirm_TradingIsDisabled = 0;
		
		uint8_t IsTradeStatisticsPublicallySharedIsSet = 0;
		uint8_t IsTradeStatisticsPublicallyShared = 0;
		uint8_t IsReadOnlyFollowingRequestsAllowedIsSet = 0;
		uint8_t IsReadOnlyFollowingRequestsAllowed = 0;
		uint8_t IsTradeAccountSharingAllowedIsSet = 0;
		uint8_t IsTradeAccountSharingAllowed = 0;
		uint8_t ReadOnlyShareToAllSCUsernamesIsSet = 0;
		uint8_t ReadOnlyShareToAllSCUsernames = 0;

		uint8_t m_UseMasterFirm_SymbolCommissionsArrayIsSet = 0;
		uint8_t m_UseMasterFirm_SymbolCommissionsArray = 0;
		uint8_t m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet = 0;
		uint8_t m_UseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginForDayTradingIsSet = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginForDayTrading = 0;

		uint8_t m_UseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet = 0;
		uint8_t m_UseMasterFirm_SymbolCommissionsArrayFullOverride = 0;
		uint8_t m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet = 0;
		uint8_t m_UseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginFullOverrideIsSet = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginFullOverride = 0;

		uint8_t m_UseMasterFirm_TradeFeesFullOverrideIsSet = 0;
		uint8_t m_UseMasterFirm_TradeFeesFullOverride = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet = 0;
		uint8_t m_UseMasterFirm_UsePercentOfMarginForDayTradingFullOverride = 0;

		uint8_t m_LiquidationOnlyModeIsSet = 0;
		uint8_t m_LiquidationOnlyMode = 0;

		uint8_t m_CustomerOrFirmIsSet = 0;
		uint8_t m_CustomerOrFirm = 0;

		uint8_t m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet = 0;
		uint8_t m_MasterFirm_FlattenCancelAccountWhenDailyLossLimitMet = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMinimumAccountValue = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMarginIntraday = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet = 0;
		uint8_t m_MasterFirm_FlattenCancelWhenUnderMarginEndOfDay = 0;
		uint32_t m_MasterFirm_MaximumOrderQuantityIsSet = 0;
		uint32_t m_MasterFirm_MaximumOrderQuantity = 0;

		uint8_t m_ExchangeTraderIdIsSet = 0;
		DTC_VLS::vls_t m_ExchangeTraderId;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsNewAccount() const;
		uint8_t GetCurrencyCodeIsSet() const;
		uint8_t GetDailyNetLossLimitInAccountCurrencyIsSet() const;
		float GetDailyNetLossLimitInAccountCurrency() const;
		uint8_t GetPercentOfCashBalanceForDailyNetLossLimitIsSet() const;
		int32_t GetPercentOfCashBalanceForDailyNetLossLimit() const;
		uint8_t GetUseTrailingAccountValueToNotAllowIncreaseInPositionsIsSet() const;
		uint8_t GetUseTrailingAccountValueToNotAllowIncreaseInPositions() const;
		uint8_t GetDoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() const;
		uint8_t GetDoNotAllowIncreaseInPositionsAtDailyLossLimit() const;
		uint8_t GetFlattenPositionsAtDailyLossLimitIsSet() const;
		uint8_t GetFlattenPositionsAtDailyLossLimit() const;
		uint8_t GetClosePositionsAtEndOfDayIsSet() const;
		uint8_t GetClosePositionsAtEndOfDay() const;
		uint8_t GetFlattenPositionsWhenUnderMarginIntradayIsSet() const;
		uint8_t GetFlattenPositionsWhenUnderMarginIntraday() const;
		uint8_t GetFlattenPositionsWhenUnderMarginAtEndOfDayIsSet() const;
		uint8_t GetFlattenPositionsWhenUnderMarginAtEndOfDay() const;
		uint8_t GetSenderSubIDIsSet() const;
		uint8_t GetSenderLocationIdIsSet() const;
		uint8_t GetSelfMatchPreventionIDIsSet() const;
		uint8_t GetCTICodeIsSet() const;
		int32_t GetCTICode() const;
		uint8_t GetTradeAccountIsReadOnlyIsSet() const;
		uint8_t GetTradeAccountIsReadOnly() const;
		uint8_t GetMaximumGlobalPositionQuantityIsSet() const;
		int32_t GetMaximumGlobalPositionQuantity() const;
		uint8_t GetTradeFeePerContractIsSet() const;
		double GetTradeFeePerContract() const;
		uint8_t GetTradeFeePerShareIsSet() const;
		double GetTradeFeePerShare() const;
		uint8_t GetRequireSufficientMarginForNewPositionsIsSet() const;
		uint8_t GetRequireSufficientMarginForNewPositions() const;
		uint8_t GetUsePercentOfMarginIsSet() const;
		int32_t GetUsePercentOfMargin() const;
		uint8_t GetUsePercentOfMarginForDayTradingIsSet() const;
		int32_t GetUsePercentOfMarginForDayTrading() const;
		uint8_t GetMaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() const;
		int32_t GetMaximumAllowedAccountBalanceForPositionsAsPercentage() const;
		uint8_t GetFirmIDIsSet() const;
		uint8_t GetTradingIsDisabledIsSet() const;
		uint8_t GetTradingIsDisabled() const;
		uint8_t GetDescriptiveNameIsSet() const;
		uint8_t GetIsMasterFirmControlAccountIsSet() const;
		uint8_t GetIsMasterFirmControlAccount() const;
		uint8_t GetMinimumRequiredAccountValueIsSet() const;
		double GetMinimumRequiredAccountValue() const;
		uint8_t GetBeginTimeForDayMarginIsSet() const;
		int64_t GetBeginTimeForDayMargin() const;
		uint8_t GetEndTimeForDayMarginIsSet() const;
		int64_t GetEndTimeForDayMargin() const;
		uint8_t GetDayMarginTimeZoneIsSet() const;

		uint8_t GetUseMasterFirm_FlattenPositionsWhenUnderMarginIntradayIsSet() const;
		uint8_t GetUseMasterFirm_FlattenPositionsWhenUnderMarginIntraday() const;
		uint8_t GetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDayIsSet() const;
		uint8_t GetUseMasterFirm_FlattenPositionsWhenUnderMarginAtEndOfDay() const;
		uint8_t GetUseMasterFirm_SymbolLimitsArrayIsSet() const;
		uint8_t GetUseMasterFirm_SymbolLimitsArray() const;
		uint8_t GetUseMasterFirm_TradeFeesIsSet() const;
		uint8_t GetUseMasterFirm_TradeFees() const;
		uint8_t GetUseMasterFirm_TradeFeePerShareIsSet() const;
		uint8_t GetUseMasterFirm_TradeFeePerShare() const;
		uint8_t GetUseMasterFirm_RequireSufficientMarginForNewPositionsIsSet() const;
		uint8_t GetUseMasterFirm_RequireSufficientMarginForNewPositions() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginIsSet() const;
		uint8_t GetUseMasterFirm_UsePercentOfMargin() const;
		uint8_t GetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentageIsSet() const;
		uint8_t GetUseMasterFirm_MaximumAllowedAccountBalanceForPositionsAsPercentage() const;
		uint8_t GetUseMasterFirm_MinimumRequiredAccountValueIsSet() const;
		uint8_t GetUseMasterFirm_MinimumRequiredAccountValue() const;
		uint8_t GetUseMasterFirm_MarginTimeSettingsIsSet() const;
		uint8_t GetUseMasterFirm_MarginTimeSettings() const;
		uint8_t GetUseMasterFirm_TradingIsDisabledIsSet() const;
		uint8_t GetUseMasterFirm_TradingIsDisabled() const;

		uint8_t GetIsTradeStatisticsPublicallySharedIsSet() const;
		uint8_t GetIsTradeStatisticsPublicallyShared() const;
		uint8_t GetIsReadOnlyFollowingRequestsAllowedIsSet() const;
		uint8_t GetIsReadOnlyFollowingRequestsAllowed() const;
		uint8_t GetIsTradeAccountSharingAllowedIsSet() const;
		uint8_t GetIsTradeAccountSharingAllowed() const;
		uint8_t GetReadOnlyShareToAllSCUsernamesIsSet() const;
		uint8_t GetReadOnlyShareToAllSCUsernames() const;

		uint8_t GetUseMasterFirm_SymbolCommissionsArrayIsSet() const;
		uint8_t GetUseMasterFirm_SymbolCommissionsArray() const;
		uint8_t GetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimitIsSet() const;
		uint8_t GetUseMasterFirm_DoNotAllowIncreaseInPositionsAtDailyLossLimit() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginForDayTradingIsSet() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginForDayTrading() const;
		uint8_t GetUseMasterFirm_SymbolCommissionsArrayFullOverrideIsSet() const;
		uint8_t GetUseMasterFirm_SymbolCommissionsArrayFullOverride() const;
		uint8_t GetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrdersIsSet() const;
		uint8_t GetUseMasterFirm_NumDaysBeforeLastTradingDateToDisallowOrders() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginFullOverrideIsSet() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginFullOverride() const;
		uint8_t GetUseMasterFirm_TradeFeesFullOverrideIsSet() const;
		uint8_t GetUseMasterFirm_TradeFeesFullOverride() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverrideIsSet() const;
		uint8_t GetUseMasterFirm_UsePercentOfMarginForDayTradingFullOverride() const;
		uint8_t GetLiquidationOnlyModeIsSet() const;
		uint8_t GetLiquidationOnlyMode() const;

		uint8_t GetCustomerOrFirmIsSet() const;
		uint8_t GetCustomerOrFirm() const;

		uint8_t GetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMetIsSet() const;
		uint8_t GetMasterFirm_FlattenCancelAccountWhenDailyLossLimitMet() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMinimumAccountValueIsSet() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMinimumAccountValue() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMarginIntradayIsSet() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMarginIntraday() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMarginEndOfDayIsSet() const;
		uint8_t GetMasterFirm_FlattenCancelWhenUnderMarginEndOfDay() const;
		uint32_t GetMasterFirm_MaximumOrderQuantityIsSet() const;
		uint32_t GetMasterFirm_MaximumOrderQuantity() const;
		uint8_t GetExchangeTraderIdIsSet() const;

		//-------------------------------------------------------------------
		const char* GetNewAccountAuthorizedUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, NewAccountAuthorizedUsername, offsetof(s_TradeAccountDataUpdate, NewAccountAuthorizedUsername));
		}

		void AddNewAccountAuthorizedUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, NewAccountAuthorizedUsername, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataUpdate, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetCurrencyCode() const
		{
			return GetVariableLengthStringField(Size, BaseSize, CurrencyCode, offsetof(s_TradeAccountDataUpdate, CurrencyCode));
		}

		void AddCurrencyCode(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, CurrencyCode, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSenderSubID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, SenderSubID, offsetof(s_TradeAccountDataUpdate, SenderSubID));
		}

		void AddSenderSubID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, SenderSubID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSenderLocationId() const
		{
			return GetVariableLengthStringField(Size, BaseSize, SenderLocationId, offsetof(s_TradeAccountDataUpdate, SenderLocationId));
		}

		void AddSenderLocationId(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, SenderLocationId, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSelfMatchPreventionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, SelfMatchPreventionID, offsetof(s_TradeAccountDataUpdate, SelfMatchPreventionID));
		}

		void AddSelfMatchPreventionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, SelfMatchPreventionID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetFirmID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FirmID, offsetof(s_TradeAccountDataUpdate, FirmID));
		}

		void AddFirmID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FirmID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetDescriptiveName() const
		{
			return GetVariableLengthStringField(Size, BaseSize, DescriptiveName, offsetof(s_TradeAccountDataUpdate, DescriptiveName));
		}

		void AddDescriptiveName(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, DescriptiveName, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetDayMarginTimeZone() const
		{
			return GetVariableLengthStringField(Size, BaseSize, DayMarginTimeZone, offsetof(s_TradeAccountDataUpdate, DayMarginTimeZone));
		}

		void AddDayMarginTimeZone(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, DayMarginTimeZone, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetExchangeTraderId() const
		{
			return GetVariableLengthStringField(Size, BaseSize, m_ExchangeTraderId, offsetof(s_TradeAccountDataUpdate, m_ExchangeTraderId));
		}

		void AddExchangeTraderId(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, m_ExchangeTraderId, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataDelete
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_DELETE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		DTC_VLS::vls_t TradeAccount;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataDelete, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_TradeAccountDataSymbolLimitsResponse
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_RESPONSE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;

		DTC_VLS::vls_t TradeAccount;

		DTC_VLS::vls_t Symbol;
		double TradePositionLimit = 0;
		double OrderQuantityLimit = 0;
		int32_t UsePercentOfMargin = 0;
		uint8_t OverrideMarginOtherAccounts = 0;
		int32_t UsePercentOfMarginForDayTrading = 0;
		int32_t NumberOfDaysBeforeLastTradingDateToDisallowOrders = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		double GetTradePositionLimit() const;
		double GetOrderQuantityLimit() const;
		int32_t GetUsePercentOfMargin() const;
		uint8_t GetOverrideMarginOtherAccounts() const;
		int32_t GetUsePercentOfMarginForDayTrading() const;
		int32_t GetNumberOfDaysBeforeLastTradingDateToDisallowOrders() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataSymbolLimitsResponse, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_TradeAccountDataSymbolLimitsResponse, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataSymbolLimitsUpdate
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		uint8_t IsDeleted = 0;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t Symbol;
		double TradePositionLimit = 0;
		double OrderQuantityLimit = 0;
		int32_t UsePercentOfMargin = 100;
		uint8_t OverrideMarginOtherAccounts = 0;
		int32_t UsePercentOfMarginForDayTrading = 0;
		int32_t NumberOfDaysBeforeLastTradingDateToDisallowOrders = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsDeleted() const;
		double GetTradePositionLimit() const;
		double GetOrderQuantityLimit() const;
		int32_t GetUsePercentOfMargin() const;
		uint8_t GetOverrideMarginOtherAccounts() const;
		int32_t GetUsePercentOfMarginForDayTrading() const;
		int32_t GetNumberOfDaysBeforeLastTradingDateToDisallowOrders() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataSymbolLimitsUpdate, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_TradeAccountDataSymbolLimitsUpdate, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataSymbolCommissionResponse
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_RESPONSE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t Symbol;
		double TradeFeePerContract = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		double GetTradeFeePerContract() const;
		
		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataSymbolCommissionResponse, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_TradeAccountDataSymbolCommissionResponse, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataSymbolCommissionUpdate
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_SYMBOL_COMMISSION_UPDATE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		uint8_t IsDeleted = 0;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t Symbol;
		double TradeFeePerContract = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsDeleted() const;
		double GetTradeFeePerContract() const;


		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataSymbolCommissionUpdate, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_TradeAccountDataSymbolCommissionUpdate, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataAuthorizedUsernameResponse
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_RESPONSE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t Username;

		uint8_t UpdateOperationComplete = 0;
		uint16_t UpdateOperationMessageType = 0;
		DTC_VLS::vls_t UpdateOperationErrorText;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetUpdateOperationComplete() const;
		uint16_t GetUpdateOperationMessageType() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataAuthorizedUsernameResponse, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_TradeAccountDataAuthorizedUsernameResponse, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUpdateOperationErrorText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, UpdateOperationErrorText, offsetof(s_TradeAccountDataAuthorizedUsernameResponse, UpdateOperationErrorText));
		}

		void AddUpdateOperationErrorText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, UpdateOperationErrorText, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataAuthorizedUsernameAdd
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_ADD;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;

		DTC_VLS::vls_t TradeAccount;

		DTC_VLS::vls_t Username;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;


		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataAuthorizedUsernameAdd, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_TradeAccountDataAuthorizedUsernameAdd, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataAuthorizedUsernameRemove
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_AUTHORIZED_USERNAME_REMOVE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t Username;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataAuthorizedUsernameRemove, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_TradeAccountDataAuthorizedUsernameRemove, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataUsernameToShareWithResponse
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_RESPONSE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;

		DTC_VLS::vls_t TradeAccount;

		DTC_VLS::vls_t Username;
		uint8_t IsReadOnly = 0;

		uint8_t UpdateOperationComplete = 0;
		uint16_t UpdateOperationMessageType = 0;
		DTC_VLS::vls_t UpdateOperationErrorText;


		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsReadOnly() const;
		uint8_t GetUpdateOperationComplete() const;
		uint16_t GetUpdateOperationMessageType() const;


		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataUsernameToShareWithResponse, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_TradeAccountDataUsernameToShareWithResponse, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUpdateOperationErrorText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, UpdateOperationErrorText, offsetof(s_TradeAccountDataUsernameToShareWithResponse, UpdateOperationErrorText));
		}

		void AddUpdateOperationErrorText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, UpdateOperationErrorText, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataUsernameToShareWithAdd
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_ADD;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;

		DTC_VLS::vls_t TradeAccount;

		DTC_VLS::vls_t Username;
		uint8_t IsReadOnly = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsReadOnly() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataUsernameToShareWithAdd, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_TradeAccountDataUsernameToShareWithAdd, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataUsernameToShareWithRemove
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_USERNAME_TO_SHARE_WITH_REMOVE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t Username;
		uint8_t IsReadOnly = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsReadOnly() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataUsernameToShareWithRemove, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_TradeAccountDataUsernameToShareWithRemove, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};
	
	/*==========================================================================*/
	struct s_TradeAccountDataResponseTrailer
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_RESPONSE_TRAILER;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		uint8_t m_IsSnapshot = 0;
		uint8_t m_IsFirstMessageInBatch = 0;
		uint8_t m_IsLastMessageInBatch = 0;
		DTC_VLS::vls_t TradeAccount;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsSnapshot() const;
		uint8_t GetIsFirstMessageInBatch() const;
		uint8_t GetIsLastMessageInBatch() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataResponseTrailer, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
	};

	/*==========================================================================*/
	struct s_TradeAccountDataUpdateOperationComplete
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::TRADE_ACCOUNT_DATA_UPDATE_OPERATION_COMPLETE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t RequestID = 0;
		uint8_t IsError = 0;
		DTC_VLS::vls_t ErrorText;
		uint8_t IsDeleteOperation = 0;
		uint8_t IsSymbolLimitsUpdateOperation = 0;
		uint8_t IsSymbolCommissionUpdateOperation = 0;
		DTC_VLS::vls_t TradeAccount;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		uint8_t GetIsError() const;
		uint8_t GetIsDeleteOperation() const;
		uint8_t GetIsSymbolLimitsUpdateOperation() const;
		uint8_t GetIsSymbolCommissionUpdateOperation() const;

		//-------------------------------------------------------------------
		const char* GetErrorText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ErrorText, offsetof(s_TradeAccountDataUpdateOperationComplete, ErrorText));
		}

		void AddErrorText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ErrorText, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountDataUpdateOperationComplete, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//--End of Get Add functions----------------------------------------

	};

	/*==========================================================================*/
	struct s_ProcessedFillIdentifier
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::PROCESSED_FILL_IDENTIFIER;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		DTC_VLS::vls_t FillIdentifier;

		uint8_t IsSnapshot = 0;
		uint8_t IsFirstMessageInBatch = 0;
		uint8_t IsLastMessageInBatch = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		//-------------------------------------------------------------------
		const char* GetFillIdentifier() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FillIdentifier, offsetof(s_ProcessedFillIdentifier, FillIdentifier));
		}

		void AddFillIdentifier(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FillIdentifier, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_FlattenPositionsForTradeAccount
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = DTC::FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t ClientOrderID;
		DTC_VLS::vls_t FreeFormText;
		uint8_t IsAutomatedOrder = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_FlattenPositionsForTradeAccount, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_FlattenPositionsForTradeAccount, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetFreeFormText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_FlattenPositionsForTradeAccount, FreeFormText));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_UserInformation
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::USER_INFORMATION;
		uint16_t BaseSize = sizeof(*this);

		// Message fields
		DTC_VLS::vls_t OperatorID;
		DTC_VLS::vls_t LocationID;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		//-------------------------------------------------------------------
		const char* GetOperatorID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, OperatorID, offsetof(s_UserInformation, OperatorID));
		}

		void AddOperatorID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, OperatorID, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetLocationID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, LocationID, offsetof(s_UserInformation, LocationID));
		}

		void AddLocationID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, LocationID, StringLength);
		}
		//-------------------------------------------------------------------
	};

	/*==========================================================================*/
	struct s_MarginDataRequest
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::MARGIN_DATA_REQUEST;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		int32_t RequestID = 0;
		DTC_VLS::vls_t TradeAccount;
		DTC_VLS::vls_t Symbol;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;

		//-------------------------------------------------------------------
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_MarginDataRequest, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_MarginDataRequest, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
		//-------------------------------------------------------------------
	};

	/*==========================================================================*/
	struct s_MarginDataResponse
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = n_DTCNonStandard::MARGIN_DATA_RESPONSE;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		int32_t RequestID = 0;
		DTC_VLS::vls_t ErrorText;
		double InitialExchangeMargin = 0;
		double MaintenanceExchangeMargin = 0;
		double InitialAccountMargin = 0;
		double MaintenanceAccountMargin = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetRequestID() const;
		double GetInitialExchangeMargin() const;
		double GetMaintenanceExchangeMargin() const;
		double GetInitialAccountMargin() const;
		double GetMaintenanceAccountMargin() const;

		//-------------------------------------------------------------------
		const char* GetErrorText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ErrorText, offsetof(s_MarginDataResponse, ErrorText));
		}

		void AddErrorText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ErrorText, StringLength);
		}
		//-------------------------------------------------------------------
	};

	
#pragma pack(pop)
}
