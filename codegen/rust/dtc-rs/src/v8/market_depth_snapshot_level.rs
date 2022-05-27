// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-27 08:32:53.747629 +0800 WITA m=+0.008238335
use super::*;

pub(crate) const MARKET_DEPTH_SNAPSHOT_LEVEL_FIXED_SIZE: usize = 56;

/// size                       u16                       = MarketDepthSnapshotLevelFixedSize  (56)
/// type                       u16                       = MARKET_DEPTH_SNAPSHOT_LEVEL  (122)
/// symbol_id                  u32                       = 0
/// side                       AtBidOrAskEnum            = BID_ASK_UNSET  (0)
/// price                      f64                       = 0
/// quantity                   f64                       = 0
/// level                      u16                       = 0
/// is_first_message_in_batch  bool                      = false
/// is_last_message_in_batch   bool                      = false
/// date_time                  DateTimeWithMilliseconds  = 0
/// num_orders                 u32                       = 0
pub(crate) const MARKET_DEPTH_SNAPSHOT_LEVEL_FIXED_DEFAULT: [u8; 56] = [
    56, 0, 122, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

/// This is a message sent by Server to provide the initial market depth data
/// entries to the Client after the Client subscribes to market data or separately
/// subscribes to market depth data. The Client will need to separately subscribe
/// to market depth data if the Server requires it.
///
/// Each message provides a single entry of depth data. Therefore, the Server
/// will send multiple MarketDepthSnapshotLevelFixed messages in a series
/// in order for the Client to build up its initial market depth book.
///
/// The first message will be identified by the IsFirstMessageInBatch field
/// being set to 1. The last message will be identified by the IsLastMessageInBatch
/// field being set to 1.
///
/// In the case where the market depth book is empty, the Server still needs
/// to send through one single message with the SymbolID set, IsFirstMessageInBatch
/// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
/// be at the default values. The Client will understand this as an empty
/// book.
pub trait MarketDepthSnapshotLevel: Message {
    type Safe: MarketDepthSnapshotLevel;
    type Unsafe: MarketDepthSnapshotLevel;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS/MarketDepthRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32;

    /// Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
    /// = 2, if this is an ask side market depth entry.
    fn side(&self) -> AtBidOrAskEnum;

    /// This is the price of the market depth entry.
    fn price(&self) -> f64;

    /// This is the quantity of orders at the Price.
    fn quantity(&self) -> f64;

    /// This indicates the level of the price within the market depth book. The
    /// minimum value is 1. There is no maximum value. A value of 1 is considered
    /// the best bid or ask data.
    fn level(&self) -> u16;

    /// Set to 1 if this is the first message in the batch of messages.
    fn is_first_message_in_batch(&self) -> bool;

    /// Set to 1 if this is the last message in a batch of messages. If there
    /// is only a single message to be sent, in case the market depth book is
    /// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
    /// will equal 1.
    fn is_last_message_in_batch(&self) -> bool;

    fn date_time(&self) -> DateTimeWithMilliseconds;

    fn num_orders(&self) -> u32;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS/MarketDepthRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self;

    /// Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
    /// = 2, if this is an ask side market depth entry.
    fn set_side(&mut self, value: AtBidOrAskEnum) -> &mut Self;

    /// This is the price of the market depth entry.
    fn set_price(&mut self, value: f64) -> &mut Self;

    /// This is the quantity of orders at the Price.
    fn set_quantity(&mut self, value: f64) -> &mut Self;

    /// This indicates the level of the price within the market depth book. The
    /// minimum value is 1. There is no maximum value. A value of 1 is considered
    /// the best bid or ask data.
    fn set_level(&mut self, value: u16) -> &mut Self;

    /// Set to 1 if this is the first message in the batch of messages.
    fn set_is_first_message_in_batch(&mut self, value: bool) -> &mut Self;

    /// Set to 1 if this is the last message in a batch of messages. If there
    /// is only a single message to be sent, in case the market depth book is
    /// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
    /// will equal 1.
    fn set_is_last_message_in_batch(&mut self, value: bool) -> &mut Self;

    fn set_date_time(&mut self, value: DateTimeWithMilliseconds) -> &mut Self;

    fn set_num_orders(&mut self, value: u32) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl MarketDepthSnapshotLevel) {
        to.set_symbol_id(self.symbol_id());
        to.set_side(self.side());
        to.set_price(self.price());
        to.set_quantity(self.quantity());
        to.set_level(self.level());
        to.set_is_first_message_in_batch(self.is_first_message_in_batch());
        to.set_is_last_message_in_batch(self.is_last_message_in_batch());
        to.set_date_time(self.date_time());
        to.set_num_orders(self.num_orders());
    }

    #[inline]
    fn parse<F: Fn(Parsed<Self::Safe, Self::Unsafe>) -> Result<(), crate::Error>>(
        data: &[u8],
        f: F,
    ) -> Result<(), crate::Error> {
        if data.len() < 4 {
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

/// This is a message sent by Server to provide the initial market depth data
/// entries to the Client after the Client subscribes to market data or separately
/// subscribes to market depth data. The Client will need to separately subscribe
/// to market depth data if the Server requires it.
///
/// Each message provides a single entry of depth data. Therefore, the Server
/// will send multiple MarketDepthSnapshotLevelFixed messages in a series
/// in order for the Client to build up its initial market depth book.
///
/// The first message will be identified by the IsFirstMessageInBatch field
/// being set to 1. The last message will be identified by the IsLastMessageInBatch
/// field being set to 1.
///
/// In the case where the market depth book is empty, the Server still needs
/// to send through one single message with the SymbolID set, IsFirstMessageInBatch
/// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
/// be at the default values. The Client will understand this as an empty
/// book.
pub struct MarketDepthSnapshotLevelFixed {
    data: *const MarketDepthSnapshotLevelFixedData,
}

pub struct MarketDepthSnapshotLevelFixedUnsafe {
    data: *const MarketDepthSnapshotLevelFixedData,
}

#[repr(packed(8), C)]
pub struct MarketDepthSnapshotLevelFixedData {
    size: u16,
    r#type: u16,
    symbol_id: u32,
    side: AtBidOrAskEnum,
    price: f64,
    quantity: f64,
    level: u16,
    is_first_message_in_batch: bool,
    is_last_message_in_batch: bool,
    date_time: DateTimeWithMilliseconds,
    num_orders: u32,
}

impl MarketDepthSnapshotLevelFixedData {
    pub fn new() -> Self {
        Self {
            size: 56u16.to_le(),
            r#type: MARKET_DEPTH_SNAPSHOT_LEVEL.to_le(),
            symbol_id: 0u32,
            side: AtBidOrAskEnum::BidAskUnset.to_le(),
            price: 0.0f64,
            quantity: 0.0f64,
            level: 0u16,
            is_first_message_in_batch: false,
            is_last_message_in_batch: false,
            date_time: 0.0f64,
            num_orders: 0u32,
        }
    }
}

unsafe impl Send for MarketDepthSnapshotLevelFixed {}
unsafe impl Send for MarketDepthSnapshotLevelFixedUnsafe {}
unsafe impl Send for MarketDepthSnapshotLevelFixedData {}

impl Drop for MarketDepthSnapshotLevelFixed {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for MarketDepthSnapshotLevelFixedUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for MarketDepthSnapshotLevelFixed {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for MarketDepthSnapshotLevelFixedUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for MarketDepthSnapshotLevelFixed {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for MarketDepthSnapshotLevelFixedUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for MarketDepthSnapshotLevelFixed {
    type Target = MarketDepthSnapshotLevelFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDepthSnapshotLevelFixed {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for MarketDepthSnapshotLevelFixedUnsafe {
    type Target = MarketDepthSnapshotLevelFixedData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for MarketDepthSnapshotLevelFixedUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::fmt::Display for MarketDepthSnapshotLevelFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDepthSnapshotLevelFixed(size: {}, type: {}, symbol_id: {}, side: {}, price: {}, quantity: {}, level: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, date_time: {}, num_orders: {})", self.size(), self.r#type(), self.symbol_id(), self.side(), self.price(), self.quantity(), self.level(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.date_time(), self.num_orders()).as_str())
    }
}

impl core::fmt::Debug for MarketDepthSnapshotLevelFixed {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDepthSnapshotLevelFixed(size: {}, type: {}, symbol_id: {}, side: {}, price: {}, quantity: {}, level: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, date_time: {}, num_orders: {})", self.size(), self.r#type(), self.symbol_id(), self.side(), self.price(), self.quantity(), self.level(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.date_time(), self.num_orders()).as_str())
    }
}

impl core::fmt::Display for MarketDepthSnapshotLevelFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDepthSnapshotLevelFixedUnsafe(size: {}, type: {}, symbol_id: {}, side: {}, price: {}, quantity: {}, level: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, date_time: {}, num_orders: {})", self.size(), self.r#type(), self.symbol_id(), self.side(), self.price(), self.quantity(), self.level(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.date_time(), self.num_orders()).as_str())
    }
}

impl core::fmt::Debug for MarketDepthSnapshotLevelFixedUnsafe {
    #[inline]
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        f.write_str(format!("MarketDepthSnapshotLevelFixedUnsafe(size: {}, type: {}, symbol_id: {}, side: {}, price: {}, quantity: {}, level: {}, is_first_message_in_batch: {}, is_last_message_in_batch: {}, date_time: {}, num_orders: {})", self.size(), self.r#type(), self.symbol_id(), self.side(), self.price(), self.quantity(), self.level(), self.is_first_message_in_batch(), self.is_last_message_in_batch(), self.date_time(), self.num_orders()).as_str())
    }
}

impl crate::Message for MarketDepthSnapshotLevelFixed {
    type Data = MarketDepthSnapshotLevelFixedData;

    const TYPE: u16 = MARKET_DEPTH_SNAPSHOT_LEVEL;
    const BASE_SIZE: usize = 56;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDepthSnapshotLevelFixedData::new()),
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
            data: data as *const MarketDepthSnapshotLevelFixedData,
        }
    }
}
impl crate::Message for MarketDepthSnapshotLevelFixedUnsafe {
    type Data = MarketDepthSnapshotLevelFixedData;

    const TYPE: u16 = MARKET_DEPTH_SNAPSHOT_LEVEL;
    const BASE_SIZE: usize = 56;
    const BASE_SIZE_OFFSET: isize = 0;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(Self::BASE_SIZE, MarketDepthSnapshotLevelFixedData::new()),
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
            data: data as *const MarketDepthSnapshotLevelFixedData,
        }
    }
}
/// This is a message sent by Server to provide the initial market depth data
/// entries to the Client after the Client subscribes to market data or separately
/// subscribes to market depth data. The Client will need to separately subscribe
/// to market depth data if the Server requires it.
///
/// Each message provides a single entry of depth data. Therefore, the Server
/// will send multiple MarketDepthSnapshotLevelFixed messages in a series
/// in order for the Client to build up its initial market depth book.
///
/// The first message will be identified by the IsFirstMessageInBatch field
/// being set to 1. The last message will be identified by the IsLastMessageInBatch
/// field being set to 1.
///
/// In the case where the market depth book is empty, the Server still needs
/// to send through one single message with the SymbolID set, IsFirstMessageInBatch
/// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
/// be at the default values. The Client will understand this as an empty
/// book.
impl MarketDepthSnapshotLevel for MarketDepthSnapshotLevelFixed {
    type Safe = MarketDepthSnapshotLevelFixed;
    type Unsafe = MarketDepthSnapshotLevelFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS/MarketDepthRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        u32::from_le(self.symbol_id)
    }

    /// Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
    /// = 2, if this is an ask side market depth entry.
    fn side(&self) -> AtBidOrAskEnum {
        AtBidOrAskEnum::from_le(self.side)
    }

    /// This is the price of the market depth entry.
    fn price(&self) -> f64 {
        f64_le(self.price)
    }

    /// This is the quantity of orders at the Price.
    fn quantity(&self) -> f64 {
        f64_le(self.quantity)
    }

    /// This indicates the level of the price within the market depth book. The
    /// minimum value is 1. There is no maximum value. A value of 1 is considered
    /// the best bid or ask data.
    fn level(&self) -> u16 {
        u16::from_le(self.level)
    }

    /// Set to 1 if this is the first message in the batch of messages.
    fn is_first_message_in_batch(&self) -> bool {
        self.is_first_message_in_batch
    }

    /// Set to 1 if this is the last message in a batch of messages. If there
    /// is only a single message to be sent, in case the market depth book is
    /// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
    /// will equal 1.
    fn is_last_message_in_batch(&self) -> bool {
        self.is_last_message_in_batch
    }

    fn date_time(&self) -> DateTimeWithMilliseconds {
        f64_le(self.date_time)
    }

    fn num_orders(&self) -> u32 {
        u32::from_le(self.num_orders)
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS/MarketDepthRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        self.symbol_id = value.to_le();
        self
    }

    /// Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
    /// = 2, if this is an ask side market depth entry.
    fn set_side(&mut self, value: AtBidOrAskEnum) -> &mut Self {
        self.side = unsafe { core::mem::transmute((value as u16).to_le()) };
        self
    }

    /// This is the price of the market depth entry.
    fn set_price(&mut self, value: f64) -> &mut Self {
        self.price = f64_le(value);
        self
    }

    /// This is the quantity of orders at the Price.
    fn set_quantity(&mut self, value: f64) -> &mut Self {
        self.quantity = f64_le(value);
        self
    }

    /// This indicates the level of the price within the market depth book. The
    /// minimum value is 1. There is no maximum value. A value of 1 is considered
    /// the best bid or ask data.
    fn set_level(&mut self, value: u16) -> &mut Self {
        self.level = value.to_le();
        self
    }

    /// Set to 1 if this is the first message in the batch of messages.
    fn set_is_first_message_in_batch(&mut self, value: bool) -> &mut Self {
        self.is_first_message_in_batch = value;
        self
    }

    /// Set to 1 if this is the last message in a batch of messages. If there
    /// is only a single message to be sent, in case the market depth book is
    /// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
    /// will equal 1.
    fn set_is_last_message_in_batch(&mut self, value: bool) -> &mut Self {
        self.is_last_message_in_batch = value;
        self
    }

    fn set_date_time(&mut self, value: DateTimeWithMilliseconds) -> &mut Self {
        self.date_time = f64_le(value);
        self
    }

    fn set_num_orders(&mut self, value: u32) -> &mut Self {
        self.num_orders = value.to_le();
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

/// This is a message sent by Server to provide the initial market depth data
/// entries to the Client after the Client subscribes to market data or separately
/// subscribes to market depth data. The Client will need to separately subscribe
/// to market depth data if the Server requires it.
///
/// Each message provides a single entry of depth data. Therefore, the Server
/// will send multiple MarketDepthSnapshotLevelFixed messages in a series
/// in order for the Client to build up its initial market depth book.
///
/// The first message will be identified by the IsFirstMessageInBatch field
/// being set to 1. The last message will be identified by the IsLastMessageInBatch
/// field being set to 1.
///
/// In the case where the market depth book is empty, the Server still needs
/// to send through one single message with the SymbolID set, IsFirstMessageInBatch
/// equal to 1 and IsLastMessageInBatch equal to 1. All other members will
/// be at the default values. The Client will understand this as an empty
/// book.
impl MarketDepthSnapshotLevel for MarketDepthSnapshotLevelFixedUnsafe {
    type Safe = MarketDepthSnapshotLevelFixed;
    type Unsafe = MarketDepthSnapshotLevelFixedUnsafe;

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS/MarketDepthRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn symbol_id(&self) -> u32 {
        if self.is_out_of_bounds(8) {
            0u32
        } else {
            u32::from_le(self.symbol_id)
        }
    }

    /// Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
    /// = 2, if this is an ask side market depth entry.
    fn side(&self) -> AtBidOrAskEnum {
        if self.is_out_of_bounds(10) {
            AtBidOrAskEnum::BidAskUnset.to_le()
        } else {
            AtBidOrAskEnum::from_le(self.side)
        }
    }

    /// This is the price of the market depth entry.
    fn price(&self) -> f64 {
        if self.is_out_of_bounds(24) {
            0.0f64
        } else {
            f64_le(self.price)
        }
    }

    /// This is the quantity of orders at the Price.
    fn quantity(&self) -> f64 {
        if self.is_out_of_bounds(32) {
            0.0f64
        } else {
            f64_le(self.quantity)
        }
    }

    /// This indicates the level of the price within the market depth book. The
    /// minimum value is 1. There is no maximum value. A value of 1 is considered
    /// the best bid or ask data.
    fn level(&self) -> u16 {
        if self.is_out_of_bounds(34) {
            0u16
        } else {
            u16::from_le(self.level)
        }
    }

    /// Set to 1 if this is the first message in the batch of messages.
    fn is_first_message_in_batch(&self) -> bool {
        if self.is_out_of_bounds(35) {
            false
        } else {
            self.is_first_message_in_batch
        }
    }

    /// Set to 1 if this is the last message in a batch of messages. If there
    /// is only a single message to be sent, in case the market depth book is
    /// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
    /// will equal 1.
    fn is_last_message_in_batch(&self) -> bool {
        if self.is_out_of_bounds(36) {
            false
        } else {
            self.is_last_message_in_batch
        }
    }

    fn date_time(&self) -> DateTimeWithMilliseconds {
        if self.is_out_of_bounds(48) {
            0.0f64
        } else {
            f64_le(self.date_time)
        }
    }

    fn num_orders(&self) -> u32 {
        if self.is_out_of_bounds(52) {
            0u32
        } else {
            u32::from_le(self.num_orders)
        }
    }

    /// This is the same SymbolID sent by the Client in the MarketDataRequestVLS/MarketDepthRequestVLS
    /// message which corresponds to the Symbol that the data in this message
    /// is for.
    fn set_symbol_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(8) {
            self.symbol_id = value.to_le();
        }
        self
    }

    /// Set to AT_BID = 1 if this is a bid side market depth entry. Set to AT_ASK
    /// = 2, if this is an ask side market depth entry.
    fn set_side(&mut self, value: AtBidOrAskEnum) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.side = unsafe { core::mem::transmute((value as u16).to_le()) };
        }
        self
    }

    /// This is the price of the market depth entry.
    fn set_price(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(24) {
            self.price = f64_le(value);
        }
        self
    }

    /// This is the quantity of orders at the Price.
    fn set_quantity(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(32) {
            self.quantity = f64_le(value);
        }
        self
    }

    /// This indicates the level of the price within the market depth book. The
    /// minimum value is 1. There is no maximum value. A value of 1 is considered
    /// the best bid or ask data.
    fn set_level(&mut self, value: u16) -> &mut Self {
        if !self.is_out_of_bounds(34) {
            self.level = value.to_le();
        }
        self
    }

    /// Set to 1 if this is the first message in the batch of messages.
    fn set_is_first_message_in_batch(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(35) {
            self.is_first_message_in_batch = value;
        }
        self
    }

    /// Set to 1 if this is the last message in a batch of messages. If there
    /// is only a single message to be sent, in case the market depth book is
    /// empty, then IsFirstMessageInBatch will equal 1 and IsLastMessageInBatch
    /// will equal 1.
    fn set_is_last_message_in_batch(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(36) {
            self.is_last_message_in_batch = value;
        }
        self
    }

    fn set_date_time(&mut self, value: DateTimeWithMilliseconds) -> &mut Self {
        if !self.is_out_of_bounds(48) {
            self.date_time = f64_le(value);
        }
        self
    }

    fn set_num_orders(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(52) {
            self.num_orders = value.to_le();
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
                56usize,
                core::mem::size_of::<MarketDepthSnapshotLevelFixedData>(),
                "MarketDepthSnapshotLevelFixedData sizeof expected {:} but was {:}",
                56usize,
                core::mem::size_of::<MarketDepthSnapshotLevelFixedData>()
            );
            assert_eq!(
                56u16,
                MarketDepthSnapshotLevelFixed::new().size(),
                "MarketDepthSnapshotLevelFixed sizeof expected {:} but was {:}",
                56u16,
                MarketDepthSnapshotLevelFixed::new().size(),
            );
            assert_eq!(
                MARKET_DEPTH_SNAPSHOT_LEVEL,
                MarketDepthSnapshotLevelFixed::new().r#type(),
                "MarketDepthSnapshotLevelFixed type expected {:} but was {:}",
                MARKET_DEPTH_SNAPSHOT_LEVEL,
                MarketDepthSnapshotLevelFixed::new().r#type(),
            );
            assert_eq!(
                122u16,
                MarketDepthSnapshotLevelFixed::new().r#type(),
                "MarketDepthSnapshotLevelFixed type expected {:} but was {:}",
                122u16,
                MarketDepthSnapshotLevelFixed::new().r#type(),
            );
            let d = MarketDepthSnapshotLevelFixedData::new();
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
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
                "symbol_id offset expected {:} but was {:}",
                4usize,
                (core::ptr::addr_of!(d.symbol_id) as usize) - p,
            );
            assert_eq!(
                8usize,
                (core::ptr::addr_of!(d.side) as usize) - p,
                "side offset expected {:} but was {:}",
                8usize,
                (core::ptr::addr_of!(d.side) as usize) - p,
            );
            assert_eq!(
                16usize,
                (core::ptr::addr_of!(d.price) as usize) - p,
                "price offset expected {:} but was {:}",
                16usize,
                (core::ptr::addr_of!(d.price) as usize) - p,
            );
            assert_eq!(
                24usize,
                (core::ptr::addr_of!(d.quantity) as usize) - p,
                "quantity offset expected {:} but was {:}",
                24usize,
                (core::ptr::addr_of!(d.quantity) as usize) - p,
            );
            assert_eq!(
                32usize,
                (core::ptr::addr_of!(d.level) as usize) - p,
                "level offset expected {:} but was {:}",
                32usize,
                (core::ptr::addr_of!(d.level) as usize) - p,
            );
            assert_eq!(
                34usize,
                (core::ptr::addr_of!(d.is_first_message_in_batch) as usize) - p,
                "is_first_message_in_batch offset expected {:} but was {:}",
                34usize,
                (core::ptr::addr_of!(d.is_first_message_in_batch) as usize) - p,
            );
            assert_eq!(
                35usize,
                (core::ptr::addr_of!(d.is_last_message_in_batch) as usize) - p,
                "is_last_message_in_batch offset expected {:} but was {:}",
                35usize,
                (core::ptr::addr_of!(d.is_last_message_in_batch) as usize) - p,
            );
            assert_eq!(
                40usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
                "date_time offset expected {:} but was {:}",
                40usize,
                (core::ptr::addr_of!(d.date_time) as usize) - p,
            );
            assert_eq!(
                48usize,
                (core::ptr::addr_of!(d.num_orders) as usize) - p,
                "num_orders offset expected {:} but was {:}",
                48usize,
                (core::ptr::addr_of!(d.num_orders) as usize) - p,
            );
        }
    }
}
