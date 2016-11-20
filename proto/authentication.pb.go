// Code generated by protoc-gen-go.
// source: authentication.proto
// DO NOT EDIT!

package Spotify

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AuthenticationType int32

const (
	AuthenticationType_AUTHENTICATION_USER_PASS                   AuthenticationType = 0
	AuthenticationType_AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS  AuthenticationType = 1
	AuthenticationType_AUTHENTICATION_STORED_FACEBOOK_CREDENTIALS AuthenticationType = 2
	AuthenticationType_AUTHENTICATION_SPOTIFY_TOKEN               AuthenticationType = 3
	AuthenticationType_AUTHENTICATION_FACEBOOK_TOKEN              AuthenticationType = 4
)

var AuthenticationType_name = map[int32]string{
	0: "AUTHENTICATION_USER_PASS",
	1: "AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS",
	2: "AUTHENTICATION_STORED_FACEBOOK_CREDENTIALS",
	3: "AUTHENTICATION_SPOTIFY_TOKEN",
	4: "AUTHENTICATION_FACEBOOK_TOKEN",
}
var AuthenticationType_value = map[string]int32{
	"AUTHENTICATION_USER_PASS":                   0,
	"AUTHENTICATION_STORED_SPOTIFY_CREDENTIALS":  1,
	"AUTHENTICATION_STORED_FACEBOOK_CREDENTIALS": 2,
	"AUTHENTICATION_SPOTIFY_TOKEN":               3,
	"AUTHENTICATION_FACEBOOK_TOKEN":              4,
}

func (x AuthenticationType) Enum() *AuthenticationType {
	p := new(AuthenticationType)
	*p = x
	return p
}
func (x AuthenticationType) String() string {
	return proto.EnumName(AuthenticationType_name, int32(x))
}
func (x *AuthenticationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AuthenticationType_value, data, "AuthenticationType")
	if err != nil {
		return err
	}
	*x = AuthenticationType(value)
	return nil
}
func (AuthenticationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type AccountCreation int32

const (
	AccountCreation_ACCOUNT_CREATION_ALWAYS_PROMPT AccountCreation = 1
	AccountCreation_ACCOUNT_CREATION_ALWAYS_CREATE AccountCreation = 3
)

var AccountCreation_name = map[int32]string{
	1: "ACCOUNT_CREATION_ALWAYS_PROMPT",
	3: "ACCOUNT_CREATION_ALWAYS_CREATE",
}
var AccountCreation_value = map[string]int32{
	"ACCOUNT_CREATION_ALWAYS_PROMPT": 1,
	"ACCOUNT_CREATION_ALWAYS_CREATE": 3,
}

func (x AccountCreation) Enum() *AccountCreation {
	p := new(AccountCreation)
	*p = x
	return p
}
func (x AccountCreation) String() string {
	return proto.EnumName(AccountCreation_name, int32(x))
}
func (x *AccountCreation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AccountCreation_value, data, "AccountCreation")
	if err != nil {
		return err
	}
	*x = AccountCreation(value)
	return nil
}
func (AccountCreation) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type CpuFamily int32

const (
	CpuFamily_CPU_UNKNOWN  CpuFamily = 0
	CpuFamily_CPU_X86      CpuFamily = 1
	CpuFamily_CPU_X86_64   CpuFamily = 2
	CpuFamily_CPU_PPC      CpuFamily = 3
	CpuFamily_CPU_PPC_64   CpuFamily = 4
	CpuFamily_CPU_ARM      CpuFamily = 5
	CpuFamily_CPU_IA64     CpuFamily = 6
	CpuFamily_CPU_SH       CpuFamily = 7
	CpuFamily_CPU_MIPS     CpuFamily = 8
	CpuFamily_CPU_BLACKFIN CpuFamily = 9
)

var CpuFamily_name = map[int32]string{
	0: "CPU_UNKNOWN",
	1: "CPU_X86",
	2: "CPU_X86_64",
	3: "CPU_PPC",
	4: "CPU_PPC_64",
	5: "CPU_ARM",
	6: "CPU_IA64",
	7: "CPU_SH",
	8: "CPU_MIPS",
	9: "CPU_BLACKFIN",
}
var CpuFamily_value = map[string]int32{
	"CPU_UNKNOWN":  0,
	"CPU_X86":      1,
	"CPU_X86_64":   2,
	"CPU_PPC":      3,
	"CPU_PPC_64":   4,
	"CPU_ARM":      5,
	"CPU_IA64":     6,
	"CPU_SH":       7,
	"CPU_MIPS":     8,
	"CPU_BLACKFIN": 9,
}

func (x CpuFamily) Enum() *CpuFamily {
	p := new(CpuFamily)
	*p = x
	return p
}
func (x CpuFamily) String() string {
	return proto.EnumName(CpuFamily_name, int32(x))
}
func (x *CpuFamily) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CpuFamily_value, data, "CpuFamily")
	if err != nil {
		return err
	}
	*x = CpuFamily(value)
	return nil
}
func (CpuFamily) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type Brand int32

const (
	Brand_BRAND_UNBRANDED Brand = 0
	Brand_BRAND_INQ       Brand = 1
	Brand_BRAND_HTC       Brand = 2
	Brand_BRAND_NOKIA     Brand = 3
)

var Brand_name = map[int32]string{
	0: "BRAND_UNBRANDED",
	1: "BRAND_INQ",
	2: "BRAND_HTC",
	3: "BRAND_NOKIA",
}
var Brand_value = map[string]int32{
	"BRAND_UNBRANDED": 0,
	"BRAND_INQ":       1,
	"BRAND_HTC":       2,
	"BRAND_NOKIA":     3,
}

func (x Brand) Enum() *Brand {
	p := new(Brand)
	*p = x
	return p
}
func (x Brand) String() string {
	return proto.EnumName(Brand_name, int32(x))
}
func (x *Brand) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Brand_value, data, "Brand")
	if err != nil {
		return err
	}
	*x = Brand(value)
	return nil
}
func (Brand) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type Os int32

const (
	Os_OS_UNKNOWN    Os = 0
	Os_OS_WINDOWS    Os = 1
	Os_OS_OSX        Os = 2
	Os_OS_IPHONE     Os = 3
	Os_OS_S60        Os = 4
	Os_OS_LINUX      Os = 5
	Os_OS_WINDOWS_CE Os = 6
	Os_OS_ANDROID    Os = 7
	Os_OS_PALM       Os = 8
	Os_OS_FREEBSD    Os = 9
	Os_OS_BLACKBERRY Os = 10
	Os_OS_SONOS      Os = 11
	Os_OS_LOGITECH   Os = 12
	Os_OS_WP7        Os = 13
	Os_OS_ONKYO      Os = 14
	Os_OS_PHILIPS    Os = 15
	Os_OS_WD         Os = 16
	Os_OS_VOLVO      Os = 17
	Os_OS_TIVO       Os = 18
	Os_OS_AWOX       Os = 19
	Os_OS_MEEGO      Os = 20
	Os_OS_QNXNTO     Os = 21
	Os_OS_BCO        Os = 22
)

