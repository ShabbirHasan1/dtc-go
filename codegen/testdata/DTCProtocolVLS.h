
// Data and Trading Communications Protocol Variable Length String Encoding (DTC Protocol VLS)

// This protocol is in the public domain and freely usable by anyone.

// Documentation: http://DTCprotocol.org/index.php?page=doc_DTCMessageDocumentation.php

// The uint8_t ordering is little endian.

#pragma once

#include "DTCProtocol.h"

namespace DTC_VLS
{
#pragma pack(8)

	// see DTCProtocol.h for constants and enum values

	/*==========================================================================*/
	enum TradeModeEnum : int8_t
	{ TRADE_MODE_UNSET = 0
	, TRADE_MODE_DEMO = 1
	, TRADE_MODE_SIMULATED = 2
	, TRADE_MODE_LIVE = 3
	};

	/*==========================================================================*/
	struct s_VariableLengthStringField
	{
		uint16_t Offset = 0;
		uint16_t Length = 0;
	};

	typedef s_VariableLengthStringField vls_t;

	/*==========================================================================*/
	inline void AddVariableLengthStringField(uint16_t& r_BaseStructureSizeField, vls_t& r_VariableLengthStringField, uint16_t SizeToAdd)
	{
		if (SizeToAdd == 0)
		{
			r_VariableLengthStringField.Offset = 0;
			r_VariableLengthStringField.Length = 0;
		}
		else
		{
			r_VariableLengthStringField.Offset = r_BaseStructureSizeField;
			r_VariableLengthStringField.Length = SizeToAdd + 1;
			r_BaseStructureSizeField += r_VariableLengthStringField.Length;
		}
	}

	/*==========================================================================*/
	inline const char* GetVariableLengthStringField
		( const uint16_t& MessageSizeField
		, uint16_t BaseStructureSizeField
		, const vls_t& VariableLengthStringField
		, const uint16_t VariableLengthStringFieldOffset
	)
	{
		if (BaseStructureSizeField < VariableLengthStringFieldOffset + sizeof(vls_t))
			return "";
		else if (VariableLengthStringField.Offset == 0 || VariableLengthStringField.Length <= 1)
			return "";
		else if ((VariableLengthStringField.Offset + VariableLengthStringField.Length) > MessageSizeField)
			return "";
		else
		{
			int Length = VariableLengthStringField.Length;
			if (Length > 4096)
			{
				Length = 4096;
			}

			char * BaseMessage = reinterpret_cast< char *>( const_cast<uint16_t *>(&MessageSizeField));
			char * TerminatorCharacter = BaseMessage + VariableLengthStringField.Offset + Length - 1;

			*TerminatorCharacter = 0;//Make sure there is a null terminator here.

			return (const char*) BaseMessage + VariableLengthStringField.Offset;

		}
	}

	/*==========================================================================*/
	inline uint16_t GetStringFieldSendLength(uint32_t SendLength)
	{
		uint16_t ActualSendLength = static_cast<uint16_t>(SendLength);
		
		++ActualSendLength;

		return ActualSendLength;
	}

	/*==========================================================================*/
	struct s_LogonRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;
		int32_t ProtocolVersion;
		vls_t Username;
		vls_t Password;
		vls_t GeneralTextData;
		int32_t Integer_1;
		int32_t Integer_2;
		int32_t HeartbeatIntervalInSeconds; 
		DTC::TradeModeEnum TradeMode;
		vls_t TradeAccount;
		vls_t HardwareIdentifier;
		vls_t ClientName;
		int32_t MarketDataTransmissionInterval;

