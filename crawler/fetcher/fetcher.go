package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"practice/crawler_distributed/config"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.QPS)

func Fetch(url string) ([]byte, error) {

	<-rateLimiter

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: status code %v \n", resp.StatusCode)
		return nil, fmt.Errorf("wrong status code: %d \n", resp.StatusCode)
	}

	//encode:=determineEncoding(resp.Body)
	//encodeReader:=transform.NewReader(resp.Body,
	//	encode.NewDecoder())
	//all, err := ioutil.ReadAll(encodeReader)

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

//检测编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error; %v \n", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
