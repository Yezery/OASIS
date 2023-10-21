package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"example.com/m/utils"
	"github.com/gin-gonic/gin"
)

type GPTController struct{}
type Message struct {
	
	Role    string `json:"role"`
	Content string `json:"content"`
}
type MakePayload struct {
	Model      string     `json:"model"`
	Messages   [1]Message `json:"messages"`
	Max_tokens int64      `json:"max_tokens"`
}

var APIKey = "Bearer sk-sXW9AAjZrbPg2zZbRw2pyMVQjCWVtGD6g5NsIkWByAfjW3si"

func (STG *GPTController) SendToGPT(c *gin.Context) {

	var messager Message
	var messageArray [1]Message
	if err := c.BindJSON(&messager); err != nil {
		utils.SendResponse(c.Writer, http.StatusBadRequest, err)
		panic(err)
	}
	messageArray[0] = messager
	data := MakePayload{
		"gpt-3.5-turbo",
		messageArray,
		500,
	}

	url := "https://api.chatanywhere.com.cn/v1/chat/completions"
	method := "POST"
	payload, _ := json.MarshalIndent(data, "", " ")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payload)))

	if err != nil {
		fmt.Println(err)
		utils.SendResponse(c.Writer, http.StatusInternalServerError, err)
		return
	}
	req.Header.Add("Authorization", APIKey)
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		utils.SendResponse(c.Writer, http.StatusInternalServerError, err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		utils.SendResponse(c.Writer, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(string(body))
	utils.SendResponse(c.Writer, http.StatusOK, string(body))
}
