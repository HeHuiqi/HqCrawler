package engine

import (
	"HqCrawler/hqfetcher"
	"log"
)

func Run(seeds ...Request)  {

	var requests []Request

	for _,r := range seeds {
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]


		paresResult,err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,paresResult.Requests...)
		for _,item := range paresResult.Items {
			log.Printf("Get item %v",item)
		}
	}

}
func worker(r Request) (ParserResult,error)  {
	log.Printf("Fetching %s\n",r.Url)

	body,err := hqfetcher.HqFetch(r.Url)
	if err != nil {
		log.Printf("Fetch error: %v",err)
		return ParserResult{},err
	}
	return r.ParserFunc(body),nil
}