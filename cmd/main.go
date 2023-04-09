package main

import (
	"github.com/if-nil/curlcolor"
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	if err := curlcolor.Run(os.Args[1:]); err != nil {
		slog.Error("error: %v", err)
	}
}
