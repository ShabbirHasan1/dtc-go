use crate::error::Error;
use crate::v8::*;
use crate::*;

pub trait Factory: Send + 'static {
    type LogonRequest: LogonRequest;
    type LogonRequestUnsafe: LogonRequest;
    type EncodingRequest: EncodingRequest;
    type EncodingRequestUnsafe: EncodingRequest;

    fn logon_request() -> Self::LogonRequest;

    fn logon_request_parse(
        data: &[u8],
    ) -> Result<crate::message::Parsed<Self::LogonRequest, Self::LogonRequestUnsafe>, Error>;
}

pub struct FixedFactory;

impl Factory for FixedFactory {
    type LogonRequest = LogonRequestFixed;
    type LogonRequestUnsafe = LogonRequestFixedUnsafe;
    type EncodingRequest = EncodingRequestSafe;
    type EncodingRequestUnsafe = EncodingRequestUnsafe;

    fn logon_request() -> Self::LogonRequest {
        Self::LogonRequest::new()
    }

    fn logon_request_parse(
        data: &[u8],
    ) -> Result<crate::message::Parsed<Self::LogonRequest, Self::LogonRequestUnsafe>, Error>
    {
        LogonRequestFixed::parse(data)
    }
}

pub struct VLSFactory;
