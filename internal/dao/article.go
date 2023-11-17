package dao

import (
	"blog_service/internal/model"
	"fmt"
	"time"
)

func (d *Dao) CreateArticle(title, content, author, category string, tags []string, state uint8, publishDate time.Time) error {
	var jsTags model.Tags
	jsTags.Scan(tags)
	article := model.Article{
		Title:       title,
		Content:     content,
		Author:      author,
		PublishDate: publishDate,
		Tags:        jsTags,
		Category:    category,
		State:       state,
	}
	fmt.Println("article:", article)
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, content string, state uint8, modifiedBy string) error {
	article := model.Article{
		Model: &model.Model{
			ID: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}

	if content != "" {
		values["content"] = content
	}

	return article.Update(d.engine, values)
}
