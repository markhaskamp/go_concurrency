package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)


func main() {
  rand.Seed(time.Now().Unix())
  p1Workers := 4
  printerWorkers := 4
  startTime := time.Now()

  p1Channel := make(chan int)
  printerChannel := make(chan int)
  wg := sync.WaitGroup{}

  startP1Workers(p1Workers, &wg, p1Channel, printerChannel)
  startPrinterWorkers(printerWorkers, printerChannel, &wg)

  // original input is generated and sent to 1st channel
  for i:=0; i<25; i++ {
    wg.Add(1)
    p1Channel <- i
  }

  wg.Wait()
  fmt.Printf("elapsed time: %v\n", time.Since(startTime))
}

func startP1Workers(p1Workers int,
                    wg *sync.WaitGroup,
                    p1Channel chan int,
                    printerChannel chan<- int) {

  for i := 0; i<p1Workers; i++ {
    go func() {
      for {
        n := <- p1Channel
        processor1(n, wg, printerChannel)
      }
    }()
  }
}

func startPrinterWorkers(printerWorkers int,
                         printerChannel chan int,
                         wg *sync.WaitGroup) {

  for i := 0; i<printerWorkers; i++ {
    go func() {
      for {
        n := <- printerChannel
        printer(n, wg)
      }
    }()
  }
}

func processor1(n int,
                wg *sync.WaitGroup,
                printerChannel chan<- int) {

  if (n % 2) == 0 {
    wg.Done()
    return
  }

  sleepMilliseconds := rand.Intn(2000)
  time.Sleep(time.Duration(sleepMilliseconds) * time.Millisecond)

  printerChannel <- n
}

func printer(n int, wg *sync.WaitGroup) {
  fmt.Printf("%d ", n)
  wg.Done()
}


