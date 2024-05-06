package main

import (
	"log"
	"os"

	"awesome.fintech.org/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	rootCmd.AddCommand(cmd.DatabaseRootCmd())
	rootCmd.AddCommand(cmd.ServerRootCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)

		os.Exit(1)
	}
}
