/*
Copyright © 2026 Henry Grigoryan henrygrigoryan5@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var port int
var origin string
var clearCache bool

var rootCmd = &cobra.Command{
	Use:   "caching-proxy",
	Short: "Caching Proxy",
	Long:  `A CLI tool that starts a caching proxy server, it will forward requests to the actual server and cache the responses.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {

		if !clearCache && origin == "" {
			return fmt.Errorf("--origin and --clear-cache are mutually exclusive")
		}

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
	rootCmd.PersistentFlags().IntVar(&port, "port", 3000, "Port to run the server on")
	rootCmd.PersistentFlags().StringVar(&origin, "origin", "", "Origin server URL")
	rootCmd.Flags().BoolVar(&clearCache, "clear-cache", false, "Clear cache for a specific URL")
}
