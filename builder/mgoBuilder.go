package builder

import (
	"github.com/mashenjun/scaffolding/template"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

type mgoBuilder struct {
	name        string
	project     string
	projectPath string
	dirs        []string
	deps        []string
	files       map[string][]byte
}

func NewMgoBuilder(pPath string) *mgoBuilder {
	dirPaths := []string{
		path.Join(pPath, "backend", "model"),
	}

	fileMap := make(map[string][]byte)
	fileMap[path.Join(pPath, "backend", "model", "db.go")] = []byte(template.MgoDBFile)

	return &mgoBuilder{projectPath: pPath, dirs: dirPaths, project: path.Base(pPath), files: fileMap, name: "mgo",
	deps:[]string{"github.com/globalsign/mgo"}}
}

func (m *mgoBuilder) PrepareDirs() error {
	for _, v := range m.dirs {
		err := os.MkdirAll(v, 0755)
		if err != nil {
			return errors.Wrap(err, "could not make dir "+ v)
		}
	}
	return nil
}

func (m *mgoBuilder) PrepareDeps() error {
	for _, v := range m.deps {
		cmd := exec.Command("go", "get", "-u", v)
		if err := cmd.Run(); err != nil {
			return errors.Wrap(err, "could not download dependency "+ v)
		}
	}
	return nil
}

func (m *mgoBuilder) PrepareFiles() error {
	for k, v := range m.files {
		err := ioutil.WriteFile(k, v, 0644)
		if err != nil {
			return errors.Wrap(err, "could not write file "+ k)
		}
	}
	return nil
}

func (m *mgoBuilder) Name() string {
	return m.name
}

func (m *mgoBuilder) Deps() []string {
	return m.deps
}