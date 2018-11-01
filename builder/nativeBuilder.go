package builder

import (
	"github.com/mashenjun/scaffolding/template"
	"io/ioutil"
	"os"
	"path"
)

type nativeBuilder struct {
	name        string
	project     string
	projectPath string
	dirs        []string
	files       map[string][]byte
	deps        []string
}

func NewNativeBuilder(pPath string) *nativeBuilder {
	dirPaths := []string{
		pPath,
		path.Join(pPath, "cmd"),
		path.Join(pPath, "com"),
	}

	fileMap := make(map[string][]byte)
	fileMap[path.Join(pPath, "main.go")] = []byte(template.NativeMainFile)
	return &nativeBuilder{projectPath: pPath, dirs: dirPaths, project: path.Base(pPath), files: fileMap, name: "native"}
}

func (n *nativeBuilder) PrepareDirs() error {
	for _, v := range n.dirs {
		err := os.MkdirAll(v, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *nativeBuilder) PrepareDeps() error {
	return nil
}

func (n *nativeBuilder) PrepareFiles() error {
	for k, v := range n.files {
		err := ioutil.WriteFile(k, v, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *nativeBuilder) Name() string {
	return n.name
}

func (n *nativeBuilder) Deps() []string {
	return n.deps
}
