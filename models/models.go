package models

type Playlist struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Creator string   `json:"creator"`
	Links   []string `json:"links"`
	Public  bool     `json:"public"`
}
