package machines

type Machine struct {
	id             uint32 `json:"id"`
	name           string `json:"name"`
	cpuCount       uint8  `json:"cpuCount"`
	totalDiskSpace uint64 `json:"totalDiskSpace"`
}
