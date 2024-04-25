package main

import (
	"fmt"
	"sort"
	"strings"
)

func compareMap(m1 map[rune]int, m2 map[rune]int) bool{
	if len(m1) != len(m2){
		return false
	}
	for key, val1 := range(m1){
		val2, ok := m2[key]
		if !ok || val1 != val2{
			return false
		} 
	}
	return true
}

func PrintMap(notSort map[string][]string)map[string][]string{
	res := make(map[string][]string)
	for key, words := range(notSort){
		sort.Strings(words)
		res[key] = words
		fmt.Printf("%v : ", key)
		fmt.Println(words)
	}
	return res
}

func Solve(words []string) map[string][]string{
	uniq := make(map[string]map[rune]int)
	res := make(map[string][]string)
	unique_name :=make(map[string]bool)
	for _, word := range(words){
		word = strings.ToLower(word)
		if _, ok := unique_name[word]; ok{
			continue
		}
		unique_name[word] = true
		wordMap := make(map[rune]int)
		for _, letter := range(word){
			if val, exist := wordMap[letter]; exist{
				wordMap[letter] = val + 1
			}else{
				wordMap[letter] = 1
			}
		}
		isExist := false
		for key, val := range(uniq){
			if compareMap(val, wordMap){
				isExist = true
				res[key] = append(res[key], word)
			}
		}
		if !isExist{
			res[word] = []string{word}

			uniq[word] = wordMap
		}
	}

	res = PrintMap(res)
	return res
}

func main(){
	Solve([]string{"пятак", "пятка", "тяпка", "листок",  "пяТка", "слиток", "столик"})
	
}