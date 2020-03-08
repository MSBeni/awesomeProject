package main

import (
	"github.com/spf13/cobra"
	"time"
)

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use:                    "curtime",
		Aliases:                nil,
		SuggestFor:             nil,
		Short:                  "",
		Long:                   "",
		Example:                "",
		ValidArgs:              nil,
		Args:                   nil,
		ArgAliases:             nil,
		BashCompletionFunction: "",
		Deprecated:             "",
		Hidden:                 false,
		Annotations:            nil,
		Version:                "",
		PersistentPreRun:       nil,
		PersistentPreRunE:      nil,
		PreRun:                 nil,
		PreRunE:                nil,
		Run:                    nil,
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettytime := now.Format(time.RubyDate)
			cmd.Println("Hello Gopher, the current time is: ", prettytime)
			return nil
		},
	}
}
