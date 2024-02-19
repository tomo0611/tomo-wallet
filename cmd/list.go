/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 変数
var print_json bool
var since, until string
var limit int

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "ボクのお財布の履歴を表示します。",
	Long:  `お財布の履歴表示`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// フラグの値を変数にバインド
	listCmd.Flags().BoolP("json", "j", false, "Print as JSON")
	listCmd.Flags().StringP("since", "s", "", "Start date")
	listCmd.Flags().StringP("until", "u", "", "End date")
	listCmd.Flags().IntP("limit", "l", 10, "Number of records to display")

	// 必須のフラグに指定
	listCmd.MarkFlagRequired("since")
}
