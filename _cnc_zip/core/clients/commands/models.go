package commands

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs/models"
	"strings"
	"sync"
)

var (
	//stores all the commands correctly
	//this will hold everysingle command inside the system
	Commands map[string]*Command = make(map[string]*Command)
	mutex    sync.Mutex
)

//stores the commands structure properly
//this will help without issues happening on requests
type Command struct {
	//stores the command name properly
	//this will be called on when the command is being executed
	//allows for better control without issues happening on request
	CommandName string //stores the commandName properly without issues
	//this will store all the different aliases
	//this will ensure its done correctly without issues
	Aliases []string
	//stores the command description properly
	//this will store the decription properly and safely
	//this will help users to view what the command does without issues happening
	CommandDescription string //stores the command decription properly
	//stores all the permissions needed
	//this will ensure its done without errors happening
	CommandPermissions []string //stored in array of strings 
	//stores the command function properly
	//this will execute when the command length isnt greater than 1
	CommandFunction func(s *sessions.Session, cmd []string) error
	//stores all the subcommands properly
	//this will ensure its correctly done without issues
	SubCommands []SubCommand //stored in array of them properly
	//stores the invalid subcommand function
	//this will render when the subcommand is invalid properly
	InvalidSubCommand func(s *sessions.Session, cmd []string) error
	//stores the custom command information
	//this will ensure its done without any errors
	CustomCommand string //stored the body correctly
	//stores the bin command context properly
	//this will ensure its done without any issues
	BinCommand *models.BinCommand //stores the information
}

//stores the subcommand structure
//this will be used within subcommands properly
type SubCommand struct {
	//stores the subcommand name properly
	//this will be used properly without issues happening
	SubcommandName string //stores the subcommandName properly
	//stores the subcommand description properly
	//this will ensure its done correctly without issues
	Description string //stores the description properly
	//stores all the permissions properly
	//this will ensure its done correctly without issues
	CommandPermissions []string
	//this will allow for subcommand name splitting
	//functions like `admin` could accept `admin=true` with the split
	CommandSplit string
	//stores the execution path for the subfunction properly
	//this will ensure its done correctly without issues happening
	SubCommandFunction func(s *sessions.Session, cmd []string) error
	//stores if the command wants to be executed properly
	//this will ensure its done without errors happening
	RenderRef bool

	//grabs the auto complete system properly
	//this will allow us to get the array for possible
	AutoComplete func(s *sessions.Session) []string
}

//makes the command correctly
//this will ensure its done correctly without issues
func MakeCommand(c *Command) {
	mutex.Lock()
	defer mutex.Unlock()
	//registers the command properly
	//this will ensure its done correctly and properly
	Commands[c.CommandName] = c
}

//tries to correctly find the command
//this will ensure its done properly without issues happening
func TryCommand(command string) *Command {
	//tries to find the main command name
	//this will ensure its done correctly without issues
	if cmd := Commands[command]; cmd != nil { //checks for the error
		return cmd //returns the command structure properly
	}
	
	//ranges through all the commands
	//this will ensure its done properly without issues
	for c := range Commands {
		//ranges through all the alises
		//this will ensure its done correctly
		for aliases := range Commands[c].Aliases {
			//correctly lowers the subcommand properly
			//this will ensure its done properly without issues happening
			if strings.ToLower(Commands[c].Aliases[aliases]) == command {
				return Commands[c] //returns the command structure
			}
		}
	}

	//returns nil properly
	//this will ensure its done correctly
	return nil
}

//tries to correctly find the subcommand
//this will ensure its done without issues happening
func (c *Command) FindSubs(inlet string) *SubCommand {
	//ranges throughout the subcommand
	//this will ensure we can find the correct information
	for pos := range c.SubCommands {
		//tries to compare the information
		//this will ensure its done properly without issues
		if c.SubCommands[pos].SubcommandName == strings.SplitAfter(inlet, c.SubCommands[pos].CommandSplit)[0] {
			return &c.SubCommands[pos] //returns the structure
		}
	}
	//returns nil as it couldn't be found
	//this will ensure its done properly without issues
	return nil //return nil and kills function
}