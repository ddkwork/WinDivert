package windivert

import (
	"fmt"
	"syscall"
	"unsafe"
)

type LPOVERLAPPED = *syscall.Overlapped

// Source: windivert.h
type (
	PWINDIVERT_LAYER        = *uint32
	PWINDIVERT_DATA_NETWORK = *struct{}
	PWINDIVERT_DATA_FLOW    = *struct{}
	PWINDIVERT_DATA_SOCKET  = *struct{}
	PWINDIVERT_DATA_REFLECT = *struct{}
	PWINDIVERT_ADDRESS      = *struct{}
	PWINDIVERT_EVENT        = *uint32
	PWINDIVERT_PARAM        = *uint32
	PWINDIVERT_SHUTDOWN     = *uint32
	PWINDIVERT_IPHDR        = *struct{}
	PWINDIVERT_IPV6HDR      = *struct{}
	PWINDIVERT_ICMPHDR      = *struct{}
	PWINDIVERT_ICMPV6HDR    = *struct{}
	PWINDIVERT_TCPHDR       = *struct{}
	PWINDIVERT_UDPHDR       = *struct{}
)

// Source: windivert.h:0 -> WINDIVERT_LAYER
type WINDIVERT_LAYER uint32

const (
	WindivertLayerNetwork WINDIVERT_LAYER = iota
	WindivertLayerNetworkForward
	WindivertLayerFlow
	WindivertLayerSocket
	WindivertLayerReflect
)

func (w WINDIVERT_LAYER) String() string {
	switch w {
	case WindivertLayerNetwork:
		return "Windivert Layer Network"
	case WindivertLayerNetworkForward:
		return "Windivert Layer Network Forward"
	case WindivertLayerFlow:
		return "Windivert Layer Flow"
	case WindivertLayerSocket:
		return "Windivert Layer Socket"
	case WindivertLayerReflect:
		return "Windivert Layer Reflect"
	default:
		return fmt.Sprintf("WINDIVERT_LAYER(0x%X)", uint32(w))
	}
}

// Source: windivert.h:0 -> WINDIVERT_EVENT
type WINDIVERT_EVENT uint32

const (
	WindivertEventNetworkPacket WINDIVERT_EVENT = iota
	WindivertEventFlowEstablished
	WindivertEventFlowDeleted
	WindivertEventSocketBind
	WindivertEventSocketConnect
	WindivertEventSocketListen
	WindivertEventSocketAccept
	WindivertEventSocketClose
	WindivertEventReflectOpen
	WindivertEventReflectClose
)

func (w WINDIVERT_EVENT) String() string {
	switch w {
	case WindivertEventNetworkPacket:
		return "Windivert Event Network Packet"
	case WindivertEventFlowEstablished:
		return "Windivert Event Flow Established"
	case WindivertEventFlowDeleted:
		return "Windivert Event Flow Deleted"
	case WindivertEventSocketBind:
		return "Windivert Event Socket Bind"
	case WindivertEventSocketConnect:
		return "Windivert Event Socket Connect"
	case WindivertEventSocketListen:
		return "Windivert Event Socket Listen"
	case WindivertEventSocketAccept:
		return "Windivert Event Socket Accept"
	case WindivertEventSocketClose:
		return "Windivert Event Socket Close"
	case WindivertEventReflectOpen:
		return "Windivert Event Reflect Open"
	case WindivertEventReflectClose:
		return "Windivert Event Reflect Close"
	default:
		return fmt.Sprintf("WINDIVERT_EVENT(0x%X)", uint32(w))
	}
}

// Source: windivert.h:0 -> WINDIVERT_PARAM
type WINDIVERT_PARAM uint32

const (
	WindivertParamQueueLength WINDIVERT_PARAM = iota
	WindivertParamQueueTime
	WindivertParamQueueSize
	WindivertParamVersionMajor
	WindivertParamVersionMinor
)

