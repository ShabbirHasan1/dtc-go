package dtc

type EncodingEnum int32

const (
	EncodingEnum_BINARY                              EncodingEnum = 0
	EncodingEnum_BINARY_WITH_VARIABLE_LENGTH_STRINGS EncodingEnum = 1
	EncodingEnum_JSON                                EncodingEnum = 2
	EncodingEnum_JSON_COMPACT                        EncodingEnum = 3
	EncodingEnum_PROTOBUF                            EncodingEnum = 4
)

func (enum EncodingEnum) String() string {
	switch enum {
	case EncodingEnum_BINARY:
		return "Binary"
	case EncodingEnum_BINARY_WITH_VARIABLE_LENGTH_STRINGS:
		return "Binary VLS"
	case EncodingEnum_JSON:
		return "JSON"
	case EncodingEnum_JSON_COMPACT:
		return "JSON Compact"
	case EncodingEnum_PROTOBUF:
		return "Protocol Buffers"
	default:
		return "Unknown"
	}
}

/*==========================================================================*/
type LogonStatusEnum int32

const (
	LogonStatusEnum_LOGON_SUCCESS               LogonStatusEnum = 1
	LogonStatusEnum_LOGON_ERROR                 LogonStatusEnum = 2
	LogonStatusEnum_LOGON_ERROR_NO_RECONNECT    LogonStatusEnum = 3
	LogonStatusEnum_LOGON_RECONNECT_NEW_ADDRESS LogonStatusEnum = 4
)

/*==========================================================================*/
type RequestActionEnum int32

const (
	RequestActionEnum_SUBSCRIBE                      RequestActionEnum = 1
	RequestActionEnum_UNSUBSCRIBE                    RequestActionEnum = 2
	RequestActionEnum_SNAPSHOT                       RequestActionEnum = 3
	RequestActionEnum_SNAPSHOT_WITH_INTERVAL_UPDATES RequestActionEnum = 4
)

/*==========================================================================*/
type UnbundledTradeIndicatorEnum int8

const (
	UnbundledTradeIndicatorEnum_UNBUNDLED_TRADE_NONE               UnbundledTradeIndicatorEnum = 0
	UnbundledTradeIndicatorEnum_FIRST_SUB_TRADE_OF_UNBUNDLED_TRADE UnbundledTradeIndicatorEnum = 1
	UnbundledTradeIndicatorEnum_LAST_SUB_TRADE_OF_UNBUNDLED_TRADE  UnbundledTradeIndicatorEnum = 2
)

/*==========================================================================*/
type OrderStatusEnum int32

const (
	OrderStatusEnum_ORDER_STATUS_UNSPECIFIED            OrderStatusEnum = 0
	OrderStatusEnum_ORDER_STATUS_ORDER_SENT             OrderStatusEnum = 1
	OrderStatusEnum_ORDER_STATUS_PENDING_OPEN           OrderStatusEnum = 2
	OrderStatusEnum_ORDER_STATUS_PENDING_CHILD          OrderStatusEnum = 3
	OrderStatusEnum_ORDER_STATUS_OPEN                   OrderStatusEnum = 4
	OrderStatusEnum_ORDER_STATUS_PENDING_CANCEL_REPLACE OrderStatusEnum = 5
	OrderStatusEnum_ORDER_STATUS_PENDING_CANCEL         OrderStatusEnum = 6
	OrderStatusEnum_ORDER_STATUS_FILLED                 OrderStatusEnum = 7
	OrderStatusEnum_ORDER_STATUS_CANCELED               OrderStatusEnum = 8
	OrderStatusEnum_ORDER_STATUS_REJECTED               OrderStatusEnum = 9
	OrderStatusEnum_ORDER_STATUS_PARTIALLY_FILLED       OrderStatusEnum = 10
)

/*==========================================================================*/
type OrderUpdateReasonEnum int32

