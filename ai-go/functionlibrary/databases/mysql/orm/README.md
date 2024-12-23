# orm映射操作数据库-mysql
Golang使用gorm操作mysql数据表时，经常需要自动设置创建和修改时间

以下几种方式

## 1.使用默认支持字段

```pgsql
CreatedAt  time.Time `gorm:"column:created_at;" json:"created_at"` // 创建时间
UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`  // 更新时间
```

这是官方文档说明：

> GORM 约定使用 `CreatedAt`、`UpdatedAt` 追踪创建/更新时间。如果您定义了这种字段，GORM 在创建、更新时会自动填充当前时间

## 2.自定义字段

使用 **autoCreateTime** 和 **autoUpdateTime** 标签

```pgsql
Created    time.Time `gorm:"column:created;autoCreateTime" json:"created"` // 自动创建当前时间
Updated    time.Time `gorm:"column:updated;autoUpdateTime" json:"updated"` // 更新时间
```

## 3.设置写入权限

数据库设置 **CURRENT_TIMESTAMP**

```pgsql
`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
```

结构体添加不可写权限 **<-:false**

```pgsql
CreatedAt  time.Time `gorm:"column:created_at;<-:false" json:"created_at"` // 创建时间
UpdatedAt  time.Time `gorm:"column:updated_at;<-:false" json:"updated_at"` // 更新时间
```

## 4.使用 GORM 内置的gorm.Model结构体[不太推荐,不方便设置标签信息]

```elm
// gorm.Model 定义 
type Model struct { 
    ID uint `gorm:"primary_key"` 
    CreatedAt time.Time 
    UpdatedAt time.Time
    DeletedAt *time.Time
}
type User struct { 
    gorm.Model 
    Name string 
}
```

这是官方文档说明：

> - **将其嵌入在您的结构体中**: 您可以直接在您的结构体中嵌入 `gorm.Model` ，以便自动包含这些字段。 这对于在不同模型之间保持一致性并利用GORM内置的约定非常有用，请参考[嵌入结构](https://link.segmentfault.com/?enc=vjk0uzg%2FiRyL3nun5h9KpQ%3D%3D.8ncGUOHQ25ZqIcWRcCdpB1IghtBX5beSiKbW9QbtlIwXN1oeq%2FBJSwMhsZbJRMJliUksL48DlNK%2F6wFDwuUMNQ%3D%3D)。
>
> - - **包含的字段**：

```autohotkey
-   `ID` ：每个记录的唯一标识符（主键）。
-   `CreatedAt` ：在创建记录时自动设置为当前时间。
-   `UpdatedAt`：每当记录更新时，自动更新为当前时间。
-   `DeletedAt`：用于软删除（将记录标记为已删除，而实际上并未从数据库中删除）。
```
