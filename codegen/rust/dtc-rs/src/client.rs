use std::time::Instant;

use ntex::{
    codec::Encoder,
    connect::rustls,
};
use ntex_bytes::*;

use crate::{*, error::{ConnectError, CloseReason, Error}, connection::{Connection, Handler}, util::Framed};

pub struct Client {
    connector: ntex::connect::Connector<String>,
}

impl Client {
    pub fn new() -> Self {
        Self { connector: connect::Connector::new() }
    }

    pub async fn start<H: Handler>(
        &self,
        host: String,
        handler: H,
    ) -> Result<(), Error> {
        let (tx, rx) = mpsc::channel();
        let mut conn = Connection::new(host.clone(), Instant::now(), tx.clone());

        handler.on_connecting(&conn);

        let io = match self.connector.connect(host.clone()).await {
            Ok(io) => io,
            Err(e) => {
                handler.on_closed(
                    &conn,
                    CloseReason::ConnectError {
                        host: host.clone(),
                        reason: e.clone().into(),
                    },
                );
                return Err(Error::Connect(e.into()));
            }
        };

        conn.connected_at = Some(Instant::now());
        handler.on_connected(&conn);

        let io_ref = io.get_ref();
        rt::spawn(async move {
            while let Some(m) = rx.recv().await {
                match io_ref.encode(m, &VecCodec) {
                    Ok(_) => {}
                    Err(_e) => {
                        return;
                    }
                }
            }
        });

        let mut framed = Framed::new();
        loop {
            match io.recv(&codec::BytesCodec).await {
                Ok(Some(ref mut src)) => {
                    framed.push(src.as_ref(), |message| {
                        match handler.on_message(message, &conn) {
                            _ => {}
                        }
                    });
                }
                Ok(None) => {}
                Err(_e) => {
                    break;
                }
            }
        }
        Ok(())
    }
}

pub struct TlsClient {
    cert_store: RootCertStore,
    config: ClientConfig,
    connector: rustls::Connector<String>,
}

impl TlsClient {
    pub fn new() -> Self {
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
            .with_root_certificates(cert_store.clone())
            .with_no_client_auth();

        let connector = rustls::Connector::new(config.clone());

        Self {
            cert_store,
            config,
            connector,
        }
    }

    pub fn cert_store(&self) -> &RootCertStore {
        &self.cert_store
    }

    pub fn config(&self) -> &ClientConfig {
        &self.config
    }

    pub async fn start<H: Handler>(
        &self,
        host: String,
        handler: H,
    ) -> Result<(), Error> {
        let (tx, rx) = mpsc::channel();
        let mut conn = Connection::new(host.clone(), Instant::now(), tx.clone());

        handler.on_connecting(&conn);

        let io = match self.connector.connect(host.clone()).await {
            Ok(io) => io,
            Err(e) => {
                handler.on_closed(
                    &conn,
                    CloseReason::ConnectError {
                        host: host.clone(),
                        reason: e.clone().into(),
                    },
                );
                return Err(Error::Connect(e.into()));
            }
        };

        conn.connected_at = Some(Instant::now());
        handler.on_connected(&conn);

        let io_ref = io.get_ref();
        rt::spawn(async move {
            while let Some(m) = rx.recv().await {
                match io_ref.encode(m, &VecCodec) {
                    Ok(_) => {}
                    Err(_e) => {
                        return;
                    }
                }
            }
        });

        let mut framed = Framed::new();
        loop {
            match io.recv(&codec::BytesCodec).await {
                Ok(Some(ref mut src)) => {
                    framed.push(src.as_ref(), |message| {
                        match handler.on_message(message, &conn) {
                            _ => {}
                        }
                    });
                }
                Ok(None) => {}
                Err(_e) => {
                    break;
                }
            }
        }
        Ok(())
    }
}

/// Bytes codec.
///
/// Reads/Writes chunks of bytes from a stream.
#[derive(Debug, Copy, Clone)]
struct VecCodec;

impl Encoder for VecCodec {
    type Item = Vec<u8>;
    type Error = io::Error;

    #[inline]
    fn encode(&self, item: Vec<u8>, dst: &mut BytesMut) -> Result<(), Self::Error> {
        dst.extend_from_slice(item.as_ref());
        Ok(())
    }
}

impl Into<ConnectError> for ntex::connect::ConnectError {
    fn into(self) -> ConnectError {
        match self {
            connect::ConnectError::Resolver(e) => ConnectError::Resolver(e),
            connect::ConnectError::NoRecords => ConnectError::NoRecords,
            connect::ConnectError::InvalidInput => ConnectError::InvalidInput,
            connect::ConnectError::Unresolved => ConnectError::Unresolved,
            connect::ConnectError::Io(e) => ConnectError::Io(e)
        }
    }
}
