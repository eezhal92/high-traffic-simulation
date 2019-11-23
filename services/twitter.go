package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/eezhal92/high-traffic/domain"
)

func GetTwitterFeeds(ctx context.Context) domain.FeedResponse {
	n := rand.Intn(4000)
	timeout := time.Duration(n) * time.Millisecond

	select {
	case <-ctx.Done():
		fmt.Println("tw cancelled")
	case <-time.After(timeout):

		feeds := make([]domain.Feed, 0)
		feedA := domain.Feed{Title: "Twitter World"}
		feedB := domain.Feed{Title: "Twitter World B"}
		feedC := domain.Feed{Title: "Twitter World C"}

		feeds = append(feeds, feedA)
		feeds = append(feeds, feedB)
		feeds = append(feeds, feedC)

		return domain.FeedResponse{
			Data: feeds,
			Err:  nil,
		}
	}

	return domain.FeedResponse{}
}
