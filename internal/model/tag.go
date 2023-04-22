package model

import "blog_service/pkg/app"

type Tag struct {
	*Model
	Name   string `json:"name"`
	State  uint8  `json:"state"`
	Conent interface{}
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
