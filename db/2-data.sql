INSERT INTO machines (name, cpu_count, total_disk_space) VALUES
  ('Ubuntu VM', 4, 17179869184),
  ('Windows VM', 2, 21474836480);

INSERT INTO disks (machine_id, capacity) VALUES
  (1, 17179869184),
  (2, 17179869184),
  (2, 4294967296),
  (NULL, 8589934592);
