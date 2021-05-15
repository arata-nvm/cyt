package cmd

import (
	"github.com/arata-nvm/cyt/internal/youtube"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use: "update",
	Run: func(cmd *cobra.Command, args []string) {
		youtube.ShowRecentVideos()
	},
}
