package shared

import (
	"os"
	"os/exec"
)

func ClearConsole() {
	cmd := exec.Command("clear") // Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
