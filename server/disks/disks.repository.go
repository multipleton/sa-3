package disks

import (
	"database/sql"
)

type DisksRepository struct {
	db *sql.DB
}

func (dr *DisksRepository) FindById(id uint32) (DisksEntity, error) {
	var result DisksEntity
	row := dr.db.QueryRow("SELECT * FROM disks WHERE id=?", id)
	err := row.Scan(&result.Id, &result.MachineId, &result.Capacity)
	return result, err
}

func (dr *DisksRepository) Update(entity DisksEntity) (DisksEntity, error) {
	var result DisksEntity
	query := "UPDATE disks SET machine_id=? capacity=? WHERE id=?"
	_, err := dr.db.Exec(query, entity.MachineId, entity.Capacity, entity.Id)
	if err == nil {
		result, err = dr.FindById(entity.Id)
	}
	return result, err
}

func NewDisksRepository(db *sql.DB) *DisksRepository {
	return &DisksRepository{db: db}
}
