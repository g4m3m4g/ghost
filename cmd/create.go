package cmd

import (
	"fmt"
	"net/http"

	"ghost/internal/client"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:  "create [roomId]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		roomId := args[0]

		resp, err := http.Post("http://localhost:8080/create/"+roomId, "", nil)
		if err != nil || resp.StatusCode != 200 {
			fmt.Println("failed to create room")
			return
		}

		fmt.Println("Created room:", roomId)

		
		client.JoinRoom(roomId)
	},
}