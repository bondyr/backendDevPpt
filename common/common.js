exports.processCmdLineArguments = function () {
  if (process.argv.length !== 4) {
    console.log('Usage: node js_multiple_requests.js shouldDisplay nrOfRequests\n', 'Example:\n', '    node js_multiple_requests.js 1 100\n')
    throw "Incorrect nr of args"
  }

  const myArgs = process.argv.slice(2)
  const shouldDisplay = myArgs[0] === '1'
  const nrOfRequests = myArgs[1]

  return {
    shouldDisplay,
    nrOfRequests
  }
}