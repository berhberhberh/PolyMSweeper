package marketmaker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	GammaAPIURL = "https://gamma-api.polymarket.com"
	CLOBURL     = "https://clob.polymarket.com"
)

// MarketMaker handles market making operations
type MarketMaker struct {
	config     *Config
	httpClient *http.Client
}

// New creates a new MarketMaker instance
func New(config *Config) *MarketMaker {
	return &MarketMaker{
		config: config,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// FetchMarkets retrieves markets from Polymarket Gamma API
func (mm *MarketMaker) FetchMarkets() ([]Market, error) {
	url := fmt.Sprintf("%s/markets?limit=%d&closed=false&order=volume24hr&ascending=false",
		GammaAPIURL, mm.config.MaxMarkets)

	resp, err := mm.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch markets: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gamma API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var markets []Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, fmt.Errorf("failed to unmarshal markets: %w", err)
	}

	return markets, nil
}

// GetOrderBook fetches the orderbook for a specific token
func (mm *MarketMaker) GetOrderBook(tokenID string) (*OrderBookResponse, error) {
	url := fmt.Sprintf("%s/book?token_id=%s", CLOBURL, tokenID)

	resp, err := mm.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch orderbook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CLOB API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read orderbook response: %w", err)
	}

	var orderbook OrderBookResponse
	if err := json.Unmarshal(body, &orderbook); err != nil {
		return nil, fmt.Errorf("failed to unmarshal orderbook: %w", err)
	}

	return &orderbook, nil
}

// parseFloat safely parses a string to float64
func parseFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, 64)
}

// parseVolume handles volume which can be string or number
func parseVolume(v interface{}) float64 {
	if v == nil {
		return 0
	}

	switch val := v.(type) {
	case float64:
		return val
	case string:
		f, _ := strconv.ParseFloat(val, 64)
		return f
	default:
		return 0
	}
}
