const http = require('http');
const { PORT, HOST } = require('../config');

const getAllMachines = () => {
  return new Promise((resolve, reject) => {
    const options = {
      hostname: HOST,
      port: PORT,
      path: '/machines',
      method: 'GET',
    };
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
    req.end();
  });
};

module.exports = { getAllMachines };
