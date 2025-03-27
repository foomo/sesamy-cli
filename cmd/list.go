package cmd

import (
	"log/slog"

	"github.com/foomo/sesamy-cli/cmd/list"
	"github.com/spf13/cobra"
)

// NewList represents the list command
func NewList(l *slog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Google Tag Manager containers",
	}

	cmd.AddCommand(
		list.NewWeb(l),
		list.NewServer(l),
	)

	return cmd
}
