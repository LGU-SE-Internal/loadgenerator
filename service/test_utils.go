package service

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/httpclient"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"sort"
	"time"
)

type SvcImpl struct {
	cli     *httpclient.HttpClient
	BaseUrl string
}

func (s *SvcImpl) ShowStats() {
	// 创建 TUI 界面
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)
	table.SetBackgroundColor(tcell.ColorDefault)
	// 设置表头
	headers := []string{"URL", "Method", "Success", "Failed", "Request Bodies", "Response Bodies"}
	for i, header := range headers {
		table.SetCell(0, i, tview.NewTableCell(header).SetTextColor(tcell.ColorYellow))
	}

	// 填充表格数据
	updateTable := func() {
		statistics := s.cli.GetRequestStats()
		// 先提取并排序 keys
		keys := make([]httpclient.RequestStatsKey, 0, len(statistics))
		for key := range s.cli.GetRequestStats() {
			keys = append(keys, key)
		}

		sort.Slice(keys, func(i, j int) bool {
			if keys[i].URL == keys[j].URL {
				return keys[i].Method < keys[j].Method
			}
			return keys[i].URL < keys[j].URL
		})
		row := 1
		for _, key := range keys {
			stats := statistics[key]
			table.SetCell(row, 0, tview.NewTableCell(key.URL))
			table.SetCell(row, 1, tview.NewTableCell(key.Method))
			table.SetCell(row, 2, tview.NewTableCell(fmt.Sprintf("%d", stats.Success)))
			table.SetCell(row, 3, tview.NewTableCell(fmt.Sprintf("%d", stats.Failed)))
			table.SetCell(row, 4, tview.NewTableCell(fmt.Sprintf("%v", stats.RequestBody)))
			table.SetCell(row, 5, tview.NewTableCell(fmt.Sprintf("%v", stats.ResponseBody)))
			row++
		}
	}

	// 启动一个定时器，每秒更新一次表格
	go func() {
		for {
			time.Sleep(1 * time.Second)
			app.QueueUpdateDraw(updateTable)
		}
	}()
	// 启动 TUI 应用
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
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
		BaseUrl: "http://10.10.10.220:32310",
	}
}
