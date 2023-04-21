package curlcolor

import (
	"os"
	"os/exec"
)

// TODO: add options and more help info
const help = "\033[34mcurlor\033[0m is a wrapper for curl, it will print the response body with color."

func Run(argv []string, version string) error {
	mgr, argv, err := ResolveManager(argv)
	if err != nil {
		return err
	}
	if mgr.Version {
		mgr.Printer.Highlight([]byte("curlor version: "+version), "yaml")
		return nil
	}
	if mgr.Help {
		mgr.Printer.Print([]byte(help))
		return nil
	}
	if !mgr.CurlParameter.GetBool("include") {
		argv = append(argv, "--include")
	}
	cmd := exec.Command(mgr.CurlCmd, argv...)
	cmd.Stdin = os.Stdin

	scheme := GetUrlScheme(mgr.CurlParameter.GetString("url"))
	if scheme != "https" && scheme != "http" || mgr.CurlParameter.GetString("output") != "" ||
		mgr.CurlParameter.GetBool("version") || mgr.CurlParameter.GetString("help") != "" {
		/* not http or output to file */
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
	cmdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	if err = ParseAndPrintError(mgr, cmdErr); err != nil {
		return err
	}
	if err = ParseAndPrintOutput(mgr, cmdOut); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
