package cmd

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.td.teradata.com/appcenter/backup-restore/handlers"
	"github.td.teradata.com/appcenter/backup-restore/internal"

	yaml "gopkg.in/yaml.v2"
)

var cfgFile string

const (
	envVarPrefix  = "DEPLOY_SERVICE"
	chartDatabase = "chart"
)

var rootCmd = &cobra.Command{
	Use:   "deploy-service",
	Short: "Deploy-service is a kubernetes chart deployment service",
	RunE: func(cmd *cobra.Command, args []string) error {
		ds := &internal.DiskStore{Status: "uri", Val: 10}
		s := handlers.NewServer(ds)
		//log.Info("Deploy service starting on %v port", config.Port)
		return http.ListenAndServe(":"+strconv.Itoa(config.Port), s.Mux())
	},
}

// Execute bootstraps the deploy-service
func Execute() error {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "configuration file for deploy-service")
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	if err := initConfigE(); err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
		return
	}
}

func initConfigE() error {
	v := viper.New()
	// set default values in viper.
	// Viper needs to know if a key exists in order to override it.
	// https://github.com/spf13/viper/issues/188
	b, err := yaml.Marshal(NewConfig())
	if err != nil {
		return err
	}
	defaultConfig := bytes.NewReader(b)
	v.SetConfigType("yaml")
	if err := v.MergeConfig(defaultConfig); err != nil {
		return err
	}
	// overwrite values from config
	v.SetConfigFile(cfgFile)
	if err := v.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return err
		}
		// dont return error if file is missing. overwrite file is optional
	}
	// tell viper to overrwire env variables
	v.AutomaticEnv()
	v.SetEnvPrefix(envVarPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// refresh configuration with all merged values
	config = &Config{}
	return v.Unmarshal(&config)
}
