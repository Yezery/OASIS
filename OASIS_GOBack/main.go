package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"example.com/m/controllers"
	"example.com/m/models"
	"github.com/dgrijalva/jwt-go"
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

// 读取配置
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

// 跨域中间件
func Cors3() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的Authorization字段
		tokenString := c.GetHeader("Authorization")
		fmt.Println(tokenString)
		// 检查Token是否存在
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 提取Token中的实际令牌部分
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// 设置JWT令牌的签名密钥
		jwtKey := []byte("0xbbcb99c61cd3d3746b8760dfc99ee23f336ef0ea")

		// 解析和验证令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// 将令牌信息存储到上下文中，以便后续处理函数使用
		// _, _ := token.Claims.(jwt.MapClaims)
		// c.Set("username", claims["username"])
		c.Next()
	}
}

func main() {

	// 读取配置文件
	config := getYamlConfig("config/config.yaml")
	fmt.Printf("config -> %x\n", config)

	// 初始化数据库
	db, _ := config.setupDatabase()

	// 将附加到 Gin 的 Context 上下文中，以便在控制器和服务中进行使用。
	router := gin.Default()
	//
	router.Use(Cors3())
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	UserMnemonicController := &controllers.UserMnemonicController{}
	SaleController := &controllers.SaleController{}
	SaleTypeController := &controllers.SaleTypeController{}
	NFTOwnerListController := &controllers.NFTOwnerListController{}
	UserTokenController := &controllers.UserTokenController{}
	// 连接FiscoBcos
	UserMnemonicController.FiscoConn()
	router.POST("/checkMnemonic", UserMnemonicController.CheckMnemonic)

	router.POST("/forgetMnemonic", UserMnemonicController.ForgetMnemonic)

	// 开启WebSocket协议
	Client := models.NewChatController()

	// 路由接口
	router.POST("/getToken", UserTokenController.GetToken)
	router.GET("/getSaleList", SaleController.GetSaleList)
	router.GET("/getTypeList", SaleTypeController.GetTypeList)
	router.POST("/setMnemonic", UserMnemonicController.SetMnemonic)
	router.POST("/Search", NFTOwnerListController.Search)

	jwtGroup := router.Group("/", JWT())
	jwtGroup.POST("/getOwnerNFTs", NFTOwnerListController.GetOwnerNFTs)
	jwtGroup.POST("/getOwnerNFTsByAddress", NFTOwnerListController.GetOwnerNFTsByAddress)
	jwtGroup.POST("/createSale", SaleController.CreateSale)
	jwtGroup.POST("/UpdateNFTOwnerList", NFTOwnerListController.UpdateNFTOwnerList)
	jwtGroup.POST("/UpdateNFTOwnerListAfterBuy", NFTOwnerListController.UpdateNFTOwnerListAfterBuy)
	jwtGroup.POST("/createNFT", NFTOwnerListController.CreateNFTInf)
	jwtGroup.POST("/DeleteSale", SaleController.DeleteSale)
	jwtGroup.POST("/GetOwnerUpSaleNFTs", NFTOwnerListController.GetOwnerUpSaleNFTs)
	jwtGroup.POST("/GetSaleListByContractAddress", SaleController.GetSaleListByContractAddress)
	router.GET("/OasisChat/:username", Client.WebSocketHandler)

	jwtGroup.POST("/setAuthenticationMetaInformation", UserMnemonicController.SetAuthenticationMetaInformation)

	// router.RunTLS(":"+config.Server.Port, "./config/cert.pem", "./config/key.pem")
	router.Run(":" + config.Server.Port)
}
