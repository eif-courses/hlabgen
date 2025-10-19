package models

type Doctor struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
}
