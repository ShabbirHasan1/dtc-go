// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const JOURNAL_ENTRY_RESPONSE_VLS_SIZE: usize = 32;

pub(crate) const JOURNAL_ENTRY_RESPONSE_FIXED_SIZE: usize = 280;

/// size               u16       = JournalEntryResponseVLSSize  (32)
/// type               u16       = JOURNAL_ENTRY_RESPONSE  (706)
/// base_size          u16       = JournalEntryResponseVLSSize  (32)
/// journal_entry      string    = ""
/// date_time          DateTime  = 0
/// is_final_response  bool      = false
pub(crate) const JOURNAL_ENTRY_RESPONSE_VLS_DEFAULT: [u8; 32] = [
    32, 0, 194, 2, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0,
];

/// size               u16        = JournalEntryResponseFixedSize  (280)
/// type               u16        = JOURNAL_ENTRY_RESPONSE  (706)
/// journal_entry      string256  = ""
/// date_time          DateTime   = 0
/// is_final_response  bool       = false
pub(crate) const JOURNAL_ENTRY_RESPONSE_FIXED_DEFAULT: [u8; 280] = [
    24, 1, 194, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait JournalEntryResponse: Message {
    type Safe: JournalEntryResponse;
    type Unsafe: JournalEntryResponse;

    fn journal_entry(&self) -> &str;

    fn date_time(&self) -> DateTime;

    fn is_final_response(&self) -> bool;

    fn set_journal_entry(&mut self, value: &str) -> &mut Self;

    fn set_date_time(&mut self, value: DateTime) -> &mut Self;

    fn set_is_final_response(&mut self, value: bool) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl JournalEntryResponse) {
        to.set_journal_entry(self.journal_entry());
        to.set_date_time(self.date_time());
        to.set_is_final_response(self.is_final_response());
    }
}

pub struct JournalEntryResponseVLS {
    data: *const JournalEntryResponseVLSData,
    capacity: usize,
}

pub struct JournalEntryResponseVLSUnsafe {
    data: *const JournalEntryResponseVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct JournalEntryResponseVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    journal_entry: VLS,
    date_time: DateTime,
    is_final_response: bool,
}

pub struct JournalEntryResponseFixed {
    data: *const JournalEntryResponseFixedData,
}

pub struct JournalEntryResponseFixedUnsafe {
    data: *const JournalEntryResponseFixedData,
}

#[repr(packed(8), C)]
pub struct JournalEntryResponseFixedData {
    size: u16,
    r#type: u16,
    journal_entry: [u8; 256],
    date_time: DateTime,
    is_final_response: bool,
}

impl JournalEntryResponseVLSData {
    pub fn new() -> Self {
        Self {
            size: 32u16.to_le(),
            r#type: JOURNAL_ENTRY_RESPONSE.to_le(),
            base_size: 32u16.to_le(),
            journal_entry: crate::message::VLS::new(),
            date_time: 0i64,
            is_final_response: false,
        }
    }
}

impl JournalEntryResponseFixedData {
    pub fn new() -> Self {
        Self {
            size: 280u16.to_le(),
            r#type: JOURNAL_ENTRY_RESPONSE.to_le(),
            journal_entry: [0; 256],
            date_time: 0i64,
            is_final_response: false,
        }
    }
}

unsafe impl Send for JournalEntryResponseFixed {}
unsafe impl Send for JournalEntryResponseFixedUnsafe {}
unsafe impl Send for JournalEntryResponseFixedData {}
unsafe impl Send for JournalEntryResponseVLS {}
unsafe impl Send for JournalEntryResponseVLSUnsafe {}
unsafe impl Send for JournalEntryResponseVLSData {}

