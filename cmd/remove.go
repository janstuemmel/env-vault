package cmd

import (
	"errors"
	"fmt"

	"github.com/janstuemmel/env-vault/storage"
	"github.com/spf13/cobra"
)

type RemoveCommand struct {
	Storage storage.Storage
}

func (c *RemoveCommand) Args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires at least one arg")
	}
	return nil
}

func (c *RemoveCommand) RunE(cmd *cobra.Command, args []string) error {
	if len(args) == 1 {
		err := c.Storage.RemoveEnv(args[0])

		if err != nil {
			return err
		}

		return nil
	}

	if len(args) == 2 {
		err := c.Storage.Remove(args[0], args[1])

		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("too many arguments")
}
