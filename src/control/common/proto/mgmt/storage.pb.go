// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package mgmt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmptyParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyParams) Reset()         { *m = EmptyParams{} }
func (m *EmptyParams) String() string { return proto.CompactTextString(m) }
func (*EmptyParams) ProtoMessage()    {}
func (*EmptyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{0}
}
func (m *EmptyParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyParams.Unmarshal(m, b)
}
func (m *EmptyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyParams.Marshal(b, m, deterministic)
}
func (dst *EmptyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyParams.Merge(dst, src)
}
func (m *EmptyParams) XXX_Size() int {
	return xxx_messageInfo_EmptyParams.Size(m)
}
func (m *EmptyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyParams.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyParams proto.InternalMessageInfo

type FeatureName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureName) Reset()         { *m = FeatureName{} }
func (m *FeatureName) String() string { return proto.CompactTextString(m) }
func (*FeatureName) ProtoMessage()    {}
func (*FeatureName) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{1}
}
func (m *FeatureName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureName.Unmarshal(m, b)
}
func (m *FeatureName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureName.Marshal(b, m, deterministic)
}
func (dst *FeatureName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureName.Merge(dst, src)
}
func (m *FeatureName) XXX_Size() int {
	return xxx_messageInfo_FeatureName.Size(m)
}
func (m *FeatureName) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureName.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureName proto.InternalMessageInfo

