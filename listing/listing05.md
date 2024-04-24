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
error
Причина в том, что переменная err внутри себя содержит структуру iface. В ней одно из полей хранит указатель на пустое значение типа customError, а другое указатель на структуру с метаданными типа customError
```
