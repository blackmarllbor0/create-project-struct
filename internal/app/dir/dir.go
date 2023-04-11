package dir

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/blackmarllboro/create-project-struct/internal/app/file"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/args"
)

const perm = 0755 // Access rights to create folders

// Constants with the name of the project subdirectories
const (
	cmdDir      = "cmd"
	pkgDir      = "pkg"
	internalDir = "internal"
	appDir      = "app"
	cfgDir      = "config"
)

type Dirs struct {
	projectName string
	currentDir  bool
	file        *file.File
}

func NewDirs() *Dirs {
	return &Dirs{file: file.NewFile()}
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
	projectDir, currentDir, err := args.GetProjectName()
	d.projectName = path.Base(projectDir)
	d.currentDir = currentDir
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
	projectDirs := [4]string{cmdDir, pkgDir, internalDir, cfgDir}
	for i := 0; i < len(projectDirs); i++ {
		currentDir := projectDirs[i]

		var dir string
		if d.currentDir {
			dir = currentDir
		} else {
			dir = fmt.Sprintf("%s/%s", d.projectName, currentDir)
		}

		if err := os.Mkdir(dir, perm); err != nil {
			return err
		}

		// в зависимости от текущего создаваемого каталога создаём файлы или подкаталоги.
		if currentDir == cmdDir {
			if err := d.file.GenerateMainFile(dir + "/" + d.projectName + ".go"); err != nil {
				return err
			}
		} else if currentDir == internalDir {
			if err := d.createInternalSubDir(); err != nil {
				return err
			}
		} else if currentDir == cfgDir {
			if err := d.file.GenerateCfgFile(dir); err != nil {
				return err
			}
		}
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := d.file.GenerateGoModFile(currentDir+"/"+d.projectName, d.currentDir); err != nil {
		return err
	}

	return nil
}

func (d *Dirs) createInternalSubDir() error {
	internalSubDirs := [2]string{pkgDir, appDir}

	for i := 0; i < len(internalSubDirs); i++ {
		currentDir := internalSubDirs[i]

		var createSubDirPath string
		if !d.currentDir {
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
