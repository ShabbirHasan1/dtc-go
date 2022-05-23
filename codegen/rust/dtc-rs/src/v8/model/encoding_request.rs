use crate::{set_fixed, Message, allocate};

pub trait EncodingRequest: Message {    
    fn protocol_version(&self) -> i32;

    fn encoding(&self) -> EncodingEnum;

    fn protocol_type(&self) -> &str;

    fn set_protocol_version(&mut self, value: i32) -> &mut Self;

    fn set_encoding(&mut self, value: EncodingEnum) -> &mut Self;

    fn set_protocol_type(&mut self, value: &str) -> &mut Self;

    fn copy_to(&self, to: &mut impl EncodingRequest) {
        to.set_protocol_version(self.protocol_version())
            .set_encoding(self.encoding())
            .set_protocol_type(self.protocol_type());
    }
}

#[derive(Clone, Copy)]
#[repr(u32)]
pub enum EncodingEnum {
    BinaryEncoding = 0,
    BinaryWithVariableLengthStrings = 1,
    // JsonEncoding = 2,
    // JsonCompactEncoding = 3,
    // ProtocolBuffers = 4,
}

impl EncodingEnum {
    #[inline]
    pub fn to_le(self) -> Self {
        unsafe { core::mem::transmute((self as u32).to_le()) }
    }

    #[inline]
    pub fn from_le(value: Self) -> Self {
        unsafe { core::mem::transmute(u32::from_le(value as u32)) }
    }
}

pub struct EncodingRequestSafe(*const EncodingRequestData);
pub struct EncodingRequestUnsafe(*const EncodingRequestData);

unsafe impl Send for EncodingRequestSafe {}
unsafe impl Send for EncodingRequestUnsafe {}
unsafe impl Send for EncodingRequestData {}

impl core::ops::Deref for EncodingRequestSafe {
    type Target = EncodingRequestData;

    fn deref(&self) -> &Self::Target {
        unsafe { &*self.0 }
    }
}

impl core::ops::DerefMut for EncodingRequestSafe {
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.0 as *mut Self::Target) }
    }
}

impl core::ops::Deref for EncodingRequestUnsafe {
    type Target = EncodingRequestData;

    fn deref(&self) -> &Self::Target {
        unsafe { &*self.0 }
    }
}

impl core::ops::DerefMut for EncodingRequestUnsafe {
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.0 as *mut Self::Target) }
    }
}

impl Message for EncodingRequestSafe {
    type Safe = EncodingRequestSafe;
    type Unsafe = EncodingRequestUnsafe;
    type Data = EncodingRequestData;
    const BASE_SIZE_OFFSET: isize = 0;

    fn new() -> Self {
        Self(allocate(Self::DEFAULT_CAPACITY, Self::Data::new()))
    }

    fn to_safe(self) -> Self::Safe {
        self
    }

    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.0
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _capacity: usize) -> Self {
        Self(data as *const EncodingRequestData)
    }

    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }
}

impl Message for EncodingRequestUnsafe {
    type Safe = EncodingRequestSafe;
    type Unsafe = EncodingRequestUnsafe;
    type Data = EncodingRequestData;
    const BASE_SIZE_OFFSET: isize = 0;

    fn new() -> Self {
        Self(allocate(Self::DEFAULT_CAPACITY, Self::Data::new()))
    }

    fn to_safe(self) -> Self::Safe {
        let mut s = Self::Safe::new();
        self.copy_to(&mut s);
        s
    }

    fn size(&self) -> u16 {
        u16::from_le(self.size)
    }

    fn base_size(&self) -> u16 {
        u16::from_le(self.size)
    }

    unsafe fn as_ptr(&self) -> *const Self::Data {
        self.0
    }

    #[inline]
    unsafe fn from_raw_parts(data: *const u8, _capacity: usize) -> Self {
        Self(data as *const EncodingRequestData)
    }

    fn r#type(&self) -> u16 {
        u16::from_le(self.r#type)
    }
}

impl Clone for EncodingRequestSafe {
    fn clone(&self) -> Self {
        let mut r = Self::new();
        self.copy_to(&mut r);
        r
    }
}

impl Clone for EncodingRequestUnsafe {
    fn clone(&self) -> Self {
        let mut r = Self::new();
        self.copy_to(&mut r);
        r
    }
}

impl Drop for EncodingRequestSafe {
    fn drop(&mut self) {
        unsafe {
            Vec::from_raw_parts(
                self.0 as *mut u8,
                self.size() as usize,
                self.size() as usize,
            );
        }
    }
}

impl Drop for EncodingRequestUnsafe {
    fn drop(&mut self) {
        unsafe {
            Vec::from_raw_parts(
                self.0 as *mut u8,
                self.size() as usize,
                self.size() as usize,
            );
        }
    }
}

#[repr(packed, C)]
pub struct EncodingRequestData {
    size: u16,
    r#type: u16,
    protocol_version: i32,
    encoding: EncodingEnum,
    protocol_type: [u8; 4],
}

impl EncodingRequestData {
    pub fn new() -> Self {
        Self {
            size: 16u16.to_le(),
            r#type: 2u16.to_le(),
            protocol_version: 8i32.to_le(),
            encoding: EncodingEnum::BinaryWithVariableLengthStrings.to_le(),
            protocol_type: [b'D', b'T', b'C', b'\0'],
        }
    }
}

impl EncodingRequest for EncodingRequestSafe {
    fn protocol_version(&self) -> i32 {
        i32::from_le(self.protocol_version)
    }

    fn encoding(&self) -> EncodingEnum {
        unsafe { core::mem::transmute(u32::from_le(self.encoding as u32)) }
    }

    fn protocol_type(&self) -> &str {
        unsafe {
            std::str::from_utf8_unchecked(
                std::ffi::CStr::from_ptr(&self.protocol_type as *const _ as *const i8)
                    .to_bytes(),
            )
        }
    }

    fn set_encoding(&mut self, value: EncodingEnum) -> &mut Self {
        self.encoding = value.to_le();
        self
    }

    fn set_protocol_type(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.protocol_type[..], value);
        self
    }

    fn set_protocol_version(&mut self, value: i32) -> &mut Self {
        self.protocol_version = value.to_le();
        self
    }
}

impl EncodingRequest for EncodingRequestUnsafe {
    fn protocol_version(&self) -> i32 {
        if self.is_out_of_bounds(4) {
            0
        } else {
            i32::from_le(self.protocol_version)
        }
    }

    fn encoding(&self) -> EncodingEnum {
        if self.is_out_of_bounds(8) {
            EncodingEnum::BinaryEncoding
        } else {
            EncodingEnum::from_le(self.encoding)
        }
    }

    fn protocol_type(&self) -> &str {
        if self.is_out_of_bounds(12) {
            ""
        } else {
            unsafe {
                std::str::from_utf8_unchecked(
                    std::ffi::CStr::from_ptr(&self.protocol_type as *const _ as *const i8)
                        .to_bytes(),
                )
            }
        }
    }

    fn set_encoding(&mut self, value: EncodingEnum) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.encoding = value;
        }
        self
    }

    fn set_protocol_type(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.protocol_type[..], value);
        self
    }

    fn set_protocol_version(&mut self, value: i32) -> &mut Self {
        self.protocol_version = value.to_le();
        self
    }
}


impl Into<Vec<u8>> for EncodingRequestSafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for EncodingRequestUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}