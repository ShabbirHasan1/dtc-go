#![allow(dead_code)]

mod client;
mod connection;
mod error;
mod message;
mod sierra_chart;
mod util;
mod v8;

pub use client::*;
use log::info;
pub use message::*;

use std::{io, time::Duration};

use ntex::{channel::mpsc, codec, connect, rt};
use rustls::{ClientConfig, OwnedTrustAnchor, RootCertStore};

use clap::Parser;

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

    use crate::connection::{Connection, Handler};
    use crate::error::Error;

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

    impl<F: Factory> Handler for ClientHandler<F> {
        fn on_message(&self, data: &[u8], conn: &Connection) -> Result<(), Error> {
            crate::v8::handle::<F, Self>(self, data, conn)
        }
    }

    impl<F: Factory> crate::v8::HandlerV8 for ClientHandler<F> {
        fn on_logon_request(
            &self,
            _ctx: &Connection,
            _request: &impl LogonRequest,
        ) -> Result<(), Error> {
            Ok(())
        }
    }

    #[test]
    fn test_client() {
        let client = TlsClient::new();
        rt::block_on(async move {
            match client
                .start(
                    crate::sierra_chart::FUTURES_TRADING_1.into(),
                    ClientHandler::<FixedFactory>::new(),
                )
                .await
            {
                Ok(()) => {
                    println!("success")
                }
                Err(_reason) => {
                    println!("error");
                    // println("error {}", reason);
                }
            }
        });
    }

    #[test]
    fn test_schema() {
        let _vls = VLSFactory;
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
