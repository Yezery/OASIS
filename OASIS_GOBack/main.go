package main

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/m/config"
	"example.com/m/controllers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

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
	gin.SetMode(gin.ReleaseMode)
	config := &config.Config{}
	// 读取配置文件
	config.GetYamlConfig("config/config.yaml")
	// 初始化Mysql数据库
	db, _ := config.SetupMysql()
	// 初始化Redis数据库
	rd, _ := config.SetupRedis()
	fs, _ := config.SetupFisco()
	// 将附加到 Gin 的 Context 上下文中，以便在控制器和服务中进行使用。
	router := gin.Default()
	//
	router.Use(Cors3())
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("rd", rd)
		c.Set("fs", fs)
	})
	UserMnemonicController := &controllers.UserMnemonicController{}
	SaleController := &controllers.SaleController{}
	NFTOwnerListController := &controllers.NFTOwnerListController{}
	UserTokenController := &controllers.UserTokenController{}
	TransactionController := &controllers.TransactionController{}
	GPTController := &controllers.GPTController{}
	// 公开路由
	// 查
	router.GET("/getSaleList", SaleController.GetSaleList)
	router.POST("/queryMnemonic", UserMnemonicController.CheckMnemonic)
	router.POST("/getOnSaleNFTByNFTAddress", SaleController.GetOnSaleNFTByNFTAddress)
	router.POST("/getToken", UserTokenController.GetToken)
	router.POST("/getSeriesByNFTAddress", NFTOwnerListController.GetSeriesByNFTAddress)
	router.POST("/Search", NFTOwnerListController.Search)
	router.POST("/scheduleDailySummary", TransactionController.ScheduleDailySummary)
	router.POST("/mainSearch", NFTOwnerListController.MainSearch)
	// 增
	router.POST("/setMnemonic", UserMnemonicController.SetMnemonic)
	router.POST("/setAuthenticationMetaInformation", UserMnemonicController.SetAuthenticationMetaInformation)
	// 改
	router.POST("/resetMnemonic", UserMnemonicController.ResetMnemonic)
	router.POST("/forgetMnemonic", UserMnemonicController.ForgetMnemonic)

	// 授权路由
	jwtGroup := router.Group("/", JWT())
	router.POST("/sendToGPT", GPTController.SendToGPT)
	// 查
	jwtGroup.POST("/getOwnerNFTs", NFTOwnerListController.GetOwnerNFTs)
	jwtGroup.POST("/getOwnerNFTsByAddress", NFTOwnerListController.GetOwnerNFTsByAddress)
	jwtGroup.POST("/GetOwnerUpSaleNFTs", NFTOwnerListController.GetOwnerUpSaleNFTs)

	// 增
	jwtGroup.POST("/makeNewTransaction", TransactionController.MakeNewTransaction)
	jwtGroup.POST("/createNFT", NFTOwnerListController.CreateNFTInf)
	jwtGroup.POST("/createSale", SaleController.CreateSale)
	// 改
	jwtGroup.POST("/UpdateNFTOwnerList", NFTOwnerListController.UpdateNFTOwnerList)
	jwtGroup.POST("/UpdateNFTOwnerListAfterBuy", NFTOwnerListController.UpdateNFTOwnerListAfterBuy)
	// 删
	jwtGroup.POST("/DeleteSale", SaleController.DeleteSale)

	router.Run(":" + config.Server.Port)
}
