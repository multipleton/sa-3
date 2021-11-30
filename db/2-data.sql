INSERT INTO machines (id, name, cpu_count, total_disk_space) VALUES
    (1, 'Ubuntu VM', 4, 17179869184),
    (2, 'Windows VM', 2, 21474836480);

INSERT INTO disks (id, machine_id, capacity) VALUES
    (1, 1, 17179869184),
    (2, 2, 17179869184),
    (3, 2, 4294967296),
    (4, NULL, 8589934592);
