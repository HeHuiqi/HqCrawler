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

		log.Printf("Fetching %s\n",r.Url)

		body,err := hqfetcher.HqFetch(r.Url)
		if err != nil {
			log.Printf("Fetch error: %v",err)
			continue
		}
		paresResuslt := r.ParserFunc(body)
		requests = append(requests,paresResuslt.Requests...)
		for _,item := range paresResuslt.Items {
			log.Printf("Get item %v",item)
		}
	}

}