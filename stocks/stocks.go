package stocks

import (
	"context"
	"log"
	ad "stocks_service/alphaAdapter"
)

type Server struct {
	UnimplementedStocksServiceServer
}

func (s *Server) GetLastPrice(ctx context.Context, message *StockSymbolMessage) (*StockPriceMessage, error) {
	log.Printf("Recieved request for code: %s\n", message.Symbol)

	API_KEY := "WAR2QUKV9B5KY2S2"
	adapter := ad.AlphaAdapter{API_KEY: API_KEY}
	resp, err := adapter.GetLastPrice(message.Symbol)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	return &StockPriceMessage{Price: resp}, nil
}

func (s *Server) GetPriceAtDate(ctx context.Context, message *PriceAtDateMessage) (*StockPriceMessage, error) {
	log.Printf("Recieved request for code: %s at date: %s\n", message.Symbol, message.Date)

	API_KEY := "WAR2QUKV9B5KY2S2"
	adapter := ad.AlphaAdapter{API_KEY: API_KEY}
	resp, err := adapter.GetPriceAtDate(message.Symbol, message.Date)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	return &StockPriceMessage{Price: resp}, nil
}
