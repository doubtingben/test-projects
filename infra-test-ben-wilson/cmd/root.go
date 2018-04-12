package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var config string

var logger *logrus.Logger

type flags struct {
	ListenPort uint16
	verbosity  int
	apiKey     string
}

var commandLineFlags = flags{}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "infura-json-endpoint",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = logrus.New()

		switch commandLineFlags.verbosity {
		case 0:
			logger.Level = logrus.ErrorLevel
			break
		case 1:
			logger.Level = logrus.WarnLevel
			break
		case 2:
			fallthrough
		case 3:
			logger.Level = logrus.InfoLevel
			break
		default:
			logger.Level = logrus.DebugLevel
			break
		}

		logger.Debug("Initialization complete")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().CountVarP(&commandLineFlags.verbosity, "verbosity", "v", "Output verbosity")
	RootCmd.PersistentFlags().StringVarP(&commandLineFlags.apiKey, "apiKey", "a", "", "API Key")
}
