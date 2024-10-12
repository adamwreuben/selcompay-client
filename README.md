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

	"github.com/jkarage/selcompay-client"
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

	cln := client.New(host, apiKey, apiSecret)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body := struct {
		Vendor          string `json:"vendor"`
		ID              string `json:"order_id"`
		BuyerEmail      string `json:"buyer_email"`
		BuyerName       string `json:"buyer_name"`
		BuyerPhone      string `json:"buyer_phone"`
		Amount          int    `json:"amount"`
		Currency        string `json:"currency"`
		RedirectURL     string `json:"redirect_url,omitempty"`
		CancelURL       string `json:"cancel_url,omitempty"`
		Webhook         string `json:"webhook,omitempty"`
		BuyerRemarks    string `json:"buyer_remarks,omitempty"`
		MerchantRemarks string `json:"merchant_remarks,omitempty"`
		NumberItems     int    `json:"no_of_items"`
		HeaderColour    string `json:"header_colour,omitempty"`
		LinkColour      string `json:"link_colour,omitempty"`
		ButtonColour    string `json:"button_colour,omitempty"`
		Expiry          int    `json:"expiry,omitempty"`
	}{
		Vendor:      "XXXXXXXXXXXXX",
		ID:          uuid.NewString(),
		BuyerEmail:  "example@gmail.com",
		BuyerName:   "Joseph",
		BuyerPhone:  "255XXXXXXXXX",
		Amount:      1000,
		Webhook:     base64.StdEncoding.EncodeToString([]byte("xxxxxxxxxxxxxx")),
		Currency:    "TZS",
		NumberItems: 1,
	}

	resp, err := cln.CreateOrderMinimal(ctx, client.OrderInputMinimal(body))
	if err != nil {
		return "", err
	}

	fmt.Println(resp)

	return nil
}
```




