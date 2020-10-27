// Package cmd contains the root configuration
package cmd

import (
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Default configuration can be set here.
var defaults = map[string]interface{}{
	// log config
	"log.level": "debug",
	// server config
	"server.host": "0.0.0.0",
	"server.port": 8080,
	"server.cors.allowedHeaders": []string{
		"Content-Type",
		"Sec-Fetch-Dest",
		"Referer",
		"accept",
		"User-Agent",
	},
	"server.cors.allowedOrigins": []string{"*"},
	"server.cors.allowedMethods": []string{"GET", "POST", "OPTIONS", "HEAD"},
	// db config
	"db.host":         "postgres",
	"db.port":         5432,
	"db.username":     "xendit",
	"db.password":     "xendit",
	"db.name":         "xendit",
	"db.pool.minOpen": 10,
	"db.pool.maxOpen": 100,
	"db.migrate":      true,
	"db.logMode":      true,
}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "xendit-ta",
	Short: "xendit-ta exercise",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		log.Error("Cannot execute root command: ", err)
		os.Exit(1)
	}
}

// init initialize the root
func init() {
	cobra.OnInitialize(initConfig)

	// load the defaults from the default map here, add new defaults on the map
	for key, val := range defaults {
		viper.SetDefault(key, val)
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xendit-ta.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Error("Could not get home directory: ", err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".xendit-ta")
	}

	viper.SetEnvPrefix("XENDIT_TA")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
