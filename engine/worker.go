package engine

import (
	"HqCrawler/hqfetcher"
	"log"
)

func Worker(r Request) (ParserResult,error)  {
	log.Printf("Fetching %s\n",r.Url)

	body,err := hqfetcher.HqFetch(r.Url)
	if err != nil {
		log.Printf("Fetch error: %v",err)
		return ParserResult{},err
	}
	return r.ParserFunc(body,r.Url),nil
}