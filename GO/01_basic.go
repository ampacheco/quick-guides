package main

import (
	"fmt"
)

type example struct {
 	flag bool
 	counter int16
 	pi float32
}

func main() {

  var e1 example
  fmt.Printf("%+v\n", e1)

  e2 := example{
   flag: true,
   counter: 10,
   pi: 3.141592,
  }
  fmt.Printf("%+v\n", e2)

  e3 := struct {
   flag bool
   counter int16
   pi float32
  }{
   flag: true,
   counter: 10,
   pi: 3.141592,
  }
  fmt.Println("Flag", e3.flag)
  fmt.Println("Counter", e3.counter)
  fmt.Println("Pi", e3.pi)
}
