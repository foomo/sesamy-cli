package cmd

import (
	"log/slog"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
