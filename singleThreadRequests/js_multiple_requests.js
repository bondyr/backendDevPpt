const httpClient = require('superagent')
const { processCmdLineArguments } = require('../common/common')

const { shouldDisplay, nrOfRequests } = processCmdLineArguments()

const url = 'localhost:3000'
// -------------------------------------------------------


console.log(`\n=== Waiting for ${nrOfRequests} http responses......\n`)

for (let i = 0; i < nrOfRequests; ++i) {
  httpClient.get(url).then( (body) => {
    if (shouldDisplay) {
      console.log('Response: ', body.text)
    }
  })
}