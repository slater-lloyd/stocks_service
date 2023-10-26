package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "stocks_service/stocks"
)

type StockClient struct {
}

var (
	stockClient pb.StocksServiceClient
)

func prepareStockClient(c *context.Context) error {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to establish connection on 9000: %v", err)
		return err
	}

	if stockClient != nil {
		conn.Close()
		return nil
	}

	stockClient = pb.NewStocksServiceClient(conn)
	return nil
}

func (sc *StockClient) GetLastPrice(code string, ctx *context.Context) (*float32, error) {
	if err := prepareStockClient(ctx); err != nil {
		return nil, err
	}

	res, err := stockClient.GetLastPrice(*ctx, &pb.StockSymbolMessage{Symbol: code})
	if err != nil {
		log.Fatalf("Error in client response: %v", err)
	}

	return &res.Price, nil
}

func (sc *StockClient) GetPriceAtDate(code string, date string, ctx *context.Context) (*float32, error) {
	if err := prepareStockClient(ctx); err != nil {
		return nil, err
	}

	res, err := stockClient.GetPriceAtDate(*ctx, &pb.SymbolAtDateMessage{Symbol: code, Date: date})
	if err != nil {
		log.Fatalf("Error in client response: %v", err)
	}

	return &res.Price, nil
}

func (sc *StockClient) GetPercentChangeAtDate(code string, date string, ctx *context.Context) (*float32, error) {
	if err := prepareStockClient(ctx); err != nil {
		return nil, err
	}

	res, err := stockClient.GetPercentChangeAtDate(*ctx, &pb.SymbolAtDateMessage{Symbol: code, Date: date})
	if err != nil {
		log.Fatalf("Error in client response: %v", err)
	}

	return &res.Price, nil
}
