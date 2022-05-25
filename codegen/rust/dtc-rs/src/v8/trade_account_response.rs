// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const TRADE_ACCOUNT_RESPONSE_VLS_SIZE: usize = 24;

pub(crate) const TRADE_ACCOUNT_RESPONSE_FIXED_SIZE: usize = 48;

/// size                   u16     = TradeAccountResponseVLSSize  (24)
/// type                   u16     = TRADE_ACCOUNT_RESPONSE  (401)
/// base_size              u16     = TradeAccountResponseVLSSize  (24)
/// total_number_messages  i32     = 0
/// message_number         i32     = 0
/// trade_account          string  = ""
/// request_id             i32     = 0
pub(crate) const TRADE_ACCOUNT_RESPONSE_VLS_DEFAULT: [u8; 24] = [
    24, 0, 145, 1, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// size                   u16       = TradeAccountResponseFixedSize  (48)
/// type                   u16       = TRADE_ACCOUNT_RESPONSE  (401)
/// total_number_messages  i32       = 0
/// message_number         i32       = 0
/// trade_account          string32  = ""
/// request_id             i32       = 0
pub(crate) const TRADE_ACCOUNT_RESPONSE_FIXED_DEFAULT: [u8; 48] = [
    48, 0, 145, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// This is a message from the Server to the Client in response to a TradeAccountsRequestFixed
/// message, providing a single trade account. There is one message for each
/// trade account.
pub trait TradeAccountResponse: Message {
    type Safe: TradeAccountResponse;
    type Unsafe: TradeAccountResponse;

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn total_number_messages(&self) -> i32;

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn message_number(&self) -> i32;

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn trade_account(&self) -> &str;

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn request_id(&self) -> i32;

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_total_number_messages(&mut self, value: i32) -> &mut Self;

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_message_number(&mut self, value: i32) -> &mut Self;

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountResponse) {
        to.set_total_number_messages(self.total_number_messages());
        to.set_message_number(self.message_number());
        to.set_trade_account(self.trade_account());
        to.set_request_id(self.request_id());
    }
}

/// This is a message from the Server to the Client in response to a TradeAccountsRequestFixed
/// message, providing a single trade account. There is one message for each
/// trade account.
pub struct TradeAccountResponseVLS {
    data: *const TradeAccountResponseVLSData,
    capacity: usize,
}

pub struct TradeAccountResponseVLSUnsafe {
    data: *const TradeAccountResponseVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct TradeAccountResponseVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    total_number_messages: i32,
    message_number: i32,
    trade_account: VLS,
    request_id: i32,
}

/// This is a message from the Server to the Client in response to a TradeAccountsRequestFixed
/// message, providing a single trade account. There is one message for each
/// trade account.
pub struct TradeAccountResponseFixed {
    data: *const TradeAccountResponseFixedData,
}

pub struct TradeAccountResponseFixedUnsafe {
    data: *const TradeAccountResponseFixedData,
}

#[repr(packed(8), C)]
pub struct TradeAccountResponseFixedData {
    size: u16,
    r#type: u16,
    total_number_messages: i32,
    message_number: i32,
    trade_account: [u8; 32],
    request_id: i32,
}

impl TradeAccountResponseVLSData {
    pub fn new() -> Self {
        Self {
            size: 24u16.to_le(),
            r#type: TRADE_ACCOUNT_RESPONSE.to_le(),
            base_size: 24u16.to_le(),
            total_number_messages: 0i32,
            message_number: 0i32,
            trade_account: crate::message::VLS::new(),
            request_id: 0i32,
        }
    }
}

impl TradeAccountResponseFixedData {
    pub fn new() -> Self {
        Self {
            size: 48u16.to_le(),
            r#type: TRADE_ACCOUNT_RESPONSE.to_le(),
            total_number_messages: 0i32,
            message_number: 0i32,
            trade_account: [0; 32],
            request_id: 0i32,
        }
    }
}

unsafe impl Send for TradeAccountResponseFixed {}
unsafe impl Send for TradeAccountResponseFixedUnsafe {}
unsafe impl Send for TradeAccountResponseFixedData {}
unsafe impl Send for TradeAccountResponseVLS {}
unsafe impl Send for TradeAccountResponseVLSUnsafe {}
unsafe impl Send for TradeAccountResponseVLSData {}

impl Drop for TradeAccountResponseFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountResponseFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountResponseVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountResponseVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountResponseFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountResponseFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountResponseVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountResponseVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountResponseFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountResponseFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountResponseVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountResponseVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountResponseFixed {
    type Target = TradeAccountResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountResponseFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountResponseFixedUnsafe {
    type Target = TradeAccountResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountResponseFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountResponseVLS {
    type Target = TradeAccountResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountResponseVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountResponseVLSUnsafe {
    type Target = TradeAccountResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountResponseVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for TradeAccountResponseFixed {
    type Data = TradeAccountResponseFixedData;

    const BASE_SIZE: usize = 48;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountResponseFixedData::new()),
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
            data: data as *const TradeAccountResponseFixedData,
        }
    }
}
impl crate::Message for TradeAccountResponseFixedUnsafe {
    type Data = TradeAccountResponseFixedData;

    const BASE_SIZE: usize = 48;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountResponseFixedData::new()),
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
            data: data as *const TradeAccountResponseFixedData,
        }
    }
}
impl crate::Message for TradeAccountResponseVLS {
    type Data = TradeAccountResponseVLSData;

    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountResponseVLSData::new()),
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
            data: data as *const TradeAccountResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountResponseVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountResponseVLSData;
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
impl crate::Message for TradeAccountResponseVLSUnsafe {
    type Data = TradeAccountResponseVLSData;

    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, TradeAccountResponseVLSData::new()),
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
            data: data as *const TradeAccountResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountResponseVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountResponseVLSData;
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
/// This is a message from the Server to the Client in response to a TradeAccountsRequestFixed
/// message, providing a single trade account. There is one message for each
/// trade account.
impl TradeAccountResponse for TradeAccountResponseVLS {
    type Safe = TradeAccountResponseVLS;
    type Unsafe = TradeAccountResponseVLSUnsafe;

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn total_number_messages(&self) -> i32 {
        i32::from_le(self.total_number_messages)
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn message_number(&self) -> i32 {
        i32::from_le(self.message_number)
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_total_number_messages(&mut self, value: i32) -> &mut Self {
        self.total_number_messages = value.to_le();
        self
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_message_number(&mut self, value: i32) -> &mut Self {
        self.message_number = value.to_le();
        self
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = set_vls(self, self.trade_account, value);
        self
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
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

/// This is a message from the Server to the Client in response to a TradeAccountsRequestFixed
/// message, providing a single trade account. There is one message for each
/// trade account.
impl TradeAccountResponse for TradeAccountResponseVLSUnsafe {
    type Safe = TradeAccountResponseVLS;
    type Unsafe = TradeAccountResponseVLSUnsafe;

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn total_number_messages(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.total_number_messages)
        }
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn message_number(&self) -> i32 {
        if self.is_out_of_bounds(16) {
            0i32
        } else {
            i32::from_le(self.message_number)
        }
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(20) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(24) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_total_number_messages(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.total_number_messages = value.to_le();
        }
        self
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_message_number(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.message_number = value.to_le();
        }
        self
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(20) {
            self.trade_account = set_vls(self, self.trade_account, value);
        }
        self
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.request_id = value.to_le();
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

/// This is a message from the Server to the Client in response to a TradeAccountsRequestFixed
/// message, providing a single trade account. There is one message for each
/// trade account.
impl TradeAccountResponse for TradeAccountResponseFixed {
    type Safe = TradeAccountResponseFixed;
    type Unsafe = TradeAccountResponseFixedUnsafe;

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn total_number_messages(&self) -> i32 {
        i32::from_le(self.total_number_messages)
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn message_number(&self) -> i32 {
        i32::from_le(self.message_number)
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn trade_account(&self) -> &str {
        get_fixed(&self.trade_account[..])
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_total_number_messages(&mut self, value: i32) -> &mut Self {
        self.total_number_messages = value.to_le();
        self
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_message_number(&mut self, value: i32) -> &mut Self {
        self.message_number = value.to_le();
        self
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.trade_account[..], value);
        self
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
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

/// This is a message from the Server to the Client in response to a TradeAccountsRequestFixed
/// message, providing a single trade account. There is one message for each
/// trade account.
impl TradeAccountResponse for TradeAccountResponseFixedUnsafe {
    type Safe = TradeAccountResponseFixed;
    type Unsafe = TradeAccountResponseFixedUnsafe;

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn total_number_messages(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.total_number_messages)
        }
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn message_number(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.message_number)
        }
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(44) {
            ""
        } else {
            get_fixed(&self.trade_account[..])
        }
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(48) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// This indicates the total number of Account list messages when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_total_number_messages(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.total_number_messages = value.to_le();
        }
        self
    }

    /// This indicates the 1-based index of the Account list message when a batch
    /// of messages is being sent. If there is only one message being sent, this
    /// will be 1.
    fn set_message_number(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.message_number = value.to_le();
        }
        self
    }

    /// The trade account identifier.
    ///
    /// In the case when there are no Trade Accounts available for the logged
    /// in Username, the Server will send a single TradeAccountResponseVLS message
    /// to the Client and leave this field empty.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(44) {
            set_fixed(&mut self.trade_account[..], value);
        }
        self
    }

    /// The RequestID sent in the corresponding request by the Client.
    ///
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(48) {
            self.request_id = value.to_le();
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
                48usize,
                core::mem::size_of::<TradeAccountResponseFixedData>(),
                "TradeAccountResponseFixedData sizeof expected {:} but was {:}",
                48usize,
                core::mem::size_of::<TradeAccountResponseFixedData>()
            );
            assert_eq!(
                48u16,
                TradeAccountResponseFixed::new().size(),
                "TradeAccountResponseFixed sizeof expected {:} but was {:}",
                48u16,
                TradeAccountResponseFixed::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_RESPONSE,
                TradeAccountResponseFixed::new().r#type(),
                "TradeAccountResponseFixed type expected {:} but was {:}",
                TRADE_ACCOUNT_RESPONSE,
                TradeAccountResponseFixed::new().r#type(),
            );
            assert_eq!(
                401u16,
                TradeAccountResponseFixed::new().r#type(),
                "TradeAccountResponseFixed type expected {:} but was {:}",
                401u16,
                TradeAccountResponseFixed::new().r#type(),
            );
            let d = TradeAccountResponseFixedData::new();
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
                (core::ptr::addr_of!(d.total_number_messages) as usize) - p,
                "total_number_messages offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.total_number_messages) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.message_number) as usize) - p,
                "message_number offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.message_number) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                44usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                44usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                24usize,
                core::mem::size_of::<TradeAccountResponseVLSData>(),
                "TradeAccountResponseVLSData sizeof expected {:} but was {:}",
                24usize,
                core::mem::size_of::<TradeAccountResponseVLSData>()
            );
            assert_eq!(
                24u16,
                TradeAccountResponseVLS::new().size(),
                "TradeAccountResponseVLS sizeof expected {:} but was {:}",
                24u16,
                TradeAccountResponseVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_RESPONSE,
                TradeAccountResponseVLS::new().r#type(),
                "TradeAccountResponseVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_RESPONSE,
                TradeAccountResponseVLS::new().r#type(),
            );
            assert_eq!(
                401u16,
                TradeAccountResponseVLS::new().r#type(),
                "TradeAccountResponseVLS type expected {:} but was {:}",
                401u16,
                TradeAccountResponseVLS::new().r#type(),
            );
            let d = TradeAccountResponseVLSData::new();
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
                (core::ptr::addr_of!(d.total_number_messages) as usize) - p,
                "total_number_messages offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.total_number_messages) as usize) - p,
            );
            assert_eq!(
                12usize,
                (core::ptr::addr_of!(d.message_number) as usize) - p,
                "message_number offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.message_number) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                20usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                20usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
        }
    }
}
