package file

import (
	"fmt"
	"os"

	"github.com/blackmarllboro/create-project-struct/pkg/version"

	"github.com/blackmarllboro/create-project-struct/internal/pkg/temp/interfaces"
)

type File struct {
	temp interfaces.Template
}

func NewFile(temp interfaces.Template) *File {
	return &File{temp: temp}
}

func (fl File) createAndWriteFile(dir, content string) error {
	file, err := os.Create(dir)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	if err != nil {
		return err
	}

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write file, err: %v", err)
	}

	return nil
}

func (fl File) getFileContent(fileName, projectName, goVersion string) string {
	switch fileName {
	case fmt.Sprintf("%s.go", projectName):
		return fl.temp.GetMain()
	case "golangci.yml":
		return fl.temp.GetLintConfig()
	case "Makefile":
		return fl.temp.GetMakefile(projectName)
	case "go.mod":
		return fl.temp.GetGoMod(projectName, goVersion)
	default:
		return ""
	}
}

func (fl File) GenerateFiles(isCurrentDir bool, projectName string) error {
	goVersion, err := version.GoVersion()
	if err != nil {
		return fmt.Errorf("failed to get go version, err: %v", err)
	}

	mainGoFile := fmt.Sprintf("%s.go", projectName)
	fileNames := []string{
		mainGoFile,
		"golangci.yml",
		"Makefile",
		"go.mod",
		"config.yml",
	}

	dirPrefix := "./"
	if !isCurrentDir {
		dirPrefix = fmt.Sprintf("%s/", projectName)
	}

	for _, fileName := range fileNames {
		filePath := dirPrefix + fileName

		switch fileName {
		case mainGoFile:
			filePath = fmt.Sprintf("%s/cmd/%s", dirPrefix, fileName)
		case "config.yml":
			filePath = fmt.Sprintf("%s/config/%s", dirPrefix, fileName)
		}

		if err := fl.createAndWriteFile(filePath, fl.getFileContent(fileName, projectName, goVersion)); err != nil {
			return fmt.Errorf("failed to write file data, err: %v", err)
		}
	}

	return nil
}
