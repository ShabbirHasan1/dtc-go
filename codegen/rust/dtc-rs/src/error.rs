use std::fmt::Debug;

#[derive(thiserror::Error, Clone, Debug)]
pub enum Error {
    #[error("Out of memory")]
    OutOfMemory,
    #[error("Overflow")]
    Overflow,
    #[error("Malformed: {0}")]
    Malformed(&'static str),
    #[error("Unknown message type: {0}")]
    UnknownType(u16),
    #[error("Connect error {0}")]
    Connect(ConnectError),
    #[error("Send error")]
    SendError,
}

#[derive(thiserror::Error, Debug)]
pub enum ConnectError {
    /// Failed to resolve the hostname
    #[error("Failed resolving hostname: {0}")]
    Resolver(std::io::Error),

    /// No dns records
    #[error("No dns records found for the input")]
    NoRecords,

    /// Invalid input
    #[error("Invalid input")]
    InvalidInput,

    /// Unresolved host name
    #[error("Connector received `Connect` method with unresolved host")]
    Unresolved,

    /// Connection io error
    #[error("{0}")]
    Io(#[from] std::io::Error),
}

impl Clone for ConnectError {
    fn clone(&self) -> Self {
        match self {
            ConnectError::Resolver(err) => {
                ConnectError::Resolver(std::io::Error::new(err.kind(), format!("{}", err)))
            }
            ConnectError::NoRecords => ConnectError::NoRecords,
            ConnectError::InvalidInput => ConnectError::InvalidInput,
            ConnectError::Unresolved => ConnectError::Unresolved,
            ConnectError::Io(err) => {
                ConnectError::Io(std::io::Error::new(err.kind(), format!("{}", err)))
            }
        }
    }
}

#[derive(thiserror::Error, Clone, Debug)]
pub enum CloseReason {
    #[error("Lost connection")]
    LostConnection,
    #[error("Logoff received")]
    LogoffReceived,
    #[error("Failed to connect to {host} because {reason}")]
    ConnectError { host: String, reason: ConnectError },
}
