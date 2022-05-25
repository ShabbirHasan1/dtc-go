// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const MARKET_DATA_UPDATE_SESSION_HIGH_FIXED_SIZE: usize = 24;

/// size                  u16            = MarketDataUpdateSessionHighFixedSize  (24)
/// type                  u16            = MARKET_DATA_UPDATE_SESSION_HIGH  (114)
/// symbol_id             u32            = 0
/// price                 f64            = 0
/// trading_session_date  DateTime4Byte  = 0
pub(crate) const MARKET_DATA_UPDATE_SESSION_HIGH_FIXED_DEFAULT: [u8; 24] = [
    24, 0, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// Sent by the Server to the Client to update the session High as the High
/// price changes throughout the session.
pub trait MarketDataUpdateSessionHigh: Message {
    type Safe: MarketDataUpdateSessionHigh;
    type Unsafe: MarketDataUpdateSessionHigh;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// The session High price.
    fn price(&self) -> f64;

    fn trading_session_date(&self) -> DateTime4Byte;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// The session High price.
    fn set_price(&mut self, value: f64) -> &mut Self;

    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketDataUpdateSessionHigh) {
        to.set_symbol_id(self.symbol_id());
        to.set_price(self.price());
        to.set_trading_session_date(self.trading_session_date());
    }
}

/// Sent by the Server to the Client to update the session High as the High
/// price changes throughout the session.
pub struct MarketDataUpdateSessionHighFixed {
    data: *const MarketDataUpdateSessionHighFixedData,
}

pub struct MarketDataUpdateSessionHighFixedUnsafe {
    data: *const MarketDataUpdateSessionHighFixedData,
}

#[repr(packed(8), C)]
pub struct MarketDataUpdateSessionHighFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    price: f64,
    trading_session_date: DateTime4Byte,
}

impl MarketDataUpdateSessionHighFixedData {
    pub fn new() -> Self {
        Self {
            size: 24u16.to_le(),
            r#type: MARKET_DATA_UPDATE_SESSION_HIGH.to_le(),
            symbol_id: 0u32,
            price: 0.0f64,
            trading_session_date: 0u32,
        }
    }
}

unsafe impl Send for MarketDataUpdateSessionHighFixed {}
unsafe impl Send for MarketDataUpdateSessionHighFixedUnsafe {}
unsafe impl Send for MarketDataUpdateSessionHighFixedData {}

impl Drop for MarketDataUpdateSessionHighFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateSessionHighFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateSessionHighFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateSessionHighFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateSessionHighFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateSessionHighFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateSessionHighFixed {
    type Target = MarketDataUpdateSessionHighFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateSessionHighFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateSessionHighFixedUnsafe {
    type Target = MarketDataUpdateSessionHighFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateSessionHighFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketDataUpdateSessionHighFixed {
    type Data = MarketDataUpdateSessionHighFixedData;

    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateSessionHighFixedData::new()),
        }
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const MarketDataUpdateSessionHighFixedData,
        }
    }
}
impl crate::Message for MarketDataUpdateSessionHighFixedUnsafe {
    type Data = MarketDataUpdateSessionHighFixedData;

    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDataUpdateSessionHighFixedData::new()),
        }
    }

    #[inline]
    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }

    #[inline]
    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        u16::from_le(self.size)
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {
        Self {
            data: data as *const MarketDataUpdateSessionHighFixedData,
        }
    }
}
/// Sent by the Server to the Client to update the session High as the High
/// price changes throughout the session.
impl MarketDataUpdateSessionHigh for MarketDataUpdateSessionHighFixed {
    type Safe = MarketDataUpdateSessionHighFixed;
    type Unsafe = MarketDataUpdateSessionHighFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// The session High price.
    fn price(&self) -> f64 {
        f64_le(self.price)
    }

    fn trading_session_date(&self) -> DateTime4Byte {
        u32::from_le(self.trading_session_date)
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    /// The session High price.
    fn set_price(&mut self, value: f64) -> &mut Self {
        self.price = f64_le(value);
        self
    }

    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self {
        self.trading_session_date = value.to_le();
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }
}

/// Sent by the Server to the Client to update the session High as the High
/// price changes throughout the session.
impl MarketDataUpdateSessionHigh for MarketDataUpdateSessionHighFixedUnsafe {
    type Safe = MarketDataUpdateSessionHighFixed;
    type Unsafe = MarketDataUpdateSessionHighFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    /// The session High price.
    fn price(&self) -> f64 {
        if self.is_out_of_bounds(16) {
            0.0f64
        } else {
            f64_le(self.price)
        }
    }

    fn trading_session_date(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(20) {
            0u32
        } else {
            u32::from_le(self.trading_session_date)
        }
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.symbol_id = value.to_le();
        }
        self
    }

    /// The session High price.
    fn set_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.price = f64_le(value);
        }
        self
    }

    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.trading_session_date = value.to_le();
        }
        self
    }

    #[inline]
    fn clone_safe(&self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }
}

#[cfg(test)]
pub(crate) mod tests {
    use super::*;

    #[test]
    pub(crate) fn layout() {
        unsafe {
            assert_eq!(
                24usize,
                core::mem::size_of::<MarketDataUpdateSessionHighFixedData>(),
                "MarketDataUpdateSessionHighFixedData sizeof expected {:} but was {:}",
                24usize,
                core::mem::size_of::<MarketDataUpdateSessionHighFixedData>()
            );
            assert_eq!(
                24u16,
                MarketDataUpdateSessionHighFixed::new().size(),
                "MarketDataUpdateSessionHighFixed sizeof expected {:} but was {:}",
                24u16,
                MarketDataUpdateSessionHighFixed::new().size(),
            );
            assert_eq!(
                MARKET_DATA_UPDATE_SESSION_HIGH,
                MarketDataUpdateSessionHighFixed::new().r#type(),
                "MarketDataUpdateSessionHighFixed type expected {:} but was {:}",
                MARKET_DATA_UPDATE_SESSION_HIGH,
                MarketDataUpdateSessionHighFixed::new().r#type(),
            );
            assert_eq!(
                114u16,
                MarketDataUpdateSessionHighFixed::new().r#type(),
                "MarketDataUpdateSessionHighFixed type expected {:} but was {:}",
                114u16,
                MarketDataUpdateSessionHighFixed::new().r#type(),
            );
            let d = MarketDataUpdateSessionHighFixedData::new();
            let p = (&d as *const _ as *const u8).offset(0) as usize;
            assert_eq!(
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
                "size offset expected {:} but was {:}",
                0usize,
                (core::ptr::addr_of!(d.size) as usize) - p,
            );
            assert_eq!(
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
                "type offset expected {:} but was {:}",
                2usize,
                (core::ptr::addr_of!(d.r#type) as usize) - p,
            );
            assert_eq!(
                4usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
                "symbol_id offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.price) as usize) - p,
                "price offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.price) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.trading_session_date) as usize) - p,
                "trading_session_date offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.trading_session_date) as usize) - p,
            );
        }
    }
}