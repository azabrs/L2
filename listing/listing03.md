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
Причина в том, что err содержит в себе структуру iface Способ исправить это, использовать при сравнении вместо nil (*os.PathError)(nil)

```
