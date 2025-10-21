package main

import (
	"fmt"
	"log"

	"fiscal/pkg/marketmaker"
)

func main() {
	fmt.Println("===========================================")
	fmt.Println("Active Market Scanner - Tradeable Spreads")
	fmt.Println("===========================================")
	fmt.Println()

	// Initialize market maker
	mm := marketmaker.New(&marketmaker.Config{
		MinSpreadPct:    0.002, // Only trade if spread > 0.2%
		TargetSpreadPct: 0.001, // Capture 0.1% per round-trip
		MaxMarkets:      100,   // Scan top 100 markets
	})

	// Find active markets with tradeable spreads
	fmt.Println("Scanning for active markets with real liquidity...")
	opportunities, err := mm.FindActiveMarkets()
	if err != nil {
		log.Fatalf("Error finding active markets: %v", err)
	}

	if len(opportunities) == 0 {
		fmt.Println("\nNo active markets found with tradeable spreads.")
		fmt.Println("All spreads are too tight (< 0.2%) or markets are illiquid.")
		return
	}

	fmt.Printf("\n[SUCCESS] Found %d active markets with tradeable spreads!\n\n", len(opportunities))

	// Show top 10
	for i, opp := range opportunities {
		if i >= 10 {
			break
		}

		mid := (opp.BestBid + opp.BestAsk) / 2
		ourSpread := opp.SuggestedSellPrice - opp.SuggestedBuyPrice
		ourSpreadPct := (ourSpread / mid) * 100

		fmt.Printf("%d. %s\n", i+1, opp.Question)
		fmt.Printf("   Volume: $%.0f | Mid: %.4f\n", opp.Volume, mid)
		fmt.Printf("   Current Market: Bid %.4f | Ask %.4f | Spread %.3f%%\n",
			opp.BestBid, opp.BestAsk, opp.SpreadPct*100)
		fmt.Printf("   Your Orders:    Bid %.4f | Ask %.4f | Spread %.3f%%\n",
			opp.SuggestedBuyPrice, opp.SuggestedSellPrice, ourSpreadPct)
		fmt.Printf("   Profit per round-trip: ~%.3f%%\n",
			((opp.SuggestedSellPrice-opp.SuggestedBuyPrice)/opp.SuggestedBuyPrice)*100)
		fmt.Printf("   Token ID: %s\n", opp.TokenID)
		fmt.Println()
	}

	// Show statistics
	fmt.Println("===========================================")
	fmt.Println("Summary:")
	fmt.Println("===========================================")

	totalVolume := 0.0
	avgSpread := 0.0
	for _, opp := range opportunities {
		totalVolume += opp.Volume
		avgSpread += opp.SpreadPct
	}

	if len(opportunities) > 0 {
		avgSpread /= float64(len(opportunities))
	}

	fmt.Printf("Markets found: %d\n", len(opportunities))
	fmt.Printf("Total 24h volume: $%.0f\n", totalVolume)
	fmt.Printf("Average spread: %.3f%%\n", avgSpread*100)
	fmt.Println()
	fmt.Println("These markets have REAL liquidity and active traders.")
	fmt.Println("You can place orders INSIDE their spread and be the best price!")
}
