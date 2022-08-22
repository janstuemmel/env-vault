package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/keyring"
	"github.com/janstuemmel/env-vault/cmd"
	"github.com/janstuemmel/env-vault/storage"
	"github.com/spf13/cobra"
)

func main() {
	// vault
	ring, err := keyring.Open(keyring.Config{
		ServiceName: "envvault",
	})

	if err != nil {
		log.Fatal(err)
	}

	// storage
	str := storage.Storage{Ring: ring}

	// commands
	rc := cmd.RootCommand{}
	cc := cmd.CreateCommand{Storage: str}
	rec := cmd.RemoveCommand{Storage: str}
	ac := cmd.AddCommand{Storage: str}
	lc := cmd.ListCommand{Storage: str}
	ec := cmd.ExecCommand{Storage: str}

	// cli
	envvault := rc.Command()
	envvault.AddCommand(&cobra.Command{
		Use:   "create <env>",
		Short: "Creates an environment",
		Args:  cc.Args,
		RunE:  cc.RunE,
	})
	envvault.AddCommand(&cobra.Command{
		Use:   "add <env> <key>=<value>",
		Short: "Adds a secret to environment",
		Args:  ac.Args,
		RunE:  ac.RunE,
	})
	envvault.AddCommand(&cobra.Command{
		Use:   "remove <env> [<var>]",
		Short: "Removes an environment or secret",
		Args:  rec.Args,
		RunE:  rec.RunE,
	})
	envvault.AddCommand(&cobra.Command{
		Use:   "list [<env>]",
		Short: "Lists items of an environment",
		Args:  lc.Args,
		RunE:  lc.RunE,
	})
	envvault.AddCommand(&cobra.Command{
		Use:   "exec <env> [<program>] [<...args>]",
		Short: "Set env vars ans execute program",
		Args:  ec.Args,
		RunE:  ec.RunE,
	})

	// exec
	if err := envvault.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
