// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const HISTORICAL_ACCOUNT_BALANCES_REQUEST_VLS_SIZE: usize = 24;

pub(crate) const HISTORICAL_ACCOUNT_BALANCES_REQUEST_FIXED_SIZE: usize = 48;

/// size             u16       = HistoricalAccountBalancesRequestVLSSize  (24)
/// type             u16       = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
/// base_size        u16       = HistoricalAccountBalancesRequestVLSSize  (24)
/// request_id       i32       = 0
/// trade_account    string    = ""
/// start_date_time  DateTime  = 0
pub(crate) const HISTORICAL_ACCOUNT_BALANCES_REQUEST_VLS_DEFAULT: [u8; 24] = [
    24, 0, 91, 2, 24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// size             u16       = HistoricalAccountBalancesRequestFixedSize  (48)
/// type             u16       = HISTORICAL_ACCOUNT_BALANCES_REQUEST  (603)
/// request_id       i32       = 0
/// trade_account    string32  = ""
/// start_date_time  DateTime  = 0
pub(crate) const HISTORICAL_ACCOUNT_BALANCES_REQUEST_FIXED_DEFAULT: [u8; 48] = [
    48, 0, 91, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// This is a message from the Client to the Server to request a history of
/// Cash Balance changes for the specified Trade Account.
///
/// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// messages or reject he request with a message.
///
/// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
pub trait HistoricalAccountBalancesRequest: Message {
    type Safe: HistoricalAccountBalancesRequest;
    type Unsafe: HistoricalAccountBalancesRequest;

    /// A unique request identifier for this request.
    fn request_id(&self) -> i32;

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn trade_account(&self) -> &str;

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn start_date_time(&self) -> DateTime;

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self;

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl HistoricalAccountBalancesRequest) {
        to.set_request_id(self.request_id());
        to.set_trade_account(self.trade_account());
        to.set_start_date_time(self.start_date_time());
    }
}

/// This is a message from the Client to the Server to request a history of
/// Cash Balance changes for the specified Trade Account.
///
/// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// messages or reject he request with a message.
///
/// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
pub struct HistoricalAccountBalancesRequestVLS {
    data: *const HistoricalAccountBalancesRequestVLSData,
    capacity: usize,
}

pub struct HistoricalAccountBalancesRequestVLSUnsafe {
    data: *const HistoricalAccountBalancesRequestVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct HistoricalAccountBalancesRequestVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: i32,
    trade_account: VLS,
    start_date_time: DateTime,
}

/// This is a message from the Client to the Server to request a history of
/// Cash Balance changes for the specified Trade Account.
///
/// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// messages or reject he request with a message.
///
/// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
pub struct HistoricalAccountBalancesRequestFixed {
    data: *const HistoricalAccountBalancesRequestFixedData,
}

pub struct HistoricalAccountBalancesRequestFixedUnsafe {
    data: *const HistoricalAccountBalancesRequestFixedData,
}

#[repr(packed(8), C)]
pub struct HistoricalAccountBalancesRequestFixedData {
    size: u16,
    r#type: u16,
    request_id: i32,
    trade_account: [u8; 32],
    start_date_time: DateTime,
}

impl HistoricalAccountBalancesRequestVLSData {
    pub fn new() -> Self {
        Self {
            size: 24u16.to_le(),
            r#type: HISTORICAL_ACCOUNT_BALANCES_REQUEST.to_le(),
            base_size: 24u16.to_le(),
            request_id: 0i32,
            trade_account: crate::message::VLS::new(),
            start_date_time: 0i64,
        }
    }
}

impl HistoricalAccountBalancesRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 48u16.to_le(),
            r#type: HISTORICAL_ACCOUNT_BALANCES_REQUEST.to_le(),
            request_id: 0i32,
            trade_account: [0; 32],
            start_date_time: 0i64,
        }
    }
}

unsafe impl Send for HistoricalAccountBalancesRequestFixed {}
unsafe impl Send for HistoricalAccountBalancesRequestFixedUnsafe {}
unsafe impl Send for HistoricalAccountBalancesRequestFixedData {}
unsafe impl Send for HistoricalAccountBalancesRequestVLS {}
unsafe impl Send for HistoricalAccountBalancesRequestVLSUnsafe {}
unsafe impl Send for HistoricalAccountBalancesRequestVLSData {}

impl Drop for HistoricalAccountBalancesRequestFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalAccountBalancesRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalAccountBalancesRequestVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for HistoricalAccountBalancesRequestVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for HistoricalAccountBalancesRequestFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalAccountBalancesRequestFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalAccountBalancesRequestVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for HistoricalAccountBalancesRequestVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for HistoricalAccountBalancesRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalAccountBalancesRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalAccountBalancesRequestVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for HistoricalAccountBalancesRequestVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for HistoricalAccountBalancesRequestFixed {
    type Target = HistoricalAccountBalancesRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalAccountBalancesRequestFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalAccountBalancesRequestFixedUnsafe {
    type Target = HistoricalAccountBalancesRequestFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalAccountBalancesRequestFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalAccountBalancesRequestVLS {
    type Target = HistoricalAccountBalancesRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalAccountBalancesRequestVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for HistoricalAccountBalancesRequestVLSUnsafe {
    type Target = HistoricalAccountBalancesRequestVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for HistoricalAccountBalancesRequestVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for HistoricalAccountBalancesRequestFixed {
    type Data = HistoricalAccountBalancesRequestFixedData;

    const BASE_SIZE: usize = 48;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                HistoricalAccountBalancesRequestFixedData::new(),
            ),
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
            data: data as *const HistoricalAccountBalancesRequestFixedData,
        }
    }
}
impl crate::Message for HistoricalAccountBalancesRequestFixedUnsafe {
    type Data = HistoricalAccountBalancesRequestFixedData;

    const BASE_SIZE: usize = 48;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                HistoricalAccountBalancesRequestFixedData::new(),
            ),
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
            data: data as *const HistoricalAccountBalancesRequestFixedData,
        }
    }
}
impl crate::Message for HistoricalAccountBalancesRequestVLS {
    type Data = HistoricalAccountBalancesRequestVLSData;

    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                HistoricalAccountBalancesRequestVLSData::new(),
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
            data: data as *const HistoricalAccountBalancesRequestVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for HistoricalAccountBalancesRequestVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalAccountBalancesRequestVLSData;
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
impl crate::Message for HistoricalAccountBalancesRequestVLSUnsafe {
    type Data = HistoricalAccountBalancesRequestVLSData;

    const BASE_SIZE: usize = 24;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                HistoricalAccountBalancesRequestVLSData::new(),
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
            data: data as *const HistoricalAccountBalancesRequestVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for HistoricalAccountBalancesRequestVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const HistoricalAccountBalancesRequestVLSData;
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
/// This is a message from the Client to the Server to request a history of
/// Cash Balance changes for the specified Trade Account.
///
/// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// messages or reject he request with a message.
///
/// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
impl HistoricalAccountBalancesRequest for HistoricalAccountBalancesRequestVLS {
    type Safe = HistoricalAccountBalancesRequestVLS;
    type Unsafe = HistoricalAccountBalancesRequestVLSUnsafe;

    /// A unique request identifier for this request.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn start_date_time(&self) -> DateTime {
        i64::from_le(self.start_date_time)
    }

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = set_vls(self, self.trade_account, value);
        self
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        self.start_date_time = value.to_le();
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

/// This is a message from the Client to the Server to request a history of
/// Cash Balance changes for the specified Trade Account.
///
/// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// messages or reject he request with a message.
///
/// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
impl HistoricalAccountBalancesRequest for HistoricalAccountBalancesRequestVLSUnsafe {
    type Safe = HistoricalAccountBalancesRequestVLS;
    type Unsafe = HistoricalAccountBalancesRequestVLSUnsafe;

    /// A unique request identifier for this request.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(12) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn start_date_time(&self) -> DateTime {
        if self.is_out_of_bounds(24) {
            0i64
        } else {
            i64::from_le(self.start_date_time)
        }
    }

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(12) {
            self.request_id = value.to_le();
        }
        self
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(16) {
            self.trade_account = set_vls(self, self.trade_account, value);
        }
        self
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.start_date_time = value.to_le();
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

/// This is a message from the Client to the Server to request a history of
/// Cash Balance changes for the specified Trade Account.
///
/// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// messages or reject he request with a message.
///
/// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
impl HistoricalAccountBalancesRequest for HistoricalAccountBalancesRequestFixed {
    type Safe = HistoricalAccountBalancesRequestFixed;
    type Unsafe = HistoricalAccountBalancesRequestFixedUnsafe;

