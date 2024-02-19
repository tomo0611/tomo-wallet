/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/

package main

import "github.com/tomo0611/tomo-wallet/cmd"

var (
	version  = "UNKNOWN"
	revision = "UNKNOWN"
)

func main() {
	cmd.Version = version
	cmd.Revision = revision

	cmd.Execute()
}
