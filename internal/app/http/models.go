package http

type AddNumberRequest struct {
	Value int64 `json:"value"`
}

type NumbersResponse struct {
	Numbers []int64 `json:"numbers"`
}
