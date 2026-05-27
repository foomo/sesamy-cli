package cli

import (
	"log/slog"

	"github.com/foomo/sesamy-cli/internal/cli/provision"
	"github.com/spf13/cobra"
)

// NewProvision represents the provision command
func NewProvision(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision Google Tag Manager containers",
	}

	cmd.AddCommand(
		provision.NewServer(l),
		provision.NewWeb(l),
	)

	return cmd
}
