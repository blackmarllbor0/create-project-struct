package dir

import (
	"errors"
	"fmt"
	"github.com/blackmarllboro/create-project-strcut/internal"
	"github.com/blackmarllboro/create-project-strcut/internal/pkg/args"
	"os"
)

const perm = 0755 // Право доступа для создания папок

// Константы с именем подкаталогов проекта
const (
	cmdDir      = "cmd"
	pkgDir      = "pkg"
	internalDir = "internal"
)

type Dirs struct {
	projectName string
	log         internal.Logger
}

func NewDirs(log internal.Logger) *Dirs {
	return &Dirs{log: log}
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

		d.log.Info("корневая директория проекта успешно создана")
	}

	return nil
}

// createProjectDirs создает структуру папок в директории с проектом.
func (d *Dirs) createProjectDirs() error {
	projectDirs := [3]string{cmdDir, pkgDir, internalDir}
	for i := 0; i < len(projectDirs); i++ {
		dir := fmt.Sprintf("%s/%s", d.projectName, projectDirs[i])
		if err := os.Mkdir(dir, perm); err != nil {
			return errors.New("такая директория уже существует")
		}
		d.log.Info("Директория ./" + projectDirs[i] + " успешно создана")
	}
	return nil
}
