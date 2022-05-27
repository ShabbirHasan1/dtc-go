// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const MARKET_DATA_UPDATE_SESSION_SETTLEMENT_FIXED_SIZE: usize = 24;

/// size       u16            = MarketDataUpdateSessionSettlementFixedSize  (24)
/// type       u16            = MARKET_DATA_UPDATE_SESSION_SETTLEMENT  (119)
/// symbol_id  u32            = 0
/// price      f64            = 0
/// date_time  DateTime4Byte  = 0
pub(crate) const MARKET_DATA_UPDATE_SESSION_SETTLEMENT_FIXED_DEFAULT: [u8; 24] = [
    24, 0, 119, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// Sent by the Server to the Client to update the session settlement price
/// when the session settlement price changes.
pub trait MarketDataUpdateSessionSettlement: Message {
    type Safe: MarketDataUpdateSessionSettlement;
    type Unsafe: MarketDataUpdateSessionSettlement;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// The settlement price.
    fn price(&self) -> f64;

    /// That trading date the settlement price is for. The time component is not
    /// normally considered relevant in this case.
    fn date_time(&self) -> DateTime4Byte;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// The settlement price.
    fn set_price(&mut self, value: f64) -> &mut Self;

    /// That trading date the settlement price is for. The time component is not
    /// normally considered relevant in this case.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketDataUpdateSessionSettlement) {
        to.set_symbol_id(self.symbol_id());
        to.set_price(self.price());
        to.set_date_time(self.date_time());
    }

    #[inline]
    fn parse<F: Fn(Parsed<Self::Safe, Self::Unsafe>) -> Result<(), crate::Error>>(
        data: &[u8],
        f: F,
    ) -> Result<(), crate::Error> {
        if data.len() < 4 {
            return Err(crate::Error::Malformed("need more data"));
        }
        let size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) };
        let base_size = if Self::BASE_SIZE_OFFSET == 0 {
            size
        } else {
            let base_size = unsafe {
                u16::from_le(*(data.as_ptr().offset(Self::BASE_SIZE_OFFSET) as *const u16))
            };
            if base_size > size {
                return Err(crate::Error::Malformed("base_size is greater than size"));
            }
            base_size
        };
        if (base_size as usize) >= Self::BASE_SIZE {
            let msg = unsafe { Self::Safe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Safe(&msg));
            core::mem::forget(msg);
            r
        } else {
            let msg = unsafe { Self::Unsafe::from_raw_parts(data.as_ptr(), size as usize) };
            let r = f(Parsed::Unsafe(&msg));
            core::mem::forget(msg);
            r
        }
    }
}

/// Sent by the Server to the Client to update the session settlement price
/// when the session settlement price changes.
pub struct MarketDataUpdateSessionSettlementFixed {
    data: *const MarketDataUpdateSessionSettlementFixedData,
}

pub struct MarketDataUpdateSessionSettlementFixedUnsafe {
    data: *const MarketDataUpdateSessionSettlementFixedData,
}

#[repr(packed(8), C)]
pub struct MarketDataUpdateSessionSettlementFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    price: f64,
    date_time: DateTime4Byte,
}

impl MarketDataUpdateSessionSettlementFixedData {
    pub fn new() -> Self {
        Self {
            size: 24u16.to_le(),
            r#type: MARKET_DATA_UPDATE_SESSION_SETTLEMENT.to_le(),
            symbol_id: 0u32,
            price: 0.0f64,
            date_time: 0u32,
        }
    }
}

unsafe impl Send for MarketDataUpdateSessionSettlementFixed {}
unsafe impl Send for MarketDataUpdateSessionSettlementFixedUnsafe {}
unsafe impl Send for MarketDataUpdateSessionSettlementFixedData {}

impl Drop for MarketDataUpdateSessionSettlementFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateSessionSettlementFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateSessionSettlementFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateSessionSettlementFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateSessionSettlementFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateSessionSettlementFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateSessionSettlementFixed {
    type Target = MarketDataUpdateSessionSettlementFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateSessionSettlementFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateSessionSettlementFixedUnsafe {
    type Target = MarketDataUpdateSessionSettlementFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateSessionSettlementFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for MarketDataUpdateSessionSettlementFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateSessionSettlementFixed(size: {}, type: {}, symbol_id: {}, price: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.price(), self.date_time()).as_str())
    }
}

impl core::fmt::Debug for MarketDataUpdateSessionSettlementFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateSessionSettlementFixed(size: {}, type: {}, symbol_id: {}, price: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.price(), self.date_time()).as_str())
    }
}

