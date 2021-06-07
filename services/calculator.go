package services

import (
	"iuv520/daily-reminder/orm"
	"log"
	"time"
)

//Calculator 计算中奖了没中了多少钱
type Calculator struct {
}

//Run 拿用户的数和系统同日前一天的数进行比对
func (c *Calculator) Run(userID string) (map[string]int, map[string][]string, error) {
	engine := orm.GetEngine()
	user := &orm.UserInfo{}
	user.Name = userID
	has, err := engine.Get(user)
	if !has || err != nil {
		log.Printf("user does not exists")
		return nil, nil, err
	}
	today := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	lotteries := make([]orm.Lottery, 0)
	err = engine.Where("date = ?", today).Find(&lotteries)
	if err != nil {
		log.Printf("lottery does not exists")
	}

	userLotteryList := make(map[string][]orm.UserLottery)

	userContentLottery := user.Content.Lottery
	//把user买的彩票列表按票种排序
	for _, content := range userContentLottery {
		userLotteryList[content.Kind] = append(userLotteryList[content.Kind], *content)
	}

	res := make(map[string]int)
	typeNumbers := make(map[string][]string)

	for _, content := range lotteries {
		v, ok := userLotteryList[content.Kind]
		//存在的话 将中奖值和用户买的列表传到下一个数组
		if ok {
			result, numbers := c.isHeWin(content.Content, v, content.Kind)
			res[content.Kind] = result
			typeNumbers[content.Kind] = numbers
		}
	}

	return res, typeNumbers, nil

}

func (c *Calculator) isHeWin(correct string, buied []orm.UserLottery, kind string) (int, []string) {
	total := 0
	winned := make([]string, 0)
	for _, userNumber := range buied {
		res := c.cal(correct, userNumber.Content, kind)
		if res == -1 {
			winned = append(winned, userNumber.Content)
			return -1, winned
		}
		if res != 0 {
			total = total + res
			winned = append(winned, userNumber.Content)
		}

	}
	return total, winned
}

func (c *Calculator) cal(correct string, current string, kind string) int {
	lotteryInstance := NewLottery(correct, current, kind)
	return lotteryInstance.AmIWin()
}
