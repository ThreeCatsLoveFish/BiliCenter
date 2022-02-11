package pull

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"subcenter/manager"
	"subcenter/service/push"
	"time"

	"github.com/gookit/config/v2"
)

const (
	AStockEnv = "A_STOCK"

	EastMoney = "EastMoney"
)

// initPush bind endpoints with config file
func initStock() {
	pushConf := config.NewWithOptions("stock", func(opt *config.Options) {
		opt.ParseEnv = true
	})
	pushConf.LoadOSEnv([]string{AStockEnv}, false)
	stock := pushConf.Get(AStockEnv).(string)
	stockList := strings.Split(stock, ",")
	addPull(EastMoney, NewEastMoneyPull(stockList))
}

type StockData struct {
	Code      string  `json:"f57"`  // 股票代码
	Name      string  `json:"f58"`  // 股票名称
	Latest    int32   `json:"f43"`  // 最新价
	Highest   int32   `json:"f44"`  // 最高价
	Lowest    int32   `json:"f45"`  // 最低价
	Open      int32   `json:"f46"`  // 今开价
	Volume    int32   `json:"f47"`  // 成交量
	Value     float32 `json:"f48"`  // 成交额
	Ratio     int32   `json:"f50"`  // 量比
	LimitUp   int32   `json:"f51"`  // 涨停
	LimitDown int32   `json:"f52"`  // 跌停
	PrevClose int32   `json:"f60"`  // 昨收价
	Turnover  int32   `json:"f168"` // 换手率
}

func (data StockData) rate() float64 {
	return float64(data.Latest-data.PrevClose) / float64(data.PrevClose) * 100.0
}

func (data StockData) open() float64 {
	return float64(data.Open) / 100.0
}

func (data StockData) now() float64 {
	return float64(data.Latest) / 100.0
}

func (data StockData) high() float64 {
	return float64(data.Highest) / 100.0
}

func (data StockData) low() float64 {
	return float64(data.Lowest) / 100.0
}

func (data StockData) title() string {
	return fmt.Sprintf("# %s %s", data.Name, data.Code)
}

func (data StockData) content() string {
	return fmt.Sprintf(
		"时间: %s\n\n最新价: %.2f 涨跌幅: %.2f%%\n\n今开价: %.2f 最高价: %.2f 最低价: %.2f",
		time.Now().In(location).Format(time.RFC1123Z),
		data.now(), data.rate(), data.open(), data.high(), data.low(),
	)
}

type EastMoneyPull struct {
	stockList []string
}

func NewEastMoneyPull(stockList []string) EastMoneyPull {
	return EastMoneyPull{stockList}
}

func (EastMoneyPull) getData(secId string) (*StockData, error) {
	rawUrl := "https://push2.eastmoney.com/api/qt/stock/get"
	params := url.Values{
		"fields": []string{"f43,f44,f45,f46,f47,f48,f50,f51,f52,f57,f58,f60,f168"},
		"secid":  []string{secId},
	}
	data, err := manager.GetWithParams(rawUrl, params)
	if err != nil {
		// FIXME: add log here
		return nil, err
	}
	type resBody struct {
		Data StockData `json:"data"`
	}
	var b resBody
	err = json.Unmarshal(data, &b)
	if err != nil {
		// FIXME: add log here
		return nil, err
	}
	return &b.Data, nil
}

func (pull EastMoneyPull) Obtain() ([]push.Data, error) {
	var data []push.Data
	for _, stock := range pull.stockList {
		stock, err := pull.getData(stock)
		if err != nil {
			// FIXME: add log here
			continue
		}
		data = append(data, push.Data{
			Title:   stock.title(),
			Content: stock.content(),
		})
	}
	return data, nil
}
