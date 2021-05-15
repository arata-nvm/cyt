package channel

import (
	"os"

	"github.com/arata-nvm/cyt/internal/config"
	"github.com/arata-nvm/cyt/internal/youtube"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		t := tablewriter.NewWriter(os.Stdout)
		t.SetHeader([]string{"Id", "Name"})

		channels := config.GetChannels()
		for _, channelId := range channels {
			channelName, err := youtube.GetChannelName(channelId)
			cobra.CheckErr(err)

			t.Append([]string{channelId, channelName})
		}

		t.Render()
	},
}
