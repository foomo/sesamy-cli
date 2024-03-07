package cmd

import (
	"bytes"
	"io"
	"os"

	"github.com/foomo/sesamy-cli/pkg/config"
	"github.com/mitchellh/mapstructure"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logger      *pterm.Logger
	verbose     bool
	cfgFilename string
	cfg         *config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sesamy",
	Short: "Server Side Tag Management System",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "output debug information")
	rootCmd.PersistentFlags().StringVarP(&cfgFilename, "config", "c", "", "config file (default is sesamy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")
	if cfgFilename == "-" {
		// do nothing
	} else if cfgFilename != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFilename)
	} else {
		// Search config in home directory with name ".sesamy" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("sesamy")
	}

	// read in environment variables that match
	//viper.EnvKeyReplacer(strings.NewReplacer(".", "_"))
	//viper.SetEnvPrefix("SESAMY")
	//viper.AutomaticEnv()

	logger = pterm.DefaultLogger.WithTime(false)
	if verbose {
		logger = logger.WithLevel(pterm.LogLevelTrace).WithCaller(true)
	}
}

func preRunReadConfig(cmd *cobra.Command, args []string) error {
	if cfgFilename == "-" {
		logger.Debug("using config from stdin")
		b, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			return err
		}
	} else {
		logger.Debug("using config file", logger.Args("filename", viper.ConfigFileUsed()))
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	logger.Debug("config", logger.ArgsFromMap(viper.AllSettings()))

	if err := viper.Unmarshal(&cfg, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	}); err != nil {
		return err
	}

	return nil
}
