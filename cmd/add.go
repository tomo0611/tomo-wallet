/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 時間 店舗名 金額 その取引終了時の残高 どの財布からの支出か メモ

var datetime, shop_name, amount, balance, memo string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "ボクのお財布に履歴を追加します。",
	Long:  `お財布に履歴を追加します。`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// フラグを追加
	addCmd.Flags().StringVarP(&datetime, "datetime", "d", "", "payment date. format is 'YYYYMMDD'")
	addCmd.Flags().StringVarP(&shop_name, "shop", "s", "", "shop name")
	addCmd.Flags().StringVarP(&amount, "amount", "a", "", "amount")
	addCmd.Flags().StringVarP(&balance, "balance", "b", "", "balance")
	addCmd.Flags().StringVarP(&memo, "memo", "m", "", "memo")

	// 必須のフラグに指定
	addCmd.MarkFlagRequired("date")
	addCmd.MarkFlagRequired("shop")
	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagRequired("balance")
	addCmd.MarkFlagRequired("wallet")
}
