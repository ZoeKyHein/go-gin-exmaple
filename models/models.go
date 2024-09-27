package models

import (
	"fmt"
	"github.com/ZoeKyHein/go-gin-example/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"` // 主键
	CreatedOn  int `json:"created_on"`            // 创建时间
	ModifiedOn int `json:"modified_on"`           // 修改时间
}

// 用于初始化数据库
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()              // 数据库类型
	dbName = sec.Key("NAME").String()              // 数据库名称
	user = sec.Key("USER").String()                // 用户名
	password = sec.Key("PASSWORD").String()        // 密码
	host = sec.Key("HOST").String()                // 主机
	tablePrefix = sec.Key("TABLE_PREFIX").String() // 表前缀

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))
	if err != nil {
		log.Fatal(2, "Fail to open database: %v", err)
	}
	// 设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)       // 禁用表名复数
	db.LogMode(true)             // 打印sql语句
	db.DB().SetMaxIdleConns(10)  // 设置空闲连接池中连接的最大数量
	db.DB().SetMaxOpenConns(100) // 设置数据库的最大打开连接数
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
