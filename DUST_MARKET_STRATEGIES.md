# Dust Market Extraction Strategies

## What Are Dust Markets?

Dust markets are prediction markets with:
- Placeholder orderbooks (bid: 0.001, ask: 0.999)
- No active market makers
- Low/zero trading volume
- Often extreme outcomes (longshots)

## Why Are They Hard?

1. **Pricing difficulty**: No reference prices to anchor to
2. **Adverse selection**: If someone trades with you, they likely know something you don't
3. **Capital inefficiency**: Money sits idle waiting for fills
4. **Information asymmetry**: These markets attract informed traders who wait for liquidity

## Strategy 1: Sports Longshots (Base Rate Pricing)

### Theory
Use base rates and public data to price rare sports outcomes.

### Example: "Will Browns win Super Bowl 2026?"

**Base rate approach:**
- 32 NFL teams → base rate = 1/32 = 3.125%
- Browns recent performance: Bad (2-8 record)
- Adjust down: ~0.5-1% fair value

**Your pricing:**
- Bid: 0.005 (0.5%)
- Ask: 0.015 (1.5%)
- Spread: 200% (wide for safety)

**Position sizing:**
- Start: $5-10 per market
- Max exposure: $50 total across all NFL longshots

**Edge:**
- If fair value is 0.8%, you have 0.3% buffer on bid side
- If filled at 0.015, you're selling at premium to fair value

**Risk:**
- Injury to star QB could shift odds suddenly
- Playoff picture changes mid-season

### Validation Sources:
- Vegas futures odds (DraftKings, FanDuel)
- FiveThirtyEight projections
- ESPN power rankings

## Strategy 2: Political Events (Field Size Analysis)

### Theory
Price based on number of credible candidates and polling data.

### Example: "Will [Candidate X] win NYC Mayoral race?"

**Field size approach:**
- 5 credible candidates
- Base rate: 1/5 = 20%
- Check polls: If candidate has 10% support → adjust to 10-15%

**Your pricing:**
- Bid: 0.10 (10%)
- Ask: 0.25 (25%)
- Spread: 150%

**Position sizing:**
- $10-20 per candidate
- Total exposure: $100 across all candidates in race

**Edge:**
- Market may be completely unpriced
- You set the narrative

**Risk:**
- Scandal breaks → candidate drops to 0%
- Election canceled/postponed

### Validation Sources:
- RealClearPolitics polls
- PredictIt (comparison markets)
- Local news coverage

## Strategy 3: Economic Events (Fed Watching)

### Theory
Use Fed dot plots, economic forecasts, and futures markets.

### Example: "Fed increases rates by 25+ bps after Oct 2025 meeting?"

**Market data approach:**
- Check CME FedWatch Tool (Fed Funds futures)
- Read Fed minutes and dot plot
- Current market pricing: 35% chance

**Your pricing:**
- Bid: 0.30 (30%)
- Ask: 0.40 (40%)
- Spread: 33%

**Position sizing:**
- $20-50 per market (more confident due to data)
- Max: $200 across all Fed events

**Edge:**
- You have access to same data as everyone
- But you're first liquidity provider

**Risk:**
- Inflation spike → probability jumps to 80%
- You're on wrong side

### Validation Sources:
- CME FedWatch Tool
- Fed meeting minutes
- Bloomberg economic forecasts

## Strategy 4: Portfolio Approach (Diversification)

### Theory
Don't try to pick winners - be the house across ALL markets.

### Implementation:
1. **Pick 20-30 uncorrelated dust markets**
   - Mix: Sports, politics, economic
   - Different timeframes
   - Different probability ranges

2. **Price conservatively across all**
   - Wide spreads (100-200%)
   - Small positions ($5-20)
   - Total capital: $200-500

3. **Expected outcomes:**
   - 70% of markets: No fills (capital tied up)
   - 20% of markets: Slow fills at fair prices (small profit)
   - 10% of markets: Bad pricing (losses)

4. **Goal: Positive expected value across portfolio**
   - Win: 20% of markets × $2 profit = $0.40 per market
   - Lose: 10% of markets × $5 loss = -$0.50 per market
   - **Net: -$0.10 per market** (UNPROFITABLE!)

**Problem:** This doesn't work unless your pricing is VERY good.

## Strategy 5: External Data Anchoring

### Theory
Never price in a vacuum - always anchor to external markets.

### Process:
1. **Find comparable markets**
   - "Browns win Super Bowl" → Check DraftKings odds
   - "Fed raises rates" → Check CME futures
   - "Candidate X wins" → Check PredictIt

2. **Convert odds to probabilities**
   - Vegas: Browns +10000 = 0.99% implied probability
   - Add vig adjustment: True probability ~0.8%

3. **Price around external market**
   - Bid: 0.005 (below external)
   - Ask: 0.015 (above external)

