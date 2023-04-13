package dir

import (
	"errors"
	"fmt"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/args"
	"os"
	"path"

	"github.com/blackmarllboro/create-project-struct/internal/app/file"
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
	file         *file.File
	name         args.GetProjectName
}

func NewDirs(f *file.File, p args.ProjectName) *Dirs {
	return &Dirs{file: f, name: p}
}

func (d *Dirs) CreateProject() error {
	if err := d.createProjectDir(); err != nil {
		return err
	}

	if err := d.createProjectDirs(); err != nil {
		return err
	}

	return nil
}

func (d *Dirs) createProjectDir() error {
	projectDir, currentDir, err := d.name.GetProjectName()
	d.projectName = path.Base(projectDir)
	d.isCurrentDir = currentDir
	if err != nil {
		return err
	}

	if !currentDir {
		if err := os.Mkdir(d.projectName, perm); err != nil {
			return errors.New("this directory already exists")
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
			return err
		}

		// depending on the current directory being created, create files or subdirectories.
		if err := d.createFilesInSubdirs(currentDir, dir); err != nil {
			return err
		}
	}

	if err := d.file.GenerateFilesInMainDir(d.projectName, d.isCurrentDir); err != nil {
		return err
	}

	return nil
}

func (d *Dirs) createInternalSubDir() error {
	internalSubDirs := [2]string{pkgDir, appDir}

	for i := 0; i < len(internalSubDirs); i++ {
		currentDir := internalSubDirs[i]

		var createSubDirPath string
		if !d.isCurrentDir {
			createSubDirPath = d.projectName + "/" + internalDir + "/" + currentDir
		} else {
			createSubDirPath = internalDir + "/" + currentDir
		}

		if err := os.Mkdir(createSubDirPath, perm); err != nil {
			return err
		}
	}

	return nil
}

func (d *Dirs) createFilesInSubdirs(currentDir, dir string) error {
	switch currentDir {
	case cmdDir:
		if err := d.file.GenerateMainFile(dir + fmt.Sprintf("/%s", d.projectName) + ".go"); err != nil {
			return err
		}
	case internalDir:
		if err := d.createInternalSubDir(); err != nil {
			return err
		}
	case cfgDir:
		if err := d.file.GenerateCfgFile(dir); err != nil {
			return err
		}
	}

	return nil
}
