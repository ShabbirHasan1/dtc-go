package model

type EncodingEnum int32

const (
	BINARY_ENCODING                     EncodingEnum = 0
	BINARY_WITH_VARIABLE_LENGTH_STRINGS EncodingEnum = 1
	JSON_ENCODING                       EncodingEnum = 2
	JSON_COMPACT_ENCODING               EncodingEnum = 3
	PROTOCOL_BUFFERS                    EncodingEnum = 4
)

type LogonStatusEnum int32

const (
	LOGON_SUCCESS               LogonStatusEnum = 1
	LOGON_ERROR                 LogonStatusEnum = 2
	LOGON_ERROR_NO_RECONNECT    LogonStatusEnum = 3
	LOGON_RECONNECT_NEW_ADDRESS LogonStatusEnum = 4
)

type TradeModeEnum int32

const (
	TRADE_MODE_UNSET     TradeModeEnum = 0
	TRADE_MODE_DEMO      TradeModeEnum = 1
	TRADE_MODE_SIMULATED TradeModeEnum = 2
	TRADE_MODE_LIVE      TradeModeEnum = 3
)

type RequestActionEnum int32

const (
	SUBSCRIBE                      RequestActionEnum = 1
	UNSUBSCRIBE                    RequestActionEnum = 2
	SNAPSHOT                       RequestActionEnum = 3
	SNAPSHOT_WITH_INTERVAL_UPDATES RequestActionEnum = 4
)

type UnbundledTradeIndicatorEnum int8

const (
	UNBUNDLED_TRADE_NONE               UnbundledTradeIndicatorEnum = 0
	FIRST_SUB_TRADE_OF_UNBUNDLED_TRADE UnbundledTradeIndicatorEnum = 1
	LAST_SUB_TRADE_OF_UNBUNDLED_TRADE  UnbundledTradeIndicatorEnum = 2
)

type OrderStatusEnum int32

const (
	ORDER_STATUS_UNSPECIFIED            OrderStatusEnum = 0
	ORDER_STATUS_ORDER_SENT             OrderStatusEnum = 1
	ORDER_STATUS_PENDING_OPEN           OrderStatusEnum = 2
	ORDER_STATUS_PENDING_CHILD          OrderStatusEnum = 3
	ORDER_STATUS_OPEN                   OrderStatusEnum = 4
	ORDER_STATUS_PENDING_CANCEL_REPLACE OrderStatusEnum = 5
	ORDER_STATUS_PENDING_CANCEL         OrderStatusEnum = 6
	ORDER_STATUS_FILLED                 OrderStatusEnum = 7
	ORDER_STATUS_CANCELED               OrderStatusEnum = 8
	ORDER_STATUS_REJECTED               OrderStatusEnum = 9
	ORDER_STATUS_PARTIALLY_FILLED       OrderStatusEnum = 10
)

type OrderUpdateReasonEnum int32

const (
	ORDER_UPDATE_REASON_UNSET     OrderUpdateReasonEnum = 0
	OPEN_ORDERS_REQUEST_RESPONSE  OrderUpdateReasonEnum = 1
	NEW_ORDER_ACCEPTED            OrderUpdateReasonEnum = 2
	GENERAL_ORDER_UPDATE          OrderUpdateReasonEnum = 3
	ORDER_FILLED                  OrderUpdateReasonEnum = 4
	ORDER_FILLED_PARTIALLY        OrderUpdateReasonEnum = 5
	ORDER_CANCELED                OrderUpdateReasonEnum = 6
	ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReasonEnum = 7
	NEW_ORDER_REJECTED            OrderUpdateReasonEnum = 8
	ORDER_CANCEL_REJECTED         OrderUpdateReasonEnum = 9
	ORDER_CANCEL_REPLACE_REJECTED OrderUpdateReasonEnum = 10
)

type AtBidOrAskEnum8 uint8

const (
	BID_ASK_UNSET_8 AtBidOrAskEnum8 = 0
	AT_BID_8        AtBidOrAskEnum8 = 1
	AT_ASK_8        AtBidOrAskEnum8 = 2
)

type AtBidOrAskEnum uint16

const (
	BID_ASK_UNSET AtBidOrAskEnum = 0
	AT_BID        AtBidOrAskEnum = 1
	AT_ASK        AtBidOrAskEnum = 2
)

type MarketDepthUpdateTypeEnum uint8

