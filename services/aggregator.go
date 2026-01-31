package services

import (
	"errors"
	"math"
	"time"

	"github.com/softdev/go-worker-data-transform-engine/models"
)

var ErrEmptyInput = errors.New("input data cannot be empty")
var ErrMissingFields = errors.New("Missing required fields: {name, price, change, change_percent, day_high, day_low, previous_close}")

// AggregateStockData processes an array of stock records and returns averaged values.
// Uses symbol/name from the first record. Returns actual UTC timestamp.
func AggregateStockData(records []models.StockRecord) (*models.AggregatedResult, error) {
	if len(records) == 0 {
		return nil, ErrEmptyInput
	}

	first := records[0]
	if err := validateRecord(first); err != nil {
		return nil, err
	}

	count := float64(len(records))

	var sumPrice, sumChange, sumChangePercent float64
	var sumDayHigh, sumDayLow, sumPreviousClose float64

	for _, r := range records {
		sumPrice += r.Price
		sumChange += r.Change
		sumChangePercent += r.ChangePercent
		sumDayHigh += r.DayHigh
		sumDayLow += r.DayLow
		sumPreviousClose += r.PreviousClose
	}

	return &models.AggregatedResult{
		Symbol:        first.Symbol,
		Name:          first.Name,
		Price:         roundTo(sumPrice/count, 2),
		Change:        roundTo(sumChange/count, 4),
		ChangePercent: roundTo(sumChangePercent/count, 6),
		DayHigh:       roundTo(sumDayHigh/count, 2),
		DayLow:        roundTo(sumDayLow/count, 2),
		PreviousClose: roundTo(sumPreviousClose/count, 2),
		Timestamp:     time.Now().UTC().Format(time.RFC3339),
	}, nil
}

// roundTo rounds a float to the specified number of decimal places.
func roundTo(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}

// validateRecord checks that required fields are present.
func validateRecord(r models.StockRecord) error {
	if r.Symbol == "" || r.Name == "" {
		return ErrMissingFields
	}
	// Numeric fields with zero values are considered missing
	if r.Price == 0 && r.Change == 0 && r.ChangePercent == 0 &&
		r.DayHigh == 0 && r.DayLow == 0 && r.PreviousClose == 0 {
		return ErrMissingFields
	}
	return nil
}
