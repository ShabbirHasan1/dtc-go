// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const TRADE_ACCOUNT_DATA_REQUEST_VLS_SIZE: usize = 14;

/// size           u16     = TradeAccountDataRequestVLSSize  (14)
/// type           u16     = TRADE_ACCOUNT_DATA_REQUEST  (10115)
/// base_size      u16     = TradeAccountDataRequestVLSSize  (14)
/// request_id     u32     = 0
/// trade_account  string  = ""
pub(crate) const TRADE_ACCOUNT_DATA_REQUEST_VLS_DEFAULT: [u8; 14] =
    [14, 0, 131, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait TradeAccountDataRequest: Message {
    type Safe: TradeAccountDataRequest;
    type Unsafe: TradeAccountDataRequest;

    fn request_id(&self) -> u32;

    fn trade_account(&self) -> &str;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountDataRequest) {
        to.set_request_id(self.request_id());
        to.set_trade_account(self.trade_account());
    }
}

pub struct TradeAccountDataRequestVLS {
    data: *const TradeAccountDataRequestVLSData,
    capacity: usize,
}

pub struct TradeAccountDataRequestVLSUnsafe {
    data: *const TradeAccountDataRequestVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct TradeAccountDataRequestVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    trade_account: VLS,
}

impl TradeAccountDataRequestVLSData {
    pub fn new() -> Self {
        Self {
            size: 14u16.to_le(),
            r#type: TRADE_ACCOUNT_DATA_REQUEST.to_le(),
            base_size: 14u16.to_le(),
            request_id: 0u32.to_le(),
            trade_account: crate::message::VLS::new(),
        }
    }
}

unsafe impl Send for TradeAccountDataRequestVLS {}
unsafe impl Send for TradeAccountDataRequestVLSUnsafe {}
unsafe impl Send for TradeAccountDataRequestVLSData {}

impl Drop for TradeAccountDataRequestVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountDataRequestVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountDataRequestVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountDataRequestVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountDataRequestVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountDataRequestVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountDataRequestVLS {
    type Target = TradeAccountDataRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataRequestVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountDataRequestVLSUnsafe {
    type Target = TradeAccountDataRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataRequestVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for TradeAccountDataRequestVLS {
    type Data = TradeAccountDataRequestVLSData;

    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountDataRequestVLSData::new()),
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
            data: data as *const TradeAccountDataRequestVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataRequestVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataRequestVLSData;
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
impl crate::Message for TradeAccountDataRequestVLSUnsafe {
    type Data = TradeAccountDataRequestVLSData;

    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountDataRequestVLSData::new()),
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
            data: data as *const TradeAccountDataRequestVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataRequestVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataRequestVLSData;
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
impl TradeAccountDataRequest for TradeAccountDataRequestVLS {
    type Safe = TradeAccountDataRequestVLS;
    type Unsafe = TradeAccountDataRequestVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = set_vls(self, self.trade_account, value);
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

impl TradeAccountDataRequest for TradeAccountDataRequestVLSUnsafe {
    type Safe = TradeAccountDataRequestVLS;
    type Unsafe = TradeAccountDataRequestVLSUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(10) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(14) {
            self.trade_account = set_vls(self, self.trade_account, value);
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
                14usize,
                core::mem::size_of::<TradeAccountDataRequestVLSData>(),
                "TradeAccountDataRequestVLSData sizeof expected {:} but was {:}",
                14usize,
                core::mem::size_of::<TradeAccountDataRequestVLSData>()
            );
            assert_eq!(
                14u16,
                TradeAccountDataRequestVLS::new().size(),
                "TradeAccountDataRequestVLS sizeof expected {:} but was {:}",
                14u16,
                TradeAccountDataRequestVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_DATA_REQUEST,
                TradeAccountDataRequestVLS::new().r#type(),
                "TradeAccountDataRequestVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_DATA_REQUEST,
                TradeAccountDataRequestVLS::new().r#type(),
            );
            assert_eq!(
                10115u16,
                TradeAccountDataRequestVLS::new().r#type(),
                "TradeAccountDataRequestVLS type expected {:} but was {:}",
                10115u16,
                TradeAccountDataRequestVLS::new().r#type(),
            );
            let d = TradeAccountDataRequestVLSData::new();
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
                6usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
            assert_eq!(
                10usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
        }
    }
}
