package marketmaker

import (
	"math"
	"strings"
)

// PricingStrategy contains logic for pricing different types of markets
type PricingStrategy struct{}

// MarketCategory represents different types of markets
type MarketCategory int

const (
	CategoryUnknown MarketCategory = iota
	CategorySports                 // Sports outcomes (Super Bowl, etc.)
	CategoryPolitics               // Elections, political events
	CategoryEconomic               // Fed rates, inflation, etc.
	CategoryLongshot              // Very unlikely events (< 5%)
	CategoryCompetitive           // 40-60% probability range
)

// CategorizeMarket attempts to categorize a market based on its question
func (ps *PricingStrategy) CategorizeMarket(question string) MarketCategory {
	lowerQ := strings.ToLower(question)

	// Sports keywords
	sportsKeywords := []string{"super bowl", "win", "championship", "nfl", "nba", "mlb", "soccer", "world cup"}
	for _, keyword := range sportsKeywords {
		if strings.Contains(lowerQ, keyword) {
			return CategorySports
		}
	}

	// Politics keywords
	politicsKeywords := []string{"election", "president", "mayor", "senator", "win the", "elected"}
	for _, keyword := range politicsKeywords {
		if strings.Contains(lowerQ, keyword) {
			return CategoryPolitics
		}
	}

	// Economic keywords
	economicKeywords := []string{"fed", "interest rate", "inflation", "gdp", "recession", "unemployment"}
	for _, keyword := range economicKeywords {
		if strings.Contains(lowerQ, keyword) {
			return CategoryEconomic
		}
	}

	return CategoryUnknown
}

// SuggestPricingForDustMarket provides intelligent pricing for illiquid/dust markets
// Returns (bidPrice, askPrice, reasoning)
func (ps *PricingStrategy) SuggestPricingForDustMarket(question string, category MarketCategory) (float64, float64, string) {
	switch category {
	case CategorySports:
		return ps.priceSportsLongshot(question)

	case CategoryPolitics:
		return ps.pricePoliticalEvent(question)

	case CategoryEconomic:
		return ps.priceEconomicEvent(question)

	default:
		// Conservative default: wide spread for safety
		return 0.10, 0.30, "Unknown category - using conservative wide spread (10-30%)"
	}
}

// priceSportsLongshot prices sports outcomes (usually longshots)
func (ps *PricingStrategy) priceSportsLongshot(question string) (float64, float64, string) {
	lowerQ := strings.ToLower(question)

	// Super Bowl winner pricing
	if strings.Contains(lowerQ, "super bowl") {
		// Most NFL teams: ~3% (1/32 teams)
		// Bad teams: ~0.5-1%
		// Good teams: ~5-15%

		// Check for bad teams
		badTeams := []string{"browns", "titans", "jets", "raiders", "panthers", "giants"}
		for _, team := range badTeams {
			if strings.Contains(lowerQ, team) {
				return 0.005, 0.015, "Bad NFL team - priced at 0.5-1.5% (conservative longshot)"
			}
		}

		// Check for good teams
		goodTeams := []string{"chiefs", "49ers", "ravens", "bills", "eagles"}
		for _, team := range goodTeams {
			if strings.Contains(lowerQ, team) {
				return 0.08, 0.12, "Good NFL team - priced at 8-12% (competitive odds)"
			}
		}

		// Average team: ~3%
		return 0.02, 0.05, "Average NFL team - priced at 2-5% (base rate 1/32 teams)"
	}

	// Default sports longshot
	return 0.01, 0.05, "Generic sports longshot - priced at 1-5%"
}

