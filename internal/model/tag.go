package model

import (
	"blog_service/pkg/app"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name   string      `json:"name"`
	State  uint8       `json:"state"`
	Conent interface{} `gorm:"-"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db.Where("name = ?", t.Name)

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
		db.Offset(pageOffset).Limit(pageSize)
	}

	if t.Name != "" {
		db.Where("name = ?", t.Name)
	}

	db = db.Where("state = ?", t.State)

	if err = db.Model(&tags).Where("is_del = ?", 0).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {

	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0).Update(values).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error
}
