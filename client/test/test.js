'use strict';

const mocha = require('mocha');
const describe = mocha.describe;
const it = mocha.it;
const assert = require('chai').assert;

const { getAllMachines } = require('../machines');
const { connectDiskToMachine } = require('../disks');

describe('Machines module', function() {
	it('should return array', async function() {
		const machines = await getAllMachines();
		assert.strictEqual(Array.isArray(machines), true);
	});

	it('should return array with non-zero length', async function() {
		const machines = await getAllMachines();
		assert.strictEqual(machines.length > 0, true);
   });
});

describe('Disks module', function() {
	it('should return non-null object', async function() {
		const disk = await connectDiskToMachine(1, 1);
		assert.notEqual(disk, null);
	});
});

