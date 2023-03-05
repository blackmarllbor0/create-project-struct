package file

import (
	"fmt"
	"github.com/blackmarllboro/create-project-struct/internal"
	"github.com/blackmarllboro/create-project-struct/pkg/version"
	"os"
	"path"
)

type File struct {
	Logger internal.Logger
}

func NewFile(logger internal.Logger) *File {
	return &File{Logger: logger}
}

// createAndWriteFile создает файл с заданным именем и вписывает в файл переданный контент.
func (fl File) createAndWriteFile(dir, content string) error {
	f, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	if err != nil {
		return err
	}

	if _, err := f.WriteString(content); err != nil {
		return err
	}

	fl.Logger.Info("Файл " + dir + " успешно создан")

	return nil
}

// GenerateMainFile создает основной файл приложения с именем проекта
func (fl File) GenerateMainFile(dir string) error {
	const content = `
package main

import "fmt"

func main() {
	fmt.Println("Hello!")
}
`

	if err := fl.createAndWriteFile(dir, content); err != nil {
		return err
	}

	return nil
}

// GenerateCfgFile создает конфигурационный файл config.yaml.
func (fl File) GenerateCfgFile(dir string) error {
	if err := fl.createAndWriteFile(dir+"/config.yaml", ""); err != nil {
		return err
	}

	return nil
}

// GenerateGoModFile создает go.mod файл с именем проекта.
func (fl File) GenerateGoModFile(dir string) error {
	goVersion, err := version.GoVersion()
	if err != nil {
		return err
	}

	projectName := path.Base(dir)

	content := fmt.Sprintf("module %s\n\n%s", projectName, goVersion)

	if err := fl.createAndWriteFile(projectName+"/go.mod", content); err != nil {
		return err
	}

	return nil
}
