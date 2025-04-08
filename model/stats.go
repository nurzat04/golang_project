package model

type StatsRequest struct {
	Code string `json:"code" binding:"required"`
}

type StatsResponse struct {
	Code  string `json:"code"`
	Count int64  `json:"count"`
}
