package disks

import "github.com/multipleton/sa-3/machines"

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
	dbDisk.MachineId = dto.MachineId
	result, err = ds.disksRepository.Update(dbDisk)
	if err != nil {
		return result, err
	}
	err = ds.updateMachineTotalDiskSpace(dto.MachineId)
	if err != nil {
		err = ds.disksRepository.Delete(id)
	}
	return result, err
}

func (ds *DisksService) updateMachineTotalDiskSpace(id uint32) error {
	machine, err := ds.machinesService.FindById(id)
	if err == nil {
		return err
	}
	disks, err := ds.disksRepository.FindByMachineId(id)
	if err == nil {
		return err
	}
	var totalDiskSpace uint64
	for _, disk := range disks {
		totalDiskSpace += disk.Capacity
	}
	machine.TotalDiskSpace = totalDiskSpace
	_, err = ds.machinesService.Update(id, machine)
	return err
}

func NewDisksService(disksRepository *DisksRepository) *DisksService {
	return &DisksService{disksRepository: disksRepository}
}
