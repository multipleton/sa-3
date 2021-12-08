'use strict';

const { getAllMachines } = require('./machines');
const { connectDiskToMachine } = require('./disks');

async function diskModuleUsage() {
	const disk = await connectDiskToMachine(1, 1);
	console.log('Usage of disk module:');
	console.log(disk);
}

async function machineModuleUsage() {
	const machines = await getAllMachines();
	console.log('Usage of machine module:');
	console.log(machines);
}

// Usage of modules

diskModuleUsage();
machineModuleUsage();

