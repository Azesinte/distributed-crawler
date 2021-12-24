package DealErr

import (
	"io/ioutil"
	"net/http"
)

func GetHtml() {
	resp, err := http.Get("https://www.haomagujia.com/")
	DealErr(err, `http.Get`)
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}
