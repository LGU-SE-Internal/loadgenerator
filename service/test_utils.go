package service

import "github.com/Lincyaw/loadgenerator/httpclient"

type SvcImpl struct {
	cli     *httpclient.HttpClient
	BaseUrl string
}

func NewSvcClients() *SvcImpl {
	cli := httpclient.NewCustomClient()
	cli.AddHeader("Proxy-Connection", "keep-alive")

	cli.AddHeader("Accept", "application/json")
	cli.AddHeader("X-Requested-With", "XMLHttpRequest")
	cli.AddHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	cli.AddHeader("Content-Type", "application/json")
	cli.AddHeader("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	cli.AddHeader("Connection", "keep-alive")
	return &SvcImpl{
		cli:     cli,
		BaseUrl: "http://10.10.10.220:30604",
	}
}
