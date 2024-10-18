# gorm 更新

```go
if err := db.Model(&model.tableObj{}).Where("number = ?", number).Updates(map[string]interface{}{
							"column1": gorm.Expr("column1 + ?", value), 
                            "column2_date": utils.NowTime(),
                            }).Error; err != nil {
							logger.Error("column1 update +1  error:%v", err)
						}
```