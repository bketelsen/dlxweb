package cmd

import (
	"os"
	"path/filepath"

	"github.com/bketelsen/dlxweb/generated/client"
	"github.com/spf13/cobra"

	wlog "github.com/dixonwille/wlog/v3"
	"github.com/spf13/viper"
)

var cfgFile string
var cl *client.Client
var log wlog.UI
var host string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dlx",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	log = wlog.New(os.Stdin, os.Stdout, os.Stderr)

	log = wlog.AddPrefix("?", wlog.Cross, "i", "-", "", "~", wlog.Check, "!", log)
	log = wlog.AddConcurrent(log)
	log = wlog.AddColor(wlog.None, wlog.Red, wlog.Blue, wlog.None, wlog.None, wlog.None, wlog.Cyan, wlog.Green, wlog.Magenta, log)
	rootCmd.PersistentFlags().StringVar(&host, "api", "", "dlx API address")
	viper.BindPFlag("api", rootCmd.PersistentFlags().Lookup("api"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dlx" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".dlx")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file: " + viper.ConfigFileUsed())
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn("Didn't find config file")

			home, err := os.UserHomeDir()
			cobra.CheckErr(err)
			viper.SafeWriteConfigAs(filepath.Join(home, ".dlx.yaml"))
			log.Info("Edit 'api' key in config file " + filepath.Join(home, ".dlx.yaml"))
		} else {
			// Config file was found but another error was produced
			log.Error(err.Error())
		}
	}

	host = viper.GetString("api")
	cl = client.New(host)
}