func (m *FeatureName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Category struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{2}
}
func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (dst *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(dst, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

// Feature represents a management task that can be performed by server.
type Feature struct {
	// The category of capabilities this feature belongs to.
	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// The name of the feature.
	Fname *FeatureName `protobuf:"bytes,2,opt,name=fname,proto3" json:"fname,omitempty"`
	// The description of the feature.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{3}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *Feature) GetFname() *FeatureName {
	if m != nil {
		return m.Fname
	}
	return nil
}

func (m *Feature) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// NvmeNamespace represent NVMe namespaces available on controller.
type NvmeNamespace struct {
	// namespace id
	Id int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// device capacity in GBytes
	Capacity             int32    `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NvmeNamespace) Reset()         { *m = NvmeNamespace{} }
func (m *NvmeNamespace) String() string { return proto.CompactTextString(m) }
func (*NvmeNamespace) ProtoMessage()    {}
func (*NvmeNamespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{4}
}
func (m *NvmeNamespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NvmeNamespace.Unmarshal(m, b)
}
func (m *NvmeNamespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NvmeNamespace.Marshal(b, m, deterministic)
}
func (dst *NvmeNamespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NvmeNamespace.Merge(dst, src)
}
func (m *NvmeNamespace) XXX_Size() int {
	return xxx_messageInfo_NvmeNamespace.Size(m)
}
func (m *NvmeNamespace) XXX_DiscardUnknown() {
	xxx_messageInfo_NvmeNamespace.DiscardUnknown(m)
}

var xxx_messageInfo_NvmeNamespace proto.InternalMessageInfo

func (m *NvmeNamespace) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NvmeNamespace) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

// NvmeController represents an NVMe Controller.
type NvmeController struct {
	// The id of the controller.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The model name of the controller.
	Model string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	// The serial number of the controller.
	Serial string `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	// The pci address of the controller.
	Pciaddr string `protobuf:"bytes,4,opt,name=pciaddr,proto3" json:"pciaddr,omitempty"`
	// The firmware revision of the controller.
	Fwrev string `protobuf:"bytes,5,opt,name=fwrev,proto3" json:"fwrev,omitempty"`
	// NvmeNamespaces created on this controller.
	Namespace            []*NvmeNamespace `protobuf:"bytes,6,rep,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NvmeController) Reset()         { *m = NvmeController{} }
func (m *NvmeController) String() string { return proto.CompactTextString(m) }
func (*NvmeController) ProtoMessage()    {}
func (*NvmeController) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{5}
}
func (m *NvmeController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NvmeController.Unmarshal(m, b)
}
func (m *NvmeController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NvmeController.Marshal(b, m, deterministic)
}
func (dst *NvmeController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NvmeController.Merge(dst, src)
}
func (m *NvmeController) XXX_Size() int {
	return xxx_messageInfo_NvmeController.Size(m)
}
func (m *NvmeController) XXX_DiscardUnknown() {
	xxx_messageInfo_NvmeController.DiscardUnknown(m)
}

var xxx_messageInfo_NvmeController proto.InternalMessageInfo

func (m *NvmeController) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NvmeController) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *NvmeController) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *NvmeController) GetPciaddr() string {
	if m != nil {
		return m.Pciaddr
	}
	return ""
}

func (m *NvmeController) GetFwrev() string {
	if m != nil {
		return m.Fwrev
	}
	return ""
}

func (m *NvmeController) GetNamespace() []*NvmeNamespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

type UpdateNvmeParams struct {
	// The pci address of the controller to update firmware on.
	Pciaddr string `protobuf:"bytes,1,opt,name=pciaddr,proto3" json:"pciaddr,omitempty"`
	// Filesystem path containing firmware image
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Firmware slot (register) to update
	Slot                 int32    `protobuf:"varint,3,opt,name=slot,proto3" json:"slot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNvmeParams) Reset()         { *m = UpdateNvmeParams{} }
func (m *UpdateNvmeParams) String() string { return proto.CompactTextString(m) }
func (*UpdateNvmeParams) ProtoMessage()    {}
func (*UpdateNvmeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{6}
}
func (m *UpdateNvmeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNvmeParams.Unmarshal(m, b)
}
func (m *UpdateNvmeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNvmeParams.Marshal(b, m, deterministic)
}
func (dst *UpdateNvmeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNvmeParams.Merge(dst, src)
}
func (m *UpdateNvmeParams) XXX_Size() int {
	return xxx_messageInfo_UpdateNvmeParams.Size(m)
}
func (m *UpdateNvmeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNvmeParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNvmeParams proto.InternalMessageInfo

func (m *UpdateNvmeParams) GetPciaddr() string {
	if m != nil {
		return m.Pciaddr
	}
	return ""
}

func (m *UpdateNvmeParams) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UpdateNvmeParams) GetSlot() int32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

type FioConfigPath struct {
	// Filesystem path containing fio job configuration
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FioConfigPath) Reset()         { *m = FioConfigPath{} }
func (m *FioConfigPath) String() string { return proto.CompactTextString(m) }
func (*FioConfigPath) ProtoMessage()    {}
func (*FioConfigPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{7}
}
func (m *FioConfigPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FioConfigPath.Unmarshal(m, b)
}
func (m *FioConfigPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FioConfigPath.Marshal(b, m, deterministic)
}
func (dst *FioConfigPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FioConfigPath.Merge(dst, src)
}
func (m *FioConfigPath) XXX_Size() int {
	return xxx_messageInfo_FioConfigPath.Size(m)
}
func (m *FioConfigPath) XXX_DiscardUnknown() {
	xxx_messageInfo_FioConfigPath.DiscardUnknown(m)
}

var xxx_messageInfo_FioConfigPath proto.InternalMessageInfo

func (m *FioConfigPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type BurnInNvmeParams struct {
	// The pci address of the controller to perform burn-in on.
	Pciaddr string `protobuf:"bytes,1,opt,name=pciaddr,proto3" json:"pciaddr,omitempty"`
	// FIO workload configuration file path
	Path                 *FioConfigPath `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BurnInNvmeParams) Reset()         { *m = BurnInNvmeParams{} }
func (m *BurnInNvmeParams) String() string { return proto.CompactTextString(m) }
func (*BurnInNvmeParams) ProtoMessage()    {}
func (*BurnInNvmeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{8}
}
func (m *BurnInNvmeParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BurnInNvmeParams.Unmarshal(m, b)
}
func (m *BurnInNvmeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BurnInNvmeParams.Marshal(b, m, deterministic)
}
func (dst *BurnInNvmeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnInNvmeParams.Merge(dst, src)
}
func (m *BurnInNvmeParams) XXX_Size() int {
	return xxx_messageInfo_BurnInNvmeParams.Size(m)
}
func (m *BurnInNvmeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnInNvmeParams.DiscardUnknown(m)
}

var xxx_messageInfo_BurnInNvmeParams proto.InternalMessageInfo

func (m *BurnInNvmeParams) GetPciaddr() string {
	if m != nil {
		return m.Pciaddr
	}
	return ""
}

func (m *BurnInNvmeParams) GetPath() *FioConfigPath {
	if m != nil {
		return m.Path
	}
	return nil
}

type BurnInNvmeReport struct {
	// Report from running burn-in reported by fio
	Report               string   `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BurnInNvmeReport) Reset()         { *m = BurnInNvmeReport{} }
func (m *BurnInNvmeReport) String() string { return proto.CompactTextString(m) }
func (*BurnInNvmeReport) ProtoMessage()    {}
func (*BurnInNvmeReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{9}
}
func (m *BurnInNvmeReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BurnInNvmeReport.Unmarshal(m, b)
}
func (m *BurnInNvmeReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BurnInNvmeReport.Marshal(b, m, deterministic)
}
func (dst *BurnInNvmeReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnInNvmeReport.Merge(dst, src)
}
func (m *BurnInNvmeReport) XXX_Size() int {
	return xxx_messageInfo_BurnInNvmeReport.Size(m)
}
func (m *BurnInNvmeReport) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnInNvmeReport.DiscardUnknown(m)
}

var xxx_messageInfo_BurnInNvmeReport proto.InternalMessageInfo

func (m *BurnInNvmeReport) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

// ScmModule represent Storage Class Memory modules installed.
type ScmModule struct {
	// The uid of the module.
	// string uid = 1;
	// The physical id of the module.
	Physicalid uint32 `protobuf:"varint,1,opt,name=physicalid,proto3" json:"physicalid,omitempty"`
	// The device handle of the module.
	// string handle = 3;
	// The channel id where module is installed.
	Channel uint32 `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
	// The channel position where module is installed.
	Channelpos uint32 `protobuf:"varint,3,opt,name=channelpos,proto3" json:"channelpos,omitempty"`
	// The memory controller id attached to module.
	Memctrlr uint32 `protobuf:"varint,4,opt,name=memctrlr,proto3" json:"memctrlr,omitempty"`
	// The socket id attached to module.
	Socket uint32 `protobuf:"varint,5,opt,name=socket,proto3" json:"socket,omitempty"`
	// The serial number of the module.
	// string serial = 8;
	// The capacity of the module.
	Capacity             uint64   `protobuf:"varint,6,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScmModule) Reset()         { *m = ScmModule{} }
func (m *ScmModule) String() string { return proto.CompactTextString(m) }
func (*ScmModule) ProtoMessage()    {}
func (*ScmModule) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{10}
}
func (m *ScmModule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScmModule.Unmarshal(m, b)
}
func (m *ScmModule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScmModule.Marshal(b, m, deterministic)
}
func (dst *ScmModule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScmModule.Merge(dst, src)
}
func (m *ScmModule) XXX_Size() int {
	return xxx_messageInfo_ScmModule.Size(m)
}
func (m *ScmModule) XXX_DiscardUnknown() {
	xxx_messageInfo_ScmModule.DiscardUnknown(m)
}

var xxx_messageInfo_ScmModule proto.InternalMessageInfo

func (m *ScmModule) GetPhysicalid() uint32 {
	if m != nil {
		return m.Physicalid
	}
	return 0
}

func (m *ScmModule) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *ScmModule) GetChannelpos() uint32 {
	if m != nil {
		return m.Channelpos
	}
	return 0
}

func (m *ScmModule) GetMemctrlr() uint32 {
	if m != nil {
		return m.Memctrlr
	}
	return 0
}

func (m *ScmModule) GetSocket() uint32 {
	if m != nil {
		return m.Socket
	}
	return 0
}

func (m *ScmModule) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type FormatStorageParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FormatStorageParams) Reset()         { *m = FormatStorageParams{} }
func (m *FormatStorageParams) String() string { return proto.CompactTextString(m) }
func (*FormatStorageParams) ProtoMessage()    {}
func (*FormatStorageParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{11}
}
func (m *FormatStorageParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormatStorageParams.Unmarshal(m, b)
}
func (m *FormatStorageParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormatStorageParams.Marshal(b, m, deterministic)
}
func (dst *FormatStorageParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormatStorageParams.Merge(dst, src)
}
func (m *FormatStorageParams) XXX_Size() int {
	return xxx_messageInfo_FormatStorageParams.Size(m)
}
func (m *FormatStorageParams) XXX_DiscardUnknown() {
	xxx_messageInfo_FormatStorageParams.DiscardUnknown(m)
}

var xxx_messageInfo_FormatStorageParams proto.InternalMessageInfo

type FormatStorageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FormatStorageResponse) Reset()         { *m = FormatStorageResponse{} }
func (m *FormatStorageResponse) String() string { return proto.CompactTextString(m) }
func (*FormatStorageResponse) ProtoMessage()    {}
func (*FormatStorageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_storage_242c73d0201305e3, []int{12}
}
func (m *FormatStorageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormatStorageResponse.Unmarshal(m, b)
}
func (m *FormatStorageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormatStorageResponse.Marshal(b, m, deterministic)
}
func (dst *FormatStorageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormatStorageResponse.Merge(dst, src)
}
func (m *FormatStorageResponse) XXX_Size() int {
	return xxx_messageInfo_FormatStorageResponse.Size(m)
}
func (m *FormatStorageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FormatStorageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FormatStorageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmptyParams)(nil), "mgmt.EmptyParams")
	proto.RegisterType((*FeatureName)(nil), "mgmt.FeatureName")
	proto.RegisterType((*Category)(nil), "mgmt.Category")
	proto.RegisterType((*Feature)(nil), "mgmt.Feature")
	proto.RegisterType((*NvmeNamespace)(nil), "mgmt.NvmeNamespace")
	proto.RegisterType((*NvmeController)(nil), "mgmt.NvmeController")
	proto.RegisterType((*UpdateNvmeParams)(nil), "mgmt.UpdateNvmeParams")
	proto.RegisterType((*FioConfigPath)(nil), "mgmt.FioConfigPath")
	proto.RegisterType((*BurnInNvmeParams)(nil), "mgmt.BurnInNvmeParams")
	proto.RegisterType((*BurnInNvmeReport)(nil), "mgmt.BurnInNvmeReport")
	proto.RegisterType((*ScmModule)(nil), "mgmt.ScmModule")
	proto.RegisterType((*FormatStorageParams)(nil), "mgmt.FormatStorageParams")
	proto.RegisterType((*FormatStorageResponse)(nil), "mgmt.FormatStorageResponse")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_storage_242c73d0201305e3) }

