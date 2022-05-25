// generated by github.com/moontrade/dtc-go/codegen/go at 2022-05-25 15:25:42.126453 +0800 WITA m=+0.007296918
use super::*;

pub(crate) const TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE_VLS_SIZE: usize = 52;

/// size                                                        u16     = TradeAccountDataSymbolLimitsUpdateVLSSize  (52)
/// type                                                        u16     = TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE  (10122)
/// base_size                                                   u16     = TradeAccountDataSymbolLimitsUpdateVLSSize  (52)
/// request_id                                                  u32     = 0
/// is_deleted                                                  bool    = false
/// trade_account                                               string  = ""
/// symbol                                                      string  = ""
/// trade_position_limit                                        f64     = 0
/// order_quantity_limit                                        f64     = 0
/// use_percent_of_margin                                       i32     = 100
/// override_margin_other_accounts                              u8      = 0
/// use_percent_of_margin_for_day_trading                       i32     = 0
/// number_of_days_before_last_trading_date_to_disallow_orders  i32     = 0
/// fixed_margin_cash_value                                     f32     = 0
pub(crate) const TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE_VLS_DEFAULT: [u8; 52] = [
    52, 0, 138, 39, 52, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 100, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
];

pub trait TradeAccountDataSymbolLimitsUpdate: Message {
    type Safe: TradeAccountDataSymbolLimitsUpdate;
    type Unsafe: TradeAccountDataSymbolLimitsUpdate;

    fn request_id(&self) -> u32;

    fn is_deleted(&self) -> bool;

    fn trade_account(&self) -> &str;

    fn symbol(&self) -> &str;

    fn trade_position_limit(&self) -> f64;

    fn order_quantity_limit(&self) -> f64;

    fn use_percent_of_margin(&self) -> i32;

    fn override_margin_other_accounts(&self) -> u8;

    fn use_percent_of_margin_for_day_trading(&self) -> i32;

    fn number_of_days_before_last_trading_date_to_disallow_orders(&self) -> i32;

    fn fixed_margin_cash_value(&self) -> f32;

    fn set_request_id(&mut self, value: u32) -> &mut Self;

    fn set_is_deleted(&mut self, value: bool) -> &mut Self;

    fn set_trade_account(&mut self, value: &str) -> &mut Self;

    fn set_symbol(&mut self, value: &str) -> &mut Self;

    fn set_trade_position_limit(&mut self, value: f64) -> &mut Self;

    fn set_order_quantity_limit(&mut self, value: f64) -> &mut Self;

    fn set_use_percent_of_margin(&mut self, value: i32) -> &mut Self;

    fn set_override_margin_other_accounts(&mut self, value: u8) -> &mut Self;

    fn set_use_percent_of_margin_for_day_trading(&mut self, value: i32) -> &mut Self;

    fn set_number_of_days_before_last_trading_date_to_disallow_orders(
        &mut self,
        value: i32,
    ) -> &mut Self;

    fn set_fixed_margin_cash_value(&mut self, value: f32) -> &mut Self;

    fn clone_safe(&self) -> Self::Safe;

    fn to_safe(self) -> Self::Safe;

    fn copy_to(&self, to: &mut impl TradeAccountDataSymbolLimitsUpdate) {
        to.set_request_id(self.request_id());
        to.set_is_deleted(self.is_deleted());
        to.set_trade_account(self.trade_account());
        to.set_symbol(self.symbol());
        to.set_trade_position_limit(self.trade_position_limit());
        to.set_order_quantity_limit(self.order_quantity_limit());
        to.set_use_percent_of_margin(self.use_percent_of_margin());
        to.set_override_margin_other_accounts(self.override_margin_other_accounts());
        to.set_use_percent_of_margin_for_day_trading(self.use_percent_of_margin_for_day_trading());
        to.set_number_of_days_before_last_trading_date_to_disallow_orders(
            self.number_of_days_before_last_trading_date_to_disallow_orders(),
        );
        to.set_fixed_margin_cash_value(self.fixed_margin_cash_value());
    }
}

