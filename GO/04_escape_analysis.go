package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func stayOnStack() user {
	u := user{
		name:  "Bill",
		email: "bill@email.com",
	}
	return u
}

func main() {
	u := stayOnStack()
	fmt.Printf("User: %+v", u)
}
