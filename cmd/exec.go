package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/janstuemmel/env-vault/storage"
	"github.com/spf13/cobra"
	osexec "golang.org/x/sys/execabs"
)

type ExecCommand struct {
	Storage storage.Storage
}

func (c *ExecCommand) Args(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("requires at least two args")
	}
	return nil
}

// mostly copied from https://github.com/99designs/aws-vault/blob/master/cli/exec.go
func (c *ExecCommand) RunE(cmd *cobra.Command, args []string) error {

	command := os.Getenv("SHELL")

	if len(args) > 1 {
		command = args[1]
	}

	items, err := c.Storage.GetEnv(args[0])

	if err != nil {
		return err
	}

	var env []string

	for _, item := range items {
		env = append(env, fmt.Sprintf("%s=%s", item.Key, item.Value))
	}

	var command_args []string

	if len(args) > 2 {
		command_args = args[2:]
	}

	exe := osexec.Command(command, command_args...)
	exe.Stdin = os.Stdin
	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	exe.Env = append(os.Environ(), env...)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	if err := exe.Start(); err != nil {
		return err
	}

	go func() {
		for {
			sig := <-sigChan
			_ = exe.Process.Signal(sig)
		}
	}()

	if err := exe.Wait(); err != nil {
		_ = exe.Process.Signal(os.Kill)
		return fmt.Errorf("failed to wait for command termination: %v", err)
	}

	// waitStatus := exe.ProcessState.Sys().(syscall.WaitStatus)
	// os.Exit(waitStatus.ExitStatus())
	return nil
}
