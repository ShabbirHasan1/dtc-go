// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const EXCHANGE_LIST_RESPONSE_VLS_SIZE: usize = 24;

pub(crate) const EXCHANGE_LIST_RESPONSE_FIXED_SIZE: usize = 76;

/// size              u16     = ExchangeListResponseVLSSize  (24)
/// type              u16     = EXCHANGE_LIST_RESPONSE  (501)
/// base_size         u16     = ExchangeListResponseVLSSize  (24)
/// request_id        i32     = 0
/// exchange          string  = ""
/// is_final_message  bool    = false
/// description       string  = ""
pub(crate) const EXCHANGE_LIST_RESPONSE_VLS_DEFAULT: [u8; 24] = [
    24, 0, 245, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// size              u16       = ExchangeListResponseFixedSize  (76)
/// type              u16       = EXCHANGE_LIST_RESPONSE  (501)
/// request_id        i32       = 0
/// exchange          string16  = ""
/// is_final_message  bool      = false
/// description       string48  = ""
pub(crate) const EXCHANGE_LIST_RESPONSE_FIXED_DEFAULT: [u8; 76] = [
    76, 0, 245, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// The server will return this message for each supported exchange.
///
/// If there are no exchanges to return in response to a request, send through
/// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
/// all other members in the default state and the Client will recognize there
/// are no Exchanges.
pub trait ExchangeListResponse: Message {
    type Safe: ExchangeListResponse;
    type Unsafe: ExchangeListResponse;

    /// The RequestID sent in the request by the Client.
    fn request_id(&self) -> i32;

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn exchange(&self) -> &str;

    /// 1 = final message in batch.
    fn is_final_message(&self) -> bool;

    /// The complete exchange description.
    fn description(&self) -> &str;

    /// The RequestID sent in the request by the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn set_exchange(&mut self, value: &str) -> &mut Self;

    /// 1 = final message in batch.
    fn set_is_final_message(&mut self, value: bool) -> &mut Self;

    /// The complete exchange description.
    fn set_description(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl ExchangeListResponse) {
        to.set_request_id(self.request_id());
        to.set_exchange(self.exchange());
        to.set_is_final_message(self.is_final_message());
        to.set_description(self.description());
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

/// The server will return this message for each supported exchange.
///
/// If there are no exchanges to return in response to a request, send through
/// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
/// all other members in the default state and the Client will recognize there
/// are no Exchanges.
pub struct ExchangeListResponseVLS {
    data: *const ExchangeListResponseVLSData,
    capacity: usize,
}

pub struct ExchangeListResponseVLSUnsafe {
    data: *const ExchangeListResponseVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct ExchangeListResponseVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    exchange: VLS,
    is_final_message: bool,
    description: VLS,
}

/// The server will return this message for each supported exchange.
///
/// If there are no exchanges to return in response to a request, send through
/// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
/// all other members in the default state and the Client will recognize there
/// are no Exchanges.
pub struct ExchangeListResponseFixed {
    data: *const ExchangeListResponseFixedData,
}

pub struct ExchangeListResponseFixedUnsafe {
    data: *const ExchangeListResponseFixedData,
}

#[repr(packed(8), C)]
pub struct ExchangeListResponseFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    exchange: [u8; 16],
    is_final_message: bool,
    description: [u8; 48],
}

impl ExchangeListResponseVLSData {
    pub fn new() -> Self {
        Self {
            size: 24u16.to_le(),
            r#type: EXCHANGE_LIST_RESPONSE.to_le(),
            base_size: 24u16.to_le(),
            request_id: 0i32,
            exchange: crate::message::VLS::new(),
            is_final_message: false,
            description: crate::message::VLS::new(),
        }
    }
}

impl ExchangeListResponseFixedData {
    pub fn new() -> Self {
        Self {
            size: 76u16.to_le(),
            r#type: EXCHANGE_LIST_RESPONSE.to_le(),
            request_id: 0i32,
            exchange: [0; 16],
            is_final_message: false,
            description: [0; 48],
        }
    }
}

unsafe impl Send for ExchangeListResponseFixed {}
unsafe impl Send for ExchangeListResponseFixedUnsafe {}
unsafe impl Send for ExchangeListResponseFixedData {}
unsafe impl Send for ExchangeListResponseVLS {}
unsafe impl Send for ExchangeListResponseVLSUnsafe {}
unsafe impl Send for ExchangeListResponseVLSData {}

impl Drop for ExchangeListResponseFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for ExchangeListResponseFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for ExchangeListResponseVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for ExchangeListResponseVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for ExchangeListResponseFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for ExchangeListResponseFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for ExchangeListResponseVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for ExchangeListResponseVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for ExchangeListResponseFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for ExchangeListResponseFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for ExchangeListResponseVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for ExchangeListResponseVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for ExchangeListResponseFixed {
    type Target = ExchangeListResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ExchangeListResponseFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for ExchangeListResponseFixedUnsafe {
    type Target = ExchangeListResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ExchangeListResponseFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for ExchangeListResponseVLS {
    type Target = ExchangeListResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ExchangeListResponseVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for ExchangeListResponseVLSUnsafe {
    type Target = ExchangeListResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for ExchangeListResponseVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for ExchangeListResponseFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseFixed(size: {}, type: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl core::fmt::Debug for ExchangeListResponseFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseFixed(size: {}, type: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl core::fmt::Display for ExchangeListResponseFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseFixedUnsafe(size: {}, type: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl core::fmt::Debug for ExchangeListResponseFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseFixedUnsafe(size: {}, type: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl core::fmt::Display for ExchangeListResponseVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseVLS(size: {}, type: {}, base_size: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl core::fmt::Debug for ExchangeListResponseVLS {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseVLS(size: {}, type: {}, base_size: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl core::fmt::Display for ExchangeListResponseVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl core::fmt::Debug for ExchangeListResponseVLSUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("ExchangeListResponseVLSUnsafe(size: {}, type: {}, base_size: {}, request_id: {}, exchange: \"{}\", is_final_message: {}, description: \"{}\")", self.size(), self.r#type(), self.base_size(), self.request_id(), self.exchange(), self.is_final_message(), self.description()).as_str())
    }
}

impl crate::Message for ExchangeListResponseFixed {
    type Data = ExchangeListResponseFixedData;

    const TYPE: u16 = EXCHANGE_LIST_RESPONSE;
    const BASE_SIZE: usize = 76;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ExchangeListResponseFixedData::new()),
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
            data: data as *const ExchangeListResponseFixedData,
        }
    }
}
impl crate::Message for ExchangeListResponseFixedUnsafe {
    type Data = ExchangeListResponseFixedData;

    const TYPE: u16 = EXCHANGE_LIST_RESPONSE;
    const BASE_SIZE: usize = 76;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ExchangeListResponseFixedData::new()),
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
            data: data as *const ExchangeListResponseFixedData,
        }
    }
}
impl crate::Message for ExchangeListResponseVLS {
    type Data = ExchangeListResponseVLSData;

    const TYPE: u16 = EXCHANGE_LIST_RESPONSE;
    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ExchangeListResponseVLSData::new()),
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
            data: data as *const ExchangeListResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for ExchangeListResponseVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const ExchangeListResponseVLSData;
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
impl crate::Message for ExchangeListResponseVLSUnsafe {
    type Data = ExchangeListResponseVLSData;

    const TYPE: u16 = EXCHANGE_LIST_RESPONSE;
    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, ExchangeListResponseVLSData::new()),
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
            data: data as *const ExchangeListResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for ExchangeListResponseVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const ExchangeListResponseVLSData;
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
/// The server will return this message for each supported exchange.
///
/// If there are no exchanges to return in response to a request, send through
/// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
/// all other members in the default state and the Client will recognize there
/// are no Exchanges.
impl ExchangeListResponse for ExchangeListResponseVLS {
    type Safe = ExchangeListResponseVLS;
    type Unsafe = ExchangeListResponseVLSUnsafe;

    /// The RequestID sent in the request by the Client.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn exchange(&self) -> &str {
        get_vls(self, self.exchange)
    }

    /// 1 = final message in batch.
    fn is_final_message(&self) -> bool {
        self.is_final_message
    }

    /// The complete exchange description.
    fn description(&self) -> &str {
        get_vls(self, self.description)
    }

    /// The RequestID sent in the request by the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        self.exchange = set_vls(self, self.exchange, value);
        self
    }

    /// 1 = final message in batch.
    fn set_is_final_message(&mut self, value: bool) -> &mut Self {
        self.is_final_message = value;
        self
    }

    /// The complete exchange description.
    fn set_description(&mut self, value: &str) -> &mut Self {
        self.description = set_vls(self, self.description, value);
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

/// The server will return this message for each supported exchange.
///
/// If there are no exchanges to return in response to a request, send through
/// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
/// all other members in the default state and the Client will recognize there
/// are no Exchanges.
impl ExchangeListResponse for ExchangeListResponseVLSUnsafe {
    type Safe = ExchangeListResponseVLS;
    type Unsafe = ExchangeListResponseVLSUnsafe;

    /// The RequestID sent in the request by the Client.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.exchange)
        }
    }

    /// 1 = final message in batch.
    fn is_final_message(&self) -> bool {
        if self.is_out_of_bounds(17) {
            false
        } else {
            self.is_final_message
        }
    }

    /// The complete exchange description.
    fn description(&self) -> &str {
        if self.is_out_of_bounds(22) {
            ""
        } else {
            get_vls(self, self.description)
        }
    }

    /// The RequestID sent in the request by the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.exchange = set_vls(self, self.exchange, value);
        }
        self
    }

    /// 1 = final message in batch.
    fn set_is_final_message(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(17) {
            self.is_final_message = value;
        }
        self
    }

    /// The complete exchange description.
    fn set_description(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(22) {
            self.description = set_vls(self, self.description, value);
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

/// The server will return this message for each supported exchange.
///
/// If there are no exchanges to return in response to a request, send through
/// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
/// all other members in the default state and the Client will recognize there
/// are no Exchanges.
impl ExchangeListResponse for ExchangeListResponseFixed {
    type Safe = ExchangeListResponseFixed;
    type Unsafe = ExchangeListResponseFixedUnsafe;

    /// The RequestID sent in the request by the Client.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn exchange(&self) -> &str {
        get_fixed(&self.exchange[..])
    }

    /// 1 = final message in batch.
    fn is_final_message(&self) -> bool {
        self.is_final_message
    }

    /// The complete exchange description.
    fn description(&self) -> &str {
        get_fixed(&self.description[..])
    }

    /// The RequestID sent in the request by the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.exchange[..], value);
        self
    }

    /// 1 = final message in batch.
    fn set_is_final_message(&mut self, value: bool) -> &mut Self {
        self.is_final_message = value;
        self
    }

    /// The complete exchange description.
    fn set_description(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.description[..], value);
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

/// The server will return this message for each supported exchange.
///
/// If there are no exchanges to return in response to a request, send through
/// one of these messages with the RequestID set and IsFinalMessage = 1. Leave
/// all other members in the default state and the Client will recognize there
/// are no Exchanges.
impl ExchangeListResponse for ExchangeListResponseFixedUnsafe {
    type Safe = ExchangeListResponseFixed;
    type Unsafe = ExchangeListResponseFixedUnsafe;

    /// The RequestID sent in the request by the Client.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(24) {
            ""
        } else {
            get_fixed(&self.exchange[..])
        }
    }

    /// 1 = final message in batch.
    fn is_final_message(&self) -> bool {
        if self.is_out_of_bounds(25) {
            false
        } else {
            self.is_final_message
        }
    }

    /// The complete exchange description.
    fn description(&self) -> &str {
        if self.is_out_of_bounds(73) {
            ""
        } else {
            get_fixed(&self.description[..])
        }
    }

    /// The RequestID sent in the request by the Client.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    /// The exchange identifier that the Server uses to identify a particular
    /// exchange.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            set_fixed(&mut self.exchange[..], value);
        }
        self
    }

    /// 1 = final message in batch.
    fn set_is_final_message(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(25) {
            self.is_final_message = value;
        }
        self
    }

    /// The complete exchange description.
    fn set_description(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(73) {
            set_fixed(&mut self.description[..], value);
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
                76usize,
                core::mem::size_of::<ExchangeListResponseFixedData>(),
                "ExchangeListResponseFixedData sizeof expected {:} but was {:}",
                76usize,
                core::mem::size_of::<ExchangeListResponseFixedData>()
            );
            assert_eq!(
                76u16,
                ExchangeListResponseFixed::new().size(),
                "ExchangeListResponseFixed sizeof expected {:} but was {:}",
                76u16,
                ExchangeListResponseFixed::new().size(),
            );
            assert_eq!(
                EXCHANGE_LIST_RESPONSE,
                ExchangeListResponseFixed::new().r#type(),
                "ExchangeListResponseFixed type expected {:} but was {:}",
                EXCHANGE_LIST_RESPONSE,
                ExchangeListResponseFixed::new().r#type(),
            );
            assert_eq!(
                501u16,
                ExchangeListResponseFixed::new().r#type(),
                "ExchangeListResponseFixed type expected {:} but was {:}",
                501u16,
                ExchangeListResponseFixed::new().r#type(),
            );
            let d = ExchangeListResponseFixedData::new();
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
                (core::ptr::addr_of!(d.exchange) as usize) - p,
                "exchange offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
            );
            assert_eq!(
                24usize,
                (core::ptr::addr_of!(d.is_final_message) as usize) - p,
                "is_final_message offset expected {:} but was {:}",
                24usize,
                (core::ptr::addr_of!(d.is_final_message) as usize) - p,
            );
            assert_eq!(
                25usize,
                (core::ptr::addr_of!(d.description) as usize) - p,
                "description offset expected {:} but was {:}",
                25usize,
                (core::ptr::addr_of!(d.description) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                24usize,
                core::mem::size_of::<ExchangeListResponseVLSData>(),
                "ExchangeListResponseVLSData sizeof expected {:} but was {:}",
                24usize,
                core::mem::size_of::<ExchangeListResponseVLSData>()
            );
            assert_eq!(
                24u16,
                ExchangeListResponseVLS::new().size(),
                "ExchangeListResponseVLS sizeof expected {:} but was {:}",
                24u16,
                ExchangeListResponseVLS::new().size(),
            );
            assert_eq!(
                EXCHANGE_LIST_RESPONSE,
                ExchangeListResponseVLS::new().r#type(),
                "ExchangeListResponseVLS type expected {:} but was {:}",
                EXCHANGE_LIST_RESPONSE,
                ExchangeListResponseVLS::new().r#type(),
            );
            assert_eq!(
                501u16,
                ExchangeListResponseVLS::new().r#type(),
                "ExchangeListResponseVLS type expected {:} but was {:}",
                501u16,
                ExchangeListResponseVLS::new().r#type(),
            );
            let d = ExchangeListResponseVLSData::new();
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
                (core::ptr::addr_of!(d.exchange) as usize) - p,
                "exchange offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.is_final_message) as usize) - p,
                "is_final_message offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.is_final_message) as usize) - p,
            );
            assert_eq!(
                18usize,
                (core::ptr::addr_of!(d.description) as usize) - p,
                "description offset expected {:} but was {:}",
                18usize,
                (core::ptr::addr_of!(d.description) as usize) - p,
            );
        }
    }
}
