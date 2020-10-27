package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the serve",
	Run:   runServe,
}

// init initialize serve
func init() {
	rootCmd.AddCommand(serveCmd)
}

func runServe(c *cobra.Command, args []string) {
	level, err := log.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		log.WithError(err).Error("xendit-ta terminated")
		os.Exit(1)
	}
	log.SetLevel(level)

	log.Info("Starting xendit-ta...")
	apiServer, err := createServer()
	if err != nil {
		log.WithError(err).Error("xendit-ta terminated")
		os.Exit(1)
	}
	if err := apiServer.Run(); err != nil {
		log.WithError(err).Error("xendit-ta terminated")
		os.Exit(1)
	}
}
