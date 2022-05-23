use log::info;
use ntex::channel::mpsc;
use std::{cell::Cell, ops::Sub, time::Instant};

use crate::{
    error::{CloseReason, Error},
    Message,
};

pub struct Connection {
    host: String,
    pub(crate) connecting_at: Instant,
    pub(crate) connected_at: Option<Instant>,
    pub(crate) closed_at: Option<Instant>,
    pub(crate) state: Cell<State>,
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

    pub fn connected_at(&self) -> Option<Instant> {
        self.connected_at
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

pub trait Handler {
    fn on_connecting(&self, conn: &Connection) {
        info!(
            "connecting to {} since {:?}",
            conn.host(),
            conn.connecting_at()
        );
    }

    fn on_connected(&self, conn: &Connection) {
        info!(
            "connected to {} took {:?}",
            conn.host(),
            conn.connected_at().unwrap().sub(conn.connecting_at())
        );
    }

    fn on_closed(&self, conn: &Connection, reason: CloseReason) {
        info!(
            "on_closed: {:?} {}",
            conn.connecting_at(),
            reason.to_string()
        );
    }

    fn on_message(&self, data: &[u8], conn: &Connection) -> Result<(), Error>;
}
