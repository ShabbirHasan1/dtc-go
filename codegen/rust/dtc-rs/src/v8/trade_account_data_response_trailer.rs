// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const TRADE_ACCOUNT_DATA_RESPONSE_TRAILER_VLS_SIZE: usize = 17;

/// size                       u16     = TradeAccountDataResponseTrailerVLSSize  (17)
/// type                       u16     = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER  (10130)
/// base_size                  u16     = TradeAccountDataResponseTrailerVLSSize  (17)
/// request_id                 u32     = 0
/// is_snapshot                bool    = false
/// is_first_message_in_batch  bool    = false
/// is_last_message_in_batch   bool    = false
/// trade_account              string  = ""
pub(crate) const TRADE_ACCOUNT_DATA_RESPONSE_TRAILER_VLS_DEFAULT: [u8; 17] =
    [17, 0, 146, 39, 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait TradeAccountDataResponseTrailer: Message {
    type Safe: TradeAccountDataResponseTrailer;
    type Unsafe: TradeAccountDataResponseTrailer;

    fn request_id(&self) -> u32;

    fn is_snapshot(&self) -> bool;

    fn is_first_message_in_batch(&self) -> bool;

    fn is_last_message_in_batch(&self) -> bool;

    fn trade_account(&self) -> &str;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_is_snapshot(&mut self, value: bool) -> &mut Self;

    fn set_is_first_message_in_batch(&mut self, value: bool) -> &mut Self;

    fn set_is_last_message_in_batch(&mut self, value: bool) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountDataResponseTrailer) {
        to.set_request_id(self.request_id());
        to.set_is_snapshot(self.is_snapshot());
        to.set_is_first_message_in_batch(self.is_first_message_in_batch());
        to.set_is_last_message_in_batch(self.is_last_message_in_batch());
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

pub struct TradeAccountDataResponseTrailerVLS {
    data: *const TradeAccountDataResponseTrailerVLSData,
    capacity: usize,
}

pub struct TradeAccountDataResponseTrailerVLSUnsafe {
    data: *const TradeAccountDataResponseTrailerVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct TradeAccountDataResponseTrailerVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    is_snapshot: bool,
    is_first_message_in_batch: bool,
    is_last_message_in_batch: bool,
    trade_account: VLS,
}

impl TradeAccountDataResponseTrailerVLSData {
    pub fn new() -> Self {
        Self {
            size: 17u16.to_le(),
            r#type: TRADE_ACCOUNT_DATA_RESPONSE_TRAILER.to_le(),
            base_size: 17u16.to_le(),
            request_id: 0u32.to_le(),
            is_snapshot: false,
            is_first_message_in_batch: false,
            is_last_message_in_batch: false,
            trade_account: crate::message::VLS::new(),
        }
    }
}

unsafe impl Send for TradeAccountDataResponseTrailerVLS {}
unsafe impl Send for TradeAccountDataResponseTrailerVLSUnsafe {}
unsafe impl Send for TradeAccountDataResponseTrailerVLSData {}

impl Drop for TradeAccountDataResponseTrailerVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountDataResponseTrailerVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountDataResponseTrailerVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountDataResponseTrailerVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountDataResponseTrailerVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountDataResponseTrailerVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountDataResponseTrailerVLS {
    type Target = TradeAccountDataResponseTrailerVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataResponseTrailerVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountDataResponseTrailerVLSUnsafe {
    type Target = TradeAccountDataResponseTrailerVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataResponseTrailerVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for TradeAccountDataResponseTrailerVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataResponseTrailerVLS(size: {}, type: {}, base_size: {}, request_id: {}, is_snapshot: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_snapshot(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.trade_account()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataResponseTrailerVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataResponseTrailerVLS(size: {}, type: {}, base_size: {}, request_id: {}, is_snapshot: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_snapshot(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.trade_account()).as_str())
    }
}

impl core::fmt::Display for TradeAccountDataResponseTrailerVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataResponseTrailerVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, is_snapshot: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_snapshot(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.trade_account()).as_str())
    }
}

impl core::fmt::Debug for TradeAccountDataResponseTrailerVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("TradeAccountDataResponseTrailerVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, is_snapshot: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, trade_account: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.is_snapshot(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.trade_account()).as_str())
    }
}

