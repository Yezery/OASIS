package utils

import (
	"encoding/json"
	"net/http"
)

// Result 结构体定义包含要发送到客户端的数据和元数据。
type Result struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// 定义一个辅助函数 SendResponse，该函数将接受 code 和 data 参数并将其封装为JSON格式通过HTTP响应返回给客户端。
func SendResponse(w http.ResponseWriter, code int, data interface{}) {
	result := Result{
		StatusCode: code,
		Data:       data,
	}

	// 如果code指示操作失败，则在结果中添加一条错误消息。
	if code != http.StatusOK {
		result.Message = "Operation failed. Please try again later."
	}

	// 指定响应类型并将数据编码为JSON格式。
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
