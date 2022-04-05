package main

import (
	"common/common"
	"fmt"
	"time"
)

func main() {
  // cmd line - go run, build, test
 
  var a int     // declaration + initialization
  a = 1

  var b = 2     // type implicitly inferred
  
  c := 3        // shorthand syntax

  const constA = 11
  const constB = "Hello"
  const constC = 'W'
  const constD = true

  if a > 0 {
    fmt.Println("Printing inside if")
  }

  for i := 0; i < b; i++ {
    fmt.Printf("This is a formated string c alike with %d\n", i)
  }

  // no while loop!
  // for {
  //   fmt.Println("This is an infinite loop")    // while equivalent
  // }

  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      fmt.Println("It's the weekend")
    default:
      fmt.Println("It's a weekday")
  }

  // arrays - std::array equivalent
  var tab [2]int
  tab[0] = 11
  tab[1] = 22

  tab2 := [3]string{"aa", "bb", "cc"}

  // iterating over array
  for i, v := range tab {
    fmt.Println(i, v)
  }

  // slices - std::vector equivalent
  mySlice2 := []string{"a", "b", "c", "d", "e"}
  fmt.Println(mySlice2[2:4])
  
  mySlice := make([]string, 1)
  mySlice[0] = "zeroElem"
  fmt.Println(mySlice)
  mySlice = append(mySlice, "firstElem")
  mySlice = append(mySlice, "secondElem")
  mySlice = append(mySlice, "thirdElem")
  fmt.Println(mySlice[1:])

  // hashmaps
  m := make(map[string]int)
  m["k1"] = 7
  m["k2"] = 13
  fmt.Println("map:", m)

  // functions
  res1, res2 := calc(11, 22)
  fmt.Println(res1, res2)

  _, err := common.Sqrt(-0.5)
  if err != nil {
    fmt.Println(err)
  } else {
    // do sth
  }

  // suppress 'unsused variable' error
  _ = c
  _ = tab2

  // structs	
  type person struct {
    name string
    age  int
  }
  per1 := person{"R. Potocki", 25}
  per2 := person{age: 33, name: "S. Lem"}
  fmt.Println(per1, per2)

  // pointers
  p := &per1
  p.age = 100
  fmt.Printf("%p\n", p) 
  fmt.Println(p)
  fmt.Println(*p)

  // receivers
  myCar := car{"opel", 100000}
  myCar.sayHello()
  myPlane := plane{10000, 4}
  myPlane.sayHello()

    // interfaces
  introduceYourself(&myCar)
  introduceYourself(&myPlane)
  
  // post OOP paradigm
  // it's very common to have implementations first and then derive abstractions, very convenient
  // c++ - bottom-up vs golang - top-down approach

  // automatic memory management
  // Go is pass-by-value
  // automatic reference to local variables - go compiler and runtime takes care of binding variable moving it from stack to heap

  c++         // ok
  // ++c      // nope! Only postincrement available bc of lacking pointer arithmetics
  
  // testing approach - no need to install any external lib,
  // *_test files run automatically functions starting with Test*
  // by default no asserions - use if instead!
  // show example

  // goroutines in more details:
  // "Do not communicate by sharing memory; instead, share memory by communicating."
  // channels

  // other communication mechanisms - mutex, atomic, conditional_variable, once_flag

  // helgrind alike tool for races detection:
  // go test -race mypkg
  // go run -race mysrc.go
}

func calc(a int, b int) (int, int) {
  return a + b, a * b
}

  type car struct {
    manufacturer string
    mileage  int
  }

  type plane struct {
    maxAltitude int
    engines int
  }

  func (c *car) sayHello() {
    fmt.Println("Hello, my brand is", c.manufacturer, "and I ran", c.mileage, "kilometers")
  }

  func (p *plane) sayHello() {
    fmt.Println("Hello, I can fly on", p.maxAltitude, "altitude and I have", p.engines, "engines")
  }

  type introducer interface {
    sayHello()
  }

  func introduceYourself(i introducer) {
    fmt.Println(i)
    i.sayHello()
  }

