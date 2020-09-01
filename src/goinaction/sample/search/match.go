package search

import "log"

// Result保存搜索结果
type Result struct {
	Field   string
	Content string
}

// Matcher定义了要实现的搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run searches concurrently
// 为每个数据源单独启动goroutine来执行这个函数，并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// write the result to channel
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they
// are received by the individual goroutines
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel
	// 当通道关闭的时候，循环终止
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