func (w WINDIVERT_PARAM) String() string {
	switch w {
	case WindivertParamQueueLength:
		return "Windivert Param Queue Length"
	case WindivertParamQueueTime:
		return "Windivert Param Queue Time"
	case WindivertParamQueueSize:
		return "Windivert Param Queue Size"
	case WindivertParamVersionMajor:
		return "Windivert Param Version Major"
	case WindivertParamVersionMinor:
		return "Windivert Param Version Minor"
	default:
		return fmt.Sprintf("WINDIVERT_PARAM(0x%X)", uint32(w))
	}
}

// Source: windivert.h:0 -> WINDIVERT_SHUTDOWN
type WINDIVERT_SHUTDOWN uint32

const (
	WindivertShutdownRecv WINDIVERT_SHUTDOWN = 1 + iota
	WindivertShutdownSend
	WindivertShutdownBoth
)

func (w WINDIVERT_SHUTDOWN) String() string {
	switch w {
	case WindivertShutdownRecv:
		return "Windivert Shutdown Recv"
	case WindivertShutdownSend:
		return "Windivert Shutdown Send"
	case WindivertShutdownBoth:
		return "Windivert Shutdown Both"
	default:
		return fmt.Sprintf("WINDIVERT_SHUTDOWN(0x%X)", uint32(w))
	}
}

func (w *WINDIVERT_ADDRESS) GetLayer() uint32 {
	return uint32(w.LayerBits & uint32(0xFF))
}

func (w *WINDIVERT_ADDRESS) SetLayer(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0xFF)) | uint32(uint32(val)&0xFF)
}

func (w *WINDIVERT_ADDRESS) GetEvent() uint32 {
	return uint32((w.LayerBits >> 8) & uint32(0xFF))
}

func (w *WINDIVERT_ADDRESS) SetEvent(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0xFF<<8)) | (uint32(uint32(val)&0xFF) << 8)
}

func (w *WINDIVERT_ADDRESS) GetSniffed() uint32 {
	return uint32((w.LayerBits >> 16) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetSniffed(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<16)) | (uint32(uint32(val)&0x1) << 16)
}

func (w *WINDIVERT_ADDRESS) GetOutbound() uint32 {
	return uint32((w.LayerBits >> 17) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetOutbound(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<17)) | (uint32(uint32(val)&0x1) << 17)
}

func (w *WINDIVERT_ADDRESS) GetLoopback() uint32 {
	return uint32((w.LayerBits >> 18) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetLoopback(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<18)) | (uint32(uint32(val)&0x1) << 18)
}

func (w *WINDIVERT_ADDRESS) GetImpostor() uint32 {
	return uint32((w.LayerBits >> 19) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetImpostor(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<19)) | (uint32(uint32(val)&0x1) << 19)
}

func (w *WINDIVERT_ADDRESS) GetIPv6() uint32 {
	return uint32((w.LayerBits >> 20) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetIPv6(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<20)) | (uint32(uint32(val)&0x1) << 20)
}

func (w *WINDIVERT_ADDRESS) GetIPChecksum() uint32 {
	return uint32((w.LayerBits >> 21) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetIPChecksum(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<21)) | (uint32(uint32(val)&0x1) << 21)
}

func (w *WINDIVERT_ADDRESS) GetTCPChecksum() uint32 {
	return uint32((w.LayerBits >> 22) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetTCPChecksum(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<22)) | (uint32(uint32(val)&0x1) << 22)
}

func (w *WINDIVERT_ADDRESS) GetUDPChecksum() uint32 {
	return uint32((w.LayerBits >> 23) & uint32(0x1))
}

func (w *WINDIVERT_ADDRESS) SetUDPChecksum(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0x1<<23)) | (uint32(uint32(val)&0x1) << 23)
}

func (w *WINDIVERT_ADDRESS) GetReserved1() uint32 {
	return uint32((w.LayerBits >> 24) & uint32(0xFF))
}

func (w *WINDIVERT_ADDRESS) SetReserved1(val uint32) {
	w.LayerBits = (w.LayerBits & ^uint32(0xFF<<24)) | (uint32(uint32(val)&0xFF) << 24)
}

func (w *WINDIVERT_IPHDR) GetHdrLength() uint8 {
	return uint8(w.HdrLengthBits & uint8(0xF))
}

