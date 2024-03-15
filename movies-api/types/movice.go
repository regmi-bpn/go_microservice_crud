package types

type MovieRequest struct {
	Name string `json:"name"`
}

type MovieResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
