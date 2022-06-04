package model

type Page struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
