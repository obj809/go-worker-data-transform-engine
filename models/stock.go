package models

// StockRecord represents a single stock data point from the input.
type StockRecord struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"change_percent"`
	DayHigh       float64 `json:"day_high"`
	DayLow        float64 `json:"day_low"`
	PreviousClose float64 `json:"previous_close"`
	Timestamp     string  `json:"timestamp,omitempty"`
}

// AggregatedResult represents the processed output with averaged values.
type AggregatedResult struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"change_percent"`
	DayHigh       float64 `json:"day_high"`
	DayLow        float64 `json:"day_low"`
	PreviousClose float64 `json:"previous_close"`
	Timestamp     string  `json:"timestamp"`
}

// ProcessRequest is the expected input format for the /process endpoint.
type ProcessRequest struct {
	Records []StockRecord `json:"records"`
}

// ProcessResponse is the output format for the /process endpoint.
type ProcessResponse struct {
	Success bool              `json:"success"`
	Data    *AggregatedResult `json:"data,omitempty"`
	Error   string            `json:"error,omitempty"`
}