impl Drop for JournalEntryResponseFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for JournalEntryResponseFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for JournalEntryResponseVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for JournalEntryResponseVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for JournalEntryResponseFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for JournalEntryResponseFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for JournalEntryResponseVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for JournalEntryResponseVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for JournalEntryResponseFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for JournalEntryResponseFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for JournalEntryResponseVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for JournalEntryResponseVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for JournalEntryResponseFixed {
    type Target = JournalEntryResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for JournalEntryResponseFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for JournalEntryResponseFixedUnsafe {
    type Target = JournalEntryResponseFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for JournalEntryResponseFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for JournalEntryResponseVLS {
    type Target = JournalEntryResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for JournalEntryResponseVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for JournalEntryResponseVLSUnsafe {
    type Target = JournalEntryResponseVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for JournalEntryResponseVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for JournalEntryResponseFixed {
    type Data = JournalEntryResponseFixedData;

    const BASE_SIZE: usize = 280;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, JournalEntryResponseFixedData::new()),
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
            data: data as *const JournalEntryResponseFixedData,
        }
    }
}
impl crate::Message for JournalEntryResponseFixedUnsafe {
    type Data = JournalEntryResponseFixedData;

    const BASE_SIZE: usize = 280;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, JournalEntryResponseFixedData::new()),
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
            data: data as *const JournalEntryResponseFixedData,
        }
    }
}
impl crate::Message for JournalEntryResponseVLS {
    type Data = JournalEntryResponseVLSData;

    const BASE_SIZE: usize = 32;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, JournalEntryResponseVLSData::new()),
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
            data: data as *const JournalEntryResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for JournalEntryResponseVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const JournalEntryResponseVLSData;
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
impl crate::Message for JournalEntryResponseVLSUnsafe {
    type Data = JournalEntryResponseVLSData;

    const BASE_SIZE: usize = 32;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, JournalEntryResponseVLSData::new()),
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
            data: data as *const JournalEntryResponseVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for JournalEntryResponseVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const JournalEntryResponseVLSData;
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
impl JournalEntryResponse for JournalEntryResponseVLS {
    type Safe = JournalEntryResponseVLS;
    type Unsafe = JournalEntryResponseVLSUnsafe;

    fn journal_entry(&self) -> &str {
        get_vls(self, self.journal_entry)
    }

    fn date_time(&self) -> DateTime {
        i64::from_le(self.date_time)
    }

    fn is_final_response(&self) -> bool {
        self.is_final_response
    }

    fn set_journal_entry(&mut self, value: &str) -> &mut Self {
        self.journal_entry = set_vls(self, self.journal_entry, value);
        self
    }

    fn set_date_time(&mut self, value: DateTime) -> &mut Self {
        self.date_time = value.to_le();
        self
    }

