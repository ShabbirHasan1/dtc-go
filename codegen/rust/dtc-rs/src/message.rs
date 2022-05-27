use crate::error::Error;
use core::ptr;
use std::alloc::Layout;

pub enum Parsed<'a, A: Message, B: Message> {
    Safe(&'a A),
    Unsafe(&'a B),
}

pub trait Message: Sized + Into<Vec<u8>> + Send + core::fmt::Display + core::fmt::Debug {
    // type Safe: Message;
    // type Unsafe: Message;
    type Data: Sized;

    const TYPE: u16;
    const BASE_SIZE: usize = core::mem::size_of::<Self::Data>();
    const BASE_SIZE_OFFSET: isize;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE;

    fn new() -> Self;

    #[inline]
    fn is_out_of_bounds(&self, min_base_size: u16) -> bool {
        self.base_size() < min_base_size
    }

    fn size(&self) -> u16;

    fn r#type(&self) -> u16;

    fn base_size(&self) -> u16;

    fn capacity(&self) -> u16 {
        self.size()
    }

    unsafe fn as_ptr(&self) -> *const Self::Data;

    unsafe fn from_raw_parts(ptr: *const u8, capacity: usize) -> Self;

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

// #[inline]
// pub fn parse<S: Message, U: Message>(data: &[u8]) -> Result<Parsed<S, U>, Error> {
//     if data.len() < 6 {
//         return Err(Error::Malformed("need more data"));
//     }
//     let size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) };
//     let base_size = if S::BASE_SIZE_OFFSET == 0 {
//         size
//     } else {
//         let base_size =
//             unsafe { u16::from_le(*(data.as_ptr().offset(S::BASE_SIZE_OFFSET) as *const u16)) };
//         if base_size > size {
//             return Err(Error::Malformed("base_size is greater than size"));
//         }
//         base_size
//     };
//     if (base_size as usize) >= S::BASE_SIZE {
//         unsafe {
//             let msg = S::from_raw_parts(data.as_ptr(), size as usize);
//         }
//         unsafe { Ok(Parsed::Safe(S::leak(data.as_ptr(), size as usize))) }
//     } else {
//         unsafe { Ok(Parsed::Unsafe(U::leak(data.as_ptr(), size as usize))) }
//     }
// }
//
// #[inline]
// pub fn parse2<S: Message, U: Message, F: Fn(Parsed<S, U>) -> Result<(), Error>>(
//     data: &[u8],
//     f: F,
// ) -> Result<(), Error> {
//     if data.len() < 6 {
//         return Err(Error::Malformed("need more data"));
//     }
//     let size = unsafe { u16::from_le(*(data.as_ptr() as *const u16)) };
//     let base_size = if S::BASE_SIZE_OFFSET == 0 {
//         size
//     } else {
//         let base_size =
//             unsafe { u16::from_le(*(data.as_ptr().offset(S::BASE_SIZE_OFFSET) as *const u16)) };
//         if base_size > size {
//             return Err(Error::Malformed("base_size is greater than size"));
//         }
//         base_size
//     };
//     if (base_size as usize) >= S::BASE_SIZE {
//         unsafe { f(Parsed::Safe(S::leak(data.as_ptr(), size as usize))) }
//     } else {
//         unsafe { f(Parsed::Unsafe(U::leak(data.as_ptr(), size as usize))) }
//     }
// }

#[inline]
pub(crate) const fn f32_le(value: f32) -> f32 {
    #[cfg(target_endian = "little")]
    {
        value
    }
    #[cfg(not(target_endian = "little"))]
    {
        let v: u32 = unsafe { core::mem::transmute(value) };
        unsafe { core::mem::transmute(v.to_le()) }
    }
}

#[inline]
pub(crate) const fn f64_le(value: f64) -> f64 {
    #[cfg(target_endian = "little")]
    {
        value
    }
    #[cfg(not(target_endian = "little"))]
    {
        let v: u64 = unsafe { core::mem::transmute(value) };
        unsafe { core::mem::transmute(v.to_le()) }
    }
}

pub(crate) trait VLSMessage: Message {
    unsafe fn set_ptr(&mut self, value: *const u8);

    fn set_capacity(&mut self, size: u16);

    fn set_size(&mut self, size: u16);
}

#[derive(Copy, Clone, Debug)]
#[repr(packed(8))]
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

pub(crate) fn get_fixed(src: &[u8]) -> &str {
    match memchr::memchr(0, src) {
        Some(idx) => unsafe { std::str::from_utf8_unchecked(&src[0..idx]) },
        None => unsafe { std::str::from_utf8_unchecked(src) },
    }
}

pub(crate) fn set_fixed(dst: &mut [u8], src: &str) {
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
    let data = unsafe { core::slice::from_raw_parts(m.as_ptr() as *mut u8, m.size() as usize) };
    if (m.size() as u32) < (offset.offset() as u32 + offset.length() as u32) || offset.length() == 0
    {
        ""
    } else {
        unsafe {
            let data = (m.as_ptr() as *const u8).offset(offset.offset() as isize);
            std::str::from_utf8_unchecked(std::slice::from_raw_parts(
                data,
                offset.length() as usize,
            ))
        }
        // let mut s = &data[offset.offset() as usize..(offset.offset() as usize + offset.length() as usize)];
        // while let Some(l) = s.last() {
        //     if *l == 0 {
        //         s = &s[0..(s.len()-1)];
        //     } else {
        //         break;
        //     }
        // }
        // let d = unsafe { std::str::from_utf8_unchecked(s) };
        // d
        // let s = s.trim_ascii_end();
        // match memchr::memchr(0, s) {
        //     Some(i) => {
        //         unsafe {
        //             std::str::from_utf8_unchecked(std::slice::from_raw_parts(
        //                 (&s[0..i]).as_ptr() as *mut u8,
        //                 i
        //             ))
        //         }
        //     }
        //     None => {
        //         unsafe {
        //             std::str::from_utf8_unchecked(std::slice::from_raw_parts(
        //                 s.as_ptr() as *mut u8,
        //                 s.len()
        //             ))
        //         }
        //     }
        // }
        // unsafe {
        //     // trim NULL terminator?
        //     if *(m.as_ptr() as *const u8)
        //         .offset((offset.offset() as isize) + (offset.length() as isize))
        //         == 0
        //     {
        //
        //         std::str::from_utf8_unchecked(std::slice::from_raw_parts(
        //             (m.as_ptr() as *const u8).offset(offset.offset() as isize),
        //             (offset.length() as usize) - 1,
        //         ))
        //     } else {
        //         std::str::from_utf8_unchecked(std::slice::from_raw_parts(
        //             (m.as_ptr() as *const u8).offset(offset.offset() as isize),
        //             offset.length() as usize,
        //         ))
        //     }
        // }
    }
}

pub(crate) fn set_vls<M: VLSMessage>(message: &mut M, mut offset: VLS, value: &str) -> VLS {
    if (message.size() as u32) < (offset.offset() as u32 + offset.length() as u32) {
        VLS::new()
    } else if offset.length() > value.len() as u16 {
        // In-place write.
        // offset.set_length((value.len() as u16) + 1);
        offset.set_length((value.len() as u16));
        unsafe {
            ptr::copy(
                value.as_ptr(),
                message.as_ptr().offset(offset.offset() as isize) as *mut u8,
                value.len(),
            );
            *(message.as_ptr() as *mut u8)
                .offset(offset.offset() as isize + value.len() as isize) = 0;
        }
        // offset.set_length(value.len() as u16);
        offset
    } else {
        // Overflow?
        let new_size = message.size() as usize + value.len() + 1;
        if new_size > u16::MAX as usize {
            return VLS::new();
        }

        // Add NULL terminator to length.
        // offset.set_length((value.len() as u16) + 1);
        offset.set_length((value.len() as u16));
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
        let r = get_vls(message, offset);
        // offset.set_length(value.len() as u16);
        offset
    }
}

pub(crate) fn allocate<T: Sized>(size: usize, value: T) -> *mut T {
    unsafe {
        let data = std::alloc::alloc(Layout::from_size_align_unchecked(size, 1)) as *mut T;
        // let v = String::with_capacity(size);
        // let data = v.as_ptr() as *mut T;
        // core::mem::forget(v);
        // let v = Vec::<u8>::with_capacity(size).leak();
        // let data = v.as_ptr() as *mut T;
        *data = value;
        // core::mem::forget(v);
        data
    }
}

pub(crate) fn deallocate(data: *mut u8, capacity: usize) {
    unsafe {
        std::alloc::dealloc(data, Layout::from_size_align_unchecked(capacity, 1));
        // String::from_raw_parts(data, capacity, capacity);
        // Vec::from_raw_parts(data, capacity, capacity);
    }
}
