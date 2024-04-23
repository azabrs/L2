package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"log"
)

type Flags struct{
	fields    []int
	delimiter string
	sep       bool

}

func FlagParser() *Flags{
	f := &Flags{}
	bufString := flag.String("f", "", "fields")
	flag.BoolVar(&f.sep, "s", false, "only lines containing a separator")
	flag.StringVar(&f.delimiter, "d", "\t", "Use custom delim")
	flag.Parse()
	bufArr := strings.Split(*bufString, ",")
	f.fields = make([]int, 0, len(bufArr))
	for _, v := range bufArr {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err)
		}
		f.fields = append(f.fields, val)
	}
	return f
}

func ReadStrings()[]string{
	res := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		str = strings.Trim(str, "\n")
		res = append(res, str)
	}
	return res
}

func Cut(lines []string, fl Flags) []string{
	res := make([]string, 0, len(lines))
	
	for _, line := range(lines){
		splitted := strings.Split(line, fl.delimiter)
		if fl.sep && len(splitted) < 2{
			continue
		}
		buf := make([]string, 0)
		isEmpty := false
		for _, val := range(fl.fields){
			if val - 1 < len(splitted){
				buf = append(buf, splitted[val - 1]) 
				isEmpty = true
			}
		}
		if isEmpty{
			
			res = append(res, strings.Join(buf, fl.delimiter))
		}
	}
	return res
}
func PrintLines(lines []string){
	fmt.Println("---------------------")
	for _, line := range(lines){
		fmt.Println(line)
	}
}

func CutSimulator(){
	fl := FlagParser()
	lines := ReadStrings()
	res := Cut(lines, *fl)
	PrintLines(res)
}

func main(){
	CutSimulator()
}