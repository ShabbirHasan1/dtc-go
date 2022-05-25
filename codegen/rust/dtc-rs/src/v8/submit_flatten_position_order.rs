// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const SUBMIT_FLATTEN_POSITION_ORDER_VLS_SIZE: usize = 28;

pub(crate) const SUBMIT_FLATTEN_POSITION_ORDER_FIXED_SIZE: usize = 198;

/// size                u16     = SubmitFlattenPositionOrderVLSSize  (28)
/// type                u16     = SUBMIT_FLATTEN_POSITION_ORDER  (209)
/// base_size           u16     = SubmitFlattenPositionOrderVLSSize  (28)
/// symbol              string  = ""
/// exchange            string  = ""
/// trade_account       string  = ""
/// client_order_id     string  = ""
/// free_form_text      string  = ""
/// is_automated_order  bool    = false
pub(crate) const SUBMIT_FLATTEN_POSITION_ORDER_VLS_DEFAULT: [u8; 28] = [
    28, 0, 209, 0, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// size                u16       = SubmitFlattenPositionOrderFixedSize  (198)
/// type                u16       = SUBMIT_FLATTEN_POSITION_ORDER  (209)
/// symbol              string64  = ""
/// exchange            string16  = ""
/// trade_account       string32  = ""
/// client_order_id     string32  = ""
/// free_form_text      string48  = ""
/// is_automated_order  bool      = false
pub(crate) const SUBMIT_FLATTEN_POSITION_ORDER_FIXED_DEFAULT: [u8; 198] = [
    198, 0, 209, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait SubmitFlattenPositionOrder: Message {
    type Safe: SubmitFlattenPositionOrder;
    type Unsafe: SubmitFlattenPositionOrder;

    /// The symbol of the Trade Position to flatten.
    fn symbol(&self) -> &str;

    /// The optional exchange for the Symbol.
    fn exchange(&self) -> &str;

    /// The trade account as a text string of the Trade Position to flatten.
    fn trade_account(&self) -> &str;

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn client_order_id(&self) -> &str;

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn free_form_text(&self) -> &str;

    /// Set to 1 for an order submitted by an automated trading system.
    fn is_automated_order(&self) -> bool;

    /// The symbol of the Trade Position to flatten.
    fn set_symbol(&mut self, value: &str) -> &mut Self;

    /// The optional exchange for the Symbol.
    fn set_exchange(&mut self, value: &str) -> &mut Self;

    /// The trade account as a text string of the Trade Position to flatten.
    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn set_client_order_id(&mut self, value: &str) -> &mut Self;

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn set_free_form_text(&mut self, value: &str) -> &mut Self;

    /// Set to 1 for an order submitted by an automated trading system.
    fn set_is_automated_order(&mut self, value: bool) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl SubmitFlattenPositionOrder) {
        to.set_symbol(self.symbol());
        to.set_exchange(self.exchange());
        to.set_trade_account(self.trade_account());
        to.set_client_order_id(self.client_order_id());
        to.set_free_form_text(self.free_form_text());
        to.set_is_automated_order(self.is_automated_order());
    }
}

pub struct SubmitFlattenPositionOrderVLS {
    data: *const SubmitFlattenPositionOrderVLSData,
    capacity: usize,
}

pub struct SubmitFlattenPositionOrderVLSUnsafe {
    data: *const SubmitFlattenPositionOrderVLSData,
    capacity: usize,
}

#[repr(packed(8), C)]
pub struct SubmitFlattenPositionOrderVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    symbol: VLS,
    exchange: VLS,
    trade_account: VLS,
    client_order_id: VLS,
    free_form_text: VLS,
    is_automated_order: bool,
}

pub struct SubmitFlattenPositionOrderFixed {
    data: *const SubmitFlattenPositionOrderFixedData,
}

pub struct SubmitFlattenPositionOrderFixedUnsafe {
    data: *const SubmitFlattenPositionOrderFixedData,
}

