package client

import (
	"context"
	"crypto/tls"
	"encoding"
	"github.com/moontrade/dtc-go/message"
	. "github.com/moontrade/dtc-go/v8/sc"
	"github.com/moontrade/log"
	"io"
	"math"
	"net"
	"sync"
	"time"
)

type Config struct {
	BufferSize       int
	HeartbeatSeconds int
	TLS              *tls.Config
	Logon            *LogonRequest
}

func NewConfig() *Config {
	return &Config{BufferSize: math.MaxUint16, HeartbeatSeconds: 15}
}

func (c *Config) SetBufferSize(size int) *Config {
	if size < 4096 {
		size = 4096
	}
	c.BufferSize = size
	return c
}

func (c *Config) SetHeartbeatSeconds(seconds int) *Config {
	if seconds < 5 {
		seconds = 5
	}
	c.HeartbeatSeconds = seconds
	if c.Logon != nil {
		c.Logon.SetHeartbeatIntervalInSeconds(int32(seconds))
	}
	return c
}

func (c *Config) SetTLS(config *tls.Config) *Config {
	if config == nil {
		config = &tls.Config{InsecureSkipVerify: true}
	}
	c.TLS = config
	return c
}

func (c *Config) SetLogon(request *LogonRequest) *Config {
	if c.HeartbeatSeconds < 10 {
		c.HeartbeatSeconds = 10
	}
	request.SetHeartbeatIntervalInSeconds(int32(c.HeartbeatSeconds))
	c.Logon = request
	return c
}

type Client struct {
	c      net.Conn
	cfg    *Config
	ctx    context.Context
	cancel context.CancelFunc
	data   MarketDataControl
	m      sync.RWMutex
}

type Historical struct{}

func Dial(addr string, config *Config) (*Client, error) {
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		return nil, err
	}
	c := &Client{c: conn, cfg: config}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	if err = c.start(); err != nil {
		_ = c.Close()
		return nil, err
	}
	return c, nil
}

func DialTLS(addr string, config *Config) (*Client, error) {
	if config.TLS == nil {
		config.TLS = &tls.Config{InsecureSkipVerify: true}
	}
	conn, err := tls.Dial("tcp4", addr, config.TLS)
	if err != nil {
		return nil, err
	}
	c := &Client{c: conn, cfg: config}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	if err = c.start(); err != nil {
		_ = c.Close()
		return nil, err
	}
	return c, nil
}

func (c *Client) Wait() {
	<-c.ctx.Done()
}

func (c *Client) Close() error {
	c.m.Lock()
	defer c.m.Unlock()
	if c.c != nil {
		err := c.c.Close()
		c.cancel()
		c.c = nil
		return err
	}
	return nil
}

func (c *Client) start() error {
	if c.cfg.BufferSize == 0 {
		c.cfg.BufferSize = 65536
	}
	if c.cfg.HeartbeatSeconds < 10 {
		c.cfg.HeartbeatSeconds = 10
	}
	request := NewEncodingRequest()
	request.SetEncoding(BINARY_WITH_VARIABLE_LENGTH_STRINGS)
	if err := c.Send(request); err != nil {
		return err
	}
	go func() {
		var (
			f   = message.Framer{}
			b   = make([]byte, c.cfg.BufferSize)
			n   int
			err error
		)
		for {
			n, err = c.c.Read(b)
			if err != nil {
				break
			}
			f.Push(b[0:n], c.onMessage)
		}

		log.Info("read loop ended")
		if err != nil {
			log.Error(err)
		}
		c.cancel()
	}()
	go func() {
		sleepFor := time.Second * time.Duration(c.cfg.HeartbeatSeconds)
	loop:
		for {
			select {
			case <-c.ctx.Done():
				break loop
			case <-time.After(sleepFor):
			}
			heartbeat := NewHeartbeat()
			heartbeat.SetCurrentDateTime(DateTime(time.Now().UnixNano() / 1000000000))
			if err := c.Send(heartbeat); err != nil {
				log.Error(err)
			}
		}
	}()
	return nil
}

func (c *Client) onMessage(b []byte) {
	_ = Handle(b, c)
}

func (c *Client) Send(msg encoding.BinaryMarshaler) error {
	b, _ := msg.MarshalBinary()
	n, e := c.c.Write(b)
	if e != nil {
		return e
	}
	if n != len(b) {
		if n <= 0 {
			return io.ErrShortWrite
		}
		b = b[n:]
		for len(b) > 0 {
			n, e = c.c.Write(b)
			if e != nil {
				return e
			}
			b = b[n:]
		}
	}
	return nil
}

func (c *Client) OnUnknown(b []byte, t uint16) error {
	log.Info("OnUnknown: size:%c type:%d", len(b), t)
	return nil
}

func (c *Client) OnError(b []byte, t uint16, err error) {
	log.Info("OnError: size:%c type:%d error:%c", len(b), t, err.Error())
}

// OnEncodingRequest Requirements: Not required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingRequest message is a message requesting to change the DTC
// encoding for messages.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
func (c *Client) OnEncodingRequest(msg EncodingRequest) error {
	log.Info("OnEncodingRequest: %s", message.JsonString(msg))
	return nil
}

// OnEncodingRequestExtended Requirements: Not required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingRequest message is a message requesting to change the DTC
// encoding for messages.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
func (c *Client) OnEncodingRequestExtended(msg EncodingRequestExtended) error {
	log.Info("OnEncodingRequestExtended: %s", message.JsonString(msg))
	return nil
}

// OnEncodingResponse Requirements: Required for Servers. Required for Clients if the Client
// needs to discover the encoding the Server uses.
//
// The EncodingResponse is a message from the Server to the Client, telling
// the Client what message encoding it must use to communicate with the Server.
// the Client what message encoding it must use to communicate with the Server.
//
// For the procedure to work with this message, refer to Encoding Request
// Sequence.
func (c *Client) OnEncodingResponse(msg EncodingResponse) error {
	log.Info("OnEncodingResponse: %s", message.JsonString(msg))
	log.Info("LogonRequest: %s", message.JsonString(c.cfg.Logon))
	return c.Send(c.cfg.Logon)
}

// OnLogonRequest The LogonRequest message is sent from the Client to the Server requesting
// to logon to the Server.
//
// This is the very first message the Client sends to the Server before being
// allowed to send any other message other than the EncodingRequest.
func (c *Client) OnLogonRequest(msg LogonRequest) error {
	log.Info("OnLogonRequest: %s", message.JsonString(msg))
	return nil
}

