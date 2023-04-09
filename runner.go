package curlcolor

import (
	"os"
	"os/exec"
)

func Run(argv []string) error {
	mgr, argv, err := ResolveManager(argv)
	if err != nil {
		return err
	}
	if !mgr.CurlParameter.GetBool("include") {
		argv = append(argv, "--include")
	}
	cmd := exec.Command(mgr.CurlCmd, argv...)
	cmd.Stdin = os.Stdin

	scheme := GetUrlScheme(mgr.CurlParameter.GetString("url"))
	if scheme != "https" && scheme != "http" {
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
	cmdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	if err = PrintError(mgr, cmdErr); err != nil {
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
