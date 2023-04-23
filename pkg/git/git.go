package git

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/blackmarllboro/create-project-struct/internal/pkg/args/interfaces"
)

func CreateLocalGitRepository(projectName interfaces.GetProjectName) error {
	name, isCurrentDir, err := projectName.GetProjectName()
	if err != nil {
		return fmt.Errorf("failed to get project name, err: %v", err)
	}

	if !isCurrentDir {
		if err := os.Chdir(fmt.Sprintf("./%s", name)); err != nil {
			return fmt.Errorf("failed to change directory, err: %v", err)
		}
	}

	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to init git repository, err: %v", err)
	}

	return nil
}