// pricePoliticalEvent prices political outcomes
func (ps *PricingStrategy) pricePoliticalEvent(question string) (float64, float64, string) {
	lowerQ := strings.ToLower(question)

	// Presidential elections - multi-candidate races
	if strings.Contains(lowerQ, "president") {
		// If it's a specific candidate in a race with many candidates
		if strings.Contains(lowerQ, "will") && strings.Contains(lowerQ, "win") {
			// Unknown candidate in large field: ~5-15%
			return 0.05, 0.15, "Presidential candidate in large field - priced at 5-15%"
		}
	}

	// Mayoral races - local elections
	if strings.Contains(lowerQ, "mayor") {
		return 0.10, 0.25, "Mayoral candidate - priced at 10-25% (assume 4-5 competitive candidates)"
	}

	// Rare political events (resignations, etc.)
	rareEvents := []string{"out in", "resign", "impeach", "remove"}
	for _, event := range rareEvents {
		if strings.Contains(lowerQ, event) {
			return 0.01, 0.05, "Rare political event - priced at 1-5%"
		}
	}

	// Default political event
	return 0.15, 0.35, "Generic political event - priced at 15-35%"
}

// priceEconomicEvent prices economic/Fed events
func (ps *PricingStrategy) priceEconomicEvent(question string) (float64, float64, string) {
	lowerQ := strings.ToLower(question)

	// Fed rate increases
	if strings.Contains(lowerQ, "fed") && strings.Contains(lowerQ, "increase") {
		// Check timeframe
		if strings.Contains(lowerQ, "2025") || strings.Contains(lowerQ, "2026") {
			return 0.20, 0.40, "Fed rate increase (future) - priced at 20-40%"
		}
		return 0.30, 0.50, "Fed rate increase - priced at 30-50%"
	}

	// Recession predictions
	if strings.Contains(lowerQ, "recession") {
		return 0.15, 0.35, "Recession prediction - priced at 15-35%"
	}

	// Default economic event
	return 0.25, 0.45, "Generic economic event - priced at 25-45%"
}

// CalculateKellyBetSize calculates optimal position size using Kelly Criterion
// edgePercent: your edge over the market (e.g., 0.05 for 5% edge)
// probability: your estimated true probability (e.g., 0.5 for 50%)
// Returns fraction of bankroll to risk
func (ps *PricingStrategy) CalculateKellyBetSize(probability float64, marketPrice float64) float64 {
	if probability <= marketPrice || probability <= 0 || probability >= 1 {
		return 0 // No edge or invalid probability
	}

	// Kelly formula: f = (bp - q) / b
	// where:
	//   b = odds (payout ratio)
	//   p = probability of winning
	//   q = probability of losing (1 - p)

	b := (1.0 - marketPrice) / marketPrice // Odds
	q := 1.0 - probability

	kelly := (b*probability - q) / b

	// Use fractional Kelly (1/4 Kelly) for safety
	fractionalKelly := kelly * 0.25

	// Cap at 10% of bankroll for risk management
	return math.Min(fractionalKelly, 0.10)
}

// SuggestPositionSize suggests position size for a market
// bankroll: total capital available
// probability: your estimated probability
// marketBid/marketAsk: current market prices
// Returns (buySize in dollars, sellSize in dollars, reasoning)
func (ps *PricingStrategy) SuggestPositionSize(bankroll float64, probability float64, marketBid float64, marketAsk float64) (float64, float64, string) {
	// For dust markets, start VERY small
	minSize := 5.0   // $5 minimum
	maxSize := 50.0  // $50 maximum for dust markets

	// Calculate Kelly sizing
	kellyFraction := ps.CalculateKellyBetSize(probability, (marketBid+marketAsk)/2)
	suggestedSize := bankroll * kellyFraction

	// Clamp to min/max for dust markets
	if suggestedSize < minSize {
		return minSize, minSize, "Using minimum size ($5) - Kelly suggests lower but testing requires minimum capital"
	}

	if suggestedSize > maxSize {
		return maxSize, maxSize, "Capped at $50 - dust markets are too risky for larger positions"
	}

	return suggestedSize, suggestedSize, "Kelly criterion sizing with 1/4 Kelly for safety"
}