const (
	OrderUpdateReasonEnum_ORDER_UPDATE_REASON_UNSET     OrderUpdateReasonEnum = 0
	OrderUpdateReasonEnum_OPEN_ORDERS_REQUEST_RESPONSE  OrderUpdateReasonEnum = 1
	OrderUpdateReasonEnum_NEW_ORDER_ACCEPTED            OrderUpdateReasonEnum = 2
	OrderUpdateReasonEnum_GENERAL_ORDER_UPDATE          OrderUpdateReasonEnum = 3
	OrderUpdateReasonEnum_ORDER_FILLED                  OrderUpdateReasonEnum = 4
	OrderUpdateReasonEnum_ORDER_FILLED_PARTIALLY        OrderUpdateReasonEnum = 5
	OrderUpdateReasonEnum_ORDER_CANCELED                OrderUpdateReasonEnum = 6
	OrderUpdateReasonEnum_ORDER_CANCEL_REPLACE_COMPLETE OrderUpdateReasonEnum = 7
	OrderUpdateReasonEnum_NEW_ORDER_REJECTED            OrderUpdateReasonEnum = 8
	OrderUpdateReasonEnum_ORDER_CANCEL_REJECTED         OrderUpdateReasonEnum = 9
	OrderUpdateReasonEnum_ORDER_CANCEL_REPLACE_REJECTED OrderUpdateReasonEnum = 10
)

/*==========================================================================*/
type AtBidOrAskEnum8 uint8

const (
	AtBidOrAskEnum8_BID_ASK_UNSET_8 AtBidOrAskEnum8 = 0
	AtBidOrAskEnum8_AT_BID_8        AtBidOrAskEnum8 = 1
	AtBidOrAskEnum8_AT_ASK_8        AtBidOrAskEnum8 = 2
)

/*==========================================================================*/
type AtBidOrAskEnum uint16

const (
	AtBidOrAskEnum_BID_ASK_UNSET AtBidOrAskEnum = 0
	AtBidOrAskEnum_AT_BID        AtBidOrAskEnum = 1
	AtBidOrAskEnum_AT_ASK        AtBidOrAskEnum = 2
)

/*==========================================================================*/
type MarketDepthUpdateTypeEnum uint8

const (
	MarketDepthUpdateTypeEnum_MARKET_DEPTH_UNSET               MarketDepthUpdateTypeEnum = 0
	MarketDepthUpdateTypeEnum_MARKET_DEPTH_INSERT_UPDATE_LEVEL MarketDepthUpdateTypeEnum = 1 // Insert or update depth at the given price level
	MarketDepthUpdateTypeEnum_MARKET_DEPTH_DELETE_LEVEL        MarketDepthUpdateTypeEnum = 2 // Delete depth at the given price level
)

/*==========================================================================*/
type FinalUpdateInBatchEnum uint8

const (
	FinalUpdateInBatchEnum_FINAL_UPDATE_UNSET       FinalUpdateInBatchEnum = 0
	FinalUpdateInBatchEnum_FINAL_UPDATE_TRUE        FinalUpdateInBatchEnum = 1
	FinalUpdateInBatchEnum_FINAL_UPDATE_FALSE       FinalUpdateInBatchEnum = 2
	FinalUpdateInBatchEnum_FINAL_UPDATE_BEGIN_BATCH FinalUpdateInBatchEnum = 3
)

/*==========================================================================*/
type MessageSetBoundaryEnum uint8

const (
	MessageSetBoundaryEnum_MESSAGE_SET_BOUNDARY_UNSET MessageSetBoundaryEnum = 0
	MessageSetBoundaryEnum_MESSAGE_SET_BOUNDARY_BEGIN MessageSetBoundaryEnum = 1
	MessageSetBoundaryEnum_MESSAGE_SET_BOUNDARY_END   MessageSetBoundaryEnum = 2
)

/*==========================================================================*/
type OrderTypeEnum int32

