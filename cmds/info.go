package cmds

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bleuhold/bh/cmd"
	"github.com/bleuhold/bh/filesys"
	"github.com/google/uuid"
	"log"
	"time"
)

var INFO *cmd.Command
var infoFilename = "info.json"

func init() {
	INFO = &cmd.Command{
		Name:        "info",
		Description: "To show the current info.",
		FlagSet:     flag.NewFlagSet("info", flag.ExitOnError),
		Execute:     infoExecute,
	}
	// Define the local usage for a flag.
	INFO.FlagSet.BoolVar(&help, "h", false, "")
	INFO.FlagSet.BoolVar(&help, "help", false, "Display the info help.")
	INFO.FlagSet.BoolVar(&list, "list", false, "List all of the current context information.")
	INFO.FlagSet.BoolVar(&set, "set", false, "Set some info parameter used within the global context of this application.")
}

// executes the info command
func infoExecute(cmd *cmd.Command) {
	switch {
	case help:
		cmd.PrintHelp()
	case list:
		ListInfo()
	default:
		cmd.PrintHelp()
	}
}

type Info struct {
	PropertyUUID uuid.UUID `json:"propertyUUID"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
}

// LoadInfo loads the info data.
func LoadInfo() *Info {
	info := &Info{
		PropertyUUID: uuid.UUID{},
		StartDate:    time.Now(),
		EndDate:      time.Now(),
	}
	filesys.LoadInterface(infoFilename, info)
	return info
}

// Print prints the info data to the CLI.
func (i *Info) Print() string {
	s := ""
	s += fmt.Sprintf("%-15s %v\n", "Property UUID:", i.PropertyUUID)
	s += fmt.Sprintf("%-15s %v\n", "Start Date:", i.StartDate.Format("2006-01-02"))
	s += fmt.Sprintf("%-15s %v\n", "End Date:", i.EndDate.Format("2006-01-02"))
	return s
}

// Save writes the data to the file system.
func (i *Info) Save() {
	xb, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("Unable to marshal info data: %v", err)
	}
	filesys.WriteFile(infoFilename, xb)
}

// ListInfo lists the info to the console.
func ListInfo() {
	i := LoadInfo()
	fmt.Printf("%s\n", i.Print())
}
