package git

import (
	"os/exec"
)

func CreateLocalGitRepository() error {
	cmd := exec.Command("git", "init")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