const (
	MARKET_DEPTH_UNSET               MarketDepthUpdateTypeEnum = 0
	MARKET_DEPTH_INSERT_UPDATE_LEVEL MarketDepthUpdateTypeEnum = 1
	MARKET_DEPTH_DELETE_LEVEL        MarketDepthUpdateTypeEnum = 2
)

type FinalUpdateInBatchEnum uint8

const (
	FINAL_UPDATE_UNSET       FinalUpdateInBatchEnum = 0
	FINAL_UPDATE_TRUE        FinalUpdateInBatchEnum = 1
	FINAL_UPDATE_FALSE       FinalUpdateInBatchEnum = 2
	FINAL_UPDATE_BEGIN_BATCH FinalUpdateInBatchEnum = 3
)

type MessageSetBoundaryEnum uint8

const (
	MESSAGE_SET_BOUNDARY_UNSET MessageSetBoundaryEnum = 0
	MESSAGE_SET_BOUNDARY_BEGIN MessageSetBoundaryEnum = 1
	MESSAGE_SET_BOUNDARY_END   MessageSetBoundaryEnum = 2
)

type OrderTypeEnum int32

const (
	ORDER_TYPE_UNSET             OrderTypeEnum = 0
	ORDER_TYPE_MARKET            OrderTypeEnum = 1
	ORDER_TYPE_LIMIT             OrderTypeEnum = 2
	ORDER_TYPE_STOP              OrderTypeEnum = 3
	ORDER_TYPE_STOP_LIMIT        OrderTypeEnum = 4
	ORDER_TYPE_MARKET_IF_TOUCHED OrderTypeEnum = 5
	ORDER_TYPE_LIMIT_IF_TOUCHED  OrderTypeEnum = 6
	ORDER_TYPE_MARKET_LIMIT      OrderTypeEnum = 7
)

type TimeInForceEnum int32

const (
	TIF_UNSET               TimeInForceEnum = 0
	TIF_DAY                 TimeInForceEnum = 1
	TIF_GOOD_TILL_CANCELED  TimeInForceEnum = 2
	TIF_GOOD_TILL_DATE_TIME TimeInForceEnum = 3
	TIF_IMMEDIATE_OR_CANCEL TimeInForceEnum = 4
	TIF_ALL_OR_NONE         TimeInForceEnum = 5
	TIF_FILL_OR_KILL        TimeInForceEnum = 6
)

type BuySellEnum int32

const (
	BUY_SELL_UNSET BuySellEnum = 0
	BUY            BuySellEnum = 1
	SELL           BuySellEnum = 2
)

type OpenCloseTradeEnum int32

const (
	TRADE_UNSET OpenCloseTradeEnum = 0
	TRADE_OPEN  OpenCloseTradeEnum = 1
	TRADE_CLOSE OpenCloseTradeEnum = 2
)

type PartialFillHandlingEnum int8

const (
	PARTIAL_FILL_UNSET                     PartialFillHandlingEnum = 0
	PARTIAL_FILL_HANDLING_REDUCE_QUANTITY  PartialFillHandlingEnum = 1
	PARTIAL_FILL_HANDLING_IMMEDIATE_CANCEL PartialFillHandlingEnum = 2
)

type MarketDataFeedStatusEnum int32

const (
	MARKET_DATA_FEED_STATUS_UNSET MarketDataFeedStatusEnum = 0
	MARKET_DATA_FEED_UNAVAILABLE  MarketDataFeedStatusEnum = 1
	MARKET_DATA_FEED_AVAILABLE    MarketDataFeedStatusEnum = 2
)

type TradingStatusEnum int8

const (
	TRADING_STATUS_UNKNOWN      TradingStatusEnum = 0
	TRADING_STATUS_PRE_OPEN     TradingStatusEnum = 1
	TRADING_STATUS_OPEN         TradingStatusEnum = 2
	TRADING_STATUS_CLOSE        TradingStatusEnum = 3
	TRADING_STATUS_TRADING_HALT TradingStatusEnum = 4
)

type PriceDisplayFormatEnum int32

