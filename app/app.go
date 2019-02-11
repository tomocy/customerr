package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/urfave/cli"
)

type App interface {
	Run(args []string) error
}

func NewApp() App {
	return &cliApp{
		App: cli.NewApp(),
	}
}

type cliApp struct {
	*cli.App
	errName string
}

func (a *cliApp) Run(args []string) error {
	a.prepare()
	return a.App.Run(args)
}

func (a *cliApp) prepare() {
	a.prepareCommands()
}

func (a *cliApp) prepareCommands() {
	a.prepareCreateCommand()
}

func (a *cliApp) prepareCreateCommand() {
	a.Commands = append(a.Commands, cli.Command{
		Name:  "create",
		Usage: "Create a custome error interface and its concreate type",
		Action: func(ctx *cli.Context) error {
			if len(ctx.Args()) != 1 {
				return fmt.Errorf("failed to create: invalid arguments: got %d, but expected 1", len(ctx.Args()))
			}

			errName := ctx.Args()[0]
			if errName == "" {
				return errors.New("failed to create: invalid custom error name: name should not be empty")
			}
			a.errName = errName

			fmt.Printf(
				"%s\n%s\n%s",
				a.generateCustomErrorInterfaceDefinition(),
				a.generateCustomConcreateTypeDefinition(),
				a.generateCustomErrorInterfaceImplementation(),
			)

			return nil
		},
	})
}

func (a cliApp) generateCustomErrorInterfaceDefinition() string {
	return fmt.Sprintf("type %s interface {\n	error\n	%s()\n}\n", a.generateInterfaceName(), a.generateConcreateTypeName())
}

func (a cliApp) generateCustomConcreateTypeDefinition() string {
	return fmt.Sprintf("type %s struct {\n}\n", a.generateConcreateTypeName())
}

func (a cliApp) generateCustomErrorInterfaceImplementation() string {
	typeName := a.generateConcreateTypeName()
	return fmt.Sprintf("func (e %s) Error() string {\n	return \"%s\"\n}\n\nfunc (e %s) %s() {\n}\n", typeName, typeName, typeName, typeName)
}

func (a cliApp) generateInterfaceName() string {
	return strings.Title(a.errName)
}

func (a cliApp) generateConcreateTypeName() string {
	b := make([]byte, 0, len(a.errName))
	for i, r := range a.errName {
		s := string(r)
		if i == 0 {
			s = strings.ToLower(s)
		}

		b = append(b, s...)
	}

	return string(b)
}
