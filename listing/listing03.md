```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false
Причина в том, что err содержит в себе структуру iface, которая хранит nil данные самого интерфейса и структуру itab, которая хранит метаданные интерфейса и nil не является. Способ исправить это - использовать при сравнении вместо nil (*os.PathError)(nil)

```
