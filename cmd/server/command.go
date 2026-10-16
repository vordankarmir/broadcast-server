package cmd

import (
	"github.com/spf13/cobra"
	// "github.com/vordankarmir/broadcast-server/server"
	websocket "github.com/vordankarmir/broadcast-server/websocket"
)

func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Broadcast Server",
		Long:  `A CLI tool that starts a broadcast server, it will act as a fan out application of events and will send to recievers`,
		RunE: func(cmd *cobra.Command, args []string) error {

			// server.Run()
			websocket.ConnectServer()
			return nil
		},
	}

	return cmd
}
