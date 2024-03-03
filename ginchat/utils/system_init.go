package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func InitMySQL() {
	//自定义日志模板,打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	//mysql数据库连接
	db, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
}

func GetDB() *gorm.DB {
	return db
}

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})

	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	} else {
		println("redis connect success!!", result)
	}
}
func GetRedis() *redis.Client {
	return rdb
}
func CloseRedis() {
	rdb.Close()
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到redis
func Publish(channel, msg string) error {
	return rdb.Publish(context.Background(), channel, msg).Err()
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, channel string) (*redis.Message, error) {
	pubsub := rdb.Subscribe(ctx, channel)
	defer pubsub.Close()
	return pubsub.ReceiveMessage(ctx)
}
