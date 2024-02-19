/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmdはサブコマンドなしでじっこうしたときに呼ばれるコマンドです。
var rootCmd = &cobra.Command{
	Use:   "tomo-wallet",
	Short: "ボクのお財布の履歴を管理するためのCLIツールです。",
	Long:  `ボクのお財布の履歴を管理するためのCLIツールです。`,
	// 実行時に何か処理をさせたければこれをコメントアウト
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Executeは全てのサブコマンドとフラグをルートコマンドに追加する。
// main.main()によって呼び出され、rootCmdに対して一度だけ呼び出される必要がある。
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// フラグと設定を定義する
	// Cobraは、ここで定義されるとアプリケーション内でグローバルになる
	// 永続フラグを使うことができる。

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tomo-wallet.yaml)")

	// Cobraはこのアクションが呼び出された時にだけ実行される
	// ローカルフラグも使うことができる。
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
