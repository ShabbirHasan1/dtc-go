#![allow(dead_code)]
#![allow(unused_variables)]

mod client;
mod connection;
mod error;
mod message;
mod sierra_chart;
mod util;
mod v8;

pub use client::*;
use log::info;

use message::*;
pub use message::{Message, Parsed};

use std::io;

use ntex::{channel::mpsc, codec, connect, rt};
use rustls::{ClientConfig, OwnedTrustAnchor, RootCertStore};

use clap::Parser;

pub use connection::Connection;
pub use error::Error;

/// DTC client/server
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    /// Name of the person to greet
    #[clap(short, long)]
    name: String,

    /// Number of times to greet
    #[clap(short, long, default_value_t = 1)]
    count: u8,
}

#[ntex::main]
async fn main() -> io::Result<()> {
    let args = Args::parse();
    println!("Hello {} : count({:})!", args.name, args.count);

    std::env::set_var("RUST_LOG", "trace");
    env_logger::init();

    rt::spawn(async {
        // ntex::time::sleep(Duration::from_millis(1000)).await;
        info!("done");
    })
    .await?;

    // ntex::time::sleep(Duration::from_millis(1000)).await;
    // info!("done");

    Ok(())
}

// fn main() {
//     let mut logon_request = LogonRequestFixed::new().unwrap();
//     let mut logon_request_vls = LogonRequestVLS::new().unwrap();
//     logon_request.set_username("iamme");
//     logon_request_vls.set_username(logon_request.username());
//     logon_request.copy_to(&mut logon_request_vls);
//     println!(
//         "{:}: {}",
//         logon_request.username().len(),
//         logon_request.username()
//     );
//     println!(
//         "{:}: {}",
//         logon_request_vls.username().len(),
//         logon_request_vls.username()
//     );
// }

#[cfg(test)]
mod tests {
    use libmdbx::{Environment, Mode, NoWriteMap, WriteFlags};
    use std::{marker::PhantomData, path::Path};

    // use crate::connection::{Handler};

    use super::v8::*;
    use super::*;

    pub struct ClientHandler<F: Factory> {
        _marker: PhantomData<F>,
    }

    impl<F: Factory> ClientHandler<F> {
        pub fn new() -> Self {
            Self {
                _marker: PhantomData,
            }
        }
    }

    impl<F: Factory> super::connection::Handler for ClientHandler<F> {
        fn on_message(&self, data: &[u8], conn: &Connection) -> Result<(), Error> {
            crate::v8::handle::<F, Self>(self, data, conn)
        }
    }

    impl<F: Factory> crate::v8::Handler for ClientHandler<F> {
        fn on_logon_request(
            &self,
            _ctx: &Connection,
            _request: &impl LogonRequest,
        ) -> Result<(), Error> {
            Ok(())
        }

