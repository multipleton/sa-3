package machines

type MachinesService struct {
	Mr MachinesRepository
}

func (ms *MachinesService) ReadAll() ([]Machine, error) {
	machines, err := ms.Mr.GetAll()
	if err != nil {
		return nil, err
	}
	return machines, nil
}

func (ms *MachinesService) FindById(id string) (Machine, error) {
	machine, err := ms.Mr.GetOne(id)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

func (ms *MachinesService) Update(id string, machine Machine) (Machine, error) {
	machine, err := ms.Mr.GetOne(id)
	if err != nil {
		return nil, err
	}
	updatedMachine, err := ms.Mr.UpdateOne(id, machine)
	if err != nil {
		return nil, err
	}
	return updatedMachine, nil
}
