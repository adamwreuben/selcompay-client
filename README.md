# Selcompay Go Client

[![CircleCI](https://dl.circleci.com/status-badge/img/circleci/Cy6tWW4wpE69Ftb8vdTAN9/E2TBj5h2YvKmwX36hcykvy/tree/main.svg?style=svg)](https://app.circleci.com/pipelines/github/Jkarage/selcompay-client)
[![Go Report Card](https://goreportcard.com/badge/github.com/jkarage/selcompay-client)](https://goreportcard.com/report/github.com/jkarage/selcompay-client)
[![go.mod Go version](https://img.shields.io/badge/Go-v1-blue)](https://github.com/Jkarage/selcompay-client)


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

	client "github.com/jkarage/selcompay-client"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	host := "https://apigw.selcommobile.com"
	apiKey := os.Getenv("SELCOM_API_KEY")
    	apiSecret := os.Getenv("SELCOM_SECRET_KEY")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logger := func(ctx context.Context, msg string, v ...any) {
		s := fmt.Sprintf("msg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		log.Println(s)
	}

	cln := client.New(logger, host, apiKey, apiSecret)

	body := client.OrderInputMinimal{
		Vendor:      "TILLXXXXXX",
		ID:          uuid.NewString(),
		BuyerEmail:  "example@gmail.com",
		BuyerName:   "Joseph",
		BuyerPhone:  "255XXXXXXXXX",
		Amount:      1000,
		Webhook:     base64.StdEncoding.EncodeToString([]byte("https://link.com/service")),
		Currency:    "TZS",
		NumberItems: 1,
	}

	resp, err := cln.CreateOrderMinimal(ctx, body)
	if err != nil {
		return "", err
	}

	fmt.Println(resp)

	return nil
}
```




