package config

import (
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/viper"
)

func init() {
	home := os.Getenv("HOME")
	configPath := filepath.Join(home, ".config/cyt/config")
	viper.AddConfigPath(filepath.Dir(configPath))
	viper.SetConfigName(filepath.Base(configPath))
	viper.SetConfigType("yaml")

	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		log.Fatal(err)
	}
	if f, err := os.OpenFile(configPath, os.O_RDONLY|os.O_CREATE, 0666); err != nil {
		log.Fatal(err)
	} else {
		f.Close()
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	viper.SetDefault("apikey", "")
	viper.SetDefault("channels", []string{})
	viper.SetDefault("player", "mpv")
}

func writeConfig() {
	if err := viper.WriteConfig(); err != nil {
		log.Fatal(err)
	}
}

func GetAPIKey() string {
	return viper.GetString("apikey")
}

func GetChannels() []string {
	var channels []string
	viper.UnmarshalKey("channels", &channels)
	return channels
}

func SetChannels(channels []string) {
	viper.Set("channels", channels)
	writeConfig()
}

func AddChannel(channelId string) {
	channels := GetChannels()
	channels = append(channels, channelId)
	SetChannels(channels)
}

func RemoveChannel(channelId string) {
	channels := GetChannels()

	idx := sort.Search(len(channels), func(i int) bool { return channels[i] == channelId })
	if idx < len(channels) && channels[idx] == channelId {
		newChannels := append(channels[:idx], channels[idx+1:]...)
		SetChannels(newChannels)
	}
}

func GetPlayer() string {
	return viper.GetString("player")
}

func SetPlayer(player string) {
	viper.Set("player", player)
}
