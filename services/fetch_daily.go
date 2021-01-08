package services

import (
	"iuv520/daily-reminder/crawer"
	"iuv520/daily-reminder/orm"

	"xorm.io/xorm"
)

//DailyReport 日常报告
type DailyReport struct {
	engine *xorm.Engine
}

//FetchToday 获取今日
func (d *DailyReport) FetchToday() error {
	jhClient := crawer.NewJuheCrawler()
	// 彩票
	cpList := []string{"ssq", "fcsd", "dlt", "qlc", "qxc", "pls", "plw"}
	for _, v := range cpList {
		res, err := jhClient.GetLottery(v)
		if err != nil {
			return err
		}
		has, err := d.engine.Exist(&orm.Lottery{
			Date: res.Date,
			Kind: res.Kind,
		})
		if err != nil {
			return err
		}
		if !has {
			d.engine.Insert(res)
		}
	}
	//天气
	cityList := []string{"北京", "上海", "广州", "深圳", "武汉", "成都", "沈阳"}
	for _, v := range cityList {
		res, err := jhClient.GetWeather(v)
		if err != nil {
			return err
		}
		has, err := d.engine.Exist(&orm.Weather{
			Date: res.Date,
			City: res.City,
		})
		if err != nil {
			return err
		}
		if !has {
			d.engine.Insert(res)
		}
	}

	txClient := crawer.NewTianxingCrawler()
	//热搜
	hotSearch, err := txClient.GetWeiboHotSearch()
	if err != nil {
		return err
	}
	for _, v := range hotSearch {
		d.engine.Insert(v)

	}
	return nil
}

//NewDailyReport 返回dailyReport实例
func NewDailyReport() *DailyReport {
	return &DailyReport{
		engine: orm.GetEngine(),
	}
}
