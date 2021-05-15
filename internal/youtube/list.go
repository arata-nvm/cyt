package youtube

import (
	"log"
	"os"
	"regexp"
	"sort"

	"github.com/arata-nvm/cyt/internal/config"
	"github.com/arata-nvm/cyt/internal/model"
	"github.com/olekukonko/tablewriter"
	"github.com/xeonx/timeago"
)

var reTag = regexp.MustCompile("【[^】]*】")

func ShowRecentVideos() {
	channels := config.GetChannels()
	videos := make([]*model.Video, 0, len(channels)*5)
	for _, channelId := range channels {
		channelsVideos, err := SearchVideoBy(channelId)
		if err != nil {
			log.Fatal(err)
		}
		videos = append(videos, channelsVideos...)
	}

	sort.Slice(videos, func(i, j int) bool { return videos[j].PublishedAt.Before(videos[i].PublishedAt) })

	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Id", "Title", "Channel", "Time"})

	maxVideos := 10
	if len(videos) < maxVideos {
		maxVideos = len(videos)
	}

	for _, video := range videos[:maxVideos] {
		title := reTag.ReplaceAllString(video.Title, "")
		publishedAt := timeago.English.Format(video.PublishedAt)
		t.Append([]string{video.Id, title, video.Channel.Name, publishedAt})
	}

	t.Render()
}
