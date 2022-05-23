use crate::message::*;

pub trait LogonRequest: Message {
    fn protocol_version(&self) -> i32;

    fn username(&self) -> &str;

    fn set_username(&mut self, value: &str);

    fn copy_to(&self, to: &mut impl LogonRequest) {
        to.set_username(self.username());
    }

    fn copy_from(&mut self, from: &impl LogonRequest) {
        from.copy_to(self);
    }
}

//     Size                            uint16      = LogonRequestFixedSize  (284)
//     Type                            uint16      = LOGON_REQUEST  (1)
//     ProtocolVersion                 int32       = CURRENT_VERSION  (8)
//     Username                        string[32]  = ""
//     Password                        string[32]  = ""
//     GeneralTextData                 string[64]  = ""
//     Integer_1                       int32       = 0
//     Integer_2                       int32       = 0
//     HeartbeatIntervalInSeconds      int32       = 0
//     Unused1                         int32       = 0
//     TradeAccount                    string[32]  = ""
//     HardwareIdentifier              string[64]  = ""
//     ClientName                      string[32]  = ""
//     MarketDataTransmissionInterval  int32       = 0

pub struct LogonRequestFixed {
    data: *const LogonRequestFixedData,
}

pub struct LogonRequestFixedUnsafe {
    data: *const LogonRequestFixedData,
}

#[repr(packed, C)]
pub struct LogonRequestFixedData {
    size: u16,
    r#type: u16,
    protocol_version: i32,
    username: [u8; 32],
    password: [u8; 32],
    general_text_data: [u8; 64],
    integer_1: i32,
    integer_2: i32,
    heartbeat_interval_seconds: i32,
    unused_1: i32,
    trade_account: [u8; 32],
    hardware_identifier: [u8; 64],
    client_name: [u8; 32],
    market_data_transmission_inteveral: i32,
}

pub struct LogonRequestVLS {
    data: *const LogonRequestVLSData,
    capacity: usize,
}

pub struct LogonRequestVLSUnsafe {
    data: *const LogonRequestVLSData,
    capacity: usize,
}

#[repr(packed, C)]
pub struct LogonRequestVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    protocol_version: i32,
    username: VLS,
    password: VLS,
    general_text_data: VLS,
    integer_1: i32,
    integer_2: i32,
    heartbeat_interval_seconds: i32,
    unused_1: i32,
    trade_account: VLS,
    hardware_identifier: VLS,
    client_name: VLS,
    market_data_transmission_inteveral: i32,
}

unsafe impl Send for LogonRequestFixed {}
unsafe impl Send for LogonRequestFixedUnsafe {}
unsafe impl Send for LogonRequestFixedData {}
unsafe impl Send for LogonRequestVLS {}
unsafe impl Send for LogonRequestVLSUnsafe {}
unsafe impl Send for LogonRequestVLSData {}

impl Drop for LogonRequestFixed {
    #[inline]
    fn drop(&mut self) {
        deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for LogonRequestFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        deallocate(self.data as *mut u8, self.capacity() as usize)
    }
}

impl Drop for LogonRequestVLS {
    #[inline]
    fn drop(&mut self) {
        deallocate(self.data as *mut u8, self.capacity() as usize)
    }
}

impl Drop for LogonRequestVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        deallocate(self.data as *mut u8, self.capacity() as usize)
    }
}

// impl<T: LogonRequest> ToOwned for T {
//     type Owned = LogonRequest;
// }

