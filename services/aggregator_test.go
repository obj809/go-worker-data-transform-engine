package services

import (
	"testing"

	"github.com/softdev/go-worker-data-transform-engine/models"
)

func TestAggregateStockData_Success(t *testing.T) {
	records := []models.StockRecord{
		{
			Symbol:        "^AXJO",
			Name:          "S&P/ASX 200",
			Price:         8929.41,
			Change:        7.53,
			ChangePercent: 0.084399,
			DayHigh:       8948.55,
			DayLow:        8903.81,
			PreviousClose: 8921.88,
		},
		{
			Symbol:        "^AXJO",
			Name:          "S&P/ASX 200",
			Price:         8971.71,
			Change:        12.24,
			ChangePercent: 0.136615,
			DayHigh:       8996.19,
			DayLow:        8949.59,
			PreviousClose: 8959.47,
		},
	}

	result, err := AggregateStockData(records)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Symbol != "^AXJO" {
		t.Errorf("expected symbol ^AXJO, got %s", result.Symbol)
	}
	if result.Name != "S&P/ASX 200" {
		t.Errorf("expected name S&P/ASX 200, got %s", result.Name)
	}

	// Check averaged price: (8929.41 + 8971.71) / 2 = 8950.56
	expectedPrice := 8950.56
	if result.Price != expectedPrice {
		t.Errorf("expected price %v, got %v", expectedPrice, result.Price)
	}

	// Check averaged change: (7.53 + 12.24) / 2 = 9.885
	expectedChange := 9.885
	if result.Change != expectedChange {
		t.Errorf("expected change %v, got %v", expectedChange, result.Change)
	}

	// Timestamp should be set
	if result.Timestamp == "" {
		t.Error("expected timestamp to be set")
	}
}

func TestAggregateStockData_EmptyInput(t *testing.T) {
	_, err := AggregateStockData([]models.StockRecord{})
	if err != ErrEmptyInput {
		t.Errorf("expected ErrEmptyInput, got %v", err)
	}
}

func TestAggregateStockData_SingleRecord(t *testing.T) {
	records := []models.StockRecord{
		{
			Symbol:        "TEST",
			Name:          "Test Stock",
			Price:         100.00,
			Change:        1.50,
			ChangePercent: 0.015,
			DayHigh:       101.00,
			DayLow:        99.00,
			PreviousClose: 98.50,
		},
	}

	result, err := AggregateStockData(records)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Price != 100.00 {
		t.Errorf("expected price 100.00, got %v", result.Price)
	}
}

func TestRoundTo(t *testing.T) {
	tests := []struct {
		value    float64
		decimals int
		expected float64
	}{
		{1.234567, 2, 1.23},
		{1.235, 2, 1.24},
		{1.2345, 4, 1.2345},
		{100.0, 2, 100.0},
	}

	for _, tc := range tests {
		result := roundTo(tc.value, tc.decimals)
		if result != tc.expected {
			t.Errorf("roundTo(%v, %d) = %v, expected %v", tc.value, tc.decimals, result, tc.expected)
		}
	}
}
