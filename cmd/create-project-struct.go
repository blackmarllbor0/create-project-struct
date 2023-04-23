package main

import (
	"github.com/blackmarllboro/create-project-struct/internal/app/dir"
	"github.com/blackmarllboro/create-project-struct/internal/app/file"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/args"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/log"
	"github.com/blackmarllboro/create-project-struct/internal/pkg/temp"
	"github.com/blackmarllboro/create-project-struct/pkg/git"

	l "github.com/charmbracelet/log"
)

func main() {
	logger := log.NewLogger(l.New())
	projName := args.NewArgs()
	template := temp.NewTemp()
	newFile := file.NewFile(template)
	project := dir.NewDirs(newFile, projName)

	if err := project.CreateProject(); err != nil {
		logger.Error(err)
		return
	}

	if err := git.CreateLocalGitRepository(projName); err != nil {
		logger.Error(err)
		return
	}

	logger.Info("The project dir has been successfully created")
}
