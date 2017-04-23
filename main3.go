package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)


func main() {
  rand.Seed(time.Now().Unix())
  p1Workers := 10
  printerWorkers := 2
  startTime := time.Now()

  p1Channel := make(chan int)
  printerChannel := make(chan int)
  wg := sync.WaitGroup{}

  // start up p1 workers
  startP1Workers(p1Workers, p1Channel, printerChannel)

  // start up printer workers
  for i := 0; i<printerWorkers; i++ {
    go func() {
      for {
        n := <- printerChannel
        printer(n, &wg)
      }
    }()
  }

  // original input is generated and sent to 1st channel
  for i:=0; i<25; i++ {
    wg.Add(1)
    p1Channel <- i
  }

  wg.Wait()
  fmt.Printf("elapsed time: %v\n", time.Since(startTime))
}

func startP1Workers(p1Workers int,
                    p1Channel chan int,
                    printerChannel chan<- int) {

  for i := 0; i<p1Workers; i++ {
    go func() {
      for {
        n := <- p1Channel
        processor1(n, printerChannel)
      }
    }()
  }
}

func processor1(n int, printerChannel chan<- int) {
  sleepMilliseconds := rand.Intn(2000)
  time.Sleep(time.Duration(sleepMilliseconds) * time.Millisecond)

  printerChannel <- n
}

func printer(n int, wg *sync.WaitGroup) {
  fmt.Println(n)
  wg.Done()
}


