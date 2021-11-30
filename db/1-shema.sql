CREATE TABLE IF NOT EXISTS machines (
  id SERIAL PRIMARY KEY,
  name VARCHAR (256) NOT NULL,
  cpu_count SMALLINT NOT NULL
    CHECK (cpu_count > 1),
  total_disk_space BIGINT NOT NULL DEFAULT 0
    CHECK (total_disk_space >= 0)
);

CREATE TABLE IF NOT EXISTS disks (
  id SERIAL PRIMARY KEY,
  machine_id INTEGER,
  capacity BIGINT NOT NULL
    CHECK (capacity > 0),
  FOREIGN KEY (machine_id) REFERENCES machines (id)
    ON DELETE SET NULL
);
