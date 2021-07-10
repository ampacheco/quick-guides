package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func escapeToHeap() *user {
	u := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	return &u
}

func main() {

	u := escapeToHeap()
	fmt.Printf("User: %+v", u)
}
