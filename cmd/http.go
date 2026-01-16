/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	server2 "github.com/mauriciovictor/curso-hexagonal/adapters/web/server"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long:  `A `,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebserver()
		server.Service = &productService

		fmt.Println("Starting server...")
		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

}
