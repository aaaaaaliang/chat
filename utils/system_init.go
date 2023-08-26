package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	ctx := context.Background()
	pong, err := Red.Ping(ctx).Result()
	if err != nil {
		fmt.Println("初始化 Redis 失败", err)
	} else {
		fmt.Println("redis 初始化成功", pong)
	}

}
func InitConfig() {
	// 设置配置的名字是 app
	viper.SetConfigName("app")
	// 设置配置的路径
	viper.AddConfigPath("config")

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件错误:", err)
	}

	// 打印配置内容
	fmt.Println("config app 里的内容", viper.Get("app"))
}

func InitMySQL() {
	//自定义日志 打印sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阀值
			LogLevel:      logger.Info, //级别
			Colorful:      true,
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	fmt.Println("数据库连接成功")
}
