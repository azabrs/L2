package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func Test1(t *testing.T){
	inputStr := `a4bc2d5e`
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, nil)
	assert.Equal(t, strOut, `aaaabccddddde`)
}
func Test2(t *testing.T){
	inputStr := `abcd`
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, nil)
	assert.Equal(t, strOut, `abcd`)
}
func Test3(t *testing.T){
	inputStr := `45`
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, fmt.Errorf("incorrect string"))
	assert.Equal(t, strOut, ``)
}
func Test4(t *testing.T){
	inputStr :=  ``
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, nil)
	assert.Equal(t, strOut, ``)
}
func Test5(t *testing.T){
	inputStr :=  `qwe\4\5`
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, nil)
	assert.Equal(t, strOut, `qwe45`)
}
func Test6(t *testing.T){
	inputStr :=  `qwe\45`
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, nil)
	assert.Equal(t, strOut, `qwe44444`)
}
func Test7(t *testing.T){
	inputStr :=  `qwe\\5`
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, nil)
	assert.Equal(t, strOut, `qwe\\\\\`)
}
func Test8(t *testing.T){
	inputStr :=  `qwe\45`
	strOut, err := Solve(inputStr)
	assert.Equal(t, err, nil)
	assert.Equal(t, strOut, `qwe44444`)
}