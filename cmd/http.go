package cmd

import (
	"fmt"

	server2 "github.com/nironwp/hexagonal-architecture/adapter/server"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebServer(&productService)
		fmt.Println("Webserver starter")
		server.Server()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
