// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/preparams.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LocalPreParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaillierSK *PrivateKey `protobuf:"bytes,1,opt,name=paillierSK,proto3" json:"paillierSK,omitempty"` // Serialized Paillier Private Key
	Ntildei    []byte      `protobuf:"bytes,2,opt,name=ntildei,proto3" json:"ntildei,omitempty"`
	H1I        []byte      `protobuf:"bytes,3,opt,name=h1i,proto3" json:"h1i,omitempty"`
	H2I        []byte      `protobuf:"bytes,4,opt,name=h2i,proto3" json:"h2i,omitempty"`
	Alpha      []byte      `protobuf:"bytes,5,opt,name=alpha,proto3" json:"alpha,omitempty"`
	Beta       []byte      `protobuf:"bytes,6,opt,name=beta,proto3" json:"beta,omitempty"`
	P          []byte      `protobuf:"bytes,7,opt,name=p,proto3" json:"p,omitempty"`
	Q          []byte      `protobuf:"bytes,8,opt,name=q,proto3" json:"q,omitempty"`
}

func (x *LocalPreParams) Reset() {
	*x = LocalPreParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_preparams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalPreParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalPreParams) ProtoMessage() {}

func (x *LocalPreParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_preparams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalPreParams.ProtoReflect.Descriptor instead.
func (*LocalPreParams) Descriptor() ([]byte, []int) {
	return file_proto_preparams_proto_rawDescGZIP(), []int{0}
}

func (x *LocalPreParams) GetPaillierSK() *PrivateKey {
	if x != nil {
		return x.PaillierSK
	}
	return nil
}

func (x *LocalPreParams) GetNtildei() []byte {
	if x != nil {
		return x.Ntildei
	}
	return nil
}

func (x *LocalPreParams) GetH1I() []byte {
	if x != nil {
		return x.H1I
	}
	return nil
}

func (x *LocalPreParams) GetH2I() []byte {
	if x != nil {
		return x.H2I
	}
	return nil
}

func (x *LocalPreParams) GetAlpha() []byte {
	if x != nil {
		return x.Alpha
	}
	return nil
}

func (x *LocalPreParams) GetBeta() []byte {
	if x != nil {
		return x.Beta
	}
	return nil
}

func (x *LocalPreParams) GetP() []byte {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *LocalPreParams) GetQ() []byte {
	if x != nil {
		return x.Q
	}
	return nil
}

type LocalSecrets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xi      []byte `protobuf:"bytes,1,opt,name=xi,proto3" json:"xi,omitempty"`
	ShareID []byte `protobuf:"bytes,2,opt,name=shareID,proto3" json:"shareID,omitempty"`
}

func (x *LocalSecrets) Reset() {
	*x = LocalSecrets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_preparams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalSecrets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalSecrets) ProtoMessage() {}

func (x *LocalSecrets) ProtoReflect() protoreflect.Message {
	mi := &file_proto_preparams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalSecrets.ProtoReflect.Descriptor instead.
func (*LocalSecrets) Descriptor() ([]byte, []int) {
	return file_proto_preparams_proto_rawDescGZIP(), []int{1}
}

func (x *LocalSecrets) GetXi() []byte {
	if x != nil {
		return x.Xi
	}
	return nil
}

func (x *LocalSecrets) GetShareID() []byte {
	if x != nil {
		return x.ShareID
	}
	return nil
}

type ECPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Curve []byte `protobuf:"bytes,1,opt,name=curve,proto3" json:"curve,omitempty"`
	X     []byte `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y     []byte `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *ECPoint) Reset() {
	*x = ECPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_preparams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECPoint) ProtoMessage() {}

func (x *ECPoint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_preparams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECPoint.ProtoReflect.Descriptor instead.
func (*ECPoint) Descriptor() ([]byte, []int) {
	return file_proto_preparams_proto_rawDescGZIP(), []int{2}
}

func (x *ECPoint) GetCurve() []byte {
	if x != nil {
		return x.Curve
	}
	return nil
}

func (x *ECPoint) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *ECPoint) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	N []byte `protobuf:"bytes,1,opt,name=n,proto3" json:"n,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_preparams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_preparams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_proto_preparams_proto_rawDescGZIP(), []int{3}
}

func (x *PublicKey) GetN() []byte {
	if x != nil {
		return x.N
	}
	return nil
}

type PrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey *PublicKey `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	LambdaN   []byte     `protobuf:"bytes,2,opt,name=lambdaN,proto3" json:"lambdaN,omitempty"`
	PhiN      []byte     `protobuf:"bytes,3,opt,name=phiN,proto3" json:"phiN,omitempty"`
	P         []byte     `protobuf:"bytes,4,opt,name=p,proto3" json:"p,omitempty"`
	Q         []byte     `protobuf:"bytes,5,opt,name=q,proto3" json:"q,omitempty"`
}

