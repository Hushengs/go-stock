package stock

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index 首页
func Index(r *gin.Context) {
	r.HTML(http.StatusOK, "stock/index.html", gin.H{})
}

//Detail 详情页
func Detail(r *gin.Context) {
	r.HTML(http.StatusOK, "stock/detail.html", gin.H{})
}
