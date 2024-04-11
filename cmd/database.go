package main

import (
	"awesome.fintech.org/core"

	"github.com/spf13/cobra"
)

func DatabaseRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "Database migration utility",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "up",
		Short: "Apply migration files",
		Run: func(cmd *cobra.Command, args []string) {
			core.RunMigrations()
		},
	})

	return cmd
}