var Os_name = map[int32]string{
	0:  "OS_UNKNOWN",
	1:  "OS_WINDOWS",
	2:  "OS_OSX",
	3:  "OS_IPHONE",
	4:  "OS_S60",
	5:  "OS_LINUX",
	6:  "OS_WINDOWS_CE",
	7:  "OS_ANDROID",
	8:  "OS_PALM",
	9:  "OS_FREEBSD",
	10: "OS_BLACKBERRY",
	11: "OS_SONOS",
	12: "OS_LOGITECH",
	13: "OS_WP7",
	14: "OS_ONKYO",
	15: "OS_PHILIPS",
	16: "OS_WD",
	17: "OS_VOLVO",
	18: "OS_TIVO",
	19: "OS_AWOX",
	20: "OS_MEEGO",
	21: "OS_QNXNTO",
	22: "OS_BCO",
}
var Os_value = map[string]int32{
	"OS_UNKNOWN":    0,
	"OS_WINDOWS":    1,
	"OS_OSX":        2,
	"OS_IPHONE":     3,
	"OS_S60":        4,
	"OS_LINUX":      5,
	"OS_WINDOWS_CE": 6,
	"OS_ANDROID":    7,
	"OS_PALM":       8,
	"OS_FREEBSD":    9,
	"OS_BLACKBERRY": 10,
	"OS_SONOS":      11,
	"OS_LOGITECH":   12,
	"OS_WP7":        13,
	"OS_ONKYO":      14,
	"OS_PHILIPS":    15,
	"OS_WD":         16,
	"OS_VOLVO":      17,
	"OS_TIVO":       18,
	"OS_AWOX":       19,
	"OS_MEEGO":      20,
	"OS_QNXNTO":     21,
	"OS_BCO":        22,
}

func (x Os) Enum() *Os {
	p := new(Os)
	*p = x
	return p
}
func (x Os) String() string {
	return proto.EnumName(Os_name, int32(x))
}
func (x *Os) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Os_value, data, "Os")
	if err != nil {
		return err
	}
	*x = Os(value)
	return nil
}
func (Os) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type AccountType int32

const (
	AccountType_Spotify  AccountType = 0
	AccountType_Facebook AccountType = 1
)

var AccountType_name = map[int32]string{
	0: "Spotify",
	1: "Facebook",
}
var AccountType_value = map[string]int32{
	"Spotify":  0,
	"Facebook": 1,
}

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}
func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}
func (x *AccountType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AccountType_value, data, "AccountType")
	if err != nil {
		return err
	}
	*x = AccountType(value)
	return nil
}
func (AccountType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type ClientResponseEncrypted struct {
	LoginCredentials    *LoginCredentials         `protobuf:"bytes,10,req,name=login_credentials,json=loginCredentials" json:"login_credentials,omitempty"`
	AccountCreation     *AccountCreation          `protobuf:"varint,20,opt,name=account_creation,json=accountCreation,enum=Spotify.AccountCreation" json:"account_creation,omitempty"`
	FingerprintResponse *FingerprintResponseUnion `protobuf:"bytes,30,opt,name=fingerprint_response,json=fingerprintResponse" json:"fingerprint_response,omitempty"`
	PeerTicket          *PeerTicketUnion          `protobuf:"bytes,40,opt,name=peer_ticket,json=peerTicket" json:"peer_ticket,omitempty"`
	SystemInfo          *SystemInfo               `protobuf:"bytes,50,req,name=system_info,json=systemInfo" json:"system_info,omitempty"`
	PlatformModel       *string                   `protobuf:"bytes,60,opt,name=platform_model,json=platformModel" json:"platform_model,omitempty"`
	VersionString       *string                   `protobuf:"bytes,70,opt,name=version_string,json=versionString" json:"version_string,omitempty"`
	Appkey              *LibspotifyAppKey         `protobuf:"bytes,80,opt,name=appkey" json:"appkey,omitempty"`
	ClientInfo          *ClientInfo               `protobuf:"bytes,90,opt,name=client_info,json=clientInfo" json:"client_info,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *ClientResponseEncrypted) Reset()                    { *m = ClientResponseEncrypted{} }
func (m *ClientResponseEncrypted) String() string            { return proto.CompactTextString(m) }
func (*ClientResponseEncrypted) ProtoMessage()               {}
func (*ClientResponseEncrypted) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ClientResponseEncrypted) GetLoginCredentials() *LoginCredentials {
	if m != nil {
		return m.LoginCredentials
	}
	return nil
}

func (m *ClientResponseEncrypted) GetAccountCreation() AccountCreation {
	if m != nil && m.AccountCreation != nil {
		return *m.AccountCreation
	}
	return AccountCreation_ACCOUNT_CREATION_ALWAYS_PROMPT
}

func (m *ClientResponseEncrypted) GetFingerprintResponse() *FingerprintResponseUnion {
	if m != nil {
		return m.FingerprintResponse
	}
	return nil
}

func (m *ClientResponseEncrypted) GetPeerTicket() *PeerTicketUnion {
	if m != nil {
		return m.PeerTicket
	}
	return nil
}

func (m *ClientResponseEncrypted) GetSystemInfo() *SystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return nil
}

func (m *ClientResponseEncrypted) GetPlatformModel() string {
	if m != nil && m.PlatformModel != nil {
		return *m.PlatformModel
	}
	return ""
}

func (m *ClientResponseEncrypted) GetVersionString() string {
	if m != nil && m.VersionString != nil {
		return *m.VersionString
	}
	return ""
}

func (m *ClientResponseEncrypted) GetAppkey() *LibspotifyAppKey {
	if m != nil {
		return m.Appkey
	}
	return nil
}

func (m *ClientResponseEncrypted) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

type LoginCredentials struct {
	Username         *string             `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Typ              *AuthenticationType `protobuf:"varint,20,req,name=typ,enum=Spotify.AuthenticationType" json:"typ,omitempty"`
	AuthData         []byte              `protobuf:"bytes,30,opt,name=auth_data,json=authData" json:"auth_data,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *LoginCredentials) Reset()                    { *m = LoginCredentials{} }
func (m *LoginCredentials) String() string            { return proto.CompactTextString(m) }
func (*LoginCredentials) ProtoMessage()               {}
func (*LoginCredentials) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *LoginCredentials) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *LoginCredentials) GetTyp() AuthenticationType {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return AuthenticationType_AUTHENTICATION_USER_PASS
}

func (m *LoginCredentials) GetAuthData() []byte {
	if m != nil {
		return m.AuthData
	}
	return nil
}

type FingerprintResponseUnion struct {
	Grain            *FingerprintGrainResponse      `protobuf:"bytes,10,opt,name=grain" json:"grain,omitempty"`
	HmacRipemd       *FingerprintHmacRipemdResponse `protobuf:"bytes,20,opt,name=hmac_ripemd,json=hmacRipemd" json:"hmac_ripemd,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *FingerprintResponseUnion) Reset()                    { *m = FingerprintResponseUnion{} }
func (m *FingerprintResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*FingerprintResponseUnion) ProtoMessage()               {}
func (*FingerprintResponseUnion) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *FingerprintResponseUnion) GetGrain() *FingerprintGrainResponse {
	if m != nil {
		return m.Grain
	}
	return nil
}

func (m *FingerprintResponseUnion) GetHmacRipemd() *FingerprintHmacRipemdResponse {
	if m != nil {
		return m.HmacRipemd
	}
	return nil
}

type FingerprintGrainResponse struct {
	EncryptedKey     []byte `protobuf:"bytes,10,req,name=encrypted_key,json=encryptedKey" json:"encrypted_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintGrainResponse) Reset()                    { *m = FingerprintGrainResponse{} }
func (m *FingerprintGrainResponse) String() string            { return proto.CompactTextString(m) }
func (*FingerprintGrainResponse) ProtoMessage()               {}
func (*FingerprintGrainResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *FingerprintGrainResponse) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

type FingerprintHmacRipemdResponse struct {
	Hmac             []byte `protobuf:"bytes,10,req,name=hmac" json:"hmac,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FingerprintHmacRipemdResponse) Reset()                    { *m = FingerprintHmacRipemdResponse{} }
func (m *FingerprintHmacRipemdResponse) String() string            { return proto.CompactTextString(m) }
func (*FingerprintHmacRipemdResponse) ProtoMessage()               {}
func (*FingerprintHmacRipemdResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *FingerprintHmacRipemdResponse) GetHmac() []byte {
	if m != nil {
		return m.Hmac
	}
	return nil
}

type PeerTicketUnion struct {
	PublicKey        *PeerTicketPublicKey `protobuf:"bytes,10,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	OldTicket        *PeerTicketOld       `protobuf:"bytes,20,opt,name=old_ticket,json=oldTicket" json:"old_ticket,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *PeerTicketUnion) Reset()                    { *m = PeerTicketUnion{} }
func (m *PeerTicketUnion) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketUnion) ProtoMessage()               {}
func (*PeerTicketUnion) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *PeerTicketUnion) GetPublicKey() *PeerTicketPublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PeerTicketUnion) GetOldTicket() *PeerTicketOld {
	if m != nil {
		return m.OldTicket
	}
	return nil
}