#[repr(packed(8), C)]
pub struct SubmitFlattenPositionOrderFixedData {
    size: u16,
    r#type: u16,
    symbol: [u8; 64],
    exchange: [u8; 16],
    trade_account: [u8; 32],
    client_order_id: [u8; 32],
    free_form_text: [u8; 48],
    is_automated_order: bool,
}

impl SubmitFlattenPositionOrderVLSData {
    pub fn new() -> Self {
        Self {
            size: 28u16.to_le(),
            r#type: SUBMIT_FLATTEN_POSITION_ORDER.to_le(),
            base_size: 28u16.to_le(),
            symbol: crate::message::VLS::new(),
            exchange: crate::message::VLS::new(),
            trade_account: crate::message::VLS::new(),
            client_order_id: crate::message::VLS::new(),
            free_form_text: crate::message::VLS::new(),
            is_automated_order: false,
        }
    }
}

impl SubmitFlattenPositionOrderFixedData {
    pub fn new() -> Self {
        Self {
            size: 198u16.to_le(),
            r#type: SUBMIT_FLATTEN_POSITION_ORDER.to_le(),
            symbol: [0; 64],
            exchange: [0; 16],
            trade_account: [0; 32],
            client_order_id: [0; 32],
            free_form_text: [0; 48],
            is_automated_order: false,
        }
    }
}

unsafe impl Send for SubmitFlattenPositionOrderFixed {}
unsafe impl Send for SubmitFlattenPositionOrderFixedUnsafe {}
unsafe impl Send for SubmitFlattenPositionOrderFixedData {}
unsafe impl Send for SubmitFlattenPositionOrderVLS {}
unsafe impl Send for SubmitFlattenPositionOrderVLSUnsafe {}
unsafe impl Send for SubmitFlattenPositionOrderVLSData {}

impl Drop for SubmitFlattenPositionOrderFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for SubmitFlattenPositionOrderFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for SubmitFlattenPositionOrderVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for SubmitFlattenPositionOrderVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for SubmitFlattenPositionOrderFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for SubmitFlattenPositionOrderFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for SubmitFlattenPositionOrderVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for SubmitFlattenPositionOrderVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for SubmitFlattenPositionOrderFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for SubmitFlattenPositionOrderFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for SubmitFlattenPositionOrderVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for SubmitFlattenPositionOrderVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for SubmitFlattenPositionOrderFixed {
    type Target = SubmitFlattenPositionOrderFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for SubmitFlattenPositionOrderFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for SubmitFlattenPositionOrderFixedUnsafe {
    type Target = SubmitFlattenPositionOrderFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for SubmitFlattenPositionOrderFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for SubmitFlattenPositionOrderVLS {
    type Target = SubmitFlattenPositionOrderVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for SubmitFlattenPositionOrderVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for SubmitFlattenPositionOrderVLSUnsafe {
    type Target = SubmitFlattenPositionOrderVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for SubmitFlattenPositionOrderVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for SubmitFlattenPositionOrderFixed {
    type Data = SubmitFlattenPositionOrderFixedData;

    const BASE_SIZE: usize = 198;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, SubmitFlattenPositionOrderFixedData::new()),
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
            data: data as *const SubmitFlattenPositionOrderFixedData,
        }
    }
}
impl crate::Message for SubmitFlattenPositionOrderFixedUnsafe {
    type Data = SubmitFlattenPositionOrderFixedData;

    const BASE_SIZE: usize = 198;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, SubmitFlattenPositionOrderFixedData::new()),
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
            data: data as *const SubmitFlattenPositionOrderFixedData,
        }
    }
}
impl crate::Message for SubmitFlattenPositionOrderVLS {
    type Data = SubmitFlattenPositionOrderVLSData;

    const BASE_SIZE: usize = 28;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, SubmitFlattenPositionOrderVLSData::new()),
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
            data: data as *const SubmitFlattenPositionOrderVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for SubmitFlattenPositionOrderVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const SubmitFlattenPositionOrderVLSData;
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
impl crate::Message for SubmitFlattenPositionOrderVLSUnsafe {
    type Data = SubmitFlattenPositionOrderVLSData;

