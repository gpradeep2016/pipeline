// Code generated by protoc-gen-go.
// source: bgp_nbr_brief_bag.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_sessions_session is a generated protocol buffer package.

It is generated from these files:
	bgp_nbr_brief_bag.proto

It has these top-level messages:
	BgpNbrBriefBag_KEYS
	BgpNbrBriefBag
	IPV4TunnelAddressType
	IPV4MDTAddressType
	RTConstraintAddressType
	IPV6AddressType
	BgpL2VpnAddrT
	L2VPNEVPNAddressType
	BgpL2VpnMspwAddrT
	IPV6MVPNAddressType
	IPV4MVPNAddressType
	LS_LSAddressType
	IPv4FlowspecAddressType
	IPv6FlowspecAddressType
	BgpAddrtype
*/
package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_sessions_session

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

// BGP Neighbor brief Information
type BgpNbrBriefBag_KEYS struct {
	InstanceName    string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	NeighborAddress string `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress" json:"neighbor_address,omitempty"`
}

func (m *BgpNbrBriefBag_KEYS) Reset()                    { *m = BgpNbrBriefBag_KEYS{} }
func (m *BgpNbrBriefBag_KEYS) String() string            { return proto.CompactTextString(m) }
func (*BgpNbrBriefBag_KEYS) ProtoMessage()               {}
func (*BgpNbrBriefBag_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BgpNbrBriefBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpNbrBriefBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type BgpNbrBriefBag struct {
	// Speaker this neighbor is allocated to
	SpeakerId uint32 `protobuf:"varint,50,opt,name=speaker_id,json=speakerId" json:"speaker_id,omitempty"`
	// Description
	Description string `protobuf:"bytes,51,opt,name=description" json:"description,omitempty"`
	// Local AS number
	LocalAs uint32 `protobuf:"varint,52,opt,name=local_as,json=localAs" json:"local_as,omitempty"`
	// Remote AS number
	RemoteAs uint32 `protobuf:"varint,53,opt,name=remote_as,json=remoteAs" json:"remote_as,omitempty"`
	// No. of msgs on receive queue
	MessagesQueuedIn uint32 `protobuf:"varint,54,opt,name=messages_queued_in,json=messagesQueuedIn" json:"messages_queued_in,omitempty"`
	// No. of messages on send queue
	MessagesQueuedOut uint32 `protobuf:"varint,55,opt,name=messages_queued_out,json=messagesQueuedOut" json:"messages_queued_out,omitempty"`
	// State of connection
	ConnectionState string `protobuf:"bytes,56,opt,name=connection_state,json=connectionState" json:"connection_state,omitempty"`
	// Local address for the connection
	ConnectionLocalAddress *BgpAddrtype `protobuf:"bytes,57,opt,name=connection_local_address,json=connectionLocalAddress" json:"connection_local_address,omitempty"`
	// Local address configured for the neighbor connection
	IsLocalAddressConfigured bool `protobuf:"varint,58,opt,name=is_local_address_configured,json=isLocalAddressConfigured" json:"is_local_address_configured,omitempty"`
	// Remote address for the connection
	ConnectionRemoteAddress *BgpAddrtype `protobuf:"bytes,59,opt,name=connection_remote_address,json=connectionRemoteAddress" json:"connection_remote_address,omitempty"`
	// Name of the VRF
	VrfName string `protobuf:"bytes,60,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	// Neighbor supports NSR
	NsrEnabled bool `protobuf:"varint,61,opt,name=nsr_enabled,json=nsrEnabled" json:"nsr_enabled,omitempty"`
	// NSR state
	NsrState string `protobuf:"bytes,62,opt,name=nsr_state,json=nsrState" json:"nsr_state,omitempty"`
	// Nbr has postits pending
	PostitPending bool `protobuf:"varint,63,opt,name=postit_pending,json=postitPending" json:"postit_pending,omitempty"`
}

func (m *BgpNbrBriefBag) Reset()                    { *m = BgpNbrBriefBag{} }
func (m *BgpNbrBriefBag) String() string            { return proto.CompactTextString(m) }
func (*BgpNbrBriefBag) ProtoMessage()               {}
func (*BgpNbrBriefBag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BgpNbrBriefBag) GetSpeakerId() uint32 {
	if m != nil {
		return m.SpeakerId
	}
	return 0
}

func (m *BgpNbrBriefBag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BgpNbrBriefBag) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *BgpNbrBriefBag) GetRemoteAs() uint32 {
	if m != nil {
		return m.RemoteAs
	}
	return 0
}

func (m *BgpNbrBriefBag) GetMessagesQueuedIn() uint32 {
	if m != nil {
		return m.MessagesQueuedIn
	}
	return 0
}

func (m *BgpNbrBriefBag) GetMessagesQueuedOut() uint32 {
	if m != nil {
		return m.MessagesQueuedOut
	}
	return 0
}

func (m *BgpNbrBriefBag) GetConnectionState() string {
	if m != nil {
		return m.ConnectionState
	}
	return ""
}

func (m *BgpNbrBriefBag) GetConnectionLocalAddress() *BgpAddrtype {
	if m != nil {
		return m.ConnectionLocalAddress
	}
	return nil
}

func (m *BgpNbrBriefBag) GetIsLocalAddressConfigured() bool {
	if m != nil {
		return m.IsLocalAddressConfigured
	}
	return false
}

func (m *BgpNbrBriefBag) GetConnectionRemoteAddress() *BgpAddrtype {
	if m != nil {
		return m.ConnectionRemoteAddress
	}
	return nil
}

func (m *BgpNbrBriefBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpNbrBriefBag) GetNsrEnabled() bool {
	if m != nil {
		return m.NsrEnabled
	}
	return false
}

func (m *BgpNbrBriefBag) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

func (m *BgpNbrBriefBag) GetPostitPending() bool {
	if m != nil {
		return m.PostitPending
	}
	return false
}

// IPV4Tunnel Address type
type IPV4TunnelAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4TunnelAddressType) Reset()                    { *m = IPV4TunnelAddressType{} }
func (m *IPV4TunnelAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4TunnelAddressType) ProtoMessage()               {}
func (*IPV4TunnelAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IPV4TunnelAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4MDT Address type
type IPV4MDTAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4MDTAddressType) Reset()                    { *m = IPV4MDTAddressType{} }
func (m *IPV4MDTAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4MDTAddressType) ProtoMessage()               {}
func (*IPV4MDTAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IPV4MDTAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4 RTConstraint Address type
type RTConstraintAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *RTConstraintAddressType) Reset()                    { *m = RTConstraintAddressType{} }
func (m *RTConstraintAddressType) String() string            { return proto.CompactTextString(m) }
func (*RTConstraintAddressType) ProtoMessage()               {}
func (*RTConstraintAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RTConstraintAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV6 Address type
type IPV6AddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV6AddressType) Reset()                    { *m = IPV6AddressType{} }
func (m *IPV6AddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV6AddressType) ProtoMessage()               {}
func (*IPV6AddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IPV6AddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpL2VpnAddrT struct {
	L2VpnAddress []byte `protobuf:"bytes,1,opt,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
}

func (m *BgpL2VpnAddrT) Reset()                    { *m = BgpL2VpnAddrT{} }
func (m *BgpL2VpnAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpL2VpnAddrT) ProtoMessage()               {}
func (*BgpL2VpnAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BgpL2VpnAddrT) GetL2VpnAddress() []byte {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

// L2VPN EVPN Address type
type L2VPNEVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *L2VPNEVPNAddressType) Reset()                    { *m = L2VPNEVPNAddressType{} }
func (m *L2VPNEVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*L2VPNEVPNAddressType) ProtoMessage()               {}
func (*L2VPNEVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *L2VPNEVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpL2VpnMspwAddrT struct {
	L2VpnAddress []byte `protobuf:"bytes,1,opt,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
}

func (m *BgpL2VpnMspwAddrT) Reset()                    { *m = BgpL2VpnMspwAddrT{} }
func (m *BgpL2VpnMspwAddrT) String() string            { return proto.CompactTextString(m) }
func (*BgpL2VpnMspwAddrT) ProtoMessage()               {}
func (*BgpL2VpnMspwAddrT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BgpL2VpnMspwAddrT) GetL2VpnAddress() []byte {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

// IPV6 MVPN Address type
type IPV6MVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV6MVPNAddressType) Reset()                    { *m = IPV6MVPNAddressType{} }
func (m *IPV6MVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV6MVPNAddressType) ProtoMessage()               {}
func (*IPV6MVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *IPV6MVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPV4 MVPN Address type
type IPV4MVPNAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPV4MVPNAddressType) Reset()                    { *m = IPV4MVPNAddressType{} }
func (m *IPV4MVPNAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPV4MVPNAddressType) ProtoMessage()               {}
func (*IPV4MVPNAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *IPV4MVPNAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// LINKSTATE LINKSTATE Address type
type LS_LSAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *LS_LSAddressType) Reset()                    { *m = LS_LSAddressType{} }
func (m *LS_LSAddressType) String() string            { return proto.CompactTextString(m) }
func (*LS_LSAddressType) ProtoMessage()               {}
func (*LS_LSAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *LS_LSAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPv4 Flowspec Address type
type IPv4FlowspecAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPv4FlowspecAddressType) Reset()                    { *m = IPv4FlowspecAddressType{} }
func (m *IPv4FlowspecAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPv4FlowspecAddressType) ProtoMessage()               {}
func (*IPv4FlowspecAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IPv4FlowspecAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// IPv6 Flowspec Address type
type IPv6FlowspecAddressType struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IPv6FlowspecAddressType) Reset()                    { *m = IPv6FlowspecAddressType{} }
func (m *IPv6FlowspecAddressType) String() string            { return proto.CompactTextString(m) }
func (*IPv6FlowspecAddressType) ProtoMessage()               {}
func (*IPv6FlowspecAddressType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *IPv6FlowspecAddressType) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type BgpAddrtype struct {
	Afi string `protobuf:"bytes,1,opt,name=afi" json:"afi,omitempty"`
	// IPv4 Addr
	Ipv4Address string `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	// IPv4 Mcast Addr
	Ipv4McastAddress string `protobuf:"bytes,3,opt,name=ipv4_mcast_address,json=ipv4McastAddress" json:"ipv4_mcast_address,omitempty"`
	// IPv4 Label Addr
	Ipv4LabelAddress string `protobuf:"bytes,4,opt,name=ipv4_label_address,json=ipv4LabelAddress" json:"ipv4_label_address,omitempty"`
	// IPv4 Tunnel
	Ipv4TunnelAddress *IPV4TunnelAddressType `protobuf:"bytes,5,opt,name=ipv4_tunnel_address,json=ipv4TunnelAddress" json:"ipv4_tunnel_address,omitempty"`
	// IPv4 MDT Addr
	Ipv4MdtAddress *IPV4MDTAddressType `protobuf:"bytes,6,opt,name=ipv4_mdt_address,json=ipv4MdtAddress" json:"ipv4_mdt_address,omitempty"`
	// IPv4 VPN Addr
	Ipv4VpnAddress string `protobuf:"bytes,7,opt,name=ipv4_vpn_address,json=ipv4VpnAddress" json:"ipv4_vpn_address,omitempty"`
	// IPv4 VPN Mcast Addr
	Ipv4VpnaMcastddress string `protobuf:"bytes,8,opt,name=ipv4_vpna_mcastddress,json=ipv4VpnaMcastddress" json:"ipv4_vpna_mcastddress,omitempty"`
	// IPV6 Addr
	Ipv6Address *IPV6AddressType `protobuf:"bytes,9,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
	// IPV6 Mcast Addr
	Ipv6McastAddress *IPV6AddressType `protobuf:"bytes,10,opt,name=ipv6_mcast_address,json=ipv6McastAddress" json:"ipv6_mcast_address,omitempty"`
	// IPv6 Label Addr
	Ipv6LabelAddress *IPV6AddressType `protobuf:"bytes,11,opt,name=ipv6_label_address,json=ipv6LabelAddress" json:"ipv6_label_address,omitempty"`
	// IPv6 VPN Addr
	Ipv6VpnAddress *IPV6AddressType `protobuf:"bytes,12,opt,name=ipv6_vpn_address,json=ipv6VpnAddress" json:"ipv6_vpn_address,omitempty"`
	// IPv6 VPN Mcast Addr
	Ipv6VpnMcastAddress *IPV6AddressType `protobuf:"bytes,13,opt,name=ipv6_vpn_mcast_address,json=ipv6VpnMcastAddress" json:"ipv6_vpn_mcast_address,omitempty"`
	// L2VPN VPLS Addr
	L2VpnvplsAddress *BgpL2VpnAddrT `protobuf:"bytes,14,opt,name=l2_vpnvpls_address,json=l2VpnvplsAddress" json:"l2_vpnvpls_address,omitempty"`
	// RT Constrt Addr
	RtConstraintAddress *RTConstraintAddressType `protobuf:"bytes,15,opt,name=rt_constraint_address,json=rtConstraintAddress" json:"rt_constraint_address,omitempty"`
	// MVPN addr
	Ipv6MvpnAddress *IPV6MVPNAddressType `protobuf:"bytes,16,opt,name=ipv6_mvpn_address,json=ipv6MvpnAddress" json:"ipv6_mvpn_address,omitempty"`
	// MVPN4 addr
	Ipv4MvpnAddress *IPV4MVPNAddressType `protobuf:"bytes,17,opt,name=ipv4_mvpn_address,json=ipv4MvpnAddress" json:"ipv4_mvpn_address,omitempty"`
	// L2VPN EVPN Addr
	L2VpnEvpnAddress *L2VPNEVPNAddressType `protobuf:"bytes,18,opt,name=l2_vpn_evpn_address,json=l2VpnEvpnAddress" json:"l2_vpn_evpn_address,omitempty"`
	// LINKSTATE LINKSTATE Addr
	LsLsAddress *LS_LSAddressType `protobuf:"bytes,19,opt,name=ls_ls_address,json=lsLsAddress" json:"ls_ls_address,omitempty"`
	// L2VPN MSPW Addr
	L2VpnMspwAddress *BgpL2VpnMspwAddrT `protobuf:"bytes,20,opt,name=l2_vpn_mspw_address,json=l2VpnMspwAddress" json:"l2_vpn_mspw_address,omitempty"`
	// IPV4 Flowspec Addr
	Ipv4FlowspecAddress *IPv4FlowspecAddressType `protobuf:"bytes,21,opt,name=ipv4_flowspec_address,json=ipv4FlowspecAddress" json:"ipv4_flowspec_address,omitempty"`
	// IPV6 Flowspec Addr
	Ipv6FlowspecAddress *IPv6FlowspecAddressType `protobuf:"bytes,22,opt,name=ipv6_flowspec_address,json=ipv6FlowspecAddress" json:"ipv6_flowspec_address,omitempty"`
	// IPV4 VPN Flowspec Addr
	Ipv4VpnFlowspecAddress *IPv4FlowspecAddressType `protobuf:"bytes,23,opt,name=ipv4_vpn_flowspec_address,json=ipv4VpnFlowspecAddress" json:"ipv4_vpn_flowspec_address,omitempty"`
	// IPV6 VPN Flowspec Addr
	Ipv6VpnFlowspecAddress *IPv6FlowspecAddressType `protobuf:"bytes,24,opt,name=ipv6_vpn_flowspec_address,json=ipv6VpnFlowspecAddress" json:"ipv6_vpn_flowspec_address,omitempty"`
}

func (m *BgpAddrtype) Reset()                    { *m = BgpAddrtype{} }
func (m *BgpAddrtype) String() string            { return proto.CompactTextString(m) }
func (*BgpAddrtype) ProtoMessage()               {}
func (*BgpAddrtype) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *BgpAddrtype) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4McastAddress() string {
	if m != nil {
		return m.Ipv4McastAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4LabelAddress() string {
	if m != nil {
		return m.Ipv4LabelAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4TunnelAddress() *IPV4TunnelAddressType {
	if m != nil {
		return m.Ipv4TunnelAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4MdtAddress() *IPV4MDTAddressType {
	if m != nil {
		return m.Ipv4MdtAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4VpnAddress() string {
	if m != nil {
		return m.Ipv4VpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4VpnaMcastddress() string {
	if m != nil {
		return m.Ipv4VpnaMcastddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6Address() *IPV6AddressType {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6McastAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6McastAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6LabelAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6LabelAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6VpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnMcastAddress() *IPV6AddressType {
	if m != nil {
		return m.Ipv6VpnMcastAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnvplsAddress() *BgpL2VpnAddrT {
	if m != nil {
		return m.L2VpnvplsAddress
	}
	return nil
}

func (m *BgpAddrtype) GetRtConstraintAddress() *RTConstraintAddressType {
	if m != nil {
		return m.RtConstraintAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6MvpnAddress() *IPV6MVPNAddressType {
	if m != nil {
		return m.Ipv6MvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4MvpnAddress() *IPV4MVPNAddressType {
	if m != nil {
		return m.Ipv4MvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnEvpnAddress() *L2VPNEVPNAddressType {
	if m != nil {
		return m.L2VpnEvpnAddress
	}
	return nil
}

func (m *BgpAddrtype) GetLsLsAddress() *LS_LSAddressType {
	if m != nil {
		return m.LsLsAddress
	}
	return nil
}

func (m *BgpAddrtype) GetL2VpnMspwAddress() *BgpL2VpnMspwAddrT {
	if m != nil {
		return m.L2VpnMspwAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4FlowspecAddress() *IPv4FlowspecAddressType {
	if m != nil {
		return m.Ipv4FlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6FlowspecAddress() *IPv6FlowspecAddressType {
	if m != nil {
		return m.Ipv6FlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4VpnFlowspecAddress() *IPv4FlowspecAddressType {
	if m != nil {
		return m.Ipv4VpnFlowspecAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6VpnFlowspecAddress() *IPv6FlowspecAddressType {
	if m != nil {
		return m.Ipv6VpnFlowspecAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpNbrBriefBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_nbr_brief_bag_KEYS")
	proto.RegisterType((*BgpNbrBriefBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_nbr_brief_bag")
	proto.RegisterType((*IPV4TunnelAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.IPV4TunnelAddressType")
	proto.RegisterType((*IPV4MDTAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.IPV4MDTAddressType")
	proto.RegisterType((*RTConstraintAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.RTConstraintAddressType")
	proto.RegisterType((*IPV6AddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.IPV6AddressType")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_l2vpn_addr_t")
	proto.RegisterType((*L2VPNEVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.L2VPNEVPNAddressType")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*IPV6MVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.IPV6MVPNAddressType")
	proto.RegisterType((*IPV4MVPNAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.IPV4MVPNAddressType")
	proto.RegisterType((*LS_LSAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.LS_LSAddressType")
	proto.RegisterType((*IPv4FlowspecAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.IPv4FlowspecAddressType")
	proto.RegisterType((*IPv6FlowspecAddressType)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.IPv6FlowspecAddressType")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_addrtype")
}

func init() { proto.RegisterFile("bgp_nbr_brief_bag.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x98, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x65, 0x96, 0x6d, 0xd3, 0x93, 0xb4, 0x4d, 0x27, 0x6d, 0xea, 0x6a, 0x85, 0x28, 0x45,
	0x88, 0x00, 0x25, 0x48, 0xd9, 0xe0, 0xe5, 0xb2, 0x0b, 0xaa, 0x96, 0x22, 0x55, 0xa4, 0x25, 0xb8,
	0x55, 0x24, 0x5e, 0x18, 0x4d, 0xec, 0x49, 0xd6, 0xc2, 0x19, 0x9b, 0x99, 0x49, 0xca, 0x7e, 0x0d,
	0xc4, 0xcb, 0x4a, 0x5c, 0x1e, 0xb8, 0x68, 0x25, 0xde, 0xf8, 0x84, 0x68, 0x66, 0x6c, 0xc7, 0x4e,
	0x5b, 0x35, 0x3c, 0x34, 0x79, 0x73, 0xce, 0xf9, 0x8f, 0xe7, 0x37, 0xe7, 0x32, 0x3e, 0x0a, 0xec,
	0xf6, 0x87, 0x31, 0x66, 0x7d, 0x8e, 0xfb, 0x3c, 0xa0, 0x03, 0xdc, 0x27, 0xc3, 0x66, 0xcc, 0x23,
	0x19, 0xa1, 0xef, 0xbc, 0x40, 0x78, 0x11, 0x0e, 0x22, 0x81, 0x7f, 0xe4, 0x38, 0x88, 0x27, 0x6d,
	0xac, 0xa4, 0x51, 0x4c, 0x79, 0xb3, 0x3f, 0x8c, 0x9b, 0x01, 0x13, 0x92, 0x30, 0x8f, 0x8a, 0xec,
	0x29, 0x7b, 0xc0, 0xc4, 0x93, 0xc1, 0x84, 0x36, 0x7d, 0x3a, 0x20, 0xe3, 0x50, 0xe2, 0x09, 0x1f,
	0x34, 0x05, 0x15, 0x22, 0x88, 0x98, 0x48, 0x1f, 0x0e, 0x9e, 0x41, 0xfd, 0xca, 0xd6, 0xf8, 0xab,
	0xe3, 0x6f, 0xcf, 0xd1, 0x9b, 0xb0, 0x9e, 0xbd, 0x89, 0x91, 0x11, 0xb5, 0xad, 0x7d, 0xab, 0xb1,
	0xe6, 0x56, 0x52, 0xe3, 0x19, 0x19, 0x51, 0xf4, 0x0e, 0x54, 0x19, 0x0d, 0x86, 0xcf, 0xfa, 0x11,
	0xc7, 0xc4, 0xf7, 0x39, 0x15, 0xc2, 0x7e, 0x45, 0xeb, 0x36, 0x53, 0xfb, 0x91, 0x31, 0x1f, 0xbc,
	0x5c, 0x81, 0xad, 0x2b, 0x5b, 0xa1, 0xd7, 0x00, 0x44, 0x4c, 0xc9, 0xf7, 0x94, 0xe3, 0xc0, 0xb7,
	0x5b, 0xfb, 0x56, 0x63, 0xdd, 0x5d, 0x4b, 0x2c, 0x27, 0x3e, 0xda, 0x87, 0xb2, 0x4f, 0x85, 0xc7,
	0x83, 0x58, 0x06, 0x11, 0xb3, 0x1f, 0xea, 0x57, 0xe7, 0x4d, 0x68, 0x0f, 0x4a, 0x61, 0xe4, 0x91,
	0x10, 0x13, 0x61, 0xb7, 0xf5, 0xf2, 0x55, 0xfd, 0xfb, 0x48, 0xa0, 0x07, 0xb0, 0xc6, 0xe9, 0x28,
	0x92, 0x54, 0xf9, 0x3e, 0xd4, 0xbe, 0x92, 0x31, 0x1c, 0x09, 0x74, 0x08, 0x68, 0x44, 0x85, 0x20,
	0x43, 0x2a, 0xf0, 0x0f, 0x63, 0x3a, 0xa6, 0x3e, 0x0e, 0x98, 0xed, 0x68, 0x55, 0x35, 0xf5, 0x7c,
	0xa3, 0x1d, 0x27, 0x0c, 0x35, 0xa1, 0x36, 0xab, 0x8e, 0xc6, 0xd2, 0x7e, 0xa4, 0xe5, 0x5b, 0x45,
	0xf9, 0xd7, 0x63, 0xa9, 0xe2, 0xe2, 0x45, 0x8c, 0x51, 0x4f, 0x31, 0x62, 0x21, 0x89, 0xa4, 0xf6,
	0x47, 0x26, 0x2e, 0x53, 0xfb, 0xb9, 0x32, 0xa3, 0xbf, 0x2d, 0xb0, 0x73, 0xda, 0xe4, 0x30, 0x49,
	0x2c, 0x3f, 0xde, 0xb7, 0x1a, 0xe5, 0x56, 0xd8, 0xbc, 0xdb, 0x2a, 0x50, 0xef, 0xd0, 0x5b, 0xca,
	0xe7, 0x31, 0x75, 0xeb, 0x53, 0x9a, 0x8e, 0x8e, 0xa4, 0x61, 0x41, 0x4f, 0xe0, 0x41, 0x20, 0x8a,
	0x7c, 0xd8, 0x8b, 0xd8, 0x20, 0x18, 0x8e, 0x39, 0xf5, 0xed, 0x4f, 0xf6, 0xad, 0x46, 0xc9, 0xb5,
	0x03, 0x91, 0x5f, 0xf4, 0x34, 0xf3, 0xa3, 0x97, 0x16, 0xec, 0xe5, 0xce, 0x99, 0x66, 0x26, 0x39,
	0xe8, 0xa7, 0x4b, 0x38, 0xe8, 0xee, 0x14, 0xc7, 0x35, 0x65, 0x91, 0x9c, 0x74, 0x0f, 0x4a, 0x13,
	0x3e, 0x30, 0x55, 0xff, 0x58, 0x67, 0x6d, 0x75, 0xc2, 0x07, 0xba, 0xe0, 0x5f, 0x87, 0x32, 0x13,
	0x1c, 0x53, 0x46, 0xfa, 0x21, 0xf5, 0xed, 0x27, 0xfa, 0xd0, 0xc0, 0x04, 0x3f, 0x36, 0x16, 0x55,
	0x74, 0x4a, 0x60, 0x52, 0xfe, 0x99, 0x5e, 0x5c, 0x62, 0x82, 0x9b, 0x5c, 0xbf, 0x05, 0x1b, 0x71,
	0x24, 0x64, 0x20, 0x71, 0x4c, 0x99, 0x1f, 0xb0, 0xa1, 0xfd, 0xb9, 0x7e, 0xc1, 0xba, 0xb1, 0x76,
	0x8d, 0xf1, 0xe0, 0x7d, 0xd8, 0x39, 0xe9, 0xf6, 0xda, 0x17, 0x63, 0xc6, 0x68, 0x1a, 0xc9, 0x8b,
	0xe7, 0x31, 0x45, 0xdb, 0x70, 0x7f, 0x42, 0xc2, 0x71, 0xda, 0x8b, 0xe6, 0xc7, 0xc1, 0xbb, 0x80,
	0x94, 0xfc, 0xf4, 0x8b, 0x8b, 0xdb, 0xb5, 0x1f, 0xc0, 0xae, 0x7b, 0xf1, 0x34, 0x62, 0x42, 0x72,
	0x12, 0x30, 0x79, 0xfb, 0x82, 0xb7, 0x61, 0xf3, 0xa4, 0xdb, 0x73, 0x6e, 0x17, 0x3e, 0x82, 0xaa,
	0x8a, 0x6e, 0xd8, 0x9a, 0xc4, 0x4c, 0xc7, 0x18, 0x4b, 0x75, 0x87, 0x4c, 0x7f, 0xab, 0x34, 0xab,
	0x15, 0x15, 0xb7, 0xa2, 0x8d, 0xe9, 0xc5, 0x70, 0x08, 0xdb, 0x9d, 0x56, 0xaf, 0x7b, 0x76, 0xdc,
	0xeb, 0x9e, 0xdd, 0xbe, 0xcd, 0x63, 0xd8, 0x99, 0x6e, 0x33, 0x12, 0xf1, 0xe5, 0xff, 0xda, 0xeb,
	0x3d, 0xa8, 0xa9, 0xd3, 0x9c, 0xce, 0xb5, 0x95, 0x11, 0xb7, 0xe7, 0x13, 0x37, 0xa0, 0xda, 0x39,
	0xc7, 0x9d, 0xf3, 0xb9, 0x52, 0x70, 0xd2, 0x9d, 0xb4, 0xbf, 0x0c, 0xa3, 0x4b, 0x11, 0x53, 0x6f,
	0xde, 0x05, 0xce, 0xfc, 0x0b, 0x5e, 0xec, 0x41, 0x25, 0x5f, 0xe9, 0xa8, 0x0a, 0xf7, 0xc8, 0x20,
	0x48, 0x44, 0xea, 0x11, 0xbd, 0x01, 0x15, 0xdd, 0x5d, 0xc5, 0x4b, 0xbb, 0xac, 0x6c, 0x69, 0x17,
	0x1c, 0x02, 0xd2, 0x92, 0x91, 0x47, 0x84, 0xcc, 0x84, 0xf7, 0xb4, 0xb0, 0xaa, 0x3c, 0xa7, 0xca,
	0x31, 0xab, 0x0e, 0x49, 0x9f, 0x4e, 0xef, 0xaf, 0x57, 0xa7, 0xea, 0x8e, 0x72, 0xa4, 0xea, 0xbf,
	0x2c, 0xa8, 0x69, 0xb9, 0xd4, 0x35, 0x9e, 0xe9, 0xef, 0xeb, 0x6b, 0x60, 0x7c, 0xd7, 0xd7, 0xc0,
	0xb5, 0xdd, 0xe5, 0x6e, 0xa9, 0x8d, 0x0a, 0x66, 0xf4, 0x8b, 0x05, 0x55, 0x13, 0x04, 0x7f, 0x1a,
	0x82, 0x15, 0x0d, 0xc9, 0x17, 0x01, 0x59, 0xec, 0x69, 0x77, 0x43, 0x87, 0xdd, 0xcf, 0x82, 0xde,
	0x48, 0xe8, 0xf2, 0x65, 0xbf, 0xaa, 0x43, 0xae, 0x95, 0xbd, 0xac, 0xf0, 0x51, 0x0b, 0x76, 0x52,
	0x25, 0x31, 0x19, 0x4d, 0xe4, 0x25, 0x2d, 0xaf, 0x25, 0x72, 0x72, 0x3a, 0x75, 0xa1, 0x9f, 0x2c,
	0x5d, 0x24, 0x4e, 0xf6, 0xea, 0x35, 0x7d, 0xf0, 0x68, 0x01, 0x07, 0xcf, 0xdf, 0x37, 0xba, 0x2a,
	0x53, 0x03, 0xfa, 0xd5, 0xd2, 0x85, 0xe6, 0xcc, 0x94, 0x25, 0x2c, 0x07, 0x4d, 0x85, 0xdf, 0x29,
	0xf4, 0x41, 0xc6, 0x57, 0x6c, 0x84, 0xf2, 0x12, 0xf9, 0x0a, 0x9d, 0xf7, 0xc2, 0x54, 0xb4, 0x53,
	0xa8, 0x99, 0xca, 0x72, 0xe8, 0x54, 0x91, 0x3a, 0xb9, 0x22, 0xfd, 0xd3, 0x82, 0x7a, 0xc6, 0x56,
	0xcc, 0xef, 0xfa, 0x72, 0x08, 0x6b, 0x09, 0x61, 0x21, 0xc5, 0xbf, 0x59, 0x80, 0xc2, 0x96, 0x82,
	0x9c, 0xc4, 0xa1, 0xc8, 0x10, 0x37, 0x34, 0x62, 0xbc, 0x88, 0x11, 0x26, 0xff, 0x91, 0x75, 0xab,
	0x61, 0xab, 0x67, 0x50, 0x52, 0xc0, 0x7f, 0x2c, 0xd8, 0xe1, 0x52, 0x0d, 0x67, 0xc9, 0x67, 0x3e,
	0x63, 0xdc, 0xd4, 0x8c, 0x97, 0x77, 0xcd, 0x78, 0xc3, 0x88, 0xe1, 0xd6, 0xb8, 0xbc, 0xe2, 0x40,
	0xbf, 0x5b, 0xb0, 0x65, 0x3a, 0x3a, 0x5f, 0x92, 0x55, 0x4d, 0x2a, 0x16, 0x91, 0xf0, 0x99, 0x0f,
	0xbc, 0xbb, 0xa9, 0x9b, 0x7a, 0x3a, 0x35, 0xa4, 0x84, 0xed, 0x22, 0xe1, 0xd6, 0xc2, 0x08, 0xdb,
	0xd7, 0x11, 0xb6, 0xf3, 0x84, 0x7f, 0x58, 0x50, 0x33, 0x25, 0x89, 0x69, 0x9e, 0x11, 0x69, 0x46,
	0x79, 0xd7, 0x8c, 0xd7, 0xcd, 0x6f, 0x49, 0x5d, 0x1e, 0xe7, 0x28, 0x7f, 0xb6, 0x60, 0x3d, 0x14,
	0x38, 0xd7, 0x33, 0xb5, 0xc5, 0xf4, 0xcc, 0xec, 0x64, 0xe6, 0x96, 0x43, 0xd1, 0x11, 0xb9, 0x6b,
	0x27, 0x0d, 0x5e, 0x36, 0x50, 0x2a, 0xb8, 0xed, 0xc5, 0x0c, 0x23, 0xd7, 0x8e, 0xb3, 0x49, 0xf4,
	0x4e, 0x45, 0x7c, 0x99, 0xef, 0x6a, 0xbd, 0xfb, 0x20, 0x19, 0x04, 0x33, 0xd0, 0x9d, 0xc5, 0x74,
	0xf5, 0x0d, 0x53, 0xab, 0x19, 0x1e, 0x66, 0x1c, 0x29, 0xad, 0x73, 0x95, 0xb6, 0xbe, 0x30, 0x5a,
	0xe7, 0x26, 0xda, 0x59, 0x07, 0xfa, 0xd7, 0x82, 0xbd, 0x6c, 0x92, 0xba, 0x42, 0xbc, 0xbb, 0xdc,
	0xf8, 0xd6, 0x93, 0xe1, 0xec, 0x06, 0x68, 0xe7, 0x7a, 0x68, 0x7b, 0xb9, 0x61, 0xae, 0x27, 0x5f,
	0xce, 0x19, 0x5f, 0x7f, 0x45, 0xff, 0xaf, 0xf5, 0xf0, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a,
	0x4d, 0xa6, 0x3f, 0xf2, 0x12, 0x00, 0x00,
}
