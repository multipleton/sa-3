const http = require('http');
const { PORT, HOST } = require('../config');

const connectDiskToMachine = (diskId, machineId) => {
  return new Promise((resolve, reject) => {
    const options = {
      hostname: HOST,
      port: PORT,
      path: `/disks/${diskId}`,
      method: 'PATCH',
    };
    const bodyData = JSON.stringify({ machineId });
    const req = http.request(options, res => {
      const chunks = [];
      res.on('data', chunk => {
        chunks.push(chunk);
      });
      res.on('end', () => {
        resolve(JSON.parse(chunks.join()));
      });
    });
    req.on('error', err => {
      reject(err);
    });
    req.write(bodyData);
    req.end();
  });
};

module.exports = { connectDiskToMachine };
