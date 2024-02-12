package cmd

import (
	"os"

	"github.com/aoskainer/cobra-basics/service"
	"github.com/spf13/cobra"
)

var (
	maxNumber int
	rootCmd   = &cobra.Command{
		Use:   "guessgame",
		Short: "数当てゲームです。",
		Long: `ゲームを始めると、プログラムはランダムな数を選び、プレイヤーにその数を推測するように促します。
プレイヤーが数を推測すると、プログラムはその値に応じてヒントを提供します。`,
		Run: func(cmd *cobra.Command, args []string) {
			RootService := service.NewRootService(maxNumber)
			RootService.Init()
			RootService.Play()
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&maxNumber, "max", "m", 100, "最大値を指定します。")
}