    const BASE_SIZE: usize = 28;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, SubmitFlattenPositionOrderVLSData::new()),
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
            data: data as *const SubmitFlattenPositionOrderVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for SubmitFlattenPositionOrderVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const SubmitFlattenPositionOrderVLSData;
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
impl SubmitFlattenPositionOrder for SubmitFlattenPositionOrderVLS {
    type Safe = SubmitFlattenPositionOrderVLS;
    type Unsafe = SubmitFlattenPositionOrderVLSUnsafe;

    /// The symbol of the Trade Position to flatten.
    fn symbol(&self) -> &str {
        get_vls(self, self.symbol)
    }

    /// The optional exchange for the Symbol.
    fn exchange(&self) -> &str {
        get_vls(self, self.exchange)
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn client_order_id(&self) -> &str {
        get_vls(self, self.client_order_id)
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn free_form_text(&self) -> &str {
        get_vls(self, self.free_form_text)
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn is_automated_order(&self) -> bool {
        self.is_automated_order
    }

    /// The symbol of the Trade Position to flatten.
    fn set_symbol(&mut self, value: &str) -> &mut Self {
        self.symbol = set_vls(self, self.symbol, value);
        self
    }

    /// The optional exchange for the Symbol.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        self.exchange = set_vls(self, self.exchange, value);
        self
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = set_vls(self, self.trade_account, value);
        self
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        self.client_order_id = set_vls(self, self.client_order_id, value);
        self
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        self.free_form_text = set_vls(self, self.free_form_text, value);
        self
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn set_is_automated_order(&mut self, value: bool) -> &mut Self {
        self.is_automated_order = value;
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

impl SubmitFlattenPositionOrder for SubmitFlattenPositionOrderVLSUnsafe {
    type Safe = SubmitFlattenPositionOrderVLS;
    type Unsafe = SubmitFlattenPositionOrderVLSUnsafe;

    /// The symbol of the Trade Position to flatten.
    fn symbol(&self) -> &str {
        if self.is_out_of_bounds(10) {
            ""
        } else {
            get_vls(self, self.symbol)
        }
    }

    /// The optional exchange for the Symbol.
    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(14) {
            ""
        } else {
            get_vls(self, self.exchange)
        }
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(18) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn client_order_id(&self) -> &str {
        if self.is_out_of_bounds(22) {
            ""
        } else {
            get_vls(self, self.client_order_id)
        }
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn free_form_text(&self) -> &str {
        if self.is_out_of_bounds(26) {
            ""
        } else {
            get_vls(self, self.free_form_text)
        }
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn is_automated_order(&self) -> bool {
        if self.is_out_of_bounds(27) {
            false
        } else {
            self.is_automated_order
        }
    }

    /// The symbol of the Trade Position to flatten.
    fn set_symbol(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.symbol = set_vls(self, self.symbol, value);
        }
        self
    }

    /// The optional exchange for the Symbol.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(14) {
            self.exchange = set_vls(self, self.exchange, value);
        }
        self
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(18) {
            self.trade_account = set_vls(self, self.trade_account, value);
        }
        self
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(22) {
            self.client_order_id = set_vls(self, self.client_order_id, value);
        }
        self
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(26) {
            self.free_form_text = set_vls(self, self.free_form_text, value);
        }
        self
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn set_is_automated_order(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(27) {
            self.is_automated_order = value;
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

impl SubmitFlattenPositionOrder for SubmitFlattenPositionOrderFixed {
    type Safe = SubmitFlattenPositionOrderFixed;
    type Unsafe = SubmitFlattenPositionOrderFixedUnsafe;

    /// The symbol of the Trade Position to flatten.
    fn symbol(&self) -> &str {
        get_fixed(&self.symbol[..])
    }

    /// The optional exchange for the Symbol.
    fn exchange(&self) -> &str {
        get_fixed(&self.exchange[..])
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn trade_account(&self) -> &str {
        get_fixed(&self.trade_account[..])
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn client_order_id(&self) -> &str {
        get_fixed(&self.client_order_id[..])
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn free_form_text(&self) -> &str {
        get_fixed(&self.free_form_text[..])
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn is_automated_order(&self) -> bool {
        self.is_automated_order
    }

    /// The symbol of the Trade Position to flatten.
    fn set_symbol(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.symbol[..], value);
        self
    }

    /// The optional exchange for the Symbol.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.exchange[..], value);
        self
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.trade_account[..], value);
        self
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.client_order_id[..], value);
        self
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        set_fixed(&mut self.free_form_text[..], value);
        self
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn set_is_automated_order(&mut self, value: bool) -> &mut Self {
        self.is_automated_order = value;
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

impl SubmitFlattenPositionOrder for SubmitFlattenPositionOrderFixedUnsafe {
    type Safe = SubmitFlattenPositionOrderFixed;
    type Unsafe = SubmitFlattenPositionOrderFixedUnsafe;

    /// The symbol of the Trade Position to flatten.
    fn symbol(&self) -> &str {
        if self.is_out_of_bounds(68) {
            ""
        } else {
            get_fixed(&self.symbol[..])
        }
    }

    /// The optional exchange for the Symbol.
    fn exchange(&self) -> &str {
        if self.is_out_of_bounds(84) {
            ""
        } else {
            get_fixed(&self.exchange[..])
        }
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(116) {
            ""
        } else {
            get_fixed(&self.trade_account[..])
        }
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn client_order_id(&self) -> &str {
        if self.is_out_of_bounds(148) {
            ""
        } else {
            get_fixed(&self.client_order_id[..])
        }
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn free_form_text(&self) -> &str {
        if self.is_out_of_bounds(196) {
            ""
        } else {
            get_fixed(&self.free_form_text[..])
        }
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn is_automated_order(&self) -> bool {
        if self.is_out_of_bounds(197) {
            false
        } else {
            self.is_automated_order
        }
    }

    /// The symbol of the Trade Position to flatten.
    fn set_symbol(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(68) {
            set_fixed(&mut self.symbol[..], value);
        }
        self
    }

    /// The optional exchange for the Symbol.
    fn set_exchange(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(84) {
            set_fixed(&mut self.exchange[..], value);
        }
        self
    }

    /// The trade account as a text string of the Trade Position to flatten.
    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(116) {
            set_fixed(&mut self.trade_account[..], value);
        }
        self
    }

    /// The Client supplied order identifier for the order which will be created
    /// to flatten the Trade Position.
    ///
    /// The Server will remember this for the life of the order.
    fn set_client_order_id(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(148) {
            set_fixed(&mut self.client_order_id[..], value);
        }
        self
    }

    /// Optional: This is an optional text string which can be set by the Client
    /// to associate text with the order which will be created to flatten the
    /// Trade Position.
    ///
    /// The Server is not under any obligation to use this text and it may place
    /// a limitation on the length of this text.
    fn set_free_form_text(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(196) {
            set_fixed(&mut self.free_form_text[..], value);
        }
        self
    }

    /// Set to 1 for an order submitted by an automated trading system.
    fn set_is_automated_order(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(197) {
            self.is_automated_order = value;
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
                198usize,
                core::mem::size_of::<SubmitFlattenPositionOrderFixedData>(),
                "SubmitFlattenPositionOrderFixedData sizeof expected {:} but was {:}",
                198usize,
                core::mem::size_of::<SubmitFlattenPositionOrderFixedData>()
            );
            assert_eq!(
                198u16,
                SubmitFlattenPositionOrderFixed::new().size(),
                "SubmitFlattenPositionOrderFixed sizeof expected {:} but was {:}",
                198u16,
                SubmitFlattenPositionOrderFixed::new().size(),
            );
            assert_eq!(
                SUBMIT_FLATTEN_POSITION_ORDER,
                SubmitFlattenPositionOrderFixed::new().r#type(),
                "SubmitFlattenPositionOrderFixed type expected {:} but was {:}",
                SUBMIT_FLATTEN_POSITION_ORDER,
                SubmitFlattenPositionOrderFixed::new().r#type(),
            );
            assert_eq!(
                209u16,
                SubmitFlattenPositionOrderFixed::new().r#type(),
                "SubmitFlattenPositionOrderFixed type expected {:} but was {:}",
                209u16,
                SubmitFlattenPositionOrderFixed::new().r#type(),
            );
            let d = SubmitFlattenPositionOrderFixedData::new();
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
                (core::ptr::addr_of!(d.symbol) as usize) - p,
                "symbol offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.symbol) as usize) - p,
            );
            assert_eq!(
                68usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
                "exchange offset expected {:} but was {:}",
                68usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
            );
            assert_eq!(
                84usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                84usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                116usize,
                (core::ptr::addr_of!(d.client_order_id) as usize) - p,
                "client_order_id offset expected {:} but was {:}",
                116usize,
                (core::ptr::addr_of!(d.client_order_id) as usize) - p,
            );
            assert_eq!(
                148usize,
                (core::ptr::addr_of!(d.free_form_text) as usize) - p,
                "free_form_text offset expected {:} but was {:}",
                148usize,
                (core::ptr::addr_of!(d.free_form_text) as usize) - p,
            );
            assert_eq!(
                196usize,
                (core::ptr::addr_of!(d.is_automated_order) as usize) - p,
                "is_automated_order offset expected {:} but was {:}",
                196usize,
                (core::ptr::addr_of!(d.is_automated_order) as usize) - p,
            );
        }
        unsafe {
            assert_eq!(
                28usize,
                core::mem::size_of::<SubmitFlattenPositionOrderVLSData>(),
                "SubmitFlattenPositionOrderVLSData sizeof expected {:} but was {:}",
                28usize,
                core::mem::size_of::<SubmitFlattenPositionOrderVLSData>()
            );
            assert_eq!(
                28u16,
                SubmitFlattenPositionOrderVLS::new().size(),
                "SubmitFlattenPositionOrderVLS sizeof expected {:} but was {:}",
                28u16,
                SubmitFlattenPositionOrderVLS::new().size(),
            );
            assert_eq!(
                SUBMIT_FLATTEN_POSITION_ORDER,
                SubmitFlattenPositionOrderVLS::new().r#type(),
                "SubmitFlattenPositionOrderVLS type expected {:} but was {:}",
                SUBMIT_FLATTEN_POSITION_ORDER,
                SubmitFlattenPositionOrderVLS::new().r#type(),
            );
            assert_eq!(
                209u16,
                SubmitFlattenPositionOrderVLS::new().r#type(),
                "SubmitFlattenPositionOrderVLS type expected {:} but was {:}",
                209u16,
                SubmitFlattenPositionOrderVLS::new().r#type(),
            );
            let d = SubmitFlattenPositionOrderVLSData::new();
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
                (core::ptr::addr_of!(d.symbol) as usize) - p,
                "symbol offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.symbol) as usize) - p,
            );
            assert_eq!(
                10usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
                "exchange offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.exchange) as usize) - p,
            );
            assert_eq!(
                14usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                14usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                18usize,
                (core::ptr::addr_of!(d.client_order_id) as usize) - p,
                "client_order_id offset expected {:} but was {:}",
                18usize,
                (core::ptr::addr_of!(d.client_order_id) as usize) - p,
            );
            assert_eq!(
                22usize,
                (core::ptr::addr_of!(d.free_form_text) as usize) - p,
                "free_form_text offset expected {:} but was {:}",
                22usize,
                (core::ptr::addr_of!(d.free_form_text) as usize) - p,
            );
            assert_eq!(
                26usize,
                (core::ptr::addr_of!(d.is_automated_order) as usize) - p,
                "is_automated_order offset expected {:} but was {:}",
                26usize,
                (core::ptr::addr_of!(d.is_automated_order) as usize) - p,
            );
        }
    }
}