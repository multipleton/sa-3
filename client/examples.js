'use strict';

const { getAllMachines } = require('./machines');
const { connectDiskToMachine } = require('./disks');

const diskModuleUsage = async () => {
  const disk = await connectDiskToMachine(1, 2);
  console.log('Usage of disk module:');
  console.log(disk);
};

const machineModuleUsage = async () => {
  const machines = await getAllMachines();
  console.log('Usage of machine module:');
  console.log(machines);
};

// Usage of modules
(async () => {
  await diskModuleUsage();
  await machineModuleUsage();
})();
