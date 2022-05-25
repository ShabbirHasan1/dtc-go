// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const MARKET_ORDERS_REQUEST_VLS_SIZE: usize = 28;

pub(crate) const MARKET_ORDERS_REQUEST_FIXED_SIZE: usize = 96;

/// size                                 u16                = MarketOrdersRequestVLSSize  (28)
/// type                                 u16                = MARKET_ORDERS_REQUEST  (150)
/// base_size                            u16                = MarketOrdersRequestVLSSize  (28)
/// request_action                       RequestActionEnum  = 0
/// symbol_id                            u32                = 0
/// symbol                               string             = ""
/// exchange                             string             = ""
/// send_quantities_greater_or_equal_to  u32                = 0
pub(crate) const MARKET_ORDERS_REQUEST_VLS_DEFAULT: [u8; 28] = [
    28, 0, 150, 0, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// size                                 u16                = MarketOrdersRequestFixedSize  (96)
/// type                                 u16                = MARKET_ORDERS_REQUEST  (150)
/// request_action                       RequestActionEnum  = SUBSCRIBE  (1)
/// symbol_id                            u32                = 0
/// symbol                               string64           = ""
/// exchange                             string16           = ""
/// send_quantities_greater_or_equal_to  u32                = 0
pub(crate) const MARKET_ORDERS_REQUEST_FIXED_DEFAULT: [u8; 96] = [
    96, 0, 150, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0,
];

pub trait MarketOrdersRequest: Message {
    type Safe: MarketOrdersRequest;
    type Unsafe: MarketOrdersRequest;

    fn request_action(&self) -> RequestActionEnum;

    fn symbol_id(&self) -> u32;

    fn symbol(&self) -> &str;

    fn exchange(&self) -> &str;

    fn send_quantities_greater_or_equal_to(&self) -> u32;

    fn set_request_action(&mut self, value: RequestActionEnum) -> &mut Self;

    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    fn set_symbol(&mut self, value: &str) -> &mut Self;

    fn set_exchange(&mut self, value: &str) -> &mut Self;

    fn set_send_quantities_greater_or_equal_to(&mut self, value: u32) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketOrdersRequest) {
        to.set_request_action(self.request_action());
        to.set_symbol_id(self.symbol_id());
        to.set_symbol(self.symbol());
        to.set_exchange(self.exchange());
        to.set_send_quantities_greater_or_equal_to(self.send_quantities_greater_or_equal_to());
    }
}

pub struct MarketOrdersRequestVLS {
    data: *const MarketOrdersRequestVLSData,
    capacity: usize,
}

pub struct MarketOrdersRequestVLSUnsafe {
    data: *const MarketOrdersRequestVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct MarketOrdersRequestVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_action: RequestActionEnum,
    symbol_id: u32,
    symbol: VLS,
    exchange: VLS,
    send_quantities_greater_or_equal_to: u32,
}

pub struct MarketOrdersRequestFixed {
    data: *const MarketOrdersRequestFixedData,
}

pub struct MarketOrdersRequestFixedUnsafe {
    data: *const MarketOrdersRequestFixedData,
}

#[repr(packed(1), C)]
pub struct MarketOrdersRequestFixedData {
    size: u16,
    r#type: u16,
    request_action: RequestActionEnum,
    symbol_id: u32,
    symbol: [u8; 64],
    exchange: [u8; 16],
    send_quantities_greater_or_equal_to: u32,
}

impl MarketOrdersRequestVLSData {
    pub fn new() -> Self {
        Self {
            size: 28u16.to_le(),
            r#type: MARKET_ORDERS_REQUEST.to_le(),
            base_size: 28u16.to_le(),
            request_action: RequestActionEnum::Subscribe.to_le(),
            symbol_id: 0u32.to_le(),
            symbol: crate::message::VLS::new(),
            exchange: crate::message::VLS::new(),
            send_quantities_greater_or_equal_to: 0u32.to_le(),
        }
    }
}

impl MarketOrdersRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 96u16.to_le(),
            r#type: MARKET_ORDERS_REQUEST.to_le(),
            request_action: RequestActionEnum::Subscribe.to_le(),
            symbol_id: 0u32.to_le(),
            symbol: [
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
                0, 0, 0, 0, 0, 0, 0, 0,
            ],
            exchange: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
            send_quantities_greater_or_equal_to: 0u32.to_le(),
        }
    }
}

unsafe impl Send for MarketOrdersRequestFixed {}
unsafe impl Send for MarketOrdersRequestFixedUnsafe {}
unsafe impl Send for MarketOrdersRequestFixedData {}
unsafe impl Send for MarketOrdersRequestVLS {}
unsafe impl Send for MarketOrdersRequestVLSUnsafe {}
unsafe impl Send for MarketOrdersRequestVLSData {}

