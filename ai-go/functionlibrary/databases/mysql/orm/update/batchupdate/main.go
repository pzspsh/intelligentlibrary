/*
@File   : main.go
@Author : pan
@Time   : 2024-11-06 15:48:40
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID   uint
	Name string
	Age  uint
}

func BatchUpdate1() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sql := "UPDATE users SET name = CASE id "
	sql += "WHEN 1 THEN 'NewName1' "
	sql += "WHEN 2 THEN 'NewName2' "
	sql += "WHEN 3 THEN 'NewName3' "
	sql += "ELSE name END WHERE id IN (1, 2, 3)"

	db.Exec(sql)

	fmt.Println("Batch update with SQL expression completed.")
}

func BatchUpdate2() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 假设我们要把所有年龄大于18的用户的名字改为"Adult"
	db.Model(&User{}).Where("age > ?", 18).Update("name", "Adult")

	// 或者使用Updates方法来更新多个字段
	db.Model(&User{}).Where("age > ?", 18).Updates(User{Name: "Adult", Age: 0}) // 注意：这会把年龄设为0，可能不是你想要的
	// 更安全的做法是指定要更新的字段
	db.Model(&User{}).Where("age > ?", 18).Updates(map[string]interface{}{"name": "Adult"})

	fmt.Println("Batch update completed.")
}

func BatchUpdate() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	tx := db.Begin()
	if tx.Error != nil {
		panic("failed to start transaction")
	}

	// 假设我们有一个用户ID和对应新名字的映射
	updates := map[uint]string{
		1: "NewName1",
		2: "NewName2",
		3: "NewName3",
	}

	for id, name := range updates {
		db.Model(&User{}).Where("id = ?", id).Update("name", name)
	}

	tx.Commit()
}

func BatchUpdateUsers(db *gorm.DB, users []User) error {
	if len(users) == 0 {
		return nil
	}

	var ids []interface{}
	var cases []string
	for _, user := range users {
		ids = append(ids, user.ID)
		cases = append(cases, fmt.Sprintf("WHEN %d THEN '%s'", user.ID, user.Name))
	}

	sql := fmt.Sprintf(
		"UPDATE users SET name = CASE %s END WHERE id IN (%s)",
		strings.Join(cases, " "),
		joinInts(ids), // joinInts 是一个将整数数组转换为字符串的函数
	)

	return db.Exec(sql).Error
}

func joinInts(ints []interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	for i, val := range ints {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(fmt.Sprintf("%v", val))
	}
	buffer.WriteString(")")
	return buffer.String()
}

type User1 struct {
	ID     uint   `gorm:"primaryKey"`
	Unique string `gorm:"unique"`
	Field1 string
	Field2 int
	// ... 其他字段
}

func updateBatch(db *gorm.DB, users []User1) {
	for _, user := range users {
		// 假设我们只更新Field1和Field2，其他字段根据实际情况添加
		db.Model(&User{}).Where("id = ?", user.ID).Updates(User1{
			Field1: user.Field1,
			Field2: user.Field2,
			// ... 更新其他字段
		})
	}
}

func BatchUpdate3() {
	// 连接数据库
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 假设我们有一个包含所有要更新数据的切片
	var usersToUpdate []User1
	// ... 填充usersToUpdate切片，这里省略了数据获取的代码

	// 分批更新
	batchSize := 1000 // 每个批次的大小
	var wg sync.WaitGroup
	for i := 0; i < len(usersToUpdate); i += batchSize {
		end := i + batchSize
		if end > len(usersToUpdate) {
			end = len(usersToUpdate)
		}

		wg.Add(1)
		go func(batch []User1) {
			defer wg.Done()
			updateBatch(db, batch)
		}(usersToUpdate[i:end])
	}

	wg.Wait() // 等待所有Goroutine完成
}

func main() {

}
