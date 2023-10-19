# Stocks Service

## Overview:
This service implements a proxy server of the Alpha Vantage API for historical stock data. Written in Go, I created an adapter to consume and adjust the API data for my own needs. There are several gRPC methods to get data related to any stock tracked on Alpha Vantage.


## Methods:
The GetLastPrice client method intakes the stock symbol as a parameter, then outputs the last price of that stock.

The GetPriceAtDate client method intakes a stock symbol and a date in "yyyy-mm-dd" format, then outputs the closing price of that stock on the given date.

