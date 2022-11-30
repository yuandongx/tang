package api

import (
	"fmt"
	"tang/api/com"
)

// func postXxx(c *Context)   {}
// func getXxx(c *Context)    {}
// func deleteXxx(c *Context) {}
// func patchXxx(c *Context)  {}

func postStock(c *Context) {
	var stock com.XStock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.String(504, err.Error())
		return
	}
	err := mgdb.Save("tang", stock.Symbol, stock)
	if err != nil {
		c.String(504, err.Error())
		return
	}
	fmt.Println(stock)
	c.String(200, "ok")
}

// func getStock(c *Context)    {}
// func deleteStock(c *Context) {}
// func patchStock(c *Context)  {}

func getXueqiuRank(c *Context) {
	args := make(map[string]string)
	args["page"] = c.DefaultQuery("page", "1")
	args["size"] = c.DefaultQuery("size", "30")
	args["order"] = c.DefaultQuery("order", "desc")
	args["order_by"] = c.DefaultQuery("order_by", "percent")
	args["exchange"] = c.DefaultQuery("exchange", "CN")
	args["market"] = c.DefaultQuery("market", "CN")
	args["type"] = c.DefaultQuery("type", "sha")
	if result, ok := com.GetXueqiuRank(args); ok {
		c.JSON(200, result)
	} else {
		c.String(501, "Error to fetch result from `xueqiu`.")
	}
}
