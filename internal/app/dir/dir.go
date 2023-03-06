package dir

import (
	"errors"
	"fmt"
	"os"

	"github.com/blackmarllboro/create-project-struct/internal"
	"github.com/blackmarllboro/create-project-struct/internal/app/file"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/args"
)

const perm = 0755 // Право доступа для создания папок

// Константы с именем подкаталогов проекта
const (
	cmdDir      = "cmd"
	pkgDir      = "pkg"
	internalDir = "internal"
	appDir      = "app"
	cfgDir      = "config"
)

type Dirs struct {
	projectName string
	file        *file.File
}

func NewDirs(log internal.Logger) *Dirs {
	return &Dirs{file: file.NewFile(log)}
}

// CreateProject создает проект.
func (d *Dirs) CreateProject() error {
	if err := d.createProjectDir(); err != nil {
		return err
	}

	if err := d.createProjectDirs(); err != nil {
		return err
	}

	return nil
}

// CreateProjectDir создает корневую папку проекта.
func (d *Dirs) createProjectDir() error {
	projectDir, currentDir, err := args.GetProjectName()
	d.projectName = projectDir
	if err != nil {
		return err
	}

	if !currentDir {
		if err := os.Mkdir(d.projectName, perm); err != nil {
			return errors.New("такая директория уже существует")
		}

		d.file.Logger.Info("корневая директория проекта успешно создана")
	}

	return nil
}

// logCreteDir выводит лог о создании переданной директории.
func (d *Dirs) logCreteDir(dir string) {
	d.file.Logger.Info("директория ./" + dir + " успешно создана")
}

// createProjectDirs создает структуру папок в директории с проектом.
func (d *Dirs) createProjectDirs() error {
	projectDirs := [4]string{cmdDir, pkgDir, internalDir, cfgDir}
	for i := 0; i < len(projectDirs); i++ {
		currentDir := projectDirs[i]

		dir := fmt.Sprintf("%s/%s", d.projectName, currentDir)
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

		d.logCreteDir(currentDir)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := d.file.GenerateGoModFile(currentDir + "/" + d.projectName); err != nil {
		return err
	}

	return nil
}

// createInternalSubDir создает подкаталоги "app" и pkg в директории internal.
func (d *Dirs) createInternalSubDir() error {
	internalSubDirs := [2]string{pkgDir, appDir}

	for i := 0; i < len(internalSubDirs); i++ {
		currentDir := internalSubDirs[i]

		createSubDirPath := d.projectName + "/" + internalDir + "/" + currentDir

		if err := os.Mkdir(createSubDirPath, perm); err != nil {
			return err
		}

		d.logCreteDir(internalDir + "/" + currentDir)
	}

	return nil
}
