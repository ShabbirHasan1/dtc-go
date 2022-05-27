// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const LOGOFF_VLS_SIZE: usize = 12;

pub(crate) const LOGOFF_FIXED_SIZE: usize = 102;

/// size              u16     = LogoffVLSSize  (12)
/// type              u16     = LOGOFF  (5)
/// base_size         u16     = LogoffVLSSize  (12)
/// reason            string  = ""
/// do_not_reconnect  u8      = 0
pub(crate) const LOGOFF_VLS_DEFAULT: [u8; 12] = [12, 0, 5, 0, 12, 0, 0, 0, 0, 0, 0, 0];

/// size              u16       = LogoffFixedSize  (102)
/// type              u16       = LOGOFF  (5)
/// reason            string96  = ""
/// do_not_reconnect  u8        = 0
pub(crate) const LOGOFF_FIXED_DEFAULT: [u8; 102] = [
    102, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0,
];

/// A LogoffVLS is a message which can be sent either by the Client or the
/// Server to the other side. It indicates that the Client or the Server is
/// logging off and going to be closing the connection.
///
/// When one side receives this message, it should expect the connection will
/// be closed. It should not be expected that any messages will follow the
/// LogoffVLS message, and it should close the network connection and consider
/// it finished. The side receiving this message can send a LogoffVLS message
/// to the other side before closing the connection.
pub trait Logoff: Message {
    type Safe: Logoff;
    type Unsafe: Logoff;

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn reason(&self) -> &str;

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn do_not_reconnect(&self) -> u8;

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn set_reason(&mut self, value: &str) -> &mut Self;

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn set_do_not_reconnect(&mut self, value: u8) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl Logoff) {
        to.set_reason(self.reason());
        to.set_do_not_reconnect(self.do_not_reconnect());
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

/// A LogoffVLS is a message which can be sent either by the Client or the
/// Server to the other side. It indicates that the Client or the Server is
/// logging off and going to be closing the connection.
///
/// When one side receives this message, it should expect the connection will
/// be closed. It should not be expected that any messages will follow the
/// LogoffVLS message, and it should close the network connection and consider
/// it finished. The side receiving this message can send a LogoffVLS message
/// to the other side before closing the connection.
pub struct LogoffVLS {
    data: *const LogoffVLSData,
    capacity: usize,
}

pub struct LogoffVLSUnsafe {
    data: *const LogoffVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct LogoffVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    reason: VLS,
    do_not_reconnect: u8,
}

/// A LogoffVLS is a message which can be sent either by the Client or the
/// Server to the other side. It indicates that the Client or the Server is
/// logging off and going to be closing the connection.
///
/// When one side receives this message, it should expect the connection will
/// be closed. It should not be expected that any messages will follow the
/// LogoffVLS message, and it should close the network connection and consider
/// it finished. The side receiving this message can send a LogoffVLS message
/// to the other side before closing the connection.
pub struct LogoffFixed {
    data: *const LogoffFixedData,
}

pub struct LogoffFixedUnsafe {
    data: *const LogoffFixedData,
}

#[repr(packed(8), C)]
pub struct LogoffFixedData {
    size: u16,
    r#type: u16,
    reason: [u8; 96],
    do_not_reconnect: u8,
}

impl LogoffVLSData {
    pub fn new() -> Self {
        Self {
            size: 12u16.to_le(),
            r#type: LOGOFF.to_le(),
            base_size: 12u16.to_le(),
            reason: crate::message::VLS::new(),
            do_not_reconnect: 0u8,
        }
    }
}

impl LogoffFixedData {
    pub fn new() -> Self {
        Self {
            size: 102u16.to_le(),
            r#type: LOGOFF.to_le(),
            reason: [0; 96],
            do_not_reconnect: 0u8,
        }
    }
}

unsafe impl Send for LogoffFixed {}
unsafe impl Send for LogoffFixedUnsafe {}
unsafe impl Send for LogoffFixedData {}
unsafe impl Send for LogoffVLS {}
unsafe impl Send for LogoffVLSUnsafe {}
unsafe impl Send for LogoffVLSData {}

impl Drop for LogoffFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for LogoffFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for LogoffVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for LogoffVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for LogoffFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for LogoffFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for LogoffVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for LogoffVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for LogoffFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for LogoffFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for LogoffVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for LogoffVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for LogoffFixed {
    type Target = LogoffFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for LogoffFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for LogoffFixedUnsafe {
    type Target = LogoffFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for LogoffFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for LogoffVLS {
    type Target = LogoffVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for LogoffVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for LogoffVLSUnsafe {
    type Target = LogoffVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for LogoffVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for LogoffFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(
            format!(
                "LogoffFixed(size: {}, type: {}, reason: \"{}\", do_not_reconnect: {})",
                self.size(),
                self.r#type(),
                self.reason(),
                self.do_not_reconnect()
            )
            .as_str(),
        )
    }
}

impl core::fmt::Debug for LogoffFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(
            format!(
                "LogoffFixed(size: {}, type: {}, reason: \"{}\", do_not_reconnect: {})",
                self.size(),
                self.r#type(),
                self.reason(),
                self.do_not_reconnect()
            )
            .as_str(),
        )
    }
}

impl core::fmt::Display for LogoffFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(
            format!(
                "LogoffFixedUnsafe(size: {}, type: {}, reason: \"{}\", do_not_reconnect: {})",
                self.size(),
                self.r#type(),
                self.reason(),
                self.do_not_reconnect()
            )
            .as_str(),
        )
    }
}