impl core::fmt::Display for MarketDataUpdateSessionSettlementFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateSessionSettlementFixedUnsafe(size: {}, type: {}, symbol_id: {}, price: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.price(), self.date_time()).as_str())
    }
}

impl core::fmt::Debug for MarketDataUpdateSessionSettlementFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDataUpdateSessionSettlementFixedUnsafe(size: {}, type: {}, symbol_id: {}, price: {}, date_time: {})", self.size(), self.r#type(), self.symbol_id(), self.price(), self.date_time()).as_str())
    }
}

impl crate::Message for MarketDataUpdateSessionSettlementFixed {
    type Data = MarketDataUpdateSessionSettlementFixedData;

    const TYPE: u16 = MARKET_DATA_UPDATE_SESSION_SETTLEMENT;
    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                MarketDataUpdateSessionSettlementFixedData::new(),
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
            data: data as *const MarketDataUpdateSessionSettlementFixedData,
        }
    }
}
impl crate::Message for MarketDataUpdateSessionSettlementFixedUnsafe {
    type Data = MarketDataUpdateSessionSettlementFixedData;

    const TYPE: u16 = MARKET_DATA_UPDATE_SESSION_SETTLEMENT;
    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                MarketDataUpdateSessionSettlementFixedData::new(),
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
            data: data as *const MarketDataUpdateSessionSettlementFixedData,
        }
    }
}
/// Sent by the Server to the Client to update the session settlement price
/// when the session settlement price changes.
impl MarketDataUpdateSessionSettlement for MarketDataUpdateSessionSettlementFixed {
    type Safe = MarketDataUpdateSessionSettlementFixed;
    type Unsafe = MarketDataUpdateSessionSettlementFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// The settlement price.
    fn price(&self) -> f64 {
        f64_le(self.price)
    }

    /// That trading date the settlement price is for. The time component is not
    /// normally considered relevant in this case.
    fn date_time(&self) -> DateTime4Byte {
        u32::from_le(self.date_time)
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    /// The settlement price.
    fn set_price(&mut self, value: f64) -> &mut Self {
        self.price = f64_le(value);
        self
    }

    /// That trading date the settlement price is for. The time component is not
    /// normally considered relevant in this case.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        self.date_time = value.to_le();
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

/// Sent by the Server to the Client to update the session settlement price
/// when the session settlement price changes.
impl MarketDataUpdateSessionSettlement for MarketDataUpdateSessionSettlementFixedUnsafe {
    type Safe = MarketDataUpdateSessionSettlementFixed;
    type Unsafe = MarketDataUpdateSessionSettlementFixedUnsafe;

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

    /// The settlement price.
    fn price(&self) -> f64 {
        if self.is_out_of_bounds(16) {
            0.0f64
        } else {
            f64_le(self.price)
        }
    }

    /// That trading date the settlement price is for. The time component is not
    /// normally considered relevant in this case.
    fn date_time(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(20) {
            0u32
        } else {
            u32::from_le(self.date_time)
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

    /// The settlement price.
    fn set_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.price = f64_le(value);
        }
        self
    }

    /// That trading date the settlement price is for. The time component is not
    /// normally considered relevant in this case.
    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.date_time = value.to_le();
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
                core::mem::size_of::<MarketDataUpdateSessionSettlementFixedData>(),
                "MarketDataUpdateSessionSettlementFixedData sizeof expected {:} but was {:}",
                24usize,
                core::mem::size_of::<MarketDataUpdateSessionSettlementFixedData>()
            );
            assert_eq!(
                24u16,
                MarketDataUpdateSessionSettlementFixed::new().size(),
                "MarketDataUpdateSessionSettlementFixed sizeof expected {:} but was {:}",
                24u16,
                MarketDataUpdateSessionSettlementFixed::new().size(),
            );
            assert_eq!(
                MARKET_DATA_UPDATE_SESSION_SETTLEMENT,
                MarketDataUpdateSessionSettlementFixed::new().r#type(),
                "MarketDataUpdateSessionSettlementFixed type expected {:} but was {:}",
                MARKET_DATA_UPDATE_SESSION_SETTLEMENT,
                MarketDataUpdateSessionSettlementFixed::new().r#type(),
            );
            assert_eq!(
                119u16,
                MarketDataUpdateSessionSettlementFixed::new().r#type(),
                "MarketDataUpdateSessionSettlementFixed type expected {:} but was {:}",
                119u16,
                MarketDataUpdateSessionSettlementFixed::new().r#type(),
            );
            let d = MarketDataUpdateSessionSettlementFixedData::new();
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
                (core::ptr::addr_of!(d.date_time) as usize) - p,
                "date_time offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
            );
        }
    }
}
