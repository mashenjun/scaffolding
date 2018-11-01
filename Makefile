GOPATH=$(shell echo $$GOPATH)
GOBIN=$(shell pwd)/bin
GOFILES=$(wildcard *.go)
GONAME=$(shell basename "$(PWD)")
PID=/tmp/go-$(GONAME).pid
HEAD_HASH=$(shell git rev-parse HEAD)
BUILDAT=$(shell date +%FT%T%z)

# ready
build:
	@echo "Building $(GOFILES) to current dir"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -race -ldflags "-X main.Version=$(HEAD_HASH) -X main.BuildAt=$(BUILDAT)" -o $(GONAME) $(GOFILES)

# ready
build-linux:
	@echo "Building main.go to current dir"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(HEAD_HASH) -X main.buildAt=$(BUILDAT)" -o $(GONAME) $(GOFILES)

# ready
get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -d gopkg.in/AlecAivazis/survey.v1
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -d github.com/spf13/cobra/cobra
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -d github.com/tj/go-spin
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -d github.com/pkg/errors

# broken
install:
	go install $(GOFILES)

# ready
run:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

# broken
run-container:
	docker run -d --name vms-backend -p 8086:8081 -v /vms-volumes:/vms-volumes vms-backend:0.1

# broken
watch:
	@$(MAKE) restart &
	@fswatch -o . -e 'bin/.*' | xargs -n1 -I{}  make restart

# broken
restart: stop clean build start

# broken
rebuild: clean build

# broken
rebuild-linux: clean build-linux

# broken
start:
	@echo "Starting bin/$(GONAME)"
	@./bin/$(GONAME) & echo $$! > $(PID)

# broken
stop:
	@echo "Stopping bin/$(GONAME) if it's running"
	@-kill `[[ -f $(PID) ]] && cat $(PID)` 2>/dev/null || true

# broken
clear:
	@clear

# broken
clean:
	@echo "Cleaning"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

print-%  : ; @echo $* = $($*)

.PHONY: build get install run watch start stop restart clean print