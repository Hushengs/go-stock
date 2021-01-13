package istock

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/djimenez/iconv-go"
	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
)

//Index 指数
func Index(r *gin.Context) {
	urlSh := "http://hq.sinajs.cn/list=s_sh000001"
	urlSz := "http://hq.sinajs.cn/list=s_sz399001"
	var resData string
	var httpCode int
	dataS := make([]string, 0)
	gout.GET(urlSh).BindBody(&resData).Code(&httpCode).Do()
	data := strings.Split(resData, "=")
	dataS = append(dataS, data[1:]...)
	gout.GET(urlSz).BindBody(&resData).Code(&httpCode).Do()
	data = strings.Split(resData, "=")
	dataS = append(dataS, data[1:]...)
	dataReturn := make([]interface{}, 0)
	for _, v := range dataS {
		dataT := make(map[string]interface{}, 0)
		sliceData := strings.Split(v, ",")
		dataT["name"], _ = iconv.ConvertString(sliceData[0], "gbk", "utf-8")
		dataT["name"] = strings.Trim(dataT["name"].(string), `"`)
		dataT["now_point"] = sliceData[1]
		dataT["now_change"] = sliceData[2]
		dataT["now_change_centage"] = sliceData[3]
		dataT["bargain_num"] = sliceData[4]
		idx := strings.Index(sliceData[5], `"`)
		dataT["bargain_total_money"] = sliceData[5][:idx+len([]byte(`"`))-1]
		dataReturn = append(dataReturn, dataT)
	}
	r.JSON(http.StatusOK,
		gin.H{
			"code": 0,
			"msg":  "success",
			"data": dataReturn,
		},
	)
}

//Detail 自选详情
func Detail(r *gin.Context) {
	strCode := "sz002697,sz000876"
	mapCode := make(map[string]int)
	mapCode["sz002697"] = 100
	mapCode["sz000876"] = 200
	url := "http://hq.sinajs.cn/list=" + strCode
	var resData string
	var httpCode int
	gout.GET(url).BindBody(&resData).Code(&httpCode).Do()
	aryData := strings.Split(resData, ";")
	dataS := make([]string, 0)
	// todayTotalProfitAndLoss := 0
	for _, v := range aryData {
		data := strings.Split(v, "=")
		dataS = append(dataS, data[1:]...)
	}
	dataReturn := make([]interface{}, 0)
	for _, v := range dataS {
		dataT := make(map[string]interface{}, 0)
		sliceData := strings.Split(v, ",")
		dataT["name"], _ = iconv.ConvertString(sliceData[0], "gbk", "utf-8")
		dataT["name"] = strings.Trim(dataT["name"].(string), `"`)
		dataT["opening_price"] = sliceData[1]
		dataT["closing_price"] = sliceData[2]
		dataT["now_price"] = sliceData[3]
		dataT["highest_price"] = sliceData[4]
		dataT["lowest_price"] = sliceData[5]
		num1, _ := strconv.ParseFloat(sliceData[3], 64)
		num2, _ := strconv.ParseFloat(sliceData[2], 64)
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num1-num2), 64)
		dataT["change_price"] = value
		if num1 != 0 {
			dataT["range"], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", ((num1-num2)/num1)*100), 64)
		} else {
			dataT["range"] = "0"
		}
		dataT["profit_and_loss"] = 0
		dataReturn = append(dataReturn, dataT)
	}
	r.JSON(http.StatusOK,
		gin.H{
			"code": 0,
			"msg":  "success",
			"data": dataReturn,
		},
	)
}
