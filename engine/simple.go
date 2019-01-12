package engine

import (
	"log"
	"reptile/fetcher"
)

type SimpleEngine struct {

}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			log.Printf("fetcher:error " + "fetching url %s:%v", r.Url, err)
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, items := range parseResult.Items {
			log.Printf("got item %v \n", items)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("got url %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error" + "fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
