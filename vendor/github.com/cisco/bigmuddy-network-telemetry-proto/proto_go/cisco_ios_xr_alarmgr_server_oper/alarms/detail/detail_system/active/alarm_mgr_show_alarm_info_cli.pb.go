// Code generated by protoc-gen-go.
// source: alarm_mgr_show_alarm_info_cli.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_system_active is a generated protocol buffer package.

It is generated from these files:
	alarm_mgr_show_alarm_info_cli.proto

It has these top-level messages:
	AlarmMgrShowAlarmInfoCli_KEYS
	AlarmMgrShowAlarmInfoCli
	AlarmOtn
	AlarmTca
	AlarmMgrShowAlarmData
*/
package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_system_active

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

// alarm mgr show alarm info for CLI
type AlarmMgrShowAlarmInfoCli_KEYS struct {
}

func (m *AlarmMgrShowAlarmInfoCli_KEYS) Reset()                    { *m = AlarmMgrShowAlarmInfoCli_KEYS{} }
func (m *AlarmMgrShowAlarmInfoCli_KEYS) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmInfoCli_KEYS) ProtoMessage()               {}
func (*AlarmMgrShowAlarmInfoCli_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AlarmMgrShowAlarmInfoCli struct {
	// Alarm List
	AlarmInfo []*AlarmMgrShowAlarmData `protobuf:"bytes,50,rep,name=alarm_info,json=alarmInfo" json:"alarm_info,omitempty"`
}

func (m *AlarmMgrShowAlarmInfoCli) Reset()                    { *m = AlarmMgrShowAlarmInfoCli{} }
func (m *AlarmMgrShowAlarmInfoCli) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmInfoCli) ProtoMessage()               {}
func (*AlarmMgrShowAlarmInfoCli) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AlarmMgrShowAlarmInfoCli) GetAlarmInfo() []*AlarmMgrShowAlarmData {
	if m != nil {
		return m.AlarmInfo
	}
	return nil
}

// Alarm transport attributes
type AlarmOtn struct {
	// Alarm direction
	Direction string `protobuf:"bytes,1,opt,name=direction" json:"direction,omitempty"`
	// Source of Alarm
	NotificationSource string `protobuf:"bytes,2,opt,name=notification_source,json=notificationSource" json:"notification_source,omitempty"`
}

func (m *AlarmOtn) Reset()                    { *m = AlarmOtn{} }
func (m *AlarmOtn) String() string            { return proto.CompactTextString(m) }
func (*AlarmOtn) ProtoMessage()               {}
func (*AlarmOtn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AlarmOtn) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *AlarmOtn) GetNotificationSource() string {
	if m != nil {
		return m.NotificationSource
	}
	return ""
}

// Alarm TCA Attributes
type AlarmTca struct {
	// Alarm Threshold
	ThresholdValue string `protobuf:"bytes,1,opt,name=threshold_value,json=thresholdValue" json:"threshold_value,omitempty"`
	// Alarm Threshold
	CurrentValue string `protobuf:"bytes,2,opt,name=current_value,json=currentValue" json:"current_value,omitempty"`
	// Timing Bucket
	BucketType string `protobuf:"bytes,3,opt,name=bucket_type,json=bucketType" json:"bucket_type,omitempty"`
}

func (m *AlarmTca) Reset()                    { *m = AlarmTca{} }
func (m *AlarmTca) String() string            { return proto.CompactTextString(m) }
func (*AlarmTca) ProtoMessage()               {}
func (*AlarmTca) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AlarmTca) GetThresholdValue() string {
	if m != nil {
		return m.ThresholdValue
	}
	return ""
}

func (m *AlarmTca) GetCurrentValue() string {
	if m != nil {
		return m.CurrentValue
	}
	return ""
}

func (m *AlarmTca) GetBucketType() string {
	if m != nil {
		return m.BucketType
	}
	return ""
}

