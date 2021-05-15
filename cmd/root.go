package cmd

import (
	"os"

	"github.com/arata-nvm/cyt/cmd/channel"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cyt [command]",
}

func init() {
	rootCmd.AddCommand(channel.ChannelCmd)
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(updateCmd)
}

func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}
