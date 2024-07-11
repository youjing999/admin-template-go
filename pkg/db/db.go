package db

import (
	"admin-template-go/common/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func SetupDBLink() error {
	var err error
	dbConfig := config.Config.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)

	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		// NamingStrategy 是用来定义数据库表名和字段名的命名策略的。
		// SingularTable: true 表示在创建表时，GORM会使用单数形式的表名
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
		// 这个设置用于在迁移时禁用外键约束。在数据库迁移过程中，有时需要调整表结构或数据，
		//这可能会违反外键约束。通过设置这个选项为 true，GORM在执行迁移时会自动禁用外键约束，这样可以避免因外键约束而导致的迁移失败。
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}
	sqlDB, err := Db.DB()
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpen)
	return nil
}
