package main

import (
	"fmt"
)

type example1 struct {
 	flag bool
 	counter int16
 	pi float32
}


type example2 struct {
 	flag bool
 	counter int16
 	pi float32
}


func main() {

e1 := example1{
 	flag: false,
 	counter: 2,
 	pi: 3.14,
} 
fmt.Printf("%+v\n", e1)

e2 := example2{
 	flag: false,
 	counter: 2,
 	pi: 3.15,
} 

fmt.Printf("%+v\n", e2)
// ./prog.go:37:3: cannot use e2 (type example2) as type example1 in assignment
// e1=e2 // e1 is not equal e2

e1=example1(e2)
fmt.Printf("e1: %+v\n", e1)

var e3 struct {
 	flag 	bool
 	counter int16
 	pi 	float32
}


// Assigment allowed  duet e3 was initialized with unnamed type
e1=e3
fmt.Printf("e1 with e3 assigment: %+v\n", e1)
