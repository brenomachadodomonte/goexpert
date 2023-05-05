package tax

import "testing"

func TestCalculateTax(t *testing.T) {

	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

// test coverage cmd: go test -coverprofile=coverage.out
// generate html coverage: go tool cover -html=coverage.out
func TestCalculateTaxBatch(t *testing.T) {
	type CalcTax struct {
		amount, expect float64
	}

	table := []CalcTax{
		{0.0, 0.0},
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)

		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

// run benchmark: go test -bench=. -run=^# -count=10 -bechtime=3s
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
