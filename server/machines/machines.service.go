package machines

type MachinesService struct {
	mr *MachinesRepository
}

func (ms *MachinesService) ReadAll() ([]*MachinesEntity, error) {
	machines, err := ms.mr.GetAll()
	if err != nil {
		return nil, err
	}
	return machines, nil
}

func (ms *MachinesService) FindById(id uint32) (*MachinesEntity, error) {
	machine, err := ms.mr.GetOne(id)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

func (ms *MachinesService) Update(id uint32, machine MachinesEntity) (*MachinesEntity, error) {
	_, err := ms.mr.GetOne(id)
	if err != nil {
		return nil, err
	}
	updatedMachine, err := ms.mr.UpdateOne(id, machine)
	if err != nil {
		return nil, err
	}
	return updatedMachine, nil
}

func NewMachinesService(mr *MachinesRepository) *MachinesService {
	return &MachinesService{mr: mr}
}
