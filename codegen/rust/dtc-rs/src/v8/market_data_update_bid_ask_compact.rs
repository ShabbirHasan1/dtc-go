// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const MARKET_DATA_UPDATE_BID_ASK_COMPACT_FIXED_SIZE: usize = 28;

/// size          u16            = MarketDataUpdateBidAskCompactFixedSize  (28)
/// type          u16            = MARKET_DATA_UPDATE_BID_ASK_COMPACT  (117)
/// bid_price     f32            = f32::MAX
/// bid_quantity  f32            = 0
/// ask_price     f32            = f32::MAX
/// ask_quantity  f32            = 0
/// date_time     DateTime4Byte  = 0
/// symbol_id     u32            = 0
pub(crate) const MARKET_DATA_UPDATE_BID_ASK_COMPACT_FIXED_DEFAULT: [u8; 28] = [
    28, 0, 117, 0, 255, 255, 127, 127, 0, 0, 0, 0, 255, 255, 127, 127, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0,
];

pub trait MarketDataUpdateBidAskCompact: Message {
    type Safe: MarketDataUpdateBidAskCompact;
    type Unsafe: MarketDataUpdateBidAskCompact;

    fn bid_price(&self) -> f32;

    fn bid_quantity(&self) -> f32;

    fn ask_price(&self) -> f32;

    fn ask_quantity(&self) -> f32;

    fn date_time(&self) -> DateTime4Byte;

    fn symbol_id(&self) -> u32;

    fn set_bid_price(&mut self, value: f32) -> &mut Self;

    fn set_bid_quantity(&mut self, value: f32) -> &mut Self;

    fn set_ask_price(&mut self, value: f32) -> &mut Self;

    fn set_ask_quantity(&mut self, value: f32) -> &mut Self;

    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self;

    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketDataUpdateBidAskCompact) {
        to.set_bid_price(self.bid_price());
        to.set_bid_quantity(self.bid_quantity());
        to.set_ask_price(self.ask_price());
        to.set_ask_quantity(self.ask_quantity());
        to.set_date_time(self.date_time());
        to.set_symbol_id(self.symbol_id());
    }
}

pub struct MarketDataUpdateBidAskCompactFixed {
    data: *const MarketDataUpdateBidAskCompactFixedData,
}

pub struct MarketDataUpdateBidAskCompactFixedUnsafe {
    data: *const MarketDataUpdateBidAskCompactFixedData,
}

#[repr(packed(8), C)]
pub struct MarketDataUpdateBidAskCompactFixedData {
    size: u16,
    r#type: u16,
    bid_price: f32,
    bid_quantity: f32,
    ask_price: f32,
    ask_quantity: f32,
    date_time: DateTime4Byte,
    symbol_id: u32,
}

impl MarketDataUpdateBidAskCompactFixedData {
    pub fn new() -> Self {
        Self {
            size: 28u16.to_le(),
            r#type: MARKET_DATA_UPDATE_BID_ASK_COMPACT.to_le(),
            bid_price: f32_le(f32::MAX),
            bid_quantity: 0.0f32,
            ask_price: f32_le(f32::MAX),
            ask_quantity: 0.0f32,
            date_time: 0u32,
            symbol_id: 0u32,
        }
    }
}

unsafe impl Send for MarketDataUpdateBidAskCompactFixed {}
unsafe impl Send for MarketDataUpdateBidAskCompactFixedUnsafe {}
unsafe impl Send for MarketDataUpdateBidAskCompactFixedData {}

impl Drop for MarketDataUpdateBidAskCompactFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDataUpdateBidAskCompactFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDataUpdateBidAskCompactFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDataUpdateBidAskCompactFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDataUpdateBidAskCompactFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDataUpdateBidAskCompactFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDataUpdateBidAskCompactFixed {
    type Target = MarketDataUpdateBidAskCompactFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateBidAskCompactFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDataUpdateBidAskCompactFixedUnsafe {
    type Target = MarketDataUpdateBidAskCompactFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDataUpdateBidAskCompactFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketDataUpdateBidAskCompactFixed {
    type Data = MarketDataUpdateBidAskCompactFixedData;

    const BASE_SIZE: usize = 28;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                MarketDataUpdateBidAskCompactFixedData::new(),
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
            data: data as *const MarketDataUpdateBidAskCompactFixedData,
        }
    }
}
impl crate::Message for MarketDataUpdateBidAskCompactFixedUnsafe {
    type Data = MarketDataUpdateBidAskCompactFixedData;

    const BASE_SIZE: usize = 28;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                MarketDataUpdateBidAskCompactFixedData::new(),
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
            data: data as *const MarketDataUpdateBidAskCompactFixedData,
        }
    }
}
impl MarketDataUpdateBidAskCompact for MarketDataUpdateBidAskCompactFixed {
    type Safe = MarketDataUpdateBidAskCompactFixed;
    type Unsafe = MarketDataUpdateBidAskCompactFixedUnsafe;

