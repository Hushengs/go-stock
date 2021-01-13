package router

import (
	"go-stock/admin/stock"
	istock "go-stock/interface/stock"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("index", stock.Index)
		admin.GET("detail", stock.Detail)
	}

	api := r.Group("/api")
	{
		api.GET("index", istock.Index)
		api.GET("detail", istock.Detail)
	}
}
