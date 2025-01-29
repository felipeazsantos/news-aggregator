package news

import "strings"

func Deduplicate(newsList []News) []News {
	seen := make(map[string]bool)
	var uniqueNews []News

	for _, news := range newsList {
		if _, exists := seen[news.Slug]; !exists {
			seen[news.Slug] = true
			uniqueNews = append(uniqueNews, news)
		}
	}

	return uniqueNews
}

func Categorize(newList []News) map[string][]News {
	categories := make(map[string][]News)

	for _, news := range newList {
		if strings.Contains(strings.ToLower(news.Title), "technology") {
			categories["technology"] = append(categories["technology"], news)
		}
	}

	return categories
}
