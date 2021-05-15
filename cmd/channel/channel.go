package channel

import "github.com/spf13/cobra"

var ChannelCmd = &cobra.Command{
	Use: "channel [command]",
}

func init() {
	ChannelCmd.AddCommand(addCmd)
	ChannelCmd.AddCommand(listCmd)
	ChannelCmd.AddCommand(rmCmd)
	ChannelCmd.AddCommand(searchCmd)
}
