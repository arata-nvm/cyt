package cmd

import (
	"fmt"

	"github.com/arata-nvm/cyt/internal/youtube"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/xeonx/timeago"
)

var updateCmd = &cobra.Command{
	Use: "update",
	Run: func(cmd *cobra.Command, args []string) {
		videos, err := youtube.GetRecentVideos()
		cobra.CheckErr(err)

		bold := color.New(color.Bold)
		for _, video := range videos[:10] {
			when := timeago.English.Format(video.PublishedAt)

			fmt.Printf("%s / ", video.Id)
			bold.Print(video.Title)
			fmt.Printf(" %s\n", when)

			fmt.Printf("    %s\n", video.Channel.Name)
		}
	},
}
