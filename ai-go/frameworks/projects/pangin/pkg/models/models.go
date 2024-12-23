/*
@File   : models.go
@Author : pan
@Time   : 2023-06-11 22:24:22
*/
package models

import (
	"pangin/pkg/database"
)

type PanDemo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (PanDemo) TableName() string {
	return "pandemo"
}

func CreatePanDemo(pandemo *PanDemo) (err error) {
	if err = database.DB.Create(&pandemo).Error; err != nil {
		return err
	}
	return
}

func GetAllPanDemo() (pandList []*PanDemo, err error) {
	if err := database.DB.Find(&pandList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAPanDemo(id string) (pand *PanDemo, err error) {
	if err := database.DB.Where("id=?", id).First(pand).Error; err != nil {
		return nil, err
	}
	return
}

func UpdatePanDemo(pandemo *PanDemo) (err error) {
	err = database.DB.Save(pandemo).Error
	return
}

func DeletePanDemo(id string) (err error) {
	err = database.DB.Where("id=?", id).Delete(PanDemo{}).Error
	return
}
