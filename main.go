package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Test-Plugin!\nArgs: '%+v'\nEnv: '%+v'", os.Args, os.Environ())
	os.Exit(0)
}
