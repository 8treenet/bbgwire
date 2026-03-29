package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/8treenet/freedom/infra/requests"
	"github.com/PuerkitoBio/goquery"
)

// Article 表示解析出的文章数据
type Article struct {
	Content   string // 内容
	URL       string // 链接
	Published string // 发布时间
}

func main() {
	seq := requests.NewHTTPRequest("https://bbg.buzzing.cc/lite/").Get()
	seq.SetHeaderValue("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	seq.SetHeaderValue("accept-language", "zh-CN,zh;q=0.9")
	seq.SetHeaderValue("sec-ch-ua", `"Chromium";v="146", "Not-A.Brand";v="24", "Google Chrome";v="146"`)
	seq.SetHeaderValue("sec-ch-ua-platform", `"macOS"`)
	seq.SetHeaderValue("sec-fetch-dest", `document`)
	seq.SetHeaderValue("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/146.0.0.0 Safari/537.36`)

	data, resp := seq.ToBytes()
	if resp.Error != nil {
		panic(resp.Error)
	}

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}

	// 存储解析结果
	var articles []Article

	// 使用CSS选择器查找前50个article元素
	doc.Find("article.card.article").Each(func(i int, s *goquery.Selection) {
		if i >= 50 {
			return // 只取前50条
		}

		article := Article{}

		// 1. 解析内容 - 从 a.p-name 元素获取文本
		titleLink := s.Find("a.p-name.entry-title")
		fullText := titleLink.Text()
		// 去掉序号部分，格式如 "1. 中石油利润下滑..."
		parts := strings.SplitN(fullText, ". ", 2)
		if len(parts) == 2 {
			article.Content = strings.TrimSpace(parts[1])
		} else {
			article.Content = strings.TrimSpace(fullText)
		}

		// 2. 解析链接 - 从 a.p-name 元素的 href 属性获取
		href, exists := titleLink.Attr("href")
		if exists {
			article.URL = href
		}

		// 3. 解析时间 - 从 time.dt-published 元素的 datetime 属性获取
		timeElem := s.Find("time.dt-published")
		datetime, exists := timeElem.Attr("datetime")
		if exists {
			article.Published = datetime
		}

		articles = append(articles, article)
	})

	jdata, _ := json.Marshal(articles)
	fmt.Println(string(jdata))
}
