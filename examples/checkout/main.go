package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	client "github.com/jkarage/selcompay-client"
)

func main() {
	host := os.Getenv("SELCOM")
	apiKey := os.Getenv("SELCOM_APIKEY")
	apiSecret := os.Getenv("SELCOM_APISECRET")

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

	input := struct {
		Vendor                string `json:"vendor"`
		ID                    string `json:"order_id"`
		BuyerEmail            string `json:"buyer_email"`
		BuyerName             string `json:"buyer_name"`
		BuyerUserID           string `json:"buyer_userid,omitempty"`
		BuyerPhone            string `json:"buyer_phone"`
		GatewayBuyerUUID      string `json:"gateway_buyer_uuid,omitempty"`
		Amount                int    `json:"amount"`
		Currency              string `json:"currency"`
		PaymentMethods        string `json:"payment_methods"`
		RedirectURL           string `json:"redirect_url,omitempty"`
		CancelURL             string `json:"cancel_url,omitempty"`
		Webhook               string `json:"webhook,omitempty"`
		BillingFirstName      string `json:"billing.firstname"`
		BillingLastName       string `json:"billing.lastname"`
		BillingAddress1       string `json:"billing.address_1"`
		BillingAddress2       string `json:"billing.address_2,omitempty"`
		BillingCity           string `json:"billing.city"`
		BillingStateRegion    string `json:"billing.state_or_region"`
		BillingPostCodePOBox  string `json:"billing.postcode_or_pobox"`
		BillingCountry        string `json:"billing.country"`
		BillingPhone          string `json:"billing.phone"`
		ShippingFirstName     string `json:"shipping.firstname,omitempty"`
		ShippingLastName      string `json:"shipping.lastname,omitempty"`
		ShippingAddress1      string `json:"shipping.address_1,omitempty"`
		ShippingAddress2      string `json:"shipping.address_2,omitempty"`
		ShippingCity          string `json:"shipping.city,omitempty"`
		ShippingStateRegion   string `json:"shipping.state_or_region,omitempty"`
		ShippingPostCodePOBox string `json:"shipping.postcode_or_pobox,omitempty"`
		ShippingCountry       string `json:"shipping.country,omitempty"`
		ShippingPhone         string `json:"shipping.phone,omitempty"`
		BuyerRemarks          string `json:"buyer_remarks,omitempty"`
		MerchantRemarks       string `json:"merchant_remarks,omitempty"`
		NumberItems           int    `json:"no_of_items,omitempty"`
		HeaderColour          string `json:"header_colour,omitempty"`
		LinkColour            string `json:"link_colour,omitempty"`
		ButtonColour          string `json:"button_colour,omitempty"`
		Expiry                int    `json:"expiry,omitempty"`
	}{
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
	}

	resp, err := cln.CreateOrder(ctx, client.OrderInput(input))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(resp)
}
