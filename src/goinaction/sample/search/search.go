package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

// Run performs the search logic
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// create an unbuffered channel
	// 创建一个无缓冲的通道，接受匹配后的结果
	results := make(chan *Result)

	// Setup a wait group，以便处理所有数据源
	var waitGroup sync.WaitGroup
	// 设置程序处理feeds（数据源）时需要等待的goroutines数量
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results
	for _, feed := range feeds {
		// Retrieve a matcher（匹配器） for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search
		// 启动一个goroutine来执行搜索
		// feed是一个指针变量
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			// 递减WaitGroup的计数
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor
	// 启动一个goroutine来监控所有工作是否都完成了
	go func() {
		// 等待其它goroutine处理完
		// 阻塞goroutine，直到waitGroup内部的计数到达0
		waitGroup.Wait()

		// 用关闭通道的方式，通知Display函数
		close(results)
	}()

	// 显示返回结果，并在最后一个结果显示完后返回
	Display(results)
}

// 注册一个matcher，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher已经注册")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
