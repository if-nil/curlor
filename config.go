package main

import (
	"fmt"
)

const (
	defaultTheme = "monokai"
	defaultCmd   = "curl"
)

type Config struct {
	Theme      string
	Parameters Parameter
	CurlCmd    string
	Version    bool
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

func ResolveConfig(argv []string) (*Config, []string, error) {
	theme, argv, err := TraversalSearchString("---theme", argv, defaultTheme)
	curlCmd, argv, err := TraversalSearchString("---curl", argv, defaultCmd)
	version, argv, err := TraversalSearchBool("---version", argv, false)
	if err != nil {
		return nil, nil, err
	}
	parameter, err := ParseArgs(argv)
	if err != nil {
		return nil, nil, err
	}
	return &Config{
		Theme:      theme,
		Parameters: parameter,
		CurlCmd:    curlCmd,
		Version:    version,
	}, argv, nil
}
