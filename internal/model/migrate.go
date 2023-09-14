package model

import "github.com/jinzhu/gorm"

// 自動migrate
func migrateDB(db *gorm.DB) {
	db.AutoMigrate(NewTag())
	db.AutoMigrate(NewArticle())
	db.AutoMigrate(NewArticleTag())
	db.AutoMigrate(NewAuth())

}
