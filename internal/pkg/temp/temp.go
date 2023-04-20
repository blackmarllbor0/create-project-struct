package temp

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const (
	MAIN = "main.go"
	MAKE = "Makefile"
	LINT = ".golangci.yml"
)

type Template struct{}

func NewTemplate() *Template {
	return &Template{}
}

func (t Template) GetTemplateByAlias(fileName string) (string, error) {
	var pathToTemp string

	tempDir, err := t.getTemplateDir()
	if err != nil {
		return "", fmt.Errorf("fail to get template dir, err: %v", err)
	}

	filesInDirectory, err := ioutil.ReadDir(path.Base(tempDir))
	if err != nil {
		return "", fmt.Errorf("could not read dir, maybe there are no files threr, err: %v", err)
	}

	for _, file := range filesInDirectory {
		if file.Name() == fileName {
			pathToTemp = filepath.Join(tempDir, file.Name())
		}

		break
	}

	if pathToTemp == "" {
		return "", fmt.Errorf("file with name \"%s\" not found, err: %v", fileName, err)
	}

	return pathToTemp, nil
}

func (Template) getTemplateDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("unable to get path to executable file, err: %v", err)
	}

	targetDir := filepath.Join(exePath, "template")

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return "", fmt.Errorf("directory \"%s\" not found, err: %v", targetDir, err)
	}

	return targetDir, nil
}
