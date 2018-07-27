package disk

import (
	"encoding/json"
)

type DiskUsageStat struct {
	Path              string
	Total             uint64
	Free              uint64
	Used              uint64
	UsedPercent       float64
	InodesTotal       uint64
	InodesUsed        uint64
	InodesFree        uint64
	InodesUsedPercent float64
}

type DiskPartitionStat struct {
	Device     string
	Mountpoint string
	Fstype     string
	Opts       string
}

type DiskIOCountersStat struct {
	ReadCount    uint64
	WriteCount   uint64
	ReadBytes    uint64
	WriteBytes   uint64
	ReadTime     uint64
	WriteTime    uint64
	Name         string
	IoTime       uint64
	SerialNumber string
}

func (d DiskUsageStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func (d DiskPartitionStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func (d DiskIOCountersStat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}
