package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/ionburst/namegen/namegen"
)

var appVersion string = "version"

func main() {
	var version bool = false
	var v bool = false

	flag.BoolVar(&version, "version", false, "Return namegen version")
	flag.BoolVar(&v, "v", false, "Return namegen version")
	flag.Parse()

	if version || v {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	fmt.Println(namegen.GetRandomName())
}
