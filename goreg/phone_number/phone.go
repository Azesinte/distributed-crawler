package main

import (
	"fmt"
	"regexp"
)

var (
	rePhone = `1[3456789]\d{9}`
)

func main() {
	html := GetHtml()
	re := regexp.MustCompile(rePhone)
	allString := re.FindAllStringSubmatch(html, -1)

	cnt := 0
	for _, x := range allString {
		cnt++
		fmt.Println(x[0])
	}
	fmt.Println("total number of phone is:", cnt)
}

func DealErr(err error, time string) {
	if err != nil {
		fmt.Println(time, err)
		os.Exit(1)
	}
}

func GetHtml() string {
	resp, err := http.Get("https://www.haomagujia.com/")
	DealErr(err, `http.Get`)
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}
