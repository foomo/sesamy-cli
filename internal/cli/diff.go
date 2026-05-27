package cli

import (
	"log/slog"

	"github.com/foomo/sesamy-cli/internal/cli/diff"
	"github.com/spf13/cobra"
)

// NewDiff represents the diff command
func NewDiff(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Print Google Tag Manager container status diff",
	}

	cmd.AddCommand(
		diff.NewWeb(l),
		diff.NewServer(l),
	)

	return cmd
}
