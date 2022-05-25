// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const ACCOUNT_BALANCE_REJECT_VLS_SIZE: usize = 16;

pub(crate) const ACCOUNT_BALANCE_REJECT_FIXED_SIZE: usize = 104;

/// size         u16     = AccountBalanceRejectVLSSize  (16)
/// type         u16     = ACCOUNT_BALANCE_REJECT  (602)
/// base_size    u16     = AccountBalanceRejectVLSSize  (16)
/// request_id   i32     = 0
/// reject_text  string  = ""
pub(crate) const ACCOUNT_BALANCE_REJECT_VLS_DEFAULT: [u8; 16] =
    [16, 0, 90, 2, 16, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size         u16       = AccountBalanceRejectFixedSize  (104)
/// type         u16       = ACCOUNT_BALANCE_REJECT  (602)
/// request_id   i32       = 0
/// reject_text  string96  = ""
pub(crate) const ACCOUNT_BALANCE_REJECT_FIXED_DEFAULT: [u8; 104] = [
    104, 0, 90, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait AccountBalanceReject: Message {
    type Safe: AccountBalanceReject;
    type Unsafe: AccountBalanceReject;

    /// The unique request identifier sent in the corresponding request.
    fn request_id(&self) -> i32;

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
    fn reject_text(&self) -> &str;

    /// The unique request identifier sent in the corresponding request.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
    fn set_reject_text(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl AccountBalanceReject) {
        to.set_request_id(self.request_id());
        to.set_reject_text(self.reject_text());
    }
}

pub struct AccountBalanceRejectVLS {
    data: *const AccountBalanceRejectVLSData,
    capacity: usize,
}

pub struct AccountBalanceRejectVLSUnsafe {
    data: *const AccountBalanceRejectVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct AccountBalanceRejectVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    reject_text: VLS,
}

pub struct AccountBalanceRejectFixed {
    data: *const AccountBalanceRejectFixedData,
}

pub struct AccountBalanceRejectFixedUnsafe {
    data: *const AccountBalanceRejectFixedData,
}

#[repr(packed(8), C)]
pub struct AccountBalanceRejectFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    reject_text: [u8; 96],
}

impl AccountBalanceRejectVLSData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: ACCOUNT_BALANCE_REJECT.to_le(),
            base_size: 16u16.to_le(),
            request_id: 0i32,
            reject_text: crate::message::VLS::new(),
        }
    }
}

impl AccountBalanceRejectFixedData {
    pub fn new() -> Self {
        Self {
            size: 104u16.to_le(),
            r#type: ACCOUNT_BALANCE_REJECT.to_le(),
            request_id: 0i32,
            reject_text: [0; 96],
        }
    }
}

unsafe impl Send for AccountBalanceRejectFixed {}
unsafe impl Send for AccountBalanceRejectFixedUnsafe {}
unsafe impl Send for AccountBalanceRejectFixedData {}
unsafe impl Send for AccountBalanceRejectVLS {}
unsafe impl Send for AccountBalanceRejectVLSUnsafe {}
unsafe impl Send for AccountBalanceRejectVLSData {}

impl Drop for AccountBalanceRejectFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AccountBalanceRejectFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AccountBalanceRejectVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AccountBalanceRejectVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for AccountBalanceRejectFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AccountBalanceRejectFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AccountBalanceRejectVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AccountBalanceRejectVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for AccountBalanceRejectFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AccountBalanceRejectFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AccountBalanceRejectVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AccountBalanceRejectVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for AccountBalanceRejectFixed {
    type Target = AccountBalanceRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AccountBalanceRejectFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AccountBalanceRejectFixedUnsafe {
    type Target = AccountBalanceRejectFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AccountBalanceRejectFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AccountBalanceRejectVLS {
    type Target = AccountBalanceRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AccountBalanceRejectVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AccountBalanceRejectVLSUnsafe {
    type Target = AccountBalanceRejectVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AccountBalanceRejectVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for AccountBalanceRejectFixed {
    type Data = AccountBalanceRejectFixedData;

    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AccountBalanceRejectFixedData::new()),
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
            data: data as *const AccountBalanceRejectFixedData,
        }
    }
}
impl crate::Message for AccountBalanceRejectFixedUnsafe {
    type Data = AccountBalanceRejectFixedData;

