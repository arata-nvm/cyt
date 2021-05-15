package channel

import (
	"fmt"

	"github.com/arata-nvm/cyt/internal/config"
	"github.com/arata-nvm/cyt/internal/youtube"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add <channel_id>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("add needs channel id"))
		}

		channelId := args[0]

		// check if a channel exists
		_, err := youtube.GetChannelName(channelId)
		cobra.CheckErr(err)

		config.AddChannel(channelId)
	},
}