impl crate::Message for TradeAccountDataResponseTrailerVLS {
    type Data = TradeAccountDataResponseTrailerVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER;
    const BASE_SIZE: usize = 17;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataResponseTrailerVLSData::new(),
            ),
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
            data: data as *const TradeAccountDataResponseTrailerVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataResponseTrailerVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataResponseTrailerVLSData;
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
impl crate::Message for TradeAccountDataResponseTrailerVLSUnsafe {
    type Data = TradeAccountDataResponseTrailerVLSData;

    const TYPE: u16 = TRADE_ACCOUNT_DATA_RESPONSE_TRAILER;
    const BASE_SIZE: usize = 17;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataResponseTrailerVLSData::new(),
            ),
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
            data: data as *const TradeAccountDataResponseTrailerVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataResponseTrailerVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataResponseTrailerVLSData;
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
impl TradeAccountDataResponseTrailer for TradeAccountDataResponseTrailerVLS {
    type Safe = TradeAccountDataResponseTrailerVLS;
    type Unsafe = TradeAccountDataResponseTrailerVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn is_snapshot(&self) -> bool {
        self.is_snapshot
    }

    fn is_first_message_in_batch(&self) -> bool {
        self.is_first_message_in_batch
    }

    fn is_last_message_in_batch(&self) -> bool {
        self.is_last_message_in_batch
    }

    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_is_snapshot(&mut self, value: bool) -> &mut Self {
        self.is_snapshot = value;
        self
    }

    fn set_is_first_message_in_batch(&mut self, value: bool) -> &mut Self {
        self.is_first_message_in_batch = value;
        self
    }

    fn set_is_last_message_in_batch(&mut self, value: bool) -> &mut Self {
        self.is_last_message_in_batch = value;
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

impl TradeAccountDataResponseTrailer for TradeAccountDataResponseTrailerVLSUnsafe {
    type Safe = TradeAccountDataResponseTrailerVLS;
    type Unsafe = TradeAccountDataResponseTrailerVLSUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(10) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn is_snapshot(&self) -> bool {
        if self.is_out_of_bounds(11) {
            false
        } else {
            self.is_snapshot
        }
    }

    fn is_first_message_in_batch(&self) -> bool {
        if self.is_out_of_bounds(12) {
            false
        } else {
            self.is_first_message_in_batch
        }
    }

    fn is_last_message_in_batch(&self) -> bool {
        if self.is_out_of_bounds(13) {
            false
        } else {
            self.is_last_message_in_batch
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(17) {
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

    fn set_is_snapshot(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(11) {
            self.is_snapshot = value;
        }
        self
    }

    fn set_is_first_message_in_batch(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.is_first_message_in_batch = value;
        }
        self
    }

    fn set_is_last_message_in_batch(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(13) {
            self.is_last_message_in_batch = value;
        }
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(17) {
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
                17usize,
                core::mem::size_of::<TradeAccountDataResponseTrailerVLSData>(),
                "TradeAccountDataResponseTrailerVLSData sizeof expected {:} but was {:}",
                17usize,
                core::mem::size_of::<TradeAccountDataResponseTrailerVLSData>()
            );
            assert_eq!(
                17u16,
                TradeAccountDataResponseTrailerVLS::new().size(),
                "TradeAccountDataResponseTrailerVLS sizeof expected {:} but was {:}",
                17u16,
                TradeAccountDataResponseTrailerVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_DATA_RESPONSE_TRAILER,
                TradeAccountDataResponseTrailerVLS::new().r#type(),
                "TradeAccountDataResponseTrailerVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_DATA_RESPONSE_TRAILER,
                TradeAccountDataResponseTrailerVLS::new().r#type(),
            );
            assert_eq!(
                10130u16,
                TradeAccountDataResponseTrailerVLS::new().r#type(),
                "TradeAccountDataResponseTrailerVLS type expected {:} but was {:}",
                10130u16,
                TradeAccountDataResponseTrailerVLS::new().r#type(),
            );
            let d = TradeAccountDataResponseTrailerVLSData::new();
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
                (core::ptr::addr_of!(d.is_snapshot) as usize) - p,
                "is_snapshot offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.is_snapshot) as usize) - p,
            );
            assert_eq!(
                11usize,
                (core::ptr::addr_of!(d.is_first_message_in_batch) as usize) - p,
                "is_first_message_in_batch offset expected {:} but was {:}",
                11usize,
                (core::ptr::addr_of!(d.is_first_message_in_batch) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.is_last_message_in_batch) as usize) - p,
                "is_last_message_in_batch offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.is_last_message_in_batch) as usize) - p,
            );
            assert_eq!(
                13usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                13usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
        }
    }
}
