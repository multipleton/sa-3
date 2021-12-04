package disks

import (
	"database/sql"
)

type DisksRepository struct {
	DB *sql.DB
}

func (dr *DisksRepository) FindById(id uint32) (DisksEntity, error) {
	var result DisksEntity
	row := dr.DB.QueryRow("SELECT * FROM disks WHERE id=?", id)
	err := row.Scan(&result.Id, &result.MachineId, &result.Capacity)
	return result, err
}

func (dr *DisksRepository) Update(entity DisksEntity) (DisksEntity, error) {
	var result DisksEntity
	query := "UPDATE disks SET machine_id=? capacity=? WHERE id=?"
	_, err := dr.DB.Exec(query, entity.MachineId, entity.Capacity, entity.Id)
	if err != nil {
		return result, err
	}
	result, err = dr.FindById(entity.Id)
	return result, err
}