impl core::fmt::Debug for LogoffFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(
            format!(
                "LogoffFixedUnsafe(size: {}, type: {}, reason: \"{}\", do_not_reconnect: {})",
                self.size(),
                self.r#type(),
                self.reason(),
                self.do_not_reconnect()
            )
            .as_str(),
        )
    }
}

impl core::fmt::Display for LogoffVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("LogoffVLS(size: {}, type: {}, base_size: {}, reason: \"{}\", do_not_reconnect: {})", self.size(), self.r#type(), self.base_size(), self.reason(), self.do_not_reconnect()).as_str())
    }
}

impl core::fmt::Debug for LogoffVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("LogoffVLS(size: {}, type: {}, base_size: {}, reason: \"{}\", do_not_reconnect: {})", self.size(), self.r#type(), self.base_size(), self.reason(), self.do_not_reconnect()).as_str())
    }
}

impl core::fmt::Display for LogoffVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("LogoffVLSUnsafe(size: {}, type: {}, base_size: {}, reason: \"{}\", do_not_reconnect: {})", self.size(), self.r#type(), self.base_size(), self.reason(), self.do_not_reconnect()).as_str())
    }
}

impl core::fmt::Debug for LogoffVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("LogoffVLSUnsafe(size: {}, type: {}, base_size: {}, reason: \"{}\", do_not_reconnect: {})", self.size(), self.r#type(), self.base_size(), self.reason(), self.do_not_reconnect()).as_str())
    }
}

impl crate::Message for LogoffFixed {
    type Data = LogoffFixedData;

    const TYPE: u16 = LOGOFF;
    const BASE_SIZE: usize = 102;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, LogoffFixedData::new()),
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
            data: data as *const LogoffFixedData,
        }
    }
}
impl crate::Message for LogoffFixedUnsafe {
    type Data = LogoffFixedData;

    const TYPE: u16 = LOGOFF;
    const BASE_SIZE: usize = 102;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, LogoffFixedData::new()),
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
            data: data as *const LogoffFixedData,
        }
    }
}
impl crate::Message for LogoffVLS {
    type Data = LogoffVLSData;

    const TYPE: u16 = LOGOFF;
    const BASE_SIZE: usize = 12;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, LogoffVLSData::new()),
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
            data: data as *const LogoffVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for LogoffVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const LogoffVLSData;
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
impl crate::Message for LogoffVLSUnsafe {
    type Data = LogoffVLSData;

