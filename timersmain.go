package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetStockDetails(input string) (string, string) {
	// url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + input + "&apikey=demo"
	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + input + "&apikey=LAZT6XYGMVEB106M"
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)

	// fmt.Println(result["Global Quote"])
	fmt.Println(result)

	if result["Global Quote"] == nil {
		fmt.Println("No output from site")
		return "", ""
	}

	myMap := result["Global Quote"].(map[string]interface{})
	// stock_name := myMap["01. symbol"]
	// stock_price := myMap["05. price"]
	// fmt.Println(stock_name, stock_price)

	name := fmt.Sprintf("%v", myMap["01. symbol"])
	price := fmt.Sprintf("%v", myMap["05. price"])
	return name, price
	//return "", ""
}

func CheckStockVal(StockName string, Price float32) {
	for _, val := range StockData {
		if val.StockName == StockName {
			//Check the price with min and max
		}
	}
}

func timer_main() {
	fmt.Println("Starting the timer")
	ticker := time.NewTicker(5 * time.Second)
	go func(ticker *time.Ticker) {
		for {
			select {
			case <-ticker.C:
				// do something every 5 minutes as define by the ticker above
				fmt.Println("Checking")
				for _, val := range StockData {
					StockName := "NSE:" + val.StockName
					name, price := GetStockDetails(StockName)
					fmt.Println(name, price)
					price_val, _ := strconv.ParseFloat(price, 32)
					//check current-val with min and max prices
					if name != "" {
						CheckStockVal(name, float32(price_val))
					}
				}
			}
		}
	}(ticker)
}
