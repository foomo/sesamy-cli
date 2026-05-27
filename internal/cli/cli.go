package cli

import (
	"log/slog"

	"github.com/spf13/cobra"
)

// NewCommand returns the fully assembled sesamy command tree.
func NewCommand(l *slog.Logger) *cobra.Command {
	root := NewRoot(l)
	root.AddCommand(
		NewConfig(l),
		NewList(l),
		NewDiff(l),
		NewOpen(l),
		NewProvision(l),
		NewTags(l),
		NewTypeScript(l),
		NewVersion(l),
	)
	return root
}