impl Drop for MarketOrdersRequestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketOrdersRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketOrdersRequestVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketOrdersRequestVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketOrdersRequestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketOrdersRequestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketOrdersRequestVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketOrdersRequestVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketOrdersRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketOrdersRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketOrdersRequestVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketOrdersRequestVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketOrdersRequestFixed {
    type Target = MarketOrdersRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRequestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketOrdersRequestFixedUnsafe {
    type Target = MarketOrdersRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRequestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketOrdersRequestVLS {
    type Target = MarketOrdersRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRequestVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketOrdersRequestVLSUnsafe {
    type Target = MarketOrdersRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketOrdersRequestVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for MarketOrdersRequestFixed {
    type Data = MarketOrdersRequestFixedData;

    const BASE_SIZE: usize = 96;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRequestFixedData::new()),
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
            data: data as *const MarketOrdersRequestFixedData,
        }
    }
}
impl crate::Message for MarketOrdersRequestFixedUnsafe {
    type Data = MarketOrdersRequestFixedData;

    const BASE_SIZE: usize = 96;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRequestFixedData::new()),
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
            data: data as *const MarketOrdersRequestFixedData,
        }
    }
}
impl crate::Message for MarketOrdersRequestVLS {
    type Data = MarketOrdersRequestVLSData;

    const BASE_SIZE: usize = 28;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRequestVLSData::new()),
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
            data: data as *const MarketOrdersRequestVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for MarketOrdersRequestVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const MarketOrdersRequestVLSData;
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
impl crate::Message for MarketOrdersRequestVLSUnsafe {
    type Data = MarketOrdersRequestVLSData;

    const BASE_SIZE: usize = 28;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketOrdersRequestVLSData::new()),
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
            data: data as *const MarketOrdersRequestVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for MarketOrdersRequestVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const MarketOrdersRequestVLSData;
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
impl MarketOrdersRequest for MarketOrdersRequestVLS {
    type Safe = MarketOrdersRequestVLS;
    type Unsafe = MarketOrdersRequestVLSUnsafe;

    fn request_action(&self) -> RequestActionEnum {
        RequestActionEnum::from_le(self.request_action)
    }

    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    fn symbol(&self) -> &str {
        get_vls(self, self.symbol)
    }

    fn exchange(&self) -> &str {
        get_vls(self, self.exchange)
    }

    fn send_quantities_greater_or_equal_to(&self) -> u32 {
        u32::from_le(self.send_quantities_greater_or_equal_to)
    }

    fn set_request_action(&mut self, value: RequestActionEnum) -> &mut Self {
        self.request_action = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        self.symbol = set_vls(self, self.symbol, value);
        self
    }

    fn set_exchange(&mut self, value: &str) -> &mut Self {
        self.exchange = set_vls(self, self.exchange, value);
        self
    }

