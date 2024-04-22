package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time_ntp, err := Time()
	time_time := time.Now()
	assert.Equal(t, time_time.Round(time.Minute), time_ntp.Round(time.Minute))
	assert.Equal(t, err, nil)
}