func (x *PrivateKey) Reset() {
	*x = PrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_preparams_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateKey) ProtoMessage() {}

func (x *PrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_preparams_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateKey.ProtoReflect.Descriptor instead.
func (*PrivateKey) Descriptor() ([]byte, []int) {
	return file_proto_preparams_proto_rawDescGZIP(), []int{4}
}

func (x *PrivateKey) GetPublicKey() *PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *PrivateKey) GetLambdaN() []byte {
	if x != nil {
		return x.LambdaN
	}
	return nil
}

func (x *PrivateKey) GetPhiN() []byte {
	if x != nil {
		return x.PhiN
	}
	return nil
}

func (x *PrivateKey) GetP() []byte {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *PrivateKey) GetQ() []byte {
	if x != nil {
		return x.Q
	}
	return nil
}

type LocalPartySaveData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalPreParams *LocalPreParams `protobuf:"bytes,1,opt,name=localPreParams,proto3" json:"localPreParams,omitempty"`
	LocalSecrets   *LocalSecrets   `protobuf:"bytes,2,opt,name=localSecrets,proto3" json:"localSecrets,omitempty"`
	Ks             [][]byte        `protobuf:"bytes,3,rep,name=ks,proto3" json:"ks,omitempty"`
	Ntildej        [][]byte        `protobuf:"bytes,4,rep,name=ntildej,proto3" json:"ntildej,omitempty"`
	H1J            [][]byte        `protobuf:"bytes,5,rep,name=h1j,proto3" json:"h1j,omitempty"`
	H2J            [][]byte        `protobuf:"bytes,6,rep,name=h2j,proto3" json:"h2j,omitempty"`
	BigXj          []*ECPoint      `protobuf:"bytes,7,rep,name=bigXj,proto3" json:"bigXj,omitempty"`
	PaillierPKs    []*PublicKey    `protobuf:"bytes,8,rep,name=paillierPKs,proto3" json:"paillierPKs,omitempty"`
	EcdsaPub       *ECPoint        `protobuf:"bytes,9,opt,name=ecdsaPub,proto3" json:"ecdsaPub,omitempty"`
}

func (x *LocalPartySaveData) Reset() {
	*x = LocalPartySaveData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_preparams_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalPartySaveData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalPartySaveData) ProtoMessage() {}

