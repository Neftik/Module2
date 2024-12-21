package main

import (
	"fmt"

	"github.com/Andreyka-coder9192/calc_go/internal/application"
)

func main() {
	app := application.New()
	fmt.Println("RunServer")
	// app.Run()
	app.RunServer()
}
