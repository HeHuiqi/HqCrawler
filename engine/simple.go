package engine

import (
	"log"
)

type SimpleEngine struct {

}


func (e SimpleEngine)Run(seeds ...Request)  {

	var requests []Request

	for _,r := range seeds {
		requests = append(requests,r)
	}
	itemCount := 0
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]


		paresResult,err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,paresResult.Requests...)
		itemCount++
		for _,item := range paresResult.Items {

			log.Printf("Got item #%d: %v",itemCount,item)
		}
	}

}
