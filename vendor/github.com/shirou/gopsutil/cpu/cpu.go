package cpu

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
)

type CPUTimesStat struct {
	CPU       string
	User      float64
	System    float64
	Idle      float64
	Nice      float64
	Iowait    float64
	Irq       float64
	Softirq   float64
	Steal     float64
	Guest     float64
	GuestNice float64
	Stolen    float64
}

type CPUInfoStat struct {
	CPU        int32
	VendorID   string
	Family     string
	Model      string
	Stepping   int32
	PhysicalID string
	CoreID     string
	Cores      int32
	ModelName  string
	Mhz        float64
	CacheSize  int32
	Flags      []string
}

var lastCPUTimes []CPUTimesStat
var lastPerCPUTimes []CPUTimesStat

func CPUCounts(logical bool) (int, error) {
	return runtime.NumCPU(), nil
}

func (c CPUTimesStat) String() string {
	v := []string{
		`"cpu":"` + c.CPU + `"`,
		`"user":` + strconv.FormatFloat(c.User, 'f', 1, 64),
		`"system":` + strconv.FormatFloat(c.System, 'f', 1, 64),
		`"idle":` + strconv.FormatFloat(c.Idle, 'f', 1, 64),
		`"nice":` + strconv.FormatFloat(c.Nice, 'f', 1, 64),
		`"iowait":` + strconv.FormatFloat(c.Iowait, 'f', 1, 64),
		`"irq":` + strconv.FormatFloat(c.Irq, 'f', 1, 64),
		`"softirq":` + strconv.FormatFloat(c.Softirq, 'f', 1, 64),
		`"steal":` + strconv.FormatFloat(c.Steal, 'f', 1, 64),
		`"guest":` + strconv.FormatFloat(c.Guest, 'f', 1, 64),
		`"guest_nice":` + strconv.FormatFloat(c.GuestNice, 'f', 1, 64),
		`"stolen":` + strconv.FormatFloat(c.Stolen, 'f', 1, 64),
	}

	return `{` + strings.Join(v, ",") + `}`
}

func (c CPUInfoStat) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

func init() {
	lastCPUTimes, _ = CPUTimes(false)
	lastPerCPUTimes, _ = CPUTimes(true)
}
