package dto

import (
	"time"
)

type GetAllBlogResDto struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	CreateTime time.Time `json:"create_time"`
	Tag        string    `json:"tag"`
	Visitor    int       `json:"visitor"`
	ImgLink    string    `json:"img_link"`
}

type GetBlogResDto struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

type SerialGetBlogResDto struct {
	ID   int    `json:"id"`
	Html string `json:"html"`
}

type PostBlogReqDto struct {
	Title   string `json:"title" binding:"required"`
	Tag     string `json:"tag" binding:"required"`
	ImgLink string `json:"img_link" binding:"required"`
	Link    string `json:"link" binding:"required"`
}

type SerialPostBlogReqDto struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	CreateTime time.Time `json:"create_time"`
	Tag        string    `json:"tag"`
	Link       string    `json:"link"`
	Visitor    int       `json:"visitor"`
	ImgLink    string    `json:"img_link"`
}

type PostBlogEmojiDto struct {
	BlogID  int `json:"blog_id"`
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

func SerializeGetBlog(data *GetBlogResDto, html string) *SerialGetBlogResDto {
	return &SerialGetBlogResDto{
		ID:   data.ID,
		Html: html,
	}
}

func SerializePostBlog(data *PostBlogReqDto, t time.Time) *SerialPostBlogReqDto {
	return &SerialPostBlogReqDto{
		Title:      data.Title,
		CreateTime: t,
		Tag:        data.Tag,
		Link:       data.Link,
		Visitor:    0,
		ImgLink:    data.ImgLink,
	}
}

type PutBlogVisitorDto struct {
	Visitor bool `json:"visitor"`
}

/*


func TransferComments(data *[]models.Comments) []GetCommentsDto {
	r := make([]GetCommentsDto, 0, len(*data))
	for _, i := range *data {
		r = append(r, GetCommentsDto{
			ID:         i.ID,
			Name:       i.Name,
			Comment:    i.Comment,
			CreateTime: i.CreateTime,
			Character:  i.Character,
		})
	}
	return r
}


func TransferEmoji(data *models.Emoji) GetEmojDto {
	return GetEmojDto{
		ID:      data.ID,
		Funny:   data.Funny,
		Sad:     data.Sad,
		Wow:     data.Wow,
		Clap:    data.Clap,
		Perfect: data.Perfect,
		Love:    data.Love,
		Hard:    data.Hard,
		Good:    data.Good,
		Mad:     data.Mad,
	}
}

func GetBlogRes(data *models.Blog) *GetBlogResDto {
	comment := TransferComments(&data.Comments)
	emoji := TransferEmoji(&data.Emoji)
	return &GetBlogResDto{
		ID:       data.ID,
		Link:     data.Link,
		Comments: comment,
		Emoji:    emoji,
	}
}
*/