type PeerTicketPublicKey struct {
	PublicKey        []byte `protobuf:"bytes,10,req,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PeerTicketPublicKey) Reset()                    { *m = PeerTicketPublicKey{} }
func (m *PeerTicketPublicKey) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketPublicKey) ProtoMessage()               {}
func (*PeerTicketPublicKey) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *PeerTicketPublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type PeerTicketOld struct {
	PeerTicket          []byte `protobuf:"bytes,10,req,name=peer_ticket,json=peerTicket" json:"peer_ticket,omitempty"`
	PeerTicketSignature []byte `protobuf:"bytes,20,req,name=peer_ticket_signature,json=peerTicketSignature" json:"peer_ticket_signature,omitempty"`
	XXX_unrecognized    []byte `json:"-"`
}

func (m *PeerTicketOld) Reset()                    { *m = PeerTicketOld{} }
func (m *PeerTicketOld) String() string            { return proto.CompactTextString(m) }
func (*PeerTicketOld) ProtoMessage()               {}
func (*PeerTicketOld) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *PeerTicketOld) GetPeerTicket() []byte {
	if m != nil {
		return m.PeerTicket
	}
	return nil
}

func (m *PeerTicketOld) GetPeerTicketSignature() []byte {
	if m != nil {
		return m.PeerTicketSignature
	}
	return nil
}

type SystemInfo struct {
	CpuFamily               *CpuFamily `protobuf:"varint,10,req,name=cpu_family,json=cpuFamily,enum=Spotify.CpuFamily" json:"cpu_family,omitempty"`
	CpuSubtype              *uint32    `protobuf:"varint,20,opt,name=cpu_subtype,json=cpuSubtype" json:"cpu_subtype,omitempty"`
	CpuExt                  *uint32    `protobuf:"varint,30,opt,name=cpu_ext,json=cpuExt" json:"cpu_ext,omitempty"`
	Brand                   *Brand     `protobuf:"varint,40,opt,name=brand,enum=Spotify.Brand" json:"brand,omitempty"`
	BrandFlags              *uint32    `protobuf:"varint,50,opt,name=brand_flags,json=brandFlags" json:"brand_flags,omitempty"`
	Os                      *Os        `protobuf:"varint,60,req,name=os,enum=Spotify.Os" json:"os,omitempty"`
	OsVersion               *uint32    `protobuf:"varint,70,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	OsExt                   *uint32    `protobuf:"varint,80,opt,name=os_ext,json=osExt" json:"os_ext,omitempty"`
	SystemInformationString *string    `protobuf:"bytes,90,opt,name=system_information_string,json=systemInformationString" json:"system_information_string,omitempty"`
	DeviceId                *string    `protobuf:"bytes,100,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	XXX_unrecognized        []byte     `json:"-"`
}

func (m *SystemInfo) Reset()                    { *m = SystemInfo{} }
func (m *SystemInfo) String() string            { return proto.CompactTextString(m) }
func (*SystemInfo) ProtoMessage()               {}
func (*SystemInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *SystemInfo) GetCpuFamily() CpuFamily {
	if m != nil && m.CpuFamily != nil {
		return *m.CpuFamily
	}
	return CpuFamily_CPU_UNKNOWN
}

func (m *SystemInfo) GetCpuSubtype() uint32 {
	if m != nil && m.CpuSubtype != nil {
		return *m.CpuSubtype
	}
	return 0
}

func (m *SystemInfo) GetCpuExt() uint32 {
	if m != nil && m.CpuExt != nil {
		return *m.CpuExt
	}
	return 0
}

func (m *SystemInfo) GetBrand() Brand {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return Brand_BRAND_UNBRANDED
}

func (m *SystemInfo) GetBrandFlags() uint32 {
	if m != nil && m.BrandFlags != nil {
		return *m.BrandFlags
	}
	return 0
}

func (m *SystemInfo) GetOs() Os {
	if m != nil && m.Os != nil {
		return *m.Os
	}
	return Os_OS_UNKNOWN
}

func (m *SystemInfo) GetOsVersion() uint32 {
	if m != nil && m.OsVersion != nil {
		return *m.OsVersion
	}
	return 0
}

func (m *SystemInfo) GetOsExt() uint32 {
	if m != nil && m.OsExt != nil {
		return *m.OsExt
	}
	return 0
}

func (m *SystemInfo) GetSystemInformationString() string {
	if m != nil && m.SystemInformationString != nil {
		return *m.SystemInformationString
	}
	return ""
}

func (m *SystemInfo) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

type LibspotifyAppKey struct {
	Version          *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	Devkey           []byte  `protobuf:"bytes,2,req,name=devkey" json:"devkey,omitempty"`
	Signature        []byte  `protobuf:"bytes,3,req,name=signature" json:"signature,omitempty"`
	Useragent        *string `protobuf:"bytes,4,req,name=useragent" json:"useragent,omitempty"`
	CallbackHash     []byte  `protobuf:"bytes,5,req,name=callback_hash,json=callbackHash" json:"callback_hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LibspotifyAppKey) Reset()                    { *m = LibspotifyAppKey{} }
func (m *LibspotifyAppKey) String() string            { return proto.CompactTextString(m) }
func (*LibspotifyAppKey) ProtoMessage()               {}
func (*LibspotifyAppKey) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *LibspotifyAppKey) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *LibspotifyAppKey) GetDevkey() []byte {
	if m != nil {
		return m.Devkey
	}
	return nil
}

