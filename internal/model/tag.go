package model

import (
	"blog_service/pkg/app"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name    string      `json:"name"`  //標籤名稱
	State   uint8       `json:"state"` //狀態 0:關閉 1:開啟
	Content interface{} `gorm:"-"`     //暫存欄位
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func NewTag() *Tag {
	return &Tag{}
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)

	}
	db = db.Where("state = ?", t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error

	if pageOffset > 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)

	if err = db.Model(&tags).Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return WithTransaction(db, func(tx *gorm.DB) error {
		return tx.Create(&t).Error
	})
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	db = db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(values)

	return db.Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error
}
