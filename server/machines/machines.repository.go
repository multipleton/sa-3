package machines

import (
	"database/sql"
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
	var result MachinesEntity
	row := mr.db.QueryRow("SELECT * FROM machines WHERE id=$1", id)
	err := row.Scan(&result.Id, &result.Name, &result.CpuCount, &result.TotalDiskSpace)
	return &result, err
}

func (mr *MachinesRepository) UpdateOne(id uint32, machine *MachinesEntity) (*MachinesEntity, error) {
	var result *MachinesEntity
	query := "UPDATE machines SET name=$1, cpu_count=$2, total_disk_space=$3 WHERE id=$4"
	_, err := mr.db.Exec(query, machine.Name, machine.CpuCount, machine.TotalDiskSpace, id)
	if err == nil {
		result, err = mr.GetOne(id)
	}
	return result, err
}

func NewMachinesRepository(db *sql.DB) *MachinesRepository {
	return &MachinesRepository{db: db}
}
