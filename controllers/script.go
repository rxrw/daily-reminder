package controllers

import (
	"fmt"
	"iuv520/daily-reminder/services"
)

//Script 脚本控制器
type Script struct {
}

//GetDaily 获取每日彩票啥的
func (s *Script) GetDaily() *JSONResponse {

	api := services.NewDailyReport()
	err := api.FetchToday()
	if err != nil {
		fmt.Printf("%v", err)
	}
	return &JSONResponse{
		Code: 0,
		Msg:  "success",
	}
}
