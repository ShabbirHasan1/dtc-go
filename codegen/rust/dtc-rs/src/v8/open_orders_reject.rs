// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const OPEN_ORDERS_REJECT_VLS_SIZE: usize = 16;

pub(crate) const OPEN_ORDERS_REJECT_FIXED_SIZE: usize = 104;

/// size         u16     = OpenOrdersRejectVLSSize  (16)
/// type         u16     = OPEN_ORDERS_REJECT  (302)
/// base_size    u16     = OpenOrdersRejectVLSSize  (16)
/// request_id   i32     = 0
/// reject_text  string  = ""
pub(crate) const OPEN_ORDERS_REJECT_VLS_DEFAULT: [u8; 16] =
    [16, 0, 46, 1, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size         u16       = OpenOrdersRejectFixedSize  (104)
/// type         u16       = OPEN_ORDERS_REJECT  (302)
/// request_id   i32       = 0
/// reject_text  string96  = ""
pub(crate) const OPEN_ORDERS_REJECT_FIXED_DEFAULT: [u8; 104] = [
    104, 0, 46, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// If the Server is unable to serve the request for an OpenOrdersRequestVLS
/// message received, for a reason other than there not being any open orders,
/// then send this message to the Client.
pub trait OpenOrdersReject: Message {
    type Safe: OpenOrdersReject;
    type Unsafe: OpenOrdersReject;

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn request_id(&self) -> i32;

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str;

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// Free-form text indicating the reason for rejection.
    fn set_reject_text(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl OpenOrdersReject) {
        to.set_request_id(self.request_id());
        to.set_reject_text(self.reject_text());
    }
}

/// If the Server is unable to serve the request for an OpenOrdersRequestVLS
/// message received, for a reason other than there not being any open orders,
/// then send this message to the Client.
pub struct OpenOrdersRejectVLS {
    data: *const OpenOrdersRejectVLSData,
    capacity: usize,
}

pub struct OpenOrdersRejectVLSUnsafe {
    data: *const OpenOrdersRejectVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct OpenOrdersRejectVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    reject_text: VLS,
}

/// If the Server is unable to serve the request for an OpenOrdersRequestVLS
/// message received, for a reason other than there not being any open orders,
/// then send this message to the Client.
pub struct OpenOrdersRejectFixed {
    data: *const OpenOrdersRejectFixedData,
}

pub struct OpenOrdersRejectFixedUnsafe {
    data: *const OpenOrdersRejectFixedData,
}

#[repr(packed(8), C)]
pub struct OpenOrdersRejectFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    reject_text: [u8; 96],
}

impl OpenOrdersRejectVLSData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: OPEN_ORDERS_REJECT.to_le(),
            base_size: 16u16.to_le(),
            request_id: 0i32,
            reject_text: crate::message::VLS::new(),
        }
    }
}

impl OpenOrdersRejectFixedData {
    pub fn new() -> Self {
        Self {
            size: 104u16.to_le(),
            r#type: OPEN_ORDERS_REJECT.to_le(),
            request_id: 0i32,
            reject_text: [0; 96],
        }
    }
}

unsafe impl Send for OpenOrdersRejectFixed {}
unsafe impl Send for OpenOrdersRejectFixedUnsafe {}
unsafe impl Send for OpenOrdersRejectFixedData {}
unsafe impl Send for OpenOrdersRejectVLS {}
unsafe impl Send for OpenOrdersRejectVLSUnsafe {}
unsafe impl Send for OpenOrdersRejectVLSData {}

impl Drop for OpenOrdersRejectFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for OpenOrdersRejectFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for OpenOrdersRejectVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for OpenOrdersRejectVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for OpenOrdersRejectFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for OpenOrdersRejectFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for OpenOrdersRejectVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for OpenOrdersRejectVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for OpenOrdersRejectFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for OpenOrdersRejectFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for OpenOrdersRejectVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for OpenOrdersRejectVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for OpenOrdersRejectFixed {
    type Target = OpenOrdersRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for OpenOrdersRejectFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for OpenOrdersRejectFixedUnsafe {
    type Target = OpenOrdersRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for OpenOrdersRejectFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for OpenOrdersRejectVLS {
    type Target = OpenOrdersRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for OpenOrdersRejectVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for OpenOrdersRejectVLSUnsafe {
    type Target = OpenOrdersRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for OpenOrdersRejectVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for OpenOrdersRejectFixed {
    type Data = OpenOrdersRejectFixedData;

    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, OpenOrdersRejectFixedData::new()),
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
            data: data as *const OpenOrdersRejectFixedData,
        }
    }
}
impl crate::Message for OpenOrdersRejectFixedUnsafe {
    type Data = OpenOrdersRejectFixedData;

    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, OpenOrdersRejectFixedData::new()),
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
            data: data as *const OpenOrdersRejectFixedData,
        }
    }
}
impl crate::Message for OpenOrdersRejectVLS {
    type Data = OpenOrdersRejectVLSData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, OpenOrdersRejectVLSData::new()),
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
            data: data as *const OpenOrdersRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for OpenOrdersRejectVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const OpenOrdersRejectVLSData;
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
impl crate::Message for OpenOrdersRejectVLSUnsafe {
    type Data = OpenOrdersRejectVLSData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, OpenOrdersRejectVLSData::new()),
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
            data: data as *const OpenOrdersRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for OpenOrdersRejectVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const OpenOrdersRejectVLSData;
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
/// If the Server is unable to serve the request for an OpenOrdersRequestVLS
/// message received, for a reason other than there not being any open orders,
/// then send this message to the Client.
impl OpenOrdersReject for OpenOrdersRejectVLS {
    type Safe = OpenOrdersRejectVLS;
    type Unsafe = OpenOrdersRejectVLSUnsafe;

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        get_vls(self, self.reject_text)
    }

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// Free-form text indicating the reason for rejection.
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