pub struct TradeAccountDataSymbolLimitsUpdateVLS {
    data: *const TradeAccountDataSymbolLimitsUpdateVLSData,
    capacity: usize,
}

pub struct TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    data: *const TradeAccountDataSymbolLimitsUpdateVLSData,
    capacity: usize,
}

#[repr(packed(1), C)]
pub struct TradeAccountDataSymbolLimitsUpdateVLSData {
    size: u16,
    r#type: u16,
    base_size: u16,
    request_id: u32,
    is_deleted: bool,
    trade_account: VLS,
    symbol: VLS,
    trade_position_limit: f64,
    order_quantity_limit: f64,
    use_percent_of_margin: i32,
    override_margin_other_accounts: u8,
    use_percent_of_margin_for_day_trading: i32,
    number_of_days_before_last_trading_date_to_disallow_orders: i32,
    fixed_margin_cash_value: f32,
}

impl TradeAccountDataSymbolLimitsUpdateVLSData {
    pub fn new() -> Self {
        Self {
            size: 52u16.to_le(),
            r#type: TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE.to_le(),
            base_size: 52u16.to_le(),
            request_id: 0u32.to_le(),
            is_deleted: false,
            trade_account: crate::message::VLS::new(),
            symbol: crate::message::VLS::new(),
            trade_position_limit: 0.0,
            order_quantity_limit: 0.0,
            use_percent_of_margin: 100i32.to_le(),
            override_margin_other_accounts: 0,
            use_percent_of_margin_for_day_trading: 0i32.to_le(),
            number_of_days_before_last_trading_date_to_disallow_orders: 0i32.to_le(),
            fixed_margin_cash_value: 0.0,
        }
    }
}

unsafe impl Send for TradeAccountDataSymbolLimitsUpdateVLS {}
unsafe impl Send for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {}
unsafe impl Send for TradeAccountDataSymbolLimitsUpdateVLSData {}

impl Drop for TradeAccountDataSymbolLimitsUpdateVLS {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Drop for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    #[inline]
    fn drop(&mut self) {
        crate::deallocate(self.data as *mut u8, self.capacity() as usize);
    }
}

impl Clone for TradeAccountDataSymbolLimitsUpdateVLS {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Clone for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    #[inline]
    fn clone(&self) -> Self {
        let mut c = Self::new();
        self.copy_to(&mut c);
        c
    }
}

impl Into<Vec<u8>> for TradeAccountDataSymbolLimitsUpdateVLS {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl Into<Vec<u8>> for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    #[inline]
    fn into(self) -> Vec<u8> {
        self.into_vec()
    }
}

impl core::ops::Deref for TradeAccountDataSymbolLimitsUpdateVLS {
    type Target = TradeAccountDataSymbolLimitsUpdateVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataSymbolLimitsUpdateVLS {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl core::ops::Deref for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    type Target = TradeAccountDataSymbolLimitsUpdateVLSData;

    #[inline]
    fn deref(&self) -> &Self::Target {
        unsafe { &*self.data }
    }
}

impl core::ops::DerefMut for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    #[inline]
    fn deref_mut(&mut self) -> &mut Self::Target {
        unsafe { &mut *(self.data as *mut Self::Target) }
    }
}

impl crate::Message for TradeAccountDataSymbolLimitsUpdateVLS {
    type Data = TradeAccountDataSymbolLimitsUpdateVLSData;