var fileDescriptor_storage_242c73d0201305e3 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0xa6, 0xbb, 0x69, 0x33, 0xab, 0x8d, 0x8a, 0x4b, 0x61, 0xc5, 0x01, 0x05, 0x23, 0xd1,
	0xa8, 0x87, 0x48, 0x94, 0x23, 0x37, 0x22, 0x22, 0x71, 0xa0, 0xaa, 0x5c, 0xfa, 0x03, 0x8c, 0xd7,
	0x49, 0x2c, 0xd6, 0x1f, 0xb2, 0x9d, 0xa0, 0xdc, 0xf8, 0x4b, 0x1c, 0xf8, 0x7f, 0x68, 0xbd, 0xce,
	0x66, 0x97, 0x13, 0xb7, 0x79, 0xe3, 0x79, 0xef, 0x8d, 0x67, 0x6c, 0x28, 0x9c, 0xd7, 0x96, 0x6e,
	0xf8, 0xc2, 0x58, 0xed, 0x35, 0x4a, 0xe5, 0x46, 0x7a, 0x5c, 0x40, 0xfe, 0x59, 0x1a, 0x7f, 0x78,
	0xa0, 0x96, 0x4a, 0x87, 0xdf, 0x40, 0xbe, 0xe2, 0xd4, 0xef, 0x2c, 0xbf, 0xa7, 0x92, 0x23, 0x04,
	0xa9, 0xa2, 0x92, 0x97, 0xc9, 0x2c, 0x99, 0x4f, 0x48, 0x88, 0xf1, 0x3b, 0xb8, 0x58, 0x52, 0xcf,
	0x37, 0xda, 0x1e, 0xd0, 0x2b, 0xb8, 0x60, 0x31, 0x8e, 0x35, 0x1d, 0xc6, 0xbf, 0x12, 0x38, 0x8f,
	0x5a, 0xe8, 0xf6, 0x9f, 0xba, 0xfc, 0x6e, 0xba, 0x68, 0xec, 0x17, 0x47, 0xa5, 0x13, 0x0f, 0xdd,
	0x40, 0xb6, 0x0e, 0xa6, 0xa3, 0x50, 0xf8, 0xac, 0x2d, 0xec, 0x75, 0x45, 0xda, 0x73, 0x34, 0x83,
	0xbc, 0xe2, 0x8e, 0x59, 0x61, 0xbc, 0xd0, 0xaa, 0x3c, 0x0b, 0xfe, 0xfd, 0x14, 0xfe, 0x08, 0xc5,
	0xfd, 0x5e, 0x06, 0x92, 0x33, 0x94, 0x71, 0x34, 0x85, 0x91, 0xa8, 0x82, 0x70, 0x46, 0x46, 0xa2,
	0x6a, 0xfb, 0x37, 0x94, 0x09, 0x7f, 0x08, 0xfc, 0x8c, 0x74, 0x18, 0xff, 0x4e, 0x60, 0xda, 0xb0,
	0x97, 0x5a, 0x79, 0xab, 0xeb, 0x9a, 0xdb, 0x48, 0x4f, 0x3a, 0xfa, 0x73, 0xc8, 0xa4, 0xae, 0x78,
	0x1d, 0x14, 0x27, 0xa4, 0x05, 0xe8, 0x05, 0x8c, 0x1d, 0xb7, 0x82, 0xd6, 0xb1, 0xa5, 0x88, 0x50,
	0x09, 0xe7, 0x86, 0x09, 0x5a, 0x55, 0xb6, 0x4c, 0xc3, 0xc1, 0x11, 0x36, 0x3a, 0xeb, 0x9f, 0x96,
	0xef, 0xcb, 0xac, 0xd5, 0x09, 0x00, 0xbd, 0x87, 0x89, 0x3a, 0x76, 0x5e, 0x8e, 0x67, 0x67, 0xf3,
	0xfc, 0xee, 0xaa, 0x1d, 0xc6, 0xe0, 0x52, 0xe4, 0x54, 0x85, 0xbf, 0xc1, 0xe5, 0x93, 0xa9, 0xa8,
	0xe7, 0x4d, 0x45, 0xbb, 0xd2, 0xbe, 0x6d, 0x32, 0xb4, 0x45, 0x90, 0x1a, 0xea, 0xb7, 0xb1, 0xfb,
	0x10, 0x37, 0x39, 0x57, 0x6b, 0x1f, 0xa7, 0x11, 0x62, 0xfc, 0x16, 0x8a, 0x95, 0xd0, 0x4b, 0xad,
	0xd6, 0x62, 0xf3, 0x10, 0x8b, 0x02, 0x31, 0x39, 0x11, 0xf1, 0x13, 0x5c, 0x7e, 0xda, 0x59, 0xf5,
	0x45, 0xfd, 0x97, 0xf5, 0x4d, 0xcf, 0xba, 0xbb, 0xd6, 0xc0, 0x24, 0xca, 0xde, 0xf6, 0x65, 0x09,
	0x37, 0xda, 0xfa, 0x66, 0xc0, 0x36, 0x44, 0x51, 0x35, 0x22, 0xfc, 0x27, 0x81, 0xc9, 0x23, 0x93,
	0x5f, 0x75, 0xb5, 0xab, 0x39, 0x7a, 0x0d, 0x60, 0xb6, 0x07, 0x27, 0x18, 0xad, 0xe3, 0xd2, 0x0a,
	0xd2, 0xcb, 0x34, 0xcd, 0xb1, 0x2d, 0x55, 0x2a, 0xae, 0xaf, 0x20, 0x47, 0xd8, 0x30, 0x63, 0x68,
	0xb4, 0x0b, 0x93, 0x28, 0x48, 0x2f, 0xd3, 0xbc, 0x1a, 0xc9, 0x25, 0xf3, 0xb6, 0x6e, 0x37, 0x59,
	0x90, 0x0e, 0x87, 0xe5, 0x6b, 0xf6, 0x83, 0xfb, 0xb0, 0xcb, 0x82, 0x44, 0x34, 0x78, 0x69, 0xe3,
	0x59, 0x32, 0x4f, 0x7b, 0x2f, 0xed, 0x1a, 0xae, 0x56, 0xda, 0x4a, 0xea, 0x1f, 0xdb, 0x0f, 0x1a,
	0xff, 0xe2, 0x4b, 0xb8, 0x1e, 0xa4, 0x09, 0x77, 0x46, 0x2b, 0xc7, 0xbf, 0x8f, 0xc3, 0x07, 0xfe,
	0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x83, 0x02, 0x6b, 0xd1, 0x03, 0x00, 0x00,
}