    /// A unique request identifier for this request.
    fn request_id(&self) -> i32 {
        i32::from_le(self.request_id)
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn trade_account(&self) -> &str {
        get_fixed(&self.trade_account[..])
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn start_date_time(&self) -> DateTime {
        i64::from_le(self.start_date_time)
    }

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.trade_account[..], value);
        self
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        self.start_date_time = value.to_le();
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

/// This is a message from the Client to the Server to request a history of
/// Cash Balance changes for the specified Trade Account.
///
/// The Server will respond with multiple HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// messages or reject he request with a message.
///
/// The Server will set the RequestID in the HISTORICAL_ACCOUNT_BALANCE_RESPONSE
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
/// message to match the RequestID in the HistoricalAccountBalancesRequestVLS.
impl HistoricalAccountBalancesRequest for HistoricalAccountBalancesRequestFixedUnsafe {
    type Safe = HistoricalAccountBalancesRequestFixed;
    type Unsafe = HistoricalAccountBalancesRequestFixedUnsafe;

    /// A unique request identifier for this request.
    fn request_id(&self) -> i32 {
        if self.is_out_of_bounds(8) {
            0i32
        } else {
            i32::from_le(self.request_id)
        }
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(40) {
            ""
        } else {
            get_fixed(&self.trade_account[..])
        }
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn start_date_time(&self) -> DateTime {
        if self.is_out_of_bounds(48) {
            0i64
        } else {
            i64::from_le(self.start_date_time)
        }
    }

    /// A unique request identifier for this request.
    fn set_request_id(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.request_id = value.to_le();
        }
        self
    }

    /// This is a required field. Set to the particular Trade Account for which
    /// to request historical Account Balance data.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            set_fixed(&mut self.trade_account[..], value);
        }
        self
    }

    /// Set this to the Date-Time that the server is to send historical cash balance
    /// updates starting with.
    fn set_start_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(48) {
            self.start_date_time = value.to_le();
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
                core::mem::size_of::<HistoricalAccountBalancesRequestFixedData>(),
                "HistoricalAccountBalancesRequestFixedData sizeof expected {:} but was {:}",
                48usize,
                core::mem::size_of::<HistoricalAccountBalancesRequestFixedData>()
            );
            assert_eq!(
                48u16,
                HistoricalAccountBalancesRequestFixed::new().size(),
                "HistoricalAccountBalancesRequestFixed sizeof expected {:} but was {:}",
                48u16,
                HistoricalAccountBalancesRequestFixed::new().size(),
            );
            assert_eq!(
                HISTORICAL_ACCOUNT_BALANCES_REQUEST,
                HistoricalAccountBalancesRequestFixed::new().r#type(),
                "HistoricalAccountBalancesRequestFixed type expected {:} but was {:}",
                HISTORICAL_ACCOUNT_BALANCES_REQUEST,
                HistoricalAccountBalancesRequestFixed::new().r#type(),
            );
            assert_eq!(
                603u16,
                HistoricalAccountBalancesRequestFixed::new().r#type(),
                "HistoricalAccountBalancesRequestFixed type expected {:} but was {:}",
                603u16,
                HistoricalAccountBalancesRequestFixed::new().r#type(),
            );
            let d = HistoricalAccountBalancesRequestFixedData::new();
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
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                40usize,
                (core::ptr::addr_of!(d.start_date_time) as usize) - p,
                "start_date_time offset expected {:} but was {:}",
                40usize,
                (core::ptr::addr_of!(d.start_date_time) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                24usize,
                core::mem::size_of::<HistoricalAccountBalancesRequestVLSData>(),
                "HistoricalAccountBalancesRequestVLSData sizeof expected {:} but was {:}",
                24usize,
                core::mem::size_of::<HistoricalAccountBalancesRequestVLSData>()
            );
            assert_eq!(
                24u16,
                HistoricalAccountBalancesRequestVLS::new().size(),
                "HistoricalAccountBalancesRequestVLS sizeof expected {:} but was {:}",
                24u16,
                HistoricalAccountBalancesRequestVLS::new().size(),
            );
            assert_eq!(
                HISTORICAL_ACCOUNT_BALANCES_REQUEST,
                HistoricalAccountBalancesRequestVLS::new().r#type(),
                "HistoricalAccountBalancesRequestVLS type expected {:} but was {:}",
                HISTORICAL_ACCOUNT_BALANCES_REQUEST,
                HistoricalAccountBalancesRequestVLS::new().r#type(),
            );
            assert_eq!(
                603u16,
                HistoricalAccountBalancesRequestVLS::new().r#type(),
                "HistoricalAccountBalancesRequestVLS type expected {:} but was {:}",
                603u16,
                HistoricalAccountBalancesRequestVLS::new().r#type(),
            );
            let d = HistoricalAccountBalancesRequestVLSData::new();
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
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                12usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.start_date_time) as usize) - p,
                "start_date_time offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.start_date_time) as usize) - p,
            );
        }
    }
}