func (w *WINDIVERT_IPHDR) SetHdrLength(val uint8) {
	w.HdrLengthBits = (w.HdrLengthBits & ^uint8(0xF)) | uint8(uint8(val)&0xF)
}

func (w *WINDIVERT_IPHDR) GetVersion() uint8 {
	return uint8((w.HdrLengthBits >> 4) & uint8(0xF))
}

func (w *WINDIVERT_IPHDR) SetVersion(val uint8) {
	w.HdrLengthBits = (w.HdrLengthBits & ^uint8(0xF<<4)) | (uint8(uint8(val)&0xF) << 4)
}

func (w *WINDIVERT_IPV6HDR) GetTrafficClass0() uint8 {
	return uint8(w.TrafficClass0Bits & uint8(0xF))
}

func (w *WINDIVERT_IPV6HDR) SetTrafficClass0(val uint8) {
	w.TrafficClass0Bits = (w.TrafficClass0Bits & ^uint8(0xF)) | uint8(uint8(val)&0xF)
}

func (w *WINDIVERT_IPV6HDR) GetVersion() uint8 {
	return uint8((w.TrafficClass0Bits >> 4) & uint8(0xF))
}

func (w *WINDIVERT_IPV6HDR) SetVersion(val uint8) {
	w.TrafficClass0Bits = (w.TrafficClass0Bits & ^uint8(0xF<<4)) | (uint8(uint8(val)&0xF) << 4)
}

func (w *WINDIVERT_IPV6HDR) GetFlowLabel0() uint8 {
	return uint8(w.FlowLabel0Bits & uint8(0xF))
}

func (w *WINDIVERT_IPV6HDR) SetFlowLabel0(val uint8) {
	w.FlowLabel0Bits = (w.FlowLabel0Bits & ^uint8(0xF)) | uint8(uint8(val)&0xF)
}

func (w *WINDIVERT_IPV6HDR) GetTrafficClass1() uint8 {
	return uint8((w.FlowLabel0Bits >> 4) & uint8(0xF))
}

func (w *WINDIVERT_IPV6HDR) SetTrafficClass1(val uint8) {
	w.FlowLabel0Bits = (w.FlowLabel0Bits & ^uint8(0xF<<4)) | (uint8(uint8(val)&0xF) << 4)
}

func (w *WINDIVERT_TCPHDR) GetReserved1() uint16 {
	return uint16(w.Reserved1Bits & uint16(0xF))
}

func (w *WINDIVERT_TCPHDR) SetReserved1(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0xF)) | uint16(uint16(val)&0xF)
}

func (w *WINDIVERT_TCPHDR) GetHdrLength() uint16 {
	return uint16((w.Reserved1Bits >> 4) & uint16(0xF))
}

func (w *WINDIVERT_TCPHDR) SetHdrLength(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0xF<<4)) | (uint16(uint16(val)&0xF) << 4)
}

func (w *WINDIVERT_TCPHDR) GetFin() uint16 {
	return uint16((w.Reserved1Bits >> 8) & uint16(0x1))
}

func (w *WINDIVERT_TCPHDR) SetFin(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0x1<<8)) | (uint16(uint16(val)&0x1) << 8)
}

func (w *WINDIVERT_TCPHDR) GetSyn() uint16 {
	return uint16((w.Reserved1Bits >> 9) & uint16(0x1))
}

func (w *WINDIVERT_TCPHDR) SetSyn(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0x1<<9)) | (uint16(uint16(val)&0x1) << 9)
}

func (w *WINDIVERT_TCPHDR) GetRst() uint16 {
	return uint16((w.Reserved1Bits >> 10) & uint16(0x1))
}

func (w *WINDIVERT_TCPHDR) SetRst(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0x1<<10)) | (uint16(uint16(val)&0x1) << 10)
}

func (w *WINDIVERT_TCPHDR) GetPsh() uint16 {
	return uint16((w.Reserved1Bits >> 11) & uint16(0x1))
}

func (w *WINDIVERT_TCPHDR) SetPsh(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0x1<<11)) | (uint16(uint16(val)&0x1) << 11)
}

