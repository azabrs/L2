package main

import "fmt"

type Visitor interface{
    VisitSushiBar(p SushiBar) string
    VisitPizzeria(p Pizzeria) string
    VisitBurgerBar(p BurgerBar) string
}

type Place interface{
    Accept(v Visitor) string
}

type People struct{

}

func (v People) VisitSushiBar(p SushiBar) string{
    return p.BuySushi()
}

func (v People) VisitPizzeria(p Pizzeria) string{
    return p.BuyPizza()
}

func (v People) VisitBurgerBar(p BurgerBar) string{
    return p.BuyBurger()
}

type City struct{
    places []Place
}

func (c *City)Add(p Place){
    c.places = append(c.places, p)
}

func (c *City)Accept(v Visitor) string{
    res := ""
    for _, p := range(c.places){
        res += p.Accept(v)
    }
    return res
}

type SushiBar struct{
}

func (p SushiBar) Accept(v Visitor) string{
    return v.VisitSushiBar(p)
}

func (p SushiBar)BuySushi() string{
    return "Buy Sushi "
}

type BurgerBar struct{
}

func (p BurgerBar) Accept(v Visitor) string{
    return v.VisitBurgerBar(p)
}

func (p BurgerBar)BuyBurger() string{
    return "Buy Burger "
}

type Pizzeria struct{
}

func (p Pizzeria) Accept(v Visitor) string{
    return v.VisitPizzeria(p)
}

func (p Pizzeria)BuyPizza() string{
    return "Buy Pizza "
}

func main(){
    city := City{}
    city.Add(SushiBar{})
    city.Add(BurgerBar{})
    city.Add(Pizzeria{})
    people := People{}
    fmt.Println(city.Accept(people))
}
