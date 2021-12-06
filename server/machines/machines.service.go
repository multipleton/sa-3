package machines

type MachinesService struct {
	machinesRepository *MachinesRepository
}

func (ms *MachinesService) ReadAll() ([]*MachinesEntity, error) {
	machines, err := ms.machinesRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return machines, nil
}

func (ms *MachinesService) FindById(id uint32) (*MachinesEntity, error) {
	machine, err := ms.machinesRepository.GetOne(id)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

func (ms *MachinesService) Update(id uint32, machine *MachinesEntity) (*MachinesEntity, error) {
	_, err := ms.machinesRepository.GetOne(id)
	if err != nil {
		return nil, err
	}
	updatedMachine, err := ms.machinesRepository.UpdateOne(id, machine)
	if err != nil {
		return nil, err
	}
	return updatedMachine, nil
}

func NewMachinesService(machinesRepository *MachinesRepository) *MachinesService {
	return &MachinesService{machinesRepository: machinesRepository}
}
