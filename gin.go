package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetPage(c *gin.Context, DefaultPageSize int) int {
	result := 0
	page := cast.ToInt(c.Query("page"))
	if page > 0 {
		result = (page - 1) * DefaultPageSize
	}

	return result
}