        fn on_encoding_request(
            &self,
            conn: &Connection,
            msg: &impl EncodingRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_encoding_response(
            &self,
            conn: &Connection,
            msg: &impl EncodingResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_logon_response(
            &self,
            conn: &Connection,
            msg: &impl LogonResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_logoff(&self, conn: &Connection, msg: &impl Logoff) -> Result<(), Error> {
            todo!()
        }

        fn on_heartbeat(&self, conn: &Connection, msg: &impl Heartbeat) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_feed_status(
            &self,
            conn: &Connection,
            msg: &impl MarketDataFeedStatus,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_feed_symbol_status(
            &self,
            conn: &Connection,
            msg: &impl MarketDataFeedSymbolStatus,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_trading_symbol_status(
            &self,
            conn: &Connection,
            msg: &impl TradingSymbolStatus,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_request(
            &self,
            conn: &Connection,
            msg: &impl MarketDataRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_depth_request(
            &self,
            conn: &Connection,
            msg: &impl MarketDepthRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_reject(
            &self,
            conn: &Connection,
            msg: &impl MarketDataReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_snapshot(
            &self,
            conn: &Connection,
            msg: &impl MarketDataSnapshot,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_depth_snapshot_level(
            &self,
            conn: &Connection,
            msg: &impl MarketDepthSnapshotLevel,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_depth_snapshot_level_float(
            &self,
            conn: &Connection,
            msg: &impl MarketDepthSnapshotLevelFloat,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_depth_update_level(
            &self,
            conn: &Connection,
            msg: &impl MarketDepthUpdateLevel,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_depth_update_level_float_with_milliseconds(
            &self,
            conn: &Connection,
            msg: &impl MarketDepthUpdateLevelFloatWithMilliseconds,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_depth_update_level_no_timestamp(
            &self,
            conn: &Connection,
            msg: &impl MarketDepthUpdateLevelNoTimestamp,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_trade_no_timestamp(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateTradeNoTimestamp,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_session_settlement(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateSessionSettlement,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_session_open(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateSessionOpen,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_session_num_trades(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateSessionNumTrades,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_trading_session_date(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateTradingSessionDate,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_depth_reject(
            &self,
            conn: &Connection,
            msg: &impl MarketDepthReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_trade(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateTrade,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_trade_with_unbundled_indicator(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateTradeWithUnbundledIndicator,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_trade_with_unbundled_indicator2(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateTradeWithUnbundledIndicator2,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_bid_ask(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateBidAsk,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_bid_ask_compact(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateBidAskCompact,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_bid_ask_float_with_microseconds(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateBidAskFloatWithMicroseconds,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_bid_ask_no_time_stamp(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateBidAskNoTimeStamp,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_trade_compact(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateTradeCompact,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_session_volume(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateSessionVolume,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_open_interest(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateOpenInterest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_session_high(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateSessionHigh,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_session_low(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateSessionLow,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_data_update_last_trade_snapshot(
            &self,
            conn: &Connection,
            msg: &impl MarketDataUpdateLastTradeSnapshot,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_orders_request(
            &self,
            conn: &Connection,
            msg: &impl MarketOrdersRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_orders_reject(
            &self,
            conn: &Connection,
            msg: &impl MarketOrdersReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_orders_add(
            &self,
            conn: &Connection,
            msg: &impl MarketOrdersAdd,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_orders_modify(
            &self,
            conn: &Connection,
            msg: &impl MarketOrdersModify,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_orders_remove(
            &self,
            conn: &Connection,
            msg: &impl MarketOrdersRemove,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_market_orders_snapshot_message_boundary(
            &self,
            conn: &Connection,
            msg: &impl MarketOrdersSnapshotMessageBoundary,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_submit_new_single_order(
            &self,
            conn: &Connection,
            msg: &impl SubmitNewSingleOrder,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_submit_flatten_position_order(
            &self,
            conn: &Connection,
            msg: &impl SubmitFlattenPositionOrder,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_cancel_replace_order(
            &self,
            conn: &Connection,
            msg: &impl CancelReplaceOrder,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_cancel_order(&self, conn: &Connection, msg: &impl CancelOrder) -> Result<(), Error> {
            todo!()
        }

        fn on_submit_new_oco_order(
            &self,
            conn: &Connection,
            msg: &impl SubmitNewOCOOrder,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_open_orders_request(
            &self,
            conn: &Connection,
            msg: &impl OpenOrdersRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_order_fills_request(
            &self,
            conn: &Connection,
            msg: &impl HistoricalOrderFillsRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_order_fills_reject(
            &self,
            conn: &Connection,
            msg: &impl HistoricalOrderFillsReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_current_positions_request(
            &self,
            conn: &Connection,
            msg: &impl CurrentPositionsRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_current_positions_reject(
            &self,
            conn: &Connection,
            msg: &impl CurrentPositionsReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_order_update(&self, conn: &Connection, msg: &impl OrderUpdate) -> Result<(), Error> {
            todo!()
        }

        fn on_open_orders_reject(
            &self,
            conn: &Connection,
            msg: &impl OpenOrdersReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_order_fill_response(
            &self,
            conn: &Connection,
            msg: &impl HistoricalOrderFillResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_position_update(
            &self,
            conn: &Connection,
            msg: &impl PositionUpdate,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_add_correcting_order_fill(
            &self,
            conn: &Connection,
            msg: &impl AddCorrectingOrderFill,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_correcting_order_fill_response(
            &self,
            conn: &Connection,
            msg: &impl CorrectingOrderFillResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_trade_accounts_request(
            &self,
            conn: &Connection,
            msg: &impl TradeAccountsRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_trade_account_response(
            &self,
            conn: &Connection,
            msg: &impl TradeAccountResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_exchange_list_request(
            &self,
            conn: &Connection,
            msg: &impl ExchangeListRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_exchange_list_response(
            &self,
            conn: &Connection,
            msg: &impl ExchangeListResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_symbols_for_exchange_request(
            &self,
            conn: &Connection,
            msg: &impl SymbolsForExchangeRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_underlying_symbols_for_exchange_request(
            &self,
            conn: &Connection,
            msg: &impl UnderlyingSymbolsForExchangeRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_symbols_for_underlying_request(
            &self,
            conn: &Connection,
            msg: &impl SymbolsForUnderlyingRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_symbol_search_request(
            &self,
            conn: &Connection,
            msg: &impl SymbolSearchRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_security_definition_for_symbol_request(
            &self,
            conn: &Connection,
            msg: &impl SecurityDefinitionForSymbolRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_security_definition_response(
            &self,
            conn: &Connection,
            msg: &impl SecurityDefinitionResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_security_definition_reject(
            &self,
            conn: &Connection,
            msg: &impl SecurityDefinitionReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_account_balance_request(
            &self,
            conn: &Connection,
            msg: &impl AccountBalanceRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_account_balance_reject(
            &self,
            conn: &Connection,
            msg: &impl AccountBalanceReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_account_balance_update(
            &self,
            conn: &Connection,
            msg: &impl AccountBalanceUpdate,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_account_balance_adjustment(
            &self,
            conn: &Connection,
            msg: &impl AccountBalanceAdjustment,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_account_balance_adjustment_reject(
            &self,
            conn: &Connection,
            msg: &impl AccountBalanceAdjustmentReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_account_balance_adjustment_complete(
            &self,
            conn: &Connection,
            msg: &impl AccountBalanceAdjustmentComplete,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_account_balances_request(
            &self,
            conn: &Connection,
            msg: &impl HistoricalAccountBalancesRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_account_balances_reject(
            &self,
            conn: &Connection,
            msg: &impl HistoricalAccountBalancesReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_account_balance_response(
            &self,
            conn: &Connection,
            msg: &impl HistoricalAccountBalanceResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_user_message(&self, conn: &Connection, msg: &impl UserMessage) -> Result<(), Error> {
            todo!()
        }

        fn on_general_log_message(
            &self,
            conn: &Connection,
            msg: &impl GeneralLogMessage,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_alert_message(
            &self,
            conn: &Connection,
            msg: &impl AlertMessage,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_journal_entry_add(
            &self,
            conn: &Connection,
            msg: &impl JournalEntryAdd,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_journal_entries_request(
            &self,
            conn: &Connection,
            msg: &impl JournalEntriesRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_journal_entries_reject(
            &self,
            conn: &Connection,
            msg: &impl JournalEntriesReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_journal_entry_response(
            &self,
            conn: &Connection,
            msg: &impl JournalEntryResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_price_data_request(
            &self,
            conn: &Connection,
            msg: &impl HistoricalPriceDataRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_price_data_response_header(
            &self,
            conn: &Connection,
            msg: &impl HistoricalPriceDataResponseHeader,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_price_data_reject(
            &self,
            conn: &Connection,
            msg: &impl HistoricalPriceDataReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_price_data_record_response(
            &self,
            conn: &Connection,
            msg: &impl HistoricalPriceDataRecordResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_price_data_tick_record_response(
            &self,
            conn: &Connection,
            msg: &impl HistoricalPriceDataTickRecordResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_price_data_response_trailer(
            &self,
            conn: &Connection,
            msg: &impl HistoricalPriceDataResponseTrailer,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_market_depth_data_request(
            &self,
            conn: &Connection,
            msg: &impl HistoricalMarketDepthDataRequest,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_market_depth_data_response_header(
            &self,
            conn: &Connection,
            msg: &impl HistoricalMarketDepthDataResponseHeader,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_market_depth_data_reject(
            &self,
            conn: &Connection,
            msg: &impl HistoricalMarketDepthDataReject,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_historical_market_depth_data_record_response(
            &self,
            conn: &Connection,
            msg: &impl HistoricalMarketDepthDataRecordResponse,
        ) -> Result<(), Error> {
            todo!()
        }

        fn on_unknown_message(&self, code: u16, data: &[u8]) -> Result<(), Error> {
            todo!()
        }
    }

    #[test]
    fn test_client() {
        std::env::set_var("RUST_LOG", "trace");
        env_logger::init();
        let client = Client::new();

        // let client = TlsClient::new();
        rt::block_on(async move {
            match client
                .start(
                    crate::sierra_chart::DENALI_1.into(),
                    ClientHandler::<crate::v8::FactoryFixed>::new(),
                )
                .await
            {
                Ok(()) => {
                    println!("success")
                }
                Err(reason) => {
                    println!("{}", reason);
                    // log::error!("connect error: {}", reason);
                    // println("error {}", reason);
                }
            }
        });
    }

    #[test]
    fn test_schema() {
        // let _vls = VLSFactory;
        let mut logon_request = LogonRequestFixed::new();
        let mut logon_request_vls = LogonRequestVLS::new();
        logon_request.set_username("iamme");
        logon_request_vls.set_username(logon_request.username());
        logon_request.copy_to(&mut logon_request_vls);

        println!(
            "{:}: {}",
            logon_request.username().len(),
            logon_request.username()
        );
        println!(
            "{:}: {}",
            logon_request_vls.username().len(),
            logon_request_vls.username()
        );
    }

    #[test]
    fn test_mdbx() {
        let env = Environment::<NoWriteMap>::new()
            .set_flags(
                Mode::ReadWrite {
                    sync_mode: libmdbx::SyncMode::SafeNoSync,
                }
                .into(),
            )
            .open(Path::new("db"))
            .unwrap();
        let txn = env.begin_rw_txn().unwrap();
        let db = txn.open_db(None).unwrap();
        txn.put(&db, b"key1", b"val11", WriteFlags::empty())
            .unwrap();
        txn.put(&db, b"key2", b"val22", WriteFlags::empty())
            .unwrap();
        txn.put(&db, b"key3", b"val33", WriteFlags::empty())
            .unwrap();
        txn.commit().unwrap();
    }
}
