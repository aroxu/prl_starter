package main

import (
	"os"
	"os/exec"
)

func main() {
	app := "bash"
	cmd := exec.Command(app, "-c", "/Applications/Parallels\\ Desktop.app/Contents/MacOS/prl_client_app")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "SYSTEM_VERSION_COMPAT=1")
	cmd.Start()
}
