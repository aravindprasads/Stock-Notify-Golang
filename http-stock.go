package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetStockDetails(input string) (string, string) {
	url := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + input + "&apikey=demo"
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

	//fmt.Println(result["Global Quote"])

	myMap := result["Global Quote"].(map[string]interface{})
	// stock_name := myMap["01. symbol"]
	// stock_price := myMap["05. price"]
	// fmt.Println(stock_name, stock_price)

	name := fmt.Sprintf("%v", myMap["01. symbol"])
	price := fmt.Sprintf("%v", myMap["05. price"])
	return name, price
	//return "", ""
}

func main() {
	name, price := GetStockDetails("IBM")
	fmt.Println(name, price)
}
