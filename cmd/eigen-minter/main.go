package main

import (
	"log/slog"
	"os"

	"github.com/NethermindEth/eigen-minter/cli"
)

func main() {
	rootCmd := cli.RootCmd()
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	if err := rootCmd.Execute(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
