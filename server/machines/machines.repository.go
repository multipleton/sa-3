package machines

import (
	"database/sql"
	"fmt"
)

type MachinesRepository struct {
	db *sql.DB
}

func (mr *MachinesRepository) GetAll() ([]*MachinesEntity, error) {
	rows, err := mr.db.Query("SELECT id, name, cpu_count, total_disk_space FROM machines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*MachinesEntity
	for rows.Next() {
		var m MachinesEntity
		if err := rows.Scan(&m.Id, &m.Name, &m.CpuCount, &m.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, &m)
	}
	if res == nil {
		res = make([]*MachinesEntity, 0)
	}
	return res, nil
}

func (mr *MachinesRepository) GetOne(id uint32) (*MachinesEntity, error) {
	query := fmt.Sprintf("SELECT id, name, cpu_count, total_disk_space FROM machines WHERE id=%d", id)
	row, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var m MachinesEntity
	if err := row.Scan(&m.Id, &m.Name, &m.CpuCount, &m.TotalDiskSpace); err != nil {
		return nil, err
	}
	return &m, nil
}

func (mr *MachinesRepository) UpdateOne(id uint32, machine *MachinesEntity) (*MachinesEntity, error) {
	query := fmt.Sprintf("UPDATE machines SET total_disk_space=%d WHERE id=%d RETURNING id, name, cpu_count, total_disk_space", machine.TotalDiskSpace, id)
	row, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var m MachinesEntity
	if err := row.Scan(&m.Id, &m.Name, &m.CpuCount, &m.TotalDiskSpace); err != nil {
		return nil, err
	}
	return &m, nil
}

func NewMachinesRepository(db *sql.DB) *MachinesRepository {
	return &MachinesRepository{db: db}
}
