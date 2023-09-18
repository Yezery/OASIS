package main

import (
	"fmt"
	"log"
	"os"

	"example.com/m/controllers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		Username   string            `yaml:"username"`
		Password   string            `yaml:"password"`
		Host       string            `yaml:"host"`
		Port       int               `yaml:"port"`
		Database   string            `yaml:"database"`
		Parameters map[string]string `yaml:"parameters"`
	} `yaml:"database"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

// 连接数据库
func (config *Config) setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)
	for k, v := range config.Database.Parameters {
		dsn += k + "=" + v + "&"
	}
	dsn = dsn[:len(dsn)-1]
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v\n", err)
	}
	return db, nil
}

func getYamlConfig(fileName string) Config {

	dataBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic("读取配置文件失败")
	}

	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		panic("解析 yaml 文件失败")
	}

	return config
}

func main() {

	// 读取配置文件
	config := getYamlConfig("config/config.yaml")
	fmt.Printf("config -> %x", config)

	// 初始化数据库
	db, _ := config.setupDatabase()

	// 将附加到 Gin 的 Context 上下文中，以便在控制器和服务中进行使用。
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, origin, accept")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	SaleController := &controllers.SaleController{}
	SaleTypeController := &controllers.SaleTypeController{}
	NFTOwnerListController := &controllers.NFTOwnerListController{}
	Client := models.NewChatController()
	router.GET("/getSaleList", SaleController.GetSaleList)
	router.GET("/getTypeList", SaleTypeController.GetTypeList)
	router.POST("/getOwnerNFTs", NFTOwnerListController.GetOwnerNFTs)
	router.POST("/getOwnerNFTsByAddress", NFTOwnerListController.GetOwnerNFTsByAddress)
	router.POST("/createSale", SaleController.CreateSale)
	router.POST("/UpdateNFTOwnerList", NFTOwnerListController.UpdateNFTOwnerList)
	router.POST("/UpdateNFTOwnerListAfterBuy", NFTOwnerListController.UpdateNFTOwnerListAfterBuy)
	router.POST("/createNFT", NFTOwnerListController.CreateNFTInf)
	router.POST("/DeleteSale", SaleController.DeleteSale)
	router.POST("/GetOwnerUpSaleNFTs", NFTOwnerListController.GetOwnerUpSaleNFTs)
	router.GET("/OasisChat/:username", Client.WebSocketHandler)
	router.POST("/Search", NFTOwnerListController.Search)
	router.Run(":" + config.Server.Port)
}
