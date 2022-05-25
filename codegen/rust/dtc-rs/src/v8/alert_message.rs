// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const ALERT_MESSAGE_VLS_SIZE: usize = 14;

pub(crate) const ALERT_MESSAGE_FIXED_SIZE: usize = 164;

/// size           u16     = AlertMessageVLSSize  (14)
/// type           u16     = ALERT_MESSAGE  (702)
/// base_size      u16     = AlertMessageVLSSize  (14)
/// message_text   string  = ""
/// trade_account  string  = ""
pub(crate) const ALERT_MESSAGE_VLS_DEFAULT: [u8; 14] =
    [14, 0, 190, 2, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0];

/// size           u16        = AlertMessageFixedSize  (164)
/// type           u16        = ALERT_MESSAGE  (702)
/// message_text   string128  = ""
/// trade_account  string32   = ""
pub(crate) const ALERT_MESSAGE_FIXED_DEFAULT: [u8; 164] = [
    164, 0, 190, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0,
];

pub trait AlertMessage: Message {
    type Safe: AlertMessage;
    type Unsafe: AlertMessage;

    fn message_text(&self) -> &str;

    fn trade_account(&self) -> &str;

    fn set_message_text(&mut self, value: &str) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl AlertMessage) {
        to.set_message_text(self.message_text());
        to.set_trade_account(self.trade_account());
    }
}

pub struct AlertMessageVLS {
    data: *const AlertMessageVLSData,
    capacity: usize,
}

pub struct AlertMessageVLSUnsafe {
    data: *const AlertMessageVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct AlertMessageVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    message_text: VLS,
    trade_account: VLS,
}

pub struct AlertMessageFixed {
    data: *const AlertMessageFixedData,
}

pub struct AlertMessageFixedUnsafe {
    data: *const AlertMessageFixedData,
}

#[repr(packed(8), C)]
pub struct AlertMessageFixedData {
    size: u16,
    r#type: u16,
    message_text: [u8; 128],
    trade_account: [u8; 32],
}

impl AlertMessageVLSData {
    pub fn new() -> Self {
        Self {
            size: 14u16.to_le(),
            r#type: ALERT_MESSAGE.to_le(),
            base_size: 14u16.to_le(),
            message_text: crate::message::VLS::new(),
            trade_account: crate::message::VLS::new(),
        }
    }
}

impl AlertMessageFixedData {
    pub fn new() -> Self {
        Self {
            size: 164u16.to_le(),
            r#type: ALERT_MESSAGE.to_le(),
            message_text: [0; 128],
            trade_account: [0; 32],
        }
    }
}

unsafe impl Send for AlertMessageFixed {}
unsafe impl Send for AlertMessageFixedUnsafe {}
unsafe impl Send for AlertMessageFixedData {}
unsafe impl Send for AlertMessageVLS {}
unsafe impl Send for AlertMessageVLSUnsafe {}
unsafe impl Send for AlertMessageVLSData {}

impl Drop for AlertMessageFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AlertMessageFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AlertMessageVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for AlertMessageVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for AlertMessageFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AlertMessageFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AlertMessageVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for AlertMessageVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for AlertMessageFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AlertMessageFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AlertMessageVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for AlertMessageVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for AlertMessageFixed {
    type Target = AlertMessageFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AlertMessageFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AlertMessageFixedUnsafe {
    type Target = AlertMessageFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AlertMessageFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AlertMessageVLS {
    type Target = AlertMessageVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AlertMessageVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for AlertMessageVLSUnsafe {
    type Target = AlertMessageVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for AlertMessageVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for AlertMessageFixed {
    type Data = AlertMessageFixedData;

    const BASE_SIZE: usize = 164;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AlertMessageFixedData::new()),
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
            data: data as *const AlertMessageFixedData,
        }
    }
}
impl crate::Message for AlertMessageFixedUnsafe {
    type Data = AlertMessageFixedData;

