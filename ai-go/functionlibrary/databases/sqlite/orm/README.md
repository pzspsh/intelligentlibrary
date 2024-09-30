# orm连接映射操作数据库-sqlite
# 

```go

package main
 
import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "time"
)
 
type Record struct {
    gorm.Model
    Data string
}
 
func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
 
    // 查询最早的数据
    var earliestRecord Record
    result := db.Order("created_at ASC").Limit(1).Find(&earliestRecord)
    if result.Error != nil {
        panic("failed to query earliest record")
    }
    fmt.Printf("最早的数据: %v\n", earliestRecord)
 
    // 查询最新的数据
    var latestRecord Record
    result = db.Order("created_at DESC").Limit(1).Find(&latestRecord)
    if result.Error != nil {
        panic("failed to query latest record")
    }
    fmt.Printf("最新的数据: %v\n", latestRecord)
}
```