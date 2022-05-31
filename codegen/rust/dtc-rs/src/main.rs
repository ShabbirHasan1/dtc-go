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
    use libmdbx::{Environment, Mode, NoWriteMap, WriteFlags};
    use log::error;
    use std::{marker::PhantomData, path::Path};
    use chrono::Utc;

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
            // info!(
            //     "on_message() size: {} type: {}",
            //     unsafe { u16::from_le(*(data[0..2].as_ptr() as *const u16)) },
            //     unsafe { u16::from_le(*(data[2..].as_ptr() as *const u16)) }
            // );
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
            // logon_request.set_general_text_data("sc_exchange2.datafeed");
            // logon_request.set_general_text_data("{D9423186-40FD-423E-A2DA-6D5E5C9C20F9}");
            logon_request.set_general_text_data("{D9423186-40FD-423E-A2DA-6D5E5C9C20F9}");
            logon_request.set_hardware_identifier("1931446566");
            // logon_request.set_general_text_data("sc_futures_direct.dtc.trading");
            // logon_request.set_general_text_data("{D9423186-40FD-423E-A2DA-6D5E5C9C20F9}");

            // logon_request.set_general_text_data("Service=sc_futures_direct.dtc.trading|SymbolSettings=sc_futures_direct.dtc.trading");
            // logon_request.set_general_text_data("sc_realtime_exchange.datafeed");
            // logon_request.set_general_text_data("sc_delayed_exchange.datafeed");
            // logon_request.set_general_text_data("sc_futures.dtc.trading");
            // logon_request.set_general_text_data("denali");
            // logon_request.set_general_text_data("sc_futures_direct.dtc.trading");
            // logon_request.set_general_text_data("sc_futures_direct.dtc.datafeed");
            // logon_request.set_market_data_transmission_interval(1);
            // logon_request.set_general_text_data("Service=cme_ilink.trading|SymbolSettings=cme_ilink_live.trading");
            // logon_request.set_integer1(4096);
            // logon_request.set_integer1(65535);
            // logon_request.set_client_name("Sierra Chart");
            /*

             */
            // logon_request.set_client_name("WebTradingPanelV2");
            // logon_request.set_trade_account("47029");
            // logon_request.set_hardware_identifier("7fcd786fa28967227b674f7e87217d8b");
            // logon_request.set_hardware_identifier("e82748b74f607de166723062d97e2b43");
            logon_request.set_client_name("Sierra Chart");
            conn.send(logon_request)
                .expect("failed to send LogonRequest");

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

            // let mut r = F::HistoricalMarketDepthDataRequest::new();
            // r.set_exchange("CME")
            //     .set_symbol("ESM22_FUT_CME")
            //     .set_request_id(2)
            //     .set_use_z_lib_compression(false)
            //     .set_start_date_time(Utc::now().timestamp_nanos() / 1000 - (1000000*60*60*24));
            // conn.send(r).expect("");

            // let mut r = F::HistoricalPriceDataRequest::new();
            // r.set_exchange("CME")
            //     .set_request_id(1)
            //     .set_symbol("ESM22_FUT_CME")
            //     .set_record_interval(HistoricalDataIntervalEnum::Interval5Minute)
            //     .set_max_days_to_return(10)
            //     .set_use_z_lib_compression(false);
            // conn.send(r).expect("");

            let mut r = F::SecurityDefinitionForSymbolRequest::new();
            r.set_request_id(10)
                .set_symbol("ESM22_FUT_CME")
                .set_exchange("CME");
            conn.send(r).expect("");

            let mut r = F::MarketDataRequest::new();
            r.set_exchange("CME")
                .set_symbol("ESM22_FUT_CME")
                .set_symbol_id(10)
                .set_request_action(RequestActionEnum::Subscribe);
            conn.send(r).expect("");
            // let mut m = F::AccountBalanceRequest::new();
            // m.set_trade_account("47029").set_request_id(1);

            // conn.send(r).expect("");
            Ok(())
        }


        fn on_market_data_snapshot(&self, conn: &Connection, msg: &impl MarketDataSnapshot) -> Result<(), Error> {
            println!("snapshot: bid:{:?} ask: {:?}", msg.bid_price() / 100.0, msg.ask_price() / 100.0);
            Ok(())
        }

        fn on_market_data_update_bid_ask(&self, conn: &Connection, msg: &impl MarketDataUpdateBidAsk) -> Result<(), Error> {
            println!("bid:{:?} ask:{:?}", msg.bid_price(), msg.ask_price());
            Ok(())
        }

        fn on_market_data_update_bid_ask_compact(&self, conn: &Connection, msg: &impl MarketDataUpdateBidAskCompact) -> Result<(), Error> {
            println!("bid:{:?} ask:{:?}", msg.bid_price(), msg.ask_price());
            Ok(())
        }

        // fn on_market_data_update_trade_compact(&self, conn: &Connection, msg: &impl MarketDataUpdateTradeCompact) -> Result<(), Error> {
        //     println!("trade price:{:?} qty:{:?} aggressor:{:?}", msg.price() / 100.0, msg.volume(), msg.at_bid_or_ask());
        //     Ok(())
        // }


        fn on_unknown_message(&self, code: u16, data: &[u8]) -> Result<(), Error> {
            info!("on_unknown_message({})", code);
            Ok(())
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
                    // format!("{}:11092", DENALI_1),
                    // format!("{}:10041", DENALI_1),
                    format!("{}:11099", "127.0.0.1"),
                    // format!("{}:11098", "127.0.0.1"),
                    // format!("{}:10060", DENALI_3),
                    // format!("{}:11091", FUTURES_TRADING_2),
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
        let mut logon_request = LogonRequestVLS::new();
        logon_request.set_username("kamaiu");
        logon_request.set_password("r323fsafqhn8");
        let passw = logon_request.password();
        // logon_request.set_password("");
        logon_request.set_heartbeat_interval_in_seconds(15);
        logon_request.set_hardware_identifier("833a19a6682f6551e6470008e83b187a");
        let passw2 = logon_request.password();
        // logon_request.set_general_text_data("");
        logon_request.set_integer1(4096);
        logon_request.set_client_name("Sierra Chart");
        let pass3 = logon_request.password();

        let data = logon_request.as_slice();
        // let data = data.to_vec();

        let pass = unsafe { core::str::from_utf8_unchecked(&data[63..(63+12)]) };


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
    fn test_it() {
        /*
        EÉÁú@Äº+ZTèÁ-À_~Pÿ ¡88?
L'a?³hs!
kamaiur6HRgT6q$hn8{D9423186-40FD-423E-A2DA-6D5E5C9C20F9}7fcd786fa28967227b674f7e87217d8cSierra Chart
         */
        let data: [u8; 205] = [2,0,0,0,69,2,0,201,193,250,64,0,128,6,0,0,127,0,0,1,127,0,0,1,196,186,43,90,84,232,193,45,137,192,95,126,80,24,159,255,32,128,0,0,161,0,1,0,56,0,0,0,8,0,0,0,56,0,7,0,63,0,13,0,76,0,39,0,0,0,0,0,97,63,179,104,20,0,0,0,0,0,0,0,0,0,0,0,115,0,33,0,148,0,13,0,0,0,0,0,107,97,109,97,105,117,0,114,54,72,82,103,84,54,113,36,104,110,56,0,123,68,57,52,50,51,49,56,54,45,52,48,70,68,45,52,50,51,69,45,65,50,68,65,45,54,68,53,69,53,67,57,67,50,48,70,57,125,0,55,102,99,100,55,56,54,102,97,50,56,57,54,55,50,50,55,98,54,55,52,102,55,101,56,55,50,49,55,100,56,99,0,83,105,101,114,114,97,32,67,104,97,114,116,0];
        match LogonRequestFixed::parse(&data, |msg| {
            match msg {
                Parsed::Safe(m) => {
                    println!("{:?}", m);
                }
                Parsed::Unsafe(m) => {
                    println!("{:?}", m);
                }
            }

            Ok(())
        }) {
            Ok(_) => {}
            Err(_) => {}
        }
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
