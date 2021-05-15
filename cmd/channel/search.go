package channel

import (
	"fmt"
	"os"

	"github.com/arata-nvm/cyt/internal/youtube"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use: "search [query]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("search needs query"))
		}

		query := args[0]

		channels, err := youtube.SearchChannel(query)
		if err != nil {
			cobra.CheckErr(err)
		}

		t := tablewriter.NewWriter(os.Stdout)
		t.SetHeader([]string{"Id", "Name"})

		for _, channel := range channels {
			t.Append([]string{channel.Id, channel.Name})
		}

		t.Render()
	},
}
