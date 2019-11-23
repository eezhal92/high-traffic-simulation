package domain

type Feed struct {
	Title string
}

type FeedResponse struct {
	Data []Feed
	Err  error
}
