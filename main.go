package main

import (
	"context"
	"os"

	"github.com/Lincyaw/loadgenerator/behaviors"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var chains = map[string]*behaviors.Chain{
	"NormalPreserveChain":               behaviors.NormalPreserveChain,
	"NormalOrderPayChain":               behaviors.NormalOrderPayChain,
	"OrderConsignChain":                 behaviors.OrderConsignChain,
	"TicketCollectAndEnterStationChain": behaviors.TicketCollectAndEnterStationChain,
	"AdvancedSearchChain":               behaviors.AdvancedSearchChain,
	"ConsignListChain":                  behaviors.ConsignListChain,
	"OrderChangeChain":                  behaviors.OrderChangeChain,
	"OrderCancelChain":                  behaviors.OrderCancelChain,
	// Admin chains
	"AdminBasicInfoChain": behaviors.AdminBasicInfoChain,
	"AdminOrderChain":     behaviors.AdminOrderChain,
	"AdminRouteChain":     behaviors.AdminRouteChain,
	"AdminTravelChain":    behaviors.AdminTravelChain,
	"AdminUserChain":      behaviors.AdminUserChain,
}

func callChain(chain *behaviors.Chain, count int) {
	client := service.NewSvcClients()
	defer client.CleanUp()

	chainCtx := behaviors.NewContext(context.Background())
	chainCtx.Set(behaviors.Client, client)
	for i := 0; i < count; i++ {
		chain.Execute(chainCtx)
	}
}

func getChainByName(name string) *behaviors.Chain {
	return chains[name]
}

func main() {
	var debug bool
	var threads int
	var sleepDuration int
	var chainName string
	var chainCount int
	var showStats bool

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A load generator application",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&logrus.TextFormatter{
				ForceColors:     true,
				FullTimestamp:   true,
				TimestampFormat: "2006-01-02 15:04:05",
			})

			if debug {
				logrus.SetLevel(logrus.DebugLevel)
				logrus.SetReportCaller(true)
			} else {
				logrus.SetLevel(logrus.WarnLevel)
			}

			composedChain := behaviors.NewChain(behaviors.NewFuncNode(func(ctx *behaviors.Context) (*behaviors.NodeResult, error) {
				return nil, nil
			}, "dummy"))
			// 用户行为 chains (总计 90%)
			composedChain.AddNextChain(behaviors.NormalPreserveChain, 20)               // 预订票务 - 最常见操作
			composedChain.AddNextChain(behaviors.NormalOrderPayChain, 15)               // 订单支付
			composedChain.AddNextChain(behaviors.AdvancedSearchChain, 18)               // 高级搜索 - 用户经常查询
			composedChain.AddNextChain(behaviors.TicketCollectAndEnterStationChain, 12) // 取票进站
			composedChain.AddNextChain(behaviors.OrderConsignChain, 8)                  // 订单托运
			composedChain.AddNextChain(behaviors.ConsignListChain, 6)                   // 托运列表查询
			composedChain.AddNextChain(behaviors.OrderChangeChain, 4)                   // 改签 - 较少
			composedChain.AddNextChain(behaviors.OrderCancelChain, 2)                   // 退票 - 最少

			// Admin chains (总计 10%) - 管理员操作较少
			composedChain.AddNextChain(behaviors.AdminBasicInfoChain, 3) // 基础信息管理
			composedChain.AddNextChain(behaviors.AdminOrderChain, 3)     // 订单管理
			composedChain.AddNextChain(behaviors.AdminTravelChain, 3)    // 行程管理
			composedChain.AddNextChain(behaviors.AdminRouteChain, 3)     // 路线管理
			composedChain.AddNextChain(behaviors.AdminUserChain, 3)      // 用户管理

			if chainName != "" {
				chain := getChainByName(chainName)
				if chain == nil {
					availableChains := make([]string, 0, len(chains))
					for name := range chains {
						availableChains = append(availableChains, name)
					}
					logrus.Errorf("chain %s not found", chainName)
					logrus.Infof("available chains: %v", availableChains)
					return
				}
				callChain(chain, chainCount)
				return
			}

			lg := behaviors.NewLoadGenerator(behaviors.WithThread(threads), behaviors.WithSleep(sleepDuration), behaviors.WithChain(composedChain))
			lg.Start()
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 1, "Number of threads")
	rootCmd.PersistentFlags().IntVarP(&sleepDuration, "sleep", "s", 1000, "Sleep duration in milliseconds")
	rootCmd.PersistentFlags().StringVar(&chainName, "chain", "", "Choose which chain to execute")
	rootCmd.PersistentFlags().IntVar(&chainCount, "count", 1, "How many times to run the chain")
	rootCmd.PersistentFlags().BoolVar(&showStats, "stats", false, "Show current latency statistics")

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
