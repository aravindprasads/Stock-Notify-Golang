package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Stock struct {
	UserName     string
	StockName    string
	Actual_value float32
	Min_value    float32
	Max_value    float32
	Email_id     string
}

var StockData []Stock
var DummyUserName = "dummy"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("stocks.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	//fmt.Println(StockData)
	err = t.Execute(w, StockData) //execute the template with filled data
	if err != nil {               // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	StockName := r.Form.Get("stock-name")
	MinVal := r.Form.Get("min-value")
	MaxVal := r.Form.Get("max-value")
	MailID := r.Form.Get("mail-id")

	MinValFloat, _ := strconv.ParseFloat(MinVal, 32)
	MaxValFloat, _ := strconv.ParseFloat(MaxVal, 32)

	//fmt.Println(StockName, MinValFloat, MaxValFloat, MailID)

	d := Stock{
		UserName:     DummyUserName,
		StockName:    StockName,
		Actual_value: 0,
		Min_value:    float32(MinValFloat),
		Max_value:    float32(MaxValFloat),
		Email_id:     MailID,
	}

	StockData = append(StockData, d)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	StockName := r.Form.Get("stock-name-del")
	// fmt.Println(StockName)

	del_index := -1
	for index, val := range StockData {
		if val.StockName == StockName && val.UserName == DummyUserName {
			fmt.Println("index to be deleted is ", index)
			del_index = index
		}
	}
	if del_index != -1 {
		if del_index == 0 {
			_, StockData = StockData[0], StockData[1:]
		} else if del_index == len(StockData)-1 {
			len1 := len(StockData)
			StockData = StockData[:len1-1]
		} else {
			StockData = append(StockData[:del_index], StockData[del_index:]...)
		}
	}

	// fmt.Println(StockData)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
