package Commands

import (
	"github.com/HutchyBean/discordbot/DCH"
	"log"
)

func init() {
	Commands = append(Commands, &DCH.Command{
		Name:        "ping",
		Aliases: []string{"pong"},
		Description: "Checks if the bot is alive",
		Usage:       "ping",
		SubCommands: nil,
		Execute: func(ctx *DCH.Ctx) {

			_, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Am woke")
			if err != nil {
				log.Println(err)
			}
		},

	})
}
