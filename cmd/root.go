package cmd

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	root *cobra.Command
)

func init() {
	pterm.Info.Prefix.Text = "⎈"
	pterm.Info.Scope.Style = &pterm.ThemeDefault.DebugMessageStyle
	pterm.Debug.Prefix.Text = "⛏︎"
	pterm.Debug.Scope.Style = &pterm.ThemeDefault.DebugMessageStyle
	pterm.Fatal.Prefix.Text = "⛔︎"
	pterm.Fatal.Scope.Style = &pterm.ThemeDefault.DebugMessageStyle
	pterm.Error.Prefix.Text = "⛌"
	pterm.Error.Scope.Style = &pterm.ThemeDefault.DebugMessageStyle
	pterm.Warning.Prefix.Text = "⚠"
	pterm.Warning.Scope.Style = &pterm.ThemeDefault.DebugMessageStyle
	pterm.Success.Prefix.Text = "✓"
	pterm.Success.Scope.Style = &pterm.ThemeDefault.DebugMessageStyle

	if scope := os.Getenv("SESAMY_SCOPE"); scope != "" {
		pterm.Info.Scope.Text = scope
		pterm.Debug.Scope.Text = scope
		pterm.Fatal.Scope.Text = scope
		pterm.Error.Scope.Text = scope
		pterm.Warning.Scope.Text = scope
		pterm.Success.Scope.Text = scope
	}

	root = NewRoot()
	NewConfig(root)
	NewVersion(root)
	NewTags(root)
	NewList(root)
	NewProvision(root)
	NewTypeScript(root)
}

// NewRoot represents the base command when called without any subcommands
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "sesamy",
		Short:         "Server Side Tag Management System",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			if viper.GetBool("verbose") {
				pterm.EnableDebugMessages()
			}
			return nil
		},
	}
	cmd.PersistentFlags().BoolP("verbose", "v", false, "output debug information")
	_ = viper.BindPFlag("verbose", cmd.PersistentFlags().Lookup("verbose"))

	cmd.PersistentFlags().StringSliceP("config", "c", []string{"sesamy.yaml"}, "config files (default is sesamy.yaml)")
	_ = viper.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))
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
		pterm.Error.Println(root.UsageString())
		code = 1
	}
}