const (
	OrderTypeEnum_ORDER_TYPE_UNSET             OrderTypeEnum = 0
	OrderTypeEnum_ORDER_TYPE_MARKET            OrderTypeEnum = 1
	OrderTypeEnum_ORDER_TYPE_LIMIT             OrderTypeEnum = 2
	OrderTypeEnum_ORDER_TYPE_STOP              OrderTypeEnum = 3
	OrderTypeEnum_ORDER_TYPE_STOP_LIMIT        OrderTypeEnum = 4
	OrderTypeEnum_ORDER_TYPE_MARKET_IF_TOUCHED OrderTypeEnum = 5
	OrderTypeEnum_ORDER_TYPE_LIMIT_IF_TOUCHED  OrderTypeEnum = 6
	OrderTypeEnum_ORDER_TYPE_MARKET_LIMIT      OrderTypeEnum = 7
)

/*==========================================================================*/
type TimeInForceEnum int32

const (
	TimeInForceEnum_TIF_UNSET               TimeInForceEnum = 0
	TimeInForceEnum_TIF_DAY                 TimeInForceEnum = 1
	TimeInForceEnum_TIF_GOOD_TILL_CANCELED  TimeInForceEnum = 2
	TimeInForceEnum_TIF_GOOD_TILL_DATE_TIME TimeInForceEnum = 3
	TimeInForceEnum_TIF_IMMEDIATE_OR_CANCEL TimeInForceEnum = 4
	TimeInForceEnum_TIF_ALL_OR_NONE         TimeInForceEnum = 5
	TimeInForceEnum_TIF_FILL_OR_KILL        TimeInForceEnum = 6
)

/*==========================================================================*/
type BuySellEnum int32

const (
	BuySellEnum_BUY_SELL_UNSET BuySellEnum = 0
	BuySellEnum_BUY            BuySellEnum = 1
	BuySellEnum_SELL           BuySellEnum = 2
)

/*==========================================================================*/
type OpenCloseTradeEnum int32

const (
	OpenCloseTradeEnum_TRADE_UNSET OpenCloseTradeEnum = 0
	OpenCloseTradeEnum_TRADE_OPEN  OpenCloseTradeEnum = 1
	OpenCloseTradeEnum_TRADE_CLOSE OpenCloseTradeEnum = 2
)

/*==========================================================================*/
type PartialFillHandlingEnum int8

const (
	PartialFillHandlingEnum_PARTIAL_FILL_UNSET                     PartialFillHandlingEnum = 0
	PartialFillHandlingEnum_PARTIAL_FILL_HANDLING_REDUCE_QUANTITY  PartialFillHandlingEnum = 1
	PartialFillHandlingEnum_PARTIAL_FILL_HANDLING_IMMEDIATE_CANCEL PartialFillHandlingEnum = 2
)

/*==========================================================================*/
type MarketDataFeedStatusEnum int32

const (
	MarketDataFeedStatusEnum_MARKET_DATA_FEED_STATUS_UNSET MarketDataFeedStatusEnum = 0
	MarketDataFeedStatusEnum_MARKET_DATA_FEED_UNAVAILABLE  MarketDataFeedStatusEnum = 1
	MarketDataFeedStatusEnum_MARKET_DATA_FEED_AVAILABLE    MarketDataFeedStatusEnum = 2
)

/*==========================================================================*/
type TradingStatusEnum int8

const (
	TradingStatusEnum_TRADING_STATUS_UNKNOWN      TradingStatusEnum = 0
	TradingStatusEnum_TRADING_STATUS_PRE_OPEN     TradingStatusEnum = 1
	TradingStatusEnum_TRADING_STATUS_OPEN         TradingStatusEnum = 2
	TradingStatusEnum_TRADING_STATUS_CLOSE        TradingStatusEnum = 3
	TradingStatusEnum_TRADING_STATUS_TRADING_HALT TradingStatusEnum = 4
)

