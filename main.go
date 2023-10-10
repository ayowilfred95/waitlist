package main

import (
	"github.com/ayowilfred95/api"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	// init Api
	err := api.Api()
	if err != nil {
		return err
	}
	return nil
}
