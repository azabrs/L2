package facade

import "fmt"

type tree struct{
}
func (t tree)grow(){
	fmt.Println("grow tree")
}

type child struct{
}
func(ch child) born(){
	fmt.Println("born child")
}

type house struct{
}
func(h house)build(){
	fmt.Println("build house")
}

type Man struct{
	tree
	house
	child
}

func(m Man)Todo(){
	m.build()
	m.grow()
	m.born()
}

func main(){
	Alex := Man{}
	Alex.Todo()

}
