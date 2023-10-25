package config

import (
	"fmt"
	"log"
	"os"

	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Mysql struct {
		Username   string            `yaml:"username"`
		Password   string            `yaml:"password"`
		Host       string            `yaml:"host"`
		Port       int               `yaml:"port"`
		Database   string            `yaml:"database"`
		Parameters map[string]string `yaml:"parameters"`
	} `yaml:"mysql"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Fisco struct {
		Cacert   string `yaml:"cacert"`
		Key      string `yaml:"key"`
		Cert     string `yaml:"cert"`
		Account  string `yaml:"account"`
		GroupId  int    `yaml:"groupid"`
		NodeURL  string `yaml:"nodeurl"`
		Http     bool   `yaml:"http"`
		ChainId  int64  `yaml:"chainid"`
		SMCrypto bool   `yaml:"smcrypto"`
	} `yaml:"fisco"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// 读取配置
func (config *Config) GetYamlConfig(fileName string) {
	dataBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Panicf("\x1b[%d;%dm ERROR \x1b[0m Read config file error \n", 41, 36)
	}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		log.Panicf("\x1b[%d;%dm ERROR \x1b[0m Yaml parsing error \n", 41, 36)
	}
	log.Printf("\x1b[%d;%dm SUCCESS \x1b[0m Success to read yaml config file \n", 42, 35)
}

// 连接Mysql数据库
func (config *Config) SetupMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.Database)
	for k, v := range config.Mysql.Parameters {
		dsn += k + "=" + v + "&"
	}
	dsn = dsn[:len(dsn)-1]
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("\x1b[%d;%dm ERROR \x1b[0m Failed to connect MySQL \n", 41, 36)
	}
	log.Printf("\x1b[%d;%dm SUCCESS \x1b[0m Success to connect MySQL \n", 42, 35)
	return db, nil
}

// 连接Redis数据库
func (config *Config) SetupRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Panicf("\x1b[%d;%dm ERROR \x1b[0m Redis connect ping failed \n", 41, 36)
	}
	log.Printf("\x1b[%d;%dm SUCCESS \x1b[0m Success to connect Redis \n", 42, 35)
	return client, nil
}

// 连接fisco节点
func (config *Config) SetupFisco() (*client.Client, error) {
	// 解析配置文件
	configs, err := conf.ParseConfigOptions(
		config.Fisco.Cacert,
		config.Fisco.Key,
		config.Fisco.Cert,
		config.Fisco.Account,
		config.Fisco.GroupId,
		config.Fisco.NodeURL,
		config.Fisco.Http,
		config.Fisco.ChainId,
		config.Fisco.SMCrypto,
	)
	// 如解析错误则返回
	if err != nil {
		log.Panicf("\x1b[%d;%dm ERROR \x1b[0m Fisco-Bcos ParseConfigOptions error \n", 41, 36)
	}
	// 连接节点
	client, err := client.Dial(configs)

	// 如连接错误则返回
	if err != nil {
		log.Panicf("\x1b[%d;%dm ERROR \x1b[0m Could not connect to FISCO-BCOS node \n", 41, 36)
	}

	log.Printf("\x1b[%d;%dm SUCCESS \x1b[0m Success to connect FISCO-BCOS node \n", 42, 35)
	return client, err
}
