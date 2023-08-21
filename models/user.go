package models

type User struct {
	Id      int      `json:"id"`
	UserID  int      `json:"user-id"`
	Name    string   `json:"name"`
	IsAdmin string   `json:"is-admin"`
	Groups  []string `json:"groups"`
}
