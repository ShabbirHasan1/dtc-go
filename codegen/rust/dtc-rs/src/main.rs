// #![warn(dead_code)]

mod message;
mod protocol;
mod v8;

pub use message::*;
pub use protocol::*;

use std::io;

use ntex::{channel::mpsc, codec, connect, io::types::PeerAddr, rt};
use rustls::{ClientConfig, OwnedTrustAnchor, RootCertStore};

#[ntex::main]
async fn main() -> io::Result<()> {
    std::env::set_var("RUST_LOG", "trace");
    env_logger::init();

    // rustls config
    let mut cert_store = RootCertStore::empty();
    cert_store.add_server_trust_anchors(webpki_roots::TLS_SERVER_ROOTS.0.iter().map(|ta| {
        OwnedTrustAnchor::from_subject_spki_name_constraints(
            ta.subject,
            ta.spki,
            ta.name_constraints,
        )
    }));
    let config = ClientConfig::builder()
        .with_safe_defaults()
        .with_root_certificates(cert_store)
        .with_no_client_auth();

    // rustls connector
    let connector = connect::rustls::Connector::new(config.clone());

    let io = connector
        .connect("futurestrading1.sierracharts.com:443")
        .await
        .unwrap();
    println!("Connected to tls server {:?}", io.query::<PeerAddr>().get());

    // let (mut tx_write, mut rx_write) = mpsc::channel();

    rt::spawn(async move {
        // loop {
        //     let m: Option<ProtoMessage<'static, Fixed>> = rx_write.recv();
        // }
        // while let Some(m) = rx_write.recv() {
        //     let r = io
        //         .send(Bytes::from_static(b"GET /\r\n\r\n"), &codec::BytesCodec)
        //         .await?;
        // }

        // Ok(())
    });

    rt::spawn(async move {
        // loop {
        //     match io
        //         .recv(&codec::BytesCodec)
        //         .await
        //         .map_err(|e| e.into_inner())?
        //         .ok_or_else(|| io::Error::new(io::ErrorKind::Other, "disconnected"))
        //     {
        //         Ok(m) => {}
        //         Err(e) => return Err(e),
        //     }
        // }
        // Ok(())
    });

    // let result = io
    //     .send(Bytes::from_static(b"GET /\r\n\r\n"), &codec::BytesCodec)
    //     .await
    //     .map_err(Either::into_inner)?;
    // println!("Send result: {:?}", result);

    // let resp = io
    //     .recv(&codec::BytesCodec)
    //     .await
    //     .map_err(|e| e.into_inner())?
    //     .ok_or_else(|| io::Error::new(io::ErrorKind::Other, "disconnected"))?;
    // println!("Received: {:?}", resp);

    io.on_disconnect().await;
    Ok(())
    // println!("disconnecting");
    // io.shutdown().await
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
    use std::{marker::PhantomData, path::Path};

    use libmdbx::{Environment, Mode, NoWriteMap, WriteFlags};

    use super::*;
    use super::v8::*;

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
        fn on_logon_request(&self, _ctx: &Connection, _request: &impl LogonRequest) -> Result<(), Error> {
            Ok(())
        }
    }

    #[test]
    fn test_client() {
        let client = Client::new();
        rt::block_on(async move {
            match client.start(FUTURES_TRADING_1.into(), ClientHandler::<FixedFactory>::new()).await {
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
        txn.put(&db, b"key1", b"val11", WriteFlags::empty()).unwrap();
        txn.put(&db, b"key2", b"val22", WriteFlags::empty()).unwrap();
        txn.put(&db, b"key3", b"val33", WriteFlags::empty()).unwrap();
        txn.commit().unwrap();
    }
}
