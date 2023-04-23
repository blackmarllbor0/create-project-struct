package dir

import (
	"fmt"
	"os"
	"path"

	fileI "github.com/blackmarllboro/create-project-struct/internal/app/file/interfaces"
	argsI "github.com/blackmarllboro/create-project-struct/internal/pkg/args/interfaces"
)

const perm = 0755 // Access rights to create folders

// Constants with the name of the project subdirectories.
const (
	cmdDir      = "cmd"
	pkgDir      = "pkg"
	internalDir = "internal"
	appDir      = "app"
	cfgDir      = "config"
)

type Dirs struct {
	projectName  string
	isCurrentDir bool

	file fileI.File
	name argsI.GetProjectName
}

func NewDirs(file fileI.File, name argsI.GetProjectName) *Dirs {
	return &Dirs{file: file, name: name}
}

func (d *Dirs) CreateProject() error {
	if err := d.createProjectDir(); err != nil {
		return fmt.Errorf("failed to create project directory: %v", err)
	}

	if err := d.createProjectDirs(); err != nil {
		return fmt.Errorf("failed to create project directories: %v", err)
	}

	return nil
}

func (d *Dirs) createProjectDir() error {
	projectDir, currentDir, err := d.name.GetProjectName()
	d.projectName = path.Base(projectDir)
	d.isCurrentDir = currentDir
	if err != nil {
		return fmt.Errorf("failed to get project name, err: %s", err)
	}

	if !currentDir {
		if err := os.Mkdir(d.projectName, perm); err != nil {
			return fmt.Errorf("failed to create dir, err: %v", err)
		}
	}

	return nil
}

func (d *Dirs) createProjectDirs() error {
	projectDirs := [3]string{cmdDir, internalDir, cfgDir}

	for i := 0; i < len(projectDirs); i++ {
		currentDir := projectDirs[i]

		var dir string
		if d.isCurrentDir {
			dir = currentDir
		} else {
			dir = fmt.Sprintf("%s/%s", d.projectName, currentDir)
		}

		if err := os.Mkdir(dir, perm); err != nil {
			return fmt.Errorf("failed to create dir, err: %v", err)
		}

	}

	if err := d.createInternalSubDir(); err != nil {
		return fmt.Errorf("failed to create internal sub dir, err: %v", err)
	}

	if err := d.file.GenerateFiles(d.isCurrentDir, d.projectName); err != nil {
		return fmt.Errorf("failed to generate files, err: %v", err)
	}

	return nil
}

func (d *Dirs) createInternalSubDir() error {
	internalSubDirs := [2]string{pkgDir, appDir}

	for i := 0; i < len(internalSubDirs); i++ {
		currentDir := internalSubDirs[i]

		var createSubDirPath string
		if !d.isCurrentDir {
			createSubDirPath = fmt.Sprintf("%s/%s/%s", d.projectName, internalDir, currentDir)
		} else {
			createSubDirPath = fmt.Sprintf("%s/%s", internalDir, currentDir)
		}

		if err := os.Mkdir(createSubDirPath, perm); err != nil {
			return fmt.Errorf("failed to create dir, err: %v", err)
		}
	}

	return nil
}
