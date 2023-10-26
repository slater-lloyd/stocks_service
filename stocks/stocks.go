package stocks

import (
	"context"
	"log"
	"os"
	ad "stocks_service/alphaAdapter"

	"github.com/joho/godotenv"
)

func getAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv("API_KEY")
}

type Server struct {
	UnimplementedStocksServiceServer
}

func (s *Server) GetLastPrice(ctx context.Context, message *StockSymbolMessage) (*StockPriceMessage, error) {
	log.Printf("Recieved request for code: %s\n", message.Symbol)

	API_KEY := getAPIKey()
	adapter := ad.AlphaAdapter{API_KEY: API_KEY}
	resp, err := adapter.GetLastPrice(message.Symbol)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	return &StockPriceMessage{Price: resp}, nil
}

func (s *Server) GetPriceAtDate(ctx context.Context, message *SymbolAtDateMessage) (*StockPriceMessage, error) {
	log.Printf("Recieved request for code: %s at date: %s\n", message.Symbol, message.Date)

	API_KEY := getAPIKey()
	adapter := ad.AlphaAdapter{API_KEY: API_KEY}
	resp, err := adapter.GetPriceAtDate(message.Symbol, message.Date)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	return &StockPriceMessage{Price: resp}, nil
}

func (s *Server) GetPercentChangeAtDate(ctx context.Context, message *SymbolAtDateMessage) (*StockPriceMessage, error) {
	log.Printf("Recieved percent change request for code: %s at date: %s\n", message.Symbol, message.Date)

	API_KEY := getAPIKey()
	adapter := ad.AlphaAdapter{API_KEY: API_KEY}
	resp, err := adapter.GetPercentChangeAtDate(message.Symbol, message.Date)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}
	return &StockPriceMessage{Price: resp}, nil
}
