package entity

type Preset struct {
	Id        string `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Spoiler   bool   `json:"spoiler"`
	Created   string `json:"created"`
	Updated   string `json:"updated"`
}
