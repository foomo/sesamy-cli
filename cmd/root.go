package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	cmdx "github.com/foomo/sesamy-cli/pkg/cmd"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	l    *slog.Logger
	root *cobra.Command
)

func init() {
	l = cmdx.NewLogger()

	root = NewRoot(l)
	root.AddCommand(
		NewConfig(l),
		NewConfig(l),
		NewList(l),
		NewProvision(l),
		NewTags(l),
		NewTypeScript(l),
		NewVersion(l),
	)
}

// NewRoot represents the base command when called without any subcommands
func NewRoot(l *slog.Logger) *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:           "sesamy",
		Short:         "Server Side Tag Management System",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			pterm.PrintDebugMessages = viper.GetBool("verbose")
		},
	}

	flags := cmd.PersistentFlags()

	flags.BoolP("verbose", "v", false, "output debug information")
	_ = c.BindPFlag("verbose", flags.Lookup("verbose"))

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	say := func(msg string) string {
		if say, cerr := cowsay.Say(msg, cowsay.BallonWidth(80)); cerr == nil {
			msg = say
		}
		return msg
	}

	code := 0
	defer func() {
		if r := recover(); r != nil {
			pterm.Error.Println(say("It's time to panic"))
			pterm.Error.Println(fmt.Sprintf("%v", r))
			pterm.Error.Println(string(debug.Stack()))
			code = 1
		}
		os.Exit(code)
	}()

	if err := root.Execute(); err != nil {
		pterm.Error.Println(say(strings.Split(errors.Cause(err).Error(), ":")[0]))
		pterm.Error.Println(err.Error())
		code = 1
	}
}
