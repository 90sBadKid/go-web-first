package initialize

import (
	"go-web/global"
	"go-web/model/book"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

func InitDataBase() {
	db := initDB()
	registerTables(db)
	initDataSource(db)

	global.GVA_DB = db
}

func initDB() (db *gorm.DB) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:sinopacs@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: false,
		},
		// 关闭物理外键
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return db
}

func registerTables(db *gorm.DB) {
	_ = db.AutoMigrate(
		&book.Book{},
	)
}

func initDataSource(db *gorm.DB) {
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