const (
	PRICE_DISPLAY_FORMAT_UNSET                   PriceDisplayFormatEnum = -1
	PRICE_DISPLAY_FORMAT_DECIMAL_0               PriceDisplayFormatEnum = 0
	PRICE_DISPLAY_FORMAT_DECIMAL_1               PriceDisplayFormatEnum = 1
	PRICE_DISPLAY_FORMAT_DECIMAL_2               PriceDisplayFormatEnum = 2
	PRICE_DISPLAY_FORMAT_DECIMAL_3               PriceDisplayFormatEnum = 3
	PRICE_DISPLAY_FORMAT_DECIMAL_4               PriceDisplayFormatEnum = 4
	PRICE_DISPLAY_FORMAT_DECIMAL_5               PriceDisplayFormatEnum = 5
	PRICE_DISPLAY_FORMAT_DECIMAL_6               PriceDisplayFormatEnum = 6
	PRICE_DISPLAY_FORMAT_DECIMAL_7               PriceDisplayFormatEnum = 7
	PRICE_DISPLAY_FORMAT_DECIMAL_8               PriceDisplayFormatEnum = 8
	PRICE_DISPLAY_FORMAT_DECIMAL_9               PriceDisplayFormatEnum = 9
	PRICE_DISPLAY_FORMAT_DENOMINATOR_256         PriceDisplayFormatEnum = 356
	PRICE_DISPLAY_FORMAT_DENOMINATOR_128         PriceDisplayFormatEnum = 228
	PRICE_DISPLAY_FORMAT_DENOMINATOR_64          PriceDisplayFormatEnum = 164
	PRICE_DISPLAY_FORMAT_DENOMINATOR_32_EIGHTHS  PriceDisplayFormatEnum = 140
	PRICE_DISPLAY_FORMAT_DENOMINATOR_32_QUARTERS PriceDisplayFormatEnum = 136
	PRICE_DISPLAY_FORMAT_DENOMINATOR_32_HALVES   PriceDisplayFormatEnum = 134
	PRICE_DISPLAY_FORMAT_DENOMINATOR_32          PriceDisplayFormatEnum = 132
	PRICE_DISPLAY_FORMAT_DENOMINATOR_16          PriceDisplayFormatEnum = 116
	PRICE_DISPLAY_FORMAT_DENOMINATOR_8           PriceDisplayFormatEnum = 108
	PRICE_DISPLAY_FORMAT_DENOMINATOR_4           PriceDisplayFormatEnum = 104
	PRICE_DISPLAY_FORMAT_DENOMINATOR_2           PriceDisplayFormatEnum = 102
)

type SecurityTypeEnum int32

const (
	SECURITY_TYPE_UNSET            SecurityTypeEnum = 0
	SECURITY_TYPE_FUTURES          SecurityTypeEnum = 1
	SECURITY_TYPE_STOCK            SecurityTypeEnum = 2
	SECURITY_TYPE_FOREX            SecurityTypeEnum = 3
	SECURITY_TYPE_INDEX            SecurityTypeEnum = 4
	SECURITY_TYPE_FUTURES_STRATEGY SecurityTypeEnum = 5
	SECURITY_TYPE_FUTURES_OPTION   SecurityTypeEnum = 7
	SECURITY_TYPE_STOCK_OPTION     SecurityTypeEnum = 6
	SECURITY_TYPE_INDEX_OPTION     SecurityTypeEnum = 8
	SECURITY_TYPE_BOND             SecurityTypeEnum = 9
	SECURITY_TYPE_MUTUAL_FUND      SecurityTypeEnum = 10
)

type PutCallEnum uint8

const (
	PC_UNSET PutCallEnum = 0
	PC_CALL  PutCallEnum = 1
	PC_PUT   PutCallEnum = 2
)

type SearchTypeEnum int32

const (
	SEARCH_TYPE_UNSET          SearchTypeEnum = 0
	SEARCH_TYPE_BY_SYMBOL      SearchTypeEnum = 1
	SEARCH_TYPE_BY_DESCRIPTION SearchTypeEnum = 2
)

type HistoricalDataIntervalEnum int32

const (
	INTERVAL_TICK       HistoricalDataIntervalEnum = 0
	INTERVAL_1_SECOND   HistoricalDataIntervalEnum = 1
	INTERVAL_2_SECONDS  HistoricalDataIntervalEnum = 2
	INTERVAL_4_SECONDS  HistoricalDataIntervalEnum = 4
	INTERVAL_5_SECONDS  HistoricalDataIntervalEnum = 5
	INTERVAL_10_SECONDS HistoricalDataIntervalEnum = 10
	INTERVAL_30_SECONDS HistoricalDataIntervalEnum = 30
	INTERVAL_1_MINUTE   HistoricalDataIntervalEnum = 60
	INTERVAL_5_MINUTE   HistoricalDataIntervalEnum = 300
	INTERVAL_10_MINUTE  HistoricalDataIntervalEnum = 600
	INTERVAL_15_MINUTE  HistoricalDataIntervalEnum = 900
	INTERVAL_30_MINUTE  HistoricalDataIntervalEnum = 1800
	INTERVAL_1_HOUR     HistoricalDataIntervalEnum = 3600
	INTERVAL_2_HOURS    HistoricalDataIntervalEnum = 7200
	INTERVAL_1_DAY      HistoricalDataIntervalEnum = 86400
	INTERVAL_1_WEEK     HistoricalDataIntervalEnum = 604800
)

