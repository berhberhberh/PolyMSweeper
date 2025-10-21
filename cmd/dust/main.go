package main

import (
	"fmt"
	"log"

	"fiscal/pkg/marketmaker"
)

func main() {
	fmt.Println("=======================================================")
	fmt.Println("Dust Market Analyzer - Intelligent Pricing for Illiquid Markets")
	fmt.Println("=======================================================")
	fmt.Println()

	// Initialize market maker
	mm := marketmaker.New(&marketmaker.Config{
		MinSpreadPct:    0.002,
		TargetSpreadPct: 0.001,
		MaxMarkets:      100,
	})

	// Find illiquid markets
	fmt.Println("Scanning for dust markets with placeholder orderbooks...")
	opportunities, err := mm.FindIlliquidMarkets()
	if err != nil {
		log.Fatalf("Error finding illiquid markets: %v", err)
	}

	if len(opportunities) == 0 {
		fmt.Println("\nNo illiquid markets found.")
		return
	}

	fmt.Printf("\n[SUCCESS] Found %d dust markets!\n", len(opportunities))
	fmt.Println("\nApplying intelligent pricing strategies...\n")
	fmt.Println("=======================================================")

	// Categorize and price each market
	ps := &marketmaker.PricingStrategy{}

	type ScoredOpportunity struct {
		Opp       marketmaker.Opportunity
		Category  marketmaker.MarketCategory
		BidPrice  float64
		AskPrice  float64
		Reasoning string
		PosSize   float64
	}

	var scoredOpps []ScoredOpportunity

	bankroll := 500.0 // Assume $500 bankroll for dust market testing

	for _, opp := range opportunities {
		category := ps.CategorizeMarket(opp.Question)
		bidPrice, askPrice, reasoning := ps.SuggestPricingForDustMarket(opp.Question, category)

		// Estimate probability from our pricing (mid-point)
		estimatedProb := (bidPrice + askPrice) / 2

		// Suggest position size
		buySize, sellSize, sizeReasoning := ps.SuggestPositionSize(bankroll, estimatedProb, bidPrice, askPrice)

		scoredOpps = append(scoredOpps, ScoredOpportunity{
			Opp:       opp,
			Category:  category,
			BidPrice:  bidPrice,
			AskPrice:  askPrice,
			Reasoning: reasoning + " | " + sizeReasoning,
			PosSize:   (buySize + sellSize) / 2,
		})
	}

	// Show categorized opportunities
	categories := map[marketmaker.MarketCategory]string{
		marketmaker.CategorySports:      "SPORTS LONGSHOTS",
		marketmaker.CategoryPolitics:    "POLITICAL EVENTS",
		marketmaker.CategoryEconomic:    "ECONOMIC EVENTS",
		marketmaker.CategoryUnknown:     "UNCATEGORIZED",
	}

	for catID, catName := range categories {
		var categoryOpps []ScoredOpportunity
		for _, so := range scoredOpps {
			if so.Category == catID {
				categoryOpps = append(categoryOpps, so)
			}
		}

		if len(categoryOpps) == 0 {
			continue
		}

		fmt.Printf("\n%s (%d markets)\n", catName, len(categoryOpps))
		fmt.Println("=======================================================")

		// Show top 5 from this category
		for i, so := range categoryOpps {
			if i >= 5 {
				break
			}

			spread := so.AskPrice - so.BidPrice
			spreadPct := (spread / so.BidPrice) * 100

			fmt.Printf("\n%d. %s\n", i+1, so.Opp.Question)
			fmt.Printf("   Category: %s\n", catName)
			fmt.Printf("   Current Market: Bid %.4f | Ask %.4f (placeholder)\n",
				so.Opp.BestBid, so.Opp.BestAsk)
			fmt.Printf("   Suggested Prices: Bid %.4f | Ask %.4f\n", so.BidPrice, so.AskPrice)
			fmt.Printf("   Your Spread: %.3f%%\n", spreadPct)
			fmt.Printf("   Position Size: $%.0f per side\n", so.PosSize)
			fmt.Printf("   Reasoning: %s\n", so.Reasoning)
			fmt.Printf("   Token ID: %s\n", so.Opp.TokenID)
		}

		fmt.Println()
	}

	// Summary and strategy advice
	fmt.Println("\n=======================================================")
	fmt.Println("DUST MARKET STRATEGY")
	fmt.Println("=======================================================")
	fmt.Println()
	fmt.Println("These markets are illiquid for a reason - they're hard to price!")
	fmt.Println()
	fmt.Println("Risk Mitigation:")
	fmt.Println("1. Start with TINY positions ($5-20 per market)")
	fmt.Println("2. Use WIDE spreads (50-200% spread) for safety")
	fmt.Println("3. Monitor fills closely - immediate fill = bad pricing")
	fmt.Println("4. Diversify across 10+ uncorrelated markets")
	fmt.Println("5. Use external data (betting odds, 538, etc.) to validate prices")
	fmt.Println()
	fmt.Println("Expected Outcomes:")
	fmt.Println("- Most orders: No fills for days/weeks (capital tied up)")
	fmt.Println("- Some orders: Slow fills at fair prices (small profit)")
	fmt.Println("- Few orders: Immediate fills (probably mispriced - LOSS)")
	fmt.Println()
	fmt.Println("Better Alternative:")
	fmt.Println("Run './active' scanner to find markets with REAL liquidity")
	fmt.Println("Those markets are easier to price and have actual trading volume!")
	fmt.Println("=======================================================")
}
