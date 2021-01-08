package orm

import "time"

//UserInfo 记录key和对应存储的东西
type UserInfo struct {
	ID        int64
	Name      string
	Content   UserContent `xorm:"json"`
	CreatedAt time.Time   `xorm:"created"`
}

//UserContent lottery/hotsearch 不同属性
type UserContent struct {
	Lottery []*UserLottery `xorm:"json" json:"lottery"`
	City    string         `json:"city"`
}

//UserLottery 彩票详情 假设每天都买吧
type UserLottery struct {
	Kind    string `json:"kind"`
	Content string `json:"content"`
}

//Lottery 彩票类 记录每天每种彩票的中奖情况
type Lottery struct {
	ID        int64
	Kind      string
	Number    int
	Content   string
	Pool      int
	Detail    []LotteryDetail
	Date      time.Time `xorm:"date"`
	CreatedAt time.Time `xorm:"created"`
}

//LotteryDetail 中奖详情
type LotteryDetail struct {
	PrizeName    string
	PrizeNum     int
	PrizeAmount  int
	PrizeRequire string
}

//Weather 天气类 记录每天各地区天气？有病吧..这个现查好一点
type Weather struct {
	ID          int64
	City        string    // beijing shanghai guangzhou shenzhen chengdu
	Date        time.Time `xorm:"date"`
	Temperature float32   `xorm:"null"`
	Direction   string    `xorm:"null"`
	Power       string    `xorm:"null"`
	Aqi         int       `xorm:"null"`
	CreatedAt   time.Time `xorm:"created"`
}

//HotSearch 微博热搜
type HotSearch struct {
	ID        int64
	Score     int
	Content   string
	CreatedAt time.Time `xorm:"created"`
}
