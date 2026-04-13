package cmd

import (
	"ghost/internal/client"

	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:  "join [roomId]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client.JoinRoom(args[0])
	},
}