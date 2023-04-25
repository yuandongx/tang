package com

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type XStock struct {
	Amount                float64 `json:"amount"`
	Amplitude             float64 `json:"amplitude"`
	Chg                   float64 `json:"chg"`
	Current               float64 `json:"current"`
	Current_year_percent  float64 `json:"current_year_percent"`
	Dividend_yield        float64 `json:"dividend_yield"`
	Eps                   float64 `json:"eps"`
	First_percent         float64 `json:"first_percent"`
	Float_market_capital  float64 `json:"float_market_capital"`
	Float_shares          float64 `json:"float_shares"`
	Followers             float64 `json:"followers"`
	Has_follow            bool    `json:"has_follow"`
	Income_cagr           float64 `json:"income_cagr"`
	Issue_date_ts         float64 `json:"issue_date_ts"`
	Limitup_days          int     `json:"limitup_days"`
	Lot_size              int     `json:"lot_size"`
	Main_net_inflows      float64 `json:"main_net_inflows"`
	Market_capital        float64 `json:"market_capital"`
	Name                  string  `json:"name"`
	Net_profit_cagr       float64 `json:"net_profit_cagr"`
	North_net_inflow      float64 `json:"north_net_inflow"`
	North_net_inflow_time float64 `json:"north_net_inflow_time"`
	Pb                    float64 `json:"pb"`
	Pb_ttm                float64 `json:"pb_ttm"`
	Pcf                   float64 `json:"pcf"`
	Pe_ttm                float64 `json:"pe_ttm"`
	Percent               float64 `json:"percent"`
	Percent5m             float64 `json:"percent5m"`
	Ps                    float64 `json:"ps"`
	Roe_ttm               float64 `json:"roe_ttm"`
	Symbol                string  `json:"symbol"`
	Tick_size             float64 `json:"tick_size"`
	Total_percent         float64 `json:"total_percent"`
	Total_shares          float64 `json:"total_shares"`
	Turnover_rate         float64 `json:"turnover_rate"`
	Type                  int     `json:"type"`
	Volume                float64 `json:"volume"`
	Volume_ratio          float64 `json:"volume_ratio"`
	Date_id               string  `json:"date_id"`
	Is_my_stock           string  `json:"is_my_stock"`
}
type data struct {
	Count    int      `json:"count"`
	List     []XStock `json:"list"`
	MyStocks []string `json:"my_stocks"`
}
type XQResult struct {
	Data             data   `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func makeCookie(s string) (http.Cookie, bool) {
	ss := strings.Split(s, ";")
	c := http.Cookie{}
	flag := false
	for i, s := range ss {
		tmp := strings.Split(s, "=")
		if i == 0 {
			c.Name = tmp[0]
			c.Value = tmp[1]
			flag = true
		} else {
			k := strings.ToLower(tmp[0])
			if k == "path" {
				c.Path = tmp[1]
			} else if k == "httponly" {
				c.HttpOnly = true
			} else if k == "max-age" {
				if i, err := strconv.Atoi(tmp[0]); err == nil {
					c.MaxAge = i
				}
			} else if k == "domain" {
				c.Domain = tmp[1]
			}

		}
	}
	return c, flag
}

func GetXueqiuRank(args map[string]string) (XQResult, bool) {
	result := XQResult{}
	c := &http.Client{}
	query := ""
	urlV := url.Values{}
	for k, v := range args {
		urlV.Add(k, v)
	}
	if len(args) > 0 {
		query = "?" + urlV.Encode()
	}
	request, err := http.NewRequest("GET", "https://xueqiu.com/", nil)
	url1 := "https://stock.xueqiu.com/v5/stock/screener/quote/list.json" + query
	request2, err2 := http.NewRequest("GET", url1, nil)
	if err == nil {
		if respone, err1 := c.Do(request); err1 == nil {
			if item, ok := respone.Header["Set-Cookie"]; ok {
				for _, v := range item {
					if c, ok := makeCookie(v); ok {
						request2.AddCookie(&c)
					}
				}
			}
		}
	}
	if err2 == nil {
		if response, err := c.Do(request2); err == nil {
			defer response.Body.Close()
			if bytes, err := io.ReadAll(response.Body); err == nil {
				err2 := json.Unmarshal(bytes, &result)
				if err2 == nil {
					return result, true
				}
			}
		}
	}
	return result, false
}
