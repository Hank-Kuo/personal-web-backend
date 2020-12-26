package dto

type GetEmojResDto struct {
	ID      int `json:"id"`
	Funny   int `json:"funny"`
	Sad     int `json:"sad"`
	Wow     int `json:"wow"`
	Clap    int `json:"clap"`
	Perfect int `json:"perfect"`
	Love    int `json:"love"`
	Hard    int `json:"hard"`
	Good    int `json:"good"`
	Mad     int `json:"mad"`
}

type PutEmojiReqDto struct {
	Funny   int `json:"funny"`
	Sad     int `json:"sad"`
	Wow     int `json:"wow"`
	Clap    int `json:"clap"`
	Perfect int `json:"perfect"`
	Love    int `json:"love"`
	Hard    int `json:"hard"`
	Good    int `json:"good"`
	Mad     int `json:"mad"`
}
