package args

import (
	"errors"
	"fmt"
	"os"
)

type Args struct{}

func NewArgs() *Args {
	return &Args{}
}

// GetProjectName function for getting the project name. The name is specified
// via a command line argument. Set the argument "." to create a project in the
// current directory. Boolean value returns true if the application is created in
// the current directory.
func (p Args) GetProjectName() (string, bool, error) {
	if len(os.Args) < 2 {
		return "", false, errors.New("the project name has not been transferred")
	}

	projectName := os.Args[1]

	if projectName == "." {
		pwd, err := os.Getwd()
		if err != nil {
			return "", false, fmt.Errorf("failed to get the current directory, err: %s", err)
		}

		return pwd, true, nil
	}

	return projectName, false, nil
}
