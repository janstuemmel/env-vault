package cmd

import (
	"github.com/spf13/cobra"
)

type RootCommand struct{}

func (rc *RootCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use: "env-vault",
	}
}
