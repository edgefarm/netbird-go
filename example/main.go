package main

import (
	"fmt"
	"time"

	"github.com/edgefarm/netbird-go"
)

var (
	token = "your personal access token"
)

func main() {

	client := netbird.NewClient(token)
	toCreate := &netbird.SetupKey{
		Name: "test",
		Type: "reusable",
		ExpiresIn: func(d time.Duration) int {
			return int(d.Seconds())
		}(31536000 * time.Second),
		Revoked:    false,
		UsageLimit: 0,
		Ephemeral:  false,
	}
	fmt.Println(toCreate)
	new, err := client.CreateSetupKey(toCreate)

	fmt.Println(new)
	if err != nil {
		panic(err)
	}
	p, err := client.ListSetupKeys()
	if err != nil {
		panic(err)
	}
	for _, k := range p {
		fmt.Printf("SetupKeys: %v\n", k)
	}
}
