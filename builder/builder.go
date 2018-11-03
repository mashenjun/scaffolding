package builder

import (
	"fmt"
	"strings"
)

type IBuilder interface {
	PrepareDirs() error
	PrepareDeps() error
	PrepareFiles() error
	Name() string
	Deps() []string
}

func BuildProject(done chan struct{}, errChan chan error, msg chan string, selectedBuilders []IBuilder, withDeps bool)  {
	defer close(done)
	defer close(errChan)
	defer close(msg)
	for _, b := range selectedBuilders {
		msg <- fmt.Sprintf("prepare for %s", b.Name())
		if err := b.PrepareDirs(); err != nil {
			msg <- "failed"
			errChan <- err
			return
		}
		if err := b.PrepareFiles(); err != nil {
			msg <- "failed"
			errChan <- err
			return
		}
		if withDeps {
			msg <- fmt.Sprintf("downloading %s", strings.Join(b.Deps(), ","))
			if err := b.PrepareDeps(); err != nil {
				msg <- "failed"
				errChan <- err
				return
			}
		}
	}
	msg <- "finish\n"
	done <- struct{}{}
	return
}