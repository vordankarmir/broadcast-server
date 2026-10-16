/*
Copyright © 2026 Henry Grigoryan henrygrigoryan5@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	client "github.com/vordankarmir/broadcast-server/cmd/client"
	server "github.com/vordankarmir/broadcast-server/cmd/server"
)

var runServer string
var runClient string

var rootCmd = &cobra.Command{
	Use:   "broadcast-server",
	Short: "Broadcast Server",
	Long:  `A CLI tool that starts a Broadcast Server server, it will forward requests to the actual server and cache the responses.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(server.NewCommand())
	rootCmd.AddCommand(client.NewCommand())
}
