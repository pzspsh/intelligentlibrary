/*
@File   : main.go
@Author : pan
@Time   : 2024-10-18 16:44:08
*/
package main

import (
	"gorm.io/gorm"
)

const (
	table1 = "test_table"
)

type TestTable struct {
	Name string `gorm:"column:name"`
	Age  int    `gorm:"column:age"`
}

func (TestTable) TableName() string {
	return table1
}

func CreateScopesDemo(db *gorm.DB) error {
	var err error
	data := &TestTable{
		Name: "pang",
		Age:  28,
	}
	/*
		if err = db.Scopes(db.Table("表名称")).Create("表示数据").Error; err != nil {
			return err
		}
	*/
	if err = db.Scopes(ObjTableName(table1)).Create(data).Error; err != nil {
		return err
	}
	return err
}

func UpdateScopesDemo(db *gorm.DB, name string) error {
	var err error
	if err = db.Scopes(ObjTableName(table1)).Where("name = ?", name).Update("age", 27).Error; err != nil {
		return err
	}
	return err
}

func ObjTableName(table string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(table)
	}
}

func main() {

}