// notes
/*
- ch <- val - send to channel
- <- ch - receive from channel
- chans are blocking - sending goroutine waits for receiving goroutine to be ready, waiting goroutine wait for the value to be sent
- close(ch) - no more values to be sent, sender goroutine indicates to receiver one, receiver can unblock and proceed with its other computation
- val, ok <- ch - ok is boolean meaning whether val is written by a write or by a close operation
- for value := range ch {} - iterate over values received from channel, loop breaks on channel closure.
  Sender should close channel when all values sent and received, will go out of range loop then
- unbuffered channels - are synchronous, receiver waits for sender to be ready and vice versa ch := make(chan Type)
- buffered channels - are given capacity, in-memory FIFO queue, asynchronous. ch := make(chan Type, capacity).
  Capacity defines nr of values that can be sent without receiver being ready
  Sender sends values until buffer is full, then blocks. Receiver receives values until buffer is empty, then blocks
- when using channels as fun params we can specify if channel is meant to send or receive values - increased type safety
  func foo(in <-chan string, out chan<- string) { }
           receive-only channel  send-only channel
  function       receives           sends                 values
- 'make' funtion alocates memory
- default value for channels is nil. Reading/writing on nil channel will block forever, closing will generate panic
  var ch chan interface{}
  <-ch - read blocks forever
  ch<- struct{} - write blocks forever
  close(ch) - panic
- Go idiom:
  owner of channel is goroutine that instantiates, writes and closes the channel.
  channel utilizers only have a read-only view to the channel
  it helps to avoid deadlocks and panics (writing/reading nil channel, closing nil channel, writing to closed channel, closing more than once)
- channels details
  internally a circular ring buffer with a mutex and waiting/sending list of goroutines
  channels are allocated on heap ad passed as a pointer between goroutiness
  channels seem to me as boosted cond variables (with a buffer)
  data is copied twice when using buffered channels and once when unbuffered once (sender writes value directly into receiver goroutine stack)
  no memory sharing! And access protected by mutex lock

  -------------------------------------------------------------
  Threads:
- thread - the smallest unit of execution that CPU accepts
- runs within a process, there can be many threads within a process
- threads share the same address space
- each process has at least 1 thread
- OS scheduler makes scheduling decisions at thread level, not process level
- threads can run concurrently or in parallel
- thread context switching is expensive - process context and thread context is copied
- C10k problem -  scheduler allocates a time slice for each process, thich is equally divided among threads.
  Let's say it's 10ms. FOr 1000 threads there's 10us for each;/
- threads are allocated fixed stack size - on ubuntu it's 8MB - it limits nr of threads to amount of memory we have
- threads need exclusive access to shared memory to avoid data races and undeterminitic behavior
- critical section guards are a developer convention, no tool to enforce that (but there are tools like a scoped lock to help)
- sharing memory between threads creates complexity and can lead to data races 

Goroutines:
- concurrency based on Communicating Sequential Processes paper, by Tony Hoare (1978) - very simple base ideas
  - each process is build for sequential execution - every process has a local state and it operates on it
  - no shared memory - data is communicated between processes - no race conditions or deadlocks
  - scale by adding more of the same
We can think of goroutines as user space threads, managed by go runtime (built into executable)
- goroutines are:
  - lightweight - 2KB of stack
  - low CPU overhead - three instructions per function call
  - can create hundreds of thousands of goroutines in the same address space
  - sharing of memory is avoided with usage of channels
  - contect swithing between goroutines is much cheaper than threads bc there's less state to be stores
  - goroutines run in the context of a single OS thread - from OS perspective nothing has changed
   
Dry run remarks:
does go run on a VM? No, compiled to bytecode
waitGroup behaves like a barrier
defer any function, not related to goroutine
c++ 100 000 threads - core dump check - When a thread object goes out of scope and it is in joinable state, the program is terminated. Thread destroyed in vector dtor
there's a similar concept in C++ to goroutines - boost fibers - Bartosz will present it in June
plus boost asio - async operation on any nr of threads
static assertions in go? No. But there're init time assertions
*/
