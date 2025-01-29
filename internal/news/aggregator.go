package news

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// News structure to store a news
type News struct {
	Title  string
	URL    string
	Source string
	Slug   string
	Author string
	Image  string
}

// ParseFunc defines the signature of a parse to be able to receive different returns from news APIs
type ParseFunc func(data map[string]interface{}) ([]News, error)

// map of registered parses
var parsers = map[string]ParseFunc{}

// RegisterParser register a parser for a source
func RegisterParser(source string, parser ParseFunc) {
	parsers[source] = parser
}

// FetchFromAPI process in a concurrently way to retrieve news from several APIs
func FetchFromApi(url string, source string, ch chan<- News, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching from %s: %v\n", source, err)
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding JSON from %s: %v\n", source, err)
		return
	}

	parser, ok := parsers[source]
	if !ok {
		fmt.Printf("No parser registered for source: %s\n", source)
		return
	}

	articles, err := parser(data)
	if !ok {
		fmt.Printf("Error parsing data from %s: %v\n", source, err)
		return
	}

	for _, article := range articles {
		ch <- article
	}
}

// AggregateNews aggregates news from multiple sources
func AggregateNews(sources map[string]string) []News {
	var wg sync.WaitGroup
	newsChan := make(chan News, len(sources)*10)

	for source, url := range sources {
		wg.Add(1)
		go FetchFromApi(url, source, newsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(newsChan)
	}()

	var newList []News
	for news := range newsChan {
		newList = append(newList, news)
	}

	return newList
}

// NewsApiParse converts the news api response to news struct
func NewsApiParse(data map[string]interface{}) ([]News, error) {
	return nil, nil
}
