package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/faizmokhtar/whatsmyip/model"
	"github.com/spf13/cobra"
)

// ipv4Cmd represents the ipv4 command
var ipv4Cmd = &cobra.Command{
	Use:   "ipv4",
	Short: "get the IPv4 address",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://api.ipify.org?format=json")
		if err != nil {
			fmt.Println("unable to retrieve ip address: ", err)
			return
		}
		var address model.Address
		d := json.NewDecoder(resp.Body)
		if err := d.Decode(&address); err != nil {
			fmt.Println("error: something's not right.")
		}

		fmt.Printf("Your public IPv4 address is %s\n", address.IP)
	},
}

func init() {
	rootCmd.AddCommand(ipv4Cmd)
}
