package dto

type EmojiDto struct {
	Funny   *int `json:"funny" binding:"required"`
	Sad     *int `json:"sad" binding:"required"`
	Wow     *int `json:"wow" binding:"required"`
	Clap    *int `json:"clap" binding:"required"`
	Perfect *int `json:"perfect" binding:"required"`
	Love    *int `json:"love" binding:"required"`
	Hard    *int `json:"hard" binding:"required"`
	Good    *int `json:"good" binding:"required"`
	Mad     *int `json:"mad" binding:"required"`
}