func (w *WINDIVERT_TCPHDR) GetAck() uint16 {
	return uint16((w.Reserved1Bits >> 12) & uint16(0x1))
}

func (w *WINDIVERT_TCPHDR) SetAck(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0x1<<12)) | (uint16(uint16(val)&0x1) << 12)
}

func (w *WINDIVERT_TCPHDR) GetUrg() uint16 {
	return uint16((w.Reserved1Bits >> 13) & uint16(0x1))
}

func (w *WINDIVERT_TCPHDR) SetUrg(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0x1<<13)) | (uint16(uint16(val)&0x1) << 13)
}

func (w *WINDIVERT_TCPHDR) GetReserved2() uint16 {
	return uint16((w.Reserved1Bits >> 14) & uint16(0x3))
}

func (w *WINDIVERT_TCPHDR) SetReserved2(val uint16) {
	w.Reserved1Bits = (w.Reserved1Bits & ^uint16(0x3<<14)) | (uint16(uint16(val)&0x3) << 14)
}

type (
	WINDIVERT_DATA_NETWORK struct {
		IfIdx    uint32
		SubIfIdx uint32
	} // windivert.h:0 -> WINDIVERT_DATA_NETWORK
	WINDIVERT_DATA_FLOW struct {
		EndpointId       uint64
		ParentEndpointId uint64
		ProcessId        uint32
		LocalAddr        [4]uint32
		RemoteAddr       [4]uint32
		LocalPort        uint16
		RemotePort       uint16
		Protocol         uint8
		_                [7]byte
	} // windivert.h:0 -> WINDIVERT_DATA_FLOW
	WINDIVERT_DATA_SOCKET struct {
		EndpointId       uint64
		ParentEndpointId uint64
		ProcessId        uint32
		LocalAddr        [4]uint32
		RemoteAddr       [4]uint32
		LocalPort        uint16
		RemotePort       uint16
		Protocol         uint8
		_                [7]byte
	} // windivert.h:0 -> WINDIVERT_DATA_SOCKET
	WINDIVERT_DATA_REFLECT struct {
		Timestamp int64
		ProcessId uint32
		Layer     WINDIVERT_LAYER
		Flags     uint64
		Priority  int16
		_         [6]byte
	} // windivert.h:0 -> WINDIVERT_DATA_REFLECT
	WINDIVERT_ADDRESS struct {
		Timestamp int64
		LayerBits uint32
		Reserved2 uint32
		Anon1     WINDIVERT_ADDRESS_Anon1Union
	} // windivert.h:0 -> WINDIVERT_ADDRESS
	WINDIVERT_ADDRESS_Anon1Union_ struct {
		Network   WINDIVERT_DATA_NETWORK
		Flow      WINDIVERT_DATA_FLOW
		Socket    WINDIVERT_DATA_SOCKET
		Reflect   WINDIVERT_DATA_REFLECT
		Reserved3 [64]uint8
	} // windivert.h:0 -> WINDIVERT_ADDRESS_Anon1Union
	WINDIVERT_ADDRESS_Anon1Union struct {
		Data WINDIVERT_DATA_FLOW
	} // windivert.h:0 -> WINDIVERT_ADDRESS_Anon1Union
	WINDIVERT_IPHDR struct {
		HdrLengthBits uint8
		TOS           uint8
		Length        uint16
		Id            uint16
		FragOff0      uint16
		TTL           uint8
		Protocol      uint8
		Checksum      uint16
		SrcAddr       uint32
		DstAddr       uint32
	} // windivert.h:0 -> WINDIVERT_IPHDR
	WINDIVERT_IPV6HDR struct {
		TrafficClass0Bits uint8
		FlowLabel0Bits    uint8
		FlowLabel1        uint16
		Length            uint16
		NextHdr           uint8
		HopLimit          uint8
		SrcAddr           [4]uint32
		DstAddr           [4]uint32
	} // windivert.h:0 -> WINDIVERT_IPV6HDR
	WINDIVERT_ICMPHDR struct {
		Type     uint8
		Code     uint8
		Checksum uint16
		Body     uint32
	} // windivert.h:0 -> WINDIVERT_ICMPHDR
	WINDIVERT_ICMPV6HDR struct {
		Type     uint8
		Code     uint8
		Checksum uint16
		Body     uint32
	} // windivert.h:0 -> WINDIVERT_ICMPV6HDR
	WINDIVERT_TCPHDR struct {
		SrcPort       uint16
		DstPort       uint16
		SeqNum        uint32
		AckNum        uint32
		Reserved1Bits uint16
		Window        uint16
		Checksum      uint16
		UrgPtr        uint16
	} // windivert.h:0 -> WINDIVERT_TCPHDR
	WINDIVERT_UDPHDR struct {
		SrcPort  uint16
		DstPort  uint16
		Length   uint16
		Checksum uint16
	} // windivert.h:0 -> WINDIVERT_UDPHDR
)

