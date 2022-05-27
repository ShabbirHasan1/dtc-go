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
    use crate::connection::Handler;
    use crate::sierra_chart::{
        DENALI_1, DENALI_3, FUTURES_TRADING_1, FUTURES_TRADING_2, FUTURES_TRADING_3,
    };
    use crate::util::Framed;
    use libmdbx::{Environment, Mode, NoWriteMap, WriteFlags};
    use log::error;
    use ntex::connect::rustls::ClientConfig;
    use std::net::ToSocketAddrs;
    use std::sync::Arc;
    use std::time::Instant;
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

    impl<F: Factory> crate::connection::Handler for ClientHandler<F> {
        fn on_connected(&self, conn: &Connection) {
            let mut request = F::EncodingRequest::new();
            request.set_encoding(EncodingEnum::BinaryEncoding);
            conn.send(request).expect("failed to send EncodingRequest");
        }

        fn on_message(&self, data: &[u8], conn: &Connection) -> Result<(), Error> {
            info!(
                "on_message() size: {} type: {}",
                unsafe { u16::from_le(*(data[0..2].as_ptr() as *const u16)) },
                unsafe { u16::from_le(*(data[2..].as_ptr() as *const u16)) }
            );
            crate::v8::handle::<F, Self>(self, data, conn)
        }
    }

    impl<F: Factory> crate::v8::Handler for ClientHandler<F> {
        fn on_encoding_response(
            &self,
            conn: &Connection,
            msg: &impl EncodingResponse,
        ) -> Result<(), Error> {
            info!("on_encoding_response: {:?}", msg);
            let mut logon_request = F::LogonRequest::new();
            logon_request.set_username("");
            logon_request.set_password("");
            logon_request.set_heartbeat_interval_in_seconds(15);
            // logon_request.set_hardware_identifier("833a19a6682f6551e6470008e83b187a");
            logon_request.set_general_text_data("sc_exchange2.datafeed");
            // logon_request.set_general_text_data("sc_futures_direct.dtc.trading");
            // logon_request.set_general_text_data("sc_realtime_exchange.datafeed");
            // logon_request.set_general_text_data("sc_delayed_exchange.datafeed");
            // logon_request.set_general_text_data("sc_futures.dtc.trading");
            // logon_request.set_general_text_data("denali");
            // logon_request.set_general_text_data("sc_futures_direct.dtc.trading");
            // logon_request.set_general_text_data("sc_futures_direct.dtc.datafeed");
            // logon_request.set_market_data_transmission_interval(1);
            // logon_request.set_general_text_data("Service=cme_ilink.trading|SymbolSettings=cme_ilink_live.trading");
            // logon_request.set_integer1(32768);
            // logon_request.set_integer1(65535);
            // logon_request.set_client_name("WebTradingPanelV2_Data");
            // logon_request.set_client_name("WebTradingPanelV2");
            // logon_request.set_trade_account("47029");
            conn.send(logon_request)
                .expect("failed to send LogonRequest");

            // let mut abr = F::AccountBalanceRequest::new();
            // abr.set_request_id(2).set_trade_account("47029");
            // conn.send(abr).expect("");

            // F::HistoricalPriceDataRequest
            // let mut r = F::HistoricalPriceDataRequest::new();
            // r.set_exchange("CME").set_symbol("6AM22_FUT_CME").set_request_id(1);//.set_symbol_id(1).set_request_action(RequestActionEnum::SnapshotWithIntervalUpdates).set_interval_for_snapshot_updates_in_milliseconds(1000);
            // conn.send(r).expect("");

            Ok(())
        }

        fn on_logon_response(
            &self,
            conn: &Connection,
            msg: &impl LogonResponse,
        ) -> Result<(), Error> {
            info!("on_logon_response: {:?}", msg);

            // let mut abr = F::AccountBalanceRequest::new();
            // abr.set_request_id(2).set_trade_account("47029");
            // conn.send(abr).expect("");
            // let mut r = F::MarketDepthRequest::new();
            // r.set_exchange("CME");
            // r.set_symbol_id(2).set_symbol("ESM22_FUT_CME").set_num_levels(20).set_request_action(RequestActionEnum::Subscribe);
            // conn.send(r).expect("");

            let mut r = F::SecurityDefinitionForSymbolRequest::new();
            r.set_request_id(10)
                .set_symbol("ESM22_FUT_CME")
                .set_exchange("CME");
            conn.send(r).expect("");

            let mut r = F::MarketDataRequest::new();
            r.set_exchange("CME")
                .set_symbol("ESM22_FUT_CME")
                .set_symbol_id(1)
                .set_request_action(RequestActionEnum::Subscribe);
            conn.send(r).expect("");
            // let mut m = F::AccountBalanceRequest::new();
            // m.set_trade_account("47029").set_request_id(1);

            // conn.send(r).expect("");
            Ok(())
        }

        fn on_unknown_message(&self, code: u16, data: &[u8]) -> Result<(), Error> {
            info!("on_unknown_message({})", code);
            Ok(())
        }
    }

    #[test]
    fn test_client() {
        std::env::set_var("RUST_LOG", "trace");
        env_logger::init();
        // let client = Client::new();
        let client = TlsClient::new();
        rt::block_on(async move {
            match client
                .start(
                    // format!("{}:11092", DENALI_1),
                    format!("{}:10041", DENALI_1),
                    // format!("{}:12097", "192.168.100.140"),
                    // format!("{}:10060", DENALI_3),
                    // format!("{}:11091", FUTURES_TRADING_1),
                    // (crate::sierra_chart::DENALI_1, 443).to_socket_addrs().unwrap().next().unwrap(),
                    ClientHandler::<crate::v8::FactoryFixed>::new(),
                )
                .await
            {
                Ok(()) => {
                    println!("success")
                }
                Err(reason) => {
                    error!("{:?}", reason);
                    // println!("{:}", reason);
                    // log::error!("connect error: {}", reason);
                    // println("error {}", reason);
                }
            }
        });
    }

    #[test]
    fn logon_message() {
        let mut logon_request = LogonRequestFixed::new();
        logon_request.set_password("");
        logon_request.set_username("");
        // logon_request.set_password("r6HRgT6qhn8");
        logon_request.set_heartbeat_interval_in_seconds(15);
        logon_request.set_hardware_identifier("833a19a6682f6551e6470008e83b187a");
        logon_request.set_general_text_data("");
        logon_request.set_integer1(4096);
        logon_request.set_client_name("WebTradingPanelV2");

        let username = logon_request.username();
        let password = logon_request.password();

        println!("password: {}", logon_request.password());
        println!("{:?}", logon_request);
    }

    fn take_ref2(msg: &impl EncodingResponse) {
        let clone = msg.clone_safe();
        println!("{:?}", msg);
        println!("{:?}", msg.clone_safe());
    }

    fn take_ref<M: EncodingResponse>(msg: &M) {
        let clone = msg.clone_safe();
        println!("{:?}", msg);
        println!("{:?}", msg.clone_safe());
    }

    #[test]
    fn test_schema() {
        std::env::set_var("RUST_LOG", "trace");
        env_logger::init();

        // sc_futures_direct.dtc.trading.
        // Service=cme_ilink.trading|SymbolSettings=cme_ilink_live.trading.

        let mut encoding_response = EncodingResponseFixed::new();
        encoding_response.set_encoding(EncodingEnum::BinaryWithVariableLengthStrings);
        encoding_response.set_protocol_version(8);

        // take_ref(&encoding_response);
        let v = encoding_response.into_vec();
        let d = v.as_slice();

        // match parse::<EncodingResponseFixed, EncodingResponseFixedUnsafe>(d) {
        //     Ok(p) => {
        //         match p {
        //             Parsed::Left(m) => {
        //                 m.clone_safe();
        //                 take_ref(m);
        //             }
        //             Parsed::Right(m) => {
        //                 take_ref(m);
        //             }
        //         }
        //     }
        //     Err(_) => {}
        // }
        EncodingResponseFixed::parse(d, |p| {
            match p {
                Parsed::Safe(m) => {
                    m.clone_safe();
                    take_ref(m);
                }
                Parsed::Unsafe(m) => {
                    take_ref(m);
                }
            }
            Ok(())
        })
        .expect("");
        // unsafe {
        //     let msg = EncodingResponseFixed::from_raw_parts(d.as_ptr(), d.len());
        //     // let msg = EncodingResponseFixed::leak(d.as_ptr(), d.len());
        //     let clone = msg.clone_safe();
        //     msg.size();
        //     std::mem::forget(msg);
        // }

        // let clone = encoding_response.clone_safe();

        // let _vls = VLSFactory;
        // let mut logon_request = LogonRequestFixed::new();
        // let mut logon_request_vls = LogonRequestVLS::new();
        // logon_request.set_username("iamme");
        // logon_request_vls.set_username(logon_request.username());
        // logon_request.copy_to(&mut logon_request_vls);
        //
        // println!("{}", &logon_request_vls);
        // info!("{}", &logon_request_vls);
        //
        // println!(
        //     "{:}: {}",
        //     logon_request.username().len(),
        //     logon_request.username()
        // );
        // println!(
        //     "{:}: {}",
        //     logon_request_vls.username().len(),
        //     logon_request_vls.username()
        // );
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
