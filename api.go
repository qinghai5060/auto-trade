package main

import (
	"crypto/hmac"
	"crypto/sha512"
	// "encoding/hex"
	// "encoding/json"
	"net/http"
	// "net/url"
	// "sort"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/golang/glog"
)

type GateApi struct {
	Key    string
	Secret string
}

// all support pairs
func (g *GateApi) getPairs() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/pairs"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Market Info
func (g *GateApi) marketinfo() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/marketinfo"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Market Details
func (g *GateApi) marketlist() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/marketlist"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// tickers
func (g *GateApi) tickers() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/tickers"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// ticker
func (g *GateApi) ticker(ticker string) string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/ticker" + "/" + ticker
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Depth
func (g *GateApi) orderBooks() string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/orderBooks"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Depth of pair
func (g *GateApi) orderBook(params string) string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/orderBook/" + params
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Trade History
func (g *GateApi) tradeHistory(params string) string {
	var method string = "GET"
	var url string = "http://data.gateio.io/api2/1/tradeHistory/" + params
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get account fund balances
func (g *GateApi) balances() string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/balances"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// get deposit address
func (g *GateApi) depositAddress(currency string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/depositAddress"
	var param string = "currency=" + currency
	var ret string = g.httpDo(method, url, param)
	return ret
}

// get deposit withdrawal history
func (g *GateApi) depositsWithdrawals(start string, end string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/depositsWithdrawals"
	var param string = "start=" + start + "&end=" + end
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Place order buy
func (g *GateApi) buy(currencyPair string, rate string, amount string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/buy"
	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Place order sell
func (g *GateApi) sell(currencyPair string, rate string, amount string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/sell"
	var param string = "currencyPair=" + currencyPair + "&rate=" + rate + "&amount=" + amount
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Cancel order
func (g *GateApi) cancelOrder(orderNumber string, currencyPair string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/cancelOrder"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Cancel all orders
func (g *GateApi) cancelAllOrders(types string, currencyPair string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/cancelAllOrders"
	var param string = "type=" + types + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get order status
func (g *GateApi) getOrder(orderNumber string, currencyPair string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/getOrder"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get my open order list
func (g *GateApi) openOrders() string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/openOrders"
	var param string = ""
	var ret string = g.httpDo(method, url, param)
	return ret
}

// 获取我的24小时内成交记录
func (g *GateApi) myTradeHistory(currencyPair string, orderNumber string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/tradeHistory"
	var param string = "orderNumber=" + orderNumber + "&currencyPair=" + currencyPair
	var ret string = g.httpDo(method, url, param)
	return ret
}

// Get my last 24h trades
func (g *GateApi) withdraw(currency string, amount string, address string) string {
	var method string = "POST"
	var url string = "https://api.gateio.io/api2/1/private/withdraw"
	var param string = "currency=" + currency + "&amount=" + amount + "address=" + address
	var ret string = g.httpDo(method, url, param)
	return ret
}

func getSign(secret, params string) string {
	key := []byte(secret)
	mac := hmac.New(sha512.New, key)
	mac.Write([]byte(params))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

/**
*  http request
 */
func (g *GateApi) httpDo(method string, url string, param string) string {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, strings.NewReader(param))
	if err != nil {
		// handle error
	}
	var sign string = getSign(g.Secret, param)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("key", g.Key)
	req.Header.Set("sign", sign)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if resp.StatusCode >300 || resp.StatusCode <200 {
		glog.Errorf("Failed to get response from server. StatusCode :%d",resp.StatusCode)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		glog.Errorf("Failed to get response from GATEIO, error: %s", err.Error())
		return ""
	}

	return string(body)
}
