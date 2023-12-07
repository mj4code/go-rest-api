package main

import (
	"fmt"
)

// Going to be responsible for instantiation
// and startup of go application
func Run() error {
	fmt.Println("Starting the application")
	return nil
}

func main() {
	fmt.Println("Go RESTful API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
