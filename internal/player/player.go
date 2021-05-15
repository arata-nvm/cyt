package player

import (
	"log"
	"os/exec"
	"strings"

	"github.com/arata-nvm/cyt/internal/config"
)

const videoBaseUrl = "https://www.youtube.com/watch?v="

func PlayVideos(videoIds []string) {
	videoUrls := make([]string, 0, len(videoIds))
	for _, videoId := range videoIds {
		videoUrls = append(videoUrls, videoBaseUrl+videoId)
	}

	player := config.GetPlayer()
	playerArgs := strings.Split(player, " ")
	playerCmd := exec.Command(playerArgs[0], append(playerArgs[1:], videoUrls...)...)
	if err := playerCmd.Run(); err != nil {
		log.Fatal(err)
	}
}
