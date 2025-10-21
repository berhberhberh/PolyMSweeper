package main

import (
	"fmt"
	"log"

	"fiscal/pkg/marketmaker"
)

func main() {
	fmt.Println("===========================================")
	fmt.Println("Polymarket Illiquid Market Maker")
	fmt.Println("===========================================")
	fmt.Println()

	// Initialize market maker
	mm := marketmaker.New(&marketmaker.Config{
		MinSpreadPct:    0.002, // Only trade if spread > 0.2%
		TargetSpreadPct: 0.001, // Capture 0.1% per round-trip
		MaxMarkets:      100,   // Scan top 100 markets
	})

	// Find illiquid markets (placeholder orderbooks)
	fmt.Println("Scanning for illiquid markets with placeholder orderbooks...")
	opportunities, err := mm.FindIlliquidMarkets()
	if err != nil {
		log.Fatalf("Error finding illiquid markets: %v", err)
	}

	if len(opportunities) == 0 {
		fmt.Println("\nNo illiquid markets found with placeholder orderbooks.")
		return
	}

	fmt.Printf("\n[SUCCESS] Found %d illiquid markets!\n\n", len(opportunities))
	fmt.Println("Top opportunities:")
	fmt.Println()

	// Show top 10
	for i, opp := range opportunities {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %s\n", i+1, opp.Question)
		fmt.Printf("   Current: Bid %.4f | Ask %.4f | Spread %.2f%%\n",
			opp.BestBid, opp.BestAsk, opp.SpreadPct*100)
		fmt.Printf("   Suggested: Buy %.4f | Sell %.4f\n",
			opp.SuggestedBuyPrice, opp.SuggestedSellPrice)
		fmt.Printf("   Token ID: %s\n", opp.TokenID)
		fmt.Println()
	}

	fmt.Println("===========================================")
	fmt.Println("Next Steps:")
	fmt.Println("1. Review the suggested prices")
	fmt.Println("2. Place initial bid/ask orders")
	fmt.Println("3. Monitor for fills")
	fmt.Println("===========================================")
}