// OnLogonResponse This is a response message indicating either success or an error logging
// on to the Server.
func (c *Client) OnLogonResponse(msg LogonResponse) error {
	log.Info("OnLogonResponse: %s", message.JsonString(msg))
	_ = c.Send(NewSecurityDefinitionForSymbolRequest().
		SetExchange("CME").
		SetSymbol("ESM22_FUT_CME").
		SetRequestID(1))
	return c.Send(NewMarketDataRequest().
		SetExchange("CME").
		SetSymbol("ESM22_FUT_CME").
		SetRequestAction(SUBSCRIBE).
		SetSymbolID(1))
}

// OnLogoff A Logoff is a message which can be sent either by the Client or the Server
// to the other side. It indicates that the Client or the Server is logging
// off and going to be closing the connection.
//
// When one side receives this message, it should expect the connection will
// be closed. It should not be expected that any messages will follow the
// Logoff message, and it should close the network connection and consider
// it finished. The side receiving this message can send a Logoff message
// to the other side before closing the connection.
func (c *Client) OnLogoff(msg Logoff) error {
	log.Info("OnLogoff: %s", message.JsonString(msg))
	return nil
}

// OnHeartbeat Both the Client and the Server need to send to the other side a heartbeat
// at the interval specified by the HeartbeatIntervalInSeconds member in
// the LogonRequest.
//
// There are no required member fields to set in this message. The purpose
// of the Heartbeat message is so that the Client or the Server can determine
// whether the other side is still connected.
//
// It is recommended that if there is a loss of Heartbeat messages from the
// other side, for twice the amount of the HeartbeatIntervalInSeconds time
// that it is safe to assume that the other side is no longer present and
// the network socket should be then gracefully closed.
//
// The Server may choose to send a heartbeat message every second to the
// Client. In this particular case, it is recommended the Client use a minimum
// time of about 5 to 10 seconds without a heartbeat to determine the loss
// of the connection rather than the standard of twice the amount of the
// heartbeat time interval.
func (c *Client) OnHeartbeat(msg Heartbeat) error {
	//log.Info("OnHeartbeat: %s", message.JsonString(msg))
	return nil
}

