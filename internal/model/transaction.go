package model

import (
	"blog_service/global"

	"github.com/jinzhu/gorm"
)

// 交易錯誤處理與提交
func WithTransaction(db *gorm.DB, fn func(*gorm.DB) error) error {
	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			rollbackTransaction(tx)
			panic(r)
		} else {
			err := tx.Commit().Error
			if err != nil {
				rollbackTransaction(tx)
			}
		}
	}()

	return fn(tx)

}

// Rollback 函數用於回滾事務
func rollbackTransaction(tx *gorm.DB) {
	err := tx.Rollback().Error
	if err != nil && err != gorm.ErrCantStartTransaction {
		global.Logger.Errorf("回滾時發生錯誤：%v", err)
	}
}
