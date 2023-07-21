package main

import (
	"net/http"
	"testing"
)

// go test -bench=. -run=^# -count=10 -bechtime=3s
func BenchmarkGetAllPeople(t *testing.B) {
	for i := 0; i < t.N; i++ {
		http.Get("http://localhost:8000")
	}
}
