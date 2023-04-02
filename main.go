package main

import (
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	if err := Run(os.Args[1:]); err != nil {
		slog.Error("error: %v", err)
	}
}
