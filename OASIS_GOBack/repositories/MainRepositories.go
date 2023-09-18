package repositories

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDb(c *gin.Context) *gorm.DB {
	return c.MustGet("db").(*gorm.DB)
}

// 包含数据模型的切片作为参数，并自动检查、创建和修改数据库表结构以匹配模型定义。
func MakeDbModel[T any](m *T, c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	db.AutoMigrate(m)
}
