package cmd

import (
	"fmt"
	"github.com/mashenjun/scaffolding/builder"
	"github.com/mashenjun/scaffolding/enums"
	"github.com/spf13/cobra"
	"github.com/tj/go-spin"
	"gopkg.in/AlecAivazis/survey.v1"
	"os"
	"strings"
	"time"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init creates the project directory and init file content under the current dir",
	Long:  `init creates the project directory and init file content under the current dir`,
	Run:   initRun,
}

func initRun(cmd *cobra.Command, args []string) {
	var (
		selectedBuilders []builder.IBuilder
		withDeps         = false
	)
	var pPath = ""
	var projectTyp string
	projectPrompt := &survey.Select{
		Message: "Select projectTyp type: ",
		Options: []string{enums.RestfulAPI, enums.Native},
	}
	err := survey.AskOne(projectPrompt, &projectTyp, nil)
	if err != nil {
		fmt.Printf("could not get projectTyp answer: %v\n", err)
		return
	}

	switch projectTyp {
	case enums.Native:
		selectedBuilders = append(selectedBuilders, builder.NewNativeBuilder(pPath))
		goto START_BUILDING
	case enums.RestfulAPI:
		frameworkOpts := []string{enums.GIN}
		var framework string
		frameworkPrompt := &survey.Select{
			Message: "Select api framework:",
			Options: frameworkOpts,
		}
		err = survey.AskOne(frameworkPrompt, &framework, nil)
		if err != nil {
			fmt.Printf("could not get framework: %v\n", err)
			return
		}
		switch framework {
		case enums.GIN:
			selectedBuilders = append(selectedBuilders, builder.NewGinBuilder(pPath))
		default:
			return
		}

		daoOpts := []string{enums.Mgo}
		var dao string
		daoPromot := &survey.Select{
			Message: "Select dao:",
			Options: daoOpts,
		}
		err = survey.AskOne(daoPromot, &dao, nil)
		if err != nil {
			fmt.Printf("could not get dao: %v\n", err)
			return
		}
		switch dao {
		case enums.Mgo:
			selectedBuilders = append(selectedBuilders, builder.NewMgoBuilder(pPath))
		default:
			return

		}

		depsPrompt := &survey.Confirm{
			Message: "Download dependencies ?",
		}
		err = survey.AskOne(depsPrompt, &withDeps, nil)
		if err != nil {
			fmt.Printf("could not get dependencies: %v\n", err)
			return
		}

	default:
		return
	}

START_BUILDING:
	names := make([]string, 0, len(selectedBuilders))
	for _, v := range selectedBuilders {
		names = append(names, v.Name())
	}
	confirm := false
	confirmPrompt := &survey.Confirm{
		Message: fmt.Sprintf("Build with %s under current dir?", strings.Join(names, ", ")),
	}
	err = survey.AskOne(confirmPrompt, &confirm, nil)
	if err != nil {
		fmt.Printf("could not get confirm: %v\n", err)
		return
	}
	if !confirm {
		return
	}
	// build projectTyp
	if _, err := os.Stat(pPath); !os.IsNotExist(err) {
		fmt.Printf("%s is already existed\n", pPath)
		return
	}
	errChan := make(chan error)
	doneChan := make(chan struct{})
	msg := make(chan string)
	go builder.BuildProject(doneChan, errChan, msg, selectedBuilders, withDeps)
	s := spin.New()
	message := "..."
	for {
		select {
		case <-doneChan:
			goto END
		case err := <-errChan:
			fmt.Printf("\nfailed: %v\n", err)
			goto END
		case m := <-msg:
			message = m
		default:
			fmt.Printf("\r  \033[36mgernating\033[m %s %s", s.Next(), message)
			time.Sleep(100 * time.Millisecond)
		}
	}
END:
}