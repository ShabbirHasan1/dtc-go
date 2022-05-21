use std::{cell::Cell, time::Instant};

use ntex::{
    codec::Encoder,
    connect::{rustls::Connector, ConnectError},
};
use ntex_bytes::*;

use crate::*;

// use async_trait::async_trait;

pub struct Connection {
    host: String,
    connecting_at: Instant,
    connected_at: Option<Instant>,
    closed_at: Option<Instant>,
    state: Cell<State>,
    tx: mpsc::Sender<Vec<u8>>,
}

impl Connection {
    #[inline]
    pub fn new(host: String, connecting_at: Instant, tx: mpsc::Sender<Vec<u8>>) -> Self {
        Self {
            host,
            connecting_at,
            connected_at: None,
            closed_at: None,
            state: Cell::new(State::Init),
            tx,
        }
    }

    #[inline]
    pub fn send<T: Message>(&self, message: T) -> Result<(), Error> {
        match self.tx.send(message.into()) {
            Ok(_) => Ok(()),
            Err(_) => Err(Error::SendError),
        }
    }

    pub fn host(&self) -> &str {
        self.host.as_ref()
    }

    pub fn connecting_at(&self) -> Instant {
        self.connecting_at
    }

    pub fn closed_at(&self) -> Option<Instant> {
        self.closed_at
    }

    pub fn state(&self) -> State {
        self.state.get()
    }
}

#[derive(Clone, Copy, Debug)]
pub enum State {
    Init,
    SentLogonRequest,
    Error,
}

#[derive(Clone, Debug)]
pub enum CloseReason {
    LostConnection,
    LogoffReceived,
    ConnectError { host: String, reason: ConnectError },
}

pub trait Handler {
    fn on_connecting(&self, _conn: &Connection) {}

    fn on_connected(&self, _conn: &Connection) {}

    fn on_closed(&self, _conn: &Connection, _reason: CloseReason) {}

    fn on_message(&self, data: &[u8], conn: &Connection) -> Result<(), Error>;
}


pub struct Framed {
    buf: Vec<u8>,
}

impl Framed {
    pub fn new() -> Self {
        Self { buf: Vec::new() }
    }

    pub fn push<T: Fn(&[u8])>(&mut self, src: &[u8], f: T) {
        let mut data = src.as_ref();
        if data.len() == 0 {
            return;
        }
        // Try
        if self.buf.is_empty() {
            match self.process(src, f) {
                Ok(Some(idx)) => {
                    data = &data[idx..];
                    if data.len() > 0 {
                        self.buf.extend_from_slice(data);
                    }
                }
                Ok(None) => {}
                Err(_e) => {}
            }
        } else {
            let size = if self.buf.len() > 1 {
                (unsafe { u16::from_le(*(self.buf.as_ptr() as *const u16)) }) as usize
            } else {
                let b: [u8; 2] = [self.buf.as_slice()[0], src.as_ref()[0]];
                (unsafe { u16::from_le(*(b.as_ptr() as *const u16)) }) as usize
            };

            // Is next message received?
            if (self.buf.len() + src.len()) < size {
                self.buf.extend_from_slice(src.as_ref());
                return;
            }

            let split_at = size - self.buf.len();
            self.buf.extend_from_slice(&data[0..split_at]);
            data = &data[split_at..];
            f(self.buf.as_slice());
            self.buf.clear();

            match self.process(data, f) {
                Ok(Some(idx)) => {
                    data = &data[idx..];
                    if data.len() > 0 {
                        self.buf.extend_from_slice(data);
                    }
                }
                Ok(None) => {}
                Err(_e) => {}
            }

            if data.len() < 4 {
                self.buf.extend_from_slice(data);
            }
        }
    }

    fn process<T: Fn(&[u8])>(&mut self, mut data: &[u8], f: T) -> Result<Option<usize>, Error> {
        let mut size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) } as usize;
        // Bad message?
        if size < 4 {
            return Err(Error::Malformed("bad size".into()));
        }
        if size > data.len() {
            // buf.extend_from_slice(data);
            return Ok(Some(0usize));
        }

        if size == data.len() {
            f(&data[..size]);
            Ok(None)
        } else {
            let mut offset = 0;
            loop {
                f(&data[..size]);

                offset += size;
                data = &data[size..];

                if data.len() == 0 {
                    return Ok(None);
                } else if data.len() < 2 {
                    return Ok(Some(offset));
                }
                size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) } as usize;
                if size > data.len() {
                    return Ok(Some(offset));
                }
            }
        }
    }
}


pub struct Client {
    _cert_store: RootCertStore,
    _config: ClientConfig,
    connector: Connector<String>,
}

pub const FUTURES_TRADING_1: &'static str = "futurestrading1.sierracharts.com:443";
pub const FUTURES_TRADING_2: &'static str = "futurestrading2.sierracharts.com:443";
pub const FUTURES_TRADING_3: &'static str = "futurestrading3.sierracharts.com:443";
pub const FUTURES_TRADING_4: &'static str = "futurestrading4.sierracharts.com:443";
pub const FUTURES_TRADING_5: &'static str = "futurestrading5.sierracharts.com:443";
pub const FUTURES_TRADING_6: &'static str = "futurestrading6.sierracharts.com:443";

impl Client {
    pub fn new() -> Client {
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

        let connector = connect::rustls::Connector::new(config.clone());

        Self {
            _cert_store: cert_store,
            _config: config,
            connector,
        }
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
                        reason: e.clone(),
                    },
                );
                return Err(Error::Connect(e));
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
