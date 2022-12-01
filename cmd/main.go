package main

import (
	"embed"
	"encoding/json"
	"findregion/dto"
	"fmt"
	"net"
	"os"
)

//go:embed ip-ranges.json
var content embed.FS

func main() {
	ipOrHost := os.Args[1]
	var ipAddress net.IP

	addr := net.ParseIP(ipOrHost)
	if addr != nil {
		ipAddress = addr
	} else {
		host := ipOrHost
		resolvedIP, err := net.ResolveIPAddr("ip", host)

		if err != nil {
			panic(err)
		}

		ipAddress = resolvedIP.IP
	}

	textBytes, _ := content.ReadFile("ip-ranges.json")

	lookupResult := dto.LookupResult{
		ResolvedIP: string(ipAddress.String()),
		Region:     "Unknown",
	}
	var result dto.IPRanges
	err := json.Unmarshal(textBytes, &result)

	if err != nil {
		panic(err)
	}

	for _, prefix := range result.Prefixes {
		_, ipnet, err := net.ParseCIDR(prefix.IPPrefix)

		if err != nil {
			panic(err)
		}

		if ipnet.Contains(ipAddress) {
			lookupResult.Region = prefix.Region
			break
		}
	}
	fmt.Printf("Region: %s\n", lookupResult.Region)
	fmt.Printf("IP: %s\n", lookupResult.ResolvedIP)
}
