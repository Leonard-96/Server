// Code generated by protoc-gen-go.
// source: DataFrame.proto
// DO NOT EDIT!

/*
Package DataFrame is a generated protocol buffer package.

It is generated from these files:
	DataFrame.proto

It has these top-level messages:
	Msg
	User
	UserDetail
	UserCustomAttr
	FriendLists
	SearchInfo
	Group
	GroupNumber
	GroupMsg
	PersonalMsg
	GroupNotice
*/
package DataFrame

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Msg struct {
	UserOpt       int32          `protobuf:"varint,1,opt,name=userOpt" json:"userOpt,omitempty"`
	OptResult     bool           `protobuf:"varint,2,opt,name=optResult" json:"optResult,omitempty"`
	ReceiveResult string         `protobuf:"bytes,3,opt,name=receiveResult" json:"receiveResult,omitempty"`
	User          *User          `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Friends       []*User        `protobuf:"bytes,5,rep,name=friends" json:"friends,omitempty"`
	Groups        []*Group       `protobuf:"bytes,6,rep,name=groups" json:"groups,omitempty"`
	GroupMsg      []*GroupMsg    `protobuf:"bytes,7,rep,name=groupMsg" json:"groupMsg,omitempty"`
	PersonalMsg   []*PersonalMsg `protobuf:"bytes,8,rep,name=personalMsg" json:"personalMsg,omitempty"`
	SrchInfo      *SearchInfo    `protobuf:"bytes,9,opt,name=srchInfo" json:"srchInfo,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Msg) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Msg) GetFriends() []*User {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *Msg) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *Msg) GetGroupMsg() []*GroupMsg {
	if m != nil {
		return m.GroupMsg
	}
	return nil
}

func (m *Msg) GetPersonalMsg() []*PersonalMsg {
	if m != nil {
		return m.PersonalMsg
	}
	return nil
}

func (m *Msg) GetSrchInfo() *SearchInfo {
	if m != nil {
		return m.SrchInfo
	}
	return nil
}

// 用户
type User struct {
	UesrName    string         `protobuf:"bytes,1,opt,name=uesrName" json:"uesrName,omitempty"`
	UserID      int32          `protobuf:"varint,2,opt,name=userID" json:"userID,omitempty"`
	UserPwd     string         `protobuf:"bytes,3,opt,name=userPwd" json:"userPwd,omitempty"`
	NickName    string         `protobuf:"bytes,4,opt,name=nickName" json:"nickName,omitempty"`
	Icon        []byte         `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	IconName    string         `protobuf:"bytes,6,opt,name=iconName" json:"iconName,omitempty"`
	UserDetail  *UserDetail    `protobuf:"bytes,7,opt,name=userDetail" json:"userDetail,omitempty"`
	IsOnline    bool           `protobuf:"varint,8,opt,name=isOnline" json:"isOnline,omitempty"`
	FriendLists []*FriendLists `protobuf:"bytes,9,rep,name=friendLists" json:"friendLists,omitempty"`
	Remark      string         `protobuf:"bytes,10,opt,name=remark" json:"remark,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetUserDetail() *UserDetail {
	if m != nil {
		return m.UserDetail
	}
	return nil
}

func (m *User) GetFriendLists() []*FriendLists {
	if m != nil {
		return m.FriendLists
	}
	return nil
}

// 用户详情信息
type UserDetail struct {
	UesrName    string            `protobuf:"bytes,1,opt,name=uesrName" json:"uesrName,omitempty"`
	Phone       string            `protobuf:"bytes,2,opt,name=phone" json:"phone,omitempty"`
	Addr        string            `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	QQ          string            `protobuf:"bytes,4,opt,name=QQ" json:"QQ,omitempty"`
	Wechat      string            `protobuf:"bytes,5,opt,name=wechat" json:"wechat,omitempty"`
	Sex         string            `protobuf:"bytes,6,opt,name=sex" json:"sex,omitempty"`
	Age         int32             `protobuf:"varint,7,opt,name=age" json:"age,omitempty"`
	IdNo        string            `protobuf:"bytes,8,opt,name=idNo" json:"idNo,omitempty"`
	Email       string            `protobuf:"bytes,9,opt,name=email" json:"email,omitempty"`
	CcNo        string            `protobuf:"bytes,10,opt,name=ccNo" json:"ccNo,omitempty"`
	ScNo        string            `protobuf:"bytes,11,opt,name=scNo" json:"scNo,omitempty"`
	StdNo       string            `protobuf:"bytes,12,opt,name=stdNo" json:"stdNo,omitempty"`
	CusteomAttr []*UserCustomAttr `protobuf:"bytes,13,rep,name=custeomAttr" json:"custeomAttr,omitempty"`
	UID         int32             `protobuf:"varint,14,opt,name=uID" json:"uID,omitempty"`
}

func (m *UserDetail) Reset()                    { *m = UserDetail{} }
func (m *UserDetail) String() string            { return proto.CompactTextString(m) }
func (*UserDetail) ProtoMessage()               {}
func (*UserDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserDetail) GetCusteomAttr() []*UserCustomAttr {
	if m != nil {
		return m.CusteomAttr
	}
	return nil
}

// 用户自定义属性
type UserCustomAttr struct {
	UserName    string   `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	AttrName    []string `protobuf:"bytes,2,rep,name=attrName" json:"attrName,omitempty"`
	AttrContent []string `protobuf:"bytes,3,rep,name=attrContent" json:"attrContent,omitempty"`
}

