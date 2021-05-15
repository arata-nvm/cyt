package youtube

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arata-nvm/cyt/internal/config"
	"github.com/arata-nvm/cyt/internal/model"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	service *youtube.Service
)

func init() {
	var err error
	service, err = youtube.NewService(context.Background(), option.WithAPIKey(config.GetAPIKey()))
	if err != nil {
		log.Fatal(err)
	}
}

func SearchChannel(q string) ([]*model.Channel, error) {
	call := service.Search.List([]string{"id", "snippet"}).Q(q).MaxResults(5).Type("channel")
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var channels []*model.Channel
	for _, item := range response.Items {
		channel := &model.Channel{
			Id:   item.Id.ChannelId,
			Name: item.Snippet.Title,
		}
		channels = append(channels, channel)
	}

	return channels, nil
}

func SearchVideoBy(channelId string) ([]*model.Video, error) {
	call := service.Search.List([]string{"id", "snippet"}).ChannelId(channelId).MaxResults(5).Type("video").Order("date")
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videos []*model.Video
	for _, item := range response.Items {
		publishedAt, err := time.Parse("2006-01-02T15:04:05Z", item.Snippet.PublishedAt)
		if err != nil {
			return nil, err
		}

		video := &model.Video{
			Id:    item.Id.VideoId,
			Title: item.Snippet.Title,
			Channel: &model.Channel{
				Id:   item.Snippet.ChannelId,
				Name: item.Snippet.ChannelTitle,
			},
			PublishedAt: publishedAt,
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func GetChannelName(channelId string) (string, error) {
	channels, err := SearchChannel(channelId)
	if err != nil {
		return "", err
	}

	if len(channels) == 0 {
		return "", fmt.Errorf("channel `%s` is not found", channelId)
	}

	return channels[0].Name, nil
}