    fn set_send_quantities_greater_or_equal_to(&mut self, value: u32) -> &mut Self {
        self.send_quantities_greater_or_equal_to = value.to_le();
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

impl MarketOrdersRequest for MarketOrdersRequestVLSUnsafe {
    type Safe = MarketOrdersRequestVLS;
    type Unsafe = MarketOrdersRequestVLSUnsafe;

    fn request_action(&self) -> RequestActionEnum {
        if self.is_out_of_bounds(12) {
            RequestActionEnum::Subscribe.to_le()
        } else {
            RequestActionEnum::from_le(self.request_action)
        }
    }

    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(16) {
            0u32.to_le()
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    fn symbol(&self) -> &str {
        if self.is_out_of_bounds(20) {
            ""
        } else {
            get_vls(self, self.symbol)
        }
    }

    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(24) {
            ""
        } else {
            get_vls(self, self.exchange)
        }
    }

    fn send_quantities_greater_or_equal_to(&self) -> u32 {
        if self.is_out_of_bounds(28) {
            0u32.to_le()
        } else {
            u32::from_le(self.send_quantities_greater_or_equal_to)
        }
    }

    fn set_request_action(&mut self, value: RequestActionEnum) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_action = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.symbol_id = value.to_le();
        }
        self
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.symbol = set_vls(self, self.symbol, value);
        }
        self
    }

    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.exchange = set_vls(self, self.exchange, value);
        }
        self
    }

    fn set_send_quantities_greater_or_equal_to(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(28) {
            self.send_quantities_greater_or_equal_to = value.to_le();
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

impl MarketOrdersRequest for MarketOrdersRequestFixed {
    type Safe = MarketOrdersRequestFixed;
    type Unsafe = MarketOrdersRequestFixedUnsafe;

    fn request_action(&self) -> RequestActionEnum {
        RequestActionEnum::from_le(self.request_action)
    }

    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    fn symbol(&self) -> &str {
        get_fixed(&self.symbol[..])
    }

    fn exchange(&self) -> &str {
        get_fixed(&self.exchange[..])
    }

    fn send_quantities_greater_or_equal_to(&self) -> u32 {
        u32::from_le(self.send_quantities_greater_or_equal_to)
    }

    fn set_request_action(&mut self, value: RequestActionEnum) -> &mut Self {
        self.request_action = unsafe { core::mem::transmute((value as i32).to_le()) };
        self
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.symbol[..], value);
        self
    }

    fn set_exchange(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.exchange[..], value);
        self
    }

    fn set_send_quantities_greater_or_equal_to(&mut self, value: u32) -> &mut Self {
        self.send_quantities_greater_or_equal_to = value.to_le();
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

impl MarketOrdersRequest for MarketOrdersRequestFixedUnsafe {
    type Safe = MarketOrdersRequestFixed;
    type Unsafe = MarketOrdersRequestFixedUnsafe;

    fn request_action(&self) -> RequestActionEnum {
        if self.is_out_of_bounds(8) {
            RequestActionEnum::Subscribe.to_le()
        } else {
            RequestActionEnum::from_le(self.request_action)
        }
    }

    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(12) {
            0u32.to_le()
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    fn symbol(&self) -> &str {
        if self.is_out_of_bounds(76) {
            ""
        } else {
            get_fixed(&self.symbol[..])
        }
    }

    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(92) {
            ""
        } else {
            get_fixed(&self.exchange[..])
        }
    }

    fn send_quantities_greater_or_equal_to(&self) -> u32 {
        if self.is_out_of_bounds(96) {
            0u32.to_le()
        } else {
            u32::from_le(self.send_quantities_greater_or_equal_to)
        }
    }

    fn set_request_action(&mut self, value: RequestActionEnum) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_action = unsafe { core::mem::transmute((value as i32).to_le()) };
        }
        self
    }

    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.symbol_id = value.to_le();
        }
        self
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(76) {
            set_fixed(&mut self.symbol[..], value);
        }
        self
    }

    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(92) {
            set_fixed(&mut self.exchange[..], value);
        }
        self
    }

    fn set_send_quantities_greater_or_equal_to(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(96) {
            self.send_quantities_greater_or_equal_to = value.to_le();
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
                96usize,
                core::mem::size_of::<MarketOrdersRequestFixedData>(),
                "MarketOrdersRequestFixedData sizeof expected {:} but was {:}",
                96usize,
                core::mem::size_of::<MarketOrdersRequestFixedData>()
            );
            assert_eq!(
                96u16,
                MarketOrdersRequestFixed::new().size(),
                "MarketOrdersRequestFixed sizeof expected {:} but was {:}",
                96u16,
                MarketOrdersRequestFixed::new().size(),
            );
            assert_eq!(
                MARKET_ORDERS_REQUEST,
                MarketOrdersRequestFixed::new().r#type(),
                "MarketOrdersRequestFixed type expected {:} but was {:}",
                MARKET_ORDERS_REQUEST,
                MarketOrdersRequestFixed::new().r#type(),
            );
            assert_eq!(
                150u16,
                MarketOrdersRequestFixed::new().r#type(),
                "MarketOrdersRequestFixed type expected {:} but was {:}",
                150u16,
                MarketOrdersRequestFixed::new().r#type(),
            );
            let d = MarketOrdersRequestFixedData::new();
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
                (core::ptr::addr_of!(d.request_action) as usize) - p,
                "request_action offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.request_action) as usize) - p,
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
                (core::ptr::addr_of!(d.symbol) as usize) - p,
                "symbol offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.symbol) as usize) - p,
            );
            assert_eq!(
                76usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
                "exchange offset expected {:} but was {:}",
                76usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
            );
            assert_eq!(
                92usize,
                (core::ptr::addr_of!(d.send_quantities_greater_or_equal_to) as usize) - p,
                "send_quantities_greater_or_equal_to offset expected {:} but was {:}",
                92usize,
                (core::ptr::addr_of!(d.send_quantities_greater_or_equal_to) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                28usize,
                core::mem::size_of::<MarketOrdersRequestVLSData>(),
                "MarketOrdersRequestVLSData sizeof expected {:} but was {:}",
                28usize,
                core::mem::size_of::<MarketOrdersRequestVLSData>()
            );
            assert_eq!(
                28u16,
                MarketOrdersRequestVLS::new().size(),
                "MarketOrdersRequestVLS sizeof expected {:} but was {:}",
                28u16,
                MarketOrdersRequestVLS::new().size(),
            );
            assert_eq!(
                MARKET_ORDERS_REQUEST,
                MarketOrdersRequestVLS::new().r#type(),
                "MarketOrdersRequestVLS type expected {:} but was {:}",
                MARKET_ORDERS_REQUEST,
                MarketOrdersRequestVLS::new().r#type(),
            );
            assert_eq!(
                150u16,
                MarketOrdersRequestVLS::new().r#type(),
                "MarketOrdersRequestVLS type expected {:} but was {:}",
                150u16,
                MarketOrdersRequestVLS::new().r#type(),
            );
            let d = MarketOrdersRequestVLSData::new();
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
                (core::ptr::addr_of!(d.request_action) as usize) - p,
                "request_action offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.request_action) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
                "symbol_id offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.symbol) as usize) - p,
                "symbol offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.symbol) as usize) - p,
            );
            assert_eq!(
                20usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
                "exchange offset expected {:} but was {:}",
                20usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
            );
            assert_eq!(
                24usize,
                (core::ptr::addr_of!(d.send_quantities_greater_or_equal_to) as usize) - p,
                "send_quantities_greater_or_equal_to offset expected {:} but was {:}",
                24usize,
                (core::ptr::addr_of!(d.send_quantities_greater_or_equal_to) as usize) - p,
            );
        }
    }
}