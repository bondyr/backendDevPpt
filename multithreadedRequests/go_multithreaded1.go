package main

import (
	"common/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const url = "http://localhost:3000"

var wg sync.WaitGroup

func makeRequest(shouldDisplay bool) {
  defer wg.Done()

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
    wg.Add(1)
    go makeRequest(shouldDisplay)
  }

  wg.Wait()
  
  fmt.Printf("\n=== %d responses received.\n", nrOfRequests)
}