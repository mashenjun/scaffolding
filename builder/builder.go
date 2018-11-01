package builder

type IBuilder interface {
	PrepareDirs() error
	PrepareDeps() error
	PrepareFiles() error
	Name() string
	Deps() []string
}