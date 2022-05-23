use crate::error::Error;
use core::ptr;

pub enum Parsed<'a, A: Message, B: Message> {
    Left(&'a A),
    Right(&'a B),
}

pub trait Message: Sized + Into<Vec<u8>> + Send {
    type Safe: Message;
    type Unsafe: Message;
    type Data: Sized;

    const BASE_SIZE: usize = core::mem::size_of::<Self::Data>();
    const BASE_SIZE_OFFSET: isize;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE;

    fn new() -> Self;

    fn to_safe(self) -> Self::Safe;

    #[inline]
    fn is_out_of_bounds(&self, offset: u16) -> bool {
        self.base_size() < offset
    }

    fn size(&self) -> u16;

    fn base_size(&self) -> u16;

    fn capacity(&self) -> u16 {
        self.size()
    }

    fn r#type(&self) -> u16;

    unsafe fn as_ptr(&self) -> *const Self::Data;

    unsafe fn from_raw_parts(ptr: *const u8, capacity: usize) -> Self;

    unsafe fn leak<'a>(ptr: *const u8, capacity: usize) -> &'a Self {
        let s = Self::from_raw_parts(ptr, capacity);
        let r = &s as *const Self;
        core::mem::forget(s);
        &*r
    }

    #[inline]
    fn parse(data: &[u8]) -> Result<Parsed<Self::Safe, Self::Unsafe>, Error> {
        if data.len() < 6 {
            return Err(Error::Malformed("need more data"));
        }
        let size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) };
        let base_size = if Self::BASE_SIZE_OFFSET == 0 {
            size
        } else {
            unsafe { u16::from_le(*(data.as_ptr().offset(Self::BASE_SIZE_OFFSET) as *const u16)) }
        };
        if (base_size as usize) >= Self::BASE_SIZE {
            unsafe { Ok(Parsed::Left(Self::Safe::leak(data.as_ptr(), size as usize))) }
        } else {
            unsafe {
                Ok(Parsed::Right(Self::Unsafe::leak(
                    data.as_ptr(),
                    size as usize,
                )))
            }
        }
    }

    #[inline]
    fn into_vec(self) -> Vec<u8> {
        let r = unsafe {
            Vec::from_raw_parts(
                self.as_ptr() as *mut u8,
                self.size() as usize,
                self.capacity() as usize,
            )
        };
        core::mem::forget(self);
        r
    }
}

pub(crate) trait VLSMessage: Message {
    unsafe fn set_ptr(&mut self, value: *const u8);

    fn set_capacity(&mut self, size: u16);

    fn set_size(&mut self, size: u16);
}

#[derive(Copy, Clone, Debug)]
pub struct VLS {
    offset: u16,
    length: u16,
}

impl VLS {
    #[inline]
    pub fn new() -> Self {
        Self {
            offset: 0,
            length: 0,
        }
    }
    #[inline]
    pub fn offset(&self) -> u16 {
        self.offset.to_le()
    }
    #[inline]
    pub fn length(&self) -> u16 {
        self.length.to_le()
    }
    #[inline]
    fn set_offset(&mut self, value: u16) {
        self.offset = u16::from_le(value)
    }
    #[inline]
    fn set_length(&mut self, value: u16) {
        self.length = u16::from_le(value)
    }
}

pub(crate) fn get_fixed_string(src: &[u8]) -> &str {
    match memchr::memchr(0, src) {
        Some(idx) => unsafe { std::str::from_utf8_unchecked(&src[0..idx]) },
        None => unsafe { std::str::from_utf8_unchecked(src) },
    }
}

pub(crate) fn set_fixed_string(src: &str, dst: &mut [u8]) {
    unsafe {
        if src.len() < dst.len() {
            ptr::copy(src.as_ptr() as *mut u8, dst.as_ptr() as *mut u8, src.len());
            dst[src.len()] = 0;
        } else {
            ptr::copy(src.as_ptr(), dst.as_ptr() as *mut u8, dst.len() - 1);
            dst[dst.len() - 1] = 0;
        }
    }
}

pub(crate) fn get_vls<M: VLSMessage>(m: &M, offset: VLS) -> &str {
    if (m.size() as u32) < (offset.offset() as u32 + offset.length() as u32) || offset.length() == 0
    {
        ""
    } else {
        unsafe {
            // trim NULL terminator?
            if *(m.as_ptr() as *const u8)
                .offset((offset.offset() as isize) + (offset.length() as isize))
                == 0
            {
                std::str::from_utf8_unchecked(std::slice::from_raw_parts(
                    (m.as_ptr() as *const u8).offset(offset.offset() as isize),
                    (offset.length() as usize) - 1,
                ))
            } else {
                std::str::from_utf8_unchecked(std::slice::from_raw_parts(
                    (m.as_ptr() as *const u8).offset(offset.offset() as isize),
                    offset.length() as usize,
                ))
            }
        }
    }
}

pub(crate) fn set_vls<M: VLSMessage>(
    message: &mut M,
    mut offset: VLS,
    value: &str,
) -> core::result::Result<VLS, Error> {
    if (message.size() as u32) < (offset.offset() as u32 + offset.length() as u32) {
        Ok(VLS::new())
    } else if offset.length() > value.len() as u16 {
        // In-place write.
        offset.set_length((value.len() as u16) + 1);
        unsafe {
            ptr::copy(
                value.as_ptr(),
                message.as_ptr().offset(offset.offset() as isize) as *mut u8,
                value.len(),
            );
            *(message.as_ptr() as *mut u8)
                .offset(offset.offset() as isize + value.len() as isize) = 0;
        }
        Ok(offset)
    } else {
        // Overflow?
        let new_size = message.size() as usize + value.len() + 1;
        if new_size > u16::MAX as usize {
            return Err(Error::Overflow);
        }

        // Add NULL terminator to length.
        offset.set_length((value.len() as u16) + 1);
        // Append to end of message.
        offset.set_offset(message.size());

        unsafe {
            // Delegate memory operations to Vec.
            let mut data = Vec::from_raw_parts(
                message.as_ptr() as *mut u8,
                message.size() as usize,
                message.capacity() as usize,
            );
            data.extend_from_slice(value.as_bytes());
            // Append NULL terminator.
            data.push(0u8);
            message.set_ptr(data.as_ptr());
            message.set_size(new_size as u16);
            message.set_capacity(data.capacity() as u16);
            core::mem::forget(data);
        }
        Ok(offset)
    }
}

pub(crate) fn allocate<T: Sized>(size: usize, value: T) -> *mut T {
    unsafe {
        let v = Vec::<u8>::with_capacity(size);
        let data = v.as_ptr() as *mut T;
        *data = value;
        core::mem::forget(v);
        data
    }
}

pub(crate) fn deallocate(data: *mut u8, capacity: usize) {
    unsafe {
        Vec::from_raw_parts(data, capacity, capacity);
    }
}