    fn bid_price(&self) -> f32 {
        f32_le(self.bid_price)
    }

    fn bid_quantity(&self) -> f32 {
        f32_le(self.bid_quantity)
    }

    fn ask_price(&self) -> f32 {
        f32_le(self.ask_price)
    }

    fn ask_quantity(&self) -> f32 {
        f32_le(self.ask_quantity)
    }

    fn date_time(&self) -> DateTime4Byte {
        u32::from_le(self.date_time)
    }

    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    fn set_bid_price(&mut self, value: f32) -> &mut Self {
        self.bid_price = f32_le(value);
        self
    }

    fn set_bid_quantity(&mut self, value: f32) -> &mut Self {
        self.bid_quantity = f32_le(value);
        self
    }

    fn set_ask_price(&mut self, value: f32) -> &mut Self {
        self.ask_price = f32_le(value);
        self
    }

    fn set_ask_quantity(&mut self, value: f32) -> &mut Self {
        self.ask_quantity = f32_le(value);
        self
    }

    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        self.date_time = value.to_le();
        self
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
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

impl MarketDataUpdateBidAskCompact for MarketDataUpdateBidAskCompactFixedUnsafe {
    type Safe = MarketDataUpdateBidAskCompactFixed;
    type Unsafe = MarketDataUpdateBidAskCompactFixedUnsafe;

    fn bid_price(&self) -> f32 {
        if self.is_out_of_bounds(8) {
            f32_le(f32::MAX)
        } else {
            f32_le(self.bid_price)
        }
    }

    fn bid_quantity(&self) -> f32 {
        if self.is_out_of_bounds(12) {
            0.0f32
        } else {
            f32_le(self.bid_quantity)
        }
    }

    fn ask_price(&self) -> f32 {
        if self.is_out_of_bounds(16) {
            f32_le(f32::MAX)
        } else {
            f32_le(self.ask_price)
        }
    }

    fn ask_quantity(&self) -> f32 {
        if self.is_out_of_bounds(20) {
            0.0f32
        } else {
            f32_le(self.ask_quantity)
        }
    }

    fn date_time(&self) -> DateTime4Byte {
        if self.is_out_of_bounds(24) {
            0u32
        } else {
            u32::from_le(self.date_time)
        }
    }

    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(28) {
            0u32
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    fn set_bid_price(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.bid_price = f32_le(value);
        }
        self
    }

    fn set_bid_quantity(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.bid_quantity = f32_le(value);
        }
        self
    }

    fn set_ask_price(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.ask_price = f32_le(value);
        }
        self
    }

    fn set_ask_quantity(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.ask_quantity = f32_le(value);
        }
        self
    }

    fn set_date_time(&mut self, value: DateTime4Byte) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.date_time = value.to_le();
        }
        self
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(28) {
            self.symbol_id = value.to_le();
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
                28usize,
                core::mem::size_of::<MarketDataUpdateBidAskCompactFixedData>(),
                "MarketDataUpdateBidAskCompactFixedData sizeof expected {:} but was {:}",
                28usize,
                core::mem::size_of::<MarketDataUpdateBidAskCompactFixedData>()
            );
            assert_eq!(
                28u16,
                MarketDataUpdateBidAskCompactFixed::new().size(),
                "MarketDataUpdateBidAskCompactFixed sizeof expected {:} but was {:}",
                28u16,
                MarketDataUpdateBidAskCompactFixed::new().size(),
            );
            assert_eq!(
                MARKET_DATA_UPDATE_BID_ASK_COMPACT,
                MarketDataUpdateBidAskCompactFixed::new().r#type(),
                "MarketDataUpdateBidAskCompactFixed type expected {:} but was {:}",
                MARKET_DATA_UPDATE_BID_ASK_COMPACT,
                MarketDataUpdateBidAskCompactFixed::new().r#type(),
            );
            assert_eq!(
                117u16,
                MarketDataUpdateBidAskCompactFixed::new().r#type(),
                "MarketDataUpdateBidAskCompactFixed type expected {:} but was {:}",
                117u16,
                MarketDataUpdateBidAskCompactFixed::new().r#type(),
            );
            let d = MarketDataUpdateBidAskCompactFixedData::new();
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
                (core::ptr::addr_of!(d.bid_price) as usize) - p,
                "bid_price offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.bid_price) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.bid_quantity) as usize) - p,
                "bid_quantity offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.bid_quantity) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.ask_price) as usize) - p,
                "ask_price offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.ask_price) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.ask_quantity) as usize) - p,
                "ask_quantity offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.ask_quantity) as usize) - p,
            );
            assert_eq!(
                20usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
                "date_time offset expected {:} but was {:}",
                20usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
            );
            assert_eq!(
                24usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
                "symbol_id offset expected {:} but was {:}",
                24usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
            );
        }
    }
}