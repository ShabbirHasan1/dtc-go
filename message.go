package dtc

const (
	// Authentication and connection monitoring
	MessageType_LOGON_REQUEST     uint16 = 1
	MessageType_LOGON_RESPONSE    uint16 = 2
	MessageType_HEARTBEAT         uint16 = 3
	MessageType_LOGOFF            uint16 = 5
	MessageType_ENCODING_REQUEST  uint16 = 6
	MessageType_ENCODING_RESPONSE uint16 = 7

	// Market data
	MessageType_MARKET_DATA_REQUEST  uint16 = 101
	MessageType_MARKET_DATA_RESPONSE uint16 = 103
	MessageType_MARKET_DATA_SNAPSHOT uint16 = 104

	MessageType_MARKET_DATA_UPDATE_TRADE         uint16 = 107
	MessageType_MARKET_DATA_UPDATE_TRADE_COMPACT uint16 = 112

	MessageType_MARKET_DATA_UPDATE_LAST_TRADE_SNAPSHOT              uint16 = 134
	MessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR   uint16 = 137
	MessageType_MARKET_DATA_UPDATE_TRADE_WITH_UNBUNDLED_INDICATOR_2 uint16 = 146
	MessageType_MARKET_DATA_UPDATE_TRADE_NO_TIMESTAMP               uint16 = 142

	MessageType_MARKET_DATA_UPDATE_BID_ASK         uint16 = 108
	MessageType_MARKET_DATA_UPDATE_BID_ASK_COMPACT uint16 = 117
	//127
	MessageType_MARKET_DATA_UPDATE_BID_ASK_NO_TIMESTAMP            uint16 = 143
	MessageType_MARKET_DATA_UPDATE_BID_ASK_FLOAT_WITH_MICROSECONDS uint16 = 144

	MessageType_MARKET_DATA_UPDATE_SESSION_OPEN uint16 = 120
	//128
	MessageType_MARKET_DATA_UPDATE_SESSION_HIGH uint16 = 114
	//129
	MessageType_MARKET_DATA_UPDATE_SESSION_LOW uint16 = 115
	//130
	MessageType_MARKET_DATA_UPDATE_SESSION_VOLUME     uint16 = 113
	MessageType_MARKET_DATA_UPDATE_OPEN_INTEREST      uint16 = 124
	MessageType_MARKET_DATA_UPDATE_SESSION_SETTLEMENT uint16 = 119
	//131
	MessageType_MARKET_DATA_UPDATE_SESSION_NUM_TRADES   uint16 = 135
	MessageType_MARKET_DATA_UPDATE_TRADING_SESSION_DATE uint16 = 136

	MessageType_MARKET_DEPTH_REQUEST        uint16 = 102
	MessageType_MARKET_DEPTH_REJECT         uint16 = 121
	MessageType_MARKET_DEPTH_SNAPSHOT_LEVEL uint16 = 122
	//132
	MessageType_MARKET_DEPTH_SNAPSHOT_LEVEL_FLOAT                 uint16 = 145
	MessageType_MARKET_DEPTH_UPDATE_LEVEL                         uint16 = 106
	MessageType_MARKET_DEPTH_UPDATE_LEVEL_FLOAT_WITH_MILLISECONDS uint16 = 140
	MessageType_MARKET_DEPTH_UPDATE_LEVEL_NO_TIMESTAMP            uint16 = 141
	//133

	MessageType_MARKET_DATA_FEED_STATUS        uint16 = 100
	MessageType_MARKET_DATA_FEED_SYMBOL_STATUS uint16 = 116
	MessageType_TRADING_SYMBOL_STATUS          uint16 = 138

	MessageType_MARKET_ORDERS_REQUEST uint16 = 150
	MessageType_MARKET_ORDERS_REJECT  uint16 = 151

	MessageType_MARKET_ORDERS_ADD                       uint16 = 152
	MessageType_MARKET_ORDERS_MODIFY                    uint16 = 153
	MessageType_MARKET_ORDERS_REMOVE                    uint16 = 154
	MessageType_MARKET_ORDERS_SNAPSHOT_MESSAGE_BOUNDARY uint16 = 155

	// Order entry and modification
	MessageType_SUBMIT_NEW_SINGLE_ORDER uint16 = 208
	//206

	MessageType_SUBMIT_NEW_OCO_ORDER uint16 = 201
	//207
	MessageType_SUBMIT_FLATTEN_POSITION_ORDER       uint16 = 209
	MessageType_FLATTEN_POSITIONS_FOR_TRADE_ACCOUNT uint16 = 210

	MessageType_CANCEL_ORDER uint16 = 203

	MessageType_CANCEL_REPLACE_ORDER uint16 = 204
	//205

	// Trading related
	MessageType_OPEN_ORDERS_REQUEST uint16 = 300
	MessageType_OPEN_ORDERS_REJECT  uint16 = 302

	MessageType_ORDER_UPDATE uint16 = 301

	MessageType_HISTORICAL_ORDER_FILLS_REQUEST uint16 = 303
	MessageType_HISTORICAL_ORDER_FILLS_REJECT  uint16 = 308
	MessageType_HISTORICAL_ORDER_FILL_RESPONSE uint16 = 304

	MessageType_CURRENT_POSITIONS_REQUEST      uint16 = 305
	MessageType_CURRENT_POSITIONS_REJECT       uint16 = 307
	MessageType_POSITION_UPDATE                uint16 = 306
	MessageType_ADD_CORRECTING_ORDER_FILL      uint16 = 309
	MessageType_CORRECTING_ORDER_FILL_RESPONSE uint16 = 310

	// Account list
	MessageType_TRADE_ACCOUNTS_REQUEST uint16 = 400
	MessageType_TRADE_ACCOUNT_RESPONSE uint16 = 401

	// Symbol discovery and security definitions
	MessageType_EXCHANGE_LIST_REQUEST  uint16 = 500
	MessageType_EXCHANGE_LIST_RESPONSE uint16 = 501

	MessageType_SYMBOLS_FOR_EXCHANGE_REQUEST            uint16 = 502
	MessageType_UNDERLYING_SYMBOLS_FOR_EXCHANGE_REQUEST uint16 = 503
	MessageType_SYMBOLS_FOR_UNDERLYING_REQUEST          uint16 = 504
	MessageType_SECURITY_DEFINITION_FOR_SYMBOL_REQUEST  uint16 = 506
	MessageType_SECURITY_DEFINITION_RESPONSE            uint16 = 507

	MessageType_SYMBOL_SEARCH_REQUEST uint16 = 508

	MessageType_SECURITY_DEFINITION_REJECT uint16 = 509

	// Account Balance Data
	MessageType_ACCOUNT_BALANCE_REQUEST             uint16 = 601
	MessageType_ACCOUNT_BALANCE_REJECT              uint16 = 602
	MessageType_ACCOUNT_BALANCE_UPDATE              uint16 = 600
	MessageType_ACCOUNT_BALANCE_ADJUSTMENT          uint16 = 607
	MessageType_ACCOUNT_BALANCE_ADJUSTMENT_REJECT   uint16 = 608
	MessageType_ACCOUNT_BALANCE_ADJUSTMENT_COMPLETE uint16 = 609
	MessageType_HISTORICAL_ACCOUNT_BALANCES_REQUEST uint16 = 603
	MessageType_HISTORICAL_ACCOUNT_BALANCES_REJECT  uint16 = 604
	MessageType_HISTORICAL_ACCOUNT_BALANCE_RESPONSE uint16 = 606

	// Logging
	MessageType_USER_MESSAGE        uint16 = 700
	MessageType_GENERAL_LOG_MESSAGE uint16 = 701
	MessageType_ALERT_MESSAGE       uint16 = 702

	MessageType_JOURNAL_ENTRY_ADD       uint16 = 703
	MessageType_JOURNAL_ENTRIES_REQUEST uint16 = 704
	MessageType_JOURNAL_ENTRIES_REJECT  uint16 = 705
	MessageType_JOURNAL_ENTRY_RESPONSE  uint16 = 706

	// Historical price data
	MessageType_HISTORICAL_PRICE_DATA_REQUEST              uint16 = 800
	MessageType_HISTORICAL_PRICE_DATA_RESPONSE_HEADER      uint16 = 801
	MessageType_HISTORICAL_PRICE_DATA_REJECT               uint16 = 802
	MessageType_HISTORICAL_PRICE_DATA_RECORD_RESPONSE      uint16 = 803
	MessageType_HISTORICAL_PRICE_DATA_TICK_RECORD_RESPONSE uint16 = 804
	//805
	//806
	MessageType_HISTORICAL_PRICE_DATA_RESPONSE_TRAILER uint16 = 807

	// Historical market depth data
	MessageType_HISTORICAL_MARKET_DEPTH_DATA_REQUEST         uint16 = 900
	MessageType_HISTORICAL_MARKET_DEPTH_DATA_RESPONSE_HEADER uint16 = 901
	MessageType_HISTORICAL_MARKET_DEPTH_DATA_REJECT          uint16 = 902
	MessageType_HISTORICAL_MARKET_DEPTH_DATA_RECORD_RESPONSE uint16 = 903
)

type Header struct {
	size uint16
	typ  uint16
}

type Message interface {
	Size() uint16
	Type() uint16
	ProtocolVersion() int32
}
