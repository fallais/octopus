package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"octopus/internal/cache/local"
	"octopus/internal/proxy"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run is a convenient function for Cobra.
func Run(cmd *cobra.Command, args []string) {
	// Flag
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		logrus.WithError(err).Fatalln("Error with the configuration file flag")
	}

	// Read configuration file
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.WithError(err).Fatalln("Error when reading configuration file")
	}

	// Initialize values with Viper
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		logrus.WithError(err).Fatalln("Error when reading configuration data")
	}

	c, err := local.NewLocalCache(viper.GetString("cache.settings.path"), 1*time.Hour)
	if err != nil {
		logrus.WithError(err).Fatalln("Error while creating the cache")
	}

	// Create the proxy
	p, err := proxy.NewProxy(c, proxy.WithWhitelist(viper.GetStringSlice("security.whitelist")), proxy.WithBlacklist(viper.GetStringSlice("security.blacklist")))
	if err != nil {
		logrus.WithError(err).Fatalln("Error while creating the proxy")
	}

	// Start
	logrus.Infoln("Starting the proxy")
	err = http.ListenAndServe(viper.GetString("general.bind"), p)
	if err != nil {
		logrus.WithError(err).Fatalln("Error while starting the server")
	}

}
