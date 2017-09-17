package main

import (
	"fmt"
	"flag"
	"os"
)

var fatalErr error

func fatal(e error)  {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}
func main() {
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()
}
