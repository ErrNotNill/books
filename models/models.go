package models

import "time"

type Books struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Date time.Time `json:"date"`
	Count int `json:"count"`
	Author []Authors `json:"author"`
}
type Authors struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Birthday time.Time `json:"birthday"`
	Death time.Time `json:"death,omitempty"`
	Books []Books `json:"books"`
}