package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	const token = "TOKEN"

	if len(os.Args) != 2 {
		fmt.Printf("USAGE: iptracker [ip]\n")
		os.Exit(0)
	}
	client := ipinfo.NewClient(nil, nil, token)

	ip_address := os.Args[1]
	info, err := client.GetIPInfo(net.ParseIP(ip_address))
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	infoBytes, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(infoBytes, &result); err != nil {
		log.Fatal(err)
	}
	for key, value := range result {
		fmt.Printf("%s: %+v\n", key, value)
	}
}
