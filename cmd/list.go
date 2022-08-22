package cmd

import (
	"fmt"

	"github.com/janstuemmel/env-vault/storage"
	"github.com/spf13/cobra"
)

type ListCommand struct {
	Storage storage.Storage
}

func (c *ListCommand) Args(cmd *cobra.Command, args []string) error {
	return nil
}

func (c *ListCommand) RunE(cmd *cobra.Command, args []string) error {

	// print envs when no env selected
	if len(args) < 1 {

		keys, err := c.Storage.GetEnvs()
		if err != nil {
			return err
		}

		for _, key := range keys {
			fmt.Println(key)
		}

		return nil
	}

	// print env vars when env selected
	items, err := c.Storage.GetEnv(args[0])

	if err != nil {
		return err
	}

	for _, item := range items {
		fmt.Println(item.Key)
	}

	return nil
}
