/*
@File   : main.go
@Author : pan
@Time   : 2024-11-08 16:39:17
*/
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"unique"`
	Name  string
	Age   int
}

func AutoCreate(db *gorm.DB, email string, name string, age int) {
	/* 如果你希望在记录不存在时自动插入一条新的记录，并且希望使用默认的字段值，可以使用FirstOrCreate */
	var user User
	db.FirstOrCreate(&user, User{Email: email})
	// 此时user已经是数据库中的记录，如果需要可以进一步更新其他字段
	db.Model(&user).Updates(User{Name: name, Age: age})
}

func AutoUpdate(db *gorm.DB, email string, name string, age int) {
	/* 使用FirstOrUpdate（GORM v2 新增的方法，需确保你使用的GORM版本支持）：
	FirstOrUpdate 方法会查找匹配条件的第一条记录，如果未找到则插入新的记录，如果找到则更新这条记录的字段： */
	// var user User
	// 注意：此方法会在未找到记录时插入新的记录，并更新所有字段
	// db.FirstOrUpdate(&user, User{Email: email, Name: name, Age: age})
	// 如果只想更新部分字段，可以配合使用 Where 和 Updates
	// db.Where("email = ?", email).FirstOrUpdate(&user, User{Name: name, Age: age})
}

func AutoInsert(db *gorm.DB, email string, name string, age int) {
	db.Transaction(func(tx *gorm.DB) error {
		var user User
		if result := tx.First(&user, "email = ?", email); result.Error != nil {
			// 记录不存在，插入新记录
			user = User{Email: email, Name: name, Age: age}
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
		} else {
			// 记录存在，更新记录
			if err := tx.Model(&user).Updates(User{Name: name, Age: age}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	db.AutoMigrate(&User{})

	// 示例操作
	email := "test@example.com"
	name := "Test User"
	age := 30

	var user User
	if result := db.First(&user, "email = ?", email); result.Error != nil {
		// 记录不存在，插入新记录
		user = User{Email: email, Name: name, Age: age}
		db.Create(&user)
	} else {
		// 记录存在，更新记录
		db.Model(&user).Updates(User{Name: name, Age: age})
	}
}
