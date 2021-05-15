package cmd

import (
	"fmt"

	"github.com/arata-nvm/cyt/internal/player"
	"github.com/spf13/cobra"
)

var (
	playCmd = &cobra.Command{
		Use: "play <id>...",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cobra.CheckErr(fmt.Errorf("play needs a video id"))
			}

			videoIds := args
			player.PlayVideos(videoIds)
		},
	}
)
