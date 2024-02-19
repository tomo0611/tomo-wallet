/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// アプリケーションのバージョンとリビジョン情報
var (
	Version  string
	Revision string
)

var (
	wallet string
	// configFile 設定ファイルyamlのパス
	configFile string
	// c 設定
	// c Config
)

// rootCmdはサブコマンドなしでじっこうしたときに呼ばれるコマンドです。
var rootCmd = &cobra.Command{
	Use:   "tomo-wallet",
	Short: "ボクのお財布の履歴を管理するためのCLIツールです。",
	Long:  `ボクのお財布の履歴を管理するためのCLIツールです。`,
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

	// Pありは通常のフラグ名(例.--name)だけでなく、省略形のフラグ名(例.-n)も引数として指定する
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&wallet, "wallet", "w", "", "wallet name")

	rootCmd.MarkFlagRequired("config")
}