// OnHeartbeatExtended Both the Client and the Server need to send to the other side a heartbeat
// at the interval specified by the HeartbeatIntervalInSeconds member in
// the LogonRequest.
//
// There are no required member fields to set in this message. The purpose
// of the Heartbeat message is so that the Client or the Server can determine
// whether the other side is still connected.
//
// It is recommended that if there is a loss of Heartbeat messages from the
// other side, for twice the amount of the HeartbeatIntervalInSeconds time
// that it is safe to assume that the other side is no longer present and
// the network socket should be then gracefully closed.
//
// The Server may choose to send a heartbeat message every second to the
// Client. In this particular case, it is recommended the Client use a minimum
// time of about 5 to 10 seconds without a heartbeat to determine the loss
// of the connection rather than the standard of twice the amount of the
// heartbeat time interval.
func (c *Client) OnHeartbeatExtended(msg HeartbeatExtended) error {
	//log.Info("OnHeartbeatExtended: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataFeedStatus The s_MarketDataFeed_STATUS message is an optional message sent by the
// Server to indicate the overall status of the market data feed. This status
// applies to all symbols that have been subscribed to for market data.
func (c *Client) OnMarketDataFeedStatus(msg MarketDataFeedStatus) error {
	log.Info("OnMarketDataFeedStatus: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketDataFeedSymbolStatus(msg MarketDataFeedSymbolStatus) error {
	log.Info("OnMarketDataFeedSymbolStatus: %s", message.JsonString(msg))
	return nil
}

// OnTradingSymbolStatus Sent by the Server to the Client to indicate the status of the symbol
// in regards to whether trading is open or closed or some other intermediate
// state.
func (c *Client) OnTradingSymbolStatus(msg TradingSymbolStatus) error {
	log.Info("OnTradingSymbolStatus: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataRequest The MarketDataRequest message will subscribe to market data for a particular
// Symbol or request a market data snapshot.
//
// The Server can also send market depth data in response to this message
// and not require a MarketDepthRequest.
func (c *Client) OnMarketDataRequest(msg MarketDataRequest) error {
	log.Info("OnMarketDataRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketDepthRequest(msg MarketDepthRequest) error {
	log.Info("OnMarketDepthRequest: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataReject The MarketDataReject message is sent by the Server to the Client to reject
// a MarketDataRequest message for any reason.
func (c *Client) OnMarketDataReject(msg MarketDataReject) error {
	a := c.data.Action(msg.SymbolID())
	if a == nil {
		//log.Warn("OnMarketDataReject: SymbolID:%d not found", msg.SymbolID())
		//return nil
	}
	log.Info("OnMarketDataReject: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataSnapshot The Server sends the MarketDataSnapshot message to the Client immediately
// after a successful MarketDataRequest message has been received from the
// Client and it has indicated to subscribe to the symbol or requested the
// snapshot of data.
//
// Any changes to the data fields within the MarketDataSnapshot message during
// the trading session will be sent by the Server to the Client through the
// corresponding MARKET_DATA_UPDATE_* messages.
//
// It is recommended that the MarketDataSnapshot be sent by the Server at
// the start of a new trading session.
//
// This message can be sent more often, however it is not intended to be
// sent frequently.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
//
// There is no need to send this when there is a new High or Low during the
// trading session. The Server should use the MarketDataUpdateSessionHigh
// or MarketDataUpdateSessionLow messages instead.
func (c *Client) OnMarketDataSnapshot(msg MarketDataSnapshot) error {
	a := c.data.Action(msg.SymbolID())
	if a == nil {
		//log.Warn("OnMarketDataSnapshot: SymbolID:%d not found", msg.SymbolID())
		//return nil
	}
	log.Info("OnMarketDataSnapshot: %s", message.JsonString(msg))
	return nil
}

// OnMarketDepthSnapshotLevel This is a message sent by Server to provide the initial market depth data
// entries to the Client after the Client subscribes to market data or separately
// subscribes to market depth data. The Client will need to separately subscribe
// to market depth data if the Server requires it.
//
// Each message provides a single entry of depth data. Therefore, the Server
// will send multiple MarketDepthSnapshotLevel messages in a series in order
// for the Client to build up its initial market depth book.
//
// The first message will be identified by the IsFirstMessageInBatch field
// being set to 1. The last message will be identified by the IsLastMessageInBatch
// field being set to 1.
//
// In the case where the market depth book is empty, the Server still needs
// to send through one single message with the SymbolID set, IsFirstMessageInBatch
// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
// be at the default values. The Client will understand this as an empty
// book.
func (c *Client) OnMarketDepthSnapshotLevel(msg MarketDepthSnapshotLevel) error {
	log.Info("OnMarketDepthSnapshotLevel: %s", message.JsonString(msg))
	return nil
}

// OnMarketDepthSnapshotLevelFloat This is a message sent by Server to provide the initial market depth data
// entries to the Client after the Client subscribes to market data or separately
// subscribes to market depth data. The Client will need to separately subscribe
// to market depth data if the Server requires it.
//
// Each message provides a single entry of depth data. Therefore, the Server
// will send multiple MarketDepthSnapshotLevelFloat messages in a series
// in order for the Client to build up its initial market depth book.
//
// The first message will be identified by the IsFirstMessageInBatch field
// being set to 1. The last message will be identified by the IsLastMessageInBatch
// field being set to 1.
//
// In the case where the market depth book is empty, the Server still needs
// to send through one single message with the SymbolID set, IsFirstMessageInBatch
// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
// be at the default values. The Client will understand this as an empty
// book.
func (c *Client) OnMarketDepthSnapshotLevelFloat(msg MarketDepthSnapshotLevelFloat) error {
	log.Info("OnMarketDepthSnapshotLevelFloat: %s", message.JsonString(msg))
	return nil
}

// OnMarketDepthUpdateLevel Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// Each MarketDepthUpdateLevel message updates one level of market depth
// on one side. An insert/update/delete model is used for market depth.
//
// The Client will need to determine the based upon the price, what particular
// market depth level is being updated, inserted or deleted.
//
// It is for this reason, that an insert/update is considered as one update
// type since it is possible to determine whether it is an insert or update
// based upon the existence of the price level in the existing market depth
// book on the Client side.
//
// What this means is that when the UpdateType field is MARKET_DEPTH_INSERT_UPDATE_LEVEL,
// it is considered an insert if the price level is not found on the particular
// side of the market depth being updated. It is considered an update, if
// the price level is found on the particular side of market depth being
// updated.
//
// This message uses a double datatype for the Price field. There is no level
// index. It is the responsibility of the Client to determine where in its
// market depth array it is maintaining where the insert/update/delete operation
// needs to occur.
//
// Since floating-point comparisons are not always precise, there should
// be a comparison made only to the number of decimal places the symbol specifies
// in its security definition. This can be determined through the SecurityDefinitionResponse::PriceDisplayFormat
// field.
func (c *Client) OnMarketDepthUpdateLevel(msg MarketDepthUpdateLevel) error {
	log.Info("OnMarketDepthUpdateLevel: %s", message.JsonString(msg))
	return nil
}

// OnMarketDepthUpdateLevelFloatWithMilliseconds Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is a more compact version of the MarketDepthUpdateLevel message.
// For the Price and Quantity fields, it uses a 4 byte float for compactness.
// It also supports millisecond precision for the timestamp.
func (c *Client) OnMarketDepthUpdateLevelFloatWithMilliseconds(msg MarketDepthUpdateLevelFloatWithMilliseconds) error {
	log.Info("OnMarketDepthUpdateLevelFloatWithMilliseconds: %s", message.JsonString(msg))
	return nil
}

// OnMarketDepthUpdateLevelNoTimestamp Sent by the Server to the Client to Update/Insert or Delete a particular
// market depth price level in the market depth book maintained by the Client.
// market depth price level in the market depth book maintained by the Client.
//
// This message is identical to the MarketDepthUpdateLevel message except
// it has no timestamp field. It needs to be sent when there is no change
// with the timestamp for the market depth update as compared to the prior
// update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received market depth update timestamp to know what the
// timestamp is for this message.
func (c *Client) OnMarketDepthUpdateLevelNoTimestamp(msg MarketDepthUpdateLevelNoTimestamp) error {
	log.Info("OnMarketDepthUpdateLevelNoTimestamp: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateTradeNoTimestamp This message is optional.
//
// Sent by the Server to the Client when a trade occurs. This message is
// identical to the MarketDataUpdateTrade_WITH_UNBUNDLED_INDICATOR_2 message
// except it does not have a timestamp. It needs to be sent when there is
// no change with the timestamp for the trade as compared to the prior trade.
// no change with the timestamp for the trade as compared to the prior trade.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received trade timestamp to know what the timestamp is for
// this message.
func (c *Client) OnMarketDataUpdateTradeNoTimestamp(msg MarketDataUpdateTradeNoTimestamp) error {
	log.Info("OnMarketDataUpdateTradeNoTimestamp: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateSessionSettlement Sent by the Server to the Client to update the session settlement price
// when the session settlement price changes.
func (c *Client) OnMarketDataUpdateSessionSettlement(msg MarketDataUpdateSessionSettlement) error {
	log.Info("OnMarketDataUpdateSessionSettlement: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateSessionOpen Sent by the Server to the Client to update the session Open.
func (c *Client) OnMarketDataUpdateSessionOpen(msg MarketDataUpdateSessionOpen) error {
	log.Info("OnMarketDataUpdateSessionOpen: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateSessionNumTrades Sent by the Server to the Client to update the trading session number
// of trades.
func (c *Client) OnMarketDataUpdateSessionNumTrades(msg MarketDataUpdateSessionNumTrades) error {
	log.Info("OnMarketDataUpdateSessionNumTrades: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateTradingSessionDate Sent by the Server to the Client to update the trading session Date.
func (c *Client) OnMarketDataUpdateTradingSessionDate(msg MarketDataUpdateTradingSessionDate) error {
	log.Info("OnMarketDataUpdateTradingSessionDate: %s", message.JsonString(msg))
	return nil
}

// OnMarketDepthReject The MarketDepthReject message is sent by the Server to the Client to reject
// a MarketDepthRequest message for any reason.
func (c *Client) OnMarketDepthReject(msg MarketDepthReject) error {
	log.Info("OnMarketDepthReject: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateTrade The Server sends this market data feed message to the Client when a trade
// occurs.
func (c *Client) OnMarketDataUpdateTrade(msg MarketDataUpdateTrade) error {
	log.Info("OnMarketDataUpdateTrade: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketDataUpdateTradeWithUnbundledIndicator(msg MarketDataUpdateTradeWithUnbundledIndicator) error {
	log.Info("OnMarketDataUpdateTradeWithUnbundledIndicator: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateTradeWithUnbundledIndicator2 Sent by the Server to the Client when a trade occurs. This message has
// additional fields as compared to the MarketDataUpdateTrade message and
// also supports microsecond time stamping.
func (c *Client) OnMarketDataUpdateTradeWithUnbundledIndicator2(msg MarketDataUpdateTradeWithUnbundledIndicator2) error {
	log.Info("OnMarketDataUpdateTradeWithUnbundledIndicator2: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateBidAsk The Server sends this market data feed message to the Client when the
// best bid or ask price or size changes.
func (c *Client) OnMarketDataUpdateBidAsk(msg MarketDataUpdateBidAsk) error {
	log.Info("OnMarketDataUpdateBidAsk: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketDataUpdateBidAskCompact(msg MarketDataUpdateBidAskCompact) error {
	log.Info("OnMarketDataUpdateBidAskCompact: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateBidAskFloatWithMicroseconds This message is optional.
//
// Sent by the Server to the Client when there is an update to the Bid Ask
// prices and/or quantities. This message is identical to the MarketDataUpdateBidAsk
// message except it does not have a timestamp. It needs to be sent when
// there is no change with the timestamp for the Bid Ask update as compared
// to the prior update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received Bid Ask update timestamp to know what the timestamp
// is for this message.
func (c *Client) OnMarketDataUpdateBidAskFloatWithMicroseconds(msg MarketDataUpdateBidAskFloatWithMicroseconds) error {
	log.Info("OnMarketDataUpdateBidAskFloatWithMicroseconds: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateBidAskNoTimeStamp This message is optional.
//
// Sent by the Server to the Client when there is an update to the Bid Ask
// prices and/or quantities. This message is identical to the MarketDataUpdateBidAsk
// message except it does not have a timestamp. It needs to be sent when
// there is no change with the timestamp for the Bid Ask update as compared
// to the prior update.
//
// When the Server sends this message to the Client, the Client needs to
// use the prior received Bid Ask update timestamp to know what the timestamp
// is for this message.
func (c *Client) OnMarketDataUpdateBidAskNoTimeStamp(msg MarketDataUpdateBidAskNoTimeStamp) error {
	log.Info("OnMarketDataUpdateBidAskNoTimeStamp: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateTradeCompact Sent by the Server to the Client when a trade occurs. This message is
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
// a more compact MarketDataUpdateTrade. For the price it uses a 4 byte float.
func (c *Client) OnMarketDataUpdateTradeCompact(msg MarketDataUpdateTradeCompact) error {
	log.Info("OnMarketDataUpdateTradeCompact: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateSessionVolume Sent by the Server to the Client when the session trade Volume needs to
// be updated.
//
// The recommended rule for the Server to notify the Client of a change with
// the session trade volume to maintain bandwidth efficiency, is as follows:
// When a trade occurs for a symbol subscribed to, the Server will send a
// MarketDataUpdateTrade message to the Client. The Client should then increment
// its session trade volume value for the symbol by the value in the Volume
// field in this message.
//
// The Server will assume the Client is doing this. Therefore, when a trade
// occurs and the session trade volume does not equal the prior session trade
// volume plus the Volume for the most recent trade sent to the Client, then
// the Server must send out a MarketDataUpdateSessionVolume message to the
// client since the client calculation of the session trade volume is no
// longer correct.
//
// It is assumed that the reason for this inconsistency is due to trades
// included within the session trade volume which have not been sent out
// as normal trades.
//
// The Server should also send this message out at the frequency that the
// Server determines, such as every minute if there also has been a trade
// at that time.
func (c *Client) OnMarketDataUpdateSessionVolume(msg MarketDataUpdateSessionVolume) error {
	log.Info("OnMarketDataUpdateSessionVolume: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateOpenInterest The MarketDataUpdateOpenInterest message is sent by the Server to the
// Client to update the OpenInterest field previously sent through the MarketDataSnapshot
// message.
func (c *Client) OnMarketDataUpdateOpenInterest(msg MarketDataUpdateOpenInterest) error {
	log.Info("OnMarketDataUpdateOpenInterest: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateSessionHigh Sent by the Server to the Client to update the session High as the High
// price changes throughout the session.
func (c *Client) OnMarketDataUpdateSessionHigh(msg MarketDataUpdateSessionHigh) error {
	log.Info("OnMarketDataUpdateSessionHigh: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateSessionLow Sent by the Server to the Client to update the session Low as the Low
// price changes throughout the session.
func (c *Client) OnMarketDataUpdateSessionLow(msg MarketDataUpdateSessionLow) error {
	log.Info("OnMarketDataUpdateSessionLow: %s", message.JsonString(msg))
	return nil
}

// OnMarketDataUpdateLastTradeSnapshot Sent by the Server to the Client to update the last trade price, volume
// and date-time fields under conditions when there is not a trade.
//
// This message type does not signify a trade has occurred. It should never
// be interpreted by the Client in that way.
func (c *Client) OnMarketDataUpdateLastTradeSnapshot(msg MarketDataUpdateLastTradeSnapshot) error {
	log.Info("OnMarketDataUpdateLastTradeSnapshot: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketOrdersRequest(msg MarketOrdersRequest) error {
	log.Info("OnMarketOrdersRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketOrdersReject(msg MarketOrdersReject) error {
	log.Info("OnMarketOrdersReject: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketOrdersAdd(msg MarketOrdersAdd) error {
	log.Info("OnMarketOrdersAdd: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketOrdersModify(msg MarketOrdersModify) error {
	log.Info("OnMarketOrdersModify: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketOrdersRemove(msg MarketOrdersRemove) error {
	log.Info("OnMarketOrdersRemove: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarketOrdersSnapshotMessageBoundary(msg MarketOrdersSnapshotMessageBoundary) error {
	log.Info("OnMarketOrdersSnapshotMessageBoundary: %s", message.JsonString(msg))
	return nil
}

// OnSubmitNewSingleOrder This message is used to submit a new single order into the market from
// the Client to the Server.
func (c *Client) OnSubmitNewSingleOrder(msg SubmitNewSingleOrder) error {
	log.Info("OnSubmitNewSingleOrder: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnSubmitFlattenPositionOrder(msg SubmitFlattenPositionOrder) error {
	log.Info("OnSubmitFlattenPositionOrder: %s", message.JsonString(msg))
	return nil
}

// OnCancelReplaceOrder This message is sent by the Client to the Server to cancel and replace
// an existing order. This is also known as an order modification.
//
// When the cancel and replace operation is completed, an OrderUpdate message
// is sent by the Server with the OrderUpdateReasonfield set to ORDER_CANCEL_REPLACE_COMPLETE.
// If the cancel and replace operation cannot be completed, an OrderUpdate
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
// message is sent by the Server with the OrderUpdateReason set to ORDER_CANCEL_REPLACE_REJECTED.
func (c *Client) OnCancelReplaceOrder(msg CancelReplaceOrder) error {
	log.Info("OnCancelReplaceOrder: %s", message.JsonString(msg))
	return nil
}

// OnCancelOrder This is a message from the Client to the Server requesting a previously
// sent order to be canceled.
func (c *Client) OnCancelOrder(msg CancelOrder) error {
	log.Info("OnCancelOrder: %s", message.JsonString(msg))
	return nil
}

// OnSubmitNewOCOOrder This is a message from the Client to the Server for submitting an order
// cancels order (OCO) pair into the market. What this means is when one
// of the orders is filled or canceled, the other order will be canceled.
// If one order partially fills, the other order will be reduced in quantity
// by the fill amount of the order that partially filled.
//
// A service provider must implement OCO orders on the server so that they
// can independently be modified (Cancel/Replace) and canceled independently
// using each order's distinct ServerOrderID. Although, if one of the orders
// is canceled by the Client, the other order will be canceled as well unless
// they have a parent order, as specified through the ParentTriggerClientOrderID
// field, in which case the other order should remain open.
//
// If the OCO order pair is rejected, this must be communicated through two
// separate OrderUpdate messages, 1 for each order, with the OrderUpdateReason
// set to NEW_ORDER_REJECTED.
func (c *Client) OnSubmitNewOCOOrder(msg SubmitNewOCOOrder) error {
	log.Info("OnSubmitNewOCOOrder: %s", message.JsonString(msg))
	return nil
}

// OnOpenOrdersRequest This is a message from the Client to the Server requesting the currently
// open orders.
//
// The Server will send open/working orders in response to this request through
// OrderUpdate messages.
//
// The Server will not return canceled or filled orders.
//
// When the Server responds to this request, it needs to respond with a separate
// OrderUpdate for each order.
//
// When the Server responds to this request, OrderUpdateReason in the OrderUpdate
// message must be set to OpenOrdersRequest_RESPONSE indicating the orders
// are being restated.
//
// If there are no Open orders, the Server will send back 1 OrderUpdate message
// with only the TotalNumberMessages, MessageNumber, RequestID, OrderUpdateReason,
// NoOrders = 1 fields set in the OrderUpdate message.
func (c *Client) OnOpenOrdersRequest(msg OpenOrdersRequest) error {
	log.Info("OnOpenOrdersRequest: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalOrderFillsRequest This is a message from the Client to the Server to request order fills/executions
// for an order or orders.
func (c *Client) OnHistoricalOrderFillsRequest(msg HistoricalOrderFillsRequest) error {
	log.Info("OnHistoricalOrderFillsRequest: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalOrderFillsReject If the Server is unable to serve the request for a HistoricalOrderFillsRequest
// message received, for a reason other than there not being any historical
// order fills, then send this message to the Client.
func (c *Client) OnHistoricalOrderFillsReject(msg HistoricalOrderFillsReject) error {
	log.Info("OnHistoricalOrderFillsReject: %s", message.JsonString(msg))
	return nil
}

// OnCurrentPositionsRequest This is a message from the Client to the Server to request the current
// open Trade Positions.
func (c *Client) OnCurrentPositionsRequest(msg CurrentPositionsRequest) error {
	log.Info("OnCurrentPositionsRequest: %s", message.JsonString(msg))
	return nil
}

// OnCurrentPositionsReject If the Server is unable to serve the request for an CurrentPositionsRequest
// message received, for a reason other than there not being any current
// Trade positions, then send this message to the Client.
//
// This must never be sent when there are actually no Trade Positions in
// the account or accounts requested.
func (c *Client) OnCurrentPositionsReject(msg CurrentPositionsReject) error {
	log.Info("OnCurrentPositionsReject: %s", message.JsonString(msg))
	return nil
}

// OnOrderUpdate The OrderUpdate is a unified message from the Server to the Client which
// communicates the complete details of an order, the Order Status, and the
// reason for sending the message (OrderUpdateReason).
//
// DTC uses this single unified message to provide an update for an order.
// The OrderUpdateReason field provides a clear indication for each reason
// this message is being sent.
func (c *Client) OnOrderUpdate(msg OrderUpdate) error {
	log.Info("OnOrderUpdate: %s", message.JsonString(msg))
	return nil
}

// OnOpenOrdersReject If the Server is unable to serve the request for an OpenOrdersRequest
// message received, for a reason other than there not being any open orders,
// then send this message to the Client.
func (c *Client) OnOpenOrdersReject(msg OpenOrdersReject) error {
	log.Info("OnOpenOrdersReject: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalOrderFillResponse This is a message from the Server to the Client providing an individual
// historical order fill in response to a HistoricalOrderFillsRequest message.
// historical order fill in response to a HistoricalOrderFillsRequest message.
//
// The Server is expected to send this message to the Client in response
// to a HistoricalOrderFillsRequest message even when there are no order
// fills to return. If there are no order fills to return, it needs to set
// the NoOrderFills field to 1.
func (c *Client) OnHistoricalOrderFillResponse(msg HistoricalOrderFillResponse) error {
	log.Info("OnHistoricalOrderFillResponse: %s", message.JsonString(msg))
	return nil
}

// OnPositionUpdate This is a message from the Server to the Client to report a Trade Position
// for a symbol in any Trade Account for the logged in Username.
//
// The Position Update message can either be solicited, in response to CurrentPositionsRequest.
// Or unsolicited as a Trade Position for a symbol changes during the connection
// to the Server. Each Trade Position is contained within a single message.
// to the Server. Each Trade Position is contained within a single message.
//
// When the server is responding with one or more PositionUpdate messages
// in response to a CurrentPositionsRequest message, it must not send any
// unsolicited PositionUpdate messages interleaved with the solicited PositionUpdate
// messages in response to the CurrentPositionsRequest message.
func (c *Client) OnPositionUpdate(msg PositionUpdate) error {
	log.Info("OnPositionUpdate: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnAddCorrectingOrderFill(msg AddCorrectingOrderFill) error {
	log.Info("OnAddCorrectingOrderFill: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnCorrectingOrderFillResponse(msg CorrectingOrderFillResponse) error {
	log.Info("OnCorrectingOrderFillResponse: %s", message.JsonString(msg))
	return nil
}

// OnTradeAccountsRequest This is a message from the Client to the Server to request all of the
// account identifiers for the logged in Username.
//
// If there are no accounts available, then the Server needs to respond with
// at least one TradeAccountResponse message containing an empty Trade Account.
// at least one TradeAccountResponse message containing an empty Trade Account.
func (c *Client) OnTradeAccountsRequest(msg TradeAccountsRequest) error {
	log.Info("OnTradeAccountsRequest: %s", message.JsonString(msg))
	return nil
}

// OnTradeAccountResponse This is a message from the Server to the Client in response to a TradeAccountsRequest
// message, providing a single trade account. There is one message for each
// trade account.
func (c *Client) OnTradeAccountResponse(msg TradeAccountResponse) error {
	log.Info("OnTradeAccountResponse: %s", message.JsonString(msg))
	return nil
}

// OnExchangeListRequest This is a message from the Client to the Server to request a list of all
// available exchanges from the Server.
//
// The server will respond with a separate ExchangeListResponse message for
// each exchange.
//
// In the case where the Server does not specify an exchange with its symbols,
// then the Server should provide a single response with an empty Exchange.
// then the Server should provide a single response with an empty Exchange.
func (c *Client) OnExchangeListRequest(msg ExchangeListRequest) error {
	log.Info("OnExchangeListRequest: %s", message.JsonString(msg))
	return nil
}

// OnExchangeListResponse The server will return this message for each supported exchange.
//
// If there are no exchanges to return in response to a request, send through
// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
// all other members in the default state and the Client will recognize there
// are no Exchanges.
func (c *Client) OnExchangeListResponse(msg ExchangeListResponse) error {
	log.Info("OnExchangeListResponse: %s", message.JsonString(msg))
	return nil
}

// OnSymbolsForExchangeRequest This is a message from the Client to the Server to request all of the
// Symbols for a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
func (c *Client) OnSymbolsForExchangeRequest(msg SymbolsForExchangeRequest) error {
	log.Info("OnSymbolsForExchangeRequest: %s", message.JsonString(msg))
	return nil
}

// OnUnderlyingSymbolsForExchangeRequest This is a message from the Client to the Server to request all of the
// underlying symbols on a particular Exchange. For example, all of the underlying
// futures symbols on a particular Exchange.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
func (c *Client) OnUnderlyingSymbolsForExchangeRequest(msg UnderlyingSymbolsForExchangeRequest) error {
	log.Info("OnUnderlyingSymbolsForExchangeRequest: %s", message.JsonString(msg))
	return nil
}

// OnSymbolsForUnderlyingRequest This is a message from the Client to the Server for requesting all of
// the symbols for a particular underlying symbol.
//
// For example, all of the futures contracts for a particular underlying
// futures symbol or all of the option symbols for a specific futures or
// stock symbol.
//
// The server will return a SecurityDefinitionResponse message to the Client
// for each Symbol returned.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
func (c *Client) OnSymbolsForUnderlyingRequest(msg SymbolsForUnderlyingRequest) error {
	log.Info("OnSymbolsForUnderlyingRequest: %s", message.JsonString(msg))
	return nil
}

// OnSymbolSearchRequest The SymbolSearchRequest message is sent by the Client to the Server to
// return Security Definitions matching the specified SecurityType and Exchange
// and where the Symbol or Description contains the specified SearchText.
//
// The SearchText can search either the Symbol or the Description field in
// the SecurityDefinitionResponse message.
//
// In either case there does not need to be an exact match. The SearchText
// only needs to be contained within the Symbol or the Description depending
// upon which field is being searched.
//
// The Server returns SecurityDefinitionResponse messages for all Symbols
// which match.
//
// If there are no matches, the Server needs to send a SecurityDefinitionResponse
// message to the Client with with all fields at their default values except
// for the RequestID and IsFinalMessage fields set. This will be a clear
// indication to the Client that the request returned no matches.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
func (c *Client) OnSymbolSearchRequest(msg SymbolSearchRequest) error {
	log.Info("OnSymbolSearchRequest: %s", message.JsonString(msg))
	return nil
}

// OnSecurityDefinitionForSymbolRequest This is a message from the Client to the Server for requesting Security
// Definition data for a specific symbol.
//
// The Server will return a single SecurityDefinitionResponse message in
// response to this request.
//
// The Client must always send a SecurityDefinitionForSymbolRequest message
// to the Server in order to obtain the IntegerToFloatPriceDivisor value
// in case the Server uses the integer market data messages.
//
// If the Server is rejecting this request, then it needs to send a SecurityDefinitionReject
// message to the Client.
func (c *Client) OnSecurityDefinitionForSymbolRequest(msg SecurityDefinitionForSymbolRequest) error {
	log.Info("OnSecurityDefinitionForSymbolRequest: %s", message.JsonString(msg))
	return nil
}

// OnSecurityDefinitionResponse This is a response from the Server in response to a SymbolsForExchangeRequest,
// UNDERLYING_SymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest message.
//
// If there are no symbols to return in response to a request, the Server
// needs to send through one of these messages with the RequestID set to
// the same RequestID value that the request message set it to, and IsFinalMessage
// = 1. Leave all other member fields in the default state and the Client
// will recognize there are no symbols.
//
// The Client must always send a SecurityDefinitionForSymbolRequest message
// to the Server to obtain the IntegerToFloatPriceDivisor and FloatToIntPriceMultiplier
// values in the Security Definition Response message when the Server uses
// the integer market data and order messages.
func (c *Client) OnSecurityDefinitionResponse(msg SecurityDefinitionResponse) error {
	log.Info("OnSecurityDefinitionResponse: %s", message.JsonString(msg))
	return nil
}

// OnSecurityDefinitionReject This is a message from the Server to the Client indicating the Server
// is rejecting one of the following messages: SymbolsForExchangeRequest,
// UnderlyingSymbolsForExchangeRequest, SymbolsForUnderlyingRequest, SecurityDefinitionForSymbolRequest,
// SymbolSearchRequest.
//
// If there are no symbols to send in response to one of these messages above,
// then the Server should not send a SecurityDefinitionReject message and
// instead send a SecurityDefinitionResponse with only the RequestID and
// IsFinalMessage fields set. This will be a clear indication to the Client
// that the request returned no Symbols.
func (c *Client) OnSecurityDefinitionReject(msg SecurityDefinitionReject) error {
	log.Info("OnSecurityDefinitionReject: %s", message.JsonString(msg))
	return nil
}

// OnAccountBalanceRequest This is a message from the Client to the Server to request Trade Account
// Balance data.
//
// The Server will respond with an AccountBalanceUpdate or reject the request.
// The Server will respond with an AccountBalanceUpdate or reject the request.
//
// The Server will set the RequestID in the AccountBalanceUpdate message
// to match the RequestID in the AccountBalanceRequest.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When unsolicited AccountBalanceUpdate messages are sent by the Server,
// the RequestID will be 0.
func (c *Client) OnAccountBalanceRequest(msg AccountBalanceRequest) error {
	log.Info("OnAccountBalanceRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnAccountBalanceReject(msg AccountBalanceReject) error {
	log.Info("OnAccountBalanceReject: %s", message.JsonString(msg))
	return nil
}

// OnAccountBalanceUpdate This is an optional message from the Server to Client to provide Account
// Balance information for a particular Trade Account. The server needs to
// provide a separate message for each Trade Account associated with the
// logged in username if it supports Account Balance updates.
//
// The Server will respond with an AccountBalanceUpdate in response to a
// AccountBalanceRequest message. The Server will set the RequestID in the
// AccountBalanceUpdate message to match the RequestID in the AccountBalanceRequest
// message.
//
// The Server will periodically send AccountBalanceUpdate messages as the
// Account Balance data changes. The frequency of the updates is determined
// by the Server. Account Balance updates are considered automatically subscribed
// to. When an unsolicited AccountBalanceUpdate message is sent, the RequestID
// field will be 0.
//
// When the server is responding with one or more AccountBalanceUpdate messages
// in response to a AccountBalanceRequest message, it must not send any unsolicited
// AccountBalanceUpdate messages interleaved with the solicited AccountBalanceUpdate
// messages in response to the AccountBalanceRequest message.
func (c *Client) OnAccountBalanceUpdate(msg AccountBalanceUpdate) error {
	log.Info("OnAccountBalanceUpdate: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnAccountBalanceAdjustment(msg AccountBalanceAdjustment) error {
	log.Info("OnAccountBalanceAdjustment: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnAccountBalanceAdjustmentReject(msg AccountBalanceAdjustmentReject) error {
	log.Info("OnAccountBalanceAdjustmentReject: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnAccountBalanceAdjustmentComplete(msg AccountBalanceAdjustmentComplete) error {
	log.Info("OnAccountBalanceAdjustmentComplete: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalAccountBalancesRequest This is a message from the Client to the Server to request a history of
// Cash Balance changes for the specified Trade Account.
//
// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// messages or reject he request with a message.
//
// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
// message to match the RequestID in the HistoricalAccountBalancesRequest.
// message to match the RequestID in the HistoricalAccountBalancesRequest.
func (c *Client) OnHistoricalAccountBalancesRequest(msg HistoricalAccountBalancesRequest) error {
	log.Info("OnHistoricalAccountBalancesRequest: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalAccountBalancesReject This is a message from the Server to the Client to reject a HistoricalAccountBalancesRequest
// request.
func (c *Client) OnHistoricalAccountBalancesReject(msg HistoricalAccountBalancesReject) error {
	log.Info("OnHistoricalAccountBalancesReject: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalAccountBalanceResponse(msg HistoricalAccountBalanceResponse) error {
	log.Info("OnHistoricalAccountBalanceResponse: %s", message.JsonString(msg))
	return nil
}

// OnUserMessage This message from the Server to the Client is for providing a message
// to the user.
//
// This message can be sent even before a LogonResponse.
func (c *Client) OnUserMessage(msg UserMessage) error {
	log.Info("OnUserMessage: %s", message.JsonString(msg))
	return nil
}

// OnGeneralLogMessage This message from the Server to the Client is a message which is to be
// added to a log file indicating information from the server. For example,
// if there are informational messages to provide during the process of a
// logon, this can be used to send those messages to a Client. A Client should
// never implement this message as a pop-up type message. Instead, it should
// be treated as a lower-level log type message.
//
// This message can be sent even before a LogonResponse is given.
func (c *Client) OnGeneralLogMessage(msg GeneralLogMessage) error {
	log.Info("OnGeneralLogMessage: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnAlertMessage(msg AlertMessage) error {
	log.Info("OnAlertMessage: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnJournalEntryAdd(msg JournalEntryAdd) error {
	log.Info("OnJournalEntryAdd: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnJournalEntriesRequest(msg JournalEntriesRequest) error {
	log.Info("OnJournalEntriesRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnJournalEntriesReject(msg JournalEntriesReject) error {
	log.Info("OnJournalEntriesReject: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnJournalEntryResponse(msg JournalEntryResponse) error {
	log.Info("OnJournalEntryResponse: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalPriceDataRequest This is a message from the Client to the Server for requesting historical
// price data.
//
// This request can be on the same or a separate network socket connection
// compared to the streaming market data. This is going to be specified by
// the Server.
func (c *Client) OnHistoricalPriceDataRequest(msg HistoricalPriceDataRequest) error {
	log.Info("OnHistoricalPriceDataRequest: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalPriceDataResponseHeader When a historical price data request is not rejected, this message header
// will begin the historical price data response from the Server. There will
// be one HistoricalPriceDataResponseHeader message sent ahead of the HistoricalPriceDataRecordResponse
// / HistoricalPriceDataTickRecordResponse messages. If the NoRecordsToReturn
// field is nonzero, then there are no further records that will be sent
// by the Server in response to the request by the Client.
//
// This message is never compressed.
func (c *Client) OnHistoricalPriceDataResponseHeader(msg HistoricalPriceDataResponseHeader) error {
	log.Info("OnHistoricalPriceDataResponseHeader: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalPriceDataReject When the Server rejects a historical price data request from the Client,
// a HistoricalPriceDataReject message will be sent.
//
// This message is never compressed.
func (c *Client) OnHistoricalPriceDataReject(msg HistoricalPriceDataReject) error {
	log.Info("OnHistoricalPriceDataReject: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalPriceDataRecordResponse The HistoricalPriceDataTickRecordResponse message is used when the RecordInterval
// field in a historical data request message is set to a value greater than
// INTERVAL_TICK. For example, if the RecordInterval is INTERVAL_1_MINUTE,
// then a message of this type will contain data for a 1 minute timeframe
// with a start time specified by the StartDateTime field.
//
// Even when RecordInterval is INTERVAL_TICK, the HistoricalPriceDataTickRecordResponse
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
// message can still be used instead of HistoricalPriceDataTickRecordResponse.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
func (c *Client) OnHistoricalPriceDataRecordResponse(msg HistoricalPriceDataRecordResponse) error {
	log.Info("OnHistoricalPriceDataRecordResponse: %s", message.JsonString(msg))
	return nil
}

// OnHistoricalPriceDataTickRecordResponse This is the response message when the RecordInterval field in a historical
// data request message is set to INTERVAL_TICK.
//
// If the Server does not support 1 Tick historical data or does not have
// 1 Tick historical data for the specified time period, it can respond with
// HistoricalPriceDataRecordResponse messages instead. The Server must only
// respond with messages of one type in response to a particular historical
// price data request.
//
// This message can be part of a compressed series of messages of this same
// type, if the Client requested compression be used.
func (c *Client) OnHistoricalPriceDataTickRecordResponse(msg HistoricalPriceDataTickRecordResponse) error {
	log.Info("OnHistoricalPriceDataTickRecordResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalPriceDataResponseTrailer(msg HistoricalPriceDataResponseTrailer) error {
	log.Info("OnHistoricalPriceDataResponseTrailer: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalMarketDepthDataRequest(msg HistoricalMarketDepthDataRequest) error {
	log.Info("OnHistoricalMarketDepthDataRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalMarketDepthDataResponseHeader(msg HistoricalMarketDepthDataResponseHeader) error {
	log.Info("OnHistoricalMarketDepthDataResponseHeader: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalMarketDepthDataReject(msg HistoricalMarketDepthDataReject) error {
	log.Info("OnHistoricalMarketDepthDataReject: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalMarketDepthDataRecordResponse(msg HistoricalMarketDepthDataRecordResponse) error {
	log.Info("OnHistoricalMarketDepthDataRecordResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalTradesRequest(msg HistoricalTradesRequest) error {
	log.Info("OnHistoricalTradesRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalTradesReject(msg HistoricalTradesReject) error {
	log.Info("OnHistoricalTradesReject: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnHistoricalTradesResponse(msg HistoricalTradesResponse) error {
	log.Info("OnHistoricalTradesResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnReplayChartData(msg ReplayChartData) error {
	log.Info("OnReplayChartData: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnReplayChartDataPerformAction(msg ReplayChartDataPerformAction) error {
	log.Info("OnReplayChartDataPerformAction: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnReplayChartDataStatus(msg ReplayChartDataStatus) error {
	log.Info("OnReplayChartDataStatus: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnRequestNumCurrentClientConnections(msg RequestNumCurrentClientConnections) error {
	log.Info("OnRequestNumCurrentClientConnections: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnNumCurrentClientConnectionsResponse(msg NumCurrentClientConnectionsResponse) error {
	log.Info("OnNumCurrentClientConnectionsResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnClientDeviceUpdate(msg ClientDeviceUpdate) error {
	log.Info("OnClientDeviceUpdate: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnInterprocessSynchronizationRemoteState(msg InterprocessSynchronizationRemoteState) error {
	log.Info("OnInterprocessSynchronizationRemoteState: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnInterprocessSynchronizationSnapshotRequest(msg InterprocessSynchronizationSnapshotRequest) error {
	log.Info("OnInterprocessSynchronizationSnapshotRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnInterprocessSynchronizationTradeActivityRequest(msg InterprocessSynchronizationTradeActivityRequest) error {
	log.Info("OnInterprocessSynchronizationTradeActivityRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnWriteIntradayDataFileSessionValue(msg WriteIntradayDataFileSessionValue) error {
	log.Info("OnWriteIntradayDataFileSessionValue: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnSCConfigurationSynchronization(msg SCConfigurationSynchronization) error {
	log.Info("OnSCConfigurationSynchronization: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnDownloadHistoricalOrderFillAndAccountBalanceData(msg DownloadHistoricalOrderFillAndAccountBalanceData) error {
	log.Info("OnDownloadHistoricalOrderFillAndAccountBalanceData: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeOrder(msg TradeOrder) error {
	log.Info("OnTradeOrder: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnIndividualTradePosition(msg IndividualTradePosition) error {
	log.Info("OnIndividualTradePosition: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradePositionConsolidated(msg TradePositionConsolidated) error {
	log.Info("OnTradePositionConsolidated: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeActivityData(msg TradeActivityData) error {
	log.Info("OnTradeActivityData: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataRequest(msg TradeAccountDataRequest) error {
	log.Info("OnTradeAccountDataRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataResponse(msg TradeAccountDataResponse) error {
	log.Info("OnTradeAccountDataResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataUpdate(msg TradeAccountDataUpdate) error {
	log.Info("OnTradeAccountDataUpdate: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataDelete(msg TradeAccountDataDelete) error {
	log.Info("OnTradeAccountDataDelete: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataSymbolLimitsResponse(msg TradeAccountDataSymbolLimitsResponse) error {
	log.Info("OnTradeAccountDataSymbolLimitsResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataSymbolLimitsUpdate(msg TradeAccountDataSymbolLimitsUpdate) error {
	log.Info("OnTradeAccountDataSymbolLimitsUpdate: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataSymbolCommissionResponse(msg TradeAccountDataSymbolCommissionResponse) error {
	log.Info("OnTradeAccountDataSymbolCommissionResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataSymbolCommissionUpdate(msg TradeAccountDataSymbolCommissionUpdate) error {
	log.Info("OnTradeAccountDataSymbolCommissionUpdate: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataAuthorizedUsernameResponse(msg TradeAccountDataAuthorizedUsernameResponse) error {
	log.Info("OnTradeAccountDataAuthorizedUsernameResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataAuthorizedUsernameAdd(msg TradeAccountDataAuthorizedUsernameAdd) error {
	log.Info("OnTradeAccountDataAuthorizedUsernameAdd: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataAuthorizedUsernameRemove(msg TradeAccountDataAuthorizedUsernameRemove) error {
	log.Info("OnTradeAccountDataAuthorizedUsernameRemove: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataUsernameToShareWithResponse(msg TradeAccountDataUsernameToShareWithResponse) error {
	log.Info("OnTradeAccountDataUsernameToShareWithResponse: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataUsernameToShareWithAdd(msg TradeAccountDataUsernameToShareWithAdd) error {
	log.Info("OnTradeAccountDataUsernameToShareWithAdd: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataUsernameToShareWithRemove(msg TradeAccountDataUsernameToShareWithRemove) error {
	log.Info("OnTradeAccountDataUsernameToShareWithRemove: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataResponseTrailer(msg TradeAccountDataResponseTrailer) error {
	log.Info("OnTradeAccountDataResponseTrailer: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnTradeAccountDataUpdateOperationComplete(msg TradeAccountDataUpdateOperationComplete) error {
	log.Info("OnTradeAccountDataUpdateOperationComplete: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnProcessedFillIdentifier(msg ProcessedFillIdentifier) error {
	log.Info("OnProcessedFillIdentifier: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnFlattenPositionsForTradeAccount(msg FlattenPositionsForTradeAccount) error {
	log.Info("OnFlattenPositionsForTradeAccount: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnUserInformation(msg UserInformation) error {
	log.Info("OnUserInformation: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarginDataRequest(msg MarginDataRequest) error {
	log.Info("OnMarginDataRequest: %s", message.JsonString(msg))
	return nil
}

func (c *Client) OnMarginDataResponse(msg MarginDataResponse) error {
	log.Info("OnMarginDataResponse: %s", message.JsonString(msg))
	return nil
}
