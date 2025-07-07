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

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A load generator application",
		Run: func(cmd *cobra.Command, args []string) {
			// 设置日志格式
			logrus.SetFormatter(&logrus.TextFormatter{
				ForceColors:     true,
				FullTimestamp:   true,
				TimestampFormat: "2006-01-02 15:04:05",
			})

			if debug {
				logrus.SetLevel(logrus.DebugLevel)
				logrus.SetReportCaller(true)
			} else {
				logrus.SetLevel(logrus.ErrorLevel)
			}

			composedChain := behaviors.NewChain(behaviors.NewFuncNode(func(ctx *behaviors.Context) (*behaviors.NodeResult, error) {
				return nil, nil
			}, "dummy"))
			composedChain.AddNextChain(behaviors.NormalPreserveChain, 10)
			composedChain.AddNextChain(behaviors.NormalOrderPayChain, 10)
			composedChain.AddNextChain(behaviors.OrderConsignChain, 10)
			composedChain.AddNextChain(behaviors.TicketCollectAndEnterStationChain, 10)

			composedChain.AddNextChain(behaviors.AdvancedSearchChain, 20)
			composedChain.AddNextChain(behaviors.ConsignListChain, 8)
			composedChain.AddNextChain(behaviors.OrderChangeChain, 3)
			composedChain.AddNextChain(behaviors.OrderCancelChain, 2)

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
				logrus.Infof("executed chain %s %d times", chainName, chainCount)
				return
			}

			lg := behaviors.NewLoadGenerator(behaviors.WithThread(threads), behaviors.WithSleep(sleepDuration), behaviors.WithChain(composedChain))
			lg.Start()
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 1, "Number of threads")
	rootCmd.PersistentFlags().IntVarP(&sleepDuration, "sleep", "s", 10000, "Sleep duration in milliseconds")
	rootCmd.PersistentFlags().StringVar(&chainName, "chain", "", "Choose which chain to execute")
	rootCmd.PersistentFlags().IntVar(&chainCount, "count", 1, "How many times to run the chain")

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
