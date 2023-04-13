package git

import (
	"fmt"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/args"
	"os"
	"os/exec"
)

func CreateLocalGitRepository(projectName args.GetProjectName) error {
	name, isCurrentDir, err := projectName.GetProjectName()
	if err != nil {
		return err
	}

	if !isCurrentDir {
		if err := os.Chdir(fmt.Sprintf("./%s", name)); err != nil {
			return err
		}
	}

	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
