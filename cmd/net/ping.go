package net

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var pingCmd = &cobra.Command{
	Use:     "ping",
	Short:   "ping subcommand of net command",
	Long:    " Ping subcommand",
	Example: "net ping --address example.com",
	Run: func(cmd *cobra.Command, args []string) {
		res, err := ping(address)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		fmt.Println("statusCode: ", res)
	},
}
var address string

func init() {
	pingCmd.Flags().BoolP("help", "h", false, "help for net ping")
	pingCmd.Flags().StringVarP(&address, "address", "a", "google.com", "host or ip address")
	err := pingCmd.MarkFlagRequired("address")
	if err != nil {
		fmt.Println(err)
	}
}

func ping(domain string) (int, error) {
	url := "http://" + domain
	request, err := http.NewRequest("GET", url, nil)

	var client = &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
