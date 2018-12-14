package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a new opinions REST server",
	Long:  `Start a new opinions REST server`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	log.Println("serve Christian Torres")
}