    const BASE_SIZE: usize = 52;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataSymbolLimitsUpdateVLSData::new(),
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
            data: data as *const TradeAccountDataSymbolLimitsUpdateVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataSymbolLimitsUpdateVLS {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataSymbolLimitsUpdateVLSData;
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
impl crate::Message for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    type Data = TradeAccountDataSymbolLimitsUpdateVLSData;

    const BASE_SIZE: usize = 52;
    const BASE_SIZE_OFFSET: isize = 4;
    const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;

    #[inline]
    fn new() -> Self {
        Self {
            data: crate::allocate(
                Self::BASE_SIZE,
                TradeAccountDataSymbolLimitsUpdateVLSData::new(),
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
            data: data as *const TradeAccountDataSymbolLimitsUpdateVLSData,
            capacity,
        }
    }
}

impl crate::VLSMessage for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    #[inline]
    unsafe fn set_ptr(&mut self, value: *const u8) {
        self.data = value as *const TradeAccountDataSymbolLimitsUpdateVLSData;
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
impl TradeAccountDataSymbolLimitsUpdate for TradeAccountDataSymbolLimitsUpdateVLS {
    type Safe = TradeAccountDataSymbolLimitsUpdateVLS;
    type Unsafe = TradeAccountDataSymbolLimitsUpdateVLSUnsafe;

    fn request_id(&self) -> u32 {
        u32::from_le(self.request_id)
    }

    fn is_deleted(&self) -> bool {
        self.is_deleted
    }

    fn trade_account(&self) -> &str {
        get_vls(self, self.trade_account)
    }

    fn symbol(&self) -> &str {
        get_vls(self, self.symbol)
    }

    fn trade_position_limit(&self) -> f64 {
        f64_le(self.trade_position_limit)
    }

    fn order_quantity_limit(&self) -> f64 {
        f64_le(self.order_quantity_limit)
    }

    fn use_percent_of_margin(&self) -> i32 {
        i32::from_le(self.use_percent_of_margin)
    }

    fn override_margin_other_accounts(&self) -> u8 {
        self.override_margin_other_accounts
    }

    fn use_percent_of_margin_for_day_trading(&self) -> i32 {
        i32::from_le(self.use_percent_of_margin_for_day_trading)
    }

    fn number_of_days_before_last_trading_date_to_disallow_orders(&self) -> i32 {
        i32::from_le(self.number_of_days_before_last_trading_date_to_disallow_orders)
    }

    fn fixed_margin_cash_value(&self) -> f32 {
        f32_le(self.fixed_margin_cash_value)
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        self.request_id = value.to_le();
        self
    }

    fn set_is_deleted(&mut self, value: bool) -> &mut Self {
        self.is_deleted = value;
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        self.trade_account = set_vls(self, self.trade_account, value);
        self
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        self.symbol = set_vls(self, self.symbol, value);
        self
    }

    fn set_trade_position_limit(&mut self, value: f64) -> &mut Self {
        self.trade_position_limit = f64_le(value);
        self
    }

    fn set_order_quantity_limit(&mut self, value: f64) -> &mut Self {
        self.order_quantity_limit = f64_le(value);
        self
    }

    fn set_use_percent_of_margin(&mut self, value: i32) -> &mut Self {
        self.use_percent_of_margin = value.to_le();
        self
    }

    fn set_override_margin_other_accounts(&mut self, value: u8) -> &mut Self {
        self.override_margin_other_accounts = value;
        self
    }

    fn set_use_percent_of_margin_for_day_trading(&mut self, value: i32) -> &mut Self {
        self.use_percent_of_margin_for_day_trading = value.to_le();
        self
    }

    fn set_number_of_days_before_last_trading_date_to_disallow_orders(
        &mut self,
        value: i32,
    ) -> &mut Self {
        self.number_of_days_before_last_trading_date_to_disallow_orders = value.to_le();
        self
    }

    fn set_fixed_margin_cash_value(&mut self, value: f32) -> &mut Self {
        self.fixed_margin_cash_value = f32_le(value);
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

impl TradeAccountDataSymbolLimitsUpdate for TradeAccountDataSymbolLimitsUpdateVLSUnsafe {
    type Safe = TradeAccountDataSymbolLimitsUpdateVLS;
    type Unsafe = TradeAccountDataSymbolLimitsUpdateVLSUnsafe;

    fn request_id(&self) -> u32 {
        if self.is_out_of_bounds(10) {
            0u32.to_le()
        } else {
            u32::from_le(self.request_id)
        }
    }

    fn is_deleted(&self) -> bool {
        if self.is_out_of_bounds(11) {
            false
        } else {
            self.is_deleted
        }
    }

    fn trade_account(&self) -> &str {
        if self.is_out_of_bounds(15) {
            ""
        } else {
            get_vls(self, self.trade_account)
        }
    }

    fn symbol(&self) -> &str {
        if self.is_out_of_bounds(19) {
            ""
        } else {
            get_vls(self, self.symbol)
        }
    }

    fn trade_position_limit(&self) -> f64 {
        if self.is_out_of_bounds(27) {
            0.0
        } else {
            f64_le(self.trade_position_limit)
        }
    }

    fn order_quantity_limit(&self) -> f64 {
        if self.is_out_of_bounds(35) {
            0.0
        } else {
            f64_le(self.order_quantity_limit)
        }
    }

    fn use_percent_of_margin(&self) -> i32 {
        if self.is_out_of_bounds(39) {
            100i32.to_le()
        } else {
            i32::from_le(self.use_percent_of_margin)
        }
    }

    fn override_margin_other_accounts(&self) -> u8 {
        if self.is_out_of_bounds(40) {
            0
        } else {
            self.override_margin_other_accounts
        }
    }

    fn use_percent_of_margin_for_day_trading(&self) -> i32 {
        if self.is_out_of_bounds(44) {
            0i32.to_le()
        } else {
            i32::from_le(self.use_percent_of_margin_for_day_trading)
        }
    }

    fn number_of_days_before_last_trading_date_to_disallow_orders(&self) -> i32 {
        if self.is_out_of_bounds(48) {
            0i32.to_le()
        } else {
            i32::from_le(self.number_of_days_before_last_trading_date_to_disallow_orders)
        }
    }

    fn fixed_margin_cash_value(&self) -> f32 {
        if self.is_out_of_bounds(52) {
            0.0
        } else {
            f32_le(self.fixed_margin_cash_value)
        }
    }

    fn set_request_id(&mut self, value: u32) -> &mut Self {
        if !self.is_out_of_bounds(10) {
            self.request_id = value.to_le();
        }
        self
    }

    fn set_is_deleted(&mut self, value: bool) -> &mut Self {
        if !self.is_out_of_bounds(11) {
            self.is_deleted = value;
        }
        self
    }

    fn set_trade_account(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(15) {
            self.trade_account = set_vls(self, self.trade_account, value);
        }
        self
    }

    fn set_symbol(&mut self, value: &str) -> &mut Self {
        if !self.is_out_of_bounds(19) {
            self.symbol = set_vls(self, self.symbol, value);
        }
        self
    }

    fn set_trade_position_limit(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(27) {
            self.trade_position_limit = f64_le(value);
        }
        self
    }

    fn set_order_quantity_limit(&mut self, value: f64) -> &mut Self {
        if !self.is_out_of_bounds(35) {
            self.order_quantity_limit = f64_le(value);
        }
        self
    }

    fn set_use_percent_of_margin(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(39) {
            self.use_percent_of_margin = value.to_le();
        }
        self
    }

    fn set_override_margin_other_accounts(&mut self, value: u8) -> &mut Self {
        if !self.is_out_of_bounds(40) {
            self.override_margin_other_accounts = value;
        }
        self
    }

    fn set_use_percent_of_margin_for_day_trading(&mut self, value: i32) -> &mut Self {
        if !self.is_out_of_bounds(44) {
            self.use_percent_of_margin_for_day_trading = value.to_le();
        }
        self
    }

    fn set_number_of_days_before_last_trading_date_to_disallow_orders(
        &mut self,
        value: i32,
    ) -> &mut Self {
        if !self.is_out_of_bounds(48) {
            self.number_of_days_before_last_trading_date_to_disallow_orders = value.to_le();
        }
        self
    }

    fn set_fixed_margin_cash_value(&mut self, value: f32) -> &mut Self {
        if !self.is_out_of_bounds(52) {
            self.fixed_margin_cash_value = f32_le(value);
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
                52usize,
                core::mem::size_of::<TradeAccountDataSymbolLimitsUpdateVLSData>(),
                "TradeAccountDataSymbolLimitsUpdateVLSData sizeof expected {:} but was {:}",
                52usize,
                core::mem::size_of::<TradeAccountDataSymbolLimitsUpdateVLSData>()
            );
            assert_eq!(
                52u16,
                TradeAccountDataSymbolLimitsUpdateVLS::new().size(),
                "TradeAccountDataSymbolLimitsUpdateVLS sizeof expected {:} but was {:}",
                52u16,
                TradeAccountDataSymbolLimitsUpdateVLS::new().size(),
            );
            assert_eq!(
                TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE,
                TradeAccountDataSymbolLimitsUpdateVLS::new().r#type(),
                "TradeAccountDataSymbolLimitsUpdateVLS type expected {:} but was {:}",
                TRADE_ACCOUNT_DATA_SYMBOL_LIMITS_UPDATE,
                TradeAccountDataSymbolLimitsUpdateVLS::new().r#type(),
            );
            assert_eq!(
                10122u16,
                TradeAccountDataSymbolLimitsUpdateVLS::new().r#type(),
                "TradeAccountDataSymbolLimitsUpdateVLS type expected {:} but was {:}",
                10122u16,
                TradeAccountDataSymbolLimitsUpdateVLS::new().r#type(),
            );
            let d = TradeAccountDataSymbolLimitsUpdateVLSData::new();
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
                (core::ptr::addr_of!(d.request_id) as usize) - p,
                "request_id offset expected {:} but was {:}",
                6usize,
                (core::ptr::addr_of!(d.request_id) as usize) - p,
            );
            assert_eq!(
                10usize,
                (core::ptr::addr_of!(d.is_deleted) as usize) - p,
                "is_deleted offset expected {:} but was {:}",
                10usize,
                (core::ptr::addr_of!(d.is_deleted) as usize) - p,
            );
            assert_eq!(
                11usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
                "trade_account offset expected {:} but was {:}",
                11usize,
                (core::ptr::addr_of!(d.trade_account) as usize) - p,
            );
            assert_eq!(
                15usize,
                (core::ptr::addr_of!(d.symbol) as usize) - p,
                "symbol offset expected {:} but was {:}",
                15usize,
                (core::ptr::addr_of!(d.symbol) as usize) - p,
            );
            assert_eq!(
                19usize,
                (core::ptr::addr_of!(d.trade_position_limit) as usize) - p,
                "trade_position_limit offset expected {:} but was {:}",
                19usize,
                (core::ptr::addr_of!(d.trade_position_limit) as usize) - p,
            );
            assert_eq!(
                27usize,
                (core::ptr::addr_of!(d.order_quantity_limit) as usize) - p,
                "order_quantity_limit offset expected {:} but was {:}",
                27usize,
                (core::ptr::addr_of!(d.order_quantity_limit) as usize) - p,
            );
            assert_eq!(
                35usize,
                (core::ptr::addr_of!(d.use_percent_of_margin) as usize) - p,
                "use_percent_of_margin offset expected {:} but was {:}",
                35usize,
                (core::ptr::addr_of!(d.use_percent_of_margin) as usize) - p,
            );
            assert_eq!(
                39usize,
                (core::ptr::addr_of!(d.override_margin_other_accounts) as usize) - p,
                "override_margin_other_accounts offset expected {:} but was {:}",
                39usize,
                (core::ptr::addr_of!(d.override_margin_other_accounts) as usize) - p,
            );
            assert_eq!(
                40usize,
                (core::ptr::addr_of!(d.use_percent_of_margin_for_day_trading) as usize) - p,
                "use_percent_of_margin_for_day_trading offset expected {:} but was {:}",
                40usize,
                (core::ptr::addr_of!(d.use_percent_of_margin_for_day_trading) as usize) - p,
            );
            assert_eq!(
                44usize,
                (core::ptr::addr_of!(d.number_of_days_before_last_trading_date_to_disallow_orders) as usize) - p,
                "number_of_days_before_last_trading_date_to_disallow_orders offset expected {:} but was {:}",
                44usize,
                (core::ptr::addr_of!(d.number_of_days_before_last_trading_date_to_disallow_orders) as usize) - p,
            );
            assert_eq!(
                48usize,
                (core::ptr::addr_of!(d.fixed_margin_cash_value) as usize) - p,
                "fixed_margin_cash_value offset expected {:} but was {:}",
                48usize,
                (core::ptr::addr_of!(d.fixed_margin_cash_value) as usize) - p,
            );
        }
    }
}
