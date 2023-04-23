package temp

import "fmt"

type Temp struct{}

func NewTemp() *Temp {
	return &Temp{}
}

func (t Temp) GetMain() string {
	return `package main

import "fmt"

func main() {
	fmt.Println("Hello!")
}
`
}

func (t Temp) GetMakefile(projectName string) string {
	return fmt.Sprintf(
		"PROJECT_NAME = %s\n"+
			"PROJECT_PATH = cmd/$(PROJECT_NAME).go\n\n"+
			".PHONY:run\nrun:\n\tgo run $(PROJECT_PATH)\n\n"+
			".PHONY:build\nbuild:\n\tgo build -o bin/$(PROGRAM_NAME) $(PROJECT_PATH)\n\n"+
			".PHONY:test\ntest:\n\tgo test ./...\n\n"+
			".PHONY:lint\nlint:\n\tgolangci-lint run",
		projectName,
	)
}

func (t Temp) GetGoMod(projectName, goVersion string) string {
	return fmt.Sprintf("module %s\n\n%s", projectName, goVersion)
}

func (t Temp) GetLintConfig() string {
	return `run:
  tests: true
  skip-dirs:
    - build

output:
  format: colored-line-number
  print-issued-lines: true
  print-linter-name: true
  unique-by-line: true
  sort-results: true

linters:
  disable-all: true
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - typecheck
    - unused
    - cyclop
    - dupl
    - dupword
    - funlen
    - lll
    - gochecknoglobals
    - goconst
    - gocritic
    - gocyclo
    - godot
    - gosec
    - nakedret
    - nestif
    - stylecheck
    - varnamelen
`
}
