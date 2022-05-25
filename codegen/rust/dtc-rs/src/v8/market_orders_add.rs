// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const MARKET_ORDERS_ADD_FIXED_SIZE: usize = 40;

/// size       u16                          = MarketOrdersAddFixedSize  (40)
/// type       u16                          = MARKET_ORDERS_ADD  (152)
/// symbol_id  u32                          = 0
/// side       BuySellEnum                  = BUY_SELL_UNSET  (0)
/// quantity   u32                          = 0
/// price      f64                          = 0
/// order_id   u64                          = 0
/// date_time  DateTimeWithMicrosecondsInt  = 0
pub(crate) const MARKET_ORDERS_ADD_FIXED_DEFAULT: [u8; 40] = [
    40, 0, 152, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait MarketOrdersAdd: Message {
    type Safe: MarketOrdersAdd;
    type Unsafe: MarketOrdersAdd;

    fn symbol_id(&self) -> u32;

    fn side(&self) -> BuySellEnum;

    fn quantity(&self) -> u32;

    fn price(&self) -> f64;

    fn order_id(&self) -> u64;

    fn date_time(&self) -> DateTimeWithMicrosecondsInt;

    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    fn set_side(&mut self, value: BuySellEnum) -> &mut Self;

    fn set_quantity(&mut self, value: u32) -> &mut Self;

    fn set_price(&mut self, value: f64) -> &mut Self;

    fn set_order_id(&mut self, value: u64) -> &mut Self;

    fn set_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketOrdersAdd) {
        to.set_symbol_id(self.symbol_id());
        to.set_side(self.side());
        to.set_quantity(self.quantity());
        to.set_price(self.price());
        to.set_order_id(self.order_id());
        to.set_date_time(self.date_time());
    }
}

pub struct MarketOrdersAddFixed {
    data: *const MarketOrdersAddFixedData,
}

pub struct MarketOrdersAddFixedUnsafe {
    data: *const MarketOrdersAddFixedData,
}

#[repr(packed(1), C)]
pub struct MarketOrdersAddFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    side: BuySellEnum,
    quantity: u32,
    price: f64,
    order_id: u64,
    date_time: DateTimeWithMicrosecondsInt,
}

impl MarketOrdersAddFixedData {
    pub fn new() -> Self {
        Self {
            size: 40u16.to_le(),
            r#type: MARKET_ORDERS_ADD.to_le(),
            symbol_id: 0u32.to_le(),
            side: BuySellEnum::BuySellUnset.to_le(),
            quantity: 0u32.to_le(),
            price: 0.0,
            order_id: 0u64.to_le(),
            date_time: 0i64.to_le(),
        }
    }
}

unsafe impl Send for MarketOrdersAddFixed {}
unsafe impl Send for MarketOrdersAddFixedUnsafe {}
unsafe impl Send for MarketOrdersAddFixedData {}

impl Drop for MarketOrdersAddFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketOrdersAddFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketOrdersAddFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketOrdersAddFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketOrdersAddFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketOrdersAddFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketOrdersAddFixed {
    type Target = MarketOrdersAddFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersAddFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketOrdersAddFixedUnsafe {
    type Target = MarketOrdersAddFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersAddFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketOrdersAddFixed {
    type Data = MarketOrdersAddFixedData;

    const BASE_SIZE: usize = 40;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersAddFixedData::new()),
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
            data: data as *const MarketOrdersAddFixedData,
        }
    }
}
impl crate::Message for MarketOrdersAddFixedUnsafe {
    type Data = MarketOrdersAddFixedData;

    const BASE_SIZE: usize = 40;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersAddFixedData::new()),
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
            data: data as *const MarketOrdersAddFixedData,
        }
    }
}
impl MarketOrdersAdd for MarketOrdersAddFixed {
    type Safe = MarketOrdersAddFixed;
    type Unsafe = MarketOrdersAddFixedUnsafe;

    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    fn side(&self) -> BuySellEnum {
        BuySellEnum::from_le(self.side)
    }

    fn quantity(&self) -> u32 {
        u32::from_le(self.quantity)
    }

    fn price(&self) -> f64 {
        f64_le(self.price)
    }

