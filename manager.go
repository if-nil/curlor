package curlcolor

import (
	"fmt"
	"os"
)

const (
	defaultTheme = "monokai"
	defaultCmd   = "curl"
)

type Manager struct {
	CurlParameter CurlParameter
	Printer       ColorPrinter
	Debug         bool
	CurlCmd       string
	Version       bool
}

func TraversalSearchBool(key string, argv []string, default2 bool) (bool, []string, error) {
	for i, v := range argv {
		if v == key {
			return true, append(argv[:i], argv[i+1:]...), nil
		}
	}
	return default2, argv, nil
}

func TraversalSearchString(key string, argv []string, default2 string) (string, []string, error) {
	for i, v := range argv {
		if v == key {
			if i+1 < len(argv) {
				return "", nil, fmt.Errorf("invalid parameter")
			}
			return argv[i+1], append(argv[:i], argv[i+2:]...), nil
		}
	}
	return default2, argv, nil
}

func ResolveManager(args []string) (*Manager, []string, error) {
	theme, args, err := TraversalSearchString("---theme", args, defaultTheme)
	curlCmd, args, err := TraversalSearchString("---cmd", args, defaultCmd)
	version, args, err := TraversalSearchBool("---version", args, false)
	debug, args, err := TraversalSearchBool("---debug", args, false)
	outWriter := os.Stdout
	errWriter := os.Stderr
	if err != nil {
		return nil, nil, err
	}
	parameter, err := ParseArgs(args)
	if err != nil {
		return nil, nil, err
	}
	printer := ColorPrinter{
		OutWriter: outWriter,
		ErrWriter: errWriter,
		Formatter: "terminal256",
		Theme:     theme,
	}
	return &Manager{
		CurlParameter: parameter,
		CurlCmd:       curlCmd,
		Version:       version,
		Debug:         debug,
		Printer:       printer,
	}, args, nil
}
