// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const MARKET_ORDERS_REJECT_VLS_SIZE: usize = 16;

pub(crate) const MARKET_ORDERS_REJECT_FIXED_SIZE: usize = 104;

/// size         u16     = MarketOrdersRejectVLSSize  (16)
/// type         u16     = MARKET_ORDERS_REJECT  (151)
/// base_size    u16     = MarketOrdersRejectVLSSize  (16)
/// symbol_id    u32     = 0
/// reject_text  string  = ""
pub(crate) const MARKET_ORDERS_REJECT_VLS_DEFAULT: [u8; 16] =
    [16, 0, 151, 0, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size         u16       = MarketOrdersRejectFixedSize  (104)
/// type         u16       = MARKET_ORDERS_REJECT  (151)
/// symbol_id    u32       = 0
/// reject_text  string96  = ""
pub(crate) const MARKET_ORDERS_REJECT_FIXED_DEFAULT: [u8; 104] = [
    104, 0, 151, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait MarketOrdersReject: Message {
    type Safe: MarketOrdersReject;
    type Unsafe: MarketOrdersReject;

    fn symbol_id(&self) -> u32;

    fn reject_text(&self) -> &str;

    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    fn set_reject_text(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketOrdersReject) {
        to.set_symbol_id(self.symbol_id());
        to.set_reject_text(self.reject_text());
    }
}

pub struct MarketOrdersRejectVLS {
    data: *const MarketOrdersRejectVLSData,
    capacity: usize,
}

pub struct MarketOrdersRejectVLSUnsafe {
    data: *const MarketOrdersRejectVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct MarketOrdersRejectVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    symbol_id: u32,
    reject_text: VLS,
}

pub struct MarketOrdersRejectFixed {
    data: *const MarketOrdersRejectFixedData,
}

pub struct MarketOrdersRejectFixedUnsafe {
    data: *const MarketOrdersRejectFixedData,
}

#[repr(packed(1), C)]
pub struct MarketOrdersRejectFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    reject_text: [u8; 96],
}

impl MarketOrdersRejectVLSData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: MARKET_ORDERS_REJECT.to_le(),
            base_size: 16u16.to_le(),
            symbol_id: 0u32.to_le(),
            reject_text: crate::message::VLS::new(),
        }
    }
}

impl MarketOrdersRejectFixedData {
    pub fn new() -> Self {
        Self {
            size: 104u16.to_le(),
            r#type: MARKET_ORDERS_REJECT.to_le(),
            symbol_id: 0u32.to_le(),
            reject_text: [
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
            ],
        }
    }
}

unsafe impl Send for MarketOrdersRejectFixed {}
unsafe impl Send for MarketOrdersRejectFixedUnsafe {}
unsafe impl Send for MarketOrdersRejectFixedData {}
unsafe impl Send for MarketOrdersRejectVLS {}
unsafe impl Send for MarketOrdersRejectVLSUnsafe {}
unsafe impl Send for MarketOrdersRejectVLSData {}

impl Drop for MarketOrdersRejectFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketOrdersRejectFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketOrdersRejectVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketOrdersRejectVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketOrdersRejectFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketOrdersRejectFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketOrdersRejectVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketOrdersRejectVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketOrdersRejectFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketOrdersRejectFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketOrdersRejectVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketOrdersRejectVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketOrdersRejectFixed {
    type Target = MarketOrdersRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRejectFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketOrdersRejectFixedUnsafe {
    type Target = MarketOrdersRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRejectFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketOrdersRejectVLS {
    type Target = MarketOrdersRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRejectVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketOrdersRejectVLSUnsafe {
    type Target = MarketOrdersRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRejectVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketOrdersRejectFixed {
    type Data = MarketOrdersRejectFixedData;

    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRejectFixedData::new()),
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
            data: data as *const MarketOrdersRejectFixedData,
        }
    }
}
impl crate::Message for MarketOrdersRejectFixedUnsafe {
    type Data = MarketOrdersRejectFixedData;

    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRejectFixedData::new()),
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
            data: data as *const MarketOrdersRejectFixedData,
        }
    }
}
impl crate::Message for MarketOrdersRejectVLS {
    type Data = MarketOrdersRejectVLSData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRejectVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY,
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
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const MarketOrdersRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for MarketOrdersRejectVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const MarketOrdersRejectVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }
}
impl crate::Message for MarketOrdersRejectVLSUnsafe {
    type Data = MarketOrdersRejectVLSData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRejectVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY,
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
        u16::from_le(self.base_size)
    }

    #[inline]
    fn capacity(&self) -> u16 {
        self.capacity as u16
    }

    #[inline]
    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {
        Self {
            data: data as *const MarketOrdersRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for MarketOrdersRejectVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const MarketOrdersRejectVLSData;
    }

    #[inline]
    fn set_capacity(&mut self, capacity: u16) {
        self.capacity = capacity as usize;
    }

    #[inline]
    fn set_size(&mut self, size: u16) {
        self.size = size.to_le();
    }
}
impl MarketOrdersReject for MarketOrdersRejectVLS {
    type Safe = MarketOrdersRejectVLS;
    type Unsafe = MarketOrdersRejectVLSUnsafe;

    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    fn reject_text(&self) -> &str {
        get_vls(self, self.reject_text)
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        self.reject_text = set_vls(self, self.reject_text, value);
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

impl MarketOrdersReject for MarketOrdersRejectVLSUnsafe {
    type Safe = MarketOrdersRejectVLS;
    type Unsafe = MarketOrdersRejectVLSUnsafe;

    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(12) {
            0u32.to_le()
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.reject_text)
        }
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.symbol_id = value.to_le();
        }
        self
    }

    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.reject_text = set_vls(self, self.reject_text, value);
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

