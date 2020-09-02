package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/LeisureGong/gogo/src/goinaction/sample/search"
	"log"
	"net/http"
	"regexp"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// 实现Matcher接口
type rssMatcher struct{}

// 在文档中查找特定的搜索项
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {

	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] Foe Uri[%s]\n", feed.Type, feed.Name, feed.URI)

	// 获取要搜索的数据
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 如果找到匹配的项， 将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field:   "标题",
				Content: channelItem.Title,
			})
		}

		// 检查描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field:   "正文描述",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}

// 将匹配器注册打动程序里
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// 发送HTTP Get请求获取rss数据源并解码
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// 	从网络获取rss数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// check http resp code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP返回错误 %d \n", resp.StatusCode)
	}

	// 将rss数据源文档解码到定义的数据结构类型里
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
