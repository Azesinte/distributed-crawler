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
	re_url      = `https?://[\w/:%#\$&\?\(\)~\.=\+\-]+`
	crawler_url = "https://www.hao123.com"
)

func main() {
	html := GetHtml()
	// fmt.Print(html)
	re_email := regexp.MustCompile(re_email)
	allString_email := re_email.FindAllStringSubmatch(html, -1)

	re_url := regexp.MustCompile(re_url)
	allString_url := re_url.FindAllStringSubmatch(html, -1)

	cnt := 0
	for _, x := range allString_url {
		cnt++
		fmt.Println(x[0])
	}
	fmt.Println("total number of urls is:", cnt)

	cnt = 0
	for _, x := range allString_email {
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
