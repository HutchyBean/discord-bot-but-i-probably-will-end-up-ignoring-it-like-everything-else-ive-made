package DCH

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CommandHandler struct {
	Prefix string
	Commands []*Command
	Session *discordgo.Session
}

func Init(prefix string, session *discordgo.Session) *CommandHandler {
	router := &CommandHandler{
		Prefix: prefix,
		Commands: make([]*Command, 0),
		Session:  session,
	}
	router.Session.AddHandler(router.Handle)
	return router
}

func (r *CommandHandler) GetCommand(name string) *Command {
	for _, cmd := range r.Commands {
		cmdNames := []string{cmd.Name}
		cmdNames = append(cmdNames, cmd.Aliases...)

		for _, cmdN := range cmdNames {
			if cmdN == name {
				return cmd
			}
		}
	}
	return nil
}

func (r *CommandHandler) AddCommand (command *Command) error {
	// Check for overlapping names or aliases
	for _, cmd := range r.Commands {
		newCmdNames := []string{command.Name}
		newCmdNames = append(newCmdNames, command.Aliases...)

		saveCmdNames := []string{cmd.Name}
		saveCmdNames = append(saveCmdNames, cmd.Aliases...)

		for _, newName := range newCmdNames {
			for _, saveName := range saveCmdNames {
				if newName == saveName {
					return errors.New("overlapping name or alias")
				}
			}
		}
	}

	r.Commands = append(r.Commands, command)
	return nil
}

func (r *CommandHandler) Handle(session *discordgo.Session, message *discordgo.MessageCreate) {
	if !strings.HasPrefix(message.Content, r.Prefix) || message.Author.ID == session.State.User.ID || message.Author.Bot {
		return
	}

	params := SplitToArguments(strings.TrimPrefix(message.Content,r.Prefix))

	cmd := r.GetCommand(params[0])
	if cmd != nil {
		cmd.Run(&Ctx{
			Session:   r.Session,
			Router:    r,
			Arguments: params[1:],
			Command:   cmd,
			Message:   message,
		})
	}
}
