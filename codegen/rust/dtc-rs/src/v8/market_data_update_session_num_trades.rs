// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const MARKET_DATA_UPDATE_SESSION_NUM_TRADES_FIXED_SIZE: usize = 16;

/// size                  u16            = MarketDataUpdateSessionNumTradesFixedSize  (16)
/// type                  u16            = MARKET_DATA_UPDATE_SESSION_NUM_TRADES  (135)
/// symbol_id             u32            = 0
/// num_trades            i32            = 0
/// trading_session_date  DateTime4Byte  = 0
pub(crate) const MARKET_DATA_UPDATE_SESSION_NUM_TRADES_FIXED_DEFAULT: [u8; 16] =
    [16, 0, 135, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// Sent by the Server to the Client to update the trading session number
/// of trades.
pub trait MarketDataUpdateSessionNumTrades: Message {
    type Safe: MarketDataUpdateSessionNumTrades;
    type Unsafe: MarketDataUpdateSessionNumTrades;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// The number of trades which have occurred during the current trading session.
    /// The number of trades which have occurred during the current trading session.
    fn num_trades(&self) -> i32;

    fn trading_session_date(&self) -> DateTime4Byte;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// The number of trades which have occurred during the current trading session.
    /// The number of trades which have occurred during the current trading session.
    fn set_num_trades(&mut self, value: i32) -> &mut Self;

    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketDataUpdateSessionNumTrades) {
        to.set_symbol_id(self.symbol_id());
        to.set_num_trades(self.num_trades());
        to.set_trading_session_date(self.trading_session_date());
    }
}

/// Sent by the Server to the Client to update the trading session number
/// of trades.
pub struct MarketDataUpdateSessionNumTradesFixed {
    data: *const MarketDataUpdateSessionNumTradesFixedData,
}

pub struct MarketDataUpdateSessionNumTradesFixedUnsafe {
    data: *const MarketDataUpdateSessionNumTradesFixedData,
}

#[repr(packed(8), C)]
pub struct MarketDataUpdateSessionNumTradesFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    num_trades: i32,
    trading_session_date: DateTime4Byte,
}

impl MarketDataUpdateSessionNumTradesFixedData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: MARKET_DATA_UPDATE_SESSION_NUM_TRADES.to_le(),
            symbol_id: 0u32,
            num_trades: 0i32,
            trading_session_date: 0u32,
        }
    }
}

unsafe impl Send for MarketDataUpdateSessionNumTradesFixed {}
unsafe impl Send for MarketDataUpdateSessionNumTradesFixedUnsafe {}
unsafe impl Send for MarketDataUpdateSessionNumTradesFixedData {}

impl Drop for MarketDataUpdateSessionNumTradesFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateSessionNumTradesFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateSessionNumTradesFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateSessionNumTradesFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateSessionNumTradesFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateSessionNumTradesFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateSessionNumTradesFixed {
    type Target = MarketDataUpdateSessionNumTradesFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateSessionNumTradesFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateSessionNumTradesFixedUnsafe {
    type Target = MarketDataUpdateSessionNumTradesFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateSessionNumTradesFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketDataUpdateSessionNumTradesFixed {
    type Data = MarketDataUpdateSessionNumTradesFixedData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                MarketDataUpdateSessionNumTradesFixedData::new(),
            ),
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
            data: data as *const MarketDataUpdateSessionNumTradesFixedData,
        }
    }
}
impl crate::Message for MarketDataUpdateSessionNumTradesFixedUnsafe {
    type Data = MarketDataUpdateSessionNumTradesFixedData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                MarketDataUpdateSessionNumTradesFixedData::new(),
            ),
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
            data: data as *const MarketDataUpdateSessionNumTradesFixedData,
        }
    }
}
/// Sent by the Server to the Client to update the trading session number
/// of trades.
impl MarketDataUpdateSessionNumTrades for MarketDataUpdateSessionNumTradesFixed {
    type Safe = MarketDataUpdateSessionNumTradesFixed;
    type Unsafe = MarketDataUpdateSessionNumTradesFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// The number of trades which have occurred during the current trading session.
    /// The number of trades which have occurred during the current trading session.
    fn num_trades(&self) -> i32 {
        i32::from_le(self.num_trades)
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

    /// The number of trades which have occurred during the current trading session.
    /// The number of trades which have occurred during the current trading session.
    fn set_num_trades(&mut self, value: i32) -> &mut Self {
        self.num_trades = value.to_le();
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

/// Sent by the Server to the Client to update the trading session number
/// of trades.
impl MarketDataUpdateSessionNumTrades for MarketDataUpdateSessionNumTradesFixedUnsafe {
    type Safe = MarketDataUpdateSessionNumTradesFixed;
    type Unsafe = MarketDataUpdateSessionNumTradesFixedUnsafe;

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

    /// The number of trades which have occurred during the current trading session.
    /// The number of trades which have occurred during the current trading session.
    fn num_trades(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.num_trades)
        }
    }

    fn trading_session_date(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(16) {
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

    /// The number of trades which have occurred during the current trading session.
    /// The number of trades which have occurred during the current trading session.
    fn set_num_trades(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.num_trades = value.to_le();
        }
        self
    }

    fn set_trading_session_date(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(16) {
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
                16usize,
                core::mem::size_of::<MarketDataUpdateSessionNumTradesFixedData>(),
                "MarketDataUpdateSessionNumTradesFixedData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<MarketDataUpdateSessionNumTradesFixedData>()
            );
            assert_eq!(
                16u16,
                MarketDataUpdateSessionNumTradesFixed::new().size(),
                "MarketDataUpdateSessionNumTradesFixed sizeof expected {:} but was {:}",
                16u16,
                MarketDataUpdateSessionNumTradesFixed::new().size(),
            );
            assert_eq!(
                MARKET_DATA_UPDATE_SESSION_NUM_TRADES,
                MarketDataUpdateSessionNumTradesFixed::new().r#type(),
                "MarketDataUpdateSessionNumTradesFixed type expected {:} but was {:}",
                MARKET_DATA_UPDATE_SESSION_NUM_TRADES,
                MarketDataUpdateSessionNumTradesFixed::new().r#type(),
            );
            assert_eq!(
                135u16,
                MarketDataUpdateSessionNumTradesFixed::new().r#type(),
                "MarketDataUpdateSessionNumTradesFixed type expected {:} but was {:}",
                135u16,
                MarketDataUpdateSessionNumTradesFixed::new().r#type(),
            );
            let d = MarketDataUpdateSessionNumTradesFixedData::new();
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
                (core::ptr::addr_of!(d.num_trades) as usize) - p,
                "num_trades offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.num_trades) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.trading_session_date) as usize) - p,
                "trading_session_date offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.trading_session_date) as usize) - p,
            );
        }
    }
}