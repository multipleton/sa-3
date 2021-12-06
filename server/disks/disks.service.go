package disks

type DisksService struct {
	disksRepository *DisksRepository
}

func (ds *DisksService) FindById(id uint32) (DisksEntity, error) {
	result, err := ds.disksRepository.FindById(id)
	return result, err
}

func (ds *DisksService) ConnectToMachine(id uint32, dto DiskConnectToMachineDto) (DisksEntity, error) {
	var result DisksEntity
	dbDisk, err := ds.FindById(id)
	if err != nil {
		return result, nil
	}
	dbDisk.MachineId = dto.MachineId
	result, err = ds.disksRepository.Update(dbDisk)
	return result, err
}

func NewDisksService(disksRepository *DisksRepository) *DisksService {
	return &DisksService{disksRepository: disksRepository}
}