// Alarm mgr show alarm data
type AlarmMgrShowAlarmData struct {
	// Alarm description
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// Alarm location
	Location string `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	// Alarm aid
	Aid string `protobuf:"bytes,3,opt,name=aid" json:"aid,omitempty"`
	// Alarm tag description
	Tag string `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	// Alarm module description
	Module string `protobuf:"bytes,5,opt,name=module" json:"module,omitempty"`
	// Alarm eid
	Eid string `protobuf:"bytes,6,opt,name=eid" json:"eid,omitempty"`
	// Reporting agent id
	ReportingAgentId uint32 `protobuf:"varint,7,opt,name=reporting_agent_id,json=reportingAgentId" json:"reporting_agent_id,omitempty"`
	// Pending async flag
	PendingSync bool `protobuf:"varint,8,opt,name=pending_sync,json=pendingSync" json:"pending_sync,omitempty"`
	// Alarm severity
	Severity string `protobuf:"bytes,9,opt,name=severity" json:"severity,omitempty"`
	// Alarm status
	Status string `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
	// Alarm group
	Group string `protobuf:"bytes,11,opt,name=group" json:"group,omitempty"`
	// Alarm set time
	SetTime string `protobuf:"bytes,12,opt,name=set_time,json=setTime" json:"set_time,omitempty"`
	// Alarm set time(timestamp format)
	SetTimestamp uint64 `protobuf:"varint,13,opt,name=set_timestamp,json=setTimestamp" json:"set_timestamp,omitempty"`
	// Alarm clear time
	ClearTime string `protobuf:"bytes,14,opt,name=clear_time,json=clearTime" json:"clear_time,omitempty"`
	// Alarm clear time(timestamp format)
	ClearTimestamp uint64 `protobuf:"varint,15,opt,name=clear_timestamp,json=clearTimestamp" json:"clear_timestamp,omitempty"`
	// Alarm service affecting
	ServiceAffecting string `protobuf:"bytes,16,opt,name=service_affecting,json=serviceAffecting" json:"service_affecting,omitempty"`
	// OTN feature specific alarm attributes
	Otn *AlarmOtn `protobuf:"bytes,17,opt,name=otn" json:"otn,omitempty"`
	// TCA feature specific alarm attributes
	Tca *AlarmTca `protobuf:"bytes,18,opt,name=tca" json:"tca,omitempty"`
	// alarm_event_type
	Type string `protobuf:"bytes,19,opt,name=type" json:"type,omitempty"`
	// Alarm interface name
	Interface string `protobuf:"bytes,20,opt,name=interface" json:"interface,omitempty"`
	// Alarm name
	AlarmName string `protobuf:"bytes,21,opt,name=alarm_name,json=alarmName" json:"alarm_name,omitempty"`
}

func (m *AlarmMgrShowAlarmData) Reset()                    { *m = AlarmMgrShowAlarmData{} }
func (m *AlarmMgrShowAlarmData) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmData) ProtoMessage()               {}
func (*AlarmMgrShowAlarmData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AlarmMgrShowAlarmData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetAid() string {
	if m != nil {
		return m.Aid
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetEid() string {
	if m != nil {
		return m.Eid
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetReportingAgentId() uint32 {
	if m != nil {
		return m.ReportingAgentId
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetPendingSync() bool {
	if m != nil {
		return m.PendingSync
	}
	return false
}

func (m *AlarmMgrShowAlarmData) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetSetTime() string {
	if m != nil {
		return m.SetTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetSetTimestamp() uint64 {
	if m != nil {
		return m.SetTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetClearTime() string {
	if m != nil {
		return m.ClearTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetClearTimestamp() uint64 {
	if m != nil {
		return m.ClearTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmData) GetServiceAffecting() string {
	if m != nil {
		return m.ServiceAffecting
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetOtn() *AlarmOtn {
	if m != nil {
		return m.Otn
	}
	return nil
}

func (m *AlarmMgrShowAlarmData) GetTca() *AlarmTca {
	if m != nil {
		return m.Tca
	}
	return nil
}

func (m *AlarmMgrShowAlarmData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *AlarmMgrShowAlarmData) GetAlarmName() string {
	if m != nil {
		return m.AlarmName
	}
	return ""
}

func init() {
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.active.alarm_mgr_show_alarm_info_cli_KEYS")
	proto.RegisterType((*AlarmMgrShowAlarmInfoCli)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.active.alarm_mgr_show_alarm_info_cli")
	proto.RegisterType((*AlarmOtn)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.active.alarm_otn")
	proto.RegisterType((*AlarmTca)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.active.alarm_tca")
	proto.RegisterType((*AlarmMgrShowAlarmData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_system.active.alarm_mgr_show_alarm_data")
}

func init() { proto.RegisterFile("alarm_mgr_show_alarm_info_cli.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6e, 0x13, 0x3b,
	0x14, 0x86, 0x95, 0x9b, 0x36, 0x6d, 0x4e, 0xd2, 0x36, 0x75, 0x7b, 0xaf, 0xdc, 0x2b, 0x2a, 0x42,
	0x8a, 0xd4, 0x48, 0xa0, 0x20, 0x95, 0x27, 0xa8, 0x10, 0x8b, 0x0a, 0xa9, 0x8b, 0x69, 0x85, 0x04,
	0x0b, 0x8c, 0xeb, 0x39, 0x99, 0x5a, 0xcc, 0xd8, 0x23, 0xfb, 0x4c, 0x20, 0x48, 0xbc, 0x01, 0xcf,
	0xc0, 0xb3, 0x22, 0xdb, 0x93, 0xb4, 0x0b, 0xe8, 0xaa, 0xab, 0xf8, 0x7c, 0xff, 0xaf, 0xff, 0x1c,
	0x3b, 0xf6, 0xc0, 0x89, 0x2c, 0xa5, 0xab, 0x44, 0x55, 0x38, 0xe1, 0x6f, 0xed, 0x57, 0x91, 0x4a,
	0x6d, 0xe6, 0x56, 0xa8, 0x52, 0xcf, 0x6a, 0x67, 0xc9, 0xb2, 0x37, 0x4a, 0x7b, 0x65, 0x85, 0xb6,
	0x5e, 0x7c, 0x73, 0xc9, 0x12, 0xfc, 0xe8, 0x16, 0xe8, 0x84, 0xad, 0xd1, 0xcd, 0x22, 0xf3, 0xb3,
	0x1c, 0x49, 0xea, 0xb2, 0xfd, 0x11, 0x7e, 0xe9, 0x09, 0xab, 0x99, 0x54, 0xa4, 0x17, 0x38, 0x79,
	0x0e, 0x93, 0x07, 0x7b, 0x89, 0x77, 0x6f, 0x3f, 0x5c, 0x4d, 0x7e, 0x75, 0xe0, 0xf8, 0x41, 0x1b,
	0xfb, 0x01, 0x70, 0x47, 0xf8, 0xd9, 0xb8, 0x3b, 0x1d, 0x9c, 0x7d, 0x9a, 0x3d, 0xc2, 0x84, 0xb3,
	0x3f, 0xf6, 0xcd, 0x25, 0xc9, 0xac, 0x1f, 0xd7, 0x17, 0x66, 0x6e, 0x27, 0x1f, 0x21, 0x15, 0xc2,
	0x92, 0x61, 0x4f, 0xa0, 0x9f, 0x6b, 0x87, 0x8a, 0xb4, 0x35, 0xbc, 0x33, 0xee, 0x4c, 0xfb, 0xd9,
	0x1d, 0x60, 0xaf, 0xe0, 0xc0, 0x58, 0xd2, 0x73, 0xad, 0x64, 0xa8, 0x85, 0xb7, 0x8d, 0x53, 0xc8,
	0xff, 0x89, 0x3e, 0x76, 0x5f, 0xba, 0x8a, 0xca, 0xe4, 0xfb, 0x2a, 0x9b, 0x94, 0x64, 0xa7, 0xb0,
	0x47, 0xb7, 0x0e, 0xfd, 0xad, 0x2d, 0x73, 0xb1, 0x90, 0x65, 0x83, 0x6d, 0x87, 0xdd, 0x35, 0x7e,
	0x1f, 0x28, 0x3b, 0x81, 0x1d, 0xd5, 0x38, 0x87, 0x86, 0x5a, 0x5b, 0x6a, 0x30, 0x6c, 0x61, 0x32,
	0x3d, 0x85, 0xc1, 0x4d, 0xa3, 0xbe, 0x20, 0x09, 0x5a, 0xd6, 0xc8, 0xbb, 0xd1, 0x02, 0x09, 0x5d,
	0x2f, 0x6b, 0x9c, 0xfc, 0xec, 0xc1, 0xd1, 0x5f, 0x0f, 0x80, 0x8d, 0x61, 0x90, 0xa3, 0x57, 0x4e,
	0xd7, 0xf7, 0xb6, 0x7a, 0x1f, 0xb1, 0xff, 0x61, 0xbb, 0xb4, 0x69, 0x37, 0xed, 0x00, 0xeb, 0x9a,
	0x8d, 0xa0, 0x2b, 0x75, 0xde, 0x36, 0x0d, 0xcb, 0x40, 0x48, 0x16, 0x7c, 0x23, 0x11, 0x92, 0x05,
	0xfb, 0x0f, 0x7a, 0x95, 0xcd, 0x9b, 0x12, 0xf9, 0x66, 0x84, 0x6d, 0x15, 0x9c, 0xa8, 0x73, 0xde,
	0x4b, 0x4e, 0xd4, 0x39, 0x7b, 0x09, 0xcc, 0x61, 0x6d, 0x1d, 0x69, 0x53, 0x08, 0x59, 0x84, 0x7d,
	0xeb, 0x9c, 0x6f, 0x8d, 0x3b, 0xd3, 0x9d, 0x6c, 0xb4, 0x56, 0xce, 0x83, 0x70, 0x91, 0xb3, 0x67,
	0x30, 0xac, 0xd1, 0xe4, 0xc1, 0xeb, 0x97, 0x46, 0xf1, 0xed, 0x71, 0x67, 0xba, 0x9d, 0x0d, 0x5a,
	0x76, 0xb5, 0x34, 0x2a, 0x8c, 0xee, 0x71, 0x81, 0x4e, 0xd3, 0x92, 0xf7, 0xd3, 0xe8, 0xab, 0x3a,
	0x8c, 0xe5, 0x49, 0x52, 0xe3, 0x39, 0xa4, 0xb1, 0x52, 0xc5, 0x0e, 0x61, 0xb3, 0x70, 0xb6, 0xa9,
	0xf9, 0x20, 0xe2, 0x54, 0xb0, 0xa3, 0x90, 0x44, 0x82, 0x74, 0x85, 0x7c, 0x18, 0x85, 0x2d, 0x8f,
	0x74, 0xad, 0xab, 0xf8, 0x2f, 0xad, 0x24, 0x4f, 0xb2, 0xaa, 0xf9, 0xce, 0xb8, 0x33, 0xdd, 0xc8,
	0x86, 0xad, 0x1e, 0x19, 0x3b, 0x06, 0x50, 0x25, 0x4a, 0x97, 0x12, 0x76, 0xd3, 0x85, 0x8a, 0x24,
	0x66, 0x9c, 0xc2, 0xde, 0x9d, 0x9c, 0x52, 0xf6, 0x62, 0xca, 0xee, 0xda, 0x93, 0x72, 0x5e, 0xc0,
	0x7e, 0xb8, 0xfb, 0x5a, 0xa1, 0x90, 0xf3, 0x79, 0xb8, 0x8e, 0xa6, 0xe0, 0xa3, 0x18, 0x37, 0x6a,
	0x85, 0xf3, 0x15, 0x67, 0x9f, 0xa1, 0x6b, 0xc9, 0xf0, 0xfd, 0x71, 0x67, 0x3a, 0x38, 0xbb, 0x7c,
	0xc4, 0x97, 0x64, 0xc9, 0x64, 0x21, 0x3a, 0x74, 0x20, 0x25, 0x39, 0x7b, 0xf4, 0x0e, 0xa4, 0x64,
	0x16, 0xa2, 0x19, 0x83, 0x8d, 0x78, 0xaf, 0x0f, 0xe2, 0x1e, 0xe3, 0x3a, 0x3c, 0x4e, 0x6d, 0x08,
	0xdd, 0x5c, 0x2a, 0xe4, 0x87, 0xe9, 0x2c, 0xd7, 0x20, 0x1c, 0x75, 0xca, 0x30, 0xb2, 0x42, 0xfe,
	0x6f, 0x92, 0x23, 0xb9, 0x94, 0x15, 0xde, 0xf4, 0xe2, 0x97, 0xef, 0xf5, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0xb6, 0x45, 0xac, 0x20, 0x05, 0x00, 0x00,
}
