// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const TRADE_ACCOUNT_DATA_DELETE_VLS_SIZE: usize = 14;

/// size           u16     = TradeAccountDataDeleteVLSSize  (14)
/// type           u16     = TRADE_ACCOUNT_DATA_DELETE  (10118)
/// base_size      u16     = TradeAccountDataDeleteVLSSize  (14)
/// request_id     u32     = 0
/// trade_account  string  = ""
pub(crate) const TRADE_ACCOUNT_DATA_DELETE_VLS_DEFAULT: [u8; 14] =
    [14, 0, 134, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait TradeAccountDataDelete: Message {
    type Safe: TradeAccountDataDelete;
    type Unsafe: TradeAccountDataDelete;

    fn request_id(&self) -> u32;

    fn trade_account(&self) -> &str;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountDataDelete) {
        to.set_request_id(self.request_id());
        to.set_trade_account(self.trade_account());
    }

    #[inline]
    fn parse<F: Fn(Parsed<Self::Safe, Self::Unsafe>) -> Result<(), crate::Error>>(
        data: &[u8],
        f: F,
    ) -> Result<(), crate::Error> {
        if data.len() < 6 {
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

pub struct TradeAccountDataDeleteVLS {
    data: *const TradeAccountDataDeleteVLSData,
    capacity: usize,
}

pub struct TradeAccountDataDeleteVLSUnsafe {
    data: *const TradeAccountDataDeleteVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct TradeAccountDataDeleteVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    trade_account: VLS,
}

impl TradeAccountDataDeleteVLSData {
    pub fn new() -> Self {
        Self {
            size: 14u16.to_le(),
            r#type: TRADE_ACCOUNT_DATA_DELETE.to_le(),
            base_size: 14u16.to_le(),
            request_id: 0u32.to_le(),
            trade_account: crate::message::VLS::new(),
        }
    }
}

unsafe impl Send for TradeAccountDataDeleteVLS {}
unsafe impl Send for TradeAccountDataDeleteVLSUnsafe {}
unsafe impl Send for TradeAccountDataDeleteVLSData {}

impl Drop for TradeAccountDataDeleteVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountDataDeleteVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountDataDeleteVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountDataDeleteVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountDataDeleteVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountDataDeleteVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountDataDeleteVLS {
    type Target = TradeAccountDataDeleteVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataDeleteVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountDataDeleteVLSUnsafe {
    type Target = TradeAccountDataDeleteVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataDeleteVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for TradeAccountDataDeleteVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataDeleteVLS(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataDeleteVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataDeleteVLS(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account()).as_str())
    }
}

impl core::fmt::Display for TradeAccountDataDeleteVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataDeleteVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataDeleteVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataDeleteVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.trade_account()).as_str())
    }
}

impl crate::Message for TradeAccountDataDeleteVLS {
    type Data = TradeAccountDataDeleteVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_DELETE;
    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountDataDeleteVLSData::new()),
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
            data: data as *const TradeAccountDataDeleteVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataDeleteVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataDeleteVLSData;
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
impl crate::Message for TradeAccountDataDeleteVLSUnsafe {
    type Data = TradeAccountDataDeleteVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_DELETE;
    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountDataDeleteVLSData::new()),
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
            data: data as *const TradeAccountDataDeleteVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataDeleteVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataDeleteVLSData;
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
impl TradeAccountDataDelete for TradeAccountDataDeleteVLS {
    type Safe = TradeAccountDataDeleteVLS;
    type Unsafe = TradeAccountDataDeleteVLSUnsafe;

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

impl TradeAccountDataDelete for TradeAccountDataDeleteVLSUnsafe {
    type Safe = TradeAccountDataDeleteVLS;
    type Unsafe = TradeAccountDataDeleteVLSUnsafe;

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
                core::mem::size_of::<TradeAccountDataDeleteVLSData>(),
                "TradeAccountDataDeleteVLSData sizeof expected {:} but was {:}",
                14usize,
                core::mem::size_of::<TradeAccountDataDeleteVLSData>()
            );
            assert_eq!(
                14u16,
                TradeAccountDataDeleteVLS::new().size(),
                "TradeAccountDataDeleteVLS sizeof expected {:} but was {:}",
                14u16,
                TradeAccountDataDeleteVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_DATA_DELETE,
                TradeAccountDataDeleteVLS::new().r#type(),
                "TradeAccountDataDeleteVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_DATA_DELETE,
                TradeAccountDataDeleteVLS::new().r#type(),
            );
            assert_eq!(
                10118u16,
                TradeAccountDataDeleteVLS::new().r#type(),
                "TradeAccountDataDeleteVLS type expected {:} but was {:}",
                10118u16,
                TradeAccountDataDeleteVLS::new().r#type(),
            );
            let d = TradeAccountDataDeleteVLSData::new();
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
