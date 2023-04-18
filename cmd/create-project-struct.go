package main

import (
	"log"

	"github.com/blackmarllboro/create-project-struct/internal/pkg/temp"
)

func main() {
	if _, err := temp.GetTemplateByAlias(temp.MAIN); err != nil {
		log.Println(err)
	}
	//logger := log.NewLogger(l.New())
	//
	//projName := args.ProjectName{}
	//
	//project := dir.NewDirs(file.NewFile(), projName)
	//if err := project.CreateProject(); err != nil {
	//	logger.Error(err)
	//}
	//
	//if err := git.CreateLocalGitRepository(projName); err != nil {
	//	logger.Error(err)
	//}
	//
	//logger.Info("the project dir has been successfully created")
}