    const TYPE: u16 = LOGOFF;
    const BASE_SIZE: usize = 12;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, LogoffVLSData::new()),
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
            data: data as *const LogoffVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for LogoffVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const LogoffVLSData;
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
/// A LogoffVLS is a message which can be sent either by the Client or the
/// Server to the other side. It indicates that the Client or the Server is
/// logging off and going to be closing the connection.
///
/// When one side receives this message, it should expect the connection will
/// be closed. It should not be expected that any messages will follow the
/// LogoffVLS message, and it should close the network connection and consider
/// it finished. The side receiving this message can send a LogoffVLS message
/// to the other side before closing the connection.
impl Logoff for LogoffVLS {
    type Safe = LogoffVLS;
    type Unsafe = LogoffVLSUnsafe;

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn reason(&self) -> &str {
        get_vls(self, self.reason)
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn do_not_reconnect(&self) -> u8 {
        self.do_not_reconnect
    }

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn set_reason(&mut self, value: &str) -> &mut Self {
        self.reason = set_vls(self, self.reason, value);
        self
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn set_do_not_reconnect(&mut self, value: u8) -> &mut Self {
        self.do_not_reconnect = value;
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

/// A LogoffVLS is a message which can be sent either by the Client or the
/// Server to the other side. It indicates that the Client or the Server is
/// logging off and going to be closing the connection.
///
/// When one side receives this message, it should expect the connection will
/// be closed. It should not be expected that any messages will follow the
/// LogoffVLS message, and it should close the network connection and consider
/// it finished. The side receiving this message can send a LogoffVLS message
/// to the other side before closing the connection.
impl Logoff for LogoffVLSUnsafe {
    type Safe = LogoffVLS;
    type Unsafe = LogoffVLSUnsafe;

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn reason(&self) -> &str {
        if self.is_out_of_bounds(10) {
            ""
        } else {
            get_vls(self, self.reason)
        }
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn do_not_reconnect(&self) -> u8 {
        if self.is_out_of_bounds(11) {
            0u8
        } else {
            self.do_not_reconnect
        }
    }

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn set_reason(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.reason = set_vls(self, self.reason, value);
        }
        self
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn set_do_not_reconnect(&mut self, value: u8) -> &mut Self {
        if !self.is_out_of_bounds(11) {
            self.do_not_reconnect = value;
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

/// A LogoffVLS is a message which can be sent either by the Client or the
/// Server to the other side. It indicates that the Client or the Server is
/// logging off and going to be closing the connection.
///
/// When one side receives this message, it should expect the connection will
/// be closed. It should not be expected that any messages will follow the
/// LogoffVLS message, and it should close the network connection and consider
/// it finished. The side receiving this message can send a LogoffVLS message
/// to the other side before closing the connection.
impl Logoff for LogoffFixed {
    type Safe = LogoffFixed;
    type Unsafe = LogoffFixedUnsafe;

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn reason(&self) -> &str {
        get_fixed(&self.reason[..])
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn do_not_reconnect(&self) -> u8 {
        self.do_not_reconnect
    }

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn set_reason(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.reason[..], value);
        self
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn set_do_not_reconnect(&mut self, value: u8) -> &mut Self {
        self.do_not_reconnect = value;
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

/// A LogoffVLS is a message which can be sent either by the Client or the
/// Server to the other side. It indicates that the Client or the Server is
/// logging off and going to be closing the connection.
///
/// When one side receives this message, it should expect the connection will
/// be closed. It should not be expected that any messages will follow the
/// LogoffVLS message, and it should close the network connection and consider
/// it finished. The side receiving this message can send a LogoffVLS message
/// to the other side before closing the connection.
impl Logoff for LogoffFixedUnsafe {
    type Safe = LogoffFixed;
    type Unsafe = LogoffFixedUnsafe;

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn reason(&self) -> &str {
        if self.is_out_of_bounds(100) {
            ""
        } else {
            get_fixed(&self.reason[..])
        }
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn do_not_reconnect(&self) -> u8 {
        if self.is_out_of_bounds(101) {
            0u8
        } else {
            self.do_not_reconnect
        }
    }

    /// Reason is a character string indicating the reason for the log off from
    /// either the Client or the Server.
    fn set_reason(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(100) {
            set_fixed(&mut self.reason[..], value);
        }
        self
    }

    /// When DoNotReconnect is set to a 1, this indicates to the other side that
    /// a reconnect to the opposite side should not occur automatically.
    fn set_do_not_reconnect(&mut self, value: u8) -> &mut Self {
        if !self.is_out_of_bounds(101) {
            self.do_not_reconnect = value;
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
                102usize,
                core::mem::size_of::<LogoffFixedData>(),
                "LogoffFixedData sizeof expected {:} but was {:}",
                102usize,
                core::mem::size_of::<LogoffFixedData>()
            );
            assert_eq!(
                102u16,
                LogoffFixed::new().size(),
                "LogoffFixed sizeof expected {:} but was {:}",
                102u16,
                LogoffFixed::new().size(),
            );
            assert_eq!(
                LOGOFF,
                LogoffFixed::new().r#type(),
                "LogoffFixed type expected {:} but was {:}",
                LOGOFF,
                LogoffFixed::new().r#type(),
            );
            assert_eq!(
                5u16,
                LogoffFixed::new().r#type(),
                "LogoffFixed type expected {:} but was {:}",
                5u16,
                LogoffFixed::new().r#type(),
            );
            let d = LogoffFixedData::new();
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
                (core::ptr::addr_of!(d.reason) as usize) - p,
                "reason offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.reason) as usize) - p,
            );
            assert_eq!(
                100usize,
                (core::ptr::addr_of!(d.do_not_reconnect) as usize) - p,
                "do_not_reconnect offset expected {:} but was {:}",
                100usize,
                (core::ptr::addr_of!(d.do_not_reconnect) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                12usize,
                core::mem::size_of::<LogoffVLSData>(),
                "LogoffVLSData sizeof expected {:} but was {:}",
                12usize,
                core::mem::size_of::<LogoffVLSData>()
            );
            assert_eq!(
                12u16,
                LogoffVLS::new().size(),
                "LogoffVLS sizeof expected {:} but was {:}",
                12u16,
                LogoffVLS::new().size(),
            );
            assert_eq!(
                LOGOFF,
                LogoffVLS::new().r#type(),
                "LogoffVLS type expected {:} but was {:}",
                LOGOFF,
                LogoffVLS::new().r#type(),
            );
            assert_eq!(
                5u16,
                LogoffVLS::new().r#type(),
                "LogoffVLS type expected {:} but was {:}",
                5u16,
                LogoffVLS::new().r#type(),
            );
            let d = LogoffVLSData::new();
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
                (core::ptr::addr_of!(d.reason) as usize) - p,
                "reason offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.reason) as usize) - p,
            );
            assert_eq!(
                10usize,
                (core::ptr::addr_of!(d.do_not_reconnect) as usize) - p,
                "do_not_reconnect offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.do_not_reconnect) as usize) - p,
            );
        }
    }
}
