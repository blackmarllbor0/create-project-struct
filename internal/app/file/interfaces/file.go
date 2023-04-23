package interfaces

type File interface {
	GenerateFiles(isCurrentDir bool, projectName string) error
}
