package main

import (
	"context"
	"log"
	"stocks_service/client/stockClient"
	"time"
)

func main() {
	var stock_client client.StockClient
	timeout := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	code := "IBM"
	date := "2023-06-14"

	res, err := stock_client.GetLastPrice(code, &ctx)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	log.Printf("Recieved price from server for code %s: $%f", code, *res)

	res, err = stock_client.GetPriceAtDate(code, date, &ctx)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	log.Printf("Recieved price from server for code %s on date %s: $%f", code, date, *res)

	res, err = stock_client.GetPercentChangeAtDate(code, date, &ctx)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	log.Printf("Recieved percent change from server for code %s on date %s: %%%f", code, date, *res)

	defer cancel()
}
