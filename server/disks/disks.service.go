package disks

import (
	"github.com/multipleton/sa-3/machines"
)

type DisksService struct {
	disksRepository *DisksRepository
	machinesService *machines.MachinesService
}

func (ds *DisksService) FindById(id uint32) (*DisksEntity, error) {
	result, err := ds.disksRepository.FindById(id)
	return result, err
}

func (ds *DisksService) ConnectToMachine(id uint32, dto *DiskConnectToMachineDto) (*DisksEntity, error) {
	var result *DisksEntity
	dbDisk, err := ds.FindById(id)
	if err != nil {
		return result, nil
	}
	var disk DisksEntity
	disk.Id = dbDisk.Id
	disk.MachineId = dto.MachineId
	disk.Capacity = dbDisk.Capacity
	result, err = ds.disksRepository.Update(&disk)
	if err != nil {
		return result, err
	}
	if dbDisk.MachineId > 0 {
		ds.updateMachineTotalDiskSpace(dbDisk.MachineId)
	}
	err = ds.updateMachineTotalDiskSpace(disk.MachineId)
	if err != nil {
		result, err = ds.disksRepository.Update(dbDisk)
	}
	return result, err
}

func (ds *DisksService) updateMachineTotalDiskSpace(id uint32) error {
	machine, err := ds.machinesService.FindById(id)
	if err != nil {
		return err
	}
	disks, err := ds.disksRepository.FindByMachineId(id)
	if err != nil {
		return err
	}
	var totalDiskSpace uint64 = 0
	for _, disk := range disks {
		totalDiskSpace += disk.Capacity
	}
	machine.TotalDiskSpace = totalDiskSpace
	_, err = ds.machinesService.Update(id, machine)
	return err
}

func NewDisksService(disksRepository *DisksRepository, machiesService *machines.MachinesService) *DisksService {
	return &DisksService{disksRepository: disksRepository, machinesService: machiesService}
}