func (m *LibspotifyAppKey) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *LibspotifyAppKey) GetUseragent() string {
	if m != nil && m.Useragent != nil {
		return *m.Useragent
	}
	return ""
}

func (m *LibspotifyAppKey) GetCallbackHash() []byte {
	if m != nil {
		return m.CallbackHash
	}
	return nil
}

type ClientInfo struct {
	Limited          *bool               `protobuf:"varint,1,opt,name=limited" json:"limited,omitempty"`
	Fb               *ClientInfoFacebook `protobuf:"bytes,2,opt,name=fb" json:"fb,omitempty"`
	Language         *string             `protobuf:"bytes,3,opt,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *ClientInfo) GetLimited() bool {
	if m != nil && m.Limited != nil {
		return *m.Limited
	}
	return false
}

func (m *ClientInfo) GetFb() *ClientInfoFacebook {
	if m != nil {
		return m.Fb
	}
	return nil
}

func (m *ClientInfo) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

type ClientInfoFacebook struct {
	MachineId        *string `protobuf:"bytes,1,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientInfoFacebook) Reset()                    { *m = ClientInfoFacebook{} }
func (m *ClientInfoFacebook) String() string            { return proto.CompactTextString(m) }
func (*ClientInfoFacebook) ProtoMessage()               {}
func (*ClientInfoFacebook) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *ClientInfoFacebook) GetMachineId() string {
	if m != nil && m.MachineId != nil {
		return *m.MachineId
	}
	return ""
}

type APWelcome struct {
	CanonicalUsername           *string              `protobuf:"bytes,10,req,name=canonical_username,json=canonicalUsername" json:"canonical_username,omitempty"`
	AccountTypeLoggedIn         *AccountType         `protobuf:"varint,20,req,name=account_type_logged_in,json=accountTypeLoggedIn,enum=Spotify.AccountType" json:"account_type_logged_in,omitempty"`
	CredentialsTypeLoggedIn     *AccountType         `protobuf:"varint,25,req,name=credentials_type_logged_in,json=credentialsTypeLoggedIn,enum=Spotify.AccountType" json:"credentials_type_logged_in,omitempty"`
	ReusableAuthCredentialsType *AuthenticationType  `protobuf:"varint,30,req,name=reusable_auth_credentials_type,json=reusableAuthCredentialsType,enum=Spotify.AuthenticationType" json:"reusable_auth_credentials_type,omitempty"`
	ReusableAuthCredentials     []byte               `protobuf:"bytes,40,req,name=reusable_auth_credentials,json=reusableAuthCredentials" json:"reusable_auth_credentials,omitempty"`
	LfsSecret                   []byte               `protobuf:"bytes,50,opt,name=lfs_secret,json=lfsSecret" json:"lfs_secret,omitempty"`
	AccountInfo                 *AccountInfo         `protobuf:"bytes,60,opt,name=account_info,json=accountInfo" json:"account_info,omitempty"`
	Fb                          *AccountInfoFacebook `protobuf:"bytes,70,opt,name=fb" json:"fb,omitempty"`
	XXX_unrecognized            []byte               `json:"-"`
}

func (m *APWelcome) Reset()                    { *m = APWelcome{} }
func (m *APWelcome) String() string            { return proto.CompactTextString(m) }
func (*APWelcome) ProtoMessage()               {}
func (*APWelcome) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *APWelcome) GetCanonicalUsername() string {
	if m != nil && m.CanonicalUsername != nil {
		return *m.CanonicalUsername
	}
	return ""
}

func (m *APWelcome) GetAccountTypeLoggedIn() AccountType {
	if m != nil && m.AccountTypeLoggedIn != nil {
		return *m.AccountTypeLoggedIn
	}
	return AccountType_Spotify
}

func (m *APWelcome) GetCredentialsTypeLoggedIn() AccountType {
	if m != nil && m.CredentialsTypeLoggedIn != nil {
		return *m.CredentialsTypeLoggedIn
	}
	return AccountType_Spotify
}

func (m *APWelcome) GetReusableAuthCredentialsType() AuthenticationType {
	if m != nil && m.ReusableAuthCredentialsType != nil {
		return *m.ReusableAuthCredentialsType
	}
	return AuthenticationType_AUTHENTICATION_USER_PASS
}

func (m *APWelcome) GetReusableAuthCredentials() []byte {
	if m != nil {
		return m.ReusableAuthCredentials
	}
	return nil
}

func (m *APWelcome) GetLfsSecret() []byte {
	if m != nil {
		return m.LfsSecret
	}
	return nil
}

