package controllers

import (
	"fmt"
	"log"
	"net/http"

	"example.com/m/contracts"
	"example.com/m/utils"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type UserMnemonicController struct{}
// 实例化合约
func makeUserMnemonicContract(c *gin.Context) *contracts.UsrMnemonic {
	contractAddress := "0xf4a5a8c9898b3070ca874d84e2022c16c47141b1"
	var UserMnemonicContract *contracts.UsrMnemonic
	UserMnemonicContract, _ = contracts.NewUsrMnemonic(common.HexToAddress(contractAddress), c.MustGet("fs").(*client.Client))
	return UserMnemonicContract
}

type userMnemonicDTO struct {
	Mnemonic    string `json:"encryptedPassword"`
	UserAddress string `json:"userAddress"`
}

type AuthenticationMetaInformationDTO struct {
	UserAddress string `json:"userAddress"`
	Sp1         string `json:"sp1"`
	Sp2         string `json:"sp2"`
	Sp3         string `json:"sp3"`
	NewMnemonic string `json:"newMnemonic"`
}

// 设置授权码
func (UMC *UserMnemonicController) SetMnemonic(c *gin.Context) {
	var userMnemonic userMnemonicDTO
	_ = c.ShouldBind(&userMnemonic) //绑定参数，将参数解析到NFT结构体中
	fmt.Println("========userMnemonic=======")
	fmt.Println(userMnemonic)
	fmt.Println("======================")
	UserMnemonicContract := makeUserMnemonicContract(c)
	Transaction, Receipt, err := UserMnemonicContract.SetMnemonic(c.MustGet("fs").(*client.Client).GetTransactOpts(), userMnemonic.Mnemonic, common.HexToAddress(userMnemonic.UserAddress))
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	fmt.Println(Transaction)
	fmt.Println(Receipt)
	utils.SendResponse(c.Writer, http.StatusOK, Receipt)
}

// 设置密保
func (UMC *UserMnemonicController) SetAuthenticationMetaInformation(c *gin.Context) {

	var authenticationMetaInformationDTO AuthenticationMetaInformationDTO
	_ = c.ShouldBind(&authenticationMetaInformationDTO) //绑定参数，将参数解析到NFT结构体中
	fmt.Println(authenticationMetaInformationDTO)
	UserMnemonicContract := makeUserMnemonicContract(c)
	Transaction, Receipt, err := UserMnemonicContract.SetAuthenticationMetaInformation(c.MustGet("fs").(*client.Client).GetTransactOpts(), authenticationMetaInformationDTO.Sp1, authenticationMetaInformationDTO.Sp2, authenticationMetaInformationDTO.Sp3, common.HexToAddress(authenticationMetaInformationDTO.UserAddress))
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	fmt.Println(Transaction)
	fmt.Println(Receipt)
	utils.SendResponse(c.Writer, http.StatusOK, Receipt)
}

// 检查授权码
func (UMC *UserMnemonicController) CheckMnemonic(c *gin.Context) {
	var userMnemonic userMnemonicDTO
	_ = c.ShouldBind(&userMnemonic) //绑定参数，将参数解析到NFT结构体中
	UserMnemonicContract := makeUserMnemonicContract(c)
	Mnemonic, err := UserMnemonicContract.UserMnemonic(c.MustGet("fs").(*client.Client).GetCallOpts(), common.HexToAddress(userMnemonic.UserAddress))
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	utils.SendResponse(c.Writer, http.StatusOK, Mnemonic)
}

// 重置授权码
func (UMC *UserMnemonicController) ForgetMnemonic(c *gin.Context) {
	var authenticationMetaInformationDTO AuthenticationMetaInformationDTO
	_ = c.ShouldBind(&authenticationMetaInformationDTO) //绑定参数，将参数解析到NFT结构体中
	UserMnemonicContract := makeUserMnemonicContract(c)
	code, err := UserMnemonicContract.ForgetMnemonic(c.MustGet("fs").(*client.Client).GetCallOpts(), authenticationMetaInformationDTO.Sp1, authenticationMetaInformationDTO.Sp2, authenticationMetaInformationDTO.Sp3, common.HexToAddress(authenticationMetaInformationDTO.UserAddress))
	if err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, code)
	}
	fmt.Println(code)
	utils.SendResponse(c.Writer, http.StatusOK, code)
}

func (UMC *UserMnemonicController) ResetMnemonic(c *gin.Context) {
	var authenticationMetaInformationDTO AuthenticationMetaInformationDTO
	_ = c.ShouldBind(&authenticationMetaInformationDTO) //绑定参数，将参数解析到NFT结构体中
	UserMnemonicContract := makeUserMnemonicContract(c)
	Transaction, _, err := UserMnemonicContract.ReSetMnemonic(c.MustGet("fs").(*client.Client).GetTransactOpts(), authenticationMetaInformationDTO.Sp1, authenticationMetaInformationDTO.Sp2, authenticationMetaInformationDTO.Sp3, common.HexToAddress(authenticationMetaInformationDTO.UserAddress), authenticationMetaInformationDTO.NewMnemonic)
	if err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, Transaction)
		// log.Fatalf("could not connect to local node: %v", err.Error())

	}
	utils.SendResponse(c.Writer, http.StatusOK, Transaction)
}
