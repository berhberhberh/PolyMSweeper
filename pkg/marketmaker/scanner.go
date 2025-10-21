package marketmaker

import (
	"encoding/json"
	"fmt"
	"time"
)

// FindIlliquidMarkets scans for markets with placeholder orderbooks (no real bids)
// These are the best opportunities for becoming the first market maker
func (mm *MarketMaker) FindIlliquidMarkets() ([]Opportunity, error) {
	// Fetch markets from Gamma API
	markets, err := mm.FetchMarkets()
	if err != nil {
		return nil, err
	}

	var opportunities []Opportunity

	fmt.Printf("Scanning %d markets for illiquid orderbooks...\n", len(markets))

	for _, market := range markets {
		// Skip closed markets
		if market.Closed {
			continue
		}

		// Parse token IDs
		var tokenIDs []string
		if market.ClobTokenIDs != "" {
			if err := json.Unmarshal([]byte(market.ClobTokenIDs), &tokenIDs); err != nil {
				continue
			}
		}

		if len(tokenIDs) == 0 {
			continue
		}

		// Check first token (YES token)
		tokenID := tokenIDs[0]

		// Get orderbook
		book, err := mm.GetOrderBook(tokenID)
		if err != nil {
			// Skip markets with errors
			time.Sleep(50 * time.Millisecond) // Rate limiting
			continue
		}

		// Check if orderbook has bids and asks
		if len(book.Bids) == 0 || len(book.Asks) == 0 {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// Parse best bid/ask
		bestBid, err := parseFloat(book.Bids[0].Price)
		if err != nil {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		bestAsk, err := parseFloat(book.Asks[0].Price)
		if err != nil {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// CORE LOGIC: Detect placeholder orderbooks
		// Placeholder: bid <= 0.01, ask >= 0.99 (99,800% spread)
		isPlaceholder := bestBid <= 0.01 && bestAsk >= 0.99

		if !isPlaceholder {
			// Not an illiquid market, skip
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// Parse volume
		volume, _ := parseFloat(market.Volume24hr)

		// Calculate spread
		spread := bestAsk - bestBid
		spreadPct := 0.0
		if bestBid > 0 {
			spreadPct = spread / bestBid
		}

		// For illiquid markets, suggest initial pricing
		// Use conservative wide spreads for safety (per RISKS_AND_MITIGATION.md)
		suggestedBuyPrice := 0.40  // Start at 40% for competitive events
		suggestedSellPrice := 0.60 // 20 cent spread for safety

		// For longshot markets (low probability), use different pricing
		// We can't tell probability from placeholder, so use conservative defaults
		// In practice, user should adjust based on external data sources

		opportunities = append(opportunities, Opportunity{
			Question:           market.Question,
			TokenID:            tokenID,
			Volume:             volume,
			BestBid:            bestBid,
			BestAsk:            bestAsk,
			SpreadPct:          spreadPct,
			SuggestedBuyPrice:  suggestedBuyPrice,
			SuggestedSellPrice: suggestedSellPrice,
			IsIlliquid:         true,
		})

		time.Sleep(50 * time.Millisecond) // Rate limiting
	}

	return opportunities, nil
}

// FindActiveMarkets finds markets with real liquidity (NOT placeholders)
// These are markets where other traders are already active
func (mm *MarketMaker) FindActiveMarkets() ([]Opportunity, error) {
	markets, err := mm.FetchMarkets()
	if err != nil {
		return nil, err
	}

	var opportunities []Opportunity

	fmt.Printf("Scanning %d markets for active liquidity...\n", len(markets))

	for _, market := range markets {
		if market.Closed {
			continue
		}

		var tokenIDs []string
		if market.ClobTokenIDs != "" {
			if err := json.Unmarshal([]byte(market.ClobTokenIDs), &tokenIDs); err != nil {
				continue
			}
		}

		if len(tokenIDs) == 0 {
			continue
		}

		tokenID := tokenIDs[0]

		book, err := mm.GetOrderBook(tokenID)
		if err != nil {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		if len(book.Bids) == 0 || len(book.Asks) == 0 {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		bestBid, err := parseFloat(book.Bids[0].Price)
		if err != nil {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		bestAsk, err := parseFloat(book.Asks[0].Price)
		if err != nil {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// Filter OUT placeholder orderbooks
		if bestBid <= 0.01 && bestAsk >= 0.99 {
			time.Sleep(50 * time.Millisecond)
			continue // Skip placeholders
		}

		// Filter extreme prices (< 5% or > 95%)
		if bestBid < 0.05 || bestAsk > 0.95 {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		volume, _ := parseFloat(market.Volume24hr)

		spread := bestAsk - bestBid
		spreadPct := 0.0
		if bestBid > 0 {
			spreadPct = spread / bestBid
		}

		// Check if spread is wide enough for market making
		if spreadPct < mm.config.MinSpreadPct {
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// Calculate suggested prices (place orders inside current spread)
		mid := (bestBid + bestAsk) / 2
		suggestedBuyPrice := mid - (mm.config.TargetSpreadPct / 2)
		suggestedSellPrice := mid + (mm.config.TargetSpreadPct / 2)

		opportunities = append(opportunities, Opportunity{
			Question:           market.Question,
			TokenID:            tokenID,
			Volume:             volume,
			BestBid:            bestBid,
			BestAsk:            bestAsk,
			SpreadPct:          spreadPct,
			SuggestedBuyPrice:  suggestedBuyPrice,
			SuggestedSellPrice: suggestedSellPrice,
			IsIlliquid:         false,
		})

		time.Sleep(50 * time.Millisecond)
	}

	return opportunities, nil
}