// Source: windivert.h -> Macro constants
const (
	WindivertFlagSniff               uint32 = 0x0001
	WindivertFlagDrop                uint32 = 0x0002
	WindivertFlagRecvOnly            uint32 = 0x0004
	WindivertFlagReadOnly            uint32 = WindivertFlagRecvOnly
	WindivertFlagSendOnly            uint32 = 0x0008
	WindivertFlagWriteOnly           uint32 = WindivertFlagSendOnly
	WindivertFlagNoInstall           uint32 = 0x0010
	WindivertFlagFragments           uint32 = 0x0020
	WindivertPriorityHighest         int32  = 30000
	WindivertPriorityLowest          int32  = (-WindivertPriorityHighest)
	WindivertParamQueueLengthDefault uint32 = 4096
	WindivertParamQueueLengthMin     uint32 = 32
	WindivertParamQueueLengthMax     uint32 = 16384
	WindivertParamQueueTimeDefault   uint32 = 2000
	WindivertParamQueueTimeMin       uint32 = 100
	WindivertParamQueueTimeMax       uint32 = 16000
	WindivertParamQueueSizeDefault   uint32 = 4194304
	WindivertParamQueueSizeMin       uint32 = 65535
	WindivertParamQueueSizeMax       uint32 = 33554432
	WindivertBatchMax                uint32 = 0xFF
	WindivertMtuMax                  uint32 = (40 + 0xFFFF)
	WindivertHelperNoIpChecksum      uint32 = 1
	WindivertHelperNoIcmpChecksum    uint32 = 2
	WindivertHelperNoIcmpv6Checksum  uint32 = 4
	WindivertHelperNoTcpChecksum     uint32 = 8
	WindivertHelperNoUdpChecksum     uint32 = 16
)

func (w *Windivert) WinDivertOpen(Filter *int8, Layer WINDIVERT_LAYER, Priority int16, Flags uint64) uintptr {
	r1, _, _ := getProc("WinDivertOpen").Call(uintptr(unsafe.Pointer(Filter)), uintptr(Layer), uintptr(Priority), *(*uintptr)(unsafe.Pointer(&Flags)))
	return uintptr(r1)
}

func (w *Windivert) WinDivertRecv(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PRecvLen *uint32, PAddr *WINDIVERT_ADDRESS) int32 {
	r1, _, _ := getProc("WinDivertRecv").Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PRecvLen)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertRecvEx(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PRecvLen *uint32, Flags uint64, PAddr *WINDIVERT_ADDRESS, PAddrLen *uint32, LpOverlapped LPOVERLAPPED) int32 {
	r1, _, _ := getProc("WinDivertRecvEx").Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PRecvLen)), *(*uintptr)(unsafe.Pointer(&Flags)), uintptr(unsafe.Pointer(PAddr)), uintptr(unsafe.Pointer(PAddrLen)), uintptr(unsafe.Pointer(LpOverlapped)))
	return int32(r1)
}

