package crawel

import (
	"container/list"
	"crawel/dao"
	"crawel/global"
	"crawel/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 基于goquery爬虫页面实现

func Getartices(url string) {
	articleList := list.New()
	info := &model.ArticleInfo{}
	// 请求
	req, err := http.NewRequest("GET", url, nil)
	global.CheckErr(err)
	req.Header.Add("user-agent", "")
	client := &http.Client{}
	resp, err := client.Do(req)
	global.CheckErr(err)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	global.CheckErr(err)
	doc.Find(".container .post-preview").Each(func(i int, s *goquery.Selection) {
		info.Name = s.Find("a").Text()
		info.Url, _ = s.Find("a").Attr("href")
		if !strings.Contains(info.Url, "http") {
			info.Url = "http://gityuan.com" + info.Url
		}
		dao.InsertData(info)
		articleList.PushBack(info)
		fmt.Println(articleList.Len())
	})
	fmt.Println(articleList)

}

func GetarticesForCSDN(url string) {
	info := &model.ArticleInfo{}
	req, err := http.NewRequest("GET", url, nil)
	global.CheckErr(err)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	global.CheckErr(err)
	doc, err := goquery.NewDocumentFromResponse(resp)
	global.CheckErr(err)
	fmt.Println(doc)
	doc.Find("main .article-list .article-item-box").Each(func(i int, s *goquery.Selection) {
		info.Name = s.Find("h4").Find("a").Text()
		span := s.Find("h4").Find("a").Find("span").Text()
		info.Url, _ = s.Find("h4").Find("a").Attr("href")
		// fmt.Println(strings.TrimSpace(info.Name))
		//fmt.Println(info.Url)
		// fmt.Println(span)
		// 判断文章辩题是否有<span>的内容，若有则去点
		if strings.HasPrefix(strings.TrimSpace(info.Name), span) {
			info.Name = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(info.Name), span))
		}
		dao.InsertData(info)

	})

}

