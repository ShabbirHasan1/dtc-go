use std::sync::Arc;
use std::time::Instant;

use ntex::codec::Encoder;
use ntex::connect::rustls::{ClientConfig, Connector};
use ntex_bytes::*;
use rustls::client::{ServerCertVerified, ServerCertVerifier};
use rustls::{Certificate, OwnedTrustAnchor, RootCertStore, ServerName};

use crate::{
    connection::{Connection, Handler},
    error::{CloseReason, ConnectError, Error},
    util::Framed,
    *,
};

pub struct Client {
    connector: ntex::connect::Connector<String>,
}

impl Client {
    pub fn new() -> Self {
        Self {
            connector: connect::Connector::new(),
        }
    }

    pub async fn start<H: Handler>(&self, host: String, handler: H) -> Result<(), Error> {
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
                    match framed.push(src.as_ref(), |message| {
                        match handler.on_message(message, &conn) {
                            _ => {}
                        }
                    }) {
                        Ok(_) => {}
                        Err(e) => {
                            log::error!("{:?}", e);
                        }
                    }
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
    // cert_store: RootCertStore,
    // config: ClientConfig,
    connector: Connector<String>,
    // connector: connect::openssl::Connector<String>,
}

struct CertVerifier;

impl ServerCertVerifier for CertVerifier {
    fn verify_server_cert(
        &self,
        end_entity: &Certificate,
        intermediates: &[Certificate],
        server_name: &ServerName,
        scts: &mut dyn Iterator<Item = &[u8]>,
        ocsp_response: &[u8],
        now: std::time::SystemTime,
    ) -> Result<ServerCertVerified, rustls::Error> {
        // let (cert, chain, trustroots) = prepare(end_entity, intermediates, &self.roots)?;
        // let webpki_now = webpki::Time::try_from(now).map_err(|_| Error::FailedToGetCurrentTime)?;
        //
        // let dns_name = match server_name {
        //     &ServerName::DnsName(ref n) => n,
        //     ServerName::IpAddress(_) => {
        //         return Err(Error::UnsupportedNameType);
        //     }
        // };
        //
        // let cert = cert
        //     .verify_is_valid_tls_server_cert(
        //         SUPPORTED_SIG_ALGS,
        //         &webpki::TlsServerTrustAnchors(&trustroots),
        //         &chain,
        //         webpki_now,
        //     )
        //     .map_err(pki_error)
        //     .map(|_| cert)?;
        //
        // if let Some(policy) = &self.ct_policy {
        //     policy.verify(end_entity, now, scts)?;
        // }
        //
        // if !ocsp_response.is_empty() {
        //     log::trace!("Unvalidated OCSP response: {:?}", ocsp_response.to_vec());
        // }
        //
        // cert.verify_is_valid_for_dns_name(dns_name.0.as_ref())
        //     .map_err(pki_error)
        //     .map(|_| ServerCertVerified::assertion())

        Ok(ServerCertVerified::assertion())
    }
}

impl TlsClient {
    pub fn new() -> Self {
        // load ssl keys
        // let mut builder = SslConnector::builder(SslMethod::tls()).unwrap();
        // // builder
        // //     .set_private_key_file("./examples/key.pem", SslFiletype::PEM)
        // //     .unwrap();
        // // builder
        // //     .set_certificate_chain_file("./examples/cert.pem")
        // //     .unwrap();
        // builder.set_verify(SslVerifyMode::NONE);

        let mut cert_store = RootCertStore::empty();
        cert_store.add_server_trust_anchors(webpki_roots::TLS_SERVER_ROOTS.0.iter().map(|ta| {
            OwnedTrustAnchor::from_subject_spki_name_constraints(
                ta.subject,
                ta.spki,
                ta.name_constraints,
            )
        }));
        // cert_store.add_server_trust_anchors(&webpki_roots::TLS_SERVER_ROOTS);

        let config = ClientConfig::builder()
            .with_safe_defaults()
            .with_custom_certificate_verifier(Arc::new(CertVerifier))
            // .with_root_certificates(cert_store)
            .with_no_client_auth();

        // let mut cert_store = RootCertStore::empty();
        // cert_store.add_server_trust_anchors(webpki_roots::TLS_SERVER_ROOTS.0.iter().map(|ta| {
        //     OwnedTrustAnchor::from_subject_spki_name_constraints(
        //         ta.subject,
        //         ta.spki,
        //         ta.name_constraints,
        //     )
        // }));
        //
        // ClientConfig::builder()
        //     .with_safe_defaults()
        //     .with_root_certificates(cert_store)
        //     .with_no_client_auth();

        let connector = Connector::new(config);
        // let connector = connect::openssl::Connector::new(builder.build());

        Self {
            // cert_store: cert_store,
            // config,
            connector,
        }
    }

    // pub fn cert_store(&self) -> &RootCertStore {
    //     &self.cert_store
    // }

    // pub fn config(&self) -> &ClientConfig {
    //     &self.config
    // }

    pub async fn start<H: Handler>(&self, host: String, handler: H) -> Result<(), Error> {
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
                log::error!("{:?}", e);
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
                    match framed.push(src.as_ref(), |message| {
                        match handler.on_message(message, &conn) {
                            _ => {}
                        }
                    }) {
                        Ok(_) => {}
                        Err(e) => {
                            log::error!("{:?}", e);
                            break;
                        }
                    }
                }
                Ok(None) => {}
                Err(e) => {
                    log::error!("{:?}", e);
                    break;
                }
            }
        }
        Ok(())
    }
}

/// Vec codec.
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
            connect::ConnectError::Io(e) => ConnectError::Io(e),
        }
    }
}
