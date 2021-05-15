package model

import "time"

type Video struct {
	Id          string
	Title       string
	Channel     *Channel
	PublishedAt time.Time
}
