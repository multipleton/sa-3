package machines

import (
	"database/sql"
	"fmt"
)

type MachinesRepository struct {
	Db *sql.DB
}

func (mr *MachinesRepository) GetAll() ([]Machine, error) {
	rows, err = r.DB.Query("SELECT id, name, cpu_count, total_disk_space FROM machines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Machine
	for rows.Next() {
		var m Machine
		if err := rows.Scan(&m.Id, &m.Name, &m.CpuCount, &m.TotalDiskSpace); err != nil {
			return nil, err
		}
		res = append(res, &m)
	}
	if res == nil {
		res = make([]*Machine, 0)
	}
	return res, nil
}

func (mr *MachinesRepository) GetOne(id string) (Machine, error) {
	query := fmt.Sprintf("SELECT id, name, cpu_count, total_disk_space FROM machines WHERE id=%s", id)
	row, err = r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var m Machine
	if err := row.Scan(&m.Id, &m.Name, &m.CpuCount, &m.TotalDiskSpace); err != nil {
		return nil, err
	}
	res = append(res, &m)
	return res, nil
}

func (mr *MachinesRepository) UpdateOne(id string, machine Machine) (Machine, error) {
	query := fmt.Sprintf("UPDATE machines SET total_disk_space=%s WHERE id=%s RETURNING id, name, cpu_count, total_disk_space", machine.TotalDiskSpace, id)
	row, err = r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var m Machine
	if err := row.Scan(&m.Id, &m.Name, &m.CpuCount, &m.TotalDiskSpace); err != nil {
		return nil, err
	}
	res = append(res, &m)
	return res, nil
}
