package main

import "github.com/felipeazsantos/news-aggregator/internal/news"

func main() {
	news.RegisterParser("TBD", news.NewsApiParse)

}