impl Clone for LogonRequestFixed {
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for LogonRequestFixedUnsafe {
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for LogonRequestVLS {
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for LogonRequestVLSUnsafe {
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl std::ops::Deref for LogonRequestFixed {
    type Target = LogonRequestFixedData;

    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl std::ops::DerefMut for LogonRequestFixed {
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl std::ops::Deref for LogonRequestFixedUnsafe {
    type Target = LogonRequestFixedData;

    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl std::ops::DerefMut for LogonRequestFixedUnsafe {
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl LogonRequestFixedData {
    pub fn new() -> Self {
        Self {
            size: 284,
            r#type: 1,
            protocol_version: 8,
            username: [0; 32],
            password: [0; 32],
            general_text_data: [0; 64],
            integer_1: 0,
            integer_2: 0,
            heartbeat_interval_seconds: 0,
            unused_1: 0,
            trade_account: [0; 32],
            hardware_identifier: [0; 64],
            client_name: [0; 32],
            market_data_transmission_inteveral: 0,
        }
    }
}

impl Message for LogonRequestFixed {
    type Safe = LogonRequestFixed;
    type Unsafe = LogonRequestFixedUnsafe;
    type Data = LogonRequestFixedData;
    const BASE_SIZE_OFFSET: isize = 0;

    fn new() -> Self {
        Self {
            data: allocate(Self::BASE_SIZE, LogonRequestFixedData::new()),
        }
    }

    fn to_safe(self) -> Self::Safe {
        self
    }

    fn size(&self) -> u16 {
        self.size
    }

    fn base_size(&self) -> u16 {
        self.size
    }

    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    unsafe fn from_raw_parts(data: *const u8, _capacity: usize) -> Self {
        Self {
            data: data as *const LogonRequestFixedData,
        }
    }

    fn r#type(&self) -> u16 {
        self.r#type
    }
}

impl Message for LogonRequestFixedUnsafe {
    type Safe = LogonRequestFixed;
    type Unsafe = LogonRequestFixedUnsafe;
    type Data = LogonRequestFixedData;
    const BASE_SIZE_OFFSET: isize = 0;

    fn new() -> Self {
        Self {
            data: allocate(Self::BASE_SIZE, LogonRequestFixedData::new()),
        }
    }

    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    fn size(&self) -> u16 {
        self.size
    }

    fn base_size(&self) -> u16 {
        self.size
    }

    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.data
    }

    unsafe fn from_raw_parts(data: *const u8, _capacity: usize) -> Self {
        Self {
            data: data as *const LogonRequestFixedData,
        }
    }

    fn r#type(&self) -> u16 {
        self.r#type
    }
}

impl LogonRequest for LogonRequestFixed {
    fn protocol_version(&self) -> i32 {
        i32::from_le(self.protocol_version)
    }

    fn username(&self) -> &str {
        get_fixed(&self.username[0..32])
    }

    fn set_username(&mut self, value: &str) {
        set_fixed(&mut self.username[..], value);
    }
}

impl LogonRequest for LogonRequestFixedUnsafe {
    fn protocol_version(&self) -> i32 {
        if self.is_out_of_bounds(10) {
            0
        } else {
            i32::from_le(self.protocol_version)
        }
    }

    fn username(&self) -> &str {
        if self.is_out_of_bounds(40) {
            ""
        } else {
            get_fixed(&self.username[..])
        }
    }

    fn set_username(&mut self, value: &str) {
        if !self.is_out_of_bounds(40) {
            set_fixed(&mut self.username[..], value);
        }
    }
}

impl LogonRequestVLSData {
    fn new() -> Self {
        Self {
            size: core::mem::size_of::<Self>() as u16,
            r#type: 1,
            base_size: core::mem::size_of::<Self>() as u16,
            protocol_version: 8,
            username: VLS::new(),
            password: VLS::new(),
            general_text_data: VLS::new(),
            integer_1: 0,
            integer_2: 0,
            heartbeat_interval_seconds: 0,
            unused_1: 0,
            trade_account: VLS::new(),
            hardware_identifier: VLS::new(),
            client_name: VLS::new(),
            market_data_transmission_inteveral: 0,
        }
    }
}

impl std::ops::Deref for LogonRequestVLS {
    type Target = LogonRequestVLSData;

    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl std::ops::DerefMut for LogonRequestVLS {
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl std::ops::Deref for LogonRequestVLSUnsafe {
    type Target = LogonRequestVLSData;

    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl std::ops::DerefMut for LogonRequestVLSUnsafe {
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl Message for LogonRequestVLS {
    type Safe = LogonRequestVLS;
    type Unsafe = LogonRequestVLSUnsafe;
    type Data = LogonRequestVLSData;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: allocate(Self::DEFAULT_CAPACITY, LogonRequestVLSData::new()),
            capacity: Self::DEFAULT_CAPACITY,
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        self
    }

    #[inline]
    fn size(&self) -> u16 {
        self.size
    }

    #[inline]
    fn base_size(&self) -> u16 {
        self.base_size
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
            data: data as *const LogonRequestVLSData,
            capacity,
        }
    }

    #[inline]
    fn r#type(&self) -> u16 {
        self.r#type
    }
}

impl Message for LogonRequestVLSUnsafe {
    type Safe = LogonRequestVLS;
    type Unsafe = LogonRequestVLSUnsafe;
    type Data = LogonRequestVLSData;
    const BASE_SIZE_OFFSET: isize = 4;

    #[inline]
    fn new() -> Self {
        Self {
            data: allocate(Self::DEFAULT_CAPACITY, Self::Data::new()),
            capacity: Self::DEFAULT_CAPACITY,
        }
    }

    #[inline]
    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    #[inline]
    fn size(&self) -> u16 {
        self.size
    }

    #[inline]
    fn base_size(&self) -> u16 {
        self.base_size
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
            data: data as *const LogonRequestVLSData,
            capacity,
        }
    }

    #[inline]
    fn r#type(&self) -> u16 {
        self.r#type
    }
}

impl VLSMessage for LogonRequestVLS {
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const LogonRequestVLSData;
    }

    fn set_capacity(&mut self, size: u16) {
        self.capacity = size as usize;
    }

    fn set_size(&mut self, size: u16) {
        self.size = size;
    }
}

impl VLSMessage for LogonRequestVLSUnsafe {
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const LogonRequestVLSData;
    }

    fn set_capacity(&mut self, size: u16) {
        self.capacity = size as usize;
    }

    fn set_size(&mut self, size: u16) {
        self.size = size;
    }
}

impl LogonRequest for LogonRequestVLS {
    fn protocol_version(&self) -> i32 {
        i32::from_le(self.protocol_version)
    }

    fn username(&self) -> &str {
        get_vls(self, self.username)
    }

    fn set_username(&mut self, value: &str) {
        self.username = set_vls(self, self.username, value);
    }
}

impl LogonRequest for LogonRequestVLSUnsafe {
    fn protocol_version(&self) -> i32 {
        if self.is_out_of_bounds(10) {
            0
        } else {
            i32::from_le(self.protocol_version)
        }
    }

    fn username(&self) -> &str {
        if self.is_out_of_bounds(16) {
            ""
        } else {
            get_vls(self, self.username)
        }
    }

    fn set_username(&mut self, value: &str) {
        if !self.is_out_of_bounds(16) {
            self.username = set_vls(self, self.username, value);
        }
    }
}

impl Into<Vec<u8>> for LogonRequestFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for LogonRequestFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for LogonRequestVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for LogonRequestVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}