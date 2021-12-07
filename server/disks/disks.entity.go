package disks

type DisksEntity struct {
	Id        uint32  `json:"id"`
	MachineId *uint32 `json:"machineId"`
	Capacity  uint64  `json:"capacity"`
}
