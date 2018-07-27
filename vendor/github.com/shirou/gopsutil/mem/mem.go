package mem

import (
	"encoding/json"
)

type VirtualMemoryStat struct {
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
	Active      uint64
	Inactive    uint64
	Buffers     uint64
	Cached      uint64
	Wired       uint64
	Shared      uint64
}

type SwapMemoryStat struct {
	Total       uint64
	Used        uint64
	Free        uint64
	UsedPercent float64
	Sin         uint64
	Sout        uint64
}

func (m VirtualMemoryStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func (m SwapMemoryStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
