package storage

import (
	"github.com/pawmart/northerntech-simpletwitter/internal/models"
	"github.com/pawmart/northerntech-simpletwitter/internal/restapi/operations"
)

type Storage interface {
	InsertTweet(entity *models.Tweet) error
	FindTweet(id string) (entity *models.Tweet, err error)
	UpdateTweet(id string, p *models.Tweet) error
	RemoveTweet(id string) error
	FindTweets(params operations.GetTweetsParams) []*models.Tweet
	Ping() error
	Drop()
}