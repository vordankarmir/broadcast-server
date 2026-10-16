package cmd

import (
	"github.com/spf13/cobra"
	websocket "github.com/vordankarmir/broadcast-server/websocket"
)

func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "client",
		Short: "Broadcast Client",
		Long:  `A CLI tool that starts a broadcast server client, it will send events and recieve from other clients`,
		RunE: func(cmd *cobra.Command, args []string) error {

			websocket.ConnectWebsocketClient()
			// cli.Read()
			return nil
		},
	}

	return cmd
}
