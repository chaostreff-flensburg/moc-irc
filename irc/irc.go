package irc

import (
	"crypto/tls"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gopkg.in/irc.v3"

	"github.com/chaostreff-flensburg/moc-go/models"
)

type IRC struct {
	Addr      string
	Nick      string
	Pass      string
	User      string
	Name      string
	Channel   string
	IRCClient *irc.Client
	Ready     chan error
}

// NewIRC create a new IRC Client
func NewIRC(addr, nick, pass, user, name, channel string) *IRC {
	return &IRC{
		Addr:    addr,
		Nick:    nick,
		Pass:    pass,
		User:    user,
		Name:    name,
		Channel: channel,
		Ready:   make(chan error),
	}
}

// Connect sets the config and connects to the irc server. If the connection is, comes via the
// channel client.Ready a nil. If an error occurs in this is the error.
func (client *IRC) Connect() {
	tlsConf := &tls.Config{
		//InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", client.Addr, tlsConf)
	if err != nil {
		client.Ready <- err
	}

	log.Info("Set Config")
	config := irc.ClientConfig{
		Nick: client.Nick,
		Pass: client.Pass,
		User: client.User,
		Name: client.Name,
		Handler: irc.HandlerFunc(func(c *irc.Client, m *irc.Message) {
			log.Info(m.String())
			if m.Command == "001" {
				// 001 is a welcome event, so we join channels there
				log.Info("Join ", client.Channel, " ...")
				c.Write(fmt.Sprintf("JOIN %s", client.Channel))
			} else if m.Command == irc.RPL_ENDOFNAMES {
				log.Info("Joined Channel ", client.Channel, "!")
				client.Ready <- nil
			}
		}),
	}

	log.Info("Start Client")
	// Create the client
	client.IRCClient = irc.NewClient(conn, config)
	err = client.IRCClient.Run()
	if err != nil {
		log.Error(err)
		client.Ready <- err
	}
}

// SendMessage send a Message to client.Channel
func (client *IRC) SendMessage(message *models.Message) {
	client.IRCClient.WriteMessage(&irc.Message{
		Command: "PRIVMSG",
		Params:  []string{client.Channel, message.Text},
	})
}
