package main

import (
	"github.com/spf13/cobra"

	"awesome.fintech.org/core"
)

func ServerRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Database migration utility",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "up",
		Short: "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			core.SetupDependencies()
		},
	})

	return cmd
}
