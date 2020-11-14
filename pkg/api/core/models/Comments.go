package models

import "time"

type Comments struct {
	ID uint `json: "id"`
	BlogID int `json: "bolg_id"`
	Comment string `json: "comment"`
	CreatTime time.Time `json: "create_time"`
	Good int `json: "good"`
	Bad int `json: "bad"`
}