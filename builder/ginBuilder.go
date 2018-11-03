package builder

import (
	"github.com/mashenjun/scaffolding/template"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

type ginBuilder struct {
	name        string
	project     string
	projectPath string
	dirs        []string
	deps        []string
	files       map[string][]byte
}

func NewGinBuilder(pPath string) *ginBuilder {
	dirPaths := []string{
		path.Join(pPath, "cmd"),
		path.Join(pPath, "cmd", "backend"),
		path.Join(pPath, "com"),
		path.Join(pPath, "backend"),
		path.Join(pPath, "backend", "api"),
		path.Join(pPath, "backend", "api", "v1"),
		path.Join(pPath, "backend", "config"),
	}

	fileMap := make(map[string][]byte)
	fileMap[path.Join(pPath, "cmd", "backend", "main.go")] = []byte(template.GinMainFile)
	fileMap[path.Join(pPath, "backend", "api", "api_router.go")] = []byte(template.GinApiRouterFile)
	fileMap[path.Join(pPath, "backend", "config", "config.go")] = []byte(template.ConfigFile)
	fileMap[path.Join(pPath, "backend", "api", "v1", "ping.go")] = []byte(template.GinPingFile)

	return &ginBuilder{projectPath: pPath, dirs: dirPaths, project: path.Base(pPath), files: fileMap, name: "gin",
	deps: []string{"github.com/gin-gonic/gin"}}
}

func (g *ginBuilder) PrepareDirs() error {
	for _, v := range g.dirs {
		err := os.MkdirAll(v, 0755)
		if err != nil {
			return errors.Wrap(err, "could not make dir "+ v)
		}
	}
	return nil
}

func (g *ginBuilder) PrepareDeps() error {
	for _, v := range g.deps {
		cmd := exec.Command("go", "get", "-u", v)
		if err := cmd.Run(); err != nil {
			return errors.Wrap(err, "could not download dependency "+ v)
		}
	}
	return nil
}

func (g *ginBuilder) PrepareFiles() error {
	for k, v := range g.files {
		err := ioutil.WriteFile(k, v, 0644)
		if err != nil {
			return errors.Wrap(err, "could not write file "+ k)
		}
	}
	return nil
}

func (g *ginBuilder) Name() string {
	return g.name
}

func (g *ginBuilder) Deps() []string {
	return g.deps
}