func (w *Windivert) WinDivertSend(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PSendLen *uint32, PAddr *WINDIVERT_ADDRESS) int32 {
	r1, _, _ := getProc("WinDivertSend").Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PSendLen)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertSendEx(Handle uintptr, PPacket unsafe.Pointer, PacketLen uint32, PSendLen *uint32, Flags uint64, PAddr *WINDIVERT_ADDRESS, AddrLen uint32, LpOverlapped LPOVERLAPPED) int32 {
	r1, _, _ := getProc("WinDivertSendEx").Call(Handle, uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PSendLen)), *(*uintptr)(unsafe.Pointer(&Flags)), uintptr(unsafe.Pointer(PAddr)), uintptr(AddrLen), uintptr(unsafe.Pointer(LpOverlapped)))
	return int32(r1)
}

func (w *Windivert) WinDivertShutdown(Handle uintptr, How WINDIVERT_SHUTDOWN) int32 {
	r1, _, _ := getProc("WinDivertShutdown").Call(Handle, uintptr(How))
	return int32(r1)
}

func (w *Windivert) WinDivertClose(Handle uintptr) int32 {
	r1, _, _ := getProc("WinDivertClose").Call(Handle)
	return int32(r1)
}

func (w *Windivert) WinDivertSetParam(Handle uintptr, Param WINDIVERT_PARAM, Value uint64) int32 {
	r1, _, _ := getProc("WinDivertSetParam").Call(Handle, uintptr(Param), *(*uintptr)(unsafe.Pointer(&Value)))
	return int32(r1)
}