    fn order_id(&self) -> u64 {
        u64::from_le(self.order_id)
    }

    fn date_time(&self) -> DateTimeWithMicrosecondsInt {
        i64::from_le(self.date_time)
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    fn set_side(&mut self, value: BuySellEnum) -> &mut Self {
        self.side = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }

    fn set_quantity(&mut self, value: u32) -> &mut Self {
        self.quantity = value.to_le();
        self
    }

    fn set_price(&mut self, value: f64) -> &mut Self {
        self.price = f64_le(value);
        self
    }

    fn set_order_id(&mut self, value: u64) -> &mut Self {
        self.order_id = value.to_le();
        self
    }

    fn set_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self {
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

impl MarketOrdersAdd for MarketOrdersAddFixedUnsafe {
    type Safe = MarketOrdersAddFixed;
    type Unsafe = MarketOrdersAddFixedUnsafe;

    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32.to_le()
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    fn side(&self) -> BuySellEnum {
        if self.is_out_of_bounds(12) {
            BuySellEnum::BuySellUnset.to_le()
        } else {
            BuySellEnum::from_le(self.side)
        }
    }

    fn quantity(&self) -> u32 {
        if self.is_out_of_bounds(16) {
            0u32.to_le()
        } else {
            u32::from_le(self.quantity)
        }
    }

    fn price(&self) -> f64 {
        if self.is_out_of_bounds(24) {
            0.0
        } else {
            f64_le(self.price)
        }
    }

    fn order_id(&self) -> u64 {
        if self.is_out_of_bounds(32) {
            0u64.to_le()
        } else {
            u64::from_le(self.order_id)
        }
    }

    fn date_time(&self) -> DateTimeWithMicrosecondsInt {
        if self.is_out_of_bounds(40) {
            0i64.to_le()
        } else {
            i64::from_le(self.date_time)
        }
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.symbol_id = value.to_le();
        }
        self
    }

    fn set_side(&mut self, value: BuySellEnum) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.side = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }

    fn set_quantity(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.quantity = value.to_le();
        }
        self
    }

    fn set_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.price = f64_le(value);
        }
        self
    }

    fn set_order_id(&mut self, value: u64) -> &mut Self {
        if !self.is_out_of_bounds(32) {
            self.order_id = value.to_le();
        }
        self
    }

    fn set_date_time(&mut self, value: DateTimeWithMicrosecondsInt) -> &mut Self {
        if !self.is_out_of_bounds(40) {
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
                40usize,
                core::mem::size_of::<MarketOrdersAddFixedData>(),
                "MarketOrdersAddFixedData sizeof expected {:} but was {:}",
                40usize,
                core::mem::size_of::<MarketOrdersAddFixedData>()
            );
            assert_eq!(
                40u16,
                MarketOrdersAddFixed::new().size(),
                "MarketOrdersAddFixed sizeof expected {:} but was {:}",
                40u16,
                MarketOrdersAddFixed::new().size(),
            );
            assert_eq!(
                MARKET_ORDERS_ADD,
                MarketOrdersAddFixed::new().r#type(),
                "MarketOrdersAddFixed type expected {:} but was {:}",
                MARKET_ORDERS_ADD,
                MarketOrdersAddFixed::new().r#type(),
            );
            assert_eq!(
                152u16,
                MarketOrdersAddFixed::new().r#type(),
                "MarketOrdersAddFixed type expected {:} but was {:}",
                152u16,
                MarketOrdersAddFixed::new().r#type(),
            );
            let d = MarketOrdersAddFixedData::new();
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
                (core::ptr::addr_of!(d.side) as usize) - p,
                "side offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.side) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.quantity) as usize) - p,
                "quantity offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.quantity) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.price) as usize) - p,
                "price offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.price) as usize) - p,
            );
            assert_eq!(
                24usize,
                (core::ptr::addr_of!(d.order_id) as usize) - p,
                "order_id offset expected {:} but was {:}",
                24usize,
                (core::ptr::addr_of!(d.order_id) as usize) - p,
            );
            assert_eq!(
                32usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
                "date_time offset expected {:} but was {:}",
                32usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
            );
        }
    }
}
