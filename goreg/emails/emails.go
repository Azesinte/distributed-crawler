package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var (
	re_email    = `[\w\.]+@\w+\.[a-zA-Z]{2,3}(\.[a-zA-Z]{2,3})?`
	crawler_url = "https://tieba.baidu.com/p/6051076813"
)

func main() {
	html := GetHtml()
	fmt.Print(html)
	re := regexp.MustCompile(re_email)
	allString := re.FindAllStringSubmatch(html, -1)

	cnt := 0
	for _, x := range allString {
		cnt++
		fmt.Println(x[0])
	}
	fmt.Println("total number of emails is:", cnt)
}

func DealErr(err error, time string) {
	if err != nil {
		fmt.Println(time, err)
		os.Exit(1)
	}
}

func GetHtml() string {
	fmt.Println("crawler url:", crawler_url)
	resp, err := http.Get(crawler_url)
	DealErr(err, `http.Get`)
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}
