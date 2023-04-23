package interfaces

type GetProjectName interface {
	GetProjectName() (string, bool, error)
}
