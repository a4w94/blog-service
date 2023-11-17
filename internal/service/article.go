package service

import (
	"blog_service/global"
	"time"
)

type CreateArticleRequest struct {
	Title       string   `json:"title" binding:"required"`   //文章標題
	Content     string   `json:"content" binding:"required"` //文章內容
	Author      string   `json:"author" binding:"required"`
	PublishDate string   `json:"publish_date" binding:"required"`     //文章作者
	Tags        []string `json:"tags" binding:"required"`             //文章標籤
	Category    string   `json:"category" binding:"required"`         //文章分類
	State       uint8    `form:"state,default=1" binding:"eq=0|eq=1"` //文章狀態
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {

	//解析時間格式
	publishDate, err := time.Parse("2006-01-02 15:04:05", param.PublishDate)

	//將param.Tags由切片轉成字串
	//tagsString := strings.Join(param.Tags, ",")

	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err: %v", err)
	}
	return svc.dao.CreateArticle(param.Title, param.Content, param.Author, param.Category, param.Tags, param.State, publishDate)
}
