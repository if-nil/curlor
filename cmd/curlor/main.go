package main

import (
	curlcolor "github.com/if-nil/curlor"
	"os"
)

// Version this is overridden on build time
var Version = "unset"

func main() {
	if err := curlcolor.Run(os.Args[1:], Version); err != nil {
		// fmt.Printf("\033[31m%s\033[0m\n", "curlcolor: "+err.Error())
		os.Exit(1)
	}
}
