package crawer

import (
	"iuv520/daily-reminder/orm"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

//GetLottery 聚合数据-彩票
func (c *JuheCrawler) GetLottery(lotteryType string) (*orm.Lottery, error) {
	uri := "/lottery/query"
	params := map[string]string{}
	params["lottery_id"] = lotteryType
	params["key"] = os.Getenv("LOTTERY_KEY")
	responseStructure, err := c.crawler.Run(uri, params)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	lotteryStructure := map[string]interface{}{}
	if responseStructure["result"] == nil {
		return nil, nil
	}
	lotteryStructure = responseStructure["result"].(map[string]interface{})
	lottoryNo, _ := strconv.Atoi(lotteryStructure["lottery_no"].(string))
	amount := lotteryStructure["lottery_pool_amount"].(string)
	amount = strings.ReplaceAll(amount, " ", "")
	amount = strings.ReplaceAll(amount, ",", "")
	amount = strings.ReplaceAll(amount, "亿", "00000000")
	amount = strings.ReplaceAll(amount, "万", "0000")
	amountInt, _ := strconv.Atoi(amount)
	lottoryDate, _ := time.Parse("2006-01-02", lotteryStructure["lottery_date"].(string))
	lottoryDetails := []orm.LotteryDetail{}
	for _, row1 := range lotteryStructure["lottery_prize"].([]interface{}) {
		row := row1.(map[string]interface{})
		priceAmount, _ := strconv.Atoi(row["prize_amount"].(string))
		priceNum, _ := strconv.Atoi(row["prize_num"].(string))
		detail := orm.LotteryDetail{
			PrizeName:    row["prize_name"].(string),
			PrizeNum:     priceNum,
			PrizeAmount:  priceAmount,
			PrizeRequire: row["prize_require"].(string),
		}
		lottoryDetails = append(lottoryDetails, detail)
	}
	return &orm.Lottery{
		Kind:    lotteryType,
		Number:  lottoryNo,
		Content: lotteryStructure["lottery_res"].(string),
		Pool:    amountInt,
		Detail:  lottoryDetails,
		Date:    lottoryDate,
	}, nil
}

//GetWeather 获取天气
func (c *JuheCrawler) GetWeather(city string) (*orm.Weather, error) {
	uri := "/simpleWeather/query"
	params := map[string]string{}
	params["city"] = city
	params["key"] = os.Getenv("WEATHER_KEY")
	responseStructure, err := c.crawler.Run(uri, params)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	weatherStructure := responseStructure["result"].(map[string]interface{})
	realTimeStructure := weatherStructure["realtime"].(map[string]interface{})
	temperature, _ := strconv.ParseFloat(realTimeStructure["temperature"].(string), 32)
	aqi, _ := strconv.Atoi(realTimeStructure["aqi"].(string))
	return &orm.Weather{
		City:        city,
		Date:        time.Now(),
		Temperature: float32(temperature),
		Direction:   realTimeStructure["direct"].(string),
		Power:       realTimeStructure["power"].(string),
		Aqi:         aqi,
	}, nil
}

//GetWeiboHotSearch 天行数据-微博热搜
func (c *TianxingCrawler) GetWeiboHotSearch() ([]orm.HotSearch, error) {
	uri := "/txapi/weibohot/index"
	responseStructure, err := c.crawler.Run(uri, *new(map[string]string))
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	hotStructure := responseStructure["newslist"].([]interface{})
	weiboEntities := []orm.HotSearch{}
	for _, item1 := range hotStructure {
		item := item1.(map[string]interface{})
		score, _ := strconv.Atoi(item["hotwordnum"].(string))
		entity := orm.HotSearch{
			Score:   score,
			Content: item["hotword"].(string),
		}
		weiboEntities = append(weiboEntities, entity)
	}
	return weiboEntities, nil
}
