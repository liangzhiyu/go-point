package point

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func SendData() {
	var s [100]int
	for i := 0; i < 100; i++ {
		s[i] = i
	}

	go printData(s[0:25])
	go printData(s[26:50])
	go printData(s[51:75])
	go printData(s[75:99])
	time.Sleep(1 * time.Second)
}

func printData(s []int) {
	for i, name := range s {
		if name==98{
			println("xxxxxxxxx",i)
		}
		println(i, name)
	}
}

func HttpGet() {
	response, err := http.Get("https://www.baidu.com")
	if  err != nil{
		println(err)
	}
	defer response.Body.Close()
	body, _:=ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}