func (w *Windivert) WinDivertGetParam(Handle uintptr, Param WINDIVERT_PARAM, PValue *uint64) int32 {
	r1, _, _ := getProc("WinDivertGetParam").Call(Handle, uintptr(Param), uintptr(unsafe.Pointer(PValue)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperHashPacket(PPacket unsafe.Pointer, PacketLen uint32, Seed uint64) uint64 {
	r1, _, _ := getProc("WinDivertHelperHashPacket").Call(uintptr(PPacket), uintptr(PacketLen), *(*uintptr)(unsafe.Pointer(&Seed)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (w *Windivert) WinDivertHelperParsePacket(PPacket unsafe.Pointer, PacketLen uint32, PpIpHdr *PWINDIVERT_IPHDR, PpIpv6Hdr *PWINDIVERT_IPV6HDR, PProtocol *uint8, PpIcmpHdr *PWINDIVERT_ICMPHDR, PpIcmpv6Hdr *PWINDIVERT_ICMPV6HDR, PpTcpHdr *PWINDIVERT_TCPHDR, PpUdpHdr *PWINDIVERT_UDPHDR, PpData *uintptr, PDataLen *uint32, PpNext *uintptr, PNextLen *uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperParsePacket").Call(uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PpIpHdr)), uintptr(unsafe.Pointer(PpIpv6Hdr)), uintptr(unsafe.Pointer(PProtocol)), uintptr(unsafe.Pointer(PpIcmpHdr)), uintptr(unsafe.Pointer(PpIcmpv6Hdr)), uintptr(unsafe.Pointer(PpTcpHdr)), uintptr(unsafe.Pointer(PpUdpHdr)), uintptr(unsafe.Pointer(PpData)), uintptr(unsafe.Pointer(PDataLen)), uintptr(unsafe.Pointer(PpNext)), uintptr(unsafe.Pointer(PNextLen)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperParseIPv4Address(AddrStr *int8, PAddr *uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperParseIPv4Address").Call(uintptr(unsafe.Pointer(AddrStr)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperParseIPv6Address(AddrStr *int8, PAddr *uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperParseIPv6Address").Call(uintptr(unsafe.Pointer(AddrStr)), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperFormatIPv4Address(Addr uint32, Buffer *int8, BufLen uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperFormatIPv4Address").Call(uintptr(Addr), uintptr(unsafe.Pointer(Buffer)), uintptr(BufLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperFormatIPv6Address(PAddr *uint32, Buffer *int8, BufLen uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperFormatIPv6Address").Call(uintptr(unsafe.Pointer(PAddr)), uintptr(unsafe.Pointer(Buffer)), uintptr(BufLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperCalcChecksums(PPacket unsafe.Pointer, PacketLen uint32, PAddr *WINDIVERT_ADDRESS, Flags uint64) int32 {
	r1, _, _ := getProc("WinDivertHelperCalcChecksums").Call(uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PAddr)), *(*uintptr)(unsafe.Pointer(&Flags)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperDecrementTTL(PPacket unsafe.Pointer, PacketLen uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperDecrementTTL").Call(uintptr(PPacket), uintptr(PacketLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperCompileFilter(Filter *int8, Layer WINDIVERT_LAYER, Object *int8, ObjLen uint32, ErrorStr **int8, ErrorPos *uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperCompileFilter").Call(uintptr(unsafe.Pointer(Filter)), uintptr(Layer), uintptr(unsafe.Pointer(Object)), uintptr(ObjLen), uintptr(unsafe.Pointer(ErrorStr)), uintptr(unsafe.Pointer(ErrorPos)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperEvalFilter(Filter *int8, PPacket unsafe.Pointer, PacketLen uint32, PAddr *WINDIVERT_ADDRESS) int32 {
	r1, _, _ := getProc("WinDivertHelperEvalFilter").Call(uintptr(unsafe.Pointer(Filter)), uintptr(PPacket), uintptr(PacketLen), uintptr(unsafe.Pointer(PAddr)))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperFormatFilter(Filter *int8, Layer WINDIVERT_LAYER, Buffer *int8, BufLen uint32) int32 {
	r1, _, _ := getProc("WinDivertHelperFormatFilter").Call(uintptr(unsafe.Pointer(Filter)), uintptr(Layer), uintptr(unsafe.Pointer(Buffer)), uintptr(BufLen))
	return int32(r1)
}

func (w *Windivert) WinDivertHelperNtohs(X uint16) uint16 {
	r1, _, _ := getProc("WinDivertHelperNtohs").Call(uintptr(X))
	return uint16(r1)
}

func (w *Windivert) WinDivertHelperHtons(X uint16) uint16 {
	r1, _, _ := getProc("WinDivertHelperHtons").Call(uintptr(X))
	return uint16(r1)
}

func (w *Windivert) WinDivertHelperNtohl(X uint32) uint32 {
	r1, _, _ := getProc("WinDivertHelperNtohl").Call(uintptr(X))
	return uint32(r1)
}

func (w *Windivert) WinDivertHelperHtonl(X uint32) uint32 {
	r1, _, _ := getProc("WinDivertHelperHtonl").Call(uintptr(X))
	return uint32(r1)
}

func (w *Windivert) WinDivertHelperNtohll(X uint64) uint64 {
	r1, _, _ := getProc("WinDivertHelperNtohll").Call(*(*uintptr)(unsafe.Pointer(&X)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (w *Windivert) WinDivertHelperHtonll(X uint64) uint64 {
	r1, _, _ := getProc("WinDivertHelperHtonll").Call(*(*uintptr)(unsafe.Pointer(&X)))
	return *(*uint64)(unsafe.Pointer(&r1))
}

func (w *Windivert) WinDivertHelperNtohIPv6Address(InAddr *uint32, OutAddr *uint32) {
	getProc("WinDivertHelperNtohIPv6Address").Call(uintptr(unsafe.Pointer(InAddr)), uintptr(unsafe.Pointer(OutAddr)))
}

func (w *Windivert) WinDivertHelperHtonIPv6Address(InAddr *uint32, OutAddr *uint32) {
	getProc("WinDivertHelperHtonIPv6Address").Call(uintptr(unsafe.Pointer(InAddr)), uintptr(unsafe.Pointer(OutAddr)))
}

func (w *Windivert) WinDivertHelperNtohIpv6Address(InAddr *uint32, OutAddr *uint32) {
	getProc("WinDivertHelperNtohIpv6Address").Call(uintptr(unsafe.Pointer(InAddr)), uintptr(unsafe.Pointer(OutAddr)))
}

func (w *Windivert) WinDivertHelperHtonIpv6Address(InAddr *uint32, OutAddr *uint32) {
	getProc("WinDivertHelperHtonIpv6Address").Call(uintptr(unsafe.Pointer(InAddr)), uintptr(unsafe.Pointer(OutAddr)))
}
