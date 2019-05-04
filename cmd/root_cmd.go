package cmd

import (
	"github.com/chaostreff-flensburg/moc-irc/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd will run the log streamer
var rootCmd = cobra.Command{
	Use:  "moc-irc",
	Long: "A service that will serve a irc message operation center endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, moc2irc)
	},
}

// RootCmd will add flags and subcommands to the different commands
func RootCmd() *cobra.Command {
	rootCmd.AddCommand(&moc2ircCmd)
	return &rootCmd
}

// execWithConfig load config from env
func execWithConfig(cmd *cobra.Command, fn func(config *config.Config)) {
	logrus.Info("Read Config...")
	config := config.ReadConfig()

	fn(config)
}