func (m *UserCustomAttr) Reset()                    { *m = UserCustomAttr{} }
func (m *UserCustomAttr) String() string            { return proto.CompactTextString(m) }
func (*UserCustomAttr) ProtoMessage()               {}
func (*UserCustomAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 好友分组
type FriendLists struct {
	ListNo   int32  `protobuf:"varint,1,opt,name=listNo" json:"listNo,omitempty"`
	ListName string `protobuf:"bytes,2,opt,name=listName" json:"listName,omitempty"`
}

func (m *FriendLists) Reset()                    { *m = FriendLists{} }
func (m *FriendLists) String() string            { return proto.CompactTextString(m) }
func (*FriendLists) ProtoMessage()               {}
func (*FriendLists) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 搜索请求包含的信息
type SearchInfo struct {
	SearchType   int32  `protobuf:"varint,1,opt,name=searchType" json:"searchType,omitempty"`
	SrchName     string `protobuf:"bytes,2,opt,name=srchName" json:"srchName,omitempty"`
	SrchAttrb    string `protobuf:"bytes,3,opt,name=srchAttrb" json:"srchAttrb,omitempty"`
	OnlyOnline   bool   `protobuf:"varint,4,opt,name=onlyOnline" json:"onlyOnline,omitempty"`
	AgeLow       int32  `protobuf:"varint,5,opt,name=ageLow" json:"ageLow,omitempty"`
	AgeHigh      int32  `protobuf:"varint,6,opt,name=ageHigh" json:"ageHigh,omitempty"`
	SelectMale   bool   `protobuf:"varint,7,opt,name=selectMale" json:"selectMale,omitempty"`
	SelectFemale bool   `protobuf:"varint,8,opt,name=selectFemale" json:"selectFemale,omitempty"`
	GroupNO      int32  `protobuf:"varint,9,opt,name=groupNO" json:"groupNO,omitempty"`
	Address      string `protobuf:"bytes,10,opt,name=address" json:"address,omitempty"`
	SinceId      int32  `protobuf:"varint,11,opt,name=sinceId" json:"sinceId,omitempty"`
}

func (m *SearchInfo) Reset()                    { *m = SearchInfo{} }
func (m *SearchInfo) String() string            { return proto.CompactTextString(m) }
func (*SearchInfo) ProtoMessage()               {}
func (*SearchInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 群
type Group struct {
	GroupID     int32          `protobuf:"varint,1,opt,name=groupID" json:"groupID,omitempty"`
	GroupName   string         `protobuf:"bytes,2,opt,name=groupName" json:"groupName,omitempty"`
	GroupIntro  string         `protobuf:"bytes,3,opt,name=groupIntro" json:"groupIntro,omitempty"`
	CreateTime  string         `protobuf:"bytes,4,opt,name=createTime" json:"createTime,omitempty"`
	GroupNumber []*GroupNumber `protobuf:"bytes,5,rep,name=groupNumber" json:"groupNumber,omitempty"`
	Notices     []*GroupNotice `protobuf:"bytes,6,rep,name=notices" json:"notices,omitempty"`
	Rank        int32          `protobuf:"varint,7,opt,name=rank" json:"rank,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Group) GetGroupNumber() []*GroupNumber {
	if m != nil {
		return m.GroupNumber
	}
	return nil
}

func (m *Group) GetNotices() []*GroupNotice {
	if m != nil {
		return m.Notices
	}
	return nil
}

// 群成员
type GroupNumber struct {
	GroupID  int32  `protobuf:"varint,1,opt,name=groupID" json:"groupID,omitempty"`
	NumberID int32  `protobuf:"varint,2,opt,name=numberID" json:"numberID,omitempty"`
	Remark   string `protobuf:"bytes,3,opt,name=remark" json:"remark,omitempty"`
	Identity int32  `protobuf:"varint,4,opt,name=identity" json:"identity,omitempty"`
}

func (m *GroupNumber) Reset()                    { *m = GroupNumber{} }
func (m *GroupNumber) String() string            { return proto.CompactTextString(m) }
func (*GroupNumber) ProtoMessage()               {}
func (*GroupNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// 群消息
type GroupMsg struct {
	GroupID    int32  `protobuf:"varint,1,opt,name=groupID" json:"groupID,omitempty"`
	SendTime   string `protobuf:"bytes,2,opt,name=sendTime" json:"sendTime,omitempty"`
	SenderID   int32  `protobuf:"varint,3,opt,name=senderID" json:"senderID,omitempty"`
	Content    string `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	ReadedTime int32  `protobuf:"varint,5,opt,name=readedTime" json:"readedTime,omitempty"`
	MsgType    int32  `protobuf:"varint,6,opt,name=msgType" json:"msgType,omitempty"`
}

func (m *GroupMsg) Reset()                    { *m = GroupMsg{} }
func (m *GroupMsg) String() string            { return proto.CompactTextString(m) }
func (*GroupMsg) ProtoMessage()               {}
func (*GroupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// 个人消息
type PersonalMsg struct {
	SenderID int32  `protobuf:"varint,1,opt,name=senderID" json:"senderID,omitempty"`
	RecverID int32  `protobuf:"varint,2,opt,name=recverID" json:"recverID,omitempty"`
	SendTime string `protobuf:"bytes,3,opt,name=sendTime" json:"sendTime,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
	MsgType  int32  `protobuf:"varint,5,opt,name=msgType" json:"msgType,omitempty"`
	IsReader bool   `protobuf:"varint,6,opt,name=isReader" json:"isReader,omitempty"`
}

func (m *PersonalMsg) Reset()                    { *m = PersonalMsg{} }
func (m *PersonalMsg) String() string            { return proto.CompactTextString(m) }
func (*PersonalMsg) ProtoMessage()               {}
func (*PersonalMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// 群公告
type GroupNotice struct {
	GroupID    int32  `protobuf:"varint,1,opt,name=groupID" json:"groupID,omitempty"`
	CreateTime string `protobuf:"bytes,2,opt,name=createTime" json:"createTime,omitempty"`
	CreateID   int32  `protobuf:"varint,3,opt,name=createID" json:"createID,omitempty"`
	Title      string `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Content    string `protobuf:"bytes,5,opt,name=content" json:"content,omitempty"`
	ModifyTime string `protobuf:"bytes,6,opt,name=modifyTime" json:"modifyTime,omitempty"`
}

func (m *GroupNotice) Reset()                    { *m = GroupNotice{} }
func (m *GroupNotice) String() string            { return proto.CompactTextString(m) }
func (*GroupNotice) ProtoMessage()               {}
func (*GroupNotice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*Msg)(nil), "Msg")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*UserDetail)(nil), "UserDetail")
	proto.RegisterType((*UserCustomAttr)(nil), "UserCustomAttr")
	proto.RegisterType((*FriendLists)(nil), "FriendLists")
	proto.RegisterType((*SearchInfo)(nil), "SearchInfo")
	proto.RegisterType((*Group)(nil), "Group")
	proto.RegisterType((*GroupNumber)(nil), "GroupNumber")
	proto.RegisterType((*GroupMsg)(nil), "GroupMsg")
	proto.RegisterType((*PersonalMsg)(nil), "PersonalMsg")
	proto.RegisterType((*GroupNotice)(nil), "GroupNotice")
}

