// Code generated by protoc-gen-go.
// source: Base.proto
// DO NOT EDIT!

/*
Package Base is a generated protocol buffer package.

It is generated from these files:
	Base.proto

It has these top-level messages:
*/
package Base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// 服务ID，对不同的协议进行归类
type ServiceID int32

const (
	ServiceID_LOGIN            ServiceID = 0
	ServiceID_FRIEND           ServiceID = 1
	ServiceID_MESSAGE          ServiceID = 2
	ServiceID_TRNASFER_SERVICE ServiceID = 3
	ServiceID_USER             ServiceID = 4
	ServiceID_OTHER            ServiceID = 5
)

var ServiceID_name = map[int32]string{
	0: "LOGIN",
	1: "FRIEND",
	2: "MESSAGE",
	3: "TRNASFER_SERVICE",
	4: "USER",
	5: "OTHER",
}
var ServiceID_value = map[string]int32{
	"LOGIN":            0,
	"FRIEND":           1,
	"MESSAGE":          2,
	"TRNASFER_SERVICE": 3,
	"USER":             4,
	"OTHER":            5,
}

func (x ServiceID) String() string {
	return proto.EnumName(ServiceID_name, int32(x))
}
func (ServiceID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 登陆相关命令
type LoginCmdID int32

const (
	LoginCmdID_LOGIN_REQ        LoginCmdID = 0
	LoginCmdID_LOGIN_RES        LoginCmdID = 1
	LoginCmdID_LOGOUT_REQ       LoginCmdID = 2
	LoginCmdID_LOGOUT_RES       LoginCmdID = 3
	LoginCmdID_DEVICE_TOKEN_REQ LoginCmdID = 4
	LoginCmdID_DEVICE_TOKEN_RES LoginCmdID = 5
	LoginCmdID_MSG_SERVER_REQ   LoginCmdID = 6
	LoginCmdID_MSG_SERVER_RES   LoginCmdID = 7
)

var LoginCmdID_name = map[int32]string{
	0: "LOGIN_REQ",
	1: "LOGIN_RES",
	2: "LOGOUT_REQ",
	3: "LOGOUT_RES",
	4: "DEVICE_TOKEN_REQ",
	5: "DEVICE_TOKEN_RES",
	6: "MSG_SERVER_REQ",
	7: "MSG_SERVER_RES",
}
var LoginCmdID_value = map[string]int32{
	"LOGIN_REQ":        0,
	"LOGIN_RES":        1,
	"LOGOUT_REQ":       2,
	"LOGOUT_RES":       3,
	"DEVICE_TOKEN_REQ": 4,
	"DEVICE_TOKEN_RES": 5,
	"MSG_SERVER_REQ":   6,
	"MSG_SERVER_RES":   7,
}

func (x LoginCmdID) String() string {
	return proto.EnumName(LoginCmdID_name, int32(x))
}
func (LoginCmdID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 好友、会话等相关命令。
type FriendCmdID int32

const (
	FriendCmdID_CID_FRIEND_REQ_ALL FriendCmdID = 0
	FriendCmdID_CID_FRIEND_RES_ALL FriendCmdID = 1
)

var FriendCmdID_name = map[int32]string{
	0: "CID_FRIEND_REQ_ALL",
	1: "CID_FRIEND_RES_ALL",
}
var FriendCmdID_value = map[string]int32{
	"CID_FRIEND_REQ_ALL": 0,
	"CID_FRIEND_RES_ALL": 1,
}

func (x FriendCmdID) String() string {
	return proto.EnumName(FriendCmdID_name, int32(x))
}
func (FriendCmdID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 消息相关命令
type MessageCmdID int32

const (
	MessageCmdID_CID_MSG_DATA MessageCmdID = 0
)

var MessageCmdID_name = map[int32]string{
	0: "CID_MSG_DATA",
}
var MessageCmdID_value = map[string]int32{
	"CID_MSG_DATA": 0,
}

func (x MessageCmdID) String() string {
	return proto.EnumName(MessageCmdID_name, int32(x))
}
func (MessageCmdID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 中转服务命令
type TransferServiceCmdID int32

const (
	TransferServiceCmdID_CID_TRANSFER TransferServiceCmdID = 0
)

var TransferServiceCmdID_name = map[int32]string{
	0: "CID_TRANSFER",
}
var TransferServiceCmdID_value = map[string]int32{
	"CID_TRANSFER": 0,
}

func (x TransferServiceCmdID) String() string {
	return proto.EnumName(TransferServiceCmdID_name, int32(x))
}
func (TransferServiceCmdID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 其他命令
type OtherCmdID int32

const (
	OtherCmdID_CID_OTHER_HEARTBEAT OtherCmdID = 0
)

var OtherCmdID_name = map[int32]string{
	0: "CID_OTHER_HEARTBEAT",
}
var OtherCmdID_value = map[string]int32{
	"CID_OTHER_HEARTBEAT": 0,
}

func (x OtherCmdID) String() string {
	return proto.EnumName(OtherCmdID_name, int32(x))
}
func (OtherCmdID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 消息类型
type MsgType int32

const (
	MsgType_MSG_TYPE_SINGLE_TEXT MsgType = 0
)

var MsgType_name = map[int32]string{
	0: "MSG_TYPE_SINGLE_TEXT",
}
var MsgType_value = map[string]int32{
	"MSG_TYPE_SINGLE_TEXT": 0,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}
func (MsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// 客户端类型
type ClientType int32

const (
	ClientType_WEB     ClientType = 0
	ClientType_ANDROID ClientType = 1
	ClientType_IOS     ClientType = 2
	ClientType_MAC     ClientType = 3
	ClientType_WAP     ClientType = 4
	ClientType_WEIXIN  ClientType = 5
	ClientType_WINDOWS ClientType = 6
)

var ClientType_name = map[int32]string{
	0: "WEB",
	1: "ANDROID",
	2: "IOS",
	3: "MAC",
	4: "WAP",
	5: "WEIXIN",
	6: "WINDOWS",
}
var ClientType_value = map[string]int32{
	"WEB":     0,
	"ANDROID": 1,
	"IOS":     2,
	"MAC":     3,
	"WAP":     4,
	"WEIXIN":  5,
	"WINDOWS": 6,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}
func (ClientType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterEnum("Base.ServiceID", ServiceID_name, ServiceID_value)
	proto.RegisterEnum("Base.LoginCmdID", LoginCmdID_name, LoginCmdID_value)
	proto.RegisterEnum("Base.FriendCmdID", FriendCmdID_name, FriendCmdID_value)
	proto.RegisterEnum("Base.MessageCmdID", MessageCmdID_name, MessageCmdID_value)
	proto.RegisterEnum("Base.TransferServiceCmdID", TransferServiceCmdID_name, TransferServiceCmdID_value)
	proto.RegisterEnum("Base.OtherCmdID", OtherCmdID_name, OtherCmdID_value)
	proto.RegisterEnum("Base.MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("Base.ClientType", ClientType_name, ClientType_value)
}

var fileDescriptor0 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x73, 0xd3, 0x30,
	0x10, 0xad, 0x5b, 0xdb, 0x21, 0xdb, 0xd2, 0x11, 0x22, 0x7c, 0x1c, 0x99, 0x61, 0x98, 0x61, 0x7c,
	0xc8, 0x85, 0x33, 0x07, 0xd9, 0x56, 0x1c, 0x0d, 0xb6, 0x55, 0x24, 0xb5, 0x2e, 0x5c, 0x3c, 0x69,
	0xab, 0x06, 0x0f, 0x24, 0xce, 0xd8, 0x81, 0x19, 0xf8, 0x31, 0xfc, 0x56, 0x56, 0x36, 0x07, 0x02,
	0xdc, 0x56, 0x6f, 0xdf, 0xee, 0xdb, 0x7d, 0x2b, 0x80, 0x78, 0xd5, 0xdb, 0xf9, 0xae, 0x6b, 0xf7,
	0x2d, 0xf5, 0x5d, 0x1c, 0x7d, 0x84, 0xa9, 0xb6, 0xdd, 0xb7, 0xe6, 0xd6, 0x8a, 0x94, 0x4e, 0x21,
	0xc8, 0x65, 0x26, 0x4a, 0x72, 0x44, 0x01, 0xc2, 0x85, 0x12, 0xbc, 0x4c, 0x89, 0x47, 0x4f, 0x61,
	0x52, 0x70, 0xad, 0x59, 0xc6, 0xc9, 0x31, 0x9d, 0x01, 0x31, 0xaa, 0x64, 0x7a, 0xc1, 0x55, 0xad,
	0xb9, 0xba, 0x12, 0x09, 0x27, 0x27, 0xf4, 0x01, 0xf8, 0x97, 0xf8, 0x22, 0xbe, 0xeb, 0x21, 0xcd,
	0x12, 0xc3, 0x20, 0xfa, 0xe9, 0x01, 0xe4, 0xed, 0xba, 0xd9, 0x26, 0x9b, 0x3b, 0xec, 0xfe, 0x10,
	0xa6, 0x43, 0xf7, 0x5a, 0xf1, 0xf7, 0xa8, 0xf0, 0xc7, 0x53, 0xa3, 0xc8, 0x39, 0x72, 0x65, 0x26,
	0x2f, 0xcd, 0x90, 0x3e, 0x3e, 0x78, 0x6b, 0x54, 0x40, 0xdd, 0x94, 0x3b, 0xb5, 0xda, 0xc8, 0x77,
	0x7c, 0x6c, 0xe2, 0xff, 0x07, 0xd5, 0x24, 0xa0, 0x14, 0xce, 0x0b, 0x9d, 0x0d, 0xe3, 0xe1, 0x94,
	0x8e, 0x19, 0xfe, 0x83, 0x69, 0x32, 0x89, 0xde, 0xc2, 0xe9, 0xa2, 0x6b, 0xec, 0xf6, 0x6e, 0x1c,
	0xf0, 0x29, 0xd0, 0x44, 0xa4, 0xf5, 0xb8, 0xb7, 0x2b, 0xab, 0x59, 0x9e, 0xe3, 0xa4, 0x7f, 0xe3,
	0x7a, 0xc0, 0xbd, 0xe8, 0x05, 0x9c, 0x15, 0xb6, 0xef, 0x57, 0x6b, 0x3b, 0xd6, 0x13, 0x38, 0x73,
	0x3c, 0x27, 0x93, 0x32, 0xc3, 0xc8, 0x51, 0xf4, 0x1a, 0x66, 0xa6, 0x5b, 0x6d, 0xfb, 0x7b, 0xdb,
	0xfd, 0x76, 0xf9, 0x80, 0x69, 0x14, 0x2b, 0x9d, 0x91, 0xc8, 0x7c, 0x05, 0x20, 0xf7, 0x9f, 0x6c,
	0x37, 0xe6, 0x9f, 0xc1, 0x63, 0x97, 0x1f, 0x8c, 0xac, 0x97, 0x9c, 0x29, 0x13, 0x73, 0x66, 0x90,
	0xf6, 0x12, 0x4f, 0xd1, 0xaf, 0xcd, 0xf7, 0x9d, 0xa5, 0xcf, 0x61, 0xe6, 0x94, 0xcc, 0x87, 0x0b,
	0x5e, 0x6b, 0x51, 0x66, 0x39, 0x7a, 0xc0, 0xaf, 0x1d, 0xe9, 0x0a, 0x20, 0xf9, 0x82, 0x6b, 0xed,
	0x07, 0xde, 0x04, 0x4e, 0x2a, 0x1e, 0xe3, 0x1a, 0x78, 0x46, 0x56, 0xa6, 0x4a, 0x0a, 0x77, 0x53,
	0x44, 0x85, 0xd4, 0xe8, 0x33, 0x06, 0x05, 0x4b, 0xd0, 0x60, 0xc7, 0x63, 0x17, 0xe8, 0x29, 0x9e,
	0xbe, 0xe2, 0xe2, 0x1a, 0xbf, 0x41, 0xe0, 0x6a, 0x2a, 0x51, 0xa6, 0xb2, 0xd2, 0x24, 0x8c, 0x9f,
	0xc0, 0xa3, 0xdb, 0x76, 0x33, 0xff, 0x31, 0xff, 0xdc, 0x8c, 0x7f, 0xe8, 0xe6, 0xeb, 0xfd, 0xd2,
	0xbb, 0x09, 0x87, 0xf8, 0xcd, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0xf6, 0x45, 0x48, 0x5d,
	0x02, 0x00, 0x00,
}
