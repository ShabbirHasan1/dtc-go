// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const USER_INFORMATION_VLS_SIZE: usize = 14;

/// size         u16     = UserInformationVLSSize  (14)
/// type         u16     = USER_INFORMATION  (10136)
/// base_size    u16     = UserInformationVLSSize  (14)
/// operator_id  string  = ""
/// location_id  string  = ""
pub(crate) const USER_INFORMATION_VLS_DEFAULT: [u8; 14] =
    [14, 0, 152, 39, 14, 0, 0, 0, 0, 0, 0, 0, 0, 0];

pub trait UserInformation: Message {
    type Safe: UserInformation;
    type Unsafe: UserInformation;

    fn operator_id(&self) -> &str;

    fn location_id(&self) -> &str;

    fn set_operator_id(&mut self, value: &str) -> &mut Self;

    fn set_location_id(&mut self, value: &str) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl UserInformation) {
        to.set_operator_id(self.operator_id());
        to.set_location_id(self.location_id());
    }
}

pub struct UserInformationVLS {
    data: *const UserInformationVLSData,
    capacity: usize,
}

pub struct UserInformationVLSUnsafe {
    data: *const UserInformationVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct UserInformationVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    operator_id: VLS,
    location_id: VLS,
}

impl UserInformationVLSData {
    pub fn new() -> Self {
        Self {
            size: 14u16.to_le(),
            r#type: USER_INFORMATION.to_le(),
            base_size: 14u16.to_le(),
            operator_id: crate::message::VLS::new(),
            location_id: crate::message::VLS::new(),
        }
    }
}

unsafe impl Send for UserInformationVLS {}
unsafe impl Send for UserInformationVLSUnsafe {}
unsafe impl Send for UserInformationVLSData {}

impl Drop for UserInformationVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for UserInformationVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for UserInformationVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for UserInformationVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for UserInformationVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for UserInformationVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for UserInformationVLS {
    type Target = UserInformationVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for UserInformationVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for UserInformationVLSUnsafe {
    type Target = UserInformationVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for UserInformationVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for UserInformationVLS {
    type Data = UserInformationVLSData;

    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, UserInformationVLSData::new()),
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
            data: data as *const UserInformationVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for UserInformationVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const UserInformationVLSData;
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
impl crate::Message for UserInformationVLSUnsafe {
    type Data = UserInformationVLSData;

    const BASE_SIZE: usize = 14;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, UserInformationVLSData::new()),
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
            data: data as *const UserInformationVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for UserInformationVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const UserInformationVLSData;
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
impl UserInformation for UserInformationVLS {
    type Safe = UserInformationVLS;
    type Unsafe = UserInformationVLSUnsafe;

    fn operator_id(&self) -> &str {
        get_vls(self, self.operator_id)
    }

    fn location_id(&self) -> &str {
        get_vls(self, self.location_id)
    }

    fn set_operator_id(&mut self, value: &str) -> &mut Self {
        self.operator_id = set_vls(self, self.operator_id, value);
        self
    }

    fn set_location_id(&mut self, value: &str) -> &mut Self {
        self.location_id = set_vls(self, self.location_id, value);
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

impl UserInformation for UserInformationVLSUnsafe {
    type Safe = UserInformationVLS;
    type Unsafe = UserInformationVLSUnsafe;

    fn operator_id(&self) -> &str {
        if self.is_out_of_bounds(10) {
            ""
        } else {
            get_vls(self, self.operator_id)
        }
    }

    fn location_id(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            get_vls(self, self.location_id)
        }
    }

    fn set_operator_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.operator_id = set_vls(self, self.operator_id, value);
        }
        self
    }

    fn set_location_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(14) {
            self.location_id = set_vls(self, self.location_id, value);
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
                14usize,
                core::mem::size_of::<UserInformationVLSData>(),
                "UserInformationVLSData sizeof expected {:} but was {:}",
                14usize,
                core::mem::size_of::<UserInformationVLSData>()
            );
            assert_eq!(
                14u16,
                UserInformationVLS::new().size(),
                "UserInformationVLS sizeof expected {:} but was {:}",
                14u16,
                UserInformationVLS::new().size(),
            );
            assert_eq!(
                USER_INFORMATION,
                UserInformationVLS::new().r#type(),
                "UserInformationVLS type expected {:} but was {:}",
                USER_INFORMATION,
                UserInformationVLS::new().r#type(),
            );
            assert_eq!(
                10136u16,
                UserInformationVLS::new().r#type(),
                "UserInformationVLS type expected {:} but was {:}",
                10136u16,
                UserInformationVLS::new().r#type(),
            );
            let d = UserInformationVLSData::new();
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
                (core::ptr::addr_of!(d.operator_id) as usize) - p,
                "operator_id offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.operator_id) as usize) - p,
            );
            assert_eq!(
                10usize,
                (core::ptr::addr_of!(d.location_id) as usize) - p,
                "location_id offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.location_id) as usize) - p,
            );
        }
    }
}
