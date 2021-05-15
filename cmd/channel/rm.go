package channel

import (
	"fmt"

	"github.com/arata-nvm/cyt/internal/config"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use: "rm",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cobra.CheckErr(fmt.Errorf("rm needs channel id"))
		}

		channelId := args[0]

		config.RemoveChannel(channelId)
	},
}
