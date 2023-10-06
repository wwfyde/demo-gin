package models

type User struct {
	Id        int      `json:"id"`
	UserID    int      `json:"user-id"`
	Name      string   `json:"name"`
	IsAdmin   bool     `json:"is-admin"`
	Groups    []string `json:"groups"`
	Tags      []Tag
	IsDeleted bool `json:"is-deleted"`
}

type Tag struct {
	ID   int
	Name string
}
