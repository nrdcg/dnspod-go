# DNSPod Go API client

[![Build Status](https://github.com/nrdcg/dnspod-go/workflows/Main/badge.svg?branch=master)](https://github.com/nrdcg/dnspod-go/actions)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/nrdcg/dnspod-go)](https://pkg.go.dev/github.com/nrdcg/dnspod-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/nrdcg/dnspod-go)](https://goreportcard.com/report/github.com/nrdcg/dnspod-go)

A Go client for the DNSPod API.

Originally inspired by [dnsimple](https://github.com/weppos/dnsimple-go/dnsimple)

## Getting Started

This library is a Go client you can use to interact with the DNSPod API.

```go
package main

import (
	"fmt"
	"log"

	"github.com/nrdcg/dnspod-go"
)

func main() {
	apiToken := "xxxxx"

	params := dnspod.CommonParams{LoginToken: apiToken, Format: "json"}
	client := dnspod.NewClient(params)

	// Get a list of your domains
	domains, _, _ := client.Domains.List()
	for _, domain := range domains {
		fmt.Printf("Domain: %s (id: %s)\n", domain.Name, domain.ID)
	}

	// Get a list of your domains (with error management)
	domains, _, err := client.Domains.List()
	if err != nil {
		log.Fatalln(err)
	}
	for _, domain := range domains {
		fmt.Printf("Domain: %s (id: %s)\n", domain.Name, domain.ID)
	}

	// Create a new Domain
	newDomain := dnspod.Domain{Name: "example.com"}
	domain, _, _ := client.Domains.Create(newDomain)
	fmt.Printf("Domain: %s\n (id: %s)", domain.Name, domain.ID)
}
```

## API documentation

- https://www.dnspod.cn/docs/index.html
- https://docs.dnspod.com/api-legacy/
- https://docs.dnspod.com/api/

## License

This is Free Software distributed under the MIT license.