4. **Arbitrage check**
   - If your bid > external ask: FREE MONEY (buy external, sell yours)
   - If your ask < external bid: FREE MONEY (buy yours, sell external)

### Example Workflow:
```
Market: "Browns win Super Bowl"
1. Check DraftKings: +10000 (0.99% implied)
2. Remove vig: ~0.8% true probability
3. Your bid: 0.005 (0.5%) ✓ Below market
4. Your ask: 0.015 (1.5%) ✓ Above market
5. Your spread: 1.0% (125% spread)
6. Position: $10
7. Edge: Captured 0.5% spread vs external market
```

## Strategy 6: The "Informed Trader Trap"

### Theory
You're providing free options to informed traders.

### Why This Matters:
When you place:
- Bid: 0.005
- Ask: 0.015

You've created a FREE OPTION:
- If Browns' odds improve (star player trade): Informed trader buys your 0.015 ask
- If Browns' odds worsen (injury): Informed trader sells your 0.005 bid

**You're always on the wrong side!**

### Mitigation:
1. **Wide spreads** (200%+) to compensate for adverse selection
2. **Small sizes** ($5-10) to limit damage
3. **Monitor fills**:
   - Immediate fill = BAD (mispriced)
   - Cancel other side immediately
4. **Dynamic repricing**:
   - Check external markets every 6 hours
   - Adjust your orders if odds shift

## Strategy 7: Kelly Criterion Position Sizing

### Theory
Use Kelly Criterion to avoid over-betting.

### Formula:
```
f = (bp - q) / b

Where:
  f = fraction of bankroll to bet
  b = odds (payout ratio)
  p = your estimated probability
  q = 1 - p
```

### Example:
```
Market: Browns win Super Bowl
Your estimate: 1% probability
Market ask: 0.015 (1.5%)

Edge: You think it's worth 1%, market offers 1.5%
→ Don't bet (no edge on this side)

Market bid: 0.005 (0.5%)
Your estimate: 1% (you'd sell at 2%+)

Edge: You'd sell at 2%, market only pays 0.5%
→ Don't sell (no edge)
```

**Result: Skip this market!** No edge on either side.

### Practical Kelly:
- Use 1/4 Kelly (fractional Kelly) for safety
- Cap at 2% of bankroll per market
- Max 10% of bankroll in dust markets total

## Strategy 8: The "Wait and See" Approach

### Theory
Don't be the FIRST market maker - be the SECOND.

### How:
1. **Wait for someone else to price the market**
   - Let another trader place first bid/ask
   - See what price they chose

2. **Analyze their pricing**
   - Did they use external data?
   - Is their spread wide or tight?
   - What size are they offering?

3. **Improve on their prices**
   - Bid 1 tick higher
   - Ask 1 tick lower
   - Capture the inside spread

4. **Let them take the adverse selection risk**
   - They're the "stale" liquidity
   - You're the "fresh" liquidity
   - Informed traders hit them first

**Problem:** You're not extracting value from dust markets - you're parasiting on other market makers.

## Realistic Assessment

### Expected Returns from Dust Markets:

**Optimistic scenario:**
- 30 markets × $10 per market = $300 capital
- 5 markets fill per month
- Average profit: $1 per fill (10% spread)
- Monthly profit: $5
- **Return: 1.7% per month**

**Realistic scenario:**
- Most markets: No fills
- Some markets: Break even after adverse selection
- Few markets: Good fills
- **Return: 0-2% per month** (NOT WORTH IT)

**Pessimistic scenario:**
- Adverse selection dominates
- Bad pricing leads to losses
- **Return: -5% to -10%**

### Why Active Markets Are Better:

**Active market example:**
- Market: "Will X happen?"
- Current: Bid 0.47, Ask 0.49 (2% spread)
- You: Bid 0.475, Ask 0.485 (1% spread)
- Position: $100
- Fills: 3-5 per day
- Daily profit: $2-3
- **Monthly profit: $60-90** (20-30% return)

**Advantage:**
- Price discovery exists (easier to price)
- Volume exists (faster fills)
- Less adverse selection (market is efficient)

## Conclusion: Should You Market Make Dust Markets?

### YES, IF:
- You have access to superior pricing data (private betting models, etc.)
- You can monitor 24/7 and reprice dynamically
- You have $1000+ bankroll and can absorb losses
- You're doing it to LEARN market making

### NO, IF:
- You're trying to make meaningful money
- You don't have external data sources
- You can't monitor fills in real-time
- You have < $500 bankroll

### BETTER ALTERNATIVE:
**Focus on active markets with real spreads**
- Run the `active` scanner
- Find markets with 0.2-2% spreads
- Place orders inside current spread
- Earn consistent profits from volume

Dust markets are a learning opportunity, not a profit opportunity.
