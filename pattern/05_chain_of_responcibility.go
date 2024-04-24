package main

import "fmt"

type Handler interface{
	SendRequest(int) string
}

type Concreteandler1 struct{
	Next Handler
}

func(h *Concreteandler1)SendRequest(message int) string{
	if message == 1{
		return "Its 1 handler"
	} else if h.Next != nil{
		s := h.Next.SendRequest(message)
		return s
	}
	return ""
}

type Concreteandler2 struct{
	Next Handler
}

func(h *Concreteandler2)SendRequest(message int) string{
	if message == 2{
		return "Its 2 handler"
	} else if h.Next != nil{
		s := h.Next.SendRequest(message)
		return s
	}
	return ""
}

type Concreteandler3 struct{
	Next Handler
}

func(h *Concreteandler3)SendRequest(message int) string{
	if message == 3{
		return "Its 3 handler"
	} else if h.Next != nil{
		s := h.Next.SendRequest(message)
		return s
	}
	return ""
}

func main(){
	handler := &Concreteandler1{
		Next : &Concreteandler2{
			Next : &Concreteandler3{},
		},
	}
	req := handler.SendRequest(2)
	fmt.Println(req)
}