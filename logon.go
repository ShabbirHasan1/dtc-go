package dtc

import (
	"fmt"
	"unsafe"
)

type LogonRequest interface {
	Message
	Username() string
	Password() string
	GeneralTextData() string
	Integer1() int32
	Integer2() int32
	HeartbeatIntervalInSeconds() int32
	Unused1() int32
	TradeAccount() string
	HardwareIdentifier() string
	ClientName() string
	MarketDataTransmissionInterval() int32
}

type LogonRequestSetter interface {
	SetUsername(value string) LogonRequestSetter
}

type LogonRequestFixed struct {
	size                           uint16
	_type                          uint16
	protocolVersion                int32
	username                       [UsernamePasswordLength]byte
	password                       [UsernamePasswordLength]byte
	generalTextData                [GeneralIdentifierLength]byte
	integer1                       int32
	integer2                       int32
	heartbeatIntervalInSeconds     int32
	unused1                        int32
	tradeAccount                   [TradeAccountLength]byte
	hardwareIdentifier             [GeneralIdentifierLength]byte
	clientName                     [32]byte
	marketDataTransmissionInterval int32
}

func (lr *LogonRequestFixed) Clear() {
	*lr = LogonRequestFixed{}
	lr.protocolVersion = CurrentVersion
}

func (lr *LogonRequestFixed) Size() uint16 {
	return lr.size
}

func (lr *LogonRequestFixed) Type() uint16 {
	return lr._type
}

func (lr *LogonRequestFixed) ProtocolVersion() int32 {
	return lr.protocolVersion
}

func (lr *LogonRequestFixed) Username() string {
	if lr.size < uint16(unsafe.Offsetof(lr.username))+uint16(UsernamePasswordLength) {
		return ""
	}
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&lr.username[0])),
		Len:  strlen(lr.username[0:cap(lr.username)]),
	}))
}

func (lr *LogonRequestFixed) Password() string {
	if lr.size < uint16(unsafe.Offsetof(lr.password))+uint16(UsernamePasswordLength) {
		return ""
	}
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&lr.password[0])),
		Len:  strlen(lr.password[0:cap(lr.password)]),
	}))
}

func (lr *LogonRequestFixed) GeneralTextData() string {
	if lr.size < uint16(unsafe.Offsetof(lr.generalTextData))+uint16(GeneralIdentifierLength) {
		return ""
	}
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&lr.generalTextData[0])),
		Len:  strlen(lr.generalTextData[0:cap(lr.generalTextData)]),
	}))
}

func (lr *LogonRequestFixed) Integer1() int32 {
	if lr.size < uint16(unsafe.Offsetof(lr.integer1)+unsafe.Sizeof(lr.integer1)) {
		return 0
	}
	return lr.integer1
}

func (lr *LogonRequestFixed) Integer2() int32 {
	if lr.size < uint16(unsafe.Offsetof(lr.integer2)+unsafe.Sizeof(lr.integer2)) {
		return 0
	}
	return lr.integer2
}

func (lr *LogonRequestFixed) HeartbeatIntervalInSeconds() int32 {
	if lr.size < uint16(unsafe.Offsetof(lr.heartbeatIntervalInSeconds)+unsafe.Sizeof(lr.heartbeatIntervalInSeconds)) {
		return 0
	}
	return lr.heartbeatIntervalInSeconds
}

func (lr *LogonRequestFixed) Unused1() int32 {
	if lr.size < uint16(unsafe.Offsetof(lr.unused1)+unsafe.Sizeof(lr.unused1)) {
		return 0
	}
	return lr.unused1
}

func (lr *LogonRequestFixed) TradeAccount() string {
	if lr.size < uint16(unsafe.Offsetof(lr.tradeAccount))+uint16(TradeAccountLength) {
		return ""
	}
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&lr.tradeAccount[0])),
		Len:  strlen(lr.tradeAccount[0:cap(lr.tradeAccount)]),
	}))
}

func (lr *LogonRequestFixed) HardwareIdentifier() string {
	if lr.size < uint16(unsafe.Offsetof(lr.hardwareIdentifier))+uint16(GeneralIdentifierLength) {
		return ""
	}
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&lr.hardwareIdentifier[0])),
		Len:  strlen(lr.hardwareIdentifier[0:cap(lr.hardwareIdentifier)]),
	}))
}