		s_LogonRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_LogonRequest));
			Size = sizeof(s_LogonRequest);
			Type = DTC::LOGON_REQUEST;
			BaseSize = Size;

			ProtocolVersion = DTC::CURRENT_VERSION;
		}

		int32_t GetProtocolVersion() const;

		const char* GetUsername() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Username, offsetof(s_LogonRequest, Username));
		}

		void AddUsername(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Username, StringLength);
		}

		const char* GetPassword() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Password, offsetof(s_LogonRequest, Password));
		}

		void AddPassword(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Password, StringLength);
		}

		const char* GetGeneralTextData() const
		{
			return GetVariableLengthStringField(Size, BaseSize, GeneralTextData, offsetof(s_LogonRequest, GeneralTextData));
		}

		void AddGeneralTextData(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, GeneralTextData, StringLength);
		}

		int32_t GetInteger_1() const;
		int32_t GetInteger_2() const;
		int32_t GetHeartbeatIntervalInSeconds() const;
		DTC::TradeModeEnum GetTradeMode() const;
		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_LogonRequest, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetHardwareIdentifier() const
		{
			return GetVariableLengthStringField(Size, BaseSize, HardwareIdentifier, offsetof(s_LogonRequest, HardwareIdentifier));
		}

		void AddHardwareIdentifier(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, HardwareIdentifier, StringLength);
		}

		const char* GetClientName() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientName, offsetof(s_LogonRequest, ClientName));
		}

		void AddClientName(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientName, StringLength);
		}

		int32_t GetMarketDataTransmissionInterval() const;

	};

	/*==========================================================================*/
	struct s_LogonResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;
		int32_t ProtocolVersion;
		DTC::LogonStatusEnum Result;
		vls_t ResultText;
		vls_t ReconnectAddress;
		int32_t Integer_1;
		vls_t ServerName;
		uint8_t MarketDepthUpdatesBestBidAndAsk;
		uint8_t TradingIsSupported;
		uint8_t OCOOrdersSupported;
		uint8_t OrderCancelReplaceSupported;
		vls_t SymbolExchangeDelimiter;
		uint8_t SecurityDefinitionsSupported;
		uint8_t HistoricalPriceDataSupported;
		uint8_t ResubscribeWhenMarketDataFeedAvailable;
		uint8_t MarketDepthIsSupported;
		uint8_t OneHistoricalPriceDataRequestPerConnection;
		uint8_t BracketOrdersSupported;
		uint8_t UseIntegerPriceOrderMessages;
		uint8_t UsesMultiplePositionsPerSymbolAndTradeAccount;
		uint8_t MarketDataSupported;

		s_LogonResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_LogonResponse));
			Size = sizeof(s_LogonResponse);
			Type = DTC::LOGON_RESPONSE;
			BaseSize = Size;

			ProtocolVersion = DTC::CURRENT_VERSION;
			OrderCancelReplaceSupported = 1;
			MarketDepthIsSupported = 1;
		}

		int32_t GetProtocolVersion() const;
		DTC::LogonStatusEnum GetResult() const;
		int32_t GetInteger_1() const;
		uint8_t GetMarketDepthUpdatesBestBidAndAsk() const;
		uint8_t GetTradingIsSupported() const;
		uint8_t GetOCOOrdersSupported() const;
		uint8_t GetOrderCancelReplaceSupported() const;
		uint8_t GetSecurityDefinitionsSupported() const;
		uint8_t GetHistoricalPriceDataSupported() const;
		uint8_t GetResubscribeWhenMarketDataFeedAvailable() const;
		uint8_t GetMarketDepthIsSupported() const;
		uint8_t GetOneHistoricalPriceDataRequestPerConnection() const;
		uint8_t GetUseIntegerPriceOrderMessages() const;
		uint8_t GetBracketOrdersSupported() const;
		uint8_t GetUsesMultiplePositionsPerSymbolAndTradeAccount() const;
		uint8_t GetMarketDataSupported() const;

		const char* GetResultText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ResultText, offsetof(s_LogonResponse, ResultText));
		}

		void AddResultText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ResultText, StringLength);
		}

		const char* GetReconnectAddress() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ReconnectAddress, offsetof(s_LogonResponse, ReconnectAddress));
		}

		void AddReconnectAddress(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ReconnectAddress, StringLength);
		}

		const char* GetServerName() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerName, offsetof(s_LogonResponse, ServerName));
		}

		void AddServerName(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerName, StringLength);
		}

		const char* GetSymbolExchangeDelimiter() const
		{
			return GetVariableLengthStringField(Size, BaseSize, SymbolExchangeDelimiter, offsetof(s_LogonResponse, SymbolExchangeDelimiter));
		}

		void AddSymbolExchangeDelimiter(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, SymbolExchangeDelimiter, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_Logoff
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;
		vls_t Reason;
		uint8_t DoNotReconnect;

		s_Logoff()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_Logoff));
			Size = sizeof(s_Logoff);
			Type = DTC::LOGOFF;
			BaseSize = Size;
		}

		const char* GetReason() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Reason, offsetof(s_Logoff, Reason));
		}

		void AddReason(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Reason, StringLength);
		}

		uint8_t GetDoNotReconnect() const;
	};


	/*==========================================================================*/
	struct s_MarketDataRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;
		DTC::RequestActionEnum RequestAction;
		uint32_t SymbolID;
		vls_t Symbol;
		vls_t Exchange;
		uint32_t IntervalForSnapshotUpdatesInMilliseconds;

		s_MarketDataRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_MarketDataRequest));
			Size = sizeof(s_MarketDataRequest);
			Type = DTC::MARKET_DATA_REQUEST;
			BaseSize = Size;

			RequestAction = DTC::SUBSCRIBE;
		}

		DTC::RequestActionEnum GetRequestAction() const;
		uint32_t GetSymbolID() const;

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_MarketDataRequest, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_MarketDataRequest, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		uint32_t GetIntervalForSnapshotUpdatesInMilliseconds() const;
	};

	/*==========================================================================*/
	struct s_MarketDepthRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;
		DTC::RequestActionEnum RequestAction;
		uint32_t SymbolID;
		vls_t Symbol;
		vls_t Exchange;
		int32_t NumLevels;

		s_MarketDepthRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_MarketDepthRequest));
			Size = sizeof(s_MarketDepthRequest);
			Type = DTC::MARKET_DEPTH_REQUEST;
			BaseSize = Size;

			RequestAction = DTC::SUBSCRIBE;
			NumLevels = 10;
		}

		DTC::RequestActionEnum GetRequestAction() const;
		uint32_t GetSymbolID() const;

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_MarketDepthRequest, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_MarketDepthRequest, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		int32_t GetNumLevels() const;
	};

	/*==========================================================================*/
	struct s_MarketDataReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;
		uint32_t SymbolID;
		vls_t RejectText;

		s_MarketDataReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_MarketDataReject));
			Size = sizeof(s_MarketDataReject);
			Type = DTC::MARKET_DATA_REJECT;
			BaseSize = Size;
		}

		uint32_t GetSymbolID() const;

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_MarketDataReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
	};


	/*==========================================================================*/
	struct s_MarketDepthReject
	{
		uint16_t Size = sizeof(*this);
		uint16_t Type = DTC::MARKET_DEPTH_REJECT;
		uint16_t BaseSize = sizeof(*this);
		uint32_t SymbolID = 0;
		vls_t RejectText;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetSymbolID() const;

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_MarketDepthReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_MarketOrdersRequest
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = DTC::MARKET_ORDERS_REQUEST;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		DTC::RequestActionEnum RequestAction;
		uint32_t SymbolID = 0;
		vls_t Symbol;
		vls_t Exchange;
		uint32_t SendQuantitiesGreaterOrEqualTo = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		DTC::RequestActionEnum GetRequestAction() const;
		uint32_t GetSymbolID() const;
		uint32_t GetSendQuantitiesGreaterOrEqualTo() const;

		//Implementing Get Add inline functions------------------------------
		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_MarketOrdersRequest, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
		//-------------------------------------------------------------------
		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_MarketOrdersRequest, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_MarketOrdersReject
	{
		//standard DTC VLS header
		uint16_t Size = sizeof(*this);
		uint16_t Type = DTC::MARKET_ORDERS_REJECT;
		uint16_t BaseSize = sizeof(*this);

		//message field variables
		uint32_t SymbolID = 0;
		vls_t RejectText;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		uint32_t GetSymbolID() const;

		//Implementing Get Add inline functions------------------------------
		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_MarketOrdersReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
		//--End of Get Add functions----------------------------------------
	};

	/*==========================================================================*/
	struct s_SubmitNewSingleOrder
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t Symbol;
		vls_t Exchange;

		vls_t TradeAccount;
		vls_t ClientOrderID;

		DTC::OrderTypeEnum OrderType;

		DTC::BuySellEnum BuySell;

		double Price1;
		double Price2;

		double Quantity;

		DTC::TimeInForceEnum TimeInForce;

		DTC::t_DateTime GoodTillDateTime;

		uint8_t IsAutomatedOrder;

		uint8_t IsParentOrder;

		vls_t FreeFormText;

		DTC::OpenCloseTradeEnum OpenOrClose;
		double MaxShowQuantity;

		vls_t Price1AsString;
		vls_t Price2AsString;

		s_SubmitNewSingleOrder()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SubmitNewSingleOrder));
			Size = sizeof(s_SubmitNewSingleOrder);
			Type = DTC::SUBMIT_NEW_SINGLE_ORDER;
			BaseSize = Size;
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SubmitNewSingleOrder, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SubmitNewSingleOrder, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_SubmitNewSingleOrder, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_SubmitNewSingleOrder, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		DTC::OrderTypeEnum GetOrderType() const;
		DTC::BuySellEnum GetBuySell() const;
		double GetPrice1() const;
		double GetPrice2() const;
		double GetQuantity() const;
		DTC::TimeInForceEnum GetTimeInForce() const;
		DTC::t_DateTime GetGoodTillDateTime() const;
		void SetTradeAccount(const char* NewValue);
		uint8_t GetIsAutomatedOrder() const;
		uint8_t GetIsParentOrder() const;

		const char* GetFreeFormText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_SubmitNewSingleOrder, FreeFormText));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}

		DTC::OpenCloseTradeEnum GetOpenOrClose() const;
		double GetMaxShowQuantity() const;

		const char* GetPrice1AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price1AsString, offsetof(s_SubmitNewSingleOrder, Price1AsString));
		}

		void AddPrice1AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price1AsString, StringLength);
		}

		const char* GetPrice2AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price2AsString, offsetof(s_SubmitNewSingleOrder, Price2AsString));
		}

		void AddPrice2AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price2AsString, StringLength);
		}

	};

	/*==========================================================================*/
	struct s_SubmitNewSingleOrderInt
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t Symbol;
		vls_t Exchange;

		vls_t TradeAccount;
		vls_t ClientOrderID;

		DTC::OrderTypeEnum OrderType;
		DTC::BuySellEnum BuySell;

		int64_t Price1;
		int64_t Price2;
		float Divisor;
		int64_t Quantity;

		DTC::TimeInForceEnum TimeInForce;
		DTC::t_DateTime GoodTillDateTime;

		uint8_t IsAutomatedOrder;
		uint8_t IsParentOrder;

		vls_t FreeFormText;

		DTC::OpenCloseTradeEnum OpenOrClose;

		s_SubmitNewSingleOrderInt()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SubmitNewSingleOrderInt));
			Size = sizeof(s_SubmitNewSingleOrderInt);
			Type = DTC::SUBMIT_NEW_SINGLE_ORDER_INT;
			BaseSize = Size;
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SubmitNewSingleOrderInt, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SubmitNewSingleOrderInt, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_SubmitNewSingleOrderInt, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_SubmitNewSingleOrderInt, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		DTC::OrderTypeEnum GetOrderType() const;
		DTC::BuySellEnum GetBuySell() const;
		int64_t GetPrice1() const;
		int64_t GetPrice2() const;
		float GetDivisor() const;
		int64_t GetQuantity() const;
		DTC::TimeInForceEnum GetTimeInForce() const;
		DTC::t_DateTime GetGoodTillDateTime() const;
		uint8_t GetIsAutomatedOrder() const;
		uint8_t GetIsParentOrder() const;

		const char* GetFreeFormText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_SubmitNewSingleOrderInt, FreeFormText));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}

		DTC::OpenCloseTradeEnum GetOpenOrClose() const;
	};

	/*==========================================================================*/
	struct s_SubmitFlattenPositionOrder
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t Symbol;
		vls_t Exchange;

		vls_t TradeAccount;
		vls_t ClientOrderID;
		vls_t FreeFormText;
		uint8_t IsAutomatedOrder;

		s_SubmitFlattenPositionOrder()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SubmitFlattenPositionOrder));
			Size = sizeof(s_SubmitFlattenPositionOrder);
			Type = DTC::SUBMIT_FLATTEN_POSITION_ORDER;
			BaseSize = Size;
		}

		const char* GetSymbol()
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SubmitFlattenPositionOrder, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange()
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SubmitFlattenPositionOrder, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetTradeAccount()
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_SubmitFlattenPositionOrder, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetClientOrderID()
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_SubmitFlattenPositionOrder, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		const char* GetFreeFormText()
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_SubmitFlattenPositionOrder, FreeFormText));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}

		uint8_t GetIsAutomatedOrder() const;
	};

	/*==========================================================================*/
	struct s_CancelReplaceOrder
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t ServerOrderID;
		vls_t ClientOrderID;

		double Price1;
		double Price2;

		double Quantity;

		int8_t Price1IsSet;
		int8_t Price2IsSet;

		int32_t Unused;
		DTC::TimeInForceEnum TimeInForce;
		DTC::t_DateTime GoodTillDateTime;
		uint8_t UpdatePrice1OffsetToParent;
		vls_t TradeAccount;

		vls_t Price1AsString;
		vls_t Price2AsString;

		s_CancelReplaceOrder()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_CancelReplaceOrder));
			Size = sizeof(s_CancelReplaceOrder);
			Type = DTC::CANCEL_REPLACE_ORDER;
			BaseSize = Size;

			Price1IsSet = 1;
			Price2IsSet = 1;
		}

		const char* GetServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerOrderID, offsetof(s_CancelReplaceOrder, ServerOrderID));
		}

		void AddServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerOrderID, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_CancelReplaceOrder, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_CancelReplaceOrder, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetPrice1AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price1AsString, offsetof(s_CancelReplaceOrder, Price1AsString));
		}

		void AddPrice1AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price1AsString, StringLength);
		}

		const char* GetPrice2AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price2AsString, offsetof(s_CancelReplaceOrder, Price2AsString));
		}

		void AddPrice2AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price2AsString, StringLength);
		}

		double GetPrice1() const;
		double GetPrice2() const;
		double GetQuantity() const;
		int8_t GetPrice1IsSet() const;
		int8_t GetPrice2IsSet() const;
		DTC::TimeInForceEnum GetTimeInForce() const;
		DTC::t_DateTime GetGoodTillDateTime() const;
		uint8_t GetUpdatePrice1OffsetToParent() const;
	};

	/*==========================================================================*/
	struct s_CancelReplaceOrderInt
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t ServerOrderID;
		vls_t ClientOrderID;

		int64_t Price1;
		int64_t Price2;
		float Divisor;
		int64_t Quantity;
		int8_t Price1IsSet;
		int8_t Price2IsSet;

		int32_t Unused;
		DTC::TimeInForceEnum TimeInForce;
		DTC::t_DateTime GoodTillDateTime;
		uint8_t UpdatePrice1OffsetToParent;

		s_CancelReplaceOrderInt()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_CancelReplaceOrderInt));
			Size = sizeof(s_CancelReplaceOrderInt);
			Type = DTC::CANCEL_REPLACE_ORDER_INT;
			BaseSize = Size;

			Divisor = 1.0f;
			Price1IsSet = 1;
			Price2IsSet = 1;
		}

		const char* GetServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerOrderID, offsetof(s_CancelReplaceOrderInt, ServerOrderID));
		}

		void AddServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerOrderID, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_CancelReplaceOrderInt, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		int64_t GetPrice1() const;
		int64_t GetPrice2() const;
		float GetDivisor() const;
		int64_t GetQuantity() const;
		int8_t GetPrice1IsSet() const;
		int8_t GetPrice2IsSet() const;
		DTC::TimeInForceEnum GetTimeInForce() const;
		DTC::t_DateTime GetGoodTillDateTime() const;
		uint8_t GetUpdatePrice1OffsetToParent() const;
	};

	/*==========================================================================*/
	struct s_CancelOrder
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t ServerOrderID;
		vls_t ClientOrderID;
		vls_t TradeAccount;

		s_CancelOrder()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_CancelOrder));
			Size = sizeof(s_CancelOrder);
			Type = DTC::CANCEL_ORDER;
			BaseSize = Size;
		}

		const char* GetServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerOrderID, offsetof(s_CancelOrder, ServerOrderID));
		}

		void AddServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerOrderID, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_CancelOrder, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_CancelOrder, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_SubmitNewOCOOrder
	{
		uint16_t Size = 0;
		uint16_t Type = 0;
		uint16_t BaseSize = 0;

		vls_t Symbol;
		vls_t Exchange;

		vls_t ClientOrderID_1;
		DTC::OrderTypeEnum OrderType_1 = DTC::ORDER_TYPE_UNSET;
		DTC::BuySellEnum BuySell_1 = DTC::BUY_SELL_UNSET;
		double Price1_1 = 0;
		double Price2_1 = 0;
		double Quantity_1 = 0;

		vls_t ClientOrderID_2;
		DTC::OrderTypeEnum OrderType_2 = DTC::ORDER_TYPE_UNSET;
		DTC::BuySellEnum BuySell_2 = DTC::BUY_SELL_UNSET;
		double Price1_2 = 0;
		double Price2_2 = 0;
		double Quantity_2 = 0;

		DTC::TimeInForceEnum TimeInForce = DTC::TIF_UNSET;
		DTC::t_DateTime GoodTillDateTime = 0;

		vls_t TradeAccount;

		uint8_t IsAutomatedOrder = 0;

		vls_t ParentTriggerClientOrderID;

		vls_t FreeFormText;

		DTC::OpenCloseTradeEnum OpenOrClose = DTC::TRADE_UNSET;

		DTC::PartialFillHandlingEnum PartialFillHandling = DTC::PARTIAL_FILL_UNSET;

		uint8_t UseOffsets = 0;
		double OffsetFromParent1 = 0;
		double OffsetFromParent2 = 0;
		uint8_t MaintainSamePricesOnParentFill = 0;

		vls_t Price1_1AsString;
		vls_t Price2_1AsString;
		vls_t Price1_2AsString;
		vls_t Price2_2AsString;

		s_SubmitNewOCOOrder()
		{
			Size = sizeof(s_SubmitNewOCOOrder);
			Type = DTC::SUBMIT_NEW_OCO_ORDER;
			BaseSize = Size;
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		const char* GetClientOrderID_1() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID_1, offsetof(s_SubmitNewOCOOrder, ClientOrderID_1));
		}

		void AddClientOrderID_1(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID_1, StringLength);
		}

		const char* GetClientOrderID_2() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID_2, offsetof(s_SubmitNewOCOOrder, ClientOrderID_2));
		}

		void AddClientOrderID_2(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID_2, StringLength);
		}

		const char* GetFreeFormText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_SubmitNewOCOOrder, FreeFormText));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SubmitNewOCOOrder, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SubmitNewOCOOrder, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		void AddParentTriggerClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ParentTriggerClientOrderID, StringLength);
		}

		const char* GetParentTriggerClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ParentTriggerClientOrderID, offsetof(s_SubmitNewOCOOrder, ParentTriggerClientOrderID));
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_SubmitNewOCOOrder, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}


		const char* GetPrice1_1AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price1_1AsString, offsetof(s_SubmitNewOCOOrder, Price1_1AsString));
		}

		void AddPrice1_1AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price1_1AsString, StringLength);
		}

		const char* GetPrice2_1AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price2_1AsString, offsetof(s_SubmitNewOCOOrder, Price2_1AsString));
		}

		void AddPrice2_1AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price2_1AsString, StringLength);
		}

		const char* GetPrice1_2AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price1_2AsString, offsetof(s_SubmitNewOCOOrder, Price1_2AsString));
		}

		void AddPrice1_2AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price1_2AsString, StringLength);
		}

		const char* GetPrice2_2AsString() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Price2_2AsString, offsetof(s_SubmitNewOCOOrder, Price2_2AsString));
		}

		void AddPrice2_2AsString(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Price2_2AsString, StringLength);
		}

		DTC::OrderTypeEnum GetOrderType_1() const;
		DTC::OrderTypeEnum GetOrderType_2() const;
		DTC::BuySellEnum GetBuySell_1() const;
		DTC::BuySellEnum GetBuySell_2() const;
		DTC::TimeInForceEnum GetTimeInForce() const;
		DTC::t_DateTime GetGoodTillDateTime() const;
		uint8_t GetIsAutomatedOrder() const;
		double GetPrice1_1() const;
		double GetPrice2_1() const;
		double GetPrice1_2() const;
		double GetPrice2_2() const;
		double GetQuantity_1() const;
		double GetQuantity_2() const;
		DTC::OpenCloseTradeEnum GetOpenOrClose() const;
		DTC::PartialFillHandlingEnum GetPartialFillHandling() const;
		uint8_t GetUseOffsets() const;
		double GetOffsetFromParent1() const;
		double GetOffsetFromParent2() const;
		uint8_t GetMaintainSamePricesOnParentFill() const;
	};

	/*==========================================================================*/
	struct s_SubmitNewOCOOrderInt
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t Symbol;
		vls_t Exchange;

		vls_t ClientOrderID_1;
		DTC::OrderTypeEnum OrderType_1;
		DTC::BuySellEnum BuySell_1;
		int64_t Price1_1;
		int64_t Price2_1;
		int64_t Quantity_1;

		vls_t ClientOrderID_2;
		DTC::OrderTypeEnum OrderType_2;
		DTC::BuySellEnum BuySell_2;
		int64_t Price1_2;
		int64_t Price2_2;
		int64_t Quantity_2;

		DTC::TimeInForceEnum TimeInForce;
		DTC::t_DateTime GoodTillDateTime;

		vls_t TradeAccount;

		uint8_t IsAutomatedOrder;

		vls_t ParentTriggerClientOrderID;

		vls_t FreeFormText;

		float Divisor;

		DTC::OpenCloseTradeEnum OpenOrClose;

		DTC::PartialFillHandlingEnum PartialFillHandling;

		s_SubmitNewOCOOrderInt()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SubmitNewOCOOrderInt));
			Size = sizeof(s_SubmitNewOCOOrderInt);
			Type = DTC::SUBMIT_NEW_OCO_ORDER_INT;
			BaseSize = Size;
		}

		void AddClientOrderID_1(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID_1, StringLength);
		}

		void AddClientOrderID_2(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID_2, StringLength);
		}

		const char* GetFreeFormText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_SubmitNewOCOOrderInt, FreeFormText));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}

		const char* GetClientOrderID_1() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID_1, offsetof(s_SubmitNewOCOOrderInt, ClientOrderID_1));
		}

		const char* GetClientOrderID_2() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID_2, offsetof(s_SubmitNewOCOOrderInt, ClientOrderID_2));
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SubmitNewOCOOrderInt, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SubmitNewOCOOrderInt, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		void AddParentTriggerClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ParentTriggerClientOrderID, StringLength);
		}

		const char* GetParentTriggerClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ParentTriggerClientOrderID, offsetof(s_SubmitNewOCOOrderInt, ParentTriggerClientOrderID));
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_SubmitNewOCOOrderInt, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		DTC::OrderTypeEnum GetOrderType_1() const;
		DTC::OrderTypeEnum GetOrderType_2() const;
		DTC::BuySellEnum GetBuySell_1() const;
		DTC::BuySellEnum GetBuySell_2() const;
		DTC::TimeInForceEnum GetTimeInForce() const;
		DTC::t_DateTime GetGoodTillDateTime() const;
		uint8_t GetIsAutomatedOrder() const;
		int64_t GetPrice1_1() const;
		int64_t GetPrice2_1() const;
		int64_t GetPrice1_2() const;
		int64_t GetPrice2_2() const;
		int64_t GetQuantity_1() const;
		int64_t GetQuantity_2() const;
		float GetDivisor() const;
		DTC::OpenCloseTradeEnum GetOpenOrClose() const;
		DTC::PartialFillHandlingEnum GetPartialFillHandling() const;
	};

	/*==========================================================================*/
	struct s_OpenOrdersRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		int32_t RequestAllOrders;

		vls_t ServerOrderID;
		vls_t TradeAccount;

		s_OpenOrdersRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_OpenOrdersRequest));
			Size = sizeof(s_OpenOrdersRequest);
			Type = DTC::OPEN_ORDERS_REQUEST;
			BaseSize = Size;

			RequestAllOrders = 1;
		}

		int32_t GetRequestID() const;
		int32_t GetRequestAllOrders() const;

		void AddServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerOrderID, StringLength);
		}

		const char* GetServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerOrderID, offsetof(s_OpenOrdersRequest, ServerOrderID));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_OpenOrdersRequest, TradeAccount));
		}
	};

	/*==========================================================================*/
	struct s_HistoricalOrderFillsRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		vls_t ServerOrderID;

		int32_t NumberOfDays;

		vls_t TradeAccount;

		DTC::t_DateTime StartDateTime;

		s_HistoricalOrderFillsRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalOrderFillsRequest));
			Size = sizeof(s_HistoricalOrderFillsRequest);
			Type = DTC::HISTORICAL_ORDER_FILLS_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		int32_t GetNumberOfDays() const;
		DTC::t_DateTime GetStartDateTime() const;

		void AddServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerOrderID, StringLength);
		}

		const char* GetServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerOrderID, offsetof(s_HistoricalOrderFillsRequest, ServerOrderID));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_HistoricalOrderFillsRequest, TradeAccount));
		}
	};

	/*==========================================================================*/
	struct s_HistoricalOrderFillsReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_HistoricalOrderFillsReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalOrderFillsReject));
			Size = sizeof(s_HistoricalOrderFillsReject);
			Type = DTC::HISTORICAL_ORDER_FILLS_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_HistoricalOrderFillsReject, RejectText));
		}
	};

	/*==========================================================================*/
	struct s_CurrentPositionsRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t  TradeAccount;

		s_CurrentPositionsRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_CurrentPositionsRequest));
			Size = sizeof(s_CurrentPositionsRequest);
			Type = DTC::CURRENT_POSITIONS_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_CurrentPositionsRequest, TradeAccount));
		}
	};

	/*==========================================================================*/
	struct s_CurrentPositionsReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_CurrentPositionsReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_CurrentPositionsReject));
			Size = sizeof(s_CurrentPositionsReject);
			Type = DTC::CURRENT_POSITIONS_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_CurrentPositionsReject, RejectText));
		}
	};

	/*==========================================================================*/
	struct s_OrderUpdate
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		int32_t TotalNumMessages;
		int32_t MessageNumber;

		vls_t Symbol;
		vls_t Exchange;

		vls_t PreviousServerOrderID;

		vls_t ServerOrderID;

		vls_t ClientOrderID;

		vls_t ExchangeOrderID;

		DTC::OrderStatusEnum OrderStatus;

		DTC::OrderUpdateReasonEnum OrderUpdateReason;

		DTC::OrderTypeEnum OrderType;

		DTC::BuySellEnum BuySell;

		double Price1;

		double Price2;

		DTC::TimeInForceEnum TimeInForce;
		DTC::t_DateTime GoodTillDateTime;
		double OrderQuantity;
		double FilledQuantity;
		double RemainingQuantity;
		double AverageFillPrice;

		double LastFillPrice;
		DTC::t_DateTimeWithMillisecondsInt LastFillDateTime;
		double LastFillQuantity;
		vls_t LastFillExecutionID;

		vls_t TradeAccount;
		vls_t InfoText;

		uint8_t NoOrders;
		vls_t ParentServerOrderID;
		vls_t OCOLinkedOrderServerOrderID;

		DTC::OpenCloseTradeEnum OpenOrClose;

		vls_t PreviousClientOrderID;
		vls_t FreeFormText;

		DTC::t_DateTimeWithMillisecondsInt OrderReceivedDateTime;
		DTC::t_DateTimeWithMilliseconds LatestTransactionDateTime;

		s_OrderUpdate()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_OrderUpdate));
			Size = sizeof(s_OrderUpdate);
			Type = DTC::ORDER_UPDATE;
			BaseSize = Size;

			//The following initializations indicate to the Client that these variables are in an unset state and their values should not be used
			Price1 = DBL_MAX;
			Price2 = DBL_MAX;

			OrderQuantity = DBL_MAX;
			FilledQuantity = DBL_MAX;
			RemainingQuantity = DBL_MAX;
			AverageFillPrice = DBL_MAX;

			LastFillPrice = DBL_MAX;
			LastFillQuantity = DBL_MAX;
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_OrderUpdate, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_OrderUpdate, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetPreviousServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, PreviousServerOrderID, offsetof(s_OrderUpdate, PreviousServerOrderID));
		}

		void AddPreviousServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, PreviousServerOrderID, StringLength);
		}

		const char* GetServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerOrderID, offsetof(s_OrderUpdate, ServerOrderID));
		}

		void AddServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerOrderID, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_OrderUpdate, ClientOrderID));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		const char* GetExchangeOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ExchangeOrderID, offsetof(s_OrderUpdate, ExchangeOrderID));
		}

		void AddExchangeOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ExchangeOrderID, StringLength);
		}

		void AddLastFillExecutionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, LastFillExecutionID, StringLength);
		}

		const char* GetLastFillExecutionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, LastFillExecutionID, offsetof(s_OrderUpdate, LastFillExecutionID));
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_OrderUpdate, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetInfoText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, InfoText, offsetof(s_OrderUpdate, InfoText));
		}

		void AddInfoText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, InfoText, StringLength);
		}

		const char* GetParentServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ParentServerOrderID, offsetof(s_OrderUpdate, ParentServerOrderID));
		}

		void AddParentServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ParentServerOrderID, StringLength);
		}

		const char* GetOCOLinkedOrderServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, OCOLinkedOrderServerOrderID, offsetof(s_OrderUpdate, OCOLinkedOrderServerOrderID));
		}

		void AddOCOLinkedOrderServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, OCOLinkedOrderServerOrderID, StringLength);
		}

		const char* GetPreviousClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, PreviousClientOrderID, offsetof(s_OrderUpdate, PreviousClientOrderID));
		}

		void AddPreviousClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, PreviousClientOrderID, StringLength);
		}

		const char* GetFreeFormText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_OrderUpdate, FreeFormText));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}

		double GetOrderQuantity() const;
		double GetFilledQuantity() const;
		double GetRemainingQuantity() const;
		double GetLastFillQuantity() const;
		int32_t GetRequestID() const;
		int32_t GetMessageNumber() const;
		int32_t GetTotalNumMessages() const;
		DTC::OrderStatusEnum GetOrderStatus() const;
		DTC::OrderUpdateReasonEnum GetOrderUpdateReason() const;
		DTC::OrderTypeEnum GetOrderType() const;
		DTC::BuySellEnum GetBuySell() const;
		double GetPrice1() const;
		double GetPrice2() const;
		DTC::TimeInForceEnum GetTimeInForce() const;
		DTC::t_DateTime GetGoodTillDateTime() const;
		double GetAverageFillPrice() const;
		double GetLastFillPrice() const;
		DTC::t_DateTime GetLastFillDateTime() const;
		uint8_t GetNoOrders() const;
		DTC::OpenCloseTradeEnum GetOpenOrClose() const;
		DTC::t_DateTime GetOrderReceivedDateTime() const;
		DTC::t_DateTimeWithMilliseconds GetLatestTransactionDateTime() const;
	};

	/*==========================================================================*/
	struct s_OpenOrdersReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_OpenOrdersReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_OpenOrdersReject));
			Size = sizeof(s_OpenOrdersReject);
			Type = DTC::OPEN_ORDERS_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_OpenOrdersReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_HistoricalOrderFillResponse
	{
		uint16_t Size = sizeof(s_HistoricalOrderFillResponse);
		uint16_t Type = DTC::HISTORICAL_ORDER_FILL_RESPONSE;
		uint16_t BaseSize = Size;

		int32_t RequestID = 0;

		int32_t TotalNumberMessages = 0;

		int32_t MessageNumber = 0;

		vls_t Symbol;
		vls_t Exchange;
		vls_t ServerOrderID;
		DTC::BuySellEnum BuySell = DTC::BUY_SELL_UNSET;
		double Price = 0;
		DTC::t_DateTime DateTime = 0;
		double Quantity = 0;
		vls_t UniqueExecutionID;
		vls_t TradeAccount;

		DTC::OpenCloseTradeEnum OpenClose = DTC::TRADE_UNSET;

		uint8_t NoOrderFills = 0;

		vls_t InfoText;

		double HighPriceDuringPosition = 0;
		double LowPriceDuringPosition = 0;
		double PositionQuantity = DBL_MAX;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_HistoricalOrderFillResponse, Symbol));
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_HistoricalOrderFillResponse, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetServerOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ServerOrderID, offsetof(s_HistoricalOrderFillResponse, ServerOrderID));
		}

		void AddServerOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ServerOrderID, StringLength);
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_HistoricalOrderFillResponse, TradeAccount));
		}

		void AddUniqueExecutionID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, UniqueExecutionID, StringLength);
		}

		const char* GetUniqueExecutionID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, UniqueExecutionID, offsetof(s_HistoricalOrderFillResponse, UniqueExecutionID));
		}

		void AddInfoText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, InfoText, StringLength);
		}

		const char* GetInfoText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, InfoText, offsetof(s_HistoricalOrderFillResponse, InfoText));
		}

		int32_t GetRequestID() const;
		int32_t GetMessageNumber() const;
		int32_t GetTotalNumberMessages() const;
		DTC::BuySellEnum GetBuySell() const;
		double GetPrice() const;
		DTC::t_DateTime GetDateTime() const;
		double GetQuantity() const;
		DTC::OpenCloseTradeEnum GetOpenClose() const;
		uint8_t GetNoOrderFills() const;
		double GetHighPriceDuringPosition() const;
		double GetLowPriceDuringPosition() const;
		double GetPositionQuantity() const;
	};

	/*==========================================================================*/
	struct s_PositionUpdate
	{
		uint16_t Size = sizeof(s_PositionUpdate);
		uint16_t Type = DTC::POSITION_UPDATE;
		uint16_t BaseSize = Size;

		int32_t RequestID = 0;

		int32_t TotalNumberMessages = 0;

		int32_t MessageNumber = 0;

		vls_t Symbol;
		vls_t Exchange;

		double Quantity = 0;
		double AveragePrice = 0;

		vls_t PositionIdentifier;

		vls_t TradeAccount;
		uint8_t NoPositions = 0;
		uint8_t Unsolicited = 0;

		double MarginRequirement = 0;
		DTC::t_DateTime4Byte EntryDateTime = 0;
		double OpenProfitLoss = 0;
		double HighPriceDuringPosition = 0;
		double LowPriceDuringPosition = 0;
		double QuantityLimit = 0;
		double MaxPotentialPostionQuantity = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_PositionUpdate, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_PositionUpdate, Exchange));
		}

		void AddPositionIdentifier(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, PositionIdentifier, StringLength);
		}

		const char* GetPositionIdentifier() const
		{
			return GetVariableLengthStringField(Size, BaseSize, PositionIdentifier, offsetof(s_PositionUpdate, PositionIdentifier));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_PositionUpdate, TradeAccount));
		}

		int32_t GetRequestID() const;
		int32_t GetTotalNumberMessages() const;
		int32_t GetMessageNumber() const;
		double GetQuantity() const;
		double GetAveragePrice() const;
		uint8_t GetNoPositions() const;
		uint8_t GetUnsolicited() const;

		double GetMarginRequirement() const;
		DTC::t_DateTime4Byte GetEntryDateTime() const;

		double GetOpenProfitLoss() const;
		double GetHighPriceDuringPosition() const;
		double GetLowPriceDuringPosition() const;
		double GetQuantityLimit() const;
		double GetMaxPotentialPostionQuantity() const;
	};

	/*==========================================================================*/
	struct s_AddCorrectingOrderFill
	{
		uint16_t Size = sizeof(s_AddCorrectingOrderFill);
		uint16_t Type = DTC::ADD_CORRECTING_ORDER_FILL;
		uint16_t BaseSize = Size;

		vls_t Symbol;
		vls_t Exchange;
		vls_t TradeAccount;
		vls_t ClientOrderID;
		DTC::BuySellEnum BuySell = DTC::BUY_SELL_UNSET;
		double FillPrice = 0;
		double FillQuantity = 0;
		vls_t FreeFormText;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_AddCorrectingOrderFill, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_AddCorrectingOrderFill, Exchange));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_AddCorrectingOrderFill, TradeAccount));
		}

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_AddCorrectingOrderFill, ClientOrderID));
		}

		void AddFreeFormText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, FreeFormText, StringLength);
		}

		const char* GetFreeFormText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, FreeFormText, offsetof(s_AddCorrectingOrderFill, FreeFormText));
		}

		DTC::BuySellEnum GetBuySell() const;
		double GetFillPrice() const;
		double GetFillQuantity() const;

	};

	/*==========================================================================*/
	struct s_CorrectingOrderFillResponse
	{
		uint16_t Size = sizeof(s_CorrectingOrderFillResponse);
		uint16_t Type = DTC::CORRECTING_ORDER_FILL_RESPONSE;
		uint16_t BaseSize = Size;

		vls_t ClientOrderID;
		vls_t ResultText;
		uint8_t IsError = 0;

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;

		void AddClientOrderID(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ClientOrderID, StringLength);
		}

		const char* GetClientOrderID() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ClientOrderID, offsetof(s_CorrectingOrderFillResponse, ClientOrderID));
		}

		void AddResultText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ResultText, StringLength);
		}

		const char* GetResultText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ResultText, offsetof(s_CorrectingOrderFillResponse, ResultText));
		}

		uint8_t GetIsError() const;

	};


	/*==========================================================================*/
	struct s_TradeAccountResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t TotalNumberMessages;

		int32_t MessageNumber;

		vls_t TradeAccount;

		int32_t RequestID;

		s_TradeAccountResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_TradeAccountResponse));
			Size = sizeof(s_TradeAccountResponse);
			Type = DTC::TRADE_ACCOUNT_RESPONSE;
			BaseSize = Size;
		}

		int32_t GetTotalNumberMessages() const;
		int32_t GetMessageNumber() const;

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_TradeAccountResponse, TradeAccount));
		}

		int32_t GetRequestID() const;
	};

	/*==========================================================================*/
	struct s_ExchangeListResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t Exchange;
		uint8_t IsFinalMessage;
		vls_t Description;

		s_ExchangeListResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_ExchangeListResponse));
			Size = sizeof(s_ExchangeListResponse);
			Type = DTC::EXCHANGE_LIST_RESPONSE;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		uint8_t GetIsFinalMessage() const;

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_ExchangeListResponse, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetDescription() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Description, offsetof(s_ExchangeListResponse, Description));
		}

		void AddDescription(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Description, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_SymbolsForExchangeRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t Exchange;

		DTC::SecurityTypeEnum SecurityType;
		DTC::RequestActionEnum RequestAction;

		vls_t Symbol;

		s_SymbolsForExchangeRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SymbolsForExchangeRequest));
			Size = sizeof(s_SymbolsForExchangeRequest);
			Type = DTC::SYMBOLS_FOR_EXCHANGE_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		DTC::SecurityTypeEnum GetSecurityType() const;
		DTC::RequestActionEnum GetRequestAction() const;

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SymbolsForExchangeRequest, Exchange));
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SymbolsForExchangeRequest, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_UnderlyingSymbolsForExchangeRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		vls_t Exchange;

		DTC::SecurityTypeEnum SecurityType;

		s_UnderlyingSymbolsForExchangeRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_UnderlyingSymbolsForExchangeRequest));
			Size = sizeof(s_UnderlyingSymbolsForExchangeRequest);
			Type = DTC::UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		DTC::SecurityTypeEnum GetSecurityType() const;

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_UnderlyingSymbolsForExchangeRequest, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_SymbolsForUnderlyingRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		vls_t UnderlyingSymbol;
		vls_t Exchange;

		DTC::SecurityTypeEnum SecurityType;

		s_SymbolsForUnderlyingRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SymbolsForUnderlyingRequest));
			Size = sizeof(s_SymbolsForUnderlyingRequest);
			Type = DTC::SYMBOLS_FOR_UNDERLYING_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		DTC::SecurityTypeEnum GetSecurityType() const;

		const char* GetUnderlyingSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, UnderlyingSymbol, offsetof(s_SymbolsForUnderlyingRequest, UnderlyingSymbol));
		}

		void AddUnderlyingSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, UnderlyingSymbol, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SymbolsForUnderlyingRequest, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_SymbolSearchRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t SearchText;
		vls_t Exchange;

		DTC::SecurityTypeEnum SecurityType;
		DTC::SearchTypeEnum SearchType;

		s_SymbolSearchRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SymbolSearchRequest));
			Size = sizeof(s_SymbolSearchRequest);
			Type = DTC::SYMBOL_SEARCH_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		DTC::SecurityTypeEnum GetSecurityType() const;
		DTC::SearchTypeEnum GetSearchType() const;

		const char* GetSearchText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, SearchText, offsetof(s_SymbolSearchRequest, SearchText));
		}

		void AddSearchText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, SearchText, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SymbolSearchRequest, Exchange));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_SecurityDefinitionForSymbolRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		vls_t Symbol;
		vls_t Exchange;


		s_SecurityDefinitionForSymbolRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SecurityDefinitionForSymbolRequest));
			Size = sizeof(s_SecurityDefinitionForSymbolRequest);
			Type = DTC::SECURITY_DEFINITION_FOR_SYMBOL_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SecurityDefinitionForSymbolRequest, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SecurityDefinitionForSymbolRequest, Exchange));
		}
	};

	/*==========================================================================*/
	struct s_SecurityDefinitionResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		vls_t Symbol;
		vls_t Exchange;

		DTC::SecurityTypeEnum SecurityType;

		vls_t Description;
		float MinPriceIncrement;
		DTC::PriceDisplayFormatEnum PriceDisplayFormat;
		float CurrencyValuePerIncrement;

		uint8_t IsFinalMessage;

		float FloatToIntPriceMultiplier;
		float IntToFloatPriceDivisor;
		vls_t UnderlyingSymbol;

		uint8_t UpdatesBidAskOnly;

		float StrikePrice;

		DTC::PutCallEnum PutOrCall;

		uint32_t ShortInterest;

		DTC::t_DateTime4Byte SecurityExpirationDate;

		float BuyRolloverInterest;
		float SellRolloverInterest;

		float EarningsPerShare;
		uint32_t SharesOutstanding;

		float IntToFloatQuantityDivisor;

		uint8_t HasMarketDepthData;

		float DisplayPriceMultiplier;

		vls_t ExchangeSymbol;

		float InitialMarginRequirement;
		float MaintenanceMarginRequirement;

		vls_t Currency;
		float ContractSize;
		uint32_t OpenInterest;
		DTC::t_DateTime4Byte RolloverDate;
		uint8_t IsDelayed;

		s_SecurityDefinitionResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SecurityDefinitionResponse));
			Size = sizeof(s_SecurityDefinitionResponse);
			Type = DTC::SECURITY_DEFINITION_RESPONSE;
			BaseSize = Size;

			FloatToIntPriceMultiplier = 1.0;
			IntToFloatPriceDivisor = 1.0;
			HasMarketDepthData = 1;
			DisplayPriceMultiplier = 1.0;
		}

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_SecurityDefinitionResponse, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_SecurityDefinitionResponse, Exchange));
		}

		void AddDescription(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Description, StringLength);
		}

		const char* GetDescription() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Description, offsetof(s_SecurityDefinitionResponse, Description));
		}

		int32_t GetRequestID() const;
		DTC::SecurityTypeEnum GetSecurityType() const;
		float GetMinPriceIncrement() const;
		DTC::PriceDisplayFormatEnum GetPriceDisplayFormat() const;
		float GetCurrencyValuePerIncrement() const;
		uint8_t GetIsFinalMessage() const;
		float GetFloatToIntPriceMultiplier() const;
		float GetIntToFloatPriceDivisor() const;

		void AddUnderlyingSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, UnderlyingSymbol, StringLength);
		}

		const char* GetUnderlyingSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, UnderlyingSymbol, offsetof(s_SecurityDefinitionResponse, UnderlyingSymbol));
		}

		uint8_t GetUpdatesBidAskOnly() const;
		float GetStrikePrice() const;
		DTC::PutCallEnum GetPutOrCall() const;
		uint32_t GetShortInterest() const;
		DTC::t_DateTime4Byte GetSecurityExpirationDate() const;
		float GetBuyRolloverInterest() const;
		float GetSellRolloverInterest() const;
		float GetEarningsPerShare() const;
		uint32_t GetSharesOutstanding() const;
		float GetIntToFloatQuantityDivisor() const;
		uint8_t GetHasMarketDepthData() const;
		float GetDisplayPriceMultiplier() const;

		void AddExchangeSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, ExchangeSymbol, StringLength);
		}

		const char* GetExchangeSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, ExchangeSymbol, offsetof(s_SecurityDefinitionResponse, ExchangeSymbol));
		}

		float GetInitialMarginRequirement() const;
		float GetMaintenanceMarginRequirement() const;

		void AddCurrency(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Currency, StringLength);
		}

		const char* GetCurrency() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Currency, offsetof(s_SecurityDefinitionResponse, Currency));
		}

		float GetContractSize() const;
		uint32_t GetOpenInterest() const;
		DTC::t_DateTime4Byte GetRolloverDate() const;
		uint8_t GetIsDelayed() const;
	};


	/*==========================================================================*/
	struct s_SecurityDefinitionReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_SecurityDefinitionReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_SecurityDefinitionReject));
			Size = sizeof(s_SecurityDefinitionReject);
			Type = DTC::SECURITY_DEFINITION_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_SecurityDefinitionReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_AccountBalanceRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t TradeAccount;

		s_AccountBalanceRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_AccountBalanceRequest));
			Size = sizeof(s_AccountBalanceRequest);
			Type = DTC::ACCOUNT_BALANCE_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_AccountBalanceRequest, TradeAccount));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_AccountBalanceReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_AccountBalanceReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_AccountBalanceReject));
			Size = sizeof(s_AccountBalanceReject);
			Type = DTC::ACCOUNT_BALANCE_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_AccountBalanceReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_AccountBalanceUpdate
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		double CashBalance;

		double BalanceAvailableForNewPositions;

		vls_t AccountCurrency;

		vls_t TradeAccount;
		double SecuritiesValue;  // Not including cash
		double MarginRequirement;

		int32_t TotalNumberMessages;
		int32_t MessageNumber;
		uint8_t NoAccountBalances;
		uint8_t Unsolicited;

		double OpenPositionsProfitLoss;
		double DailyProfitLoss;

		vls_t InfoText;

		uint64_t TransactionIdentifier;
		double DailyNetLossLimit;
		double TrailingAccountValueToLimitPositions;

		uint8_t DailyNetLossLimitReached = 0;
		uint8_t IsUnderRequiredMargin = 0;

		uint8_t ClosePositionsAtEndOfDay = 0;
		uint8_t TradingIsDisabled = 0;
		vls_t Description;
		uint8_t IsUnderRequiredAccountValue = 0;

		DTC::t_DateTimeWithMicrosecondsInt TransactionDateTime = 0;

		double MarginRequirementFull = 0;
		double MarginRequirementFullPositionsOnly = 0;
		double PeakMarginRequirement = 0;

		s_AccountBalanceUpdate()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_AccountBalanceUpdate));
			Size = sizeof(s_AccountBalanceUpdate);
			Type = DTC::ACCOUNT_BALANCE_UPDATE;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddAccountCurrency(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AccountCurrency, StringLength);
		}

		const char* GetAccountCurrency() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AccountCurrency, offsetof(s_AccountBalanceUpdate, AccountCurrency));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_AccountBalanceUpdate, TradeAccount));
		}

		double GetCashBalance() const;
		double GetBalanceAvailableForNewPositions() const;
		double GetSecuritiesValue() const;
		double GetMarginRequirement() const;

		int32_t GetTotalNumberMessages() const;
		int32_t GetMessageNumber() const;
		uint8_t GetNoAccountBalances() const;
		uint8_t GetUnsolicited() const;

		double GetOpenPositionsProfitLoss() const;
		double GetDailyProfitLoss() const;

		void AddInfoText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, InfoText, StringLength);
		}

		const char* GetInfoText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, InfoText, offsetof(s_AccountBalanceUpdate, InfoText));
		}

		uint64_t GetTransactionIdentifier() const;
		double GetDailyNetLossLimit() const;
		double GetTrailingAccountValueToLimitPositions() const;

		uint8_t GetDailyNetLossLimitReached() const;
		uint8_t GetIsUnderRequiredMargin() const;

		uint8_t GetClosePositionsAtEndOfDay() const;
		uint8_t GetTradingIsDisabled() const;

		void AddDescription(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Description, StringLength);
		}

		const char* GetDescription() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Description, offsetof(s_AccountBalanceUpdate, Description));
		}

		uint8_t GetIsUnderRequiredAccountValue() const;
		DTC::t_DateTimeWithMicrosecondsInt GetTransactionDateTime() const;

		double GetMarginRequirementFull() const;
		double GetMarginRequirementFullPositionsOnly() const;
		double GetPeakMarginRequirement() const;
	};


	/*==========================================================================*/
	struct s_AccountBalanceAdjustment
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t TradeAccount;
		double CreditAmount;
		double DebitAmount;
		vls_t Currency;

		vls_t Reason;

		s_AccountBalanceAdjustment()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_AccountBalanceAdjustment));
			Size = sizeof(s_AccountBalanceAdjustment);
			Type = DTC::ACCOUNT_BALANCE_ADJUSTMENT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_AccountBalanceAdjustment, TradeAccount));
		}

		double GetCreditAmount() const;
		double GetDebitAmount() const;

		void AddCurrency(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Currency, StringLength);
		}

		const char* GetCurrency() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Currency, offsetof(s_AccountBalanceAdjustment, Currency));
		}

		void AddReason(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Reason, StringLength);
		}

		const char* GetReason() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Reason, offsetof(s_AccountBalanceAdjustment, Reason));
		}
	};

	/*==========================================================================*/
	struct s_AccountBalanceAdjustmentReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_AccountBalanceAdjustmentReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_AccountBalanceAdjustmentReject));
			Size = sizeof(s_AccountBalanceAdjustmentReject);
			Type = DTC::ACCOUNT_BALANCE_ADJUSTMENT_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_AccountBalanceAdjustmentReject, RejectText));
		}

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}
	};

	/*==========================================================================*/
	struct s_HistoricalAccountBalancesRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t TradeAccount;

		DTC::t_DateTime StartDateTime;

		s_HistoricalAccountBalancesRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalAccountBalancesRequest));
			Size = sizeof(s_HistoricalAccountBalancesRequest);
			Type = DTC::HISTORICAL_ACCOUNT_BALANCES_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_HistoricalAccountBalancesRequest, TradeAccount));
		}

		DTC::t_DateTime GetStartDateTime() const;
	};

	/*==========================================================================*/
	struct s_HistoricalAccountBalancesReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_HistoricalAccountBalancesReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalAccountBalancesReject));
			Size = sizeof(s_HistoricalAccountBalancesReject);
			Type = DTC::HISTORICAL_ACCOUNT_BALANCES_REJECT;
			BaseSize = Size;
		}

		uint32_t GetRequestID() const;

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_HistoricalAccountBalancesReject, RejectText));
		}
	};

	/*==========================================================================*/
	struct s_HistoricalAccountBalanceResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		DTC::t_DateTimeWithMilliseconds DateTime;
		double CashBalance;
		vls_t AccountCurrency;
		vls_t TradeAccount;
		uint8_t IsFinalResponse;
		uint8_t NoAccountBalances;
		vls_t InfoText;
		vls_t TransactionId;

		s_HistoricalAccountBalanceResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalAccountBalanceResponse));
			Size = sizeof(s_HistoricalAccountBalanceResponse);
			Type = DTC::HISTORICAL_ACCOUNT_BALANCE_RESPONSE;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		DTC::t_DateTimeWithMilliseconds GetDateTime() const;
		double GetCashBalance() const;

		void AddAccountCurrency(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, AccountCurrency, StringLength);
		}

		const char* GetAccountCurrency() const
		{
			return GetVariableLengthStringField(Size, BaseSize, AccountCurrency, offsetof(s_HistoricalAccountBalanceResponse, AccountCurrency));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_HistoricalAccountBalanceResponse, TradeAccount));
		}

		uint8_t GetIsFinalResponse() const;
		uint8_t GetNoAccountBalances() const;

		void AddInfoText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, InfoText, StringLength);
		}

		const char* GetInfoText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, InfoText, offsetof(s_HistoricalAccountBalanceResponse, InfoText));
		}

		void AddTransactionId(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TransactionId, StringLength);
		}

		const char* GetTransactionId() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TransactionId, offsetof(s_HistoricalAccountBalanceResponse, TransactionId));
		}
	};

	/*==========================================================================*/
	struct s_UserMessage
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t UserMessage;

		uint8_t IsPopupMessage;

		s_UserMessage()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_UserMessage));
			Size = sizeof(s_UserMessage);
			Type = DTC::USER_MESSAGE;
			BaseSize = Size;

			IsPopupMessage = 1;
		}

		uint8_t GetIsPopupMessage() const;

		void AddUserMessage(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, UserMessage, StringLength);
		}

		const char* GetUserMessage() const
		{
			return GetVariableLengthStringField(Size, BaseSize, UserMessage, offsetof(s_UserMessage, UserMessage));
		}
	};

	/*==========================================================================*/
	struct s_GeneralLogMessage
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t MessageText;

		s_GeneralLogMessage()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_GeneralLogMessage));
			Size = sizeof(s_GeneralLogMessage);
			Type = DTC::GENERAL_LOG_MESSAGE;
			BaseSize = Size;
		}

		void AddMessageText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, MessageText, StringLength);
		}

		const char* GetMessageText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, MessageText, offsetof(s_GeneralLogMessage, MessageText));
		}
	};

	/*==========================================================================*/
	struct s_AlertMessage
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t MessageText;
		vls_t TradeAccount;

		s_AlertMessage()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_AlertMessage));
			Size = sizeof(s_AlertMessage);
			Type = DTC::ALERT_MESSAGE;
			BaseSize = Size;
		}

		void AddMessageText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, MessageText, StringLength);
		}

		const char* GetMessageText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, MessageText, offsetof(s_AlertMessage, MessageText));
		}

		void AddTradeAccount(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, TradeAccount, StringLength);
		}

		const char* GetTradeAccount() const
		{
			return GetVariableLengthStringField(Size, BaseSize, TradeAccount, offsetof(s_AlertMessage, TradeAccount));
		}
	};

	/*==========================================================================*/
	struct s_JournalEntryAdd
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t JournalEntry;
		DTC::t_DateTime DateTime;

		s_JournalEntryAdd()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_JournalEntryAdd));
			Size = sizeof(s_JournalEntryAdd);
			Type = DTC::JOURNAL_ENTRY_ADD;
			BaseSize = Size;
		}

		void AddJournalEntry(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, JournalEntry, StringLength);
		}

		const char* GetJournalEntry() const
		{
			return GetVariableLengthStringField(Size, BaseSize, JournalEntry, offsetof(s_JournalEntryAdd, JournalEntry));
		}

		DTC::t_DateTime GetDateTime() const;
	};

	/*==========================================================================*/
	struct s_JournalEntriesReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t RejectText;

		s_JournalEntriesReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_JournalEntriesReject));
			Size = sizeof(s_JournalEntriesReject);
			Type = DTC::JOURNAL_ENTRIES_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_JournalEntriesReject, RejectText));
		}
	};

	/*==========================================================================*/
	struct s_JournalEntryResponse
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		vls_t JournalEntry;
		DTC::t_DateTime DateTime;
		uint8_t IsFinalResponse;

		s_JournalEntryResponse()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_JournalEntryResponse));
			Size = sizeof(s_JournalEntryResponse);
			Type = DTC::JOURNAL_ENTRY_RESPONSE;
			BaseSize = Size;
		}

		void AddJournalEntry(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, JournalEntry, StringLength);
		}

		const char* GetJournalEntry() const
		{
			return GetVariableLengthStringField(Size, BaseSize, JournalEntry, offsetof(s_JournalEntryResponse, JournalEntry));
		}

		DTC::t_DateTime GetDateTime() const;
		uint8_t GetIsFinalResponse() const;
	};

	/*==========================================================================*/
	struct s_HistoricalPriceDataRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t Symbol;
		vls_t Exchange;
		DTC::HistoricalDataIntervalEnum RecordInterval;
		DTC::t_DateTime StartDateTime;
		DTC::t_DateTime EndDateTime;
		uint32_t MaxDaysToReturn;
		uint8_t UseZLibCompression;
		uint8_t RequestDividendAdjustedStockData;
		uint16_t Integer_1;

		s_HistoricalPriceDataRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalPriceDataRequest));
			Size = sizeof(s_HistoricalPriceDataRequest);
			Type = DTC::HISTORICAL_PRICE_DATA_REQUEST;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_HistoricalPriceDataRequest, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_HistoricalPriceDataRequest, Exchange));
		}

		DTC::HistoricalDataIntervalEnum GetRecordInterval() const;
		DTC::t_DateTime GetStartDateTime() const;
		DTC::t_DateTime GetEndDateTime() const;
		uint32_t GetMaxDaysToReturn() const;
		uint8_t GetUseZLibCompression() const;
		uint8_t GetRequestDividendAdjustedStockData() const;
		uint16_t GetInteger_1() const;
	};

	/*==========================================================================*/
	struct s_HistoricalPriceDataReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		vls_t RejectText;

		DTC::HistoricalPriceDataRejectReasonCodeEnum RejectReasonCode;
		uint16_t RetryTimeInSeconds;

		s_HistoricalPriceDataReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(s_HistoricalPriceDataReject));
			Size = sizeof(s_HistoricalPriceDataReject);
			Type = DTC::HISTORICAL_PRICE_DATA_REJECT;
			BaseSize = Size;
		}

		int32_t GetRequestID() const;
		DTC::HistoricalPriceDataRejectReasonCodeEnum GetRejectReasonCode() const;
		uint16_t GetRetryTimeInSeconds() const;

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_HistoricalPriceDataReject, RejectText));
		}
	};

	/*==========================================================================*/
	struct s_HistoricalMarketDepthDataRequest
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;
		vls_t Symbol;
		vls_t Exchange;
		DTC::t_DateTimeWithMicrosecondsInt StartDateTime;
		DTC::t_DateTimeWithMicrosecondsInt EndDateTime;
		uint8_t  UseZLibCompression;
		uint8_t Integer_1;

		s_HistoricalMarketDepthDataRequest()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(*this));
			Type = DTC::HISTORICAL_MARKET_DEPTH_DATA_REQUEST;
			Size = sizeof(*this);
			BaseSize = Size;
		}

		int32_t GetRequestID() const;

		void AddSymbol(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Symbol, StringLength);
		}

		const char* GetSymbol() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Symbol, offsetof(s_HistoricalMarketDepthDataRequest, Symbol));
		}

		void AddExchange(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, Exchange, StringLength);
		}

		const char* GetExchange() const
		{
			return GetVariableLengthStringField(Size, BaseSize, Exchange, offsetof(s_HistoricalMarketDepthDataRequest, Exchange));
		}

		DTC::t_DateTimeWithMicrosecondsInt GetStartDateTime() const;
		DTC::t_DateTimeWithMicrosecondsInt GetEndDateTime() const;
		uint8_t GetUseZLibCompression() const;
		uint8_t GetInteger_1() const;
	};

	/*==========================================================================*/
	struct s_HistoricalMarketDepthDataReject
	{
		uint16_t Size;
		uint16_t Type;
		uint16_t BaseSize;

		int32_t RequestID;

		vls_t RejectText;

		DTC::HistoricalPriceDataRejectReasonCodeEnum RejectReasonCode;

		s_HistoricalMarketDepthDataReject()
		{
			Clear();
		}

		uint16_t GetMessageSize() const;
		uint16_t GetBaseSize() const;
		void Clear()
		{
			memset(this, 0, sizeof(*this));
			Type = DTC::HISTORICAL_MARKET_DEPTH_DATA_REJECT;
			Size = sizeof(*this);
			BaseSize = Size;
			RejectReasonCode = DTC::HPDR_UNSET;
		}

		int32_t GetRequestID() const;

		void AddRejectText(uint16_t StringLength)
		{
			AddVariableLengthStringField(Size, RejectText, StringLength);
		}

		const char* GetRejectText() const
		{
			return GetVariableLengthStringField(Size, BaseSize, RejectText, offsetof(s_HistoricalMarketDepthDataReject, RejectText));
		}

		DTC::HistoricalPriceDataRejectReasonCodeEnum GetRejectReasonCode() const;
	};

#pragma pack()
}
