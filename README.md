# Selcompay Go Client

Copyright 2024 Tausi Africa

### Description

This Module provides functionality developed to simplify interfacing with [SelcomPay API](https://developers.selcommobile.com/) in Go.

### Requirements

To access the API, contact [SelcomPay](https://www.selcom.net/selcom-pay-)

### Usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Jkarage/selcompay-client"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	host := "https://apigw.selcommobile.com"
	apiKey := os.Getenv("SELCOMAPIKEY")
    apiSecret := os.Getenv("SELCOMSECRETKEY")

	cln := client.New(host, apiKey, apiSecret)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := cln.WalletPayment(ctx, "xxxxxxxxxxxxx", "xxxxxxxxxxx", "255xxxxxxxxx")
	if err != nil {
		return fmt.Errorf("ERROR: %w", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil
}
```




