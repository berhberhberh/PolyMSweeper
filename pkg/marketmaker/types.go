package marketmaker

// Config holds market maker configuration
type Config struct {
	MinSpreadPct    float64 // Minimum spread to participate (default 0.2%)
	TargetSpreadPct float64 // Your target spread inside theirs (default 0.1%)
	MaxMarkets      int     // Maximum number of markets to scan
}

// Market represents a Polymarket market
type Market struct {
	Question      string      `json:"question"`
	ClobTokenIDs  string      `json:"clobTokenIds"`
	Volume24hr    interface{} `json:"volume24hr"` // Can be string or number
	Closed        bool        `json:"closed"`
	Active        bool        `json:"active"`
}

// OrderBookResponse represents the CLOB orderbook response
type OrderBookResponse struct {
	Market string  `json:"market"`
	Asset  string  `json:"asset_id"`
	Bids   []Order `json:"bids"`
	Asks   []Order `json:"asks"`
}

// Order represents a single order in the orderbook
type Order struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

// Opportunity represents a market making opportunity
type Opportunity struct {
	Question           string
	TokenID            string
	Volume             float64
	BestBid            float64
	BestAsk            float64
	SpreadPct          float64
	SuggestedBuyPrice  float64
	SuggestedSellPrice float64
	IsIlliquid         bool // True if placeholder orderbook (0.001/0.999)
}