/// If the Server is unable to serve the request for an OpenOrdersRequestVLS
/// message received, for a reason other than there not being any open orders,
/// then send this message to the Client.
impl OpenOrdersReject for OpenOrdersRejectVLSUnsafe {
    type Safe = OpenOrdersRejectVLS;
    type Unsafe = OpenOrdersRejectVLSUnsafe;

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.reject_text)
        }
    }

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }

    /// Free-form text indicating the reason for rejection.
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

/// If the Server is unable to serve the request for an OpenOrdersRequestVLS
/// message received, for a reason other than there not being any open orders,
/// then send this message to the Client.
impl OpenOrdersReject for OpenOrdersRejectFixed {
    type Safe = OpenOrdersRejectFixed;
    type Unsafe = OpenOrdersRejectFixedUnsafe;

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        get_fixed(&self.reject_text[..])
    }

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// Free-form text indicating the reason for rejection.
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

/// If the Server is unable to serve the request for an OpenOrdersRequestVLS
/// message received, for a reason other than there not being any open orders,
/// then send this message to the Client.
impl OpenOrdersReject for OpenOrdersRejectFixedUnsafe {
    type Safe = OpenOrdersRejectFixed;
    type Unsafe = OpenOrdersRejectFixedUnsafe;

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// Free-form text indicating the reason for rejection.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(104) {
            ""
        } else {
            get_fixed(&self.reject_text[..])
        }
    }

    /// This is set to the RequestID field sent in the OpenOrdersRequestVLS message
    /// from the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    /// Free-form text indicating the reason for rejection.
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
                core::mem::size_of::<OpenOrdersRejectFixedData>(),
                "OpenOrdersRejectFixedData sizeof expected {:} but was {:}",
                104usize,
                core::mem::size_of::<OpenOrdersRejectFixedData>()
            );
            assert_eq!(
                104u16,
                OpenOrdersRejectFixed::new().size(),
                "OpenOrdersRejectFixed sizeof expected {:} but was {:}",
                104u16,
                OpenOrdersRejectFixed::new().size(),
            );
            assert_eq!(
                OPEN_ORDERS_REJECT,
                OpenOrdersRejectFixed::new().r#type(),
                "OpenOrdersRejectFixed type expected {:} but was {:}",
                OPEN_ORDERS_REJECT,
                OpenOrdersRejectFixed::new().r#type(),
            );
            assert_eq!(
                302u16,
                OpenOrdersRejectFixed::new().r#type(),
                "OpenOrdersRejectFixed type expected {:} but was {:}",
                302u16,
                OpenOrdersRejectFixed::new().r#type(),
            );
            let d = OpenOrdersRejectFixedData::new();
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
                16usize,
                core::mem::size_of::<OpenOrdersRejectVLSData>(),
                "OpenOrdersRejectVLSData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<OpenOrdersRejectVLSData>()
            );
            assert_eq!(
                16u16,
                OpenOrdersRejectVLS::new().size(),
                "OpenOrdersRejectVLS sizeof expected {:} but was {:}",
                16u16,
                OpenOrdersRejectVLS::new().size(),
            );
            assert_eq!(
                OPEN_ORDERS_REJECT,
                OpenOrdersRejectVLS::new().r#type(),
                "OpenOrdersRejectVLS type expected {:} but was {:}",
                OPEN_ORDERS_REJECT,
                OpenOrdersRejectVLS::new().r#type(),
            );
            assert_eq!(
                302u16,
                OpenOrdersRejectVLS::new().r#type(),
                "OpenOrdersRejectVLS type expected {:} but was {:}",
                302u16,
                OpenOrdersRejectVLS::new().r#type(),
            );
            let d = OpenOrdersRejectVLSData::new();
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
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
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