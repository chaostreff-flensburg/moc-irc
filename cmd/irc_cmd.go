package cmd

import (
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	api "github.com/chaostreff-flensburg/moc-go"
	"github.com/chaostreff-flensburg/moc-irc/config"
	"github.com/chaostreff-flensburg/moc-irc/irc"
)

var moc2ircCmd = cobra.Command{
	Use:   "moc2irc",
	Short: "MOC2IRC",
	Long:  "Transfer moc messages to irc.",
	Run: func(cmd *cobra.Command, args []string) {
		execWithConfig(cmd, moc2irc)
	},
}

// moc2irc start moc2irc
func moc2irc(config *config.Config) {
	log.Info("Start")
	ircClient := irc.NewIRC(config.Addr, config.Nick, config.Pass, config.User, config.Name, config.Channel)
	go ircClient.Connect()

	err := <-ircClient.Ready
	if err != nil {
		log.Fatalln(err)
	}

	log.Info("IRC Running!")

	apiClient := api.NewClient(config.Endpoint)
	apiClient.Loop(20 * time.Second)

	for message := range apiClient.NewMessages {
		log.Info(message.ID)
		ircClient.SendMessage(message)
	}
}
