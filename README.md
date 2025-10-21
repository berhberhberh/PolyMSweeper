# PolyMSweeper

Go-based market maker for illiquid Polymarket markets.

## Overview

PolyMSweeper identifies and analyzes illiquid prediction markets on Polymarket that have placeholder orderbooks (bid: 0.001, ask: 0.999). These markets have no active market makers and present opportunities to become the first liquidity provider.

## Features

- **Illiquid Market Detection**: Scans top markets and identifies those with placeholder orderbooks
- **Price Suggestions**: Recommends initial bid/ask prices for market making
- **High-Volume Scanning**: Analyzes top 100 markets by 24hr volume
- **Fast API Integration**: Direct integration with Polymarket's Gamma API and CLOB

## Usage

```bash
# Build
go build -o fiscal.exe

# Run scanner
./fiscal.exe
```

## How It Works

1. Fetches top markets by volume from Polymarket Gamma API
2. Checks orderbook for each market via CLOB API
3. Identifies markets with placeholder spreads (≥99% spread)
4. Suggests conservative initial pricing (40% bid, 60% ask)

## Market Making Strategy

The tool targets markets where:
- Best bid ≤ 0.01 AND best ask ≥ 0.99 (placeholder orderbook)
- No active market makers exist
- Opportunity to set initial prices and capture spreads

## Output

For each illiquid market found, displays:
- Market question
- Current bid/ask prices
- Spread percentage
- Suggested buy/sell prices
- Token ID for order placement

## Next Steps

After identifying opportunities:
1. Research the market to determine fair pricing
2. Adjust suggested prices based on external data
3. Place limit orders on both sides
4. Monitor fills and adjust as needed

## License

MIT
