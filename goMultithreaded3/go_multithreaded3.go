/*
  Go idiom - Consumer-Producer:
  Owner of channel is a function that instantiates, writes and closes the channel.
*/

package main

import (
	"common/common"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://localhost:3000"

func makeRequest(ch chan<-string) {
  resp, _ := http.Get(url)

  body, _ := ioutil.ReadAll(resp.Body)

  ch <- string(body)
}

func producer(nrOfRequests int) <-chan string {
  ch := make(chan string)

  for i := 0; i < nrOfRequests; i++ {
    go makeRequest(ch)
  }

  return ch
}

func consumer(ch <-chan string, nrOfRequests int, shouldDisplay bool) {
  for i := 0; i < nrOfRequests; i++ {
    val := <-ch
    if shouldDisplay {
      fmt.Println(val)
    }
  }
}

func main() {
  shouldDisplay, nrOfRequests := common.ProcessCmdLineArguments()
  
  fmt.Printf("\n=== Waiting for %d responses.....\n\n", nrOfRequests)

  ch := producer(nrOfRequests)
  consumer(ch, nrOfRequests, shouldDisplay)

  fmt.Printf("\n=== %d responses received.\n", nrOfRequests)
}