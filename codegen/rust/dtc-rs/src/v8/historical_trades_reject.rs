// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const HISTORICAL_TRADES_REJECT_VLS_SIZE: usize = 14;

pub(crate) const HISTORICAL_TRADES_REJECT_FIXED_SIZE: usize = 104;

/// size         u16     = HistoricalTradesRejectVLSSize  (14)
/// type         u16     = HISTORICAL_TRADES_REJECT  (10101)
/// base_size    u16     = HistoricalTradesRejectVLSSize  (14)
/// request_id   i32     = 0
/// reject_text  string  = ""
pub(crate) const HISTORICAL_TRADES_REJECT_VLS_DEFAULT: [u8; 14] =
    [14, 0, 117, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size         u16       = HistoricalTradesRejectFixedSize  (104)
/// type         u16       = HISTORICAL_TRADES_REJECT  (10101)
/// request_id   i32       = 0
/// reject_text  string96  = ""
pub(crate) const HISTORICAL_TRADES_REJECT_FIXED_DEFAULT: [u8; 104] = [
    104, 0, 117, 39, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait HistoricalTradesReject: Message {
    type Safe: HistoricalTradesReject;
    type Unsafe: HistoricalTradesReject;

    fn request_id(&self) -> i32;

    fn reject_text(&self) -> &str;

    fn set_request_id(&mut self, value: i32) -> &mut Self;

    fn set_reject_text(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl HistoricalTradesReject) {
        to.set_request_id(self.request_id());
        to.set_reject_text(self.reject_text());
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

pub struct HistoricalTradesRejectVLS {
    data: *const HistoricalTradesRejectVLSData,
    capacity: usize,
}

pub struct HistoricalTradesRejectVLSUnsafe {
    data: *const HistoricalTradesRejectVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct HistoricalTradesRejectVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    reject_text: VLS,
}

pub struct HistoricalTradesRejectFixed {
    data: *const HistoricalTradesRejectFixedData,
}

pub struct HistoricalTradesRejectFixedUnsafe {
    data: *const HistoricalTradesRejectFixedData,
}

#[repr(packed(1), C)]
pub struct HistoricalTradesRejectFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    reject_text: [u8; 96],
}

impl HistoricalTradesRejectVLSData {
    pub fn new() -> Self {
        Self {
            size: 14u16.to_le(),
            r#type: HISTORICAL_TRADES_REJECT.to_le(),
            base_size: 14u16.to_le(),
            request_id: 0i32,
            reject_text: crate::message::VLS::new(),
        }
    }
}

impl HistoricalTradesRejectFixedData {
    pub fn new() -> Self {
        Self {
            size: 104u16.to_le(),
            r#type: HISTORICAL_TRADES_REJECT.to_le(),
            request_id: 0i32,
            reject_text: [0; 96],
        }
    }
}

unsafe impl Send for HistoricalTradesRejectFixed {}
unsafe impl Send for HistoricalTradesRejectFixedUnsafe {}
unsafe impl Send for HistoricalTradesRejectFixedData {}
unsafe impl Send for HistoricalTradesRejectVLS {}
unsafe impl Send for HistoricalTradesRejectVLSUnsafe {}
unsafe impl Send for HistoricalTradesRejectVLSData {}

impl Drop for HistoricalTradesRejectFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalTradesRejectFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalTradesRejectVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalTradesRejectVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HistoricalTradesRejectFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalTradesRejectFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalTradesRejectVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalTradesRejectVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HistoricalTradesRejectFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalTradesRejectFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalTradesRejectVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalTradesRejectVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HistoricalTradesRejectFixed {
    type Target = HistoricalTradesRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalTradesRejectFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalTradesRejectFixedUnsafe {
    type Target = HistoricalTradesRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalTradesRejectFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalTradesRejectVLS {
    type Target = HistoricalTradesRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalTradesRejectVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalTradesRejectVLSUnsafe {
    type Target = HistoricalTradesRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalTradesRejectVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for HistoricalTradesRejectFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectFixed(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalTradesRejectFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectFixed(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for HistoricalTradesRejectFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectFixedUnsafe(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalTradesRejectFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectFixedUnsafe(size: {}, type: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for HistoricalTradesRejectVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectVLS(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalTradesRejectVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectVLS(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Display for HistoricalTradesRejectVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl core::fmt::Debug for HistoricalTradesRejectVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("HistoricalTradesRejectVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, reject_text: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.reject_text()).as_str())
    }
}

impl crate::Message for HistoricalTradesRejectFixed {
    type Data = HistoricalTradesRejectFixedData;

    const TYPE: u16 = HISTORICAL_TRADES_REJECT;
    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalTradesRejectFixedData::new()),
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
            data: data as *const HistoricalTradesRejectFixedData,
        }
    }
}
impl crate::Message for HistoricalTradesRejectFixedUnsafe {
    type Data = HistoricalTradesRejectFixedData;

    const TYPE: u16 = HISTORICAL_TRADES_REJECT;
    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalTradesRejectFixedData::new()),
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
            data: data as *const HistoricalTradesRejectFixedData,
        }
    }
}
impl crate::Message for HistoricalTradesRejectVLS {
    type Data = HistoricalTradesRejectVLSData;

    const TYPE: u16 = HISTORICAL_TRADES_REJECT;
    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalTradesRejectVLSData::new()),
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
            data: data as *const HistoricalTradesRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for HistoricalTradesRejectVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalTradesRejectVLSData;
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
impl crate::Message for HistoricalTradesRejectVLSUnsafe {
    type Data = HistoricalTradesRejectVLSData;

    const TYPE: u16 = HISTORICAL_TRADES_REJECT;
    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, HistoricalTradesRejectVLSData::new()),
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
            data: data as *const HistoricalTradesRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for HistoricalTradesRejectVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalTradesRejectVLSData;
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
impl HistoricalTradesReject for HistoricalTradesRejectVLS {
    type Safe = HistoricalTradesRejectVLS;
    type Unsafe = HistoricalTradesRejectVLSUnsafe;

    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    fn reject_text(&self) -> &str {
        get_vls(self, self.reject_text)
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
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

impl HistoricalTradesReject for HistoricalTradesRejectVLSUnsafe {
    type Safe = HistoricalTradesRejectVLS;
    type Unsafe = HistoricalTradesRejectVLSUnsafe;

    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(10) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            get_vls(self, self.reject_text)
        }
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_reject_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(14) {
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

impl HistoricalTradesReject for HistoricalTradesRejectFixed {
    type Safe = HistoricalTradesRejectFixed;
    type Unsafe = HistoricalTradesRejectFixedUnsafe;

    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    fn reject_text(&self) -> &str {
        get_fixed(&self.reject_text[..])
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
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

impl HistoricalTradesReject for HistoricalTradesRejectFixedUnsafe {
    type Safe = HistoricalTradesRejectFixed;
    type Unsafe = HistoricalTradesRejectFixedUnsafe;

    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(104) {
            ""
        } else {
            get_fixed(&self.reject_text[..])
        }
    }

    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
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
                core::mem::size_of::<HistoricalTradesRejectFixedData>(),
                "HistoricalTradesRejectFixedData sizeof expected {:} but was {:}",
                104usize,
                core::mem::size_of::<HistoricalTradesRejectFixedData>()
            );
            assert_eq!(
                104u16,
                HistoricalTradesRejectFixed::new().size(),
                "HistoricalTradesRejectFixed sizeof expected {:} but was {:}",
                104u16,
                HistoricalTradesRejectFixed::new().size(),
            );
            assert_eq!(
                HISTORICAL_TRADES_REJECT,
                HistoricalTradesRejectFixed::new().r#type(),
                "HistoricalTradesRejectFixed type expected {:} but was {:}",
                HISTORICAL_TRADES_REJECT,
                HistoricalTradesRejectFixed::new().r#type(),
            );
            assert_eq!(
                10101u16,
                HistoricalTradesRejectFixed::new().r#type(),
                "HistoricalTradesRejectFixed type expected {:} but was {:}",
                10101u16,
                HistoricalTradesRejectFixed::new().r#type(),
            );
            let d = HistoricalTradesRejectFixedData::new();
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
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
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
                14usize,
                core::mem::size_of::<HistoricalTradesRejectVLSData>(),
                "HistoricalTradesRejectVLSData sizeof expected {:} but was {:}",
                14usize,
                core::mem::size_of::<HistoricalTradesRejectVLSData>()
            );
            assert_eq!(
                14u16,
                HistoricalTradesRejectVLS::new().size(),
                "HistoricalTradesRejectVLS sizeof expected {:} but was {:}",
                14u16,
                HistoricalTradesRejectVLS::new().size(),
            );
            assert_eq!(
                HISTORICAL_TRADES_REJECT,
                HistoricalTradesRejectVLS::new().r#type(),
                "HistoricalTradesRejectVLS type expected {:} but was {:}",
                HISTORICAL_TRADES_REJECT,
                HistoricalTradesRejectVLS::new().r#type(),
            );
            assert_eq!(
                10101u16,
                HistoricalTradesRejectVLS::new().r#type(),
                "HistoricalTradesRejectVLS type expected {:} but was {:}",
                10101u16,
                HistoricalTradesRejectVLS::new().r#type(),
            );
            let d = HistoricalTradesRejectVLSData::new();
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
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
                "reject_text offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.reject_text) as usize) - p,
            );
        }
    }
}
