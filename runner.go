package main

import (
	"github.com/alecthomas/chroma/quick"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func Run(argv []string) error {
	config, argv, err := ResolveConfig(argv)
	if err != nil {
		return err
	}
	if !config.Parameters.MustGet("include").BoolVal() {
		argv = append(argv, "-i")
	}
	cmd := exec.Command("curl", argv...)
	cmd.Stdin = os.Stdin

	if config.Parameters.Protocol() != "https" && config.Parameters.Protocol() != "http" {
		/* not http */
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			return err
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
		return nil
	}

	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	b, err := io.ReadAll(cmdOut)
	if err != nil {
		return err
	}
	text := string(b)
	section := strings.SplitN(strings.ReplaceAll(text, "\r\n", "\n"), "\n\n", 2)
	if config.Parameters.MustGet("include").BoolVal() {
		quick.Highlight(os.Stdout, section[0]+"\n\n", "http", "terminal16m", "monokai")
	}
	url, _ := config.Parameters.Get("url")
	resp, err := ParseResponse(string(b), url.StringVal())
	if err != nil {
		return err
	}
	for resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound {
		/* redirect */
		section = strings.SplitN(strings.ReplaceAll(section[1], "\r\n", "\n"), "\n\n", 2)
		if config.Parameters.MustGet("include").BoolVal() {
			quick.Highlight(os.Stdout, section[0]+"\n\n", "http", "terminal16m", "monokai")
		}
		resp, err = ParseResponse(string(b), url.StringVal())
		if err != nil {
			return err
		}
	}
	typ := getFormatType(resp.Header.Get("Content-Type"), "")
	quick.Highlight(os.Stdout, section[1], typ, "terminal16m", "monokai")
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
