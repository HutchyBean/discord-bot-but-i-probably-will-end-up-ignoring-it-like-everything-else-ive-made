package DCH

import (
	"github.com/bwmarrin/discordgo"
)

type Execute func(*Ctx)

type Ctx struct {
	Session   *discordgo.Session
	Router    *CommandHandler
	Arguments []string
	Command   *Command
	Message   *discordgo.MessageCreate
}

type Command struct {
	Name string
	Aliases []string
	Description string
	Usage string
	SubCommands []*Command
	Execute Execute
}

func (cmd *Command) GetSubcommand(name string) *Command{
	for _, sCmd := range cmd.SubCommands {
		names := make([]string, 0)
		names = append(names, sCmd.Name)
		names = append(names, sCmd.Aliases...)

		for _, n := range names {
			if name == n {
				return sCmd
			}
		}

	}
	return nil
}

func (cmd *Command) Run(ctx *Ctx) {
	if len(ctx.Arguments) > 0 {
		subCmd := cmd.GetSubcommand(ctx.Arguments[0])
		if subCmd != nil {
			subCmd.Run(&Ctx{
				Session:   ctx.Session,
				Router:    ctx.Router,
				Arguments: ctx.Arguments[1:],
				Command:   subCmd,
				Message:   ctx.Message,
			})
			return
		}
	}

	cmd.Execute(ctx)
}