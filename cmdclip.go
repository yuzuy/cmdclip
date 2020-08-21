package cmdclip

import (
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
)

func Clip(cmd string, args ...string) (string, error) {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return "", err
	}
	v := strings.TrimRight(string(out), "\n")
	if err := clipboard.WriteAll(v); err != nil {
		return "", err
	}

	return v, nil
}
