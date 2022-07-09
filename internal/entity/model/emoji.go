package model

import "time"

type Emoji struct {
	ID        int       `json:"id" db:"id"`
	BlogID    int       `json:"-" db:"blog_id"`
	Funny     int       `json:"funny" db:"funny"`
	Sad       int       `json:"sad" db:"sad"`
	Wow       int       `json:"wow" db:"wow"`
	Clap      int       `json:"clap" db:"clap"`
	Perfect   int       `json:"perfect" db:"perfect"`
	Love      int       `json:"love" db:"love"`
	Hard      int       `json:"hard" db:"hard"`
	Good      int       `json:"good" db:"good"`
	Mad       int       `json:"mad" db:"mad"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"updated_at" db:"updated_at"`
}
