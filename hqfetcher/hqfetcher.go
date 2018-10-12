package hqfetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)
var rateLimiter  = time.Tick(10*time.Millisecond)
func HqFetch(url string) ([]byte,error)  {

	//限速执行
	<-rateLimiter

	req,_:=http.NewRequest(http.MethodGet,url,nil)
	req.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//resp,err :=http.DefaultClient.Do(req)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			fmt.Println("Redict:",req.URL)
			return nil
		},
	}
	resp,err :=client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		fmt.Println("Err:",resp.StatusCode)
		return nil,fmt.Errorf("Wrong status code: %d",resp.StatusCode)
	}

	bd := bufio.NewReader(resp.Body)
	e := determineEncoding(bd)
	//GBK转UTF-8
	//utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(bd,e.NewDecoder())

	all,err := ioutil.ReadAll(utf8Reader)
	return all,err

}
func determineEncoding(r *bufio.Reader) encoding.Encoding  {
	bytes, err := r.Peek(1024)
	if err != nil{
		log.Printf("Fetcher eror: %v",err)
		return unicode.UTF8
	}
	e,_,_ :=charset.DetermineEncoding(bytes,"")
	return e

}