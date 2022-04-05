package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"common/common"
)

const url = "http://localhost:3000"

func makeRequest(shouldDisplay bool) {
  resp, _ := http.Get(url)

  body, _ := ioutil.ReadAll(resp.Body)
  
  if shouldDisplay {
    fmt.Println(string(body))
  }
}

func main() {
  shouldDisplay, nrOfRequests := common.ProcessCmdLineArguments()
  // -----------------------------------------------------------

  fmt.Printf("\n=== Waiting for %d responses.....\n\n", nrOfRequests)
  
  for i := 0; i < nrOfRequests; i++ {
    makeRequest(shouldDisplay)
  }
  
  fmt.Printf("\n=== %d responses received.\n", nrOfRequests)
}
