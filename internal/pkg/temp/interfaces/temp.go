package interfaces

type Template interface {
	GetMain() string
	GetMakefile(projectName string) string
	GetGoMod(projectName, goVersion string) string
	GetLintConfig() string
}
