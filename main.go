package main

import (
	"os"

	"github.com/Lincyaw/loadgenerator/behaviors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	// 定义命令行选项变量
	var debug bool
	var threads int
	var sleepDuration int

	// 创建一个新的cobra命令
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A load generator application",
		Run: func(cmd *cobra.Command, args []string) {
			if debug {
				logrus.SetLevel(logrus.DebugLevel)
				logrus.SetReportCaller(true)
			} else {
				logrus.SetLevel(logrus.InfoLevel)
			}

			logrus.SetFormatter(&logrus.TextFormatter{
				FullTimestamp: true,
			})

			lg := &behaviors.LoadGenerator{}
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
			lg = behaviors.NewLoadGenerator(behaviors.WithThread(threads), behaviors.WithSleep(sleepDuration), behaviors.WithChain(composedChain))
			lg.Start()
		},
	}

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
	rootCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 3, "Number of threads")
	rootCmd.PersistentFlags().IntVarP(&sleepDuration, "sleep", "s", 100, "Sleep duration in milliseconds")

	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
