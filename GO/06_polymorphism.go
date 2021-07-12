package main

import (
	"fmt"
	"testing"
)

type pasear interface {
	pasear()		
}

type historicalPark struct {
	name string
}


func (hp historicalPark) pasear() {
	fmt.Printf("\nAprendo de la historia de USA en %s Park", hp.name)
}
 


type naturalPark struct {
	name string
} 

func (np naturalPark) pasear() {
	fmt.Printf("\nAprendo de la naturaleza de USA en %s Park", np.name)
}


func TestLastIndex(t *testing.T) {

	np := naturalPark {"Yosemite"}
	hp := historicalPark{"Independence National"}
	
	
	np.pasear()
	hp.pasear()
}