func (m *APWelcome) GetAccountInfo() *AccountInfo {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

func (m *APWelcome) GetFb() *AccountInfoFacebook {
	if m != nil {
		return m.Fb
	}
	return nil
}

type AccountInfo struct {
	Spotify          *AccountInfoSpotify  `protobuf:"bytes,1,opt,name=spotify" json:"spotify,omitempty"`
	Facebook         *AccountInfoFacebook `protobuf:"bytes,2,opt,name=facebook" json:"facebook,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *AccountInfo) Reset()                    { *m = AccountInfo{} }
func (m *AccountInfo) String() string            { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()               {}
func (*AccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *AccountInfo) GetSpotify() *AccountInfoSpotify {
	if m != nil {
		return m.Spotify
	}
	return nil
}

func (m *AccountInfo) GetFacebook() *AccountInfoFacebook {
	if m != nil {
		return m.Facebook
	}
	return nil
}

type AccountInfoSpotify struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AccountInfoSpotify) Reset()                    { *m = AccountInfoSpotify{} }
func (m *AccountInfoSpotify) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoSpotify) ProtoMessage()               {}
func (*AccountInfoSpotify) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

type AccountInfoFacebook struct {
	AccessToken      *string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	MachineId        *string `protobuf:"bytes,2,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AccountInfoFacebook) Reset()                    { *m = AccountInfoFacebook{} }
func (m *AccountInfoFacebook) String() string            { return proto.CompactTextString(m) }
func (*AccountInfoFacebook) ProtoMessage()               {}
func (*AccountInfoFacebook) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *AccountInfoFacebook) GetAccessToken() string {
	if m != nil && m.AccessToken != nil {
		return *m.AccessToken
	}
	return ""
}

func (m *AccountInfoFacebook) GetMachineId() string {
	if m != nil && m.MachineId != nil {
		return *m.MachineId
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientResponseEncrypted)(nil), "Spotify.ClientResponseEncrypted")
	proto.RegisterType((*LoginCredentials)(nil), "Spotify.LoginCredentials")
	proto.RegisterType((*FingerprintResponseUnion)(nil), "Spotify.FingerprintResponseUnion")
	proto.RegisterType((*FingerprintGrainResponse)(nil), "Spotify.FingerprintGrainResponse")
	proto.RegisterType((*FingerprintHmacRipemdResponse)(nil), "Spotify.FingerprintHmacRipemdResponse")
	proto.RegisterType((*PeerTicketUnion)(nil), "Spotify.PeerTicketUnion")
	proto.RegisterType((*PeerTicketPublicKey)(nil), "Spotify.PeerTicketPublicKey")
	proto.RegisterType((*PeerTicketOld)(nil), "Spotify.PeerTicketOld")
	proto.RegisterType((*SystemInfo)(nil), "Spotify.SystemInfo")
	proto.RegisterType((*LibspotifyAppKey)(nil), "Spotify.LibspotifyAppKey")
	proto.RegisterType((*ClientInfo)(nil), "Spotify.ClientInfo")
	proto.RegisterType((*ClientInfoFacebook)(nil), "Spotify.ClientInfoFacebook")
	proto.RegisterType((*APWelcome)(nil), "Spotify.APWelcome")
	proto.RegisterType((*AccountInfo)(nil), "Spotify.AccountInfo")
	proto.RegisterType((*AccountInfoSpotify)(nil), "Spotify.AccountInfoSpotify")
	proto.RegisterType((*AccountInfoFacebook)(nil), "Spotify.AccountInfoFacebook")
	proto.RegisterEnum("Spotify.AuthenticationType", AuthenticationType_name, AuthenticationType_value)
	proto.RegisterEnum("Spotify.AccountCreation", AccountCreation_name, AccountCreation_value)
	proto.RegisterEnum("Spotify.CpuFamily", CpuFamily_name, CpuFamily_value)
	proto.RegisterEnum("Spotify.Brand", Brand_name, Brand_value)
	proto.RegisterEnum("Spotify.Os", Os_name, Os_value)
	proto.RegisterEnum("Spotify.AccountType", AccountType_name, AccountType_value)
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x56, 0xcb, 0x72, 0xe3, 0x58,
	0x19, 0xc6, 0xce, 0xd5, 0xbf, 0xe3, 0x44, 0x7d, 0x92, 0xee, 0xb8, 0xa7, 0x2f, 0xf4, 0x98, 0x4b,
	0x85, 0x00, 0x5d, 0x10, 0x9a, 0x9e, 0x01, 0xa6, 0x8a, 0x52, 0x14, 0xb9, 0xa3, 0x8a, 0x23, 0xa9,
	0x8f, 0x94, 0x4e, 0x32, 0x1b, 0x21, 0xcb, 0xb2, 0xa3, 0x8a, 0x2c, 0xa9, 0x24, 0x79, 0x6a, 0xc2,
	0x82, 0x1d, 0x6f, 0x01, 0x3b, 0x78, 0x20, 0x36, 0x3c, 0x05, 0x8f, 0xc0, 0x82, 0xff, 0x1c, 0x1d,
	0x49, 0xb6, 0x93, 0x0e, 0x2b, 0xeb, 0xbf, 0x7f, 0xff, 0xf5, 0x18, 0xf6, 0xdc, 0x59, 0x7e, 0xe3,
	0x47, 0x79, 0xe0, 0xb9, 0x79, 0x10, 0x47, 0x6f, 0x93, 0x34, 0xce, 0x63, 0xb2, 0x61, 0x25, 0x71,
	0x1e, 0x8c, 0xef, 0x7a, 0xff, 0x58, 0x85, 0x7d, 0x25, 0x0c, 0x50, 0x81, 0xfa, 0x59, 0x12, 0x47,
	0x99, 0xaf, 0x46, 0x5e, 0x7a, 0x97, 0xe4, 0xfe, 0x88, 0xf4, 0xe1, 0x49, 0x18, 0x4f, 0x82, 0xc8,
	0xf1, 0x52, 0x7f, 0xc4, 0x5c, 0xb8, 0x61, 0xd6, 0x85, 0x37, 0xcd, 0x83, 0xf6, 0xd1, 0xf3, 0xb7,
	0xc2, 0xc1, 0xdb, 0x01, 0xd3, 0x50, 0x6a, 0x05, 0x2a, 0x85, 0x4b, 0x1c, 0xa2, 0x80, 0xe4, 0x7a,
	0x5e, 0x3c, 0x8b, 0x72, 0xe6, 0x89, 0xc3, 0xe8, 0xee, 0xbd, 0x69, 0x1c, 0x6c, 0x1f, 0x75, 0x2b,
	0x37, 0x72, 0xa1, 0xa0, 0x08, 0x39, 0xdd, 0x71, 0x17, 0x19, 0xc4, 0x86, 0xbd, 0x71, 0x10, 0x4d,
	0xfc, 0x34, 0x49, 0x03, 0x74, 0x94, 0x0a, 0xb4, 0xdd, 0xd7, 0xe8, 0xa8, 0x7d, 0xf4, 0x65, 0xe5,
	0xa8, 0x5f, 0x2b, 0x95, 0x19, 0x5d, 0x44, 0xcc, 0xe3, 0xee, 0xf8, 0xbe, 0x84, 0xfc, 0x0e, 0xda,
	0x89, 0xef, 0xa7, 0x0e, 0x96, 0xe7, 0xd6, 0xcf, 0xbb, 0x07, 0xdc, 0x59, 0x8d, 0xca, 0x44, 0x99,
	0xcd, 0x45, 0x85, 0x0f, 0x48, 0x2a, 0x06, 0x79, 0x07, 0xed, 0xec, 0x2e, 0xcb, 0xfd, 0xa9, 0x13,
	0x44, 0xe3, 0xb8, 0x7b, 0xc4, 0xeb, 0xb2, 0x5b, 0x99, 0x5a, 0x5c, 0xa6, 0xa1, 0x88, 0x42, 0x56,
	0x7d, 0x93, 0x9f, 0xc0, 0x76, 0x12, 0xba, 0xf9, 0x38, 0x4e, 0xa7, 0xce, 0x34, 0x1e, 0xf9, 0x61,
	0xf7, 0x1b, 0x8c, 0xd9, 0xa2, 0x9d, 0x92, 0x7b, 0xce, 0x98, 0x4c, 0xed, 0x3b, 0x3f, 0xcd, 0x30,
	0xa6, 0x93, 0xe5, 0x88, 0x78, 0xd2, 0xed, 0x17, 0x6a, 0x82, 0x6b, 0x71, 0x26, 0xf9, 0x35, 0xac,
	0xbb, 0x49, 0x72, 0xeb, 0xdf, 0x75, 0x4d, 0x8e, 0x7c, 0xae, 0x2d, 0xc1, 0x30, 0x2b, 0x3e, 0xe5,
	0x24, 0x39, 0xf3, 0xef, 0xa8, 0x50, 0x64, 0xb0, 0x3d, 0xde, 0xef, 0x02, 0xf6, 0xb7, 0xdc, 0xae,
	0x86, 0x5d, 0xcc, 0x42, 0x01, 0xdb, 0xab, 0xbe, 0x7b, 0x7f, 0x06, 0x69, 0xb9, 0xd1, 0xe4, 0x0b,
	0xd8, 0x9c, 0x65, 0x7e, 0x1a, 0xb9, 0x53, 0x1f, 0xa7, 0x82, 0xa1, 0xab, 0x68, 0xf2, 0x4b, 0x58,
	0xc9, 0xef, 0x12, 0xec, 0x72, 0x13, 0xbb, 0xfc, 0xa2, 0xee, 0xf2, 0xc2, 0x2c, 0xda, 0x77, 0x89,
	0x4f, 0x99, 0x1e, 0x79, 0x01, 0x2d, 0x36, 0xa6, 0xce, 0xc8, 0xcd, 0x5d, 0xde, 0xd1, 0x2d, 0xba,
	0xc9, 0x18, 0x27, 0x48, 0xf7, 0xfe, 0xd6, 0x80, 0xee, 0xe7, 0xba, 0x4a, 0xbe, 0x82, 0xb5, 0x49,
	0xea, 0x06, 0x11, 0x47, 0xf0, 0x99, 0x39, 0xf8, 0xc0, 0x14, 0x4a, 0x33, 0x5a, 0xe8, 0x93, 0x0f,
	0xd0, 0xbe, 0x99, 0xba, 0x9e, 0x93, 0x06, 0x89, 0x3f, 0x1d, 0xf1, 0x79, 0x6c, 0x1f, 0xfd, 0xf4,
	0x21, 0xf3, 0x53, 0x54, 0xa3, 0x5c, 0xab, 0xf2, 0x01, 0x37, 0x15, 0xaf, 0xf7, 0xc7, 0x05, 0x74,
	0x0b, 0xb1, 0xc8, 0x8f, 0xa0, 0xe3, 0x97, 0xeb, 0xe4, 0xb0, 0x36, 0xb1, 0xed, 0xd9, 0xa2, 0x5b,
	0x15, 0x13, 0x3b, 0xd3, 0xfb, 0x0d, 0xbc, 0x7a, 0x34, 0x1a, 0x21, 0xb0, 0xca, 0xe2, 0x09, 0x63,
	0xfe, 0xdd, 0xfb, 0x6b, 0x03, 0x76, 0x96, 0xa6, 0x93, 0xfc, 0x01, 0x20, 0x99, 0x0d, 0xc3, 0xc0,
	0x13, 0xa1, 0x58, 0x46, 0x2f, 0x1f, 0x98, 0x65, 0x93, 0x2b, 0xb1, 0xa1, 0x68, 0x25, 0xe5, 0x27,
	0xf9, 0x2d, 0x40, 0x1c, 0x8e, 0xca, 0x45, 0x28, 0xca, 0xf1, 0xec, 0x01, 0x63, 0x23, 0x1c, 0xd1,
	0x16, 0x6a, 0x16, 0x54, 0xef, 0x1d, 0xec, 0x3e, 0xe0, 0x98, 0xbc, 0x5a, 0x82, 0xc2, 0x80, 0xd7,
	0xc1, 0x7a, 0x23, 0xe8, 0x2c, 0x78, 0x24, 0x3f, 0x5c, 0xdc, 0xc3, 0xc2, 0x60, 0x7e, 0xdb, 0x8e,
	0xe0, 0xe9, 0x9c, 0x82, 0x93, 0x05, 0x93, 0xc8, 0xcd, 0x67, 0xa9, 0xcf, 0x47, 0x6c, 0x8b, 0xee,
	0xd6, 0xaa, 0x56, 0x29, 0xea, 0xfd, 0xa7, 0x09, 0x50, 0xaf, 0x21, 0x2e, 0x0b, 0x78, 0xc9, 0xcc,
	0x19, 0xbb, 0xd3, 0x20, 0x2c, 0x30, 0x6d, 0x1f, 0x91, 0x7a, 0xf0, 0x93, 0x59, 0x9f, 0x4b, 0x68,
	0xcb, 0x2b, 0x3f, 0x19, 0x2c, 0x66, 0x92, 0xcd, 0x86, 0x38, 0xa5, 0x3e, 0xaf, 0x4a, 0x87, 0x32,
	0x2f, 0x56, 0xc1, 0x21, 0xfb, 0xb0, 0xc1, 0x14, 0xfc, 0xef, 0x73, 0x3e, 0xb6, 0x1d, 0xba, 0x8e,
	0xa4, 0xfa, 0x7d, 0x4e, 0x7e, 0x0c, 0x6b, 0xc3, 0xd4, 0x8d, 0x46, 0xfc, 0xa4, 0x6c, 0x1f, 0x6d,
	0x57, 0x71, 0x8e, 0x19, 0x97, 0x16, 0x42, 0xe6, 0x9f, 0x7f, 0x38, 0xe3, 0xd0, 0x9d, 0x64, 0x78,
	0x43, 0xb8, 0x7f, 0xce, 0xea, 0x33, 0x0e, 0x2e, 0x46, 0x33, 0xce, 0xf0, 0x44, 0x30, 0xac, 0xed,
	0xca, 0x87, 0x91, 0x51, 0x64, 0xb3, 0x22, 0xc7, 0x99, 0x23, 0x2e, 0x02, 0x3f, 0x10, 0x1d, 0x6c,
	0x4d, 0xf6, 0xa9, 0x60, 0x90, 0xa7, 0xb0, 0x8e, 0x62, 0x06, 0xcd, 0xe4, 0xa2, 0xb5, 0x38, 0x63,
	0xc8, 0x7e, 0x0f, 0xcf, 0xe7, 0xee, 0x56, 0x3a, 0xe5, 0xab, 0x58, 0x5e, 0x99, 0x6f, 0xf9, 0x1e,
	0xef, 0xd7, 0x07, 0x4b, 0xc8, 0xc5, 0xbd, 0xc1, 0x3d, 0x1d, 0xf9, 0xdf, 0x05, 0x9e, 0xef, 0x04,
	0xa3, 0xee, 0xa8, 0xd8, 0xf9, 0x82, 0xa1, 0x8d, 0x7a, 0xff, 0x6c, 0xe0, 0x91, 0x58, 0x3a, 0x3b,
	0xa4, 0x0b, 0x1b, 0x25, 0xc0, 0x06, 0x66, 0xd1, 0xa1, 0x25, 0x49, 0x9e, 0xc1, 0x3a, 0x9a, 0xb2,
	0xf1, 0x68, 0xf2, 0x16, 0x0a, 0x8a, 0xbc, 0x84, 0x56, 0xdd, 0xdd, 0x95, 0x62, 0x72, 0x2a, 0x06,
	0x93, 0xb2, 0x23, 0xe3, 0x4e, 0xf0, 0x8c, 0x74, 0x57, 0x51, 0xda, 0xa2, 0x35, 0x83, 0xed, 0x9b,
	0xe7, 0x86, 0xe1, 0xd0, 0xf5, 0x6e, 0x9d, 0x1b, 0x37, 0xbb, 0xe9, 0xae, 0x15, 0xfb, 0x56, 0x32,
	0x4f, 0x91, 0xd7, 0x8b, 0x01, 0xea, 0x2b, 0xc7, 0x00, 0x86, 0xc1, 0x34, 0xc0, 0x5d, 0x44, 0x80,
	0x8d, 0x83, 0x4d, 0x5a, 0x92, 0xe4, 0xe7, 0xd0, 0x1c, 0x0f, 0x11, 0x1c, 0xdb, 0x84, 0x17, 0x0f,
	0x1c, 0xc8, 0xbe, 0xeb, 0xf9, 0xc3, 0x38, 0xbe, 0xa5, 0xa8, 0xc6, 0x8e, 0x61, 0xe8, 0x46, 0x93,
	0x19, 0xe2, 0x40, 0xd0, 0xbc, 0x30, 0x25, 0x8d, 0x0b, 0x4e, 0xee, 0x5b, 0xb1, 0xee, 0xe1, 0x22,
	0xdf, 0x04, 0x11, 0x2f, 0x66, 0x83, 0xdb, 0xb4, 0x04, 0x07, 0xab, 0xf9, 0xdf, 0x15, 0x68, 0xc9,
	0xe6, 0xa5, 0x1f, 0x7a, 0x31, 0xbf, 0xa7, 0xc4, 0x73, 0xa3, 0x38, 0xc2, 0xcb, 0x19, 0x3a, 0x73,
	0x57, 0x97, 0xe5, 0xff, 0xa4, 0x92, 0x5c, 0x94, 0xe7, 0x57, 0x83, 0x67, 0xe5, 0x8b, 0xcb, 0xc6,
	0xd4, 0xc1, 0x27, 0x79, 0x82, 0x17, 0x28, 0x88, 0xc4, 0x45, 0xde, 0x5b, 0x7e, 0x77, 0xf9, 0x29,
	0xde, 0x75, 0x6b, 0x62, 0xc0, 0x2d, 0xb4, 0x88, 0x7c, 0x84, 0x2f, 0xe6, 0x9e, 0xff, 0x65, 0x77,
	0xcf, 0x1f, 0x71, 0xb7, 0x3f, 0x67, 0xb7, 0xe0, 0xf2, 0x4f, 0xf0, 0x3a, 0xf5, 0x67, 0x99, 0x3b,
	0x0c, 0x7d, 0x87, 0x9f, 0xfd, 0xe5, 0x00, 0xb8, 0x4b, 0xff, 0xf7, 0xdd, 0x78, 0x51, 0xba, 0x60,
	0x32, 0x65, 0x31, 0x12, 0x9b, 0xf1, 0xcf, 0x46, 0xc0, 0x8d, 0x64, 0x33, 0xb1, 0xff, 0x19, 0x7b,
	0xd6, 0x97, 0x70, 0x9c, 0x39, 0x99, 0x8f, 0x36, 0x39, 0x5f, 0x49, 0x1c, 0x40, 0xe4, 0x58, 0x9c,
	0x81, 0x0f, 0xce, 0x56, 0x59, 0x5a, 0xfe, 0x80, 0x7e, 0xc3, 0xe7, 0xe3, 0x5e, 0x05, 0xf8, 0x0b,
	0xda, 0x76, 0x6b, 0x82, 0xfc, 0x82, 0x8f, 0x53, 0x7f, 0xe9, 0x2a, 0xcf, 0xa9, 0xcf, 0xcf, 0x53,
	0xef, 0x2f, 0xd0, 0x9e, 0x13, 0xe1, 0x75, 0xde, 0x10, 0x7b, 0xc5, 0x27, 0x65, 0x7e, 0x20, 0xe7,
	0xd4, 0x04, 0x8b, 0x96, 0xba, 0xe4, 0x6b, 0xd8, 0x1c, 0x0b, 0xaf, 0x62, 0x90, 0x1f, 0x8f, 0x5c,
	0x69, 0xf7, 0xf6, 0x80, 0xdc, 0x77, 0xdc, 0xbb, 0x84, 0xdd, 0x07, 0xcc, 0xc8, 0x97, 0xbc, 0x26,
	0x7e, 0x86, 0xdd, 0x8b, 0x6f, 0xfd, 0x48, 0x0c, 0x73, 0xbb, 0xe0, 0xd9, 0x8c, 0xb5, 0x34, 0xed,
	0xcd, 0xa5, 0x69, 0x3f, 0xfc, 0x77, 0x03, 0xe3, 0xdd, 0x6b, 0x32, 0x6e, 0x7b, 0x57, 0xbe, 0xb0,
	0x4f, 0x55, 0xdd, 0xd6, 0x14, 0xd9, 0xd6, 0x0c, 0xdd, 0xb9, 0xb0, 0x54, 0xea, 0x98, 0xb2, 0x65,
	0x49, 0x3f, 0xc0, 0xa5, 0xf8, 0xd9, 0x92, 0xd4, 0xb2, 0x0d, 0xaa, 0x9e, 0x38, 0x96, 0x69, 0xd8,
	0x5a, 0xff, 0xda, 0x51, 0x90, 0x60, 0x52, 0x79, 0x60, 0x49, 0x0d, 0xf2, 0x16, 0x0e, 0x1f, 0x56,
	0xef, 0xcb, 0x8a, 0x7a, 0x6c, 0x18, 0x67, 0x0b, 0xfa, 0x4d, 0xf2, 0x06, 0x5e, 0x2e, 0xeb, 0x0b,
	0xbf, 0xb6, 0x71, 0xa6, 0xea, 0xd2, 0x0a, 0xe6, 0xfd, 0x6a, 0x49, 0xa3, 0x72, 0x55, 0xa8, 0xac,
	0x1e, 0x5e, 0xc3, 0xce, 0xd2, 0x5f, 0x5b, 0xd2, 0x83, 0xd7, 0xb2, 0xa2, 0x18, 0x17, 0xba, 0xcd,
	0x02, 0x16, 0x76, 0xf2, 0xe0, 0x52, 0xbe, 0xb6, 0x1c, 0x93, 0x1a, 0xe7, 0xa6, 0x8d, 0x58, 0x1f,
	0xd1, 0xe1, 0xb4, 0x2a, 0xad, 0x1c, 0xfe, 0xbd, 0x01, 0xad, 0xea, 0xd5, 0x22, 0x3b, 0xd0, 0x56,
	0xcc, 0x0b, 0xe7, 0x42, 0x3f, 0xd3, 0x8d, 0x4b, 0x1d, 0xab, 0xd3, 0x86, 0x0d, 0xc6, 0xb8, 0xfa,
	0xfa, 0x3d, 0xfa, 0xdb, 0xc6, 0x9b, 0x57, 0x10, 0xce, 0xfb, 0x77, 0x98, 0x9b, 0x10, 0x9a, 0xa6,
	0x82, 0x69, 0x08, 0x21, 0x12, 0x4c, 0xb8, 0x5a, 0x0a, 0x65, 0x7a, 0x2e, 0xad, 0x91, 0x2d, 0xd8,
	0x64, 0x84, 0x26, 0xa3, 0x68, 0x9d, 0x00, 0xac, 0x33, 0xca, 0x3a, 0x95, 0x36, 0x4a, 0xc9, 0xb9,
	0x66, 0x5a, 0xd2, 0x26, 0x91, 0x60, 0x8b, 0x51, 0xc7, 0x03, 0x59, 0x39, 0xeb, 0x6b, 0xba, 0xd4,
	0x3a, 0x3c, 0x83, 0x35, 0xfe, 0xd8, 0x91, 0x5d, 0xd8, 0x39, 0xa6, 0xb2, 0x7e, 0x82, 0xe0, 0xf8,
	0xaf, 0x7a, 0x82, 0xf0, 0x3a, 0xd0, 0x2a, 0x98, 0x9a, 0xfe, 0x11, 0x01, 0x56, 0xe4, 0xa9, 0xad,
	0x20, 0x3e, 0xcc, 0xa6, 0x20, 0x75, 0xe3, 0x4c, 0x93, 0x31, 0xd9, 0x7f, 0x35, 0xa1, 0x69, 0x64,
	0x0c, 0xaa, 0x61, 0xcd, 0x25, 0x59, 0xd0, 0x97, 0x9a, 0x7e, 0x62, 0x5c, 0xb2, 0x1e, 0x23, 0x3e,
	0xa4, 0x0d, 0xeb, 0x0a, 0x7d, 0xa0, 0x4b, 0xfc, 0xd6, 0xcc, 0x53, 0x43, 0xc7, 0x72, 0x09, 0x91,
	0xf5, 0xfe, 0x57, 0x98, 0x21, 0x42, 0xc7, 0xef, 0x81, 0xa6, 0x5f, 0x5c, 0x61, 0x8a, 0x4f, 0xa0,
	0x53, 0x3b, 0x71, 0x14, 0x15, 0xf3, 0x2c, 0xfc, 0x22, 0x00, 0x6a, 0x68, 0x27, 0x98, 0x2b, 0x96,
	0x04, 0x69, 0x53, 0x1e, 0x9c, 0x63, 0xaa, 0x85, 0xb0, 0x4f, 0x55, 0xf5, 0xd8, 0x3a, 0x91, 0x5a,
	0xc2, 0x9e, 0x67, 0x7e, 0xac, 0x52, 0x7a, 0x2d, 0x81, 0x08, 0x60, 0x19, 0xba, 0x61, 0x49, 0x6d,
	0x96, 0x0d, 0x0b, 0x67, 0x7c, 0xd0, 0x6c, 0x55, 0x39, 0x95, 0xb6, 0x04, 0x96, 0x4b, 0xf3, 0x2b,
	0xa9, 0x23, 0x54, 0x0d, 0xfd, 0xec, 0xda, 0x90, 0xb6, 0x85, 0x6f, 0xf3, 0x54, 0x1b, 0xb0, 0xb2,
	0xee, 0x90, 0x16, 0xac, 0x31, 0xcd, 0x13, 0x49, 0x12, 0x8a, 0x9f, 0x8c, 0xc1, 0x27, 0x43, 0x7a,
	0x22, 0x10, 0xd9, 0x1a, 0x12, 0x44, 0x10, 0xf2, 0xa5, 0x71, 0x25, 0xed, 0x0a, 0xbd, 0x73, 0x55,
	0xfd, 0x60, 0x48, 0x7b, 0xa2, 0x0a, 0x1f, 0xf5, 0x2b, 0xdd, 0x36, 0xa4, 0xa7, 0x22, 0xf2, 0xb1,
	0x62, 0x48, 0xcf, 0x0e, 0x0f, 0xaa, 0x1b, 0xc3, 0x97, 0x0d, 0x9d, 0x88, 0x3d, 0xc7, 0xc2, 0xa2,
	0x93, 0x72, 0xbd, 0xa5, 0xc6, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x5f, 0x12, 0x81, 0x45,
	0x0e, 0x00, 0x00,
}
