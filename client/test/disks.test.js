'use strict';

const mocha = require('mocha');
const { describe, it } = mocha;
const assert = require('chai').assert;

const { connectDiskToMachine } = require('../disks');

describe('Disks module', () => {
	it('should return non-null object', async () => {
		const disk = await connectDiskToMachine(1, 1);
		assert.isNotNull(disk);
	});
});