    const BASE_SIZE: usize = 104;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AccountBalanceRejectFixedData::new()),
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
            data: data as *const AccountBalanceRejectFixedData,
        }
    }
}
impl crate::Message for AccountBalanceRejectVLS {
    type Data = AccountBalanceRejectVLSData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AccountBalanceRejectVLSData::new()),
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
            data: data as *const AccountBalanceRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for AccountBalanceRejectVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const AccountBalanceRejectVLSData;
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
impl crate::Message for AccountBalanceRejectVLSUnsafe {
    type Data = AccountBalanceRejectVLSData;

    const BASE_SIZE: usize = 16;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AccountBalanceRejectVLSData::new()),
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
            data: data as *const AccountBalanceRejectVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for AccountBalanceRejectVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const AccountBalanceRejectVLSData;
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
impl AccountBalanceReject for AccountBalanceRejectVLS {
    type Safe = AccountBalanceRejectVLS;
    type Unsafe = AccountBalanceRejectVLSUnsafe;

    /// The unique request identifier sent in the corresponding request.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
    fn reject_text(&self) -> &str {
        get_vls(self, self.reject_text)
    }

    /// The unique request identifier sent in the corresponding request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
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

impl AccountBalanceReject for AccountBalanceRejectVLSUnsafe {
    type Safe = AccountBalanceRejectVLS;
    type Unsafe = AccountBalanceRejectVLSUnsafe;

    /// The unique request identifier sent in the corresponding request.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.reject_text)
        }
    }

    /// The unique request identifier sent in the corresponding request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
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

impl AccountBalanceReject for AccountBalanceRejectFixed {
    type Safe = AccountBalanceRejectFixed;
    type Unsafe = AccountBalanceRejectFixedUnsafe;

    /// The unique request identifier sent in the corresponding request.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
    fn reject_text(&self) -> &str {
        get_fixed(&self.reject_text[..])
    }

    /// The unique request identifier sent in the corresponding request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
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

impl AccountBalanceReject for AccountBalanceRejectFixedUnsafe {
    type Safe = AccountBalanceRejectFixed;
    type Unsafe = AccountBalanceRejectFixedUnsafe;

    /// The unique request identifier sent in the corresponding request.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
    fn reject_text(&self) -> &str {
        if self.is_out_of_bounds(104) {
            ""
        } else {
            get_fixed(&self.reject_text[..])
        }
    }

    /// The unique request identifier sent in the corresponding request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    /// The text reason the ACCOUNT_BALANCE_REQUEST message was rejected.
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
                core::mem::size_of::<AccountBalanceRejectFixedData>(),
                "AccountBalanceRejectFixedData sizeof expected {:} but was {:}",
                104usize,
                core::mem::size_of::<AccountBalanceRejectFixedData>()
            );
            assert_eq!(
                104u16,
                AccountBalanceRejectFixed::new().size(),
                "AccountBalanceRejectFixed sizeof expected {:} but was {:}",
                104u16,
                AccountBalanceRejectFixed::new().size(),
            );
            assert_eq!(
                ACCOUNT_BALANCE_REJECT,
                AccountBalanceRejectFixed::new().r#type(),
                "AccountBalanceRejectFixed type expected {:} but was {:}",
                ACCOUNT_BALANCE_REJECT,
                AccountBalanceRejectFixed::new().r#type(),
            );
            assert_eq!(
                602u16,
                AccountBalanceRejectFixed::new().r#type(),
                "AccountBalanceRejectFixed type expected {:} but was {:}",
                602u16,
                AccountBalanceRejectFixed::new().r#type(),
            );
            let d = AccountBalanceRejectFixedData::new();
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
                core::mem::size_of::<AccountBalanceRejectVLSData>(),
                "AccountBalanceRejectVLSData sizeof expected {:} but was {:}",
                16usize,
                core::mem::size_of::<AccountBalanceRejectVLSData>()
            );
            assert_eq!(
                16u16,
                AccountBalanceRejectVLS::new().size(),
                "AccountBalanceRejectVLS sizeof expected {:} but was {:}",
                16u16,
                AccountBalanceRejectVLS::new().size(),
            );
            assert_eq!(
                ACCOUNT_BALANCE_REJECT,
                AccountBalanceRejectVLS::new().r#type(),
                "AccountBalanceRejectVLS type expected {:} but was {:}",
                ACCOUNT_BALANCE_REJECT,
                AccountBalanceRejectVLS::new().r#type(),
            );
            assert_eq!(
                602u16,
                AccountBalanceRejectVLS::new().r#type(),
                "AccountBalanceRejectVLS type expected {:} but was {:}",
                602u16,
                AccountBalanceRejectVLS::new().r#type(),
            );
            let d = AccountBalanceRejectVLSData::new();
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
