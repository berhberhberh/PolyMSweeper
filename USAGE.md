# Usage Guide

## Available Scanners

### 1. Main Scanner (Illiquid Markets)
Scans for markets with placeholder orderbooks (0.001/0.999 spreads).

```bash
go build -o fiscal.exe
./fiscal.exe
```

**Output:** List of 80+ illiquid markets with suggested conservative pricing (40% bid, 60% ask).

**Use Case:** Identify markets with no active market makers.

---

### 2. Active Market Scanner
Scans for markets with real liquidity and tradeable spreads (>0.2%).

```bash
go build -o active.exe ./cmd/active
./active.exe
```

**Output:** Markets with actual bids/asks where you can place orders inside the spread.

**Use Case:** Find markets with existing liquidity and trading volume.

---

### 3. Dust Market Analyzer (Intelligent Pricing)
Analyzes illiquid markets with intelligent categorization and pricing strategies.

```bash
go build -o dust.exe ./cmd/dust
./dust.exe
```

**Output:**
- Categorized markets (Sports, Politics, Economic, Unknown)
- Intelligent pricing based on category
- Position sizing recommendations
- Reasoning for each price suggestion

**Categories:**
- **Sports:** NFL Super Bowl winners, championships (priced 0.5-5%)
- **Politics:** Elections, candidates (priced 5-35%)
- **Economic:** Fed rates, recession (priced 20-50%)
- **Unknown:** Conservative wide spreads (10-30%)

**Use Case:** Get smart pricing suggestions for dust markets based on market type.

---

## Example Workflow

### Scenario: You want to market make on Polymarket

#### Step 1: Run active scanner first
```bash
./active.exe
```

Look for markets with:
- Spread > 0.2% (profitable)
- High volume ($1000+ daily)
- Prices 5-95% (easier to price)

**If found:** These are your BEST opportunities. Place orders inside their spread.

#### Step 2: If no active markets, run dust analyzer
```bash
./dust.exe
```

Review intelligent pricing for each category:
- **Sports longshots:** Check Vegas odds, ESPN power rankings
- **Political events:** Check polls, PredictIt
- **Economic events:** Check CME FedWatch, Fed minutes

#### Step 3: Validate external data
Before placing orders:
- Browns win Super Bowl → Check DraftKings (+10000 = ~1%)
- NYC Mayor race → Check RealClearPolitics polls
- Fed rate hike → Check CME FedWatch Tool

#### Step 4: Start small
- Position size: $5-20 per market
- Wide spreads: 100-200% for safety
- Diversify: 10+ uncorrelated markets
- Monitor: Check fills every 6 hours

---

## Scanner Comparison

| Feature | Main Scanner | Active Scanner | Dust Analyzer |
|---------|-------------|----------------|---------------|
| Markets Found | 80+ | 0-5 | 80+ |
| Pricing | Conservative (40/60) | Inside spread | Intelligent by category |
| Categorization | No | No | Yes (4 categories) |
| Position Sizing | No | No | Yes (Kelly Criterion) |
| External Data | No | No | Suggested sources |
| Best For | Quick scan | Real trading | Learning/analysis |

---

## Expected Results

### Active Markets (if found)
- **Frequency:** 0-5 markets with >0.2% spread
- **Volume:** $1000+ daily
- **Fills:** 3-10 per day
- **Profit:** 0.1-0.5% per fill
- **Monthly Return:** 10-30%

### Dust Markets
- **Frequency:** 80+ markets available
- **Volume:** $0-100 daily
- **Fills:** 0-2 per week
- **Profit:** Variable (10-200% spreads)
- **Monthly Return:** -5% to +5% (high variance)

---

## Risk Management

### For Active Markets:
- Position size: $100-500 per market
- Spread: 0.1-0.5%
- Max exposure: 20% of bankroll
- Stop loss: Exit if spread inverts

### For Dust Markets:
- Position size: $5-20 per market
- Spread: 100-200%
- Max exposure: 10% of bankroll
- Stop loss: Cancel if filled immediately

---

## Next Steps

1. **Run all three scanners** to understand market landscape
2. **Read DUST_MARKET_STRATEGIES.md** for deep dive on dust markets
3. **Pick 3-5 markets** you understand well
4. **Validate pricing** with external data sources
5. **Place small test orders** ($5-10)
6. **Monitor for fills** and adjust pricing
7. **Scale up** only after proving profitable

---

## External Data Sources

### Sports
- **NFL:** DraftKings, FanDuel futures
- **General:** ESPN, FiveThirtyEight

### Politics
- **Elections:** RealClearPolitics, PredictIt
- **Approval:** Gallup, 538

### Economic
- **Fed:** CME FedWatch Tool
- **Macro:** Bloomberg, FRED

### Crypto
- **Price:** CoinGecko, Messari
- **Events:** Twitter, Discord

---

## Troubleshooting

### "No active markets found"
**Cause:** Polymarket spreads are very tight (<0.2%)
**Solution:** Lower MinSpreadPct in config or focus on dust markets

### "Immediate fill on dust market"
**Cause:** Mispriced - you offered too good a deal
**Solution:** Cancel other side, reprice wider

### "No fills after 1 week"
**Cause:** Market is truly illiquid or prices too wide
**Solution:** Narrow spread slightly or exit market

---

## Configuration

Edit `main.go` or command files to adjust:

```go
mm := marketmaker.New(&marketmaker.Config{
    MinSpreadPct:    0.002,  // 0.2% minimum spread
    TargetSpreadPct: 0.001,  // 0.1% target spread
    MaxMarkets:      100,    // Scan top 100 markets
})
```

- **MinSpreadPct:** Only show markets with spread > this value
- **TargetSpreadPct:** Your desired spread when placing orders
- **MaxMarkets:** How many markets to scan (more = slower)

---

## Build All Scanners

```bash
# Main scanner
go build -o fiscal.exe

# Active market scanner
go build -o active.exe ./cmd/active

# Dust market analyzer
go build -o dust.exe ./cmd/dust
```

---

## Current Market State (Oct 2025)

Based on test runs:
- **Active markets with >0.2% spread:** 0
- **Dust markets (placeholders):** 82
- **Conclusion:** Most Polymarket markets are either very tight (<0.1% spread) or completely illiquid

**Implication:** Focus on dust markets with intelligent pricing, or wait for market volatility to widen spreads.