type HistoricalPriceDataRejectReasonCodeEnum int16

const (
	HPDR_UNSET                                           HistoricalPriceDataRejectReasonCodeEnum = 0
	HPDR_UNABLE_TO_SERVE_DATA_RETRY_IN_SPECIFIED_SECONDS HistoricalPriceDataRejectReasonCodeEnum = 1
	HPDR_UNABLE_TO_SERVE_DATA_DO_NOT_RETRY               HistoricalPriceDataRejectReasonCodeEnum = 2
	HPDR_DATA_REQUEST_OUTSIDE_BOUNDS_OF_AVAILABLE_DATA   HistoricalPriceDataRejectReasonCodeEnum = 3
	HPDR_GENERAL_REJECT_ERROR                            HistoricalPriceDataRejectReasonCodeEnum = 4
)

type TradeConditionEnum int8

const (
	TRADE_CONDITION_NONE                         TradeConditionEnum = 0
	TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE TradeConditionEnum = 1
	TRADE_CONDITION_ODD_LOT_EQUITY_TRADE         TradeConditionEnum = 2
)

type ReplayChartDataActionEnum int32

const (
	REPLAY_CHART_DATA_ACTION_NONE         ReplayChartDataActionEnum = 0
	REPLAY_CHART_DATA_ACTION_STOP         ReplayChartDataActionEnum = 1
	REPLAY_CHART_DATA_ACTION_PAUSE        ReplayChartDataActionEnum = 2
	REPLAY_CHART_DATA_ACTION_RESUME       ReplayChartDataActionEnum = 3
	REPLAY_CHART_DATA_ACTION_FINISH       ReplayChartDataActionEnum = 4
	REPLAY_CHART_DATA_ACTION_CHANGE_SPEED ReplayChartDataActionEnum = 5
)

type ReplayChartDataStatusEnum int32

const (
	REPLAY_CHART_DATA_STATUS_UNSET    ReplayChartDataStatusEnum = 0
	REPLAY_CHART_DATA_STATUS_STARTED  ReplayChartDataStatusEnum = 1
	REPLAY_CHART_DATA_STATUS_ERROR    ReplayChartDataStatusEnum = 2
	REPLAY_CHART_DATA_STATUS_COMPLETE ReplayChartDataStatusEnum = 3
)

type UseCompressionEnum int32

const (
	USE_COMPRESSION_DISABLED              UseCompressionEnum = 0
	USE_COMPRESSION_BLOCK_COMPRESSION     UseCompressionEnum = 1
	USE_COMPRESSION_STREAMING_COMPRESSION UseCompressionEnum = 2
)

type WriteIntradayDataFileSessionValueTypeEnum int32

const (
	INTRADAY_DATA_FILE_SESSION_VALUE_UNSET                        WriteIntradayDataFileSessionValueTypeEnum = 0
	INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_OPEN                   WriteIntradayDataFileSessionValueTypeEnum = 1
	INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_HIGH                   WriteIntradayDataFileSessionValueTypeEnum = 2
	INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_LOW                    WriteIntradayDataFileSessionValueTypeEnum = 3
	INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_VOLUME                 WriteIntradayDataFileSessionValueTypeEnum = 4
	INTRADAY_DATA_FILE_SESSION_VALUE_DAILY_SETTLEMENT_PRICE       WriteIntradayDataFileSessionValueTypeEnum = 5
	INTRADAY_DATA_FILE_SESSION_VALUE_OPEN_INTEREST                WriteIntradayDataFileSessionValueTypeEnum = 6
	INTRADAY_DATA_FILE_SESSION_VALUE_ODD_LOT_TRADE                WriteIntradayDataFileSessionValueTypeEnum = 7
	INTRADAY_DATA_FILE_SESSION_VALUE_NON_LAST_UPDATE_EQUITY_TRADE WriteIntradayDataFileSessionValueTypeEnum = 8
)
