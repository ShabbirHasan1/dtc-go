use async_trait::async_trait;

use crate::protocol::{Connection, Handler};
use crate::v8::model::*;
use crate::{Error, Parsed};

pub fn handle<F: Factory, H: HandlerV8>(
    handler: &H,
    data: &[u8],
    conn: &Connection,
) -> Result<(), Error> {
    if data.len() < 4
        || (unsafe { u16::from_le(*(data[2..].as_ptr() as *const u16)) } as usize) != data.len()
    {
        return Err(Error::Malformed(None));
    }
    let r#type = unsafe { u16::from_le(*(data[2..].as_ptr() as *const u16)) };
    match r#type {
        1 => match F::logon_request_parse(data)? {
            Parsed::Left(m) => handler.on_logon_request(conn, m),
            Parsed::Right(m) => handler.on_logon_request(conn, m),
        },

        code => handler.on_unknown_message(code, data),
    }
}

#[async_trait]
pub trait HandlerV8: Sized + Handler {
    fn on_unknown_message(&self, code: u16, _data: &[u8]) -> Result<(), Error> {
        Err(Error::UnknownType(code))
    }

    fn on_logon_request(
        &self,
        _conn: &Connection,
        _request: &impl LogonRequest,
    ) -> Result<(), Error> {
        Ok(())
    }
}