func (lr *LogonRequestFixed) ClientName() string {
	if lr.size < uint16(unsafe.Offsetof(lr.clientName))+32 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&_string{
		Data: uintptr(unsafe.Pointer(&lr.clientName[0])),
		Len:  strlen(lr.clientName[0:cap(lr.clientName)]),
	}))
}

func (lr *LogonRequestFixed) MarketDataTransmissionInterval() int32 {
	if lr.size < uint16(unsafe.Offsetof(lr.marketDataTransmissionInterval)+unsafe.Sizeof(lr.marketDataTransmissionInterval)) {
		return 0
	}
	return lr.marketDataTransmissionInterval
}

func (lr *LogonRequestFixed) String() string {
	return fmt.Sprintf("LogonRequestFixed{"+
		"size = %d,\n"+
		"type = %d,\n"+
		"}", lr.size, lr._type)
}

type LogonRequestVLS struct {
	size                           uint16
	_type                          uint16
	baseSize                       uint16
	protocolVersion                int32
	username                       VLS
	password                       VLS
	generalTextData                VLS
	integer1                       int32
	integer2                       int32
	heartbeatIntervalInSeconds     int32
	unused1                        int32
	tradeAccount                   VLS
	hardwareIdentifier             VLS
	clientName                     VLS
	marketDataTransmissionInterval int32
}

func (lr *LogonRequestVLS) Size() uint16 {
	return lr.size
}

func (lr *LogonRequestVLS) Type() uint16 {
	return lr._type
}

func (lr *LogonRequestVLS) BaseSize() uint16 {
	return lr.baseSize
}

func (lr *LogonRequestVLS) ProtocolVersion() int32 {
	return lr.protocolVersion
}

func (lr *LogonRequestVLS) Integer1() int32 {
	return lr.integer1
}

func (lr *LogonRequestVLS) Integer2() int32 {
	return lr.integer2
}

func (lr *LogonRequestVLS) HeartbeatIntervalInSeconds() int32 {
	return lr.heartbeatIntervalInSeconds
}

func (lr *LogonRequestVLS) Unused1() int32 {
	return lr.unused1
}

func (lr *LogonRequestVLS) MarketDataTransmissionInterval() int32 {
	return lr.marketDataTransmissionInterval
}

func (lr *LogonRequestVLS) Username() string {
	return getVLSField(&lr.size, lr.baseSize, &lr.username, unsafe.Offsetof(lr.username))
}

func (lr *LogonRequestVLS) Password() string {
	return getVLSField(&lr.size, lr.baseSize, &lr.password, unsafe.Offsetof(lr.password))
}

func (lr *LogonRequestVLS) GeneralTextData() string {
	return getVLSField(&lr.size, lr.baseSize, &lr.generalTextData, unsafe.Offsetof(lr.generalTextData))
}

func (lr *LogonRequestVLS) TradeAccount() string {
	return getVLSField(&lr.size, lr.baseSize, &lr.tradeAccount, unsafe.Offsetof(lr.tradeAccount))
}

func (lr *LogonRequestVLS) HardwareIdentifier() string {
	return getVLSField(&lr.size, lr.baseSize, &lr.hardwareIdentifier, unsafe.Offsetof(lr.hardwareIdentifier))
}

func (lr *LogonRequestVLS) ClientName() string {
	return getVLSField(&lr.size, lr.baseSize, &lr.clientName, unsafe.Offsetof(lr.clientName))
}

func (lr *LogonRequestVLS) SetSize(size uint16) {
	lr.size = size
}

func (lr *LogonRequestVLS) SetType(_type uint16) {
	lr._type = _type
}

func (lr *LogonRequestVLS) SetBaseSize(baseSize uint16) {
	lr.baseSize = baseSize
}

func (lr *LogonRequestVLS) SetProtocolVersion(protocolVersion int32) {
	lr.protocolVersion = protocolVersion
}

func (lr *LogonRequestVLS) SetInteger1(integer1 int32) {
	lr.integer1 = integer1
}

func (lr *LogonRequestVLS) SetInteger2(integer2 int32) {
	lr.integer2 = integer2
}

func (lr *LogonRequestVLS) SetHeartbeatIntervalInSeconds(heartbeatIntervalInSeconds int32) {
	lr.heartbeatIntervalInSeconds = heartbeatIntervalInSeconds
}

func (lr *LogonRequestVLS) SetUnused1(unused1 int32) {
	lr.unused1 = unused1
}

func (lr *LogonRequestVLS) SetMarketDataTransmissionInterval(marketDataTransmissionInterval int32) {
	lr.marketDataTransmissionInterval = marketDataTransmissionInterval
}
