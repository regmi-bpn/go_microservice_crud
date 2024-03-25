package types

type MovieRequest struct {
	Name string `json:"name"`
}

type MovieResponse struct {
	Id     string              `json:"id"`
	Name   string              `json:"name"`
	Rating MovieRatingResponse `json:"rating"`
}

type MovieRatingResponse struct {
	Like    int64 `json:"like"`
	Dislike int64 `json:"dislike"`
}
