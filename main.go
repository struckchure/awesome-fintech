package main

import (
	"log"
	"os"

	"awesome.fintech.org/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Short: `
		awesome-fintech
		------------------------------
		This project was bootstrap with awesome.fintech.org, Open Source Ledger Framework,
		built to support core fintech features and scale seemlessly.
		
		https://docs.awesome.fintech.orgserver.xyz
		`,
	}

	rootCmd.AddCommand(cmd.DatabaseRootCmd())
	rootCmd.AddCommand(cmd.ServerRootCmd())

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)

		os.Exit(1)
	}
}