func (tse TradingStatusEnum) String() string {
	switch tse {
	case TradingStatusEnum_TRADING_STATUS_PRE_OPEN:
		return "Pre-open"
	case TradingStatusEnum_TRADING_STATUS_OPEN:
		return "Open"
	case TradingStatusEnum_TRADING_STATUS_CLOSE:
		return "Closed"
	case TradingStatusEnum_TRADING_STATUS_TRADING_HALT:
		return "Trading Halt"
	//case TradingStatusEnum_TRADING_STATUS_UNKNOWN:
	//	return "Unknown"
	default:
		return "Unknown"
	}
}

/*==========================================================================*/
type PriceDisplayFormatEnum int32

const (
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_UNSET PriceDisplayFormatEnum = -1
	//The following formats indicate the number of decimal places to be displayed
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_0 PriceDisplayFormatEnum = 0
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_1 PriceDisplayFormatEnum = 1
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_2 PriceDisplayFormatEnum = 2
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_3 PriceDisplayFormatEnum = 3
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_4 PriceDisplayFormatEnum = 4
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_5 PriceDisplayFormatEnum = 5
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_6 PriceDisplayFormatEnum = 6
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_7 PriceDisplayFormatEnum = 7
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_8 PriceDisplayFormatEnum = 8
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DECIMAL_9 PriceDisplayFormatEnum = 9
	//The following formats are fractional formats
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_256         PriceDisplayFormatEnum = 356
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_128         PriceDisplayFormatEnum = 228
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_64          PriceDisplayFormatEnum = 164
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_32_EIGHTHS  PriceDisplayFormatEnum = 140
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_32_QUARTERS PriceDisplayFormatEnum = 136
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_32_HALVES   PriceDisplayFormatEnum = 134
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_32          PriceDisplayFormatEnum = 132
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_16          PriceDisplayFormatEnum = 116
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_8           PriceDisplayFormatEnum = 108
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_4           PriceDisplayFormatEnum = 104
	PriceDisplayFormatEnum_PRICE_DISPLAY_FORMAT_DENOMINATOR_2           PriceDisplayFormatEnum = 102
)

/*==========================================================================*/
type SecurityTypeEnum int32

const (
	SecurityTypeEnum_SECURITY_TYPE_UNSET            SecurityTypeEnum = 0
	SecurityTypeEnum_SECURITY_TYPE_FUTURES          SecurityTypeEnum = 1
	SecurityTypeEnum_SECURITY_TYPE_STOCK            SecurityTypeEnum = 2
	SecurityTypeEnum_SECURITY_TYPE_FOREX            SecurityTypeEnum = 3 // CryptoCurrencies also go into this category
	SecurityTypeEnum_SECURITY_TYPE_INDEX            SecurityTypeEnum = 4
	SecurityTypeEnum_SECURITY_TYPE_FUTURES_STRATEGY SecurityTypeEnum = 5
	SecurityTypeEnum_SECURITY_TYPE_FUTURES_OPTION   SecurityTypeEnum = 7
	SecurityTypeEnum_SECURITY_TYPE_STOCK_OPTION     SecurityTypeEnum = 6
	SecurityTypeEnum_SECURITY_TYPE_INDEX_OPTION     SecurityTypeEnum = 8
	SecurityTypeEnum_SECURITY_TYPE_BOND             SecurityTypeEnum = 9
	SecurityTypeEnum_SECURITY_TYPE_MUTUAL_FUND      SecurityTypeEnum = 10
)

type PutCallEnum uint8

const (
	PutCallEnum_PC_UNSET PutCallEnum = 0
	PutCallEnum_PC_CALL  PutCallEnum = 1
	PutCallEnum_PC_PUT   PutCallEnum = 2
)

type SearchTypeEnum int32