    fn set_is_final_response(&mut self, value: bool) -> &mut Self {
        self.is_final_response = value;
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

impl JournalEntryResponse for JournalEntryResponseVLSUnsafe {
    type Safe = JournalEntryResponseVLS;
    type Unsafe = JournalEntryResponseVLSUnsafe;

    fn journal_entry(&self) -> &str {
        if self.is_out_of_bounds(10) {
            ""
        } else {
            get_vls(self, self.journal_entry)
        }
    }

    fn date_time(&self) -> DateTime {
        if self.is_out_of_bounds(24) {
            0i64
        } else {
            i64::from_le(self.date_time)
        }
    }

    fn is_final_response(&self) -> bool {
        if self.is_out_of_bounds(25) {
            false
        } else {
            self.is_final_response
        }
    }

    fn set_journal_entry(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.journal_entry = set_vls(self, self.journal_entry, value);
        }
        self
    }

    fn set_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.date_time = value.to_le();
        }
        self
    }

    fn set_is_final_response(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(25) {
            self.is_final_response = value;
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

impl JournalEntryResponse for JournalEntryResponseFixed {
    type Safe = JournalEntryResponseFixed;
    type Unsafe = JournalEntryResponseFixedUnsafe;

    fn journal_entry(&self) -> &str {
        get_fixed(&self.journal_entry[..])
    }

    fn date_time(&self) -> DateTime {
        i64::from_le(self.date_time)
    }

    fn is_final_response(&self) -> bool {
        self.is_final_response
    }

    fn set_journal_entry(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.journal_entry[..], value);
        self
    }

    fn set_date_time(&mut self, value: DateTime) -> &mut Self {
        self.date_time = value.to_le();
        self
    }

    fn set_is_final_response(&mut self, value: bool) -> &mut Self {
        self.is_final_response = value;
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

impl JournalEntryResponse for JournalEntryResponseFixedUnsafe {
    type Safe = JournalEntryResponseFixed;
    type Unsafe = JournalEntryResponseFixedUnsafe;

    fn journal_entry(&self) -> &str {
        if self.is_out_of_bounds(260) {
            ""
        } else {
            get_fixed(&self.journal_entry[..])
        }
    }

    fn date_time(&self) -> DateTime {
        if self.is_out_of_bounds(272) {
            0i64
        } else {
            i64::from_le(self.date_time)
        }
    }

    fn is_final_response(&self) -> bool {
        if self.is_out_of_bounds(273) {
            false
        } else {
            self.is_final_response
        }
    }

    fn set_journal_entry(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(260) {
            set_fixed(&mut self.journal_entry[..], value);
        }
        self
    }

    fn set_date_time(&mut self, value: DateTime) -> &mut Self {
        if !self.is_out_of_bounds(272) {
            self.date_time = value.to_le();
        }
        self
    }

    fn set_is_final_response(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(273) {
            self.is_final_response = value;
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
                280usize,
                core::mem::size_of::<JournalEntryResponseFixedData>(),
                "JournalEntryResponseFixedData sizeof expected {:} but was {:}",
                280usize,
                core::mem::size_of::<JournalEntryResponseFixedData>()
            );
            assert_eq!(
                280u16,
                JournalEntryResponseFixed::new().size(),
                "JournalEntryResponseFixed sizeof expected {:} but was {:}",
                280u16,
                JournalEntryResponseFixed::new().size(),
            );
            assert_eq!(
                JOURNAL_ENTRY_RESPONSE,
                JournalEntryResponseFixed::new().r#type(),
                "JournalEntryResponseFixed type expected {:} but was {:}",
                JOURNAL_ENTRY_RESPONSE,
                JournalEntryResponseFixed::new().r#type(),
            );
            assert_eq!(
                706u16,
                JournalEntryResponseFixed::new().r#type(),
                "JournalEntryResponseFixed type expected {:} but was {:}",
                706u16,
                JournalEntryResponseFixed::new().r#type(),
            );
            let d = JournalEntryResponseFixedData::new();
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
                (core::ptr::addr_of!(d.journal_entry) as usize) - p,
                "journal_entry offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.journal_entry) as usize) - p,
            );
            assert_eq!(
                264usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
                "date_time offset expected {:} but was {:}",
                264usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
            );
            assert_eq!(
                272usize,
                (core::ptr::addr_of!(d.is_final_response) as usize) - p,
                "is_final_response offset expected {:} but was {:}",
                272usize,
                (core::ptr::addr_of!(d.is_final_response) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                32usize,
                core::mem::size_of::<JournalEntryResponseVLSData>(),
                "JournalEntryResponseVLSData sizeof expected {:} but was {:}",
                32usize,
                core::mem::size_of::<JournalEntryResponseVLSData>()
            );
            assert_eq!(
                32u16,
                JournalEntryResponseVLS::new().size(),
                "JournalEntryResponseVLS sizeof expected {:} but was {:}",
                32u16,
                JournalEntryResponseVLS::new().size(),
            );
            assert_eq!(
                JOURNAL_ENTRY_RESPONSE,
                JournalEntryResponseVLS::new().r#type(),
                "JournalEntryResponseVLS type expected {:} but was {:}",
                JOURNAL_ENTRY_RESPONSE,
                JournalEntryResponseVLS::new().r#type(),
            );
            assert_eq!(
                706u16,
                JournalEntryResponseVLS::new().r#type(),
                "JournalEntryResponseVLS type expected {:} but was {:}",
                706u16,
                JournalEntryResponseVLS::new().r#type(),
            );
            let d = JournalEntryResponseVLSData::new();
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
                (core::ptr::addr_of!(d.journal_entry) as usize) - p,
                "journal_entry offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.journal_entry) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
                "date_time offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
            );
            assert_eq!(
                24usize,
                (core::ptr::addr_of!(d.is_final_response) as usize) - p,
                "is_final_response offset expected {:} but was {:}",
                24usize,
                (core::ptr::addr_of!(d.is_final_response) as usize) - p,
            );
        }
    }
}