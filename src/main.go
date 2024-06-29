package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
)

// command line interface
// func main() {
// 	LoadConfig()
// 	LoadSkills()
// 	user, err := user.Current()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var localPlayer player

// 	var command []string // = "stats"
// 	if len(os.Args) <= 1 {
// 		command = append(command, "stats")
// 	} else {
// 		for i := 1; i < len(os.Args); i++ {
// 			command = append(command, strings.ToLower(os.Args[i]))
// 		}
// 	}

// 	localPlayer.loadPlayer(user.Name)

// 	commandHandler(&localPlayer, command)

// 	localPlayer.savePlayer()
// }

var botInf botInfo

type botInfo struct {
	APIKey    string
	Name      string
	CmdPrefix string
}

// discord bot
func main() {
	LoadConfig()
	LoadSkills()

	ReadJsonFile(&botInf, "Discord.json")
	WriteJsonFile(&botInf, "Discord.json")

	slog.Info("Starting up " + botInf.Name)
	slog.Info("disgo version", slog.String("version", disgo.Version))

	client, err := disgo.New(botInf.APIKey,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentMessageContent,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
			),
		),
		bot.WithEventListenerFunc(onMessageCreate),
	)
	if err != nil {
		slog.Error("Error building bot", slog.Any("err", err))
		return
	}

	defer client.Close(context.TODO())

	if err = client.OpenGateway(context.TODO()); err != nil {
		slog.Error("Error connecting to gateway", slog.Any("err", err))
		return
	}

	slog.Info(botInf.Name + " is not running, press CTRL-C to exit.")

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

func onMessageCreate(event *events.MessageCreate) {
	if event.Message.Author.Bot {
		return
	}

	arguments := strings.Split(event.Message.Content, " ")

	if len(arguments) < 2 || strings.ToLower(arguments[0]) != strings.ToLower(botInf.CmdPrefix) {
		return
	}

	arguments = append(arguments[:0], arguments[1:]...)
	LoadConfig()
	LoadSkills()
	var localPlayer player

	localPlayer.loadPlayer(event.Message.Author.ID.String())
	localPlayer.Name = event.Message.Author.Username

	message := commandHandler(&localPlayer, arguments)

	if message != "" {
		_, _ = event.Client().Rest().CreateMessage(event.ChannelID, discord.NewMessageCreateBuilder().SetContent(message).Build())
	}

	localPlayer.savePlayer()
}
