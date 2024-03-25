package types

type RatingConsumeMessage struct {
	MovieId string `json:"movie_id"`
	Like    bool   `json:"like"`
}