    const BASE_SIZE: usize = 164;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AlertMessageFixedData::new()),
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
            data: data as *const AlertMessageFixedData,
        }
    }
}
impl crate::Message for AlertMessageVLS {
    type Data = AlertMessageVLSData;

    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AlertMessageVLSData::new()),
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
            data: data as *const AlertMessageVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for AlertMessageVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const AlertMessageVLSData;
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
impl crate::Message for AlertMessageVLSUnsafe {
    type Data = AlertMessageVLSData;

    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, AlertMessageVLSData::new()),
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
            data: data as *const AlertMessageVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for AlertMessageVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const AlertMessageVLSData;
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
impl AlertMessage for AlertMessageVLS {
    type Safe = AlertMessageVLS;
    type Unsafe = AlertMessageVLSUnsafe;

    fn message_text(&self) -> &str {
        get_vls(self, self.message_text)
    }

    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    fn set_message_text(&mut self, value: &str) -> &mut Self {
        self.message_text = set_vls(self, self.message_text, value);
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

impl AlertMessage for AlertMessageVLSUnsafe {
    type Safe = AlertMessageVLS;
    type Unsafe = AlertMessageVLSUnsafe;

    fn message_text(&self) -> &str {
        if self.is_out_of_bounds(10) {
            ""
        } else {
            get_vls(self, self.message_text)
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    fn set_message_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.message_text = set_vls(self, self.message_text, value);
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

impl AlertMessage for AlertMessageFixed {
    type Safe = AlertMessageFixed;
    type Unsafe = AlertMessageFixedUnsafe;

    fn message_text(&self) -> &str {
        get_fixed(&self.message_text[..])
    }

    fn trade_account(&self) -> &str {
        get_fixed(&self.trade_account[..])
    }

    fn set_message_text(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.message_text[..], value);
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.trade_account[..], value);
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

impl AlertMessage for AlertMessageFixedUnsafe {
    type Safe = AlertMessageFixed;
    type Unsafe = AlertMessageFixedUnsafe;

    fn message_text(&self) -> &str {
        if self.is_out_of_bounds(132) {
            ""
        } else {
            get_fixed(&self.message_text[..])
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(164) {
            ""
        } else {
            get_fixed(&self.trade_account[..])
        }
    }

    fn set_message_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(132) {
            set_fixed(&mut self.message_text[..], value);
        }
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(164) {
            set_fixed(&mut self.trade_account[..], value);
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
                164usize,
                core::mem::size_of::<AlertMessageFixedData>(),
                "AlertMessageFixedData sizeof expected {:} but was {:}",
                164usize,
                core::mem::size_of::<AlertMessageFixedData>()
            );
            assert_eq!(
                164u16,
                AlertMessageFixed::new().size(),
                "AlertMessageFixed sizeof expected {:} but was {:}",
                164u16,
                AlertMessageFixed::new().size(),
            );
            assert_eq!(
                ALERT_MESSAGE,
                AlertMessageFixed::new().r#type(),
                "AlertMessageFixed type expected {:} but was {:}",
                ALERT_MESSAGE,
                AlertMessageFixed::new().r#type(),
            );
            assert_eq!(
                702u16,
                AlertMessageFixed::new().r#type(),
                "AlertMessageFixed type expected {:} but was {:}",
                702u16,
                AlertMessageFixed::new().r#type(),
            );
            let d = AlertMessageFixedData::new();
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
                (core::ptr::addr_of!(d.message_text) as usize) - p,
                "message_text offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.message_text) as usize) - p,
            );
            assert_eq!(
                132usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                132usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                14usize,
                core::mem::size_of::<AlertMessageVLSData>(),
                "AlertMessageVLSData sizeof expected {:} but was {:}",
                14usize,
                core::mem::size_of::<AlertMessageVLSData>()
            );
            assert_eq!(
                14u16,
                AlertMessageVLS::new().size(),
                "AlertMessageVLS sizeof expected {:} but was {:}",
                14u16,
                AlertMessageVLS::new().size(),
            );
            assert_eq!(
                ALERT_MESSAGE,
                AlertMessageVLS::new().r#type(),
                "AlertMessageVLS type expected {:} but was {:}",
                ALERT_MESSAGE,
                AlertMessageVLS::new().r#type(),
            );
            assert_eq!(
                702u16,
                AlertMessageVLS::new().r#type(),
                "AlertMessageVLS type expected {:} but was {:}",
                702u16,
                AlertMessageVLS::new().r#type(),
            );
            let d = AlertMessageVLSData::new();
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
                (core::ptr::addr_of!(d.message_text) as usize) - p,
                "message_text offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.message_text) as usize) - p,
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
