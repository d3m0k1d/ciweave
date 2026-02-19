package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ciweave",
	Short: "automatic pipeline generator",
	Long:  "ciweave is a automatic pipeline generator for ci/cd written in golang",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
