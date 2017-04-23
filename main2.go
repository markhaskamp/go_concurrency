package main

import (
  "fmt"
  "math/rand"
  "sync"
  "time"
)


func main() {
  rand.Seed(time.Now().Unix())
  concurrentWorkers := 5
  startTime := time.Now()

  c := make(chan int)
  printerWG := sync.WaitGroup{}

  for i := 0; i<concurrentWorkers; i++ {
    go func() {
      for {
        n := <- c
        printer(n, &printerWG)
      }
    }()
  }

  for i:=0; i<25; i++ {
    printerWG.Add(1)
    c <- i
  }

  printerWG.Wait()
  fmt.Printf("elapsed time: %v\n", time.Since(startTime))
}

func printer(n int, printerWG *sync.WaitGroup) {
  sleepMilliseconds := rand.Intn(2000)
  time.Sleep(time.Duration(sleepMilliseconds) * time.Millisecond)
  fmt.Printf("%v \t%d\n", sleepMilliseconds, n)   
  printerWG.Done()
}