impl MarketOrdersReject for MarketOrdersRejectFixed {
    type Safe = MarketOrdersRejectFixed;
    type Unsafe = MarketOrdersRejectFixedUnsafe;

    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    fn reject_text(&self) -> &str {
        get_fixed(&self.reject_text[..])
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.reject_text[..], value);
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

impl MarketOrdersReject for MarketOrdersRejectFixedUnsafe {
    type Safe = MarketOrdersRejectFixed;
    type Unsafe = MarketOrdersRejectFixedUnsafe;

    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32.to_le()
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(104) {
            ""
        } else {
            get_fixed(&self.reject_text[..])
        }
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.symbol_id = value.to_le();
        }
        self
    }

    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(104) {
            set_fixed(&mut self.reject_text[..], value);
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
                104usize,
                core::mem::size_of::<MarketOrdersRejectFixedData>(),
                "MarketOrdersRejectFixedData sizeof expected {:} but was {:}",
                104usize,
                core::mem::size_of::<MarketOrdersRejectFixedData>()
            );
            assert_eq!(
                104u16,
                MarketOrdersRejectFixed::new().size(),
                "MarketOrdersRejectFixed sizeof expected {:} but was {:}",
                104u16,
                MarketOrdersRejectFixed::new().size(),
            );
            assert_eq!(
                MARKET_ORDERS_REJECT,
                MarketOrdersRejectFixed::new().r#type(),
                "MarketOrdersRejectFixed type expected {:} but was {:}",
                MARKET_ORDERS_REJECT,
                MarketOrdersRejectFixed::new().r#type(),
            );
            assert_eq!(
                151u16,
                MarketOrdersRejectFixed::new().r#type(),
                "MarketOrdersRejectFixed type expected {:} but was {:}",
                151u16,
                MarketOrdersRejectFixed::new().r#type(),
            );
            let d = MarketOrdersRejectFixedData::new();
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
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
                "reject_text offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                16usize,
                core::mem::size_of::<MarketOrdersRejectVLSData>(),
                "MarketOrdersRejectVLSData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<MarketOrdersRejectVLSData>()
            );
            assert_eq!(
                16u16,
                MarketOrdersRejectVLS::new().size(),
                "MarketOrdersRejectVLS sizeof expected {:} but was {:}",
                16u16,
                MarketOrdersRejectVLS::new().size(),
            );
            assert_eq!(
                MARKET_ORDERS_REJECT,
                MarketOrdersRejectVLS::new().r#type(),
                "MarketOrdersRejectVLS type expected {:} but was {:}",
                MARKET_ORDERS_REJECT,
                MarketOrdersRejectVLS::new().r#type(),
            );
            assert_eq!(
                151u16,
                MarketOrdersRejectVLS::new().r#type(),
                "MarketOrdersRejectVLS type expected {:} but was {:}",
                151u16,
                MarketOrdersRejectVLS::new().r#type(),
            );
            let d = MarketOrdersRejectVLSData::new();
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
                (core::ptr::addr_of!(d.base_size) as usize) - p,
                "base_size offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.base_size) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
                "symbol_id offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
                "reject_text offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
            );
        }
    }
}
