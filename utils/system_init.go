package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

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

	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println("数据库连接失败:", err)
		return
	}
	fmt.Println("数据库连接成功")
}
