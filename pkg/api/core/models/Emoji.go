package models

type Emoji struct {
	ID uint `json: "id"`
	BlogID int `json: "bolg_id"`
	Perfect int `json: "perfect"`
	Interesting int`json: "interesting"`
	Mad int `json: "mad"`
	Good int`json: "good"`
	Bad int `json: "bad"`
}