package alphaAdapter

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
)

const (
	DAILY    = "TIME_SERIES_DAILY"
	INTRADAY = "TIME_SERIES_INTRADAY"
)

type AlphaAdapter struct {
	API_KEY string
}

type TimeSeriesData struct {
	MetaData   *MetaData
	TimeSeries *TimeSeries
}

type MetaData map[string]string

func (a *AlphaAdapter) GetTimeSeriesDaily(sym string) (*TimeSeriesData, error) {
	vars := "function=" + DAILY + "&symbol=" + sym + "&apikey=" + a.API_KEY
	URL := "https://www.alphavantage.co/query?" + vars
	stockRes, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Error getting stock data: %v", err)
	}

	body, err := io.ReadAll(stockRes.Body)
	metaData := gjson.Get(string(body), "Meta Data")
	timeSeries := gjson.Get(string(body), "Time Series (Daily)")
	timeSeriesMetaData, err := buildTimeSeriesData(metaData.String(), timeSeries.String())
	if err != nil {
		return nil, err
	}
	stockRes.Body.Close()

	return timeSeriesMetaData, nil
}

func (a *AlphaAdapter) GetLastPrice(sym string) (float32, error) {
	tsd, err := a.GetTimeSeriesDaily(sym)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	ts := *tsd.TimeSeries
	lastDate := ts.TimeStamps()[0]

	strVal := ts[lastDate].Close()

	flVal, err := strconv.ParseFloat(strVal, 32)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}

	return float32(flVal), nil
}

func (a *AlphaAdapter) GetPriceAtDate(sym string, date string) (float32, error) {
	tsd, err := a.GetTimeSeriesDaily(sym)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	ts := *tsd.TimeSeries
	strVal := ts[date].Close()

	flVal, err := strconv.ParseFloat(strVal, 32)
	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}

	return float32(flVal), nil
}

func (a *AlphaAdapter) GetPercentChangeAtDate(sym string, date string) (float32, error) {
	tsd, err := a.GetTimeSeriesDaily(sym)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	ts := *tsd.TimeSeries
	closeStrVal := ts[date].Close()
	openStrVal := ts[date].Open()

	closeFlVal, err := strconv.ParseFloat(closeStrVal, 32)
	openFlVal, err := strconv.ParseFloat(openStrVal, 32)

	if err != nil {
		log.Fatalf("Error in response: %v", err)
	}

	flVal := closeFlVal / openFlVal
	return float32(flVal), nil
}

func buildTimeSeriesData(metaData, timeSeries string) (*TimeSeriesData, error) {

	var meta MetaData
	err := json.Unmarshal([]byte(metaData), &meta)
	if err != nil {
		return nil, err
	}

	var series TimeSeries
	err = json.Unmarshal([]byte(timeSeries), &series)
	if err != nil {
		return nil, err
	}

	return &TimeSeriesData{&meta, &series}, nil
}

type TimeSeries map[string]Series

func (t TimeSeries) TimeStamps() []string {
	var timeStamps []string
	for key := range t {
		timeStamps = append(timeStamps, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(timeStamps)))
	return timeStamps
}

type Series map[string]string

func (s Series) Open() string {
	return s["1. open"]
}

func (s Series) High() string {
	return s["2. high"]
}

func (s Series) Low() string {
	return s["3. low"]
}

func (s Series) Close() string {
	return s["4. close"]
}

func (s Series) Volume() string {
	return s["5. volume"]
}
