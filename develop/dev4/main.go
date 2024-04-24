package main

import "fmt"

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

func PrintMap(res map[string][]string){
	for key, words := range(res){
		fmt.Printf("%v : ", key)
		fmt.Println(words)
	}
}

func Solve(words []string){
	uniq := make(map[string]map[rune]int)
	res := make(map[string][]string)
	for _, word := range(words){
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
		if isExist == false{
			res[word] = []string{word}
			uniq[word] = wordMap
		}
	}

	PrintMap(res)
}

func main(){
	Solve([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"})
	
}