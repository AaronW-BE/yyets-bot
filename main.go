package main

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

type YBot struct {
	client   *http.Client
	cookies  *cookiejar.Jar
	dryRun   bool
	savePath string
}

func (bot *YBot) New() *YBot {
	bot.cookies, _ = cookiejar.New(nil)
	bot.client = &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * 3,
	}
	return bot
}

func (bot *YBot) parse() {

}

func (bot *YBot) login() {

}

func (bot *YBot) logout() {

}

/**
搜索全站资源
*/
func (bot *YBot) seekResource(delayInterval int) {

}

type Resource struct {
	id           string
	title        string   // 标题
	cover        string   // 封面
	coverUri     string   //封面地址
	point        string   // 评分
	premiereDate string   // 首播时间
	detailUri    string   // 详情页url
	tags         []string // 地区
}

func (r *Resource) detail() {

}

func (r *Resource) downloadLinks() {

}

func main() {

}
