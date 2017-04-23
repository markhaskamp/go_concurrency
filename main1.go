package main


import (
  "fmt"
)

func main() {

  c := genNumbers(10)
  out := squareNumbers(c)

  fmt.Println("got to here")

  for i := 0; i<10; i++ {
    fmt.Println(<-out)
  }
}

// generate N numbers,
//  send into a channel,
//  return the channel
func genNumbers(N int) <-chan int {
  out := make(chan int)

  go func() {
    for i := 0; i<N; i++ {
      out <- i 
    }
    close(out)
  }()

  return out
}

// read N numbers from passed in channel,
//  square them,
//  send into a channel
//  return the channel
func squareNumbers(in <- chan int) <-chan int {
  out := make(chan int)
  go func() {
    for n := range in {
      out <- n * n
    }
    close(out)
  }()
  return out
}

