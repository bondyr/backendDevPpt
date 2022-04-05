package main

import (
	"common/common"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:3000"

func makeRequest(ch chan<- string) {
  resp, _ := http.Get(url)

  body, _ := ioutil.ReadAll(resp.Body)

  ch <- string(body)
}

func main() {
  shouldDisplay, nrOfRequests := common.ProcessCmdLineArguments()
  
  fmt.Printf("\n=== Waiting for %d responses.....\n\n", nrOfRequests)

  ch := make(chan string)

  for i := 0; i < nrOfRequests; i++ {
    go makeRequest(ch)
  }

  for i := 0; i < nrOfRequests; i++ {
    val := <-ch
    if shouldDisplay {
      fmt.Println(val)
    }
  }

  fmt.Printf("\n=== %d responses received.\n", nrOfRequests)
}

// unbiffered channels are blocking
// writer channel copies data directly to receiver's channel stack