func (x *LocalPartySaveData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_preparams_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalPartySaveData.ProtoReflect.Descriptor instead.
func (*LocalPartySaveData) Descriptor() ([]byte, []int) {
	return file_proto_preparams_proto_rawDescGZIP(), []int{5}
}

func (x *LocalPartySaveData) GetLocalPreParams() *LocalPreParams {
	if x != nil {
		return x.LocalPreParams
	}
	return nil
}

func (x *LocalPartySaveData) GetLocalSecrets() *LocalSecrets {
	if x != nil {
		return x.LocalSecrets
	}
	return nil
}

func (x *LocalPartySaveData) GetKs() [][]byte {
	if x != nil {
		return x.Ks
	}
	return nil
}

func (x *LocalPartySaveData) GetNtildej() [][]byte {
	if x != nil {
		return x.Ntildej
	}
	return nil
}

func (x *LocalPartySaveData) GetH1J() [][]byte {
	if x != nil {
		return x.H1J
	}
	return nil
}

func (x *LocalPartySaveData) GetH2J() [][]byte {
	if x != nil {
		return x.H2J
	}
	return nil
}

func (x *LocalPartySaveData) GetBigXj() []*ECPoint {
	if x != nil {
		return x.BigXj
	}
	return nil
}

func (x *LocalPartySaveData) GetPaillierPKs() []*PublicKey {
	if x != nil {
		return x.PaillierPKs
	}
	return nil
}

func (x *LocalPartySaveData) GetEcdsaPub() *ECPoint {
	if x != nil {
		return x.EcdsaPub
	}
	return nil
}

var File_proto_preparams_proto protoreflect.FileDescriptor

var file_proto_preparams_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x50, 0x72, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x70, 0x61,
	0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x53, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x69,
	0x6c, 0x6c, 0x69, 0x65, 0x72, 0x53, 0x4b, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x74, 0x69, 0x6c, 0x64,
	0x65, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x74, 0x69, 0x6c, 0x64, 0x65,
	0x69, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x31, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x68, 0x31, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x32, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x68, 0x32, 0x69, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x65, 0x74, 0x61, 0x12,
	0x0c, 0x0a, 0x01, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x70, 0x12, 0x0c, 0x0a,
	0x01, 0x71, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x71, 0x22, 0x38, 0x0a, 0x0c, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x78,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x78, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x07, 0x45, 0x43, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x01, 0x79, 0x22, 0x19, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x6e, 0x22, 0x80, 0x01,
	0x0a, 0x0a, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61,
	0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x62, 0x64, 0x61, 0x4e,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x68, 0x69, 0x4e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x70, 0x68, 0x69, 0x4e, 0x12, 0x0c, 0x0a, 0x01, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x01, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x71,
	0x22, 0xc2, 0x02, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53,
	0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x50, 0x72, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x72, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x31, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x02, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x74, 0x69, 0x6c, 0x64, 0x65, 0x6a, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x74, 0x69, 0x6c, 0x64, 0x65, 0x6a, 0x12, 0x10, 0x0a,
	0x03, 0x68, 0x31, 0x6a, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x68, 0x31, 0x6a, 0x12,
	0x10, 0x0a, 0x03, 0x68, 0x32, 0x6a, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x03, 0x68, 0x32,
	0x6a, 0x12, 0x1e, 0x0a, 0x05, 0x62, 0x69, 0x67, 0x58, 0x6a, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x45, 0x43, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x62, 0x69, 0x67, 0x58,
	0x6a, 0x12, 0x2c, 0x0a, 0x0b, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x4b, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x52, 0x0b, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x4b, 0x73, 0x12,
	0x24, 0x0a, 0x08, 0x65, 0x63, 0x64, 0x73, 0x61, 0x50, 0x75, 0x62, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x45, 0x43, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x63, 0x64,
	0x73, 0x61, 0x50, 0x75, 0x62, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_preparams_proto_rawDescOnce sync.Once
	file_proto_preparams_proto_rawDescData = file_proto_preparams_proto_rawDesc
)

func file_proto_preparams_proto_rawDescGZIP() []byte {
	file_proto_preparams_proto_rawDescOnce.Do(func() {
		file_proto_preparams_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_preparams_proto_rawDescData)
	})
	return file_proto_preparams_proto_rawDescData
}

var file_proto_preparams_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_preparams_proto_goTypes = []interface{}{
	(*LocalPreParams)(nil),     // 0: LocalPreParams
	(*LocalSecrets)(nil),       // 1: LocalSecrets
	(*ECPoint)(nil),            // 2: ECPoint
	(*PublicKey)(nil),          // 3: PublicKey
	(*PrivateKey)(nil),         // 4: PrivateKey
	(*LocalPartySaveData)(nil), // 5: LocalPartySaveData
}
var file_proto_preparams_proto_depIdxs = []int32{
	4, // 0: LocalPreParams.paillierSK:type_name -> PrivateKey
	3, // 1: PrivateKey.publicKey:type_name -> PublicKey
	0, // 2: LocalPartySaveData.localPreParams:type_name -> LocalPreParams
	1, // 3: LocalPartySaveData.localSecrets:type_name -> LocalSecrets
	2, // 4: LocalPartySaveData.bigXj:type_name -> ECPoint
	3, // 5: LocalPartySaveData.paillierPKs:type_name -> PublicKey
	2, // 6: LocalPartySaveData.ecdsaPub:type_name -> ECPoint
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_preparams_proto_init() }
func file_proto_preparams_proto_init() {
	if File_proto_preparams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_preparams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalPreParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_preparams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalSecrets); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_preparams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_preparams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_preparams_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_preparams_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalPartySaveData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_preparams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_preparams_proto_goTypes,
		DependencyIndexes: file_proto_preparams_proto_depIdxs,
		MessageInfos:      file_proto_preparams_proto_msgTypes,
	}.Build()
	File_proto_preparams_proto = out.File
	file_proto_preparams_proto_rawDesc = nil
	file_proto_preparams_proto_goTypes = nil
	file_proto_preparams_proto_depIdxs = nil
}