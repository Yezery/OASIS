package controllers

import (
	"fmt"
	"log"
	"net/http"

	"example.com/m/contracts"
	"example.com/m/utils"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type UserMnemonicController struct{}

// 创建与节点的连接
var bcosClient *client.Client

func (UMC *UserMnemonicController) FiscoConn() error {
	keyName := "manage_key_0xbbcb99c61cd3d3746b8760dfc99ee23f336ef0ea"
	// 解析配置文件
	configs, err := conf.ParseConfigOptions(
		"config/sdk/ca.crt",
		"config/sdk/sdk.key",
		"config/sdk/sdk.crt",
		"config/accounts/"+keyName+".pem",
		1,
		"127.0.0.1:20200",
		false,
		1,
		false,
	)
	// 如解析错误则返回
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 连接节点
	client, err := client.Dial(configs)

	// 如连接错误则返回
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
		return err
	}
	// 把client赋值给全局变量bcosClient
	bcosClient = client

	fmt.Println("connect to local geth node...", client)
	return nil
}

// 实例化合约
func makeUserMnemonicContract() *contracts.UsrMnemonic {
	contractAddress := "0x27a0bf96e92ea917dd7f72797ca93a9457b7375a"
	var UserMnemonicContract *contracts.UsrMnemonic
	if bcosClient != nil {
		UserMnemonicContract, _ = contracts.NewUsrMnemonic(common.HexToAddress(contractAddress), bcosClient)
		return UserMnemonicContract
	}
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

func (UMC *UserMnemonicController) SetMnemonic(c *gin.Context) {
	var userMnemonic userMnemonicDTO
	_ = c.ShouldBind(&userMnemonic) //绑定参数，将参数解析到NFT结构体中
	fmt.Println("========userMnemonic=======")
	fmt.Println(userMnemonic)
	fmt.Println("======================")
	UserMnemonicContract := makeUserMnemonicContract()
	Transaction, Receipt, err := UserMnemonicContract.SetMnemonic(bcosClient.GetTransactOpts(), userMnemonic.Mnemonic, common.HexToAddress(userMnemonic.UserAddress))
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	fmt.Println(Transaction)
	fmt.Println(Receipt)
	utils.SendResponse(c.Writer, http.StatusOK, Receipt)
}

func (UMC *UserMnemonicController) SetAuthenticationMetaInformation(c *gin.Context) {

	var authenticationMetaInformationDTO AuthenticationMetaInformationDTO
	_ = c.ShouldBind(&authenticationMetaInformationDTO) //绑定参数，将参数解析到NFT结构体中
	fmt.Println(authenticationMetaInformationDTO)
	UserMnemonicContract := makeUserMnemonicContract()
	Transaction, Receipt, err := UserMnemonicContract.SetAuthenticationMetaInformation(bcosClient.GetTransactOpts(), authenticationMetaInformationDTO.Sp1, authenticationMetaInformationDTO.Sp2, authenticationMetaInformationDTO.Sp3, common.HexToAddress(authenticationMetaInformationDTO.UserAddress))
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	fmt.Println(Transaction)
	fmt.Println(Receipt)
	utils.SendResponse(c.Writer, http.StatusOK, Receipt)
}

func (UMC *UserMnemonicController) CheckMnemonic(c *gin.Context) {
	var userMnemonic userMnemonicDTO
	_ = c.ShouldBind(&userMnemonic) //绑定参数，将参数解析到NFT结构体中
	UserMnemonicContract := makeUserMnemonicContract()
	Mnemonic, err := UserMnemonicContract.UserMnemonic(bcosClient.GetCallOpts(), common.HexToAddress(userMnemonic.UserAddress))
	if err != nil {
		log.Fatalf("could not connect to local node: %v", err)
	}
	utils.SendResponse(c.Writer, http.StatusOK, Mnemonic)
}

func (UMC *UserMnemonicController) ForgetMnemonic(c *gin.Context) {
	var authenticationMetaInformationDTO AuthenticationMetaInformationDTO
	_ = c.ShouldBind(&authenticationMetaInformationDTO) //绑定参数，将参数解析到NFT结构体中
	UserMnemonicContract := makeUserMnemonicContract()
	code, err := UserMnemonicContract.ForgetMnemonic(bcosClient.GetCallOpts(), authenticationMetaInformationDTO.Sp1, authenticationMetaInformationDTO.Sp2, authenticationMetaInformationDTO.Sp3, common.HexToAddress(authenticationMetaInformationDTO.UserAddress))
	if err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, code)
	}
	fmt.Println(code)
	utils.SendResponse(c.Writer, http.StatusOK, code)
}

func (UMC *UserMnemonicController) ResetMnemonic(c *gin.Context) {
	var authenticationMetaInformationDTO AuthenticationMetaInformationDTO
	_ = c.ShouldBind(&authenticationMetaInformationDTO) //绑定参数，将参数解析到NFT结构体中
	UserMnemonicContract := makeUserMnemonicContract()
	Transaction, _, err := UserMnemonicContract.ReSetMnemonic(bcosClient.GetTransactOpts(), authenticationMetaInformationDTO.Sp1, authenticationMetaInformationDTO.Sp2, authenticationMetaInformationDTO.Sp3, common.HexToAddress(authenticationMetaInformationDTO.UserAddress), authenticationMetaInformationDTO.NewMnemonic)
	if err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, Transaction)
		// log.Fatalf("could not connect to local node: %v", err.Error())

	}
	utils.SendResponse(c.Writer, http.StatusOK, Transaction)
}
