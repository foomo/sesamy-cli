package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/foomo/sesamy-cli/cmd"
	cmdx "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/pkg/errors"
)

func main() {
	l := cmdx.NewLogger()

	root := cmd.NewRoot(l)
	root.AddCommand(
		cmd.NewConfig(l),
		cmd.NewList(l),
		cmd.NewDiff(l),
		cmd.NewOpen(l),
		cmd.NewProvision(l),
		cmd.NewTags(l),
		cmd.NewTypeScript(l),
		cmd.NewVersion(l),
	)

	say := func(msg string) string {
		if say, cerr := cowsay.Say(msg, cowsay.BallonWidth(80)); cerr == nil {
			msg = say
		}
		return msg
	}

	code := 0
	defer func() {
		if r := recover(); r != nil {
			l.Error(say("It's time to panic"))
			l.Error(fmt.Sprintf("%v", r))
			l.Error(string(debug.Stack()))
			code = 1
		}
		os.Exit(code)
	}()

	if err := root.Execute(); err != nil {
		l.Error(say(strings.Split(errors.Cause(err).Error(), ":")[0]))
		l.Error(err.Error())
		code = 1
	}
}
