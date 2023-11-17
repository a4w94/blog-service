package model

import (
	"blog_service/pkg/app"
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	PublishDate time.Time `json:"publish_date"`
	Tags        Tags      `json:"tags" gorm:"type:text json"`
	Category    string    `json:"category"`
	State       uint8     `json:"state"` //狀態 0:關閉 1:開啟
}

type Tags []string

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func NewArticle() *Article {
	return &Article{}
}

func (a Article) TableName() string {
	return "blog_article"
}

// 建立文章
func (a Article) Create(db *gorm.DB) error {
	return WithTransaction(db, func(tx *gorm.DB) error {
		return db.Create(&a).Error
	})
}

// 更新文章
func (a Article) Update(db *gorm.DB, values interface{}) error {
	db.Model(&a).Where("id = ? AND is_del = ?", a.ID, 0).Update(values)
	return db.Error
}

// 删除文章
func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error
}

func (t *Tags) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tags) Value() (driver.Value, error) {
	return json.Marshal(t)
}
