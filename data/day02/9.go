package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadNet(url string) (ret string, num int) {

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		ret = ""
		num = -100
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		ret = ""
		num = resp.StatusCode
		return
	}

	ret = string(data)
	num = resp.StatusCode
	return
}

func main() {

	ret, num := ReadNet("https://www.baidu.com")

	fmt.Println("ret = ", ret)
	fmt.Println("num = ", num)

}
