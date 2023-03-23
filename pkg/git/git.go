package git

import "os/exec"

// GitInit инициализирует локальный репозиторий.
func GitInit() error {
	cmd := exec.Command("git", "init")
	_, err := cmd.Output()
	if err != nil {
		return err
	}

	return nil
}
