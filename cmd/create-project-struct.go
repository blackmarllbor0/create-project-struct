package main

import (
	"github.com/blackmarllboro/create-project-strcut/internal/app/dir"
	"github.com/charmbracelet/log"
)

func main() {
	logger := log.New()

	project := dir.NewDirs(logger)
	if err := project.CreateProject(); err != nil {
		logger.Error(err)
	}
}
