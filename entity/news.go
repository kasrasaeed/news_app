package entity

import "time"

type News struct {
	Author      string
	Content     string
	CreatedDate time.Time
	Description string
	Id          uint32
	Tag         string
	Title       string
}

func NewNews(author, content, description, tag, title string, id uint32) (*News, error) {
	news := &News{
		Author:      author,
		Content:     content,
		CreatedDate: time.Now(),
		Description: description,
		Id:          id,
		Tag:         tag,
		Title:       title,
	}
	err := news.validate()
	if err != nil {
		return nil, err
	}
	return news, nil
}

func (n *News) validate() error {
	if n.Content == "" ||
		n.Author == "" ||
		n.Description == "" ||
		n.Tag == "" ||
		n.Title == "" {
		return InvalidNewsEntity
	}
	return nil
}
