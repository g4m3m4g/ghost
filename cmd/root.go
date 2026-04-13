package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "ghost",
}

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(joinCmd)
}