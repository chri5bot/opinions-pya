package cmd

import (
	"log"

	"github.com/chri5bot/opinions-pya/services/opinions/conf"
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
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %+v", err)
	}
	log.Println("config", config)

}
