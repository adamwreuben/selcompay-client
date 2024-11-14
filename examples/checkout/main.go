package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	client "github.com/Golang-Tanzania/selcompay-client"
	"github.com/google/uuid"
)

func main() {
	host := os.Getenv("SELCOM_HOST")
	apiKey := os.Getenv("SELCOM_API_KEY")
	apiSecret := os.Getenv("SELCOM_API_SECRET")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	logger := func(ctx context.Context, msg string, v ...any) {
		s := fmt.Sprintf("msg: %s", msg)
		for i := 0; i < len(v); i = i + 2 {
			s = s + fmt.Sprintf(", %s: %v", v[i], v[i+1])
		}
		log.Println(s)
	}

	cln := client.New(logger, host, apiKey, apiSecret)

	input := client.OrderInput{
		Vendor:               "TILL61073446",
		ID:                   uuid.NewString(),
		BuyerEmail:           "josephbkarage@gmail.com",
		BuyerName:            "Joseph Karage",
		BuyerPhone:           "255713507067",
		Amount:               5000,
		Currency:             "TZS",
		PaymentMethods:       "MOBILEMONEYPULL",
		BillingFirstName:     "Joseph",
		BillingLastName:      "Karage",
		BillingAddress1:      "Africana",
		BillingCity:          "Dar es salaam",
		BillingStateRegion:   "Tanzania",
		BillingPostCodePOBox: "00000",
		BillingCountry:       "TZ",
		BillingPhone:         "255713507067",
		NumberItems:          1,
		Webhook:              "https://play.svix.com/memesample",
	}

	resp, err := cln.CreateOrder(ctx, input)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(resp)
}
