/*
  Buffered channel
*/

package main

import (
	"common/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "http://localhost:3000"

func makeRequest(ch chan<-string) {
  resp, _ := http.Get(url)

  body, _ := ioutil.ReadAll(resp.Body)
  
  // blocking operation, waiting for consumer to be ready
  ch <- string(body)
  
  fmt.Println("Response written to channel")
}

func producer(nrOfRequests int) <-chan string {
    ch := make(chan string)
  // ch := make(chan string, nrOfRequests)

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

  // delay creating a consumer
  const sleepTime = 5
  fmt.Printf("Waiting for consumer to come up for %d seconds\n", sleepTime)
  time.Sleep(sleepTime * time.Second)
  fmt.Println("Consumer ready for reading channel")
  
  consumer(ch, nrOfRequests, shouldDisplay)


  time.Sleep(time.Second)
  fmt.Printf("\n=== %d responses received.\n", nrOfRequests)
}
