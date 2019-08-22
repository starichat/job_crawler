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
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	global.CheckErr(err)
	doc, err := goquery.NewDocumentFromResponse(resp)
	global.CheckErr(err)
	// doc.Find(".my_tab_page_con .tab_page_list dt").Each(func(i int, s *goquery.Selection) {
	// 	info.Name = s.Find("h3").Text()
	// 	info.Url, _ = s.Find("h3").Find("a").Attr("href")
	// 	fmt.Println(strings.TrimSpace(info.Name))
	// 	dao.InsertData(info)

	// })
	doc.Find(".container .clearfix .pt0 .article-list .article-item-box .csdn-tracking-statistics .article-item-box .csdn-tracking-statistics").Each(func(i int, s *goquery.Selection) {
		info.Name = s.Find("h4").Find("a").Text()
		fmt.Println(info.Name)
	})
}
