package osbox

import (
	"log"
	"os/exec"
	"strings"
)

// ExecCmd executes the given command
func ExecCmd(c string, args ...string) string {
	cmd := exec.Command(c, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Println("failed to exec cmd:", err)
	}
	if len(out) == 0 {
		return ""
	}
	s := string(out)
	return strings.ReplaceAll(s, "\n", "")
}