const (
	SearchTypeEnum_SEARCH_TYPE_UNSET          SearchTypeEnum = 0
	SearchTypeEnum_SEARCH_TYPE_BY_SYMBOL      SearchTypeEnum = 1
	SearchTypeEnum_SEARCH_TYPE_BY_DESCRIPTION SearchTypeEnum = 2
)

/*==========================================================================*/
type HistoricalDataIntervalEnum int32

const (
	HistoricalDataIntervalEnum_INTERVAL_TICK       HistoricalDataIntervalEnum = 0
	HistoricalDataIntervalEnum_INTERVAL_1_SECOND   HistoricalDataIntervalEnum = 1
	HistoricalDataIntervalEnum_INTERVAL_2_SECONDS  HistoricalDataIntervalEnum = 2
	HistoricalDataIntervalEnum_INTERVAL_4_SECONDS  HistoricalDataIntervalEnum = 4
	HistoricalDataIntervalEnum_INTERVAL_5_SECONDS  HistoricalDataIntervalEnum = 5
	HistoricalDataIntervalEnum_INTERVAL_10_SECONDS HistoricalDataIntervalEnum = 10
	HistoricalDataIntervalEnum_INTERVAL_30_SECONDS HistoricalDataIntervalEnum = 30
	HistoricalDataIntervalEnum_INTERVAL_1_MINUTE   HistoricalDataIntervalEnum = 60
	HistoricalDataIntervalEnum_INTERVAL_5_MINUTE   HistoricalDataIntervalEnum = 300
	HistoricalDataIntervalEnum_INTERVAL_10_MINUTE  HistoricalDataIntervalEnum = 600
	HistoricalDataIntervalEnum_INTERVAL_15_MINUTE  HistoricalDataIntervalEnum = 900
	HistoricalDataIntervalEnum_INTERVAL_30_MINUTE  HistoricalDataIntervalEnum = 1800
	HistoricalDataIntervalEnum_INTERVAL_1_HOUR     HistoricalDataIntervalEnum = 3600
	HistoricalDataIntervalEnum_INTERVAL_2_HOURS    HistoricalDataIntervalEnum = 7200
	HistoricalDataIntervalEnum_INTERVAL_1_DAY      HistoricalDataIntervalEnum = 86400
	HistoricalDataIntervalEnum_INTERVAL_1_WEEK     HistoricalDataIntervalEnum = 604800
)

type HistoricalPriceDataRejectReasonCodeEnum int16

const (
	HistoricalPriceDataRejectReasonCodeEnum_HPDR_UNSET                                           HistoricalPriceDataRejectReasonCodeEnum = 0
	HistoricalPriceDataRejectReasonCodeEnum_HPDR_UNABLE_TO_SERVE_DATA_RETRY_IN_SPECIFIED_SECONDS HistoricalPriceDataRejectReasonCodeEnum = 1
	HistoricalPriceDataRejectReasonCodeEnum_HPDR_UNABLE_TO_SERVE_DATA_DO_NOT_RETRY               HistoricalPriceDataRejectReasonCodeEnum = 2
	HistoricalPriceDataRejectReasonCodeEnum_HPDR_DATA_REQUEST_OUTSIDE_BOUNDS_OF_AVAILABLE_DATA   HistoricalPriceDataRejectReasonCodeEnum = 3
	HistoricalPriceDataRejectReasonCodeEnum_HPDR_GENERAL_REJECT_ERROR                            HistoricalPriceDataRejectReasonCodeEnum = 4
)

/*==========================================================================*/
type TradeConditionEnum int8

const (
	TradeConditionEnum_TRADE_CONDITION_NONE                         TradeConditionEnum = 0
	TradeConditionEnum_TRADE_CONDITION_NON_LAST_UPDATE_EQUITY_TRADE TradeConditionEnum = 1
	TradeConditionEnum_TRADE_CONDITION_ODD_LOT_EQUITY_TRADE         TradeConditionEnum = 2
)
