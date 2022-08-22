package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/janstuemmel/env-vault/storage"
	"github.com/spf13/cobra"
)

type AddCommand struct {
	Storage storage.Storage
}

func (c *AddCommand) Args(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("requires at least one arg")
	}
	return nil
}

func (c *AddCommand) RunE(cmd *cobra.Command, args []string) error {
	env := args[0]
	data := strings.Split(args[1], "=")

	if len(data) != 2 {
		return fmt.Errorf("wrong format")
	}

	c.Storage.Add(env, data[0], data[1])

	return nil
}
