package file

import (
	"fmt"
	"os"
	"path"

	"github.com/blackmarllboro/create-project-struct/pkg/version"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

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

	return nil
}

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

func (fl File) GenerateCfgFile(dir string) error {
	if err := fl.createAndWriteFile(dir+"/config.yaml", ""); err != nil {
		return err
	}

	return nil
}

func (fl File) GenerateGoModFile(dir string, isCurrentDir bool) error {
	goVersion, err := version.GoVersion()
	if err != nil {
		return err
	}

	projectName := path.Base(dir)

	content := fmt.Sprintf("module %s\n\n%s", projectName, goVersion)

	var creatingFile string
	if isCurrentDir {
		creatingFile = "go.mod"
	} else {
		creatingFile = projectName + "/go.mod"
	}

	if err := fl.createAndWriteFile(creatingFile, content); err != nil {
		return err
	}

	return nil
}

func (fl File) GenerateMakefile(projectName string, isCurrentDir bool) error {
	content := fmt.Sprintf(
		"PROJECT_NAME = %s\n"+
			"PROJECT_PATH = cmd/$(PROJECT_NAME).go\n\n"+
			".PHONY:run\nrun:\n\tgo run $(PROJECT_PATH)\n\n"+
			".PHONY:build\nbuild:\n\tgo build -o bin/$(PROGRAM_NAME) $(PROJECT_PATH)\n\n"+
			".PHONY:test\ntest:\n\tgo test ./...\n\n"+
			".PHONY:lint\nlint:\n\tgolangci-lint run",
		projectName,
	)

	var creatingFile string
	if isCurrentDir {
		creatingFile = "Makefile"
	} else {
		creatingFile = projectName + "/Makefile"
	}

	if err := fl.createAndWriteFile(creatingFile, content); err != nil {
		return err
	}

	return nil
}
