```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Программа выведет error т. к. err содержит в себе структуру iface одно из полей которой хранит указатель на нулевое значение типа customError, а другой указатель на структуру хронящую метаданные этого типа

```
