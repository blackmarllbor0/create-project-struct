package args

type GetProjectName interface {
	GetProjectName() (string, bool, error)
}
