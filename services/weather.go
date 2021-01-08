package services

import (
	"errors"
	"iuv520/daily-reminder/orm"
	"log"
	"time"
)

//Weather 天气情况
type Weather struct{}

//Run run
func (w *Weather) Run(userID string) (*orm.Weather, error) {
	engine := orm.GetEngine()
	user := &orm.UserInfo{}
	user.Name = userID
	has, err := engine.Get(user)
	if !has || err != nil {
		log.Printf("user does not exists")
		return nil, err
	}
	weather := &orm.Weather{
		City: user.Content.City,
		Date: time.Now(),
	}

	has, err = engine.Get(weather)

	if !has {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, err
	}
	return weather, nil
}
