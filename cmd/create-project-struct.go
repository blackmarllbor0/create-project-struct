package main

import (
	"github.com/blackmarllboro/create-project-struct/internal/app/dir"
	"github.com/blackmarllboro/create-project-struct/internal/app/file"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/args"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/git"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/log"

	l "github.com/charmbracelet/log"
)

func main() {
	logger := log.NewLogger(l.New())

	projName := args.ProjectName{}

	project := dir.NewDirs(file.NewFile(), projName)
	if err := project.CreateProject(); err != nil {
		logger.Error(err)
	}

	if err := git.CreateLocalGitRepository(projName); err != nil {
		logger.Error(err)
	}

	logger.Info("the project dir has been successfully created")
}
