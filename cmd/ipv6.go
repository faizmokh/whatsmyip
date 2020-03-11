package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/faizmokhtar/whatsmyip/model"
	"github.com/spf13/cobra"
)

// ipv6Cmd represents the ipv6 command
var ipv6Cmd = &cobra.Command{
	Use:   "ipv6",
	Short: "get the IPv6 address",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://api6.ipify.org?format=json")
		if err != nil {
			fmt.Println("unable to retrieve ip address: ", err)
			return
		}

		var address model.Address
		d := json.NewDecoder(resp.Body)
		if err := d.Decode(&address); err != nil {
			fmt.Println("error: something's not right.")
		}

		fmt.Printf("Your public IPv6 address is %s\n", address.IP)
	},
}

func init() {
	rootCmd.AddCommand(ipv6Cmd)
}
