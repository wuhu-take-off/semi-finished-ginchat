package sql

import (
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func ConnectMysql() *gorm.DB {
	//连接数据库

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info, //级别
			Colorful: true,        //彩色
		},
	)
	//mysql数据库连接
	db, err := gorm.Open(mysql.Open("root:112154ZhouM..@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("filed to connect mysql")
	}
	db.AutoMigrate(&models.UserBasic{})

	//println(db.First(user, 1))

	//db.Model(user).Update("Password", "1234")
	return db
}
