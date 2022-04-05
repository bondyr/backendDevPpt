package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func doSth() {
  defer wg.Done()

  // expensive calculations
  i := 3
  i *= 234423423
}

const nrOfThreads =          100 * 1000 
// const nrOfThreads = 10 * 1000 * 1000

func main() {
  fmt.Printf("Starting work with %d threads...\n", nrOfThreads)

  for i := 0; i < nrOfThreads; i++ {
    wg.Add(1)
    go doSth()
  }

  wg.Wait()

  fmt.Print("Working done\n")
}
