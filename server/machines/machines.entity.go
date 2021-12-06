package machines

type MachinesEntity struct {
	Id             uint32 `json:"id"`
	Name           string `json:"name"`
	CpuCount       uint8  `json:"cpuCount"`
	TotalDiskSpace uint64 `json:"totalDiskSpace"`
}
