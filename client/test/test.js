'use strict';

const mocha = require('mocha');
const { describe, it } = mocha;
const assert = require('chai').assert;

const { getAllMachines } = require('../machines');
const { connectDiskToMachine } = require('../disks');

describe('Machines module', () => {
	it('should return array', async () => {
		const machines = await getAllMachines();
		assert.isArray(machines);
	});

	it('should return array with non-zero length', async () => {
		const machines = await getAllMachines();
		assert.isTrue(machines.length > 0);
   });
});

describe('Disks module', () => {
	it('should return non-null object', async () => {
		const disk = await connectDiskToMachine(1, 1);
		assert.isNotNull(disk);
	});
});

