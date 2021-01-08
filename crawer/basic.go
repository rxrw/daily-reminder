package crawer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//Crawler 通用接口
type Crawler interface {
	Init(key string, url string) error
}

//GeneralCrawler 通用爬取接口
type GeneralCrawler struct {
	key string
	url string
}

//NewGeneralCrawler 通用调用工具
func NewGeneralCrawler(key string, url string) *GeneralCrawler {
	return &GeneralCrawler{
		key: key,
		url: url,
	}
}

//Init function
func (c *GeneralCrawler) Init(key string, url string) error {
	c.key = key
	c.url = url
	return nil
}

//Run 跑
func (c *GeneralCrawler) Run(uri string, params map[string]string) (map[string]interface{}, error) {
	client := &http.Client{}
	requestParams := url.Values{}

	v, exists := params["key"]
	if !exists || v == "" {
		requestParams.Add("key", c.key)
	}
	for k, v := range params {
		requestParams.Add(k, v)
	}
	httpURL, _ := url.Parse(fmt.Sprintf("%s%s", c.url, uri))
	httpURL.RawQuery = requestParams.Encode()

	res, err := client.Get(httpURL.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	returned, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	responseStructure := make(map[string]interface{})

	err = json.Unmarshal(returned, &responseStructure)
	if err != nil {
		return nil, err
	}
	return responseStructure, nil
}

//JuheCrawler 聚合数据接口
type JuheCrawler struct {
	crawler *GeneralCrawler
}

//NewJuheCrawler 返回聚合实例
func NewJuheCrawler() *JuheCrawler {
	return &JuheCrawler{crawler: NewGeneralCrawler("", "http://apis.juhe.cn")}
}

//TianxingCrawler 天行数据接口
type TianxingCrawler struct {
	crawler *GeneralCrawler
}

//NewTianxingCrawler 返回天行实例
func NewTianxingCrawler() *TianxingCrawler {
	return &TianxingCrawler{crawler: NewGeneralCrawler(os.Getenv("TIANXING_KEY"), "http://api.tianapi.com")}
}
