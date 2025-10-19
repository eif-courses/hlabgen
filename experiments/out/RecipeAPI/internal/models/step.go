package models

type Step struct {
	ID    int    `json:"id"`
	Order int    `json:"order"`
	Desc  string `json:"description"`
}
