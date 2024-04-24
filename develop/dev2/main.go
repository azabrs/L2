package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)
const (
	FIRST_ELEM = iota
	ELEM_AFTER_ESCAPE
	ELEM_AFTER_NUM
	ELEM_AFTER_CHAR
	)

func Solve(s string)(string, error){
	var uni = make(map[string]int)
	keys := []string{}
	res := ""
	var prev string
	state := FIRST_ELEM
	for _, val := range(s){
		switch state{
		case FIRST_ELEM:
			if string(val) == "\\"{
				state = ELEM_AFTER_ESCAPE

			} else if unicode.IsDigit(val){
				return "", fmt.Errorf("incorrect string")
			} else{
				state = ELEM_AFTER_CHAR
				prev = string(val)
				uni[string(val)] = 1
				keys = append(keys, string(val))
			}
		case ELEM_AFTER_CHAR:
			if string(val) == "\\"{
				state = ELEM_AFTER_ESCAPE
			} else if unicode.IsDigit(val){
				state = ELEM_AFTER_NUM
				uni[prev] , _= strconv.Atoi(string(val))
			} else if unicode.IsLetter(val){
				state = ELEM_AFTER_CHAR
				prev = string(val)
				uni[string(val)] = 1
				keys = append(keys, string(val))
			}
		case ELEM_AFTER_NUM:
			if string(val) == "\\"{
				state = ELEM_AFTER_ESCAPE
			} else if unicode.IsDigit(val){
				return "", fmt.Errorf("incorrect string")
			} else if unicode.IsLetter(val){
				state = ELEM_AFTER_CHAR
				prev = string(val)
				uni[string(val)] = 1
				keys = append(keys, string(val))
			}
		case ELEM_AFTER_ESCAPE:
				prev = string(val)
				state = ELEM_AFTER_CHAR
				uni[string(val)] = 1
				keys = append(keys, string(val))
		}
	}
	for _, key := range(keys){
		for i := 0; i < uni[key]; i++{
			res += key
		}
	}
	return res, nil
}

func main(){
	buf, err := Solve(`45`)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(buf)

}