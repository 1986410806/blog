package database

import (
	"blog/app/models"
	"blog/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type GormModel struct {
	Id int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id" form:"id"`
}

var db *gorm.DB

func OpenDB(url string) (err error) {
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}

	if db, err = gorm.Open("mysql", url); err != nil {
		log.Errorf("opens database failed: %s", err.Error())
		return
	}

	db.LogMode(config.Conf.ShowSql)
	db.SingularTable(true) // 禁用表名负数
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)

	if err = db.AutoMigrate(models.Models...).Error; nil != err {
		log.Errorf("auto migrate tables failed: %s", err.Error())
	}
	return
}

// 获取数据库链接
func DB() *gorm.DB {
	return db
}

// 关闭连接
func CloseDB() {
	if db == nil {
		return
	}
	if err := db.Close(); nil != err {
		log.Errorf("Disconnect from database failed: %s", err.Error())
	}
}

// 事务环绕
func Tx(db *gorm.DB, txFunc func(tx *gorm.DB) error) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = txFunc(tx)
	return err
}
