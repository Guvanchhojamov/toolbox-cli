package net

import (
	"github.com/spf13/cobra"
)

var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "Net operations commands",
	Long:  "For use net operations use net commands",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	NetCmd.Flags().BoolP("help", "h", true, "show this help message")
	NetCmd.AddCommand(pingCmd)
}
