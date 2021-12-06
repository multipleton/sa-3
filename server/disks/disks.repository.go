package disks

import (
	"database/sql"
)

type DisksRepository struct {
	db *sql.DB
}

func (dr *DisksRepository) FindById(id uint32) (*DisksEntity, error) {
	var result DisksEntity
	row := dr.db.QueryRow("SELECT * FROM disks WHERE id=$1", id)
	err := row.Scan(&result.Id, &result.MachineId, &result.Capacity)
	return &result, err
}

func (dr *DisksRepository) FindByMachineId(machineId uint32) ([]*DisksEntity, error) {
	result := make([]*DisksEntity, 0)
	rows, err := dr.db.Query("SELECT * FROM disks WHERE machine_id=$1", machineId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var entry DisksEntity
		err = rows.Scan(&entry.Id, &entry.MachineId, &entry.Capacity)
		if err != nil {
			return nil, err
		}
		result = append(result, &entry)
	}
	return result, err
}

func (dr *DisksRepository) Update(entity *DisksEntity) (*DisksEntity, error) {
	var result *DisksEntity
	query := "UPDATE disks SET machine_id=$1, capacity=$2 WHERE id=$3"
	_, err := dr.db.Exec(query, entity.MachineId, entity.Capacity, entity.Id)
	if err == nil {
		result, err = dr.FindById(entity.Id)
	}
	return result, err
}

func (dr *DisksRepository) Delete(id uint32) error {
	query := "DELETE FROM disks WHERE id=$1"
	_, err := dr.db.Exec(query, id)
	return err
}

func NewDisksRepository(db *sql.DB) *DisksRepository {
	return &DisksRepository{db: db}
}
