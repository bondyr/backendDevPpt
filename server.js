const http = require('http');

const hostname = '127.0.0.1';
const port = 3000;

let i = 0

const myArgs = process.argv.slice(2)
const sleepTimeInMs = myArgs[0] | 50

const server = http.createServer(async (req, res) => {
  await sleep(sleepTimeInMs)
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end(`Hello World ${i++}`);
});

async function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});
