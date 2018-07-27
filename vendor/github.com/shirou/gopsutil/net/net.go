package net

import (
	"encoding/json"
	"net"
)

type NetIOCountersStat struct {
	Name        string // interface name
	BytesSent   uint64 // number of bytes sent
	BytesRecv   uint64 // number of bytes received
	PacketsSent uint64 // number of packets sent
	PacketsRecv uint64 // number of packets received
	Errin       uint64 // total number of errors while receiving
	Errout      uint64 // total number of errors while sending
	Dropin      uint64 // total number of incoming packets which were dropped
	Dropout     uint64 // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
}

// Addr is implemented compatibility to psutil
type Addr struct {
	IP   string
	Port uint32
}

type NetConnectionStat struct {
	Fd     uint32
	Family uint32
	Type   uint32
	Laddr  Addr
	Raddr  Addr
	Status string
	Pid    int32
}

// NetInterfaceAddr is designed for represent interface addresses
type NetInterfaceAddr struct {
	Addr string
}

type NetInterfaceStat struct {
	MTU          int      // maximum transmission unit
	Name         string   // e.g., "en0", "lo0", "eth0.100"
	HardwareAddr string   // IEEE MAC-48, EUI-48 and EUI-64 form
	Flags        []string // e.g., FlagUp, FlagLoopback, FlagMulticast
	Addrs        []NetInterfaceAddr
}

func (n NetIOCountersStat) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}

func (n NetConnectionStat) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}

func (a Addr) String() string {
	s, _ := json.Marshal(a)
	return string(s)
}

func (n NetInterfaceStat) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}

func (n NetInterfaceAddr) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}

func NetInterfaces() ([]NetInterfaceStat, error) {
	is, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ret := make([]NetInterfaceStat, 0, len(is))
	for _, ifi := range is {

		var flags []string
		if ifi.Flags&net.FlagUp != 0 {
			flags = append(flags, "up")
		}
		if ifi.Flags&net.FlagBroadcast != 0 {
			flags = append(flags, "broadcast")
		}
		if ifi.Flags&net.FlagLoopback != 0 {
			flags = append(flags, "loopback")
		}
		if ifi.Flags&net.FlagPointToPoint != 0 {
			flags = append(flags, "pointtopoint")
		}
		if ifi.Flags&net.FlagMulticast != 0 {
			flags = append(flags, "multicast")
		}

		r := NetInterfaceStat{
			Name:         ifi.Name,
			MTU:          ifi.MTU,
			HardwareAddr: ifi.HardwareAddr.String(),
			Flags:        flags,
		}
		addrs, err := ifi.Addrs()
		if err == nil {
			r.Addrs = make([]NetInterfaceAddr, 0, len(addrs))
			for _, addr := range addrs {
				r.Addrs = append(r.Addrs, NetInterfaceAddr{
					Addr: addr.String(),
				})
			}

		}
		ret = append(ret, r)
	}

	return ret, nil
}

func getNetIOCountersAll(n []NetIOCountersStat) ([]NetIOCountersStat, error) {
	r := NetIOCountersStat{
		Name: "all",
	}
	for _, nic := range n {
		r.BytesRecv += nic.BytesRecv
		r.PacketsRecv += nic.PacketsRecv
		r.Errin += nic.Errin
		r.Dropin += nic.Dropin
		r.BytesSent += nic.BytesSent
		r.PacketsSent += nic.PacketsSent
		r.Errout += nic.Errout
		r.Dropout += nic.Dropout
	}

	return []NetIOCountersStat{r}, nil
}
