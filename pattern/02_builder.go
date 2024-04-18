package main

import "fmt"

type Builder interface{
	MakeHeader(header string)
	MakeBody(body string)
	MakeFooter(footer string)
}

type ConcreteBuilder struct{
	product *Product
}

func(cb *ConcreteBuilder)MakeHeader(header string){
	cb.product.Content += "<header>" + header + "</header>"

}

func(cb *ConcreteBuilder)MakeBody(body string){
	cb.product.Content += "<body>" + body + "</body>"

}

func(cb *ConcreteBuilder)MakeFooter(footer string){
	cb.product.Content += "<footer>" + footer + "</footer>"

}

type Director struct{
	builder Builder
}

func(d Director)Construct(){
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

type Product struct{
	Content string
}

func(p *Product)ShowResult(){
	fmt.Println(p.Content)

}

func main(){
	Cb := &ConcreteBuilder{product : &Product{} }
	Dir := Director{builder : Cb}
	Cb.product.ShowResult()
	Dir.Construct()
	Cb.product.ShowResult()

}



