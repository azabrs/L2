```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
test() выведет 2
anotherTest() выведет 1
Причина в том, что в первом случае мы работаем с именованной возвращаемой переменной. А во втором - с копией, изменение которой не влияет на возвращаемое значение.
```