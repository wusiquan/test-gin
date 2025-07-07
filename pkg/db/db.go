package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func InitDbConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("未找到 .env 文件，将使用系统环境变量")
		} else {
			log.Printf("读取 .env 文件失败: %v", err)
		}
		return err
	} else {
		log.Println("已加载 .env 文件")
	}

	// 2. 从环境变量或 .env 文件中获取 APP_ENV
	env := viper.GetString("APP_ENV")

	if env == "" {
		env = "development"
		log.Printf("未设置 APP_ENV, 默认使用: %s", env)
	} else {
		log.Printf("当前环境: %s", env)
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func Connect() {
	dbPath := viper.GetString("db.path")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	// db.Create(&Product{Code: "D42", Price: 100})

}
