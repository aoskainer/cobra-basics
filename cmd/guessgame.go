package cmd

import (
	"github.com/aoskainer/cobra-basics/service"
	"github.com/spf13/cobra"
)

var (
	maxNumber    int
	guessgameCmd = &cobra.Command{
		Use:   "guessgame",
		Short: "数当てゲームです。",
		Long: `ゲームを始めると、プログラムはランダムな数を選び、プレイヤーにその数を推測するように促します。
プレイヤーが数を推測すると、プログラムはその値に応じてヒントを提供します。`,
		Run: func(cmd *cobra.Command, args []string) {
			guessGameService := service.NewGuessGameService(maxNumber)
			guessGameService.Init()
			guessGameService.Play()
		},
	}
)

func init() {
	rootCmd.AddCommand(guessgameCmd)
	guessgameCmd.Flags().IntVarP(&maxNumber, "max", "m", 100, "最大値を指定します。")
}
