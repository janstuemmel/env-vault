package cmd

import (
	"errors"

	"github.com/janstuemmel/env-vault/storage"
	"github.com/spf13/cobra"
)

type CreateCommand struct {
	Storage storage.Storage
}

func (c *CreateCommand) Args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires at least one arg")
	}
	return nil
}

func (c *CreateCommand) RunE(cmd *cobra.Command, args []string) error {
	err := c.Storage.CreateEnv(args[0])

	if err != nil {
		return err
	}

	return nil
}