var fileDescriptor0 = []byte{
	// 980 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x56, 0xff, 0x6e, 0xdc, 0x44,
	0x10, 0x96, 0x7d, 0xe7, 0xcb, 0xdd, 0x38, 0x6d, 0xd1, 0x0a, 0x21, 0x83, 0x10, 0x44, 0x16, 0x3f,
	0x2a, 0x21, 0x9d, 0x04, 0x3c, 0x41, 0xd5, 0x28, 0x25, 0x52, 0x9b, 0x34, 0x4b, 0xf3, 0x00, 0x8e,
	0xbd, 0xb9, 0x98, 0xf8, 0xec, 0xc3, 0xbb, 0xd7, 0x90, 0xb7, 0xa2, 0xe2, 0x01, 0xf8, 0x1f, 0x1e,
	0x85, 0x97, 0x60, 0x66, 0x76, 0xd7, 0x5e, 0xb7, 0x6a, 0xfe, 0xca, 0x7c, 0xdf, 0xcc, 0xcd, 0xce,
	0x7c, 0x3b, 0x3b, 0x0e, 0x3c, 0x39, 0x2e, 0x4c, 0x71, 0xd2, 0x17, 0x5b, 0xb5, 0xde, 0xf5, 0x9d,
	0xe9, 0xf2, 0xbf, 0x63, 0x98, 0xbd, 0xd2, 0x1b, 0x91, 0xc1, 0xc1, 0x5e, 0xab, 0xfe, 0x7c, 0x67,
	0xb2, 0xe8, 0x28, 0x7a, 0x9a, 0x48, 0x0f, 0xc5, 0x97, 0xb0, 0xea, 0x76, 0x46, 0x2a, 0xbd, 0x6f,
	0x4c, 0x16, 0xa3, 0x6f, 0x29, 0x47, 0x42, 0x7c, 0x03, 0x8f, 0x7a, 0x55, 0xaa, 0xfa, 0xad, 0x72,
	0x11, 0x33, 0x8c, 0x58, 0xc9, 0x29, 0x29, 0x3e, 0x87, 0x39, 0xa5, 0xcb, 0xe6, 0xe8, 0x4c, 0x7f,
	0x4a, 0xd6, 0x97, 0x08, 0x24, 0x53, 0xe2, 0x6b, 0x38, 0xb8, 0xee, 0x6b, 0xd5, 0x56, 0x3a, 0x4b,
	0x8e, 0x66, 0xa3, 0xd7, 0xb3, 0xe2, 0x2b, 0x58, 0x6c, 0xfa, 0x6e, 0xbf, 0xd3, 0xd9, 0x82, 0xfd,
	0x8b, 0xf5, 0x0b, 0x82, 0xd2, 0xb1, 0xe2, 0x5b, 0x58, 0xb2, 0x85, 0x5d, 0x64, 0x07, 0x1c, 0xb1,
	0xb2, 0x11, 0x48, 0xc8, 0xc1, 0x25, 0xd6, 0x90, 0xee, 0x54, 0xaf, 0xbb, 0xb6, 0x68, 0x28, 0x72,
	0xc9, 0x91, 0x87, 0xeb, 0xd7, 0x23, 0x27, 0xc3, 0x00, 0xf1, 0x3d, 0x2c, 0x75, 0x5f, 0xde, 0x9c,
	0xb6, 0xd7, 0x5d, 0xb6, 0xe2, 0xb2, 0xd3, 0xf5, 0xaf, 0xaa, 0x70, 0x94, 0x1c, 0x9c, 0xf9, 0xbb,
	0x18, 0xe6, 0x54, 0xb1, 0xf8, 0x02, 0x96, 0x7b, 0xa5, 0xfb, 0x33, 0x14, 0x97, 0x35, 0x5c, 0xc9,
	0x01, 0x8b, 0xcf, 0x60, 0x41, 0xdd, 0x9e, 0x1e, 0xb3, 0x82, 0x89, 0x74, 0xc8, 0xcb, 0xfe, 0xfa,
	0xae, 0x72, 0xc2, 0x79, 0x48, 0xd9, 0xda, 0xba, 0xbc, 0xe5, 0x6c, 0x73, 0x9b, 0xcd, 0x63, 0x21,
	0x60, 0x5e, 0x97, 0x5d, 0x8b, 0x82, 0x45, 0x4f, 0x0f, 0x25, 0xdb, 0x14, 0x4f, 0x7f, 0x39, 0x7e,
	0x61, 0xe3, 0x3d, 0x16, 0x3f, 0x00, 0x50, 0xda, 0x63, 0x65, 0x8a, 0xba, 0x41, 0x91, 0x6c, 0x37,
	0x97, 0x03, 0x25, 0x03, 0x37, 0x27, 0xd2, 0xe7, 0x6d, 0x53, 0xb7, 0x0a, 0x55, 0xa2, 0xeb, 0x1e,
	0x30, 0x89, 0x68, 0xaf, 0xe5, 0x65, 0xad, 0x8d, 0x46, 0x5d, 0xac, 0x88, 0x27, 0x23, 0x27, 0xc3,
	0x00, 0x6a, 0xbb, 0x57, 0xdb, 0xa2, 0xbf, 0xcd, 0x80, 0x4b, 0x72, 0x28, 0xff, 0x37, 0x06, 0xb8,
	0x9c, 0x1c, 0xf9, 0x51, 0xe5, 0x3e, 0x85, 0x64, 0x77, 0xd3, 0x61, 0x2d, 0x31, 0x3b, 0x2c, 0x20,
	0x05, 0x8a, 0xaa, 0xea, 0x9d, 0x68, 0x6c, 0x8b, 0xc7, 0x10, 0x5f, 0x5c, 0x38, 0xad, 0xe2, 0xdf,
	0x2f, 0xe8, 0xf0, 0x3b, 0x55, 0xde, 0x14, 0x86, 0x75, 0xc2, 0xc3, 0x2d, 0x12, 0x9f, 0xc0, 0x4c,
	0xab, 0x3f, 0x9c, 0x48, 0x64, 0x12, 0x53, 0x6c, 0x14, 0x0b, 0x93, 0x48, 0x32, 0x59, 0xe1, 0xea,
	0xac, 0x63, 0x01, 0x30, 0x3f, 0xd9, 0x54, 0x09, 0x56, 0x8f, 0x02, 0xae, 0x6c, 0x25, 0x0c, 0x28,
	0xb2, 0x2c, 0x31, 0xd2, 0x36, 0xc8, 0x36, 0x71, 0x9a, 0xb8, 0xd4, 0x72, 0x64, 0xd3, 0xaf, 0xb5,
	0xa1, 0x94, 0x87, 0xf6, 0xd7, 0x0c, 0xc4, 0x8f, 0x90, 0x96, 0x7b, 0x6d, 0x54, 0xb7, 0x7d, 0x66,
	0x4c, 0x9f, 0x3d, 0x62, 0x41, 0x9f, 0xf0, 0xd5, 0x3c, 0x47, 0xde, 0xd2, 0x32, 0x8c, 0xa1, 0x62,
	0xf7, 0x38, 0x47, 0x8f, 0x6d, 0xb1, 0x68, 0xe6, 0xbf, 0xc1, 0xe3, 0xe9, 0x0f, 0x58, 0x50, 0x64,
	0x26, 0x82, 0x3a, 0x4c, 0xbe, 0x02, 0x63, 0xd8, 0x17, 0xe3, 0x79, 0xe8, 0xf3, 0x58, 0x1c, 0x41,
	0x4a, 0xf6, 0xf3, 0xae, 0x35, 0xaa, 0xa5, 0xb7, 0x4c, 0xee, 0x90, 0xca, 0x9f, 0x41, 0x7a, 0x32,
	0xbd, 0xe0, 0x06, 0x0d, 0x6c, 0xcb, 0x6e, 0x0d, 0x87, 0xe8, 0x10, 0xb6, 0xec, 0x21, 0x5c, 0x80,
	0xc7, 0xf9, 0x3f, 0x78, 0xf9, 0xe3, 0x4b, 0xc2, 0xf7, 0x0d, 0x9a, 0xd1, 0x9b, 0xfb, 0x9d, 0x72,
	0x69, 0x02, 0x86, 0x52, 0xd1, 0x5b, 0x0b, 0x53, 0x79, 0x4c, 0xbb, 0x89, 0x6c, 0xea, 0xf9, 0xca,
	0xcd, 0xc2, 0x48, 0x50, 0xe6, 0xae, 0x6d, 0xee, 0xdd, 0x2c, 0xcf, 0x79, 0x96, 0x03, 0x86, 0x8a,
	0xc7, 0xbb, 0x7e, 0xd9, 0xdd, 0xf1, 0x80, 0x60, 0xf1, 0x16, 0xd1, 0xa3, 0x44, 0xeb, 0x97, 0x7a,
	0x73, 0xc3, 0x43, 0x82, 0xbb, 0xd0, 0x41, 0x5b, 0x6b, 0xa3, 0x4a, 0xf3, 0xaa, 0x68, 0xec, 0xbc,
	0x2c, 0x65, 0xc0, 0x88, 0x1c, 0x0e, 0x2d, 0x3a, 0xc1, 0xe1, 0x68, 0xfc, 0xfb, 0x99, 0x70, 0x94,
	0x9d, 0x97, 0xd2, 0xd9, 0x39, 0x0f, 0x12, 0x66, 0x77, 0x90, 0xcf, 0xc5, 0x41, 0x56, 0x5a, 0xbb,
	0x69, 0xf2, 0x90, 0x3c, 0xba, 0x6e, 0x4b, 0x75, 0x5a, 0xf1, 0x4c, 0xe1, 0x6f, 0x1c, 0xcc, 0xff,
	0x8b, 0x20, 0xe1, 0x6d, 0x37, 0xe4, 0xc5, 0xd9, 0x88, 0x82, 0xbc, 0xb8, 0x64, 0x50, 0x25, 0x7b,
	0xc4, 0x28, 0xe1, 0x48, 0x50, 0x4f, 0x36, 0xb0, 0x35, 0x7d, 0xe7, 0x44, 0x0c, 0x18, 0xf2, 0x97,
	0xbd, 0x2a, 0x8c, 0x7a, 0x53, 0x0f, 0xab, 0x28, 0x60, 0x68, 0x27, 0xd8, 0x64, 0xfb, 0xed, 0x15,
	0xae, 0xf8, 0xc4, 0xed, 0x84, 0x17, 0x23, 0x27, 0xc3, 0x00, 0xf1, 0x1d, 0x1c, 0xb4, 0x9d, 0xa9,
	0x4b, 0xe5, 0x17, 0xba, 0x8f, 0x65, 0x52, 0x7a, 0x27, 0x3d, 0xa2, 0xbe, 0x68, 0x6f, 0xdd, 0xab,
	0x64, 0x3b, 0xbf, 0x83, 0x34, 0xc8, 0xfb, 0x40, 0xcb, 0xb4, 0x3d, 0x39, 0x66, 0xd8, 0xb8, 0x03,
	0x0e, 0x96, 0xd2, 0x2c, 0x5c, 0x4a, 0xbc, 0xf8, 0x2a, 0x9c, 0xf1, 0xda, 0xdc, 0x73, 0x9b, 0xf8,
	0x1b, 0x8f, 0xf3, 0x3f, 0x23, 0x58, 0xfa, 0x8f, 0xca, 0xc3, 0xc7, 0x6a, 0x7c, 0x1b, 0xac, 0x94,
	0x9f, 0x55, 0x87, 0xbd, 0x8f, 0x4b, 0x9a, 0xd9, 0xf4, 0x1e, 0x53, 0xc6, 0xd2, 0xbd, 0x39, 0x2b,
	0xb0, 0x87, 0xa4, 0x3e, 0x4a, 0x5d, 0x29, 0x9b, 0xd3, 0xce, 0x69, 0xc0, 0xd0, 0x2f, 0xb7, 0x7a,
	0xc3, 0x4f, 0xc7, 0xcd, 0xaa, 0x83, 0xf9, 0xbb, 0x08, 0xd2, 0xe0, 0xeb, 0x36, 0x39, 0x3f, 0x7a,
	0xef, 0x7c, 0xf4, 0xe1, 0x07, 0xfb, 0x6d, 0x28, 0x97, 0xc7, 0x93, 0x9e, 0x66, 0xef, 0xf5, 0xf4,
	0xf1, 0xba, 0x83, 0xba, 0x92, 0x49, 0x5d, 0xf6, 0xfb, 0x22, 0xa9, 0x83, 0x9e, 0x4b, 0xe6, 0xef,
	0x8b, 0xc5, 0xf9, 0x5f, 0x91, 0xbf, 0x60, 0x1e, 0x82, 0x07, 0x94, 0x9e, 0x4e, 0x65, 0xfc, 0xc1,
	0x54, 0xe2, 0x29, 0x16, 0x8d, 0x6a, 0x7b, 0x4c, 0xab, 0x18, 0x2f, 0xb5, 0xf1, 0xc3, 0x6c, 0x41,
	0xd8, 0x4b, 0xf2, 0xc1, 0x1d, 0x6c, 0xbb, 0xaa, 0xbe, 0xbe, 0xe7, 0xb3, 0xec, 0x77, 0x23, 0x60,
	0xae, 0x16, 0xfc, 0xaf, 0xd4, 0xcf, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x88, 0xdf, 0xae, 0x9c,
	0x5d, 0x09, 0x00, 0x00,
}
