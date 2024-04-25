package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)
func TestGetAnagramSet(t *testing.T) {
	strs := []string{"пятка", "тяпкА", "слиток", "столик", "СТОЛИК", "Листок", "ТЯпка", "пятак"}
	mapCheck := map[string][]string{"пятка": {"пятак", "пятка", "тяпка"}, "слиток": {"листок", "слиток", "столик"}}
	require.Equal(t, Solve(strs), mapCheck)
}
