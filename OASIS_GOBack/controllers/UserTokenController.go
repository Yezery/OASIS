package controllers

import (
	"fmt"
	"net/http"
	"time"

	"example.com/m/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type UserTokenController struct{}

func (UTC *UserTokenController) GetToken(c *gin.Context) {
	type User struct {
		Address           string `json:"userAddress"`
		EncryptedPassword string `json:"encryptedPassword"`
	}
	var user User
	if err := c.BindJSON(&user); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	EncryptedPassword, err := makeUserMnemonicContract().UsrMnemonicCaller.UserMnemonic(bcosClient.GetCallOpts(), common.HexToAddress(user.Address))
	if EncryptedPassword == user.EncryptedPassword {
		fmt.Println(EncryptedPassword == user.EncryptedPassword)
		// 生成JWT令牌
		jwtKey := []byte("0xbbcb99c61cd3d3746b8760dfc99ee23f336ef0ea")
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["Address"] = user.Address
		claims["EncryptedPassword"] = user.EncryptedPassword

		// 设置其他需要的声明

		// 设置令牌的有效期
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// 签名令牌
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		fmt.Println("===========tokenString===============")
		fmt.Println(tokenString)
		fmt.Println(EncryptedPassword)
		fmt.Println(user.EncryptedPassword)
		fmt.Println("==========================")
		utils.SendResponse(c.Writer, http.StatusOK, tokenString)
	} else {
		utils.SendResponse(c.Writer, http.StatusBadGateway, err)
	}

}
