package Commands

import (
	"github.com/HutchyBean/discordbot/DCH"
	"log"
)

// So i had major big brain move. using init for each command to append to a list.


var Commands []*DCH.Command

func Load(handler *DCH.CommandHandler) {
	for _, cmd := range Commands {
		err := handler.AddCommand(cmd)
		if err != nil {
			log.Println(err)
		}
	}
}
